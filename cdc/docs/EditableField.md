# EditableField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**[]SelectOptions**](SelectOptions.md) |  | [optional] 
**Range** | Pointer to [**NumberRange**](NumberRange.md) |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**FieldType**](FieldType.md) |  | [optional] [default to FIELDTYPE_SELECT]
**Validation** | Pointer to [**Validation**](Validation.md) |  | [optional] 

## Methods

### NewEditableField

`func NewEditableField() *EditableField`

NewEditableField instantiates a new EditableField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditableFieldWithDefaults

`func NewEditableFieldWithDefaults() *EditableField`

NewEditableFieldWithDefaults instantiates a new EditableField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EditableField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditableField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditableField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditableField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EditableField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditableField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditableField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditableField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *EditableField) GetOptions() []SelectOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *EditableField) GetOptionsOk() (*[]SelectOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *EditableField) SetOptions(v []SelectOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *EditableField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetRange

`func (o *EditableField) GetRange() NumberRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *EditableField) GetRangeOk() (*NumberRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *EditableField) SetRange(v NumberRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *EditableField) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRequired

`func (o *EditableField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EditableField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EditableField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EditableField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetType

`func (o *EditableField) GetType() FieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EditableField) GetTypeOk() (*FieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EditableField) SetType(v FieldType)`

SetType sets Type field to given value.

### HasType

`func (o *EditableField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidation

`func (o *EditableField) GetValidation() Validation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *EditableField) GetValidationOk() (*Validation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *EditableField) SetValidation(v Validation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *EditableField) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


