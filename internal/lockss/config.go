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

// New attempts to automatically provide an initialized configuration using
// the default path to the local LOCKSS node configuration file. Any error
// that occurs is returned.
func New() (*Config, error) {

	myFuncName := caller.GetFuncName()

	cfg := Config{
		LocalConfigFile: DefaultConfigFile,
	}

	// fetch the props URL from the local config file
	propsURL, err := getLOCKSSPropsURL(DefaultConfigFile, PropertiesURLVarName, ConfigFileCommentChar)
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error occurred retrieving the value for %q from %q: %w",
			myFuncName,
			PropertiesURLVarName,
			DefaultConfigFile,
			err,
		)
	}

	return cfg.loadFromPropsURL(propsURL)
}

// NewFromFile attempts to provide an initialized configuration using a
// user-specified path to a local LOCKSS Properties (Parameters) XML file. Any
// error that occurs is returned. This function is usually reserved for
// testing purposes as live nodes will read their configuration using a
// provided URL.
func NewFromFile(filename string) (*Config, error) {
	cfg := Config{
		LocalConfigFile: DefaultConfigFile,
	}

	return cfg.loadFromPropsFile(filename)
}

// NewFromURL attempts to provide an initialized configuration using a
// user-specified URL to a LOCKSS Properties (Parameters) XML file. Any error
// that occurs is returned.
func NewFromURL(url string) (*Config, error) {
	cfg := Config{
		LocalConfigFile: DefaultConfigFile,
	}

	return cfg.loadFromPropsURL(url)
}

// NewFromURLWithContext attempts to provide an initialized configuration
// using a user-specified context and a URL to a LOCKSS Properties
// (Parameters) XML file. Any error that occurs is returned.
func NewFromURLWithContext(ctx context.Context, url string) (*Config, error) {
	cfg := Config{
		LocalConfigFile: DefaultConfigFile,
	}

	return cfg.loadFromPropsURLWithContext(ctx, url)
}

// InitConfig is called by config loading methods to set required fields
func (c *Config) InitConfig(xmlDoc *xmlquery.Node) {

	c.xmlDoc = xmlDoc
	c.IDInitialV3Peers = IDInitialV3Peers{
		xmlDoc: xmlDoc,
	}

}
