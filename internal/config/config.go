// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/atc0005/go-lockss/internal/caller"
)

// version reflects the application version. This is overridden via Makefile
// for release builds.
var version = "dev build"

// myAppName is the branded name of this application/project. This value will
// be used in user-facing output.
const myAppName string = "go-lockss"

// myAppURL is the branded homepage or project repo location. This value will
// be used in user-facing output.
const myAppURL string = "https://github.com/atc0005/" + myAppName

const (
	versionFlagHelp            = "Whether to display application version and then immediately exit application."
	logLevelFlagHelp           = "Log message priority filter. Log messages with a lower level are ignored."
	logFormatFlagHelp          = "Log messages are written in this format"
	configFileFlagHelp         = "Fully-qualified path to the fully-qualified path to an on-disk copy of the LOCKSS configuration/property XML file, usually named lockss.xml. This is NOT the same file as the /etc/lockss/config.dat file used to bootstrap the LOCKSS daemon at startup time."
	configServerURLFlagHelp    = "Fully-qualified URL to the LOCKSS configuration/property XML file."
	configReadTimeoutFlagHelp  = "Maximum number of seconds allowed for a request for the LOCKSS configuration XML file before timing out."
	portConnectTimeoutFlagHelp = "Maximum number of seconds allowed for a connection test against a remote TCP port before timing out."
	nodePortFlagHelp           = "Additional TCP port to connect to on remote LOCKSS nodes to verify connectivity. This flag may be repeated for each additional TCP port to check. If not set, this application connects only to the port (usually 9729) specified in the LOCKSS configuration/property XML file."
)

// shorthandFlagSuffix is appended to short flag help text to emphasize that
// the flag is a shorthand version of a longer flag.
const shorthandFlagSuffix = " (shorthand)"

// Default flag settings if not overridden by user input. Some constants are
// untyped in order to allow type promotion as needed.
const (
	defaultLogLevel              string = "info"
	defaultLogFormat             string = "text"
	defaultDisplayVersionAndExit bool   = false
	defaultConfigFile            string = ""
	defaultConfigServerURL       string = ""
	defaultConfigReadTimeout            = 10
	defaultPortConnectTimeout           = 2
)

// multiIntValueFlag is a custom type that satisfies the flag.Value interface
// in order to accept multiple integer values.
type multiValueIntFlag []int64

// String returns a comma separated string consisting of all slice elements
func (i *multiValueIntFlag) String() string {

	// From the `flag` package docs:
	// "The flag package may call the String method with a zero-valued
	// receiver, such as a nil pointer."
	if i == nil {
		return ""
	}

	numStrs := make([]string, 5)
	for _, item := range *i {
		numStrs = append(numStrs, strconv.FormatInt(item, 10))
	}

	return strings.Join(numStrs, ",")
}

// Set is called once by the flag package, in command line order, for each
// flag present
func (i *multiValueIntFlag) Set(value string) error {

	myFuncName := caller.GetFuncName()

	num, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return fmt.Errorf(
			"%s: error converting provided value %q to int: %w",
			myFuncName,
			value,
			err,
		)
	}

	*i = append(*i, num)

	return nil
}

// Config is a unified set of configuration values for this application. This
// struct is configured via command-line flags. The majority of values held by
// this object are intended to be retrieved via "Getter" methods.
type Config struct {

	// showVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application
	showVersion bool

	// nodePorts is a list of the TCP ports that should be open to each LOCKSS
	// node in the network.
	nodePorts multiValueIntFlag

	// logLevel is the chosen logging level
	logLevel string

	// logFormat controls which output format is used for log messages
	// generated by this application. This value is from a smaller subset
	// of the formats supported by the third-party leveled-logging package
	// used by this application.
	logFormat string

	// portConnectTimeout is the maximum number of seconds allowed for a
	// connection test against a remote TCP port before timing out.
	portConnectTimeout int

	// configServerReadTimeout is the maximum number of seconds allowed for a
	// request for the LOCKSS configuration XML file before timing out.
	configServerReadTimeout int

	// configServerURL is the fully-qualified URL to the LOCKSS
	// configuration/property XML file.
	configServerURL string

	// configFile is the fully-qualified path to an on-disk copy of the LOCKSS
	// configuration/property XML file, usually named lockss.xml. This is NOT
	// the same file as the /etc/lockss/config.dat file used to bootstrap the
	// LOCKSS daemon at startup time.
	configFile string
}

func (c Config) String() string {
	return fmt.Sprintf(
		"Config: { ConfigFile: %v, ConfigServerURL: %v, ConfigServerReadTimeout: %v, "+
			"PortConnectTimeout: %v, LogFormat: %v, LogLevel: %v, UserNodePorts: %v, "+
			"ShowVersion: %v}",
		c.ConfigFile(),
		c.ConfigServerURL(),
		c.ConfigServerReadTimeout(),
		c.PortConnectTimeout(),
		c.LogFormat(),
		c.LogLevel(),
		c.UserNodePorts(),
		c.ShowVersion(),
	)
}

// Version emits application version string.
func Version() string {
	return version
}

// Branding is responsible for emitting application name, version and origin.
func Branding() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// MyBinaryName returns the name of this binary
func MyBinaryName() string {
	return filepath.Base(os.Args[0])
}

// flagsUsage displays branding information and general usage details
func flagsUsage() func() {

	return func() {

		Branding()

		_, _ = fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage of \"%s\":\n",
			MyBinaryName(),
		)
		flag.PrintDefaults()

	}
}

// NewConfig is a factory function that produces a new Config object based
// on user provided flag and config file values.
func NewConfig() (*Config, error) {

	var config Config

	config.handleFlagsConfig()

	// Apply initial logging settings based on any provided CLI flags
	config.configureLogging()

	log.Debugf("After parsing flags: %v", config.String())

	// Return immediately if user just wants version details
	if config.ShowVersion() {
		return &config, nil
	}

	//
	// Attempt to load requested config file, fallback to known alternates
	// if user did not specify a config file
	//

	// Apply logging settings based on any provided config file settings
	config.configureLogging()

	log.Debug("Validating configuration ...")
	if err := config.Validate(); err != nil {
		// flag.Usage()
		// Let app handle this directly
		return nil, err
	}
	log.Debug("Configuration validated")

	// log.Debugf("Config object: %v", config.String())

	return &config, nil

}
