# HostAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_. | [optional] 
**EnableDhcp** | Pointer to **bool** | The _enable_dhcp_ field controls whether the DHCP server provides an address to the client using this assignment. When false, the address is reserved in IPAM and can be used for DNS registration and other IP address management features without being handed out via DHCP. This is for nios hosts. | [optional] 
**MacAddr** | Pointer to **string** |  | [optional] 
**Ref** | Pointer to **string** | The resource identifier. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 
**UsageType** | Pointer to **string** | The _usage_type_ field indicates how the associated _Address_ is being used. The value is derived from the _Address_ usage and is one of:  * _DHCP FIXEDADDRESS_: the address has a DHCP fixed address assignment.  * _IPAM RESERVED_: the address is reserved in IPAM and is not a fixed address.  * empty string: the address has neither a fixed address assignment nor a reservation.  This field is read-only and is set by the server. | [optional] [readonly] 

## Methods

### NewHostAddress

`func NewHostAddress() *HostAddress`

NewHostAddress instantiates a new HostAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAddressWithDefaults

`func NewHostAddressWithDefaults() *HostAddress`

NewHostAddressWithDefaults instantiates a new HostAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *HostAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HostAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HostAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HostAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetEnableDhcp

`func (o *HostAddress) GetEnableDhcp() bool`

GetEnableDhcp returns the EnableDhcp field if non-nil, zero value otherwise.

### GetEnableDhcpOk

`func (o *HostAddress) GetEnableDhcpOk() (*bool, bool)`

GetEnableDhcpOk returns a tuple with the EnableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDhcp

`func (o *HostAddress) SetEnableDhcp(v bool)`

SetEnableDhcp sets EnableDhcp field to given value.

### HasEnableDhcp

`func (o *HostAddress) HasEnableDhcp() bool`

HasEnableDhcp returns a boolean if a field has been set.

### GetMacAddr

`func (o *HostAddress) GetMacAddr() string`

GetMacAddr returns the MacAddr field if non-nil, zero value otherwise.

### GetMacAddrOk

`func (o *HostAddress) GetMacAddrOk() (*string, bool)`

GetMacAddrOk returns a tuple with the MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddr

`func (o *HostAddress) SetMacAddr(v string)`

SetMacAddr sets MacAddr field to given value.

### HasMacAddr

`func (o *HostAddress) HasMacAddr() bool`

HasMacAddr returns a boolean if a field has been set.

### GetRef

`func (o *HostAddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *HostAddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *HostAddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *HostAddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSpace

`func (o *HostAddress) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *HostAddress) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *HostAddress) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *HostAddress) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetUsageType

`func (o *HostAddress) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *HostAddress) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *HostAddress) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *HostAddress) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


