# SharedNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the shared network. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the shared network. Must contain 1 to 256 characters. Can include UTF-8. | 
**Space** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewSharedNetwork

`func NewSharedNetwork(name string, ) *SharedNetwork`

NewSharedNetwork instantiates a new SharedNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedNetworkWithDefaults

`func NewSharedNetworkWithDefaults() *SharedNetwork`

NewSharedNetworkWithDefaults instantiates a new SharedNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SharedNetwork) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SharedNetwork) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SharedNetwork) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SharedNetwork) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *SharedNetwork) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SharedNetwork) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SharedNetwork) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SharedNetwork) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SharedNetwork) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedNetwork) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedNetwork) SetName(v string)`

SetName sets Name field to given value.


### GetSpace

`func (o *SharedNetwork) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *SharedNetwork) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *SharedNetwork) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *SharedNetwork) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


