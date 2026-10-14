// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !onedux

package tsconst

// Identifiers that tailscaled claims on a Windows host. Two tailscaled
// installations that share any of them step on each other, so the OneDux Soar
// build (windows_onedux.go) gives every one a value of its own.
const (
	// WindowsServiceName is the name of the Windows service and of its
	// event log source.
	WindowsServiceName = "Tailscale"

	// WindowsDirName is the directory under %ProgramData% (state, logs,
	// audit log, syspolicy.json) and under %LocalAppData% (per-user
	// fallbacks and legacy files).
	WindowsDirName = "Tailscale"

	// WindowsPipeDir is the directory of the default LocalAPI named pipe
	// under paths.WindowsProtectedPipePrefix.
	WindowsPipeDir = "Tailscale"

	// WindowsTunName is the default name of the Wintun adapter and its
	// Wintun tunnel type. The adapter's description is
	// WintunInterfaceDesc, which Wintun derives from the tunnel type.
	WindowsTunName = "Tailscale"

	// WintunInterfaceDesc is the description attached to Tailscale
	// interfaces on Windows. This is set by the WinTun driver.
	WintunInterfaceDesc = WindowsTunName + " Tunnel"

	// WindowsTunGUID is the GUID requested for the Wintun adapter.
	WindowsTunGUID = "{37217669-42da-4657-a55b-0d995d328250}"

	// WindowsDefaultPort is the default UDP port for WireGuard and
	// peer-to-peer traffic on Windows.
	WindowsDefaultPort = 41641

	// WindowsRegBase and WindowsRegPolicyBase are the HKLM keys holding
	// tailscaled's own settings and the administrator's policy settings.
	WindowsRegBase       = `SOFTWARE\Tailscale IPN`
	WindowsRegPolicyBase = `SOFTWARE\Policies\Tailscale`

	// WindowsNRPTRuleID is the ID of the single (legacy) NRPT rule.
	WindowsNRPTRuleID = `{5abe529b-675b-4486-8459-25a634dacc23}`

	// WindowsFirewallRulePrefix is the prefix of the Windows Firewall rule
	// names ("<prefix>-In", "<prefix>-Process") and of the WFP provider
	// and sublayer display names.
	WindowsFirewallRulePrefix = "Tailscale"
)
