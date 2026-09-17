# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AccountInIdentityV2User**](AccountInIdentityV2User.md) |  | [optional] 
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
**Groups** | Pointer to [**[]GroupInIdentityV2User**](GroupInIdentityV2User.md) |  | [optional] 
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

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *User) GetAccount() AccountInIdentityV2User`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *User) GetAccountOk() (*AccountInIdentityV2User, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *User) SetAccount(v AccountInIdentityV2User)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *User) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountCspId

`func (o *User) GetAccountCspId() int32`

GetAccountCspId returns the AccountCspId field if non-nil, zero value otherwise.

### GetAccountCspIdOk

`func (o *User) GetAccountCspIdOk() (*int32, bool)`

GetAccountCspIdOk returns a tuple with the AccountCspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCspId

`func (o *User) SetAccountCspId(v int32)`

SetAccountCspId sets AccountCspId field to given value.

### HasAccountCspId

`func (o *User) HasAccountCspId() bool`

HasAccountCspId returns a boolean if a field has been set.

### GetAccountId

`func (o *User) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *User) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountInfobloxId

`func (o *User) GetAccountInfobloxId() string`

GetAccountInfobloxId returns the AccountInfobloxId field if non-nil, zero value otherwise.

### GetAccountInfobloxIdOk

`func (o *User) GetAccountInfobloxIdOk() (*string, bool)`

GetAccountInfobloxIdOk returns a tuple with the AccountInfobloxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfobloxId

`func (o *User) SetAccountInfobloxId(v string)`

SetAccountInfobloxId sets AccountInfobloxId field to given value.

### HasAccountInfobloxId

`func (o *User) HasAccountInfobloxId() bool`

HasAccountInfobloxId returns a boolean if a field has been set.

### GetAgreementsAccepted

`func (o *User) GetAgreementsAccepted() bool`

GetAgreementsAccepted returns the AgreementsAccepted field if non-nil, zero value otherwise.

### GetAgreementsAcceptedOk

`func (o *User) GetAgreementsAcceptedOk() (*bool, bool)`

GetAgreementsAcceptedOk returns a tuple with the AgreementsAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementsAccepted

`func (o *User) SetAgreementsAccepted(v bool)`

SetAgreementsAccepted sets AgreementsAccepted field to given value.

### HasAgreementsAccepted

`func (o *User) HasAgreementsAccepted() bool`

HasAgreementsAccepted returns a boolean if a field has been set.

### GetApiKey

`func (o *User) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *User) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *User) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *User) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApiKeys

`func (o *User) GetApiKeys() []string`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### GetApiKeysOk

`func (o *User) GetApiKeysOk() (*[]string, bool)`

GetApiKeysOk returns a tuple with the ApiKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeys

`func (o *User) SetApiKeys(v []string)`

SetApiKeys sets ApiKeys field to given value.

### HasApiKeys

`func (o *User) HasApiKeys() bool`

HasApiKeys returns a boolean if a field has been set.

### GetApproved

`func (o *User) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *User) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *User) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *User) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetAuthenticator

`func (o *User) GetAuthenticator() string`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *User) GetAuthenticatorOk() (*string, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *User) SetAuthenticator(v string)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *User) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetConfirmToken

`func (o *User) GetConfirmToken() string`

GetConfirmToken returns the ConfirmToken field if non-nil, zero value otherwise.

### GetConfirmTokenOk

`func (o *User) GetConfirmTokenOk() (*string, bool)`

GetConfirmTokenOk returns a tuple with the ConfirmToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmToken

`func (o *User) SetConfirmToken(v string)`

SetConfirmToken sets ConfirmToken field to given value.

### HasConfirmToken

`func (o *User) HasConfirmToken() bool`

HasConfirmToken returns a boolean if a field has been set.

### GetConfirmationSentAt

`func (o *User) GetConfirmationSentAt() time.Time`

GetConfirmationSentAt returns the ConfirmationSentAt field if non-nil, zero value otherwise.

### GetConfirmationSentAtOk

`func (o *User) GetConfirmationSentAtOk() (*time.Time, bool)`

GetConfirmationSentAtOk returns a tuple with the ConfirmationSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationSentAt

`func (o *User) SetConfirmationSentAt(v time.Time)`

