// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"flag"

	"github.com/apex/log"
)

// handleFlagsConfig wraps flag setup code into a bundle for potential ease of
// use and future testability
func (c *Config) handleFlagsConfig() {

	log.Debugf("Before parsing flags: %v", c.String())

	flag.Var(&c.nodePorts, "p", nodePortFlagHelp+" (shorthand)")
	flag.Var(&c.nodePorts, "port", nodePortFlagHelp)

	flag.StringVar(&c.configServerURL, "cs", "", configServerURLFlagHelp+" (shorthand)")
	flag.StringVar(&c.configServerURL, "config-server", "", configServerURLFlagHelp)

	flag.IntVar(&c.configServerReadTimeout, "ct", defaultConfigReadTimeout, configReadTimeoutFlagHelp+" (shorthand)")
	flag.IntVar(&c.configServerReadTimeout, "config-read-timeout", defaultConfigReadTimeout, configReadTimeoutFlagHelp)

	flag.IntVar(&c.portConnectTimeout, "pt", defaultPortConnectTimeout, portConnectTimeoutFlagHelp+" (shorthand)")
	flag.IntVar(&c.portConnectTimeout, "port-timeout", defaultPortConnectTimeout, portConnectTimeoutFlagHelp)

	flag.StringVar(&c.configFile, "cf", defaultConfigFile, configFileFlagHelp+" (shorthand)")
	flag.StringVar(&c.configFile, "config-file", defaultConfigFile, configFileFlagHelp)

	flag.BoolVar(&c.showVersion, "version", defaultDisplayVersionAndExit, versionFlagHelp)
	flag.BoolVar(&c.showVersion, "v", defaultDisplayVersionAndExit, versionFlagHelp+" (shorthand)")

	// create shorter and longer logging level flag options
	flag.StringVar(&c.logLevel, "ll", defaultLogLevel, logLevelFlagHelp+" (shorthand)")
	flag.StringVar(&c.logLevel, "log-level", defaultLogLevel, logLevelFlagHelp)

	// create shorter and longer logging format flag options
	flag.StringVar(&c.logFormat, "lf", defaultLogFormat, logFormatFlagHelp+" (shorthand)")
	flag.StringVar(&c.logFormat, "log-format", defaultLogFormat, logFormatFlagHelp)

	flag.Usage = flagsUsage()
	flag.Parse()
}
