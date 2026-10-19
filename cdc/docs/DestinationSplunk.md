# DestinationSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCertificate** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CertificateKeyPassphrase** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IndexerName** | Pointer to **string** |  | [optional] 
**InsecureMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutputDataFormat** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDestinationSplunk

`func NewDestinationSplunk() *DestinationSplunk`

NewDestinationSplunk instantiates a new DestinationSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSplunkWithDefaults

`func NewDestinationSplunkWithDefaults() *DestinationSplunk`

NewDestinationSplunkWithDefaults instantiates a new DestinationSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationSplunk) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationSplunk) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationSplunk) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationSplunk) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCertificate

`func (o *DestinationSplunk) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *DestinationSplunk) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *DestinationSplunk) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *DestinationSplunk) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificate

`func (o *DestinationSplunk) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DestinationSplunk) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DestinationSplunk) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *DestinationSplunk) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateKeyPassphrase

`func (o *DestinationSplunk) GetCertificateKeyPassphrase() string`

GetCertificateKeyPassphrase returns the CertificateKeyPassphrase field if non-nil, zero value otherwise.

### GetCertificateKeyPassphraseOk

`func (o *DestinationSplunk) GetCertificateKeyPassphraseOk() (*string, bool)`

GetCertificateKeyPassphraseOk returns a tuple with the CertificateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKeyPassphrase

`func (o *DestinationSplunk) SetCertificateKeyPassphrase(v string)`

SetCertificateKeyPassphrase sets CertificateKeyPassphrase field to given value.

### HasCertificateKeyPassphrase

`func (o *DestinationSplunk) HasCertificateKeyPassphrase() bool`

HasCertificateKeyPassphrase returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationSplunk) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationSplunk) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationSplunk) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationSplunk) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationSplunk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationSplunk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationSplunk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationSplunk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationSplunk) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationSplunk) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationSplunk) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationSplunk) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationSplunk) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationSplunk) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationSplunk) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationSplunk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexerName

`func (o *DestinationSplunk) GetIndexerName() string`

GetIndexerName returns the IndexerName field if non-nil, zero value otherwise.

### GetIndexerNameOk

`func (o *DestinationSplunk) GetIndexerNameOk() (*string, bool)`

GetIndexerNameOk returns a tuple with the IndexerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerName

`func (o *DestinationSplunk) SetIndexerName(v string)`

SetIndexerName sets IndexerName field to given value.

### HasIndexerName

`func (o *DestinationSplunk) HasIndexerName() bool`

HasIndexerName returns a boolean if a field has been set.

### GetInsecureMode

`func (o *DestinationSplunk) GetInsecureMode() bool`

GetInsecureMode returns the InsecureMode field if non-nil, zero value otherwise.

### GetInsecureModeOk

`func (o *DestinationSplunk) GetInsecureModeOk() (*bool, bool)`

GetInsecureModeOk returns a tuple with the InsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureMode

`func (o *DestinationSplunk) SetInsecureMode(v bool)`

SetInsecureMode sets InsecureMode field to given value.

### HasInsecureMode

`func (o *DestinationSplunk) HasInsecureMode() bool`

HasInsecureMode returns a boolean if a field has been set.

### GetName

`func (o *DestinationSplunk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSplunk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSplunk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationSplunk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *DestinationSplunk) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *DestinationSplunk) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *DestinationSplunk) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *DestinationSplunk) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetTags

`func (o *DestinationSplunk) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationSplunk) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationSplunk) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationSplunk) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationSplunk) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationSplunk) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationSplunk) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationSplunk) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


