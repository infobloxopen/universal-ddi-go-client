# HostStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | Pointer to [**DestinationStatus**](DestinationStatus.md) | Destination-side status for the flow on this host. | [optional] 
**Etl** | Pointer to [**EtlStatus**](EtlStatus.md) | ETL/filter pipeline status for the flow on this host. | [optional] 
**HaState** | Pointer to **string** | High-availability role of this host for the flow (e.g. active, standby) when destination HA is enabled. Empty when HA is not applicable. | [optional] 
**Id** | Pointer to **int64** | Numeric identifier of the on-prem host as known to the data connector service. | [optional] [readonly] 
**Message** | Pointer to **string** | Human-readable status message for the flow on this host. | [optional] 
**Ophid** | Pointer to **string** | On-prem host identifier (ophid) assigned by the platform. | [optional] 
**Source** | Pointer to [**SourceStatus**](SourceStatus.md) | Source-side status for the flow on this host. | [optional] 
**State** | Pointer to **string** | High-level lifecycle state of the flow on this host (e.g. deployed, deploying, error). Agent-emitted string; not a stable enum. | [optional] 
**Status** | Pointer to **string** | Aggregated runtime status of the flow on this host, as reported by the on-prem agent. Agent-emitted string; not a stable enum. | [optional] 

## Methods

### NewHostStatus

`func NewHostStatus() *HostStatus`

NewHostStatus instantiates a new HostStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostStatusWithDefaults

`func NewHostStatusWithDefaults() *HostStatus`

NewHostStatusWithDefaults instantiates a new HostStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *HostStatus) GetDestination() DestinationStatus`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *HostStatus) GetDestinationOk() (*DestinationStatus, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *HostStatus) SetDestination(v DestinationStatus)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *HostStatus) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEtl

`func (o *HostStatus) GetEtl() EtlStatus`

GetEtl returns the Etl field if non-nil, zero value otherwise.

### GetEtlOk

`func (o *HostStatus) GetEtlOk() (*EtlStatus, bool)`

GetEtlOk returns a tuple with the Etl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtl

`func (o *HostStatus) SetEtl(v EtlStatus)`

SetEtl sets Etl field to given value.

### HasEtl

`func (o *HostStatus) HasEtl() bool`

HasEtl returns a boolean if a field has been set.

### GetHaState

`func (o *HostStatus) GetHaState() string`

GetHaState returns the HaState field if non-nil, zero value otherwise.

### GetHaStateOk

`func (o *HostStatus) GetHaStateOk() (*string, bool)`

GetHaStateOk returns a tuple with the HaState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaState

`func (o *HostStatus) SetHaState(v string)`

SetHaState sets HaState field to given value.

### HasHaState

`func (o *HostStatus) HasHaState() bool`

HasHaState returns a boolean if a field has been set.

### GetId

`func (o *HostStatus) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostStatus) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostStatus) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *HostStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *HostStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HostStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HostStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HostStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOphid

`func (o *HostStatus) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *HostStatus) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *HostStatus) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *HostStatus) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetSource

`func (o *HostStatus) GetSource() SourceStatus`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *HostStatus) GetSourceOk() (*SourceStatus, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *HostStatus) SetSource(v SourceStatus)`

SetSource sets Source field to given value.

### HasSource

`func (o *HostStatus) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *HostStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HostStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HostStatus) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HostStatus) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *HostStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


