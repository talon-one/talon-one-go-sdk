# Reward

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
**VisibilityConditions** | Pointer to [**Rule**](Rule.md) | An optional rule that manages who can see this reward. If not specified, the reward is visible to all customers.  **Note:** Only the &#x60;condition&#x60; field is evaluated within this rule. The &#x60;effects&#x60; field must be an empty array, and &#x60;bindings&#x60; are not supported.  | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) | Rule to apply.  **Note**: The &#x60;bindings&#x60; field inside the rule must not be used in this endpoint. All bindings should be defined at the reward level via the top-level &#x60;bindings&#x60; field.  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated.  Each binding pairs a name with a talang expression. The expression is evaluated once  and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 
**Modified** | Pointer to **time.Time** | The timestamp when the reward was last updated in RFC3339 format. | [optional] 
**Status** | **string** | The status of the reward. | 

## Methods

### NewReward

`func NewReward(id int64, created time.Time, accountId int64, name string, apiName string, applicationIds []int64, sandbox bool, status string, ) *Reward`

NewReward instantiates a new Reward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithDefaults

`func NewRewardWithDefaults() *Reward`

NewRewardWithDefaults instantiates a new Reward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reward) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reward) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reward) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Reward) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Reward) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Reward) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetAccountId

`func (o *Reward) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Reward) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Reward) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *Reward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reward) SetName(v string)`

SetName sets Name field to given value.


### GetApiName

`func (o *Reward) GetApiName() string`

GetApiName returns the ApiName field if non-nil, zero value otherwise.

### GetApiNameOk

`func (o *Reward) GetApiNameOk() (*string, bool)`

GetApiNameOk returns a tuple with the ApiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiName

`func (o *Reward) SetApiName(v string)`

SetApiName sets ApiName field to given value.


### GetDescription

`func (o *Reward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Reward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Reward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Reward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApplicationIds

`func (o *Reward) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *Reward) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *Reward) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.


### GetSandbox

`func (o *Reward) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Reward) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Reward) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.


### GetVisibilityConditions

`func (o *Reward) GetVisibilityConditions() Rule`

GetVisibilityConditions returns the VisibilityConditions field if non-nil, zero value otherwise.

### GetVisibilityConditionsOk

`func (o *Reward) GetVisibilityConditionsOk() (*Rule, bool)`

GetVisibilityConditionsOk returns a tuple with the VisibilityConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityConditions

`func (o *Reward) SetVisibilityConditions(v Rule)`

SetVisibilityConditions sets VisibilityConditions field to given value.

### HasVisibilityConditions

`func (o *Reward) HasVisibilityConditions() bool`

HasVisibilityConditions returns a boolean if a field has been set.

### GetRule

`func (o *Reward) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *Reward) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *Reward) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *Reward) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetBindings

`func (o *Reward) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *Reward) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *Reward) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *Reward) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetModified

`func (o *Reward) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Reward) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Reward) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Reward) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetStatus

`func (o *Reward) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Reward) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Reward) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


