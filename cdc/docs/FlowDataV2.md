# FlowDataV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdcServices** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**DestinationHa** | Pointer to **bool** |  | [optional] [readonly] 
**Destinations** | Pointer to **[]int64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**EtlFilters** | Pointer to **[]int64** |  | [optional] 
**FilterExpression** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ScheduleOption** | Pointer to **string** |  | [optional] 
**ScriptSchedule** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **int64** |  | [optional] 
**SourceDataTypes** | Pointer to **[]string** |  | [optional] 
**SourceLogTypes** | Pointer to [**[]SourceLogType**](SourceLogType.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewFlowDataV2

`func NewFlowDataV2() *FlowDataV2`

NewFlowDataV2 instantiates a new FlowDataV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDataV2WithDefaults

`func NewFlowDataV2WithDefaults() *FlowDataV2`

NewFlowDataV2WithDefaults instantiates a new FlowDataV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdcServices

`func (o *FlowDataV2) GetCdcServices() []string`

GetCdcServices returns the CdcServices field if non-nil, zero value otherwise.

### GetCdcServicesOk

`func (o *FlowDataV2) GetCdcServicesOk() (*[]string, bool)`

GetCdcServicesOk returns a tuple with the CdcServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcServices

`func (o *FlowDataV2) SetCdcServices(v []string)`

SetCdcServices sets CdcServices field to given value.

### HasCdcServices

`func (o *FlowDataV2) HasCdcServices() bool`

HasCdcServices returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FlowDataV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlowDataV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlowDataV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FlowDataV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FlowDataV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowDataV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowDataV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowDataV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationHa

`func (o *FlowDataV2) GetDestinationHa() bool`

GetDestinationHa returns the DestinationHa field if non-nil, zero value otherwise.

### GetDestinationHaOk

`func (o *FlowDataV2) GetDestinationHaOk() (*bool, bool)`

GetDestinationHaOk returns a tuple with the DestinationHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHa

`func (o *FlowDataV2) SetDestinationHa(v bool)`

SetDestinationHa sets DestinationHa field to given value.

### HasDestinationHa

`func (o *FlowDataV2) HasDestinationHa() bool`

HasDestinationHa returns a boolean if a field has been set.

### GetDestinations

`func (o *FlowDataV2) GetDestinations() []int64`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *FlowDataV2) GetDestinationsOk() (*[]int64, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *FlowDataV2) SetDestinations(v []int64)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *FlowDataV2) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetEnabled

`func (o *FlowDataV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlowDataV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlowDataV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlowDataV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEtlFilters

`func (o *FlowDataV2) GetEtlFilters() []int64`

GetEtlFilters returns the EtlFilters field if non-nil, zero value otherwise.

### GetEtlFiltersOk

`func (o *FlowDataV2) GetEtlFiltersOk() (*[]int64, bool)`

GetEtlFiltersOk returns a tuple with the EtlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFilters

`func (o *FlowDataV2) SetEtlFilters(v []int64)`

SetEtlFilters sets EtlFilters field to given value.

### HasEtlFilters

`func (o *FlowDataV2) HasEtlFilters() bool`

HasEtlFilters returns a boolean if a field has been set.

### GetFilterExpression

`func (o *FlowDataV2) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *FlowDataV2) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *FlowDataV2) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *FlowDataV2) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetId

`func (o *FlowDataV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowDataV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowDataV2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FlowDataV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowDataV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowDataV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowDataV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowDataV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleOption

`func (o *FlowDataV2) GetScheduleOption() string`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *FlowDataV2) GetScheduleOptionOk() (*string, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *FlowDataV2) SetScheduleOption(v string)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *FlowDataV2) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetScriptSchedule

`func (o *FlowDataV2) GetScriptSchedule() string`

GetScriptSchedule returns the ScriptSchedule field if non-nil, zero value otherwise.

### GetScriptScheduleOk

`func (o *FlowDataV2) GetScriptScheduleOk() (*string, bool)`

GetScriptScheduleOk returns a tuple with the ScriptSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSchedule

`func (o *FlowDataV2) SetScriptSchedule(v string)`

SetScriptSchedule sets ScriptSchedule field to given value.

### HasScriptSchedule

`func (o *FlowDataV2) HasScriptSchedule() bool`

HasScriptSchedule returns a boolean if a field has been set.

### GetSource

`func (o *FlowDataV2) GetSource() int64`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FlowDataV2) GetSourceOk() (*int64, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FlowDataV2) SetSource(v int64)`

SetSource sets Source field to given value.

### HasSource

`func (o *FlowDataV2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceDataTypes

`func (o *FlowDataV2) GetSourceDataTypes() []string`

GetSourceDataTypes returns the SourceDataTypes field if non-nil, zero value otherwise.

### GetSourceDataTypesOk

`func (o *FlowDataV2) GetSourceDataTypesOk() (*[]string, bool)`

GetSourceDataTypesOk returns a tuple with the SourceDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataTypes

`func (o *FlowDataV2) SetSourceDataTypes(v []string)`

SetSourceDataTypes sets SourceDataTypes field to given value.

### HasSourceDataTypes

`func (o *FlowDataV2) HasSourceDataTypes() bool`

HasSourceDataTypes returns a boolean if a field has been set.

### GetSourceLogTypes

`func (o *FlowDataV2) GetSourceLogTypes() []SourceLogType`

GetSourceLogTypes returns the SourceLogTypes field if non-nil, zero value otherwise.

### GetSourceLogTypesOk

`func (o *FlowDataV2) GetSourceLogTypesOk() (*[]SourceLogType, bool)`

GetSourceLogTypesOk returns a tuple with the SourceLogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLogTypes

`func (o *FlowDataV2) SetSourceLogTypes(v []SourceLogType)`

SetSourceLogTypes sets SourceLogTypes field to given value.

### HasSourceLogTypes

`func (o *FlowDataV2) HasSourceLogTypes() bool`

HasSourceLogTypes returns a boolean if a field has been set.

### GetTags

`func (o *FlowDataV2) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlowDataV2) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlowDataV2) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlowDataV2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FlowDataV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlowDataV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlowDataV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FlowDataV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


