# Validation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 

## Methods

### NewValidation

`func NewValidation() *Validation`

NewValidation instantiates a new Validation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationWithDefaults

`func NewValidationWithDefaults() *Validation`

NewValidationWithDefaults instantiates a new Validation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Validation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Validation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Validation) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Validation) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPattern

`func (o *Validation) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Validation) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Validation) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *Validation) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


