# EffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float32** | The current progress of the customer in the achievement. | 
**Id** | **int64** | The id of the referral code that was redeemed. | 
**RejectionReason** | **string** | The reason why the code was rejected.  - &#x60;AdvocateNotFound&#x60;: The advocate was not found. - &#x60;CampaignLimitReached&#x60;: The campaign-wide referral code redemption limit has been reached. - &#x60;EffectCouldNotBeApplied&#x60;: One of the effects in the campaign wasn&#39;t applied because a limit for that effect was reached (most common use case will be &#x60;setDiscount&#x60; can not be applied because a discount limit is reached). - &#x60;ProfileLimitReached&#x60;: The profile-specific referral code redemption limit has been reached. - &#x60;ReferralCustomerAlreadyReferred&#x60;: The friend is already referred. - &#x60;ReferralExpired&#x60;: The transferred referral code is expired. - &#x60;ReferralLimitReached&#x60;: The referral code redemption limit has been reached. - &#x60;ReferralNotFound&#x60;: The transferred referral code is wrong. - &#x60;ReferralPartOfNotRunningCampaign&#x60;: The campaign the referral code belongs to is currently not active. The campaign ID field shows the ID of that campaign. - &#x60;ReferralRecipientDoesNotMatch&#x60;: The given referral code value does not match the recipient. - &#x60;ReferralRecipientIdSameAsAdvocate&#x60;: The recipient (friend) has the same id as the advocate. - &#x60;ReferralRejectedByCondition&#x60;: The referral code is valid and in an active campaign, but there were other conditions in that campaign&#39;s rules that were not met. - &#x60;ReferralStartDateInFuture&#x60;: The transferred referral code isn&#39;t active yet. - &#x60;ReferralPartOfNotTriggeredCampaign&#x60;: The campaign the referral code belongs to was not triggered during evaluation (an exclusive or stackable campaign). The campaign ID field shows the ID of that campaign. | 
**ConditionIndex** | Pointer to **int64** | The index of the condition that caused the rejection of the referral. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that caused the rejection of the referral. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**CampaignExclusionReason** | Pointer to **string** | The reason why the campaign the referral belongs to was excluded during [campaign evaluation](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation), when &#x60;rejectionReason&#x60; was &#x60;CouponPartOfNotTriggeredCampaign&#x60;. Its possible values are:  - &#x60;CampaignGaveLowerDiscount&#x60;: The required campaign and referral conditions were met, but another campaign in a [Highest discount value](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group offered a higher discount value. - &#x60;CampaignIsNotFirst&#x60;: The campaign was not evaluated because another campaign in a [First campaign](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group was picked and evaluated first. - &#x60;CampaignNotInEvaluationSet&#x60;: The campaign did not meet other evaluation requirements, for example, because the referral is part of an archived campaign. | [optional] 
**ProfileId** | **int64** | The internal ID of the customer profile. | 
**Name** | **string** | The description of this discount. &#x60;#number&#x60; is appended to the name. It is equal to the &#x60;position&#x60; property. | 
**Scope** | Pointer to **string** | The scope of the rolled back discount.  - For a discount per session, it can be one of &#x60;cartItems&#x60;, &#x60;additionalCosts&#x60; or &#x60;sessionTotal&#x60; - For a discount per item, it can be one of &#x60;price&#x60;, &#x60;additionalCosts&#x60; or &#x60;itemTotal&#x60; | [optional] 
**DesiredValue** | Pointer to **float32** | _[(Partial discounts enabled only)](https://docs.talon.one/docs/product/applications/manage-general-settings#partial-discounts)_. The monetary value of the discount to be applied to the additional cost without considering budget limitations. | [optional] 
**Position** | **float32** | The index of the item in the &#x60;cartItem&#x60; object containing the additional cost that this discount applies to. | 
**SubPosition** | Pointer to **float32** | The index of the item unit in its line item. | [optional] 
**TotalDiscount** | Pointer to **float32** | _(Pro rata discounts only)_ The monetary value of the total effective discount | [optional] 
**DesiredTotalDiscount** | Pointer to **float32** | _(Pro rata discounts only)_ The monetary value of the total discount to be applied without considering budget limitations | [optional] 
**BundleIndex** | Pointer to **int64** | The position of the bundle in a list of item bundles created from the same bundle definition. | [optional] 
**BundleName** | Pointer to **string** | The name of the bundle definition. | [optional] 
**TargetedItemPosition** | Pointer to **float32** | _(Discounting individual item in bundles only)_ The index of the targeted bundle item on which the applied discount is based. | [optional] 
**TargetedItemSubPosition** | Pointer to **float32** | _(Discounting individual item in bundles only)_ The sub-position of the targeted bundle item on which the applied discount is based. | [optional] 
**ExcludedFromPriceHistory** | Pointer to **bool** | When set to &#x60;true&#x60;, the applied discount is excluded from the item&#39;s price history. | [optional] 
**AdditionalCostId** | **int64** | The identifier of the additional cost to be discounted. | 
**AdditionalCost** | **string** | The API name of the additional cost to be discounted. | 
**WebhookId** | **float32** | The internal ID of the webhook. | 
**WebhookName** | **string** | The name of the webhook. | 
**ProgramId** | **int64** | ID of the loyalty program that contains these points. | 
**SubLedgerId** | **string** | API name of the loyalty program subledger that contains these points. | 
**RecipientIntegrationId** | **string** | The integration ID of the customer that receives the giveaway. | 
**StartDate** | Pointer to **time.Time** | The date after which the reimbursed points will be valid. | [optional] 
**ExpiryDate** | Pointer to **time.Time** | The date after which the reimbursed points will expire. | [optional] 
**TransactionUUID** | **string** | The identifier of this loyalty point transaction. | 
**CartItemPosition** | Pointer to **float32** | The index of the item in the cart item list to which the custom effect is applied. | [optional] 
**CartItemSubPosition** | Pointer to **float32** | For cart items with quantity &gt; 1, the sub position indicates to which item unit the custom effect is applied.  | [optional] 
**CardIdentifier** | Pointer to **string** | The identifier of the card from which these points were originally deducted. | [optional] 
**AwaitsActivation** | Pointer to **bool** | Indicates whether the points have an action-based start date. This property is returned only for point transactions with an action-based start date. | [optional] 
**ValidityDuration** | Pointer to **string** | The duration for which the points remain active, calculated relative to their start date. | [optional] 
**RuleTitle** | **string** | The title of the rule that triggered the tier upgrade. | 
**PreviousTierName** | Pointer to **string** | The name of the tier from which the user was upgraded. | [optional] 
**NewTierName** | **string** | The name of the tier to which the user has been upgraded. | 
**Sku** | **string** | SKU of the item that needs to be added. | 
**DesiredQuantity** | Pointer to **int64** | The original quantity in case a partial reward was applied. | [optional] 
**NotificationType** | **string** | The type of notification. | 
**Title** | **string** | The title of the notification. | 
**Body** | **string** | The body of the notification. | 
**Path** | **string** | The entity type and the attribute name. | 
**Description** | **string** | Description of the product bundle. | 
**BundleAttributes** | **[]string** | The cart item attributes that determined which items are being bundled together. | 
**ItemsIndices** | **[]float32** | The indices in the cart items array of the bundled items. | 
**PoolId** | **int64** | The internal ID of the giveaway pool. | 
**PoolName** | **string** | The name of the giveaway pool. | 
**GiveawayId** | **int64** | The internal ID of the giveaway. | 
**Code** | **string** | The giveaway code to be rewarded. | 
**Message** | **string** | The error message. | 
**EffectId** | **int64** | The ID of the custom effect that was triggered. | 
**Payload** | **map[string]interface{}** | The JSON payload of the custom effect. | 
**CouponValue** | **string** | The coupon code that was created. | 
**ProfileIntegrationId** | **string** | The ID of the customer profile in the third-party integration platform. | 
**IsNewReservation** | **bool** | Indicates whether this is a new coupon reservation or not. | 
**AudienceId** | Pointer to **int64** | The internal ID of the audience. | [optional] 
**AudienceName** | Pointer to **string** | The name of the audience. | [optional] 
**AchievementId** | **int64** | The internal ID of the achievement. | 
**AchievementName** | **string** | The name of the achievement. | 
**ProgressTrackerId** | **int64** | The internal ID of the achievement progress tracker. | 
**Delta** | **float32** | The value by which the customer&#39;s current progress in the achievement has increased. | 
**Target** | **float32** | The target value to complete the achievement. | 
**IsJustCompleted** | **bool** | Indicates if the customer has completed the achievement in the current session. | 
**DecreaseProgressBy** | **float32** | The value by which the customer&#39;s current progress in the achievement has decreased. | 
**CurrentProgress** | **float32** | The current progress of the customer in the achievement. | 
**ExtensionDuration** | **string** | Time frame by which the expiry date extends.  The time format is either: - immediate, or - an **integer** followed by a letter indicating the time unit.  Examples: &#x60;immediate&#x60;, &#x60;30s&#x60;, &#x60;40m&#x60;, &#x60;1h&#x60;, &#x60;5D&#x60;, &#x60;7W&#x60;, &#x60;10M&#x60;, &#x60;15Y&#x60;.  Available units:  - &#x60;s&#x60;: seconds - &#x60;m&#x60;: minutes - &#x60;h&#x60;: hours - &#x60;D&#x60;: days - &#x60;W&#x60;: weeks - &#x60;M&#x60;: months - &#x60;Y&#x60;: years  You can round certain units up or down: - &#x60;_D&#x60; for rounding down days only. Signifies the start of the day. - &#x60;_U&#x60; for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.  | 
**AffectedTransactions** | Pointer to [**[]LoyaltyLedgerEntryExpiryDateChange**](LoyaltyLedgerEntryExpiryDateChange.md) | List of transactions affected by the expiry date update. | [optional] 
**NewExpiryDate** | **time.Time** | The specified expiry date and time for all active and pending point transactions in the loyalty program subledger. | 

## Methods

### NewEffectProps

`func NewEffectProps(value float32, id int64, rejectionReason string, profileId int64, name string, position float32, additionalCostId int64, additionalCost string, webhookId float32, webhookName string, programId int64, subLedgerId string, recipientIntegrationId string, transactionUUID string, ruleTitle string, newTierName string, sku string, notificationType string, title string, body string, path string, description string, bundleAttributes []string, itemsIndices []float32, poolId int64, poolName string, giveawayId int64, code string, message string, effectId int64, payload map[string]interface{}, couponValue string, profileIntegrationId string, isNewReservation bool, achievementId int64, achievementName string, progressTrackerId int64, delta float32, target float32, isJustCompleted bool, decreaseProgressBy float32, currentProgress float32, extensionDuration string, newExpiryDate time.Time, ) *EffectProps`

NewEffectProps instantiates a new EffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectPropsWithDefaults

`func NewEffectPropsWithDefaults() *EffectProps`

NewEffectPropsWithDefaults instantiates a new EffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *EffectProps) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EffectProps) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EffectProps) SetValue(v float32)`

SetValue sets Value field to given value.


### GetId

`func (o *EffectProps) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EffectProps) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EffectProps) SetId(v int64)`

SetId sets Id field to given value.


### GetRejectionReason

`func (o *EffectProps) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *EffectProps) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *EffectProps) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.


