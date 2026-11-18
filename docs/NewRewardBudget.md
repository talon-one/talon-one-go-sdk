# NewRewardBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | The customer action that counts towards the budget. | [optional] [default to "unlocks"]
**Scope** | **string** | The type of count that the budget keeps of how many times the reward has  been unlocked.   - &#x60;perReward&#x60;: Counts how many times the reward has been unlocked by all customers. - &#x60;perProfile&#x60;: Counts how many times each customer profile has unlocked the reward.  | 
**Period** | **string** | The time frame after which the budget count resets to &#x60;0&#x60;. If you select &#x60;overall&#x60;, the budget never resets.  The count resets at the start of each calendar time frame, not on the day the budget was created. For example, a &#x60;monthly&#x60; budget created on 20 October counts the unlocks from 20 to 31 October. On 1 November, the count resets to &#x60;0&#x60;.  | 
**Limit** | **int64** | The maximum number of times the reward can be unlocked in the time frame. | 

## Methods

### NewNewRewardBudget

`func NewNewRewardBudget(scope string, period string, limit int64, ) *NewRewardBudget`

NewNewRewardBudget instantiates a new NewRewardBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewRewardBudgetWithDefaults

`func NewNewRewardBudgetWithDefaults() *NewRewardBudget`

NewNewRewardBudgetWithDefaults instantiates a new NewRewardBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *NewRewardBudget) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *NewRewardBudget) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *NewRewardBudget) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *NewRewardBudget) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetScope

`func (o *NewRewardBudget) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NewRewardBudget) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NewRewardBudget) SetScope(v string)`

SetScope sets Scope field to given value.


### GetPeriod

`func (o *NewRewardBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *NewRewardBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *NewRewardBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetLimit

`func (o *NewRewardBudget) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NewRewardBudget) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NewRewardBudget) SetLimit(v int64)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


