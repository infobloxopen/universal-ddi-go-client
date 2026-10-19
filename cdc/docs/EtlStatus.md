# EtlStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Human-readable status message for the ETL/filter pipeline. | [optional] 
**Status** | Pointer to **string** | Aggregated ETL/filter pipeline status on this host, as reported by the on-prem agent. Values are agent-emitted strings and are not part of a stable enum. | [optional] 

## Methods

### NewEtlStatus

`func NewEtlStatus() *EtlStatus`

NewEtlStatus instantiates a new EtlStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEtlStatusWithDefaults

`func NewEtlStatusWithDefaults() *EtlStatus`

NewEtlStatusWithDefaults instantiates a new EtlStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *EtlStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EtlStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EtlStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EtlStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *EtlStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EtlStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EtlStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EtlStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


