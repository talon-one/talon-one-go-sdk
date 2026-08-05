# AwardDiscountTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A target discriminator of type &#x60;cart&#x60;. | 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the bundle&#39;s items. | [optional] 
**Name** | **string** | Name of the bundle binding the discount targets. | 
**Item** | Pointer to [**AwardDiscountBundleItem**](AwardDiscountBundleItem.md) |  | [optional] 
**AdditionalCost** | [**AdditionalCostReference**](AdditionalCostReference.md) |  | 
**Target** | [**AwardDiscountAdditionalCostTargetTarget**](AwardDiscountAdditionalCostTargetTarget.md) |  | 

## Methods

### NewAwardDiscountTarget

`func NewAwardDiscountTarget(type_ string, name string, additionalCost AdditionalCostReference, target AwardDiscountAdditionalCostTargetTarget, ) *AwardDiscountTarget`

NewAwardDiscountTarget instantiates a new AwardDiscountTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountTargetWithDefaults

`func NewAwardDiscountTargetWithDefaults() *AwardDiscountTarget`

NewAwardDiscountTargetWithDefaults instantiates a new AwardDiscountTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountTarget) SetType(v string)`

SetType sets Type field to given value.


### GetProrated

`func (o *AwardDiscountTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.

### GetName

`func (o *AwardDiscountTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountTarget) SetName(v string)`

SetName sets Name field to given value.


### GetItem

`func (o *AwardDiscountTarget) GetItem() AwardDiscountBundleItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AwardDiscountTarget) GetItemOk() (*AwardDiscountBundleItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AwardDiscountTarget) SetItem(v AwardDiscountBundleItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *AwardDiscountTarget) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetAdditionalCost

`func (o *AwardDiscountTarget) GetAdditionalCost() AdditionalCostReference`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *AwardDiscountTarget) GetAdditionalCostOk() (*AdditionalCostReference, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *AwardDiscountTarget) SetAdditionalCost(v AdditionalCostReference)`

SetAdditionalCost sets AdditionalCost field to given value.


### GetTarget

`func (o *AwardDiscountTarget) GetTarget() AwardDiscountAdditionalCostTargetTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AwardDiscountTarget) GetTargetOk() (*AwardDiscountAdditionalCostTargetTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AwardDiscountTarget) SetTarget(v AwardDiscountAdditionalCostTargetTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


