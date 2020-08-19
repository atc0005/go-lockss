// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

import (
	"context"
	"fmt"

	"github.com/antchfx/xmlquery"
	"github.com/atc0005/go-lockss/internal/caller"
)

// Config represents the values retrieved from provided LOCKSS configuration
// files. Due to the large number of possible parameters, we only pull in the
// settings that we actually require for our use.
type Config struct {

	// xmlDoc is the parse tree generated from the LOCKSS
	// properties/configuration XML file by the xmlquery package.
	xmlDoc *xmlquery.Node

	// IDInitialV3Peers is the list of V3 peers for this LOCKSS node. This
	// field also provides a user-facing API surface for peer nodes.
	IDInitialV3Peers IDInitialV3Peers

	// propsURL is the URL for the LOCKSS properties/configuration file. This
	// configuration file provides settings required for LOCKSS nodes to
	// participate in a network. Among other settings, this file specifies
	// peer nodes that each LOCKSS node will need to vote/poll/etc.
	propsURL string
}

// New attempts to automatically provide an initialized configuration using
// the default path to the local LOCKSS node configuration file. Any error
// that occurs is returned.
func New() (*Config, error) {

	myFuncName := caller.GetFuncName()

	// parse the local LOCKSS Daemon configuration file for available settings
	localDaemonCfg, err := getLocalDaemonConfig(DefaultConfigFile, ConfigFileCommentChar)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving local LOCKSS daemon configuration settings from %q",
			myFuncName,
			DefaultConfigFile,
		)
	}

	// Retrieve the URL configured within the local LOCKSS daemon
	// configuration file
	propsURL, err := localDaemonCfg.get(PropertiesURLVarName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesURLVarName,
			err,
		)
	}

	preservationGroup, err := localDaemonCfg.get(PropertiesGroupName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesGroupName,
			err,
		)
	}

	cfg := Config{
		propsURL: propsURL,
		IDInitialV3Peers: IDInitialV3Peers{
			preservationGroup: preservationGroup,
		},
	}

	return cfg.loadFromPropsURL(propsURL)
}

// NewFromFile attempts to provide an initialized configuration using a
// user-specified path to a local LOCKSS Properties (Parameters) XML file. Any
// error that occurs is returned. This function is usually reserved for
// testing purposes as live nodes will read their configuration using a
// provided URL.
func NewFromFile(filename string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	// parse the local LOCKSS Daemon configuration file for available settings
	localDaemonCfg, err := getLocalDaemonConfig(DefaultConfigFile, ConfigFileCommentChar)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving local LOCKSS daemon configuration settings from %q",
			myFuncName,
			DefaultConfigFile,
		)
	}

	// Retrieve the URL configured within the local LOCKSS daemon
	// configuration file
	propsURL, err := localDaemonCfg.get(PropertiesURLVarName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesURLVarName,
			err,
		)
	}

	preservationGroup, err := localDaemonCfg.get(PropertiesGroupName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesGroupName,
			err,
		)
	}

	cfg := Config{
		propsURL: propsURL,
		IDInitialV3Peers: IDInitialV3Peers{
			preservationGroup: preservationGroup,
		},
	}

	return cfg.loadFromPropsFile(filename)
}

// NewFromURL attempts to provide an initialized configuration using a
// user-specified URL to a LOCKSS Properties (Parameters) XML file. Any error
// that occurs is returned.
func NewFromURL(url string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	// parse the local LOCKSS Daemon configuration file for available settings
	localDaemonCfg, err := getLocalDaemonConfig(DefaultConfigFile, ConfigFileCommentChar)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving local LOCKSS daemon configuration settings from %q",
			myFuncName,
			DefaultConfigFile,
		)
	}

	preservationGroup, err := localDaemonCfg.get(PropertiesGroupName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesGroupName,
			err,
		)
	}

	cfg := Config{
		// use the URL provided by the user, not the one defined within the
		// local LOCKSS daemon configuration file
		propsURL: url,
		IDInitialV3Peers: IDInitialV3Peers{
			preservationGroup: preservationGroup,
		},
	}

	return cfg.loadFromPropsURL(url)
}

// NewFromURLWithContext attempts to provide an initialized configuration
// using a user-specified context and a URL to a LOCKSS Properties
// (Parameters) XML file. Any error that occurs is returned.
func NewFromURLWithContext(ctx context.Context, url string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	// parse the local LOCKSS Daemon configuration file for available settings
	localDaemonCfg, err := getLocalDaemonConfig(DefaultConfigFile, ConfigFileCommentChar)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving local LOCKSS daemon configuration settings from %q",
			myFuncName,
			DefaultConfigFile,
		)
	}

	preservationGroup, err := localDaemonCfg.get(PropertiesGroupName)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q: %w",
			myFuncName,
			PropertiesGroupName,
			err,
		)
	}

	cfg := Config{
		// use the URL provided by the user, not the one defined within the
		// local LOCKSS daemon configuration file
		propsURL: url,
		IDInitialV3Peers: IDInitialV3Peers{
			preservationGroup: preservationGroup,
		},
	}

	return cfg.loadFromPropsURLWithContext(ctx, url)
}

// SetXMLDoc is used by config loading methods to set the parse tree for later
// processing
func (c *Config) SetXMLDoc(xmlDoc *xmlquery.Node) {
	c.xmlDoc = xmlDoc
	c.IDInitialV3Peers.xmlDoc = xmlDoc
}

// LocalConfigFile is the local copy of the configuration file required
// for the LOCKSS daemon to operate. This file can be thought of as the
// "bootstrap" configuration file which enables the daemon to startup and
// retrieve further configuration settings from the network
// property/configuration server.
func (c Config) LocalConfigFile() string {
	return DefaultConfigFile
}

// PropsURL returns the URL for the LOCKSS properties/configuration file. If
// the user provided the URL, that URL is returned, otherwise the URL
// specified by the local LOCKSS daemon configuration file is returned.
func (c Config) PropsURL() string {
	return c.propsURL
}

// PreservationGroup is the group defined in the local LOCKSS daemon
// configuration file that the node is a member of. Membership in this group
// is required in order to receive certain settings, including collections of
// peer nodes. Not all LOCKSS networks will filter settings based on this
// group.
func (c Config) PreservationGroup() string {
	return c.IDInitialV3Peers.preservationGroup
}