### GetConditionIndex

`func (o *EffectProps) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *EffectProps) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *EffectProps) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *EffectProps) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEffectIndex

`func (o *EffectProps) GetEffectIndex() int64`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *EffectProps) GetEffectIndexOk() (*int64, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectIndex

`func (o *EffectProps) SetEffectIndex(v int64)`

SetEffectIndex sets EffectIndex field to given value.

### HasEffectIndex

`func (o *EffectProps) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### GetDetails

`func (o *EffectProps) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *EffectProps) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *EffectProps) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *EffectProps) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCampaignExclusionReason

`func (o *EffectProps) GetCampaignExclusionReason() string`

GetCampaignExclusionReason returns the CampaignExclusionReason field if non-nil, zero value otherwise.

### GetCampaignExclusionReasonOk

`func (o *EffectProps) GetCampaignExclusionReasonOk() (*string, bool)`

GetCampaignExclusionReasonOk returns a tuple with the CampaignExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignExclusionReason

`func (o *EffectProps) SetCampaignExclusionReason(v string)`

SetCampaignExclusionReason sets CampaignExclusionReason field to given value.

### HasCampaignExclusionReason

`func (o *EffectProps) HasCampaignExclusionReason() bool`

HasCampaignExclusionReason returns a boolean if a field has been set.

### GetProfileId

`func (o *EffectProps) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *EffectProps) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *EffectProps) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.


