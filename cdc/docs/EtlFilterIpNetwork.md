# EtlFilterIpNetwork

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

### NewEtlFilterIpNetwork

`func NewEtlFilterIpNetwork() *EtlFilterIpNetwork`

NewEtlFilterIpNetwork instantiates a new EtlFilterIpNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlFilterIpNetworkWithDefaults

`func NewEtlFilterIpNetworkWithDefaults() *EtlFilterIpNetwork`

NewEtlFilterIpNetworkWithDefaults instantiates a new EtlFilterIpNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EtlFilterIpNetwork) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EtlFilterIpNetwork) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EtlFilterIpNetwork) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EtlFilterIpNetwork) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *EtlFilterIpNetwork) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EtlFilterIpNetwork) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EtlFilterIpNetwork) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *EtlFilterIpNetwork) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *EtlFilterIpNetwork) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EtlFilterIpNetwork) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EtlFilterIpNetwork) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EtlFilterIpNetwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *EtlFilterIpNetwork) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EtlFilterIpNetwork) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EtlFilterIpNetwork) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EtlFilterIpNetwork) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilterSource

`func (o *EtlFilterIpNetwork) GetFilterSource() FilterSource`

GetFilterSource returns the FilterSource field if non-nil, zero value otherwise.

### GetFilterSourceOk

`func (o *EtlFilterIpNetwork) GetFilterSourceOk() (*FilterSource, bool)`

GetFilterSourceOk returns a tuple with the FilterSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSource

`func (o *EtlFilterIpNetwork) SetFilterSource(v FilterSource)`

SetFilterSource sets FilterSource field to given value.

### HasFilterSource

`func (o *EtlFilterIpNetwork) HasFilterSource() bool`

HasFilterSource returns a boolean if a field has been set.

### GetId

`func (o *EtlFilterIpNetwork) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EtlFilterIpNetwork) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EtlFilterIpNetwork) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *EtlFilterIpNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIngressPolicyAction

`func (o *EtlFilterIpNetwork) GetIngressPolicyAction() string`

GetIngressPolicyAction returns the IngressPolicyAction field if non-nil, zero value otherwise.

### GetIngressPolicyActionOk

`func (o *EtlFilterIpNetwork) GetIngressPolicyActionOk() (*string, bool)`

GetIngressPolicyActionOk returns a tuple with the IngressPolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPolicyAction

`func (o *EtlFilterIpNetwork) SetIngressPolicyAction(v string)`

SetIngressPolicyAction sets IngressPolicyAction field to given value.

### HasIngressPolicyAction

`func (o *EtlFilterIpNetwork) HasIngressPolicyAction() bool`

HasIngressPolicyAction returns a boolean if a field has been set.

### GetName

`func (o *EtlFilterIpNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EtlFilterIpNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EtlFilterIpNetwork) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EtlFilterIpNetwork) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *EtlFilterIpNetwork) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EtlFilterIpNetwork) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EtlFilterIpNetwork) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *EtlFilterIpNetwork) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EtlFilterIpNetwork) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EtlFilterIpNetwork) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EtlFilterIpNetwork) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EtlFilterIpNetwork) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


