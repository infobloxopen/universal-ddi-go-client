# CurrentUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code. | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**MarketingOptIn** | Pointer to **bool** | Marketing communications opt-in preference. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**StateRegion** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrentUserUpdateRequest

`func NewCurrentUserUpdateRequest() *CurrentUserUpdateRequest`

NewCurrentUserUpdateRequest instantiates a new CurrentUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserUpdateRequestWithDefaults

`func NewCurrentUserUpdateRequestWithDefaults() *CurrentUserUpdateRequest`

NewCurrentUserUpdateRequestWithDefaults instantiates a new CurrentUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CurrentUserUpdateRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CurrentUserUpdateRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CurrentUserUpdateRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CurrentUserUpdateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFirstName

`func (o *CurrentUserUpdateRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CurrentUserUpdateRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CurrentUserUpdateRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CurrentUserUpdateRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetJobTitle

`func (o *CurrentUserUpdateRequest) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *CurrentUserUpdateRequest) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *CurrentUserUpdateRequest) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *CurrentUserUpdateRequest) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastName

`func (o *CurrentUserUpdateRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CurrentUserUpdateRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CurrentUserUpdateRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CurrentUserUpdateRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMarketingOptIn

`func (o *CurrentUserUpdateRequest) GetMarketingOptIn() bool`

GetMarketingOptIn returns the MarketingOptIn field if non-nil, zero value otherwise.

### GetMarketingOptInOk

`func (o *CurrentUserUpdateRequest) GetMarketingOptInOk() (*bool, bool)`

GetMarketingOptInOk returns a tuple with the MarketingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIn

`func (o *CurrentUserUpdateRequest) SetMarketingOptIn(v bool)`

SetMarketingOptIn sets MarketingOptIn field to given value.

### HasMarketingOptIn

`func (o *CurrentUserUpdateRequest) HasMarketingOptIn() bool`

HasMarketingOptIn returns a boolean if a field has been set.

### GetName

`func (o *CurrentUserUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUserUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUserUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CurrentUserUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CurrentUserUpdateRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CurrentUserUpdateRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CurrentUserUpdateRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CurrentUserUpdateRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStateRegion

`func (o *CurrentUserUpdateRequest) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *CurrentUserUpdateRequest) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *CurrentUserUpdateRequest) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *CurrentUserUpdateRequest) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### GetTimezone

`func (o *CurrentUserUpdateRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CurrentUserUpdateRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CurrentUserUpdateRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CurrentUserUpdateRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


