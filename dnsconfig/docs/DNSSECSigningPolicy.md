# DNSSECSigningPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]DNSSECSigningKeyPolicy**](DNSSECSigningKeyPolicy.md) | Key settings. This defines configuration for DNSSEC keys. combination of both Key-Signing Keys (KSK) and Zone-Signing Keys (ZSK).  Defaults to empty. | [optional] 
**KskAutomaticRolloverEnabled** | Pointer to **bool** | Flag indicating if KSK rollover should be automatic.  Defaults to _false_. | [optional] 
**KskNotificationEventTrigger** | Pointer to **string** | Option controls when notifications are sent for KSK rollover events.  Valid values are: * _NO_EVENTS_ - no notifications are sent * _ALL_EVENTS_ - any time KSK is rolled over, a notification is sent * _MANUAL_DS_UPDATE_EVENTS_ - a notification is sent only when DS record needs to be updated manually  Defaults to _NO_EVENTS_ | [optional] 
**KskRolloverInterval** | Pointer to **int64** | KSK rollover interval in seconds.  Used to determine how often the Key-Signing Keys should be rotated. Examples: 31536000 (1 year), 7776000 (90 days), 2592000 (30 days)  Unsigned integer, min 0.  Defaults to 31536000 (1 year). | [optional] 
**Nsec3Iterations** | Pointer to **int64** | Optional. Number of additional hash iterations to perform. Increasing this value slows down both authoritative server when signing and recursive servers when verifying, but also slows down attacker&#39;s dictionary attacks.  IMPORTANT: Changing the settings for the NSEC3 number of iterations is not recommended.  Unsigned integer, min 0 max 65535.  Defaults to _0_. | [optional] 
**Nsec3SaltLength** | Pointer to **int64** | Optional. Minimum length for NSEC3 salt in octets. Used to add entropy to the hash function to defend against pre-calculated attacks.  Unsigned integer, min 0 max 65535.  Defaults to _0_. | [optional] 
**NsecType** | Pointer to **string** | DNSSEC resource record type for nonexistent proof. This controls which type will be used to provide proof of nonexistence.  Allowed values: * _NSEC_ * _NSEC3_  Defaults to _NSEC_ | [optional] 
**ZskRolloverInterval** | Pointer to **int64** | ZSK rollover interval in seconds.  Used to determine how often the Zone-Signing Keys should be rotated. Examples: 2592000 (30 days), 1209600 (14 days)  Unsigned integer, min 0.  Defaults to 2592000 (30 days). | [optional] 
**ZskSignatureValidity** | Pointer to **int64** | ZSK signature validity period in seconds.  Determines how long DNSSEC signatures remain valid. Examples: 1209600 (14 days), 604800 (7 days)  Unsigned integer, min 0.  Defaults to 1209600 (14 days). | [optional] 

## Methods

### NewDNSSECSigningPolicy

`func NewDNSSECSigningPolicy() *DNSSECSigningPolicy`

NewDNSSECSigningPolicy instantiates a new DNSSECSigningPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSSECSigningPolicyWithDefaults

`func NewDNSSECSigningPolicyWithDefaults() *DNSSECSigningPolicy`