### GetName

`func (o *EffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *EffectProps) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *EffectProps) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *EffectProps) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *EffectProps) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetDesiredValue

`func (o *EffectProps) GetDesiredValue() float32`

GetDesiredValue returns the DesiredValue field if non-nil, zero value otherwise.

### GetDesiredValueOk

`func (o *EffectProps) GetDesiredValueOk() (*float32, bool)`

GetDesiredValueOk returns a tuple with the DesiredValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredValue

`func (o *EffectProps) SetDesiredValue(v float32)`

SetDesiredValue sets DesiredValue field to given value.

### HasDesiredValue

`func (o *EffectProps) HasDesiredValue() bool`

HasDesiredValue returns a boolean if a field has been set.

### GetPosition

`func (o *EffectProps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EffectProps) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EffectProps) SetPosition(v float32)`

SetPosition sets Position field to given value.


### GetSubPosition

`func (o *EffectProps) GetSubPosition() float32`

GetSubPosition returns the SubPosition field if non-nil, zero value otherwise.

### GetSubPositionOk

`func (o *EffectProps) GetSubPositionOk() (*float32, bool)`

GetSubPositionOk returns a tuple with the SubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubPosition

`func (o *EffectProps) SetSubPosition(v float32)`

SetSubPosition sets SubPosition field to given value.

### HasSubPosition

`func (o *EffectProps) HasSubPosition() bool`

HasSubPosition returns a boolean if a field has been set.

### GetTotalDiscount

`func (o *EffectProps) GetTotalDiscount() float32`

GetTotalDiscount returns the TotalDiscount field if non-nil, zero value otherwise.

### GetTotalDiscountOk

`func (o *EffectProps) GetTotalDiscountOk() (*float32, bool)`

GetTotalDiscountOk returns a tuple with the TotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiscount

`func (o *EffectProps) SetTotalDiscount(v float32)`

SetTotalDiscount sets TotalDiscount field to given value.

### HasTotalDiscount

`func (o *EffectProps) HasTotalDiscount() bool`

HasTotalDiscount returns a boolean if a field has been set.

### GetDesiredTotalDiscount

`func (o *EffectProps) GetDesiredTotalDiscount() float32`

GetDesiredTotalDiscount returns the DesiredTotalDiscount field if non-nil, zero value otherwise.

### GetDesiredTotalDiscountOk

`func (o *EffectProps) GetDesiredTotalDiscountOk() (*float32, bool)`

GetDesiredTotalDiscountOk returns a tuple with the DesiredTotalDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTotalDiscount

`func (o *EffectProps) SetDesiredTotalDiscount(v float32)`

SetDesiredTotalDiscount sets DesiredTotalDiscount field to given value.

### HasDesiredTotalDiscount

`func (o *EffectProps) HasDesiredTotalDiscount() bool`

HasDesiredTotalDiscount returns a boolean if a field has been set.

### GetBundleIndex

`func (o *EffectProps) GetBundleIndex() int64`

GetBundleIndex returns the BundleIndex field if non-nil, zero value otherwise.

### GetBundleIndexOk

`func (o *EffectProps) GetBundleIndexOk() (*int64, bool)`

GetBundleIndexOk returns a tuple with the BundleIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleIndex

`func (o *EffectProps) SetBundleIndex(v int64)`

SetBundleIndex sets BundleIndex field to given value.

### HasBundleIndex

`func (o *EffectProps) HasBundleIndex() bool`

HasBundleIndex returns a boolean if a field has been set.

### GetBundleName

`func (o *EffectProps) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *EffectProps) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *EffectProps) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *EffectProps) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetTargetedItemPosition

