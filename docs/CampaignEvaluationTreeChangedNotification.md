# CampaignEvaluationTreeChangedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | **string** | The type of the notification | 
**TotalResultSize** | **int64** | The total size of the result set. | 
**Data** | [**[]CampaignEvaluationTreeChangedNotificationData**](CampaignEvaluationTreeChangedNotificationData.md) | The array of changes. | 

## Methods

### NewCampaignEvaluationTreeChangedNotification

`func NewCampaignEvaluationTreeChangedNotification(notificationType string, totalResultSize int64, data []CampaignEvaluationTreeChangedNotificationData, ) *CampaignEvaluationTreeChangedNotification`

NewCampaignEvaluationTreeChangedNotification instantiates a new CampaignEvaluationTreeChangedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEvaluationTreeChangedNotificationWithDefaults

`func NewCampaignEvaluationTreeChangedNotificationWithDefaults() *CampaignEvaluationTreeChangedNotification`

NewCampaignEvaluationTreeChangedNotificationWithDefaults instantiates a new CampaignEvaluationTreeChangedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *CampaignEvaluationTreeChangedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *CampaignEvaluationTreeChangedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTotalResultSize

`func (o *CampaignEvaluationTreeChangedNotification) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CampaignEvaluationTreeChangedNotification) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CampaignEvaluationTreeChangedNotification) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *CampaignEvaluationTreeChangedNotification) GetData() []CampaignEvaluationTreeChangedNotificationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CampaignEvaluationTreeChangedNotification) GetDataOk() (*[]CampaignEvaluationTreeChangedNotificationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CampaignEvaluationTreeChangedNotification) SetData(v []CampaignEvaluationTreeChangedNotificationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


