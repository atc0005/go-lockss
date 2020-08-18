// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

///////////////////////////////////////////////////////////////////////////////
// NOTE: This content is not currently used (stubbed out early on) and may be
// removed in a future release.
///////////////////////////////////////////////////////////////////////////////

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
