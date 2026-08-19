# SelectorGroupBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | A block discriminator of type &#x60;group&#x60;. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | Logical operator applied across child blocks. &#x60;all&#x60; requires every child to pass, &#x60;atLeastOne&#x60; requires at least one, &#x60;none&#x60; requires all to fail. | 
**Blocks** | [**[]SelectorBlock**](SelectorBlock.md) | Child predicate blocks evaluated according to the operator. | 

## Methods

### NewSelectorGroupBlock

`func NewSelectorGroupBlock(id string, type_ string, operator string, blocks []SelectorBlock, ) *SelectorGroupBlock`

NewSelectorGroupBlock instantiates a new SelectorGroupBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorGroupBlockWithDefaults

`func NewSelectorGroupBlockWithDefaults() *SelectorGroupBlock`

NewSelectorGroupBlockWithDefaults instantiates a new SelectorGroupBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SelectorGroupBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SelectorGroupBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SelectorGroupBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SelectorGroupBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectorGroupBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectorGroupBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *SelectorGroupBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SelectorGroupBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SelectorGroupBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SelectorGroupBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *SelectorGroupBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SelectorGroupBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SelectorGroupBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *SelectorGroupBlock) GetBlocks() []SelectorBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *SelectorGroupBlock) GetBlocksOk() (*[]SelectorBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *SelectorGroupBlock) SetBlocks(v []SelectorBlock)`

SetBlocks sets Blocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


