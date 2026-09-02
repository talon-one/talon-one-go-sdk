# IntegrationRewardsCatalog200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | [**IntegrationRewardsCatalog200ResponseCatalog**](IntegrationRewardsCatalog200ResponseCatalog.md) |  | 
**Loyalty** | Pointer to [**map[string]LoyaltyBalances**](LoyaltyBalances.md) | The customer&#39;s loyalty balances for the specified loyalty program. Returned only when &#x60;loyaltyProgramId&#x60; is provided together with &#x60;profileIntegrationId&#x60; or &#x60;loyaltyCardId&#x60;.  | [optional] 

## Methods

### NewIntegrationRewardsCatalog200Response

`func NewIntegrationRewardsCatalog200Response(catalog IntegrationRewardsCatalog200ResponseCatalog, ) *IntegrationRewardsCatalog200Response`

NewIntegrationRewardsCatalog200Response instantiates a new IntegrationRewardsCatalog200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationRewardsCatalog200ResponseWithDefaults

`func NewIntegrationRewardsCatalog200ResponseWithDefaults() *IntegrationRewardsCatalog200Response`

NewIntegrationRewardsCatalog200ResponseWithDefaults instantiates a new IntegrationRewardsCatalog200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *IntegrationRewardsCatalog200Response) GetCatalog() IntegrationRewardsCatalog200ResponseCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *IntegrationRewardsCatalog200Response) GetCatalogOk() (*IntegrationRewardsCatalog200ResponseCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *IntegrationRewardsCatalog200Response) SetCatalog(v IntegrationRewardsCatalog200ResponseCatalog)`

SetCatalog sets Catalog field to given value.


### GetLoyalty

`func (o *IntegrationRewardsCatalog200Response) GetLoyalty() map[string]LoyaltyBalances`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *IntegrationRewardsCatalog200Response) GetLoyaltyOk() (*map[string]LoyaltyBalances, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *IntegrationRewardsCatalog200Response) SetLoyalty(v map[string]LoyaltyBalances)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *IntegrationRewardsCatalog200Response) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


