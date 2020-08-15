// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

import "github.com/antchfx/xmlquery"

// Config represents the values retrieved from provided LOCKSS configuration
// files. Due to the large number of possible parameters, we only pull in the
// settings that we actually require for our use.
type Config struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	xmlDoc *xmlquery.Node

	// IDInitialV3Peers is the list of V3 peers for this LOCKSS node
	IDInitialV3Peers IDInitialV3Peers

	// LocalConfigFile is the local copy of the configuration file required
	// for the LOCKSS daemon to operate. This file can be thought of as the
	// "bootstrap" configuration file which enables the daemon to startup and
	// retrieve further configuration settings from the network
	// property/configuration server.
	LocalConfigFile string

	// PropsURL is the URL for the LOCKSS properties/configuration file. This
	// configuration file provides settings required for LOCKSS nodes to
	// participate in a network. Among other settings, this file specifies
	// peer nodes that each LOCKSS node will need to vote/poll/etc.
	PropsURL string
}

// Proxy represents the parent proxy element containing settings related to
// the content proxy support in LOCKSS
type Proxy struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	//xmlDoc *xmlquery.Node

	NoManifestIndexResponses string
	Port                     int
	AuditPort                int
	Access                   ProxyAccess
}

// ProxyAccess represents access control settings for the content proxy
// support in LOCKSS
type ProxyAccess struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	//xmlDoc *xmlquery.Node

	IPLogForbidden bool
	IPInclude      []string
}

// IDInitialV3Peers represents the initial list of V3 peers for a LOCKSS node.
type IDInitialV3Peers struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	xmlDoc *xmlquery.Node
}
