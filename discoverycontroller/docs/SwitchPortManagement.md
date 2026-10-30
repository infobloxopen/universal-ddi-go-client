# SwitchPortManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**PollingInterval** | Pointer to **int64** | Polling interval in minutes. | [optional] 
**PortControlPreference** | Pointer to **string** | Port control preference corresponds to one of SNMP or CLI. | [optional] 

## Methods

### NewSwitchPortManagement

`func NewSwitchPortManagement() *SwitchPortManagement`

NewSwitchPortManagement instantiates a new SwitchPortManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchPortManagementWithDefaults

`func NewSwitchPortManagementWithDefaults() *SwitchPortManagement`

NewSwitchPortManagementWithDefaults instantiates a new SwitchPortManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SwitchPortManagement) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SwitchPortManagement) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SwitchPortManagement) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SwitchPortManagement) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPollingInterval

`func (o *SwitchPortManagement) GetPollingInterval() int64`

GetPollingInterval returns the PollingInterval field if non-nil, zero value otherwise.

### GetPollingIntervalOk

`func (o *SwitchPortManagement) GetPollingIntervalOk() (*int64, bool)`

GetPollingIntervalOk returns a tuple with the PollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingInterval

`func (o *SwitchPortManagement) SetPollingInterval(v int64)`

SetPollingInterval sets PollingInterval field to given value.

### HasPollingInterval

`func (o *SwitchPortManagement) HasPollingInterval() bool`

HasPollingInterval returns a boolean if a field has been set.

### GetPortControlPreference

`func (o *SwitchPortManagement) GetPortControlPreference() string`

GetPortControlPreference returns the PortControlPreference field if non-nil, zero value otherwise.

### GetPortControlPreferenceOk

`func (o *SwitchPortManagement) GetPortControlPreferenceOk() (*string, bool)`

GetPortControlPreferenceOk returns a tuple with the PortControlPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortControlPreference

`func (o *SwitchPortManagement) SetPortControlPreference(v string)`

SetPortControlPreference sets PortControlPreference field to given value.

### HasPortControlPreference

`func (o *SwitchPortManagement) HasPortControlPreference() bool`

HasPortControlPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


