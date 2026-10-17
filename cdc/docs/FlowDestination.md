# FlowDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the destination is enabled. | [optional] 
**Id** | Pointer to **int64** | Numeric identifier of the destination. | [optional] [readonly] 
**Name** | Pointer to **string** | Display name of the destination. | [optional] 
**OutputDataFormat** | Pointer to **string** | Wire format the destination is configured to emit (e.g. json, cef, leef). | [optional] 
**TransportProtocol** | Pointer to **string** | Transport protocol used by the destination (e.g. tcp, udp, https). | [optional] 
**Type** | Pointer to **string** | Destination type (e.g. syslog, splunk, splunkcloud, http, application, bloxone, reporting). | [optional] 

## Methods

### NewFlowDestination

`func NewFlowDestination() *FlowDestination`

NewFlowDestination instantiates a new FlowDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowDestinationWithDefaults

`func NewFlowDestinationWithDefaults() *FlowDestination`

NewFlowDestinationWithDefaults instantiates a new FlowDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *FlowDestination) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FlowDestination) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FlowDestination) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FlowDestination) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *FlowDestination) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowDestination) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowDestination) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FlowDestination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FlowDestination) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FlowDestination) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FlowDestination) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FlowDestination) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *FlowDestination) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *FlowDestination) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *FlowDestination) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *FlowDestination) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *FlowDestination) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *FlowDestination) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *FlowDestination) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *FlowDestination) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetType

`func (o *FlowDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FlowDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FlowDestination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FlowDestination) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


