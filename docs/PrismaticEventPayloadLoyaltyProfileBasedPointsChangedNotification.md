# PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIntegrationID** | **string** |  | 
**LoyaltyProgramID** | **int64** |  | 
**SubledgerID** | **string** |  | 
**SourceOfEvent** | **string** |  | 
**EmployeeName** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **int64** |  | [optional] 
**CurrentPoints** | **float32** |  | 
**Actions** | Pointer to [**[]PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction**](PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction.md) |  | [optional] 
**PublishedAt** | **time.Time** | Timestamp when the event was published. | 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification

`func NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification instantiates a new PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.

### HasEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasEmployeeName() bool`

HasEmployeeName returns a boolean if a field has been set.

### GetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetActions() []PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetActionsOk() (*[]PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetActions(v []PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


