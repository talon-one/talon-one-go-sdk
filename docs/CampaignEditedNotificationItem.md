# CampaignEditedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | **interface{}** | The campaign whose state changed. | 
**OldCampaign** | **interface{}** | The campaign before the change. | 
**Ruleset** | Pointer to **interface{}** | The current ruleset. | [optional] 

## Methods

### NewCampaignEditedNotificationItem

`func NewCampaignEditedNotificationItem(event string, campaign interface{}, oldCampaign interface{}, ) *CampaignEditedNotificationItem`

NewCampaignEditedNotificationItem instantiates a new CampaignEditedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEditedNotificationItemWithDefaults

`func NewCampaignEditedNotificationItemWithDefaults() *CampaignEditedNotificationItem`

NewCampaignEditedNotificationItemWithDefaults instantiates a new CampaignEditedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignEditedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignEditedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignEditedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignEditedNotificationItem) GetCampaign() interface{}`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignEditedNotificationItem) GetCampaignOk() (*interface{}, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignEditedNotificationItem) SetCampaign(v interface{})`

SetCampaign sets Campaign field to given value.


### SetCampaignNil

`func (o *CampaignEditedNotificationItem) SetCampaignNil(b bool)`

 SetCampaignNil sets the value for Campaign to be an explicit nil

### UnsetCampaign
`func (o *CampaignEditedNotificationItem) UnsetCampaign()`

UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
### GetOldCampaign

`func (o *CampaignEditedNotificationItem) GetOldCampaign() interface{}`

GetOldCampaign returns the OldCampaign field if non-nil, zero value otherwise.

### GetOldCampaignOk

`func (o *CampaignEditedNotificationItem) GetOldCampaignOk() (*interface{}, bool)`

GetOldCampaignOk returns a tuple with the OldCampaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldCampaign

`func (o *CampaignEditedNotificationItem) SetOldCampaign(v interface{})`

SetOldCampaign sets OldCampaign field to given value.


### SetOldCampaignNil

`func (o *CampaignEditedNotificationItem) SetOldCampaignNil(b bool)`

 SetOldCampaignNil sets the value for OldCampaign to be an explicit nil

### UnsetOldCampaign
`func (o *CampaignEditedNotificationItem) UnsetOldCampaign()`

UnsetOldCampaign ensures that no value is present for OldCampaign, not even an explicit nil
### GetRuleset

`func (o *CampaignEditedNotificationItem) GetRuleset() interface{}`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignEditedNotificationItem) GetRulesetOk() (*interface{}, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignEditedNotificationItem) SetRuleset(v interface{})`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignEditedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRulesetNil

`func (o *CampaignEditedNotificationItem) SetRulesetNil(b bool)`

 SetRulesetNil sets the value for Ruleset to be an explicit nil

### UnsetRuleset
`func (o *CampaignEditedNotificationItem) UnsetRuleset()`

UnsetRuleset ensures that no value is present for Ruleset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


