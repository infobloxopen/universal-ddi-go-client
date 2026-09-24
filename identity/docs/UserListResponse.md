# UserListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]User**](User.md) |  | [optional] 

## Methods

### NewUserListResponse

`func NewUserListResponse() *UserListResponse`

NewUserListResponse instantiates a new UserListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListResponseWithDefaults

`func NewUserListResponseWithDefaults() *UserListResponse`

NewUserListResponseWithDefaults instantiates a new UserListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *UserListResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *UserListResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *UserListResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *UserListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *UserListResponse) GetResults() []User`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *UserListResponse) GetResultsOk() (*[]User, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *UserListResponse) SetResults(v []User)`

SetResults sets Results field to given value.

### HasResults

`func (o *UserListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


