# LoyaltyCardPointsAddedDeductedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardIdentifier** | **string** | Loyalty card identification number. | 
**EmployeeName** | **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | **int64** | The ID of the loyalty program. | 
**NotificationType** | **string** | The type of notification. | 
**ProfileIntegrationIDs** | **[]string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | **string** | The ID of the subledger within the loyalty program where these points were added or deducted. | 
**TypeOfChange** | **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes:  - [The Campaign Manager](/docs/product/getting-started)  - [Management API](/management-api#tag/Loyalty)  - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | **int64** | The ID of the employee who added or deducted points. | 
**UsersPerCardLimit** | **int64** | The max amount of user profiles with whom a card can be shared. This can be set to &#x60;0&#x60; for no limit. | 
**Amount** | **float32** | The amount of added or deducted loyalty points. | 
**ExpiryDate** | Pointer to **time.Time** | The expiration date for loyalty points. | [optional] 
**Operation** | **string** | The action (addition or subtraction) made with loyalty points. | 
**Reason** | **string** | The reason for the points addition or deduction. | 
**StartDate** | Pointer to **time.Time** | The start date for loyalty points. | [optional] 
**TransactionUUID** | **string** | The identifier of the transaction in the loyalty ledger. | 
**Created** | Pointer to **time.Time** | The timestamp when the loyalty transaction occurred. | [optional] 
**AwaitsActivation** | Pointer to **bool** | Indicates whether the points have an action-based start date. This property is returned only for point transactions with an action-based start date. | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, relative to the activation date.  **Note**: This only applies to points for which &#x60;AwaitsActivation&#x60; is &#x60;true&#x60; and &#x60;ExpiryDate&#x60; is not set.  | [optional] 

## Methods

### NewLoyaltyCardPointsAddedDeductedNotification

`func NewLoyaltyCardPointsAddedDeductedNotification(cardIdentifier string, employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationIDs []string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, usersPerCardLimit int64, amount float32, operation string, reason string, transactionUUID string, ) *LoyaltyCardPointsAddedDeductedNotification`

NewLoyaltyCardPointsAddedDeductedNotification instantiates a new LoyaltyCardPointsAddedDeductedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardPointsAddedDeductedNotificationWithDefaults

`func NewLoyaltyCardPointsAddedDeductedNotificationWithDefaults() *LoyaltyCardPointsAddedDeductedNotification`

NewLoyaltyCardPointsAddedDeductedNotificationWithDefaults instantiates a new LoyaltyCardPointsAddedDeductedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIdentifier

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetEmployeeName

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationIDs

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetProfileIntegrationIDs() []string`

GetProfileIntegrationIDs returns the ProfileIntegrationIDs field if non-nil, zero value otherwise.

### GetProfileIntegrationIDsOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetProfileIntegrationIDsOk() (*[]string, bool)`

GetProfileIntegrationIDsOk returns a tuple with the ProfileIntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationIDs

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetProfileIntegrationIDs(v []string)`

SetProfileIntegrationIDs sets ProfileIntegrationIDs field to given value.


### GetSessionIntegrationID

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetUsersPerCardLimit

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetAmount

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetExpiryDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetOperation

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LoyaltyCardPointsAddedDeductedNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LoyaltyCardPointsAddedDeductedNotification) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetAwaitsActivation

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *LoyaltyCardPointsAddedDeductedNotification) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetValidityDuration

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LoyaltyCardPointsAddedDeductedNotification) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *LoyaltyCardPointsAddedDeductedNotification) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *LoyaltyCardPointsAddedDeductedNotification) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