`func (o *EffectProps) GetTargetedItemPosition() float32`

GetTargetedItemPosition returns the TargetedItemPosition field if non-nil, zero value otherwise.

### GetTargetedItemPositionOk

`func (o *EffectProps) GetTargetedItemPositionOk() (*float32, bool)`

GetTargetedItemPositionOk returns a tuple with the TargetedItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedItemPosition

`func (o *EffectProps) SetTargetedItemPosition(v float32)`

SetTargetedItemPosition sets TargetedItemPosition field to given value.

### HasTargetedItemPosition

`func (o *EffectProps) HasTargetedItemPosition() bool`

HasTargetedItemPosition returns a boolean if a field has been set.

### GetTargetedItemSubPosition

`func (o *EffectProps) GetTargetedItemSubPosition() float32`

GetTargetedItemSubPosition returns the TargetedItemSubPosition field if non-nil, zero value otherwise.

### GetTargetedItemSubPositionOk

`func (o *EffectProps) GetTargetedItemSubPositionOk() (*float32, bool)`

GetTargetedItemSubPositionOk returns a tuple with the TargetedItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetedItemSubPosition

`func (o *EffectProps) SetTargetedItemSubPosition(v float32)`

SetTargetedItemSubPosition sets TargetedItemSubPosition field to given value.

