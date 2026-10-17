# ScriptApiAPI

All URIs are relative to *http://localhost/api/cdc-flow*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationsProductData**](ScriptApiAPI.md#CreateApplicationsProductData) | **Post** /v1/product_applications | Create Product Script configurations
[**CreateScriptData**](ScriptApiAPI.md#CreateScriptData) | **Post** /v1/scripts | Create Application Script configurations
[**DeleteApplicationsProductData**](ScriptApiAPI.md#DeleteApplicationsProductData) | **Delete** /v1/product_applications/{product_id} | Delete a product-application mapping
[**DeleteScriptData**](ScriptApiAPI.md#DeleteScriptData) | **Delete** /v1/scripts/{file_id} | Delete scripts configurations
[**ReadApplicationsProductData**](ScriptApiAPI.md#ReadApplicationsProductData) | **Get** /v1/product_applications/{product_id} | Read script ID for product ID
[**ReadScriptData**](ScriptApiAPI.md#ReadScriptData) | **Get** /v1/scripts/{file_id} | Read scripts configurations
[**UpdateApplicationsProductData**](ScriptApiAPI.md#UpdateApplicationsProductData) | **Put** /v1/product_applications/{product_id} | Update a product-application mapping
[**UpdateScriptData**](ScriptApiAPI.md#UpdateScriptData) | **Put** /v1/scripts/{file_id} | Update scripts configurations



## CreateApplicationsProductData

> ApplicationsProductDataCreateResponse CreateApplicationsProductData(ctx).Body(body).Execute()

Create Product Script configurations



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
	body := *cdc.NewApplicationsProductData() // ApplicationsProductData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.CreateApplicationsProductData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.CreateApplicationsProductData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationsProductData`: ApplicationsProductDataCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.CreateApplicationsProductData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPICreateApplicationsProductDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ApplicationsProductData**](ApplicationsProductData.md) |  | 

### Return type

[**ApplicationsProductDataCreateResponse**](ApplicationsProductDataCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateScriptData

> ScriptDataCreateResponse CreateScriptData(ctx).Body(body).Execute()

Create Application Script configurations



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
	body := *cdc.NewScriptData() // ScriptData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.CreateScriptData(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.CreateScriptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScriptData`: ScriptDataCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.CreateScriptData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPICreateScriptDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ScriptData**](ScriptData.md) |  | 

### Return type

[**ScriptDataCreateResponse**](ScriptDataCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationsProductData

> ApplicationsProductDataDeleteResponse DeleteApplicationsProductData(ctx, productId).Execute()

Delete a product-application mapping



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
	productId := int64(789) // int64 | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.DeleteApplicationsProductData(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.DeleteApplicationsProductData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteApplicationsProductData`: ApplicationsProductDataDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.DeleteApplicationsProductData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIDeleteApplicationsProductDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationsProductDataDeleteResponse**](ApplicationsProductDataDeleteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScriptData

> ScriptDataDeleteResponse DeleteScriptData(ctx, fileId).Execute()

Delete scripts configurations



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
	fileId := int64(789) // int64 | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.DeleteScriptData(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.DeleteScriptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteScriptData`: ScriptDataDeleteResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.DeleteScriptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIDeleteScriptDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ScriptDataDeleteResponse**](ScriptDataDeleteResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadApplicationsProductData

> ApplicationsProductDataReadResponse ReadApplicationsProductData(ctx, productId).Execute()

Read script ID for product ID



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
	productId := int64(789) // int64 | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.ReadApplicationsProductData(context.Background(), productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.ReadApplicationsProductData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadApplicationsProductData`: ApplicationsProductDataReadResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.ReadApplicationsProductData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIReadApplicationsProductDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ApplicationsProductDataReadResponse**](ApplicationsProductDataReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadScriptData

> ScriptDataReadResponse ReadScriptData(ctx, fileId).Execute()

Read scripts configurations



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
	fileId := int64(789) // int64 | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.ReadScriptData(context.Background(), fileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.ReadScriptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadScriptData`: ScriptDataReadResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.ReadScriptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIReadScriptDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ScriptDataReadResponse**](ScriptDataReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationsProductData

> ApplicationsProductDataUpdateResponse UpdateApplicationsProductData(ctx, productId).Body(body).Execute()

Update a product-application mapping



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
	productId := int64(789) // int64 | 
	body := *cdc.NewApplicationsProductData() // ApplicationsProductData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.UpdateApplicationsProductData(context.Background(), productId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.UpdateApplicationsProductData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicationsProductData`: ApplicationsProductDataUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.UpdateApplicationsProductData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIUpdateApplicationsProductDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ApplicationsProductData**](ApplicationsProductData.md) |  | 

### Return type

[**ApplicationsProductDataUpdateResponse**](ApplicationsProductDataUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScriptData

> ScriptDataUpdateResponse UpdateScriptData(ctx, fileId).Body(body).Execute()

Update scripts configurations



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
	fileId := int64(789) // int64 | 
	body := *cdc.NewScriptData() // ScriptData | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ScriptApiAPI.UpdateScriptData(context.Background(), fileId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScriptApiAPI.UpdateScriptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScriptData`: ScriptDataUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ScriptApiAPI.UpdateScriptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `ScriptApiAPIUpdateScriptDataRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ScriptData**](ScriptData.md) |  | 

### Return type

[**ScriptDataUpdateResponse**](ScriptDataUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

