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

## [v0.2.15] - 2024-08-21

### Changed

#### Dependency Updates

- (GH-493) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.4 to go-ci-oldstable-build-v0.21.5 in /dependabot/docker/builds
- (GH-495) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.5 to go-ci-oldstable-build-v0.21.6 in /dependabot/docker/builds
- (GH-498) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.6 to go-ci-oldstable-build-v0.21.7 in /dependabot/docker/builds
- (GH-508) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.7 to go-ci-oldstable-build-v0.21.8 in /dependabot/docker/builds
- (GH-515) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.8 to go-ci-oldstable-build-v0.21.9 in /dependabot/docker/builds
- (GH-505) Go Dependency: Bump golang.org/x/net from 0.27.0 to 0.28.0
- (GH-501) Go Dependency: Bump golang.org/x/sys from 0.22.0 to 0.23.0
- (GH-511) Go Dependency: Bump golang.org/x/sys from 0.23.0 to 0.24.0
- (GH-506) Go Dependency: Bump golang.org/x/text from 0.16.0 to 0.17.0
- (GH-518) Go Runtime: Bump golang from 1.21.12 to 1.22.6 in /dependabot/docker/go
- (GH-517) Update project to Go 1.22 series

#### Other

- (GH-499) Push `REPO_VERSION` var into containers for builds

### Fixed

- (GH-520) Fix govet linting error raised by updated linter

## [v0.2.14] - 2024-07-10

### Changed

#### Dependency Updates

- (GH-467) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.7 to go-ci-oldstable-build-v0.20.8 in /dependabot/docker/builds
- (GH-475) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.8 to go-ci-oldstable-build-v0.21.2 in /dependabot/docker/builds
- (GH-479) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.2 to go-ci-oldstable-build-v0.21.3 in /dependabot/docker/builds
- (GH-481) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.21.3 to go-ci-oldstable-build-v0.21.4 in /dependabot/docker/builds
- (GH-472) Go Dependency: Bump github.com/antchfx/xmlquery from 1.4.0 to 1.4.1
- (GH-487) Go Dependency: Bump golang.org/x/net from 0.26.0 to 0.27.0
- (GH-484) Go Dependency: Bump golang.org/x/sys from 0.21.0 to 0.22.0
- (GH-478) Go Runtime: Bump golang from 1.21.11 to 1.21.12 in /dependabot/docker/go

## [v0.2.13] - 2024-06-07

### Changed

#### Dependency Updates

- (GH-448) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.4 to go-ci-oldstable-build-v0.20.5 in /dependabot/docker/builds
- (GH-450) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.5 to go-ci-oldstable-build-v0.20.6 in /dependabot/docker/builds
- (GH-463) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.6 to go-ci-oldstable-build-v0.20.7 in /dependabot/docker/builds
- (GH-445) Go Dependency: Bump github.com/fatih/color from 1.16.0 to 1.17.0
- (GH-461) Go Dependency: Bump golang.org/x/net from 0.25.0 to 0.26.0
- (GH-459) Go Dependency: Bump golang.org/x/sys from 0.20.0 to 0.21.0
- (GH-460) Go Dependency: Bump golang.org/x/text from 0.15.0 to 0.16.0
- (GH-454) Go Runtime: Bump golang from 1.21.10 to 1.21.11 in /dependabot/docker/go

### Fixed

- (GH-451) Remove inactive maligned linter
- (GH-452) Fix errcheck linting errors

## [v0.2.12] - 2024-05-13

### Changed

#### Dependency Updates

- (GH-428) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.1 to go-ci-oldstable-build-v0.20.2 in /dependabot/docker/builds
- (GH-438) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.2 to go-ci-oldstable-build-v0.20.3 in /dependabot/docker/builds
- (GH-441) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.3 to go-ci-oldstable-build-v0.20.4 in /dependabot/docker/builds
- (GH-434) Go Dependency: Bump golang.org/x/net from 0.24.0 to 0.25.0
- (GH-429) Go Dependency: Bump golang.org/x/sys from 0.19.0 to 0.20.0
- (GH-430) Go Dependency: Bump golang.org/x/text from 0.14.0 to 0.15.0
- (GH-437) Go Runtime: Bump golang from 1.21.9 to 1.21.10 in /dependabot/docker/go

## [v0.2.11] - 2024-04-11

### Changed

#### Dependency Updates

