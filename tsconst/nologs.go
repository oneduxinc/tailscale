// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !onedux

package tsconst

// NoLogsNoSupportByDefault is whether this build opts out of log uploads to
// log.tailscale.com (and of Tailscale's technical support) without being told
// to with --no-logs-no-support or TS_NO_LOGS_NO_SUPPORT.
const NoLogsNoSupportByDefault = false
