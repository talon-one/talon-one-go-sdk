# CheckAttributeBlock

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

## Methods

### NewCheckAttributeBlock

`func NewCheckAttributeBlock(id string, type_ string, operator string, attribute interface{}, ) *CheckAttributeBlock`

NewCheckAttributeBlock instantiates a new CheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAttributeBlockWithDefaults

`func NewCheckAttributeBlockWithDefaults() *CheckAttributeBlock`

NewCheckAttributeBlockWithDefaults instantiates a new CheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAttributeBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAttributeBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAttributeBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CheckAttributeBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckAttributeBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckAttributeBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckAttributeBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckAttributeBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckAttributeBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckAttributeBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *CheckAttributeBlock) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *CheckAttributeBlock) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *CheckAttributeBlock) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.


### SetAttributeNil

`func (o *CheckAttributeBlock) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *CheckAttributeBlock) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *CheckAttributeBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckAttributeBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckAttributeBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CheckAttributeBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CheckAttributeBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CheckAttributeBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *CheckAttributeBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *CheckAttributeBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *CheckAttributeBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *CheckAttributeBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *CheckAttributeBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *CheckAttributeBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *CheckAttributeBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *CheckAttributeBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *CheckAttributeBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *CheckAttributeBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *CheckAttributeBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *CheckAttributeBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStart

`func (o *CheckAttributeBlock) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CheckAttributeBlock) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CheckAttributeBlock) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *CheckAttributeBlock) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *CheckAttributeBlock) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *CheckAttributeBlock) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *CheckAttributeBlock) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CheckAttributeBlock) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CheckAttributeBlock) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *CheckAttributeBlock) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *CheckAttributeBlock) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *CheckAttributeBlock) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *CheckAttributeBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *CheckAttributeBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *CheckAttributeBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *CheckAttributeBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *CheckAttributeBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *CheckAttributeBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *CheckAttributeBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *CheckAttributeBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *CheckAttributeBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *CheckAttributeBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *CheckAttributeBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *CheckAttributeBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *CheckAttributeBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CheckAttributeBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CheckAttributeBlock) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *CheckAttributeBlock) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *CheckAttributeBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *CheckAttributeBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *CheckAttributeBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CheckAttributeBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CheckAttributeBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *CheckAttributeBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CheckAttributeBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CheckAttributeBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


