# ControllerDiscoveryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**DiscoveryConfig** | Pointer to [**ControllerDiscoveryConfig**](ControllerDiscoveryConfig.md) | Config of the Discovery Service. | [optional] 
**HostId** | Pointer to **string** | Required. Id of the onprem host. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Ophid** | Pointer to **string** |  | [optional] 
**PoolId** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewControllerDiscoveryService

`func NewControllerDiscoveryService() *ControllerDiscoveryService`

NewControllerDiscoveryService instantiates a new ControllerDiscoveryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerDiscoveryServiceWithDefaults

`func NewControllerDiscoveryServiceWithDefaults() *ControllerDiscoveryService`

NewControllerDiscoveryServiceWithDefaults instantiates a new ControllerDiscoveryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ControllerDiscoveryService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerDiscoveryService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerDiscoveryService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerDiscoveryService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDiscoveryConfig

`func (o *ControllerDiscoveryService) GetDiscoveryConfig() ControllerDiscoveryConfig`

GetDiscoveryConfig returns the DiscoveryConfig field if non-nil, zero value otherwise.

### GetDiscoveryConfigOk

`func (o *ControllerDiscoveryService) GetDiscoveryConfigOk() (*ControllerDiscoveryConfig, bool)`

GetDiscoveryConfigOk returns a tuple with the DiscoveryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryConfig

`func (o *ControllerDiscoveryService) SetDiscoveryConfig(v ControllerDiscoveryConfig)`

SetDiscoveryConfig sets DiscoveryConfig field to given value.

### HasDiscoveryConfig

`func (o *ControllerDiscoveryService) HasDiscoveryConfig() bool`

HasDiscoveryConfig returns a boolean if a field has been set.

### GetHostId

`func (o *ControllerDiscoveryService) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ControllerDiscoveryService) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ControllerDiscoveryService) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ControllerDiscoveryService) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *ControllerDiscoveryService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerDiscoveryService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerDiscoveryService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerDiscoveryService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOphid

`func (o *ControllerDiscoveryService) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ControllerDiscoveryService) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ControllerDiscoveryService) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ControllerDiscoveryService) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPoolId

`func (o *ControllerDiscoveryService) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ControllerDiscoveryService) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ControllerDiscoveryService) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ControllerDiscoveryService) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *ControllerDiscoveryService) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ControllerDiscoveryService) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ControllerDiscoveryService) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ControllerDiscoveryService) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerDiscoveryService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerDiscoveryService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerDiscoveryService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerDiscoveryService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


