# SNMPV3Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | Pointer to [**KeyStoreCredential**](KeyStoreCredential.md) |  | [optional] 
**Priority** | Pointer to **int64** |  | [optional] 

## Methods

### NewSNMPV3Credentials

`func NewSNMPV3Credentials() *SNMPV3Credentials`

NewSNMPV3Credentials instantiates a new SNMPV3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPV3CredentialsWithDefaults

`func NewSNMPV3CredentialsWithDefaults() *SNMPV3Credentials`

NewSNMPV3CredentialsWithDefaults instantiates a new SNMPV3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *SNMPV3Credentials) GetCreds() KeyStoreCredential`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *SNMPV3Credentials) GetCredsOk() (*KeyStoreCredential, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *SNMPV3Credentials) SetCreds(v KeyStoreCredential)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *SNMPV3Credentials) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetPriority

`func (o *SNMPV3Credentials) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SNMPV3Credentials) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SNMPV3Credentials) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SNMPV3Credentials) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


