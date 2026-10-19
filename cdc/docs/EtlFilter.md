# EtlFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Data** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IngressPolicyAction** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEtlFilter

`func NewEtlFilter() *EtlFilter`

NewEtlFilter instantiates a new EtlFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlFilterWithDefaults

`func NewEtlFilterWithDefaults() *EtlFilter`

NewEtlFilterWithDefaults instantiates a new EtlFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EtlFilter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EtlFilter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EtlFilter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EtlFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EtlFilter) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EtlFilter) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EtlFilter) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *EtlFilter) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *EtlFilter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtlFilter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtlFilter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtlFilter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EtlFilter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EtlFilter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EtlFilter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EtlFilter) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *EtlFilter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EtlFilter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EtlFilter) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EtlFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressPolicyAction

`func (o *EtlFilter) GetIngressPolicyAction() string`

GetIngressPolicyAction returns the IngressPolicyAction field if non-nil, zero value otherwise.

### GetIngressPolicyActionOk

`func (o *EtlFilter) GetIngressPolicyActionOk() (*string, bool)`

GetIngressPolicyActionOk returns a tuple with the IngressPolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPolicyAction

`func (o *EtlFilter) SetIngressPolicyAction(v string)`

SetIngressPolicyAction sets IngressPolicyAction field to given value.

### HasIngressPolicyAction

`func (o *EtlFilter) HasIngressPolicyAction() bool`

HasIngressPolicyAction returns a boolean if a field has been set.

### GetName

`func (o *EtlFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlFilter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtlFilter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EtlFilter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtlFilter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtlFilter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtlFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *EtlFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EtlFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EtlFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EtlFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EtlFilter) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EtlFilter) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EtlFilter) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EtlFilter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


