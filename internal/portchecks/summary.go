// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package portchecks

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/apex/log"
)

// PrintSummary generates a table of all collected port check results
func (rs Results) PrintSummary() {
	w := new(tabwriter.Writer)
	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)

	// Format in tab-separated columns
	w.Init(os.Stdout, 16, 8, 8, '\t', 0)
	// w.Init(os.Stdout, 4, 4, 4, ' ', 0)

	// Add some lead-in spacing to better separate any earlier log messages from
	// summary output
	_, _ = fmt.Fprintf(w, "\n")

	// Header row in output
	_, _ = fmt.Fprintf(w, "Peer\tPort\tOpen\tError\t\n")

	// Separator row; I'm sure this can be handled better
	_, _ = fmt.Fprintln(w, "----\t----\t----\t-----\t")

	sort.Slice(rs, func(i, _ int) bool {
		// return rs[i].Open < rs[j].Open
		return rs[i].Open

	})

	for _, item := range rs {

		// if any errors were recorded when querying DNS server show those
		// instead of attempting to show real results
		var errText string
		if item.Error != nil {
			errText = item.Error.Error()
		}
		_, _ = fmt.Fprintf(w,
			"%s\t%d\t%t\t%s\t\n",
			item.Host,
			item.Port,
			item.Open,
			errText,
		)
	}

	_, _ = fmt.Fprintln(w)

	hosts := rs.Hosts()
	hostsReachable := rs.HostsReachable()
	pHostsReachable := float32(hostsReachable) / float32(hosts) * 100

	ports := rs.Ports()
	portsScanned := rs.PortsScanned()
	portsReachable := rs.PortsReachable()
	pPortsReachable := float32(portsReachable) / float32(portsScanned) * 100

	fmt.Print("Summary:\n\n")

	fmt.Printf(
		"- %d unique ports checked on each of %d hosts.\n",
		ports,
		hosts,
	)

	fmt.Printf(
		"- %.0f%% (%d/%d) of peer nodes are reachable (at least one open port) from this system.\n",
		pHostsReachable,
		hostsReachable,
		hosts,
	)
	fmt.Printf(
		"- %.0f%% (%d/%d) of ports scanned are reachable from this system.\n",
		pPortsReachable,
		portsReachable,
		portsScanned,
	)
	fmt.Println()

	if err := w.Flush(); err != nil {
		log.Errorf("Error flushing tabwriter: %v", err.Error())
	}
}