SetConfirmationSentAt sets ConfirmationSentAt field to given value.

### HasConfirmationSentAt

`func (o *User) HasConfirmationSentAt() bool`

HasConfirmationSentAt returns a boolean if a field has been set.

### GetConfirmedAt

`func (o *User) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *User) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *User) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *User) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### GetCountry

`func (o *User) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *User) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *User) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *User) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCspId

`func (o *User) GetCspId() int32`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *User) GetCspIdOk() (*int32, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *User) SetCspId(v int32)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *User) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetDeletedAt

`func (o *User) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *User) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *User) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *User) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *User) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *User) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *User) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *User) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFailedAttempts

`func (o *User) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *User) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *User) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.

### HasFailedAttempts

`func (o *User) HasFailedAttempts() bool`

HasFailedAttempts returns a boolean if a field has been set.

### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetGroupIds

`func (o *User) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *User) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *User) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *User) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetGroups

`func (o *User) GetGroups() []GroupInIdentityV2User`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *User) GetGroupsOk() (*[]GroupInIdentityV2User, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *User) SetGroups(v []GroupInIdentityV2User)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *User) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobTitle

`func (o *User) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *User) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *User) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *User) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLastValidatedAt

`func (o *User) GetLastValidatedAt() time.Time`

GetLastValidatedAt returns the LastValidatedAt field if non-nil, zero value otherwise.

### GetLastValidatedAtOk

`func (o *User) GetLastValidatedAtOk() (*time.Time, bool)`

GetLastValidatedAtOk returns a tuple with the LastValidatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidatedAt

`func (o *User) SetLastValidatedAt(v time.Time)`

SetLastValidatedAt sets LastValidatedAt field to given value.

### HasLastValidatedAt

`func (o *User) HasLastValidatedAt() bool`

HasLastValidatedAt returns a boolean if a field has been set.

### GetLockedAt

`func (o *User) GetLockedAt() time.Time`

GetLockedAt returns the LockedAt field if non-nil, zero value otherwise.

### GetLockedAtOk

`func (o *User) GetLockedAtOk() (*time.Time, bool)`

GetLockedAtOk returns a tuple with the LockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedAt

`func (o *User) SetLockedAt(v time.Time)`

SetLockedAt sets LockedAt field to given value.

### HasLockedAt

`func (o *User) HasLockedAt() bool`

HasLockedAt returns a boolean if a field has been set.

### GetMarketingOptIn

`func (o *User) GetMarketingOptIn() bool`

GetMarketingOptIn returns the MarketingOptIn field if non-nil, zero value otherwise.

### GetMarketingOptInOk

`func (o *User) GetMarketingOptInOk() (*bool, bool)`

GetMarketingOptInOk returns a tuple with the MarketingOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingOptIn

`func (o *User) SetMarketingOptIn(v bool)`

SetMarketingOptIn sets MarketingOptIn field to given value.

### HasMarketingOptIn

`func (o *User) HasMarketingOptIn() bool`

HasMarketingOptIn returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRevalidationDueAt

`func (o *User) GetNextRevalidationDueAt() time.Time`

GetNextRevalidationDueAt returns the NextRevalidationDueAt field if non-nil, zero value otherwise.

### GetNextRevalidationDueAtOk

`func (o *User) GetNextRevalidationDueAtOk() (*time.Time, bool)`

GetNextRevalidationDueAtOk returns a tuple with the NextRevalidationDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRevalidationDueAt

`func (o *User) SetNextRevalidationDueAt(v time.Time)`

SetNextRevalidationDueAt sets NextRevalidationDueAt field to given value.

### HasNextRevalidationDueAt

`func (o *User) HasNextRevalidationDueAt() bool`

HasNextRevalidationDueAt returns a boolean if a field has been set.

### GetOrigin

`func (o *User) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *User) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *User) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *User) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPasswordHash

`func (o *User) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *User) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *User) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *User) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *User) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetProfileValidationRequired

`func (o *User) GetProfileValidationRequired() bool`

GetProfileValidationRequired returns the ProfileValidationRequired field if non-nil, zero value otherwise.

### GetProfileValidationRequiredOk

`func (o *User) GetProfileValidationRequiredOk() (*bool, bool)`

