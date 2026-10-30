# ListNIServiceConfigsV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]NIDiscoveryServiceV2**](NIDiscoveryServiceV2.md) |  | [optional] 

## Methods

### NewListNIServiceConfigsV2Response

`func NewListNIServiceConfigsV2Response() *ListNIServiceConfigsV2Response`

NewListNIServiceConfigsV2Response instantiates a new ListNIServiceConfigsV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNIServiceConfigsV2ResponseWithDefaults

`func NewListNIServiceConfigsV2ResponseWithDefaults() *ListNIServiceConfigsV2Response`

NewListNIServiceConfigsV2ResponseWithDefaults instantiates a new ListNIServiceConfigsV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ListNIServiceConfigsV2Response) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListNIServiceConfigsV2Response) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListNIServiceConfigsV2Response) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *ListNIServiceConfigsV2Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *ListNIServiceConfigsV2Response) GetResults() []NIDiscoveryServiceV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ListNIServiceConfigsV2Response) GetResultsOk() (*[]NIDiscoveryServiceV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ListNIServiceConfigsV2Response) SetResults(v []NIDiscoveryServiceV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *ListNIServiceConfigsV2Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


