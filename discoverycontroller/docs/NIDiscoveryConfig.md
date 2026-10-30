# NIDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutSchedule** | Pointer to [**[]ScheduleTime**](ScheduleTime.md) |  | [optional] 
**Interfaces** | Pointer to [**[]ScanInterface**](ScanInterface.md) | List of interfaces on which discovery should be performed. | [optional] 
**Networks** | Pointer to [**[]ScanNetwork**](ScanNetwork.md) | List of objects that represent either static IP&#39;s or a given range of IP&#39;s for the discovery engine to probe. | [optional] 
**SdnConfigurations** | Pointer to [**[]SDNConfiguration**](SDNConfiguration.md) |  | [optional] 
**SdnPollingSettings** | Pointer to [**SDNPollingSettings**](SDNPollingSettings.md) |  | [optional] 
**SeedRouterIpAddresses** | Pointer to **[]string** | Represents a list of IP addresses used as seed routers during discovery. | [optional] 

## Methods

### NewNIDiscoveryConfig

`func NewNIDiscoveryConfig() *NIDiscoveryConfig`

NewNIDiscoveryConfig instantiates a new NIDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNIDiscoveryConfigWithDefaults

`func NewNIDiscoveryConfigWithDefaults() *NIDiscoveryConfig`

NewNIDiscoveryConfigWithDefaults instantiates a new NIDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutSchedule

`func (o *NIDiscoveryConfig) GetBlackoutSchedule() []ScheduleTime`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *NIDiscoveryConfig) GetBlackoutScheduleOk() (*[]ScheduleTime, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *NIDiscoveryConfig) SetBlackoutSchedule(v []ScheduleTime)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *NIDiscoveryConfig) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.

### GetInterfaces

`func (o *NIDiscoveryConfig) GetInterfaces() []ScanInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *NIDiscoveryConfig) GetInterfacesOk() (*[]ScanInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *NIDiscoveryConfig) SetInterfaces(v []ScanInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *NIDiscoveryConfig) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetNetworks

`func (o *NIDiscoveryConfig) GetNetworks() []ScanNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NIDiscoveryConfig) GetNetworksOk() (*[]ScanNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NIDiscoveryConfig) SetNetworks(v []ScanNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NIDiscoveryConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSdnConfigurations

`func (o *NIDiscoveryConfig) GetSdnConfigurations() []SDNConfiguration`

GetSdnConfigurations returns the SdnConfigurations field if non-nil, zero value otherwise.

### GetSdnConfigurationsOk

`func (o *NIDiscoveryConfig) GetSdnConfigurationsOk() (*[]SDNConfiguration, bool)`

GetSdnConfigurationsOk returns a tuple with the SdnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigurations

`func (o *NIDiscoveryConfig) SetSdnConfigurations(v []SDNConfiguration)`

SetSdnConfigurations sets SdnConfigurations field to given value.

### HasSdnConfigurations

`func (o *NIDiscoveryConfig) HasSdnConfigurations() bool`

HasSdnConfigurations returns a boolean if a field has been set.

### GetSdnPollingSettings

`func (o *NIDiscoveryConfig) GetSdnPollingSettings() SDNPollingSettings`

GetSdnPollingSettings returns the SdnPollingSettings field if non-nil, zero value otherwise.

### GetSdnPollingSettingsOk

`func (o *NIDiscoveryConfig) GetSdnPollingSettingsOk() (*SDNPollingSettings, bool)`

GetSdnPollingSettingsOk returns a tuple with the SdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnPollingSettings

`func (o *NIDiscoveryConfig) SetSdnPollingSettings(v SDNPollingSettings)`

SetSdnPollingSettings sets SdnPollingSettings field to given value.

### HasSdnPollingSettings

`func (o *NIDiscoveryConfig) HasSdnPollingSettings() bool`

HasSdnPollingSettings returns a boolean if a field has been set.

### GetSeedRouterIpAddresses

`func (o *NIDiscoveryConfig) GetSeedRouterIpAddresses() []string`

GetSeedRouterIpAddresses returns the SeedRouterIpAddresses field if non-nil, zero value otherwise.

### GetSeedRouterIpAddressesOk

`func (o *NIDiscoveryConfig) GetSeedRouterIpAddressesOk() (*[]string, bool)`

GetSeedRouterIpAddressesOk returns a tuple with the SeedRouterIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterIpAddresses

`func (o *NIDiscoveryConfig) SetSeedRouterIpAddresses(v []string)`

SetSeedRouterIpAddresses sets SeedRouterIpAddresses field to given value.

### HasSeedRouterIpAddresses

`func (o *NIDiscoveryConfig) HasSeedRouterIpAddresses() bool`

HasSeedRouterIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


