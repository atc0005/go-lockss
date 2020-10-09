# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/go-lockss/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.1.3] - 2020-10-09

### Added

- static binaries

### Changed

- builds
- deps

### Fixed

- YYYY-MM-DD format of changelog entries
  - previous release

- builds
- config handling

## [v0.1.2] - 2020-8-20

### Added

- Use a set of XPath expressions vs a single, hard-coded expression
  - first check to see if the preservation group defined for the node is
    in-use within the LOCKSS network configuration/properties file
  - if it is, use a XPath template that searches for peer nodes restricted to
    the set preservation group
  - if it is not, use a fixed XPath template that assumes a flat peer list is
    used
    - using a fixed XPath query is intended to help prevent false-positive
      matches for peers restricted to a specific host which might occur if the
      XPath query were more liberal
  - if all query attempts fail, bail out with an error

- Add sample file for use in future tests
  - `docs/lockss-group-based-peers-sample.xml`

- Add additional logging at points where the application can experience delays
  in operation

### Changed

- Dependencies
  - upgrade `apex/log`
    - `v1.8.0` to `v1.9.0`

- Default port connection timeout adjusted
  - `10s` in `pre-v0.1.0` release
  - `1s` in `v0.1.0` release
  - `2s` now in `v0.1.2` release
    - the intent is to better balance between slower networks and waiting too
      long for a response

- Refactored parsing of the local LOCKSS daemon configuration file
  (`/etc/lockss/config.dat`)
  - instead of cherry-picking one or two values, we now parse the entire file
    for later use, skipping any blank lines or comments
  - we also implement an internal method to wrap the process of retrieving
    needed settings in an effort to provide a more stable/reliable result

- Collapsed/simplified some debugging output

- `internal/lockss` package logging enabled if `debug` log level enabled for
  application

### Fixed

- Typo in Stringer implementation for `config.Config`
- Skip blank lines (as intended) when parsing the local LOCKSS daemon
  configuration file (`/etc/lockss/config.dat`)

## [v0.1.1] - 2020-08-16

### Fixed

- Makefile
  - Remove stray parenthesis (a typo) preventing version embedding from
    working properly

## [v0.1.0] - 2020-08-16

### Added

Planned features:

- `n2n` binary (or of another name) to serve as a Nagios plugin for regular
  monitoring of node-to-node communication (voting, polling, repairs)

Features of the initial prototype:

- `hayoh` binary
  - User configurable logging levels
  - User configurable logging format
  - User configurable timeouts (config file retrieval, port connection tests)
  - User configurable location of LOCKSS configuration/property settings
    (custom file or URL)
  - User configurable *additional* TCP ports to check
    - the default is to scan the LCAP port provided in the LOCKSS
      configuration/property XML file

Worth noting (in no particular order):

- Command-line flags support via `flag` standard library package
- Go modules (vs classic `GOPATH` setup)
- GitHub Actions Workflows which apply linting and build checks
- Makefile for general use cases (including local linting)
  - Note: See README for available options if building on Windows
- Dependabot updates monitoring
- Vendored dependencies
- README
  - Link badges to applicable GitHub Actions workflows results

[Unreleased]: https://github.com/atc0005/go-lockss/compare/v0.1.2...HEAD
[v0.1.2]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.0
