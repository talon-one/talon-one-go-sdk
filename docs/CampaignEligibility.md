# CampaignEligibility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Unique ID of Campaign. | 
**ApplicationId** | **int64** | The ID of the Application that owns this entity. | 
**Name** | **string** | The name of the campaign. | 
**Description** | Pointer to **string** | A detailed description of the campaign. | [optional] 
**StartTime** | Pointer to **time.Time** | Timestamp when the campaign will become active. | [optional] 
**EndTime** | Pointer to **time.Time** | Timestamp when the campaign will become inactive. | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Arbitrary properties associated with this campaign. | [optional] 
**State** | **string** | The state of the campaign.  | [default to "enabled"]
**Tags** | **[]string** | A list of tags for the campaign. | 
**Features** | **[]string** | The features enabled in this campaign. | 
**Rules** | Pointer to [**[]RuleMetadata**](RuleMetadata.md) | A list of rules containing customer-facing details of the rewards defined in the campaign. | [optional] 
**Eligibility** | [**[]CampaignEligibilityDetails**](CampaignEligibilityDetails.md) | The customer&#39;s eligibility for each campaign in the current customer session. | 

## Methods

### NewCampaignEligibility

`func NewCampaignEligibility(id int64, applicationId int64, name string, state string, tags []string, features []string, eligibility []CampaignEligibilityDetails, ) *CampaignEligibility`

NewCampaignEligibility instantiates a new CampaignEligibility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignEligibilityWithDefaults

`func NewCampaignEligibilityWithDefaults() *CampaignEligibility`

NewCampaignEligibilityWithDefaults instantiates a new CampaignEligibility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignEligibility) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignEligibility) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignEligibility) SetId(v int64)`

SetId sets Id field to given value.


### GetApplicationId

`func (o *CampaignEligibility) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CampaignEligibility) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CampaignEligibility) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetName

`func (o *CampaignEligibility) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignEligibility) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignEligibility) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CampaignEligibility) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignEligibility) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignEligibility) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignEligibility) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartTime

`func (o *CampaignEligibility) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CampaignEligibility) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CampaignEligibility) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CampaignEligibility) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *CampaignEligibility) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CampaignEligibility) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CampaignEligibility) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CampaignEligibility) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetAttributes

`func (o *CampaignEligibility) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CampaignEligibility) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CampaignEligibility) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CampaignEligibility) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetState

`func (o *CampaignEligibility) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CampaignEligibility) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CampaignEligibility) SetState(v string)`

SetState sets State field to given value.


### GetTags

`func (o *CampaignEligibility) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CampaignEligibility) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CampaignEligibility) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetFeatures

`func (o *CampaignEligibility) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CampaignEligibility) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CampaignEligibility) SetFeatures(v []string)`

SetFeatures sets Features field to given value.


### GetRules

`func (o *CampaignEligibility) GetRules() []RuleMetadata`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CampaignEligibility) GetRulesOk() (*[]RuleMetadata, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CampaignEligibility) SetRules(v []RuleMetadata)`

SetRules sets Rules field to given value.

### HasRules

`func (o *CampaignEligibility) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetEligibility

`func (o *CampaignEligibility) GetEligibility() []CampaignEligibilityDetails`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *CampaignEligibility) GetEligibilityOk() (*[]CampaignEligibilityDetails, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *CampaignEligibility) SetEligibility(v []CampaignEligibilityDetails)`

SetEligibility sets Eligibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


