# ControllerNIProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSettings** | Pointer to [**ControllerAdvancedSettings**](ControllerAdvancedSettings.md) | NI advanced settings. | [optional] 
**CliCredentials** | Pointer to [**[]ControllerCLICredentials**](ControllerCLICredentials.md) | CLI Credential Settings. | [optional] 
**PollingSettings** | Pointer to [**ControllerPollingSettings**](ControllerPollingSettings.md) | NI polling settings. | [optional] 
**ServicePorts** | Pointer to [**[]ControllerServicePort**](ControllerServicePort.md) | Service port configuration. | [optional] 
**SnmpCredentials** | Pointer to [**[]ControllerSNMPCredentials**](ControllerSNMPCredentials.md) | SNMP Credentials Settings. | [optional] 
**SnmpV3Credentials** | Pointer to [**[]ControllerSNMPV3Credentials**](ControllerSNMPV3Credentials.md) | SNMPv3 Credentials Settings. | [optional] 

## Methods

### NewControllerNIProperties

`func NewControllerNIProperties() *ControllerNIProperties`

NewControllerNIProperties instantiates a new ControllerNIProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNIPropertiesWithDefaults

`func NewControllerNIPropertiesWithDefaults() *ControllerNIProperties`

NewControllerNIPropertiesWithDefaults instantiates a new ControllerNIProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSettings

`func (o *ControllerNIProperties) GetAdvancedSettings() ControllerAdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *ControllerNIProperties) GetAdvancedSettingsOk() (*ControllerAdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *ControllerNIProperties) SetAdvancedSettings(v ControllerAdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *ControllerNIProperties) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetCliCredentials

`func (o *ControllerNIProperties) GetCliCredentials() []ControllerCLICredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *ControllerNIProperties) GetCliCredentialsOk() (*[]ControllerCLICredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *ControllerNIProperties) SetCliCredentials(v []ControllerCLICredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *ControllerNIProperties) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetPollingSettings

`func (o *ControllerNIProperties) GetPollingSettings() ControllerPollingSettings`

GetPollingSettings returns the PollingSettings field if non-nil, zero value otherwise.

### GetPollingSettingsOk

`func (o *ControllerNIProperties) GetPollingSettingsOk() (*ControllerPollingSettings, bool)`

GetPollingSettingsOk returns a tuple with the PollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingSettings

`func (o *ControllerNIProperties) SetPollingSettings(v ControllerPollingSettings)`

SetPollingSettings sets PollingSettings field to given value.

### HasPollingSettings

`func (o *ControllerNIProperties) HasPollingSettings() bool`

HasPollingSettings returns a boolean if a field has been set.

### GetServicePorts

`func (o *ControllerNIProperties) GetServicePorts() []ControllerServicePort`

GetServicePorts returns the ServicePorts field if non-nil, zero value otherwise.

### GetServicePortsOk

`func (o *ControllerNIProperties) GetServicePortsOk() (*[]ControllerServicePort, bool)`

GetServicePortsOk returns a tuple with the ServicePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePorts

`func (o *ControllerNIProperties) SetServicePorts(v []ControllerServicePort)`

SetServicePorts sets ServicePorts field to given value.

### HasServicePorts

`func (o *ControllerNIProperties) HasServicePorts() bool`

HasServicePorts returns a boolean if a field has been set.

### GetSnmpCredentials

`func (o *ControllerNIProperties) GetSnmpCredentials() []ControllerSNMPCredentials`

GetSnmpCredentials returns the SnmpCredentials field if non-nil, zero value otherwise.

### GetSnmpCredentialsOk

`func (o *ControllerNIProperties) GetSnmpCredentialsOk() (*[]ControllerSNMPCredentials, bool)`

GetSnmpCredentialsOk returns a tuple with the SnmpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredentials

`func (o *ControllerNIProperties) SetSnmpCredentials(v []ControllerSNMPCredentials)`

SetSnmpCredentials sets SnmpCredentials field to given value.

### HasSnmpCredentials

`func (o *ControllerNIProperties) HasSnmpCredentials() bool`

HasSnmpCredentials returns a boolean if a field has been set.

### GetSnmpV3Credentials

`func (o *ControllerNIProperties) GetSnmpV3Credentials() []ControllerSNMPV3Credentials`

GetSnmpV3Credentials returns the SnmpV3Credentials field if non-nil, zero value otherwise.

### GetSnmpV3CredentialsOk

`func (o *ControllerNIProperties) GetSnmpV3CredentialsOk() (*[]ControllerSNMPV3Credentials, bool)`

GetSnmpV3CredentialsOk returns a tuple with the SnmpV3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpV3Credentials

`func (o *ControllerNIProperties) SetSnmpV3Credentials(v []ControllerSNMPV3Credentials)`

SetSnmpV3Credentials sets SnmpV3Credentials field to given value.

### HasSnmpV3Credentials

`func (o *ControllerNIProperties) HasSnmpV3Credentials() bool`

HasSnmpV3Credentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


