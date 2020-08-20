// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/atc0005/go-lockss/internal/caller"
)

// Allow for potential leading and trailing whitespace, but leave it out
// of the match groups.
const v3PeerRegex string = `^\s*(TCP):\[([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)\]:([0-9]+)\s*$`
const v3PeerExpectedFormat string = "TCP:[1.2.3.4]:9729"
const v3PeerRegexExpectedMatches int = 4

// IDInitialV3Peers represents the initial list of V3 peers for a LOCKSS node.
type IDInitialV3Peers struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	xmlDoc *xmlquery.Node

	// preservationGroup is the group defined in the local LOCKSS daemon
	// configuration file that the node is a member of. Membership in this
	// group is required in order to receive certain settings, including
	// collections of peer nodes. Not all LOCKSS networks will filter settings
	// based on this group.
	preservationGroup string
}

// V3Peer represents a peer LOCKSS node in the network. This type implements
// the net.Addr interface (https://golang.org/pkg/net/#Addr).
type V3Peer struct {

	// Protocol represents the network type (for example, "tcp", "udp")
	Protocol string

	// IPAddress is the IP Address of the LOCKSS node/peer that this node will
	// try to communicate with
	IPAddress string

	// LCAPPort is the LCAP (Library Content Audit Protocol) TCP port used to
	// make connections to remote LOCKSS nodes for content voting, polling and
	// repairing purposes.
	LCAPPort int
}

// Network returns the name of the network (for example, "tcp", "udp") in
// order to implement the net.Addr interface.
func (v3p V3Peer) Network() string {
	return strings.ToLower(v3p.Protocol)
}

// String returns the string from of the address, port pair in order to
// implement the net.Addr interface.
func (v3p V3Peer) String() string {
	return fmt.Sprintf("%s:%d", v3p.IPAddress, v3p.LCAPPort)
}

