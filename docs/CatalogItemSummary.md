# CatalogItemSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of the catalog item. | 
**CatalogId** | **int64** | The ID of the catalog the item belongs to. | 
**Sku** | **string** | The stock keeping unit of the item. | 
**Price** | Pointer to **float32** | The price of the item. | [optional] 
**ProductName** | Pointer to **string** | The name of the product the item belongs to. | [optional] 
**ProductId** | Pointer to **int64** | The ID of the product the item is linked to.  **Note:** Only returned when the product catalog integration is enabled.  | [optional] 
**SkuId** | Pointer to **int64** | The ID of the SKU the item is linked to.  **Note:** Only returned when the product catalog integration is enabled.  | [optional] 

## Methods

### NewCatalogItemSummary

`func NewCatalogItemSummary(id int64, catalogId int64, sku string, ) *CatalogItemSummary`

NewCatalogItemSummary instantiates a new CatalogItemSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemSummaryWithDefaults

`func NewCatalogItemSummaryWithDefaults() *CatalogItemSummary`

NewCatalogItemSummaryWithDefaults instantiates a new CatalogItemSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogItemSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItemSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItemSummary) SetId(v int64)`

SetId sets Id field to given value.


### GetCatalogId

`func (o *CatalogItemSummary) GetCatalogId() int64`

GetCatalogId returns the CatalogId field if non-nil, zero value otherwise.

### GetCatalogIdOk

`func (o *CatalogItemSummary) GetCatalogIdOk() (*int64, bool)`

GetCatalogIdOk returns a tuple with the CatalogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogId

`func (o *CatalogItemSummary) SetCatalogId(v int64)`

SetCatalogId sets CatalogId field to given value.


### GetSku

`func (o *CatalogItemSummary) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CatalogItemSummary) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CatalogItemSummary) SetSku(v string)`

SetSku sets Sku field to given value.


### GetPrice

`func (o *CatalogItemSummary) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemSummary) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItemSummary) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItemSummary) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetProductName

`func (o *CatalogItemSummary) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *CatalogItemSummary) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *CatalogItemSummary) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *CatalogItemSummary) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### GetProductId

`func (o *CatalogItemSummary) GetProductId() int64`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CatalogItemSummary) GetProductIdOk() (*int64, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CatalogItemSummary) SetProductId(v int64)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *CatalogItemSummary) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetSkuId

`func (o *CatalogItemSummary) GetSkuId() int64`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *CatalogItemSummary) GetSkuIdOk() (*int64, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *CatalogItemSummary) SetSkuId(v int64)`

SetSkuId sets SkuId field to given value.

### HasSkuId

`func (o *CatalogItemSummary) HasSkuId() bool`

HasSkuId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


