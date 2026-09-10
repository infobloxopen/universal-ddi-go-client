# AuthZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone configuration. | [optional] 
**CompartmentId** | Pointer to **string** | The access view associated with the object. If no access view is associated with the object, the value defaults to empty. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**DnssecKeys** | Pointer to [**[]DNSSECKey**](DNSSECKey.md) | The list of DNSSEC keys used by the _AuthZone_ for zone signing. | [optional] [readonly] 
**DnssecSigningPolicy** | Pointer to [**DNSSECSigningPolicy**](DNSSECSigningPolicy.md) | Optional. DNSSEC zone signing settings. | [optional] 
**DnssecStatus** | Pointer to **string** | Read Only.  DNSSEC status indicates the current DNSSEC signing status of the zone.  Possible values: - _UNSIGNED_: The zone is not signed with DNSSEC - _SIGNED_: The zone is fully signed with DNSSEC - _ROLLOVER_IN_PROGRESS_: DNSSEC key rollover is currently in progress - _SIGN_IN_PROGRESS_: The zone is currently being signed with DNSSEC - _UNSIGN_IN_PROGRESS_: The zone is currently being unsigned (DNSSEC removal in progress) | [optional] [readonly] 
**ExternalPrimaries** | Pointer to [**[]ExternalPrimary**](ExternalPrimary.md) | Optional. DNS primaries external to Universal DDI. Order is not significant. | [optional] 
**ExternalProviders** | Pointer to [**[]AuthZoneExternalProvider**](AuthZoneExternalProvider.md) | list of external providers for the auth zone. | [optional] [readonly] 
**ExternalProvidersMetadata** | Pointer to **map[string]interface{}** | External DNS providers metadata. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ExternalSecondary**](ExternalSecondary.md) | DNS secondaries external to Universal DDI. Order is not significant. | [optional] 
**Fqdn** | Pointer to **string** | Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation. | [optional] 
**GridPrimaries** | Pointer to [**[]MemberServer**](MemberServer.md) | Optional. The list of the NIOS Grid Primaries assigned to an AuthZone, only applicable for the NIOS Zones. | [optional] 
**GridSecondaries** | Pointer to [**[]MemberServer**](MemberServer.md) | Optional. The list of the NIOS Grid Secondaries assigned to an AuthZone, only applicable for the NIOS Zones. | [optional] 
**GssTsigEnabled** | Pointer to **bool** | _gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]Inheritance2AssignedHost**](Inheritance2AssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**AuthZoneInheritance**](AuthZoneInheritance.md) | Optional. Inheritance configuration. | [optional] 
**InitialSoaSerial** | Pointer to **int64** | On-create-only. SOA serial is allowed to be set when the authoritative zone is created. | [optional] 
**InternalSecondaries** | Pointer to [**[]InternalSecondary**](InternalSecondary.md) | Optional. Universal DDI hosts acting as internal secondaries. Order is not significant. | [optional] 
**MappedSubnet** | Pointer to **string** | Reverse zone network address in the following format: \&quot;ip-address/cidr\&quot;. Defaults to empty. | [optional] [readonly] 
**Mapping** | Pointer to **string** | Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to forward. | [optional] [readonly] 
**MaxRecordsPerType** | Pointer to **int64** | The maximum number of records that can be stored in an RRset (records of same name and type), to prevent a slowdown in query processing due to an excessive number of those RRsets. The limit is enforced when serving the zone on-prem, not at the time of record creation or update. Exceeding the limit will result in the zone failing to load or to be updated. If 0, it means there is no limit. Defaults to _2000_. | [optional] 
**MaxTypesPerName** | Pointer to **int64** | The maximum number of record types that can be stored for an owner name, to prevent a slowdown in query processing due to an excessive number of those records. The limit is enforced when serving the zone on-prem, not at the time of record creation or update. Exceeding the limit will result in the zone failing to load or to be updated. If 0, it means there is no limit. Defaults to _100_. | [optional] 
**Nameservers** | Pointer to [**[]Nameserver**](Nameserver.md) | Optional. A list of DNS Nameservers of various roles. Cannot be configured if _nsg_ is configured. | [optional] 
**NiosGridsMetadata** | Pointer to **map[string]interface{}** | NIOS Grids Metadata holds multiple NIOS grids data. | [optional] 
**Notify** | Pointer to **bool** | Also notify all external secondary DNS servers if enabled.  Defaults to _false_. | [optional] 
**Nsg** | Pointer to **string** | The resource identifier. | [optional] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**PrimaryType** | Pointer to **string** | Primary type for an authoritative zone. Read only after creation. Allowed values:  * _external_: zone data owned by an external nameserver,  * _cloud_: zone data is owned by a Universal DDI host. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Zone FQDN in punycode. | [optional] [readonly] 
**QueryAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty. | [optional] 
**SecondaryZoneRecordsSync** | Pointer to **bool** | Optional. Defines if secondary zone records should be synchronized.  Defaults to _false_. Only allowed to update when primary_type is \&quot;external\&quot;. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**TransferAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Clients must match this ACL to receive zone transfers. | [optional] 
**UpdateAcl** | Pointer to [**[]ACLItem**](ACLItem.md) | Optional. Specifies which hosts are allowed to submit Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**UseForwardersForSubzones** | Pointer to **bool** | Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_. | [optional] 
**Version** | Pointer to **string** | Read Only.  Version indicates the version of the zone in context of assigned DNS NSGs and nameservers.  Possible values: - _v1_: The zone uses original NSG model - _v2_: The zone uses new \&quot;Unified Nameservers\&quot; model | [optional] [readonly] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**Warnings** | Pointer to [**[]Warning**](Warning.md) | The list of an auth zone warnings. | [optional] [readonly] 
**ZoneAuthority** | Pointer to [**ZoneAuthority**](ZoneAuthority.md) | Optional. ZoneAuthority. | [optional] 

