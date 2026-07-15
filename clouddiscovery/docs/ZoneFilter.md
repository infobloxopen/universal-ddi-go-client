# ZoneFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Wildcards** | Pointer to **[]string** |  | [optional] 

## Methods

### NewZoneFilter

`func NewZoneFilter() *ZoneFilter`

NewZoneFilter instantiates a new ZoneFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneFilterWithDefaults

`func NewZoneFilterWithDefaults() *ZoneFilter`

NewZoneFilterWithDefaults instantiates a new ZoneFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ZoneFilter) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ZoneFilter) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ZoneFilter) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ZoneFilter) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetWildcards

`func (o *ZoneFilter) GetWildcards() []string`

GetWildcards returns the Wildcards field if non-nil, zero value otherwise.

### GetWildcardsOk

`func (o *ZoneFilter) GetWildcardsOk() (*[]string, bool)`

GetWildcardsOk returns a tuple with the Wildcards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcards

`func (o *ZoneFilter) SetWildcards(v []string)`

SetWildcards sets Wildcards field to given value.

### HasWildcards

`func (o *ZoneFilter) HasWildcards() bool`

HasWildcards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


