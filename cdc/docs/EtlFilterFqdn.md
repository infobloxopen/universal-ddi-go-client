# EtlFilterFqdn

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

### NewEtlFilterFqdn

`func NewEtlFilterFqdn() *EtlFilterFqdn`

NewEtlFilterFqdn instantiates a new EtlFilterFqdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlFilterFqdnWithDefaults

`func NewEtlFilterFqdnWithDefaults() *EtlFilterFqdn`

NewEtlFilterFqdnWithDefaults instantiates a new EtlFilterFqdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EtlFilterFqdn) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EtlFilterFqdn) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EtlFilterFqdn) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EtlFilterFqdn) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EtlFilterFqdn) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EtlFilterFqdn) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EtlFilterFqdn) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *EtlFilterFqdn) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *EtlFilterFqdn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtlFilterFqdn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtlFilterFqdn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtlFilterFqdn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EtlFilterFqdn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EtlFilterFqdn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EtlFilterFqdn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EtlFilterFqdn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilterSource

`func (o *EtlFilterFqdn) GetFilterSource() FilterSource`

GetFilterSource returns the FilterSource field if non-nil, zero value otherwise.

### GetFilterSourceOk

`func (o *EtlFilterFqdn) GetFilterSourceOk() (*FilterSource, bool)`

GetFilterSourceOk returns a tuple with the FilterSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSource

`func (o *EtlFilterFqdn) SetFilterSource(v FilterSource)`

SetFilterSource sets FilterSource field to given value.

### HasFilterSource

`func (o *EtlFilterFqdn) HasFilterSource() bool`

HasFilterSource returns a boolean if a field has been set.

### GetId

`func (o *EtlFilterFqdn) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EtlFilterFqdn) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EtlFilterFqdn) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EtlFilterFqdn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressPolicyAction

`func (o *EtlFilterFqdn) GetIngressPolicyAction() string`

GetIngressPolicyAction returns the IngressPolicyAction field if non-nil, zero value otherwise.

### GetIngressPolicyActionOk

`func (o *EtlFilterFqdn) GetIngressPolicyActionOk() (*string, bool)`

GetIngressPolicyActionOk returns a tuple with the IngressPolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPolicyAction

`func (o *EtlFilterFqdn) SetIngressPolicyAction(v string)`

SetIngressPolicyAction sets IngressPolicyAction field to given value.

### HasIngressPolicyAction

`func (o *EtlFilterFqdn) HasIngressPolicyAction() bool`

HasIngressPolicyAction returns a boolean if a field has been set.

### GetName

`func (o *EtlFilterFqdn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlFilterFqdn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlFilterFqdn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtlFilterFqdn) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EtlFilterFqdn) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtlFilterFqdn) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtlFilterFqdn) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtlFilterFqdn) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EtlFilterFqdn) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EtlFilterFqdn) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EtlFilterFqdn) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EtlFilterFqdn) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


