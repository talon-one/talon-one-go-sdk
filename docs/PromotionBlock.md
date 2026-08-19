# PromotionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for this block. | 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] 
**Operator** | **string** | An indicator of how the block compares its elements. | 
**Blocks** | [**[]PromotionBlock**](PromotionBlock.md) | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to [**[]PromotionBlock**](PromotionBlock.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]PromotionBlock**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 
**Name** | **string** | A custom description recorded as the reason for the point deduction. | 
**Value** | [**RedeemLoyaltyPointsBlock1Value**](RedeemLoyaltyPointsBlock1Value.md) |  | 
**Partial** | **bool** | When set to &#x60;true&#x60;, applies a partial item reward if the remaining budget is insufficient to award the full reward. | 
**Target** | [**TriggerCustomEffectBlock1Target**](TriggerCustomEffectBlock1Target.md) |  | 
**Expression** | **[]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 
**NotificationType** | **string** | The type of notification to display. | 
**Title** | **string** | The notification heading shown to the customer. | 
**Body** | Pointer to **string** | The notification body text. Supports template placeholders (e.g. \&quot;{{$Session.Total}}\&quot;) evaluated at rule execution time. | [optional] 
**Sku** | **string** | The stock keeping unit of the item to award. | 
**Quantity** | **string** | The number of items to award. Supports template placeholders (e.g. \&quot;{{$Session.Total / 2}}\&quot;) for dynamic quantities. | 
**GiveawayPool** | [**AwardGiveawayBlock1GiveawayPool**](AwardGiveawayBlock1GiveawayPool.md) |  | 
**Profile** | **string** | The customer profile to add or remove from the audience. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Attribute** | [**UpdateAttributeValueBlock1Attribute**](UpdateAttributeValueBlock1Attribute.md) |  | 
**Min** | Pointer to **interface{}** |  | [optional] 
**Max** | Pointer to **interface{}** |  | [optional] 
**Start** | Pointer to **interface{}** |  | [optional] 
**End** | Pointer to **interface{}** |  | [optional] 
**StartInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;start&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**EndInclusive** | Pointer to **bool** | When &#x60;true&#x60;, the &#x60;end&#x60; value is included in the range for the &#x60;within&#x60; operator. | [optional] 
**TimezoneInsensitive** | Pointer to **bool** | Indicates whether the &#x60;within&#x60; operator ignores time zones and compares the wall-clock time only. When &#x60;false&#x60;, time zones are taken into account. | [optional] 
**Values** | Pointer to **interface{}** |  | [optional] 
**Count** | Pointer to **interface{}** |  | [optional] 
**Audience** | [**UpdateAudienceMembershipBlock1Audience**](UpdateAudienceMembershipBlock1Audience.md) |  | 
**Program** | [**RedeemLoyaltyPointsBlock1Program**](RedeemLoyaltyPointsBlock1Program.md) |  | 
**Subledger** | **string** | The name of the subledger to deduct points from. Can be empty if this block deducts from the loyalty program&#39;s main ledger instead of a subledger. | 
**Balance** | **string** | The type of balance to check:  - &#x60;current&#x60; is the sum of currently active points  - &#x60;pending&#x60; is the sum of pending points.  - &#x60;negative&#x60; is the sum of negative points.  - &#x60;tentativeCurrent&#x60; is the tentative points balance within the current open customer session. | 
**Redeem** | **bool** | When &#x60;true&#x60;, the referral code is redeemed. | 
**Achievement** | [**CheckAchievementBlock1Achievement**](CheckAchievementBlock1Achievement.md) |  | 
**Webhook** | [**TriggerWebhookBlock1Webhook**](TriggerWebhookBlock1Webhook.md) |  | 
**Params** | Pointer to **map[string]interface{}** | The custom effect&#39;s parameters, in configured order. Each property name is the parameter&#39;s title, lowercased with spaces replaced by underscores (for example, &#x60;Order ID&#x60; becomes &#x60;order_id&#x60;); falls back to &#x60;param_0&#x60;, &#x60;param_1&#x60;, and so on if a title is blank or collides with another. | [optional] 
**CustomEffect** | [**TriggerCustomEffectBlock1CustomEffect**](TriggerCustomEffectBlock1CustomEffect.md) |  | 
**EventType** | **string** | The event type to check against. | 
**Matchers** | Pointer to [**[]PromotionBlock**](PromotionBlock.md) |  | [optional] 
**Action** | **string** | The limitable action to check. | 
**CampaignId** | [**CreateReferralBlock1CampaignId**](CreateReferralBlock1CampaignId.md) |  | 
**RecipientId** | **string** | The integration ID of the customer that is allowed to redeem this coupon. | 
**StoreInSession** | **bool** | When &#x60;true&#x60;, the referral code is stored in the session. | 
**UsageLimit** | Pointer to [**CreateReferralBlock1UsageLimit**](CreateReferralBlock1UsageLimit.md) |  | [optional] 
**DiscountLimit** | Pointer to [**CreateCouponBlock1DiscountLimit**](CreateCouponBlock1DiscountLimit.md) |  | [optional] 
**StartDate** | Pointer to **interface{}** |  | [optional] 
**ExpiryDate** | Pointer to **interface{}** |  | [optional] 
**Attributes** | Pointer to **interface{}** |  | [optional] 
**ValidCharacters** | Pointer to **string** | Characters used to generate the random parts of a code. | [optional] 
**Pattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 
**FriendId** | **string** | An optional integration ID of the friend&#39;s profile. | 
**Tier** | [**CheckTierBlock1Tier**](CheckTierBlock1Tier.md) |  | 

