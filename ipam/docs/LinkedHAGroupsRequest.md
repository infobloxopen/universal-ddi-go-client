# LinkedHAGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HaGroupId** | Pointer to **string** | The resource identifier. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**OrderBy** | Pointer to **string** | atlas.api.sorting | [optional] 
**Paging** | Pointer to **string** | atlas.api.paging | [optional] 

## Methods

### NewLinkedHAGroupsRequest

`func NewLinkedHAGroupsRequest() *LinkedHAGroupsRequest`

NewLinkedHAGroupsRequest instantiates a new LinkedHAGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedHAGroupsRequestWithDefaults

`func NewLinkedHAGroupsRequestWithDefaults() *LinkedHAGroupsRequest`

NewLinkedHAGroupsRequestWithDefaults instantiates a new LinkedHAGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHaGroupId

`func (o *LinkedHAGroupsRequest) GetHaGroupId() string`

GetHaGroupId returns the HaGroupId field if non-nil, zero value otherwise.

### GetHaGroupIdOk

`func (o *LinkedHAGroupsRequest) GetHaGroupIdOk() (*string, bool)`

GetHaGroupIdOk returns a tuple with the HaGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaGroupId

`func (o *LinkedHAGroupsRequest) SetHaGroupId(v string)`

SetHaGroupId sets HaGroupId field to given value.

### HasHaGroupId

`func (o *LinkedHAGroupsRequest) HasHaGroupId() bool`

HasHaGroupId returns a boolean if a field has been set.

### GetHosts

`func (o *LinkedHAGroupsRequest) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *LinkedHAGroupsRequest) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *LinkedHAGroupsRequest) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *LinkedHAGroupsRequest) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetOrderBy

`func (o *LinkedHAGroupsRequest) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *LinkedHAGroupsRequest) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *LinkedHAGroupsRequest) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *LinkedHAGroupsRequest) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPaging

`func (o *LinkedHAGroupsRequest) GetPaging() string`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *LinkedHAGroupsRequest) GetPagingOk() (*string, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *LinkedHAGroupsRequest) SetPaging(v string)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *LinkedHAGroupsRequest) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


