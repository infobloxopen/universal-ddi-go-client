# SourceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTimestamp** | Pointer to **string** | Timestamp of the last data event processed from the source, as reported by the on-prem agent. | [optional] 
**LogTypes** | Pointer to [**[]LogType**](LogType.md) | Per-log-type status entries reported by the on-prem agent. | [optional] 
**Status** | Pointer to **string** | Aggregated status of the source on this host, as reported by the on-prem agent. Values are agent-emitted strings and are not part of a stable enum. | [optional] 

## Methods

### NewSourceStatus

`func NewSourceStatus() *SourceStatus`

NewSourceStatus instantiates a new SourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceStatusWithDefaults

`func NewSourceStatusWithDefaults() *SourceStatus`

NewSourceStatusWithDefaults instantiates a new SourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTimestamp

`func (o *SourceStatus) GetDataTimestamp() string`

GetDataTimestamp returns the DataTimestamp field if non-nil, zero value otherwise.

### GetDataTimestampOk

`func (o *SourceStatus) GetDataTimestampOk() (*string, bool)`

GetDataTimestampOk returns a tuple with the DataTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimestamp

`func (o *SourceStatus) SetDataTimestamp(v string)`

SetDataTimestamp sets DataTimestamp field to given value.

### HasDataTimestamp

`func (o *SourceStatus) HasDataTimestamp() bool`

HasDataTimestamp returns a boolean if a field has been set.

### GetLogTypes

`func (o *SourceStatus) GetLogTypes() []LogType`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *SourceStatus) GetLogTypesOk() (*[]LogType, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *SourceStatus) SetLogTypes(v []LogType)`

SetLogTypes sets LogTypes field to given value.

### HasLogTypes

`func (o *SourceStatus) HasLogTypes() bool`

HasLogTypes returns a boolean if a field has been set.

### GetStatus

`func (o *SourceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SourceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SourceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SourceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