## Methods

### NewAuthZone

`func NewAuthZone() *AuthZone`

NewAuthZone instantiates a new AuthZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthZoneWithDefaults

`func NewAuthZoneWithDefaults() *AuthZone`

NewAuthZoneWithDefaults instantiates a new AuthZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *AuthZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AuthZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AuthZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AuthZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCompartmentId

`func (o *AuthZone) GetCompartmentId() string`

GetCompartmentId returns the CompartmentId field if non-nil, zero value otherwise.

### GetCompartmentIdOk

`func (o *AuthZone) GetCompartmentIdOk() (*string, bool)`

GetCompartmentIdOk returns a tuple with the CompartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompartmentId

`func (o *AuthZone) SetCompartmentId(v string)`

SetCompartmentId sets CompartmentId field to given value.

### HasCompartmentId

`func (o *AuthZone) HasCompartmentId() bool`

HasCompartmentId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *AuthZone) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AuthZone) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AuthZone) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AuthZone) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDnssecKeys

`func (o *AuthZone) GetDnssecKeys() []DNSSECKey`

GetDnssecKeys returns the DnssecKeys field if non-nil, zero value otherwise.

### GetDnssecKeysOk

`func (o *AuthZone) GetDnssecKeysOk() (*[]DNSSECKey, bool)`

GetDnssecKeysOk returns a tuple with the DnssecKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeys

`func (o *AuthZone) SetDnssecKeys(v []DNSSECKey)`

SetDnssecKeys sets DnssecKeys field to given value.

### HasDnssecKeys

`func (o *AuthZone) HasDnssecKeys() bool`

HasDnssecKeys returns a boolean if a field has been set.

### GetDnssecSigningPolicy

`func (o *AuthZone) GetDnssecSigningPolicy() DNSSECSigningPolicy`

GetDnssecSigningPolicy returns the DnssecSigningPolicy field if non-nil, zero value otherwise.

### GetDnssecSigningPolicyOk

