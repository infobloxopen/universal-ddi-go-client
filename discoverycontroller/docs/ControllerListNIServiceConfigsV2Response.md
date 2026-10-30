# ControllerListNIServiceConfigsV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]ControllerNIDiscoveryServiceV2**](ControllerNIDiscoveryServiceV2.md) |  | [optional] 

## Methods

### NewControllerListNIServiceConfigsV2Response

`func NewControllerListNIServiceConfigsV2Response() *ControllerListNIServiceConfigsV2Response`

NewControllerListNIServiceConfigsV2Response instantiates a new ControllerListNIServiceConfigsV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerListNIServiceConfigsV2ResponseWithDefaults

`func NewControllerListNIServiceConfigsV2ResponseWithDefaults() *ControllerListNIServiceConfigsV2Response`

NewControllerListNIServiceConfigsV2ResponseWithDefaults instantiates a new ControllerListNIServiceConfigsV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ControllerListNIServiceConfigsV2Response) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ControllerListNIServiceConfigsV2Response) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ControllerListNIServiceConfigsV2Response) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *ControllerListNIServiceConfigsV2Response) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *ControllerListNIServiceConfigsV2Response) GetResults() []ControllerNIDiscoveryServiceV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ControllerListNIServiceConfigsV2Response) GetResultsOk() (*[]ControllerNIDiscoveryServiceV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ControllerListNIServiceConfigsV2Response) SetResults(v []ControllerNIDiscoveryServiceV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *ControllerListNIServiceConfigsV2Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


