# GroupsAPI

All URIs are relative to *http://localhost/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUser**](GroupsAPI.md#AddUser) | **Post** /groups/{id}/users/{id} | AddUser adds a user to the user group with the specified ID.
[**Create**](GroupsAPI.md#Create) | **Post** /groups | Create creates a new user group.
[**Delete**](GroupsAPI.md#Delete) | **Delete** /groups/{id} | Delete deletes the user group with the specified ID.
[**List**](GroupsAPI.md#List) | **Get** /groups | List returns all user groups.
[**ListUsers**](GroupsAPI.md#ListUsers) | **Get** /groups/{id}/users | ListUsers returns all users in the user group with the specified ID.
[**Read**](GroupsAPI.md#Read) | **Get** /groups/{id} | Read returns the user group with the specified ID.
[**RemoveUser**](GroupsAPI.md#RemoveUser) | **Delete** /groups/{id}/users/{id} | RemoveUser removes a user from the user group with the specified ID.
[**Update**](GroupsAPI.md#Update) | **Put** /groups/{id} | Update updates the user group with the specified ID.
[**Update2**](GroupsAPI.md#Update2) | **Patch** /groups/{id} | Update updates the user group with the specified ID.



## AddUser

> GroupAddUserResponse AddUser(ctx, id, id2).Execute()

AddUser adds a user to the user group with the specified ID.

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

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.AddUser(context.Background(), id, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUser`: GroupAddUserResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 
**id2** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIAddUserRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GroupAddUserResponse**](GroupAddUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> GroupCreateResponse Create(ctx).Body(body).Execute()

Create creates a new user group.

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
	body := *identity.NewGroupCreateRequest() // GroupCreateRequest | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.Create(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: GroupCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**GroupCreateRequest**](GroupCreateRequest.md) |  | 

### Return type

[**GroupCreateResponse**](GroupCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id).CheckFeatureFlag(checkFeatureFlag).Execute()

Delete deletes the user group with the specified ID.

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
	r, err := apiClient.GroupsAPI.Delete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `GroupsAPIDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**checkFeatureFlag** | **bool** | When true, groups gated behind a feature flag are filtered by whether that flag is enabled for the owning account. Defaults to false (no FF filtering). NOTE: Only takes effect on private (SaaS) endpoints; silently ignored on public endpoints that share this request type. | 

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


## List

> GroupListResponse List(ctx).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Fields(fields).CheckFeatureFlag(checkFeatureFlag).Execute()

List returns all user groups.

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
	resp, r, err := apiClient.GroupsAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: GroupListResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**checkFeatureFlag** | **bool** | When true, groups gated behind a feature flag are filtered by whether that flag is enabled for the owning account. Defaults to false (no FF filtering). NOTE: Only takes effect on private (SaaS) endpoints; silently ignored on public endpoints that share this request type. | 

### Return type

[**GroupListResponse**](GroupListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> UserListResponse ListUsers(ctx, id).Filter(filter).OrderBy(orderBy).Offset(offset).Limit(limit).PageToken(pageToken).Fields(fields).CheckFeatureFlag(checkFeatureFlag).Execute()

ListUsers returns all users in the user group with the specified ID.

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
	id := "a5183192-1e00-475f-b334-38e1f0bb1bc7" // string | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.ListUsers(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: UserListResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIListUsersRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**orderBy** | **string** |   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.         | 
**offset** | **int32** |   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.          | 
**limit** | **int32** |   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.          | 
**pageToken** | **string** |   The service-defined string used to identify a page of resources. A null value indicates the first page.          | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 
**checkFeatureFlag** | **bool** | When true, groups gated behind a feature flag are filtered by whether that flag is enabled for the owning account. Defaults to false (no FF filtering). NOTE: Only takes effect on private (SaaS) endpoints; silently ignored on public endpoints that share this request type. | 

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GroupGetResponse Read(ctx, id).CheckFeatureFlag(checkFeatureFlag).Execute()

Read returns the user group with the specified ID.

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
	resp, r, err := apiClient.GroupsAPI.Read(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GroupGetResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**checkFeatureFlag** | **bool** | When true, groups gated behind a feature flag are filtered by whether that flag is enabled for the owning account. Defaults to false (no FF filtering). NOTE: Only takes effect on private (SaaS) endpoints; silently ignored on public endpoints that share this request type. | 

### Return type

[**GroupGetResponse**](GroupGetResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUser

> GroupRemoveUserResponse RemoveUser(ctx, id, id2).Execute()

RemoveUser removes a user from the user group with the specified ID.

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

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.RemoveUser(context.Background(), id, id2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUser`: GroupRemoveUserResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.RemoveUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 
**id2** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIRemoveUserRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**GroupRemoveUserResponse**](GroupRemoveUserResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> GroupUpdateResponse Update(ctx, id).Body(body).Execute()

Update updates the user group with the specified ID.

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
	body := *identity.NewGroup() // Group | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.Update(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: GroupUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Group**](Group.md) |  | 

### Return type

[**GroupUpdateResponse**](GroupUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update2

> GroupUpdateResponse Update2(ctx, id).Body(body).Execute()

Update updates the user group with the specified ID.

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
	body := *identity.NewGroup() // Group | 

	apiClient := identity.NewAPIClient()
	resp, r, err := apiClient.GroupsAPI.Update2(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.Update2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update2`: GroupUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.Update2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | An application specific resource identity of a resource | 

### Other Parameters

Other parameters are passed through a pointer to a `GroupsAPIUpdate2Request` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**Group**](Group.md) |  | 

### Return type

[**GroupUpdateResponse**](GroupUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

