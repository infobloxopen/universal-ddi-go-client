# PauseSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int64** | The time of the day in minutes that indicates the end of the blackout window. | [optional] 
**StartTime** | Pointer to **int64** | The time of the day in minutes that indicates the start of the blackout window. | [optional] 
**Weekday** | Pointer to **int64** |  | [optional] 

## Methods

### NewPauseSchedule

`func NewPauseSchedule() *PauseSchedule`

NewPauseSchedule instantiates a new PauseSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPauseScheduleWithDefaults

`func NewPauseScheduleWithDefaults() *PauseSchedule`

NewPauseScheduleWithDefaults instantiates a new PauseSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *PauseSchedule) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *PauseSchedule) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *PauseSchedule) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *PauseSchedule) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetStartTime

`func (o *PauseSchedule) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *PauseSchedule) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *PauseSchedule) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *PauseSchedule) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetWeekday

`func (o *PauseSchedule) GetWeekday() int64`

GetWeekday returns the Weekday field if non-nil, zero value otherwise.

### GetWeekdayOk

`func (o *PauseSchedule) GetWeekdayOk() (*int64, bool)`

GetWeekdayOk returns a tuple with the Weekday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekday

`func (o *PauseSchedule) SetWeekday(v int64)`

SetWeekday sets Weekday field to given value.

### HasWeekday

`func (o *PauseSchedule) HasWeekday() bool`

HasWeekday returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


