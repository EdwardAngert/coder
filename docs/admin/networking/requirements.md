# Network Requirements

This page lists the ports and protocols each Coder component needs, so you
can configure firewalls, security groups, and network policies before you
deploy. For a conceptual overview of how connections are established, see
[Networking](./index.md).

> [!NOTE]
> The only rows below that reach the public internet by default are the
> STUN rows (Google's public STUN servers) and coderd's telemetry, update
> check, and Terraform provider downloads. All of them can be disabled.
> Every other row is either inbound to coderd or entirely internal to your
> deployment. See [Air-gapped Deployments](../../install/airgap.md) for the
> full air-gapped configuration.

## Port and protocol summary

| Component                    | Direction | Port / protocol                            | Required?                                   |
|-------------------------------|-----------|---------------------------------------------|----------------------------------------------|
| Coder server (coderd)          | Inbound   | `443` HTTPS/WebSocket (or your configured [`--tls-address`](../../reference/cli/server.md#--tls-address) / [`--http-address`](../../reference/cli/server.md#--http-address)) | Required, from clients, agents, and provisioners |
| Coder server (coderd)          | Outbound  | PostgreSQL (default `5432`)                 | Required                                      |
| Coder server (coderd)          | Outbound  | UDP `3478` to STUN servers                  | Optional; health checks probe the configured STUN servers, only if STUN isn't [disabled](../../reference/cli/server.md#--derp-server-stun-addresses) |
| Coder server (coderd)          | Outbound  | HTTPS to `coder.com`, `github.com`, `releases.hashicorp.com`, `registry.terraform.io` | Optional; telemetry, update checks, and Terraform downloads, all [can be disabled](../../install/airgap.md) |
| Workspace agent                | Outbound  | `443` HTTPS/WebSocket to the Coder access URL | Required                                      |
| Workspace agent                | Outbound  | UDP `3478` to STUN servers                  | Optional; only for [direct (P2P) connections](./index.md#direct-connections) |
| Workspace agent                | Outbound  | UDP, ephemeral high ports                   | Optional; only for direct connections         |
| Client (CLI, Coder Desktop, browser) | Outbound | `443` HTTPS/WebSocket to the Coder access URL | Required                                      |
| Client (CLI, Coder Desktop)    | Outbound  | UDP `3478` to STUN servers                  | Optional; only for direct connections         |
| Client (CLI, Coder Desktop)    | Outbound  | UDP, ephemeral high ports                   | Optional; only for direct connections         |
| External provisioner           | Outbound  | `443` HTTPS/WebSocket to the Coder access URL | Required; no inbound ports needed             |
| Workspace proxy                | Inbound   | `443` HTTPS/WebSocket, from clients         | Required, if the proxy is in use              |
| Workspace proxy                | Outbound  | `443` HTTPS/WebSocket to the primary Coder access URL | Required                            |

Only coderd and any workspace proxies accept inbound connections.
Workspace agents, clients, and provisioners only ever dial out.

## Coder server (coderd)

The Coder server must have an inbound address reachable by users, workspace
agents, and provisioners, set via
[`--access-url`](../../reference/cli/server.md#--access-url). Any reverse
proxy or ingress between coderd and its clients must support WebSockets.

coderd also makes several outbound connections:

- **PostgreSQL**: required for all deployments.
- **STUN servers**: used to help clients and agents discover direct
  connection paths. Enabled by default; see
  [`--derp-server-stun-addresses`](../../reference/cli/server.md#--derp-server-stun-addresses)
  to change or disable it.
- **Telemetry, update checks, and Terraform provider downloads**: all
  optional and can be turned off for [air-gapped deployments](../../install/airgap.md).

## Workspace agents and clients

Workspace agents and clients (the CLI, Coder Desktop, or a browser) only
need outbound access to the Coder access URL over HTTPS/WebSocket on `443`.
No inbound ports are required on either side.

If a direct (peer-to-peer) connection is possible, the agent and client
each send outbound UDP to the configured STUN servers on `3478` to discover
a reachable `ip:port` pair, then exchange UDP traffic with each other on
ephemeral high ports. This traffic never needs an inbound firewall rule
specifically for Coder, since it uses NAT traversal; see
[STUN and NAT](./stun.md) for how that process works. If this UDP traffic
is blocked, connections fall back to a relayed (DERP) connection over the
same `443` HTTPS/WebSocket path used for everything else.

## External provisioners

External provisioners behave like clients on the network: they only need
outbound access to the Coder access URL over `443`. No inbound ports are
required.

## Workspace proxies

A [workspace proxy](./workspace-proxies.md) needs an inbound `443` from
clients in its region, and an outbound `443` connection back to the
primary Coder server's access URL. Workspace proxies never connect to
Postgres directly.

## Up next

- Learn about [Networking](./index.md)
- Learn about [Workspace Proxies](./workspace-proxies.md)
- Troubleshoot [Networking Issues](./troubleshooting.md)
