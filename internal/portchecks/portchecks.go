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
func CheckPort(netAddr net.Addr, port int, timeout time.Duration) Result {

	myFuncName := caller.GetFuncName()

	// split out net.Addr into host, port values
	// host, portStr, err := net.SplitHostPort(netAddr.String())
	host, _, err := net.SplitHostPort(netAddr.String())
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
	// portInt, strConvErr := strconv.ParseInt(portStr, 10, 32)
	// if strConvErr != nil {
	// 	return Result{
	// 		Host: host,
	// 		Error: fmt.Errorf(
	// 			"%s: error occurred converting port string: %w",
	// 			myFuncName,
	// 			strConvErr,
	// 		),
	// 	}
	// }

	// // force int64 to int type in order to fit into our Result struct
	// port := int(portInt)

	log.Debugf("%s: Host: %s", myFuncName, host)
	log.Debugf("%s: Port: %d", myFuncName, port)

	conn, err := net.DialTimeout(netAddr.Network(), netAddr.String(), timeout)
	if err != nil {

		log.Debugf(
			"%s: error connecting to port %d/%s on %s: %v",
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

// UniquePorts accepts one or more ports of potentially duplicated ports and
// returns a slice of unique ports
func UniquePorts(ports ...int) []int {

	uniquePorts := make(map[int]int)

	for _, port := range ports {
		uniquePorts[port] = port
	}

	portsList := make([]int, 0, len(uniquePorts))
	for _, port := range uniquePorts {
		portsList = append(portsList, port)
	}

	return portsList

}

// Hosts returns the number of unique hosts in the results set.
func (rs Results) Hosts() int {

	hosts := make(map[string][]int, 10)

	for _, i := range rs {
		hosts[i.Host] = append(hosts[i.Host], i.Port)
	}

	return len(hosts)
}

// HostsReachable returns the number of hosts with at least one open port.
func (rs Results) HostsReachable() int {

	hosts := make(map[string][]int, 10)

	for _, i := range rs {
		if i.Open {
			hosts[i.Host] = append(hosts[i.Host], i.Port)
		}
	}

	return len(hosts)
}

// Ports returns the number of unique ports in the results set.
func (rs Results) Ports() int {

	ports := make(map[int]int, 10)

	for _, i := range rs {
		ports[i.Port] = i.Port
	}

	return len(ports)
}

// PortsScanned returns the total port scan attempts in the results set.
func (rs Results) PortsScanned() int {
	return len(rs)
}

// PortsReachable returns the number of open ports found.
func (rs Results) PortsReachable() int {

	var r int

	for _, i := range rs {
		if i.Open {
			r++
		}
	}

	return r
}
