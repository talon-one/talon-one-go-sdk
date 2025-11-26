# ItemAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributeid** | **int64** | The ID of the attribute of the item. | 
**Name** | **string** | The name of the attribute. | 
**Value** | **interface{}** |  | 

## Methods

### NewItemAttribute

`func NewItemAttribute(attributeid int64, name string, value interface{}, ) *ItemAttribute`

NewItemAttribute instantiates a new ItemAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAttributeWithDefaults

`func NewItemAttributeWithDefaults() *ItemAttribute`

NewItemAttributeWithDefaults instantiates a new ItemAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeid

`func (o *ItemAttribute) GetAttributeid() int64`

GetAttributeid returns the Attributeid field if non-nil, zero value otherwise.

### GetAttributeidOk

`func (o *ItemAttribute) GetAttributeidOk() (*int64, bool)`

GetAttributeidOk returns a tuple with the Attributeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeid

`func (o *ItemAttribute) SetAttributeid(v int64)`

SetAttributeid sets Attributeid field to given value.


### GetName

`func (o *ItemAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ItemAttribute) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ItemAttribute) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ItemAttribute) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ItemAttribute) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ItemAttribute) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


