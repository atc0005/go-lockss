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

	if c.NodePorts() == nil || len(c.NodePorts()) == 0 {
		return fmt.Errorf(
			"%s: one or more TCP ports not provided",
			myFuncName,
		)
	}
	log.Debugf("c.NodePorts() validates: (%d entries) %#v", len(c.NodePorts()), c.NodePorts())

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
