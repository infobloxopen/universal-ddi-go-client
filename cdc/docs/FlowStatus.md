# FlowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventStatus** | Pointer to [**TestEventStatus**](TestEventStatus.md) | Status of the most recent test event triggered for this flow, if any. | [optional] 
**Host** | Pointer to [**[]HostStatus**](HostStatus.md) | Per-host runtime status entries for this flow. | [optional] 
**Status** | Pointer to **string** | Aggregated flow status across all hosts, as computed by the service. Agent-emitted vocabulary; not a stable enum. | [optional] 
**Timestamp** | Pointer to **time.Time** | Time at which the aggregated status was computed. | [optional] 

## Methods

### NewFlowStatus

`func NewFlowStatus() *FlowStatus`

NewFlowStatus instantiates a new FlowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowStatusWithDefaults

`func NewFlowStatusWithDefaults() *FlowStatus`

NewFlowStatusWithDefaults instantiates a new FlowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventStatus

`func (o *FlowStatus) GetEventStatus() TestEventStatus`

GetEventStatus returns the EventStatus field if non-nil, zero value otherwise.

### GetEventStatusOk

`func (o *FlowStatus) GetEventStatusOk() (*TestEventStatus, bool)`

GetEventStatusOk returns a tuple with the EventStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventStatus

`func (o *FlowStatus) SetEventStatus(v TestEventStatus)`

SetEventStatus sets EventStatus field to given value.

### HasEventStatus

`func (o *FlowStatus) HasEventStatus() bool`

HasEventStatus returns a boolean if a field has been set.

### GetHost

`func (o *FlowStatus) GetHost() []HostStatus`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *FlowStatus) GetHostOk() (*[]HostStatus, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *FlowStatus) SetHost(v []HostStatus)`

SetHost sets Host field to given value.

### HasHost

`func (o *FlowStatus) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetStatus

`func (o *FlowStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FlowStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FlowStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FlowStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *FlowStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FlowStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FlowStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FlowStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


