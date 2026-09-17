# DNSSECSigningKeyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **int64** | Algorithm used for the key.  Allowed values: * _5_ - RSASHA1 * _7_ - NSEC3RSASHA1 (RSASHA1-NSEC3-SHA1) * _8_ - RSASHA256 * _10_ - RSASHA512 * _13_ - ECDSAP256SHA256 * _14_ - ECDSAP384SHA384 * _15_ - ED25519 * _16_ - ED448  Defaults to _8_ (RSASHA256). | [optional] 
**Size** | Pointer to **int64** | Key size in bits.  Value should be within allowed range for _algorithm_: * _RSASHA1_: 1024..4096 * _NSEC3RSASHA1_: 1024..4096 * _RSASHA256_: 1024..4096 * _RSASHA512_: 1024..4096 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456  Defaults are based on the _algorithm_ and _type_: For KSK:  * _RSASHA1_: 2048 * _NSEC3RSASHA1_: 2048 * _RSASHA256_: 2048 * _RSASHA512_: 2048 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456  For ZSK:  * _RSASHA1_: 1024 * _NSEC3RSASHA1_: 1024 * _RSASHA256_: 1024 * _RSASHA512_: 1024 * _ECDSAP256SHA256_: 256 * _ECDSAP384SHA384_: 384 * _ED25519_: 256 * _ED448_: 456 | [optional] 
**Type** | **string** | Key type.  Allowed values: * _KSK_: Key-Signing Key, used to sign DNSKEY records. * _ZSK_: Zone-Signing Key, used to sign all other records in the zone. | 

## Methods

### NewDNSSECSigningKeyPolicy

`func NewDNSSECSigningKeyPolicy(type_ string, ) *DNSSECSigningKeyPolicy`

NewDNSSECSigningKeyPolicy instantiates a new DNSSECSigningKeyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECSigningKeyPolicyWithDefaults

`func NewDNSSECSigningKeyPolicyWithDefaults() *DNSSECSigningKeyPolicy`

NewDNSSECSigningKeyPolicyWithDefaults instantiates a new DNSSECSigningKeyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *DNSSECSigningKeyPolicy) GetAlgorithm() int64`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DNSSECSigningKeyPolicy) GetAlgorithmOk() (*int64, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DNSSECSigningKeyPolicy) SetAlgorithm(v int64)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *DNSSECSigningKeyPolicy) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetSize

`func (o *DNSSECSigningKeyPolicy) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DNSSECSigningKeyPolicy) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DNSSECSigningKeyPolicy) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DNSSECSigningKeyPolicy) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *DNSSECSigningKeyPolicy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSSECSigningKeyPolicy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSSECSigningKeyPolicy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


