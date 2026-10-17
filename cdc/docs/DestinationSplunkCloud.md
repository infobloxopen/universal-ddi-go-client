# DestinationSplunkCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCertificate** | Pointer to **string** |  | [optional] 
**Certificate** | Pointer to **string** |  | [optional] 
**CertificateKeyPassphrase** | Pointer to **string** |  | [optional] 
**ConfigFilename** | Pointer to **string** |  | [optional] 
**ConfigOption** | Pointer to [**ConfigOption**](ConfigOption.md) |  | [optional] [default to CONFIGOPTION_DEFAULT]
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

### NewDestinationSplunkCloud

`func NewDestinationSplunkCloud() *DestinationSplunkCloud`

NewDestinationSplunkCloud instantiates a new DestinationSplunkCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSplunkCloudWithDefaults

`func NewDestinationSplunkCloudWithDefaults() *DestinationSplunkCloud`

NewDestinationSplunkCloudWithDefaults instantiates a new DestinationSplunkCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationSplunkCloud) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationSplunkCloud) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationSplunkCloud) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationSplunkCloud) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCertificate

`func (o *DestinationSplunkCloud) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *DestinationSplunkCloud) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *DestinationSplunkCloud) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *DestinationSplunkCloud) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCertificate

`func (o *DestinationSplunkCloud) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *DestinationSplunkCloud) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *DestinationSplunkCloud) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *DestinationSplunkCloud) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateKeyPassphrase

`func (o *DestinationSplunkCloud) GetCertificateKeyPassphrase() string`

GetCertificateKeyPassphrase returns the CertificateKeyPassphrase field if non-nil, zero value otherwise.

### GetCertificateKeyPassphraseOk

`func (o *DestinationSplunkCloud) GetCertificateKeyPassphraseOk() (*string, bool)`

GetCertificateKeyPassphraseOk returns a tuple with the CertificateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateKeyPassphrase

`func (o *DestinationSplunkCloud) SetCertificateKeyPassphrase(v string)`

SetCertificateKeyPassphrase sets CertificateKeyPassphrase field to given value.

### HasCertificateKeyPassphrase

`func (o *DestinationSplunkCloud) HasCertificateKeyPassphrase() bool`

HasCertificateKeyPassphrase returns a boolean if a field has been set.

### GetConfigFilename

`func (o *DestinationSplunkCloud) GetConfigFilename() string`

GetConfigFilename returns the ConfigFilename field if non-nil, zero value otherwise.

### GetConfigFilenameOk

`func (o *DestinationSplunkCloud) GetConfigFilenameOk() (*string, bool)`

GetConfigFilenameOk returns a tuple with the ConfigFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFilename

`func (o *DestinationSplunkCloud) SetConfigFilename(v string)`

SetConfigFilename sets ConfigFilename field to given value.

### HasConfigFilename

`func (o *DestinationSplunkCloud) HasConfigFilename() bool`

HasConfigFilename returns a boolean if a field has been set.

### GetConfigOption

`func (o *DestinationSplunkCloud) GetConfigOption() ConfigOption`

GetConfigOption returns the ConfigOption field if non-nil, zero value otherwise.

### GetConfigOptionOk

`func (o *DestinationSplunkCloud) GetConfigOptionOk() (*ConfigOption, bool)`

GetConfigOptionOk returns a tuple with the ConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOption

`func (o *DestinationSplunkCloud) SetConfigOption(v ConfigOption)`

SetConfigOption sets ConfigOption field to given value.

### HasConfigOption

`func (o *DestinationSplunkCloud) HasConfigOption() bool`

HasConfigOption returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationSplunkCloud) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationSplunkCloud) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationSplunkCloud) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationSplunkCloud) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationSplunkCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationSplunkCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationSplunkCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationSplunkCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationSplunkCloud) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationSplunkCloud) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationSplunkCloud) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationSplunkCloud) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationSplunkCloud) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationSplunkCloud) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationSplunkCloud) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationSplunkCloud) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexerName

`func (o *DestinationSplunkCloud) GetIndexerName() string`

GetIndexerName returns the IndexerName field if non-nil, zero value otherwise.

### GetIndexerNameOk

`func (o *DestinationSplunkCloud) GetIndexerNameOk() (*string, bool)`

GetIndexerNameOk returns a tuple with the IndexerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexerName

`func (o *DestinationSplunkCloud) SetIndexerName(v string)`

SetIndexerName sets IndexerName field to given value.

### HasIndexerName

`func (o *DestinationSplunkCloud) HasIndexerName() bool`

HasIndexerName returns a boolean if a field has been set.

### GetInsecureMode

`func (o *DestinationSplunkCloud) GetInsecureMode() bool`

GetInsecureMode returns the InsecureMode field if non-nil, zero value otherwise.

### GetInsecureModeOk

`func (o *DestinationSplunkCloud) GetInsecureModeOk() (*bool, bool)`

GetInsecureModeOk returns a tuple with the InsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureMode

`func (o *DestinationSplunkCloud) SetInsecureMode(v bool)`

SetInsecureMode sets InsecureMode field to given value.

### HasInsecureMode

`func (o *DestinationSplunkCloud) HasInsecureMode() bool`

HasInsecureMode returns a boolean if a field has been set.

### GetName

`func (o *DestinationSplunkCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSplunkCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSplunkCloud) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationSplunkCloud) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *DestinationSplunkCloud) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *DestinationSplunkCloud) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *DestinationSplunkCloud) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *DestinationSplunkCloud) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetTags

`func (o *DestinationSplunkCloud) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationSplunkCloud) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationSplunkCloud) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationSplunkCloud) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationSplunkCloud) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationSplunkCloud) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationSplunkCloud) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationSplunkCloud) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


