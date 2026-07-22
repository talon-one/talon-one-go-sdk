# CatalogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **interface{}** |  | 
**Payload** | [**AddPriceAdjustmentCatalogAction**](AddPriceAdjustmentCatalogAction.md) |  | 

## Methods

### NewCatalogAction

`func NewCatalogAction(type_ interface{}, payload AddPriceAdjustmentCatalogAction, ) *CatalogAction`

NewCatalogAction instantiates a new CatalogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogActionWithDefaults

`func NewCatalogActionWithDefaults() *CatalogAction`

NewCatalogActionWithDefaults instantiates a new CatalogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CatalogAction) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogAction) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogAction) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *CatalogAction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CatalogAction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPayload

`func (o *CatalogAction) GetPayload() AddPriceAdjustmentCatalogAction`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *CatalogAction) GetPayloadOk() (*AddPriceAdjustmentCatalogAction, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *CatalogAction) SetPayload(v AddPriceAdjustmentCatalogAction)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


