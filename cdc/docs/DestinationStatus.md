# DestinationStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataTimestamp** | Pointer to **string** | Timestamp of the last data event forwarded to the destination, as reported by the on-prem agent. | [optional] 
**LogTypes** | Pointer to [**[]LogType**](LogType.md) | Per-log-type status entries reported by the on-prem agent. | [optional] 
**Status** | Pointer to **string** | Aggregated status of the destination on this host, as reported by the on-prem agent. Values are agent-emitted strings and are not part of a stable enum. | [optional] 

## Methods

### NewDestinationStatus

`func NewDestinationStatus() *DestinationStatus`

NewDestinationStatus instantiates a new DestinationStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationStatusWithDefaults

`func NewDestinationStatusWithDefaults() *DestinationStatus`

NewDestinationStatusWithDefaults instantiates a new DestinationStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataTimestamp

`func (o *DestinationStatus) GetDataTimestamp() string`

GetDataTimestamp returns the DataTimestamp field if non-nil, zero value otherwise.

### GetDataTimestampOk

`func (o *DestinationStatus) GetDataTimestampOk() (*string, bool)`

GetDataTimestampOk returns a tuple with the DataTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTimestamp

`func (o *DestinationStatus) SetDataTimestamp(v string)`

SetDataTimestamp sets DataTimestamp field to given value.

### HasDataTimestamp

`func (o *DestinationStatus) HasDataTimestamp() bool`

HasDataTimestamp returns a boolean if a field has been set.

### GetLogTypes

`func (o *DestinationStatus) GetLogTypes() []LogType`

GetLogTypes returns the LogTypes field if non-nil, zero value otherwise.

### GetLogTypesOk

`func (o *DestinationStatus) GetLogTypesOk() (*[]LogType, bool)`

GetLogTypesOk returns a tuple with the LogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypes

`func (o *DestinationStatus) SetLogTypes(v []LogType)`

SetLogTypes sets LogTypes field to given value.

### HasLogTypes

`func (o *DestinationStatus) HasLogTypes() bool`

HasLogTypes returns a boolean if a field has been set.

### GetStatus

`func (o *DestinationStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DestinationStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DestinationStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DestinationStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


