# FlowV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CdcHostNames** | Pointer to **[]string** | Resolved names of the on-prem hosts referenced by &#39;cdc_hosts&#39;. Read-only; not filterable. | [optional] 
**CdcHosts** | Pointer to **[]int64** | Identifiers (ophids) of the on-prem hosts on which this flow is deployed. | [optional] 
**CdcServiceNames** | Pointer to **[]string** | Resolved names of the CDC services referenced by &#39;cdc_services&#39;. Read-only; not filterable. | [optional] 
**CdcServices** | Pointer to **[]string** | Identifiers of the on-prem CDC services carrying this flow. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp at which the flow was created. | [optional] [readonly] 
**Description** | Pointer to **string** | Free-form description of the flow. | [optional] 
**DestinationHa** | Pointer to **bool** | Whether destination high availability is enabled for this flow. When true, one host acts as active and another as standby (see HostStatus.ha_state). | [optional] 
**Destinations** | Pointer to [**[]FlowDestination**](FlowDestination.md) | Destinations attached to the flow. Each entry carries the destination id and its resolved name, type, transport and output format. Not filterable. | [optional] 
**Enabled** | Pointer to **bool** | Whether the flow is enabled. Disabled flows are configured but do not process data. | [optional] 
**EtlFilterNames** | Pointer to **[]string** | Resolved names of the ETL filters referenced by &#39;etl_filters&#39;. Read-only; not filterable. | [optional] 
**EtlFilters** | Pointer to **[]int64** | Identifiers of the ETL filters attached to this flow. | [optional] 
**FilterExpression** | Pointer to **string** | Optional filter expression applied to the flow. Not filterable. | [optional] 
**Id** | Pointer to **int64** | Numeric identifier of the flow. | [optional] [readonly] 
**Name** | Pointer to **string** | Human-readable name of the flow. | [optional] 
**ScheduleOption** | Pointer to **string** | Named schedule option applied to the flow. | [optional] 
**ScriptSchedule** | Pointer to **string** | Optional cron-style schedule applied to script-based flows. Not filterable. | [optional] 
**Source** | Pointer to **int64** | Numeric identifier of the source resource attached to this flow. | [optional] 
**SourceDataTypes** | Pointer to **[]string** | Source data type identifiers selected on the source for this flow. | [optional] 
**SourceEnabled** | Pointer to **bool** | Resolved &#39;enabled&#39; state of the referenced source. Read-only. | [optional] 
**SourceLogTypes** | Pointer to [**[]SourceLogType**](SourceLogType.md) | Source log types (with optional subtypes) selected for this flow. | [optional] 
**SourceName** | Pointer to **string** | Resolved name of the source referenced by &#39;source&#39;. Read-only; not filterable. | [optional] 
**SourceType** | Pointer to **string** | Resolved type of the source referenced by &#39;source&#39; (e.g. bloxone, nios, schedule). Read-only; not filterable. | [optional] 
**Status** | Pointer to [**FlowStatus**](FlowStatus.md) | Live runtime status of the flow aggregated across all hosts on which it is deployed. Values are agent-emitted and are not part of a stable enum. | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to **string** | Category of the flow, derived from the underlying source type. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp at which the flow was last updated. | [optional] [readonly] 

## Methods

### NewFlowV2

`func NewFlowV2() *FlowV2`

NewFlowV2 instantiates a new FlowV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowV2WithDefaults

`func NewFlowV2WithDefaults() *FlowV2`

NewFlowV2WithDefaults instantiates a new FlowV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdcHostNames

`func (o *FlowV2) GetCdcHostNames() []string`

GetCdcHostNames returns the CdcHostNames field if non-nil, zero value otherwise.

### GetCdcHostNamesOk

`func (o *FlowV2) GetCdcHostNamesOk() (*[]string, bool)`

GetCdcHostNamesOk returns a tuple with the CdcHostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcHostNames

`func (o *FlowV2) SetCdcHostNames(v []string)`

SetCdcHostNames sets CdcHostNames field to given value.

### HasCdcHostNames

`func (o *FlowV2) HasCdcHostNames() bool`

HasCdcHostNames returns a boolean if a field has been set.

### GetCdcHosts

`func (o *FlowV2) GetCdcHosts() []int64`

GetCdcHosts returns the CdcHosts field if non-nil, zero value otherwise.

### GetCdcHostsOk

