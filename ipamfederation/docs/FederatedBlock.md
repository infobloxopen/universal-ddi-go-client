# FederatedBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**AllocationV4** | Pointer to [**Allocation**](Allocation.md) | The percentage of the Federated Block’s total address space that is consumed by Leaf Terminals. | [optional] 
**Cidr** | Pointer to **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the federated block. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**FederatedPoolId** | Pointer to **string** | The resource identifier. | [optional] 
**FederatedRealm** | **string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | The metadata for the federated block in JSON format. | [optional] 
**Name** | Pointer to **string** | The name of the federated block. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**NetworkCompliance** | Pointer to [**NetworkCompliance**](NetworkCompliance.md) | The network compliance policy for child objects. This defines the minimum, default, and maximum netmask lengths that can be used when creating child blocks. | [optional] 
**NetworkCompliant** | Pointer to **bool** | Indicates if this block is compliant with its parent&#39;s network compliance policy. When false, a trouble dot should be displayed in the UI to indicate non-compliance. | [optional] [readonly] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol of federated block (_ip4_ or _ip6_). | [optional] [readonly] 
**Region** | Pointer to **string** | The region where the federated block is located. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated block in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Utilization** | Pointer to **int64** | The percentage of Federated Block utilization. | [optional] 
**UtilizationV6** | Pointer to [**UtilizationV6**](UtilizationV6.md) | The IPv6 utilization metrics for the federated block. | [optional] [readonly] 

## Methods

### NewFederatedBlock

`func NewFederatedBlock(address string, federatedRealm string, ) *FederatedBlock`

NewFederatedBlock instantiates a new FederatedBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederatedBlockWithDefaults

`func NewFederatedBlockWithDefaults() *FederatedBlock`

NewFederatedBlockWithDefaults instantiates a new FederatedBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FederatedBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FederatedBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FederatedBlock) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAllocationV4

`func (o *FederatedBlock) GetAllocationV4() Allocation`

GetAllocationV4 returns the AllocationV4 field if non-nil, zero value otherwise.

### GetAllocationV4Ok

`func (o *FederatedBlock) GetAllocationV4Ok() (*Allocation, bool)`

GetAllocationV4Ok returns a tuple with the AllocationV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationV4

`func (o *FederatedBlock) SetAllocationV4(v Allocation)`

SetAllocationV4 sets AllocationV4 field to given value.

### HasAllocationV4

`func (o *FederatedBlock) HasAllocationV4() bool`

HasAllocationV4 returns a boolean if a field has been set.

### GetCidr

`func (o *FederatedBlock) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FederatedBlock) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FederatedBlock) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FederatedBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *FederatedBlock) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederatedBlock) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederatedBlock) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederatedBlock) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FederatedBlock) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederatedBlock) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederatedBlock) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FederatedBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFederatedPoolId

`func (o *FederatedBlock) GetFederatedPoolId() string`

GetFederatedPoolId returns the FederatedPoolId field if non-nil, zero value otherwise.

### GetFederatedPoolIdOk

`func (o *FederatedBlock) GetFederatedPoolIdOk() (*string, bool)`

GetFederatedPoolIdOk returns a tuple with the FederatedPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedPoolId

`func (o *FederatedBlock) SetFederatedPoolId(v string)`

SetFederatedPoolId sets FederatedPoolId field to given value.

### HasFederatedPoolId

`func (o *FederatedBlock) HasFederatedPoolId() bool`

HasFederatedPoolId returns a boolean if a field has been set.

### GetFederatedRealm

`func (o *FederatedBlock) GetFederatedRealm() string`

GetFederatedRealm returns the FederatedRealm field if non-nil, zero value otherwise.

### GetFederatedRealmOk

`func (o *FederatedBlock) GetFederatedRealmOk() (*string, bool)`

GetFederatedRealmOk returns a tuple with the FederatedRealm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealm

`func (o *FederatedBlock) SetFederatedRealm(v string)`

SetFederatedRealm sets FederatedRealm field to given value.


### GetId

`func (o *FederatedBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederatedBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederatedBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederatedBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *FederatedBlock) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FederatedBlock) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FederatedBlock) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FederatedBlock) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *FederatedBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederatedBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederatedBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederatedBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCompliance

`func (o *FederatedBlock) GetNetworkCompliance() NetworkCompliance`

GetNetworkCompliance returns the NetworkCompliance field if non-nil, zero value otherwise.

### GetNetworkComplianceOk

`func (o *FederatedBlock) GetNetworkComplianceOk() (*NetworkCompliance, bool)`

GetNetworkComplianceOk returns a tuple with the NetworkCompliance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCompliance

`func (o *FederatedBlock) SetNetworkCompliance(v NetworkCompliance)`

SetNetworkCompliance sets NetworkCompliance field to given value.

### HasNetworkCompliance

`func (o *FederatedBlock) HasNetworkCompliance() bool`

HasNetworkCompliance returns a boolean if a field has been set.

### GetNetworkCompliant

`func (o *FederatedBlock) GetNetworkCompliant() bool`

GetNetworkCompliant returns the NetworkCompliant field if non-nil, zero value otherwise.

### GetNetworkCompliantOk

`func (o *FederatedBlock) GetNetworkCompliantOk() (*bool, bool)`

GetNetworkCompliantOk returns a tuple with the NetworkCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCompliant

`func (o *FederatedBlock) SetNetworkCompliant(v bool)`

SetNetworkCompliant sets NetworkCompliant field to given value.

### HasNetworkCompliant

`func (o *FederatedBlock) HasNetworkCompliant() bool`

HasNetworkCompliant returns a boolean if a field has been set.

### GetParent

`func (o *FederatedBlock) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FederatedBlock) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FederatedBlock) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FederatedBlock) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *FederatedBlock) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *FederatedBlock) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *FederatedBlock) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *FederatedBlock) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRegion

`func (o *FederatedBlock) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FederatedBlock) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FederatedBlock) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FederatedBlock) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetState

`func (o *FederatedBlock) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FederatedBlock) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FederatedBlock) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *FederatedBlock) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *FederatedBlock) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederatedBlock) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederatedBlock) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederatedBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FederatedBlock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FederatedBlock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FederatedBlock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FederatedBlock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUtilization

`func (o *FederatedBlock) GetUtilization() int64`

GetUtilization returns the Utilization field if non-nil, zero value otherwise.

### GetUtilizationOk

`func (o *FederatedBlock) GetUtilizationOk() (*int64, bool)`

GetUtilizationOk returns a tuple with the Utilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilization

`func (o *FederatedBlock) SetUtilization(v int64)`

SetUtilization sets Utilization field to given value.

### HasUtilization

`func (o *FederatedBlock) HasUtilization() bool`

HasUtilization returns a boolean if a field has been set.

### GetUtilizationV6

`func (o *FederatedBlock) GetUtilizationV6() UtilizationV6`

GetUtilizationV6 returns the UtilizationV6 field if non-nil, zero value otherwise.

### GetUtilizationV6Ok

`func (o *FederatedBlock) GetUtilizationV6Ok() (*UtilizationV6, bool)`

GetUtilizationV6Ok returns a tuple with the UtilizationV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationV6

`func (o *FederatedBlock) SetUtilizationV6(v UtilizationV6)`

SetUtilizationV6 sets UtilizationV6 field to given value.

### HasUtilizationV6

`func (o *FederatedBlock) HasUtilizationV6() bool`

HasUtilizationV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


