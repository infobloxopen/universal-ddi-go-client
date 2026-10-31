# ExportTableAPI

All URIs are relative to *http://csp.infoblox.com/api/ddi/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportTable**](ExportTableAPI.md#ExportTable) | **Post** /ipam/export/table | Initiate an asynchronous table export.



## ExportTable

> ExportTableResponse ExportTable(ctx).Body(body).Execute()

Initiate an asynchronous table export.



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
	body := *ipam.NewExportTableRequest([]string{"Fields_example"}) // ExportTableRequest | 

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.ExportTableAPI.ExportTable(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportTableAPI.ExportTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportTable`: ExportTableResponse
	fmt.Fprintf(os.Stdout, "Response from `ExportTableAPI.ExportTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ExportTableAPIExportTableRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ExportTableRequest**](ExportTableRequest.md) |  | 

### Return type

[**ExportTableResponse**](ExportTableResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

