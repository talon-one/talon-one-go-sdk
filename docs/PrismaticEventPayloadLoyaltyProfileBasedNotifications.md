# PrismaticEventPayloadLoyaltyProfileBasedNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIntegrationID** | **string** |  | 
**LoyaltyProgramID** | **int64** |  | 
**SubledgerID** | **string** |  | 
**SourceOfEvent** | **string** |  | 
**EmployeeName** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **int64** |  | [optional] 
**SessionIntegrationID** | Pointer to **string** |  | [optional] 
**CurrentTier** | Pointer to **string** |  | [optional] 
**CurrentPoints** | **float32** |  | 
**OldTier** | Pointer to **string** |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimestampOfTierChange** | Pointer to **time.Time** |  | [optional] 
**Actions** | Pointer to [**[]PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction**](PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction.md) |  | [optional] 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedNotifications

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotifications(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, ) *PrismaticEventPayloadLoyaltyProfileBasedNotifications`

NewPrismaticEventPayloadLoyaltyProfileBasedNotifications instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedNotifications`

NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.

### HasEmployeeName

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasEmployeeName() bool`

HasEmployeeName returns a boolean if a field has been set.

### GetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetSessionIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.

### HasSessionIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasSessionIntegrationID() bool`

HasSessionIntegrationID returns a boolean if a field has been set.

### GetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetActions() []PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) GetActionsOk() (*[]PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) SetActions(v []PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotifications) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


