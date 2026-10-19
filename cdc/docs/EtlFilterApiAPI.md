# EtlFilterApiAPI

All URIs are relative to *http://localhost/api/cdc-flow*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEtlFilterFqdn**](EtlFilterApiAPI.md#CreateEtlFilterFqdn) | **Post** /v1/etls/filters/fqdn | Create ETL FQDN filter
[**CreateEtlFilterHostname**](EtlFilterApiAPI.md#CreateEtlFilterHostname) | **Post** /v1/etls/filters/hostname | Create an ETL hostname filter
[**CreateEtlFilterIpNetwork**](EtlFilterApiAPI.md#CreateEtlFilterIpNetwork) | **Post** /v1/etls/filters/ip_network | Create an ETL IP network filter
[**CreateEtlFilterOphid**](EtlFilterApiAPI.md#CreateEtlFilterOphid) | **Post** /v1/etls/filters/ophid | Create an ETL ophid filter
[**CreateEtlFilterRecordType**](EtlFilterApiAPI.md#CreateEtlFilterRecordType) | **Post** /v1/etls/filters/record_type | Create an ETL DNS Record Type filter
[**CreateEtlFilterRegex**](EtlFilterApiAPI.md#CreateEtlFilterRegex) | **Post** /v1/etls/filters/regex | Create ETL regex filter
[**CreateEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#CreateEtlFilterThreatClassProperty) | **Post** /v1/etls/filters/threat_class_property | Create an ETL Threat Class Property filter
[**DeleteEtlFilterFqdn**](EtlFilterApiAPI.md#DeleteEtlFilterFqdn) | **Delete** /v1/etls/filters/fqdn/{id} | Delete an ETL FQDN filter
[**DeleteEtlFilterHostname**](EtlFilterApiAPI.md#DeleteEtlFilterHostname) | **Delete** /v1/etls/filters/hostname/{id} | Delete ETL hostname filters
[**DeleteEtlFilterIpNetwork**](EtlFilterApiAPI.md#DeleteEtlFilterIpNetwork) | **Delete** /v1/etls/filters/ip_network/{id} | Delete an ETL IP network filter
[**DeleteEtlFilterOphid**](EtlFilterApiAPI.md#DeleteEtlFilterOphid) | **Delete** /v1/etls/filters/ophid/{id} | Delete ETL ophid filters
[**DeleteEtlFilterRecordType**](EtlFilterApiAPI.md#DeleteEtlFilterRecordType) | **Delete** /v1/etls/filters/record_type/{id} | Delete an ETL DNS Record Type filter
[**DeleteEtlFilterRegex**](EtlFilterApiAPI.md#DeleteEtlFilterRegex) | **Delete** /v1/etls/filters/regex/{id} | Delete ETL regex filters
[**DeleteEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#DeleteEtlFilterThreatClassProperty) | **Delete** /v1/etls/filters/threat_class_property/{id} | Delete ETL Threat Class Property filters
[**ListDeleteEtlFilterFqdn**](EtlFilterApiAPI.md#ListDeleteEtlFilterFqdn) | **Delete** /v1/etls/filters/fqdn | Delete a collection of ETL FQDN filters
[**ListDeleteEtlFilterHostname**](EtlFilterApiAPI.md#ListDeleteEtlFilterHostname) | **Delete** /v1/etls/filters/hostname | Delete a collection of ETL hostname filters
[**ListDeleteEtlFilterIpNetwork**](EtlFilterApiAPI.md#ListDeleteEtlFilterIpNetwork) | **Delete** /v1/etls/filters/ip_network | Delete a collection of ETL IP network filters
[**ListDeleteEtlFilterOphid**](EtlFilterApiAPI.md#ListDeleteEtlFilterOphid) | **Delete** /v1/etls/filters/ophid | Delete a collection of ETL ophid filters
[**ListDeleteEtlFilterRecordType**](EtlFilterApiAPI.md#ListDeleteEtlFilterRecordType) | **Delete** /v1/etls/filters/record_type | Delete a collection of ETL DNS Record Type filters
[**ListDeleteEtlFilterRegex**](EtlFilterApiAPI.md#ListDeleteEtlFilterRegex) | **Delete** /v1/etls/filters/regex | Delete a collection of ETL regex filters
[**ListDeleteEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#ListDeleteEtlFilterThreatClassProperty) | **Delete** /v1/etls/filters/threat_class_property | Delete a collection of ETL Threat Class Property filters
[**ListEtlFilter**](EtlFilterApiAPI.md#ListEtlFilter) | **Get** /v1/display/etls/filters | List ETL filters defined for a given account
[**ListEtlFilterFqdn**](EtlFilterApiAPI.md#ListEtlFilterFqdn) | **Get** /v1/etls/filters/fqdn | List ETL FQDN filters
[**ListEtlFilterHostname**](EtlFilterApiAPI.md#ListEtlFilterHostname) | **Get** /v1/etls/filters/hostname | List ETL hostname filters
[**ListEtlFilterIngress**](EtlFilterApiAPI.md#ListEtlFilterIngress) | **Get** /v1/display/etls/ingress_filters | List ingress ETL filters
[**ListEtlFilterIpNetwork**](EtlFilterApiAPI.md#ListEtlFilterIpNetwork) | **Get** /v1/etls/filters/ip_network | List ETL IP network  filters
[**ListEtlFilterLocalZones**](EtlFilterApiAPI.md#ListEtlFilterLocalZones) | **Get** /v1/etls/filters/local_zones | List ETL Local Zone filters
[**ListEtlFilterOphid**](EtlFilterApiAPI.md#ListEtlFilterOphid) | **Get** /v1/etls/filters/ophid | List ETL ophid filters
[**ListEtlFilterRecordType**](EtlFilterApiAPI.md#ListEtlFilterRecordType) | **Get** /v1/etls/filters/record_type | List ETL DNS Record Type filters
[**ListEtlFilterRegex**](EtlFilterApiAPI.md#ListEtlFilterRegex) | **Get** /v1/etls/filters/regex | List ETL regex  filters
[**ListEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#ListEtlFilterThreatClassProperty) | **Get** /v1/etls/filters/threat_class_property | List ETL Threat Class Property filters
[**ReadEtlFilterFqdn**](EtlFilterApiAPI.md#ReadEtlFilterFqdn) | **Get** /v1/etls/filters/fqdn/{id} | Read ETL FQDN filter configurations
[**ReadEtlFilterHostname**](EtlFilterApiAPI.md#ReadEtlFilterHostname) | **Get** /v1/etls/filters/hostname/{id} | Read ETL hostname filter configurations
[**ReadEtlFilterIpNetwork**](EtlFilterApiAPI.md#ReadEtlFilterIpNetwork) | **Get** /v1/etls/filters/ip_network/{id} | Read ETL IP network filter configurations
[**ReadEtlFilterLocalZones**](EtlFilterApiAPI.md#ReadEtlFilterLocalZones) | **Get** /v1/etls/filters/local_zones/{id} | Read ETL Local Zones filter configurations
[**ReadEtlFilterOphid**](EtlFilterApiAPI.md#ReadEtlFilterOphid) | **Get** /v1/etls/filters/ophid/{id} | Read ETL ophid filter configurations
[**ReadEtlFilterRecordType**](EtlFilterApiAPI.md#ReadEtlFilterRecordType) | **Get** /v1/etls/filters/record_type/{id} | Read ETL DNS Record Type filter configurations
[**ReadEtlFilterRegex**](EtlFilterApiAPI.md#ReadEtlFilterRegex) | **Get** /v1/etls/filters/regex/{id} | Read ETL regex filter configurations
[**ReadEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#ReadEtlFilterThreatClassProperty) | **Get** /v1/etls/filters/threat_class_property/{id} | Read ETL Threat Class Property filter configurations
[**UpdateEtlFilterFqdn**](EtlFilterApiAPI.md#UpdateEtlFilterFqdn) | **Put** /v1/etls/filters/fqdn/{id} | Update an ETL FQDN filter
[**UpdateEtlFilterHostname**](EtlFilterApiAPI.md#UpdateEtlFilterHostname) | **Put** /v1/etls/filters/hostname/{id} | Update an ETL hostname filter
[**UpdateEtlFilterIpNetwork**](EtlFilterApiAPI.md#UpdateEtlFilterIpNetwork) | **Put** /v1/etls/filters/ip_network/{id} | Update an ETL IP network filter
[**UpdateEtlFilterOphid**](EtlFilterApiAPI.md#UpdateEtlFilterOphid) | **Put** /v1/etls/filters/ophid/{id} | Update an ETL ophid filter
[**UpdateEtlFilterRecordType**](EtlFilterApiAPI.md#UpdateEtlFilterRecordType) | **Put** /v1/etls/filters/record_type/{id} | Update an ETL DNS Record Type filter
[**UpdateEtlFilterRegex**](EtlFilterApiAPI.md#UpdateEtlFilterRegex) | **Put** /v1/etls/filters/regex/{id} | Update an ETL regex filter
[**UpdateEtlFilterThreatClassProperty**](EtlFilterApiAPI.md#UpdateEtlFilterThreatClassProperty) | **Put** /v1/etls/filters/threat_class_property/{id} | Update an ETL Threat Class Property filter



## CreateEtlFilterFqdn

> EtlFilterFqdnCreateResponse CreateEtlFilterFqdn(ctx).Body(body).Execute()

Create ETL FQDN filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterFqdn() // EtlFilterFqdn | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterFqdn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterFqdn`: EtlFilterFqdnCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterFqdn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterFqdn**](EtlFilterFqdn.md) |  | 

### Return type

[**EtlFilterFqdnCreateResponse**](EtlFilterFqdnCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterHostname

> EtlFilterHostnameCreateResponse CreateEtlFilterHostname(ctx).Body(body).Execute()

Create an ETL hostname filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterHostname() // EtlFilterHostname | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterHostname(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterHostname`: EtlFilterHostnameCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterHostname`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterHostname**](EtlFilterHostname.md) |  | 

### Return type

[**EtlFilterHostnameCreateResponse**](EtlFilterHostnameCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterIpNetwork

> EtlFilterIpNetworkCreateResponse CreateEtlFilterIpNetwork(ctx).Body(body).Execute()

Create an ETL IP network filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterIpNetwork() // EtlFilterIpNetwork | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterIpNetwork(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterIpNetwork`: EtlFilterIpNetworkCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterIpNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterIpNetwork**](EtlFilterIpNetwork.md) |  | 

### Return type

[**EtlFilterIpNetworkCreateResponse**](EtlFilterIpNetworkCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterOphid

> EtlFilterOphidCreateResponse CreateEtlFilterOphid(ctx).Body(body).Execute()

Create an ETL ophid filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterOphid() // EtlFilterOphid | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterOphid(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterOphid`: EtlFilterOphidCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterOphid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterOphid**](EtlFilterOphid.md) |  | 

### Return type

[**EtlFilterOphidCreateResponse**](EtlFilterOphidCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterRecordType

> EtlFilterRecordTypeCreateResponse CreateEtlFilterRecordType(ctx).Body(body).Execute()

Create an ETL DNS Record Type filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterRecordType() // EtlFilterRecordType | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterRecordType(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterRecordType`: EtlFilterRecordTypeCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterRecordType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRecordType**](EtlFilterRecordType.md) |  | 

### Return type

[**EtlFilterRecordTypeCreateResponse**](EtlFilterRecordTypeCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterRegex

> EtlFilterRegexCreateResponse CreateEtlFilterRegex(ctx).Body(body).Execute()

Create ETL regex filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterRegex() // EtlFilterRegex | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterRegex(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterRegex`: EtlFilterRegexCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterRegex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRegex**](EtlFilterRegex.md) |  | 

### Return type

[**EtlFilterRegexCreateResponse**](EtlFilterRegexCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEtlFilterThreatClassProperty

> EtlFilterThreatClassPropertyCreateResponse CreateEtlFilterThreatClassProperty(ctx).Body(body).Execute()

Create an ETL Threat Class Property filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterThreatClassProperty() // EtlFilterThreatClassProperty | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.CreateEtlFilterThreatClassProperty(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.CreateEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEtlFilterThreatClassProperty`: EtlFilterThreatClassPropertyCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.CreateEtlFilterThreatClassProperty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPICreateEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterThreatClassProperty**](EtlFilterThreatClassProperty.md) |  | 

### Return type

[**EtlFilterThreatClassPropertyCreateResponse**](EtlFilterThreatClassPropertyCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterFqdn

> DeleteEtlFilterFqdn(ctx, id).Execute()

Delete an ETL FQDN filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterFqdn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterHostname

> DeleteEtlFilterHostname(ctx, id).Execute()

Delete ETL hostname filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterHostname(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterIpNetwork

> DeleteEtlFilterIpNetwork(ctx, id).Execute()

Delete an ETL IP network filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterIpNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterOphid

> DeleteEtlFilterOphid(ctx, id).Execute()

Delete ETL ophid filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterOphid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterRecordType

> DeleteEtlFilterRecordType(ctx, id).Execute()

Delete an ETL DNS Record Type filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterRecordType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterRegex

> DeleteEtlFilterRegex(ctx, id).Execute()

Delete ETL regex filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterRegex(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEtlFilterThreatClassProperty

> DeleteEtlFilterThreatClassProperty(ctx, id).Execute()

Delete ETL Threat Class Property filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.DeleteEtlFilterThreatClassProperty(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.DeleteEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIDeleteEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterFqdn

> ListDeleteEtlFilterFqdn(ctx).Body(body).Execute()

Delete a collection of ETL FQDN filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterFqdnListDeleteRequest() // EtlFilterFqdnListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterFqdn(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterFqdnListDeleteRequest**](EtlFilterFqdnListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterHostname

> ListDeleteEtlFilterHostname(ctx).Body(body).Execute()

Delete a collection of ETL hostname filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterHostnameListDeleteRequest() // EtlFilterHostnameListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterHostname(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterHostnameListDeleteRequest**](EtlFilterHostnameListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterIpNetwork

> ListDeleteEtlFilterIpNetwork(ctx).Body(body).Execute()

Delete a collection of ETL IP network filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterIpNetworkListDeleteRequest() // EtlFilterIpNetworkListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterIpNetwork(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterIpNetworkListDeleteRequest**](EtlFilterIpNetworkListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterOphid

> ListDeleteEtlFilterOphid(ctx).Body(body).Execute()

Delete a collection of ETL ophid filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterOphidListDeleteRequest() // EtlFilterOphidListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterOphid(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterOphidListDeleteRequest**](EtlFilterOphidListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterRecordType

> ListDeleteEtlFilterRecordType(ctx).Body(body).Execute()

Delete a collection of ETL DNS Record Type filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterRecordTypeListDeleteRequest() // EtlFilterRecordTypeListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterRecordType(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRecordTypeListDeleteRequest**](EtlFilterRecordTypeListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterRegex

> ListDeleteEtlFilterRegex(ctx).Body(body).Execute()

Delete a collection of ETL regex filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterRegexListDeleteRequest() // EtlFilterRegexListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterRegex(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRegexListDeleteRequest**](EtlFilterRegexListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeleteEtlFilterThreatClassProperty

> ListDeleteEtlFilterThreatClassProperty(ctx).Body(body).Execute()

Delete a collection of ETL Threat Class Property filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	body := *cdc.NewEtlFilterThreatClassPropertyListDeleteRequest() // EtlFilterThreatClassPropertyListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.EtlFilterApiAPI.ListDeleteEtlFilterThreatClassProperty(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListDeleteEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListDeleteEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterThreatClassPropertyListDeleteRequest**](EtlFilterThreatClassPropertyListDeleteRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilter

> EtlFilterListResponse ListEtlFilter(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL filters defined for a given account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilter(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilter`: EtlFilterListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterListResponse**](EtlFilterListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterFqdn

> EtlFilterFqdnListResponse ListEtlFilterFqdn(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL FQDN filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterFqdn(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterFqdn`: EtlFilterFqdnListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterFqdn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterFqdnListResponse**](EtlFilterFqdnListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterHostname

> EtlFilterHostnameListResponse ListEtlFilterHostname(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL hostname filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterHostname(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterHostname`: EtlFilterHostnameListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterHostname`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterHostnameListResponse**](EtlFilterHostnameListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterIngress

> EtlFilterListResponse ListEtlFilterIngress(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ingress ETL filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterIngress(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterIngress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterIngress`: EtlFilterListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterIngress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterIngressRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterListResponse**](EtlFilterListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterIpNetwork

> EtlFilterIpNetworkListResponse ListEtlFilterIpNetwork(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL IP network  filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterIpNetwork(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterIpNetwork`: EtlFilterIpNetworkListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterIpNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterIpNetworkListResponse**](EtlFilterIpNetworkListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterLocalZones

> EtlFilterLocalZonesListResponse ListEtlFilterLocalZones(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL Local Zone filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterLocalZones(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterLocalZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterLocalZones`: EtlFilterLocalZonesListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterLocalZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterLocalZonesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterLocalZonesListResponse**](EtlFilterLocalZonesListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterOphid

> EtlFilterOphidListResponse ListEtlFilterOphid(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL ophid filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterOphid(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterOphid`: EtlFilterOphidListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterOphid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterOphidListResponse**](EtlFilterOphidListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterRecordType

> EtlFilterRecordTypeListResponse ListEtlFilterRecordType(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL DNS Record Type filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterRecordType(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterRecordType`: EtlFilterRecordTypeListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterRecordType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterRecordTypeListResponse**](EtlFilterRecordTypeListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterRegex

> EtlFilterRegexListResponse ListEtlFilterRegex(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL regex  filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterRegex(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterRegex`: EtlFilterRegexListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterRegex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterRegexListResponse**](EtlFilterRegexListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEtlFilterThreatClassProperty

> EtlFilterThreatClassPropertyListResponse ListEtlFilterThreatClassProperty(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List ETL Threat Class Property filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ListEtlFilterThreatClassProperty(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ListEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEtlFilterThreatClassProperty`: EtlFilterThreatClassPropertyListResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ListEtlFilterThreatClassProperty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIListEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**tfilter** | **string** | This parameter is used for filtering by tags. | 
**torderBy** | **string** | This parameter is used for sorting by tags. | 

### Return type

[**EtlFilterThreatClassPropertyListResponse**](EtlFilterThreatClassPropertyListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterFqdn

> EtlFilterFqdnReadResponse ReadEtlFilterFqdn(ctx, id).Fields(fields).Execute()

Read ETL FQDN filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterFqdn(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterFqdn`: EtlFilterFqdnReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterFqdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterFqdnReadResponse**](EtlFilterFqdnReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterHostname

> EtlFilterHostnameReadResponse ReadEtlFilterHostname(ctx, id).Fields(fields).Execute()

Read ETL hostname filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterHostname(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterHostname`: EtlFilterHostnameReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterHostnameReadResponse**](EtlFilterHostnameReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterIpNetwork

> EtlFilterIpNetworkReadResponse ReadEtlFilterIpNetwork(ctx, id).Fields(fields).Execute()

Read ETL IP network filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterIpNetwork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterIpNetwork`: EtlFilterIpNetworkReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterIpNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterIpNetworkReadResponse**](EtlFilterIpNetworkReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterLocalZones

> EtlFilterLocalZonesReadResponse ReadEtlFilterLocalZones(ctx, id).Fields(fields).Execute()

Read ETL Local Zones filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterLocalZones(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterLocalZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterLocalZones`: EtlFilterLocalZonesReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterLocalZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterLocalZonesRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterLocalZonesReadResponse**](EtlFilterLocalZonesReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterOphid

> EtlFilterOphidReadResponse ReadEtlFilterOphid(ctx, id).Fields(fields).Execute()

Read ETL ophid filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterOphid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterOphid`: EtlFilterOphidReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterOphid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterOphidReadResponse**](EtlFilterOphidReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterRecordType

> EtlFilterRecordTypeReadResponse ReadEtlFilterRecordType(ctx, id).Fields(fields).Execute()

Read ETL DNS Record Type filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterRecordType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterRecordType`: EtlFilterRecordTypeReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterRecordType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterRecordTypeReadResponse**](EtlFilterRecordTypeReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterRegex

> EtlFilterRegexReadResponse ReadEtlFilterRegex(ctx, id).Fields(fields).Execute()

Read ETL regex filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterRegex(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterRegex`: EtlFilterRegexReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterRegex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterRegexReadResponse**](EtlFilterRegexReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadEtlFilterThreatClassProperty

> EtlFilterThreatClassPropertyReadResponse ReadEtlFilterThreatClassProperty(ctx, id).Fields(fields).Execute()

Read ETL Threat Class Property filter configurations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.ReadEtlFilterThreatClassProperty(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.ReadEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadEtlFilterThreatClassProperty`: EtlFilterThreatClassPropertyReadResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.ReadEtlFilterThreatClassProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIReadEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**EtlFilterThreatClassPropertyReadResponse**](EtlFilterThreatClassPropertyReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterFqdn

> EtlFilterFqdnUpdateResponse UpdateEtlFilterFqdn(ctx, id).Body(body).Execute()

Update an ETL FQDN filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterFqdn() // EtlFilterFqdn | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterFqdn(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterFqdn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterFqdn`: EtlFilterFqdnUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterFqdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterFqdnRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterFqdn**](EtlFilterFqdn.md) |  | 

### Return type

[**EtlFilterFqdnUpdateResponse**](EtlFilterFqdnUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterHostname

> EtlFilterHostnameUpdateResponse UpdateEtlFilterHostname(ctx, id).Body(body).Execute()

Update an ETL hostname filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterHostname() // EtlFilterHostname | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterHostname(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterHostname``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterHostname`: EtlFilterHostnameUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterHostname`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterHostnameRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterHostname**](EtlFilterHostname.md) |  | 

### Return type

[**EtlFilterHostnameUpdateResponse**](EtlFilterHostnameUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterIpNetwork

> EtlFilterIpNetworkUpdateResponse UpdateEtlFilterIpNetwork(ctx, id).Body(body).Execute()

Update an ETL IP network filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterIpNetwork() // EtlFilterIpNetwork | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterIpNetwork(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterIpNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterIpNetwork`: EtlFilterIpNetworkUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterIpNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterIpNetworkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterIpNetwork**](EtlFilterIpNetwork.md) |  | 

### Return type

[**EtlFilterIpNetworkUpdateResponse**](EtlFilterIpNetworkUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterOphid

> EtlFilterOphidUpdateResponse UpdateEtlFilterOphid(ctx, id).Body(body).Execute()

Update an ETL ophid filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterOphid() // EtlFilterOphid | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterOphid(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterOphid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterOphid`: EtlFilterOphidUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterOphid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterOphidRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterOphid**](EtlFilterOphid.md) |  | 

### Return type

[**EtlFilterOphidUpdateResponse**](EtlFilterOphidUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterRecordType

> EtlFilterRecordTypeUpdateResponse UpdateEtlFilterRecordType(ctx, id).Body(body).Execute()

Update an ETL DNS Record Type filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterRecordType() // EtlFilterRecordType | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterRecordType(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterRecordType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterRecordType`: EtlFilterRecordTypeUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterRecordType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterRecordTypeRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRecordType**](EtlFilterRecordType.md) |  | 

### Return type

[**EtlFilterRecordTypeUpdateResponse**](EtlFilterRecordTypeUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterRegex

> EtlFilterRegexUpdateResponse UpdateEtlFilterRegex(ctx, id).Body(body).Execute()

Update an ETL regex filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterRegex() // EtlFilterRegex | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterRegex(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterRegex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterRegex`: EtlFilterRegexUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterRegex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterRegexRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterRegex**](EtlFilterRegex.md) |  | 

### Return type

[**EtlFilterRegexUpdateResponse**](EtlFilterRegexUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEtlFilterThreatClassProperty

> EtlFilterThreatClassPropertyUpdateResponse UpdateEtlFilterThreatClassProperty(ctx, id).Body(body).Execute()

Update an ETL Threat Class Property filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

func main() {
	id := int64(789) // int64 | ETL Filter Identifier
	body := *cdc.NewEtlFilterThreatClassProperty() // EtlFilterThreatClassProperty | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.EtlFilterApiAPI.UpdateEtlFilterThreatClassProperty(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EtlFilterApiAPI.UpdateEtlFilterThreatClassProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEtlFilterThreatClassProperty`: EtlFilterThreatClassPropertyUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `EtlFilterApiAPI.UpdateEtlFilterThreatClassProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ETL Filter Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `EtlFilterApiAPIUpdateEtlFilterThreatClassPropertyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**EtlFilterThreatClassProperty**](EtlFilterThreatClassProperty.md) |  | 

### Return type

[**EtlFilterThreatClassPropertyUpdateResponse**](EtlFilterThreatClassPropertyUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

