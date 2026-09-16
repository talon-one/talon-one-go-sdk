# Block

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | **string** | An indicator of how the block compares its elements. | 
**Blocks** | [**[]Block**](Block.md) | Child blocks evaluated according to the operator. | 
**OnFailure** | Pointer to [**[]Block**](Block.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]Block**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 
**Name** | **string** | A custom description recorded as the reason for the point deduction. | 
**Value** | [**RedeemLoyaltyPointsBlock1Value**](RedeemLoyaltyPointsBlock1Value.md) |  | 
**Partial** | **bool** | When &#x60;true&#x60;, applies a partial points reward when the requested value exceeds the configured budget. | 
**Target** | [**AwardLoyaltyPointsTarget**](AwardLoyaltyPointsTarget.md) |  | 
**Expression** | **[]interface{}** | The raw Talang expression as an array. For a function call, the first element is the function name and subsequent elements are its arguments. For any other expression (for example a bare attribute path or a literal value), this is a single-element array containing that value. | 
**NotificationType** | **string** | The type of notification to display. | 
**Title** | **string** | The notification heading shown to the customer. | 
**Body** | Pointer to **string** | The notification body text. Supports template placeholders (e.g. \&quot;{{$Session.Total}}\&quot;) evaluated at rule execution time. | [optional] 
**Sku** | **string** | The stock keeping unit of the item to award. | 
**Quantity** | **string** | The number of items to award. Supports template placeholders (e.g. \&quot;{{$Session.Total / 2}}\&quot;) for dynamic quantities. | 
**GiveawayPool** | [**GiveawayPoolBlockReference**](GiveawayPoolBlockReference.md) | The giveaway pool from which an item is awarded. | 
**Profile** | **string** | The customer profile to add or remove from the audience. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Audience** | [**AudienceBlockReference**](AudienceBlockReference.md) | The audience to add the customer to or remove them from. | 
**Program** | [**RedeemLoyaltyPointsBlock1Program**](RedeemLoyaltyPointsBlock1Program.md) |  | 
**Subledger** | **string** | The name of the subledger to deduct points from. Can be empty if this block deducts from the loyalty program&#39;s main ledger instead of a subledger. | 
**Balance** | **string** | The type of balance to check:  - &#x60;current&#x60; is the sum of currently active points  - &#x60;pending&#x60; is the sum of pending points.  - &#x60;negative&#x60; is the sum of negative points.  - &#x60;tentativeCurrent&#x60; is the tentative points balance within the current open customer session. | 
**Redeem** | **bool** | When &#x60;true&#x60;, the referral code is redeemed. | 
**Achievement** | [**AchievementBlockReference**](AchievementBlockReference.md) | The achievement to check for. | 
**Attribute** | [**AttributeBlockReference**](AttributeBlockReference.md) | The attribute being updated. | 
**Webhook** | [**WebhookBlockReference**](WebhookBlockReference.md) | The webhook to trigger. | 
**Params** | Pointer to **map[string]interface{}** | The custom effect&#39;s parameters, in configured order. Each property name is the parameter&#39;s title, lowercased with spaces replaced by underscores (for example, &#x60;Order ID&#x60; becomes &#x60;order_id&#x60;); falls back to &#x60;param_0&#x60;, &#x60;param_1&#x60;, and so on if a title is blank or collides with another. | [optional] 
**CustomEffect** | [**CustomEffectBlockReference**](CustomEffectBlockReference.md) | The custom effect to trigger. | 
**EventType** | **string** | The event type to check against. | 
**Matchers** | Pointer to [**[]Block**](Block.md) |  | [optional] 
**Action** | **string** | The limitable action to check. | 
**CampaignId** | [**CreateReferralBlock1CampaignId**](CreateReferralBlock1CampaignId.md) |  | 
**RecipientId** | **string** | The integration ID of the customer that is allowed to redeem this coupon. | 
**StoreInSession** | **bool** | When &#x60;true&#x60;, the referral code is stored in the session. | 
**UsageLimit** | Pointer to [**CreateReferralBlock1UsageLimit**](CreateReferralBlock1UsageLimit.md) |  | [optional] 
**DiscountLimit** | Pointer to [**CreateCouponBlock1DiscountLimit**](CreateCouponBlock1DiscountLimit.md) |  | [optional] 
**StartDate** | Pointer to **interface{}** | Timestamp at which the awarded points become active. Mutually exclusive with &#x60;awaitsActivation&#x60;. | [optional] 
**ExpiryDate** | Pointer to **interface{}** | Timestamp at which the awarded points expire. Mutually exclusive with &#x60;validityDuration&#x60;. | [optional] 
**Attributes** | Pointer to **interface{}** | Custom attributes associated with this referral code. | [optional] 
**ValidCharacters** | Pointer to **string** | Characters used to generate the random parts of a code. | [optional] 
**Pattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 
**FriendId** | **string** | An optional integration ID of the friend&#39;s profile. | 
**Recipient** | **string** | The customer profile that receives the points. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Tier** | [**TierBlockReference**](TierBlockReference.md) | The tier to check for. | 
**AwaitsActivation** | Pointer to **bool** | When &#x60;true&#x60;, the awarded points require manual or delayed activation before becoming active. Mutually exclusive with &#x60;startDate&#x60;. | [optional] 
**ValidityDuration** | Pointer to **string** | Relative duration (e.g. &#x60;30D&#x60;) after which the awarded points expire. Mutually exclusive with &#x60;expiryDate&#x60;. | [optional] 
**PendingDuration** | Pointer to **string** | Relative duration (e.g. &#x60;3D&#x60;) the awarded points remain pending before activation. | [optional] 

## Methods

### NewBlock

`func NewBlock(type_ string, operator string, blocks []Block, name string, value RedeemLoyaltyPointsBlock1Value, partial bool, target AwardLoyaltyPointsTarget, expression []interface{}, notificationType string, title string, sku string, quantity string, giveawayPool GiveawayPoolBlockReference, profile string, audience AudienceBlockReference, program RedeemLoyaltyPointsBlock1Program, subledger string, balance string, redeem bool, achievement AchievementBlockReference, attribute AttributeBlockReference, webhook WebhookBlockReference, customEffect CustomEffectBlockReference, eventType string, action string, campaignId CreateReferralBlock1CampaignId, recipientId string, storeInSession bool, friendId string, recipient string, tier TierBlockReference, ) *Block`

NewBlock instantiates a new Block object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWithDefaults

`func NewBlockWithDefaults() *Block`

NewBlockWithDefaults instantiates a new Block object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Block) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Block) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Block) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Block) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Block) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Block) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Block) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *Block) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Block) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Block) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Block) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *Block) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Block) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Block) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetBlocks

