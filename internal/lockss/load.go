// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package lockss

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/atc0005/go-lockss/internal/caller"
)

// configVar represents the configuration variables and values found in
// the local LOCKSS configuration file.
type configVar struct {

	// Name is the variable name found in the LOCKSS configuration file.
	Name string

	// Value is the value for a variable found in the LOCKSS configuration
	// file.
	Value string
}

// daemonConfig represents the key=value settings found within the local
// LOCKSS daemon configuration file.
type daemonConfig map[string]string

// get is a helper method for retrieving specific configuration settings found
// within the local LOCKSS daemon configuration file.
func (dc daemonConfig) get(setting string) (string, error) {

	myFuncName := caller.GetFuncName()

	if _, ok := dc[setting]; !ok {
		return "", fmt.Errorf(
			"%s: requested setting %q not found",
			myFuncName,
			setting,
		)
	}

	return dc[setting], nil
}

// loadFromPropsFile attempts to update the current configuration settings by
// using the user-specified path to an on-disk copy of the LOCKSS Properties
// (Parameters) XML file. Any error that occurs is returned. This function is
// usually reserved for testing purposes as live nodes will read their
// configuration using a provided URL.
func (c *Config) loadFromPropsFile(filename string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	logger.Printf(
		"%s: called at %s for filename %q \n",
		myFuncName,
		time.Now().Format(logTimeLayout),
		filename,
	)

	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, fmt.Errorf("error occurred opening file: %w", err)
	}

	// #nosec G307
	// Believed to be a false-positive from recent gosec release
	// https://github.com/securego/gosec/issues/714
	defer func() {
		if err := f.Close(); err != nil {
			logger.Printf(
				"Error occurred closing file %q: %v\n",
				filename,
				err,
			)
		}
	}()

	doc, err := xmlquery.Parse(f)
	if err != nil {
		return nil, fmt.Errorf("error occurred parsing XML file: %w", err)
	}

	// Update our config object with parsed XML data
	c.SetXMLDoc(doc)

	return c, nil

}

// loadFromPropsURL attempts to update the current configuration settings by
// using the user-specified URL to the LOCKSS Properties (Parameters) XML
// file. Any error that occurs is returned. This method is a convenience
// wrapper around the LoadFromURLWithContext method that uses a default
// timeout. If you wish to use a different timeout value, you should use the
// LoadFromURLWithContext method directly with an appropriate context to apply
// the desired timeout settings.
func (c *Config) loadFromPropsURL(url string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	logger.Printf(
		"%s: called at %s for url %q \n",
		myFuncName,
		time.Now().Format(logTimeLayout),
		url,
	)

	ctx, cancel := context.WithTimeout(context.Background(), DefaultConfigLoadURLTimeout)
	defer cancel()

	return c.loadFromPropsURLWithContext(ctx, url)

}

// loadFromURLWithContext accepts a context and a user-specified URL to a
// LOCKSS Properties (Parameters) XML file. This method attempts to update the
// current configuration settings by using that XML file. Any error that
// occurs is returned.
func (c *Config) loadFromPropsURLWithContext(ctx context.Context, url string) (*Config, error) {

	myFuncName := caller.GetFuncName()

	logger.Printf(
		"%s: called at %s for url %q \n",
		myFuncName,
		time.Now().Format(logTimeLayout),
		url,
	)

	if ctx.Err() != nil {
		logger.Println("Context has expired before validation:", time.Now().Format(logTimeLayout))
	}

	// Validate input data
	if valid, err := isValidURL(url); !valid {
		return nil, err
	}

	var configFileBuffer bytes.Buffer

	// prepare request
	req, newReqErr := http.NewRequestWithContext(ctx, http.MethodGet, url, &configFileBuffer)
	if newReqErr != nil {
		return nil, newReqErr
	}

	if ctx.Err() != nil {
		logger.Println("Context has expired before Do(req):", time.Now().Format(logTimeLayout))
	}

	var httpClient http.Client
	resp, httpClientErr := httpClient.Do(req)
	if ctx.Err() != nil {
		logger.Println("Context has expired after Do(req):", time.Now().Format(logTimeLayout))
	}
	if httpClientErr != nil {
		return nil, httpClientErr
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logger.Printf(
				"Error occurred closing response body: %v\n",
				err,
			)
		}
	}()

	// Get the response body directly. We'll create a bytes buffer later to
	// wrap the response data in order to provide an io.Reader where needed.
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	switch {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
	//
	// Informational responses (100-199)
	// Successful responses (200-299)
	// Redirects (300-399)
	// Client errors (400-499)
	// Server errors (500-599)
	case resp.StatusCode > 399:
		invalidStatusCodeErr := fmt.Errorf(
			"%s: error occurred requesting URL %q: %v, %q",
			myFuncName,
			url,
			resp.Status,
			string(responseData),
		)
		logger.Println(invalidStatusCodeErr)
		return nil, invalidStatusCodeErr

	case resp.StatusCode >= 300 && resp.StatusCode < 400:
		// TODO: Should we treat a redirect as anything of consequence?
		warningStatusCodeMsg := fmt.Sprintf(
			"%s: server issued redirect response when requesting URL %q: %v, %q",
			myFuncName,
			url,
			resp.Status,
			string(responseData),
		)
		logger.Println(warningStatusCodeMsg)
		// return nil, warningStatusCodeMsg

	}

	respBodyBuffer := bytes.NewBuffer(responseData)
	doc, err := xmlquery.Parse(respBodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("error occurred parsing XML file: %w", err)
	}

	if ctx.Err() != nil {
		logger.Println("Context has expired after parsing response body:", time.Now().Format(logTimeLayout))
	}

	// Update our config object with parsed XML data
	c.SetXMLDoc(doc)

	return c, nil
}