`func (o *AuthZone) GetDnssecSigningPolicyOk() (*DNSSECSigningPolicy, bool)`

GetDnssecSigningPolicyOk returns a tuple with the DnssecSigningPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecSigningPolicy

`func (o *AuthZone) SetDnssecSigningPolicy(v DNSSECSigningPolicy)`

SetDnssecSigningPolicy sets DnssecSigningPolicy field to given value.

### HasDnssecSigningPolicy

`func (o *AuthZone) HasDnssecSigningPolicy() bool`

HasDnssecSigningPolicy returns a boolean if a field has been set.

### GetDnssecStatus

`func (o *AuthZone) GetDnssecStatus() string`

GetDnssecStatus returns the DnssecStatus field if non-nil, zero value otherwise.

### GetDnssecStatusOk

`func (o *AuthZone) GetDnssecStatusOk() (*string, bool)`

GetDnssecStatusOk returns a tuple with the DnssecStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecStatus

`func (o *AuthZone) SetDnssecStatus(v string)`

SetDnssecStatus sets DnssecStatus field to given value.

### HasDnssecStatus

`func (o *AuthZone) HasDnssecStatus() bool`

HasDnssecStatus returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *AuthZone) GetExternalPrimaries() []ExternalPrimary`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *AuthZone) GetExternalPrimariesOk() (*[]ExternalPrimary, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *AuthZone) SetExternalPrimaries(v []ExternalPrimary)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *AuthZone) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalProviders

`func (o *AuthZone) GetExternalProviders() []AuthZoneExternalProvider`

GetExternalProviders returns the ExternalProviders field if non-nil, zero value otherwise.

### GetExternalProvidersOk

`func (o *AuthZone) GetExternalProvidersOk() (*[]AuthZoneExternalProvider, bool)`

GetExternalProvidersOk returns a tuple with the ExternalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProviders

`func (o *AuthZone) SetExternalProviders(v []AuthZoneExternalProvider)`

SetExternalProviders sets ExternalProviders field to given value.

### HasExternalProviders

`func (o *AuthZone) HasExternalProviders() bool`

HasExternalProviders returns a boolean if a field has been set.

### GetExternalProvidersMetadata

`func (o *AuthZone) GetExternalProvidersMetadata() map[string]interface{}`

GetExternalProvidersMetadata returns the ExternalProvidersMetadata field if non-nil, zero value otherwise.

### GetExternalProvidersMetadataOk

`func (o *AuthZone) GetExternalProvidersMetadataOk() (*map[string]interface{}, bool)`

GetExternalProvidersMetadataOk returns a tuple with the ExternalProvidersMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProvidersMetadata

`func (o *AuthZone) SetExternalProvidersMetadata(v map[string]interface{})`

SetExternalProvidersMetadata sets ExternalProvidersMetadata field to given value.

### HasExternalProvidersMetadata

`func (o *AuthZone) HasExternalProvidersMetadata() bool`

HasExternalProvidersMetadata returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *AuthZone) GetExternalSecondaries() []ExternalSecondary`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *AuthZone) GetExternalSecondariesOk() (*[]ExternalSecondary, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *AuthZone) SetExternalSecondaries(v []ExternalSecondary)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *AuthZone) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetFqdn

`func (o *AuthZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *AuthZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *AuthZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *AuthZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGridPrimaries

`func (o *AuthZone) GetGridPrimaries() []MemberServer`

GetGridPrimaries returns the GridPrimaries field if non-nil, zero value otherwise.

### GetGridPrimariesOk

`func (o *AuthZone) GetGridPrimariesOk() (*[]MemberServer, bool)`

GetGridPrimariesOk returns a tuple with the GridPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimaries

`func (o *AuthZone) SetGridPrimaries(v []MemberServer)`

SetGridPrimaries sets GridPrimaries field to given value.

### HasGridPrimaries

`func (o *AuthZone) HasGridPrimaries() bool`

HasGridPrimaries returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *AuthZone) GetGridSecondaries() []MemberServer`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *AuthZone) GetGridSecondariesOk() (*[]MemberServer, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *AuthZone) SetGridSecondaries(v []MemberServer)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *AuthZone) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetGssTsigEnabled

