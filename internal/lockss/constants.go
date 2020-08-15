// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// References:
//
// https://github.com/lockss/lockss-daemon/blob/3cd40b6915b21424fd9d03eb9ab33d6813bc0271/rpms/BUILD/lockss-daemon-root/etc/lockss/hostconfig
// https://github.com/lockss/lockss-daemon/blob/3cd40b6915b21424fd9d03eb9ab33d6813bc0271/rpms/BUILD/lockss-daemon-root/etc/init.d/lockss
// https://github.com/lockss/lockss-daemon/blob/3cd40b6915b21424fd9d03eb9ab33d6813bc0271/test/frameworks/clean_up_daemon/bin/clean_cache_cron

package lockss

import "time"

// DefaultConfigLoadURLTimeout specifies the default timeout for retrieving
// the LOCKSS properties/configuration file from the LOCKSS network
// configuration server.
const DefaultConfigLoadURLTimeout = 7 * time.Second

const (

	// DefaultConfigFile represents the path to the local LOCKSS configuration
	// file. This is not the properties or parameters file, but is instead the
	// configuration file that the LOCKSS daemon reads at startup to determine
	// core operational parameters.
	//
	// Among other settings, this file specifies the LOCKSS_PROPS_URL value
	// which the daemon uses to retrieve the LOCKSS Properties (Parameters)
	// file for further configuration. This path is hard-coded in LOCKSS v1
	// tooling, so this path should be fairly stable.
	DefaultConfigFile string = "/etc/lockss/config.dat"

	// PropertiesURLVarName is the name of the configuration file variable that we
	// will use to retrieve the LOCKSS Properties (Parameters) URL.
	PropertiesURLVarName string = "LOCKSS_PROPS_URL"

	// ConfigFileCommentChar is the character used to indicate that a line is
	// commented out in the local LOCKSS configuration file.
	ConfigFileCommentChar string = "#"
)