- (GH-404) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.4 to go-ci-oldstable-build-v0.16.0 in /dependabot/docker/builds
- (GH-405) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.0 to go-ci-oldstable-build-v0.16.1 in /dependabot/docker/builds
- (GH-407) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.16.1 to go-ci-oldstable-build-v0.19.0 in /dependabot/docker/builds
- (GH-410) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.19.0 to go-ci-oldstable-build-v0.20.0 in /dependabot/docker/builds
- (GH-415) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.20.0 to go-ci-oldstable-build-v0.20.1 in /dependabot/docker/builds
- (GH-423) Go Dependency: Bump github.com/antchfx/xmlquery from 1.3.18 to 1.4.0
- (GH-418) Go Dependency: Bump github.com/antchfx/xpath from 1.2.5 to 1.3.0
- (GH-419) Go Dependency: Bump golang.org/x/net from 0.22.0 to 0.24.0
- (GH-417) Go Dependency: Bump golang.org/x/sys from 0.18.0 to 0.19.0
- (GH-412) Go Runtime: Bump golang from 1.21.8 to 1.21.9 in /dependabot/docker/go

## [v0.2.10] - 2024-03-08

### Changed

#### Dependency Updates

- (GH-399) Add todo/release label to "Go Runtime" PRs
- (GH-390) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.2 to go-ci-oldstable-build-v0.15.3 in /dependabot/docker/builds
- (GH-398) Build Image: Bump atc0005/go-ci from go-ci-oldstable-build-v0.15.3 to go-ci-oldstable-build-v0.15.4 in /dependabot/docker/builds
- (GH-386) canary: bump golang from 1.21.6 to 1.21.7 in /dependabot/docker/go
- (GH-383) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.15.0 to go-ci-oldstable-build-v0.15.2 in /dependabot/docker/builds
- (GH-394) Go Dependency: Bump golang.org/x/net from 0.21.0 to 0.22.0
- (GH-393) Go Dependency: Bump golang.org/x/sys from 0.17.0 to 0.18.0
- (GH-395) Go Runtime: Bump golang from 1.21.7 to 1.21.8 in /dependabot/docker/go
- (GH-388) Update Dependabot PR prefixes (redux)
- (GH-387) Update Dependabot PR prefixes
- (GH-385) Update project to Go 1.21 series

## [v0.2.9] - 2024-02-15

### Changed

#### Dependency Updates

- (GH-370) canary: bump golang from 1.20.13 to 1.20.14 in /dependabot/docker/go
- (GH-355) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.3 to go-ci-oldstable-build-v0.14.4 in /dependabot/docker/builds
- (GH-359) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.4 to go-ci-oldstable-build-v0.14.5 in /dependabot/docker/builds
- (GH-361) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.5 to go-ci-oldstable-build-v0.14.6 in /dependabot/docker/builds
- (GH-373) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.6 to go-ci-oldstable-build-v0.14.9 in /dependabot/docker/builds
- (GH-376) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.9 to go-ci-oldstable-build-v0.15.0 in /dependabot/docker/builds
- (GH-368) go.mod: bump golang.org/x/net from 0.20.0 to 0.21.0
- (GH-369) go.mod: bump golang.org/x/sys from 0.16.0 to 0.17.0

### Fixed

- (GH-379) Fix `unused-parameter` revive linting error

## [v0.2.8] - 2024-01-19

### Changed

#### Dependency Updates

- (GH-350) canary: bump golang from 1.20.12 to 1.20.13 in /dependabot/docker/go
- (GH-352) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.2 to go-ci-oldstable-build-v0.14.3 in /dependabot/docker/builds
- (GH-343) ghaw: bump github/codeql-action from 2 to 3
- (GH-348) go.mod: bump golang.org/x/net from 0.19.0 to 0.20.0

## [v0.2.7] - 2023-12-09

### Changed

#### Dependency Updates

- (GH-337) canary: bump golang from 1.20.11 to 1.20.12 in /dependabot/docker/go
- (GH-338) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.14.1 to go-ci-oldstable-build-v0.14.2 in /dependabot/docker/builds
- (GH-332) go.mod: bump golang.org/x/net from 0.18.0 to 0.19.0
- (GH-333) go.mod: bump golang.org/x/sys from 0.14.0 to 0.15.0

## [v0.2.6] - 2023-11-16

### Changed

#### Dependency Updates

