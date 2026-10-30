# ControllerAdvancedSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArpAggregateLimit** | Pointer to **int32** | Number of MAC addresses. | [optional] 
**ArpCacheRefreshPeriod** | Pointer to **int32** | ARP Cache Refresh in seconds or minutes. | [optional] 
**DeviceExpirationDays** | Pointer to **int32** | Days after which expired device data needs to be purged. | [optional] 
**DnsLookupOptions** | Pointer to **string** | Specify whether to perform a DNS lookup when discovering the devices. Valid values are: * _disable_   : Turn off DNS lookup when discovering the devices. * _infraonly_ : Collect hostnames for the network devices using DNS lookup. * _all_       : Collect hostnames for network devices (switches and routers) and end host devices using DNS lookup. | [optional] 
**DnsLookupThrottle** | Pointer to **int32** | DNS Lookup Throttle. This is the value in percentage that throttles the traffic on the DNS server. | [optional] 
**EndHostExpirationDays** | Pointer to **int32** | Days after which expired assets data needs to be purged. | [optional] 
**PingRetries** | Pointer to **int32** | Ping Sweep Attempts. | [optional] 
**PingSweepFrequency** | Pointer to **int32** | Ping Sweep Frequency in days or hours. | [optional] 
**PingTimeout** | Pointer to **int32** | Ping Sweep timeout limit in seconds. | [optional] 
**PortScanningEnabled** | Pointer to **bool** | Controls the Port Scanning for the discovered devices.  Valid values are: * _enabled_: Enables the port scanning for the services specified in the Service Ports. * _disabled_: Disables the port scanning for the services specified in the Service Ports. | [optional] 
**RouteLimit** | Pointer to **int32** | number of routes that can be configured in the discovery. | [optional] 
**SnmpDiscoveryVersion** | Pointer to **string** | Authenticate and poll using one of these versions of the SNMP protocol. Valid values are: * _snmpv1_  : Use SNMPV1 protocol. * _snmpv2c_ : Use SNMPV2C protocol. | [optional] 
**Snmpv1DataCollectionFallback** | Pointer to **string** | Version of SNMP protocol for credential discovery and device polling. Valid values are: * _disabled_ : Exclusively use SNMPv2c and up. Do not fallback to SNMPV1. * _enabled_  : Support fallback to SNMPv1, if SNMPV2C and up are not available. | [optional] 
**TcpScanTechnique** | Pointer to **string** | Configure the TCP scan technique to be used for the discovery.  Valid values are: * _syn_:     Perform scans on thousands of TCP ports per system, never completing connections across any well-known port.              SYN packets are sent and the poller waits for a response while continuing to scan other ports.              A SYN/ACK response indicates the protocol port is listening while a RST indicates it is not listening.              The SYN option presents less impact on the network. * _connect_: Perform scan on IPv6 networks. Unlike the SYN option, complete connections are attempted on the              scanned system and each successive TCP protocol port being scanned. | [optional] 

## Methods

### NewControllerAdvancedSettings

`func NewControllerAdvancedSettings() *ControllerAdvancedSettings`

NewControllerAdvancedSettings instantiates a new ControllerAdvancedSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerAdvancedSettingsWithDefaults

`func NewControllerAdvancedSettingsWithDefaults() *ControllerAdvancedSettings`

NewControllerAdvancedSettingsWithDefaults instantiates a new ControllerAdvancedSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArpAggregateLimit

`func (o *ControllerAdvancedSettings) GetArpAggregateLimit() int32`

GetArpAggregateLimit returns the ArpAggregateLimit field if non-nil, zero value otherwise.

### GetArpAggregateLimitOk

`func (o *ControllerAdvancedSettings) GetArpAggregateLimitOk() (*int32, bool)`

GetArpAggregateLimitOk returns a tuple with the ArpAggregateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpAggregateLimit

`func (o *ControllerAdvancedSettings) SetArpAggregateLimit(v int32)`

SetArpAggregateLimit sets ArpAggregateLimit field to given value.

### HasArpAggregateLimit

`func (o *ControllerAdvancedSettings) HasArpAggregateLimit() bool`

HasArpAggregateLimit returns a boolean if a field has been set.

### GetArpCacheRefreshPeriod

`func (o *ControllerAdvancedSettings) GetArpCacheRefreshPeriod() int32`

GetArpCacheRefreshPeriod returns the ArpCacheRefreshPeriod field if non-nil, zero value otherwise.

### GetArpCacheRefreshPeriodOk

`func (o *ControllerAdvancedSettings) GetArpCacheRefreshPeriodOk() (*int32, bool)`

GetArpCacheRefreshPeriodOk returns a tuple with the ArpCacheRefreshPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpCacheRefreshPeriod

`func (o *ControllerAdvancedSettings) SetArpCacheRefreshPeriod(v int32)`

SetArpCacheRefreshPeriod sets ArpCacheRefreshPeriod field to given value.

