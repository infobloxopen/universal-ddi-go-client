# PassiveDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsUnmanaged** | Pointer to **bool** | Indicates whether tagging is enabled for this passive discovery configuration. | [optional] 
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**EnabledSources** | Pointer to [**[]PassiveDiscoverySource**](PassiveDiscoverySource.md) | List of passive discovery sources for which job is enabled. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IsSocInsightEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**IsTdEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewPassiveDiscoveryConfig

`func NewPassiveDiscoveryConfig() *PassiveDiscoveryConfig`

NewPassiveDiscoveryConfig instantiates a new PassiveDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassiveDiscoveryConfigWithDefaults

`func NewPassiveDiscoveryConfigWithDefaults() *PassiveDiscoveryConfig`

NewPassiveDiscoveryConfigWithDefaults instantiates a new PassiveDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsUnmanaged

`func (o *PassiveDiscoveryConfig) GetAssetsUnmanaged() bool`

GetAssetsUnmanaged returns the AssetsUnmanaged field if non-nil, zero value otherwise.

### GetAssetsUnmanagedOk

`func (o *PassiveDiscoveryConfig) GetAssetsUnmanagedOk() (*bool, bool)`

GetAssetsUnmanagedOk returns a tuple with the AssetsUnmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsUnmanaged

`func (o *PassiveDiscoveryConfig) SetAssetsUnmanaged(v bool)`

SetAssetsUnmanaged sets AssetsUnmanaged field to given value.

### HasAssetsUnmanaged

`func (o *PassiveDiscoveryConfig) HasAssetsUnmanaged() bool`

HasAssetsUnmanaged returns a boolean if a field has been set.

### GetConfigId

`func (o *PassiveDiscoveryConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PassiveDiscoveryConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PassiveDiscoveryConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PassiveDiscoveryConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PassiveDiscoveryConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PassiveDiscoveryConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PassiveDiscoveryConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PassiveDiscoveryConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabledSources

`func (o *PassiveDiscoveryConfig) GetEnabledSources() []PassiveDiscoverySource`

GetEnabledSources returns the EnabledSources field if non-nil, zero value otherwise.

### GetEnabledSourcesOk

`func (o *PassiveDiscoveryConfig) GetEnabledSourcesOk() (*[]PassiveDiscoverySource, bool)`

GetEnabledSourcesOk returns a tuple with the EnabledSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledSources

`func (o *PassiveDiscoveryConfig) SetEnabledSources(v []PassiveDiscoverySource)`

SetEnabledSources sets EnabledSources field to given value.

### HasEnabledSources

`func (o *PassiveDiscoveryConfig) HasEnabledSources() bool`

HasEnabledSources returns a boolean if a field has been set.

### GetId

`func (o *PassiveDiscoveryConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PassiveDiscoveryConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PassiveDiscoveryConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PassiveDiscoveryConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSocInsightEnabled

`func (o *PassiveDiscoveryConfig) GetIsSocInsightEnabled() bool`

GetIsSocInsightEnabled returns the IsSocInsightEnabled field if non-nil, zero value otherwise.

### GetIsSocInsightEnabledOk

`func (o *PassiveDiscoveryConfig) GetIsSocInsightEnabledOk() (*bool, bool)`

GetIsSocInsightEnabledOk returns a tuple with the IsSocInsightEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSocInsightEnabled

`func (o *PassiveDiscoveryConfig) SetIsSocInsightEnabled(v bool)`

SetIsSocInsightEnabled sets IsSocInsightEnabled field to given value.

### HasIsSocInsightEnabled

`func (o *PassiveDiscoveryConfig) HasIsSocInsightEnabled() bool`

HasIsSocInsightEnabled returns a boolean if a field has been set.

### GetIsTdEnabled

`func (o *PassiveDiscoveryConfig) GetIsTdEnabled() bool`

GetIsTdEnabled returns the IsTdEnabled field if non-nil, zero value otherwise.

### GetIsTdEnabledOk

`func (o *PassiveDiscoveryConfig) GetIsTdEnabledOk() (*bool, bool)`

GetIsTdEnabledOk returns a tuple with the IsTdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTdEnabled

`func (o *PassiveDiscoveryConfig) SetIsTdEnabled(v bool)`

SetIsTdEnabled sets IsTdEnabled field to given value.

### HasIsTdEnabled

`func (o *PassiveDiscoveryConfig) HasIsTdEnabled() bool`

HasIsTdEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PassiveDiscoveryConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PassiveDiscoveryConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PassiveDiscoveryConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PassiveDiscoveryConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


