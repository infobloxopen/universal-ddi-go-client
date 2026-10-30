# ControllerService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ophid** | Pointer to **string** | Required. On-Prem Host ID on which the job is to run. | [optional] 
**ServiceId** | Pointer to **string** | Required. Unique Identifier of the discovery service associated with this job config. | [optional] 
**ServiceName** | Pointer to **string** | Read-only. Name of the discovery service. | [optional] [readonly] 

## Methods

### NewControllerService

`func NewControllerService() *ControllerService`

NewControllerService instantiates a new ControllerService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerServiceWithDefaults

`func NewControllerServiceWithDefaults() *ControllerService`

NewControllerServiceWithDefaults instantiates a new ControllerService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOphid

`func (o *ControllerService) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ControllerService) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ControllerService) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ControllerService) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetServiceId

`func (o *ControllerService) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ControllerService) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ControllerService) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ControllerService) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ControllerService) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ControllerService) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ControllerService) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ControllerService) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


