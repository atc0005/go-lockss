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

// UserNodePorts returns the user-provided LOCKSS node ports to check or an
// empty slice if no ports were specified.
func (c Config) UserNodePorts() []int {

	// force conversion of int64 to int in order to be more flexible (e.g.,
	// the net standard library tends to work with `int` and `string` type
	// instead of `int64`)

	switch {
	case c.nodePorts != nil:
		var ports []int
		for i := range c.nodePorts {
			ports = append(ports, int(c.nodePorts[i]))
		}
		return ports
	default:
		return []int{}
	}
}

// LogLevel returns the user-provided logging level or default log level if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogLevel() string {

	switch {
	case c.logLevel != "":
		return c.logLevel
	default:
		return defaultLogLevel
	}
}

// LogFormat returns the user-provided logging format or default log format if
// not provided. CLI flag values take precedence if provided.
func (c Config) LogFormat() string {

	switch {
	case c.logFormat != "":
		return c.logFormat
	default:
		return defaultLogFormat
	}
}

// ConfigServerURL returns the user-provided URL to the LOCKSS
// configuration/property XML file or the default value if not provided.
func (c Config) ConfigServerURL() string {

	switch {
	case c.configServerURL != "":
		return c.configServerURL
	default:
		return defaultConfigServerURL
	}
}

// ConfigFile returns the user-provided path to the on-disk LOCKSS
// configuration/property XML file or the default value if not provided.
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
