# ControllerSwitchPortManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**PollingInterval** | Pointer to **int32** | Polling interval in minutes. | [optional] 
**PortControlPreference** | Pointer to **string** | Port control preference corresponds to one of SNMP or CLI. | [optional] 

## Methods

### NewControllerSwitchPortManagement

`func NewControllerSwitchPortManagement() *ControllerSwitchPortManagement`

NewControllerSwitchPortManagement instantiates a new ControllerSwitchPortManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerSwitchPortManagementWithDefaults

`func NewControllerSwitchPortManagementWithDefaults() *ControllerSwitchPortManagement`

NewControllerSwitchPortManagementWithDefaults instantiates a new ControllerSwitchPortManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ControllerSwitchPortManagement) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ControllerSwitchPortManagement) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ControllerSwitchPortManagement) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ControllerSwitchPortManagement) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPollingInterval

`func (o *ControllerSwitchPortManagement) GetPollingInterval() int32`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *ControllerSwitchPortManagement) GetPollingIntervalOk() (*int32, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *ControllerSwitchPortManagement) SetPollingInterval(v int32)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *ControllerSwitchPortManagement) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPortControlPreference

`func (o *ControllerSwitchPortManagement) GetPortControlPreference() string`

GetPortControlPreference returns the PortControlPreference field if non-nil, zero value otherwise.

### GetPortControlPreferenceOk

`func (o *ControllerSwitchPortManagement) GetPortControlPreferenceOk() (*string, bool)`

GetPortControlPreferenceOk returns a tuple with the PortControlPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlPreference

`func (o *ControllerSwitchPortManagement) SetPortControlPreference(v string)`

SetPortControlPreference sets PortControlPreference field to given value.

### HasPortControlPreference

`func (o *ControllerSwitchPortManagement) HasPortControlPreference() bool`

HasPortControlPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


