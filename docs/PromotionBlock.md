# PromotionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | The update operation applied to the attribute. | 
**Blocks** | [**[]PromotionBlock**](PromotionBlock.md) | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to [**[]PromotionBlock**](PromotionBlock.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]PromotionBlock**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 
**Expression** | **[]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 
**NotificationType** | **string** | The type of notification to display. | 
**Title** | **string** | The notification heading shown to the customer. | 
**Body** | Pointer to **string** | The notification body text. Supports template placeholders (e.g. \&quot;{{$Session.Total}}\&quot;) evaluated at rule execution time. | [optional] 
**Sku** | **string** | The stock keeping unit of the item to award. | 
**Name** | **string** | The display name of the item to award. | 
**Quantity** | **string** | The number of items to award. Supports template placeholders (e.g. \&quot;{{$Session.Total / 2}}\&quot;) for dynamic quantities. | 
**Partial** | Pointer to **bool** | When set to &#x60;true&#x60;, applies a partial item reward if the remaining budget is insufficient to award the full reward. | [optional] 
**GiveawayPool** | [**AwardGiveawayBlock1GiveawayPool**](AwardGiveawayBlock1GiveawayPool.md) |  | 
**Profile** | **string** | The customer profile to add or remove from the audience. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Attribute** | [**UpdateAttributeValueBlock1Attribute**](UpdateAttributeValueBlock1Attribute.md) |  | 
**Value** | **interface{}** |  | 
**Min** | Pointer to **interface{}** |  | [optional] 
**Max** | Pointer to **interface{}** |  | [optional] 
**Values** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **interface{}** |  | [optional] 
**Audience** | [**UpdateAudienceMembershipBlock1Audience**](UpdateAudienceMembershipBlock1Audience.md) |  | 
**Redeem** | **bool** | When &#x60;true&#x60;, the referral code is redeemed. | 
**Achievement** | [**UpdateAchievementProgressBlock1Achievement**](UpdateAchievementProgressBlock1Achievement.md) |  | 
**Target** | [**UpdateAttributeValueBlock1Target**](UpdateAttributeValueBlock1Target.md) |  | 

## Methods

### NewPromotionBlock

`func NewPromotionBlock(id string, type_ string, operator string, blocks []PromotionBlock, expression []interface{}, notificationType string, title string, sku string, name string, quantity string, giveawayPool AwardGiveawayBlock1GiveawayPool, profile string, attribute UpdateAttributeValueBlock1Attribute, value interface{}, audience UpdateAudienceMembershipBlock1Audience, redeem bool, achievement UpdateAchievementProgressBlock1Achievement, target UpdateAttributeValueBlock1Target, ) *PromotionBlock`

NewPromotionBlock instantiates a new PromotionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromotionBlockWithDefaults

`func NewPromotionBlockWithDefaults() *PromotionBlock`

NewPromotionBlockWithDefaults instantiates a new PromotionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PromotionBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PromotionBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PromotionBlock) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PromotionBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PromotionBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PromotionBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *PromotionBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PromotionBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PromotionBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PromotionBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *PromotionBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PromotionBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PromotionBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *PromotionBlock) GetBlocks() []PromotionBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *PromotionBlock) GetBlocksOk() (*[]PromotionBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *PromotionBlock) SetBlocks(v []PromotionBlock)`

SetBlocks sets Blocks field to given value.


### GetOnFailure

`func (o *PromotionBlock) GetOnFailure() []PromotionBlock`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *PromotionBlock) GetOnFailureOk() (*[]PromotionBlock, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *PromotionBlock) SetOnFailure(v []PromotionBlock)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *PromotionBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *PromotionBlock) GetOnError() map[string][]PromotionBlock`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *PromotionBlock) GetOnErrorOk() (*map[string][]PromotionBlock, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *PromotionBlock) SetOnError(v map[string][]PromotionBlock)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *PromotionBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetExpression

`func (o *PromotionBlock) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *PromotionBlock) GetExpressionOk() (*[]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *PromotionBlock) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.


### GetNotificationType

`func (o *PromotionBlock) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *PromotionBlock) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *PromotionBlock) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTitle

