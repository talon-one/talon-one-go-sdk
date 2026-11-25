# AchievementProgressWithDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The status of the achievement. | 
**Progress** | **float32** | The current progress of the customer in the achievement. | 
**StartDate** | Pointer to **time.Time** | Timestamp at which the customer started the achievement. | [optional] 
**CompletionDate** | Pointer to **time.Time** | Timestamp at which point the customer completed the achievement. | [optional] 
**EndDate** | Pointer to **time.Time** | Timestamp at which point the achievement ends and resets for the customer. | [optional] 
**AchievementId** | **int64** | The internal ID of the achievement. | 
**Name** | **string** | The internal name of the achievement used in API requests.  | 
**Title** | **string** | The display name of the achievement in the Campaign Manager. | 
**Description** | **string** | The description of the achievement in the Campaign Manager. | 
**CampaignId** | Pointer to **int64** | This property is **deprecated**. Use &#x60;campaignIds&#x60; (Integration API) or &#x60;referencedByCampaigns&#x60; (Management API) instead. This field contains the first campaign ID from the related &#x60;campaignIds&#x60;, and is omitted when &#x60;campaignIds&#x60; is empty. | [optional] 
**CampaignIds** | **[]int64** | The IDs of the campaigns that reference this achievement, in ascending order. | 
**ReferencedByCampaigns** | [**[]CampaignReference**](CampaignReference.md) | The campaigns that reference this achievement, in ascending order of their &#x60;id&#x60;. | 
**Target** | Pointer to **float32** | The required number of actions or the transactional milestone to complete the achievement. | [optional] 
**AchievementPeriod** | Pointer to **string** | The relative duration after which the achievement ends and resets for a particular customer profile. This property is omitted for on-completion achievements which have no specified period.  The period is a positive real number followed by one letter indicating the time unit, and an optional rounding suffix.  Examples: &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;, &#x60;30D_D&#x60;, &#x60;1W_U&#x60;.  The letters indicate the following time units:  - &#x60;s&#x60;: Seconds.  - &#x60;m&#x60;: Minutes.  - &#x60;h&#x60;: Hours.  - &#x60;D&#x60;: Days.  - &#x60;W&#x60;: Weeks.  - &#x60;M&#x60;: Months.  - &#x60;Y&#x60;: Years.  If a suffix is present, it indicates the following rounding:  - &#x60;_D&#x60;: Rounds down, to the start of the day. It applies to days only. For example, &#x60;30D_D&#x60; ends at the start of the day, 30 days after the customer starts the achievement.  - &#x60;_U&#x60;: Rounds up, to the end of the day, week, month or year. It applies to days, weeks, months and years. For example, &#x60;1W_U&#x60; ends at the end of the week in which the one-week period would end.  | [optional] 
**AchievementRecurrencePolicy** | **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again. - &#x60;on_completion&#x60;: When the customer progress status reaches &#x60;completed&#x60;, the achievement resets and becomes available again.  | 
**AchievementActivationPolicy** | **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | 
**AchievementFixedStartDate** | Pointer to **time.Time** | The achievement&#39;s start date when &#x60;achievementActivationPolicy&#x60; is equal to &#x60;fixed_schedule&#x60;.  **Note:** It is an RFC3339 timestamp string.  | [optional] 
**AchievementEndDate** | Pointer to **time.Time** | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It is an RFC3339 timestamp string.  | [optional] 
**AchievementAllowRollbackAfterCompletion** | Pointer to **bool** | When &#x60;true&#x60;, customer progress can be rolled back in completed achievements. | [optional] 

## Methods

### NewAchievementProgressWithDefinition

`func NewAchievementProgressWithDefinition(status string, progress float32, achievementId int64, name string, title string, description string, campaignIds []int64, referencedByCampaigns []CampaignReference, achievementRecurrencePolicy string, achievementActivationPolicy string, ) *AchievementProgressWithDefinition`

NewAchievementProgressWithDefinition instantiates a new AchievementProgressWithDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAchievementProgressWithDefinitionWithDefaults

`func NewAchievementProgressWithDefinitionWithDefaults() *AchievementProgressWithDefinition`

NewAchievementProgressWithDefinitionWithDefaults instantiates a new AchievementProgressWithDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AchievementProgressWithDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AchievementProgressWithDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AchievementProgressWithDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *AchievementProgressWithDefinition) GetProgress() float32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *AchievementProgressWithDefinition) GetProgressOk() (*float32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *AchievementProgressWithDefinition) SetProgress(v float32)`

SetProgress sets Progress field to given value.


### GetStartDate

`func (o *AchievementProgressWithDefinition) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AchievementProgressWithDefinition) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AchievementProgressWithDefinition) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AchievementProgressWithDefinition) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetCompletionDate

`func (o *AchievementProgressWithDefinition) GetCompletionDate() time.Time`

GetCompletionDate returns the CompletionDate field if non-nil, zero value otherwise.

### GetCompletionDateOk

`func (o *AchievementProgressWithDefinition) GetCompletionDateOk() (*time.Time, bool)`

GetCompletionDateOk returns a tuple with the CompletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDate

`func (o *AchievementProgressWithDefinition) SetCompletionDate(v time.Time)`

SetCompletionDate sets CompletionDate field to given value.

### HasCompletionDate

`func (o *AchievementProgressWithDefinition) HasCompletionDate() bool`

HasCompletionDate returns a boolean if a field has been set.

### GetEndDate

`func (o *AchievementProgressWithDefinition) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *AchievementProgressWithDefinition) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *AchievementProgressWithDefinition) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *AchievementProgressWithDefinition) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAchievementId

`func (o *AchievementProgressWithDefinition) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *AchievementProgressWithDefinition) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *AchievementProgressWithDefinition) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetName

`func (o *AchievementProgressWithDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AchievementProgressWithDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AchievementProgressWithDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *AchievementProgressWithDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AchievementProgressWithDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AchievementProgressWithDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AchievementProgressWithDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AchievementProgressWithDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AchievementProgressWithDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCampaignId

`func (o *AchievementProgressWithDefinition) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *AchievementProgressWithDefinition) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *AchievementProgressWithDefinition) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *AchievementProgressWithDefinition) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### GetCampaignIds

`func (o *AchievementProgressWithDefinition) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *AchievementProgressWithDefinition) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *AchievementProgressWithDefinition) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.


### GetReferencedByCampaigns

`func (o *AchievementProgressWithDefinition) GetReferencedByCampaigns() []CampaignReference`

GetReferencedByCampaigns returns the ReferencedByCampaigns field if non-nil, zero value otherwise.

### GetReferencedByCampaignsOk

`func (o *AchievementProgressWithDefinition) GetReferencedByCampaignsOk() (*[]CampaignReference, bool)`

GetReferencedByCampaignsOk returns a tuple with the ReferencedByCampaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedByCampaigns

`func (o *AchievementProgressWithDefinition) SetReferencedByCampaigns(v []CampaignReference)`

SetReferencedByCampaigns sets ReferencedByCampaigns field to given value.


### GetTarget

`func (o *AchievementProgressWithDefinition) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AchievementProgressWithDefinition) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AchievementProgressWithDefinition) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *AchievementProgressWithDefinition) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetAchievementPeriod

`func (o *AchievementProgressWithDefinition) GetAchievementPeriod() string`

GetAchievementPeriod returns the AchievementPeriod field if non-nil, zero value otherwise.

### GetAchievementPeriodOk

`func (o *AchievementProgressWithDefinition) GetAchievementPeriodOk() (*string, bool)`

GetAchievementPeriodOk returns a tuple with the AchievementPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementPeriod

`func (o *AchievementProgressWithDefinition) SetAchievementPeriod(v string)`

SetAchievementPeriod sets AchievementPeriod field to given value.

### HasAchievementPeriod

`func (o *AchievementProgressWithDefinition) HasAchievementPeriod() bool`

HasAchievementPeriod returns a boolean if a field has been set.

