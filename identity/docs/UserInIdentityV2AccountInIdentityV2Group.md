# UserInIdentityV2AccountInIdentityV2Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountCspId** | Pointer to **int32** |  | [optional] 
**AccountId** | Pointer to **string** | The resource identifier. | [optional] 
**AccountInfobloxId** | Pointer to **string** |  | [optional] 
**AgreementsAccepted** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to **string** |  | [optional] 
**ApiKeys** | Pointer to **[]string** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**Authenticator** | Pointer to **string** |  | [optional] 
**ConfirmToken** | Pointer to **string** |  | [optional] 
**ConfirmationSentAt** | Pointer to **time.Time** |  | [optional] 
**ConfirmedAt** | Pointer to **time.Time** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CspId** | Pointer to **int32** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedBy** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FailedAttempts** | Pointer to **int32** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**GroupIds** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JobTitle** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **time.Time** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**LastValidatedAt** | Pointer to **time.Time** |  | [optional] 
**LockedAt** | Pointer to **time.Time** |  | [optional] 
**MarketingOptIn** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextRevalidationDueAt** | Pointer to **time.Time** |  | [optional] 
**Origin** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**ProfileValidationRequired** | Pointer to **bool** |  | [optional] 
**ResetPasswordSentAt** | Pointer to **time.Time** |  | [optional] 
**ResetPasswordToken** | Pointer to **string** |  | [optional] 
**SignInCount** | Pointer to **int32** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateChangedAt** | Pointer to **time.Time** |  | [optional] 
**StateRegion** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnlockToken** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUserInIdentityV2AccountInIdentityV2Group

`func NewUserInIdentityV2AccountInIdentityV2Group() *UserInIdentityV2AccountInIdentityV2Group`

NewUserInIdentityV2AccountInIdentityV2Group instantiates a new UserInIdentityV2AccountInIdentityV2Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInIdentityV2AccountInIdentityV2GroupWithDefaults

`func NewUserInIdentityV2AccountInIdentityV2GroupWithDefaults() *UserInIdentityV2AccountInIdentityV2Group`

NewUserInIdentityV2AccountInIdentityV2GroupWithDefaults instantiates a new UserInIdentityV2AccountInIdentityV2Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountCspId() int32`

GetAccountCspId returns the AccountCspId field if non-nil, zero value otherwise.

### GetAccountCspIdOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountCspIdOk() (*int32, bool)`

GetAccountCspIdOk returns a tuple with the AccountCspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetAccountCspId(v int32)`

SetAccountCspId sets AccountCspId field to given value.

### HasAccountCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasAccountCspId() bool`

HasAccountCspId returns a boolean if a field has been set.

### GetAccountId

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountInfobloxId

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountInfobloxId() string`

GetAccountInfobloxId returns the AccountInfobloxId field if non-nil, zero value otherwise.

### GetAccountInfobloxIdOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAccountInfobloxIdOk() (*string, bool)`

GetAccountInfobloxIdOk returns a tuple with the AccountInfobloxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfobloxId

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetAccountInfobloxId(v string)`

SetAccountInfobloxId sets AccountInfobloxId field to given value.

### HasAccountInfobloxId

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasAccountInfobloxId() bool`

HasAccountInfobloxId returns a boolean if a field has been set.

### GetAgreementsAccepted

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAgreementsAccepted() bool`

GetAgreementsAccepted returns the AgreementsAccepted field if non-nil, zero value otherwise.

### GetAgreementsAcceptedOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAgreementsAcceptedOk() (*bool, bool)`

GetAgreementsAcceptedOk returns a tuple with the AgreementsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementsAccepted

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetAgreementsAccepted(v bool)`

SetAgreementsAccepted sets AgreementsAccepted field to given value.

### HasAgreementsAccepted

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasAgreementsAccepted() bool`

HasAgreementsAccepted returns a boolean if a field has been set.

