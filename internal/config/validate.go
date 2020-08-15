// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Package config provides types and functions to collect, validate and apply
// user-provided settings.
package config

import (
	"fmt"

	"github.com/apex/log"
	"github.com/atc0005/go-lockss/internal/caller"
)

// TCP port ranges
// http://www.iana.org/assignments/port-numbers
// Port numbers are assigned in various ways, based on three ranges: System
// Ports (0-1023), User Ports (1024-49151), and the Dynamic and/or Private
// Ports (49152-65535)
const (
	TCPReservedPort            int = 0
	TCPSystemPortStart         int = 1
	TCPSystemPortEnd           int = 1023
	TCPUserPortStart           int = 1024
	TCPUserPortEnd             int = 49151
	TCPDynamicPrivatePortStart int = 49152
	TCPDynamicPrivatePortEnd   int = 65535
)

// Validate verifies all struct fields have been provided acceptable values
func (c Config) Validate() error {

	myFuncName := caller.GetFuncName()

	// Re-enable if switching to the alexflint/go-arg package.
	// if c.ConfigFile() == "" {
	// 	return fmt.Errorf(
	// 		"%s: missing fully-qualified path to config file to load",
	// 		myFuncName,
	// 	)
	// }
	// log.Debugf("c.ConfigFile() validates: %#v", c.ConfigFile())

	if len(c.UserNodePorts()) > 0 {

		for _, port := range c.UserNodePorts() {
			switch {

			case (port >= TCPSystemPortStart) && (port <= TCPDynamicPrivatePortEnd):
				log.Debugf(
					"Port %d is within the range of %d and %d",
					port,
					TCPSystemPortStart,
					TCPDynamicPrivatePortEnd,
				)

			default:
				errMsg := fmt.Sprintf(
					"invalid port %d specified; outside of the range of %d and %d",
					port,
					TCPSystemPortStart,
					TCPDynamicPrivatePortEnd,
				)
				log.Debugf("%s: %v", myFuncName, errMsg)

				return fmt.Errorf(errMsg)
			}
		}

	}
	log.Debugf("c.UserNodePorts() validates: (%d entries) %#v", len(c.UserNodePorts()), c.UserNodePorts())

	if c.PortConnectTimeout() == 0 {
		return fmt.Errorf(
			"%s: invalid port connect timeout value provided: %v",
			myFuncName,
			c.PortConnectTimeout(),
		)
	}
	log.Debugf("c.PortConnectTimeout() validates: %#v", c.PortConnectTimeout())

	if c.ConfigServerReadTimeout() == 0 {
		return fmt.Errorf(
			"%s: invalid port connect timeout value provided: %v",
			myFuncName,
			c.ConfigServerReadTimeout(),
		)
	}
	log.Debugf("c.ConfigServerReadTimeout() validates: %#v", c.ConfigServerReadTimeout())

	switch c.LogLevel() {
	case LogLevelFatal:
	case LogLevelError:
	case LogLevelWarn:
	case LogLevelInfo:
	case LogLevelDebug:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for log level",
			myFuncName,
			c.LogLevel(),
		)
	}
	log.Debugf("c.LogLevel() validates: %#v", c.LogLevel())

	switch c.LogFormat() {
	case LogFormatCLI:
	case LogFormatJSON:
	case LogFormatLogFmt:
	case LogFormatText:
	case LogFormatDiscard:
	default:
		return fmt.Errorf(
			"%s: invalid option %q provided for log format",
			myFuncName,
			c.LogFormat(),
		)
	}
	log.Debugf("c.LogFormat() validates: %#v", c.LogFormat())

	// Optimist
	log.Debug("All validation checks pass")
	return nil

}
