# ControllerNIDiscoveryConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutSchedule** | Pointer to [**[]ControllerScheduleTime**](ControllerScheduleTime.md) |  | [optional] 
**Interfaces** | Pointer to [**[]ControllerScanInterface**](ControllerScanInterface.md) | List of interfaces on which discovery should be performed. | [optional] 
**Networks** | Pointer to [**[]ControllerScanNetwork**](ControllerScanNetwork.md) | List of objects that represent either static IP&#39;s or a given range of IP&#39;s for the discovery engine to probe. | [optional] 
**SdnConfigurations** | Pointer to [**[]ControllerSDNConfiguration**](ControllerSDNConfiguration.md) |  | [optional] 
**SdnPollingSettings** | Pointer to [**ControllerSDNPollingSettings**](ControllerSDNPollingSettings.md) |  | [optional] 
**SeedRouterIpAddresses** | Pointer to **[]string** | Represents a list of IP addresses used as seed routers during discovery. | [optional] 

## Methods

### NewControllerNIDiscoveryConfig

`func NewControllerNIDiscoveryConfig() *ControllerNIDiscoveryConfig`

NewControllerNIDiscoveryConfig instantiates a new ControllerNIDiscoveryConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNIDiscoveryConfigWithDefaults

`func NewControllerNIDiscoveryConfigWithDefaults() *ControllerNIDiscoveryConfig`

NewControllerNIDiscoveryConfigWithDefaults instantiates a new ControllerNIDiscoveryConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutSchedule

`func (o *ControllerNIDiscoveryConfig) GetBlackoutSchedule() []ControllerScheduleTime`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *ControllerNIDiscoveryConfig) GetBlackoutScheduleOk() (*[]ControllerScheduleTime, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *ControllerNIDiscoveryConfig) SetBlackoutSchedule(v []ControllerScheduleTime)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *ControllerNIDiscoveryConfig) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.

### GetInterfaces

`func (o *ControllerNIDiscoveryConfig) GetInterfaces() []ControllerScanInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *ControllerNIDiscoveryConfig) GetInterfacesOk() (*[]ControllerScanInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *ControllerNIDiscoveryConfig) SetInterfaces(v []ControllerScanInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *ControllerNIDiscoveryConfig) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetNetworks

`func (o *ControllerNIDiscoveryConfig) GetNetworks() []ControllerScanNetwork`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *ControllerNIDiscoveryConfig) GetNetworksOk() (*[]ControllerScanNetwork, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *ControllerNIDiscoveryConfig) SetNetworks(v []ControllerScanNetwork)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *ControllerNIDiscoveryConfig) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetSdnConfigurations

`func (o *ControllerNIDiscoveryConfig) GetSdnConfigurations() []ControllerSDNConfiguration`

GetSdnConfigurations returns the SdnConfigurations field if non-nil, zero value otherwise.

### GetSdnConfigurationsOk

`func (o *ControllerNIDiscoveryConfig) GetSdnConfigurationsOk() (*[]ControllerSDNConfiguration, bool)`

GetSdnConfigurationsOk returns a tuple with the SdnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigurations

`func (o *ControllerNIDiscoveryConfig) SetSdnConfigurations(v []ControllerSDNConfiguration)`

SetSdnConfigurations sets SdnConfigurations field to given value.

### HasSdnConfigurations

`func (o *ControllerNIDiscoveryConfig) HasSdnConfigurations() bool`

HasSdnConfigurations returns a boolean if a field has been set.

### GetSdnPollingSettings

`func (o *ControllerNIDiscoveryConfig) GetSdnPollingSettings() ControllerSDNPollingSettings`

GetSdnPollingSettings returns the SdnPollingSettings field if non-nil, zero value otherwise.

### GetSdnPollingSettingsOk

`func (o *ControllerNIDiscoveryConfig) GetSdnPollingSettingsOk() (*ControllerSDNPollingSettings, bool)`

GetSdnPollingSettingsOk returns a tuple with the SdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnPollingSettings

`func (o *ControllerNIDiscoveryConfig) SetSdnPollingSettings(v ControllerSDNPollingSettings)`

SetSdnPollingSettings sets SdnPollingSettings field to given value.

### HasSdnPollingSettings

`func (o *ControllerNIDiscoveryConfig) HasSdnPollingSettings() bool`

HasSdnPollingSettings returns a boolean if a field has been set.

### GetSeedRouterIpAddresses

`func (o *ControllerNIDiscoveryConfig) GetSeedRouterIpAddresses() []string`

GetSeedRouterIpAddresses returns the SeedRouterIpAddresses field if non-nil, zero value otherwise.

### GetSeedRouterIpAddressesOk

`func (o *ControllerNIDiscoveryConfig) GetSeedRouterIpAddressesOk() (*[]string, bool)`

GetSeedRouterIpAddressesOk returns a tuple with the SeedRouterIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedRouterIpAddresses

`func (o *ControllerNIDiscoveryConfig) SetSeedRouterIpAddresses(v []string)`

SetSeedRouterIpAddresses sets SeedRouterIpAddresses field to given value.

### HasSeedRouterIpAddresses

`func (o *ControllerNIDiscoveryConfig) HasSeedRouterIpAddresses() bool`

HasSeedRouterIpAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


