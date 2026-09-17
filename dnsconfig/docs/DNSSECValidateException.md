# DNSSECValidateException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | Domain name to exclude from DNSSEC validation. Treated as absolute (fully qualified) with or without a trailing dot. Must not be empty. It can contain U-labels for IDNs and NR-LDH labels but is not supposed to contain A-labels. | 
**ProtocolDomain** | Pointer to **string** | Standard DNS presentation of &#39;domain&#39;, only containing A-labels and NR-LDH labels. Will always end with a trailing dot. | [optional] [readonly] 

## Methods

### NewDNSSECValidateException

`func NewDNSSECValidateException(domain string, ) *DNSSECValidateException`

NewDNSSECValidateException instantiates a new DNSSECValidateException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECValidateExceptionWithDefaults

`func NewDNSSECValidateExceptionWithDefaults() *DNSSECValidateException`

NewDNSSECValidateExceptionWithDefaults instantiates a new DNSSECValidateException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DNSSECValidateException) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DNSSECValidateException) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DNSSECValidateException) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetProtocolDomain

`func (o *DNSSECValidateException) GetProtocolDomain() string`

GetProtocolDomain returns the ProtocolDomain field if non-nil, zero value otherwise.

### GetProtocolDomainOk

`func (o *DNSSECValidateException) GetProtocolDomainOk() (*string, bool)`

GetProtocolDomainOk returns a tuple with the ProtocolDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolDomain

`func (o *DNSSECValidateException) SetProtocolDomain(v string)`

SetProtocolDomain sets ProtocolDomain field to given value.

### HasProtocolDomain

`func (o *DNSSECValidateException) HasProtocolDomain() bool`

HasProtocolDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


