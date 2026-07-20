# Coder Connect

This page explains how Coder Connect, the VPN-based tunnel behind
[Coder Desktop](../../user-guides/desktop/index.md), fits into Coder's
networking model, so you can account for it in DNS, firewall, and
endpoint-security policy. For installation, end-user setup, and
troubleshooting, see [Coder Desktop](../../user-guides/desktop/index.md) and
[Coder Desktop Connect and Sync](../../user-guides/desktop/desktop-connect-sync.md).

## How it fits into the networking stack

Coder Connect doesn't introduce a separate networking system: it uses the
same [WireGuard tunnel, coordinator, STUN, and DERP fallback](./establishing-connections.md)
as the CLI and IDE extensions. The difference is where that tunnel lives.
Instead of a per-process connection scoped to a single `coder` CLI
invocation, Coder Desktop installs a system-wide VPN network extension
(via a TUN interface), and enabling Coder Connect keeps it connected to
all of your workspaces at once, so any application on your machine, an
SSH client, a browser, an IDE, can reach a workspace by hostname without
going through the `coder` CLI first.

Because it shares the same underlying tailnet connection, the same
[direct-connection](./establishing-connections.md#establishing-a-direct-connection)
and
[relayed-fallback](./establishing-connections.md#falling-back-to-a-relayed-connection)
logic applies: Coder Connect attempts a direct connection to each
workspace agent and falls back to DERP (through coderd or the nearest
[workspace proxy](./workspace-proxies.md)) when a direct path isn't
available.

The VPN extension only routes traffic destined for Coder workspaces; all
other network traffic on your machine is unaffected.

## DNS resolution

Coder Connect uses split DNS, supported natively on macOS and Windows, so
only queries for Coder-owned hostnames are resolved through the tunnel.
Every workspace agent gets a hostname under the fixed `.coder` suffix:

- `<agent>.<workspace>.me.coder`, for an agent in a workspace you own
- `<agent>.<workspace>.<username>.coder`, the fully qualified form that
  also works for workspaces shared with you
- `<workspace>.coder`, a shorthand that exists only when the workspace has
  exactly one agent

## IPv6 addressing

Each workspace agent reachable through Coder Connect is assigned an IPv6
address from a Coder-owned unique local address (ULA) range,
`fd60:627a:a42b::/48`, deterministically derived from the agent's ID. This
is distinct from the `fd7a:115c:a1e0::/48` range Tailscale's own software
uses elsewhere in the stack.

This addressing only exists inside the WireGuard tunnel; it isn't traffic
your server-side firewall or a Kubernetes `NetworkPolicy` needs to account
for. It's relevant on the client side: if endpoint security software or an
EDR agent inspects or filters traffic on the VPN network extension, allow
destinations in `fd60:627a:a42b::/48` so Coder Connect isn't blocked.

## Air-gapped deployments

Coder Connect's tunnel works the same way in an air-gapped deployment as
described in
[Establishing Connections](./establishing-connections.md#air-gapped-deployments):
it uses whatever DERP and STUN configuration your deployment has, including
a fully relayed-only setup with STUN disabled.

Installing Coder Desktop itself needs separate consideration: the
Homebrew and WinGet install paths reach public package repositories, so
for an air-gapped fleet, distribute the [manually downloaded
installer](../../user-guides/desktop/index.md#installation) instead, and
use MDM or group policy to [disable automatic
updates](../../user-guides/desktop/index.md#disable-automatic-updates)
so the app doesn't attempt to reach its update servers.

## Up next

- Learn about [Establishing Connections](./establishing-connections.md)
- Learn about [Workspace Proxies](./workspace-proxies.md)
- Set up [Coder Desktop](../../user-guides/desktop/index.md)
