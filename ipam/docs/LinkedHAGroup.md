# LinkedHAGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to [**[]LinkedHAGroupHost**](LinkedHAGroupHost.md) | The list of hosts. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Mode** | Pointer to **string** | The mode of the HA group. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the HA group. | [optional] [readonly] 

## Methods

### NewLinkedHAGroup

`func NewLinkedHAGroup() *LinkedHAGroup`

NewLinkedHAGroup instantiates a new LinkedHAGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedHAGroupWithDefaults

`func NewLinkedHAGroupWithDefaults() *LinkedHAGroup`

NewLinkedHAGroupWithDefaults instantiates a new LinkedHAGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *LinkedHAGroup) GetHosts() []LinkedHAGroupHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *LinkedHAGroup) GetHostsOk() (*[]LinkedHAGroupHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *LinkedHAGroup) SetHosts(v []LinkedHAGroupHost)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *LinkedHAGroup) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *LinkedHAGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedHAGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedHAGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedHAGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *LinkedHAGroup) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *LinkedHAGroup) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *LinkedHAGroup) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *LinkedHAGroup) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *LinkedHAGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedHAGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedHAGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkedHAGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


