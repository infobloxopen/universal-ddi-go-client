# AccountOrganizationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The resource identifier. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** | The resource identifier. | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**ParentAccountOrganizationGroupId** | Pointer to **string** | The resource identifier. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAccountOrganizationGroup

`func NewAccountOrganizationGroup() *AccountOrganizationGroup`

NewAccountOrganizationGroup instantiates a new AccountOrganizationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOrganizationGroupWithDefaults

`func NewAccountOrganizationGroupWithDefaults() *AccountOrganizationGroup`

NewAccountOrganizationGroupWithDefaults instantiates a new AccountOrganizationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountOrganizationGroup) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountOrganizationGroup) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountOrganizationGroup) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountOrganizationGroup) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountOrganizationGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountOrganizationGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountOrganizationGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountOrganizationGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AccountOrganizationGroup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccountOrganizationGroup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccountOrganizationGroup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AccountOrganizationGroup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeletedAt

`func (o *AccountOrganizationGroup) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *AccountOrganizationGroup) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *AccountOrganizationGroup) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *AccountOrganizationGroup) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetId

`func (o *AccountOrganizationGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountOrganizationGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountOrganizationGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AccountOrganizationGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountOrganizationGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountOrganizationGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountOrganizationGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountOrganizationGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentAccountOrganizationGroupId

`func (o *AccountOrganizationGroup) GetParentAccountOrganizationGroupId() string`

GetParentAccountOrganizationGroupId returns the ParentAccountOrganizationGroupId field if non-nil, zero value otherwise.

### GetParentAccountOrganizationGroupIdOk

`func (o *AccountOrganizationGroup) GetParentAccountOrganizationGroupIdOk() (*string, bool)`

GetParentAccountOrganizationGroupIdOk returns a tuple with the ParentAccountOrganizationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountOrganizationGroupId

`func (o *AccountOrganizationGroup) SetParentAccountOrganizationGroupId(v string)`

SetParentAccountOrganizationGroupId sets ParentAccountOrganizationGroupId field to given value.

### HasParentAccountOrganizationGroupId

`func (o *AccountOrganizationGroup) HasParentAccountOrganizationGroupId() bool`

HasParentAccountOrganizationGroupId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AccountOrganizationGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountOrganizationGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountOrganizationGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountOrganizationGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


