# KeyStoreCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] [readonly] 
**Notes** | Pointer to **string** |  | [optional] [readonly] 
**Values** | Pointer to [**[]KeyStoreValue**](KeyStoreValue.md) |  | [optional] [readonly] 

## Methods

### NewKeyStoreCredential

`func NewKeyStoreCredential() *KeyStoreCredential`

NewKeyStoreCredential instantiates a new KeyStoreCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyStoreCredentialWithDefaults

`func NewKeyStoreCredentialWithDefaults() *KeyStoreCredential`

NewKeyStoreCredentialWithDefaults instantiates a new KeyStoreCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyStoreCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyStoreCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyStoreCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyStoreCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyStoreCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyStoreCredential) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyStoreCredential) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyStoreCredential) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *KeyStoreCredential) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *KeyStoreCredential) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *KeyStoreCredential) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *KeyStoreCredential) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetValues

`func (o *KeyStoreCredential) GetValues() []KeyStoreValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KeyStoreCredential) GetValuesOk() (*[]KeyStoreValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KeyStoreCredential) SetValues(v []KeyStoreValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *KeyStoreCredential) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


