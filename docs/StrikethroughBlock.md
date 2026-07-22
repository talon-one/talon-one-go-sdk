# StrikethroughBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | The comparison operator applied to the attribute. | 
**Blocks** | [**[]StrikethroughBlock**](StrikethroughBlock.md) | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to [**[]StrikethroughBlock**](StrikethroughBlock.md) | Strikethrough blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]StrikethroughBlock**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 
**Expression** | **[]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 
**Attribute** | **string** | The attribute path identifier (e.g. \&quot;$Session.Total\&quot;). | 
**Value** | Pointer to **interface{}** |  | [optional] 
**Min** | Pointer to **interface{}** |  | [optional] 
**Max** | Pointer to **interface{}** |  | [optional] 
**Values** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewStrikethroughBlock

`func NewStrikethroughBlock(id string, type_ string, operator string, blocks []StrikethroughBlock, expression []interface{}, attribute string, ) *StrikethroughBlock`

NewStrikethroughBlock instantiates a new StrikethroughBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughBlockWithDefaults

`func NewStrikethroughBlockWithDefaults() *StrikethroughBlock`

NewStrikethroughBlockWithDefaults instantiates a new StrikethroughBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrikethroughBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrikethroughBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *StrikethroughBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StrikethroughBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StrikethroughBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *StrikethroughBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StrikethroughBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StrikethroughBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StrikethroughBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *StrikethroughBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StrikethroughBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StrikethroughBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *StrikethroughBlock) GetBlocks() []StrikethroughBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *StrikethroughBlock) GetBlocksOk() (*[]StrikethroughBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *StrikethroughBlock) SetBlocks(v []StrikethroughBlock)`

SetBlocks sets Blocks field to given value.


### GetOnFailure

`func (o *StrikethroughBlock) GetOnFailure() []StrikethroughBlock`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *StrikethroughBlock) GetOnFailureOk() (*[]StrikethroughBlock, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *StrikethroughBlock) SetOnFailure(v []StrikethroughBlock)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *StrikethroughBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *StrikethroughBlock) GetOnError() map[string][]StrikethroughBlock`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *StrikethroughBlock) GetOnErrorOk() (*map[string][]StrikethroughBlock, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *StrikethroughBlock) SetOnError(v map[string][]StrikethroughBlock)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *StrikethroughBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetExpression

`func (o *StrikethroughBlock) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *StrikethroughBlock) GetExpressionOk() (*[]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *StrikethroughBlock) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.


### GetAttribute

`func (o *StrikethroughBlock) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *StrikethroughBlock) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *StrikethroughBlock) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetValue

`func (o *StrikethroughBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StrikethroughBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StrikethroughBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *StrikethroughBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StrikethroughBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StrikethroughBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *StrikethroughBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *StrikethroughBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *StrikethroughBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *StrikethroughBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *StrikethroughBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *StrikethroughBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *StrikethroughBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *StrikethroughBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *StrikethroughBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *StrikethroughBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *StrikethroughBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *StrikethroughBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetValues

`func (o *StrikethroughBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StrikethroughBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StrikethroughBlock) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *StrikethroughBlock) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *StrikethroughBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *StrikethroughBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *StrikethroughBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StrikethroughBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StrikethroughBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *StrikethroughBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *StrikethroughBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *StrikethroughBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


