# StrikethroughRuleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the rule. | [optional] 
**ParentId** | Pointer to **string** | ID of the parent rule, if any. | [optional] 
**Title** | **string** | A short description of the rule. | 
**Description** | Pointer to **string** | A longer description of the rule. | [optional] 
**Blocks** | **[]interface{}** | The condition and effect blocks that make up this strikethrough rule. | 

## Methods

### NewStrikethroughRuleV2

`func NewStrikethroughRuleV2(title string, blocks []interface{}, ) *StrikethroughRuleV2`

NewStrikethroughRuleV2 instantiates a new StrikethroughRuleV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughRuleV2WithDefaults

`func NewStrikethroughRuleV2WithDefaults() *StrikethroughRuleV2`

NewStrikethroughRuleV2WithDefaults instantiates a new StrikethroughRuleV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrikethroughRuleV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughRuleV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrikethroughRuleV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StrikethroughRuleV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *StrikethroughRuleV2) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *StrikethroughRuleV2) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *StrikethroughRuleV2) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *StrikethroughRuleV2) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *StrikethroughRuleV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StrikethroughRuleV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StrikethroughRuleV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *StrikethroughRuleV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StrikethroughRuleV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StrikethroughRuleV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StrikethroughRuleV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBlocks

`func (o *StrikethroughRuleV2) GetBlocks() []interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *StrikethroughRuleV2) GetBlocksOk() (*[]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *StrikethroughRuleV2) SetBlocks(v []interface{})`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


