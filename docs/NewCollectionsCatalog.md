# NewCollectionsCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The display name for the collection catalog. | 
**Description** | Pointer to **string** | A longer, more detailed description of the collection catalog. | [optional] 
**Category** | Pointer to **string** | Category used to group collection catalogs. | [optional] [default to "custom"]
**Rules** | Pointer to [**[]CatalogRule**](CatalogRule.md) | Array of rules to store in this collection catalog. Rules should only contain title (no description, as description is at the collection catalog level). At least one rule or cart item filter is required. | [optional] 
**CartItemFilters** | Pointer to [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Array of cart item filters to store in this collection catalog. If not provided, will be extracted from the rules. Cart item filters should only contain name (no description, as description is at the collection catalog level). | [optional] 

## Methods

### NewNewCollectionsCatalog

`func NewNewCollectionsCatalog(title string, ) *NewCollectionsCatalog`

NewNewCollectionsCatalog instantiates a new NewCollectionsCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCollectionsCatalogWithDefaults

`func NewNewCollectionsCatalogWithDefaults() *NewCollectionsCatalog`

NewNewCollectionsCatalogWithDefaults instantiates a new NewCollectionsCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NewCollectionsCatalog) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewCollectionsCatalog) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewCollectionsCatalog) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *NewCollectionsCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewCollectionsCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewCollectionsCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewCollectionsCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NewCollectionsCatalog) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NewCollectionsCatalog) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NewCollectionsCatalog) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NewCollectionsCatalog) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetRules

`func (o *NewCollectionsCatalog) GetRules() []CatalogRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *NewCollectionsCatalog) GetRulesOk() (*[]CatalogRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *NewCollectionsCatalog) SetRules(v []CatalogRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *NewCollectionsCatalog) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetCartItemFilters

`func (o *NewCollectionsCatalog) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *NewCollectionsCatalog) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *NewCollectionsCatalog) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.

### HasCartItemFilters

`func (o *NewCollectionsCatalog) HasCartItemFilters() bool`

HasCartItemFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


