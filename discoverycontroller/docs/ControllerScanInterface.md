# ControllerScanInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Represents the interface to bind to. | [optional] 
**IpAddress** | Pointer to **string** | Represents the ip address configured on the interface. | [optional] 
**Name** | Pointer to **string** | Readable name given for this configuration. | [optional] 

## Methods

### NewControllerScanInterface

`func NewControllerScanInterface() *ControllerScanInterface`

NewControllerScanInterface instantiates a new ControllerScanInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerScanInterfaceWithDefaults

`func NewControllerScanInterfaceWithDefaults() *ControllerScanInterface`

NewControllerScanInterfaceWithDefaults instantiates a new ControllerScanInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *ControllerScanInterface) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ControllerScanInterface) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ControllerScanInterface) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ControllerScanInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpAddress

`func (o *ControllerScanInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ControllerScanInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ControllerScanInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ControllerScanInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *ControllerScanInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ControllerScanInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ControllerScanInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ControllerScanInterface) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


