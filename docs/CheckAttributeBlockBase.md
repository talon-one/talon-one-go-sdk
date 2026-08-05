# CheckAttributeBlockBase

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

### NewCheckAttributeBlockBase

`func NewCheckAttributeBlockBase(id string, type_ string, operator string, attribute interface{}, ) *CheckAttributeBlockBase`

NewCheckAttributeBlockBase instantiates a new CheckAttributeBlockBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAttributeBlockBaseWithDefaults

`func NewCheckAttributeBlockBaseWithDefaults() *CheckAttributeBlockBase`

NewCheckAttributeBlockBaseWithDefaults instantiates a new CheckAttributeBlockBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAttributeBlockBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAttributeBlockBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAttributeBlockBase) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CheckAttributeBlockBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckAttributeBlockBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckAttributeBlockBase) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckAttributeBlockBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckAttributeBlockBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckAttributeBlockBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckAttributeBlockBase) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckAttributeBlockBase) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckAttributeBlockBase) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckAttributeBlockBase) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *CheckAttributeBlockBase) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *CheckAttributeBlockBase) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *CheckAttributeBlockBase) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.


### SetAttributeNil

`func (o *CheckAttributeBlockBase) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *CheckAttributeBlockBase) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *CheckAttributeBlockBase) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CheckAttributeBlockBase) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CheckAttributeBlockBase) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CheckAttributeBlockBase) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CheckAttributeBlockBase) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CheckAttributeBlockBase) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *CheckAttributeBlockBase) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *CheckAttributeBlockBase) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *CheckAttributeBlockBase) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *CheckAttributeBlockBase) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *CheckAttributeBlockBase) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *CheckAttributeBlockBase) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *CheckAttributeBlockBase) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *CheckAttributeBlockBase) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *CheckAttributeBlockBase) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *CheckAttributeBlockBase) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *CheckAttributeBlockBase) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *CheckAttributeBlockBase) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStart

`func (o *CheckAttributeBlockBase) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CheckAttributeBlockBase) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CheckAttributeBlockBase) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *CheckAttributeBlockBase) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *CheckAttributeBlockBase) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *CheckAttributeBlockBase) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *CheckAttributeBlockBase) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *CheckAttributeBlockBase) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *CheckAttributeBlockBase) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *CheckAttributeBlockBase) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *CheckAttributeBlockBase) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *CheckAttributeBlockBase) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *CheckAttributeBlockBase) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *CheckAttributeBlockBase) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *CheckAttributeBlockBase) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *CheckAttributeBlockBase) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *CheckAttributeBlockBase) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *CheckAttributeBlockBase) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *CheckAttributeBlockBase) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *CheckAttributeBlockBase) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *CheckAttributeBlockBase) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *CheckAttributeBlockBase) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *CheckAttributeBlockBase) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *CheckAttributeBlockBase) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *CheckAttributeBlockBase) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CheckAttributeBlockBase) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CheckAttributeBlockBase) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *CheckAttributeBlockBase) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *CheckAttributeBlockBase) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *CheckAttributeBlockBase) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *CheckAttributeBlockBase) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CheckAttributeBlockBase) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CheckAttributeBlockBase) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *CheckAttributeBlockBase) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *CheckAttributeBlockBase) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *CheckAttributeBlockBase) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


