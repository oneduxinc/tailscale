# OneDux Soar daemon build

This branch (`oneduxsoar`) is upstream `tailscale/tailscale` plus a small patch so that
the OneDux Soar daemon can be installed next to an upstream Tailscale client on the same
Linux host. Our desktop app (separate repository) talks to this `tailscaled` over
LocalAPI; nothing here changes the UI, protocol or control-plane behaviour. The one
local-authorization addition is described under "Who may operate the daemon".

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

### Windows

On Windows the identifiers are different but the problem is the same, with one twist: when
`tailscaled.exe` runs as a service it ignores `--socket`, `--statedir`, `--tun` and `--port`
(the engine runs in a `/subproc` child that does not get them), so they have to be right by
default. `tsconst/windows.go` (upstream) and `tsconst/windows_onedux.go` (`onedux`) hold them:

| Identifier | Upstream | `onedux` build |
|---|---|---|
| service name / event log source | `Tailscale` | `OneDuxSoar` |
| LocalAPI pipe | `\\.\pipe\ProtectedPrefix\Administrators\Tailscale\tailscaled` | `…\Administrators\OneDuxSoar\tailscaled` |
| `%ProgramData%` / `%LocalAppData%` directory (state, logs, audit log, `syspolicy.json`, `tailscaled-env.txt`) | `Tailscale` | `OneDux Soar` |
| Wintun adapter name / type / description | `Tailscale` / `Tailscale Tunnel` | `OneDux Soar` / `OneDux Soar Tunnel` |
| Wintun adapter GUID | `{37217669-…}` | `{4ba78abe-…}` |
| default UDP port | 41641 | 41642 |
| registry (settings / policy) | `SOFTWARE\Tailscale IPN`, `SOFTWARE\Policies\Tailscale` | `SOFTWARE\OneDux Soar IPN`, `SOFTWARE\Policies\OneDux Soar` |
| legacy NRPT rule ID | `{5abe529b-…}` | `{17c94c2b-…}` |
| Windows Firewall rules, WFP display names | `Tailscale-In`, `Tailscale-Process`, `Tailscale …` | `OneDuxSoar-In`, `OneDuxSoar-Process`, `OneDuxSoar …` |

With these, `tailscaled.exe install-system-daemon` installs a second, independent service.
Tested on Windows 11 next to the official 1.104.1 client: both logged in, each restarted,
stopped, used as the one with an exit node, and ours uninstalled, without the other's
adapter, addresses, firewall rules, NRPT rule or state changing.

Not removed by `uninstall-system-daemon` (an installer has to): the `OneDuxSoar-Process`
firewall rule, the `SOFTWARE\OneDux Soar IPN` key, `%ProgramData%\OneDux Soar`.

### No log upload by default

The `onedux` build never uploads logs to log.tailscale.com (`tsconst.NoLogsNoSupportByDefault`,
read by `envknob.NoLogsNoSupport`). Upstream needs `--no-logs-no-support` for that, which the
Windows service cannot be given and which every other invocation (`tailscaled --cleanup`)
would have to repeat.

### Who may operate the daemon

Upstream lets root and one `Prefs.OperatorUser` change the node's state over LocalAPI.
That pref is stored in the current login profile: it is lost on logout and on a restart
while logged out, after which a desktop user cannot log back in from the GUI.

The `onedux` build additionally accepts members of the local group **`onedux-soar`**
(`ipn/ipnauth/operatorgroup_onedux.go`; checked per connection in
`ConnIdentity.IsReadonlyConn`). Without the tag the group name is empty and the check is
a no-op. `OperatorUser` keeps working as upstream.

Still shared, by design: while **both** daemons are Running, each one's
"drop CGNAT-sourced traffic not arriving on my interface" rule drops tailnet traffic on
the other's interface. The product rule is "one connected at a time"; this patch only
guarantees that starting, stopping, restarting or uninstalling one daemon never breaks
the other.

## Build

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -tags onedux,ts_omit_captiveportal,ts_omit_clientupdate -o tailscaled ./cmd/tailscaled
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
  go build -tags onedux,ts_omit_captiveportal,ts_omit_clientupdate -o tailscale  ./cmd/tailscale
```

`ts_omit_clientupdate` removes upstream update checks (our app ships the daemon).
`ts_omit_captiveportal` removes captive-portal detection, which otherwise probes
`controlplane.tailscale.com` / `login.tailscale.com` regardless of `--login-server`.

Run with our own identifiers and no log upload, e.g.:

```sh
tailscaled --statedir=/var/lib/oneduxsoar --socket=/run/oneduxsoar/tailscaled.sock \
  --tun=odxsoar0 --port=41642 --no-logs-no-support
```

## Tests

- CI: `.github/workflows/oneduxsoar-linux.yml` (despite the name it also builds windows/amd64; push to `oneduxsoar` = dev build; manual run with `release=N` = release `v<upstream>-oneduxsoar.N`, tagged). Upstream workflows are disabled in the repository settings.
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