### GetApiKey

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeys

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApiKeys() []string`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApiKeysOk() (*[]string, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetApiKeys(v []string)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetApproved

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAuthenticator() string`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetAuthenticatorOk() (*string, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetAuthenticator(v string)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetConfirmToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmToken() string`

GetConfirmToken returns the ConfirmToken field if non-nil, zero value otherwise.

### GetConfirmTokenOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmTokenOk() (*string, bool)`

GetConfirmTokenOk returns a tuple with the ConfirmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetConfirmToken(v string)`

SetConfirmToken sets ConfirmToken field to given value.

### HasConfirmToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasConfirmToken() bool`

HasConfirmToken returns a boolean if a field has been set.

### GetConfirmationSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmationSentAt() time.Time`

GetConfirmationSentAt returns the ConfirmationSentAt field if non-nil, zero value otherwise.

### GetConfirmationSentAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmationSentAtOk() (*time.Time, bool)`

GetConfirmationSentAtOk returns a tuple with the ConfirmationSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetConfirmationSentAt(v time.Time)`

SetConfirmationSentAt sets ConfirmationSentAt field to given value.

### HasConfirmationSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasConfirmationSentAt() bool`

HasConfirmationSentAt returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetCountry

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetEmail

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFailedAttempts

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### GetFirstName

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroupIds

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetId

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastValidatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastValidatedAt() time.Time`

GetLastValidatedAt returns the LastValidatedAt field if non-nil, zero value otherwise.

### GetLastValidatedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLastValidatedAtOk() (*time.Time, bool)`

GetLastValidatedAtOk returns a tuple with the LastValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetLastValidatedAt(v time.Time)`

SetLastValidatedAt sets LastValidatedAt field to given value.

### HasLastValidatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasLastValidatedAt() bool`

HasLastValidatedAt returns a boolean if a field has been set.

### GetLockedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLockedAt() time.Time`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetLockedAtOk() (*time.Time, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetLockedAt(v time.Time)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetMarketingOptIn

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetMarketingOptIn() bool`

GetMarketingOptIn returns the MarketingOptIn field if non-nil, zero value otherwise.

### GetMarketingOptInOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetMarketingOptInOk() (*bool, bool)`

GetMarketingOptInOk returns a tuple with the MarketingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIn

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetMarketingOptIn(v bool)`

SetMarketingOptIn sets MarketingOptIn field to given value.

### HasMarketingOptIn

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasMarketingOptIn() bool`

HasMarketingOptIn returns a boolean if a field has been set.

### GetName

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRevalidationDueAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetNextRevalidationDueAt() time.Time`

GetNextRevalidationDueAt returns the NextRevalidationDueAt field if non-nil, zero value otherwise.

### GetNextRevalidationDueAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetNextRevalidationDueAtOk() (*time.Time, bool)`

GetNextRevalidationDueAtOk returns a tuple with the NextRevalidationDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRevalidationDueAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetNextRevalidationDueAt(v time.Time)`

SetNextRevalidationDueAt sets NextRevalidationDueAt field to given value.

### HasNextRevalidationDueAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasNextRevalidationDueAt() bool`

HasNextRevalidationDueAt returns a boolean if a field has been set.

### GetOrigin

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPasswordHash

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetProfileValidationRequired

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetProfileValidationRequired() bool`

GetProfileValidationRequired returns the ProfileValidationRequired field if non-nil, zero value otherwise.

### GetProfileValidationRequiredOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetProfileValidationRequiredOk() (*bool, bool)`

GetProfileValidationRequiredOk returns a tuple with the ProfileValidationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileValidationRequired

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetProfileValidationRequired(v bool)`

SetProfileValidationRequired sets ProfileValidationRequired field to given value.

### HasProfileValidationRequired

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasProfileValidationRequired() bool`

HasProfileValidationRequired returns a boolean if a field has been set.

### GetResetPasswordSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetResetPasswordSentAt() time.Time`

GetResetPasswordSentAt returns the ResetPasswordSentAt field if non-nil, zero value otherwise.

### GetResetPasswordSentAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetResetPasswordSentAtOk() (*time.Time, bool)`

GetResetPasswordSentAtOk returns a tuple with the ResetPasswordSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetResetPasswordSentAt(v time.Time)`

SetResetPasswordSentAt sets ResetPasswordSentAt field to given value.

### HasResetPasswordSentAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasResetPasswordSentAt() bool`

HasResetPasswordSentAt returns a boolean if a field has been set.

### GetResetPasswordToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetResetPasswordToken() string`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetResetPasswordTokenOk() (*string, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetResetPasswordToken(v string)`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### GetSignInCount

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetSignInCount() int32`

GetSignInCount returns the SignInCount field if non-nil, zero value otherwise.

### GetSignInCountOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetSignInCountOk() (*int32, bool)`

GetSignInCountOk returns a tuple with the SignInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCount

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetSignInCount(v int32)`

SetSignInCount sets SignInCount field to given value.

### HasSignInCount

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasSignInCount() bool`

HasSignInCount returns a boolean if a field has been set.

### GetState

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStateRegion

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### GetTags

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetType

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnlockToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetUnlockToken() string`

GetUnlockToken returns the UnlockToken field if non-nil, zero value otherwise.

### GetUnlockTokenOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetUnlockTokenOk() (*string, bool)`

GetUnlockTokenOk returns a tuple with the UnlockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetUnlockToken(v string)`

SetUnlockToken sets UnlockToken field to given value.

### HasUnlockToken

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasUnlockToken() bool`

HasUnlockToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserInIdentityV2AccountInIdentityV2Group) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserInIdentityV2AccountInIdentityV2Group) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


