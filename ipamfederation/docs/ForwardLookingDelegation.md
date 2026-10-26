# ForwardLookingDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”. | 
**Cidr** | Pointer to **int64** | The CIDR of the delegation. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the delegation. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**FederatedPoolId** | Pointer to **string** | The resource identifier. | [optional] 
**FederatedRealms** | **[]string** | The resource identifier. | 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the delegation. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**NetworkCompliant** | Pointer to **bool** | The compliance status of the forward looking delegation, as determined by the federation service. | [optional] [readonly] 
**Protocol** | Pointer to **string** | The type of protocol of delegation (_ip4_ or _ip6_). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the delegation in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewForwardLookingDelegation

`func NewForwardLookingDelegation(address string, federatedRealms []string, ) *ForwardLookingDelegation`

NewForwardLookingDelegation instantiates a new ForwardLookingDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardLookingDelegationWithDefaults

`func NewForwardLookingDelegationWithDefaults() *ForwardLookingDelegation`

NewForwardLookingDelegationWithDefaults instantiates a new ForwardLookingDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ForwardLookingDelegation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ForwardLookingDelegation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ForwardLookingDelegation) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCidr

`func (o *ForwardLookingDelegation) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ForwardLookingDelegation) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ForwardLookingDelegation) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ForwardLookingDelegation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *ForwardLookingDelegation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ForwardLookingDelegation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ForwardLookingDelegation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ForwardLookingDelegation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ForwardLookingDelegation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForwardLookingDelegation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForwardLookingDelegation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ForwardLookingDelegation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFederatedPoolId

`func (o *ForwardLookingDelegation) GetFederatedPoolId() string`

GetFederatedPoolId returns the FederatedPoolId field if non-nil, zero value otherwise.

### GetFederatedPoolIdOk

`func (o *ForwardLookingDelegation) GetFederatedPoolIdOk() (*string, bool)`

GetFederatedPoolIdOk returns a tuple with the FederatedPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedPoolId

`func (o *ForwardLookingDelegation) SetFederatedPoolId(v string)`

SetFederatedPoolId sets FederatedPoolId field to given value.

### HasFederatedPoolId

`func (o *ForwardLookingDelegation) HasFederatedPoolId() bool`

HasFederatedPoolId returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *ForwardLookingDelegation) GetFederatedRealms() []string`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *ForwardLookingDelegation) GetFederatedRealmsOk() (*[]string, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *ForwardLookingDelegation) SetFederatedRealms(v []string)`

SetFederatedRealms sets FederatedRealms field to given value.


### GetId

`func (o *ForwardLookingDelegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForwardLookingDelegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForwardLookingDelegation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ForwardLookingDelegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ForwardLookingDelegation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForwardLookingDelegation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForwardLookingDelegation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForwardLookingDelegation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCompliant

`func (o *ForwardLookingDelegation) GetNetworkCompliant() bool`

GetNetworkCompliant returns the NetworkCompliant field if non-nil, zero value otherwise.

### GetNetworkCompliantOk

`func (o *ForwardLookingDelegation) GetNetworkCompliantOk() (*bool, bool)`

GetNetworkCompliantOk returns a tuple with the NetworkCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCompliant

`func (o *ForwardLookingDelegation) SetNetworkCompliant(v bool)`

SetNetworkCompliant sets NetworkCompliant field to given value.

### HasNetworkCompliant

`func (o *ForwardLookingDelegation) HasNetworkCompliant() bool`

HasNetworkCompliant returns a boolean if a field has been set.

### GetProtocol

`func (o *ForwardLookingDelegation) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ForwardLookingDelegation) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ForwardLookingDelegation) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ForwardLookingDelegation) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *ForwardLookingDelegation) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ForwardLookingDelegation) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ForwardLookingDelegation) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ForwardLookingDelegation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ForwardLookingDelegation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ForwardLookingDelegation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ForwardLookingDelegation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ForwardLookingDelegation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


