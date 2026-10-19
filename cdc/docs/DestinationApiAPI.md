# DestinationApiAPI

All URIs are relative to *http://localhost/api/cdc-flow*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteDestinationSplunkCloud**](DestinationApiAPI.md#BulkDeleteDestinationSplunkCloud) | **Delete** /v1/destinations/splunkcloud | Delete a collection of Splunk Cloud destination servers
[**CreateDestinationApplication**](DestinationApiAPI.md#CreateDestinationApplication) | **Post** /v1/destinations/application | Create an application destination server
[**CreateDestinationHttp**](DestinationApiAPI.md#CreateDestinationHttp) | **Post** /v1/destinations/http | Create Http destination server
[**CreateDestinationReporting**](DestinationApiAPI.md#CreateDestinationReporting) | **Post** /v1/destinations/reporting | Create a NIOS reporting destination server
[**CreateDestinationSplunk**](DestinationApiAPI.md#CreateDestinationSplunk) | **Post** /v1/destinations/splunk | Create a Splunk destination server
[**CreateDestinationSplunkCloud**](DestinationApiAPI.md#CreateDestinationSplunkCloud) | **Post** /v1/destinations/splunkcloud | Create a Splunk Cloud destination server
[**CreateDestinationSyslog**](DestinationApiAPI.md#CreateDestinationSyslog) | **Post** /v1/destinations/syslog | Create a Syslog destination server
[**DeleteDestinationApplication**](DestinationApiAPI.md#DeleteDestinationApplication) | **Delete** /v1/destinations/application/{id} | Delete an application destination server
[**DeleteDestinationHttp**](DestinationApiAPI.md#DeleteDestinationHttp) | **Delete** /v1/destinations/http/{id} | Delete Http destination server
[**DeleteDestinationReporting**](DestinationApiAPI.md#DeleteDestinationReporting) | **Delete** /v1/destinations/reporting/{id} | Delete a NIOS reporting destination server
[**DeleteDestinationSplunk**](DestinationApiAPI.md#DeleteDestinationSplunk) | **Delete** /v1/destinations/splunk/{id} | Delete a Splunk destination server
[**DeleteDestinationSplunkCloud**](DestinationApiAPI.md#DeleteDestinationSplunkCloud) | **Delete** /v1/destinations/splunkcloud/{id} | Delete a Splunk Cloud destination server
[**DeleteDestinationSyslog**](DestinationApiAPI.md#DeleteDestinationSyslog) | **Delete** /v1/destinations/syslog/{id} | Delete a Syslog destination server
[**ListDeleteDestinationApplication**](DestinationApiAPI.md#ListDeleteDestinationApplication) | **Delete** /v1/destinations/application | Delete a collection of application destination servers
[**ListDeleteDestinationHttp**](DestinationApiAPI.md#ListDeleteDestinationHttp) | **Delete** /v1/destinations/http | Delete collection of Http destination servers
[**ListDeleteDestinationReporting**](DestinationApiAPI.md#ListDeleteDestinationReporting) | **Delete** /v1/destinations/reporting | Delete a collection of NIOS reporting destination servers
[**ListDeleteDestinationSplunk**](DestinationApiAPI.md#ListDeleteDestinationSplunk) | **Delete** /v1/destinations/splunk | Delete a collection of Splunk destination servers
[**ListDeleteDestinationSyslog**](DestinationApiAPI.md#ListDeleteDestinationSyslog) | **Delete** /v1/destinations/syslog | Delete a collection of Syslog destination servers
[**ListDestination**](DestinationApiAPI.md#ListDestination) | **Get** /v1/display/destinations | List destinations defined for a given account
[**ListDestinationApplication**](DestinationApiAPI.md#ListDestinationApplication) | **Get** /v1/destinations/application | List application destination servers
[**ListDestinationBloxone**](DestinationApiAPI.md#ListDestinationBloxone) | **Get** /v1/destinations/bloxone | List Bloxone destinations
[**ListDestinationHttp**](DestinationApiAPI.md#ListDestinationHttp) | **Get** /v1/destinations/http | List Http destination servers
[**ListDestinationReporting**](DestinationApiAPI.md#ListDestinationReporting) | **Get** /v1/destinations/reporting | List NIOS reporting destinations
[**ListDestinationSplunk**](DestinationApiAPI.md#ListDestinationSplunk) | **Get** /v1/destinations/splunk | List Splunk destination servers
[**ListDestinationSplunkCloud**](DestinationApiAPI.md#ListDestinationSplunkCloud) | **Get** /v1/destinations/splunkcloud | List Splunk Cloud destination servers
[**ListDestinationSyslog**](DestinationApiAPI.md#ListDestinationSyslog) | **Get** /v1/destinations/syslog | List Syslog destination servers
[**ReadDestinationApplication**](DestinationApiAPI.md#ReadDestinationApplication) | **Get** /v1/destinations/application/{id} | Read application destination server configurations
[**ReadDestinationBloxone**](DestinationApiAPI.md#ReadDestinationBloxone) | **Get** /v1/destinations/bloxone/{id} | Read Bloxone destination configurations
[**ReadDestinationHttp**](DestinationApiAPI.md#ReadDestinationHttp) | **Get** /v1/destinations/http/{id} | Read Http destination server configurations
[**ReadDestinationReporting**](DestinationApiAPI.md#ReadDestinationReporting) | **Get** /v1/destinations/reporting/{id} | Read NIOS reporting destination server configurations
[**ReadDestinationSplunk**](DestinationApiAPI.md#ReadDestinationSplunk) | **Get** /v1/destinations/splunk/{id} | Read Splunk destination server configurations
[**ReadDestinationSplunkCloud**](DestinationApiAPI.md#ReadDestinationSplunkCloud) | **Get** /v1/destinations/splunkcloud/{id} | Read Splunk Cloud destination server configurations
[**ReadDestinationSyslog**](DestinationApiAPI.md#ReadDestinationSyslog) | **Get** /v1/destinations/syslog/{id} | Read Syslog destination server configurations
[**UpdateDestinationApplication**](DestinationApiAPI.md#UpdateDestinationApplication) | **Put** /v1/destinations/application/{id} | Update an application destination server
[**UpdateDestinationHttp**](DestinationApiAPI.md#UpdateDestinationHttp) | **Put** /v1/destinations/http/{id} | Update Http destination server
[**UpdateDestinationReporting**](DestinationApiAPI.md#UpdateDestinationReporting) | **Put** /v1/destinations/reporting/{id} | Update a NIOS reporting destination server
[**UpdateDestinationSplunk**](DestinationApiAPI.md#UpdateDestinationSplunk) | **Put** /v1/destinations/splunk/{id} | Update a Splunk destination server
[**UpdateDestinationSplunkCloud**](DestinationApiAPI.md#UpdateDestinationSplunkCloud) | **Put** /v1/destinations/splunkcloud/{id} | Update a Splunk Cloud destination server
[**UpdateDestinationSyslog**](DestinationApiAPI.md#UpdateDestinationSyslog) | **Put** /v1/destinations/syslog/{id} | Update a Syslog destination server



## BulkDeleteDestinationSplunkCloud

> BulkDeleteDestinationSplunkCloud(ctx).Body(body).Execute()

Delete a collection of Splunk Cloud destination servers



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
	body := *cdc.NewDestinationSplunkCloudBulkDeleteRequest() // DestinationSplunkCloudBulkDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.BulkDeleteDestinationSplunkCloud(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.BulkDeleteDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIBulkDeleteDestinationSplunkCloudRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunkCloudBulkDeleteRequest**](DestinationSplunkCloudBulkDeleteRequest.md) |  | 

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


## CreateDestinationApplication

> DestinationApplicationCreateResponse CreateDestinationApplication(ctx).Body(body).Execute()

Create an application destination server



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
	body := *cdc.NewDestinationApplication() // DestinationApplication | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationApplication`: DestinationApplicationCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationApplicationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationApplication**](DestinationApplication.md) |  | 

### Return type

[**DestinationApplicationCreateResponse**](DestinationApplicationCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestinationHttp

> DestinationHttpCreateResponse CreateDestinationHttp(ctx).Body(body).Execute()

Create Http destination server



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
	body := *cdc.NewDestinationHttp() // DestinationHttp | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationHttp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationHttp`: DestinationHttpCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationHttp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationHttpRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationHttp**](DestinationHttp.md) |  | 

### Return type

[**DestinationHttpCreateResponse**](DestinationHttpCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestinationReporting

> DestinationReportingCreateResponse CreateDestinationReporting(ctx).Body(body).Execute()

Create a NIOS reporting destination server



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
	body := *cdc.NewDestinationReporting() // DestinationReporting | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationReporting(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationReporting`: DestinationReportingCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationReporting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationReportingRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationReporting**](DestinationReporting.md) |  | 

### Return type

[**DestinationReportingCreateResponse**](DestinationReportingCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestinationSplunk

> DestinationSplunkCreateResponse CreateDestinationSplunk(ctx).Body(body).Execute()

Create a Splunk destination server



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
	body := *cdc.NewDestinationSplunk() // DestinationSplunk | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationSplunk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationSplunk`: DestinationSplunkCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationSplunk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationSplunkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunk**](DestinationSplunk.md) |  | 

### Return type

[**DestinationSplunkCreateResponse**](DestinationSplunkCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestinationSplunkCloud

> DestinationSplunkCloudCreateResponse CreateDestinationSplunkCloud(ctx).Body(body).Execute()

Create a Splunk Cloud destination server



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
	body := *cdc.NewDestinationSplunkCloud() // DestinationSplunkCloud | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationSplunkCloud(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationSplunkCloud`: DestinationSplunkCloudCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationSplunkCloud`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationSplunkCloudRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunkCloud**](DestinationSplunkCloud.md) |  | 

### Return type

[**DestinationSplunkCloudCreateResponse**](DestinationSplunkCloudCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestinationSyslog

> DestinationSyslogCreateResponse CreateDestinationSyslog(ctx).Body(body).Execute()

Create a Syslog destination server



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
	body := *cdc.NewDestinationSyslog() // DestinationSyslog | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.CreateDestinationSyslog(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.CreateDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDestinationSyslog`: DestinationSyslogCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.CreateDestinationSyslog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPICreateDestinationSyslogRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSyslog**](DestinationSyslog.md) |  | 

### Return type

[**DestinationSyslogCreateResponse**](DestinationSyslogCreateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestinationApplication

> DeleteDestinationApplication(ctx, id).Execute()

Delete an application destination server



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationApplicationRequest` struct via the builder pattern


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


## DeleteDestinationHttp

> DeleteDestinationHttp(ctx, id).Execute()

Delete Http destination server



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationHttp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationHttpRequest` struct via the builder pattern


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


## DeleteDestinationReporting

> DeleteDestinationReporting(ctx, id).Execute()

Delete a NIOS reporting destination server



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationReporting(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationReportingRequest` struct via the builder pattern


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


## DeleteDestinationSplunk

> DeleteDestinationSplunk(ctx, id).Execute()

Delete a Splunk destination server



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationSplunk(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationSplunkRequest` struct via the builder pattern


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


## DeleteDestinationSplunkCloud

> DeleteDestinationSplunkCloud(ctx, id).Execute()

Delete a Splunk Cloud destination server



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
	id := int64(789) // int64 | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationSplunkCloud(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationSplunkCloudRequest` struct via the builder pattern


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


## DeleteDestinationSyslog

> DeleteDestinationSyslog(ctx, id).Execute()

Delete a Syslog destination server



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.DeleteDestinationSyslog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.DeleteDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIDeleteDestinationSyslogRequest` struct via the builder pattern


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


## ListDeleteDestinationApplication

> ListDeleteDestinationApplication(ctx).Body(body).Execute()

Delete a collection of application destination servers



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
	body := *cdc.NewDestinationApplicationListDeleteRequest() // DestinationApplicationListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.ListDeleteDestinationApplication(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDeleteDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDeleteDestinationApplicationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationApplicationListDeleteRequest**](DestinationApplicationListDeleteRequest.md) |  | 

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


## ListDeleteDestinationHttp

> ListDeleteDestinationHttp(ctx).Body(body).Execute()

Delete collection of Http destination servers



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
	body := *cdc.NewDestinationHttpListDeleteRequest() // DestinationHttpListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.ListDeleteDestinationHttp(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDeleteDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDeleteDestinationHttpRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationHttpListDeleteRequest**](DestinationHttpListDeleteRequest.md) |  | 

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


## ListDeleteDestinationReporting

> ListDeleteDestinationReporting(ctx).Body(body).Execute()

Delete a collection of NIOS reporting destination servers



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
	body := *cdc.NewDestinationReportingListDeleteRequest() // DestinationReportingListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.ListDeleteDestinationReporting(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDeleteDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDeleteDestinationReportingRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationReportingListDeleteRequest**](DestinationReportingListDeleteRequest.md) |  | 

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


## ListDeleteDestinationSplunk

> ListDeleteDestinationSplunk(ctx).Body(body).Execute()

Delete a collection of Splunk destination servers



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
	body := *cdc.NewDestinationSplunkListDeleteRequest() // DestinationSplunkListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.ListDeleteDestinationSplunk(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDeleteDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDeleteDestinationSplunkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunkListDeleteRequest**](DestinationSplunkListDeleteRequest.md) |  | 

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


## ListDeleteDestinationSyslog

> ListDeleteDestinationSyslog(ctx).Body(body).Execute()

Delete a collection of Syslog destination servers



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
	body := *cdc.NewDestinationSyslogListDeleteRequest() // DestinationSyslogListDeleteRequest | 

	apiClient := cdc.NewAPIClient()
	r, err := apiClient.DestinationApiAPI.ListDeleteDestinationSyslog(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDeleteDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDeleteDestinationSyslogRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSyslogListDeleteRequest**](DestinationSyslogListDeleteRequest.md) |  | 

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


## ListDestination

> DestinationListResponse ListDestination(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List destinations defined for a given account



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestination(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestination``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestination`: DestinationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationRequest` struct via the builder pattern


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

[**DestinationListResponse**](DestinationListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationApplication

> DestinationApplicationListResponse ListDestinationApplication(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List application destination servers



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationApplication(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationApplication`: DestinationApplicationListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationApplicationRequest` struct via the builder pattern


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

[**DestinationApplicationListResponse**](DestinationApplicationListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationBloxone

> DestinationBloxoneListResponse ListDestinationBloxone(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Bloxone destinations



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationBloxone(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationBloxone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationBloxone`: DestinationBloxoneListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationBloxone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationBloxoneRequest` struct via the builder pattern


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

[**DestinationBloxoneListResponse**](DestinationBloxoneListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationHttp

> DestinationHttpListResponse ListDestinationHttp(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Http destination servers



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationHttp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationHttp`: DestinationHttpListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationHttp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationHttpRequest` struct via the builder pattern


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

[**DestinationHttpListResponse**](DestinationHttpListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationReporting

> DestinationReportingListResponse ListDestinationReporting(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List NIOS reporting destinations



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationReporting(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationReporting`: DestinationReportingListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationReporting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationReportingRequest` struct via the builder pattern


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

[**DestinationReportingListResponse**](DestinationReportingListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationSplunk

> DestinationSplunkListResponse ListDestinationSplunk(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Splunk destination servers



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationSplunk(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationSplunk`: DestinationSplunkListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationSplunk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationSplunkRequest` struct via the builder pattern


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

[**DestinationSplunkListResponse**](DestinationSplunkListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationSplunkCloud

> DestinationSplunkCloudListResponse ListDestinationSplunkCloud(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Splunk Cloud destination servers



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationSplunkCloud(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationSplunkCloud`: DestinationSplunkCloudListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationSplunkCloud`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationSplunkCloudRequest` struct via the builder pattern


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

[**DestinationSplunkCloudListResponse**](DestinationSplunkCloudListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationSyslog

> DestinationSyslogListResponse ListDestinationSyslog(ctx).Filter(filter).Fields(fields).Offset(offset).Limit(limit).PageToken(pageToken).OrderBy(orderBy).Tfilter(tfilter).TorderBy(torderBy).Execute()

List Syslog destination servers



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
	resp, r, err := apiClient.DestinationApiAPI.ListDestinationSyslog(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ListDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDestinationSyslog`: DestinationSyslogListResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ListDestinationSyslog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIListDestinationSyslogRequest` struct via the builder pattern


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

[**DestinationSyslogListResponse**](DestinationSyslogListResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationApplication

> DestinationApplicationReadResponse ReadDestinationApplication(ctx, id).Fields(fields).Execute()

Read application destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationApplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationApplication`: DestinationApplicationReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationApplicationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationApplicationReadResponse**](DestinationApplicationReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationBloxone

> DestinationBloxoneReadResponse ReadDestinationBloxone(ctx, id).Filter(filter).Fields(fields).Execute()

Read Bloxone destination configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationBloxone(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationBloxone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationBloxone`: DestinationBloxoneReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationBloxone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationBloxoneRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationBloxoneReadResponse**](DestinationBloxoneReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationHttp

> DestinationHttpReadResponse ReadDestinationHttp(ctx, id).Fields(fields).Execute()

Read Http destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationHttp(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationHttp`: DestinationHttpReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationHttp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationHttpRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationHttpReadResponse**](DestinationHttpReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationReporting

> DestinationReportingReadResponse ReadDestinationReporting(ctx, id).Filter(filter).Fields(fields).Execute()

Read NIOS reporting destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationReporting(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationReporting`: DestinationReportingReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationReporting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationReportingRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**filter** | **string** |   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |         | 
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationReportingReadResponse**](DestinationReportingReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationSplunk

> DestinationSplunkReadResponse ReadDestinationSplunk(ctx, id).Fields(fields).Execute()

Read Splunk destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationSplunk(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationSplunk`: DestinationSplunkReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationSplunkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationSplunkReadResponse**](DestinationSplunkReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationSplunkCloud

> DestinationSplunkCloudReadResponse ReadDestinationSplunkCloud(ctx, id).Fields(fields).Execute()

Read Splunk Cloud destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationSplunkCloud(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationSplunkCloud`: DestinationSplunkCloudReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationSplunkCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationSplunkCloudRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationSplunkCloudReadResponse**](DestinationSplunkCloudReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDestinationSyslog

> DestinationSyslogReadResponse ReadDestinationSyslog(ctx, id).Fields(fields).Execute()

Read Syslog destination server configurations



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
	id := int64(789) // int64 | Destination Identifier

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.ReadDestinationSyslog(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.ReadDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadDestinationSyslog`: DestinationSyslogReadResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.ReadDestinationSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIReadDestinationSyslogRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fields** | **string** |   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a â€œflatâ€ resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.         | 

### Return type

[**DestinationSyslogReadResponse**](DestinationSyslogReadResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationApplication

> DestinationApplicationUpdateResponse UpdateDestinationApplication(ctx, id).Body(body).Execute()

Update an application destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationApplication() // DestinationApplication | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationApplication(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationApplication`: DestinationApplicationUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationApplicationRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationApplication**](DestinationApplication.md) |  | 

### Return type

[**DestinationApplicationUpdateResponse**](DestinationApplicationUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationHttp

> DestinationHttpUpdateResponse UpdateDestinationHttp(ctx, id).Body(body).Execute()

Update Http destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationHttp() // DestinationHttp | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationHttp(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationHttp`: DestinationHttpUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationHttp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationHttpRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationHttp**](DestinationHttp.md) |  | 

### Return type

[**DestinationHttpUpdateResponse**](DestinationHttpUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationReporting

> DestinationReportingUpdateResponse UpdateDestinationReporting(ctx, id).Body(body).Execute()

Update a NIOS reporting destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationReporting() // DestinationReporting | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationReporting(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationReporting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationReporting`: DestinationReportingUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationReporting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationReportingRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationReporting**](DestinationReporting.md) |  | 

### Return type

[**DestinationReportingUpdateResponse**](DestinationReportingUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationSplunk

> DestinationSplunkUpdateResponse UpdateDestinationSplunk(ctx, id).Body(body).Execute()

Update a Splunk destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationSplunk() // DestinationSplunk | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationSplunk(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationSplunk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationSplunk`: DestinationSplunkUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationSplunk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationSplunkRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunk**](DestinationSplunk.md) |  | 

### Return type

[**DestinationSplunkUpdateResponse**](DestinationSplunkUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationSplunkCloud

> DestinationSplunkCloudUpdateResponse UpdateDestinationSplunkCloud(ctx, id).Body(body).Execute()

Update a Splunk Cloud destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationSplunkCloud() // DestinationSplunkCloud | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationSplunkCloud(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationSplunkCloud``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationSplunkCloud`: DestinationSplunkCloudUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationSplunkCloud`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationSplunkCloudRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSplunkCloud**](DestinationSplunkCloud.md) |  | 

### Return type

[**DestinationSplunkCloudUpdateResponse**](DestinationSplunkCloudUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationSyslog

> DestinationSyslogUpdateResponse UpdateDestinationSyslog(ctx, id).Body(body).Execute()

Update a Syslog destination server



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
	id := int64(789) // int64 | Destination Identifier
	body := *cdc.NewDestinationSyslog() // DestinationSyslog | 

	apiClient := cdc.NewAPIClient()
	resp, r, err := apiClient.DestinationApiAPI.UpdateDestinationSyslog(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DestinationApiAPI.UpdateDestinationSyslog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDestinationSyslog`: DestinationSyslogUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DestinationApiAPI.UpdateDestinationSyslog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Destination Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a `DestinationApiAPIUpdateDestinationSyslogRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**body** | [**DestinationSyslog**](DestinationSyslog.md) |  | 

### Return type

[**DestinationSyslogUpdateResponse**](DestinationSyslogUpdateResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

