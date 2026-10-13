// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !onedux

package ipnauth

// operatorGroup is the name of a local group whose members may operate
// tailscaled without being root. Upstream has no such group; see
// operatorgroup_onedux.go.
const operatorGroup = ""
