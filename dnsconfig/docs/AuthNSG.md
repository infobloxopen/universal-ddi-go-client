# AuthNSG

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for the object. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ExternalPrimary**](ExternalPrimary.md) | Optional. DNS primaries external to Universal DDI. Order is not significant. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ExternalSecondary**](ExternalSecondary.md) | DNS secondaries external to Universal DDI. Order is not significant. | [optional] 
**GridPrimaries** | Pointer to [**[]MemberServer**](MemberServer.md) | Optional. The list of the NIOS Grid Primaries assigned to an AuthNSG, only applicable for the NIOS. | [optional] 
**GridSecondaries** | Pointer to [**[]MemberServer**](MemberServer.md) | Optional. The list of the NIOS Grid Secondaries assigned to an AuthNSG, only applicable for the NIOS. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalSecondaries** | Pointer to [**[]InternalSecondary**](InternalSecondary.md) | Optional. Universal DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**Name** | **string** | Name of the object. | 
**Nameservers** | Pointer to [**[]Nameserver**](Nameserver.md) | Optional. A list of DNS Nameservers of various roles. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**Version** | Pointer to **string** | Read Only.  Version indicates the version of the Authoritative DNS Server Group in context of DNS NSGs and nameservers that are used.  Possible values: - _v1_: The Authoritative DNS Server Group uses original NSG model - _v2_: The Authoritative DNS Server Group uses new \&quot;Unified Nameservers\&quot; model | [optional] [readonly] 

## Methods

### NewAuthNSG

`func NewAuthNSG(name string, ) *AuthNSG`

NewAuthNSG instantiates a new AuthNSG object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthNSGWithDefaults

`func NewAuthNSGWithDefaults() *AuthNSG`

NewAuthNSGWithDefaults instantiates a new AuthNSG object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *AuthNSG) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuthNSG) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuthNSG) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuthNSG) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *AuthNSG) GetExternalPrimaries() []ExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *AuthNSG) GetExternalPrimariesOk() (*[]ExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *AuthNSG) SetExternalPrimaries(v []ExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *AuthNSG) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *AuthNSG) GetExternalSecondaries() []ExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *AuthNSG) GetExternalSecondariesOk() (*[]ExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *AuthNSG) SetExternalSecondaries(v []ExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *AuthNSG) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetGridPrimaries

`func (o *AuthNSG) GetGridPrimaries() []MemberServer`

GetGridPrimaries returns the GridPrimaries field if non-nil, zero value otherwise.

### GetGridPrimariesOk

`func (o *AuthNSG) GetGridPrimariesOk() (*[]MemberServer, bool)`

GetGridPrimariesOk returns a tuple with the GridPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimaries

`func (o *AuthNSG) SetGridPrimaries(v []MemberServer)`

SetGridPrimaries sets GridPrimaries field to given value.

### HasGridPrimaries

`func (o *AuthNSG) HasGridPrimaries() bool`

HasGridPrimaries returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *AuthNSG) GetGridSecondaries() []MemberServer`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *AuthNSG) GetGridSecondariesOk() (*[]MemberServer, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *AuthNSG) SetGridSecondaries(v []MemberServer)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *AuthNSG) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetId

`func (o *AuthNSG) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthNSG) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthNSG) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthNSG) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *AuthNSG) GetInternalSecondaries() []InternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *AuthNSG) GetInternalSecondariesOk() (*[]InternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *AuthNSG) SetInternalSecondaries(v []InternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *AuthNSG) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetName

`func (o *AuthNSG) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthNSG) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthNSG) SetName(v string)`

SetName sets Name field to given value.


### GetNameservers

`func (o *AuthNSG) GetNameservers() []Nameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *AuthNSG) GetNameserversOk() (*[]Nameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *AuthNSG) SetNameservers(v []Nameserver)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *AuthNSG) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetNsgs

`func (o *AuthNSG) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *AuthNSG) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *AuthNSG) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *AuthNSG) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetTags

`func (o *AuthNSG) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuthNSG) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuthNSG) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuthNSG) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *AuthNSG) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthNSG) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthNSG) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthNSG) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


