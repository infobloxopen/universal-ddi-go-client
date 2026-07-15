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
