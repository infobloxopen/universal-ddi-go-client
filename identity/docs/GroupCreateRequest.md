# GroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AllowedCidrs** | Pointer to **[]string** | The list of allowed CIDRs for the group. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnableAllowedCidrs** | Pointer to **bool** | The flag to enable allowed CIDRs for the group. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | The name of the group. | [optional] 
**UserIds** | Pointer to **[]string** | The resource identifier. | [optional] 

## Methods

### NewGroupCreateRequest

`func NewGroupCreateRequest() *GroupCreateRequest`

NewGroupCreateRequest instantiates a new GroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCreateRequestWithDefaults

`func NewGroupCreateRequestWithDefaults() *GroupCreateRequest`

NewGroupCreateRequestWithDefaults instantiates a new GroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GroupCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GroupCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GroupCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GroupCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAllowedCidrs

`func (o *GroupCreateRequest) GetAllowedCidrs() []string`

GetAllowedCidrs returns the AllowedCidrs field if non-nil, zero value otherwise.

### GetAllowedCidrsOk

`func (o *GroupCreateRequest) GetAllowedCidrsOk() (*[]string, bool)`

GetAllowedCidrsOk returns a tuple with the AllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCidrs

`func (o *GroupCreateRequest) SetAllowedCidrs(v []string)`

SetAllowedCidrs sets AllowedCidrs field to given value.

### HasAllowedCidrs

`func (o *GroupCreateRequest) HasAllowedCidrs() bool`

HasAllowedCidrs returns a boolean if a field has been set.

### GetDescription

`func (o *GroupCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAllowedCidrs

`func (o *GroupCreateRequest) GetEnableAllowedCidrs() bool`

GetEnableAllowedCidrs returns the EnableAllowedCidrs field if non-nil, zero value otherwise.

### GetEnableAllowedCidrsOk

`func (o *GroupCreateRequest) GetEnableAllowedCidrsOk() (*bool, bool)`

GetEnableAllowedCidrsOk returns a tuple with the EnableAllowedCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowedCidrs

`func (o *GroupCreateRequest) SetEnableAllowedCidrs(v bool)`

SetEnableAllowedCidrs sets EnableAllowedCidrs field to given value.

### HasEnableAllowedCidrs

`func (o *GroupCreateRequest) HasEnableAllowedCidrs() bool`

HasEnableAllowedCidrs returns a boolean if a field has been set.

### GetId

`func (o *GroupCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *GroupCreateRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GroupCreateRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GroupCreateRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GroupCreateRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *GroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserIds

`func (o *GroupCreateRequest) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *GroupCreateRequest) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *GroupCreateRequest) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *GroupCreateRequest) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


