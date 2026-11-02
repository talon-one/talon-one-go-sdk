# StrikethroughPriceNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | The version of the strikethrough pricing notification. Set for **scheduled** strikethrough pricing updates only.  | [optional] 
**ValidFrom** | Pointer to **time.Time** | Timestamp at which the strikethrough pricing update becomes valid. Set for **scheduled** strikethrough pricing updates (version: v2) only.  | [optional] 
**ApplicationId** | **int64** | The ID of the Application to which the catalog items labels belongs. | 
**CurrentBatch** | **int64** | The batch number of the notification. Notifications might be sent in different batches. | 
**TotalBatches** | **int64** | The total number of batches for the notification. | 
**Trigger** | [**StrikethroughTrigger**](StrikethroughTrigger.md) |  | 
**ChangedItems** | [**[]StrikethroughChangedItem**](StrikethroughChangedItem.md) |  | 
**NotificationType** | **string** | The type of notification. | 
**SentAt** | **time.Time** | Timestamp when the notification was sent by Talon.One. There may be a delay before the notification is delivered to the user. | 

## Methods

### NewStrikethroughPriceNotification

`func NewStrikethroughPriceNotification(applicationId int64, currentBatch int64, totalBatches int64, trigger StrikethroughTrigger, changedItems []StrikethroughChangedItem, notificationType string, sentAt time.Time, ) *StrikethroughPriceNotification`

NewStrikethroughPriceNotification instantiates a new StrikethroughPriceNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughPriceNotificationWithDefaults

`func NewStrikethroughPriceNotificationWithDefaults() *StrikethroughPriceNotification`

NewStrikethroughPriceNotificationWithDefaults instantiates a new StrikethroughPriceNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *StrikethroughPriceNotification) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StrikethroughPriceNotification) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StrikethroughPriceNotification) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StrikethroughPriceNotification) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetValidFrom

`func (o *StrikethroughPriceNotification) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *StrikethroughPriceNotification) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *StrikethroughPriceNotification) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *StrikethroughPriceNotification) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetApplicationId

`func (o *StrikethroughPriceNotification) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *StrikethroughPriceNotification) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *StrikethroughPriceNotification) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetCurrentBatch

`func (o *StrikethroughPriceNotification) GetCurrentBatch() int64`

GetCurrentBatch returns the CurrentBatch field if non-nil, zero value otherwise.

### GetCurrentBatchOk

`func (o *StrikethroughPriceNotification) GetCurrentBatchOk() (*int64, bool)`

GetCurrentBatchOk returns a tuple with the CurrentBatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBatch

`func (o *StrikethroughPriceNotification) SetCurrentBatch(v int64)`

SetCurrentBatch sets CurrentBatch field to given value.


### GetTotalBatches

`func (o *StrikethroughPriceNotification) GetTotalBatches() int64`

GetTotalBatches returns the TotalBatches field if non-nil, zero value otherwise.

### GetTotalBatchesOk

`func (o *StrikethroughPriceNotification) GetTotalBatchesOk() (*int64, bool)`

GetTotalBatchesOk returns a tuple with the TotalBatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBatches

`func (o *StrikethroughPriceNotification) SetTotalBatches(v int64)`

SetTotalBatches sets TotalBatches field to given value.


### GetTrigger

`func (o *StrikethroughPriceNotification) GetTrigger() StrikethroughTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *StrikethroughPriceNotification) GetTriggerOk() (*StrikethroughTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *StrikethroughPriceNotification) SetTrigger(v StrikethroughTrigger)`

SetTrigger sets Trigger field to given value.


### GetChangedItems

`func (o *StrikethroughPriceNotification) GetChangedItems() []StrikethroughChangedItem`

GetChangedItems returns the ChangedItems field if non-nil, zero value otherwise.

### GetChangedItemsOk

`func (o *StrikethroughPriceNotification) GetChangedItemsOk() (*[]StrikethroughChangedItem, bool)`

GetChangedItemsOk returns a tuple with the ChangedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedItems

`func (o *StrikethroughPriceNotification) SetChangedItems(v []StrikethroughChangedItem)`

SetChangedItems sets ChangedItems field to given value.


### GetNotificationType

`func (o *StrikethroughPriceNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *StrikethroughPriceNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *StrikethroughPriceNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetSentAt

`func (o *StrikethroughPriceNotification) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *StrikethroughPriceNotification) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *StrikethroughPriceNotification) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