## Methods

### NewPromotionBlock

`func NewPromotionBlock(id string, type_ string, operator string, blocks []PromotionBlock, name string, value RedeemLoyaltyPointsBlock1Value, partial bool, target TriggerCustomEffectBlock1Target, expression []interface{}, notificationType string, title string, sku string, quantity string, giveawayPool AwardGiveawayBlock1GiveawayPool, profile string, attribute UpdateAttributeValueBlock1Attribute, audience UpdateAudienceMembershipBlock1Audience, program RedeemLoyaltyPointsBlock1Program, subledger string, balance string, redeem bool, achievement CheckAchievementBlock1Achievement, webhook TriggerWebhookBlock1Webhook, customEffect TriggerCustomEffectBlock1CustomEffect, eventType string, action string, campaignId CreateReferralBlock1CampaignId, recipientId string, storeInSession bool, friendId string, tier CheckTierBlock1Tier, ) *PromotionBlock`

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


### GetValue

`func (o *PromotionBlock) GetValue() RedeemLoyaltyPointsBlock1Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PromotionBlock) GetValueOk() (*RedeemLoyaltyPointsBlock1Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PromotionBlock) SetValue(v RedeemLoyaltyPointsBlock1Value)`

SetValue sets Value field to given value.


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


### GetTarget

`func (o *PromotionBlock) GetTarget() TriggerCustomEffectBlock1Target`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PromotionBlock) GetTargetOk() (*TriggerCustomEffectBlock1Target, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PromotionBlock) SetTarget(v TriggerCustomEffectBlock1Target)`

SetTarget sets Target field to given value.


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
### GetStart

`func (o *PromotionBlock) GetStart() interface{}`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PromotionBlock) GetStartOk() (*interface{}, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PromotionBlock) SetStart(v interface{})`

SetStart sets Start field to given value.

### HasStart

`func (o *PromotionBlock) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *PromotionBlock) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *PromotionBlock) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetEnd

`func (o *PromotionBlock) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *PromotionBlock) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *PromotionBlock) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *PromotionBlock) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *PromotionBlock) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *PromotionBlock) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetStartInclusive

`func (o *PromotionBlock) GetStartInclusive() bool`

