# GroupInIdentityV2UserInIdentityV2Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewGroupInIdentityV2UserInIdentityV2Account

`func NewGroupInIdentityV2UserInIdentityV2Account() *GroupInIdentityV2UserInIdentityV2Account`

NewGroupInIdentityV2UserInIdentityV2Account instantiates a new GroupInIdentityV2UserInIdentityV2Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupInIdentityV2UserInIdentityV2AccountWithDefaults

`func NewGroupInIdentityV2UserInIdentityV2AccountWithDefaults() *GroupInIdentityV2UserInIdentityV2Account`

NewGroupInIdentityV2UserInIdentityV2AccountWithDefaults instantiates a new GroupInIdentityV2UserInIdentityV2Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetAllowedCidrs() []string`

GetAllowedCidrs returns the AllowedCidrs field if non-nil, zero value otherwise.

### GetAllowedCidrsOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetAllowedCidrsOk() (*[]string, bool)`

GetAllowedCidrsOk returns a tuple with the AllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetAllowedCidrs(v []string)`

SetAllowedCidrs sets AllowedCidrs field to given value.

### HasAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasAllowedCidrs() bool`

HasAllowedCidrs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetEnableAllowedCidrs() bool`

GetEnableAllowedCidrs returns the EnableAllowedCidrs field if non-nil, zero value otherwise.

### GetEnableAllowedCidrsOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetEnableAllowedCidrsOk() (*bool, bool)`

GetEnableAllowedCidrsOk returns a tuple with the EnableAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetEnableAllowedCidrs(v bool)`

SetEnableAllowedCidrs sets EnableAllowedCidrs field to given value.

### HasEnableAllowedCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasEnableAllowedCidrs() bool`

HasEnableAllowedCidrs returns a boolean if a field has been set.

### GetEntitledFeatures

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetEntitledFeatures() map[string]map[string]interface{}`

GetEntitledFeatures returns the EntitledFeatures field if non-nil, zero value otherwise.

### GetEntitledFeaturesOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetEntitledFeaturesOk() (*map[string]map[string]interface{}, bool)`

GetEntitledFeaturesOk returns a tuple with the EntitledFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitledFeatures

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetEntitledFeatures(v map[string]map[string]interface{})`

SetEntitledFeatures sets EntitledFeatures field to given value.

### HasEntitledFeatures

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasEntitledFeatures() bool`

HasEntitledFeatures returns a boolean if a field has been set.

### GetId

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrigin

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetResetCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetResetCidrs() bool`

GetResetCidrs returns the ResetCidrs field if non-nil, zero value otherwise.

### GetResetCidrsOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetResetCidrsOk() (*bool, bool)`

GetResetCidrsOk returns a tuple with the ResetCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetResetCidrs(v bool)`

SetResetCidrs sets ResetCidrs field to given value.

### HasResetCidrs

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasResetCidrs() bool`

HasResetCidrs returns a boolean if a field has been set.

### GetScope

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserIds

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GroupInIdentityV2UserInIdentityV2Account) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GroupInIdentityV2UserInIdentityV2Account) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GroupInIdentityV2UserInIdentityV2Account) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


