// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"time"

	"github.com/apex/log"
)

// NodePorts returns the user-provided LOCKSS node ports to check or the
// default port if no ports were specified.
func (c Config) NodePorts() []int64 {

	switch {
	case c.nodePorts != nil:
		return c.nodePorts
	default:
		var defaultPort multiValueIntFlag
		defaultPort = append(defaultPort, defaultNodePort)
		return defaultPort
	}
}

// LogLevel returns the user-provided logging level or empty string if not
// provided. CLI flag values take precedence if provided.
func (c Config) LogLevel() string {

	switch {
	case c.logLevel != "":
		return c.logLevel
	default:
		return ""
	}
}

// LogFormat returns the user-provided logging format or empty string if not
// provided. CLI flag values take precedence if provided.
func (c Config) LogFormat() string {

	switch {
	case c.logFormat != "":
		return c.logFormat
	default:
		return ""
	}
}

// ConfigServerURL returns the user-provided URL to the LOCKSS
// configuration/property XML file.
func (c Config) ConfigServerURL() string {

	switch {
	case c.configServerURL != "":
		return c.configServerURL
	default:
		return ""
	}
}

// ConfigFile returns the user-provided path to the on-disk LOCKSS
// configuration/property XML file.
func (c Config) ConfigFile() string {

	switch {
	case c.configFile != "":
		return c.configFile
	default:
		return defaultConfigFile
	}
}

// ShowVersion returns the user-provided choice of displaying the application
// version and exiting or the default value for this choice.
func (c Config) ShowVersion() bool {
	return c.showVersion
}

// PortConnectTimeout returns the user-provided choice of what timeout value
// to use for port connection attempts queries. If not set, returns the
// default value for our application.
func (c Config) PortConnectTimeout() time.Duration {

	switch {
	case c.portConnectTimeout != 0:
		return time.Duration(c.portConnectTimeout) * time.Second
	default:
		log.Debugf(
			"Requested port connect timeout value not specified, using default: %v",
			defaultPortConnectTimeout,
		)
		return time.Duration(defaultPortConnectTimeout) * time.Second
	}
}

// ConfigServerReadTimeout returns the user-provided choice of what timeout
// value to use for attempts to read the remote LOCKSS configuration file. If
// not set, returns the default value for our application.
func (c Config) ConfigServerReadTimeout() time.Duration {

	switch {
	case c.configServerReadTimeout != 0:
		return time.Duration(c.configServerReadTimeout) * time.Second
	default:
		log.Debugf(
			"Requested config read timeout value not specified, using default: %v",
			defaultConfigReadTimeout,
		)
		return time.Duration(defaultConfigReadTimeout) * time.Second
	}
}
