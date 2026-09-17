# UserCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The resource identifier. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code (required for interactive users). | [optional] 
**Email** | Pointer to **string** | The email address of the user. | [optional] 
**FirstName** | Pointer to **string** | The first name of the user. | [optional] 
**GroupIds** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JobTitle** | Pointer to **string** | Job title (required for interactive users). | [optional] 
**LastName** | Pointer to **string** | The last name of the user. | [optional] 
**MarketingOptIn** | Pointer to **bool** | Marketing communications opt-in preference (required for interactive users). | [optional] 
**Name** | Pointer to **string** | The name of the user. | [optional] 
**StateRegion** | Pointer to **string** | ISO 3166-2 subdivision code, or free-text for countries without subdivisions (required for interactive users). | [optional] 
**Type** | Pointer to **string** | The type of the user (HostAppUser/ATEP/local/service/interactive). | [optional] 

## Methods

### NewUserCreateRequest

`func NewUserCreateRequest() *UserCreateRequest`

NewUserCreateRequest instantiates a new UserCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateRequestWithDefaults

`func NewUserCreateRequestWithDefaults() *UserCreateRequest`

NewUserCreateRequestWithDefaults instantiates a new UserCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCountry

`func (o *UserCreateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserCreateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserCreateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserCreateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreateRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserCreateRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreateRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreateRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCreateRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroupIds

`func (o *UserCreateRequest) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UserCreateRequest) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UserCreateRequest) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UserCreateRequest) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetId

`func (o *UserCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *UserCreateRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserCreateRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserCreateRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserCreateRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *UserCreateRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreateRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreateRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCreateRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMarketingOptIn

`func (o *UserCreateRequest) GetMarketingOptIn() bool`

GetMarketingOptIn returns the MarketingOptIn field if non-nil, zero value otherwise.

### GetMarketingOptInOk

`func (o *UserCreateRequest) GetMarketingOptInOk() (*bool, bool)`

GetMarketingOptInOk returns a tuple with the MarketingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIn

`func (o *UserCreateRequest) SetMarketingOptIn(v bool)`

SetMarketingOptIn sets MarketingOptIn field to given value.

### HasMarketingOptIn

`func (o *UserCreateRequest) HasMarketingOptIn() bool`

HasMarketingOptIn returns a boolean if a field has been set.

### GetName

`func (o *UserCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStateRegion

`func (o *UserCreateRequest) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *UserCreateRequest) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *UserCreateRequest) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *UserCreateRequest) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### GetType

`func (o *UserCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserCreateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


