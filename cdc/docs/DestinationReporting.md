# DestinationReporting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ReportingApplianceAddress** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewDestinationReporting

`func NewDestinationReporting() *DestinationReporting`

NewDestinationReporting instantiates a new DestinationReporting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationReportingWithDefaults

`func NewDestinationReportingWithDefaults() *DestinationReporting`

NewDestinationReportingWithDefaults instantiates a new DestinationReporting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DestinationReporting) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DestinationReporting) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DestinationReporting) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DestinationReporting) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DestinationReporting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DestinationReporting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DestinationReporting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DestinationReporting) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DestinationReporting) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DestinationReporting) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DestinationReporting) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DestinationReporting) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DestinationReporting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DestinationReporting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DestinationReporting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DestinationReporting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DestinationReporting) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DestinationReporting) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DestinationReporting) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DestinationReporting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DestinationReporting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationReporting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationReporting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationReporting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *DestinationReporting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DestinationReporting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DestinationReporting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DestinationReporting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetReportingApplianceAddress

`func (o *DestinationReporting) GetReportingApplianceAddress() string`

GetReportingApplianceAddress returns the ReportingApplianceAddress field if non-nil, zero value otherwise.

### GetReportingApplianceAddressOk

`func (o *DestinationReporting) GetReportingApplianceAddressOk() (*string, bool)`

GetReportingApplianceAddressOk returns a tuple with the ReportingApplianceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingApplianceAddress

`func (o *DestinationReporting) SetReportingApplianceAddress(v string)`

SetReportingApplianceAddress sets ReportingApplianceAddress field to given value.

### HasReportingApplianceAddress

`func (o *DestinationReporting) HasReportingApplianceAddress() bool`

HasReportingApplianceAddress returns a boolean if a field has been set.

### GetTags

`func (o *DestinationReporting) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DestinationReporting) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DestinationReporting) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DestinationReporting) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DestinationReporting) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DestinationReporting) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DestinationReporting) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DestinationReporting) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsername

`func (o *DestinationReporting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DestinationReporting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DestinationReporting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DestinationReporting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