GetStartInclusive returns the StartInclusive field if non-nil, zero value otherwise.

### GetStartInclusiveOk

`func (o *PromotionBlock) GetStartInclusiveOk() (*bool, bool)`

GetStartInclusiveOk returns a tuple with the StartInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartInclusive

`func (o *PromotionBlock) SetStartInclusive(v bool)`

SetStartInclusive sets StartInclusive field to given value.

### HasStartInclusive

`func (o *PromotionBlock) HasStartInclusive() bool`

HasStartInclusive returns a boolean if a field has been set.

### GetEndInclusive

`func (o *PromotionBlock) GetEndInclusive() bool`

GetEndInclusive returns the EndInclusive field if non-nil, zero value otherwise.

### GetEndInclusiveOk

`func (o *PromotionBlock) GetEndInclusiveOk() (*bool, bool)`

GetEndInclusiveOk returns a tuple with the EndInclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndInclusive

`func (o *PromotionBlock) SetEndInclusive(v bool)`

SetEndInclusive sets EndInclusive field to given value.

### HasEndInclusive

`func (o *PromotionBlock) HasEndInclusive() bool`

HasEndInclusive returns a boolean if a field has been set.

### GetTimezoneInsensitive

`func (o *PromotionBlock) GetTimezoneInsensitive() bool`

GetTimezoneInsensitive returns the TimezoneInsensitive field if non-nil, zero value otherwise.

### GetTimezoneInsensitiveOk

`func (o *PromotionBlock) GetTimezoneInsensitiveOk() (*bool, bool)`

GetTimezoneInsensitiveOk returns a tuple with the TimezoneInsensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneInsensitive

`func (o *PromotionBlock) SetTimezoneInsensitive(v bool)`

SetTimezoneInsensitive sets TimezoneInsensitive field to given value.

### HasTimezoneInsensitive

`func (o *PromotionBlock) HasTimezoneInsensitive() bool`

HasTimezoneInsensitive returns a boolean if a field has been set.

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


### GetProgram

`func (o *PromotionBlock) GetProgram() RedeemLoyaltyPointsBlock1Program`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *PromotionBlock) GetProgramOk() (*RedeemLoyaltyPointsBlock1Program, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *PromotionBlock) SetProgram(v RedeemLoyaltyPointsBlock1Program)`

SetProgram sets Program field to given value.


### GetSubledger

`func (o *PromotionBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *PromotionBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *PromotionBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetBalance

`func (o *PromotionBlock) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PromotionBlock) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PromotionBlock) SetBalance(v string)`

SetBalance sets Balance field to given value.


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

`func (o *PromotionBlock) GetAchievement() CheckAchievementBlock1Achievement`

GetAchievement returns the Achievement field if non-nil, zero value otherwise.

### GetAchievementOk

`func (o *PromotionBlock) GetAchievementOk() (*CheckAchievementBlock1Achievement, bool)`

GetAchievementOk returns a tuple with the Achievement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievement

`func (o *PromotionBlock) SetAchievement(v CheckAchievementBlock1Achievement)`

SetAchievement sets Achievement field to given value.


### GetWebhook

`func (o *PromotionBlock) GetWebhook() TriggerWebhookBlock1Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *PromotionBlock) GetWebhookOk() (*TriggerWebhookBlock1Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *PromotionBlock) SetWebhook(v TriggerWebhookBlock1Webhook)`

SetWebhook sets Webhook field to given value.


### GetParams

`func (o *PromotionBlock) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *PromotionBlock) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *PromotionBlock) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *PromotionBlock) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCustomEffect

`func (o *PromotionBlock) GetCustomEffect() TriggerCustomEffectBlock1CustomEffect`

GetCustomEffect returns the CustomEffect field if non-nil, zero value otherwise.

### GetCustomEffectOk

`func (o *PromotionBlock) GetCustomEffectOk() (*TriggerCustomEffectBlock1CustomEffect, bool)`

GetCustomEffectOk returns a tuple with the CustomEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffect

`func (o *PromotionBlock) SetCustomEffect(v TriggerCustomEffectBlock1CustomEffect)`

SetCustomEffect sets CustomEffect field to given value.


### GetEventType

`func (o *PromotionBlock) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PromotionBlock) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PromotionBlock) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMatchers

`func (o *PromotionBlock) GetMatchers() []PromotionBlock`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *PromotionBlock) GetMatchersOk() (*[]PromotionBlock, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *PromotionBlock) SetMatchers(v []PromotionBlock)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *PromotionBlock) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetAction

`func (o *PromotionBlock) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PromotionBlock) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PromotionBlock) SetAction(v string)`

SetAction sets Action field to given value.


### GetCampaignId

`func (o *PromotionBlock) GetCampaignId() CreateReferralBlock1CampaignId`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *PromotionBlock) GetCampaignIdOk() (*CreateReferralBlock1CampaignId, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *PromotionBlock) SetCampaignId(v CreateReferralBlock1CampaignId)`

SetCampaignId sets CampaignId field to given value.


### GetRecipientId

`func (o *PromotionBlock) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PromotionBlock) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PromotionBlock) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetStoreInSession

`func (o *PromotionBlock) GetStoreInSession() bool`

GetStoreInSession returns the StoreInSession field if non-nil, zero value otherwise.

### GetStoreInSessionOk

`func (o *PromotionBlock) GetStoreInSessionOk() (*bool, bool)`

GetStoreInSessionOk returns a tuple with the StoreInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInSession

`func (o *PromotionBlock) SetStoreInSession(v bool)`

SetStoreInSession sets StoreInSession field to given value.


### GetUsageLimit

`func (o *PromotionBlock) GetUsageLimit() CreateReferralBlock1UsageLimit`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *PromotionBlock) GetUsageLimitOk() (*CreateReferralBlock1UsageLimit, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *PromotionBlock) SetUsageLimit(v CreateReferralBlock1UsageLimit)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *PromotionBlock) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetDiscountLimit

`func (o *PromotionBlock) GetDiscountLimit() CreateCouponBlock1DiscountLimit`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *PromotionBlock) GetDiscountLimitOk() (*CreateCouponBlock1DiscountLimit, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *PromotionBlock) SetDiscountLimit(v CreateCouponBlock1DiscountLimit)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *PromotionBlock) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *PromotionBlock) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PromotionBlock) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PromotionBlock) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PromotionBlock) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PromotionBlock) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PromotionBlock) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetExpiryDate

`func (o *PromotionBlock) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PromotionBlock) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PromotionBlock) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PromotionBlock) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *PromotionBlock) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *PromotionBlock) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetAttributes

`func (o *PromotionBlock) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PromotionBlock) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PromotionBlock) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PromotionBlock) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *PromotionBlock) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *PromotionBlock) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetValidCharacters

`func (o *PromotionBlock) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *PromotionBlock) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *PromotionBlock) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *PromotionBlock) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetPattern

`func (o *PromotionBlock) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *PromotionBlock) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *PromotionBlock) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *PromotionBlock) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetFriendId

`func (o *PromotionBlock) GetFriendId() string`

GetFriendId returns the FriendId field if non-nil, zero value otherwise.

### GetFriendIdOk

`func (o *PromotionBlock) GetFriendIdOk() (*string, bool)`

GetFriendIdOk returns a tuple with the FriendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendId

`func (o *PromotionBlock) SetFriendId(v string)`

SetFriendId sets FriendId field to given value.


### GetTier

`func (o *PromotionBlock) GetTier() CheckTierBlock1Tier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PromotionBlock) GetTierOk() (*CheckTierBlock1Tier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PromotionBlock) SetTier(v CheckTierBlock1Tier)`

SetTier sets Tier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


