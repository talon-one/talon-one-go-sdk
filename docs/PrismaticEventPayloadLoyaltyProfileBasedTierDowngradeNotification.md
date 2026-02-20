# PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification

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
**TierExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimestampOfTierChange** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | **time.Time** | Timestamp when the event was published. | 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification

`func NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification(profileIntegrationID string, loyaltyProgramID int64, subledgerID string, sourceOfEvent string, currentPoints float32, publishedAt time.Time, ) *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification instantiates a new PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotificationWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotificationWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification`

NewPrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotificationWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.

### HasCurrentTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) HasCurrentTier() bool`

HasCurrentTier returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventPayloadLoyaltyProfileBasedTierDowngradeNotification) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


