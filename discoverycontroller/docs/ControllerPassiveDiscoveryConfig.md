# ControllerPassiveDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsUnmanaged** | Pointer to **bool** | Indicates whether tagging is enabled for this passive discovery configuration. | [optional] 
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**EnabledSources** | Pointer to [**[]ControllerPassiveDiscoverySource**](ControllerPassiveDiscoverySource.md) | List of passive discovery sources for which job is enabled. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IsSocInsightEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**IsTdEnabled** | Pointer to **bool** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewControllerPassiveDiscoveryConfig

`func NewControllerPassiveDiscoveryConfig() *ControllerPassiveDiscoveryConfig`

NewControllerPassiveDiscoveryConfig instantiates a new ControllerPassiveDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerPassiveDiscoveryConfigWithDefaults

`func NewControllerPassiveDiscoveryConfigWithDefaults() *ControllerPassiveDiscoveryConfig`

NewControllerPassiveDiscoveryConfigWithDefaults instantiates a new ControllerPassiveDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsUnmanaged

`func (o *ControllerPassiveDiscoveryConfig) GetAssetsUnmanaged() bool`

GetAssetsUnmanaged returns the AssetsUnmanaged field if non-nil, zero value otherwise.

### GetAssetsUnmanagedOk

`func (o *ControllerPassiveDiscoveryConfig) GetAssetsUnmanagedOk() (*bool, bool)`

GetAssetsUnmanagedOk returns a tuple with the AssetsUnmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsUnmanaged

`func (o *ControllerPassiveDiscoveryConfig) SetAssetsUnmanaged(v bool)`

SetAssetsUnmanaged sets AssetsUnmanaged field to given value.

### HasAssetsUnmanaged

`func (o *ControllerPassiveDiscoveryConfig) HasAssetsUnmanaged() bool`

HasAssetsUnmanaged returns a boolean if a field has been set.

### GetConfigId

`func (o *ControllerPassiveDiscoveryConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ControllerPassiveDiscoveryConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ControllerPassiveDiscoveryConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ControllerPassiveDiscoveryConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControllerPassiveDiscoveryConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerPassiveDiscoveryConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerPassiveDiscoveryConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerPassiveDiscoveryConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabledSources

`func (o *ControllerPassiveDiscoveryConfig) GetEnabledSources() []ControllerPassiveDiscoverySource`

GetEnabledSources returns the EnabledSources field if non-nil, zero value otherwise.

### GetEnabledSourcesOk

`func (o *ControllerPassiveDiscoveryConfig) GetEnabledSourcesOk() (*[]ControllerPassiveDiscoverySource, bool)`

GetEnabledSourcesOk returns a tuple with the EnabledSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledSources

`func (o *ControllerPassiveDiscoveryConfig) SetEnabledSources(v []ControllerPassiveDiscoverySource)`

SetEnabledSources sets EnabledSources field to given value.

### HasEnabledSources

`func (o *ControllerPassiveDiscoveryConfig) HasEnabledSources() bool`

HasEnabledSources returns a boolean if a field has been set.

### GetId

`func (o *ControllerPassiveDiscoveryConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerPassiveDiscoveryConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerPassiveDiscoveryConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerPassiveDiscoveryConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSocInsightEnabled

`func (o *ControllerPassiveDiscoveryConfig) GetIsSocInsightEnabled() bool`

GetIsSocInsightEnabled returns the IsSocInsightEnabled field if non-nil, zero value otherwise.

### GetIsSocInsightEnabledOk

`func (o *ControllerPassiveDiscoveryConfig) GetIsSocInsightEnabledOk() (*bool, bool)`

GetIsSocInsightEnabledOk returns a tuple with the IsSocInsightEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSocInsightEnabled

`func (o *ControllerPassiveDiscoveryConfig) SetIsSocInsightEnabled(v bool)`

SetIsSocInsightEnabled sets IsSocInsightEnabled field to given value.

### HasIsSocInsightEnabled

`func (o *ControllerPassiveDiscoveryConfig) HasIsSocInsightEnabled() bool`

HasIsSocInsightEnabled returns a boolean if a field has been set.

### GetIsTdEnabled

`func (o *ControllerPassiveDiscoveryConfig) GetIsTdEnabled() bool`

GetIsTdEnabled returns the IsTdEnabled field if non-nil, zero value otherwise.

### GetIsTdEnabledOk

`func (o *ControllerPassiveDiscoveryConfig) GetIsTdEnabledOk() (*bool, bool)`

GetIsTdEnabledOk returns a tuple with the IsTdEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTdEnabled

`func (o *ControllerPassiveDiscoveryConfig) SetIsTdEnabled(v bool)`

SetIsTdEnabled sets IsTdEnabled field to given value.

### HasIsTdEnabled

`func (o *ControllerPassiveDiscoveryConfig) HasIsTdEnabled() bool`

HasIsTdEnabled returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerPassiveDiscoveryConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerPassiveDiscoveryConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerPassiveDiscoveryConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerPassiveDiscoveryConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


