# AddrBlockApprovals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int32** | The account id of the addr block. | [optional] 
**Address** | Pointer to **string** | The address of the addr block. | [optional] 
**ApprovalStatus** | Pointer to [**AddrBlockApprovalsApprovalStatus**](AddrBlockApprovalsApprovalStatus.md) | The approval status of the addr block. | [optional] [default to ADDRBLOCKAPPROVALSAPPROVALSTATUS_VERIFIED]
**Comments** | Pointer to **string** | Any comments about the address block. | [optional] 
**CompanyName** | Pointer to **string** | The company name of the addr block. | [optional] 
**RequestedAt** | Pointer to **time.Time** | The time the approval request was created, which is based on the address block&#39;s creation time. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time this Address Block object was updated. | [optional] [readonly] 
**Username** | Pointer to **string** | The user who created the address block. | [optional] 

## Methods

### NewAddrBlockApprovals

`func NewAddrBlockApprovals() *AddrBlockApprovals`

NewAddrBlockApprovals instantiates a new AddrBlockApprovals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddrBlockApprovalsWithDefaults

`func NewAddrBlockApprovalsWithDefaults() *AddrBlockApprovals`

NewAddrBlockApprovalsWithDefaults instantiates a new AddrBlockApprovals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AddrBlockApprovals) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddrBlockApprovals) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddrBlockApprovals) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddrBlockApprovals) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAddress

`func (o *AddrBlockApprovals) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddrBlockApprovals) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddrBlockApprovals) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddrBlockApprovals) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *AddrBlockApprovals) GetApprovalStatus() AddrBlockApprovalsApprovalStatus`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *AddrBlockApprovals) GetApprovalStatusOk() (*AddrBlockApprovalsApprovalStatus, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *AddrBlockApprovals) SetApprovalStatus(v AddrBlockApprovalsApprovalStatus)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *AddrBlockApprovals) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetComments

`func (o *AddrBlockApprovals) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *AddrBlockApprovals) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *AddrBlockApprovals) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *AddrBlockApprovals) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCompanyName

`func (o *AddrBlockApprovals) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *AddrBlockApprovals) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *AddrBlockApprovals) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *AddrBlockApprovals) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetRequestedAt

`func (o *AddrBlockApprovals) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *AddrBlockApprovals) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *AddrBlockApprovals) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *AddrBlockApprovals) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AddrBlockApprovals) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AddrBlockApprovals) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AddrBlockApprovals) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AddrBlockApprovals) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUsername

`func (o *AddrBlockApprovals) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddrBlockApprovals) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddrBlockApprovals) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddrBlockApprovals) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


