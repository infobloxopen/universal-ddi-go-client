# ControllerSNMPV3Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | Pointer to [**ControllerKeyStoreCredential**](ControllerKeyStoreCredential.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewControllerSNMPV3Credentials

`func NewControllerSNMPV3Credentials() *ControllerSNMPV3Credentials`

NewControllerSNMPV3Credentials instantiates a new ControllerSNMPV3Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerSNMPV3CredentialsWithDefaults

`func NewControllerSNMPV3CredentialsWithDefaults() *ControllerSNMPV3Credentials`

NewControllerSNMPV3CredentialsWithDefaults instantiates a new ControllerSNMPV3Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *ControllerSNMPV3Credentials) GetCreds() ControllerKeyStoreCredential`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ControllerSNMPV3Credentials) GetCredsOk() (*ControllerKeyStoreCredential, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ControllerSNMPV3Credentials) SetCreds(v ControllerKeyStoreCredential)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ControllerSNMPV3Credentials) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetPriority

`func (o *ControllerSNMPV3Credentials) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ControllerSNMPV3Credentials) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ControllerSNMPV3Credentials) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ControllerSNMPV3Credentials) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