- (GH-324) canary: bump golang from 1.20.8 to 1.20.11 in /dependabot/docker/go
- (GH-327) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.12 to go-ci-oldstable-build-v0.14.1 in /dependabot/docker/builds
- (GH-307) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.9 to go-ci-oldstable-build-v0.13.12 in /dependabot/docker/builds
- (GH-291) go.mod: bump github.com/antchfx/xmlquery from 1.3.17 to 1.3.18
- (GH-313) go.mod: bump github.com/antchfx/xpath from 1.2.4 to 1.2.5
- (GH-321) go.mod: bump github.com/fatih/color from 1.15.0 to 1.16.0
- (GH-309) go.mod: bump github.com/mattn/go-isatty from 0.0.19 to 0.0.20
- (GH-303) go.mod: bump golang.org/x/net from 0.15.0 to 0.17.0
- (GH-325) go.mod: bump golang.org/x/net from 0.17.0 to 0.18.0
- (GH-296) go.mod: bump golang.org/x/sys from 0.12.0 to 0.13.0
- (GH-318) go.mod: bump golang.org/x/sys from 0.13.0 to 0.14.0
- (GH-317) go.mod: bump golang.org/x/text from 0.13.0 to 0.14.0

### Fixed

- (GH-315) Fix goconst linting errors

## [v0.2.5] - 2023-10-06

### Changed

#### Dependency Updates

- (GH-278) canary: bump golang from 1.20.7 to 1.20.8 in /dependabot/docker/go
- (GH-265) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.4 to go-ci-oldstable-build-v0.13.5 in /dependabot/docker/builds
- (GH-266) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.5 to go-ci-oldstable-build-v0.13.6 in /dependabot/docker/builds
- (GH-269) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.6 to go-ci-oldstable-build-v0.13.7 in /dependabot/docker/builds
- (GH-279) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.7 to go-ci-oldstable-build-v0.13.8 in /dependabot/docker/builds
- (GH-286) docker: bump atc0005/go-ci from go-ci-oldstable-build-v0.13.8 to go-ci-oldstable-build-v0.13.9 in /dependabot/docker/builds
- (GH-274) ghaw: bump actions/checkout from 3 to 4
- (GH-276) go.mod: bump golang.org/x/net from 0.14.0 to 0.15.0
- (GH-271) go.mod: bump golang.org/x/sys from 0.11.0 to 0.12.0
- (GH-270) go.mod: bump golang.org/x/text from 0.12.0 to 0.13.0

## [v0.2.4] - 2023-08-18

### Added

- (GH-231) Add initial automated release notes config
- (GH-233) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.4` to `go-ci-oldstable-build-v0.13.4`
  - `golang.org/x/net`
    - `v0.12.0` to `v0.14.0`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
  - `golang.org/x/text`
    - `v0.11.0` to `v0.12.0`
- (GH-235) Update Dependabot config to monitor both branches
- (GH-259) Update project to Go 1.20 series

## [v0.2.3] - 2023-07-14

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.9` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.5` to `go-ci-oldstable-build-v0.11.4`
  - `antchfx/xmlquery`
    - `v1.3.15` to `v1.3.17`
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
  - `golang.org/x/sys`
    - `v0.8.0` to `v0.10.0`
  - `golang.org/x/net`
    - `v0.10.0` to `v0.12.0`
  - `golang.org/x/text`
    - `v0.9.0` to `v0.11.0`
- (GH-216) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-212) Disable depguard linter
- (GH-220) Restore local CodeQL workflow

## [v0.2.2] - 2023-05-11

### Overview

- Dependency updates
- built using Go 1.19.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.9`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.4` to `go-ci-oldstable-build-v0.10.5`
  - `golang.org/x/net`
    - `v0.9.0` to `v0.10.0`

## [v0.2.1] - 2023-04-13

### Overview

- Bug fixes
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Fixed

- (GH-203) Wrong version number info outputted by latest release

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

[Unreleased]: https://github.com/atc0005/go-lockss/compare/v0.2.15...HEAD
[v0.2.15]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.15
[v0.2.14]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.14
[v0.2.13]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.13
[v0.2.12]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.12
[v0.2.11]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.11
[v0.2.10]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.10
[v0.2.9]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.9
[v0.2.8]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.8
[v0.2.7]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.7
[v0.2.6]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.6
[v0.2.5]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.5
[v0.2.4]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.4
[v0.2.3]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.3
[v0.2.2]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.2
[v0.2.1]: https://github.com/atc0005/go-lockss/releases/tag/v0.2.1
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
