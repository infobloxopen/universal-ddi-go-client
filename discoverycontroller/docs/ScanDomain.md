# ScanDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. Unique identifier of this scan domain. | [optional] [readonly] 
**InterfaceName** | Pointer to **string** | Host interface associated with this scan domain (e.g., \&quot;eth0\&quot;). Defaults to \&quot;eth0\&quot; if empty. | [optional] 
**IpSpaceId** | Pointer to **string** | The resource identifier. | [optional] 
**IpSpaceName** | Pointer to **string** | Read-only. Name of the IP space, resolved from IPAM. | [optional] [readonly] 
**Name** | Pointer to **string** | User-given name for this scan domain (e.g., \&quot;default\&quot;). Must be unique within a service config. | [optional] 
**Networks** | Pointer to [**[]ScanNetwork**](ScanNetwork.md) | Scan networks belonging to this scan domain. | [optional] 
**SeedRouterIpAddresses** | Pointer to **[]string** | Seed router IP addresses belonging to this scan domain. | [optional] 
**VrfHint** | Pointer to **string** | Regex to match VRF discovered over the interface. Defaults to \&quot;.*\&quot; (match all). | [optional] 

## Methods

### NewScanDomain

`func NewScanDomain() *ScanDomain`

NewScanDomain instantiates a new ScanDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanDomainWithDefaults

`func NewScanDomainWithDefaults() *ScanDomain`

NewScanDomainWithDefaults instantiates a new ScanDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScanDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScanDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScanDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ScanDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *ScanDomain) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *ScanDomain) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *ScanDomain) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *ScanDomain) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIpSpaceId

`func (o *ScanDomain) GetIpSpaceId() string`

GetIpSpaceId returns the IpSpaceId field if non-nil, zero value otherwise.

### GetIpSpaceIdOk

`func (o *ScanDomain) GetIpSpaceIdOk() (*string, bool)`

GetIpSpaceIdOk returns a tuple with the IpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceId

`func (o *ScanDomain) SetIpSpaceId(v string)`

SetIpSpaceId sets IpSpaceId field to given value.

### HasIpSpaceId

`func (o *ScanDomain) HasIpSpaceId() bool`

HasIpSpaceId returns a boolean if a field has been set.

### GetIpSpaceName

`func (o *ScanDomain) GetIpSpaceName() string`

GetIpSpaceName returns the IpSpaceName field if non-nil, zero value otherwise.

### GetIpSpaceNameOk

`func (o *ScanDomain) GetIpSpaceNameOk() (*string, bool)`

GetIpSpaceNameOk returns a tuple with the IpSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceName

`func (o *ScanDomain) SetIpSpaceName(v string)`

SetIpSpaceName sets IpSpaceName field to given value.

### HasIpSpaceName

`func (o *ScanDomain) HasIpSpaceName() bool`

HasIpSpaceName returns a boolean if a field has been set.

### GetName

`func (o *ScanDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanDomain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScanDomain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworks

`func (o *ScanDomain) GetNetworks() []ScanNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ScanDomain) GetNetworksOk() (*[]ScanNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ScanDomain) SetNetworks(v []ScanNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ScanDomain) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSeedRouterIpAddresses

`func (o *ScanDomain) GetSeedRouterIpAddresses() []string`

GetSeedRouterIpAddresses returns the SeedRouterIpAddresses field if non-nil, zero value otherwise.

### GetSeedRouterIpAddressesOk

`func (o *ScanDomain) GetSeedRouterIpAddressesOk() (*[]string, bool)`

GetSeedRouterIpAddressesOk returns a tuple with the SeedRouterIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterIpAddresses

`func (o *ScanDomain) SetSeedRouterIpAddresses(v []string)`

SetSeedRouterIpAddresses sets SeedRouterIpAddresses field to given value.

### HasSeedRouterIpAddresses

`func (o *ScanDomain) HasSeedRouterIpAddresses() bool`

HasSeedRouterIpAddresses returns a boolean if a field has been set.

### GetVrfHint

`func (o *ScanDomain) GetVrfHint() string`

GetVrfHint returns the VrfHint field if non-nil, zero value otherwise.

### GetVrfHintOk

`func (o *ScanDomain) GetVrfHintOk() (*string, bool)`

GetVrfHintOk returns a tuple with the VrfHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfHint

`func (o *ScanDomain) SetVrfHint(v string)`

SetVrfHint sets VrfHint field to given value.

### HasVrfHint

`func (o *ScanDomain) HasVrfHint() bool`

HasVrfHint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


