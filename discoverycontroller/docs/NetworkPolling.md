# NetworkPolling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpCollectionPriority** | Pointer to **string** | Instructs whether to perform ARP collection using SNMP or CLI. | [optional] 
**AutoArpRefresh** | Pointer to **bool** | When true, enables automatic ARP Refresh across all switch ports. | [optional] 
**CliCollection** | Pointer to **bool** | Instructs whether data collection for discovery is to be carried out using CLI. | [optional] 
**NetBiosScanning** | Pointer to **bool** | When true, enables NetBios Scanning. | [optional] 
**PollingFrequencyModifier** | Pointer to **int64** |  | [optional] 
**RouteCollectionPriority** | Pointer to **string** | Instructs whether to perform Route collection using SNMP or CLI. | [optional] 
**SmartSubnetPingSweep** | Pointer to **bool** | Instructs whether data collection for discovery is to be carried out during a Smart Subnet Ping Sweep. | [optional] 
**SnmpCollection** | Pointer to **bool** | Instructs whether data collection for discovery is to be carried out using SNMP protocol. | [optional] 

## Methods

### NewNetworkPolling

`func NewNetworkPolling() *NetworkPolling`

NewNetworkPolling instantiates a new NetworkPolling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPollingWithDefaults

`func NewNetworkPollingWithDefaults() *NetworkPolling`

NewNetworkPollingWithDefaults instantiates a new NetworkPolling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpCollectionPriority

`func (o *NetworkPolling) GetArpCollectionPriority() string`

GetArpCollectionPriority returns the ArpCollectionPriority field if non-nil, zero value otherwise.

### GetArpCollectionPriorityOk

`func (o *NetworkPolling) GetArpCollectionPriorityOk() (*string, bool)`

GetArpCollectionPriorityOk returns a tuple with the ArpCollectionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCollectionPriority

`func (o *NetworkPolling) SetArpCollectionPriority(v string)`

SetArpCollectionPriority sets ArpCollectionPriority field to given value.

### HasArpCollectionPriority

`func (o *NetworkPolling) HasArpCollectionPriority() bool`

HasArpCollectionPriority returns a boolean if a field has been set.

### GetAutoArpRefresh

`func (o *NetworkPolling) GetAutoArpRefresh() bool`

GetAutoArpRefresh returns the AutoArpRefresh field if non-nil, zero value otherwise.

### GetAutoArpRefreshOk

`func (o *NetworkPolling) GetAutoArpRefreshOk() (*bool, bool)`

GetAutoArpRefreshOk returns a tuple with the AutoArpRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArpRefresh

`func (o *NetworkPolling) SetAutoArpRefresh(v bool)`

SetAutoArpRefresh sets AutoArpRefresh field to given value.

### HasAutoArpRefresh

`func (o *NetworkPolling) HasAutoArpRefresh() bool`

HasAutoArpRefresh returns a boolean if a field has been set.

### GetCliCollection

`func (o *NetworkPolling) GetCliCollection() bool`

GetCliCollection returns the CliCollection field if non-nil, zero value otherwise.

### GetCliCollectionOk

`func (o *NetworkPolling) GetCliCollectionOk() (*bool, bool)`

GetCliCollectionOk returns a tuple with the CliCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollection

`func (o *NetworkPolling) SetCliCollection(v bool)`

SetCliCollection sets CliCollection field to given value.

### HasCliCollection

`func (o *NetworkPolling) HasCliCollection() bool`

HasCliCollection returns a boolean if a field has been set.

### GetNetBiosScanning

`func (o *NetworkPolling) GetNetBiosScanning() bool`

GetNetBiosScanning returns the NetBiosScanning field if non-nil, zero value otherwise.

### GetNetBiosScanningOk

`func (o *NetworkPolling) GetNetBiosScanningOk() (*bool, bool)`

GetNetBiosScanningOk returns a tuple with the NetBiosScanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBiosScanning

`func (o *NetworkPolling) SetNetBiosScanning(v bool)`

SetNetBiosScanning sets NetBiosScanning field to given value.

### HasNetBiosScanning

`func (o *NetworkPolling) HasNetBiosScanning() bool`

HasNetBiosScanning returns a boolean if a field has been set.

### GetPollingFrequencyModifier

`func (o *NetworkPolling) GetPollingFrequencyModifier() int64`

GetPollingFrequencyModifier returns the PollingFrequencyModifier field if non-nil, zero value otherwise.

### GetPollingFrequencyModifierOk

`func (o *NetworkPolling) GetPollingFrequencyModifierOk() (*int64, bool)`

GetPollingFrequencyModifierOk returns a tuple with the PollingFrequencyModifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingFrequencyModifier

`func (o *NetworkPolling) SetPollingFrequencyModifier(v int64)`

SetPollingFrequencyModifier sets PollingFrequencyModifier field to given value.

### HasPollingFrequencyModifier

`func (o *NetworkPolling) HasPollingFrequencyModifier() bool`

HasPollingFrequencyModifier returns a boolean if a field has been set.

### GetRouteCollectionPriority

`func (o *NetworkPolling) GetRouteCollectionPriority() string`

GetRouteCollectionPriority returns the RouteCollectionPriority field if non-nil, zero value otherwise.

### GetRouteCollectionPriorityOk

`func (o *NetworkPolling) GetRouteCollectionPriorityOk() (*string, bool)`

GetRouteCollectionPriorityOk returns a tuple with the RouteCollectionPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteCollectionPriority

`func (o *NetworkPolling) SetRouteCollectionPriority(v string)`

SetRouteCollectionPriority sets RouteCollectionPriority field to given value.

### HasRouteCollectionPriority

`func (o *NetworkPolling) HasRouteCollectionPriority() bool`

HasRouteCollectionPriority returns a boolean if a field has been set.

### GetSmartSubnetPingSweep

`func (o *NetworkPolling) GetSmartSubnetPingSweep() bool`

GetSmartSubnetPingSweep returns the SmartSubnetPingSweep field if non-nil, zero value otherwise.

### GetSmartSubnetPingSweepOk

`func (o *NetworkPolling) GetSmartSubnetPingSweepOk() (*bool, bool)`

GetSmartSubnetPingSweepOk returns a tuple with the SmartSubnetPingSweep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartSubnetPingSweep

`func (o *NetworkPolling) SetSmartSubnetPingSweep(v bool)`

SetSmartSubnetPingSweep sets SmartSubnetPingSweep field to given value.

### HasSmartSubnetPingSweep

`func (o *NetworkPolling) HasSmartSubnetPingSweep() bool`

HasSmartSubnetPingSweep returns a boolean if a field has been set.

### GetSnmpCollection

`func (o *NetworkPolling) GetSnmpCollection() bool`

GetSnmpCollection returns the SnmpCollection field if non-nil, zero value otherwise.

### GetSnmpCollectionOk

`func (o *NetworkPolling) GetSnmpCollectionOk() (*bool, bool)`

GetSnmpCollectionOk returns a tuple with the SnmpCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollection

`func (o *NetworkPolling) SetSnmpCollection(v bool)`

SetSnmpCollection sets SnmpCollection field to given value.

### HasSnmpCollection

`func (o *NetworkPolling) HasSnmpCollection() bool`

HasSnmpCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


