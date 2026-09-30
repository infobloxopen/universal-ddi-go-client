# HostHAGroupCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The number of HA groups the host belongs to. | [optional] 
**Host** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewHostHAGroupCount

`func NewHostHAGroupCount() *HostHAGroupCount`

NewHostHAGroupCount instantiates a new HostHAGroupCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostHAGroupCountWithDefaults

`func NewHostHAGroupCountWithDefaults() *HostHAGroupCount`

NewHostHAGroupCountWithDefaults instantiates a new HostHAGroupCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *HostHAGroupCount) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *HostHAGroupCount) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *HostHAGroupCount) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *HostHAGroupCount) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetHost

`func (o *HostHAGroupCount) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HostHAGroupCount) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HostHAGroupCount) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *HostHAGroupCount) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


