# SharedNetworkSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IPv4 or IPv6 address of the subnet. | [optional] 
**Cidr** | Pointer to **int64** | The CIDR prefix length of the subnet. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the subnet. May contain 1 to 256 characters. Can include UTF-8. | [optional] 

## Methods

### NewSharedNetworkSubnet

`func NewSharedNetworkSubnet() *SharedNetworkSubnet`

NewSharedNetworkSubnet instantiates a new SharedNetworkSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedNetworkSubnetWithDefaults

`func NewSharedNetworkSubnetWithDefaults() *SharedNetworkSubnet`

NewSharedNetworkSubnetWithDefaults instantiates a new SharedNetworkSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SharedNetworkSubnet) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SharedNetworkSubnet) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SharedNetworkSubnet) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SharedNetworkSubnet) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCidr

`func (o *SharedNetworkSubnet) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *SharedNetworkSubnet) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *SharedNetworkSubnet) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *SharedNetworkSubnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetId

`func (o *SharedNetworkSubnet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedNetworkSubnet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedNetworkSubnet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharedNetworkSubnet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SharedNetworkSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedNetworkSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedNetworkSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SharedNetworkSubnet) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


