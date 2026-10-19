# DestinationHttp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Authentication** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CredentialId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**InsecureMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutputDataFormat** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**TransportProtocol** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDestinationHttp

`func NewDestinationHttp() *DestinationHttp`

NewDestinationHttp instantiates a new DestinationHttp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationHttpWithDefaults

`func NewDestinationHttpWithDefaults() *DestinationHttp`

NewDestinationHttpWithDefaults instantiates a new DestinationHttp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationHttp) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationHttp) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationHttp) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationHttp) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAuthentication

`func (o *DestinationHttp) GetAuthentication() string`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *DestinationHttp) GetAuthenticationOk() (*string, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *DestinationHttp) SetAuthentication(v string)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *DestinationHttp) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationHttp) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationHttp) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationHttp) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationHttp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialId

`func (o *DestinationHttp) GetCredentialId() string`

GetCredentialId returns the CredentialId field if non-nil, zero value otherwise.

### GetCredentialIdOk

`func (o *DestinationHttp) GetCredentialIdOk() (*string, bool)`

GetCredentialIdOk returns a tuple with the CredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialId

`func (o *DestinationHttp) SetCredentialId(v string)`

SetCredentialId sets CredentialId field to given value.

### HasCredentialId

`func (o *DestinationHttp) HasCredentialId() bool`

HasCredentialId returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationHttp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationHttp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationHttp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationHttp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationHttp) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationHttp) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationHttp) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationHttp) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationHttp) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationHttp) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationHttp) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationHttp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsecureMode

`func (o *DestinationHttp) GetInsecureMode() bool`

GetInsecureMode returns the InsecureMode field if non-nil, zero value otherwise.

### GetInsecureModeOk

`func (o *DestinationHttp) GetInsecureModeOk() (*bool, bool)`

GetInsecureModeOk returns a tuple with the InsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureMode

`func (o *DestinationHttp) SetInsecureMode(v bool)`

SetInsecureMode sets InsecureMode field to given value.

### HasInsecureMode

`func (o *DestinationHttp) HasInsecureMode() bool`

HasInsecureMode returns a boolean if a field has been set.

### GetName

`func (o *DestinationHttp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationHttp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationHttp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationHttp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *DestinationHttp) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *DestinationHttp) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *DestinationHttp) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *DestinationHttp) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetTags

`func (o *DestinationHttp) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationHttp) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationHttp) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationHttp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *DestinationHttp) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *DestinationHttp) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *DestinationHttp) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *DestinationHttp) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationHttp) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationHttp) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationHttp) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationHttp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


