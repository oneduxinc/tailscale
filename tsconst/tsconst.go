// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package tsconst exports some constants used elsewhere in the
// codebase.
package tsconst

// WintunInterfaceDesc0_14 is the description Wintun 0.14 attached to
// interfaces on Windows. WintunInterfaceDesc, the current one, is in
// windows.go with the other identifiers tailscaled claims on a Windows host.
const WintunInterfaceDesc0_14 = "Wintun Userspace Tunnel"

// TailnetLockNotTrustedMsg is the error message used by tailnet lock
// and sniffed (via substring) out of an error sent over the network.
const TailnetLockNotTrustedMsg = "this node is not trusted by Tailnet Lock"
