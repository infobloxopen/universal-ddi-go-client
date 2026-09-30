# SessionRenewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expires** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Jwt** | Pointer to **string** |  | [optional] 

## Methods

### NewSessionRenewResponse

`func NewSessionRenewResponse() *SessionRenewResponse`

NewSessionRenewResponse instantiates a new SessionRenewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRenewResponseWithDefaults

`func NewSessionRenewResponseWithDefaults() *SessionRenewResponse`

NewSessionRenewResponseWithDefaults instantiates a new SessionRenewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpires

`func (o *SessionRenewResponse) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *SessionRenewResponse) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *SessionRenewResponse) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *SessionRenewResponse) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExpiresAt

`func (o *SessionRenewResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *SessionRenewResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *SessionRenewResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *SessionRenewResponse) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetJwt

`func (o *SessionRenewResponse) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *SessionRenewResponse) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *SessionRenewResponse) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *SessionRenewResponse) HasJwt() bool`

HasJwt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


