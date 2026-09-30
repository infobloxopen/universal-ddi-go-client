# ExportQueryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | cidr_block filters the export results to objects within the given CIDR notation (e.g. \&quot;10.0.0.0/8\&quot;). | [optional] 
**Fts** | Pointer to **string** | fts is an optional full-text search query string used to restrict export results to objects whose searchable fields contain the given value. | [optional] 
**Node** | Pointer to **string** | node is the resource identifier used to scope the export query. For the _subnet_ view it should be the IP space ID or realm ID. | [optional] 
**PerspectiveType** | Pointer to **string** | perspective_type specifies how IP space objects are grouped. Allowed values are: * _network_, * _location_, * _address_. | [optional] 
**State** | Pointer to **string** | state filters address results by allocation state. Allowed values are: * _free_, * _used_, * _any_. | [optional] 
**Tfilter** | Pointer to **string** | tfilter is an optional tag filter expression (for example, &#x60;key&#x3D;&#x3D;\&quot;value\&quot;&#x60;) used to restrict export results to objects whose tags match the expression. | [optional] 
**View** | Pointer to **string** | view specifies the perspective table view type. Allowed values are: * _subnet_, * _rangeaddress_, * _flataddress_, * _all_. Defaults to _subnet_ when not specified. | [optional] 

## Methods

### NewExportQueryParams

`func NewExportQueryParams() *ExportQueryParams`

NewExportQueryParams instantiates a new ExportQueryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportQueryParamsWithDefaults

`func NewExportQueryParamsWithDefaults() *ExportQueryParams`

NewExportQueryParamsWithDefaults instantiates a new ExportQueryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *ExportQueryParams) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ExportQueryParams) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ExportQueryParams) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *ExportQueryParams) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetFts

`func (o *ExportQueryParams) GetFts() string`

GetFts returns the Fts field if non-nil, zero value otherwise.

### GetFtsOk

`func (o *ExportQueryParams) GetFtsOk() (*string, bool)`

GetFtsOk returns a tuple with the Fts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFts

`func (o *ExportQueryParams) SetFts(v string)`

SetFts sets Fts field to given value.

### HasFts

`func (o *ExportQueryParams) HasFts() bool`

HasFts returns a boolean if a field has been set.

### GetNode

`func (o *ExportQueryParams) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *ExportQueryParams) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *ExportQueryParams) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *ExportQueryParams) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetPerspectiveType

`func (o *ExportQueryParams) GetPerspectiveType() string`

GetPerspectiveType returns the PerspectiveType field if non-nil, zero value otherwise.

### GetPerspectiveTypeOk

`func (o *ExportQueryParams) GetPerspectiveTypeOk() (*string, bool)`

GetPerspectiveTypeOk returns a tuple with the PerspectiveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerspectiveType

`func (o *ExportQueryParams) SetPerspectiveType(v string)`

SetPerspectiveType sets PerspectiveType field to given value.

### HasPerspectiveType

`func (o *ExportQueryParams) HasPerspectiveType() bool`

HasPerspectiveType returns a boolean if a field has been set.

### GetState

`func (o *ExportQueryParams) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExportQueryParams) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExportQueryParams) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ExportQueryParams) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTfilter

`func (o *ExportQueryParams) GetTfilter() string`

GetTfilter returns the Tfilter field if non-nil, zero value otherwise.

### GetTfilterOk

`func (o *ExportQueryParams) GetTfilterOk() (*string, bool)`

GetTfilterOk returns a tuple with the Tfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfilter

`func (o *ExportQueryParams) SetTfilter(v string)`

SetTfilter sets Tfilter field to given value.

### HasTfilter

`func (o *ExportQueryParams) HasTfilter() bool`

HasTfilter returns a boolean if a field has been set.

### GetView

`func (o *ExportQueryParams) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ExportQueryParams) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ExportQueryParams) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ExportQueryParams) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