`func (o *PromotionBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PromotionBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PromotionBlock) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *PromotionBlock) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PromotionBlock) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PromotionBlock) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PromotionBlock) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSku

`func (o *PromotionBlock) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *PromotionBlock) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *PromotionBlock) SetSku(v string)`

SetSku sets Sku field to given value.


### GetName

`func (o *PromotionBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PromotionBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PromotionBlock) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *PromotionBlock) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PromotionBlock) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PromotionBlock) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetPartial

`func (o *PromotionBlock) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *PromotionBlock) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *PromotionBlock) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *PromotionBlock) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetGiveawayPool

`func (o *PromotionBlock) GetGiveawayPool() AwardGiveawayBlock1GiveawayPool`

GetGiveawayPool returns the GiveawayPool field if non-nil, zero value otherwise.

### GetGiveawayPoolOk

`func (o *PromotionBlock) GetGiveawayPoolOk() (*AwardGiveawayBlock1GiveawayPool, bool)`

GetGiveawayPoolOk returns a tuple with the GiveawayPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayPool

`func (o *PromotionBlock) SetGiveawayPool(v AwardGiveawayBlock1GiveawayPool)`

SetGiveawayPool sets GiveawayPool field to given value.


### GetProfile

`func (o *PromotionBlock) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *PromotionBlock) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *PromotionBlock) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAttribute

`func (o *PromotionBlock) GetAttribute() UpdateAttributeValueBlock1Attribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *PromotionBlock) GetAttributeOk() (*UpdateAttributeValueBlock1Attribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *PromotionBlock) SetAttribute(v UpdateAttributeValueBlock1Attribute)`

SetAttribute sets Attribute field to given value.


### GetValue

`func (o *PromotionBlock) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PromotionBlock) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PromotionBlock) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *PromotionBlock) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PromotionBlock) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *PromotionBlock) GetMin() interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PromotionBlock) GetMinOk() (*interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PromotionBlock) SetMin(v interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *PromotionBlock) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *PromotionBlock) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *PromotionBlock) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *PromotionBlock) GetMax() interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PromotionBlock) GetMaxOk() (*interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PromotionBlock) SetMax(v interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *PromotionBlock) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *PromotionBlock) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *PromotionBlock) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetValues

`func (o *PromotionBlock) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PromotionBlock) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PromotionBlock) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *PromotionBlock) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *PromotionBlock) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *PromotionBlock) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetCount

`func (o *PromotionBlock) GetCount() interface{}`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PromotionBlock) GetCountOk() (*interface{}, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PromotionBlock) SetCount(v interface{})`

SetCount sets Count field to given value.

### HasCount

`func (o *PromotionBlock) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PromotionBlock) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PromotionBlock) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetAudience

`func (o *PromotionBlock) GetAudience() UpdateAudienceMembershipBlock1Audience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *PromotionBlock) GetAudienceOk() (*UpdateAudienceMembershipBlock1Audience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *PromotionBlock) SetAudience(v UpdateAudienceMembershipBlock1Audience)`

SetAudience sets Audience field to given value.


### GetRedeem

`func (o *PromotionBlock) GetRedeem() bool`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *PromotionBlock) GetRedeemOk() (*bool, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *PromotionBlock) SetRedeem(v bool)`

SetRedeem sets Redeem field to given value.


### GetAchievement

`func (o *PromotionBlock) GetAchievement() UpdateAchievementProgressBlock1Achievement`

GetAchievement returns the Achievement field if non-nil, zero value otherwise.

### GetAchievementOk

`func (o *PromotionBlock) GetAchievementOk() (*UpdateAchievementProgressBlock1Achievement, bool)`

GetAchievementOk returns a tuple with the Achievement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievement

`func (o *PromotionBlock) SetAchievement(v UpdateAchievementProgressBlock1Achievement)`

SetAchievement sets Achievement field to given value.


### GetTarget

`func (o *PromotionBlock) GetTarget() UpdateAttributeValueBlock1Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PromotionBlock) GetTargetOk() (*UpdateAttributeValueBlock1Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PromotionBlock) SetTarget(v UpdateAttributeValueBlock1Target)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


