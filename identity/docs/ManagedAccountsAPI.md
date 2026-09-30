# ManagedAccountsAPI

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ManagedAccountsAPI.md#Create) | **Post** /managed_accounts | 
[**Delete**](ManagedAccountsAPI.md#Delete) | **Delete** /managed_accounts/{id} | 
[**Hierarchy**](ManagedAccountsAPI.md#Hierarchy) | **Get** /managed_accounts/tree/hierarchy | 
[**Hierarchy2**](ManagedAccountsAPI.md#Hierarchy2) | **Get** /managed_accounts/tree/hierarchy/{id} | 
[**List**](ManagedAccountsAPI.md#List) | **Get** /managed_accounts | 
[**Read**](ManagedAccountsAPI.md#Read) | **Get** /managed_accounts/{id} | 
[**Update**](ManagedAccountsAPI.md#Update) | **Put** /managed_accounts/{id} | 
[**Update2**](ManagedAccountsAPI.md#Update2) | **Patch** /managed_accounts/{id} | 
[**UpdateAddress**](ManagedAccountsAPI.md#UpdateAddress) | **Put** /managed_accounts/{id}/address/{id} | 



## Create

> ManagedAccountCreateResponse Create(ctx).Body(body).Execute()



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
	body := *identity.NewManagedAccount() // ManagedAccount | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: ManagedAccountCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ManagedAccount**](ManagedAccount.md) |  | 

### Return type

[**ManagedAccountCreateResponse**](ManagedAccountCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := identity.NewAPIClient()
	r, err := apiClient.ManagedAccountsAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIDeleteRequest` struct via the builder pattern


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


## Hierarchy

> ManagedAccountHierarchyResponse Hierarchy(ctx).Deep(deep).ExcludeDeleted(excludeDeleted).FilterType(filterType).Execute()



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
	resp, r, err := apiClient.ManagedAccountsAPI.Hierarchy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Hierarchy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Hierarchy`: ManagedAccountHierarchyResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Hierarchy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIHierarchyRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deep** | **bool** |  | 
**excludeDeleted** | **bool** |  | 
**filterType** | **string** |  | [default to &quot;ALL&quot;]

### Return type

[**ManagedAccountHierarchyResponse**](ManagedAccountHierarchyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Hierarchy2

> ManagedAccountHierarchyResponse Hierarchy2(ctx, id).Deep(deep).ExcludeDeleted(excludeDeleted).FilterType(filterType).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.Hierarchy2(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Hierarchy2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Hierarchy2`: ManagedAccountHierarchyResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Hierarchy2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIHierarchy2Request` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**deep** | **bool** |  | 
**excludeDeleted** | **bool** |  | 
**filterType** | **string** |  | [default to &quot;ALL&quot;]

### Return type

[**ManagedAccountHierarchyResponse**](ManagedAccountHierarchyResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ManagedAccountListResponse List(ctx).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Fields(fields).ExcludeDeleted(excludeDeleted).Execute()



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
	resp, r, err := apiClient.ManagedAccountsAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ManagedAccountListResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**excludeDeleted** | **bool** |  | 

### Return type

[**ManagedAccountListResponse**](ManagedAccountListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> ManagedAccountGetResponse Read(ctx, id).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.Read(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: ManagedAccountGetResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ManagedAccountGetResponse**](ManagedAccountGetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ManagedAccountUpdateResponse Update(ctx, id).Body(body).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *identity.NewManagedAccount() // ManagedAccount | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: ManagedAccountUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ManagedAccount**](ManagedAccount.md) |  | 

### Return type

[**ManagedAccountUpdateResponse**](ManagedAccountUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update2

> ManagedAccountUpdateResponse Update2(ctx, id).Body(body).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	body := *identity.NewManagedAccount() // ManagedAccount | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.Update2(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update2`: ManagedAccountUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.Update2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIUpdate2Request` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**ManagedAccount**](ManagedAccount.md) |  | 

### Return type

[**ManagedAccountUpdateResponse**](ManagedAccountUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAddress

> AccountAddressUpdateResponse UpdateAddress(ctx, id, id2).Body(body).Execute()



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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | An application specific resource identity of a resource
	id2 := "id_example" // string | An application specific resource identity of a resource
	body := *identity.NewAccountAddress() // AccountAddress | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.ManagedAccountsAPI.UpdateAddress(context.Background(), id, id2).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAccountsAPI.UpdateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAddress`: AccountAddressUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagedAccountsAPI.UpdateAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 
**id2** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `ManagedAccountsAPIUpdateAddressRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**AccountAddress**](AccountAddress.md) |  | 

### Return type

[**AccountAddressUpdateResponse**](AccountAddressUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