### HasTargetedItemSubPosition

`func (o *EffectProps) HasTargetedItemSubPosition() bool`

HasTargetedItemSubPosition returns a boolean if a field has been set.

### GetExcludedFromPriceHistory

`func (o *EffectProps) GetExcludedFromPriceHistory() bool`

GetExcludedFromPriceHistory returns the ExcludedFromPriceHistory field if non-nil, zero value otherwise.

### GetExcludedFromPriceHistoryOk

`func (o *EffectProps) GetExcludedFromPriceHistoryOk() (*bool, bool)`

GetExcludedFromPriceHistoryOk returns a tuple with the ExcludedFromPriceHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromPriceHistory

`func (o *EffectProps) SetExcludedFromPriceHistory(v bool)`

SetExcludedFromPriceHistory sets ExcludedFromPriceHistory field to given value.

### HasExcludedFromPriceHistory

`func (o *EffectProps) HasExcludedFromPriceHistory() bool`

HasExcludedFromPriceHistory returns a boolean if a field has been set.

### GetAdditionalCostId

`func (o *EffectProps) GetAdditionalCostId() int64`

GetAdditionalCostId returns the AdditionalCostId field if non-nil, zero value otherwise.

### GetAdditionalCostIdOk

`func (o *EffectProps) GetAdditionalCostIdOk() (*int64, bool)`

GetAdditionalCostIdOk returns a tuple with the AdditionalCostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCostId

`func (o *EffectProps) SetAdditionalCostId(v int64)`

SetAdditionalCostId sets AdditionalCostId field to given value.


### GetAdditionalCost

`func (o *EffectProps) GetAdditionalCost() string`

GetAdditionalCost returns the AdditionalCost field if non-nil, zero value otherwise.

### GetAdditionalCostOk

`func (o *EffectProps) GetAdditionalCostOk() (*string, bool)`

GetAdditionalCostOk returns a tuple with the AdditionalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCost

`func (o *EffectProps) SetAdditionalCost(v string)`

SetAdditionalCost sets AdditionalCost field to given value.


### GetWebhookId

`func (o *EffectProps) GetWebhookId() float32`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *EffectProps) GetWebhookIdOk() (*float32, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *EffectProps) SetWebhookId(v float32)`

SetWebhookId sets WebhookId field to given value.


### GetWebhookName

`func (o *EffectProps) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *EffectProps) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *EffectProps) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.


### GetProgramId

`func (o *EffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *EffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *EffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *EffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *EffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *EffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetRecipientIntegrationId

`func (o *EffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *EffectProps) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *EffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.


### GetStartDate

`func (o *EffectProps) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *EffectProps) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *EffectProps) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *EffectProps) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *EffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *EffectProps) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *EffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *EffectProps) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetTransactionUUID

`func (o *EffectProps) GetTransactionUUID() string`

GetTransactionUUID returns the TransactionUUID field if non-nil, zero value otherwise.

### GetTransactionUUIDOk

`func (o *EffectProps) GetTransactionUUIDOk() (*string, bool)`

GetTransactionUUIDOk returns a tuple with the TransactionUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionUUID

`func (o *EffectProps) SetTransactionUUID(v string)`

SetTransactionUUID sets TransactionUUID field to given value.


### GetCartItemPosition

`func (o *EffectProps) GetCartItemPosition() float32`

GetCartItemPosition returns the CartItemPosition field if non-nil, zero value otherwise.

### GetCartItemPositionOk

`func (o *EffectProps) GetCartItemPositionOk() (*float32, bool)`

GetCartItemPositionOk returns a tuple with the CartItemPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemPosition

`func (o *EffectProps) SetCartItemPosition(v float32)`

SetCartItemPosition sets CartItemPosition field to given value.

### HasCartItemPosition

`func (o *EffectProps) HasCartItemPosition() bool`

HasCartItemPosition returns a boolean if a field has been set.

### GetCartItemSubPosition

`func (o *EffectProps) GetCartItemSubPosition() float32`

GetCartItemSubPosition returns the CartItemSubPosition field if non-nil, zero value otherwise.

### GetCartItemSubPositionOk

`func (o *EffectProps) GetCartItemSubPositionOk() (*float32, bool)`

GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCartItemSubPosition

`func (o *EffectProps) SetCartItemSubPosition(v float32)`

SetCartItemSubPosition sets CartItemSubPosition field to given value.

### HasCartItemSubPosition

`func (o *EffectProps) HasCartItemSubPosition() bool`

HasCartItemSubPosition returns a boolean if a field has been set.

### GetCardIdentifier

`func (o *EffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *EffectProps) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *EffectProps) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *EffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### GetAwaitsActivation