### HasArpCacheRefreshPeriod

`func (o *ControllerAdvancedSettings) HasArpCacheRefreshPeriod() bool`

HasArpCacheRefreshPeriod returns a boolean if a field has been set.

### GetDeviceExpirationDays

`func (o *ControllerAdvancedSettings) GetDeviceExpirationDays() int32`

GetDeviceExpirationDays returns the DeviceExpirationDays field if non-nil, zero value otherwise.

### GetDeviceExpirationDaysOk

`func (o *ControllerAdvancedSettings) GetDeviceExpirationDaysOk() (*int32, bool)`

GetDeviceExpirationDaysOk returns a tuple with the DeviceExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceExpirationDays

`func (o *ControllerAdvancedSettings) SetDeviceExpirationDays(v int32)`

SetDeviceExpirationDays sets DeviceExpirationDays field to given value.

### HasDeviceExpirationDays

`func (o *ControllerAdvancedSettings) HasDeviceExpirationDays() bool`

HasDeviceExpirationDays returns a boolean if a field has been set.

### GetDnsLookupOptions

`func (o *ControllerAdvancedSettings) GetDnsLookupOptions() string`

GetDnsLookupOptions returns the DnsLookupOptions field if non-nil, zero value otherwise.

### GetDnsLookupOptionsOk

`func (o *ControllerAdvancedSettings) GetDnsLookupOptionsOk() (*string, bool)`

GetDnsLookupOptionsOk returns a tuple with the DnsLookupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupOptions

`func (o *ControllerAdvancedSettings) SetDnsLookupOptions(v string)`

SetDnsLookupOptions sets DnsLookupOptions field to given value.

### HasDnsLookupOptions

`func (o *ControllerAdvancedSettings) HasDnsLookupOptions() bool`

HasDnsLookupOptions returns a boolean if a field has been set.

### GetDnsLookupThrottle

`func (o *ControllerAdvancedSettings) GetDnsLookupThrottle() int32`

GetDnsLookupThrottle returns the DnsLookupThrottle field if non-nil, zero value otherwise.

### GetDnsLookupThrottleOk

`func (o *ControllerAdvancedSettings) GetDnsLookupThrottleOk() (*int32, bool)`

GetDnsLookupThrottleOk returns a tuple with the DnsLookupThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsLookupThrottle

`func (o *ControllerAdvancedSettings) SetDnsLookupThrottle(v int32)`

SetDnsLookupThrottle sets DnsLookupThrottle field to given value.

### HasDnsLookupThrottle

`func (o *ControllerAdvancedSettings) HasDnsLookupThrottle() bool`

HasDnsLookupThrottle returns a boolean if a field has been set.

### GetEndHostExpirationDays

`func (o *ControllerAdvancedSettings) GetEndHostExpirationDays() int32`

GetEndHostExpirationDays returns the EndHostExpirationDays field if non-nil, zero value otherwise.

### GetEndHostExpirationDaysOk

`func (o *ControllerAdvancedSettings) GetEndHostExpirationDaysOk() (*int32, bool)`

GetEndHostExpirationDaysOk returns a tuple with the EndHostExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostExpirationDays

`func (o *ControllerAdvancedSettings) SetEndHostExpirationDays(v int32)`

SetEndHostExpirationDays sets EndHostExpirationDays field to given value.

### HasEndHostExpirationDays

`func (o *ControllerAdvancedSettings) HasEndHostExpirationDays() bool`

HasEndHostExpirationDays returns a boolean if a field has been set.

### GetPingRetries

`func (o *ControllerAdvancedSettings) GetPingRetries() int32`

GetPingRetries returns the PingRetries field if non-nil, zero value otherwise.

### GetPingRetriesOk

`func (o *ControllerAdvancedSettings) GetPingRetriesOk() (*int32, bool)`

GetPingRetriesOk returns a tuple with the PingRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingRetries

`func (o *ControllerAdvancedSettings) SetPingRetries(v int32)`

SetPingRetries sets PingRetries field to given value.

### HasPingRetries

`func (o *ControllerAdvancedSettings) HasPingRetries() bool`

HasPingRetries returns a boolean if a field has been set.

### GetPingSweepFrequency

`func (o *ControllerAdvancedSettings) GetPingSweepFrequency() int32`

GetPingSweepFrequency returns the PingSweepFrequency field if non-nil, zero value otherwise.

### GetPingSweepFrequencyOk

`func (o *ControllerAdvancedSettings) GetPingSweepFrequencyOk() (*int32, bool)`

GetPingSweepFrequencyOk returns a tuple with the PingSweepFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingSweepFrequency

`func (o *ControllerAdvancedSettings) SetPingSweepFrequency(v int32)`

SetPingSweepFrequency sets PingSweepFrequency field to given value.

### HasPingSweepFrequency

`func (o *ControllerAdvancedSettings) HasPingSweepFrequency() bool`

