# PassthroughBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | The type discriminator for this block. | 
**Expression** | **[]interface{}** | The raw Talang expression as an array. The first element is the function name; subsequent elements are its arguments, which may themselves be nested expressions. | 

## Methods

### NewPassthroughBlock

`func NewPassthroughBlock(id string, type_ string, expression []interface{}, ) *PassthroughBlock`

NewPassthroughBlock instantiates a new PassthroughBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPassthroughBlockWithDefaults

`func NewPassthroughBlockWithDefaults() *PassthroughBlock`

NewPassthroughBlockWithDefaults instantiates a new PassthroughBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PassthroughBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PassthroughBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PassthroughBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PassthroughBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PassthroughBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PassthroughBlock) SetType(v string)`

SetType sets Type field to given value.


### GetExpression

`func (o *PassthroughBlock) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PassthroughBlock) GetExpressionOk() (*[]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PassthroughBlock) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


