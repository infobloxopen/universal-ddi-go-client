# NIDiscoveryService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**DestinationIpSpaceId** | Pointer to **string** | The resource identifier. | [optional] 
**DestinationIpSpaceName** | Pointer to **string** | Read-only. Name of the Read-Write IPSpace where discovered objects would be created. | [optional] 
**HostId** | Pointer to **string** | Required. Id of the onprem host. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**NiDiscoveryConfig** | Pointer to [**NIDiscoveryConfig**](NIDiscoveryConfig.md) | Configuration of the NI Discovery Service. | [optional] 
**NiJobConfigId** | Pointer to **string** |  | [optional] 
**NiJobConfigName** | Pointer to **string** | Read-only. Name of the Job Configuration. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Ophid of Host on which the job is to run. | [optional] 
**PoolId** | Pointer to **string** | Id of the pool of which the service and host belongs. | [optional] 
**ServiceId** | Pointer to **string** | The resource identifier. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewNIDiscoveryService

`func NewNIDiscoveryService() *NIDiscoveryService`

NewNIDiscoveryService instantiates a new NIDiscoveryService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNIDiscoveryServiceWithDefaults

`func NewNIDiscoveryServiceWithDefaults() *NIDiscoveryService`

NewNIDiscoveryServiceWithDefaults instantiates a new NIDiscoveryService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *NIDiscoveryService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NIDiscoveryService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NIDiscoveryService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NIDiscoveryService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDestinationIpSpaceId

`func (o *NIDiscoveryService) GetDestinationIpSpaceId() string`

GetDestinationIpSpaceId returns the DestinationIpSpaceId field if non-nil, zero value otherwise.

### GetDestinationIpSpaceIdOk

`func (o *NIDiscoveryService) GetDestinationIpSpaceIdOk() (*string, bool)`

GetDestinationIpSpaceIdOk returns a tuple with the DestinationIpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceId

`func (o *NIDiscoveryService) SetDestinationIpSpaceId(v string)`

SetDestinationIpSpaceId sets DestinationIpSpaceId field to given value.

### HasDestinationIpSpaceId

`func (o *NIDiscoveryService) HasDestinationIpSpaceId() bool`

HasDestinationIpSpaceId returns a boolean if a field has been set.

### GetDestinationIpSpaceName

`func (o *NIDiscoveryService) GetDestinationIpSpaceName() string`

GetDestinationIpSpaceName returns the DestinationIpSpaceName field if non-nil, zero value otherwise.

### GetDestinationIpSpaceNameOk

`func (o *NIDiscoveryService) GetDestinationIpSpaceNameOk() (*string, bool)`

GetDestinationIpSpaceNameOk returns a tuple with the DestinationIpSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceName

`func (o *NIDiscoveryService) SetDestinationIpSpaceName(v string)`

SetDestinationIpSpaceName sets DestinationIpSpaceName field to given value.

### HasDestinationIpSpaceName

`func (o *NIDiscoveryService) HasDestinationIpSpaceName() bool`

HasDestinationIpSpaceName returns a boolean if a field has been set.

### GetHostId

`func (o *NIDiscoveryService) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *NIDiscoveryService) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *NIDiscoveryService) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *NIDiscoveryService) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetId

`func (o *NIDiscoveryService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NIDiscoveryService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NIDiscoveryService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NIDiscoveryService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNiDiscoveryConfig

`func (o *NIDiscoveryService) GetNiDiscoveryConfig() NIDiscoveryConfig`

GetNiDiscoveryConfig returns the NiDiscoveryConfig field if non-nil, zero value otherwise.

### GetNiDiscoveryConfigOk

`func (o *NIDiscoveryService) GetNiDiscoveryConfigOk() (*NIDiscoveryConfig, bool)`

GetNiDiscoveryConfigOk returns a tuple with the NiDiscoveryConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiDiscoveryConfig

`func (o *NIDiscoveryService) SetNiDiscoveryConfig(v NIDiscoveryConfig)`

SetNiDiscoveryConfig sets NiDiscoveryConfig field to given value.

### HasNiDiscoveryConfig

`func (o *NIDiscoveryService) HasNiDiscoveryConfig() bool`

HasNiDiscoveryConfig returns a boolean if a field has been set.

### GetNiJobConfigId

`func (o *NIDiscoveryService) GetNiJobConfigId() string`

GetNiJobConfigId returns the NiJobConfigId field if non-nil, zero value otherwise.

### GetNiJobConfigIdOk

`func (o *NIDiscoveryService) GetNiJobConfigIdOk() (*string, bool)`

GetNiJobConfigIdOk returns a tuple with the NiJobConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigId

`func (o *NIDiscoveryService) SetNiJobConfigId(v string)`

SetNiJobConfigId sets NiJobConfigId field to given value.

### HasNiJobConfigId

`func (o *NIDiscoveryService) HasNiJobConfigId() bool`

HasNiJobConfigId returns a boolean if a field has been set.

### GetNiJobConfigName

`func (o *NIDiscoveryService) GetNiJobConfigName() string`

GetNiJobConfigName returns the NiJobConfigName field if non-nil, zero value otherwise.

### GetNiJobConfigNameOk

`func (o *NIDiscoveryService) GetNiJobConfigNameOk() (*string, bool)`

GetNiJobConfigNameOk returns a tuple with the NiJobConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigName

`func (o *NIDiscoveryService) SetNiJobConfigName(v string)`

SetNiJobConfigName sets NiJobConfigName field to given value.

### HasNiJobConfigName

`func (o *NIDiscoveryService) HasNiJobConfigName() bool`

HasNiJobConfigName returns a boolean if a field has been set.

### GetOphid

`func (o *NIDiscoveryService) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *NIDiscoveryService) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *NIDiscoveryService) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *NIDiscoveryService) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPoolId

`func (o *NIDiscoveryService) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *NIDiscoveryService) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *NIDiscoveryService) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *NIDiscoveryService) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetServiceId

`func (o *NIDiscoveryService) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *NIDiscoveryService) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *NIDiscoveryService) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *NIDiscoveryService) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NIDiscoveryService) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NIDiscoveryService) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NIDiscoveryService) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NIDiscoveryService) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


