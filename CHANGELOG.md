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

## [v0.1.11] - 2022-03-03

### Overview

- Dependency updates
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-111) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-112) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-114) var-declaration: should omit type string from declaration of var
  (revive)

## [v0.1.10] - 2022-01-25

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-107) Update go.mod file, canary Dockerfile to reflect current
      dependencies

## [v0.1.9] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `antchfx/xmlquery`
    - `v1.3.8` to `v1.3.9`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.8] - 2021-11-09

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `antchfx/xmlquery`
    - `v1.3.7` to `v1.3.8`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

### Fixed

- (GH-97) False positive `G307: Deferring unsafe method "Close" on type
  "*os.File" (gosec)` linting error
- (GH-96) `regexpMust: for const patterns like v3PeerRegex, use
  regexp.MustCompile (gocritic)` linting error

## [v0.1.7] - 2021-09-23

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`
  - `antchfx/xmlquery`
    - `v1.3.6` to `v1.3.7`

## [v0.1.6] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

## [v0.1.5] - 2021-07-16

### Overview

- Dependency updates
- built using Go 1.16.6
  - Statically linked
  - Linux (x86, x64)

### Changed

- Dependency updates
  - `Go`
    - `1.16.5` to `1.16.6`
  - `actions/setup-node`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

## [v0.1.4] - 2021-07-04

### Overview

- Dependency updates
- built using Go 1.16.5
  - Statically linked
  - Linux (x86, x64)

### Added

- Create "canary" Dockerfile to track stable Go releases, serve as a reminder
  to generate fresh binaries

### Changed

- Swap out GoDoc badge for pkg.go.dev badge
- Dependencies
  - `Go`
    - `1.15.2` to `1.16.5`
  - `antchfx/xmlquery`
    - `v1.3.3` to `v1.3.6`
  - `actions/setup-node`
    - `v1.3.3` to `v1.3.4`
  - `actions/checkout`
    - `v2.3.3` to `v2.3.4`
  - `actions/setup-node`
    - `v2.1.2` to `v2.2.0`

## [v0.1.3] - 2020-10-09

### Added

- Statically linked binary release
  - Built using Go 1.15.2
  - Native Go binaries (no cgo)
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

### Changed

- Dependencies
  - `antchfx/xmlquery`
    - `v1.2.4` to `v1.3.3`
  - `actions/checkout`
    - `v2.3.2` to `v2.3.3`
  - `actions/setup-node`
    - `v2.1.1` to `v2.1.2`

- Build options updated
  - Add `-trimpath` build flag
  - Explicitly disable cgo
  - Apply `osusergo` and `netgo` build tags
    - help ensure static builds that are not dependent on glibc

### Fixed

- YYYY-MM-DD format of changelog entries
  - previous release
- gocritic commentFormatting linting errors
- Add missing shorthand suffix in flags help text
- Makefile generates checksums with qualified path
- Makefile build options do not generate static binaries
- (Some) getter methods do not appear to return intended default values

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

[Unreleased]: https://github.com/atc0005/go-lockss/compare/v0.1.11...HEAD
[v0.1.11]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.11
[v0.1.10]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.10
[v0.1.9]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.9
[v0.1.8]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.8
[v0.1.7]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.0
