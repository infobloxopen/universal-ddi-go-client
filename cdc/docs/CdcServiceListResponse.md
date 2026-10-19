# CdcServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]CdcService**](CdcService.md) |  | [optional] 

## Methods

### NewCdcServiceListResponse

`func NewCdcServiceListResponse() *CdcServiceListResponse`

NewCdcServiceListResponse instantiates a new CdcServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdcServiceListResponseWithDefaults

`func NewCdcServiceListResponseWithDefaults() *CdcServiceListResponse`

NewCdcServiceListResponseWithDefaults instantiates a new CdcServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *CdcServiceListResponse) GetResults() []CdcService`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *CdcServiceListResponse) GetResultsOk() (*[]CdcService, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *CdcServiceListResponse) SetResults(v []CdcService)`

SetResults sets Results field to given value.

### HasResults

`func (o *CdcServiceListResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


