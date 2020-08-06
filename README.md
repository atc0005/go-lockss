<!-- omit in toc -->
# go-lockss

Go-based tooling for working with LOCKSS systems

<!-- omit in toc -->
## Table of contents

- [Project home](#project-home)
- [Overview](#overview)
- [Status](#status)
- [Features](#features)
- [Known Issues](#known-issues)
- [Changelog](#changelog)
- [Requirements](#requirements)
- [How to install it](#how-to-install-it)
- [Configuration](#configuration)
  - [Command-line arguments](#command-line-arguments)
    - [Threshold calculations](#threshold-calculations)
    - [Shared](#shared)
    - [`hayoh`](#hayoh)
    - [`n2n`](#n2n)
    - [Worth noting](#worth-noting)
- [Examples](#examples)
- [References](#references)

## Project home

See [our GitHub repo](https://github.com/atc0005/go-lockss) for the latest code,
to file an issue or submit improvements for review and potential inclusion
into the project.

## Overview

This project provides a set of library packages and tools to help with
monitoring and troubleshooting LOCKSS nodes.

## Status

Alpha quality.

In the current state, this project provides a single (usable) application that
attempts to automatically obtain the list of peer nodes from a central LOCKSS
property/configuration server and check access to 9729/tcp (LCAP) to determine
whether the node is accessible for polling, voting and repair purposes.

The plan is to expose underlying libraries for use elsewhere once some
additional "bake time" has occurred.

## Features

- Single binary
  - a nice side-effect is that this makes deployment to LOCKSS nodes easier
- User configurable logging levels
- User configurable logging format
- User configurable times (config file retrieval, port connection tests)
- User configurable location of LOCKSS configuration/property settings (custom
  file or URL)

## Known Issues

- The prototype `cmd/n2n` binary is a stub application, not usable in its
  current form.
- The `--ports` flag is not currently used; the provided help text is
  misleading
- Missing documentation/examples

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet in an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

- Go 1.13+ (for building)
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

Tested using:

- Go 1.13+
- Windows 10 Version 1903
  - native
  - WSL
- Ubuntu Linux 16.04+

## How to install it

1. [Download](https://golang.org/dl/) Go
1. [Install](https://golang.org/doc/install) Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/go-lockss`
   1. `cd go-lockss`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - for current operating system
     - `go build -mod=vendor ./cmd/hayoh/`
       - *forces build to use bundled dependencies in top-level `vendor`
         folder*
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for Windows
      - `make windows`
   - for Linux
     - `make linux`
1. Copy the applicable binary to whatever systems needs to run it
   - if using `Makefile`: look in the subdirectories within
     `/tmp/go-lockss/release_assets/`
   - if using `go build`: look in `/tmp/go-lockss/`

## Configuration

### Command-line arguments

#### Threshold calculations

- Placeholder

This section is intended to cover the calculations used to determine WARNING
and CRITICAL Nagios service check threshold values. As of this writing no such
Nagios plugin exists in this repo, though it is likely that the `n2n` stub
application will be repurposed for that.

#### Shared

- Flags marked as **`required`** must be set via CLI flag
- Flags *not* marked as required are for settings where a useful default is
  already defined, *or* automatically obtained **when run on a LOCKSS node**

| Flag                        | Required | Default                  | Repeat  | Possible                                   | Description                                                                                                                                                                                                                                                    |
| --------------------------- | -------- | ------------------------ | ------- | ------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`                 | No       | `false`                  | No      | `h`, `help`                                | Show Help text along with the list of supported flags.                                                                                                                                                                                                         |
| `p`, `port`                 | No       | `9729`                   | **Yes** | *one valid TCP port per flag invocation*   | TCP port to connect to on remote LOCKSS nodes to verify connectivity. This flag may be repeated for each additional TCP port to check. The flag defaults to 9729 if not specified.                                                                             |
| `cf`, `config-file`         | No       | *empty string*           | No      | *valid filename*                           | Fully-qualified path to the fully-qualified path to an on-disk copy of the LOCKSS configuration/property XML file, usually named lockss.xml. This is NOT the same file as the /etc/lockss/config.dat file used to bootstrap the LOCKSS daemon at startup time. |
| `cs`, `config-server`       | No       | [*Maybe*](#worth-noting) | No      | *valid URL*                                | Fully-qualified URL to the LOCKSS configuration/property XML file.                                                                                                                                                                                             |
| `v`, `version`              | No       | `false`                  | No      | `v`, `version`                             | Whether to display application version and then immediately exit application.                                                                                                                                                                                  |
| `ll`, `log-level`           | No       | `info`                   | No      | `fatal`, `error`, `warn`, `info`, `debug`  | Log message priority filter. Log messages with a lower level are ignored.                                                                                                                                                                                      |
| `lf`, `log-format`          | No       | `text`                   | No      | `cli`, `json`, `logfmt`, `text`, `discard` | Use the specified `apex/log` package "handler" to output log messages in that handler's format.                                                                                                                                                                |
| `pt`, `port-timeout`        | No       | `1`                      | No      | *any positive whole number*                | Maximum number of seconds allowed for a connection test against a remote TCP port before timing out.                                                                                                                                                           |
| `ct`, `config-read-timeout` | No       | `10`                     | No      | *any positive whole number*                | Maximum number of seconds allowed for a request for the LOCKSS configuration XML file before timing out.                                                                                                                                                       |

#### `hayoh`

- placeholder

#### `n2n`

- placeholder

#### Worth noting

- When run on a LOCKSS node, and if the `config-server` setting is not
  specified, an attempt is made to automatically lookup the central
  configuration/properties server from the local `/etc/lockss/config.dat`
  file.

- Log format names map directly to the Handlers provided by the `apex/log`
  package. Their descriptions are copied from the [official
  README](https://github.com/apex/log/blob/master/Readme.md) and provided
  below for reference:

  | Log Format ("Handler") | Description                        |
  | ---------------------- | ---------------------------------- |
  | `cli`                  | human-friendly CLI output          |
  | `json`                 | provides log output in JSON format |
  | `logfmt`               | plain-text logfmt output           |
  | `text`                 | human-friendly colored output      |
  | `discard`              | discards all logs                  |

## Examples

Placeholder :(

## References

Various references used when developing this project can be found in our
[references](docs/references.md) doc.
