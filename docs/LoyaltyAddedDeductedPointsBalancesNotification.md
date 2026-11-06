# LoyaltyAddedDeductedPointsBalancesNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployeeName** | **string** | The name of the employee who added or deducted points. | 
**LoyaltyProgramID** | **int64** | The ID of the loyalty program. | 
**NotificationType** | **string** | The type of notification. | 
**ProfileIntegrationID** | **string** | The integration ID of the customer profile to whom points were added or deducted. | 
**SessionIntegrationID** | **string** | The integration ID of the session through which the points were earned or lost. | 
**SubledgerID** | **string** | The ID of the subledger within the loyalty program where these points were added. | 
**TypeOfChange** | **string** | The notification source, that is, it indicates whether the points were added or deducted via one of the following routes: - [The Campaign Manager](/docs/product/getting-started) - [Management API](/management-api#tag/Loyalty) - [Rule Engine](/docs/product/applications/evaluation-order-for-rules-and-filters)  | 
**UserID** | **int64** | The ID of the employee who added or deducted points. | 
**Actions** | [**[]AddedDeductedPointsBalancesAction**](AddedDeductedPointsBalancesAction.md) | The list of actions that have been triggered in the loyalty program. | 
**CurrentPoints** | **float32** | The current points balance. | 

## Methods

### NewLoyaltyAddedDeductedPointsBalancesNotification

`func NewLoyaltyAddedDeductedPointsBalancesNotification(employeeName string, loyaltyProgramID int64, notificationType string, profileIntegrationID string, sessionIntegrationID string, subledgerID string, typeOfChange string, userID int64, actions []AddedDeductedPointsBalancesAction, currentPoints float32, ) *LoyaltyAddedDeductedPointsBalancesNotification`

NewLoyaltyAddedDeductedPointsBalancesNotification instantiates a new LoyaltyAddedDeductedPointsBalancesNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyAddedDeductedPointsBalancesNotificationWithDefaults

`func NewLoyaltyAddedDeductedPointsBalancesNotificationWithDefaults() *LoyaltyAddedDeductedPointsBalancesNotification`

NewLoyaltyAddedDeductedPointsBalancesNotificationWithDefaults instantiates a new LoyaltyAddedDeductedPointsBalancesNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployeeName

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetLoyaltyProgramID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetNotificationType

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetProfileIntegrationID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetSessionIntegrationID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.


### GetSubledgerID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetTypeOfChange

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetUserID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.


### GetActions

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetActions() []AddedDeductedPointsBalancesAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetActionsOk() (*[]AddedDeductedPointsBalancesAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetActions(v []AddedDeductedPointsBalancesAction)`

SetActions sets Actions field to given value.


### GetCurrentPoints

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *LoyaltyAddedDeductedPointsBalancesNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