`func (o *AuthZone) GetGssTsigEnabled() bool`

GetGssTsigEnabled returns the GssTsigEnabled field if non-nil, zero value otherwise.

### GetGssTsigEnabledOk

`func (o *AuthZone) GetGssTsigEnabledOk() (*bool, bool)`

GetGssTsigEnabledOk returns a tuple with the GssTsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGssTsigEnabled

`func (o *AuthZone) SetGssTsigEnabled(v bool)`

SetGssTsigEnabled sets GssTsigEnabled field to given value.

### HasGssTsigEnabled

`func (o *AuthZone) HasGssTsigEnabled() bool`

HasGssTsigEnabled returns a boolean if a field has been set.

### GetId

`func (o *AuthZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *AuthZone) GetInheritanceAssignedHosts() []Inheritance2AssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *AuthZone) GetInheritanceAssignedHostsOk() (*[]Inheritance2AssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *AuthZone) SetInheritanceAssignedHosts(v []Inheritance2AssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *AuthZone) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *AuthZone) GetInheritanceSources() AuthZoneInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *AuthZone) GetInheritanceSourcesOk() (*AuthZoneInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *AuthZone) SetInheritanceSources(v AuthZoneInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *AuthZone) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetInitialSoaSerial

`func (o *AuthZone) GetInitialSoaSerial() int64`

GetInitialSoaSerial returns the InitialSoaSerial field if non-nil, zero value otherwise.

### GetInitialSoaSerialOk

`func (o *AuthZone) GetInitialSoaSerialOk() (*int64, bool)`

GetInitialSoaSerialOk returns a tuple with the InitialSoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSoaSerial

`func (o *AuthZone) SetInitialSoaSerial(v int64)`

SetInitialSoaSerial sets InitialSoaSerial field to given value.

### HasInitialSoaSerial

`func (o *AuthZone) HasInitialSoaSerial() bool`

HasInitialSoaSerial returns a boolean if a field has been set.

### GetInternalSecondaries

`func (o *AuthZone) GetInternalSecondaries() []InternalSecondary`

GetInternalSecondaries returns the InternalSecondaries field if non-nil, zero value otherwise.

### GetInternalSecondariesOk

`func (o *AuthZone) GetInternalSecondariesOk() (*[]InternalSecondary, bool)`

GetInternalSecondariesOk returns a tuple with the InternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalSecondaries

`func (o *AuthZone) SetInternalSecondaries(v []InternalSecondary)`

SetInternalSecondaries sets InternalSecondaries field to given value.

### HasInternalSecondaries

`func (o *AuthZone) HasInternalSecondaries() bool`

HasInternalSecondaries returns a boolean if a field has been set.

### GetMappedSubnet

`func (o *AuthZone) GetMappedSubnet() string`

GetMappedSubnet returns the MappedSubnet field if non-nil, zero value otherwise.

### GetMappedSubnetOk

`func (o *AuthZone) GetMappedSubnetOk() (*string, bool)`

GetMappedSubnetOk returns a tuple with the MappedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedSubnet

`func (o *AuthZone) SetMappedSubnet(v string)`

SetMappedSubnet sets MappedSubnet field to given value.

### HasMappedSubnet

`func (o *AuthZone) HasMappedSubnet() bool`

HasMappedSubnet returns a boolean if a field has been set.

### GetMapping

`func (o *AuthZone) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *AuthZone) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *AuthZone) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *AuthZone) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetMaxRecordsPerType

`func (o *AuthZone) GetMaxRecordsPerType() int64`

GetMaxRecordsPerType returns the MaxRecordsPerType field if non-nil, zero value otherwise.

### GetMaxRecordsPerTypeOk

