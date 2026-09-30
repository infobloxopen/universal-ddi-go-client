# UserInIdentityV2Account

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
**Groups** | Pointer to [**[]GroupInIdentityV2UserInIdentityV2Account**](GroupInIdentityV2UserInIdentityV2Account.md) |  | [optional] 
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

### NewUserInIdentityV2Account

`func NewUserInIdentityV2Account() *UserInIdentityV2Account`

NewUserInIdentityV2Account instantiates a new UserInIdentityV2Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInIdentityV2AccountWithDefaults

`func NewUserInIdentityV2AccountWithDefaults() *UserInIdentityV2Account`

NewUserInIdentityV2AccountWithDefaults instantiates a new UserInIdentityV2Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountCspId

`func (o *UserInIdentityV2Account) GetAccountCspId() int32`

GetAccountCspId returns the AccountCspId field if non-nil, zero value otherwise.

### GetAccountCspIdOk

`func (o *UserInIdentityV2Account) GetAccountCspIdOk() (*int32, bool)`

GetAccountCspIdOk returns a tuple with the AccountCspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCspId

`func (o *UserInIdentityV2Account) SetAccountCspId(v int32)`

SetAccountCspId sets AccountCspId field to given value.

### HasAccountCspId

`func (o *UserInIdentityV2Account) HasAccountCspId() bool`

HasAccountCspId returns a boolean if a field has been set.

### GetAccountId

`func (o *UserInIdentityV2Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserInIdentityV2Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserInIdentityV2Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserInIdentityV2Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountInfobloxId

`func (o *UserInIdentityV2Account) GetAccountInfobloxId() string`

GetAccountInfobloxId returns the AccountInfobloxId field if non-nil, zero value otherwise.

### GetAccountInfobloxIdOk

`func (o *UserInIdentityV2Account) GetAccountInfobloxIdOk() (*string, bool)`

GetAccountInfobloxIdOk returns a tuple with the AccountInfobloxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfobloxId

`func (o *UserInIdentityV2Account) SetAccountInfobloxId(v string)`

SetAccountInfobloxId sets AccountInfobloxId field to given value.

### HasAccountInfobloxId

`func (o *UserInIdentityV2Account) HasAccountInfobloxId() bool`

HasAccountInfobloxId returns a boolean if a field has been set.

### GetAgreementsAccepted

`func (o *UserInIdentityV2Account) GetAgreementsAccepted() bool`

GetAgreementsAccepted returns the AgreementsAccepted field if non-nil, zero value otherwise.

### GetAgreementsAcceptedOk

`func (o *UserInIdentityV2Account) GetAgreementsAcceptedOk() (*bool, bool)`

GetAgreementsAcceptedOk returns a tuple with the AgreementsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementsAccepted

`func (o *UserInIdentityV2Account) SetAgreementsAccepted(v bool)`

SetAgreementsAccepted sets AgreementsAccepted field to given value.

### HasAgreementsAccepted

`func (o *UserInIdentityV2Account) HasAgreementsAccepted() bool`

HasAgreementsAccepted returns a boolean if a field has been set.

### GetApiKey

`func (o *UserInIdentityV2Account) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserInIdentityV2Account) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserInIdentityV2Account) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserInIdentityV2Account) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeys

`func (o *UserInIdentityV2Account) GetApiKeys() []string`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *UserInIdentityV2Account) GetApiKeysOk() (*[]string, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *UserInIdentityV2Account) SetApiKeys(v []string)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *UserInIdentityV2Account) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetApproved

`func (o *UserInIdentityV2Account) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *UserInIdentityV2Account) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *UserInIdentityV2Account) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *UserInIdentityV2Account) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetAuthenticator

`func (o *UserInIdentityV2Account) GetAuthenticator() string`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *UserInIdentityV2Account) GetAuthenticatorOk() (*string, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *UserInIdentityV2Account) SetAuthenticator(v string)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *UserInIdentityV2Account) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetConfirmToken

`func (o *UserInIdentityV2Account) GetConfirmToken() string`

GetConfirmToken returns the ConfirmToken field if non-nil, zero value otherwise.

### GetConfirmTokenOk

`func (o *UserInIdentityV2Account) GetConfirmTokenOk() (*string, bool)`

GetConfirmTokenOk returns a tuple with the ConfirmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmToken

`func (o *UserInIdentityV2Account) SetConfirmToken(v string)`

SetConfirmToken sets ConfirmToken field to given value.

### HasConfirmToken

`func (o *UserInIdentityV2Account) HasConfirmToken() bool`

HasConfirmToken returns a boolean if a field has been set.

### GetConfirmationSentAt

`func (o *UserInIdentityV2Account) GetConfirmationSentAt() time.Time`

GetConfirmationSentAt returns the ConfirmationSentAt field if non-nil, zero value otherwise.

### GetConfirmationSentAtOk

`func (o *UserInIdentityV2Account) GetConfirmationSentAtOk() (*time.Time, bool)`

GetConfirmationSentAtOk returns a tuple with the ConfirmationSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSentAt

`func (o *UserInIdentityV2Account) SetConfirmationSentAt(v time.Time)`

SetConfirmationSentAt sets ConfirmationSentAt field to given value.

### HasConfirmationSentAt

`func (o *UserInIdentityV2Account) HasConfirmationSentAt() bool`

HasConfirmationSentAt returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *UserInIdentityV2Account) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *UserInIdentityV2Account) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *UserInIdentityV2Account) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *UserInIdentityV2Account) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetCountry

`func (o *UserInIdentityV2Account) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserInIdentityV2Account) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserInIdentityV2Account) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserInIdentityV2Account) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserInIdentityV2Account) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserInIdentityV2Account) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserInIdentityV2Account) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserInIdentityV2Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCspId

