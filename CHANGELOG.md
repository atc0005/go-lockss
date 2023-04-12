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

## [v0.2.0] - 2023-04-12

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-198) Generate RPM/DEB packages using nFPM
- (GH-196) Add version details to Windows executables

### Changed

- (GH-197) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-195) Makefile: Compress binaries & use fixed filenames
- (GH-193) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-194) Build dev/stable releases using go-ci Docker image

### Fixed

- (GH-199) Remove prototype n2n binary (redux)

## [v0.1.17] - 2023-04-12

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-165) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
  - `antchfx/xmlquery`
    - `v1.3.13` to `v1.3.15`
  - `antchfx/xpath`
    - `v1.2.1` to `v1.2.4`
  - `pelletier/go-toml`
    - `v2.0.6` to `v2.0.7`
  - `fatih/color`
    - `v1.13.0` to `v1.15.0`
  - `go-logfmt/logfmt`
    - `v0.5.1` to `v0.6.0`
  - `mattn/go-isatty`
    - `v0.0.16` to `v0.0.18`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.7.0`
  - `golang.org/x/net`
    - `v0.4.0` to `v0.9.0`
  - `golang.org/x/text`
    - `v0.5.0` to `v0.9.0`
- CI
  - (GH-177) Drop `Push Validation` workflow
  - (GH-178) Rework workflow scheduling
  - (GH-180) Remove `Push Validation` workflow status badge

### Fixed

- (GH-186) Update vuln analysis GHAW to use on.push hook

## [v0.1.16] - 2022-12-09

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
  - `github.com/antchfx/xmlquery`
    - `v1.3.12` to `v1.3.13`
  - `github.com/mattn/go-colorable`
    - `v0.1.2` to `v0.1.13`
  - `github.com/mattn/go-isatty`
    - `v0.0.8` to `v0.0.16`
  - `golang.org/x/sys`
    - `v0.0.0-20211216021012-1d35b9e2eb4e` to `v0.3.0`
  - `golang.org/x/text`
    - `v0.3.7` to `v0.5.0`
  - `golang.org/x/net`
    - `v0.0.0-20220127200216-cd36cc0744dd` to `v0.4.0`
  - `github.com/golang/groupcache`
    - `v0.0.0-20200121045136-8c9f03a8e57e` to
      `v0.0.0-20210331224755-41bb18bfe9da`
  - `github.com/golang/gddo`
    - `v0.0.0-20200715224205-051695c33a3f` to
      `v0.0.0-20210115222349-20d68f94ee1f`
  - `github.com/fatih/color`
    - `v1.7.0` to `v1.13.0`
  - `github.com/go-logfmt/logfmt`
    - `v0.4.0` to `v0.5.1`
  - `github.com/pkg/errors`
    - `v0.8.1` to `v0.9.1`
  - `github.com/kr/logfmt`
    - `v0.0.0-20140226030751-b84e30acd515` to
      `v0.0.0-20210122060352-19f9bcb100e6`
- (GH-149) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-153) Fix Makefile Go module base path detection

## [v0.1.15] - 2022-09-22

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
  - `github/codeql-action`
    - `v2.1.22` to `v2.1.25`
- (GH-139) Update project to Go 1.19
- (GH-140) Update Makefile and GitHub Actions Workflows
- (GH-141) Remove prototype `cmd/n2n` binary
- (GH-142) Add CodeQL GitHub Action Workflow

### Fixed

- (GH-138) Add missing cmd doc files

## [v0.1.14] - 2022-08-24

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.12` to `1.17.13`
  - `antchfx/xmlquery`
    - `v1.3.11` to `v1.3.12`

### Fixed

- (GH-135) Apply Go 1.19 specific doc comments linting fixes
- (GH-136) Swap io/ioutil package for io package

## [v0.1.13] - 2022-07-21

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.12`
  - `antchfx/xmlquery`
    - `v1.3.10` to `v1.3.11`

### Fixed

- (GH-130) Update lintinstall Makefile recipe
- (GH-132) Fix Markdownlint references

## [v0.1.12] - 2022-05-06

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`
  - `antchfx/xmlquery`
    - `v1.3.9` to `v1.3.10`

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

[Unreleased]: https://github.com/atc0005/go-lockss/compare/v0.2.0...HEAD
[v0.2.0]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.0
[v0.1.17]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.17
[v0.1.16]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.16
[v0.1.15]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.15
[v0.1.14]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.14
[v0.1.13]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.13
[v0.1.12]: https://github.com/atc0005/go-lockss/releases/tag/v0.1.12
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