`func (o *EffectProps) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *EffectProps) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *EffectProps) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *EffectProps) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetValidityDuration

`func (o *EffectProps) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *EffectProps) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *EffectProps) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *EffectProps) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetRuleTitle

`func (o *EffectProps) GetRuleTitle() string`

GetRuleTitle returns the RuleTitle field if non-nil, zero value otherwise.

### GetRuleTitleOk

`func (o *EffectProps) GetRuleTitleOk() (*string, bool)`

GetRuleTitleOk returns a tuple with the RuleTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleTitle

`func (o *EffectProps) SetRuleTitle(v string)`

SetRuleTitle sets RuleTitle field to given value.


### GetPreviousTierName

`func (o *EffectProps) GetPreviousTierName() string`

GetPreviousTierName returns the PreviousTierName field if non-nil, zero value otherwise.

### GetPreviousTierNameOk

`func (o *EffectProps) GetPreviousTierNameOk() (*string, bool)`

GetPreviousTierNameOk returns a tuple with the PreviousTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTierName

`func (o *EffectProps) SetPreviousTierName(v string)`

SetPreviousTierName sets PreviousTierName field to given value.

### HasPreviousTierName

`func (o *EffectProps) HasPreviousTierName() bool`

HasPreviousTierName returns a boolean if a field has been set.

### GetNewTierName

`func (o *EffectProps) GetNewTierName() string`

GetNewTierName returns the NewTierName field if non-nil, zero value otherwise.

### GetNewTierNameOk

`func (o *EffectProps) GetNewTierNameOk() (*string, bool)`

GetNewTierNameOk returns a tuple with the NewTierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTierName

`func (o *EffectProps) SetNewTierName(v string)`

SetNewTierName sets NewTierName field to given value.


### GetSku

`func (o *EffectProps) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *EffectProps) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *EffectProps) SetSku(v string)`

SetSku sets Sku field to given value.


### GetDesiredQuantity

`func (o *EffectProps) GetDesiredQuantity() int64`

GetDesiredQuantity returns the DesiredQuantity field if non-nil, zero value otherwise.

### GetDesiredQuantityOk

`func (o *EffectProps) GetDesiredQuantityOk() (*int64, bool)`

GetDesiredQuantityOk returns a tuple with the DesiredQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantity

`func (o *EffectProps) SetDesiredQuantity(v int64)`

SetDesiredQuantity sets DesiredQuantity field to given value.

### HasDesiredQuantity

`func (o *EffectProps) HasDesiredQuantity() bool`

HasDesiredQuantity returns a boolean if a field has been set.

### GetNotificationType

`func (o *EffectProps) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *EffectProps) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *EffectProps) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetTitle

`func (o *EffectProps) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EffectProps) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EffectProps) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *EffectProps) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *EffectProps) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *EffectProps) SetBody(v string)`

SetBody sets Body field to given value.


### GetPath

`func (o *EffectProps) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EffectProps) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EffectProps) SetPath(v string)`

SetPath sets Path field to given value.


### GetDescription

`func (o *EffectProps) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EffectProps) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EffectProps) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBundleAttributes

`func (o *EffectProps) GetBundleAttributes() []string`

GetBundleAttributes returns the BundleAttributes field if non-nil, zero value otherwise.

