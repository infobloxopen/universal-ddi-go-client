# ExportTableResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | The error message if the export request failed, otherwise empty. | [optional] [readonly] 
**ReqId** | Pointer to **string** | The resource identifier. | [optional] 
**Status** | Pointer to **string** | The current status of the export request.  Valid values are: * _pending_ * _processing_ * _complete_ * _failed_. | [optional] [readonly] 

## Methods

### NewExportTableResult

`func NewExportTableResult() *ExportTableResult`

NewExportTableResult instantiates a new ExportTableResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportTableResultWithDefaults

`func NewExportTableResultWithDefaults() *ExportTableResult`

NewExportTableResultWithDefaults instantiates a new ExportTableResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ExportTableResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ExportTableResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ExportTableResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ExportTableResult) HasError() bool`

HasError returns a boolean if a field has been set.

### GetReqId

`func (o *ExportTableResult) GetReqId() string`

GetReqId returns the ReqId field if non-nil, zero value otherwise.

### GetReqIdOk

`func (o *ExportTableResult) GetReqIdOk() (*string, bool)`

GetReqIdOk returns a tuple with the ReqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqId

`func (o *ExportTableResult) SetReqId(v string)`

SetReqId sets ReqId field to given value.

### HasReqId

`func (o *ExportTableResult) HasReqId() bool`

HasReqId returns a boolean if a field has been set.

### GetStatus

`func (o *ExportTableResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportTableResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportTableResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExportTableResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


