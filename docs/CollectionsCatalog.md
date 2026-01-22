# CollectionsCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this collection. | 
**AccountId** | **int64** | The ID of the account that owns this collection. | 
**ApplicationId** | **int64** | The ID of the Application that owns this collection. | 
**Title** | **string** | A short description of the collection catalog. | 
**Description** | Pointer to **string** | A longer, more detailed description of the collection catalog. | [optional] 
**Category** | **string** | Category used to group collection catalogs. | 
**Source** | **string** | Indicates whether the collection is custom or shipped by Talon.One. | 
**Rules** | [**[]CatalogRule**](CatalogRule.md) | Array of rule templates in this collection catalog. Rules only contain title (no description, as description is at the collection catalog level). | 
**CartItemFilters** | [**[]CartItemFilterTemplate**](CartItemFilterTemplate.md) | Array of cart item filter templates in this collection catalog. Cart item filters only contain name (no description, as description is at the collection catalog level). | 
**Created** | **time.Time** | Timestamp when the collection was created. | 
**CreatedBy** | **int64** | ID of the user who created the collection. | 
**Modified** | Pointer to **time.Time** | Timestamp when the collection was last updated. | [optional] 
**ModifiedBy** | Pointer to **int64** | ID of the user who last updated the collection. | [optional] 

## Methods

### NewCollectionsCatalog

`func NewCollectionsCatalog(id int64, accountId int64, applicationId int64, title string, category string, source string, rules []CatalogRule, cartItemFilters []CartItemFilterTemplate, created time.Time, createdBy int64, ) *CollectionsCatalog`

NewCollectionsCatalog instantiates a new CollectionsCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionsCatalogWithDefaults

`func NewCollectionsCatalogWithDefaults() *CollectionsCatalog`

NewCollectionsCatalogWithDefaults instantiates a new CollectionsCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CollectionsCatalog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CollectionsCatalog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CollectionsCatalog) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *CollectionsCatalog) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CollectionsCatalog) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CollectionsCatalog) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetApplicationId

`func (o *CollectionsCatalog) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CollectionsCatalog) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CollectionsCatalog) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetTitle

`func (o *CollectionsCatalog) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CollectionsCatalog) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CollectionsCatalog) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CollectionsCatalog) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CollectionsCatalog) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CollectionsCatalog) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CollectionsCatalog) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *CollectionsCatalog) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CollectionsCatalog) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CollectionsCatalog) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSource

`func (o *CollectionsCatalog) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CollectionsCatalog) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CollectionsCatalog) SetSource(v string)`

SetSource sets Source field to given value.


### GetRules

`func (o *CollectionsCatalog) GetRules() []CatalogRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *CollectionsCatalog) GetRulesOk() (*[]CatalogRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *CollectionsCatalog) SetRules(v []CatalogRule)`

SetRules sets Rules field to given value.


### GetCartItemFilters

`func (o *CollectionsCatalog) GetCartItemFilters() []CartItemFilterTemplate`

GetCartItemFilters returns the CartItemFilters field if non-nil, zero value otherwise.

### GetCartItemFiltersOk

`func (o *CollectionsCatalog) GetCartItemFiltersOk() (*[]CartItemFilterTemplate, bool)`

GetCartItemFiltersOk returns a tuple with the CartItemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemFilters

`func (o *CollectionsCatalog) SetCartItemFilters(v []CartItemFilterTemplate)`

SetCartItemFilters sets CartItemFilters field to given value.


### GetCreated

`func (o *CollectionsCatalog) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CollectionsCatalog) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CollectionsCatalog) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCreatedBy

`func (o *CollectionsCatalog) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CollectionsCatalog) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CollectionsCatalog) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetModified

`func (o *CollectionsCatalog) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *CollectionsCatalog) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *CollectionsCatalog) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *CollectionsCatalog) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetModifiedBy

`func (o *CollectionsCatalog) GetModifiedBy() int64`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *CollectionsCatalog) GetModifiedByOk() (*int64, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *CollectionsCatalog) SetModifiedBy(v int64)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *CollectionsCatalog) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


