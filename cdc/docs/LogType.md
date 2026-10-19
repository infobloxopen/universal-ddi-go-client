# LogType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human-readable status message for the log type. | [optional] 
**Type** | Pointer to **string** | Log type identifier (source data type) that this entry refers to. | [optional] 

## Methods

### NewLogType

`func NewLogType() *LogType`

NewLogType instantiates a new LogType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogTypeWithDefaults

`func NewLogTypeWithDefaults() *LogType`

NewLogTypeWithDefaults instantiates a new LogType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *LogType) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogType) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogType) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogType) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *LogType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LogType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LogType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LogType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


