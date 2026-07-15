# Changelog

All notable changes to `universal-ddi-go-client` are documented here.

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
