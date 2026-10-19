# EtlFilterRegex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Data** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**FilterSource** | Pointer to [**FilterSource**](FilterSource.md) |  | [optional] [default to FILTERSOURCE_DEFAULT]
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IngressPolicyAction** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEtlFilterRegex

`func NewEtlFilterRegex() *EtlFilterRegex`

NewEtlFilterRegex instantiates a new EtlFilterRegex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlFilterRegexWithDefaults

`func NewEtlFilterRegexWithDefaults() *EtlFilterRegex`

NewEtlFilterRegexWithDefaults instantiates a new EtlFilterRegex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EtlFilterRegex) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EtlFilterRegex) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EtlFilterRegex) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EtlFilterRegex) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EtlFilterRegex) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EtlFilterRegex) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EtlFilterRegex) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *EtlFilterRegex) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *EtlFilterRegex) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtlFilterRegex) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtlFilterRegex) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtlFilterRegex) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EtlFilterRegex) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EtlFilterRegex) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EtlFilterRegex) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EtlFilterRegex) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilterSource

`func (o *EtlFilterRegex) GetFilterSource() FilterSource`

GetFilterSource returns the FilterSource field if non-nil, zero value otherwise.

### GetFilterSourceOk

`func (o *EtlFilterRegex) GetFilterSourceOk() (*FilterSource, bool)`

GetFilterSourceOk returns a tuple with the FilterSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSource

`func (o *EtlFilterRegex) SetFilterSource(v FilterSource)`

SetFilterSource sets FilterSource field to given value.

### HasFilterSource

`func (o *EtlFilterRegex) HasFilterSource() bool`

HasFilterSource returns a boolean if a field has been set.

### GetId

`func (o *EtlFilterRegex) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EtlFilterRegex) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EtlFilterRegex) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EtlFilterRegex) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressPolicyAction

`func (o *EtlFilterRegex) GetIngressPolicyAction() string`

GetIngressPolicyAction returns the IngressPolicyAction field if non-nil, zero value otherwise.

### GetIngressPolicyActionOk

`func (o *EtlFilterRegex) GetIngressPolicyActionOk() (*string, bool)`

GetIngressPolicyActionOk returns a tuple with the IngressPolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPolicyAction

`func (o *EtlFilterRegex) SetIngressPolicyAction(v string)`

SetIngressPolicyAction sets IngressPolicyAction field to given value.

### HasIngressPolicyAction

`func (o *EtlFilterRegex) HasIngressPolicyAction() bool`

HasIngressPolicyAction returns a boolean if a field has been set.

### GetName

`func (o *EtlFilterRegex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlFilterRegex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlFilterRegex) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtlFilterRegex) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EtlFilterRegex) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtlFilterRegex) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtlFilterRegex) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtlFilterRegex) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EtlFilterRegex) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EtlFilterRegex) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EtlFilterRegex) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EtlFilterRegex) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


