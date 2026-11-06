# LoyaltyPointsAddedDeductedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeName** | **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | **int64** | The ID of the loyalty program. | 
**NotificationType** | **string** | The type of notification. | 
**ProfileIntegrationID** | **string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | **string** | The ID of the subledger within the loyalty program where these points were added. | 
**TypeOfChange** | **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes:  - [The Campaign Manager](/docs/product/getting-started)  - [Management API](/management-api#tag/Loyalty)  - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | **int64** | The ID of the employee who added or deducted points. | 
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

### NewLoyaltyPointsAddedDeductedNotification

`func NewLoyaltyPointsAddedDeductedNotification(employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationID string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, amount float32, operation string, reason string, transactionUUID string, ) *LoyaltyPointsAddedDeductedNotification`

NewLoyaltyPointsAddedDeductedNotification instantiates a new LoyaltyPointsAddedDeductedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyPointsAddedDeductedNotificationWithDefaults

`func NewLoyaltyPointsAddedDeductedNotificationWithDefaults() *LoyaltyPointsAddedDeductedNotification`

NewLoyaltyPointsAddedDeductedNotificationWithDefaults instantiates a new LoyaltyPointsAddedDeductedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeName

`func (o *LoyaltyPointsAddedDeductedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *LoyaltyPointsAddedDeductedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *LoyaltyPointsAddedDeductedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *LoyaltyPointsAddedDeductedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *LoyaltyPointsAddedDeductedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyPointsAddedDeductedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationID

`func (o *LoyaltyPointsAddedDeductedNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *LoyaltyPointsAddedDeductedNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetSessionIntegrationID

`func (o *LoyaltyPointsAddedDeductedNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *LoyaltyPointsAddedDeductedNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *LoyaltyPointsAddedDeductedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *LoyaltyPointsAddedDeductedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *LoyaltyPointsAddedDeductedNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *LoyaltyPointsAddedDeductedNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *LoyaltyPointsAddedDeductedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoyaltyPointsAddedDeductedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetAmount

`func (o *LoyaltyPointsAddedDeductedNotification) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *LoyaltyPointsAddedDeductedNotification) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetExpiryDate

`func (o *LoyaltyPointsAddedDeductedNotification) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *LoyaltyPointsAddedDeductedNotification) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *LoyaltyPointsAddedDeductedNotification) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetOperation

`func (o *LoyaltyPointsAddedDeductedNotification) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *LoyaltyPointsAddedDeductedNotification) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetReason

`func (o *LoyaltyPointsAddedDeductedNotification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *LoyaltyPointsAddedDeductedNotification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetStartDate

`func (o *LoyaltyPointsAddedDeductedNotification) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *LoyaltyPointsAddedDeductedNotification) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *LoyaltyPointsAddedDeductedNotification) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *LoyaltyPointsAddedDeductedNotification) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *LoyaltyPointsAddedDeductedNotification) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCreated

`func (o *LoyaltyPointsAddedDeductedNotification) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoyaltyPointsAddedDeductedNotification) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *LoyaltyPointsAddedDeductedNotification) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetAwaitsActivation

`func (o *LoyaltyPointsAddedDeductedNotification) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *LoyaltyPointsAddedDeductedNotification) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *LoyaltyPointsAddedDeductedNotification) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetValidityDuration

`func (o *LoyaltyPointsAddedDeductedNotification) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *LoyaltyPointsAddedDeductedNotification) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *LoyaltyPointsAddedDeductedNotification) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *LoyaltyPointsAddedDeductedNotification) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


