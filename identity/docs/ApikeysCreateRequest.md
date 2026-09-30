# ApikeysCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompartmentId** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** | Expiry timestamp in RFC 3339 / ISO 8601 format (e.g. 2026-01-31T23:59:59Z). | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewApikeysCreateRequest

`func NewApikeysCreateRequest() *ApikeysCreateRequest`

NewApikeysCreateRequest instantiates a new ApikeysCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApikeysCreateRequestWithDefaults

`func NewApikeysCreateRequestWithDefaults() *ApikeysCreateRequest`

NewApikeysCreateRequestWithDefaults instantiates a new ApikeysCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompartmentId

`func (o *ApikeysCreateRequest) GetCompartmentId() string`

GetCompartmentId returns the CompartmentId field if non-nil, zero value otherwise.

### GetCompartmentIdOk

`func (o *ApikeysCreateRequest) GetCompartmentIdOk() (*string, bool)`

GetCompartmentIdOk returns a tuple with the CompartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentId

`func (o *ApikeysCreateRequest) SetCompartmentId(v string)`

SetCompartmentId sets CompartmentId field to given value.

### HasCompartmentId

`func (o *ApikeysCreateRequest) HasCompartmentId() bool`

HasCompartmentId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApikeysCreateRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApikeysCreateRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApikeysCreateRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApikeysCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetName

`func (o *ApikeysCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApikeysCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApikeysCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApikeysCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ApikeysCreateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApikeysCreateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApikeysCreateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApikeysCreateRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


