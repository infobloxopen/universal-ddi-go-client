# JobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**DiscoverNow** | Pointer to **bool** | Flag to indicate on-demand discovery scan when Configuration is disabled. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Required. On-Prem Host ID on which the job is to run. | [optional] 
**PauseSchedules** | Pointer to [**[]PauseSchedule**](PauseSchedule.md) | List of pause schedules. | [optional] 
**PoolId** | Pointer to **string** | Unique Identifier of the pool to which the service and host belongs. | [optional] 
**ServiceId** | Pointer to **string** | Required. Unique Identifier of the discovery service associated to this job config. | [optional] 
**ServiceName** | Pointer to **string** | Read-only. Name of the discovery service. | [optional] [readonly] 
**SourceOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewJobConfig

`func NewJobConfig() *JobConfig`

NewJobConfig instantiates a new JobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobConfigWithDefaults

`func NewJobConfigWithDefaults() *JobConfig`

NewJobConfigWithDefaults instantiates a new JobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *JobConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *JobConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *JobConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *JobConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *JobConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *JobConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDiscoverNow

`func (o *JobConfig) GetDiscoverNow() bool`

GetDiscoverNow returns the DiscoverNow field if non-nil, zero value otherwise.

### GetDiscoverNowOk

`func (o *JobConfig) GetDiscoverNowOk() (*bool, bool)`

GetDiscoverNowOk returns a tuple with the DiscoverNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNow

`func (o *JobConfig) SetDiscoverNow(v bool)`

SetDiscoverNow sets DiscoverNow field to given value.

### HasDiscoverNow

`func (o *JobConfig) HasDiscoverNow() bool`

HasDiscoverNow returns a boolean if a field has been set.

### GetId

`func (o *JobConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *JobConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOphid

`func (o *JobConfig) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *JobConfig) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *JobConfig) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *JobConfig) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPauseSchedules

`func (o *JobConfig) GetPauseSchedules() []PauseSchedule`

GetPauseSchedules returns the PauseSchedules field if non-nil, zero value otherwise.

### GetPauseSchedulesOk

`func (o *JobConfig) GetPauseSchedulesOk() (*[]PauseSchedule, bool)`

GetPauseSchedulesOk returns a tuple with the PauseSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseSchedules

`func (o *JobConfig) SetPauseSchedules(v []PauseSchedule)`

SetPauseSchedules sets PauseSchedules field to given value.

### HasPauseSchedules

`func (o *JobConfig) HasPauseSchedules() bool`

HasPauseSchedules returns a boolean if a field has been set.

### GetPoolId

`func (o *JobConfig) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *JobConfig) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *JobConfig) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *JobConfig) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *JobConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *JobConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *JobConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *JobConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *JobConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *JobConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *JobConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *JobConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSourceOptions

`func (o *JobConfig) GetSourceOptions() map[string]interface{}`

GetSourceOptions returns the SourceOptions field if non-nil, zero value otherwise.

### GetSourceOptionsOk

`func (o *JobConfig) GetSourceOptionsOk() (*map[string]interface{}, bool)`

GetSourceOptionsOk returns a tuple with the SourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOptions

`func (o *JobConfig) SetSourceOptions(v map[string]interface{})`

SetSourceOptions sets SourceOptions field to given value.

### HasSourceOptions

`func (o *JobConfig) HasSourceOptions() bool`

HasSourceOptions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *JobConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JobConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JobConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *JobConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