`func (o *Block) GetBlocks() []Block`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *Block) GetBlocksOk() (*[]Block, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *Block) SetBlocks(v []Block)`

SetBlocks sets Blocks field to given value.


### GetOnFailure

`func (o *Block) GetOnFailure() []Block`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *Block) GetOnFailureOk() (*[]Block, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *Block) SetOnFailure(v []Block)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *Block) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *Block) GetOnError() map[string][]Block`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *Block) GetOnErrorOk() (*map[string][]Block, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *Block) SetOnError(v map[string][]Block)`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *Block) HasOnError() bool`

HasOnError returns a boolean if a field has been set.

### GetName

`func (o *Block) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Block) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Block) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Block) GetValue() RedeemLoyaltyPointsBlock1Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Block) GetValueOk() (*RedeemLoyaltyPointsBlock1Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Block) SetValue(v RedeemLoyaltyPointsBlock1Value)`

SetValue sets Value field to given value.


### GetPartial

`func (o *Block) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *Block) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *Block) SetPartial(v bool)`

SetPartial sets Partial field to given value.


### GetTarget

`func (o *Block) GetTarget() AwardLoyaltyPointsTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Block) GetTargetOk() (*AwardLoyaltyPointsTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Block) SetTarget(v AwardLoyaltyPointsTarget)`

