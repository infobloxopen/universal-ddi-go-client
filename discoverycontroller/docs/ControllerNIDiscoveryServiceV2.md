# ControllerNIDiscoveryServiceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at. | [optional] [readonly] 
**HostId** | Pointer to **string** | Required. Id of the onprem host. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**NiDiscoveryConfig** | Pointer to [**ControllerNIDiscoveryConfigV2**](ControllerNIDiscoveryConfigV2.md) | V2 Configuration of the NI Discovery Service with scan domains. | [optional] 
**NiJobConfigId** | Pointer to **string** | Reference to the NI Job Configuration. | [optional] 
**NiJobConfigName** | Pointer to **string** | Read-only. Name of the Job Configuration. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Ophid of Host on which the job is to run. | [optional] 
**PoolId** | Pointer to **string** | Id of the pool of which the service and host belongs. | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at. | [optional] [readonly] 

## Methods

### NewControllerNIDiscoveryServiceV2

`func NewControllerNIDiscoveryServiceV2() *ControllerNIDiscoveryServiceV2`

NewControllerNIDiscoveryServiceV2 instantiates a new ControllerNIDiscoveryServiceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNIDiscoveryServiceV2WithDefaults

`func NewControllerNIDiscoveryServiceV2WithDefaults() *ControllerNIDiscoveryServiceV2`

NewControllerNIDiscoveryServiceV2WithDefaults instantiates a new ControllerNIDiscoveryServiceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ControllerNIDiscoveryServiceV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerNIDiscoveryServiceV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerNIDiscoveryServiceV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerNIDiscoveryServiceV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHostId

`func (o *ControllerNIDiscoveryServiceV2) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ControllerNIDiscoveryServiceV2) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ControllerNIDiscoveryServiceV2) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ControllerNIDiscoveryServiceV2) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *ControllerNIDiscoveryServiceV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerNIDiscoveryServiceV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerNIDiscoveryServiceV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerNIDiscoveryServiceV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNiDiscoveryConfig

`func (o *ControllerNIDiscoveryServiceV2) GetNiDiscoveryConfig() ControllerNIDiscoveryConfigV2`

GetNiDiscoveryConfig returns the NiDiscoveryConfig field if non-nil, zero value otherwise.

### GetNiDiscoveryConfigOk

`func (o *ControllerNIDiscoveryServiceV2) GetNiDiscoveryConfigOk() (*ControllerNIDiscoveryConfigV2, bool)`

GetNiDiscoveryConfigOk returns a tuple with the NiDiscoveryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiDiscoveryConfig

`func (o *ControllerNIDiscoveryServiceV2) SetNiDiscoveryConfig(v ControllerNIDiscoveryConfigV2)`

SetNiDiscoveryConfig sets NiDiscoveryConfig field to given value.

### HasNiDiscoveryConfig

`func (o *ControllerNIDiscoveryServiceV2) HasNiDiscoveryConfig() bool`

HasNiDiscoveryConfig returns a boolean if a field has been set.

### GetNiJobConfigId

`func (o *ControllerNIDiscoveryServiceV2) GetNiJobConfigId() string`

GetNiJobConfigId returns the NiJobConfigId field if non-nil, zero value otherwise.

### GetNiJobConfigIdOk

`func (o *ControllerNIDiscoveryServiceV2) GetNiJobConfigIdOk() (*string, bool)`

GetNiJobConfigIdOk returns a tuple with the NiJobConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigId

`func (o *ControllerNIDiscoveryServiceV2) SetNiJobConfigId(v string)`

SetNiJobConfigId sets NiJobConfigId field to given value.

### HasNiJobConfigId

`func (o *ControllerNIDiscoveryServiceV2) HasNiJobConfigId() bool`

HasNiJobConfigId returns a boolean if a field has been set.

### GetNiJobConfigName

`func (o *ControllerNIDiscoveryServiceV2) GetNiJobConfigName() string`

GetNiJobConfigName returns the NiJobConfigName field if non-nil, zero value otherwise.

### GetNiJobConfigNameOk

`func (o *ControllerNIDiscoveryServiceV2) GetNiJobConfigNameOk() (*string, bool)`

GetNiJobConfigNameOk returns a tuple with the NiJobConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigName

`func (o *ControllerNIDiscoveryServiceV2) SetNiJobConfigName(v string)`

SetNiJobConfigName sets NiJobConfigName field to given value.

### HasNiJobConfigName

`func (o *ControllerNIDiscoveryServiceV2) HasNiJobConfigName() bool`

HasNiJobConfigName returns a boolean if a field has been set.

### GetOphid

`func (o *ControllerNIDiscoveryServiceV2) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ControllerNIDiscoveryServiceV2) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ControllerNIDiscoveryServiceV2) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ControllerNIDiscoveryServiceV2) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPoolId

`func (o *ControllerNIDiscoveryServiceV2) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ControllerNIDiscoveryServiceV2) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ControllerNIDiscoveryServiceV2) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ControllerNIDiscoveryServiceV2) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *ControllerNIDiscoveryServiceV2) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ControllerNIDiscoveryServiceV2) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ControllerNIDiscoveryServiceV2) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ControllerNIDiscoveryServiceV2) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerNIDiscoveryServiceV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerNIDiscoveryServiceV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerNIDiscoveryServiceV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerNIDiscoveryServiceV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