`func (o *AuthZone) GetMaxRecordsPerTypeOk() (*int64, bool)`

GetMaxRecordsPerTypeOk returns a tuple with the MaxRecordsPerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecordsPerType

`func (o *AuthZone) SetMaxRecordsPerType(v int64)`

SetMaxRecordsPerType sets MaxRecordsPerType field to given value.

### HasMaxRecordsPerType

`func (o *AuthZone) HasMaxRecordsPerType() bool`

HasMaxRecordsPerType returns a boolean if a field has been set.

### GetMaxTypesPerName

`func (o *AuthZone) GetMaxTypesPerName() int64`

GetMaxTypesPerName returns the MaxTypesPerName field if non-nil, zero value otherwise.

### GetMaxTypesPerNameOk

`func (o *AuthZone) GetMaxTypesPerNameOk() (*int64, bool)`

GetMaxTypesPerNameOk returns a tuple with the MaxTypesPerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTypesPerName

`func (o *AuthZone) SetMaxTypesPerName(v int64)`

SetMaxTypesPerName sets MaxTypesPerName field to given value.

### HasMaxTypesPerName

`func (o *AuthZone) HasMaxTypesPerName() bool`

HasMaxTypesPerName returns a boolean if a field has been set.

### GetNameservers

`func (o *AuthZone) GetNameservers() []Nameserver`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *AuthZone) GetNameserversOk() (*[]Nameserver, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *AuthZone) SetNameservers(v []Nameserver)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *AuthZone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetNiosGridsMetadata

`func (o *AuthZone) GetNiosGridsMetadata() map[string]interface{}`

GetNiosGridsMetadata returns the NiosGridsMetadata field if non-nil, zero value otherwise.

### GetNiosGridsMetadataOk

`func (o *AuthZone) GetNiosGridsMetadataOk() (*map[string]interface{}, bool)`

GetNiosGridsMetadataOk returns a tuple with the NiosGridsMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNiosGridsMetadata

`func (o *AuthZone) SetNiosGridsMetadata(v map[string]interface{})`

SetNiosGridsMetadata sets NiosGridsMetadata field to given value.

### HasNiosGridsMetadata

`func (o *AuthZone) HasNiosGridsMetadata() bool`

HasNiosGridsMetadata returns a boolean if a field has been set.

### GetNotify

`func (o *AuthZone) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *AuthZone) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *AuthZone) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *AuthZone) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetNsg

`func (o *AuthZone) GetNsg() string`

GetNsg returns the Nsg field if non-nil, zero value otherwise.

### GetNsgOk

`func (o *AuthZone) GetNsgOk() (*string, bool)`

GetNsgOk returns a tuple with the Nsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsg

`func (o *AuthZone) SetNsg(v string)`

SetNsg sets Nsg field to given value.

### HasNsg

`func (o *AuthZone) HasNsg() bool`

HasNsg returns a boolean if a field has been set.

### GetNsgs

`func (o *AuthZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *AuthZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *AuthZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *AuthZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetParent

`func (o *AuthZone) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *AuthZone) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *AuthZone) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *AuthZone) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrimaryType

`func (o *AuthZone) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *AuthZone) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *AuthZone) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *AuthZone) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *AuthZone) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *AuthZone) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *AuthZone) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *AuthZone) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetQueryAcl

`func (o *AuthZone) GetQueryAcl() []ACLItem`

GetQueryAcl returns the QueryAcl field if non-nil, zero value otherwise.

### GetQueryAclOk

`func (o *AuthZone) GetQueryAclOk() (*[]ACLItem, bool)`

GetQueryAclOk returns a tuple with the QueryAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryAcl

`func (o *AuthZone) SetQueryAcl(v []ACLItem)`

SetQueryAcl sets QueryAcl field to given value.

### HasQueryAcl

`func (o *AuthZone) HasQueryAcl() bool`

HasQueryAcl returns a boolean if a field has been set.

### GetSecondaryZoneRecordsSync

