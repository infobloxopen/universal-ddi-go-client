# NetworkCompliance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultNetmaskLength** | Pointer to **int64** | The default prefix length suggested when allocating new child address blocks or subnets. 0 means the server auto-fills this to max(block CIDR + 1, minimum_netmask_length). | [optional] 
**MaximumNetmaskLength** | Pointer to **int64** | The maximum (most specific) prefix length allowed for child address blocks or subnets. A child address block or subnet with a prefix length greater than this value is non-compliant. 0 means the server auto-fills this to the protocol ceiling (32 for IPv4, 128 for IPv6). | [optional] 
**MinimumNetmaskLength** | Pointer to **int64** | The minimum (most general) prefix length allowed for child __AddressBlock__ or __Subnet__ object. A child address block or subnet with a prefix length smaller than this value is non-compliant. 0 means the server auto-fills this to the parent block&#39;s CIDR + 1. | [optional] 

## Methods

### NewNetworkCompliance

`func NewNetworkCompliance() *NetworkCompliance`

NewNetworkCompliance instantiates a new NetworkCompliance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkComplianceWithDefaults

`func NewNetworkComplianceWithDefaults() *NetworkCompliance`

NewNetworkComplianceWithDefaults instantiates a new NetworkCompliance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultNetmaskLength

`func (o *NetworkCompliance) GetDefaultNetmaskLength() int64`

GetDefaultNetmaskLength returns the DefaultNetmaskLength field if non-nil, zero value otherwise.

### GetDefaultNetmaskLengthOk

`func (o *NetworkCompliance) GetDefaultNetmaskLengthOk() (*int64, bool)`

GetDefaultNetmaskLengthOk returns a tuple with the DefaultNetmaskLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetmaskLength

`func (o *NetworkCompliance) SetDefaultNetmaskLength(v int64)`

SetDefaultNetmaskLength sets DefaultNetmaskLength field to given value.

### HasDefaultNetmaskLength

`func (o *NetworkCompliance) HasDefaultNetmaskLength() bool`

HasDefaultNetmaskLength returns a boolean if a field has been set.

### GetMaximumNetmaskLength

`func (o *NetworkCompliance) GetMaximumNetmaskLength() int64`

GetMaximumNetmaskLength returns the MaximumNetmaskLength field if non-nil, zero value otherwise.

### GetMaximumNetmaskLengthOk

`func (o *NetworkCompliance) GetMaximumNetmaskLengthOk() (*int64, bool)`

GetMaximumNetmaskLengthOk returns a tuple with the MaximumNetmaskLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumNetmaskLength

`func (o *NetworkCompliance) SetMaximumNetmaskLength(v int64)`

SetMaximumNetmaskLength sets MaximumNetmaskLength field to given value.

### HasMaximumNetmaskLength

`func (o *NetworkCompliance) HasMaximumNetmaskLength() bool`

HasMaximumNetmaskLength returns a boolean if a field has been set.

### GetMinimumNetmaskLength

`func (o *NetworkCompliance) GetMinimumNetmaskLength() int64`

GetMinimumNetmaskLength returns the MinimumNetmaskLength field if non-nil, zero value otherwise.

### GetMinimumNetmaskLengthOk

`func (o *NetworkCompliance) GetMinimumNetmaskLengthOk() (*int64, bool)`

GetMinimumNetmaskLengthOk returns a tuple with the MinimumNetmaskLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumNetmaskLength

`func (o *NetworkCompliance) SetMinimumNetmaskLength(v int64)`

SetMinimumNetmaskLength sets MinimumNetmaskLength field to given value.

### HasMinimumNetmaskLength

`func (o *NetworkCompliance) HasMinimumNetmaskLength() bool`

HasMinimumNetmaskLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


