# ServicePort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int64** |  | [optional] 
**Protocol** | Pointer to **string** | Service protocol correspond to one of TCP or UDP. | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewServicePort

`func NewServicePort() *ServicePort`

NewServicePort instantiates a new ServicePort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortWithDefaults

`func NewServicePortWithDefaults() *ServicePort`

NewServicePortWithDefaults instantiates a new ServicePort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ServicePort) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServicePort) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServicePort) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServicePort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *ServicePort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ServicePort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ServicePort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ServicePort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *ServicePort) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServicePort) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServicePort) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServicePort) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


