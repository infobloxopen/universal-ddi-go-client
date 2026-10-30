# SDNConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCert** | Pointer to **string** |  | [optional] 
**Creds** | Pointer to [**KeyStoreCredential**](KeyStoreCredential.md) |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**OnPrem** | Pointer to **bool** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SdnType** | Pointer to **string** |  | [optional] 
**UseGlobalProxy** | Pointer to **bool** |  | [optional] 

## Methods

### NewSDNConfiguration

`func NewSDNConfiguration() *SDNConfiguration`

NewSDNConfiguration instantiates a new SDNConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDNConfigurationWithDefaults

`func NewSDNConfigurationWithDefaults() *SDNConfiguration`

NewSDNConfigurationWithDefaults instantiates a new SDNConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SDNConfiguration) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SDNConfiguration) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SDNConfiguration) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SDNConfiguration) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCert

`func (o *SDNConfiguration) GetCaCert() string`

GetCaCert returns the CaCert field if non-nil, zero value otherwise.

### GetCaCertOk

`func (o *SDNConfiguration) GetCaCertOk() (*string, bool)`

GetCaCertOk returns a tuple with the CaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCert

`func (o *SDNConfiguration) SetCaCert(v string)`

SetCaCert sets CaCert field to given value.

### HasCaCert

`func (o *SDNConfiguration) HasCaCert() bool`

HasCaCert returns a boolean if a field has been set.

### GetCreds

`func (o *SDNConfiguration) GetCreds() KeyStoreCredential`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *SDNConfiguration) GetCredsOk() (*KeyStoreCredential, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *SDNConfiguration) SetCreds(v KeyStoreCredential)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *SDNConfiguration) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetHandle

`func (o *SDNConfiguration) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SDNConfiguration) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SDNConfiguration) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SDNConfiguration) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetOnPrem

`func (o *SDNConfiguration) GetOnPrem() bool`

GetOnPrem returns the OnPrem field if non-nil, zero value otherwise.

### GetOnPremOk

`func (o *SDNConfiguration) GetOnPremOk() (*bool, bool)`

GetOnPremOk returns a tuple with the OnPrem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnPrem

`func (o *SDNConfiguration) SetOnPrem(v bool)`

SetOnPrem sets OnPrem field to given value.

### HasOnPrem

`func (o *SDNConfiguration) HasOnPrem() bool`

HasOnPrem returns a boolean if a field has been set.

### GetProtocol

`func (o *SDNConfiguration) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SDNConfiguration) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SDNConfiguration) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SDNConfiguration) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSdnType

`func (o *SDNConfiguration) GetSdnType() string`

GetSdnType returns the SdnType field if non-nil, zero value otherwise.

### GetSdnTypeOk

`func (o *SDNConfiguration) GetSdnTypeOk() (*string, bool)`

GetSdnTypeOk returns a tuple with the SdnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnType

`func (o *SDNConfiguration) SetSdnType(v string)`

SetSdnType sets SdnType field to given value.

### HasSdnType

`func (o *SDNConfiguration) HasSdnType() bool`

HasSdnType returns a boolean if a field has been set.

### GetUseGlobalProxy

`func (o *SDNConfiguration) GetUseGlobalProxy() bool`

GetUseGlobalProxy returns the UseGlobalProxy field if non-nil, zero value otherwise.

### GetUseGlobalProxyOk

`func (o *SDNConfiguration) GetUseGlobalProxyOk() (*bool, bool)`

GetUseGlobalProxyOk returns a tuple with the UseGlobalProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGlobalProxy

`func (o *SDNConfiguration) SetUseGlobalProxy(v bool)`

SetUseGlobalProxy sets UseGlobalProxy field to given value.

### HasUseGlobalProxy

`func (o *SDNConfiguration) HasUseGlobalProxy() bool`

HasUseGlobalProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


