# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**AdminUser** | Pointer to [**UserInIdentityV2Account**](UserInIdentityV2Account.md) |  | [optional] 
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

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Account) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Account) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Account) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Account) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminUser

`func (o *Account) GetAdminUser() UserInIdentityV2Account`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *Account) GetAdminUserOk() (*UserInIdentityV2Account, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *Account) SetAdminUser(v UserInIdentityV2Account)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *Account) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetAdminUserId

`func (o *Account) GetAdminUserId() string`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *Account) GetAdminUserIdOk() (*string, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *Account) SetAdminUserId(v string)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *Account) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetAxurAccountId

`func (o *Account) GetAxurAccountId() string`

GetAxurAccountId returns the AxurAccountId field if non-nil, zero value otherwise.

### GetAxurAccountIdOk

`func (o *Account) GetAxurAccountIdOk() (*string, bool)`

GetAxurAccountIdOk returns a tuple with the AxurAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxurAccountId

`func (o *Account) SetAxurAccountId(v string)`

SetAxurAccountId sets AxurAccountId field to given value.

### HasAxurAccountId

`func (o *Account) HasAxurAccountId() bool`

HasAxurAccountId returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *Account) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *Account) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *Account) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *Account) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCompanyNumber

`func (o *Account) GetCompanyNumber() int32`

GetCompanyNumber returns the CompanyNumber field if non-nil, zero value otherwise.

### GetCompanyNumberOk

`func (o *Account) GetCompanyNumberOk() (*int32, bool)`

GetCompanyNumberOk returns a tuple with the CompanyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNumber

`func (o *Account) SetCompanyNumber(v int32)`

SetCompanyNumber sets CompanyNumber field to given value.

### HasCompanyNumber

`func (o *Account) HasCompanyNumber() bool`

HasCompanyNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Account) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Account) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Account) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Account) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCspId

`func (o *Account) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *Account) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *Account) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *Account) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *Account) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Account) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Account) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Account) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Account) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Account) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Account) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Account) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpAuthnEnabled

`func (o *Account) GetIdpAuthnEnabled() bool`

GetIdpAuthnEnabled returns the IdpAuthnEnabled field if non-nil, zero value otherwise.

### GetIdpAuthnEnabledOk

`func (o *Account) GetIdpAuthnEnabledOk() (*bool, bool)`

GetIdpAuthnEnabledOk returns a tuple with the IdpAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAuthnEnabled

`func (o *Account) SetIdpAuthnEnabled(v bool)`

SetIdpAuthnEnabled sets IdpAuthnEnabled field to given value.

### HasIdpAuthnEnabled

`func (o *Account) HasIdpAuthnEnabled() bool`

HasIdpAuthnEnabled returns a boolean if a field has been set.

### GetLabels

`func (o *Account) GetLabels() map[string]map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Account) GetLabelsOk() (*map[string]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Account) SetLabels(v map[string]map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Account) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentAccountId

`func (o *Account) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *Account) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *Account) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *Account) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetSfdcAccountId

`func (o *Account) GetSfdcAccountId() string`

GetSfdcAccountId returns the SfdcAccountId field if non-nil, zero value otherwise.

### GetSfdcAccountIdOk

`func (o *Account) GetSfdcAccountIdOk() (*string, bool)`

GetSfdcAccountIdOk returns a tuple with the SfdcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcAccountId

`func (o *Account) SetSfdcAccountId(v string)`

SetSfdcAccountId sets SfdcAccountId field to given value.

### HasSfdcAccountId

`func (o *Account) HasSfdcAccountId() bool`

HasSfdcAccountId returns a boolean if a field has been set.

### GetState

`func (o *Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Account) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Account) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Account) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *Account) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *Account) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *Account) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *Account) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStorageId

`func (o *Account) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *Account) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *Account) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *Account) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetSubresourceName

`func (o *Account) GetSubresourceName() string`

GetSubresourceName returns the SubresourceName field if non-nil, zero value otherwise.

### GetSubresourceNameOk

`func (o *Account) GetSubresourceNameOk() (*string, bool)`

GetSubresourceNameOk returns a tuple with the SubresourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubresourceName

`func (o *Account) SetSubresourceName(v string)`

SetSubresourceName sets SubresourceName field to given value.

### HasSubresourceName

`func (o *Account) HasSubresourceName() bool`

HasSubresourceName returns a boolean if a field has been set.

### GetSupportEnabled

`func (o *Account) GetSupportEnabled() bool`

GetSupportEnabled returns the SupportEnabled field if non-nil, zero value otherwise.

### GetSupportEnabledOk

`func (o *Account) GetSupportEnabledOk() (*bool, bool)`

GetSupportEnabledOk returns a tuple with the SupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEnabled

`func (o *Account) SetSupportEnabled(v bool)`

SetSupportEnabled sets SupportEnabled field to given value.

### HasSupportEnabled

`func (o *Account) HasSupportEnabled() bool`

HasSupportEnabled returns a boolean if a field has been set.

### GetTags

`func (o *Account) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Account) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Account) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Account) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


