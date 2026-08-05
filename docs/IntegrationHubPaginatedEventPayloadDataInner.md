# IntegrationHubPaginatedEventPayloadDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **int64** | The ID of the integration hub event. Return this value in the delivery-status callback to mark the event delivered or failed. | 
**ProfileIntegrationID** | **string** |  | 
**LoyaltyProgramID** | **int64** |  | 
**LoyaltyProgramName** | **string** | The name of the loyalty program. | 
**SubledgerID** | **string** |  | 
**SourceOfEvent** | **string** |  | 
**CurrentTier** | **string** | The name of the customer&#39;s current tier. | 
**SessionIntegrationID** | Pointer to **string** | The integration ID of the session through which the points were earned or lost. Only set when the change results from a rule engine execution; empty otherwise. | [optional] 
**EmployeeName** | **string** |  | 
**UserID** | Pointer to **int64** |  | [optional] 
**CurrentPoints** | **float32** |  | 
**Actions** | Pointer to [**[]IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction**](IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction.md) |  | [optional] 
**PublishedAt** | **time.Time** | Timestamp when the event was published. | 
**OldTier** | Pointer to **string** |  | [optional] 
**TierExpirationDate** | Pointer to **time.Time** |  | [optional] 
**TimestampOfTierChange** | Pointer to **time.Time** |  | [optional] 
**PointsRequiredToTheNextTier** | Pointer to **float32** |  | [optional] 
**NextTier** | Pointer to **string** |  | [optional] 
**Id** | **int64** |  | 
**Created** | **time.Time** |  | 
**CampaignId** | **int64** |  | 
**Value** | **string** |  | 
**UsageLimit** | **int64** |  | 
**DiscountLimit** | Pointer to **float32** |  | [optional] 
**ReservationLimit** | Pointer to **int64** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ExpiryDate** | Pointer to **time.Time** |  | [optional] 
**UsageCounter** | **int64** |  | 
**DiscountCounter** | Pointer to **float32** |  | [optional] 
**DiscountRemainder** | Pointer to **float32** |  | [optional] 
**ReferralId** | Pointer to **int64** |  | [optional] 
**RecipientIntegrationId** | Pointer to **string** |  | [optional] 
**ImportId** | Pointer to **int64** |  | [optional] 
**BatchId** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**Limits** | Pointer to [**[]IntegrationHubEventPayloadCouponBasedNotificationsLimits**](IntegrationHubEventPayloadCouponBasedNotificationsLimits.md) |  | [optional] 

## Methods

### NewIntegrationHubPaginatedEventPayloadDataInner

`func NewIntegrationHubPaginatedEventPayloadDataInner(eventId int64, profileIntegrationID string, loyaltyProgramID int64, loyaltyProgramName string, subledgerID string, sourceOfEvent string, currentTier string, employeeName string, currentPoints float32, publishedAt time.Time, id int64, created time.Time, campaignId int64, value string, usageLimit int64, usageCounter int64, ) *IntegrationHubPaginatedEventPayloadDataInner`

NewIntegrationHubPaginatedEventPayloadDataInner instantiates a new IntegrationHubPaginatedEventPayloadDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubPaginatedEventPayloadDataInnerWithDefaults

`func NewIntegrationHubPaginatedEventPayloadDataInnerWithDefaults() *IntegrationHubPaginatedEventPayloadDataInner`

NewIntegrationHubPaginatedEventPayloadDataInnerWithDefaults instantiates a new IntegrationHubPaginatedEventPayloadDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetProfileIntegrationID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetProfileIntegrationID() string`

GetProfileIntegrationID returns the ProfileIntegrationID field if non-nil, zero value otherwise.

### GetProfileIntegrationIDOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetProfileIntegrationIDOk() (*string, bool)`

GetProfileIntegrationIDOk returns a tuple with the ProfileIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetProfileIntegrationID(v string)`

SetProfileIntegrationID sets ProfileIntegrationID field to given value.


### GetLoyaltyProgramID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLoyaltyProgramID() int64`

GetLoyaltyProgramID returns the LoyaltyProgramID field if non-nil, zero value otherwise.

### GetLoyaltyProgramIDOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLoyaltyProgramIDOk() (*int64, bool)`

GetLoyaltyProgramIDOk returns a tuple with the LoyaltyProgramID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetLoyaltyProgramID(v int64)`

SetLoyaltyProgramID sets LoyaltyProgramID field to given value.


### GetLoyaltyProgramName

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLoyaltyProgramName() string`

GetLoyaltyProgramName returns the LoyaltyProgramName field if non-nil, zero value otherwise.

### GetLoyaltyProgramNameOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLoyaltyProgramNameOk() (*string, bool)`

GetLoyaltyProgramNameOk returns a tuple with the LoyaltyProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramName

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetLoyaltyProgramName(v string)`

SetLoyaltyProgramName sets LoyaltyProgramName field to given value.


### GetSubledgerID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSubledgerID() string`

GetSubledgerID returns the SubledgerID field if non-nil, zero value otherwise.

### GetSubledgerIDOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSubledgerIDOk() (*string, bool)`

GetSubledgerIDOk returns a tuple with the SubledgerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetSubledgerID(v string)`

SetSubledgerID sets SubledgerID field to given value.


### GetSourceOfEvent

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetCurrentTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCurrentTier() string`

GetCurrentTier returns the CurrentTier field if non-nil, zero value otherwise.

### GetCurrentTierOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCurrentTierOk() (*string, bool)`

GetCurrentTierOk returns a tuple with the CurrentTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetCurrentTier(v string)`

SetCurrentTier sets CurrentTier field to given value.


### GetSessionIntegrationID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSessionIntegrationID() string`

GetSessionIntegrationID returns the SessionIntegrationID field if non-nil, zero value otherwise.

### GetSessionIntegrationIDOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetSessionIntegrationIDOk() (*string, bool)`

GetSessionIntegrationIDOk returns a tuple with the SessionIntegrationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetSessionIntegrationID(v string)`

SetSessionIntegrationID sets SessionIntegrationID field to given value.

### HasSessionIntegrationID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasSessionIntegrationID() bool`

HasSessionIntegrationID returns a boolean if a field has been set.

### GetEmployeeName

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetUserID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUserID() int64`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUserIDOk() (*int64, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetUserID(v int64)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentPoints

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCurrentPoints() float32`

GetCurrentPoints returns the CurrentPoints field if non-nil, zero value otherwise.

### GetCurrentPointsOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCurrentPointsOk() (*float32, bool)`

GetCurrentPointsOk returns a tuple with the CurrentPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPoints

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetCurrentPoints(v float32)`

SetCurrentPoints sets CurrentPoints field to given value.


### GetActions

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetActions() []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetActionsOk() (*[]IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetActions(v []IntegrationHubEventPayloadLoyaltyProfileBasedPointsChangedNotificationAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPublishedAt

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetOldTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetOldTier() string`

GetOldTier returns the OldTier field if non-nil, zero value otherwise.

### GetOldTierOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetOldTierOk() (*string, bool)`

GetOldTierOk returns a tuple with the OldTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetOldTier(v string)`

SetOldTier sets OldTier field to given value.

### HasOldTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasOldTier() bool`

HasOldTier returns a boolean if a field has been set.

### GetTierExpirationDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetTierExpirationDate() time.Time`

GetTierExpirationDate returns the TierExpirationDate field if non-nil, zero value otherwise.

### GetTierExpirationDateOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetTierExpirationDateOk() (*time.Time, bool)`

GetTierExpirationDateOk returns a tuple with the TierExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierExpirationDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetTierExpirationDate(v time.Time)`

SetTierExpirationDate sets TierExpirationDate field to given value.

### HasTierExpirationDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasTierExpirationDate() bool`

HasTierExpirationDate returns a boolean if a field has been set.

### GetTimestampOfTierChange

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetTimestampOfTierChange() time.Time`

GetTimestampOfTierChange returns the TimestampOfTierChange field if non-nil, zero value otherwise.

### GetTimestampOfTierChangeOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetTimestampOfTierChangeOk() (*time.Time, bool)`

GetTimestampOfTierChangeOk returns a tuple with the TimestampOfTierChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampOfTierChange

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetTimestampOfTierChange(v time.Time)`

SetTimestampOfTierChange sets TimestampOfTierChange field to given value.

### HasTimestampOfTierChange

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasTimestampOfTierChange() bool`

HasTimestampOfTierChange returns a boolean if a field has been set.

### GetPointsRequiredToTheNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetPointsRequiredToTheNextTier() float32`

GetPointsRequiredToTheNextTier returns the PointsRequiredToTheNextTier field if non-nil, zero value otherwise.

### GetPointsRequiredToTheNextTierOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetPointsRequiredToTheNextTierOk() (*float32, bool)`

GetPointsRequiredToTheNextTierOk returns a tuple with the PointsRequiredToTheNextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequiredToTheNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetPointsRequiredToTheNextTier(v float32)`

SetPointsRequiredToTheNextTier sets PointsRequiredToTheNextTier field to given value.

### HasPointsRequiredToTheNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasPointsRequiredToTheNextTier() bool`

HasPointsRequiredToTheNextTier returns a boolean if a field has been set.

### GetNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetNextTier() string`

GetNextTier returns the NextTier field if non-nil, zero value otherwise.

### GetNextTierOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetNextTierOk() (*string, bool)`

GetNextTierOk returns a tuple with the NextTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetNextTier(v string)`

SetNextTier sets NextTier field to given value.

### HasNextTier

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasNextTier() bool`

HasNextTier returns a boolean if a field has been set.

### GetId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetValue

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetUsageLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageCounter

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetDiscountCounter

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountCounterOk() (*float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCounter

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetDiscountCounter(v float32)`

SetDiscountCounter sets DiscountCounter field to given value.

### HasDiscountCounter

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### GetDiscountRemainder

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetDiscountRemainderOk() (*float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRemainder

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetDiscountRemainder(v float32)`

SetDiscountRemainder sets DiscountRemainder field to given value.

### HasDiscountRemainder

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### GetReferralId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetReferralId() int64`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetReferralIdOk() (*int64, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetReferralId(v int64)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetImportId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetBatchId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetAttributes

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLimits

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLimits() []IntegrationHubEventPayloadCouponBasedNotificationsLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IntegrationHubPaginatedEventPayloadDataInner) GetLimitsOk() (*[]IntegrationHubEventPayloadCouponBasedNotificationsLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IntegrationHubPaginatedEventPayloadDataInner) SetLimits(v []IntegrationHubEventPayloadCouponBasedNotificationsLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IntegrationHubPaginatedEventPayloadDataInner) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