GetProfileValidationRequiredOk returns a tuple with the ProfileValidationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileValidationRequired

`func (o *User) SetProfileValidationRequired(v bool)`

SetProfileValidationRequired sets ProfileValidationRequired field to given value.

### HasProfileValidationRequired

`func (o *User) HasProfileValidationRequired() bool`

HasProfileValidationRequired returns a boolean if a field has been set.

### GetResetPasswordSentAt

`func (o *User) GetResetPasswordSentAt() time.Time`

GetResetPasswordSentAt returns the ResetPasswordSentAt field if non-nil, zero value otherwise.

### GetResetPasswordSentAtOk

`func (o *User) GetResetPasswordSentAtOk() (*time.Time, bool)`

GetResetPasswordSentAtOk returns a tuple with the ResetPasswordSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordSentAt

`func (o *User) SetResetPasswordSentAt(v time.Time)`

SetResetPasswordSentAt sets ResetPasswordSentAt field to given value.

### HasResetPasswordSentAt

`func (o *User) HasResetPasswordSentAt() bool`

HasResetPasswordSentAt returns a boolean if a field has been set.

### GetResetPasswordToken

`func (o *User) GetResetPasswordToken() string`

GetResetPasswordToken returns the ResetPasswordToken field if non-nil, zero value otherwise.

### GetResetPasswordTokenOk

`func (o *User) GetResetPasswordTokenOk() (*string, bool)`

GetResetPasswordTokenOk returns a tuple with the ResetPasswordToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPasswordToken

`func (o *User) SetResetPasswordToken(v string)`

SetResetPasswordToken sets ResetPasswordToken field to given value.

### HasResetPasswordToken

`func (o *User) HasResetPasswordToken() bool`

HasResetPasswordToken returns a boolean if a field has been set.

### GetSignInCount

`func (o *User) GetSignInCount() int32`

GetSignInCount returns the SignInCount field if non-nil, zero value otherwise.

### GetSignInCountOk

`func (o *User) GetSignInCountOk() (*int32, bool)`

GetSignInCountOk returns a tuple with the SignInCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInCount

`func (o *User) SetSignInCount(v int32)`

SetSignInCount sets SignInCount field to given value.

### HasSignInCount

`func (o *User) HasSignInCount() bool`

HasSignInCount returns a boolean if a field has been set.

### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateChangedAt

`func (o *User) GetStateChangedAt() time.Time`

GetStateChangedAt returns the StateChangedAt field if non-nil, zero value otherwise.

### GetStateChangedAtOk

`func (o *User) GetStateChangedAtOk() (*time.Time, bool)`

GetStateChangedAtOk returns a tuple with the StateChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChangedAt

`func (o *User) SetStateChangedAt(v time.Time)`

SetStateChangedAt sets StateChangedAt field to given value.

### HasStateChangedAt

`func (o *User) HasStateChangedAt() bool`

HasStateChangedAt returns a boolean if a field has been set.

### GetStateRegion

`func (o *User) GetStateRegion() string`

GetStateRegion returns the StateRegion field if non-nil, zero value otherwise.

### GetStateRegionOk

`func (o *User) GetStateRegionOk() (*string, bool)`

GetStateRegionOk returns a tuple with the StateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRegion

`func (o *User) SetStateRegion(v string)`

SetStateRegion sets StateRegion field to given value.

### HasStateRegion

`func (o *User) HasStateRegion() bool`

HasStateRegion returns a boolean if a field has been set.

### GetTags

`func (o *User) GetTags() map[string]map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *User) GetTagsOk() (*map[string]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *User) SetTags(v map[string]map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *User) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimezone

`func (o *User) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *User) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *User) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *User) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetType

`func (o *User) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *User) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *User) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *User) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnlockToken

`func (o *User) GetUnlockToken() string`

GetUnlockToken returns the UnlockToken field if non-nil, zero value otherwise.

### GetUnlockTokenOk

`func (o *User) GetUnlockTokenOk() (*string, bool)`

GetUnlockTokenOk returns a tuple with the UnlockToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockToken

`func (o *User) SetUnlockToken(v string)`

SetUnlockToken sets UnlockToken field to given value.

### HasUnlockToken

`func (o *User) HasUnlockToken() bool`

HasUnlockToken returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


