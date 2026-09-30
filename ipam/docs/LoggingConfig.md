# LoggingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyslogServers** | Pointer to [**[]SyslogServer**](SyslogServer.md) | Optional. List of syslog servers. | [optional] 

## Methods

### NewLoggingConfig

`func NewLoggingConfig() *LoggingConfig`

NewLoggingConfig instantiates a new LoggingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingConfigWithDefaults

`func NewLoggingConfigWithDefaults() *LoggingConfig`

NewLoggingConfigWithDefaults instantiates a new LoggingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyslogServers

`func (o *LoggingConfig) GetSyslogServers() []SyslogServer`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *LoggingConfig) GetSyslogServersOk() (*[]SyslogServer, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *LoggingConfig) SetSyslogServers(v []SyslogServer)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *LoggingConfig) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