SetTarget sets Target field to given value.


### GetExpression

`func (o *Block) GetExpression() []interface{}`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Block) GetExpressionOk() (*[]interface{}, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Block) SetExpression(v []interface{})`

SetExpression sets Expression field to given value.


### GetNotificationType

`func (o *Block) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *Block) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *Block) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTitle

`func (o *Block) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Block) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Block) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *Block) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Block) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Block) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Block) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetSku

`func (o *Block) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Block) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Block) SetSku(v string)`

SetSku sets Sku field to given value.


### GetQuantity

`func (o *Block) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Block) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Block) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetGiveawayPool

`func (o *Block) GetGiveawayPool() GiveawayPoolBlockReference`

GetGiveawayPool returns the GiveawayPool field if non-nil, zero value otherwise.

### GetGiveawayPoolOk

`func (o *Block) GetGiveawayPoolOk() (*GiveawayPoolBlockReference, bool)`

GetGiveawayPoolOk returns a tuple with the GiveawayPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayPool

`func (o *Block) SetGiveawayPool(v GiveawayPoolBlockReference)`

SetGiveawayPool sets GiveawayPool field to given value.


### GetProfile

`func (o *Block) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Block) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Block) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAudience

`func (o *Block) GetAudience() AudienceBlockReference`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Block) GetAudienceOk() (*AudienceBlockReference, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Block) SetAudience(v AudienceBlockReference)`

SetAudience sets Audience field to given value.


### GetProgram

`func (o *Block) GetProgram() RedeemLoyaltyPointsBlock1Program`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *Block) GetProgramOk() (*RedeemLoyaltyPointsBlock1Program, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *Block) SetProgram(v RedeemLoyaltyPointsBlock1Program)`

SetProgram sets Program field to given value.


### GetSubledger

`func (o *Block) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *Block) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *Block) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetBalance

`func (o *Block) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Block) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Block) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetRedeem

`func (o *Block) GetRedeem() bool`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *Block) GetRedeemOk() (*bool, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *Block) SetRedeem(v bool)`

SetRedeem sets Redeem field to given value.


### GetAchievement

`func (o *Block) GetAchievement() AchievementBlockReference`

GetAchievement returns the Achievement field if non-nil, zero value otherwise.

### GetAchievementOk

`func (o *Block) GetAchievementOk() (*AchievementBlockReference, bool)`

GetAchievementOk returns a tuple with the Achievement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievement

`func (o *Block) SetAchievement(v AchievementBlockReference)`

SetAchievement sets Achievement field to given value.


### GetAttribute

`func (o *Block) GetAttribute() AttributeBlockReference`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Block) GetAttributeOk() (*AttributeBlockReference, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Block) SetAttribute(v AttributeBlockReference)`

SetAttribute sets Attribute field to given value.


### GetWebhook

`func (o *Block) GetWebhook() WebhookBlockReference`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *Block) GetWebhookOk() (*WebhookBlockReference, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *Block) SetWebhook(v WebhookBlockReference)`

SetWebhook sets Webhook field to given value.


### GetParams

`func (o *Block) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Block) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Block) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *Block) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCustomEffect

`func (o *Block) GetCustomEffect() CustomEffectBlockReference`

GetCustomEffect returns the CustomEffect field if non-nil, zero value otherwise.

### GetCustomEffectOk

`func (o *Block) GetCustomEffectOk() (*CustomEffectBlockReference, bool)`

GetCustomEffectOk returns a tuple with the CustomEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEffect

`func (o *Block) SetCustomEffect(v CustomEffectBlockReference)`

SetCustomEffect sets CustomEffect field to given value.


### GetEventType

`func (o *Block) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Block) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Block) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetMatchers

