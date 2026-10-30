# NIDiscoveryServiceV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at. | [optional] [readonly] 
**HostId** | Pointer to **string** | Required. Id of the onprem host. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**NiDiscoveryConfig** | Pointer to [**NIDiscoveryConfigV2**](NIDiscoveryConfigV2.md) | V2 Configuration of the NI Discovery Service with scan domains. | [optional] 
**NiJobConfigId** | Pointer to **string** | Reference to the NI Job Configuration. | [optional] 
**NiJobConfigName** | Pointer to **string** | Read-only. Name of the Job Configuration. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Ophid of Host on which the job is to run. | [optional] 
**PoolId** | Pointer to **string** | Id of the pool of which the service and host belongs. | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at. | [optional] [readonly] 

## Methods

### NewNIDiscoveryServiceV2

`func NewNIDiscoveryServiceV2() *NIDiscoveryServiceV2`

NewNIDiscoveryServiceV2 instantiates a new NIDiscoveryServiceV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNIDiscoveryServiceV2WithDefaults

`func NewNIDiscoveryServiceV2WithDefaults() *NIDiscoveryServiceV2`

NewNIDiscoveryServiceV2WithDefaults instantiates a new NIDiscoveryServiceV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NIDiscoveryServiceV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NIDiscoveryServiceV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NIDiscoveryServiceV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NIDiscoveryServiceV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHostId

`func (o *NIDiscoveryServiceV2) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *NIDiscoveryServiceV2) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *NIDiscoveryServiceV2) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *NIDiscoveryServiceV2) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *NIDiscoveryServiceV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NIDiscoveryServiceV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NIDiscoveryServiceV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NIDiscoveryServiceV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNiDiscoveryConfig

`func (o *NIDiscoveryServiceV2) GetNiDiscoveryConfig() NIDiscoveryConfigV2`

GetNiDiscoveryConfig returns the NiDiscoveryConfig field if non-nil, zero value otherwise.

### GetNiDiscoveryConfigOk

`func (o *NIDiscoveryServiceV2) GetNiDiscoveryConfigOk() (*NIDiscoveryConfigV2, bool)`

GetNiDiscoveryConfigOk returns a tuple with the NiDiscoveryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiDiscoveryConfig

`func (o *NIDiscoveryServiceV2) SetNiDiscoveryConfig(v NIDiscoveryConfigV2)`

SetNiDiscoveryConfig sets NiDiscoveryConfig field to given value.

### HasNiDiscoveryConfig

`func (o *NIDiscoveryServiceV2) HasNiDiscoveryConfig() bool`

HasNiDiscoveryConfig returns a boolean if a field has been set.

### GetNiJobConfigId

`func (o *NIDiscoveryServiceV2) GetNiJobConfigId() string`

GetNiJobConfigId returns the NiJobConfigId field if non-nil, zero value otherwise.

### GetNiJobConfigIdOk

`func (o *NIDiscoveryServiceV2) GetNiJobConfigIdOk() (*string, bool)`

GetNiJobConfigIdOk returns a tuple with the NiJobConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigId

`func (o *NIDiscoveryServiceV2) SetNiJobConfigId(v string)`

SetNiJobConfigId sets NiJobConfigId field to given value.

### HasNiJobConfigId

`func (o *NIDiscoveryServiceV2) HasNiJobConfigId() bool`

HasNiJobConfigId returns a boolean if a field has been set.

### GetNiJobConfigName

`func (o *NIDiscoveryServiceV2) GetNiJobConfigName() string`

GetNiJobConfigName returns the NiJobConfigName field if non-nil, zero value otherwise.

### GetNiJobConfigNameOk

`func (o *NIDiscoveryServiceV2) GetNiJobConfigNameOk() (*string, bool)`

GetNiJobConfigNameOk returns a tuple with the NiJobConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigName

`func (o *NIDiscoveryServiceV2) SetNiJobConfigName(v string)`

SetNiJobConfigName sets NiJobConfigName field to given value.

### HasNiJobConfigName

`func (o *NIDiscoveryServiceV2) HasNiJobConfigName() bool`

HasNiJobConfigName returns a boolean if a field has been set.

### GetOphid

`func (o *NIDiscoveryServiceV2) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *NIDiscoveryServiceV2) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *NIDiscoveryServiceV2) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *NIDiscoveryServiceV2) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPoolId

`func (o *NIDiscoveryServiceV2) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *NIDiscoveryServiceV2) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *NIDiscoveryServiceV2) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *NIDiscoveryServiceV2) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *NIDiscoveryServiceV2) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *NIDiscoveryServiceV2) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *NIDiscoveryServiceV2) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *NIDiscoveryServiceV2) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NIDiscoveryServiceV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NIDiscoveryServiceV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NIDiscoveryServiceV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NIDiscoveryServiceV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


