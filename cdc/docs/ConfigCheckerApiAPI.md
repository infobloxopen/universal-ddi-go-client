# ConfigCheckerApiAPI

All URIs are relative to *http://localhost/api/cdc-flow*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckConfig**](ConfigCheckerApiAPI.md#CheckConfig) | **Post** /v2/config/check | Runs a config check



## CheckConfig

> TypesConfigCheckResponse CheckConfig(ctx).Body(body).Execute()

Runs a config check



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
	body := *cdc.NewTypesConfigCheckRequest() // TypesConfigCheckRequest | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.ConfigCheckerApiAPI.CheckConfig(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigCheckerApiAPI.CheckConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckConfig`: TypesConfigCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigCheckerApiAPI.CheckConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ConfigCheckerApiAPICheckConfigRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**TypesConfigCheckRequest**](TypesConfigCheckRequest.md) |  | 

### Return type

[**TypesConfigCheckResponse**](TypesConfigCheckResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

