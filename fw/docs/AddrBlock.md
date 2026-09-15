# AddrBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The CIDR of the address block. | [optional] 
**Description** | Pointer to **string** | End-user description for the address block. | [optional] 

## Methods

### NewAddrBlock

`func NewAddrBlock() *AddrBlock`

NewAddrBlock instantiates a new AddrBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddrBlockWithDefaults

`func NewAddrBlockWithDefaults() *AddrBlock`

NewAddrBlockWithDefaults instantiates a new AddrBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddrBlock) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddrBlock) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddrBlock) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddrBlock) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *AddrBlock) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddrBlock) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddrBlock) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddrBlock) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


