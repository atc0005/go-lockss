// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/apex/log"

	"github.com/atc0005/go-lockss/internal/config"
	"github.com/atc0005/go-lockss/internal/lockss"
	"github.com/atc0005/go-lockss/internal/portchecks"
)

func main() {

	// config.EnableLogging()
	lockss.DisableLogging()

	// setup application configuration
	appCfg, err := config.NewConfig()
	switch {
	// TODO: How else to guard against nil cfg object?
	case appCfg != nil && appCfg.ShowVersion():
		config.Branding()
		os.Exit(0)
	case err == nil:
		// do nothing for this one
	case errors.Is(err, flag.ErrHelp):
		os.Exit(0)
	default:
		log.Fatalf("failed to initialize application: %s", err)
	}

	// If user supplied values, we should use those to retrieve the LOCKSS
	// configuration from the central LOCKSS configuration server, otherwise
	// try to automatically determine values and go from there.
	var lockssCfg *lockss.Config
	var cfgSource string
	switch {
	case appCfg.ConfigServerURL() != "":

		log.Debugf(
			"ConfigServerURL() is non-empty, using value %q",
			appCfg.ConfigServerURL(),
		)

		cfgSource = appCfg.ConfigServerURL()

		var err error
		lockssCfg, err = lockss.NewFromURL(appCfg.ConfigServerURL())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case appCfg.ConfigFile() != "":

		log.Debugf(
			"ConfigFile() is non-empty, using value %q",
			appCfg.ConfigFile(),
		)

		cfgSource = appCfg.ConfigFile()

		var err error
		lockssCfg, err = lockss.NewFromFile(appCfg.ConfigFile())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:

		log.Debug("ConfigServerURL() is empty")
		log.Debug("Attempting to automatically retrieve config server value")

		var err error
		lockssCfg, err = lockss.New()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cfgSource = lockssCfg.PropsURL
	}

	log.Debugf("Full LOCKSS config object: %+v", lockssCfg)
	log.Debugf("Full App config object: %+v", appCfg)

	peersList, err := lockssCfg.IDInitialV3Peers.List()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(peersList) == 0 {
		fmt.Println("ERROR: No peers found in LOCKSS configuration file!")
		fmt.Println("No peers to check, exiting.")
		os.Exit(1)
	}

	log.Debugf("%d peers listed in %s", len(peersList), cfgSource)

	if appCfg.LogLevel() == config.LogLevelDebug {

		for idx, peer := range peersList {
			fmt.Printf("\n##########################################\n")
			fmt.Printf("Peer %d: %+v\n", idx, peer)
			fmt.Printf(
				"Protocol: %q, IP Address: %q, Port: %d\n",
				peer.Protocol,
				peer.IPAddress,
				peer.Port,
			)

			fmt.Println("peer.Network():", peer.Network())
			fmt.Println("peer.String():", peer.String())
		}

	}

	// spin off X goroutines, where X is len(peersList) and check the peer.Port
	// returning the result on a channel?
	// create a context, pass to the goroutines
	// inside of the goroutines setup a deadline and then a select block
	// the select block can listen on a result and a timeout

	expectedResponses := len(peersList)

	// collect results here that are pulled off the channel used by
	// goroutines as they complete their work
	results := make(portchecks.Results, 0, expectedResponses)

	// setup a channel to funnel each result from a port check. The capacity
	// is set to mirror the number of peers in the network to reduce
	// collection delay.
	resultsChan := make(chan portchecks.Result, expectedResponses)

	for _, peer := range peersList {
		go func(peer lockss.V3Peer, connTimeout time.Duration) {
			log.Infof("Checking port %d on %s ...", peer.Port, peer.IPAddress)
			resultsChan <- portchecks.CheckPort(peer, connTimeout)
		}(peer, appCfg.PortConnectTimeout())
	}

	// Collect all responses, continue until we exhaust the number of expected
	// responses calculated earlier as our signal to stop collecting responses
	remainingResponses := expectedResponses
	for remainingResponses > 0 {
		result := <-resultsChan
		results = append(results, result)
		remainingResponses--
	}

	results.PrintSummary()

}
