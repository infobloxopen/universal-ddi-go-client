# ControllerPauseSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int32** | The time of the day in minutes that indicates the end of the blackout window. | [optional] 
**StartTime** | Pointer to **int32** | The time of the day in minutes that indicates the start of the blackout window. | [optional] 
**Weekday** | Pointer to **int32** |  | [optional] 

## Methods

### NewControllerPauseSchedule

`func NewControllerPauseSchedule() *ControllerPauseSchedule`

NewControllerPauseSchedule instantiates a new ControllerPauseSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerPauseScheduleWithDefaults

`func NewControllerPauseScheduleWithDefaults() *ControllerPauseSchedule`

NewControllerPauseScheduleWithDefaults instantiates a new ControllerPauseSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ControllerPauseSchedule) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ControllerPauseSchedule) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ControllerPauseSchedule) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ControllerPauseSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ControllerPauseSchedule) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ControllerPauseSchedule) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ControllerPauseSchedule) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ControllerPauseSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWeekday

`func (o *ControllerPauseSchedule) GetWeekday() int32`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *ControllerPauseSchedule) GetWeekdayOk() (*int32, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *ControllerPauseSchedule) SetWeekday(v int32)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *ControllerPauseSchedule) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


