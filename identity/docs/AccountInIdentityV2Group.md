# AccountInIdentityV2Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**AdminUser** | Pointer to [**UserInIdentityV2AccountInIdentityV2Group**](UserInIdentityV2AccountInIdentityV2Group.md) |  | [optional] 
**AdminUserId** | Pointer to **string** | The resource identifier. | [optional] 
**AxurAccountId** | Pointer to **string** |  | [optional] 
**CompanyDomain** | Pointer to **string** |  | [optional] 
**CompanyNumber** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CspId** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IdpAuthnEnabled** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentAccountId** | Pointer to **string** | The resource identifier. | [optional] 
**SfdcAccountId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateChangedAt** | Pointer to **time.Time** |  | [optional] 
**StorageId** | Pointer to **int32** |  | [optional] 
**SubresourceName** | Pointer to **string** |  | [optional] 
**SupportEnabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountInIdentityV2Group

`func NewAccountInIdentityV2Group() *AccountInIdentityV2Group`

NewAccountInIdentityV2Group instantiates a new AccountInIdentityV2Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInIdentityV2GroupWithDefaults

`func NewAccountInIdentityV2GroupWithDefaults() *AccountInIdentityV2Group`

NewAccountInIdentityV2GroupWithDefaults instantiates a new AccountInIdentityV2Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *AccountInIdentityV2Group) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountInIdentityV2Group) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountInIdentityV2Group) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountInIdentityV2Group) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminUser

`func (o *AccountInIdentityV2Group) GetAdminUser() UserInIdentityV2AccountInIdentityV2Group`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *AccountInIdentityV2Group) GetAdminUserOk() (*UserInIdentityV2AccountInIdentityV2Group, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *AccountInIdentityV2Group) SetAdminUser(v UserInIdentityV2AccountInIdentityV2Group)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *AccountInIdentityV2Group) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetAdminUserId

`func (o *AccountInIdentityV2Group) GetAdminUserId() string`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *AccountInIdentityV2Group) GetAdminUserIdOk() (*string, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *AccountInIdentityV2Group) SetAdminUserId(v string)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *AccountInIdentityV2Group) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetAxurAccountId

`func (o *AccountInIdentityV2Group) GetAxurAccountId() string`

GetAxurAccountId returns the AxurAccountId field if non-nil, zero value otherwise.

### GetAxurAccountIdOk

`func (o *AccountInIdentityV2Group) GetAxurAccountIdOk() (*string, bool)`

GetAxurAccountIdOk returns a tuple with the AxurAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxurAccountId

`func (o *AccountInIdentityV2Group) SetAxurAccountId(v string)`

SetAxurAccountId sets AxurAccountId field to given value.

### HasAxurAccountId

`func (o *AccountInIdentityV2Group) HasAxurAccountId() bool`

HasAxurAccountId returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *AccountInIdentityV2Group) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *AccountInIdentityV2Group) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *AccountInIdentityV2Group) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *AccountInIdentityV2Group) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCompanyNumber

`func (o *AccountInIdentityV2Group) GetCompanyNumber() int32`

GetCompanyNumber returns the CompanyNumber field if non-nil, zero value otherwise.

### GetCompanyNumberOk

`func (o *AccountInIdentityV2Group) GetCompanyNumberOk() (*int32, bool)`

GetCompanyNumberOk returns a tuple with the CompanyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNumber

`func (o *AccountInIdentityV2Group) SetCompanyNumber(v int32)`

SetCompanyNumber sets CompanyNumber field to given value.

### HasCompanyNumber

`func (o *AccountInIdentityV2Group) HasCompanyNumber() bool`

HasCompanyNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountInIdentityV2Group) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountInIdentityV2Group) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountInIdentityV2Group) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountInIdentityV2Group) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AccountInIdentityV2Group) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountInIdentityV2Group) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountInIdentityV2Group) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccountInIdentityV2Group) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCspId

