# CurrentApikeysCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | Pointer to **string** | Expiry timestamp in RFC 3339 / ISO 8601 format (e.g. 2026-01-31T23:59:59Z). | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrentApikeysCreateRequest

`func NewCurrentApikeysCreateRequest() *CurrentApikeysCreateRequest`

NewCurrentApikeysCreateRequest instantiates a new CurrentApikeysCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentApikeysCreateRequestWithDefaults

`func NewCurrentApikeysCreateRequestWithDefaults() *CurrentApikeysCreateRequest`

NewCurrentApikeysCreateRequestWithDefaults instantiates a new CurrentApikeysCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *CurrentApikeysCreateRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CurrentApikeysCreateRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CurrentApikeysCreateRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CurrentApikeysCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *CurrentApikeysCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentApikeysCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentApikeysCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentApikeysCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


