# ScanNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationIpSpaceId** | Pointer to **string** | The resource identifier. | [optional] 
**DestinationIpSpaceName** | Pointer to **string** | Read-only. Name of the Read-Write IPSpace where discovered objects would be created. | [optional] 
**ExclusionIpAddressRange** | Pointer to **string** | Represents the ipaddress or range that has to be excluded from the discovery scan. | [optional] 
**ExclusionMode** | Pointer to **string** |  | [optional] 
**IpAddressRange** | Pointer to **string** | Represents the network to be scanned in the format of cidr/mask or range of IPs. | [optional] 
**IpSpaceAction** | Pointer to **string** | The inheritance setting for IP Space Valid values are: * _inherit_: Use the inherited value from parent IPSpace * _override_: Use the value set in the object.  Defaults to _inherit_. | [optional] 
**PingSweep** | Pointer to **bool** |  | [optional] 

## Methods

### NewScanNetwork

`func NewScanNetwork() *ScanNetwork`

NewScanNetwork instantiates a new ScanNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanNetworkWithDefaults

`func NewScanNetworkWithDefaults() *ScanNetwork`

NewScanNetworkWithDefaults instantiates a new ScanNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIpSpaceId

`func (o *ScanNetwork) GetDestinationIpSpaceId() string`

GetDestinationIpSpaceId returns the DestinationIpSpaceId field if non-nil, zero value otherwise.

### GetDestinationIpSpaceIdOk

`func (o *ScanNetwork) GetDestinationIpSpaceIdOk() (*string, bool)`

GetDestinationIpSpaceIdOk returns a tuple with the DestinationIpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceId

`func (o *ScanNetwork) SetDestinationIpSpaceId(v string)`

SetDestinationIpSpaceId sets DestinationIpSpaceId field to given value.

### HasDestinationIpSpaceId

`func (o *ScanNetwork) HasDestinationIpSpaceId() bool`

HasDestinationIpSpaceId returns a boolean if a field has been set.

### GetDestinationIpSpaceName

`func (o *ScanNetwork) GetDestinationIpSpaceName() string`

GetDestinationIpSpaceName returns the DestinationIpSpaceName field if non-nil, zero value otherwise.

### GetDestinationIpSpaceNameOk

`func (o *ScanNetwork) GetDestinationIpSpaceNameOk() (*string, bool)`

GetDestinationIpSpaceNameOk returns a tuple with the DestinationIpSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceName

`func (o *ScanNetwork) SetDestinationIpSpaceName(v string)`

SetDestinationIpSpaceName sets DestinationIpSpaceName field to given value.

### HasDestinationIpSpaceName

`func (o *ScanNetwork) HasDestinationIpSpaceName() bool`

HasDestinationIpSpaceName returns a boolean if a field has been set.

### GetExclusionIpAddressRange

`func (o *ScanNetwork) GetExclusionIpAddressRange() string`

GetExclusionIpAddressRange returns the ExclusionIpAddressRange field if non-nil, zero value otherwise.

### GetExclusionIpAddressRangeOk

`func (o *ScanNetwork) GetExclusionIpAddressRangeOk() (*string, bool)`

GetExclusionIpAddressRangeOk returns a tuple with the ExclusionIpAddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionIpAddressRange

`func (o *ScanNetwork) SetExclusionIpAddressRange(v string)`

SetExclusionIpAddressRange sets ExclusionIpAddressRange field to given value.

### HasExclusionIpAddressRange

`func (o *ScanNetwork) HasExclusionIpAddressRange() bool`

HasExclusionIpAddressRange returns a boolean if a field has been set.

### GetExclusionMode

`func (o *ScanNetwork) GetExclusionMode() string`

GetExclusionMode returns the ExclusionMode field if non-nil, zero value otherwise.

### GetExclusionModeOk

`func (o *ScanNetwork) GetExclusionModeOk() (*string, bool)`

GetExclusionModeOk returns a tuple with the ExclusionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionMode

`func (o *ScanNetwork) SetExclusionMode(v string)`

SetExclusionMode sets ExclusionMode field to given value.

### HasExclusionMode

`func (o *ScanNetwork) HasExclusionMode() bool`

HasExclusionMode returns a boolean if a field has been set.

### GetIpAddressRange

`func (o *ScanNetwork) GetIpAddressRange() string`

GetIpAddressRange returns the IpAddressRange field if non-nil, zero value otherwise.

### GetIpAddressRangeOk

`func (o *ScanNetwork) GetIpAddressRangeOk() (*string, bool)`

GetIpAddressRangeOk returns a tuple with the IpAddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressRange

`func (o *ScanNetwork) SetIpAddressRange(v string)`

SetIpAddressRange sets IpAddressRange field to given value.

### HasIpAddressRange

`func (o *ScanNetwork) HasIpAddressRange() bool`

HasIpAddressRange returns a boolean if a field has been set.

### GetIpSpaceAction

`func (o *ScanNetwork) GetIpSpaceAction() string`

GetIpSpaceAction returns the IpSpaceAction field if non-nil, zero value otherwise.

### GetIpSpaceActionOk

`func (o *ScanNetwork) GetIpSpaceActionOk() (*string, bool)`

GetIpSpaceActionOk returns a tuple with the IpSpaceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceAction

`func (o *ScanNetwork) SetIpSpaceAction(v string)`

SetIpSpaceAction sets IpSpaceAction field to given value.

### HasIpSpaceAction

`func (o *ScanNetwork) HasIpSpaceAction() bool`

HasIpSpaceAction returns a boolean if a field has been set.

### GetPingSweep

`func (o *ScanNetwork) GetPingSweep() bool`

GetPingSweep returns the PingSweep field if non-nil, zero value otherwise.

### GetPingSweepOk

`func (o *ScanNetwork) GetPingSweepOk() (*bool, bool)`

GetPingSweepOk returns a tuple with the PingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweep

`func (o *ScanNetwork) SetPingSweep(v bool)`

SetPingSweep sets PingSweep field to given value.

### HasPingSweep

`func (o *ScanNetwork) HasPingSweep() bool`

HasPingSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


