# GroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ApiPageInfo**](ApiPageInfo.md) |  | [optional] 
**Results** | Pointer to [**[]Group**](Group.md) |  | [optional] 

## Methods

### NewGroupListResponse

`func NewGroupListResponse() *GroupListResponse`

NewGroupListResponse instantiates a new GroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupListResponseWithDefaults

`func NewGroupListResponseWithDefaults() *GroupListResponse`

NewGroupListResponseWithDefaults instantiates a new GroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *GroupListResponse) GetPage() ApiPageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GroupListResponse) GetPageOk() (*ApiPageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GroupListResponse) SetPage(v ApiPageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *GroupListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetResults

`func (o *GroupListResponse) GetResults() []Group`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GroupListResponse) GetResultsOk() (*[]Group, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GroupListResponse) SetResults(v []Group)`

SetResults sets Results field to given value.

### HasResults

`func (o *GroupListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


