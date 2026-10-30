# AdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpAggregateLimit** | Pointer to **int64** | Number of MAC addresses. | [optional] 
**ArpCacheRefreshPeriod** | Pointer to **int64** | ARP Cache Refresh in seconds or minutes. | [optional] 
**DeviceExpirationDays** | Pointer to **int64** | Days after which expired device data needs to be purged. | [optional] 
**DnsLookupOptions** | Pointer to **string** | Specify whether to perform a DNS lookup when discovering the devices. Valid values are: * _disable_   : Turn off DNS lookup when discovering the devices. * _infraonly_ : Collect hostnames for the network devices using DNS lookup. * _all_       : Collect hostnames for network devices (switches and routers) and end host devices using DNS lookup. | [optional] 
**DnsLookupThrottle** | Pointer to **int64** | DNS Lookup Throttle. This is the value in percentage that throttles the traffic on the DNS server. | [optional] 
**EndHostExpirationDays** | Pointer to **int64** | Days after which expired assets data needs to be purged. | [optional] 
**PingRetries** | Pointer to **int64** | Ping Sweep Attempts. | [optional] 
**PingSweepFrequency** | Pointer to **int64** | Ping Sweep Frequency in days or hours. | [optional] 
**PingTimeout** | Pointer to **int64** | Ping Sweep timeout limit in seconds. | [optional] 
**PortScanningEnabled** | Pointer to **bool** | Controls the Port Scanning for the discovered devices.  Valid values are: * _enabled_: Enables the port scanning for the services specified in the Service Ports. * _disabled_: Disables the port scanning for the services specified in the Service Ports. | [optional] 
**RouteLimit** | Pointer to **int64** | number of routes that can be configured in the discovery. | [optional] 
**SnmpDiscoveryVersion** | Pointer to **string** | Authenticate and poll using one of these versions of the SNMP protocol. Valid values are: * _snmpv1_  : Use SNMPV1 protocol. * _snmpv2c_ : Use SNMPV2C protocol. | [optional] 
**Snmpv1DataCollectionFallback** | Pointer to **string** | Version of SNMP protocol for credential discovery and device polling. Valid values are: * _disabled_ : Exclusively use SNMPv2c and up. Do not fallback to SNMPV1. * _enabled_  : Support fallback to SNMPv1, if SNMPV2C and up are not available. | [optional] 
**TcpScanTechnique** | Pointer to **string** | Configure the TCP scan technique to be used for the discovery.  Valid values are: * _syn_:     Perform scans on thousands of TCP ports per system, never completing connections across any well-known port.              SYN packets are sent and the poller waits for a response while continuing to scan other ports.              A SYN/ACK response indicates the protocol port is listening while a RST indicates it is not listening.              The SYN option presents less impact on the network. * _connect_: Perform scan on IPv6 networks. Unlike the SYN option, complete connections are attempted on the              scanned system and each successive TCP protocol port being scanned. | [optional] 

## Methods

### NewAdvancedSettings

`func NewAdvancedSettings() *AdvancedSettings`

NewAdvancedSettings instantiates a new AdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedSettingsWithDefaults

`func NewAdvancedSettingsWithDefaults() *AdvancedSettings`

NewAdvancedSettingsWithDefaults instantiates a new AdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpAggregateLimit

`func (o *AdvancedSettings) GetArpAggregateLimit() int64`

GetArpAggregateLimit returns the ArpAggregateLimit field if non-nil, zero value otherwise.

### GetArpAggregateLimitOk

`func (o *AdvancedSettings) GetArpAggregateLimitOk() (*int64, bool)`

GetArpAggregateLimitOk returns a tuple with the ArpAggregateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpAggregateLimit

`func (o *AdvancedSettings) SetArpAggregateLimit(v int64)`

SetArpAggregateLimit sets ArpAggregateLimit field to given value.

### HasArpAggregateLimit

`func (o *AdvancedSettings) HasArpAggregateLimit() bool`

HasArpAggregateLimit returns a boolean if a field has been set.

### GetArpCacheRefreshPeriod

`func (o *AdvancedSettings) GetArpCacheRefreshPeriod() int64`

GetArpCacheRefreshPeriod returns the ArpCacheRefreshPeriod field if non-nil, zero value otherwise.

### GetArpCacheRefreshPeriodOk

`func (o *AdvancedSettings) GetArpCacheRefreshPeriodOk() (*int64, bool)`

GetArpCacheRefreshPeriodOk returns a tuple with the ArpCacheRefreshPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCacheRefreshPeriod

`func (o *AdvancedSettings) SetArpCacheRefreshPeriod(v int64)`

SetArpCacheRefreshPeriod sets ArpCacheRefreshPeriod field to given value.

### HasArpCacheRefreshPeriod

`func (o *AdvancedSettings) HasArpCacheRefreshPeriod() bool`

HasArpCacheRefreshPeriod returns a boolean if a field has been set.

### GetDeviceExpirationDays

`func (o *AdvancedSettings) GetDeviceExpirationDays() int64`

GetDeviceExpirationDays returns the DeviceExpirationDays field if non-nil, zero value otherwise.

### GetDeviceExpirationDaysOk

`func (o *AdvancedSettings) GetDeviceExpirationDaysOk() (*int64, bool)`

GetDeviceExpirationDaysOk returns a tuple with the DeviceExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExpirationDays

`func (o *AdvancedSettings) SetDeviceExpirationDays(v int64)`

SetDeviceExpirationDays sets DeviceExpirationDays field to given value.

### HasDeviceExpirationDays

`func (o *AdvancedSettings) HasDeviceExpirationDays() bool`

HasDeviceExpirationDays returns a boolean if a field has been set.

### GetDnsLookupOptions

`func (o *AdvancedSettings) GetDnsLookupOptions() string`

GetDnsLookupOptions returns the DnsLookupOptions field if non-nil, zero value otherwise.

### GetDnsLookupOptionsOk

`func (o *AdvancedSettings) GetDnsLookupOptionsOk() (*string, bool)`

GetDnsLookupOptionsOk returns a tuple with the DnsLookupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupOptions

`func (o *AdvancedSettings) SetDnsLookupOptions(v string)`

SetDnsLookupOptions sets DnsLookupOptions field to given value.

### HasDnsLookupOptions

`func (o *AdvancedSettings) HasDnsLookupOptions() bool`

HasDnsLookupOptions returns a boolean if a field has been set.

### GetDnsLookupThrottle

`func (o *AdvancedSettings) GetDnsLookupThrottle() int64`

GetDnsLookupThrottle returns the DnsLookupThrottle field if non-nil, zero value otherwise.

### GetDnsLookupThrottleOk

`func (o *AdvancedSettings) GetDnsLookupThrottleOk() (*int64, bool)`

GetDnsLookupThrottleOk returns a tuple with the DnsLookupThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupThrottle

`func (o *AdvancedSettings) SetDnsLookupThrottle(v int64)`

SetDnsLookupThrottle sets DnsLookupThrottle field to given value.

### HasDnsLookupThrottle

`func (o *AdvancedSettings) HasDnsLookupThrottle() bool`

HasDnsLookupThrottle returns a boolean if a field has been set.

### GetEndHostExpirationDays

`func (o *AdvancedSettings) GetEndHostExpirationDays() int64`

GetEndHostExpirationDays returns the EndHostExpirationDays field if non-nil, zero value otherwise.

### GetEndHostExpirationDaysOk

`func (o *AdvancedSettings) GetEndHostExpirationDaysOk() (*int64, bool)`

GetEndHostExpirationDaysOk returns a tuple with the EndHostExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostExpirationDays

`func (o *AdvancedSettings) SetEndHostExpirationDays(v int64)`

SetEndHostExpirationDays sets EndHostExpirationDays field to given value.

### HasEndHostExpirationDays

`func (o *AdvancedSettings) HasEndHostExpirationDays() bool`

HasEndHostExpirationDays returns a boolean if a field has been set.

### GetPingRetries

`func (o *AdvancedSettings) GetPingRetries() int64`

GetPingRetries returns the PingRetries field if non-nil, zero value otherwise.

### GetPingRetriesOk

`func (o *AdvancedSettings) GetPingRetriesOk() (*int64, bool)`

GetPingRetriesOk returns a tuple with the PingRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRetries

`func (o *AdvancedSettings) SetPingRetries(v int64)`

SetPingRetries sets PingRetries field to given value.

### HasPingRetries

`func (o *AdvancedSettings) HasPingRetries() bool`

HasPingRetries returns a boolean if a field has been set.

### GetPingSweepFrequency

`func (o *AdvancedSettings) GetPingSweepFrequency() int64`

GetPingSweepFrequency returns the PingSweepFrequency field if non-nil, zero value otherwise.

### GetPingSweepFrequencyOk

`func (o *AdvancedSettings) GetPingSweepFrequencyOk() (*int64, bool)`

GetPingSweepFrequencyOk returns a tuple with the PingSweepFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweepFrequency

