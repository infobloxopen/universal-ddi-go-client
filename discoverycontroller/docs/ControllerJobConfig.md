# ControllerJobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**DiscoverNow** | Pointer to **bool** | Flag to indicate on-demand discovery scan when Configuration is disabled. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Required. On-Prem Host ID on which the job is to run. | [optional] 
**PauseSchedules** | Pointer to [**[]ControllerPauseSchedule**](ControllerPauseSchedule.md) | List of pause schedules. | [optional] 
**PoolId** | Pointer to **string** | Unique Identifier of the pool to which the service and host belongs. | [optional] 
**ServiceId** | Pointer to **string** | Required. Unique Identifier of the discovery service associated to this job config. | [optional] 
**ServiceName** | Pointer to **string** | Read-only. Name of the discovery service. | [optional] [readonly] 
**SourceOptions** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewControllerJobConfig

`func NewControllerJobConfig() *ControllerJobConfig`

NewControllerJobConfig instantiates a new ControllerJobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerJobConfigWithDefaults

`func NewControllerJobConfigWithDefaults() *ControllerJobConfig`

NewControllerJobConfigWithDefaults instantiates a new ControllerJobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *ControllerJobConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ControllerJobConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ControllerJobConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ControllerJobConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControllerJobConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerJobConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerJobConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerJobConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDiscoverNow

`func (o *ControllerJobConfig) GetDiscoverNow() bool`

GetDiscoverNow returns the DiscoverNow field if non-nil, zero value otherwise.

### GetDiscoverNowOk

`func (o *ControllerJobConfig) GetDiscoverNowOk() (*bool, bool)`

GetDiscoverNowOk returns a tuple with the DiscoverNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverNow

`func (o *ControllerJobConfig) SetDiscoverNow(v bool)`

SetDiscoverNow sets DiscoverNow field to given value.

### HasDiscoverNow

`func (o *ControllerJobConfig) HasDiscoverNow() bool`

HasDiscoverNow returns a boolean if a field has been set.

### GetId

`func (o *ControllerJobConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerJobConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerJobConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerJobConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOphid

`func (o *ControllerJobConfig) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ControllerJobConfig) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ControllerJobConfig) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ControllerJobConfig) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPauseSchedules

`func (o *ControllerJobConfig) GetPauseSchedules() []ControllerPauseSchedule`

GetPauseSchedules returns the PauseSchedules field if non-nil, zero value otherwise.

### GetPauseSchedulesOk

`func (o *ControllerJobConfig) GetPauseSchedulesOk() (*[]ControllerPauseSchedule, bool)`

GetPauseSchedulesOk returns a tuple with the PauseSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseSchedules

`func (o *ControllerJobConfig) SetPauseSchedules(v []ControllerPauseSchedule)`

SetPauseSchedules sets PauseSchedules field to given value.

### HasPauseSchedules

`func (o *ControllerJobConfig) HasPauseSchedules() bool`

HasPauseSchedules returns a boolean if a field has been set.

### GetPoolId

`func (o *ControllerJobConfig) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ControllerJobConfig) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ControllerJobConfig) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ControllerJobConfig) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *ControllerJobConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ControllerJobConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ControllerJobConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ControllerJobConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ControllerJobConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ControllerJobConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ControllerJobConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ControllerJobConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSourceOptions

`func (o *ControllerJobConfig) GetSourceOptions() map[string]interface{}`

GetSourceOptions returns the SourceOptions field if non-nil, zero value otherwise.

### GetSourceOptionsOk

`func (o *ControllerJobConfig) GetSourceOptionsOk() (*map[string]interface{}, bool)`

GetSourceOptionsOk returns a tuple with the SourceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOptions

`func (o *ControllerJobConfig) SetSourceOptions(v map[string]interface{})`

SetSourceOptions sets SourceOptions field to given value.

### HasSourceOptions

`func (o *ControllerJobConfig) HasSourceOptions() bool`

HasSourceOptions returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerJobConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerJobConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerJobConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerJobConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