`func (o *AccountInIdentityV2Group) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *AccountInIdentityV2Group) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *AccountInIdentityV2Group) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *AccountInIdentityV2Group) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AccountInIdentityV2Group) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AccountInIdentityV2Group) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AccountInIdentityV2Group) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AccountInIdentityV2Group) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AccountInIdentityV2Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountInIdentityV2Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountInIdentityV2Group) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountInIdentityV2Group) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *AccountInIdentityV2Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountInIdentityV2Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountInIdentityV2Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountInIdentityV2Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpAuthnEnabled

`func (o *AccountInIdentityV2Group) GetIdpAuthnEnabled() bool`

GetIdpAuthnEnabled returns the IdpAuthnEnabled field if non-nil, zero value otherwise.

### GetIdpAuthnEnabledOk

`func (o *AccountInIdentityV2Group) GetIdpAuthnEnabledOk() (*bool, bool)`

GetIdpAuthnEnabledOk returns a tuple with the IdpAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAuthnEnabled

`func (o *AccountInIdentityV2Group) SetIdpAuthnEnabled(v bool)`

SetIdpAuthnEnabled sets IdpAuthnEnabled field to given value.

### HasIdpAuthnEnabled

`func (o *AccountInIdentityV2Group) HasIdpAuthnEnabled() bool`

HasIdpAuthnEnabled returns a boolean if a field has been set.

### GetLabels

`func (o *AccountInIdentityV2Group) GetLabels() map[string]map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AccountInIdentityV2Group) GetLabelsOk() (*map[string]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AccountInIdentityV2Group) SetLabels(v map[string]map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AccountInIdentityV2Group) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *AccountInIdentityV2Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountInIdentityV2Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountInIdentityV2Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountInIdentityV2Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentAccountId

`func (o *AccountInIdentityV2Group) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *AccountInIdentityV2Group) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *AccountInIdentityV2Group) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *AccountInIdentityV2Group) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetSfdcAccountId

`func (o *AccountInIdentityV2Group) GetSfdcAccountId() string`

GetSfdcAccountId returns the SfdcAccountId field if non-nil, zero value otherwise.

### GetSfdcAccountIdOk

`func (o *AccountInIdentityV2Group) GetSfdcAccountIdOk() (*string, bool)`

GetSfdcAccountIdOk returns a tuple with the SfdcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcAccountId

`func (o *AccountInIdentityV2Group) SetSfdcAccountId(v string)`

SetSfdcAccountId sets SfdcAccountId field to given value.

### HasSfdcAccountId

`func (o *AccountInIdentityV2Group) HasSfdcAccountId() bool`

HasSfdcAccountId returns a boolean if a field has been set.

### GetState

`func (o *AccountInIdentityV2Group) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccountInIdentityV2Group) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccountInIdentityV2Group) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *AccountInIdentityV2Group) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *AccountInIdentityV2Group) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *AccountInIdentityV2Group) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *AccountInIdentityV2Group) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *AccountInIdentityV2Group) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStorageId

`func (o *AccountInIdentityV2Group) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *AccountInIdentityV2Group) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *AccountInIdentityV2Group) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *AccountInIdentityV2Group) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetSubresourceName

`func (o *AccountInIdentityV2Group) GetSubresourceName() string`

GetSubresourceName returns the SubresourceName field if non-nil, zero value otherwise.

### GetSubresourceNameOk

`func (o *AccountInIdentityV2Group) GetSubresourceNameOk() (*string, bool)`

GetSubresourceNameOk returns a tuple with the SubresourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceName

`func (o *AccountInIdentityV2Group) SetSubresourceName(v string)`

SetSubresourceName sets SubresourceName field to given value.

### HasSubresourceName

`func (o *AccountInIdentityV2Group) HasSubresourceName() bool`

HasSubresourceName returns a boolean if a field has been set.

### GetSupportEnabled

`func (o *AccountInIdentityV2Group) GetSupportEnabled() bool`

GetSupportEnabled returns the SupportEnabled field if non-nil, zero value otherwise.

### GetSupportEnabledOk

`func (o *AccountInIdentityV2Group) GetSupportEnabledOk() (*bool, bool)`

GetSupportEnabledOk returns a tuple with the SupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEnabled

`func (o *AccountInIdentityV2Group) SetSupportEnabled(v bool)`

SetSupportEnabled sets SupportEnabled field to given value.

### HasSupportEnabled

`func (o *AccountInIdentityV2Group) HasSupportEnabled() bool`

HasSupportEnabled returns a boolean if a field has been set.

### GetTags

`func (o *AccountInIdentityV2Group) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AccountInIdentityV2Group) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AccountInIdentityV2Group) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AccountInIdentityV2Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountInIdentityV2Group) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountInIdentityV2Group) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountInIdentityV2Group) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountInIdentityV2Group) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


