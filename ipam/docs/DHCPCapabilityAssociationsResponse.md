# DHCPCapabilityAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressBlocks** | Pointer to [**[]AddressBlock**](AddressBlock.md) | The list of address blocks to which the associated subnets and ranges belong. | [optional] 
**IpSpace** | Pointer to [**IPSpace**](IPSpace.md) | The ip_space to which the address blocks, subnets and ranges belong. | [optional] 
**Ranges** | Pointer to [**[]Range**](Range.md) | The list of associated ranges. | [optional] 
**Subnets** | Pointer to [**[]Subnet**](Subnet.md) | The list of associated subnets. | [optional] 

## Methods

### NewDHCPCapabilityAssociationsResponse

`func NewDHCPCapabilityAssociationsResponse() *DHCPCapabilityAssociationsResponse`

NewDHCPCapabilityAssociationsResponse instantiates a new DHCPCapabilityAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPCapabilityAssociationsResponseWithDefaults

`func NewDHCPCapabilityAssociationsResponseWithDefaults() *DHCPCapabilityAssociationsResponse`

NewDHCPCapabilityAssociationsResponseWithDefaults instantiates a new DHCPCapabilityAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressBlocks

`func (o *DHCPCapabilityAssociationsResponse) GetAddressBlocks() []AddressBlock`

GetAddressBlocks returns the AddressBlocks field if non-nil, zero value otherwise.

### GetAddressBlocksOk

`func (o *DHCPCapabilityAssociationsResponse) GetAddressBlocksOk() (*[]AddressBlock, bool)`

GetAddressBlocksOk returns a tuple with the AddressBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressBlocks

`func (o *DHCPCapabilityAssociationsResponse) SetAddressBlocks(v []AddressBlock)`

SetAddressBlocks sets AddressBlocks field to given value.

### HasAddressBlocks

`func (o *DHCPCapabilityAssociationsResponse) HasAddressBlocks() bool`

HasAddressBlocks returns a boolean if a field has been set.

### GetIpSpace

`func (o *DHCPCapabilityAssociationsResponse) GetIpSpace() IPSpace`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *DHCPCapabilityAssociationsResponse) GetIpSpaceOk() (*IPSpace, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *DHCPCapabilityAssociationsResponse) SetIpSpace(v IPSpace)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *DHCPCapabilityAssociationsResponse) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetRanges

`func (o *DHCPCapabilityAssociationsResponse) GetRanges() []Range`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *DHCPCapabilityAssociationsResponse) GetRangesOk() (*[]Range, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *DHCPCapabilityAssociationsResponse) SetRanges(v []Range)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *DHCPCapabilityAssociationsResponse) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetSubnets

`func (o *DHCPCapabilityAssociationsResponse) GetSubnets() []Subnet`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *DHCPCapabilityAssociationsResponse) GetSubnetsOk() (*[]Subnet, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *DHCPCapabilityAssociationsResponse) SetSubnets(v []Subnet)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *DHCPCapabilityAssociationsResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


