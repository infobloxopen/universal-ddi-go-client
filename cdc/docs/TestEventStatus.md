# TestEventStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogType** | Pointer to **string** | Log type against which the test event was executed. | [optional] 
**Message** | Pointer to **string** | Human-readable message describing the test event outcome. | [optional] 
**Status** | Pointer to **string** | Outcome of the most recent test event triggered for the flow. Agent-emitted string; not a stable enum. | [optional] 
**Timestamp** | Pointer to [**StrTimestamp**](StrTimestamp.md) | Timestamp at which the test event was recorded. | [optional] 

## Methods

### NewTestEventStatus

`func NewTestEventStatus() *TestEventStatus`

NewTestEventStatus instantiates a new TestEventStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestEventStatusWithDefaults

`func NewTestEventStatusWithDefaults() *TestEventStatus`

NewTestEventStatusWithDefaults instantiates a new TestEventStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogType

`func (o *TestEventStatus) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *TestEventStatus) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *TestEventStatus) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *TestEventStatus) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetMessage

`func (o *TestEventStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TestEventStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TestEventStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TestEventStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *TestEventStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestEventStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestEventStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestEventStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *TestEventStatus) GetTimestamp() StrTimestamp`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TestEventStatus) GetTimestampOk() (*StrTimestamp, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TestEventStatus) SetTimestamp(v StrTimestamp)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TestEventStatus) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


