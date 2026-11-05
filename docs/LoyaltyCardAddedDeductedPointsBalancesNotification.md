# LoyaltyCardAddedDeductedPointsBalancesNotification

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
**TypeOfChange** | **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes: - [The Campaign Manager](/docs/product/getting-started) - [Management API](/management-api#tag/Loyalty) - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | **int64** | The ID of the employee who added or deducted points. | 
**UsersPerCardLimit** | **int64** | The max amount of user profiles with whom a card can be shared. This can be set to &#x60;0&#x60; for no limit. | 
**Actions** | [**[]AddedDeductedPointsBalancesAction**](AddedDeductedPointsBalancesAction.md) | The list of actions that have been triggered in the loyalty program. | 
**CurrentPoints** | **float32** | The current points balance. | 

## Methods

### NewLoyaltyCardAddedDeductedPointsBalancesNotification

`func NewLoyaltyCardAddedDeductedPointsBalancesNotification(cardIdentifier string, employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationIDs []string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, usersPerCardLimit int64, actions []AddedDeductedPointsBalancesAction, currentPoints float32, ) *LoyaltyCardAddedDeductedPointsBalancesNotification`

NewLoyaltyCardAddedDeductedPointsBalancesNotification instantiates a new LoyaltyCardAddedDeductedPointsBalancesNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardAddedDeductedPointsBalancesNotificationWithDefaults

`func NewLoyaltyCardAddedDeductedPointsBalancesNotificationWithDefaults() *LoyaltyCardAddedDeductedPointsBalancesNotification`

NewLoyaltyCardAddedDeductedPointsBalancesNotificationWithDefaults instantiates a new LoyaltyCardAddedDeductedPointsBalancesNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardIdentifier

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.


### GetEmployeeName

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationIDs

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetProfileIntegrationIDs() []string`

GetProfileIntegrationIDs returns the ProfileIntegrationIDs field if non-nil, zero value otherwise.

### GetProfileIntegrationIDsOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetProfileIntegrationIDsOk() (*[]string, bool)`

GetProfileIntegrationIDsOk returns a tuple with the ProfileIntegrationIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationIDs

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetProfileIntegrationIDs(v []string)`

SetProfileIntegrationIDs sets ProfileIntegrationIDs field to given value.


### GetSessionIntegrationID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetUsersPerCardLimit

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetUsersPerCardLimit() int64`

GetUsersPerCardLimit returns the UsersPerCardLimit field if non-nil, zero value otherwise.

### GetUsersPerCardLimitOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetUsersPerCardLimitOk() (*int64, bool)`

GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersPerCardLimit

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetUsersPerCardLimit(v int64)`

SetUsersPerCardLimit sets UsersPerCardLimit field to given value.


### GetActions

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetActions() []AddedDeductedPointsBalancesAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetActionsOk() (*[]AddedDeductedPointsBalancesAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetActions(v []AddedDeductedPointsBalancesAction)`

SetActions sets Actions field to given value.


### GetCurrentPoints

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *LoyaltyCardAddedDeductedPointsBalancesNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


