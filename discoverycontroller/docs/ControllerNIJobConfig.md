# ControllerNIJobConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**ForceServiceListOverride** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**NiJobConfigName** | Pointer to **string** | ni_job_config_name represents the name of the NI Job Configuration. | [optional] [readonly] 
**Ophid** | Pointer to **string** | Deprecated. On-Prem Host ID on which the job is to run. | [optional] 
**PoolId** | Pointer to **string** | Unique Identifier of the pool to which the service and host belongs. | [optional] 
**Properties** | Pointer to [**ControllerNIProperties**](ControllerNIProperties.md) |  | [optional] 
**ServiceId** | Pointer to **string** | Deprecated. Unique Identifier of the discovery service associated with this job config. | [optional] 
**ServiceName** | Pointer to **string** | Deprecated. Read-only. Name of the discovery service. | [optional] [readonly] 
**Services** | Pointer to [**[]ControllerService**](ControllerService.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewControllerNIJobConfig

`func NewControllerNIJobConfig() *ControllerNIJobConfig`

NewControllerNIJobConfig instantiates a new ControllerNIJobConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNIJobConfigWithDefaults

`func NewControllerNIJobConfigWithDefaults() *ControllerNIJobConfig`

NewControllerNIJobConfigWithDefaults instantiates a new ControllerNIJobConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *ControllerNIJobConfig) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *ControllerNIJobConfig) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *ControllerNIJobConfig) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *ControllerNIJobConfig) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ControllerNIJobConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ControllerNIJobConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ControllerNIJobConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ControllerNIJobConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetForceServiceListOverride

`func (o *ControllerNIJobConfig) GetForceServiceListOverride() bool`

GetForceServiceListOverride returns the ForceServiceListOverride field if non-nil, zero value otherwise.

### GetForceServiceListOverrideOk

`func (o *ControllerNIJobConfig) GetForceServiceListOverrideOk() (*bool, bool)`

GetForceServiceListOverrideOk returns a tuple with the ForceServiceListOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceServiceListOverride

`func (o *ControllerNIJobConfig) SetForceServiceListOverride(v bool)`

SetForceServiceListOverride sets ForceServiceListOverride field to given value.

### HasForceServiceListOverride

`func (o *ControllerNIJobConfig) HasForceServiceListOverride() bool`

HasForceServiceListOverride returns a boolean if a field has been set.

### GetId

`func (o *ControllerNIJobConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerNIJobConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerNIJobConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerNIJobConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNiJobConfigName

`func (o *ControllerNIJobConfig) GetNiJobConfigName() string`

GetNiJobConfigName returns the NiJobConfigName field if non-nil, zero value otherwise.

### GetNiJobConfigNameOk

`func (o *ControllerNIJobConfig) GetNiJobConfigNameOk() (*string, bool)`

GetNiJobConfigNameOk returns a tuple with the NiJobConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiJobConfigName

`func (o *ControllerNIJobConfig) SetNiJobConfigName(v string)`

SetNiJobConfigName sets NiJobConfigName field to given value.

### HasNiJobConfigName

`func (o *ControllerNIJobConfig) HasNiJobConfigName() bool`

HasNiJobConfigName returns a boolean if a field has been set.

### GetOphid

`func (o *ControllerNIJobConfig) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ControllerNIJobConfig) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ControllerNIJobConfig) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ControllerNIJobConfig) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetPoolId

`func (o *ControllerNIJobConfig) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ControllerNIJobConfig) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ControllerNIJobConfig) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ControllerNIJobConfig) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetProperties

`func (o *ControllerNIJobConfig) GetProperties() ControllerNIProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ControllerNIJobConfig) GetPropertiesOk() (*ControllerNIProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ControllerNIJobConfig) SetProperties(v ControllerNIProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ControllerNIJobConfig) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetServiceId

`func (o *ControllerNIJobConfig) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ControllerNIJobConfig) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ControllerNIJobConfig) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *ControllerNIJobConfig) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *ControllerNIJobConfig) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ControllerNIJobConfig) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ControllerNIJobConfig) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ControllerNIJobConfig) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServices

`func (o *ControllerNIJobConfig) GetServices() []ControllerService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ControllerNIJobConfig) GetServicesOk() (*[]ControllerService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ControllerNIJobConfig) SetServices(v []ControllerService)`

SetServices sets Services field to given value.

### HasServices

`func (o *ControllerNIJobConfig) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ControllerNIJobConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ControllerNIJobConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ControllerNIJobConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ControllerNIJobConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


