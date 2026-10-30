# SDNPollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSdnVirtualNetwork** | Pointer to **string** |  | [optional] 
**EndHostPollingEnabled** | Pointer to **bool** |  | [optional] 
**EndHostPollingIntervalMinutes** | Pointer to **int64** |  | [optional] 
**SdnDiscovery** | Pointer to **bool** |  | [optional] 

## Methods

### NewSDNPollingSettings

`func NewSDNPollingSettings() *SDNPollingSettings`

NewSDNPollingSettings instantiates a new SDNPollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSDNPollingSettingsWithDefaults

`func NewSDNPollingSettingsWithDefaults() *SDNPollingSettings`

NewSDNPollingSettingsWithDefaults instantiates a new SDNPollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSdnVirtualNetwork

`func (o *SDNPollingSettings) GetDefaultSdnVirtualNetwork() string`

GetDefaultSdnVirtualNetwork returns the DefaultSdnVirtualNetwork field if non-nil, zero value otherwise.

### GetDefaultSdnVirtualNetworkOk

`func (o *SDNPollingSettings) GetDefaultSdnVirtualNetworkOk() (*string, bool)`

GetDefaultSdnVirtualNetworkOk returns a tuple with the DefaultSdnVirtualNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSdnVirtualNetwork

`func (o *SDNPollingSettings) SetDefaultSdnVirtualNetwork(v string)`

SetDefaultSdnVirtualNetwork sets DefaultSdnVirtualNetwork field to given value.

### HasDefaultSdnVirtualNetwork

`func (o *SDNPollingSettings) HasDefaultSdnVirtualNetwork() bool`

HasDefaultSdnVirtualNetwork returns a boolean if a field has been set.

### GetEndHostPollingEnabled

`func (o *SDNPollingSettings) GetEndHostPollingEnabled() bool`

GetEndHostPollingEnabled returns the EndHostPollingEnabled field if non-nil, zero value otherwise.

### GetEndHostPollingEnabledOk

`func (o *SDNPollingSettings) GetEndHostPollingEnabledOk() (*bool, bool)`

GetEndHostPollingEnabledOk returns a tuple with the EndHostPollingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostPollingEnabled

`func (o *SDNPollingSettings) SetEndHostPollingEnabled(v bool)`

SetEndHostPollingEnabled sets EndHostPollingEnabled field to given value.

### HasEndHostPollingEnabled

`func (o *SDNPollingSettings) HasEndHostPollingEnabled() bool`

HasEndHostPollingEnabled returns a boolean if a field has been set.

### GetEndHostPollingIntervalMinutes

`func (o *SDNPollingSettings) GetEndHostPollingIntervalMinutes() int64`

GetEndHostPollingIntervalMinutes returns the EndHostPollingIntervalMinutes field if non-nil, zero value otherwise.

### GetEndHostPollingIntervalMinutesOk

`func (o *SDNPollingSettings) GetEndHostPollingIntervalMinutesOk() (*int64, bool)`

GetEndHostPollingIntervalMinutesOk returns a tuple with the EndHostPollingIntervalMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostPollingIntervalMinutes

`func (o *SDNPollingSettings) SetEndHostPollingIntervalMinutes(v int64)`

SetEndHostPollingIntervalMinutes sets EndHostPollingIntervalMinutes field to given value.

### HasEndHostPollingIntervalMinutes

`func (o *SDNPollingSettings) HasEndHostPollingIntervalMinutes() bool`

HasEndHostPollingIntervalMinutes returns a boolean if a field has been set.

### GetSdnDiscovery

`func (o *SDNPollingSettings) GetSdnDiscovery() bool`

GetSdnDiscovery returns the SdnDiscovery field if non-nil, zero value otherwise.

### GetSdnDiscoveryOk

`func (o *SDNPollingSettings) GetSdnDiscoveryOk() (*bool, bool)`

GetSdnDiscoveryOk returns a tuple with the SdnDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnDiscovery

`func (o *SDNPollingSettings) SetSdnDiscovery(v bool)`

SetSdnDiscovery sets SdnDiscovery field to given value.

### HasSdnDiscovery

`func (o *SDNPollingSettings) HasSdnDiscovery() bool`

HasSdnDiscovery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