### GetBundleAttributesOk

`func (o *EffectProps) GetBundleAttributesOk() (*[]string, bool)`

GetBundleAttributesOk returns a tuple with the BundleAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleAttributes

`func (o *EffectProps) SetBundleAttributes(v []string)`

SetBundleAttributes sets BundleAttributes field to given value.


### GetItemsIndices

`func (o *EffectProps) GetItemsIndices() []float32`

GetItemsIndices returns the ItemsIndices field if non-nil, zero value otherwise.

### GetItemsIndicesOk

`func (o *EffectProps) GetItemsIndicesOk() (*[]float32, bool)`

GetItemsIndicesOk returns a tuple with the ItemsIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsIndices

`func (o *EffectProps) SetItemsIndices(v []float32)`

SetItemsIndices sets ItemsIndices field to given value.


### GetPoolId

`func (o *EffectProps) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *EffectProps) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *EffectProps) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolName

`func (o *EffectProps) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *EffectProps) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *EffectProps) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetGiveawayId

`func (o *EffectProps) GetGiveawayId() int64`

GetGiveawayId returns the GiveawayId field if non-nil, zero value otherwise.

### GetGiveawayIdOk

`func (o *EffectProps) GetGiveawayIdOk() (*int64, bool)`

GetGiveawayIdOk returns a tuple with the GiveawayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayId

`func (o *EffectProps) SetGiveawayId(v int64)`

SetGiveawayId sets GiveawayId field to given value.


### GetCode

`func (o *EffectProps) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EffectProps) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EffectProps) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *EffectProps) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EffectProps) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EffectProps) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEffectId

`func (o *EffectProps) GetEffectId() int64`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *EffectProps) GetEffectIdOk() (*int64, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *EffectProps) SetEffectId(v int64)`

SetEffectId sets EffectId field to given value.


### GetPayload

`func (o *EffectProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *EffectProps) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *EffectProps) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetCouponValue

`func (o *EffectProps) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *EffectProps) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *EffectProps) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.


### GetProfileIntegrationId

