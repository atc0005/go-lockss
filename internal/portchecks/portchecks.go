// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package portchecks

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/apex/log"
	"github.com/atc0005/go-lockss/internal/caller"
)

// Results is a collection of port check results. Intended for aggregation
// before bulk processing of some kind.
type Results []Result

// Result is the outcome of a port check.
type Result struct {
	Host  string
	Port  int
	Open  bool
	Error error
}

// CheckPort accepts a net.Addr interface type and a timeout. A connection
// check is performed to determine whether the port is reachable. The result
// of this check is returned as a Result value for later processing. Any
// errors which occur as part of this check are recorded as a partial check
// Result.
func CheckPort(netAddr net.Addr, timeout time.Duration) Result {

	myFuncName := caller.GetFuncName()

	// split out net.Addr into host, port values
	host, portStr, err := net.SplitHostPort(netAddr.String())
	if err != nil {
		return Result{
			Host: host,
			Error: fmt.Errorf(
				"%s: failed to split net.Addr into host/port values: %w",
				myFuncName,
				err,
			),
		}
	}

	// convert port string into int64
	portInt, strConvErr := strconv.ParseInt(portStr, 10, 32)
	if strConvErr != nil {
		return Result{
			Host: host,
			Error: fmt.Errorf(
				"%s: error occurred converting port string: %w",
				myFuncName,
				strConvErr,
			),
		}
	}

	// force int64 to int type in order to fit into our Result struct
	port := int(portInt)

	log.Debugf("Host: %s", host)
	log.Debugf("Port: %d", port)

	conn, err := net.DialTimeout(netAddr.Network(), netAddr.String(), timeout)
	if err != nil {

		log.Debugf(
			"%s: error connecting to port %d/%s on %s: %w",
			myFuncName,
			port,
			netAddr.Network(),
			host,
			err,
		)

		// the net package errors are usually pretty expressive, so there
		// isn't a lot that we need to add right now.
		return Result{
			Host:  host,
			Port:  port,
			Error: err,
		}
	}

	// deferring this so that we can return our result based on the connection
	// check and handle the connection close attempt separately.
	defer func() {
		if err := conn.Close(); err != nil {
			log.Errorf(
				"%s: failed to close connection to port %d/%s on %s: %v",
				myFuncName,
				port,
				netAddr.Network(),
				host,
				err,
			)
		}
	}()

	return Result{
		Host:  host,
		Port:  port,
		Open:  true,
		Error: nil,
	}

}

// Reachable returns the number of open ports found.
func (rs Results) Reachable() int {

	var r int

	for _, i := range rs {
		if i.Open {
			r++
		}
	}

	return r
}
