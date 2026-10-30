# ControllerSDNConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCert** | Pointer to **string** |  | [optional] 
**Creds** | Pointer to [**ControllerKeyStoreCredential**](ControllerKeyStoreCredential.md) |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**OnPrem** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SdnType** | Pointer to **string** |  | [optional] 
**UseGlobalProxy** | Pointer to **bool** |  | [optional] 

## Methods

### NewControllerSDNConfiguration

`func NewControllerSDNConfiguration() *ControllerSDNConfiguration`

NewControllerSDNConfiguration instantiates a new ControllerSDNConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerSDNConfigurationWithDefaults

`func NewControllerSDNConfigurationWithDefaults() *ControllerSDNConfiguration`

NewControllerSDNConfigurationWithDefaults instantiates a new ControllerSDNConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ControllerSDNConfiguration) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ControllerSDNConfiguration) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ControllerSDNConfiguration) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ControllerSDNConfiguration) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *ControllerSDNConfiguration) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *ControllerSDNConfiguration) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *ControllerSDNConfiguration) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *ControllerSDNConfiguration) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCreds

`func (o *ControllerSDNConfiguration) GetCreds() ControllerKeyStoreCredential`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ControllerSDNConfiguration) GetCredsOk() (*ControllerKeyStoreCredential, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ControllerSDNConfiguration) SetCreds(v ControllerKeyStoreCredential)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ControllerSDNConfiguration) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetHandle

`func (o *ControllerSDNConfiguration) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *ControllerSDNConfiguration) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *ControllerSDNConfiguration) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *ControllerSDNConfiguration) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetOnPrem

`func (o *ControllerSDNConfiguration) GetOnPrem() bool`

GetOnPrem returns the OnPrem field if non-nil, zero value otherwise.

### GetOnPremOk

`func (o *ControllerSDNConfiguration) GetOnPremOk() (*bool, bool)`

GetOnPremOk returns a tuple with the OnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPrem

`func (o *ControllerSDNConfiguration) SetOnPrem(v bool)`

SetOnPrem sets OnPrem field to given value.

### HasOnPrem

`func (o *ControllerSDNConfiguration) HasOnPrem() bool`

HasOnPrem returns a boolean if a field has been set.

### GetProtocol

`func (o *ControllerSDNConfiguration) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ControllerSDNConfiguration) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ControllerSDNConfiguration) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ControllerSDNConfiguration) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSdnType

`func (o *ControllerSDNConfiguration) GetSdnType() string`

GetSdnType returns the SdnType field if non-nil, zero value otherwise.

### GetSdnTypeOk

`func (o *ControllerSDNConfiguration) GetSdnTypeOk() (*string, bool)`

GetSdnTypeOk returns a tuple with the SdnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnType

`func (o *ControllerSDNConfiguration) SetSdnType(v string)`

SetSdnType sets SdnType field to given value.

### HasSdnType

`func (o *ControllerSDNConfiguration) HasSdnType() bool`

HasSdnType returns a boolean if a field has been set.

### GetUseGlobalProxy

`func (o *ControllerSDNConfiguration) GetUseGlobalProxy() bool`

GetUseGlobalProxy returns the UseGlobalProxy field if non-nil, zero value otherwise.

### GetUseGlobalProxyOk

`func (o *ControllerSDNConfiguration) GetUseGlobalProxyOk() (*bool, bool)`

GetUseGlobalProxyOk returns a tuple with the UseGlobalProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalProxy

`func (o *ControllerSDNConfiguration) SetUseGlobalProxy(v bool)`

SetUseGlobalProxy sets UseGlobalProxy field to given value.

### HasUseGlobalProxy

`func (o *ControllerSDNConfiguration) HasUseGlobalProxy() bool`

HasUseGlobalProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


