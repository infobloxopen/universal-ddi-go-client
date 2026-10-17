# DataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourceFields** | Pointer to [**[]SourceField**](SourceField.md) |  | [optional] 
**SourceSubtypes** | Pointer to **[]string** |  | [optional] 
**SourceType** | Pointer to [**SourceType**](SourceType.md) |  | [optional] [default to SOURCETYPE_UNKNOWN]
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDataSource

`func NewDataSource() *DataSource`

NewDataSource instantiates a new DataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataSourceWithDefaults

`func NewDataSourceWithDefaults() *DataSource`

NewDataSourceWithDefaults instantiates a new DataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DataSource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DataSource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DataSource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DataSource) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DataSource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DataSource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DataSource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DataSource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DataSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DataSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DataSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DataSource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceFields

`func (o *DataSource) GetSourceFields() []SourceField`

GetSourceFields returns the SourceFields field if non-nil, zero value otherwise.

### GetSourceFieldsOk

`func (o *DataSource) GetSourceFieldsOk() (*[]SourceField, bool)`

GetSourceFieldsOk returns a tuple with the SourceFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFields

`func (o *DataSource) SetSourceFields(v []SourceField)`

SetSourceFields sets SourceFields field to given value.

### HasSourceFields

`func (o *DataSource) HasSourceFields() bool`

HasSourceFields returns a boolean if a field has been set.

### GetSourceSubtypes

`func (o *DataSource) GetSourceSubtypes() []string`

GetSourceSubtypes returns the SourceSubtypes field if non-nil, zero value otherwise.

### GetSourceSubtypesOk

`func (o *DataSource) GetSourceSubtypesOk() (*[]string, bool)`

GetSourceSubtypesOk returns a tuple with the SourceSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSubtypes

`func (o *DataSource) SetSourceSubtypes(v []string)`

SetSourceSubtypes sets SourceSubtypes field to given value.

### HasSourceSubtypes

`func (o *DataSource) HasSourceSubtypes() bool`

HasSourceSubtypes returns a boolean if a field has been set.

### GetSourceType

`func (o *DataSource) GetSourceType() SourceType`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *DataSource) GetSourceTypeOk() (*SourceType, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *DataSource) SetSourceType(v SourceType)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *DataSource) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *DataSource) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DataSource) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DataSource) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DataSource) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DataSource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DataSource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DataSource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DataSource) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


