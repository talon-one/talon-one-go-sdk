# RewardBudgetWithLimitCounter

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
**LimitCounter** | Pointer to [**[]RewardBudgetLimitCounter**](RewardBudgetLimitCounter.md) | The count of times the reward has been unlocked in the current period, and how many times it can still be unlocked before the limit is reached. It is only returned for budgets with the &#x60;perReward&#x60; scope, because every customer profile has their own count in a &#x60;perProfile&#x60; budget.  The count is stored for each Application. A reward is connected to one Application, so the list always has exactly one item.  | [optional] 

## Methods

### NewRewardBudgetWithLimitCounter

`func NewRewardBudgetWithLimitCounter(id int64, created time.Time, rewardId int64, metric string, scope string, period string, limit int64, ) *RewardBudgetWithLimitCounter`

NewRewardBudgetWithLimitCounter instantiates a new RewardBudgetWithLimitCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardBudgetWithLimitCounterWithDefaults

`func NewRewardBudgetWithLimitCounterWithDefaults() *RewardBudgetWithLimitCounter`

NewRewardBudgetWithLimitCounterWithDefaults instantiates a new RewardBudgetWithLimitCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardBudgetWithLimitCounter) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardBudgetWithLimitCounter) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardBudgetWithLimitCounter) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RewardBudgetWithLimitCounter) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RewardBudgetWithLimitCounter) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RewardBudgetWithLimitCounter) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetRewardId

`func (o *RewardBudgetWithLimitCounter) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *RewardBudgetWithLimitCounter) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *RewardBudgetWithLimitCounter) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.


### GetMetric

`func (o *RewardBudgetWithLimitCounter) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *RewardBudgetWithLimitCounter) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *RewardBudgetWithLimitCounter) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetScope

`func (o *RewardBudgetWithLimitCounter) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RewardBudgetWithLimitCounter) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RewardBudgetWithLimitCounter) SetScope(v string)`

SetScope sets Scope field to given value.


### GetPeriod

`func (o *RewardBudgetWithLimitCounter) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RewardBudgetWithLimitCounter) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RewardBudgetWithLimitCounter) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetLimit

`func (o *RewardBudgetWithLimitCounter) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RewardBudgetWithLimitCounter) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RewardBudgetWithLimitCounter) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetLimitCounter

`func (o *RewardBudgetWithLimitCounter) GetLimitCounter() []RewardBudgetLimitCounter`

GetLimitCounter returns the LimitCounter field if non-nil, zero value otherwise.

### GetLimitCounterOk

`func (o *RewardBudgetWithLimitCounter) GetLimitCounterOk() (*[]RewardBudgetLimitCounter, bool)`

GetLimitCounterOk returns a tuple with the LimitCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitCounter

`func (o *RewardBudgetWithLimitCounter) SetLimitCounter(v []RewardBudgetLimitCounter)`

SetLimitCounter sets LimitCounter field to given value.

### HasLimitCounter

`func (o *RewardBudgetWithLimitCounter) HasLimitCounter() bool`

HasLimitCounter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


