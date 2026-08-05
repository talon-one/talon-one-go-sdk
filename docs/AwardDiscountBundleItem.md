# AwardDiscountBundleItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A bundle-item selector of type &#x60;byIndex&#x60;. | 
**Value** | **int64** | The zero-based index of the slot within the bundle. | 
**Attribute** | **string** | A per-item attribute expression used to rank bundle items. | 
**Direction** | **string** | Ranking direction. &#x60;highest&#x60; picks the item with the largest attribute value, &#x60;lowest&#x60; the smallest. | 

## Methods

### NewAwardDiscountBundleItem

`func NewAwardDiscountBundleItem(type_ string, value int64, attribute string, direction string, ) *AwardDiscountBundleItem`

NewAwardDiscountBundleItem instantiates a new AwardDiscountBundleItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountBundleItemWithDefaults

`func NewAwardDiscountBundleItemWithDefaults() *AwardDiscountBundleItem`

NewAwardDiscountBundleItemWithDefaults instantiates a new AwardDiscountBundleItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountBundleItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountBundleItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountBundleItem) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *AwardDiscountBundleItem) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AwardDiscountBundleItem) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AwardDiscountBundleItem) SetValue(v int64)`

SetValue sets Value field to given value.


### GetAttribute

`func (o *AwardDiscountBundleItem) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AwardDiscountBundleItem) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AwardDiscountBundleItem) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetDirection

`func (o *AwardDiscountBundleItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *AwardDiscountBundleItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *AwardDiscountBundleItem) SetDirection(v string)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


