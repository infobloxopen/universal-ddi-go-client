# DestinationSyslog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CaCertificate** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**InsecureMode** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OutputDataFormat** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int64** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**TransportProtocol** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDestinationSyslog

`func NewDestinationSyslog() *DestinationSyslog`

NewDestinationSyslog instantiates a new DestinationSyslog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSyslogWithDefaults

`func NewDestinationSyslogWithDefaults() *DestinationSyslog`

NewDestinationSyslogWithDefaults instantiates a new DestinationSyslog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationSyslog) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationSyslog) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationSyslog) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationSyslog) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCaCertificate

`func (o *DestinationSyslog) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *DestinationSyslog) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *DestinationSyslog) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *DestinationSyslog) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationSyslog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationSyslog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationSyslog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationSyslog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationSyslog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationSyslog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationSyslog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationSyslog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationSyslog) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationSyslog) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationSyslog) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationSyslog) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationSyslog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationSyslog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationSyslog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationSyslog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsecureMode

`func (o *DestinationSyslog) GetInsecureMode() bool`

GetInsecureMode returns the InsecureMode field if non-nil, zero value otherwise.

### GetInsecureModeOk

`func (o *DestinationSyslog) GetInsecureModeOk() (*bool, bool)`

GetInsecureModeOk returns a tuple with the InsecureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureMode

`func (o *DestinationSyslog) SetInsecureMode(v bool)`

SetInsecureMode sets InsecureMode field to given value.

### HasInsecureMode

`func (o *DestinationSyslog) HasInsecureMode() bool`

HasInsecureMode returns a boolean if a field has been set.

### GetName

`func (o *DestinationSyslog) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSyslog) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSyslog) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationSyslog) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputDataFormat

`func (o *DestinationSyslog) GetOutputDataFormat() string`

GetOutputDataFormat returns the OutputDataFormat field if non-nil, zero value otherwise.

### GetOutputDataFormatOk

`func (o *DestinationSyslog) GetOutputDataFormatOk() (*string, bool)`

GetOutputDataFormatOk returns a tuple with the OutputDataFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDataFormat

`func (o *DestinationSyslog) SetOutputDataFormat(v string)`

SetOutputDataFormat sets OutputDataFormat field to given value.

### HasOutputDataFormat

`func (o *DestinationSyslog) HasOutputDataFormat() bool`

HasOutputDataFormat returns a boolean if a field has been set.

### GetPort

`func (o *DestinationSyslog) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DestinationSyslog) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DestinationSyslog) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *DestinationSyslog) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTags

`func (o *DestinationSyslog) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationSyslog) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationSyslog) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationSyslog) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransportProtocol

`func (o *DestinationSyslog) GetTransportProtocol() string`

GetTransportProtocol returns the TransportProtocol field if non-nil, zero value otherwise.

### GetTransportProtocolOk

`func (o *DestinationSyslog) GetTransportProtocolOk() (*string, bool)`

GetTransportProtocolOk returns a tuple with the TransportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportProtocol

`func (o *DestinationSyslog) SetTransportProtocol(v string)`

SetTransportProtocol sets TransportProtocol field to given value.

### HasTransportProtocol

`func (o *DestinationSyslog) HasTransportProtocol() bool`

HasTransportProtocol returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationSyslog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationSyslog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationSyslog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationSyslog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


