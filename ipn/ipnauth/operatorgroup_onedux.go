// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build onedux

package ipnauth

// operatorGroup is the name of a local group whose members may operate
// tailscaled without being root, in addition to the single Prefs.OperatorUser.
//
// OneDux Soar build: Prefs.OperatorUser lives in the current login profile, so
// it is lost on logout (the profile is deleted) and on a restart while logged
// out (it was never persisted). A desktop user would then be unable to log
// back in from the GUI. Group membership is looked up for every connection, so
// adding or removing a member takes effect immediately.
const operatorGroup = "onedux-soar"
