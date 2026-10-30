# Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsUnmanaged** | Pointer to **bool** | Indicates whether discovered assets should be tagged as unmanaged. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Read-only. Timestamp created_at, equals to time when particular object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | Optional. Description of the configuration. | [optional] 
**Destination** | Pointer to **string** | Read-only. Name of the Read-Write IPSpace where discovered objects would be created. | [optional] 
**DestinationId** | Pointer to **string** | The resource identifier. | [optional] 
**Enable** | Pointer to **bool** | Flag to indicate if this Configuration is enabled. Defaults to _true_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | Required. Name of the configuration. | [optional] 
**SourceId** | Pointer to **string** | Source id is the unique identifier of the source. Example: in case of NGC, it would be NIOS grid id. | [optional] 
**SourceName** | Pointer to **string** | In case of NGC, it would be the grid name. The grid name can have conflicts for multiple grids in same account. So grid-name should be suffixed with last 4 chars of grid_id. Example: Infoblox-c35f This can be done only if name is default (Infoblox) In case of Active Discovery, it would be NMAP. | [optional] 
**SourceType** | Pointer to **string** | Required. Defines the source type for the configuration. Valid values are:  * _ngc_: NIOS Grid Connector source type.  * _active_: Active Discovery (NMAP) source type. **Deprecated** - retained for backwards compatibility only. Use _ni_ for new deployments.  * _ni_: Network Insight active discovery source type.  * _passive_: Passive Discovery source type. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at. | [optional] [readonly] 

## Methods

### NewConfiguration

`func NewConfiguration() *Configuration`

NewConfiguration instantiates a new Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationWithDefaults

`func NewConfigurationWithDefaults() *Configuration`

NewConfigurationWithDefaults instantiates a new Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsUnmanaged

`func (o *Configuration) GetAssetsUnmanaged() bool`

GetAssetsUnmanaged returns the AssetsUnmanaged field if non-nil, zero value otherwise.

### GetAssetsUnmanagedOk

`func (o *Configuration) GetAssetsUnmanagedOk() (*bool, bool)`

GetAssetsUnmanagedOk returns a tuple with the AssetsUnmanaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsUnmanaged

`func (o *Configuration) SetAssetsUnmanaged(v bool)`

SetAssetsUnmanaged sets AssetsUnmanaged field to given value.

### HasAssetsUnmanaged

`func (o *Configuration) HasAssetsUnmanaged() bool`

HasAssetsUnmanaged returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Configuration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Configuration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Configuration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Configuration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Configuration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Configuration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Configuration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Configuration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestination

`func (o *Configuration) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Configuration) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Configuration) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Configuration) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationId

`func (o *Configuration) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *Configuration) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *Configuration) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *Configuration) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetEnable

`func (o *Configuration) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Configuration) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Configuration) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *Configuration) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *Configuration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Configuration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Configuration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Configuration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Configuration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Configuration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Configuration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Configuration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceId

`func (o *Configuration) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Configuration) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Configuration) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *Configuration) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceName

`func (o *Configuration) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Configuration) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Configuration) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *Configuration) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetSourceType

`func (o *Configuration) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *Configuration) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *Configuration) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *Configuration) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Configuration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Configuration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Configuration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Configuration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