### GetAchievementRecurrencePolicy

`func (o *AchievementProgressWithDefinition) GetAchievementRecurrencePolicy() string`

GetAchievementRecurrencePolicy returns the AchievementRecurrencePolicy field if non-nil, zero value otherwise.

### GetAchievementRecurrencePolicyOk

`func (o *AchievementProgressWithDefinition) GetAchievementRecurrencePolicyOk() (*string, bool)`

GetAchievementRecurrencePolicyOk returns a tuple with the AchievementRecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementRecurrencePolicy

`func (o *AchievementProgressWithDefinition) SetAchievementRecurrencePolicy(v string)`

SetAchievementRecurrencePolicy sets AchievementRecurrencePolicy field to given value.


### GetAchievementActivationPolicy

`func (o *AchievementProgressWithDefinition) GetAchievementActivationPolicy() string`

GetAchievementActivationPolicy returns the AchievementActivationPolicy field if non-nil, zero value otherwise.

### GetAchievementActivationPolicyOk

`func (o *AchievementProgressWithDefinition) GetAchievementActivationPolicyOk() (*string, bool)`

GetAchievementActivationPolicyOk returns a tuple with the AchievementActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementActivationPolicy

`func (o *AchievementProgressWithDefinition) SetAchievementActivationPolicy(v string)`

SetAchievementActivationPolicy sets AchievementActivationPolicy field to given value.


### GetAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) GetAchievementFixedStartDate() time.Time`

GetAchievementFixedStartDate returns the AchievementFixedStartDate field if non-nil, zero value otherwise.

### GetAchievementFixedStartDateOk

`func (o *AchievementProgressWithDefinition) GetAchievementFixedStartDateOk() (*time.Time, bool)`

GetAchievementFixedStartDateOk returns a tuple with the AchievementFixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) SetAchievementFixedStartDate(v time.Time)`

SetAchievementFixedStartDate sets AchievementFixedStartDate field to given value.

### HasAchievementFixedStartDate

`func (o *AchievementProgressWithDefinition) HasAchievementFixedStartDate() bool`

HasAchievementFixedStartDate returns a boolean if a field has been set.

### GetAchievementEndDate

`func (o *AchievementProgressWithDefinition) GetAchievementEndDate() time.Time`

GetAchievementEndDate returns the AchievementEndDate field if non-nil, zero value otherwise.

### GetAchievementEndDateOk

`func (o *AchievementProgressWithDefinition) GetAchievementEndDateOk() (*time.Time, bool)`

GetAchievementEndDateOk returns a tuple with the AchievementEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementEndDate

`func (o *AchievementProgressWithDefinition) SetAchievementEndDate(v time.Time)`

SetAchievementEndDate sets AchievementEndDate field to given value.

### HasAchievementEndDate

`func (o *AchievementProgressWithDefinition) HasAchievementEndDate() bool`

HasAchievementEndDate returns a boolean if a field has been set.

### GetAchievementAllowRollbackAfterCompletion

`func (o *AchievementProgressWithDefinition) GetAchievementAllowRollbackAfterCompletion() bool`

GetAchievementAllowRollbackAfterCompletion returns the AchievementAllowRollbackAfterCompletion field if non-nil, zero value otherwise.

### GetAchievementAllowRollbackAfterCompletionOk

`func (o *AchievementProgressWithDefinition) GetAchievementAllowRollbackAfterCompletionOk() (*bool, bool)`

GetAchievementAllowRollbackAfterCompletionOk returns a tuple with the AchievementAllowRollbackAfterCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementAllowRollbackAfterCompletion

`func (o *AchievementProgressWithDefinition) SetAchievementAllowRollbackAfterCompletion(v bool)`

SetAchievementAllowRollbackAfterCompletion sets AchievementAllowRollbackAfterCompletion field to given value.

### HasAchievementAllowRollbackAfterCompletion

`func (o *AchievementProgressWithDefinition) HasAchievementAllowRollbackAfterCompletion() bool`

HasAchievementAllowRollbackAfterCompletion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


