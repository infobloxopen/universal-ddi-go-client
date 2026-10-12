# Changelog

All notable changes to `universal-ddi-go-client` are documented here.

---

## [v0.5.0] - 2026-10-12

### New Features

- **`dtc`** — Added full CRUD support for DNS Traffic Control objects: Server, Pool, Policy, Topology, and Health Checks (HTTP, ICMP, TCP, SIP, PDP, SNMP), plus SNMP User Security; schema since refreshed to match the latest upstream spec. ([#9](https://github.com/infobloxopen/universal-ddi-go-client/pull/9), [#16](https://github.com/infobloxopen/universal-ddi-go-client/pull/16))
- **`internal`, `option`, `client`** — Added a token-bucket rate limiter, enabled by default (25 requests/sec, burst 25); configure via `WithRateLimit(rps, burst)`, disable via `WithRateLimitDisabled()`, or override via `INFOBLOX_RATE_LIMIT` / `INFOBLOX_RATE_LIMIT_BURST` environment variables. ([#15](https://github.com/infobloxopen/universal-ddi-go-client/pull/15))
- **`dnsconfig`** — Added Unified Nameservers support: new `Nameserver` and `AuthNSG` models, DNSSEC signing policy/key models, and DNSSEC/external-provider metadata on `View`, while retaining backward compatibility for the old model. ([#18](https://github.com/infobloxopen/universal-ddi-go-client/pull/18))
- **`fw`** — Added `AddrBlock`/`AddrBlockApprovals` for address-block-level Network List approvals, and per-item `expiry_time`/`status` tracking on Named List items. ([#21](https://github.com/infobloxopen/universal-ddi-go-client/pull/21))
- **`anycast`** — Added a `Filter` query parameter to `GetAnycastConfigList` for filtering results by logical expression. ([#24](https://github.com/infobloxopen/universal-ddi-go-client/pull/24))

### Bug Fixes

- **`internal`** — Redacted `Authorization`, `Cookie`, and `Set-Cookie` headers in debug request/response logs. ([#17](https://github.com/infobloxopen/universal-ddi-go-client/pull/17))

### Maintenance

- Bumped the Go version in `go.mod` from 1.23 to 1.26. The minimum required Go version is now **1.26**. ([#28](https://github.com/infobloxopen/universal-ddi-go-client/pull/28))

---

## [v0.4.0] - 2026-07-15

### Breaking Changes

- **All packages** — Renamed all SDK identifiers, package names, module paths, client constructors, API types, and documentation from **BloxOne** to **Universal DDI**. ([#1](https://github.com/infobloxopen/universal-ddi-go-client/pull/1))
- **Environment variables** — Introduced `INFOBLOX_PORTAL_KEY` and `INFOBLOX_PORTAL_URL`; `BLOXONE_API_KEY` and `BLOXONE_CSP_URL` are still supported as fallbacks but are deprecated and will be removed in a future release.

### New Features

- **`clouddiscovery`** — Added `ZoneFilter` model and extended `DiscoveryConfig` and `DNSConfig` with zone filtering fields, reflecting upstream API additions. ([#4](https://github.com/infobloxopen/universal-ddi-go-client/pull/4))
- **`dnsconfig`, `dnsdata`** — Added `compartment_id` field to `View`, `AuthZone`, `ForwardZone`, `SecondaryZone`, `Delegation`, and `ACL` in DNS Config; added `compartment_id` to `Record` in DNS Data. ([bloxone-go-client#61](https://github.com/infobloxopen/bloxone-go-client/pull/61))
- **`ipam`** — Extended `DHCPConfig` and `InheritedDHCPConfig` with additional fields for DHCP Config Profile inherited sources. ([bloxone-go-client#60](https://github.com/infobloxopen/bloxone-go-client/pull/60))

### Maintenance

- Bumped Go version in `go.mod`. The minimum required Go version is now **1.23**. ([#6](https://github.com/infobloxopen/universal-ddi-go-client/pull/6))
