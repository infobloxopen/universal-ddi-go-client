# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountInIdentityV2Group**](AccountInIdentityV2Group.md) |  | [optional] 
**AccountId** | Pointer to **string** | The resource identifier. | [optional] 
**AllowedCidrs** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnableAllowedCidrs** | Pointer to **bool** |  | [optional] 
**EntitledFeatures** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**ResetCidrs** | Pointer to **bool** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserIds** | Pointer to **[]string** | The resource identifier. | [optional] 
**Users** | Pointer to [**[]UserInIdentityV2Group**](UserInIdentityV2Group.md) |  | [optional] 

## Methods

### NewGroup

`func NewGroup() *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Group) GetAccount() AccountInIdentityV2Group`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Group) GetAccountOk() (*AccountInIdentityV2Group, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Group) SetAccount(v AccountInIdentityV2Group)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Group) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountId

`func (o *Group) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Group) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Group) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Group) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllowedCidrs

`func (o *Group) GetAllowedCidrs() []string`

GetAllowedCidrs returns the AllowedCidrs field if non-nil, zero value otherwise.

### GetAllowedCidrsOk

`func (o *Group) GetAllowedCidrsOk() (*[]string, bool)`

GetAllowedCidrsOk returns a tuple with the AllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCidrs

`func (o *Group) SetAllowedCidrs(v []string)`

SetAllowedCidrs sets AllowedCidrs field to given value.

### HasAllowedCidrs

`func (o *Group) HasAllowedCidrs() bool`

HasAllowedCidrs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Group) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Group) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Group) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Group) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAllowedCidrs

`func (o *Group) GetEnableAllowedCidrs() bool`

GetEnableAllowedCidrs returns the EnableAllowedCidrs field if non-nil, zero value otherwise.

### GetEnableAllowedCidrsOk

`func (o *Group) GetEnableAllowedCidrsOk() (*bool, bool)`

GetEnableAllowedCidrsOk returns a tuple with the EnableAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowedCidrs

`func (o *Group) SetEnableAllowedCidrs(v bool)`

SetEnableAllowedCidrs sets EnableAllowedCidrs field to given value.

### HasEnableAllowedCidrs

`func (o *Group) HasEnableAllowedCidrs() bool`

HasEnableAllowedCidrs returns a boolean if a field has been set.

### GetEntitledFeatures

`func (o *Group) GetEntitledFeatures() map[string]map[string]interface{}`

GetEntitledFeatures returns the EntitledFeatures field if non-nil, zero value otherwise.

### GetEntitledFeaturesOk

`func (o *Group) GetEntitledFeaturesOk() (*map[string]map[string]interface{}, bool)`

GetEntitledFeaturesOk returns a tuple with the EntitledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledFeatures

`func (o *Group) SetEntitledFeatures(v map[string]map[string]interface{})`

SetEntitledFeatures sets EntitledFeatures field to given value.

### HasEntitledFeatures

`func (o *Group) HasEntitledFeatures() bool`

HasEntitledFeatures returns a boolean if a field has been set.

### GetId

`func (o *Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *Group) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Group) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Group) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Group) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *Group) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Group) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Group) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *Group) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetResetCidrs

`func (o *Group) GetResetCidrs() bool`

GetResetCidrs returns the ResetCidrs field if non-nil, zero value otherwise.

### GetResetCidrsOk

`func (o *Group) GetResetCidrsOk() (*bool, bool)`

GetResetCidrsOk returns a tuple with the ResetCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCidrs

`func (o *Group) SetResetCidrs(v bool)`

SetResetCidrs sets ResetCidrs field to given value.

### HasResetCidrs

`func (o *Group) HasResetCidrs() bool`

HasResetCidrs returns a boolean if a field has been set.

### GetScope

`func (o *Group) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Group) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Group) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Group) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Group) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Group) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Group) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Group) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserIds

`func (o *Group) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *Group) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *Group) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *Group) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### GetUsers

`func (o *Group) GetUsers() []UserInIdentityV2Group`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Group) GetUsersOk() (*[]UserInIdentityV2Group, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Group) SetUsers(v []UserInIdentityV2Group)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Group) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


