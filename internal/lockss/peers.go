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

	peers := make([]V3Peer, 0, 5)

	re, regExCompileErr := regexp.Compile(v3PeerRegex)
	if regExCompileErr != nil {
		return nil, fmt.Errorf("error compiling regex: %w", regExCompileErr)
	}

	xpathExpression := "//property[@name='id.initialV3PeerList']//list/value"

	xmlqueryNodes, err := xmlquery.QueryAll(l.xmlDoc, xpathExpression)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred running XPath query: %w",
			myFuncName,
			err,
		)
	}

	// If the query failed to find a match then this will be nil. We should
	// halt further attempts to retrieve more specific values.
	if xmlqueryNodes == nil {
		return nil, fmt.Errorf(
			"%s: unable to query value using xpathExpression: %q",
			myFuncName,
			xpathExpression,
		)
	}

	for _, prop := range xmlqueryNodes {
		// fmt.Printf("Raw XML: %q\n", prop.OutputXML(true))
		// fmt.Printf("Inner Text: %q\n", prop.InnerText())

		// Trim away protocol/port to obtain IP Addresses
		// TODO: Perhaps create a type instead that contains all values?
		// if we split on ':', we can capture:
		//
		// TCP
		// [1.2.3.4]
		// 9729
		//
		// For the bracketed IP Address, we can just trim off the brackets
		elementText := prop.InnerText()
		// elementText = strings.TrimPrefix(elementText, "TCP:[")
		// elementText = strings.TrimSuffix(elementText, "]:9729")

		matches := re.FindStringSubmatch(elementText)

		// we should only have one
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

		peers = append(peers, peer)

	}

	return peers, nil

}