// isValidURL performs validation checks on the provided URL. Any encountered
// error is returned.
func isValidURL(configURL string) (bool, error) {

	myFuncName := caller.GetFuncName()

	logger.Printf(
		"%s: called at %s for url %q \n",
		myFuncName,
		time.Now().Format(logTimeLayout),
		configURL,
	)

	u, err := url.Parse(configURL)
	if err != nil {
		return false, fmt.Errorf(
			"%s: unable to parse provided URL %q: %w",
			myFuncName,
			configURL,
			err,
		)
	}

	logger.Printf("%s: Parsed URL: %q\n", myFuncName, u)

	return true, nil
}

// getLocalDaemonConfig parses the local LOCKSS daemon configuration file
// (usually /etc/lockss/config.dat) and returns a custom type wrapping a map
// of key/value pairs representing available settings. Empty configuration
// settings are returned.
func getLocalDaemonConfig(filename string, ignorePrefix string) (daemonConfig, error) {

	myFuncName := caller.GetFuncName()
	cleanFilename := filepath.Clean(filename)

	logger.Printf("%s: Request to open %q received", myFuncName, filename)
	logger.Printf("%s: Attempting to open sanitized version of file %q",
		myFuncName, cleanFilename)

	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, fmt.Errorf(
			"%s: error encountered opening file %q as %q: %w",
			myFuncName,
			filename,
			cleanFilename,
			err,
		)
	}

	// #nosec G307
	// Believed to be a false-positive from recent gosec release
	// https://github.com/securego/gosec/issues/714
	defer func() {
		if err := f.Close(); err != nil {
			// Ignore "file already closed" errors
			if !errors.Is(err, os.ErrClosed) {
				logger.Printf(
					"%s: failed to close file %q: %s",
					myFuncName,
					cleanFilename,
					err.Error(),
				)
			}
		}
	}()

	s := bufio.NewScanner(f)
	var lineno int
	var cfgVar configVar
	daemonConfig := make(daemonConfig)

	// TODO: Does Scan() perform any whitespace manipulation already?
	for s.Scan() {
		lineno++
		currentLine := s.Text()
		logger.Printf(
			"%s: Scanned line %d from %q: %q",
			myFuncName,
			lineno,
			filename,
			currentLine,
		)

		currentLine = strings.TrimSpace(currentLine)
		logger.Printf(
			"%s: Line %d from %q after lowercasing and whitespace removal: %q",
			myFuncName,
			lineno,
			filename,
			currentLine,
		)

		// explicitly ignore blank lines
		if currentLine == "" {
			logger.Printf(
				"%s: Line %d from %q appears to only contain whitespace, skipping ...",
				myFuncName,
				lineno,
				filename,
			)
			continue
		}

		// explicitly ignore lines beginning with specified pattern, if
		// provided
		if ignorePrefix != "" {
			if strings.HasPrefix(currentLine, ignorePrefix) {
				logger.Printf(
					"%s: Ignoring line %d due to leading %q",
					myFuncName,
					lineno,
					ignorePrefix,
				)
				continue
			}
		}

		// need to split on '='
		fields := strings.Split(currentLine, "=")

		// we're expecting a length of 2
		// VARIABLE=VALUE, split becomes [VARIABLE VALUE]
		if len(fields) != 2 {
			return nil, fmt.Errorf(
				"%s: error parsing daemon config file %q",
				myFuncName,
				filepath.Clean(filename),
			)
		}

		cfgVar.Name = strings.TrimSpace(fields[0])
		cfgVar.Value = strings.TrimSpace(fields[1])
		cfgVar.Value = strings.Trim(cfgVar.Value, `"'`)

		daemonConfig[cfgVar.Name] = cfgVar.Value

	}

	logger.Printf("%s: Exited s.Scan() loop", myFuncName)

	// report any errors encountered while scanning the input file
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf(
			"%s: error encountered while scanning %q: %w",
			myFuncName,
			filepath.Clean(filename),
			err,
		)
	}

	// explicitly close file, bail if failure occurs
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf(
			"%s: failed to close file %q: %w",
			myFuncName,
			filepath.Clean(filename),
			err,
		)
	}

	// otherwise, report that the requested param was not found
	return daemonConfig, nil

}
