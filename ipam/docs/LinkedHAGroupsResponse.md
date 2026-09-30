# LinkedHAGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostHaGroupCounts** | Pointer to [**[]HostHAGroupCount**](HostHAGroupCount.md) | The number of HA groups each host belongs to. Host is provided in the request, either implicitly via HA group reference or directly via list of host references. | [optional] 
**Results** | Pointer to [**[]LinkedHAGroup**](LinkedHAGroup.md) | The list of linked HA group objects. | [optional] 

## Methods

### NewLinkedHAGroupsResponse

`func NewLinkedHAGroupsResponse() *LinkedHAGroupsResponse`

NewLinkedHAGroupsResponse instantiates a new LinkedHAGroupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedHAGroupsResponseWithDefaults

`func NewLinkedHAGroupsResponseWithDefaults() *LinkedHAGroupsResponse`

NewLinkedHAGroupsResponseWithDefaults instantiates a new LinkedHAGroupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostHaGroupCounts

`func (o *LinkedHAGroupsResponse) GetHostHaGroupCounts() []HostHAGroupCount`

GetHostHaGroupCounts returns the HostHaGroupCounts field if non-nil, zero value otherwise.

### GetHostHaGroupCountsOk

`func (o *LinkedHAGroupsResponse) GetHostHaGroupCountsOk() (*[]HostHAGroupCount, bool)`

GetHostHaGroupCountsOk returns a tuple with the HostHaGroupCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHaGroupCounts

`func (o *LinkedHAGroupsResponse) SetHostHaGroupCounts(v []HostHAGroupCount)`

SetHostHaGroupCounts sets HostHaGroupCounts field to given value.

### HasHostHaGroupCounts

`func (o *LinkedHAGroupsResponse) HasHostHaGroupCounts() bool`

HasHostHaGroupCounts returns a boolean if a field has been set.

### GetResults

`func (o *LinkedHAGroupsResponse) GetResults() []LinkedHAGroup`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *LinkedHAGroupsResponse) GetResultsOk() (*[]LinkedHAGroup, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *LinkedHAGroupsResponse) SetResults(v []LinkedHAGroup)`

SetResults sets Results field to given value.

### HasResults

`func (o *LinkedHAGroupsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


