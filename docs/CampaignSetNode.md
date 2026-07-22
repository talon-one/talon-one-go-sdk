# CampaignSetNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Name** | **string** | Name of the set. | 
**Operator** | **string** | An indicator of how the set operates on its elements. | 
**Elements** | [**[]CampaignSetNode**](CampaignSetNode.md) | Child elements of this set. | 
**GroupId** | **int64** | The ID of the campaign set. | 
**Locked** | **bool** | An indicator of whether the campaign set is locked for modification. | 
**Description** | Pointer to **string** | A description of the campaign set. | [optional] 
**EvaluationMode** | **string** | The mode by which campaigns in the campaign evaluation group are evaluated. | 
**EvaluationScope** | **string** | The evaluation scope of the campaign evaluation group. | 
**CampaignId** | **int64** | ID of the campaign | 

## Methods

### NewCampaignSetNode

`func NewCampaignSetNode(type_ string, name string, operator string, elements []CampaignSetNode, groupId int64, locked bool, evaluationMode string, evaluationScope string, campaignId int64, ) *CampaignSetNode`

NewCampaignSetNode instantiates a new CampaignSetNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSetNodeWithDefaults

`func NewCampaignSetNodeWithDefaults() *CampaignSetNode`

NewCampaignSetNodeWithDefaults instantiates a new CampaignSetNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CampaignSetNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSetNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSetNode) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CampaignSetNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSetNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignSetNode) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *CampaignSetNode) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CampaignSetNode) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CampaignSetNode) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetElements

`func (o *CampaignSetNode) GetElements() []CampaignSetNode`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *CampaignSetNode) GetElementsOk() (*[]CampaignSetNode, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *CampaignSetNode) SetElements(v []CampaignSetNode)`

SetElements sets Elements field to given value.


### GetGroupId

`func (o *CampaignSetNode) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *CampaignSetNode) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *CampaignSetNode) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetLocked

`func (o *CampaignSetNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CampaignSetNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CampaignSetNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetDescription

`func (o *CampaignSetNode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignSetNode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignSetNode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignSetNode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEvaluationMode

`func (o *CampaignSetNode) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *CampaignSetNode) GetEvaluationModeOk() (*string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMode

`func (o *CampaignSetNode) SetEvaluationMode(v string)`

SetEvaluationMode sets EvaluationMode field to given value.


### GetEvaluationScope

`func (o *CampaignSetNode) GetEvaluationScope() string`

GetEvaluationScope returns the EvaluationScope field if non-nil, zero value otherwise.

### GetEvaluationScopeOk

`func (o *CampaignSetNode) GetEvaluationScopeOk() (*string, bool)`

GetEvaluationScopeOk returns a tuple with the EvaluationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationScope

`func (o *CampaignSetNode) SetEvaluationScope(v string)`

SetEvaluationScope sets EvaluationScope field to given value.


### GetCampaignId

`func (o *CampaignSetNode) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignSetNode) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignSetNode) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


