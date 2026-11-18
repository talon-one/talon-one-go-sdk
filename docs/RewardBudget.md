# RewardBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**RewardId** | **int64** | The ID of the reward this budget applies to. | 
**Metric** | **string** | The customer action that counts towards the budget. | 
**Scope** | **string** | The type of count that the budget keeps of how many times the reward has  been unlocked.  - &#x60;perReward&#x60;: Counts how many times the reward has been unlocked by all customers. - &#x60;perProfile&#x60;: Counts how many times each customer profile has unlocked the reward.  | 
**Period** | **string** | The time frame after which the budget count resets to &#x60;0&#x60;. If you select &#x60;overall&#x60;, the budget never resets.  The count resets at the start of each calendar time frame, not on the day the budget was created. For example, a &#x60;monthly&#x60; budget created on 20 October counts each time the reward is unlocked from 20 to 31 October. On 1 November,  resets to &#x60;0&#x60;.  | 
**Limit** | **int64** | The maximum number of times the reward can be unlocked in the time frame. | 

## Methods

### NewRewardBudget

`func NewRewardBudget(id int64, created time.Time, rewardId int64, metric string, scope string, period string, limit int64, ) *RewardBudget`

NewRewardBudget instantiates a new RewardBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardBudgetWithDefaults

`func NewRewardBudgetWithDefaults() *RewardBudget`

NewRewardBudgetWithDefaults instantiates a new RewardBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardBudget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardBudget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardBudget) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RewardBudget) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RewardBudget) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RewardBudget) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetRewardId

`func (o *RewardBudget) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *RewardBudget) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *RewardBudget) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.


### GetMetric

`func (o *RewardBudget) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RewardBudget) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RewardBudget) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetScope

`func (o *RewardBudget) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RewardBudget) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RewardBudget) SetScope(v string)`

SetScope sets Scope field to given value.


### GetPeriod

`func (o *RewardBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RewardBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RewardBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetLimit

`func (o *RewardBudget) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RewardBudget) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RewardBudget) SetLimit(v int64)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


