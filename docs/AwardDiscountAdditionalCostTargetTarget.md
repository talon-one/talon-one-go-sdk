# AwardDiscountAdditionalCostTargetTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A target discriminator of type &#x60;cart&#x60;. | 
**Prorated** | Pointer to **bool** | Whether to distribute the discount proportionally across the selected items. | [optional] 
**Name** | **string** | The name of the selector binding the discount targets. | 

## Methods

### NewAwardDiscountAdditionalCostTargetTarget

`func NewAwardDiscountAdditionalCostTargetTarget(type_ string, name string, ) *AwardDiscountAdditionalCostTargetTarget`

NewAwardDiscountAdditionalCostTargetTarget instantiates a new AwardDiscountAdditionalCostTargetTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardDiscountAdditionalCostTargetTargetWithDefaults

`func NewAwardDiscountAdditionalCostTargetTargetWithDefaults() *AwardDiscountAdditionalCostTargetTarget`

NewAwardDiscountAdditionalCostTargetTargetWithDefaults instantiates a new AwardDiscountAdditionalCostTargetTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AwardDiscountAdditionalCostTargetTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardDiscountAdditionalCostTargetTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardDiscountAdditionalCostTargetTarget) SetType(v string)`

SetType sets Type field to given value.


### GetProrated

`func (o *AwardDiscountAdditionalCostTargetTarget) GetProrated() bool`

GetProrated returns the Prorated field if non-nil, zero value otherwise.

### GetProratedOk

`func (o *AwardDiscountAdditionalCostTargetTarget) GetProratedOk() (*bool, bool)`

GetProratedOk returns a tuple with the Prorated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrated

`func (o *AwardDiscountAdditionalCostTargetTarget) SetProrated(v bool)`

SetProrated sets Prorated field to given value.

### HasProrated

`func (o *AwardDiscountAdditionalCostTargetTarget) HasProrated() bool`

HasProrated returns a boolean if a field has been set.

### GetName

`func (o *AwardDiscountAdditionalCostTargetTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardDiscountAdditionalCostTargetTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardDiscountAdditionalCostTargetTarget) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


