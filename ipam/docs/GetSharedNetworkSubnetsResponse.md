# GetSharedNetworkSubnetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SharedNetworkSubnet**](SharedNetworkSubnet.md) | The list of subnets that are members of the shared network. | [optional] 

## Methods

### NewGetSharedNetworkSubnetsResponse

`func NewGetSharedNetworkSubnetsResponse() *GetSharedNetworkSubnetsResponse`

NewGetSharedNetworkSubnetsResponse instantiates a new GetSharedNetworkSubnetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSharedNetworkSubnetsResponseWithDefaults

`func NewGetSharedNetworkSubnetsResponseWithDefaults() *GetSharedNetworkSubnetsResponse`

NewGetSharedNetworkSubnetsResponseWithDefaults instantiates a new GetSharedNetworkSubnetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GetSharedNetworkSubnetsResponse) GetResults() []SharedNetworkSubnet`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetSharedNetworkSubnetsResponse) GetResultsOk() (*[]SharedNetworkSubnet, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetSharedNetworkSubnetsResponse) SetResults(v []SharedNetworkSubnet)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetSharedNetworkSubnetsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


