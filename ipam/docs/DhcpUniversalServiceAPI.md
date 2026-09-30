# DhcpUniversalServiceAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCapabilityAssociationsCount**](DhcpUniversalServiceAPI.md#GetCapabilityAssociationsCount) | **Get** /dhcp/universal_service/{id}/associations/count | Retrieve DHCP capability associations count. Use this method to retrieve count of associated subnets/ranges for a given universal service id.
[**ListCapabilityAssociations**](DhcpUniversalServiceAPI.md#ListCapabilityAssociations) | **Get** /dhcp/universal_service/{id}/associations | Retrieve DHCP capability associations. Use this method to retrieve associated subnets/ranges with an ancestors chain for a given universal service id.



## GetCapabilityAssociationsCount

> DHCPCapabilityAssociationsCountResponse GetCapabilityAssociationsCount(ctx, id).Execute()

Retrieve DHCP capability associations count. Use this method to retrieve count of associated subnets/ranges for a given universal service id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/ipam"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpUniversalServiceAPI.GetCapabilityAssociationsCount(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpUniversalServiceAPI.GetCapabilityAssociationsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCapabilityAssociationsCount`: DHCPCapabilityAssociationsCountResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpUniversalServiceAPI.GetCapabilityAssociationsCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpUniversalServiceAPIGetCapabilityAssociationsCountRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**DHCPCapabilityAssociationsCountResponse**](DHCPCapabilityAssociationsCountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCapabilityAssociations

> DHCPCapabilityAssociationsResponse ListCapabilityAssociations(ctx, id).Execute()

Retrieve DHCP capability associations. Use this method to retrieve associated subnets/ranges with an ancestors chain for a given universal service id.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/ipam"
)

func main() {
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.DhcpUniversalServiceAPI.ListCapabilityAssociations(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpUniversalServiceAPI.ListCapabilityAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCapabilityAssociations`: DHCPCapabilityAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpUniversalServiceAPI.ListCapabilityAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpUniversalServiceAPIListCapabilityAssociationsRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**DHCPCapabilityAssociationsResponse**](DHCPCapabilityAssociationsResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

