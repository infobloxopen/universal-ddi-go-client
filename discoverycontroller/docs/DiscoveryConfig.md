# DiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to **[]string** | List of networks in the form of address/cidr to be scanned. | [optional] 

## Methods

### NewDiscoveryConfig

`func NewDiscoveryConfig() *DiscoveryConfig`

NewDiscoveryConfig instantiates a new DiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryConfigWithDefaults

`func NewDiscoveryConfigWithDefaults() *DiscoveryConfig`

NewDiscoveryConfigWithDefaults instantiates a new DiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *DiscoveryConfig) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DiscoveryConfig) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DiscoveryConfig) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *DiscoveryConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