HasPingSweepFrequency returns a boolean if a field has been set.

### GetPingTimeout

`func (o *ControllerAdvancedSettings) GetPingTimeout() int32`

GetPingTimeout returns the PingTimeout field if non-nil, zero value otherwise.

### GetPingTimeoutOk

`func (o *ControllerAdvancedSettings) GetPingTimeoutOk() (*int32, bool)`

GetPingTimeoutOk returns a tuple with the PingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingTimeout

`func (o *ControllerAdvancedSettings) SetPingTimeout(v int32)`

SetPingTimeout sets PingTimeout field to given value.

### HasPingTimeout

`func (o *ControllerAdvancedSettings) HasPingTimeout() bool`

HasPingTimeout returns a boolean if a field has been set.

### GetPortScanningEnabled

`func (o *ControllerAdvancedSettings) GetPortScanningEnabled() bool`

GetPortScanningEnabled returns the PortScanningEnabled field if non-nil, zero value otherwise.

### GetPortScanningEnabledOk

`func (o *ControllerAdvancedSettings) GetPortScanningEnabledOk() (*bool, bool)`

GetPortScanningEnabledOk returns a tuple with the PortScanningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortScanningEnabled

`func (o *ControllerAdvancedSettings) SetPortScanningEnabled(v bool)`

SetPortScanningEnabled sets PortScanningEnabled field to given value.

### HasPortScanningEnabled

`func (o *ControllerAdvancedSettings) HasPortScanningEnabled() bool`

HasPortScanningEnabled returns a boolean if a field has been set.

### GetRouteLimit

`func (o *ControllerAdvancedSettings) GetRouteLimit() int32`

GetRouteLimit returns the RouteLimit field if non-nil, zero value otherwise.

### GetRouteLimitOk

`func (o *ControllerAdvancedSettings) GetRouteLimitOk() (*int32, bool)`

GetRouteLimitOk returns a tuple with the RouteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteLimit

`func (o *ControllerAdvancedSettings) SetRouteLimit(v int32)`

SetRouteLimit sets RouteLimit field to given value.

### HasRouteLimit

`func (o *ControllerAdvancedSettings) HasRouteLimit() bool`

HasRouteLimit returns a boolean if a field has been set.

### GetSnmpDiscoveryVersion

`func (o *ControllerAdvancedSettings) GetSnmpDiscoveryVersion() string`

GetSnmpDiscoveryVersion returns the SnmpDiscoveryVersion field if non-nil, zero value otherwise.

### GetSnmpDiscoveryVersionOk

`func (o *ControllerAdvancedSettings) GetSnmpDiscoveryVersionOk() (*string, bool)`

GetSnmpDiscoveryVersionOk returns a tuple with the SnmpDiscoveryVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpDiscoveryVersion

`func (o *ControllerAdvancedSettings) SetSnmpDiscoveryVersion(v string)`

SetSnmpDiscoveryVersion sets SnmpDiscoveryVersion field to given value.

### HasSnmpDiscoveryVersion

`func (o *ControllerAdvancedSettings) HasSnmpDiscoveryVersion() bool`

HasSnmpDiscoveryVersion returns a boolean if a field has been set.

### GetSnmpv1DataCollectionFallback

`func (o *ControllerAdvancedSettings) GetSnmpv1DataCollectionFallback() string`

GetSnmpv1DataCollectionFallback returns the Snmpv1DataCollectionFallback field if non-nil, zero value otherwise.

### GetSnmpv1DataCollectionFallbackOk

`func (o *ControllerAdvancedSettings) GetSnmpv1DataCollectionFallbackOk() (*string, bool)`

GetSnmpv1DataCollectionFallbackOk returns a tuple with the Snmpv1DataCollectionFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv1DataCollectionFallback

`func (o *ControllerAdvancedSettings) SetSnmpv1DataCollectionFallback(v string)`

SetSnmpv1DataCollectionFallback sets Snmpv1DataCollectionFallback field to given value.

### HasSnmpv1DataCollectionFallback

`func (o *ControllerAdvancedSettings) HasSnmpv1DataCollectionFallback() bool`

HasSnmpv1DataCollectionFallback returns a boolean if a field has been set.

### GetTcpScanTechnique

`func (o *ControllerAdvancedSettings) GetTcpScanTechnique() string`

GetTcpScanTechnique returns the TcpScanTechnique field if non-nil, zero value otherwise.

### GetTcpScanTechniqueOk

`func (o *ControllerAdvancedSettings) GetTcpScanTechniqueOk() (*string, bool)`

GetTcpScanTechniqueOk returns a tuple with the TcpScanTechnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpScanTechnique

`func (o *ControllerAdvancedSettings) SetTcpScanTechnique(v string)`

SetTcpScanTechnique sets TcpScanTechnique field to given value.

### HasTcpScanTechnique

`func (o *ControllerAdvancedSettings) HasTcpScanTechnique() bool`

HasTcpScanTechnique returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


