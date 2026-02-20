# PrismaticEventPayloadCouponBasedNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Limits** | Pointer to [**[]PrismaticEventPayloadCouponBasedNotificationsLimits**](PrismaticEventPayloadCouponBasedNotificationsLimits.md) |  | [optional] 
**PublishedAt** | **time.Time** | Timestamp when the event was published. | 
**SourceOfEvent** | **string** |  | 
**EmployeeName** | **string** |  | 

## Methods

### NewPrismaticEventPayloadCouponBasedNotifications

`func NewPrismaticEventPayloadCouponBasedNotifications(id int64, created time.Time, campaignId int64, value string, usageLimit int64, usageCounter int64, publishedAt time.Time, sourceOfEvent string, employeeName string, ) *PrismaticEventPayloadCouponBasedNotifications`

NewPrismaticEventPayloadCouponBasedNotifications instantiates a new PrismaticEventPayloadCouponBasedNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadCouponBasedNotificationsWithDefaults

`func NewPrismaticEventPayloadCouponBasedNotificationsWithDefaults() *PrismaticEventPayloadCouponBasedNotifications`

NewPrismaticEventPayloadCouponBasedNotificationsWithDefaults instantiates a new PrismaticEventPayloadCouponBasedNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetValue

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetValue(v string)`

SetValue sets Value field to given value.


### GetUsageLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageCounter

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetDiscountCounter

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountCounterOk() (*float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCounter

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetDiscountCounter(v float32)`

SetDiscountCounter sets DiscountCounter field to given value.

### HasDiscountCounter

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### GetDiscountRemainder

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetDiscountRemainderOk() (*float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRemainder

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetDiscountRemainder(v float32)`

SetDiscountRemainder sets DiscountRemainder field to given value.

### HasDiscountRemainder

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### GetReferralId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetReferralId() int64`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetReferralIdOk() (*int64, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetReferralId(v int64)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetImportId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetBatchId

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetAttributes

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLimits

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetLimits() []PrismaticEventPayloadCouponBasedNotificationsLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetLimitsOk() (*[]PrismaticEventPayloadCouponBasedNotificationsLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetLimits(v []PrismaticEventPayloadCouponBasedNotificationsLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *PrismaticEventPayloadCouponBasedNotifications) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPublishedAt

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetSourceOfEvent

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetSourceOfEvent() string`

GetSourceOfEvent returns the SourceOfEvent field if non-nil, zero value otherwise.

### GetSourceOfEventOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetSourceOfEventOk() (*string, bool)`

GetSourceOfEventOk returns a tuple with the SourceOfEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfEvent

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetSourceOfEvent(v string)`

SetSourceOfEvent sets SourceOfEvent field to given value.


### GetEmployeeName

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *PrismaticEventPayloadCouponBasedNotifications) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *PrismaticEventPayloadCouponBasedNotifications) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


