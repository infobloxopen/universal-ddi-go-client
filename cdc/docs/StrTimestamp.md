# StrTimestamp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Seconds** | Pointer to **string** | Timestamp expressed as seconds since Unix epoch, encoded as a string. | [optional] 

## Methods

### NewStrTimestamp

`func NewStrTimestamp() *StrTimestamp`

NewStrTimestamp instantiates a new StrTimestamp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrTimestampWithDefaults

`func NewStrTimestampWithDefaults() *StrTimestamp`

NewStrTimestampWithDefaults instantiates a new StrTimestamp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeconds

`func (o *StrTimestamp) GetSeconds() string`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *StrTimestamp) GetSecondsOk() (*string, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *StrTimestamp) SetSeconds(v string)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *StrTimestamp) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


