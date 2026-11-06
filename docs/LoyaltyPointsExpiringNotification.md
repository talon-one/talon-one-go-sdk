# LoyaltyPointsExpiringNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]ExpiringPointsData**](ExpiringPointsData.md) | The array of expiring points. | 
**NotificationType** | **string** | The type of notification. | 

## Methods

### NewLoyaltyPointsExpiringNotification

`func NewLoyaltyPointsExpiringNotification(totalResultSize int64, data []ExpiringPointsData, notificationType string, ) *LoyaltyPointsExpiringNotification`

NewLoyaltyPointsExpiringNotification instantiates a new LoyaltyPointsExpiringNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyPointsExpiringNotificationWithDefaults

`func NewLoyaltyPointsExpiringNotificationWithDefaults() *LoyaltyPointsExpiringNotification`

NewLoyaltyPointsExpiringNotificationWithDefaults instantiates a new LoyaltyPointsExpiringNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *LoyaltyPointsExpiringNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *LoyaltyPointsExpiringNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *LoyaltyPointsExpiringNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *LoyaltyPointsExpiringNotification) GetData() []ExpiringPointsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoyaltyPointsExpiringNotification) GetDataOk() (*[]ExpiringPointsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoyaltyPointsExpiringNotification) SetData(v []ExpiringPointsData)`

SetData sets Data field to given value.


### GetNotificationType

`func (o *LoyaltyPointsExpiringNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *LoyaltyPointsExpiringNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *LoyaltyPointsExpiringNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


