# CampaignCreatedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | **interface{}** | The campaign whose state changed. | 
**Ruleset** | Pointer to **interface{}** | The current ruleset. | [optional] 
**EvaluationPosition** | **interface{}** | The campaign position within the evaluation tree. | 

## Methods

### NewCampaignCreatedNotificationItem

`func NewCampaignCreatedNotificationItem(event string, campaign interface{}, evaluationPosition interface{}, ) *CampaignCreatedNotificationItem`

NewCampaignCreatedNotificationItem instantiates a new CampaignCreatedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCreatedNotificationItemWithDefaults

`func NewCampaignCreatedNotificationItemWithDefaults() *CampaignCreatedNotificationItem`

NewCampaignCreatedNotificationItemWithDefaults instantiates a new CampaignCreatedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignCreatedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignCreatedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignCreatedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignCreatedNotificationItem) GetCampaign() interface{}`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignCreatedNotificationItem) GetCampaignOk() (*interface{}, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignCreatedNotificationItem) SetCampaign(v interface{})`

SetCampaign sets Campaign field to given value.


### SetCampaignNil

`func (o *CampaignCreatedNotificationItem) SetCampaignNil(b bool)`

 SetCampaignNil sets the value for Campaign to be an explicit nil

### UnsetCampaign
`func (o *CampaignCreatedNotificationItem) UnsetCampaign()`

UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
### GetRuleset

`func (o *CampaignCreatedNotificationItem) GetRuleset() interface{}`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignCreatedNotificationItem) GetRulesetOk() (*interface{}, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignCreatedNotificationItem) SetRuleset(v interface{})`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignCreatedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRulesetNil

`func (o *CampaignCreatedNotificationItem) SetRulesetNil(b bool)`

 SetRulesetNil sets the value for Ruleset to be an explicit nil

### UnsetRuleset
`func (o *CampaignCreatedNotificationItem) UnsetRuleset()`

UnsetRuleset ensures that no value is present for Ruleset, not even an explicit nil
### GetEvaluationPosition

`func (o *CampaignCreatedNotificationItem) GetEvaluationPosition() interface{}`

GetEvaluationPosition returns the EvaluationPosition field if non-nil, zero value otherwise.

### GetEvaluationPositionOk

`func (o *CampaignCreatedNotificationItem) GetEvaluationPositionOk() (*interface{}, bool)`

GetEvaluationPositionOk returns a tuple with the EvaluationPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationPosition

`func (o *CampaignCreatedNotificationItem) SetEvaluationPosition(v interface{})`

SetEvaluationPosition sets EvaluationPosition field to given value.


### SetEvaluationPositionNil

`func (o *CampaignCreatedNotificationItem) SetEvaluationPositionNil(b bool)`

 SetEvaluationPositionNil sets the value for EvaluationPosition to be an explicit nil

### UnsetEvaluationPosition
`func (o *CampaignCreatedNotificationItem) UnsetEvaluationPosition()`

UnsetEvaluationPosition ensures that no value is present for EvaluationPosition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


