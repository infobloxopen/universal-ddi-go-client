# SessionGetCurrentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewSessionGetCurrentResponse

`func NewSessionGetCurrentResponse() *SessionGetCurrentResponse`

NewSessionGetCurrentResponse instantiates a new SessionGetCurrentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionGetCurrentResponseWithDefaults

`func NewSessionGetCurrentResponseWithDefaults() *SessionGetCurrentResponse`

NewSessionGetCurrentResponseWithDefaults instantiates a new SessionGetCurrentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *SessionGetCurrentResponse) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SessionGetCurrentResponse) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SessionGetCurrentResponse) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SessionGetCurrentResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetGroups

`func (o *SessionGetCurrentResponse) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *SessionGetCurrentResponse) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *SessionGetCurrentResponse) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *SessionGetCurrentResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUser

`func (o *SessionGetCurrentResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SessionGetCurrentResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SessionGetCurrentResponse) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *SessionGetCurrentResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


