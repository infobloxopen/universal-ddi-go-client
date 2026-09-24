# ManagedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountOrganizationGroupId** | Pointer to **string** | The resource identifier. | [optional] 
**AccountOrganizationGroupName** | Pointer to **string** |  | [optional] 
**AccountOrganizationGroupPath** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Addresses** | Pointer to [**[]AccountAddress**](AccountAddress.md) |  | [optional] 
**AdminUser** | Pointer to [**User**](User.md) |  | [optional] 
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
**ManagementModel** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentAccountId** | Pointer to **string** | The resource identifier. | [optional] 
**ParentAccountOrganizationGroupId** | Pointer to **string** | The resource identifier. | [optional] 
**PrimaryContact** | Pointer to [**AccountContact**](AccountContact.md) |  | [optional] 
**SfdcAccountId** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateChangedAt** | Pointer to **time.Time** |  | [optional] 
**StorageId** | Pointer to **int32** |  | [optional] 
**SupportEnabled** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewManagedAccount

`func NewManagedAccount() *ManagedAccount`

NewManagedAccount instantiates a new ManagedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountWithDefaults

`func NewManagedAccountWithDefaults() *ManagedAccount`

NewManagedAccountWithDefaults instantiates a new ManagedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountOrganizationGroupId

`func (o *ManagedAccount) GetAccountOrganizationGroupId() string`

GetAccountOrganizationGroupId returns the AccountOrganizationGroupId field if non-nil, zero value otherwise.

### GetAccountOrganizationGroupIdOk

`func (o *ManagedAccount) GetAccountOrganizationGroupIdOk() (*string, bool)`

GetAccountOrganizationGroupIdOk returns a tuple with the AccountOrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOrganizationGroupId

`func (o *ManagedAccount) SetAccountOrganizationGroupId(v string)`

SetAccountOrganizationGroupId sets AccountOrganizationGroupId field to given value.

### HasAccountOrganizationGroupId

`func (o *ManagedAccount) HasAccountOrganizationGroupId() bool`

HasAccountOrganizationGroupId returns a boolean if a field has been set.

### GetAccountOrganizationGroupName

`func (o *ManagedAccount) GetAccountOrganizationGroupName() string`

GetAccountOrganizationGroupName returns the AccountOrganizationGroupName field if non-nil, zero value otherwise.

### GetAccountOrganizationGroupNameOk

`func (o *ManagedAccount) GetAccountOrganizationGroupNameOk() (*string, bool)`

GetAccountOrganizationGroupNameOk returns a tuple with the AccountOrganizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOrganizationGroupName

`func (o *ManagedAccount) SetAccountOrganizationGroupName(v string)`

SetAccountOrganizationGroupName sets AccountOrganizationGroupName field to given value.

### HasAccountOrganizationGroupName

`func (o *ManagedAccount) HasAccountOrganizationGroupName() bool`

HasAccountOrganizationGroupName returns a boolean if a field has been set.

### GetAccountOrganizationGroupPath

`func (o *ManagedAccount) GetAccountOrganizationGroupPath() string`

GetAccountOrganizationGroupPath returns the AccountOrganizationGroupPath field if non-nil, zero value otherwise.

### GetAccountOrganizationGroupPathOk

`func (o *ManagedAccount) GetAccountOrganizationGroupPathOk() (*string, bool)`

GetAccountOrganizationGroupPathOk returns a tuple with the AccountOrganizationGroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOrganizationGroupPath

`func (o *ManagedAccount) SetAccountOrganizationGroupPath(v string)`

SetAccountOrganizationGroupPath sets AccountOrganizationGroupPath field to given value.

### HasAccountOrganizationGroupPath

`func (o *ManagedAccount) HasAccountOrganizationGroupPath() bool`

HasAccountOrganizationGroupPath returns a boolean if a field has been set.

### GetAccountType

`func (o *ManagedAccount) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *ManagedAccount) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *ManagedAccount) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *ManagedAccount) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAddresses

`func (o *ManagedAccount) GetAddresses() []AccountAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ManagedAccount) GetAddressesOk() (*[]AccountAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ManagedAccount) SetAddresses(v []AccountAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ManagedAccount) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAdminUser

