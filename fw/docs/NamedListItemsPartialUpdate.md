# NamedListItemsPartialUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The list of ItemStructs structures that contains items and their descriptions. | [optional] 
**Id** | Pointer to **int32** | The Named List object identifier. | [optional] [readonly] 
**InsertedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The list of ItemStructs structures that contains items and their descriptions. | [optional] 
**UpdatedItemsDescribed** | Pointer to [**[]ItemStructs**](ItemStructs.md) | The list of ItemStructs structures that contains items, descriptions, statuses, status detail | [optional] 

## Methods

### NewNamedListItemsPartialUpdate

`func NewNamedListItemsPartialUpdate() *NamedListItemsPartialUpdate`

NewNamedListItemsPartialUpdate instantiates a new NamedListItemsPartialUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedListItemsPartialUpdateWithDefaults

`func NewNamedListItemsPartialUpdateWithDefaults() *NamedListItemsPartialUpdate`

NewNamedListItemsPartialUpdateWithDefaults instantiates a new NamedListItemsPartialUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedItemsDescribed

`func (o *NamedListItemsPartialUpdate) GetDeletedItemsDescribed() []ItemStructs`

GetDeletedItemsDescribed returns the DeletedItemsDescribed field if non-nil, zero value otherwise.

### GetDeletedItemsDescribedOk

`func (o *NamedListItemsPartialUpdate) GetDeletedItemsDescribedOk() (*[]ItemStructs, bool)`

GetDeletedItemsDescribedOk returns a tuple with the DeletedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItemsDescribed

`func (o *NamedListItemsPartialUpdate) SetDeletedItemsDescribed(v []ItemStructs)`

SetDeletedItemsDescribed sets DeletedItemsDescribed field to given value.

### HasDeletedItemsDescribed

`func (o *NamedListItemsPartialUpdate) HasDeletedItemsDescribed() bool`

HasDeletedItemsDescribed returns a boolean if a field has been set.

### GetId

`func (o *NamedListItemsPartialUpdate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NamedListItemsPartialUpdate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NamedListItemsPartialUpdate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NamedListItemsPartialUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInsertedItemsDescribed

`func (o *NamedListItemsPartialUpdate) GetInsertedItemsDescribed() []ItemStructs`

GetInsertedItemsDescribed returns the InsertedItemsDescribed field if non-nil, zero value otherwise.

### GetInsertedItemsDescribedOk

`func (o *NamedListItemsPartialUpdate) GetInsertedItemsDescribedOk() (*[]ItemStructs, bool)`

GetInsertedItemsDescribedOk returns a tuple with the InsertedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItemsDescribed

`func (o *NamedListItemsPartialUpdate) SetInsertedItemsDescribed(v []ItemStructs)`

SetInsertedItemsDescribed sets InsertedItemsDescribed field to given value.

### HasInsertedItemsDescribed

`func (o *NamedListItemsPartialUpdate) HasInsertedItemsDescribed() bool`

HasInsertedItemsDescribed returns a boolean if a field has been set.

### GetUpdatedItemsDescribed

`func (o *NamedListItemsPartialUpdate) GetUpdatedItemsDescribed() []ItemStructs`

GetUpdatedItemsDescribed returns the UpdatedItemsDescribed field if non-nil, zero value otherwise.

### GetUpdatedItemsDescribedOk

`func (o *NamedListItemsPartialUpdate) GetUpdatedItemsDescribedOk() (*[]ItemStructs, bool)`

GetUpdatedItemsDescribedOk returns a tuple with the UpdatedItemsDescribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedItemsDescribed

`func (o *NamedListItemsPartialUpdate) SetUpdatedItemsDescribed(v []ItemStructs)`

SetUpdatedItemsDescribed sets UpdatedItemsDescribed field to given value.

### HasUpdatedItemsDescribed

`func (o *NamedListItemsPartialUpdate) HasUpdatedItemsDescribed() bool`

HasUpdatedItemsDescribed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