`func (o *AuthZone) GetSecondaryZoneRecordsSync() bool`

GetSecondaryZoneRecordsSync returns the SecondaryZoneRecordsSync field if non-nil, zero value otherwise.

### GetSecondaryZoneRecordsSyncOk

`func (o *AuthZone) GetSecondaryZoneRecordsSyncOk() (*bool, bool)`

GetSecondaryZoneRecordsSyncOk returns a tuple with the SecondaryZoneRecordsSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryZoneRecordsSync

`func (o *AuthZone) SetSecondaryZoneRecordsSync(v bool)`

SetSecondaryZoneRecordsSync sets SecondaryZoneRecordsSync field to given value.

### HasSecondaryZoneRecordsSync

`func (o *AuthZone) HasSecondaryZoneRecordsSync() bool`

HasSecondaryZoneRecordsSync returns a boolean if a field has been set.

### GetTags

`func (o *AuthZone) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AuthZone) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AuthZone) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AuthZone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransferAcl

`func (o *AuthZone) GetTransferAcl() []ACLItem`

GetTransferAcl returns the TransferAcl field if non-nil, zero value otherwise.

### GetTransferAclOk

`func (o *AuthZone) GetTransferAclOk() (*[]ACLItem, bool)`

GetTransferAclOk returns a tuple with the TransferAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAcl

`func (o *AuthZone) SetTransferAcl(v []ACLItem)`

SetTransferAcl sets TransferAcl field to given value.

### HasTransferAcl

`func (o *AuthZone) HasTransferAcl() bool`

HasTransferAcl returns a boolean if a field has been set.

### GetUpdateAcl

`func (o *AuthZone) GetUpdateAcl() []ACLItem`

GetUpdateAcl returns the UpdateAcl field if non-nil, zero value otherwise.

### GetUpdateAclOk

`func (o *AuthZone) GetUpdateAclOk() (*[]ACLItem, bool)`

GetUpdateAclOk returns a tuple with the UpdateAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAcl

`func (o *AuthZone) SetUpdateAcl(v []ACLItem)`

SetUpdateAcl sets UpdateAcl field to given value.

### HasUpdateAcl

`func (o *AuthZone) HasUpdateAcl() bool`

HasUpdateAcl returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseForwardersForSubzones

`func (o *AuthZone) GetUseForwardersForSubzones() bool`

GetUseForwardersForSubzones returns the UseForwardersForSubzones field if non-nil, zero value otherwise.

### GetUseForwardersForSubzonesOk

`func (o *AuthZone) GetUseForwardersForSubzonesOk() (*bool, bool)`

GetUseForwardersForSubzonesOk returns a tuple with the UseForwardersForSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseForwardersForSubzones

`func (o *AuthZone) SetUseForwardersForSubzones(v bool)`

SetUseForwardersForSubzones sets UseForwardersForSubzones field to given value.

### HasUseForwardersForSubzones

`func (o *AuthZone) HasUseForwardersForSubzones() bool`

HasUseForwardersForSubzones returns a boolean if a field has been set.

### GetVersion

`func (o *AuthZone) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AuthZone) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AuthZone) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AuthZone) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetView

`func (o *AuthZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *AuthZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *AuthZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *AuthZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWarnings

`func (o *AuthZone) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AuthZone) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AuthZone) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AuthZone) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetZoneAuthority

`func (o *AuthZone) GetZoneAuthority() ZoneAuthority`

GetZoneAuthority returns the ZoneAuthority field if non-nil, zero value otherwise.

### GetZoneAuthorityOk

`func (o *AuthZone) GetZoneAuthorityOk() (*ZoneAuthority, bool)`

GetZoneAuthorityOk returns a tuple with the ZoneAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAuthority

`func (o *AuthZone) SetZoneAuthority(v ZoneAuthority)`

SetZoneAuthority sets ZoneAuthority field to given value.

### HasZoneAuthority

`func (o *AuthZone) HasZoneAuthority() bool`

HasZoneAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