`func (o *Block) GetMatchers() []Block`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *Block) GetMatchersOk() (*[]Block, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *Block) SetMatchers(v []Block)`

SetMatchers sets Matchers field to given value.

### HasMatchers

`func (o *Block) HasMatchers() bool`

HasMatchers returns a boolean if a field has been set.

### GetAction

`func (o *Block) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Block) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Block) SetAction(v string)`

SetAction sets Action field to given value.


### GetCampaignId

`func (o *Block) GetCampaignId() CreateReferralBlock1CampaignId`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Block) GetCampaignIdOk() (*CreateReferralBlock1CampaignId, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Block) SetCampaignId(v CreateReferralBlock1CampaignId)`

SetCampaignId sets CampaignId field to given value.


### GetRecipientId

`func (o *Block) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *Block) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *Block) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetStoreInSession

`func (o *Block) GetStoreInSession() bool`

GetStoreInSession returns the StoreInSession field if non-nil, zero value otherwise.

### GetStoreInSessionOk

`func (o *Block) GetStoreInSessionOk() (*bool, bool)`

GetStoreInSessionOk returns a tuple with the StoreInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInSession

`func (o *Block) SetStoreInSession(v bool)`

SetStoreInSession sets StoreInSession field to given value.


### GetUsageLimit

`func (o *Block) GetUsageLimit() CreateReferralBlock1UsageLimit`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Block) GetUsageLimitOk() (*CreateReferralBlock1UsageLimit, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *Block) SetUsageLimit(v CreateReferralBlock1UsageLimit)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *Block) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetDiscountLimit

`func (o *Block) GetDiscountLimit() CreateCouponBlock1DiscountLimit`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *Block) GetDiscountLimitOk() (*CreateCouponBlock1DiscountLimit, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *Block) SetDiscountLimit(v CreateCouponBlock1DiscountLimit)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *Block) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *Block) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Block) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Block) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Block) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *Block) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Block) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetExpiryDate

`func (o *Block) GetExpiryDate() interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Block) GetExpiryDateOk() (*interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Block) SetExpiryDate(v interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Block) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDateNil

`func (o *Block) SetExpiryDateNil(b bool)`

 SetExpiryDateNil sets the value for ExpiryDate to be an explicit nil

### UnsetExpiryDate
`func (o *Block) UnsetExpiryDate()`

UnsetExpiryDate ensures that no value is present for ExpiryDate, not even an explicit nil
### GetAttributes

`func (o *Block) GetAttributes() interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Block) GetAttributesOk() (*interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Block) SetAttributes(v interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Block) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *Block) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *Block) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetValidCharacters

`func (o *Block) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *Block) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *Block) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *Block) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetPattern

`func (o *Block) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *Block) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *Block) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *Block) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetFriendId

`func (o *Block) GetFriendId() string`

GetFriendId returns the FriendId field if non-nil, zero value otherwise.

### GetFriendIdOk

`func (o *Block) GetFriendIdOk() (*string, bool)`

GetFriendIdOk returns a tuple with the FriendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendId

`func (o *Block) SetFriendId(v string)`

SetFriendId sets FriendId field to given value.


### GetRecipient

`func (o *Block) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Block) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Block) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetTier

`func (o *Block) GetTier() TierBlockReference`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *Block) GetTierOk() (*TierBlockReference, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *Block) SetTier(v TierBlockReference)`

SetTier sets Tier field to given value.


### GetAwaitsActivation

`func (o *Block) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *Block) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *Block) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *Block) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetValidityDuration

`func (o *Block) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *Block) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *Block) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *Block) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetPendingDuration

`func (o *Block) GetPendingDuration() string`

GetPendingDuration returns the PendingDuration field if non-nil, zero value otherwise.

### GetPendingDurationOk

`func (o *Block) GetPendingDurationOk() (*string, bool)`

GetPendingDurationOk returns a tuple with the PendingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDuration

`func (o *Block) SetPendingDuration(v string)`

SetPendingDuration sets PendingDuration field to given value.

### HasPendingDuration

`func (o *Block) HasPendingDuration() bool`

HasPendingDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


