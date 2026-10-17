# SubscriptionApiAPI

All URIs are relative to *http://localhost/api/cdc-flow*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscriptionData**](SubscriptionApiAPI.md#CreateSubscriptionData) | **Post** /v1/subscriptions | Create subscription configurations
[**DeleteSubscriptionData**](SubscriptionApiAPI.md#DeleteSubscriptionData) | **Delete** /v1/subscriptions/{id} | Delete subscription configurations
[**UpdateSubscriptionData**](SubscriptionApiAPI.md#UpdateSubscriptionData) | **Put** /v1/subscriptions/{id} | Update subscription configurations



## CreateSubscriptionData

> SubscriptionDataCreateResponse CreateSubscriptionData(ctx).Body(body).Execute()

Create subscription configurations



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
	body := *cdc.NewSubscriptionData() // SubscriptionData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.SubscriptionApiAPI.CreateSubscriptionData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.CreateSubscriptionData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubscriptionData`: SubscriptionDataCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionApiAPI.CreateSubscriptionData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SubscriptionApiAPICreateSubscriptionDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**SubscriptionData**](SubscriptionData.md) |  | 

### Return type

[**SubscriptionDataCreateResponse**](SubscriptionDataCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriptionData

> DeleteSubscriptionData(ctx, id).Execute()

Delete subscription configurations



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | Subscription Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.SubscriptionApiAPI.DeleteSubscriptionData(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.DeleteSubscriptionData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `SubscriptionApiAPIDeleteSubscriptionDataRequest` struct via the builder pattern


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


## UpdateSubscriptionData

> SubscriptionDataUpdateResponse UpdateSubscriptionData(ctx, id).Body(body).Execute()

Update subscription configurations



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | Subscription Identifier
	body := *cdc.NewSubscriptionData() // SubscriptionData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.SubscriptionApiAPI.UpdateSubscriptionData(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApiAPI.UpdateSubscriptionData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubscriptionData`: SubscriptionDataUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionApiAPI.UpdateSubscriptionData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Subscription Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `SubscriptionApiAPIUpdateSubscriptionDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**SubscriptionData**](SubscriptionData.md) |  | 

### Return type

[**SubscriptionDataUpdateResponse**](SubscriptionDataUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

