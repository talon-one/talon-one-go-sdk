# SelectorBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | Logical operator applied across child blocks. &#x60;all&#x60; requires every child to pass, &#x60;atLeastOne&#x60; requires at least one, &#x60;none&#x60; requires all to fail. | 
**Blocks** | [**[]SelectorBlock**](SelectorBlock.md) | Child predicate blocks evaluated according to the operator. | 
**Expression** | **[]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 

## Methods

### NewSelectorBlock

`func NewSelectorBlock(id string, type_ string, operator string, blocks []SelectorBlock, expression []interface{}, ) *SelectorBlock`

NewSelectorBlock instantiates a new SelectorBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorBlockWithDefaults

`func NewSelectorBlockWithDefaults() *SelectorBlock`

NewSelectorBlockWithDefaults instantiates a new SelectorBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectorBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectorBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectorBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SelectorBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectorBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectorBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *SelectorBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SelectorBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SelectorBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SelectorBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *SelectorBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SelectorBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SelectorBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *SelectorBlock) GetBlocks() []SelectorBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SelectorBlock) GetBlocksOk() (*[]SelectorBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SelectorBlock) SetBlocks(v []SelectorBlock)`

SetBlocks sets Blocks field to given value.


### GetExpression

`func (o *SelectorBlock) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SelectorBlock) GetExpressionOk() (*[]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SelectorBlock) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


