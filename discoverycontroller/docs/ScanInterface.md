# ScanInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | Pointer to **string** | Represents the interface to bind to. | [optional] 
**IpAddress** | Pointer to **string** | Represents the ip address configured on the interface. | [optional] 
**Name** | Pointer to **string** | Readable name given for this configuration. | [optional] 

## Methods

### NewScanInterface

`func NewScanInterface() *ScanInterface`

NewScanInterface instantiates a new ScanInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanInterfaceWithDefaults

`func NewScanInterfaceWithDefaults() *ScanInterface`

NewScanInterfaceWithDefaults instantiates a new ScanInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *ScanInterface) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ScanInterface) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ScanInterface) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ScanInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetIpAddress

`func (o *ScanInterface) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ScanInterface) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ScanInterface) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ScanInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *ScanInterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScanInterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScanInterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScanInterface) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


