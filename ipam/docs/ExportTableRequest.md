# ExportTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | Pointer to **string** | The filter expression used to select the objects to export. | [optional] 
**Fields** | **[]string** | The list of field names to include in the export output. | 
**QueryParams** | Pointer to [**ExportQueryParams**](ExportQueryParams.md) | query_params holds optional perspective-table query parameters. Storing them as a nested message allows future parameters to be added without changing this field list. | [optional] 
**SelectedObjs** | Pointer to **[]string** | The list of resource types of the objects selected for export. | [optional] 

## Methods

### NewExportTableRequest

`func NewExportTableRequest(fields []string, ) *ExportTableRequest`

NewExportTableRequest instantiates a new ExportTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTableRequestWithDefaults

`func NewExportTableRequestWithDefaults() *ExportTableRequest`

NewExportTableRequestWithDefaults instantiates a new ExportTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *ExportTableRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ExportTableRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ExportTableRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ExportTableRequest) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetFields

`func (o *ExportTableRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ExportTableRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ExportTableRequest) SetFields(v []string)`

SetFields sets Fields field to given value.


### GetQueryParams

`func (o *ExportTableRequest) GetQueryParams() ExportQueryParams`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *ExportTableRequest) GetQueryParamsOk() (*ExportQueryParams, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *ExportTableRequest) SetQueryParams(v ExportQueryParams)`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *ExportTableRequest) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetSelectedObjs

`func (o *ExportTableRequest) GetSelectedObjs() []string`

GetSelectedObjs returns the SelectedObjs field if non-nil, zero value otherwise.

### GetSelectedObjsOk

`func (o *ExportTableRequest) GetSelectedObjsOk() (*[]string, bool)`

GetSelectedObjsOk returns a tuple with the SelectedObjs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedObjs

`func (o *ExportTableRequest) SetSelectedObjs(v []string)`

SetSelectedObjs sets SelectedObjs field to given value.

### HasSelectedObjs

`func (o *ExportTableRequest) HasSelectedObjs() bool`

HasSelectedObjs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


