# CreateSharedNetworkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**SharedNetwork**](SharedNetwork.md) | The created SharedNetwork object. | [optional] 

## Methods

### NewCreateSharedNetworkResponse

`func NewCreateSharedNetworkResponse() *CreateSharedNetworkResponse`

NewCreateSharedNetworkResponse instantiates a new CreateSharedNetworkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSharedNetworkResponseWithDefaults

`func NewCreateSharedNetworkResponseWithDefaults() *CreateSharedNetworkResponse`

NewCreateSharedNetworkResponseWithDefaults instantiates a new CreateSharedNetworkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CreateSharedNetworkResponse) GetResult() SharedNetwork`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CreateSharedNetworkResponse) GetResultOk() (*SharedNetwork, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CreateSharedNetworkResponse) SetResult(v SharedNetwork)`

SetResult sets Result field to given value.

### HasResult

`func (o *CreateSharedNetworkResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


