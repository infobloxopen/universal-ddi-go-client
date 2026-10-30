# PollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Polling** | Pointer to [**NetworkPolling**](NetworkPolling.md) |  | [optional] 
**SwitchPortManagement** | Pointer to [**SwitchPortManagement**](SwitchPortManagement.md) |  | [optional] 

## Methods

### NewPollingSettings

`func NewPollingSettings() *PollingSettings`

NewPollingSettings instantiates a new PollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollingSettingsWithDefaults

`func NewPollingSettingsWithDefaults() *PollingSettings`

NewPollingSettingsWithDefaults instantiates a new PollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolling

`func (o *PollingSettings) GetPolling() NetworkPolling`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *PollingSettings) GetPollingOk() (*NetworkPolling, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *PollingSettings) SetPolling(v NetworkPolling)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *PollingSettings) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetSwitchPortManagement

`func (o *PollingSettings) GetSwitchPortManagement() SwitchPortManagement`

GetSwitchPortManagement returns the SwitchPortManagement field if non-nil, zero value otherwise.

### GetSwitchPortManagementOk

`func (o *PollingSettings) GetSwitchPortManagementOk() (*SwitchPortManagement, bool)`

GetSwitchPortManagementOk returns a tuple with the SwitchPortManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortManagement

`func (o *PollingSettings) SetSwitchPortManagement(v SwitchPortManagement)`

SetSwitchPortManagement sets SwitchPortManagement field to given value.

### HasSwitchPortManagement

`func (o *PollingSettings) HasSwitchPortManagement() bool`

HasSwitchPortManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


