# NgcConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**Filters** | Pointer to [**[]Filter**](Filter.md) |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IncludeAllDiscoveredObjs** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewNgcConfig

`func NewNgcConfig() *NgcConfig`

NewNgcConfig instantiates a new NgcConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgcConfigWithDefaults

`func NewNgcConfigWithDefaults() *NgcConfig`

NewNgcConfigWithDefaults instantiates a new NgcConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *NgcConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *NgcConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *NgcConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *NgcConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NgcConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NgcConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NgcConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NgcConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFilters

`func (o *NgcConfig) GetFilters() []Filter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NgcConfig) GetFiltersOk() (*[]Filter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NgcConfig) SetFilters(v []Filter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *NgcConfig) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetId

`func (o *NgcConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NgcConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NgcConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NgcConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeAllDiscoveredObjs

`func (o *NgcConfig) GetIncludeAllDiscoveredObjs() bool`

GetIncludeAllDiscoveredObjs returns the IncludeAllDiscoveredObjs field if non-nil, zero value otherwise.

### GetIncludeAllDiscoveredObjsOk

`func (o *NgcConfig) GetIncludeAllDiscoveredObjsOk() (*bool, bool)`

GetIncludeAllDiscoveredObjsOk returns a tuple with the IncludeAllDiscoveredObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllDiscoveredObjs

`func (o *NgcConfig) SetIncludeAllDiscoveredObjs(v bool)`

SetIncludeAllDiscoveredObjs sets IncludeAllDiscoveredObjs field to given value.

### HasIncludeAllDiscoveredObjs

`func (o *NgcConfig) HasIncludeAllDiscoveredObjs() bool`

HasIncludeAllDiscoveredObjs returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NgcConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NgcConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NgcConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NgcConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


