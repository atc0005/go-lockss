/*
Go-based tooling for working with LOCKSS systems.

# PROJECT HOME

See our GitHub repo (https://github.com/atc0005/go-lockss) for the latest
code, to file an issue or submit improvements for review and potential
inclusion into the project.

# PURPOSE

Provide a set of library packages and tools to help with monitoring and
troubleshooting LOCKSS nodes.

In the current state, this project provides a single (usable) executable that
attempts to automatically obtain the list of peer nodes from a central LOCKSS
property/configuration server and check access to 9729/tcp (LCAP) to determine
whether the node is accessible for polling, voting and repair purposes.

# FEATURES

  - Single binary (which makes deployment to LOCKSS nodes easier)
  - User configurable logging levels
  - User configurable logging format
  - User configurable timeouts (config file retrieval, port connection tests)
  - User configurable location of LOCKSS configuration/property settings
    (custom file or URL)
  - User configurable *additional* TCP ports to check; the default is to scan
    the LCAP port provided in the LOCKSS configuration/property XML file

# KNOWN ISSUES

  - The prototype `cmd/n2n` binary is a stub application, not usable in its
    current form.

# USAGE

  - Use the `--help` flag for current options
*/
package main