`func (o *UserInIdentityV2Account) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *UserInIdentityV2Account) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *UserInIdentityV2Account) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *UserInIdentityV2Account) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *UserInIdentityV2Account) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UserInIdentityV2Account) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UserInIdentityV2Account) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UserInIdentityV2Account) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *UserInIdentityV2Account) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *UserInIdentityV2Account) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *UserInIdentityV2Account) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *UserInIdentityV2Account) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetEmail

`func (o *UserInIdentityV2Account) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserInIdentityV2Account) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserInIdentityV2Account) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserInIdentityV2Account) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFailedAttempts

`func (o *UserInIdentityV2Account) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *UserInIdentityV2Account) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *UserInIdentityV2Account) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *UserInIdentityV2Account) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### GetFirstName

`func (o *UserInIdentityV2Account) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserInIdentityV2Account) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserInIdentityV2Account) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserInIdentityV2Account) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroupIds

`func (o *UserInIdentityV2Account) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *UserInIdentityV2Account) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *UserInIdentityV2Account) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *UserInIdentityV2Account) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetGroups

`func (o *UserInIdentityV2Account) GetGroups() []GroupInIdentityV2UserInIdentityV2Account`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserInIdentityV2Account) GetGroupsOk() (*[]GroupInIdentityV2UserInIdentityV2Account, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserInIdentityV2Account) SetGroups(v []GroupInIdentityV2UserInIdentityV2Account)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserInIdentityV2Account) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *UserInIdentityV2Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserInIdentityV2Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserInIdentityV2Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserInIdentityV2Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *UserInIdentityV2Account) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *UserInIdentityV2Account) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *UserInIdentityV2Account) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *UserInIdentityV2Account) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastLogin

`func (o *UserInIdentityV2Account) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *UserInIdentityV2Account) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *UserInIdentityV2Account) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *UserInIdentityV2Account) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *UserInIdentityV2Account) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserInIdentityV2Account) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserInIdentityV2Account) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserInIdentityV2Account) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastValidatedAt

`func (o *UserInIdentityV2Account) GetLastValidatedAt() time.Time`

GetLastValidatedAt returns the LastValidatedAt field if non-nil, zero value otherwise.

### GetLastValidatedAtOk

`func (o *UserInIdentityV2Account) GetLastValidatedAtOk() (*time.Time, bool)`

GetLastValidatedAtOk returns a tuple with the LastValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidatedAt

`func (o *UserInIdentityV2Account) SetLastValidatedAt(v time.Time)`

SetLastValidatedAt sets LastValidatedAt field to given value.

### HasLastValidatedAt

`func (o *UserInIdentityV2Account) HasLastValidatedAt() bool`

HasLastValidatedAt returns a boolean if a field has been set.

### GetLockedAt

`func (o *UserInIdentityV2Account) GetLockedAt() time.Time`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *UserInIdentityV2Account) GetLockedAtOk() (*time.Time, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *UserInIdentityV2Account) SetLockedAt(v time.Time)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *UserInIdentityV2Account) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetMarketingOptIn

`func (o *UserInIdentityV2Account) GetMarketingOptIn() bool`

GetMarketingOptIn returns the MarketingOptIn field if non-nil, zero value otherwise.

### GetMarketingOptInOk

`func (o *UserInIdentityV2Account) GetMarketingOptInOk() (*bool, bool)`

GetMarketingOptInOk returns a tuple with the MarketingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIn

`func (o *UserInIdentityV2Account) SetMarketingOptIn(v bool)`

SetMarketingOptIn sets MarketingOptIn field to given value.

### HasMarketingOptIn

`func (o *UserInIdentityV2Account) HasMarketingOptIn() bool`

HasMarketingOptIn returns a boolean if a field has been set.

### GetName

