# ExportTableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**ExportTableResult**](ExportTableResult.md) | The export request result. | [optional] [readonly] 

## Methods

### NewExportTableResponse

`func NewExportTableResponse() *ExportTableResponse`

NewExportTableResponse instantiates a new ExportTableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTableResponseWithDefaults

`func NewExportTableResponseWithDefaults() *ExportTableResponse`

NewExportTableResponseWithDefaults instantiates a new ExportTableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ExportTableResponse) GetResult() ExportTableResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExportTableResponse) GetResultOk() (*ExportTableResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExportTableResponse) SetResult(v ExportTableResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ExportTableResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