// List returns a slice of initial peers.
func (l IDInitialV3Peers) List() ([]V3Peer, error) {

	myFuncName := caller.GetFuncName()

	if l.xmlDoc == nil {
		return nil, fmt.Errorf(
			"%s: nil pointer dereference. Potential method call before config load",
			myFuncName,
		)
	}

	re, regExCompileErr := regexp.Compile(v3PeerRegex)
	if regExCompileErr != nil {
		return nil, fmt.Errorf("error compiling regex: %w", regExCompileErr)
	}

	/*
		    All of these XPath expressions appear to give the intended results
		    of restricting peers to the 'test' element with a 'prod' attribute
		    (the Preservation Group name).

			---

		    These expressions (first exact, the second more loose) apply a
		    preceding sibling filter to a 'then' element, then descend the
		    (exact) path to the peer list values.

		   /lockss-config/property[@name='org.lockss']/if/then[preceding-sibling::or/test[@group='prod']]/property[@name='id.initialV3PeerList']/list/value
		   //then[preceding-sibling::or/test[@group='prod']]/property[@name='id.initialV3PeerList']/list/value

			---

		   These expressions first target the 'test' element with a 'prod'
		   attribute then walk back a specified number of steps (2 in this
		   case) then descend the (exact) path to the peer list values.

		   /lockss-config/property[@name='org.lockss']/if/or/test[@group='prod']/../../then/property[@name='id.initialV3PeerList']/list/value
		   //test[@group='prod']/../../then/property[@name='id.initialV3PeerList']/list/value

	*/

	// xpathExpPresGroupUsed is an expression intended to determine if the
	// Preservation Group is used as a restriction within the LOCKSS
	// properties/configuration XML file.
	xpathExpPresGroupUsed := fmt.Sprintf(
		"//test[@group='%s']",
		l.preservationGroup,
	)

	// xpathExpPeersNoGroup is used if the preservation group defined for this
	// node is not in use within the LOCKSS properties/configuration XML file
	xpathExpPeersNoGroup := "//property[@name='org.lockss']" +
		"/property[@name='id.initialV3PeerList']/list/value"

	// xpathExpPeersWithGroup is used if the preservation group defined for
	// this node IS in use within the LOCKSS properties/configuration XML file
	xpathExpPeersWithGroup := fmt.Sprintf(
		"/lockss-config/property[@name='org.lockss']/if"+
			"/then[preceding-sibling::or/test[@group='%s']]"+
			"/property[@name='id.initialV3PeerList']/list/value",
		l.preservationGroup,
	)

	logger.Printf(
		"%s: using %q XPath expression to determine whether the %q preservation group is in use",
		myFuncName,
		xpathExpPresGroupUsed,
		l.preservationGroup,
	)
	xmlQueryNodeGroupUsed, xmlQueryErr := xmlquery.Query(l.xmlDoc, xpathExpPresGroupUsed)
	if xmlQueryErr != nil {
		return nil, fmt.Errorf(
			"%s: error occurred running XPath query %q: %w",
			myFuncName,
			xpathExpPresGroupUsed,
			xmlQueryErr,
		)
	}

	// match preallocated slice size for []V3Peer)
	xmlqueryNodes := make([]*xmlquery.Node, 0, 10)

	switch {
	// if the preservation group is in use ...
	case xmlQueryNodeGroupUsed != nil:

		logger.Printf("%s: %q preservation group IS in use", myFuncName, l.preservationGroup)

		logger.Printf(
			"%s: using %q XPath expression to retrieve group-based peer nodes",
			myFuncName,
			xpathExpPeersWithGroup,
		)
		xmlqueryNodes, xmlQueryErr = xmlquery.QueryAll(l.xmlDoc, xpathExpPeersWithGroup)
		if xmlQueryErr != nil {
			return nil, fmt.Errorf(
				"%s: error occurred running XPath query %q: %w",
				myFuncName,
				xpathExpPeersWithGroup,
				xmlQueryErr,
			)
		}

		// If we didn't find any peer nodes restricted to this LOCKSS node's
		// preservation group, maybe the preservation group is used to
		// restrict other settings and *not* the peer nodes. Try again, this
		// time without the preservation group restriction.
		if xmlqueryNodes == nil {

			logger.Printf(
				"%s: Unable to retrieve group-based peer nodes using %q XPath expression",
				myFuncName,
				xpathExpPeersWithGroup,
			)

			logger.Printf(
				"%s: using %q XPath expression to retrieve non-group peer nodes",
				myFuncName,
				xpathExpPeersNoGroup,
			)
			xmlqueryNodes, xmlQueryErr = xmlquery.QueryAll(l.xmlDoc, xpathExpPeersNoGroup)
			if xmlQueryErr != nil {
				return nil, fmt.Errorf(
					"%s: error occurred running XPath query %q: %w",
					myFuncName,
					xpathExpPeersNoGroup,
					xmlQueryErr,
				)
			}

			// Give up if we were unable to find a match using either
			// expression.
			if xmlqueryNodes == nil {

				logger.Printf(
					"%s: unable to retrieve non-group peer nodes using %q XPath expression",
					myFuncName,
					xpathExpPeersNoGroup,
				)

				return nil, fmt.Errorf(
					"%s: unable to retrieve peer nodes using XPath expression %q or %q",
					myFuncName,
					xpathExpPeersWithGroup,
					xpathExpPeersNoGroup,
				)
			}
		}

	// if the preservation group is *NOT* in use ...
	case xmlQueryNodeGroupUsed == nil:

		logger.Printf("%s: %q preservation group is NOT in use", myFuncName, l.preservationGroup)

		logger.Printf(
			"%s: using %q XPath expression to retrieve non-group peer nodes",
			myFuncName,
			xpathExpPeersNoGroup,
		)
		xmlqueryNodes, xmlQueryErr = xmlquery.QueryAll(l.xmlDoc, xpathExpPeersNoGroup)
		if xmlQueryErr != nil {
			return nil, fmt.Errorf(
				"%s: error occurred running XPath query %q: %w",
				myFuncName,
				xpathExpPeersNoGroup,
				xmlQueryErr,
			)
		}

		// Give up if we were unable to find a match.
		if xmlqueryNodes == nil {

			logger.Printf(
				"%s: unable to retrieve non-group peer nodes using %q XPath expression",
				myFuncName,
				xpathExpPeersNoGroup,
			)
			return nil, fmt.Errorf(
				"%s: unable to query value using xpathExpression: %q",
				myFuncName,
				xpathExpPeersNoGroup,
			)
		}
	}

	logger.Printf(
		"%s: Successfully retrieved %d peer node search results",
		myFuncName,
		len(xmlqueryNodes),
	)

	peers := make([]V3Peer, 0, 10)
	for _, prop := range xmlqueryNodes {
		elementText := prop.InnerText()
		matches := re.FindStringSubmatch(elementText)

		// each "node" result should have a very specific number of fields
		// that we will use to populate our V3Peers list
		if len(matches) != v3PeerRegexExpectedMatches {
			return nil, fmt.Errorf(
				"%s: peer entry %q did not match expected format of '%s'",
				myFuncName,
				elementText,
				v3PeerExpectedFormat,
			)
		}

		peerProtocol := matches[1]
		peerIPAddress := matches[2]
		peerPort, strConvErr := strconv.ParseInt(matches[3], 10, 32)
		if strConvErr != nil {
			return nil, fmt.Errorf(
				"%s: error occurred converting peer port: %w",
				myFuncName,
				strConvErr,
			)
		}

		peer := V3Peer{
			Protocol:  peerProtocol,
			IPAddress: peerIPAddress,
			LCAPPort:  int(peerPort),
		}

		logger.Printf(
			"%s: parsed peer [%s, %s, %d] from search results",
			myFuncName,
			peer.Protocol,
			peer.IPAddress,
			peer.LCAPPort,
		)
		peers = append(peers, peer)

	}

	return peers, nil

}
