# ControllerNIDiscoveryConfigV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutSchedule** | Pointer to [**[]ControllerScheduleTime**](ControllerScheduleTime.md) | Discovery blackout settings. | [optional] 
**ScanDomains** | Pointer to [**[]ControllerScanDomain**](ControllerScanDomain.md) | Scan domains grouping scan networks and seed routers by interface. | [optional] 
**SdnConfigurations** | Pointer to [**[]ControllerSDNConfiguration**](ControllerSDNConfiguration.md) | List of SDN connection credentials. | [optional] 
**SdnPollingSettings** | Pointer to [**ControllerSDNPollingSettings**](ControllerSDNPollingSettings.md) | Settings for SDN endpoint polling behavior. | [optional] 

## Methods

### NewControllerNIDiscoveryConfigV2

`func NewControllerNIDiscoveryConfigV2() *ControllerNIDiscoveryConfigV2`

NewControllerNIDiscoveryConfigV2 instantiates a new ControllerNIDiscoveryConfigV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerNIDiscoveryConfigV2WithDefaults

`func NewControllerNIDiscoveryConfigV2WithDefaults() *ControllerNIDiscoveryConfigV2`

NewControllerNIDiscoveryConfigV2WithDefaults instantiates a new ControllerNIDiscoveryConfigV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutSchedule

`func (o *ControllerNIDiscoveryConfigV2) GetBlackoutSchedule() []ControllerScheduleTime`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *ControllerNIDiscoveryConfigV2) GetBlackoutScheduleOk() (*[]ControllerScheduleTime, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *ControllerNIDiscoveryConfigV2) SetBlackoutSchedule(v []ControllerScheduleTime)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *ControllerNIDiscoveryConfigV2) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.

### GetScanDomains

`func (o *ControllerNIDiscoveryConfigV2) GetScanDomains() []ControllerScanDomain`

GetScanDomains returns the ScanDomains field if non-nil, zero value otherwise.

### GetScanDomainsOk

`func (o *ControllerNIDiscoveryConfigV2) GetScanDomainsOk() (*[]ControllerScanDomain, bool)`

GetScanDomainsOk returns a tuple with the ScanDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDomains

`func (o *ControllerNIDiscoveryConfigV2) SetScanDomains(v []ControllerScanDomain)`

SetScanDomains sets ScanDomains field to given value.

### HasScanDomains

`func (o *ControllerNIDiscoveryConfigV2) HasScanDomains() bool`

HasScanDomains returns a boolean if a field has been set.

### GetSdnConfigurations

`func (o *ControllerNIDiscoveryConfigV2) GetSdnConfigurations() []ControllerSDNConfiguration`

GetSdnConfigurations returns the SdnConfigurations field if non-nil, zero value otherwise.

### GetSdnConfigurationsOk

`func (o *ControllerNIDiscoveryConfigV2) GetSdnConfigurationsOk() (*[]ControllerSDNConfiguration, bool)`

GetSdnConfigurationsOk returns a tuple with the SdnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigurations

`func (o *ControllerNIDiscoveryConfigV2) SetSdnConfigurations(v []ControllerSDNConfiguration)`

SetSdnConfigurations sets SdnConfigurations field to given value.

### HasSdnConfigurations

`func (o *ControllerNIDiscoveryConfigV2) HasSdnConfigurations() bool`

HasSdnConfigurations returns a boolean if a field has been set.

### GetSdnPollingSettings

`func (o *ControllerNIDiscoveryConfigV2) GetSdnPollingSettings() ControllerSDNPollingSettings`

GetSdnPollingSettings returns the SdnPollingSettings field if non-nil, zero value otherwise.

### GetSdnPollingSettingsOk

`func (o *ControllerNIDiscoveryConfigV2) GetSdnPollingSettingsOk() (*ControllerSDNPollingSettings, bool)`

GetSdnPollingSettingsOk returns a tuple with the SdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnPollingSettings

`func (o *ControllerNIDiscoveryConfigV2) SetSdnPollingSettings(v ControllerSDNPollingSettings)`

SetSdnPollingSettings sets SdnPollingSettings field to given value.

### HasSdnPollingSettings

`func (o *ControllerNIDiscoveryConfigV2) HasSdnPollingSettings() bool`

HasSdnPollingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


