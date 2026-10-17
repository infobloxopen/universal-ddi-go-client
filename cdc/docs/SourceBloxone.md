# SourceBloxone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewSourceBloxone

`func NewSourceBloxone() *SourceBloxone`

NewSourceBloxone instantiates a new SourceBloxone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceBloxoneWithDefaults

`func NewSourceBloxoneWithDefaults() *SourceBloxone`

NewSourceBloxoneWithDefaults instantiates a new SourceBloxone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SourceBloxone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SourceBloxone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SourceBloxone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SourceBloxone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SourceBloxone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SourceBloxone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SourceBloxone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SourceBloxone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *SourceBloxone) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceBloxone) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceBloxone) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SourceBloxone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SourceBloxone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceBloxone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceBloxone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceBloxone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *SourceBloxone) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SourceBloxone) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SourceBloxone) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *SourceBloxone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SourceBloxone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SourceBloxone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SourceBloxone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SourceBloxone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


