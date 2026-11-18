# RewardWithBudgets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**AccountId** | **int64** | The ID of the account that owns this entity. | 
**Name** | **string** | The name of the reward. | 
**ApiName** | **string** | A unique identifier used to reference the reward in API integrations. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**ApplicationIds** | **[]int64** | The IDs of the Applications this reward is connected to.   **Note**: Currently, a reward can only be connected to one Application.  | 
**Sandbox** | **bool** | Indicates if this is a live or sandbox reward. Rewards of a given type can only be connected to Applications of the same type. | 
**EligibilityConditions** | Pointer to [**Rule**](Rule.md) | An optional rule that manages who can see this reward. If not specified, the reward is visible to all customers.  **Note:** Only the &#x60;condition&#x60; field is evaluated within this rule. The &#x60;effects&#x60; field must be an empty array, and &#x60;bindings&#x60; are not supported.  | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) | Rule to apply.  **Note**: The &#x60;bindings&#x60; field inside the rule must not be used in this endpoint. All bindings should be defined at the reward level via the top-level &#x60;bindings&#x60; field.  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated. Each binding pairs a name with a talang expression. The expression is evaluated once and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. Each object defines the specific loyalty program and subledger from which points are deducted when activating the reward.  **Note:** When creating a reward, the &#x60;id&#x60; of each entry is ignored and a new entry is always created.  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the customer rewards unlocked from this reward remain active, relative to when the customer reward was unlocked.  The time format is an **integer** followed by one letter indicating the time unit.  Examples: &#x60;30m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  This field cannot be used in combination with &#x60;expiryDate&#x60;. If both &#x60;validityDuration&#x60; and &#x60;expiryDate&#x60; are not set, customer rewards have unlimited validity.  | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the customer rewards unlocked from this reward. Customer rewards never expire if this is omitted and &#x60;validityDuration&#x60; is not set.  **Note:** - The date must be in the future. - This field cannot be used in combination with &#x60;validityDuration&#x60;.  | [optional] 
**Modified** | Pointer to **time.Time** | The timestamp when the reward was last updated in RFC3339 format. | [optional] 
**Status** | **string** | The status of the reward. | 
**Budgets** | Pointer to [**[]RewardBudgetWithLimitCounter**](RewardBudgetWithLimitCounter.md) | The budgets associated with the reward. This property is returned only when &#x60;includeBudgets&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### NewRewardWithBudgets

`func NewRewardWithBudgets(id int64, created time.Time, accountId int64, name string, apiName string, applicationIds []int64, sandbox bool, status string, ) *RewardWithBudgets`

NewRewardWithBudgets instantiates a new RewardWithBudgets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithBudgetsWithDefaults

`func NewRewardWithBudgetsWithDefaults() *RewardWithBudgets`

NewRewardWithBudgetsWithDefaults instantiates a new RewardWithBudgets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardWithBudgets) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardWithBudgets) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardWithBudgets) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RewardWithBudgets) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RewardWithBudgets) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RewardWithBudgets) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *RewardWithBudgets) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RewardWithBudgets) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RewardWithBudgets) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *RewardWithBudgets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RewardWithBudgets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RewardWithBudgets) SetName(v string)`

SetName sets Name field to given value.


### GetApiName

`func (o *RewardWithBudgets) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *RewardWithBudgets) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *RewardWithBudgets) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetDescription

`func (o *RewardWithBudgets) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardWithBudgets) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardWithBudgets) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardWithBudgets) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *RewardWithBudgets) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *RewardWithBudgets) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *RewardWithBudgets) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetSandbox

`func (o *RewardWithBudgets) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *RewardWithBudgets) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *RewardWithBudgets) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetEligibilityConditions

`func (o *RewardWithBudgets) GetEligibilityConditions() Rule`

GetEligibilityConditions returns the EligibilityConditions field if non-nil, zero value otherwise.

### GetEligibilityConditionsOk

`func (o *RewardWithBudgets) GetEligibilityConditionsOk() (*Rule, bool)`

GetEligibilityConditionsOk returns a tuple with the EligibilityConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibilityConditions

`func (o *RewardWithBudgets) SetEligibilityConditions(v Rule)`

SetEligibilityConditions sets EligibilityConditions field to given value.

### HasEligibilityConditions

`func (o *RewardWithBudgets) HasEligibilityConditions() bool`

HasEligibilityConditions returns a boolean if a field has been set.

### GetRule

`func (o *RewardWithBudgets) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RewardWithBudgets) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RewardWithBudgets) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *RewardWithBudgets) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetBindings

`func (o *RewardWithBudgets) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *RewardWithBudgets) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *RewardWithBudgets) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *RewardWithBudgets) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetPointsRequired

`func (o *RewardWithBudgets) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *RewardWithBudgets) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *RewardWithBudgets) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *RewardWithBudgets) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.

### GetValidityDuration

`func (o *RewardWithBudgets) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *RewardWithBudgets) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *RewardWithBudgets) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *RewardWithBudgets) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetExpiryDate

`func (o *RewardWithBudgets) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *RewardWithBudgets) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *RewardWithBudgets) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *RewardWithBudgets) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetModified

`func (o *RewardWithBudgets) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RewardWithBudgets) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RewardWithBudgets) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *RewardWithBudgets) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetStatus

`func (o *RewardWithBudgets) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RewardWithBudgets) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RewardWithBudgets) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBudgets

`func (o *RewardWithBudgets) GetBudgets() []RewardBudgetWithLimitCounter`

GetBudgets returns the Budgets field if non-nil, zero value otherwise.

### GetBudgetsOk

`func (o *RewardWithBudgets) GetBudgetsOk() (*[]RewardBudgetWithLimitCounter, bool)`

GetBudgetsOk returns a tuple with the Budgets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgets

`func (o *RewardWithBudgets) SetBudgets(v []RewardBudgetWithLimitCounter)`

SetBudgets sets Budgets field to given value.

### HasBudgets

`func (o *RewardWithBudgets) HasBudgets() bool`

HasBudgets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