`func (o *EffectProps) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *EffectProps) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *EffectProps) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetIsNewReservation

`func (o *EffectProps) GetIsNewReservation() bool`

GetIsNewReservation returns the IsNewReservation field if non-nil, zero value otherwise.

### GetIsNewReservationOk

`func (o *EffectProps) GetIsNewReservationOk() (*bool, bool)`

GetIsNewReservationOk returns a tuple with the IsNewReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewReservation

`func (o *EffectProps) SetIsNewReservation(v bool)`

SetIsNewReservation sets IsNewReservation field to given value.


### GetAudienceId

`func (o *EffectProps) GetAudienceId() int64`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *EffectProps) GetAudienceIdOk() (*int64, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *EffectProps) SetAudienceId(v int64)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *EffectProps) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetAudienceName

`func (o *EffectProps) GetAudienceName() string`

GetAudienceName returns the AudienceName field if non-nil, zero value otherwise.

### GetAudienceNameOk

`func (o *EffectProps) GetAudienceNameOk() (*string, bool)`

GetAudienceNameOk returns a tuple with the AudienceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceName

`func (o *EffectProps) SetAudienceName(v string)`

SetAudienceName sets AudienceName field to given value.

### HasAudienceName

`func (o *EffectProps) HasAudienceName() bool`

HasAudienceName returns a boolean if a field has been set.

### GetAchievementId

`func (o *EffectProps) GetAchievementId() int64`

GetAchievementId returns the AchievementId field if non-nil, zero value otherwise.

### GetAchievementIdOk

`func (o *EffectProps) GetAchievementIdOk() (*int64, bool)`

GetAchievementIdOk returns a tuple with the AchievementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementId

`func (o *EffectProps) SetAchievementId(v int64)`

SetAchievementId sets AchievementId field to given value.


### GetAchievementName

`func (o *EffectProps) GetAchievementName() string`

GetAchievementName returns the AchievementName field if non-nil, zero value otherwise.

### GetAchievementNameOk

`func (o *EffectProps) GetAchievementNameOk() (*string, bool)`

GetAchievementNameOk returns a tuple with the AchievementName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementName

`func (o *EffectProps) SetAchievementName(v string)`

SetAchievementName sets AchievementName field to given value.


### GetProgressTrackerId

`func (o *EffectProps) GetProgressTrackerId() int64`

GetProgressTrackerId returns the ProgressTrackerId field if non-nil, zero value otherwise.

### GetProgressTrackerIdOk

`func (o *EffectProps) GetProgressTrackerIdOk() (*int64, bool)`

GetProgressTrackerIdOk returns a tuple with the ProgressTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressTrackerId

`func (o *EffectProps) SetProgressTrackerId(v int64)`

SetProgressTrackerId sets ProgressTrackerId field to given value.


### GetDelta

`func (o *EffectProps) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *EffectProps) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *EffectProps) SetDelta(v float32)`

SetDelta sets Delta field to given value.


### GetTarget

`func (o *EffectProps) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EffectProps) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EffectProps) SetTarget(v float32)`

SetTarget sets Target field to given value.


### GetIsJustCompleted

`func (o *EffectProps) GetIsJustCompleted() bool`

GetIsJustCompleted returns the IsJustCompleted field if non-nil, zero value otherwise.

### GetIsJustCompletedOk

`func (o *EffectProps) GetIsJustCompletedOk() (*bool, bool)`

GetIsJustCompletedOk returns a tuple with the IsJustCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsJustCompleted

`func (o *EffectProps) SetIsJustCompleted(v bool)`

SetIsJustCompleted sets IsJustCompleted field to given value.


### GetDecreaseProgressBy

`func (o *EffectProps) GetDecreaseProgressBy() float32`

GetDecreaseProgressBy returns the DecreaseProgressBy field if non-nil, zero value otherwise.

### GetDecreaseProgressByOk

`func (o *EffectProps) GetDecreaseProgressByOk() (*float32, bool)`

GetDecreaseProgressByOk returns a tuple with the DecreaseProgressBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecreaseProgressBy

`func (o *EffectProps) SetDecreaseProgressBy(v float32)`

SetDecreaseProgressBy sets DecreaseProgressBy field to given value.


### GetCurrentProgress

`func (o *EffectProps) GetCurrentProgress() float32`

GetCurrentProgress returns the CurrentProgress field if non-nil, zero value otherwise.

### GetCurrentProgressOk

`func (o *EffectProps) GetCurrentProgressOk() (*float32, bool)`

GetCurrentProgressOk returns a tuple with the CurrentProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentProgress

`func (o *EffectProps) SetCurrentProgress(v float32)`

SetCurrentProgress sets CurrentProgress field to given value.


### GetExtensionDuration

`func (o *EffectProps) GetExtensionDuration() string`

GetExtensionDuration returns the ExtensionDuration field if non-nil, zero value otherwise.

### GetExtensionDurationOk

`func (o *EffectProps) GetExtensionDurationOk() (*string, bool)`

GetExtensionDurationOk returns a tuple with the ExtensionDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionDuration

`func (o *EffectProps) SetExtensionDuration(v string)`

SetExtensionDuration sets ExtensionDuration field to given value.


### GetAffectedTransactions

`func (o *EffectProps) GetAffectedTransactions() []LoyaltyLedgerEntryExpiryDateChange`

GetAffectedTransactions returns the AffectedTransactions field if non-nil, zero value otherwise.

### GetAffectedTransactionsOk

`func (o *EffectProps) GetAffectedTransactionsOk() (*[]LoyaltyLedgerEntryExpiryDateChange, bool)`

GetAffectedTransactionsOk returns a tuple with the AffectedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedTransactions

`func (o *EffectProps) SetAffectedTransactions(v []LoyaltyLedgerEntryExpiryDateChange)`

SetAffectedTransactions sets AffectedTransactions field to given value.

### HasAffectedTransactions

`func (o *EffectProps) HasAffectedTransactions() bool`

HasAffectedTransactions returns a boolean if a field has been set.

### GetNewExpiryDate

`func (o *EffectProps) GetNewExpiryDate() time.Time`

GetNewExpiryDate returns the NewExpiryDate field if non-nil, zero value otherwise.

### GetNewExpiryDateOk

`func (o *EffectProps) GetNewExpiryDateOk() (*time.Time, bool)`

GetNewExpiryDateOk returns a tuple with the NewExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExpiryDate

`func (o *EffectProps) SetNewExpiryDate(v time.Time)`

SetNewExpiryDate sets NewExpiryDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