`func (o *ManagedAccount) GetAdminUser() User`

GetAdminUser returns the AdminUser field if non-nil, zero value otherwise.

### GetAdminUserOk

`func (o *ManagedAccount) GetAdminUserOk() (*User, bool)`

GetAdminUserOk returns a tuple with the AdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUser

`func (o *ManagedAccount) SetAdminUser(v User)`

SetAdminUser sets AdminUser field to given value.

### HasAdminUser

`func (o *ManagedAccount) HasAdminUser() bool`

HasAdminUser returns a boolean if a field has been set.

### GetAdminUserId

`func (o *ManagedAccount) GetAdminUserId() string`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *ManagedAccount) GetAdminUserIdOk() (*string, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *ManagedAccount) SetAdminUserId(v string)`

SetAdminUserId sets AdminUserId field to given value.

### HasAdminUserId

`func (o *ManagedAccount) HasAdminUserId() bool`

HasAdminUserId returns a boolean if a field has been set.

### GetAxurAccountId

`func (o *ManagedAccount) GetAxurAccountId() string`

GetAxurAccountId returns the AxurAccountId field if non-nil, zero value otherwise.

### GetAxurAccountIdOk

`func (o *ManagedAccount) GetAxurAccountIdOk() (*string, bool)`

GetAxurAccountIdOk returns a tuple with the AxurAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxurAccountId

`func (o *ManagedAccount) SetAxurAccountId(v string)`

SetAxurAccountId sets AxurAccountId field to given value.

### HasAxurAccountId

`func (o *ManagedAccount) HasAxurAccountId() bool`

HasAxurAccountId returns a boolean if a field has been set.

### GetCompanyDomain

`func (o *ManagedAccount) GetCompanyDomain() string`

GetCompanyDomain returns the CompanyDomain field if non-nil, zero value otherwise.

### GetCompanyDomainOk

`func (o *ManagedAccount) GetCompanyDomainOk() (*string, bool)`

GetCompanyDomainOk returns a tuple with the CompanyDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyDomain

`func (o *ManagedAccount) SetCompanyDomain(v string)`

SetCompanyDomain sets CompanyDomain field to given value.

### HasCompanyDomain

`func (o *ManagedAccount) HasCompanyDomain() bool`

HasCompanyDomain returns a boolean if a field has been set.

### GetCompanyNumber

`func (o *ManagedAccount) GetCompanyNumber() int32`

GetCompanyNumber returns the CompanyNumber field if non-nil, zero value otherwise.

### GetCompanyNumberOk

`func (o *ManagedAccount) GetCompanyNumberOk() (*int32, bool)`

GetCompanyNumberOk returns a tuple with the CompanyNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyNumber

`func (o *ManagedAccount) SetCompanyNumber(v int32)`

SetCompanyNumber sets CompanyNumber field to given value.

### HasCompanyNumber

`func (o *ManagedAccount) HasCompanyNumber() bool`

HasCompanyNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ManagedAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedAccount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ManagedAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ManagedAccount) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ManagedAccount) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ManagedAccount) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ManagedAccount) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCspId

`func (o *ManagedAccount) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *ManagedAccount) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *ManagedAccount) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *ManagedAccount) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ManagedAccount) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ManagedAccount) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ManagedAccount) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ManagedAccount) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ManagedAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedAccount) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ManagedAccount) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ManagedAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdpAuthnEnabled

`func (o *ManagedAccount) GetIdpAuthnEnabled() bool`

GetIdpAuthnEnabled returns the IdpAuthnEnabled field if non-nil, zero value otherwise.

### GetIdpAuthnEnabledOk

`func (o *ManagedAccount) GetIdpAuthnEnabledOk() (*bool, bool)`

GetIdpAuthnEnabledOk returns a tuple with the IdpAuthnEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAuthnEnabled

`func (o *ManagedAccount) SetIdpAuthnEnabled(v bool)`

SetIdpAuthnEnabled sets IdpAuthnEnabled field to given value.

### HasIdpAuthnEnabled

`func (o *ManagedAccount) HasIdpAuthnEnabled() bool`

HasIdpAuthnEnabled returns a boolean if a field has been set.

