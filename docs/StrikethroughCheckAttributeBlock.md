# StrikethroughCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | The comparison operator applied to the attribute. | 
**Attribute** | **interface{}** |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**Min** | Pointer to **interface{}** |  | [optional] 
**Max** | Pointer to **interface{}** |  | [optional] 
**Start** | Pointer to **interface{}** |  | [optional] 
**End** | Pointer to **interface{}** |  | [optional] 
**StartInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;start&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**EndInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;end&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**TimezoneInsensitive** | Pointer to **bool** | Indicates whether the &#x60;within&#x60; operator ignores time zones and compares the wall-clock time only. When &#x60;false&#x60;, time zones are taken into account. | [optional] 
**Values** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **interface{}** |  | [optional] 
**OnFailure** | Pointer to [**[]StrikethroughBlock**](StrikethroughBlock.md) | Strikethrough blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewStrikethroughCheckAttributeBlock

`func NewStrikethroughCheckAttributeBlock(id string, type_ string, operator string, attribute interface{}, ) *StrikethroughCheckAttributeBlock`

NewStrikethroughCheckAttributeBlock instantiates a new StrikethroughCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughCheckAttributeBlockWithDefaults

`func NewStrikethroughCheckAttributeBlockWithDefaults() *StrikethroughCheckAttributeBlock`

NewStrikethroughCheckAttributeBlockWithDefaults instantiates a new StrikethroughCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrikethroughCheckAttributeBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughCheckAttributeBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrikethroughCheckAttributeBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *StrikethroughCheckAttributeBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StrikethroughCheckAttributeBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StrikethroughCheckAttributeBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *StrikethroughCheckAttributeBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StrikethroughCheckAttributeBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StrikethroughCheckAttributeBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StrikethroughCheckAttributeBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *StrikethroughCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *StrikethroughCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *StrikethroughCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *StrikethroughCheckAttributeBlock) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *StrikethroughCheckAttributeBlock) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *StrikethroughCheckAttributeBlock) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.


### SetAttributeNil

`func (o *StrikethroughCheckAttributeBlock) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *StrikethroughCheckAttributeBlock) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *StrikethroughCheckAttributeBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StrikethroughCheckAttributeBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StrikethroughCheckAttributeBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *StrikethroughCheckAttributeBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *StrikethroughCheckAttributeBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StrikethroughCheckAttributeBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *StrikethroughCheckAttributeBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *StrikethroughCheckAttributeBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *StrikethroughCheckAttributeBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *StrikethroughCheckAttributeBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *StrikethroughCheckAttributeBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *StrikethroughCheckAttributeBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *StrikethroughCheckAttributeBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *StrikethroughCheckAttributeBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *StrikethroughCheckAttributeBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *StrikethroughCheckAttributeBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *StrikethroughCheckAttributeBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *StrikethroughCheckAttributeBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStart

`func (o *StrikethroughCheckAttributeBlock) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *StrikethroughCheckAttributeBlock) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *StrikethroughCheckAttributeBlock) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *StrikethroughCheckAttributeBlock) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *StrikethroughCheckAttributeBlock) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *StrikethroughCheckAttributeBlock) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *StrikethroughCheckAttributeBlock) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *StrikethroughCheckAttributeBlock) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *StrikethroughCheckAttributeBlock) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *StrikethroughCheckAttributeBlock) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *StrikethroughCheckAttributeBlock) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *StrikethroughCheckAttributeBlock) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *StrikethroughCheckAttributeBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *StrikethroughCheckAttributeBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *StrikethroughCheckAttributeBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *StrikethroughCheckAttributeBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *StrikethroughCheckAttributeBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *StrikethroughCheckAttributeBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *StrikethroughCheckAttributeBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *StrikethroughCheckAttributeBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *StrikethroughCheckAttributeBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *StrikethroughCheckAttributeBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *StrikethroughCheckAttributeBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *StrikethroughCheckAttributeBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *StrikethroughCheckAttributeBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *StrikethroughCheckAttributeBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *StrikethroughCheckAttributeBlock) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *StrikethroughCheckAttributeBlock) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *StrikethroughCheckAttributeBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *StrikethroughCheckAttributeBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *StrikethroughCheckAttributeBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *StrikethroughCheckAttributeBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *StrikethroughCheckAttributeBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *StrikethroughCheckAttributeBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *StrikethroughCheckAttributeBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *StrikethroughCheckAttributeBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOnFailure

`func (o *StrikethroughCheckAttributeBlock) GetOnFailure() []StrikethroughBlock`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *StrikethroughCheckAttributeBlock) GetOnFailureOk() (*[]StrikethroughBlock, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *StrikethroughCheckAttributeBlock) SetOnFailure(v []StrikethroughBlock)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *StrikethroughCheckAttributeBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


