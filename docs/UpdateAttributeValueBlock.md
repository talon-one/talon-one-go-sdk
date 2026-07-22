# UpdateAttributeValueBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | The update operation applied to the attribute. | 
**Attribute** | [**UpdateAttributeValueBlock1Attribute**](UpdateAttributeValueBlock1Attribute.md) |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**Target** | [**UpdateAttributeValueBlock1Target**](UpdateAttributeValueBlock1Target.md) |  | 

## Methods

### NewUpdateAttributeValueBlock

`func NewUpdateAttributeValueBlock(id string, type_ string, operator string, attribute UpdateAttributeValueBlock1Attribute, target UpdateAttributeValueBlock1Target, ) *UpdateAttributeValueBlock`

NewUpdateAttributeValueBlock instantiates a new UpdateAttributeValueBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAttributeValueBlockWithDefaults

`func NewUpdateAttributeValueBlockWithDefaults() *UpdateAttributeValueBlock`

NewUpdateAttributeValueBlockWithDefaults instantiates a new UpdateAttributeValueBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAttributeValueBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAttributeValueBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAttributeValueBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UpdateAttributeValueBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAttributeValueBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAttributeValueBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *UpdateAttributeValueBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateAttributeValueBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateAttributeValueBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateAttributeValueBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateAttributeValueBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateAttributeValueBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateAttributeValueBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetAttribute

`func (o *UpdateAttributeValueBlock) GetAttribute() UpdateAttributeValueBlock1Attribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *UpdateAttributeValueBlock) GetAttributeOk() (*UpdateAttributeValueBlock1Attribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *UpdateAttributeValueBlock) SetAttribute(v UpdateAttributeValueBlock1Attribute)`

SetAttribute sets Attribute field to given value.


### GetValue

`func (o *UpdateAttributeValueBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateAttributeValueBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateAttributeValueBlock) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateAttributeValueBlock) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *UpdateAttributeValueBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *UpdateAttributeValueBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetTarget

`func (o *UpdateAttributeValueBlock) GetTarget() UpdateAttributeValueBlock1Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UpdateAttributeValueBlock) GetTargetOk() (*UpdateAttributeValueBlock1Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UpdateAttributeValueBlock) SetTarget(v UpdateAttributeValueBlock1Target)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


