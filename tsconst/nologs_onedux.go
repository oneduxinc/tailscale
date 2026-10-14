// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build onedux

package tsconst

// NoLogsNoSupportByDefault is true in the OneDux Soar build: it never uploads
// logs to log.tailscale.com. Relying on the --no-logs-no-support flag is not
// enough: the Windows service runs without flags, and every other invocation
// of the binary (tailscaled --cleanup, for one) would have to repeat it.
const NoLogsNoSupportByDefault = true
