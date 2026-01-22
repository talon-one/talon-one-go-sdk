# UpdateCollectionsCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display name for the collection catalog. | [optional] 
**Description** | Pointer to **string** | A longer, more detailed description of the collection catalog. | [optional] 
**Category** | Pointer to **string** | Category used to group collection catalogs. | [optional] 
**Rules** | Pointer to [**[]CatalogRule**](CatalogRule.md) | Replaces the stored rules. Rules should only contain title (no description, as description is at the collection catalog level). | [optional] 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Replaces the stored cart item filters. Cart item filters should only contain name (no description, as description is at the collection catalog level). | [optional] 

## Methods

### NewUpdateCollectionsCatalog

`func NewUpdateCollectionsCatalog() *UpdateCollectionsCatalog`

NewUpdateCollectionsCatalog instantiates a new UpdateCollectionsCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCollectionsCatalogWithDefaults

`func NewUpdateCollectionsCatalogWithDefaults() *UpdateCollectionsCatalog`

NewUpdateCollectionsCatalogWithDefaults instantiates a new UpdateCollectionsCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdateCollectionsCatalog) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateCollectionsCatalog) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateCollectionsCatalog) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateCollectionsCatalog) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCollectionsCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCollectionsCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCollectionsCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCollectionsCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *UpdateCollectionsCatalog) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateCollectionsCatalog) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateCollectionsCatalog) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateCollectionsCatalog) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetRules

`func (o *UpdateCollectionsCatalog) GetRules() []CatalogRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *UpdateCollectionsCatalog) GetRulesOk() (*[]CatalogRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *UpdateCollectionsCatalog) SetRules(v []CatalogRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *UpdateCollectionsCatalog) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCartItemFilters

`func (o *UpdateCollectionsCatalog) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *UpdateCollectionsCatalog) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *UpdateCollectionsCatalog) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.

### HasCartItemFilters

`func (o *UpdateCollectionsCatalog) HasCartItemFilters() bool`

HasCartItemFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