`func (o *UserInIdentityV2Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserInIdentityV2Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserInIdentityV2Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserInIdentityV2Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRevalidationDueAt

`func (o *UserInIdentityV2Account) GetNextRevalidationDueAt() time.Time`

GetNextRevalidationDueAt returns the NextRevalidationDueAt field if non-nil, zero value otherwise.

### GetNextRevalidationDueAtOk

`func (o *UserInIdentityV2Account) GetNextRevalidationDueAtOk() (*time.Time, bool)`

GetNextRevalidationDueAtOk returns a tuple with the NextRevalidationDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRevalidationDueAt

`func (o *UserInIdentityV2Account) SetNextRevalidationDueAt(v time.Time)`

SetNextRevalidationDueAt sets NextRevalidationDueAt field to given value.

### HasNextRevalidationDueAt

`func (o *UserInIdentityV2Account) HasNextRevalidationDueAt() bool`

HasNextRevalidationDueAt returns a boolean if a field has been set.

### GetOrigin

`func (o *UserInIdentityV2Account) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *UserInIdentityV2Account) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *UserInIdentityV2Account) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *UserInIdentityV2Account) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPasswordHash

`func (o *UserInIdentityV2Account) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *UserInIdentityV2Account) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *UserInIdentityV2Account) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *UserInIdentityV2Account) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *UserInIdentityV2Account) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserInIdentityV2Account) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserInIdentityV2Account) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserInIdentityV2Account) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetProfileValidationRequired

`func (o *UserInIdentityV2Account) GetProfileValidationRequired() bool`

GetProfileValidationRequired returns the ProfileValidationRequired field if non-nil, zero value otherwise.

### GetProfileValidationRequiredOk

`func (o *UserInIdentityV2Account) GetProfileValidationRequiredOk() (*bool, bool)`

GetProfileValidationRequiredOk returns a tuple with the ProfileValidationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileValidationRequired

`func (o *UserInIdentityV2Account) SetProfileValidationRequired(v bool)`

SetProfileValidationRequired sets ProfileValidationRequired field to given value.

### HasProfileValidationRequired

`func (o *UserInIdentityV2Account) HasProfileValidationRequired() bool`

HasProfileValidationRequired returns a boolean if a field has been set.

### GetResetPasswordSentAt

`func (o *UserInIdentityV2Account) GetResetPasswordSentAt() time.Time`

GetResetPasswordSentAt returns the ResetPasswordSentAt field if non-nil, zero value otherwise.

### GetResetPasswordSentAtOk

`func (o *UserInIdentityV2Account) GetResetPasswordSentAtOk() (*time.Time, bool)`

GetResetPasswordSentAtOk returns a tuple with the ResetPasswordSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordSentAt

`func (o *UserInIdentityV2Account) SetResetPasswordSentAt(v time.Time)`

SetResetPasswordSentAt sets ResetPasswordSentAt field to given value.

### HasResetPasswordSentAt

`func (o *UserInIdentityV2Account) HasResetPasswordSentAt() bool`

HasResetPasswordSentAt returns a boolean if a field has been set.

### GetResetPasswordToken

`func (o *UserInIdentityV2Account) GetResetPasswordToken() string`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *UserInIdentityV2Account) GetResetPasswordTokenOk() (*string, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *UserInIdentityV2Account) SetResetPasswordToken(v string)`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *UserInIdentityV2Account) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### GetSignInCount

`func (o *UserInIdentityV2Account) GetSignInCount() int32`

GetSignInCount returns the SignInCount field if non-nil, zero value otherwise.

### GetSignInCountOk

`func (o *UserInIdentityV2Account) GetSignInCountOk() (*int32, bool)`

GetSignInCountOk returns a tuple with the SignInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCount

`func (o *UserInIdentityV2Account) SetSignInCount(v int32)`

SetSignInCount sets SignInCount field to given value.

### HasSignInCount

`func (o *UserInIdentityV2Account) HasSignInCount() bool`

HasSignInCount returns a boolean if a field has been set.

### GetState

`func (o *UserInIdentityV2Account) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserInIdentityV2Account) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserInIdentityV2Account) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UserInIdentityV2Account) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *UserInIdentityV2Account) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *UserInIdentityV2Account) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *UserInIdentityV2Account) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *UserInIdentityV2Account) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStateRegion

`func (o *UserInIdentityV2Account) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *UserInIdentityV2Account) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *UserInIdentityV2Account) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *UserInIdentityV2Account) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### GetTags

`func (o *UserInIdentityV2Account) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UserInIdentityV2Account) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UserInIdentityV2Account) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UserInIdentityV2Account) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *UserInIdentityV2Account) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserInIdentityV2Account) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserInIdentityV2Account) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserInIdentityV2Account) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetType

`func (o *UserInIdentityV2Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserInIdentityV2Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserInIdentityV2Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserInIdentityV2Account) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnlockToken

`func (o *UserInIdentityV2Account) GetUnlockToken() string`

GetUnlockToken returns the UnlockToken field if non-nil, zero value otherwise.

### GetUnlockTokenOk

`func (o *UserInIdentityV2Account) GetUnlockTokenOk() (*string, bool)`

GetUnlockTokenOk returns a tuple with the UnlockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockToken

`func (o *UserInIdentityV2Account) SetUnlockToken(v string)`

SetUnlockToken sets UnlockToken field to given value.

### HasUnlockToken

`func (o *UserInIdentityV2Account) HasUnlockToken() bool`

HasUnlockToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserInIdentityV2Account) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserInIdentityV2Account) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserInIdentityV2Account) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserInIdentityV2Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


