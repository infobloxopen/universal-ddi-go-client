# DNSSECKeyDisplayStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSince** | Pointer to **time.Time** | Time when a key became active. | [optional] [readonly] 
**Algorithm** | Pointer to **int64** | Algorithm used for a key. | [optional] [readonly] 
**InactiveSince** | Pointer to **time.Time** | Time when a key became inactive. | [optional] [readonly] 
**KeyId** | Pointer to **int64** | Key id (tag). | [optional] [readonly] 
**PublishedSince** | Pointer to **time.Time** | Time when a key was published. | [optional] [readonly] 
**Size** | Pointer to **int64** | Key size in bits. | [optional] [readonly] 
**Status** | Pointer to **string** | Key status. Can be _UNKNOWN_, _PUBLISHED_, _ACTIVE_, _INACTIVE_. | [optional] [readonly] 
**Type** | Pointer to **string** | Key type (could be _KSK_ or _ZSK_). | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | Time when a key status was last updated by a host. | [optional] [readonly] 

## Methods

### NewDNSSECKeyDisplayStatus

`func NewDNSSECKeyDisplayStatus() *DNSSECKeyDisplayStatus`

NewDNSSECKeyDisplayStatus instantiates a new DNSSECKeyDisplayStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECKeyDisplayStatusWithDefaults

`func NewDNSSECKeyDisplayStatusWithDefaults() *DNSSECKeyDisplayStatus`

NewDNSSECKeyDisplayStatusWithDefaults instantiates a new DNSSECKeyDisplayStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSince

`func (o *DNSSECKeyDisplayStatus) GetActiveSince() time.Time`

GetActiveSince returns the ActiveSince field if non-nil, zero value otherwise.

### GetActiveSinceOk

`func (o *DNSSECKeyDisplayStatus) GetActiveSinceOk() (*time.Time, bool)`

GetActiveSinceOk returns a tuple with the ActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSince

`func (o *DNSSECKeyDisplayStatus) SetActiveSince(v time.Time)`

SetActiveSince sets ActiveSince field to given value.

### HasActiveSince

`func (o *DNSSECKeyDisplayStatus) HasActiveSince() bool`

HasActiveSince returns a boolean if a field has been set.

### GetAlgorithm

`func (o *DNSSECKeyDisplayStatus) GetAlgorithm() int64`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *DNSSECKeyDisplayStatus) GetAlgorithmOk() (*int64, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *DNSSECKeyDisplayStatus) SetAlgorithm(v int64)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *DNSSECKeyDisplayStatus) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetInactiveSince

`func (o *DNSSECKeyDisplayStatus) GetInactiveSince() time.Time`

GetInactiveSince returns the InactiveSince field if non-nil, zero value otherwise.

### GetInactiveSinceOk

`func (o *DNSSECKeyDisplayStatus) GetInactiveSinceOk() (*time.Time, bool)`

GetInactiveSinceOk returns a tuple with the InactiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveSince

`func (o *DNSSECKeyDisplayStatus) SetInactiveSince(v time.Time)`

SetInactiveSince sets InactiveSince field to given value.

### HasInactiveSince

`func (o *DNSSECKeyDisplayStatus) HasInactiveSince() bool`

HasInactiveSince returns a boolean if a field has been set.

### GetKeyId

`func (o *DNSSECKeyDisplayStatus) GetKeyId() int64`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *DNSSECKeyDisplayStatus) GetKeyIdOk() (*int64, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *DNSSECKeyDisplayStatus) SetKeyId(v int64)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *DNSSECKeyDisplayStatus) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetPublishedSince

`func (o *DNSSECKeyDisplayStatus) GetPublishedSince() time.Time`

GetPublishedSince returns the PublishedSince field if non-nil, zero value otherwise.

### GetPublishedSinceOk

`func (o *DNSSECKeyDisplayStatus) GetPublishedSinceOk() (*time.Time, bool)`

GetPublishedSinceOk returns a tuple with the PublishedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedSince

`func (o *DNSSECKeyDisplayStatus) SetPublishedSince(v time.Time)`

SetPublishedSince sets PublishedSince field to given value.

### HasPublishedSince

`func (o *DNSSECKeyDisplayStatus) HasPublishedSince() bool`

HasPublishedSince returns a boolean if a field has been set.

### GetSize

`func (o *DNSSECKeyDisplayStatus) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DNSSECKeyDisplayStatus) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DNSSECKeyDisplayStatus) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *DNSSECKeyDisplayStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *DNSSECKeyDisplayStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DNSSECKeyDisplayStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DNSSECKeyDisplayStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DNSSECKeyDisplayStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DNSSECKeyDisplayStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DNSSECKeyDisplayStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DNSSECKeyDisplayStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DNSSECKeyDisplayStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DNSSECKeyDisplayStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DNSSECKeyDisplayStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DNSSECKeyDisplayStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DNSSECKeyDisplayStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


