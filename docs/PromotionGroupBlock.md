# PromotionGroupBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | Logical operator applied across child blocks. &#x60;all&#x60; requires every child to pass, &#x60;atLeastOne&#x60; requires at least one, &#x60;none&#x60; requires all to fail. | 
**Blocks** | **[]interface{}** | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to **[]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to **map[string][]interface{}** | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewPromotionGroupBlock

`func NewPromotionGroupBlock(id string, type_ string, operator string, blocks []interface{}, ) *PromotionGroupBlock`

NewPromotionGroupBlock instantiates a new PromotionGroupBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionGroupBlockWithDefaults

`func NewPromotionGroupBlockWithDefaults() *PromotionGroupBlock`

NewPromotionGroupBlockWithDefaults instantiates a new PromotionGroupBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionGroupBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionGroupBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionGroupBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromotionGroupBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionGroupBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionGroupBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *PromotionGroupBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromotionGroupBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromotionGroupBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromotionGroupBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *PromotionGroupBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PromotionGroupBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PromotionGroupBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *PromotionGroupBlock) GetBlocks() []interface{}`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *PromotionGroupBlock) GetBlocksOk() (*[]interface{}, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *PromotionGroupBlock) SetBlocks(v []interface{})`

SetBlocks sets Blocks field to given value.


### GetOnFailure

`func (o *PromotionGroupBlock) GetOnFailure() []interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *PromotionGroupBlock) GetOnFailureOk() (*[]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *PromotionGroupBlock) SetOnFailure(v []interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *PromotionGroupBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *PromotionGroupBlock) GetOnError() map[string][]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *PromotionGroupBlock) GetOnErrorOk() (*map[string][]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *PromotionGroupBlock) SetOnError(v map[string][]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *PromotionGroupBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


