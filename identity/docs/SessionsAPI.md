# SessionsAPI

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionDelete**](SessionsAPI.md#SessionDelete) | **Delete** /session/users/sign_out | 
[**SessionGetCurrent**](SessionsAPI.md#SessionGetCurrent) | **Get** /session/current | 
[**SessionRenew**](SessionsAPI.md#SessionRenew) | **Post** /session/users/renew | 



## SessionDelete

> SessionDelete(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/identity"
)

func main() {

	apiClient := identity.NewAPIClient()
	r, err := apiClient.SessionsAPI.SessionDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.SessionDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a `SessionsAPISessionDeleteRequest` struct via the builder pattern


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


## SessionGetCurrent

> SessionGetCurrentResponse SessionGetCurrent(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/identity"
)

func main() {

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.SessionsAPI.SessionGetCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.SessionGetCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionGetCurrent`: SessionGetCurrentResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.SessionGetCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a `SessionsAPISessionGetCurrentRequest` struct via the builder pattern


### Return type

[**SessionGetCurrentResponse**](SessionGetCurrentResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SessionRenew

> SessionRenewResponse SessionRenew(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/universal-ddi-go-client/identity"
)

func main() {

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.SessionsAPI.SessionRenew(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionsAPI.SessionRenew``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SessionRenew`: SessionRenewResponse
	fmt.Fprintf(os.Stdout, "Response from `SessionsAPI.SessionRenew`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a `SessionsAPISessionRenewRequest` struct via the builder pattern


### Return type

[**SessionRenewResponse**](SessionRenewResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

