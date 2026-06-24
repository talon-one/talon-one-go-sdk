# CustomerAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of the achievement. | 
**Name** | **string** | The internal name of the achievement used in API requests.  | 
**Title** | **string** | The display name of the achievement in the Campaign Manager. | 
**Description** | **string** | The description of the achievement in the Campaign Manager. | 
**Target** | **float32** | The required number of actions or the transactional milestone to complete the achievement. | 
**RecurrencePolicy** | **string** | The policy that determines if and how the achievement recurs. - &#x60;no_recurrence&#x60;: The achievement can be completed only once. - &#x60;on_expiration&#x60;: The achievement resets after it expires and becomes available again. - &#x60;on_completion&#x60;: When the customer progress status reaches &#x60;completed&#x60;, the achievement resets and becomes available again.  | 
**ActivationPolicy** | **string** | The policy that determines how the achievement starts, ends, or resets. - &#x60;user_action&#x60;: The achievement ends or resets relative to when the customer started the achievement. - &#x60;fixed_schedule&#x60;: The achievement starts, ends, or resets for all customers following a fixed schedule.  | 
**FixedStartDate** | Pointer to **time.Time** | The achievement&#39;s start date when &#x60;activationPolicy&#x60; is equal to &#x60;fixed_schedule&#x60;.  **Note:** It is an RFC3339 timestamp string.  | [optional] 
**EndDate** | Pointer to **time.Time** | The achievement&#39;s end date. If defined, customers cannot participate in the achievement after this date.  **Note:** It is an RFC3339 timestamp string.  | [optional] 
**AllowRollbackAfterCompletion** | **bool** | When &#x60;true&#x60;, customer progress can be rolled back in completed achievements. | 
**CurrentProgress** | Pointer to [**AchievementProgress**](AchievementProgress.md) |  | [optional] 

## Methods

### NewCustomerAchievement

`func NewCustomerAchievement(id int64, name string, title string, description string, target float32, recurrencePolicy string, activationPolicy string, allowRollbackAfterCompletion bool, ) *CustomerAchievement`

NewCustomerAchievement instantiates a new CustomerAchievement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerAchievementWithDefaults

`func NewCustomerAchievementWithDefaults() *CustomerAchievement`

NewCustomerAchievementWithDefaults instantiates a new CustomerAchievement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerAchievement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerAchievement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerAchievement) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CustomerAchievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomerAchievement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomerAchievement) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CustomerAchievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CustomerAchievement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CustomerAchievement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CustomerAchievement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerAchievement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerAchievement) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTarget

`func (o *CustomerAchievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CustomerAchievement) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CustomerAchievement) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetRecurrencePolicy

`func (o *CustomerAchievement) GetRecurrencePolicy() string`

GetRecurrencePolicy returns the RecurrencePolicy field if non-nil, zero value otherwise.

### GetRecurrencePolicyOk

`func (o *CustomerAchievement) GetRecurrencePolicyOk() (*string, bool)`

GetRecurrencePolicyOk returns a tuple with the RecurrencePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurrencePolicy

`func (o *CustomerAchievement) SetRecurrencePolicy(v string)`

SetRecurrencePolicy sets RecurrencePolicy field to given value.


### GetActivationPolicy

`func (o *CustomerAchievement) GetActivationPolicy() string`

GetActivationPolicy returns the ActivationPolicy field if non-nil, zero value otherwise.

### GetActivationPolicyOk

`func (o *CustomerAchievement) GetActivationPolicyOk() (*string, bool)`

GetActivationPolicyOk returns a tuple with the ActivationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPolicy

`func (o *CustomerAchievement) SetActivationPolicy(v string)`

SetActivationPolicy sets ActivationPolicy field to given value.


### GetFixedStartDate

`func (o *CustomerAchievement) GetFixedStartDate() time.Time`

GetFixedStartDate returns the FixedStartDate field if non-nil, zero value otherwise.

### GetFixedStartDateOk

`func (o *CustomerAchievement) GetFixedStartDateOk() (*time.Time, bool)`

GetFixedStartDateOk returns a tuple with the FixedStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedStartDate

`func (o *CustomerAchievement) SetFixedStartDate(v time.Time)`

SetFixedStartDate sets FixedStartDate field to given value.

### HasFixedStartDate

`func (o *CustomerAchievement) HasFixedStartDate() bool`

HasFixedStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CustomerAchievement) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CustomerAchievement) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CustomerAchievement) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CustomerAchievement) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetAllowRollbackAfterCompletion

`func (o *CustomerAchievement) GetAllowRollbackAfterCompletion() bool`

GetAllowRollbackAfterCompletion returns the AllowRollbackAfterCompletion field if non-nil, zero value otherwise.

### GetAllowRollbackAfterCompletionOk

`func (o *CustomerAchievement) GetAllowRollbackAfterCompletionOk() (*bool, bool)`

GetAllowRollbackAfterCompletionOk returns a tuple with the AllowRollbackAfterCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRollbackAfterCompletion

`func (o *CustomerAchievement) SetAllowRollbackAfterCompletion(v bool)`

SetAllowRollbackAfterCompletion sets AllowRollbackAfterCompletion field to given value.


### GetCurrentProgress

`func (o *CustomerAchievement) GetCurrentProgress() AchievementProgress`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *CustomerAchievement) GetCurrentProgressOk() (*AchievementProgress, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgress

`func (o *CustomerAchievement) SetCurrentProgress(v AchievementProgress)`

SetCurrentProgress sets CurrentProgress field to given value.

### HasCurrentProgress

`func (o *CustomerAchievement) HasCurrentProgress() bool`

HasCurrentProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


