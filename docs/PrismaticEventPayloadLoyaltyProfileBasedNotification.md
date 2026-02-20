# PrismaticEventPayloadLoyaltyProfileBasedNotification

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
**CurrentTier** | Pointer to **string** |  | [optional] 
**OldTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimestampOfTierChange** | Pointer to **time.Time** |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedNotification

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *PrismaticEventPayloadLoyaltyProfileBasedNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedNotification instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedNotificationWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotificationWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedNotificationWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.

### HasEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasEmployeeName() bool`

HasEmployeeName returns a boolean if a field has been set.

### GetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetActions() []PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetActionsOk() (*[]PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetActions(v []PrismaticEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotification) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


