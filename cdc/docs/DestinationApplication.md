# DestinationApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **int64** |  | [optional] 
**Context** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**OutputDataFormat** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDestinationApplication

`func NewDestinationApplication() *DestinationApplication`

NewDestinationApplication instantiates a new DestinationApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationWithDefaults

`func NewDestinationApplicationWithDefaults() *DestinationApplication`

NewDestinationApplicationWithDefaults instantiates a new DestinationApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DestinationApplication) GetApplication() int64`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplication) GetApplicationOk() (*int64, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplication) SetApplication(v int64)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DestinationApplication) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetContext

`func (o *DestinationApplication) GetContext() map[string]map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationApplication) GetContextOk() (*map[string]map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationApplication) SetContext(v map[string]map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *DestinationApplication) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationApplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationApplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationApplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationApplication) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationApplication) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationApplication) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationApplication) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationApplication) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationApplication) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationApplication) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationApplication) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DestinationApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *DestinationApplication) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *DestinationApplication) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *DestinationApplication) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *DestinationApplication) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetTags

`func (o *DestinationApplication) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationApplication) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationApplication) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationApplication) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


