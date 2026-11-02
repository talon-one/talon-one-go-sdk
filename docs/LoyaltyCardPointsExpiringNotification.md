# LoyaltyCardPointsExpiringNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]ExpiringCardPointsData**](ExpiringCardPointsData.md) | The array of expiring points. | 
**NotificationType** | **string** | The type of notification. | 

## Methods

### NewLoyaltyCardPointsExpiringNotification

`func NewLoyaltyCardPointsExpiringNotification(totalResultSize int64, data []ExpiringCardPointsData, notificationType string, ) *LoyaltyCardPointsExpiringNotification`

NewLoyaltyCardPointsExpiringNotification instantiates a new LoyaltyCardPointsExpiringNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardPointsExpiringNotificationWithDefaults

`func NewLoyaltyCardPointsExpiringNotificationWithDefaults() *LoyaltyCardPointsExpiringNotification`

NewLoyaltyCardPointsExpiringNotificationWithDefaults instantiates a new LoyaltyCardPointsExpiringNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *LoyaltyCardPointsExpiringNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *LoyaltyCardPointsExpiringNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *LoyaltyCardPointsExpiringNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *LoyaltyCardPointsExpiringNotification) GetData() []ExpiringCardPointsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoyaltyCardPointsExpiringNotification) GetDataOk() (*[]ExpiringCardPointsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoyaltyCardPointsExpiringNotification) SetData(v []ExpiringCardPointsData)`

SetData sets Data field to given value.


### GetNotificationType

`func (o *LoyaltyCardPointsExpiringNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyCardPointsExpiringNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyCardPointsExpiringNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


