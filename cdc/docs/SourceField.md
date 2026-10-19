# SourceField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 

## Methods

### NewSourceField

`func NewSourceField() *SourceField`

NewSourceField instantiates a new SourceField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceFieldWithDefaults

`func NewSourceFieldWithDefaults() *SourceField`

NewSourceFieldWithDefaults instantiates a new SourceField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SourceField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SourceField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SourceField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SourceField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SourceField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *SourceField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SourceField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SourceField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SourceField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


