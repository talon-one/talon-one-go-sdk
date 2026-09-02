# ListCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **string** | The list membership operator applied to the attribute. | [optional] 
**Values** | **interface{}** | The set of values to match against. | 

## Methods

### NewListCheckAttributeBlock

`func NewListCheckAttributeBlock(values interface{}, ) *ListCheckAttributeBlock`

NewListCheckAttributeBlock instantiates a new ListCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCheckAttributeBlockWithDefaults

`func NewListCheckAttributeBlockWithDefaults() *ListCheckAttributeBlock`

NewListCheckAttributeBlockWithDefaults instantiates a new ListCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *ListCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ListCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ListCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ListCheckAttributeBlock) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *ListCheckAttributeBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ListCheckAttributeBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ListCheckAttributeBlock) SetValues(v interface{})`

SetValues sets Values field to given value.


### SetValuesNil

`func (o *ListCheckAttributeBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *ListCheckAttributeBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


