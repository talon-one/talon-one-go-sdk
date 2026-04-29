# UpdateReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the reward. | 
**Description** | Pointer to **string** | A description of the reward. | [optional] 
**Status** | **string** | The status of the reward. | 
**Rule** | Pointer to [**[]Rule**](Rule.md) | Rule to apply.  **Note**: The &#x60;bindings&#x60; field inside the rule must not be used in this endpoint. All bindings should be defined at the reward level via the top-level &#x60;bindings&#x60; field.  | [optional] 
**Bindings** | Pointer to [**[]Binding**](Binding.md) | A list of named variables created before the reward&#39;s rules are evaluated.  Each binding pairs a name with a talang expression. The expression is evaluated once  and its result is available by name in any rule condition or effect. Bindings must be defined outside of individual rules. | [optional] 

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


### GetRule

`func (o *UpdateReward) GetRule() []Rule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UpdateReward) GetRuleOk() (*[]Rule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UpdateReward) SetRule(v []Rule)`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


