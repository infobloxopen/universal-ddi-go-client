# CompartmentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewCompartmentCreateRequest

`func NewCompartmentCreateRequest() *CompartmentCreateRequest`

NewCompartmentCreateRequest instantiates a new CompartmentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompartmentCreateRequestWithDefaults

`func NewCompartmentCreateRequestWithDefaults() *CompartmentCreateRequest`

NewCompartmentCreateRequestWithDefaults instantiates a new CompartmentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CompartmentCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CompartmentCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CompartmentCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CompartmentCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CompartmentCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompartmentCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompartmentCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompartmentCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *CompartmentCreateRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CompartmentCreateRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CompartmentCreateRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CompartmentCreateRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTags

`func (o *CompartmentCreateRequest) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CompartmentCreateRequest) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CompartmentCreateRequest) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *CompartmentCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


