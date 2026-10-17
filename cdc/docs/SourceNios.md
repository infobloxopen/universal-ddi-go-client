# SourceNios

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCertificate** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**InsecureMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NiosPassword** | Pointer to **string** |  | [optional] 
**NiosUsername** | Pointer to **string** |  | [optional] 
**RpzLogInsecureMode** | Pointer to **bool** |  | [optional] 
**ScpPassword** | Pointer to **string** |  | [optional] 
**ScpUsername** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSourceNios

`func NewSourceNios() *SourceNios`

NewSourceNios instantiates a new SourceNios object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceNiosWithDefaults

`func NewSourceNiosWithDefaults() *SourceNios`

NewSourceNiosWithDefaults instantiates a new SourceNios object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SourceNios) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SourceNios) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SourceNios) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SourceNios) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCertificate

`func (o *SourceNios) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *SourceNios) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *SourceNios) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *SourceNios) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificate

`func (o *SourceNios) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SourceNios) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SourceNios) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SourceNios) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SourceNios) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SourceNios) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SourceNios) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SourceNios) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *SourceNios) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceNios) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceNios) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceNios) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SourceNios) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SourceNios) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SourceNios) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SourceNios) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *SourceNios) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceNios) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceNios) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SourceNios) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsecureMode

`func (o *SourceNios) GetInsecureMode() bool`

GetInsecureMode returns the InsecureMode field if non-nil, zero value otherwise.

### GetInsecureModeOk

`func (o *SourceNios) GetInsecureModeOk() (*bool, bool)`

GetInsecureModeOk returns a tuple with the InsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureMode

`func (o *SourceNios) SetInsecureMode(v bool)`

SetInsecureMode sets InsecureMode field to given value.

### HasInsecureMode

`func (o *SourceNios) HasInsecureMode() bool`

HasInsecureMode returns a boolean if a field has been set.

### GetName

`func (o *SourceNios) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceNios) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceNios) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceNios) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNiosPassword

`func (o *SourceNios) GetNiosPassword() string`

GetNiosPassword returns the NiosPassword field if non-nil, zero value otherwise.

### GetNiosPasswordOk

`func (o *SourceNios) GetNiosPasswordOk() (*string, bool)`

GetNiosPasswordOk returns a tuple with the NiosPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiosPassword

`func (o *SourceNios) SetNiosPassword(v string)`

SetNiosPassword sets NiosPassword field to given value.

### HasNiosPassword

`func (o *SourceNios) HasNiosPassword() bool`

HasNiosPassword returns a boolean if a field has been set.

### GetNiosUsername

`func (o *SourceNios) GetNiosUsername() string`

GetNiosUsername returns the NiosUsername field if non-nil, zero value otherwise.

### GetNiosUsernameOk

`func (o *SourceNios) GetNiosUsernameOk() (*string, bool)`

GetNiosUsernameOk returns a tuple with the NiosUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiosUsername

`func (o *SourceNios) SetNiosUsername(v string)`

SetNiosUsername sets NiosUsername field to given value.

### HasNiosUsername

`func (o *SourceNios) HasNiosUsername() bool`

HasNiosUsername returns a boolean if a field has been set.

### GetRpzLogInsecureMode

`func (o *SourceNios) GetRpzLogInsecureMode() bool`

GetRpzLogInsecureMode returns the RpzLogInsecureMode field if non-nil, zero value otherwise.

### GetRpzLogInsecureModeOk

`func (o *SourceNios) GetRpzLogInsecureModeOk() (*bool, bool)`

GetRpzLogInsecureModeOk returns a tuple with the RpzLogInsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzLogInsecureMode

`func (o *SourceNios) SetRpzLogInsecureMode(v bool)`

SetRpzLogInsecureMode sets RpzLogInsecureMode field to given value.

### HasRpzLogInsecureMode

`func (o *SourceNios) HasRpzLogInsecureMode() bool`

HasRpzLogInsecureMode returns a boolean if a field has been set.

### GetScpPassword

`func (o *SourceNios) GetScpPassword() string`

GetScpPassword returns the ScpPassword field if non-nil, zero value otherwise.

### GetScpPasswordOk

`func (o *SourceNios) GetScpPasswordOk() (*string, bool)`

GetScpPasswordOk returns a tuple with the ScpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpPassword

`func (o *SourceNios) SetScpPassword(v string)`

SetScpPassword sets ScpPassword field to given value.

### HasScpPassword

`func (o *SourceNios) HasScpPassword() bool`

HasScpPassword returns a boolean if a field has been set.

### GetScpUsername

`func (o *SourceNios) GetScpUsername() string`

GetScpUsername returns the ScpUsername field if non-nil, zero value otherwise.

### GetScpUsernameOk

`func (o *SourceNios) GetScpUsernameOk() (*string, bool)`

GetScpUsernameOk returns a tuple with the ScpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpUsername

`func (o *SourceNios) SetScpUsername(v string)`

SetScpUsername sets ScpUsername field to given value.

### HasScpUsername

`func (o *SourceNios) HasScpUsername() bool`

HasScpUsername returns a boolean if a field has been set.

### GetTags

`func (o *SourceNios) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SourceNios) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SourceNios) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SourceNios) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SourceNios) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SourceNios) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SourceNios) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SourceNios) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


