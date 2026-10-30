# NIProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSettings** | Pointer to [**AdvancedSettings**](AdvancedSettings.md) | NI advanced settings. | [optional] 
**CliCredentials** | Pointer to [**[]CLICredentials**](CLICredentials.md) | CLI Credential Settings. | [optional] 
**PollingSettings** | Pointer to [**PollingSettings**](PollingSettings.md) | NI polling settings. | [optional] 
**ServicePorts** | Pointer to [**[]ServicePort**](ServicePort.md) | Service port configuration. | [optional] 
**SnmpCredentials** | Pointer to [**[]SNMPCredentials**](SNMPCredentials.md) | SNMP Credentials Settings. | [optional] 
**SnmpV3Credentials** | Pointer to [**[]SNMPV3Credentials**](SNMPV3Credentials.md) | SNMPv3 Credentials Settings. | [optional] 

## Methods

### NewNIProperties

`func NewNIProperties() *NIProperties`

NewNIProperties instantiates a new NIProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNIPropertiesWithDefaults

`func NewNIPropertiesWithDefaults() *NIProperties`

NewNIPropertiesWithDefaults instantiates a new NIProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSettings

`func (o *NIProperties) GetAdvancedSettings() AdvancedSettings`

GetAdvancedSettings returns the AdvancedSettings field if non-nil, zero value otherwise.

### GetAdvancedSettingsOk

`func (o *NIProperties) GetAdvancedSettingsOk() (*AdvancedSettings, bool)`

GetAdvancedSettingsOk returns a tuple with the AdvancedSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSettings

`func (o *NIProperties) SetAdvancedSettings(v AdvancedSettings)`

SetAdvancedSettings sets AdvancedSettings field to given value.

### HasAdvancedSettings

`func (o *NIProperties) HasAdvancedSettings() bool`

HasAdvancedSettings returns a boolean if a field has been set.

### GetCliCredentials

`func (o *NIProperties) GetCliCredentials() []CLICredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *NIProperties) GetCliCredentialsOk() (*[]CLICredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *NIProperties) SetCliCredentials(v []CLICredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *NIProperties) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetPollingSettings

`func (o *NIProperties) GetPollingSettings() PollingSettings`

GetPollingSettings returns the PollingSettings field if non-nil, zero value otherwise.

### GetPollingSettingsOk

`func (o *NIProperties) GetPollingSettingsOk() (*PollingSettings, bool)`

GetPollingSettingsOk returns a tuple with the PollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingSettings

`func (o *NIProperties) SetPollingSettings(v PollingSettings)`

SetPollingSettings sets PollingSettings field to given value.

### HasPollingSettings

`func (o *NIProperties) HasPollingSettings() bool`

HasPollingSettings returns a boolean if a field has been set.

### GetServicePorts

`func (o *NIProperties) GetServicePorts() []ServicePort`

GetServicePorts returns the ServicePorts field if non-nil, zero value otherwise.

### GetServicePortsOk

`func (o *NIProperties) GetServicePortsOk() (*[]ServicePort, bool)`

GetServicePortsOk returns a tuple with the ServicePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePorts

`func (o *NIProperties) SetServicePorts(v []ServicePort)`

SetServicePorts sets ServicePorts field to given value.

### HasServicePorts

`func (o *NIProperties) HasServicePorts() bool`

HasServicePorts returns a boolean if a field has been set.

### GetSnmpCredentials

`func (o *NIProperties) GetSnmpCredentials() []SNMPCredentials`

GetSnmpCredentials returns the SnmpCredentials field if non-nil, zero value otherwise.

### GetSnmpCredentialsOk

`func (o *NIProperties) GetSnmpCredentialsOk() (*[]SNMPCredentials, bool)`

GetSnmpCredentialsOk returns a tuple with the SnmpCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredentials

`func (o *NIProperties) SetSnmpCredentials(v []SNMPCredentials)`

SetSnmpCredentials sets SnmpCredentials field to given value.

### HasSnmpCredentials

`func (o *NIProperties) HasSnmpCredentials() bool`

HasSnmpCredentials returns a boolean if a field has been set.

### GetSnmpV3Credentials

`func (o *NIProperties) GetSnmpV3Credentials() []SNMPV3Credentials`

GetSnmpV3Credentials returns the SnmpV3Credentials field if non-nil, zero value otherwise.

### GetSnmpV3CredentialsOk

`func (o *NIProperties) GetSnmpV3CredentialsOk() (*[]SNMPV3Credentials, bool)`

GetSnmpV3CredentialsOk returns a tuple with the SnmpV3Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpV3Credentials

`func (o *NIProperties) SetSnmpV3Credentials(v []SNMPV3Credentials)`

SetSnmpV3Credentials sets SnmpV3Credentials field to given value.

### HasSnmpV3Credentials

`func (o *NIProperties) HasSnmpV3Credentials() bool`

HasSnmpV3Credentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


