# DNSSECKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to **int64** | Algorithm used for the key. | [optional] [readonly] 
**KeyId** | Pointer to **int64** | Key ID (also known as Key Tag). | [optional] [readonly] 
**NextRolloverEvent** | Pointer to **time.Time** | Next Rollover Event Time. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | Public key in Base64 format. | [optional] [readonly] 
**Size** | Pointer to **int64** | Key size in bits. | [optional] [readonly] 
**Type** | Pointer to **string** | Key type.  Allowed values: * _KSK_: Key-Signing Key, used to sign DNSKEY records. * _ZSK_: Zone-Signing Key, used to sign all other records in the zone. | [optional] [readonly] 

## Methods

### NewDNSSECKey

`func NewDNSSECKey() *DNSSECKey`

NewDNSSECKey instantiates a new DNSSECKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECKeyWithDefaults

`func NewDNSSECKeyWithDefaults() *DNSSECKey`

NewDNSSECKeyWithDefaults instantiates a new DNSSECKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *DNSSECKey) GetAlgorithm() int64`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DNSSECKey) GetAlgorithmOk() (*int64, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DNSSECKey) SetAlgorithm(v int64)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *DNSSECKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKeyId

`func (o *DNSSECKey) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *DNSSECKey) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *DNSSECKey) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *DNSSECKey) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetNextRolloverEvent

`func (o *DNSSECKey) GetNextRolloverEvent() time.Time`

GetNextRolloverEvent returns the NextRolloverEvent field if non-nil, zero value otherwise.

### GetNextRolloverEventOk

`func (o *DNSSECKey) GetNextRolloverEventOk() (*time.Time, bool)`

GetNextRolloverEventOk returns a tuple with the NextRolloverEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRolloverEvent

`func (o *DNSSECKey) SetNextRolloverEvent(v time.Time)`

SetNextRolloverEvent sets NextRolloverEvent field to given value.

### HasNextRolloverEvent

`func (o *DNSSECKey) HasNextRolloverEvent() bool`

HasNextRolloverEvent returns a boolean if a field has been set.

### GetPublicKey

`func (o *DNSSECKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *DNSSECKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *DNSSECKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *DNSSECKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSize

`func (o *DNSSECKey) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DNSSECKey) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DNSSECKey) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DNSSECKey) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *DNSSECKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSSECKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSSECKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSSECKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


