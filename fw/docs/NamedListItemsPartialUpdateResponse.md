# NamedListItemsPartialUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedItems** | Pointer to [**[]ItemStructs**](ItemStructs.md) | Deleted items list | [optional] 
**InsertedItems** | Pointer to [**[]ItemStructs**](ItemStructs.md) | Inserted items list | [optional] 
**UpdatedItems** | Pointer to [**[]ItemStructs**](ItemStructs.md) | Updated items list | [optional] 

## Methods

### NewNamedListItemsPartialUpdateResponse

`func NewNamedListItemsPartialUpdateResponse() *NamedListItemsPartialUpdateResponse`

NewNamedListItemsPartialUpdateResponse instantiates a new NamedListItemsPartialUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedListItemsPartialUpdateResponseWithDefaults

`func NewNamedListItemsPartialUpdateResponseWithDefaults() *NamedListItemsPartialUpdateResponse`

NewNamedListItemsPartialUpdateResponseWithDefaults instantiates a new NamedListItemsPartialUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedItems

`func (o *NamedListItemsPartialUpdateResponse) GetDeletedItems() []ItemStructs`

GetDeletedItems returns the DeletedItems field if non-nil, zero value otherwise.

### GetDeletedItemsOk

`func (o *NamedListItemsPartialUpdateResponse) GetDeletedItemsOk() (*[]ItemStructs, bool)`

GetDeletedItemsOk returns a tuple with the DeletedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedItems

`func (o *NamedListItemsPartialUpdateResponse) SetDeletedItems(v []ItemStructs)`

SetDeletedItems sets DeletedItems field to given value.

### HasDeletedItems

`func (o *NamedListItemsPartialUpdateResponse) HasDeletedItems() bool`

HasDeletedItems returns a boolean if a field has been set.

### GetInsertedItems

`func (o *NamedListItemsPartialUpdateResponse) GetInsertedItems() []ItemStructs`

GetInsertedItems returns the InsertedItems field if non-nil, zero value otherwise.

### GetInsertedItemsOk

`func (o *NamedListItemsPartialUpdateResponse) GetInsertedItemsOk() (*[]ItemStructs, bool)`

GetInsertedItemsOk returns a tuple with the InsertedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedItems

`func (o *NamedListItemsPartialUpdateResponse) SetInsertedItems(v []ItemStructs)`

SetInsertedItems sets InsertedItems field to given value.

### HasInsertedItems

`func (o *NamedListItemsPartialUpdateResponse) HasInsertedItems() bool`

HasInsertedItems returns a boolean if a field has been set.

### GetUpdatedItems

`func (o *NamedListItemsPartialUpdateResponse) GetUpdatedItems() []ItemStructs`

GetUpdatedItems returns the UpdatedItems field if non-nil, zero value otherwise.

### GetUpdatedItemsOk

`func (o *NamedListItemsPartialUpdateResponse) GetUpdatedItemsOk() (*[]ItemStructs, bool)`

GetUpdatedItemsOk returns a tuple with the UpdatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedItems

`func (o *NamedListItemsPartialUpdateResponse) SetUpdatedItems(v []ItemStructs)`

SetUpdatedItems sets UpdatedItems field to given value.

### HasUpdatedItems

`func (o *NamedListItemsPartialUpdateResponse) HasUpdatedItems() bool`

HasUpdatedItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


