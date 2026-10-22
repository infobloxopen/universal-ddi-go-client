# ManagedAccountHierarchyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CspId** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**HasChildren** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**NodeType** | Pointer to [**ManagedAccountHierarchyNodeNodeType**](ManagedAccountHierarchyNodeNodeType.md) |  | [optional] [default to MANAGEDACCOUNTHIERARCHYNODENODETYPE_UNKNOWN]
**ParentNodeId** | Pointer to **string** | The resource identifier. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StorageId** | Pointer to **int32** |  | [optional] 

## Methods

### NewManagedAccountHierarchyNode

`func NewManagedAccountHierarchyNode() *ManagedAccountHierarchyNode`

NewManagedAccountHierarchyNode instantiates a new ManagedAccountHierarchyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedAccountHierarchyNodeWithDefaults

`func NewManagedAccountHierarchyNodeWithDefaults() *ManagedAccountHierarchyNode`

NewManagedAccountHierarchyNodeWithDefaults instantiates a new ManagedAccountHierarchyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCspId

`func (o *ManagedAccountHierarchyNode) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *ManagedAccountHierarchyNode) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *ManagedAccountHierarchyNode) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *ManagedAccountHierarchyNode) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ManagedAccountHierarchyNode) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ManagedAccountHierarchyNode) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ManagedAccountHierarchyNode) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ManagedAccountHierarchyNode) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetHasChildren

`func (o *ManagedAccountHierarchyNode) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *ManagedAccountHierarchyNode) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *ManagedAccountHierarchyNode) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *ManagedAccountHierarchyNode) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetId

`func (o *ManagedAccountHierarchyNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedAccountHierarchyNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedAccountHierarchyNode) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedAccountHierarchyNode) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagedAccountHierarchyNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedAccountHierarchyNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedAccountHierarchyNode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ManagedAccountHierarchyNode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeType

`func (o *ManagedAccountHierarchyNode) GetNodeType() ManagedAccountHierarchyNodeNodeType`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ManagedAccountHierarchyNode) GetNodeTypeOk() (*ManagedAccountHierarchyNodeNodeType, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ManagedAccountHierarchyNode) SetNodeType(v ManagedAccountHierarchyNodeNodeType)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ManagedAccountHierarchyNode) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetParentNodeId

`func (o *ManagedAccountHierarchyNode) GetParentNodeId() string`

GetParentNodeId returns the ParentNodeId field if non-nil, zero value otherwise.

### GetParentNodeIdOk

`func (o *ManagedAccountHierarchyNode) GetParentNodeIdOk() (*string, bool)`

GetParentNodeIdOk returns a tuple with the ParentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNodeId

`func (o *ManagedAccountHierarchyNode) SetParentNodeId(v string)`

SetParentNodeId sets ParentNodeId field to given value.

### HasParentNodeId

`func (o *ManagedAccountHierarchyNode) HasParentNodeId() bool`

HasParentNodeId returns a boolean if a field has been set.

### GetState

`func (o *ManagedAccountHierarchyNode) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ManagedAccountHierarchyNode) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ManagedAccountHierarchyNode) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ManagedAccountHierarchyNode) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageId

`func (o *ManagedAccountHierarchyNode) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *ManagedAccountHierarchyNode) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *ManagedAccountHierarchyNode) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *ManagedAccountHierarchyNode) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