`func (o *FlowV2) GetCdcHostsOk() (*[]int64, bool)`

GetCdcHostsOk returns a tuple with the CdcHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcHosts

`func (o *FlowV2) SetCdcHosts(v []int64)`

SetCdcHosts sets CdcHosts field to given value.

### HasCdcHosts

`func (o *FlowV2) HasCdcHosts() bool`

HasCdcHosts returns a boolean if a field has been set.

### GetCdcServiceNames

`func (o *FlowV2) GetCdcServiceNames() []string`

GetCdcServiceNames returns the CdcServiceNames field if non-nil, zero value otherwise.

### GetCdcServiceNamesOk

`func (o *FlowV2) GetCdcServiceNamesOk() (*[]string, bool)`

GetCdcServiceNamesOk returns a tuple with the CdcServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcServiceNames

`func (o *FlowV2) SetCdcServiceNames(v []string)`

SetCdcServiceNames sets CdcServiceNames field to given value.

### HasCdcServiceNames

`func (o *FlowV2) HasCdcServiceNames() bool`

HasCdcServiceNames returns a boolean if a field has been set.

### GetCdcServices

`func (o *FlowV2) GetCdcServices() []string`

GetCdcServices returns the CdcServices field if non-nil, zero value otherwise.

### GetCdcServicesOk

`func (o *FlowV2) GetCdcServicesOk() (*[]string, bool)`

GetCdcServicesOk returns a tuple with the CdcServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdcServices

`func (o *FlowV2) SetCdcServices(v []string)`

SetCdcServices sets CdcServices field to given value.

### HasCdcServices

`func (o *FlowV2) HasCdcServices() bool`

HasCdcServices returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FlowV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlowV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlowV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FlowV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FlowV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FlowV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FlowV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FlowV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestinationHa

`func (o *FlowV2) GetDestinationHa() bool`

GetDestinationHa returns the DestinationHa field if non-nil, zero value otherwise.

### GetDestinationHaOk

`func (o *FlowV2) GetDestinationHaOk() (*bool, bool)`

GetDestinationHaOk returns a tuple with the DestinationHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHa

`func (o *FlowV2) SetDestinationHa(v bool)`

SetDestinationHa sets DestinationHa field to given value.

### HasDestinationHa

`func (o *FlowV2) HasDestinationHa() bool`

HasDestinationHa returns a boolean if a field has been set.

### GetDestinations

