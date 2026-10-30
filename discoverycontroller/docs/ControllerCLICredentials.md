# ControllerCLICredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | Pointer to [**ControllerKeyStoreCredential**](ControllerKeyStoreCredential.md) |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 

## Methods

### NewControllerCLICredentials

`func NewControllerCLICredentials() *ControllerCLICredentials`

NewControllerCLICredentials instantiates a new ControllerCLICredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerCLICredentialsWithDefaults

`func NewControllerCLICredentialsWithDefaults() *ControllerCLICredentials`

NewControllerCLICredentialsWithDefaults instantiates a new ControllerCLICredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *ControllerCLICredentials) GetCreds() ControllerKeyStoreCredential`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ControllerCLICredentials) GetCredsOk() (*ControllerKeyStoreCredential, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ControllerCLICredentials) SetCreds(v ControllerKeyStoreCredential)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ControllerCLICredentials) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetPriority

`func (o *ControllerCLICredentials) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ControllerCLICredentials) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ControllerCLICredentials) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ControllerCLICredentials) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


