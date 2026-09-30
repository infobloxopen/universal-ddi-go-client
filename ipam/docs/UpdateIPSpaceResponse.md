# UpdateIPSpaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | A unique ID to identify access view reassignment operation. | [optional] 
**Result** | Pointer to [**IPSpace**](IPSpace.md) | The IPSpace object. | [optional] 

## Methods

### NewUpdateIPSpaceResponse

`func NewUpdateIPSpaceResponse() *UpdateIPSpaceResponse`

NewUpdateIPSpaceResponse instantiates a new UpdateIPSpaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIPSpaceResponseWithDefaults

`func NewUpdateIPSpaceResponseWithDefaults() *UpdateIPSpaceResponse`

NewUpdateIPSpaceResponseWithDefaults instantiates a new UpdateIPSpaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *UpdateIPSpaceResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *UpdateIPSpaceResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *UpdateIPSpaceResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *UpdateIPSpaceResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetResult

`func (o *UpdateIPSpaceResponse) GetResult() IPSpace`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateIPSpaceResponse) GetResultOk() (*IPSpace, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateIPSpaceResponse) SetResult(v IPSpace)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateIPSpaceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


