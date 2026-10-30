# ControllerListJobConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ControllerJobConfig**](ControllerJobConfig.md) |  | [optional] 

## Methods

### NewControllerListJobConfigResponse

`func NewControllerListJobConfigResponse() *ControllerListJobConfigResponse`

NewControllerListJobConfigResponse instantiates a new ControllerListJobConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerListJobConfigResponseWithDefaults

`func NewControllerListJobConfigResponseWithDefaults() *ControllerListJobConfigResponse`

NewControllerListJobConfigResponseWithDefaults instantiates a new ControllerListJobConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ControllerListJobConfigResponse) GetResults() []ControllerJobConfig`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ControllerListJobConfigResponse) GetResultsOk() (*[]ControllerJobConfig, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ControllerListJobConfigResponse) SetResults(v []ControllerJobConfig)`

SetResults sets Results field to given value.

### HasResults

`func (o *ControllerListJobConfigResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


