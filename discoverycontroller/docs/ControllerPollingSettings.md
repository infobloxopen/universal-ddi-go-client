# ControllerPollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Polling** | Pointer to [**ControllerNetworkPolling**](ControllerNetworkPolling.md) |  | [optional] 
**SwitchPortManagement** | Pointer to [**ControllerSwitchPortManagement**](ControllerSwitchPortManagement.md) |  | [optional] 

## Methods

### NewControllerPollingSettings

`func NewControllerPollingSettings() *ControllerPollingSettings`

NewControllerPollingSettings instantiates a new ControllerPollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerPollingSettingsWithDefaults

`func NewControllerPollingSettingsWithDefaults() *ControllerPollingSettings`

NewControllerPollingSettingsWithDefaults instantiates a new ControllerPollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolling

`func (o *ControllerPollingSettings) GetPolling() ControllerNetworkPolling`

GetPolling returns the Polling field if non-nil, zero value otherwise.

### GetPollingOk

`func (o *ControllerPollingSettings) GetPollingOk() (*ControllerNetworkPolling, bool)`

GetPollingOk returns a tuple with the Polling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolling

`func (o *ControllerPollingSettings) SetPolling(v ControllerNetworkPolling)`

SetPolling sets Polling field to given value.

### HasPolling

`func (o *ControllerPollingSettings) HasPolling() bool`

HasPolling returns a boolean if a field has been set.

### GetSwitchPortManagement

`func (o *ControllerPollingSettings) GetSwitchPortManagement() ControllerSwitchPortManagement`

GetSwitchPortManagement returns the SwitchPortManagement field if non-nil, zero value otherwise.

### GetSwitchPortManagementOk

`func (o *ControllerPollingSettings) GetSwitchPortManagementOk() (*ControllerSwitchPortManagement, bool)`

GetSwitchPortManagementOk returns a tuple with the SwitchPortManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchPortManagement

`func (o *ControllerPollingSettings) SetSwitchPortManagement(v ControllerSwitchPortManagement)`

SetSwitchPortManagement sets SwitchPortManagement field to given value.

### HasSwitchPortManagement

`func (o *ControllerPollingSettings) HasSwitchPortManagement() bool`

HasSwitchPortManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


