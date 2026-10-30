# ControllerScanNetwork

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

### NewControllerScanNetwork

`func NewControllerScanNetwork() *ControllerScanNetwork`

NewControllerScanNetwork instantiates a new ControllerScanNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerScanNetworkWithDefaults

`func NewControllerScanNetworkWithDefaults() *ControllerScanNetwork`

NewControllerScanNetworkWithDefaults instantiates a new ControllerScanNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationIpSpaceId

`func (o *ControllerScanNetwork) GetDestinationIpSpaceId() string`

GetDestinationIpSpaceId returns the DestinationIpSpaceId field if non-nil, zero value otherwise.

### GetDestinationIpSpaceIdOk

`func (o *ControllerScanNetwork) GetDestinationIpSpaceIdOk() (*string, bool)`

GetDestinationIpSpaceIdOk returns a tuple with the DestinationIpSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceId

`func (o *ControllerScanNetwork) SetDestinationIpSpaceId(v string)`

SetDestinationIpSpaceId sets DestinationIpSpaceId field to given value.

### HasDestinationIpSpaceId

`func (o *ControllerScanNetwork) HasDestinationIpSpaceId() bool`

HasDestinationIpSpaceId returns a boolean if a field has been set.

### GetDestinationIpSpaceName

`func (o *ControllerScanNetwork) GetDestinationIpSpaceName() string`

GetDestinationIpSpaceName returns the DestinationIpSpaceName field if non-nil, zero value otherwise.

### GetDestinationIpSpaceNameOk

`func (o *ControllerScanNetwork) GetDestinationIpSpaceNameOk() (*string, bool)`

GetDestinationIpSpaceNameOk returns a tuple with the DestinationIpSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIpSpaceName

`func (o *ControllerScanNetwork) SetDestinationIpSpaceName(v string)`

SetDestinationIpSpaceName sets DestinationIpSpaceName field to given value.

### HasDestinationIpSpaceName

`func (o *ControllerScanNetwork) HasDestinationIpSpaceName() bool`

HasDestinationIpSpaceName returns a boolean if a field has been set.

### GetExclusionIpAddressRange

`func (o *ControllerScanNetwork) GetExclusionIpAddressRange() string`

GetExclusionIpAddressRange returns the ExclusionIpAddressRange field if non-nil, zero value otherwise.

### GetExclusionIpAddressRangeOk

`func (o *ControllerScanNetwork) GetExclusionIpAddressRangeOk() (*string, bool)`

GetExclusionIpAddressRangeOk returns a tuple with the ExclusionIpAddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionIpAddressRange

`func (o *ControllerScanNetwork) SetExclusionIpAddressRange(v string)`

SetExclusionIpAddressRange sets ExclusionIpAddressRange field to given value.

### HasExclusionIpAddressRange

`func (o *ControllerScanNetwork) HasExclusionIpAddressRange() bool`

HasExclusionIpAddressRange returns a boolean if a field has been set.

### GetExclusionMode

`func (o *ControllerScanNetwork) GetExclusionMode() string`

GetExclusionMode returns the ExclusionMode field if non-nil, zero value otherwise.

### GetExclusionModeOk

`func (o *ControllerScanNetwork) GetExclusionModeOk() (*string, bool)`

GetExclusionModeOk returns a tuple with the ExclusionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclusionMode

`func (o *ControllerScanNetwork) SetExclusionMode(v string)`

SetExclusionMode sets ExclusionMode field to given value.

### HasExclusionMode

`func (o *ControllerScanNetwork) HasExclusionMode() bool`

HasExclusionMode returns a boolean if a field has been set.

### GetIpAddressRange

`func (o *ControllerScanNetwork) GetIpAddressRange() string`

GetIpAddressRange returns the IpAddressRange field if non-nil, zero value otherwise.

### GetIpAddressRangeOk

`func (o *ControllerScanNetwork) GetIpAddressRangeOk() (*string, bool)`

GetIpAddressRangeOk returns a tuple with the IpAddressRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressRange

`func (o *ControllerScanNetwork) SetIpAddressRange(v string)`

SetIpAddressRange sets IpAddressRange field to given value.

### HasIpAddressRange

`func (o *ControllerScanNetwork) HasIpAddressRange() bool`

HasIpAddressRange returns a boolean if a field has been set.

### GetIpSpaceAction

`func (o *ControllerScanNetwork) GetIpSpaceAction() string`

GetIpSpaceAction returns the IpSpaceAction field if non-nil, zero value otherwise.

### GetIpSpaceActionOk

`func (o *ControllerScanNetwork) GetIpSpaceActionOk() (*string, bool)`

GetIpSpaceActionOk returns a tuple with the IpSpaceAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpaceAction

`func (o *ControllerScanNetwork) SetIpSpaceAction(v string)`

SetIpSpaceAction sets IpSpaceAction field to given value.

### HasIpSpaceAction

`func (o *ControllerScanNetwork) HasIpSpaceAction() bool`

HasIpSpaceAction returns a boolean if a field has been set.

### GetPingSweep

`func (o *ControllerScanNetwork) GetPingSweep() bool`

GetPingSweep returns the PingSweep field if non-nil, zero value otherwise.

### GetPingSweepOk

`func (o *ControllerScanNetwork) GetPingSweepOk() (*bool, bool)`

GetPingSweepOk returns a tuple with the PingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweep

`func (o *ControllerScanNetwork) SetPingSweep(v bool)`

SetPingSweep sets PingSweep field to given value.

### HasPingSweep

`func (o *ControllerScanNetwork) HasPingSweep() bool`

HasPingSweep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


