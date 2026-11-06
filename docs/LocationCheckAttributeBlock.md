# LocationCheckAttributeBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | **string** | A block discriminator of type &#x60;checkAttribute&#x60;. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | **string** | The location membership operator applied to the attribute. | 
**Attribute** | **interface{}** | The attribute path identifier (e.g. \&quot;$Session.Total\&quot;). | 
**Value** | Pointer to **interface{}** | The comparison value for scalar operators. | [optional] 
**Min** | Pointer to **interface{}** | The minimum value allowed for the &#x60;between&#x60; operator. | [optional] 
**Max** | Pointer to **interface{}** | The maximum value allowed for the &#x60;between&#x60; operator. | [optional] 
**Start** | Pointer to **interface{}** | The start value for the &#x60;within&#x60; operator. | [optional] 
**End** | Pointer to **interface{}** | The end value for the &#x60;within&#x60; operator. | [optional] 
**StartInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;start&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**EndInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;end&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**TimezoneInsensitive** | Pointer to **bool** | Indicates whether the &#x60;within&#x60; operator ignores time zones and compares the wall-clock time only. When &#x60;false&#x60;, time zones are taken into account. | [optional] 
**Values** | [**LocationCheckAttributeBlock1Values**](LocationCheckAttributeBlock1Values.md) |  | 
**Count** | Pointer to **interface{}** | The count threshold for &#x60;containsAtLeast&#x60; and &#x60;containsExactly&#x60; operators. | [optional] 
**OnFailure** | Pointer to [**[]Block**](Block.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewLocationCheckAttributeBlock

`func NewLocationCheckAttributeBlock(type_ string, operator string, attribute interface{}, values LocationCheckAttributeBlock1Values, ) *LocationCheckAttributeBlock`

NewLocationCheckAttributeBlock instantiates a new LocationCheckAttributeBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationCheckAttributeBlockWithDefaults

`func NewLocationCheckAttributeBlockWithDefaults() *LocationCheckAttributeBlock`

NewLocationCheckAttributeBlockWithDefaults instantiates a new LocationCheckAttributeBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocationCheckAttributeBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationCheckAttributeBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationCheckAttributeBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationCheckAttributeBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *LocationCheckAttributeBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationCheckAttributeBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationCheckAttributeBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *LocationCheckAttributeBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LocationCheckAttributeBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LocationCheckAttributeBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LocationCheckAttributeBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *LocationCheckAttributeBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *LocationCheckAttributeBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *LocationCheckAttributeBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *LocationCheckAttributeBlock) GetAttribute() interface{}`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *LocationCheckAttributeBlock) GetAttributeOk() (*interface{}, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *LocationCheckAttributeBlock) SetAttribute(v interface{})`

SetAttribute sets Attribute field to given value.


### SetAttributeNil

`func (o *LocationCheckAttributeBlock) SetAttributeNil(b bool)`

 SetAttributeNil sets the value for Attribute to be an explicit nil

### UnsetAttribute
`func (o *LocationCheckAttributeBlock) UnsetAttribute()`

UnsetAttribute ensures that no value is present for Attribute, not even an explicit nil
### GetValue

`func (o *LocationCheckAttributeBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LocationCheckAttributeBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LocationCheckAttributeBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *LocationCheckAttributeBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *LocationCheckAttributeBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LocationCheckAttributeBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *LocationCheckAttributeBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *LocationCheckAttributeBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *LocationCheckAttributeBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *LocationCheckAttributeBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *LocationCheckAttributeBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *LocationCheckAttributeBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *LocationCheckAttributeBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *LocationCheckAttributeBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *LocationCheckAttributeBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *LocationCheckAttributeBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *LocationCheckAttributeBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *LocationCheckAttributeBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStart

`func (o *LocationCheckAttributeBlock) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LocationCheckAttributeBlock) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LocationCheckAttributeBlock) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *LocationCheckAttributeBlock) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *LocationCheckAttributeBlock) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LocationCheckAttributeBlock) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *LocationCheckAttributeBlock) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *LocationCheckAttributeBlock) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *LocationCheckAttributeBlock) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *LocationCheckAttributeBlock) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *LocationCheckAttributeBlock) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *LocationCheckAttributeBlock) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *LocationCheckAttributeBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *LocationCheckAttributeBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *LocationCheckAttributeBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *LocationCheckAttributeBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *LocationCheckAttributeBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *LocationCheckAttributeBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *LocationCheckAttributeBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *LocationCheckAttributeBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *LocationCheckAttributeBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *LocationCheckAttributeBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *LocationCheckAttributeBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *LocationCheckAttributeBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

### GetValues

`func (o *LocationCheckAttributeBlock) GetValues() LocationCheckAttributeBlock1Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *LocationCheckAttributeBlock) GetValuesOk() (*LocationCheckAttributeBlock1Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *LocationCheckAttributeBlock) SetValues(v LocationCheckAttributeBlock1Values)`

SetValues sets Values field to given value.


### GetCount

`func (o *LocationCheckAttributeBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LocationCheckAttributeBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LocationCheckAttributeBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *LocationCheckAttributeBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *LocationCheckAttributeBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *LocationCheckAttributeBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOnFailure

`func (o *LocationCheckAttributeBlock) GetOnFailure() []Block`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *LocationCheckAttributeBlock) GetOnFailureOk() (*[]Block, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *LocationCheckAttributeBlock) SetOnFailure(v []Block)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *LocationCheckAttributeBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


