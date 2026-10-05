# FederatedPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocation** | Pointer to [**Allocation**](Allocation.md) | The allocation details for the __FederatedPool__. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Description** | Pointer to **string** | The description for the federated pool. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**FederatedRealm** | **string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | The metadata for the federated pool in JSON format. | [optional] 
**Name** | Pointer to **string** | The name of the federated pool. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**NetworkCompliance** | Pointer to [**NetworkCompliance**](NetworkCompliance.md) | The network compliance of the __FederatedPool__. | [optional] 
**NetworkCompliant** | Pointer to **bool** | Indicates if this pool is compliant with its parent&#39;s network compliance policy. When false, a trouble dot should be displayed in the UI to indicate non-compliance. | [optional] [readonly] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | **string** | The address family of the pool (&#39;ip4&#39;, &#39;ip6&#39;, or &#39;ip4/ip6&#39; for dual mode support on NIOS_X pools only). | 
**Provider** | [**ProviderType**](ProviderType.md) | The cloud provider type this pool is associated with. | [default to PROVIDERTYPE_NIOS_X]
**Region** | **string** | The region/locale this pool is associated with (e.g., &#39;us-west-1&#39;, &#39;eu-central-1&#39;). | 
**State** | Pointer to **string** | The current state of the federated pool (e.g., &#39;create-complete&#39;, &#39;create-in-progress&#39;, &#39;delete-in-progress&#39;). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated pool in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The IPv4 utilization percentage of the __FederatedPool__. | [optional] [readonly] 
**UtilizationV6** | Pointer to [**UtilizationV6**](UtilizationV6.md) | The IPv6 utilization metrics for the __FederatedPool__. | [optional] [readonly] 

## Methods

### NewFederatedPool

`func NewFederatedPool(federatedRealm string, protocol string, provider ProviderType, region string, ) *FederatedPool`

NewFederatedPool instantiates a new FederatedPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedPoolWithDefaults

`func NewFederatedPoolWithDefaults() *FederatedPool`

NewFederatedPoolWithDefaults instantiates a new FederatedPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocation

`func (o *FederatedPool) GetAllocation() Allocation`

GetAllocation returns the Allocation field if non-nil, zero value otherwise.

### GetAllocationOk

`func (o *FederatedPool) GetAllocationOk() (*Allocation, bool)`

GetAllocationOk returns a tuple with the Allocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocation

`func (o *FederatedPool) SetAllocation(v Allocation)`

SetAllocation sets Allocation field to given value.

### HasAllocation

`func (o *FederatedPool) HasAllocation() bool`

HasAllocation returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederatedPool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederatedPool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederatedPool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederatedPool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *FederatedPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederatedPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederatedPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederatedPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFederatedRealm

`func (o *FederatedPool) GetFederatedRealm() string`

GetFederatedRealm returns the FederatedRealm field if non-nil, zero value otherwise.

### GetFederatedRealmOk

`func (o *FederatedPool) GetFederatedRealmOk() (*string, bool)`

GetFederatedRealmOk returns a tuple with the FederatedRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealm

`func (o *FederatedPool) SetFederatedRealm(v string)`

SetFederatedRealm sets FederatedRealm field to given value.


### GetId

`func (o *FederatedPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedPool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FederatedPool) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FederatedPool) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FederatedPool) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FederatedPool) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *FederatedPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCompliance

`func (o *FederatedPool) GetNetworkCompliance() NetworkCompliance`

GetNetworkCompliance returns the NetworkCompliance field if non-nil, zero value otherwise.

### GetNetworkComplianceOk

`func (o *FederatedPool) GetNetworkComplianceOk() (*NetworkCompliance, bool)`

GetNetworkComplianceOk returns a tuple with the NetworkCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCompliance

`func (o *FederatedPool) SetNetworkCompliance(v NetworkCompliance)`

SetNetworkCompliance sets NetworkCompliance field to given value.

### HasNetworkCompliance

`func (o *FederatedPool) HasNetworkCompliance() bool`

HasNetworkCompliance returns a boolean if a field has been set.

### GetNetworkCompliant

`func (o *FederatedPool) GetNetworkCompliant() bool`

GetNetworkCompliant returns the NetworkCompliant field if non-nil, zero value otherwise.

### GetNetworkCompliantOk

`func (o *FederatedPool) GetNetworkCompliantOk() (*bool, bool)`

GetNetworkCompliantOk returns a tuple with the NetworkCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCompliant

`func (o *FederatedPool) SetNetworkCompliant(v bool)`

SetNetworkCompliant sets NetworkCompliant field to given value.

### HasNetworkCompliant

`func (o *FederatedPool) HasNetworkCompliant() bool`

HasNetworkCompliant returns a boolean if a field has been set.

### GetParent

`func (o *FederatedPool) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FederatedPool) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FederatedPool) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FederatedPool) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *FederatedPool) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederatedPool) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederatedPool) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetProvider

`func (o *FederatedPool) GetProvider() ProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FederatedPool) GetProviderOk() (*ProviderType, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FederatedPool) SetProvider(v ProviderType)`

SetProvider sets Provider field to given value.


### GetRegion

`func (o *FederatedPool) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FederatedPool) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FederatedPool) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetState

`func (o *FederatedPool) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FederatedPool) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FederatedPool) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FederatedPool) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *FederatedPool) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederatedPool) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederatedPool) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederatedPool) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederatedPool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederatedPool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederatedPool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederatedPool) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *FederatedPool) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *FederatedPool) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *FederatedPool) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *FederatedPool) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *FederatedPool) GetUtilizationV6() UtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *FederatedPool) GetUtilizationV6Ok() (*UtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *FederatedPool) SetUtilizationV6(v UtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *FederatedPool) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