`func (o *AdvancedSettings) SetPingSweepFrequency(v int64)`

SetPingSweepFrequency sets PingSweepFrequency field to given value.

### HasPingSweepFrequency

`func (o *AdvancedSettings) HasPingSweepFrequency() bool`

HasPingSweepFrequency returns a boolean if a field has been set.

### GetPingTimeout

`func (o *AdvancedSettings) GetPingTimeout() int64`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *AdvancedSettings) GetPingTimeoutOk() (*int64, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *AdvancedSettings) SetPingTimeout(v int64)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *AdvancedSettings) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetPortScanningEnabled

`func (o *AdvancedSettings) GetPortScanningEnabled() bool`

GetPortScanningEnabled returns the PortScanningEnabled field if non-nil, zero value otherwise.

### GetPortScanningEnabledOk

`func (o *AdvancedSettings) GetPortScanningEnabledOk() (*bool, bool)`

GetPortScanningEnabledOk returns a tuple with the PortScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanningEnabled

`func (o *AdvancedSettings) SetPortScanningEnabled(v bool)`

SetPortScanningEnabled sets PortScanningEnabled field to given value.

### HasPortScanningEnabled

`func (o *AdvancedSettings) HasPortScanningEnabled() bool`

HasPortScanningEnabled returns a boolean if a field has been set.

### GetRouteLimit

`func (o *AdvancedSettings) GetRouteLimit() int64`

GetRouteLimit returns the RouteLimit field if non-nil, zero value otherwise.

### GetRouteLimitOk

`func (o *AdvancedSettings) GetRouteLimitOk() (*int64, bool)`

GetRouteLimitOk returns a tuple with the RouteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteLimit

`func (o *AdvancedSettings) SetRouteLimit(v int64)`

SetRouteLimit sets RouteLimit field to given value.

### HasRouteLimit

`func (o *AdvancedSettings) HasRouteLimit() bool`

HasRouteLimit returns a boolean if a field has been set.

### GetSnmpDiscoveryVersion

`func (o *AdvancedSettings) GetSnmpDiscoveryVersion() string`

GetSnmpDiscoveryVersion returns the SnmpDiscoveryVersion field if non-nil, zero value otherwise.

### GetSnmpDiscoveryVersionOk

`func (o *AdvancedSettings) GetSnmpDiscoveryVersionOk() (*string, bool)`

GetSnmpDiscoveryVersionOk returns a tuple with the SnmpDiscoveryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpDiscoveryVersion

`func (o *AdvancedSettings) SetSnmpDiscoveryVersion(v string)`

SetSnmpDiscoveryVersion sets SnmpDiscoveryVersion field to given value.

### HasSnmpDiscoveryVersion

`func (o *AdvancedSettings) HasSnmpDiscoveryVersion() bool`

HasSnmpDiscoveryVersion returns a boolean if a field has been set.

### GetSnmpv1DataCollectionFallback

`func (o *AdvancedSettings) GetSnmpv1DataCollectionFallback() string`

GetSnmpv1DataCollectionFallback returns the Snmpv1DataCollectionFallback field if non-nil, zero value otherwise.

### GetSnmpv1DataCollectionFallbackOk

`func (o *AdvancedSettings) GetSnmpv1DataCollectionFallbackOk() (*string, bool)`

GetSnmpv1DataCollectionFallbackOk returns a tuple with the Snmpv1DataCollectionFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1DataCollectionFallback

`func (o *AdvancedSettings) SetSnmpv1DataCollectionFallback(v string)`

SetSnmpv1DataCollectionFallback sets Snmpv1DataCollectionFallback field to given value.

### HasSnmpv1DataCollectionFallback

`func (o *AdvancedSettings) HasSnmpv1DataCollectionFallback() bool`

HasSnmpv1DataCollectionFallback returns a boolean if a field has been set.

### GetTcpScanTechnique

`func (o *AdvancedSettings) GetTcpScanTechnique() string`

GetTcpScanTechnique returns the TcpScanTechnique field if non-nil, zero value otherwise.

### GetTcpScanTechniqueOk

`func (o *AdvancedSettings) GetTcpScanTechniqueOk() (*string, bool)`

GetTcpScanTechniqueOk returns a tuple with the TcpScanTechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanTechnique

`func (o *AdvancedSettings) SetTcpScanTechnique(v string)`

SetTcpScanTechnique sets TcpScanTechnique field to given value.

### HasTcpScanTechnique

`func (o *AdvancedSettings) HasTcpScanTechnique() bool`

HasTcpScanTechnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


