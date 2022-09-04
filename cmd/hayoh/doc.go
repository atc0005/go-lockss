// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/go-lockss
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// CLI application that attempts to automatically obtain the list of peer
// nodes from a central LOCKSS property/configuration server and check access
// to 9729/tcp (LCAP) to determine whether the node is accessible for polling,
// voting and repair purposes.
//
// See our [GitHub repo]:
//
//   - to review documentation (including examples)
//   - for the latest code
//   - to file an issue or submit improvements for review and potential
//     inclusion into the project
//
// [GitHub repo]: https://github.com/atc0005/go-lockss
package main