### GetLabels

`func (o *ManagedAccount) GetLabels() map[string]map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ManagedAccount) GetLabelsOk() (*map[string]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ManagedAccount) SetLabels(v map[string]map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ManagedAccount) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetManagementModel

`func (o *ManagedAccount) GetManagementModel() string`

GetManagementModel returns the ManagementModel field if non-nil, zero value otherwise.

### GetManagementModelOk

`func (o *ManagedAccount) GetManagementModelOk() (*string, bool)`

GetManagementModelOk returns a tuple with the ManagementModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementModel

`func (o *ManagedAccount) SetManagementModel(v string)`

SetManagementModel sets ManagementModel field to given value.

### HasManagementModel

`func (o *ManagedAccount) HasManagementModel() bool`

HasManagementModel returns a boolean if a field has been set.

### GetName

`func (o *ManagedAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentAccountId

`func (o *ManagedAccount) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *ManagedAccount) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *ManagedAccount) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *ManagedAccount) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetParentAccountOrganizationGroupId

`func (o *ManagedAccount) GetParentAccountOrganizationGroupId() string`

GetParentAccountOrganizationGroupId returns the ParentAccountOrganizationGroupId field if non-nil, zero value otherwise.

### GetParentAccountOrganizationGroupIdOk

`func (o *ManagedAccount) GetParentAccountOrganizationGroupIdOk() (*string, bool)`

GetParentAccountOrganizationGroupIdOk returns a tuple with the ParentAccountOrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountOrganizationGroupId

`func (o *ManagedAccount) SetParentAccountOrganizationGroupId(v string)`

SetParentAccountOrganizationGroupId sets ParentAccountOrganizationGroupId field to given value.

### HasParentAccountOrganizationGroupId

`func (o *ManagedAccount) HasParentAccountOrganizationGroupId() bool`

HasParentAccountOrganizationGroupId returns a boolean if a field has been set.

### GetPrimaryContact

`func (o *ManagedAccount) GetPrimaryContact() AccountContact`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *ManagedAccount) GetPrimaryContactOk() (*AccountContact, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *ManagedAccount) SetPrimaryContact(v AccountContact)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *ManagedAccount) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.

### GetSfdcAccountId

`func (o *ManagedAccount) GetSfdcAccountId() string`

GetSfdcAccountId returns the SfdcAccountId field if non-nil, zero value otherwise.

### GetSfdcAccountIdOk

`func (o *ManagedAccount) GetSfdcAccountIdOk() (*string, bool)`

GetSfdcAccountIdOk returns a tuple with the SfdcAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSfdcAccountId

`func (o *ManagedAccount) SetSfdcAccountId(v string)`

SetSfdcAccountId sets SfdcAccountId field to given value.

### HasSfdcAccountId

`func (o *ManagedAccount) HasSfdcAccountId() bool`

HasSfdcAccountId returns a boolean if a field has been set.

### GetState

`func (o *ManagedAccount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedAccount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedAccount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagedAccount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *ManagedAccount) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *ManagedAccount) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *ManagedAccount) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *ManagedAccount) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStorageId

`func (o *ManagedAccount) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *ManagedAccount) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *ManagedAccount) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *ManagedAccount) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetSupportEnabled

`func (o *ManagedAccount) GetSupportEnabled() bool`

GetSupportEnabled returns the SupportEnabled field if non-nil, zero value otherwise.

### GetSupportEnabledOk

`func (o *ManagedAccount) GetSupportEnabledOk() (*bool, bool)`

GetSupportEnabledOk returns a tuple with the SupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEnabled

`func (o *ManagedAccount) SetSupportEnabled(v bool)`

SetSupportEnabled sets SupportEnabled field to given value.

### HasSupportEnabled

`func (o *ManagedAccount) HasSupportEnabled() bool`

HasSupportEnabled returns a boolean if a field has been set.

### GetTags

`func (o *ManagedAccount) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ManagedAccount) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ManagedAccount) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ManagedAccount) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ManagedAccount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedAccount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedAccount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ManagedAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserId

`func (o *ManagedAccount) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ManagedAccount) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ManagedAccount) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ManagedAccount) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


