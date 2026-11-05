# RewardCatalogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The unique ID of the reward. | 
**Name** | **string** | The customer-facing name of the reward. | 
**Description** | Pointer to **string** | The customer-facing description of the reward. | [optional] 
**PointsRequired** | Pointer to [**[]RewardPointsRequired**](RewardPointsRequired.md) | The loyalty points required to activate the reward. | [optional] 
**Rule** | [**RuleMetadata**](RuleMetadata.md) | Customer-facing rule metadata for the reward. | 
**Eligibility** | Pointer to [**RewardEligibility**](RewardEligibility.md) | The customer&#39;s eligibility for the reward. Returned only when the request includes a &#x60;profileIntegrationId&#x60; or &#x60;loyaltyCardId&#x60;.  | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the customer rewards unlocked from this reward remain active, relative to when the customer reward was unlocked.  The time format is an **integer** followed by one letter indicating the time unit.  Examples: &#x60;30m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  Certain units can be rounded up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  This field is not returned in combination with &#x60;expiryDate&#x60;. If both &#x60;validityDuration&#x60; and &#x60;expiryDate&#x60; are not returned, customer rewards have unlimited validity.  | [optional] 
**ExpiryDate** | Pointer to **time.Time** | Expiry date of the customer rewards unlocked from this reward. Customer rewards never expire if this is omitted and &#x60;validityDuration&#x60; is not set.  **Note:** - This field is not returned in combination with &#x60;validityDuration&#x60;.  | [optional] 

## Methods

### NewRewardCatalogItem

`func NewRewardCatalogItem(id int64, name string, rule RuleMetadata, ) *RewardCatalogItem`

NewRewardCatalogItem instantiates a new RewardCatalogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardCatalogItemWithDefaults

`func NewRewardCatalogItemWithDefaults() *RewardCatalogItem`

NewRewardCatalogItemWithDefaults instantiates a new RewardCatalogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardCatalogItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardCatalogItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardCatalogItem) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *RewardCatalogItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RewardCatalogItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RewardCatalogItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RewardCatalogItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardCatalogItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardCatalogItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardCatalogItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPointsRequired

`func (o *RewardCatalogItem) GetPointsRequired() []RewardPointsRequired`

GetPointsRequired returns the PointsRequired field if non-nil, zero value otherwise.

### GetPointsRequiredOk

`func (o *RewardCatalogItem) GetPointsRequiredOk() (*[]RewardPointsRequired, bool)`

GetPointsRequiredOk returns a tuple with the PointsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointsRequired

`func (o *RewardCatalogItem) SetPointsRequired(v []RewardPointsRequired)`

SetPointsRequired sets PointsRequired field to given value.

### HasPointsRequired

`func (o *RewardCatalogItem) HasPointsRequired() bool`

HasPointsRequired returns a boolean if a field has been set.

### GetRule

`func (o *RewardCatalogItem) GetRule() RuleMetadata`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RewardCatalogItem) GetRuleOk() (*RuleMetadata, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RewardCatalogItem) SetRule(v RuleMetadata)`

SetRule sets Rule field to given value.


### GetEligibility

`func (o *RewardCatalogItem) GetEligibility() RewardEligibility`

GetEligibility returns the Eligibility field if non-nil, zero value otherwise.

### GetEligibilityOk

`func (o *RewardCatalogItem) GetEligibilityOk() (*RewardEligibility, bool)`

GetEligibilityOk returns a tuple with the Eligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibility

`func (o *RewardCatalogItem) SetEligibility(v RewardEligibility)`

SetEligibility sets Eligibility field to given value.

### HasEligibility

`func (o *RewardCatalogItem) HasEligibility() bool`

HasEligibility returns a boolean if a field has been set.

### GetValidityDuration

`func (o *RewardCatalogItem) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *RewardCatalogItem) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *RewardCatalogItem) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *RewardCatalogItem) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetExpiryDate

`func (o *RewardCatalogItem) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *RewardCatalogItem) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *RewardCatalogItem) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *RewardCatalogItem) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