`func (o *FlowV2) GetDestinations() []FlowDestination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *FlowV2) GetDestinationsOk() (*[]FlowDestination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *FlowV2) SetDestinations(v []FlowDestination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *FlowV2) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetEnabled

`func (o *FlowV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlowV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlowV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlowV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEtlFilterNames

`func (o *FlowV2) GetEtlFilterNames() []string`

GetEtlFilterNames returns the EtlFilterNames field if non-nil, zero value otherwise.

### GetEtlFilterNamesOk

`func (o *FlowV2) GetEtlFilterNamesOk() (*[]string, bool)`

GetEtlFilterNamesOk returns a tuple with the EtlFilterNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFilterNames

`func (o *FlowV2) SetEtlFilterNames(v []string)`

SetEtlFilterNames sets EtlFilterNames field to given value.

### HasEtlFilterNames

`func (o *FlowV2) HasEtlFilterNames() bool`

HasEtlFilterNames returns a boolean if a field has been set.

### GetEtlFilters

`func (o *FlowV2) GetEtlFilters() []int64`

GetEtlFilters returns the EtlFilters field if non-nil, zero value otherwise.

### GetEtlFiltersOk

`func (o *FlowV2) GetEtlFiltersOk() (*[]int64, bool)`

GetEtlFiltersOk returns a tuple with the EtlFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtlFilters

`func (o *FlowV2) SetEtlFilters(v []int64)`

SetEtlFilters sets EtlFilters field to given value.

### HasEtlFilters

`func (o *FlowV2) HasEtlFilters() bool`

HasEtlFilters returns a boolean if a field has been set.

### GetFilterExpression

`func (o *FlowV2) GetFilterExpression() string`

GetFilterExpression returns the FilterExpression field if non-nil, zero value otherwise.

### GetFilterExpressionOk

`func (o *FlowV2) GetFilterExpressionOk() (*string, bool)`

GetFilterExpressionOk returns a tuple with the FilterExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterExpression

`func (o *FlowV2) SetFilterExpression(v string)`

SetFilterExpression sets FilterExpression field to given value.

### HasFilterExpression

`func (o *FlowV2) HasFilterExpression() bool`

HasFilterExpression returns a boolean if a field has been set.

### GetId

`func (o *FlowV2) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowV2) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowV2) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FlowV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduleOption

`func (o *FlowV2) GetScheduleOption() string`

GetScheduleOption returns the ScheduleOption field if non-nil, zero value otherwise.

### GetScheduleOptionOk

`func (o *FlowV2) GetScheduleOptionOk() (*string, bool)`

GetScheduleOptionOk returns a tuple with the ScheduleOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleOption

`func (o *FlowV2) SetScheduleOption(v string)`

SetScheduleOption sets ScheduleOption field to given value.

### HasScheduleOption

`func (o *FlowV2) HasScheduleOption() bool`

HasScheduleOption returns a boolean if a field has been set.

### GetScriptSchedule

`func (o *FlowV2) GetScriptSchedule() string`

GetScriptSchedule returns the ScriptSchedule field if non-nil, zero value otherwise.

### GetScriptScheduleOk

`func (o *FlowV2) GetScriptScheduleOk() (*string, bool)`

GetScriptScheduleOk returns a tuple with the ScriptSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSchedule

`func (o *FlowV2) SetScriptSchedule(v string)`

SetScriptSchedule sets ScriptSchedule field to given value.

### HasScriptSchedule

`func (o *FlowV2) HasScriptSchedule() bool`

HasScriptSchedule returns a boolean if a field has been set.

### GetSource

`func (o *FlowV2) GetSource() int64`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FlowV2) GetSourceOk() (*int64, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FlowV2) SetSource(v int64)`

SetSource sets Source field to given value.

### HasSource

`func (o *FlowV2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSourceDataTypes

`func (o *FlowV2) GetSourceDataTypes() []string`

GetSourceDataTypes returns the SourceDataTypes field if non-nil, zero value otherwise.

### GetSourceDataTypesOk

`func (o *FlowV2) GetSourceDataTypesOk() (*[]string, bool)`

GetSourceDataTypesOk returns a tuple with the SourceDataTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataTypes

`func (o *FlowV2) SetSourceDataTypes(v []string)`

SetSourceDataTypes sets SourceDataTypes field to given value.

### HasSourceDataTypes

`func (o *FlowV2) HasSourceDataTypes() bool`

HasSourceDataTypes returns a boolean if a field has been set.

### GetSourceEnabled

`func (o *FlowV2) GetSourceEnabled() bool`

GetSourceEnabled returns the SourceEnabled field if non-nil, zero value otherwise.

### GetSourceEnabledOk

`func (o *FlowV2) GetSourceEnabledOk() (*bool, bool)`

GetSourceEnabledOk returns a tuple with the SourceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEnabled

`func (o *FlowV2) SetSourceEnabled(v bool)`

SetSourceEnabled sets SourceEnabled field to given value.

### HasSourceEnabled

`func (o *FlowV2) HasSourceEnabled() bool`

HasSourceEnabled returns a boolean if a field has been set.

### GetSourceLogTypes

`func (o *FlowV2) GetSourceLogTypes() []SourceLogType`

GetSourceLogTypes returns the SourceLogTypes field if non-nil, zero value otherwise.

### GetSourceLogTypesOk

`func (o *FlowV2) GetSourceLogTypesOk() (*[]SourceLogType, bool)`

GetSourceLogTypesOk returns a tuple with the SourceLogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLogTypes

`func (o *FlowV2) SetSourceLogTypes(v []SourceLogType)`

SetSourceLogTypes sets SourceLogTypes field to given value.

### HasSourceLogTypes

`func (o *FlowV2) HasSourceLogTypes() bool`

HasSourceLogTypes returns a boolean if a field has been set.

### GetSourceName

`func (o *FlowV2) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FlowV2) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FlowV2) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *FlowV2) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceType

`func (o *FlowV2) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *FlowV2) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *FlowV2) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *FlowV2) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetStatus

`func (o *FlowV2) GetStatus() FlowStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowV2) GetStatusOk() (*FlowStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowV2) SetStatus(v FlowStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *FlowV2) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FlowV2) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FlowV2) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FlowV2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *FlowV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FlowV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlowV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlowV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FlowV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


