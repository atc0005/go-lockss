// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"os"

	"github.com/apex/log"

	"github.com/atc0005/go-lockss/internal/config"
	"github.com/atc0005/go-lockss/internal/lockss"
)

func main() {

	testFunc := func(c *lockss.Config, note string) {

		fmt.Println(note)

		// fmt.Println("Proxy Access Included IPs:", cfg.Proxy.Access.IPInclude)
		// fmt.Println("Proxy Access IP Log Forbidden:", cfg.Proxy.Access.IPLogForbidden)

		fmt.Printf("Full config object: %+v\n", c)

		peersList, err := c.IDInitialV3Peers.List()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for idx, peer := range peersList {
			fmt.Printf("\n##########################################\n")
			fmt.Printf("Peer %d: %+v\n", idx, peer)
			fmt.Printf(
				"Protocol: %q, IP Address: %q, Port: %d\n",
				peer.Protocol,
				peer.IPAddress,
				peer.LCAPPort,
			)

			fmt.Println("peer.Network():", peer.Network())
			fmt.Println("peer.String():", peer.String())
		}
	}

	// config.EnableLogging()
	lockss.DisableLogging()

	// setup application configuration
	appCfg, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// If user supplied values, we should use those to retrieve the LOCKSS
	// configuration from the central LOCKSS configuration server, otherwise
	// try to automatically determine values and go from there.
	var lockssCfg *lockss.Config
	switch {
	case appCfg.ConfigServerURL() != "":

		log.Debugf(
			"ConfigServerURL() is non-empty, using value %q",
			appCfg.ConfigServerURL(),
		)

		var err error
		lockssCfg, err = lockss.NewFromURL(appCfg.ConfigServerURL())
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
	}

	testFunc(lockssCfg, "LOCKSS configuration loaded")

	// fileCfg, err := lockss.NewFromFile(inputFile)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// testFunc(fileCfg, "Loading config from file")

	// urlCfg, err := lockss.NewFromURL(inputURL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// testFunc(urlCfg, "Loading config from URL")

	// connect to the

}
