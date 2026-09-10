# Nameserver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Optional. Required only if _origin_ is _external_. IP Address of the nameserver. | [optional] 
**Fqdn** | Pointer to **string** | Optional. Required only if _origin_ is _external_. FQDN of the nameserver. | [optional] 
**Host** | Pointer to **string** | The resource identifier. | [optional] 
**Origin** | Pointer to **string** | Allowed values: * _external_, * _nios-x_, * _nios_ | [optional] 
**ProtocolFqdn** | Pointer to **string** | FQDN of the nameserver in punycode. | [optional] [readonly] 
**Role** | **string** | Allowed values: * _primary_, * _secondary_ | 
**Stealth** | Pointer to **bool** | If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_ | [optional] 
**TsigEnabled** | Pointer to **bool** | Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from a primary. | [optional] 
**TsigKey** | Pointer to [**TSIGKey**](TSIGKey.md) | Optional. TSIG key.  Error if empty while _tsig_enabled_ is _true_. | [optional] 

## Methods

### NewNameserver

`func NewNameserver(role string, ) *Nameserver`

NewNameserver instantiates a new Nameserver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNameserverWithDefaults

`func NewNameserverWithDefaults() *Nameserver`

NewNameserverWithDefaults instantiates a new Nameserver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Nameserver) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Nameserver) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Nameserver) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Nameserver) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *Nameserver) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Nameserver) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Nameserver) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Nameserver) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetHost

`func (o *Nameserver) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Nameserver) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Nameserver) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Nameserver) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOrigin

`func (o *Nameserver) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Nameserver) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Nameserver) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Nameserver) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *Nameserver) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *Nameserver) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *Nameserver) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *Nameserver) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetRole

`func (o *Nameserver) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Nameserver) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Nameserver) SetRole(v string)`

SetRole sets Role field to given value.


### GetStealth

`func (o *Nameserver) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *Nameserver) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *Nameserver) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *Nameserver) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *Nameserver) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *Nameserver) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *Nameserver) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *Nameserver) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *Nameserver) GetTsigKey() TSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *Nameserver) GetTsigKeyOk() (*TSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *Nameserver) SetTsigKey(v TSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *Nameserver) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


