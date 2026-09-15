# ItemStructs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the item. | [optional] 
**ExpiryTime** | Pointer to **time.Time** | The time at which this list item expires, as an RFC 3339 timestamp string. May be specified in any timezone. Unset (null) means no expiry.  Write semantics: - Insert/replace (POST /named_lists/{id}/items): set when present;   NULL when absent in payload. - Patch update (PATCH /named_lists/{id}/items, updated_items_described):   set when present; unchanged when absent in payload. Clearing an existing expiry_time via request field mask is not supported. | [optional] 
**Item** | Pointer to **string** | The data of the item. | [optional] 
**Status** | Pointer to [**ItemStructsItemStatus**](ItemStructsItemStatus.md) | The status of the item. Applicable to TI domains only | [optional] [default to ITEMSTRUCTSITEMSTATUS_ACTIVE]
**StatusDetails** | Pointer to **string** | The status details of the item. Applicable to TI domains only | [optional] 

## Methods

### NewItemStructs

`func NewItemStructs() *ItemStructs`

NewItemStructs instantiates a new ItemStructs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStructsWithDefaults

`func NewItemStructsWithDefaults() *ItemStructs`

NewItemStructsWithDefaults instantiates a new ItemStructs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ItemStructs) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemStructs) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemStructs) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemStructs) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ItemStructs) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ItemStructs) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ItemStructs) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ItemStructs) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetItem

`func (o *ItemStructs) GetItem() string`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemStructs) GetItemOk() (*string, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemStructs) SetItem(v string)`

SetItem sets Item field to given value.

### HasItem

`func (o *ItemStructs) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetStatus

`func (o *ItemStructs) GetStatus() ItemStructsItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemStructs) GetStatusOk() (*ItemStructsItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemStructs) SetStatus(v ItemStructsItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemStructs) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *ItemStructs) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *ItemStructs) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *ItemStructs) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *ItemStructs) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


