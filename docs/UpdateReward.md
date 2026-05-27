# UpdateReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the reward. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**Status** | **string** | The status of the reward. | 
**VisibilityConditions** | Pointer to [**Rule**](Rule.md) | An optional rule that manages who can see this reward. If not specified, the reward is visible to all customers.  **Note:** Only the &#x60;condition&#x60; field is evaluated within this rule. The &#x60;effects&#x60; field must be an empty array, and &#x60;bindings&#x60; are not supported.  | [optional] 
**Rule** | Pointer to [**Rule**](Rule.md) | Rule to apply.  **Note**: The &#x60;bindings&#x60; field inside the rule must not be used in this endpoint. All bindings should be defined at the reward level via the top-level &#x60;bindings&#x60; field.  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated.  Each binding pairs a name with a talang expression. The expression is evaluated once  and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. Each object defines the specific loyalty program and subledger from which points are deducted when activating the reward.  **Note:** - Objects with an &#x60;id&#x60; are updated. - Objects without an &#x60;id&#x60; are created. - Existing objects omitted from the payload are deleted.  | [optional] 

## Methods

### NewUpdateReward

`func NewUpdateReward(name string, status string, ) *UpdateReward`

NewUpdateReward instantiates a new UpdateReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRewardWithDefaults

`func NewUpdateRewardWithDefaults() *UpdateReward`

NewUpdateRewardWithDefaults instantiates a new UpdateReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateReward) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateReward) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateReward) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateReward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateReward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateReward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateReward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateReward) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateReward) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateReward) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVisibilityConditions

`func (o *UpdateReward) GetVisibilityConditions() Rule`

GetVisibilityConditions returns the VisibilityConditions field if non-nil, zero value otherwise.

### GetVisibilityConditionsOk

`func (o *UpdateReward) GetVisibilityConditionsOk() (*Rule, bool)`

GetVisibilityConditionsOk returns a tuple with the VisibilityConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibilityConditions

`func (o *UpdateReward) SetVisibilityConditions(v Rule)`

SetVisibilityConditions sets VisibilityConditions field to given value.

### HasVisibilityConditions

`func (o *UpdateReward) HasVisibilityConditions() bool`

HasVisibilityConditions returns a boolean if a field has been set.

### GetRule

`func (o *UpdateReward) GetRule() Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UpdateReward) GetRuleOk() (*Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UpdateReward) SetRule(v Rule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *UpdateReward) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetBindings

`func (o *UpdateReward) GetBindings() []Binding`

GetBindings returns the Bindings field if non-nil, zero value otherwise.

### GetBindingsOk

`func (o *UpdateReward) GetBindingsOk() (*[]Binding, bool)`

GetBindingsOk returns a tuple with the Bindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindings

`func (o *UpdateReward) SetBindings(v []Binding)`

SetBindings sets Bindings field to given value.

### HasBindings

`func (o *UpdateReward) HasBindings() bool`

HasBindings returns a boolean if a field has been set.

### GetPointsRequired

`func (o *UpdateReward) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *UpdateReward) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *UpdateReward) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *UpdateReward) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


