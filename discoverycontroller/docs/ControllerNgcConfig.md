# ControllerNgcConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**Filters** | Pointer to [**[]ControllerFilter**](ControllerFilter.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IncludeAllDiscoveredObjs** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewControllerNgcConfig

`func NewControllerNgcConfig() *ControllerNgcConfig`

NewControllerNgcConfig instantiates a new ControllerNgcConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNgcConfigWithDefaults

`func NewControllerNgcConfigWithDefaults() *ControllerNgcConfig`

NewControllerNgcConfigWithDefaults instantiates a new ControllerNgcConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *ControllerNgcConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ControllerNgcConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ControllerNgcConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ControllerNgcConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControllerNgcConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerNgcConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerNgcConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerNgcConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFilters

`func (o *ControllerNgcConfig) GetFilters() []ControllerFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ControllerNgcConfig) GetFiltersOk() (*[]ControllerFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ControllerNgcConfig) SetFilters(v []ControllerFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ControllerNgcConfig) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *ControllerNgcConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerNgcConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerNgcConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerNgcConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeAllDiscoveredObjs

`func (o *ControllerNgcConfig) GetIncludeAllDiscoveredObjs() bool`

GetIncludeAllDiscoveredObjs returns the IncludeAllDiscoveredObjs field if non-nil, zero value otherwise.

### GetIncludeAllDiscoveredObjsOk

`func (o *ControllerNgcConfig) GetIncludeAllDiscoveredObjsOk() (*bool, bool)`

GetIncludeAllDiscoveredObjsOk returns a tuple with the IncludeAllDiscoveredObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDiscoveredObjs

`func (o *ControllerNgcConfig) SetIncludeAllDiscoveredObjs(v bool)`

SetIncludeAllDiscoveredObjs sets IncludeAllDiscoveredObjs field to given value.

### HasIncludeAllDiscoveredObjs

`func (o *ControllerNgcConfig) HasIncludeAllDiscoveredObjs() bool`

HasIncludeAllDiscoveredObjs returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerNgcConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerNgcConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerNgcConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerNgcConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


