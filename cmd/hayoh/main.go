// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

//go:generate go-winres make --product-version=git-tag --file-version=git-tag

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

	log.Debug("Initializing application")

	// lockss.EnableLogging()
	lockss.DisableLogging()

	// setup application configuration
	appCfg, err := config.NewConfig()
	switch {
	// TODO: How else to guard against nil cfg object?
	case appCfg != nil && appCfg.ShowVersion():
		_, _ = fmt.Fprintln(
			flag.CommandLine.Output(),
			config.Branding(),
		)
		os.Exit(0)
	case err == nil:
		// do nothing for this one
	case errors.Is(err, flag.ErrHelp):
		os.Exit(0)
	default:
		log.Errorf("failed to initialize application: %s", err)
		flag.Usage()
		os.Exit(1)
	}

	// if we have set the app logging level to Debug, enable lockss package
	// logging too
	// if logger, ok := log.Log.(*log.Logger); ok {
	// 	if logger.Level == log.DebugLevel {
	// 		lockss.EnableLogging()
	// 	}
	// }
	if appCfg.LogLevel() == config.LogLevelDebug {
		lockss.EnableLogging()
	}

	fmt.Printf(
		"\n[%v] Starting %s version %q ...\n",
		time.Now().Format("2006-01-02 15.04:05"),
		config.MyBinaryName(),
		config.Version(),
	)

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

		cfgSource = lockssCfg.PropsURL()

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
			log.Debugf(
				"Peer %d: [Protocol: %q, IP Address: %q, Port: %d, peer.Network(): %q, peer.String(): %q]",
				idx,
				peer.Protocol,
				peer.IPAddress,
				peer.LCAPPort,
				peer.Network(),
				peer.String(),
			)
		}
	}

	ports := append(appCfg.UserNodePorts(), peersList[0].LCAPPort)
	ports = portchecks.UniquePorts(ports...)

	numPorts := len(ports)
	numPeers := len(peersList)
	expectedResponses := numPeers * numPorts

	log.Debugf("Expected responses: %d", expectedResponses)

	// collect results here that are pulled off the channel used by
	// goroutines as they complete their work
	results := make(portchecks.Results, 0, expectedResponses)

	// setup a channel to funnel each result from a port check. The capacity
	// is set to mirror the number of peers in the network to reduce
	// collection delay.
	resultsChan := make(chan portchecks.Result, expectedResponses)

	fmt.Printf(
		"[%v] Checking %d ports on %d peer nodes ...\n",
		time.Now().Format("2006-01-02 15.04:05"),
		numPorts,
		numPeers,
	)

	for _, peer := range peersList {
		go func(peer lockss.V3Peer, ports []int, connTimeout time.Duration) {
			for _, port := range ports {
				log.Debugf("Checking port %d on %s ...", port, peer.IPAddress)
				resultsChan <- portchecks.CheckPort(peer, port, connTimeout)
			}
		}(peer, ports, appCfg.PortConnectTimeout())
	}

	// Collect all responses, continue until we exhaust the number of expected
	// responses calculated earlier as our signal to stop collecting responses
	remainingResponses := expectedResponses
	for remainingResponses > 0 {
		result := <-resultsChan
		results = append(results, result)
		remainingResponses--

		if remainingResponses > 0 {
			// skip emitting "Waiting" message if we're no longe waiting
			log.Debugf("Waiting on %d responses ...", remainingResponses)
			continue
		}

		log.Debug("All responses received")

	}

	results.PrintSummary()

}
