// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build onedux

package tsconst

// OneDux Soar build: the same identifiers as windows.go with values that do
// not overlap upstream's, so that this tailscaled and an upstream Tailscale
// client can be installed on the same Windows host. See windows.go for what
// each one means.
const (
	WindowsServiceName = "OneDuxSoar"
	WindowsDirName     = "OneDux Soar"
	WindowsPipeDir     = "OneDuxSoar"

	WindowsTunName      = "OneDux Soar"
	WintunInterfaceDesc = WindowsTunName + " Tunnel"
	WindowsTunGUID      = "{4ba78abe-22f6-4391-960b-d9fd1a5519d6}"

	WindowsDefaultPort = 41642

	WindowsRegBase       = `SOFTWARE\OneDux Soar IPN`
	WindowsRegPolicyBase = `SOFTWARE\Policies\OneDux Soar`

	WindowsNRPTRuleID = `{17c94c2b-59fb-49ed-b42d-a72c675d1ba8}`

	WindowsFirewallRulePrefix = "OneDuxSoar"
)
