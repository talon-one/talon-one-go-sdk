# CouponCodeExpiringNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]CouponCodeExpiringNotificationData**](CouponCodeExpiringNotificationData.md) | The array of expiring coupon notifications. | 
**NotificationType** | **string** | The type of notification. | 

## Methods

### NewCouponCodeExpiringNotification

`func NewCouponCodeExpiringNotification(totalResultSize int64, data []CouponCodeExpiringNotificationData, notificationType string, ) *CouponCodeExpiringNotification`

NewCouponCodeExpiringNotification instantiates a new CouponCodeExpiringNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCodeExpiringNotificationWithDefaults

`func NewCouponCodeExpiringNotificationWithDefaults() *CouponCodeExpiringNotification`

NewCouponCodeExpiringNotificationWithDefaults instantiates a new CouponCodeExpiringNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *CouponCodeExpiringNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CouponCodeExpiringNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CouponCodeExpiringNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CouponCodeExpiringNotification) GetData() []CouponCodeExpiringNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CouponCodeExpiringNotification) GetDataOk() (*[]CouponCodeExpiringNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CouponCodeExpiringNotification) SetData(v []CouponCodeExpiringNotificationData)`

SetData sets Data field to given value.


### GetNotificationType

`func (o *CouponCodeExpiringNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CouponCodeExpiringNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CouponCodeExpiringNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


