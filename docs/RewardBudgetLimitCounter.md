# RewardBudgetLimitCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counter** | **int64** | The number of times the reward has been unlocked in the current period. It is &#x60;0&#x60; if the reward has not been unlocked by any profile yet. | 
**Remaining** | **int64** | The number of times the reward can still be unlocked. It is never negative. | 

## Methods

### NewRewardBudgetLimitCounter

`func NewRewardBudgetLimitCounter(counter int64, remaining int64, ) *RewardBudgetLimitCounter`

NewRewardBudgetLimitCounter instantiates a new RewardBudgetLimitCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardBudgetLimitCounterWithDefaults

`func NewRewardBudgetLimitCounterWithDefaults() *RewardBudgetLimitCounter`

NewRewardBudgetLimitCounterWithDefaults instantiates a new RewardBudgetLimitCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounter

`func (o *RewardBudgetLimitCounter) GetCounter() int64`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *RewardBudgetLimitCounter) GetCounterOk() (*int64, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *RewardBudgetLimitCounter) SetCounter(v int64)`

SetCounter sets Counter field to given value.


### GetRemaining

`func (o *RewardBudgetLimitCounter) GetRemaining() int64`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *RewardBudgetLimitCounter) GetRemainingOk() (*int64, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *RewardBudgetLimitCounter) SetRemaining(v int64)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


