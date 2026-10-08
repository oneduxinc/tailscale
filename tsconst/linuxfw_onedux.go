// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build onedux

package tsconst

// OneDux Soar build: the same constants as linuxfw.go with values that do not
// overlap upstream's, so that this tailscaled and an upstream tailscaled can
// be installed on the same Linux host. See linuxfw.go for what each one means.
//
// Packet marks use the fourth byte (bits 24:31) instead of the third: the
// connmark save/restore rules are matched by content (mask included), so
// sharing the mask would make the two daemons share — and delete — one rule.
const (
	LinuxFwmarkMask    = "0xff000000"
	LinuxFwmarkMaskNum = 0xff000000

	LinuxSubnetRouteMark    = "0x4000000"
	LinuxSubnetRouteMarkNum = 0x4000000

	LinuxBypassMark    = "0x8000000"
	LinuxBypassMarkNum = 0x8000000
)

const (
	LinuxChainPrefix    = "odx-"
	LinuxRouteTable     = 152
	LinuxIPRulePrefBase = 5300
)