NewDNSSECSigningPolicyWithDefaults instantiates a new DNSSECSigningPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *DNSSECSigningPolicy) GetKeys() []DNSSECSigningKeyPolicy`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *DNSSECSigningPolicy) GetKeysOk() (*[]DNSSECSigningKeyPolicy, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *DNSSECSigningPolicy) SetKeys(v []DNSSECSigningKeyPolicy)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *DNSSECSigningPolicy) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetKskAutomaticRolloverEnabled

`func (o *DNSSECSigningPolicy) GetKskAutomaticRolloverEnabled() bool`

GetKskAutomaticRolloverEnabled returns the KskAutomaticRolloverEnabled field if non-nil, zero value otherwise.

### GetKskAutomaticRolloverEnabledOk

`func (o *DNSSECSigningPolicy) GetKskAutomaticRolloverEnabledOk() (*bool, bool)`

GetKskAutomaticRolloverEnabledOk returns a tuple with the KskAutomaticRolloverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskAutomaticRolloverEnabled

`func (o *DNSSECSigningPolicy) SetKskAutomaticRolloverEnabled(v bool)`

SetKskAutomaticRolloverEnabled sets KskAutomaticRolloverEnabled field to given value.

### HasKskAutomaticRolloverEnabled

`func (o *DNSSECSigningPolicy) HasKskAutomaticRolloverEnabled() bool`

HasKskAutomaticRolloverEnabled returns a boolean if a field has been set.

### GetKskNotificationEventTrigger

`func (o *DNSSECSigningPolicy) GetKskNotificationEventTrigger() string`

GetKskNotificationEventTrigger returns the KskNotificationEventTrigger field if non-nil, zero value otherwise.

### GetKskNotificationEventTriggerOk

`func (o *DNSSECSigningPolicy) GetKskNotificationEventTriggerOk() (*string, bool)`

GetKskNotificationEventTriggerOk returns a tuple with the KskNotificationEventTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskNotificationEventTrigger

`func (o *DNSSECSigningPolicy) SetKskNotificationEventTrigger(v string)`

SetKskNotificationEventTrigger sets KskNotificationEventTrigger field to given value.

### HasKskNotificationEventTrigger

`func (o *DNSSECSigningPolicy) HasKskNotificationEventTrigger() bool`

HasKskNotificationEventTrigger returns a boolean if a field has been set.

### GetKskRolloverInterval

`func (o *DNSSECSigningPolicy) GetKskRolloverInterval() int64`

GetKskRolloverInterval returns the KskRolloverInterval field if non-nil, zero value otherwise.

### GetKskRolloverIntervalOk

`func (o *DNSSECSigningPolicy) GetKskRolloverIntervalOk() (*int64, bool)`

GetKskRolloverIntervalOk returns a tuple with the KskRolloverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskRolloverInterval

`func (o *DNSSECSigningPolicy) SetKskRolloverInterval(v int64)`

SetKskRolloverInterval sets KskRolloverInterval field to given value.

### HasKskRolloverInterval

`func (o *DNSSECSigningPolicy) HasKskRolloverInterval() bool`

HasKskRolloverInterval returns a boolean if a field has been set.

### GetNsec3Iterations

`func (o *DNSSECSigningPolicy) GetNsec3Iterations() int64`

GetNsec3Iterations returns the Nsec3Iterations field if non-nil, zero value otherwise.

### GetNsec3IterationsOk

`func (o *DNSSECSigningPolicy) GetNsec3IterationsOk() (*int64, bool)`

GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3Iterations

`func (o *DNSSECSigningPolicy) SetNsec3Iterations(v int64)`

SetNsec3Iterations sets Nsec3Iterations field to given value.

### HasNsec3Iterations

`func (o *DNSSECSigningPolicy) HasNsec3Iterations() bool`

HasNsec3Iterations returns a boolean if a field has been set.

### GetNsec3SaltLength

`func (o *DNSSECSigningPolicy) GetNsec3SaltLength() int64`

GetNsec3SaltLength returns the Nsec3SaltLength field if non-nil, zero value otherwise.

### GetNsec3SaltLengthOk

`func (o *DNSSECSigningPolicy) GetNsec3SaltLengthOk() (*int64, bool)`

GetNsec3SaltLengthOk returns a tuple with the Nsec3SaltLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltLength

`func (o *DNSSECSigningPolicy) SetNsec3SaltLength(v int64)`

SetNsec3SaltLength sets Nsec3SaltLength field to given value.

### HasNsec3SaltLength

`func (o *DNSSECSigningPolicy) HasNsec3SaltLength() bool`

HasNsec3SaltLength returns a boolean if a field has been set.

### GetNsecType

`func (o *DNSSECSigningPolicy) GetNsecType() string`

GetNsecType returns the NsecType field if non-nil, zero value otherwise.

### GetNsecTypeOk

`func (o *DNSSECSigningPolicy) GetNsecTypeOk() (*string, bool)`

GetNsecTypeOk returns a tuple with the NsecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsecType

`func (o *DNSSECSigningPolicy) SetNsecType(v string)`

SetNsecType sets NsecType field to given value.

### HasNsecType

`func (o *DNSSECSigningPolicy) HasNsecType() bool`

HasNsecType returns a boolean if a field has been set.

### GetZskRolloverInterval

`func (o *DNSSECSigningPolicy) GetZskRolloverInterval() int64`

GetZskRolloverInterval returns the ZskRolloverInterval field if non-nil, zero value otherwise.

### GetZskRolloverIntervalOk

`func (o *DNSSECSigningPolicy) GetZskRolloverIntervalOk() (*int64, bool)`

GetZskRolloverIntervalOk returns a tuple with the ZskRolloverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskRolloverInterval

`func (o *DNSSECSigningPolicy) SetZskRolloverInterval(v int64)`

SetZskRolloverInterval sets ZskRolloverInterval field to given value.

### HasZskRolloverInterval

`func (o *DNSSECSigningPolicy) HasZskRolloverInterval() bool`

HasZskRolloverInterval returns a boolean if a field has been set.

### GetZskSignatureValidity

`func (o *DNSSECSigningPolicy) GetZskSignatureValidity() int64`

GetZskSignatureValidity returns the ZskSignatureValidity field if non-nil, zero value otherwise.

### GetZskSignatureValidityOk

`func (o *DNSSECSigningPolicy) GetZskSignatureValidityOk() (*int64, bool)`

GetZskSignatureValidityOk returns a tuple with the ZskSignatureValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskSignatureValidity

`func (o *DNSSECSigningPolicy) SetZskSignatureValidity(v int64)`

SetZskSignatureValidity sets ZskSignatureValidity field to given value.

### HasZskSignatureValidity

`func (o *DNSSECSigningPolicy) HasZskSignatureValidity() bool`

HasZskSignatureValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


