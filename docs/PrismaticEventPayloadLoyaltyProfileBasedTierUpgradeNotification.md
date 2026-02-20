# PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIntegrationID** | **string** |  | 
**LoyaltyProgramID** | **int64** |  | 
**SubledgerID** | **string** |  | 
**SourceOfEvent** | **string** |  | 
**CurrentTier** | Pointer to **string** |  | [optional] 
**CurrentPoints** | **float32** |  | 
**OldTier** | Pointer to **string** |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimestampOfTierChange** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | **time.Time** | Timestamp when the event was published. | 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification

`func NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification instantiates a new PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotificationWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierUpgradeNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


