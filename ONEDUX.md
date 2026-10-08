# OneDux Soar daemon build

This branch (`oneduxsoar`) is upstream `tailscale/tailscale` plus a small patch so that
the OneDux Soar daemon can be installed next to an upstream Tailscale client on the same
Linux host. Our desktop app (separate repository) talks to this `tailscaled` over
LocalAPI; nothing here changes the UI, protocol or control-plane behaviour.

Without the `onedux` build tag this tree builds exactly upstream's `tailscaled`.

## What the patch changes

Two `tailscaled` processes on one Linux host used to delete and rewrite each other's
state, because every one of these identifiers was a fixed literal:

| Identifier | Upstream | `onedux` build |
|---|---|---|
| netfilter chain prefix (`<p>input`, `<p>forward`, `<p>postrouting`, `<p>clamp`) and nftables connmark rule labels | `ts-` | `odx-` |
| policy routing table | 52 | 152 |
| `ip rule` priority base | 5200 | 5300 |
| packet mark byte (mask / subnet-route / bypass) | `0xff0000` / `0x40000` / `0x80000` | `0xff000000` / `0x4000000` / `0x8000000` |

The values live in `tsconst/linuxfw.go` (upstream, `!onedux`) and
`tsconst/linuxfw_onedux.go` (`onedux`); `util/linuxfw`, `wgengine/router/osrouter` and
`net/netmon` read them from there instead of repeating the literals.

The mark byte has to differ as well: the connmark save/restore rules sit in the shared
`mangle` chains and are found again by content, mask included.

Still shared, by design: while **both** daemons are Running, each one's
"drop CGNAT-sourced traffic not arriving on my interface" rule drops tailnet traffic on
the other's interface. The product rule is "one connected at a time"; this patch only
guarantees that starting, stopping, restarting or uninstalling one daemon never breaks
the other.

## Build

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -tags onedux,ts_omit_captiveportal -o tailscaled ./cmd/tailscaled
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -tags onedux,ts_omit_captiveportal -o tailscale  ./cmd/tailscale
```

`ts_omit_captiveportal` removes captive-portal detection, which otherwise probes
`controlplane.tailscale.com` / `login.tailscale.com` regardless of `--login-server`.

Run with our own identifiers and no log upload, e.g.:

```sh
tailscaled --statedir=/var/lib/oneduxsoar --socket=/run/oneduxsoar/tailscaled.sock \
  --tun=odxsoar0 --port=41642 --no-logs-no-support
```

## Tests

- Upstream unit tests: run **without** the `onedux` tag (several linuxfw / osrouter tests
  assert the upstream literals such as `ts-input`, `pref 5210`, `table 52`).
- Coexistence: install upstream `tailscale` and this build on one Linux host, bring both
  up, then start / stop / restart / uninstall each in turn and check the other's chains,
  `ip rule`s and routing table are untouched. Record of the first run (2026-10-08, Ubuntu
  24.04, iptables-nft and nftables mode) is in the OneDux planning docs
  (`infra/devops/ide` §1b).

## Rebasing

Rebase `oneduxsoar` onto each upstream release tag we adopt. Conflicts, if any, are
limited to the files above; after a rebase, grep for new `"ts-` literals and new uses of
table 52 / priority 5200 in `util/linuxfw`, `wgengine/router` and `net/netmon`.