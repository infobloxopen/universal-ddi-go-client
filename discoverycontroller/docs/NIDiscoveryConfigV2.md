# NIDiscoveryConfigV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutSchedule** | Pointer to [**[]ScheduleTime**](ScheduleTime.md) | Discovery blackout settings. | [optional] 
**ScanDomains** | Pointer to [**[]ScanDomain**](ScanDomain.md) | Scan domains grouping scan networks and seed routers by interface. | [optional] 
**SdnConfigurations** | Pointer to [**[]SDNConfiguration**](SDNConfiguration.md) | List of SDN connection credentials. | [optional] 
**SdnPollingSettings** | Pointer to [**SDNPollingSettings**](SDNPollingSettings.md) | Settings for SDN endpoint polling behavior. | [optional] 

## Methods

### NewNIDiscoveryConfigV2

`func NewNIDiscoveryConfigV2() *NIDiscoveryConfigV2`

NewNIDiscoveryConfigV2 instantiates a new NIDiscoveryConfigV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNIDiscoveryConfigV2WithDefaults

`func NewNIDiscoveryConfigV2WithDefaults() *NIDiscoveryConfigV2`

NewNIDiscoveryConfigV2WithDefaults instantiates a new NIDiscoveryConfigV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlackoutSchedule

`func (o *NIDiscoveryConfigV2) GetBlackoutSchedule() []ScheduleTime`

GetBlackoutSchedule returns the BlackoutSchedule field if non-nil, zero value otherwise.

### GetBlackoutScheduleOk

`func (o *NIDiscoveryConfigV2) GetBlackoutScheduleOk() (*[]ScheduleTime, bool)`

GetBlackoutScheduleOk returns a tuple with the BlackoutSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackoutSchedule

`func (o *NIDiscoveryConfigV2) SetBlackoutSchedule(v []ScheduleTime)`

SetBlackoutSchedule sets BlackoutSchedule field to given value.

### HasBlackoutSchedule

`func (o *NIDiscoveryConfigV2) HasBlackoutSchedule() bool`

HasBlackoutSchedule returns a boolean if a field has been set.

### GetScanDomains

`func (o *NIDiscoveryConfigV2) GetScanDomains() []ScanDomain`

GetScanDomains returns the ScanDomains field if non-nil, zero value otherwise.

### GetScanDomainsOk

`func (o *NIDiscoveryConfigV2) GetScanDomainsOk() (*[]ScanDomain, bool)`

GetScanDomainsOk returns a tuple with the ScanDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanDomains

`func (o *NIDiscoveryConfigV2) SetScanDomains(v []ScanDomain)`

SetScanDomains sets ScanDomains field to given value.

### HasScanDomains

`func (o *NIDiscoveryConfigV2) HasScanDomains() bool`

HasScanDomains returns a boolean if a field has been set.

### GetSdnConfigurations

`func (o *NIDiscoveryConfigV2) GetSdnConfigurations() []SDNConfiguration`

GetSdnConfigurations returns the SdnConfigurations field if non-nil, zero value otherwise.

### GetSdnConfigurationsOk

`func (o *NIDiscoveryConfigV2) GetSdnConfigurationsOk() (*[]SDNConfiguration, bool)`

GetSdnConfigurationsOk returns a tuple with the SdnConfigurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnConfigurations

`func (o *NIDiscoveryConfigV2) SetSdnConfigurations(v []SDNConfiguration)`

SetSdnConfigurations sets SdnConfigurations field to given value.

### HasSdnConfigurations

`func (o *NIDiscoveryConfigV2) HasSdnConfigurations() bool`

HasSdnConfigurations returns a boolean if a field has been set.

### GetSdnPollingSettings

`func (o *NIDiscoveryConfigV2) GetSdnPollingSettings() SDNPollingSettings`

GetSdnPollingSettings returns the SdnPollingSettings field if non-nil, zero value otherwise.

### GetSdnPollingSettingsOk

`func (o *NIDiscoveryConfigV2) GetSdnPollingSettingsOk() (*SDNPollingSettings, bool)`

GetSdnPollingSettingsOk returns a tuple with the SdnPollingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnPollingSettings

`func (o *NIDiscoveryConfigV2) SetSdnPollingSettings(v SDNPollingSettings)`

SetSdnPollingSettings sets SdnPollingSettings field to given value.

### HasSdnPollingSettings

`func (o *NIDiscoveryConfigV2) HasSdnPollingSettings() bool`

HasSdnPollingSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


