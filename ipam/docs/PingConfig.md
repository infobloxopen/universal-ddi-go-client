# PingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PingCheck** | Pointer to **bool** | Indicates if Ping before Offer is enabled. Default is _false_. | [optional] 
**PingNumberOfRequests** | Pointer to **int64** | Number of ping requests. Default is 1. | [optional] 
**PingTimeOut** | Pointer to **int64** | Ping time out in milliseconds. Default is 100ms. | [optional] 

## Methods

### NewPingConfig

`func NewPingConfig() *PingConfig`

NewPingConfig instantiates a new PingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingConfigWithDefaults

`func NewPingConfigWithDefaults() *PingConfig`

NewPingConfigWithDefaults instantiates a new PingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPingCheck

`func (o *PingConfig) GetPingCheck() bool`

GetPingCheck returns the PingCheck field if non-nil, zero value otherwise.

### GetPingCheckOk

`func (o *PingConfig) GetPingCheckOk() (*bool, bool)`

GetPingCheckOk returns a tuple with the PingCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingCheck

`func (o *PingConfig) SetPingCheck(v bool)`

SetPingCheck sets PingCheck field to given value.

### HasPingCheck

`func (o *PingConfig) HasPingCheck() bool`

HasPingCheck returns a boolean if a field has been set.

### GetPingNumberOfRequests

`func (o *PingConfig) GetPingNumberOfRequests() int64`

GetPingNumberOfRequests returns the PingNumberOfRequests field if non-nil, zero value otherwise.

### GetPingNumberOfRequestsOk

`func (o *PingConfig) GetPingNumberOfRequestsOk() (*int64, bool)`

GetPingNumberOfRequestsOk returns a tuple with the PingNumberOfRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingNumberOfRequests

`func (o *PingConfig) SetPingNumberOfRequests(v int64)`

SetPingNumberOfRequests sets PingNumberOfRequests field to given value.

### HasPingNumberOfRequests

`func (o *PingConfig) HasPingNumberOfRequests() bool`

HasPingNumberOfRequests returns a boolean if a field has been set.

### GetPingTimeOut

`func (o *PingConfig) GetPingTimeOut() int64`

GetPingTimeOut returns the PingTimeOut field if non-nil, zero value otherwise.

### GetPingTimeOutOk

`func (o *PingConfig) GetPingTimeOutOk() (*int64, bool)`

GetPingTimeOutOk returns a tuple with the PingTimeOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeOut

`func (o *PingConfig) SetPingTimeOut(v int64)`

SetPingTimeOut sets PingTimeOut field to given value.

### HasPingTimeOut

`func (o *PingConfig) HasPingTimeOut() bool`

HasPingTimeOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


