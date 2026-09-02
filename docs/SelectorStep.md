# SelectorStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A step discriminator of type &#x60;filter&#x60;. | 
**Predicate** | [**Block**](Block.md) |  | 
**Fields** | [**[]SortSelectorStepField**](SortSelectorStepField.md) | One or more fields to sort by, applied in order. Each field has its own direction. | 
**Operator** | **string** | The aggregation operator applied to the items produced by the preceding step: - &#x60;max&#x60;, &#x60;min&#x60;, and &#x60;sum&#x60; operate on numeric values. - &#x60;count&#x60; returns the number of items. - &#x60;empty&#x60; reports whether the list is empty.  | 
**From** | Pointer to [**SelectSelectorStepFrom**](SelectSelectorStepFrom.md) |  | [optional] 
**To** | Pointer to **int32** | The end index for the &#x60;between&#x60; operator. The item at this index is not included. | [optional] 
**Count** | Pointer to **int32** | The maximum number of items to select for the &#x60;many&#x60; operator. | [optional] 
**Index** | Pointer to **int32** | The exact position of the item to select for the &#x60;one&#x60; operator. | [optional] 
**Partial** | Pointer to **bool** | Indicates if the step returns fewer items than requested when the source list is shorter than the range needs. Always &#x60;true&#x60; for the &#x60;many&#x60; and &#x60;between&#x60; operators; not present for &#x60;one&#x60;, which fails instead of returning a partial result. | [optional] 
**Expression** | **string** | The attribute path each item is mapped to. | 
**ValueMap** | [**SelectorValueMapRef**](SelectorValueMapRef.md) |  | 

## Methods

### NewSelectorStep

`func NewSelectorStep(type_ string, predicate Block, fields []SortSelectorStepField, operator string, expression string, valueMap SelectorValueMapRef, ) *SelectorStep`

NewSelectorStep instantiates a new SelectorStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorStepWithDefaults

`func NewSelectorStepWithDefaults() *SelectorStep`

NewSelectorStepWithDefaults instantiates a new SelectorStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SelectorStep) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectorStep) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectorStep) SetType(v string)`

SetType sets Type field to given value.


### GetPredicate

`func (o *SelectorStep) GetPredicate() Block`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *SelectorStep) GetPredicateOk() (*Block, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *SelectorStep) SetPredicate(v Block)`

SetPredicate sets Predicate field to given value.


### GetFields

`func (o *SelectorStep) GetFields() []SortSelectorStepField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *SelectorStep) GetFieldsOk() (*[]SortSelectorStepField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *SelectorStep) SetFields(v []SortSelectorStepField)`

SetFields sets Fields field to given value.


### GetOperator

`func (o *SelectorStep) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SelectorStep) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SelectorStep) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetFrom

`func (o *SelectorStep) GetFrom() SelectSelectorStepFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SelectorStep) GetFromOk() (*SelectSelectorStepFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SelectorStep) SetFrom(v SelectSelectorStepFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *SelectorStep) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *SelectorStep) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SelectorStep) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SelectorStep) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *SelectorStep) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetCount

`func (o *SelectorStep) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SelectorStep) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SelectorStep) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SelectorStep) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetIndex

`func (o *SelectorStep) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *SelectorStep) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *SelectorStep) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *SelectorStep) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetPartial

`func (o *SelectorStep) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *SelectorStep) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *SelectorStep) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *SelectorStep) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetExpression

`func (o *SelectorStep) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SelectorStep) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SelectorStep) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetValueMap

`func (o *SelectorStep) GetValueMap() SelectorValueMapRef`

GetValueMap returns the ValueMap field if non-nil, zero value otherwise.

### GetValueMapOk

`func (o *SelectorStep) GetValueMapOk() (*SelectorValueMapRef, bool)`

GetValueMapOk returns a tuple with the ValueMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMap

`func (o *SelectorStep) SetValueMap(v SelectorValueMapRef)`

SetValueMap sets ValueMap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


