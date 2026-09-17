# Overview

This library enables you to interact with the Infoblox Universal DDI APIs using Go. The library is generated using the [OpenAPI Generator](https://openapi-generator.tech) project. 

The following Universal DDI APIs are supported:

## Infoblox Cloud
- [Infrastructure Management](infra_mgmt/README.md)
- [Infrastructure Provision (HostActivation API)](infra_provision/README.md)
- [Anycast Configuration Manager](anycast/README.md)
- [Upgrade Policy](upgradePolicy/README.md)

## Infoblox Threat Defense
- [Threat Defense Cloud (FW API)](fw/README.md)
- [DNS Forwarding Proxy (DFP API)](dfp/README.md)
- [Redirect](redirect/README.md)

## Universal DDI
- [IP Address Management](ipam/README.md)
- [DNS Configuration](dns_config/README.md)
- [DNS Data](dns_data/README.md)
- [Keys](keys/README.md)
- [DNS Traffic Control](dtc/README.md- [Keys](keys/README.md) 

# Migrating from bloxone-go-client

`universal-ddi-go-client` is the official successor to [`bloxone-go-client`](https://github.com/infobloxopen/bloxone-go-client). The API surface, authentication model, and supported endpoints are identical — the only breaking change is the Go module path and package import prefix.

### Migration steps

1. Update your `go.mod` dependency:
```bash
go get github.com/infobloxopen/universal-ddi-go-client
```

2. Replace all occurrences of `github.com/infobloxopen/bloxone-go-client` with `github.com/infobloxopen/universal-ddi-go-client` in your Go source files.

3. Run `go mod tidy` to remove the old dependency.

All existing configuration options (`option.WithAPIKey`, `option.WithCSPUrl`, `option.WithClientName`, `option.WithDefaultTags`) remain unchanged.

### Environment Variable Changes

The environment variables have been renamed as part of this migration:

| Old (`bloxone-go-client`) | New (`universal-ddi-go-client`) |
|---|---|
| `BLOXONE_API_KEY` | `INFOBLOX_PORTAL_KEY` |
| `BLOXONE_CSP_URL` | `INFOBLOX_PORTAL_URL` |

> **Note:** The `BLOXONE_*` environment variables are supported as fallbacks, but are deprecated and will be removed in a future release. Please update your environment to use the new `INFOBLOX_*` variable names.

# Installation

To install `universal-ddi-go-client` use `go get` command:

```bash
go get github.com/infobloxopen/universal-ddi-go-client
```

# Usage

You can either use an aggregated client to interact with multiple Universal DDI APIs or create a client for a specific API.

#### Aggregated Client
You can use an aggregated client to interact with multiple Universal DDI APIs. The aggregated client is available in the `client` package.

Import the package in your code:

```go
import uddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
```

To create a new API client, you can use the `NewAPIClient` function as shown below
```go
client := uddiclient.NewAPIClient()
// Now you can access the API clients using the client object, e.g.:
dnsConfigClient := client.DNSConfigurationAPI
```

#### Specific API Client
Alternatively, you can create a client for a specific API using the API package. For example, to create a client for the DNS Configuration API:

```go
import "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
client := dnsconfig.NewAPIClient()
```

# Configuration

The `NewAPIClient` function accepts a variadic list of `option.ClientOption` functions that can be used to configure the client.
It requires the `option` package to be imported. You can import the package using:
```go
import "github.com/infobloxopen/universal-ddi-go-client/option"
```

### Client Name
The client name is used to identify the client in the logs. By default, the client name is set to `universal-ddi-go-client`. You can change this using the `option.WithClientName` option. For example:
```go
client := uddiclient.NewAPIClient(option.WithClientName("my-client"))
```

### Server URL

The default URL for the Cloud Services Portal is `https://csp.infoblox.com`. If you need to change this, you can use `option.WithCSPUrl` to set the URL. For example:

```go
client := uddiclient.NewAPIClient(option.WithCSPUrl("https://csp.eu.infoblox.com"))
```

You can also set the URL using the environment variable `INFOBLOX_PORTAL_URL`

### Authorization

An API key is required to access Universal DDI API. You can obtain an API key by following the instructions in the guide for [Configuring User API Keys](https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys).

To use an API key with Universal DDI API, you can use the `option.WithAPIKey` option. For example:

```go
client := uddiclient.NewAPIClient(option.WithAPIKey("YOUR_API_KEY"))
```

You can also set the API key using the environment variable `INFOBLOX_PORTAL_KEY`

Note: The API key is a secret and should be handled securely. Hardcoding the API key in your code is not recommended.

### Default Tags

You can set default tags for all API requests using the `option.WithDefaultTags` option. For example:

```go
client := uddiclient.NewAPIClient(option.WithDefaultTags(map[string]string{"tag1": "value1", "tag2": "value2"}))
```
This will add the tags `tag1=value1` and `tag2=value2` to all API requests that support tags in the request body.

# Unified Nameservers

## Deprecation Notice: DNS Server Group (NSG) Assignment

As part of the Unified Nameservers initiative, the following DNS Server Group configurations are deprecated and will no longer be allowed. Requests that use them will be rejected:

- **Nested DNS Server Groups** — A DNS Server Group cannot be nested inside an authoritative DNS Server Group (AuthNSG).
- **Dual-role DNS Server Groups** — The same DNS Server Group cannot be assigned to both a primary zone and a secondary zone.
- **Multiple DNS Server Groups on an authoritative zone** — Only one DNS Server Group may be assigned to an authoritative zone (AuthZone).
- **DNS Server Groups combined with internal secondaries** — A DNS Server Group and internal secondaries cannot coexist on an authoritative zone.

Review your configurations for any of the above patterns and update them before they are enforced.

### Rate Limiting

The SDK includes a built-in token-bucket rate limiter that throttles outbound API requests. This prevents overwhelming the API server with too many concurrent requests, which can result in `429 Too Many Requests` or `500 Internal Server Error` responses — particularly when used with tools like Terraform that make parallel API calls.

**The rate limiter is enabled by default at 25 requests per second with a burst size of 25.** No configuration is needed for most use cases.

> **Note:** Read-only operations (`GET` requests) are excluded from rate limiting. Only mutating operations (`POST`, `PUT`, `PATCH`, `DELETE`) are throttled.

#### Environment Variables

You can tune or disable the rate limiter using environment variables:

| Variable | Type | Default | Description |
|---|---|---|---|
| `INFOBLOX_RATE_LIMIT` | float | `25` | Requests per second. Set to `0` to disable rate limiting. |
| `INFOBLOX_RATE_LIMIT_BURST` | int | Same as rate limit | Maximum number of requests that can be made instantly before throttling kicks in. |

```bash
# Override the default rate limit
export INFOBLOX_RATE_LIMIT=10
export INFOBLOX_RATE_LIMIT_BURST=5

# Disable rate limiting entirely
export INFOBLOX_RATE_LIMIT=0
```

#### Programmatic Configuration

You can configure the rate limiter using functional options:

```go
// Set a rate limit of 10 requests/second with a burst of 5
client := uddiclient.NewAPIClient(option.WithRateLimit(10, 5))

// Disable rate limiting entirely
client := uddiclient.NewAPIClient(option.WithRateLimitDisabled())
```

When using the aggregated client (`client.NewAPIClient`), the rate limiter is shared across all services, so the configured rate applies to the total request throughput across all API endpoints.

You can also provide a custom rate limiter implementation using `option.WithRateLimiter`:

```go
client := uddiclient.NewAPIClient(option.WithRateLimiter(myCustomLimiter))
```

The custom limiter must implement the `option.RateLimiter` interface:

```go
type RateLimiter interface {
    Wait(ctx context.Context) error
}
```

### Retry

The SDK automatically retries requests that fail with transient HTTP errors. This handles temporary server-side issues without requiring callers to implement their own retry logic.

**Retry is enabled by default with 3 retries, exponential backoff (1s–30s), and full jitter.** The following HTTP status codes are retried:
- `429 Too Many Requests`
- `500 Internal Server Error`
- `502 Bad Gateway`
- `503 Service Unavailable`
- `504 Gateway Timeout`

When the server returns a `Retry-After` header with a `429` response, the SDK respects the indicated wait duration (capped to the configured maximum wait).

Permanent transport errors — such as DNS `NXDOMAIN` (host not found) and invalid TLS certificates — are not retried.

#### Environment Variables

| Variable | Type | Default | Description |
|---|---|---|---|
| `INFOBLOX_MAX_RETRIES` | int | `3` | Maximum number of retry attempts. Set to `0` to disable retry. |
| `INFOBLOX_RETRY_MIN_WAIT` | duration | `1s` | Minimum backoff between retries. |
| `INFOBLOX_RETRY_MAX_WAIT` | duration | `30s` | Maximum backoff between retries. |

```bash
# Override retry settings
export INFOBLOX_MAX_RETRIES=5
export INFOBLOX_RETRY_MIN_WAIT=2s
export INFOBLOX_RETRY_MAX_WAIT=1m

# Disable retry entirely
export INFOBLOX_MAX_RETRIES=0
```

#### Programmatic Configuration

```go
// Custom retry settings
client := uddiclient.NewAPIClient(option.WithRetry(5, 2*time.Second, 1*time.Minute))

// Disable retry entirely
client := uddiclient.NewAPIClient(option.WithRetryDisabled())
```
