# PromotionCheckAttributeBlock

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
**OnFailure** | Pointer to [**[]PromotionBlock**](PromotionBlock.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewPromotionCheckAttributeBlock

`func NewPromotionCheckAttributeBlock(id string, type_ string, operator string, attribute interface{}, ) *PromotionCheckAttributeBlock`

NewPromotionCheckAttributeBlock instantiates a new PromotionCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionCheckAttributeBlockWithDefaults

`func NewPromotionCheckAttributeBlockWithDefaults() *PromotionCheckAttributeBlock`

NewPromotionCheckAttributeBlockWithDefaults instantiates a new PromotionCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionCheckAttributeBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionCheckAttributeBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionCheckAttributeBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromotionCheckAttributeBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionCheckAttributeBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionCheckAttributeBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *PromotionCheckAttributeBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromotionCheckAttributeBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromotionCheckAttributeBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromotionCheckAttributeBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *PromotionCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PromotionCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PromotionCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *PromotionCheckAttributeBlock) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *PromotionCheckAttributeBlock) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *PromotionCheckAttributeBlock) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.


### SetAttributeNil

`func (o *PromotionCheckAttributeBlock) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *PromotionCheckAttributeBlock) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *PromotionCheckAttributeBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PromotionCheckAttributeBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PromotionCheckAttributeBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PromotionCheckAttributeBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PromotionCheckAttributeBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PromotionCheckAttributeBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *PromotionCheckAttributeBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PromotionCheckAttributeBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PromotionCheckAttributeBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *PromotionCheckAttributeBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *PromotionCheckAttributeBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *PromotionCheckAttributeBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *PromotionCheckAttributeBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PromotionCheckAttributeBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PromotionCheckAttributeBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *PromotionCheckAttributeBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *PromotionCheckAttributeBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *PromotionCheckAttributeBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStart

`func (o *PromotionCheckAttributeBlock) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PromotionCheckAttributeBlock) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PromotionCheckAttributeBlock) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *PromotionCheckAttributeBlock) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *PromotionCheckAttributeBlock) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *PromotionCheckAttributeBlock) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *PromotionCheckAttributeBlock) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PromotionCheckAttributeBlock) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PromotionCheckAttributeBlock) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *PromotionCheckAttributeBlock) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *PromotionCheckAttributeBlock) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *PromotionCheckAttributeBlock) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *PromotionCheckAttributeBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *PromotionCheckAttributeBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *PromotionCheckAttributeBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *PromotionCheckAttributeBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *PromotionCheckAttributeBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *PromotionCheckAttributeBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *PromotionCheckAttributeBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *PromotionCheckAttributeBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *PromotionCheckAttributeBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *PromotionCheckAttributeBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *PromotionCheckAttributeBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *PromotionCheckAttributeBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *PromotionCheckAttributeBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PromotionCheckAttributeBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PromotionCheckAttributeBlock) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *PromotionCheckAttributeBlock) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *PromotionCheckAttributeBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *PromotionCheckAttributeBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *PromotionCheckAttributeBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PromotionCheckAttributeBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PromotionCheckAttributeBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *PromotionCheckAttributeBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PromotionCheckAttributeBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PromotionCheckAttributeBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOnFailure

`func (o *PromotionCheckAttributeBlock) GetOnFailure() []PromotionBlock`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *PromotionCheckAttributeBlock) GetOnFailureOk() (*[]PromotionBlock, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *PromotionCheckAttributeBlock) SetOnFailure(v []PromotionBlock)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *PromotionCheckAttributeBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


