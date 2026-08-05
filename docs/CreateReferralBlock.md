# CreateReferralBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**CampaignId** | [**CreateReferralBlock1CampaignId**](CreateReferralBlock1CampaignId.md) |  | 
**FriendId** | **string** | An optional integration ID of the friend&#39;s profile. | 
**StoreInSession** | **bool** | When &#x60;true&#x60;, the referral code is stored in the session. | 
**UsageLimit** | Pointer to [**CreateReferralBlock1UsageLimit**](CreateReferralBlock1UsageLimit.md) |  | [optional] 
**StartDate** | Pointer to **interface{}** |  | [optional] 
**ExpiryDate** | Pointer to **interface{}** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**ValidCharacters** | Pointer to **string** | Characters used to generate the random parts of a code. | [optional] 
**Pattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### NewCreateReferralBlock

`func NewCreateReferralBlock(id string, type_ string, campaignId CreateReferralBlock1CampaignId, friendId string, storeInSession bool, ) *CreateReferralBlock`

NewCreateReferralBlock instantiates a new CreateReferralBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReferralBlockWithDefaults

`func NewCreateReferralBlockWithDefaults() *CreateReferralBlock`

NewCreateReferralBlockWithDefaults instantiates a new CreateReferralBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReferralBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReferralBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReferralBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CreateReferralBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateReferralBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateReferralBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CreateReferralBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateReferralBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateReferralBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateReferralBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignId

`func (o *CreateReferralBlock) GetCampaignId() CreateReferralBlock1CampaignId`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreateReferralBlock) GetCampaignIdOk() (*CreateReferralBlock1CampaignId, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreateReferralBlock) SetCampaignId(v CreateReferralBlock1CampaignId)`

SetCampaignId sets CampaignId field to given value.


### GetFriendId

`func (o *CreateReferralBlock) GetFriendId() string`

GetFriendId returns the FriendId field if non-nil, zero value otherwise.

### GetFriendIdOk

`func (o *CreateReferralBlock) GetFriendIdOk() (*string, bool)`

GetFriendIdOk returns a tuple with the FriendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendId

`func (o *CreateReferralBlock) SetFriendId(v string)`

SetFriendId sets FriendId field to given value.


### GetStoreInSession

`func (o *CreateReferralBlock) GetStoreInSession() bool`

GetStoreInSession returns the StoreInSession field if non-nil, zero value otherwise.

### GetStoreInSessionOk

`func (o *CreateReferralBlock) GetStoreInSessionOk() (*bool, bool)`

GetStoreInSessionOk returns a tuple with the StoreInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInSession

`func (o *CreateReferralBlock) SetStoreInSession(v bool)`

SetStoreInSession sets StoreInSession field to given value.


### GetUsageLimit

`func (o *CreateReferralBlock) GetUsageLimit() CreateReferralBlock1UsageLimit`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CreateReferralBlock) GetUsageLimitOk() (*CreateReferralBlock1UsageLimit, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CreateReferralBlock) SetUsageLimit(v CreateReferralBlock1UsageLimit)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *CreateReferralBlock) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateReferralBlock) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateReferralBlock) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateReferralBlock) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateReferralBlock) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *CreateReferralBlock) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *CreateReferralBlock) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetExpiryDate

`func (o *CreateReferralBlock) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateReferralBlock) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CreateReferralBlock) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CreateReferralBlock) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *CreateReferralBlock) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *CreateReferralBlock) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetAttributes

`func (o *CreateReferralBlock) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateReferralBlock) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateReferralBlock) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateReferralBlock) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *CreateReferralBlock) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *CreateReferralBlock) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetValidCharacters

`func (o *CreateReferralBlock) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *CreateReferralBlock) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *CreateReferralBlock) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *CreateReferralBlock) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetPattern

`func (o *CreateReferralBlock) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateReferralBlock) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateReferralBlock) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateReferralBlock) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


