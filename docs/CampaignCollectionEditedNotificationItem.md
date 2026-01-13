# CampaignCollectionEditedNotificationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | **string** | The type of the event. Can be one of the following: [&#39;campaign_state_changed&#39;, &#39;campaign_ruleset_changed&#39;, &#39;campaign_edited&#39;, &#39;campaign_created&#39;, &#39;campaign_deleted&#39;]  | 
**Campaign** | **interface{}** | The current campaign. | 
**Ruleset** | Pointer to **interface{}** | The current ruleset. | [optional] 
**Collection** | **interface{}** | The collection that was edited. | 

## Methods

### NewCampaignCollectionEditedNotificationItem

`func NewCampaignCollectionEditedNotificationItem(event string, campaign interface{}, collection interface{}, ) *CampaignCollectionEditedNotificationItem`

NewCampaignCollectionEditedNotificationItem instantiates a new CampaignCollectionEditedNotificationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignCollectionEditedNotificationItemWithDefaults

`func NewCampaignCollectionEditedNotificationItemWithDefaults() *CampaignCollectionEditedNotificationItem`

NewCampaignCollectionEditedNotificationItemWithDefaults instantiates a new CampaignCollectionEditedNotificationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *CampaignCollectionEditedNotificationItem) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *CampaignCollectionEditedNotificationItem) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *CampaignCollectionEditedNotificationItem) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCampaign

`func (o *CampaignCollectionEditedNotificationItem) GetCampaign() interface{}`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *CampaignCollectionEditedNotificationItem) GetCampaignOk() (*interface{}, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *CampaignCollectionEditedNotificationItem) SetCampaign(v interface{})`

SetCampaign sets Campaign field to given value.


### SetCampaignNil

`func (o *CampaignCollectionEditedNotificationItem) SetCampaignNil(b bool)`

 SetCampaignNil sets the value for Campaign to be an explicit nil

### UnsetCampaign
`func (o *CampaignCollectionEditedNotificationItem) UnsetCampaign()`

UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
### GetRuleset

`func (o *CampaignCollectionEditedNotificationItem) GetRuleset() interface{}`

GetRuleset returns the Ruleset field if non-nil, zero value otherwise.

### GetRulesetOk

`func (o *CampaignCollectionEditedNotificationItem) GetRulesetOk() (*interface{}, bool)`

GetRulesetOk returns a tuple with the Ruleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleset

`func (o *CampaignCollectionEditedNotificationItem) SetRuleset(v interface{})`

SetRuleset sets Ruleset field to given value.

### HasRuleset

`func (o *CampaignCollectionEditedNotificationItem) HasRuleset() bool`

HasRuleset returns a boolean if a field has been set.

### SetRulesetNil

`func (o *CampaignCollectionEditedNotificationItem) SetRulesetNil(b bool)`

 SetRulesetNil sets the value for Ruleset to be an explicit nil

### UnsetRuleset
`func (o *CampaignCollectionEditedNotificationItem) UnsetRuleset()`

UnsetRuleset ensures that no value is present for Ruleset, not even an explicit nil
### GetCollection

`func (o *CampaignCollectionEditedNotificationItem) GetCollection() interface{}`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *CampaignCollectionEditedNotificationItem) GetCollectionOk() (*interface{}, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *CampaignCollectionEditedNotificationItem) SetCollection(v interface{})`

SetCollection sets Collection field to given value.


### SetCollectionNil

`func (o *CampaignCollectionEditedNotificationItem) SetCollectionNil(b bool)`

 SetCollectionNil sets the value for Collection to be an explicit nil

### UnsetCollection
`func (o *CampaignCollectionEditedNotificationItem) UnsetCollection()`

UnsetCollection ensures that no value is present for Collection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


