# CampaignEvaluationTreeChangedNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **int64** | The ID of the Application whose campaign evaluation tree changed. | 
**OldEvaluationTree** | Pointer to [**CampaignSet**](CampaignSet.md) | The previous campaign evaluation tree. | [optional] 
**EvaluationTree** | [**CampaignSet**](CampaignSet.md) | The new campaign evaluation tree. | 

## Methods

### NewCampaignEvaluationTreeChangedNotificationData

`func NewCampaignEvaluationTreeChangedNotificationData(applicationId int64, evaluationTree CampaignSet, ) *CampaignEvaluationTreeChangedNotificationData`

NewCampaignEvaluationTreeChangedNotificationData instantiates a new CampaignEvaluationTreeChangedNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEvaluationTreeChangedNotificationDataWithDefaults

`func NewCampaignEvaluationTreeChangedNotificationDataWithDefaults() *CampaignEvaluationTreeChangedNotificationData`

NewCampaignEvaluationTreeChangedNotificationDataWithDefaults instantiates a new CampaignEvaluationTreeChangedNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CampaignEvaluationTreeChangedNotificationData) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEvaluationTreeChangedNotificationData) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignEvaluationTreeChangedNotificationData) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotificationData) GetOldEvaluationTree() CampaignSet`

GetOldEvaluationTree returns the OldEvaluationTree field if non-nil, zero value otherwise.

### GetOldEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotificationData) GetOldEvaluationTreeOk() (*CampaignSet, bool)`

GetOldEvaluationTreeOk returns a tuple with the OldEvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotificationData) SetOldEvaluationTree(v CampaignSet)`

SetOldEvaluationTree sets OldEvaluationTree field to given value.

### HasOldEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotificationData) HasOldEvaluationTree() bool`

HasOldEvaluationTree returns a boolean if a field has been set.

### GetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotificationData) GetEvaluationTree() CampaignSet`

GetEvaluationTree returns the EvaluationTree field if non-nil, zero value otherwise.

### GetEvaluationTreeOk

`func (o *CampaignEvaluationTreeChangedNotificationData) GetEvaluationTreeOk() (*CampaignSet, bool)`

GetEvaluationTreeOk returns a tuple with the EvaluationTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationTree

`func (o *CampaignEvaluationTreeChangedNotificationData) SetEvaluationTree(v CampaignSet)`

SetEvaluationTree sets EvaluationTree field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


