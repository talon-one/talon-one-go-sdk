# \ManagementAPI

All URIs are relative to *https://yourbaseurl.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUserByEmail**](ManagementAPI.md#ActivateUserByEmail) | **Post** /v1/users/activate | Enable user by email address
[**AddLoyaltyCardPoints**](ManagementAPI.md#AddLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/add_points | Add points to card
[**AddLoyaltyPoints**](ManagementAPI.md#AddLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/add_points | Add points to customer profile
[**CopyCampaignToApplications**](ManagementAPI.md#CopyCampaignToApplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into the specified Application
[**CreateAccountCollection**](ManagementAPI.md#CreateAccountCollection) | **Post** /v1/collections | Create account-level collection
[**CreateAchievement**](ManagementAPI.md#CreateAchievement) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | Create achievement
[**CreateAdditionalCost**](ManagementAPI.md#CreateAdditionalCost) | **Post** /v1/additional_costs | Create additional cost
[**CreateAttribute**](ManagementAPI.md#CreateAttribute) | **Post** /v1/attributes | Create custom attribute
[**CreateBatchLoyaltyCards**](ManagementAPI.md#CreateBatchLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards/batch | Create loyalty cards
[**CreateCampaignFromTemplate**](ManagementAPI.md#CreateCampaignFromTemplate) | **Post** /v1/applications/{applicationId}/create_campaign_from_template | Create campaign from campaign template
[**CreateCampaignStoreBudget**](ManagementAPI.md#CreateCampaignStoreBudget) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | Create campaign store budget
[**CreateCollection**](ManagementAPI.md#CreateCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | Create campaign-level collection
[**CreateCoupons**](ManagementAPI.md#CreateCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
[**CreateCouponsAsync**](ManagementAPI.md#CreateCouponsAsync) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_async | Create coupons asynchronously
[**CreateCouponsDeletionJob**](ManagementAPI.md#CreateCouponsDeletionJob) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_deletion_jobs | Creates a coupon deletion job
[**CreateCouponsForMultipleRecipients**](ManagementAPI.md#CreateCouponsForMultipleRecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
[**CreateInviteEmail**](ManagementAPI.md#CreateInviteEmail) | **Post** /v1/invite_emails | Resend invitation email
[**CreateInviteV2**](ManagementAPI.md#CreateInviteV2) | **Post** /v2/invites | Invite user
[**CreatePasswordRecoveryEmail**](ManagementAPI.md#CreatePasswordRecoveryEmail) | **Post** /v1/password_recovery_emails | Request a password reset
[**CreateSession**](ManagementAPI.md#CreateSession) | **Post** /v1/sessions | Create session
[**CreateStore**](ManagementAPI.md#CreateStore) | **Post** /v1/applications/{applicationId}/stores | Create store
[**DeactivateUserByEmail**](ManagementAPI.md#DeactivateUserByEmail) | **Post** /v1/users/deactivate | Disable user by email address
[**DeductLoyaltyCardPoints**](ManagementAPI.md#DeductLoyaltyCardPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/deduct_points | Deduct points from card
[**DeleteAccountCollection**](ManagementAPI.md#DeleteAccountCollection) | **Delete** /v1/collections/{collectionId} | Delete account-level collection
[**DeleteAchievement**](ManagementAPI.md#DeleteAchievement) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Delete achievement
[**DeleteCampaign**](ManagementAPI.md#DeleteCampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
[**DeleteCampaignStoreBudgets**](ManagementAPI.md#DeleteCampaignStoreBudgets) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | Delete campaign store budgets
[**DeleteCollection**](ManagementAPI.md#DeleteCollection) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Delete campaign-level collection
[**DeleteCoupon**](ManagementAPI.md#DeleteCoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
[**DeleteCoupons**](ManagementAPI.md#DeleteCoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
[**DeleteLoyaltyCard**](ManagementAPI.md#DeleteLoyaltyCard) | **Delete** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Delete loyalty card
[**DeleteReferral**](ManagementAPI.md#DeleteReferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
[**DeleteStore**](ManagementAPI.md#DeleteStore) | **Delete** /v1/applications/{applicationId}/stores/{storeId} | Delete store
[**DeleteUser**](ManagementAPI.md#DeleteUser) | **Delete** /v1/users/{userId} | Delete user
[**DeleteUserByEmail**](ManagementAPI.md#DeleteUserByEmail) | **Post** /v1/users/delete | Delete user by email address
[**DestroySession**](ManagementAPI.md#DestroySession) | **Delete** /v1/sessions | Destroy session
[**DisconnectCampaignStores**](ManagementAPI.md#DisconnectCampaignStores) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores | Disconnect stores
[**ExportAccountCollectionItems**](ManagementAPI.md#ExportAccountCollectionItems) | **Get** /v1/collections/{collectionId}/export | Export account-level collection&#39;s items
[**ExportAchievements**](ManagementAPI.md#ExportAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId}/export | Export achievement customer data
[**ExportAudiencesMemberships**](ManagementAPI.md#ExportAudiencesMemberships) | **Get** /v1/audiences/{audienceId}/memberships/export | Export audience members
[**ExportCampaignStoreBudgets**](ManagementAPI.md#ExportCampaignStoreBudgets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/export | Export campaign store budgets
[**ExportCampaignStores**](ManagementAPI.md#ExportCampaignStores) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/export | Export stores
[**ExportCollectionItems**](ManagementAPI.md#ExportCollectionItems) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/export | Export campaign-level collection&#39;s items
[**ExportCoupons**](ManagementAPI.md#ExportCoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons
[**ExportCustomerSessions**](ManagementAPI.md#ExportCustomerSessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions
[**ExportCustomersTiers**](ManagementAPI.md#ExportCustomersTiers) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customers_tiers | Export customers&#39; tier data
[**ExportEffects**](ManagementAPI.md#ExportEffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects
[**ExportLoyaltyBalance**](ManagementAPI.md#ExportLoyaltyBalance) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balance | Export customer loyalty balance to CSV
[**ExportLoyaltyBalances**](ManagementAPI.md#ExportLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balances | Export customer loyalty balances
[**ExportLoyaltyCardBalances**](ManagementAPI.md#ExportLoyaltyCardBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_card_balances | Export all card transaction logs
[**ExportLoyaltyCardLedger**](ManagementAPI.md#ExportLoyaltyCardLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/export_log | Export card&#39;s ledger log
[**ExportLoyaltyCards**](ManagementAPI.md#ExportLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/export | Export loyalty cards
[**ExportLoyaltyLedger**](ManagementAPI.md#ExportLoyaltyLedger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/export_log | Export customer&#39;s transaction logs
[**ExportPoolGiveaways**](ManagementAPI.md#ExportPoolGiveaways) | **Get** /v1/giveaways/pools/{poolId}/export | Export giveaway codes of a giveaway pool
[**ExportReferrals**](ManagementAPI.md#ExportReferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals
[**GenerateCouponRejections**](ManagementAPI.md#GenerateCouponRejections) | **Get** /v1/coupon_rejections | Summarize coupon redemption failures in session
[**GetAccessLogsWithoutTotalCount**](ManagementAPI.md#GetAccessLogsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
[**GetAccount**](ManagementAPI.md#GetAccount) | **Get** /v1/accounts/{accountId} | Get account details
[**GetAccountAnalytics**](ManagementAPI.md#GetAccountAnalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
[**GetAccountCollection**](ManagementAPI.md#GetAccountCollection) | **Get** /v1/collections/{collectionId} | Get account-level collection
[**GetAchievement**](ManagementAPI.md#GetAchievement) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Get achievement
[**GetAdditionalCost**](ManagementAPI.md#GetAdditionalCost) | **Get** /v1/additional_costs/{additionalCostId} | Get additional cost
[**GetAdditionalCosts**](ManagementAPI.md#GetAdditionalCosts) | **Get** /v1/additional_costs | List additional costs
[**GetApplication**](ManagementAPI.md#GetApplication) | **Get** /v1/applications/{applicationId} | Get Application
[**GetApplicationApiHealth**](ManagementAPI.md#GetApplicationApiHealth) | **Get** /v1/applications/{applicationId}/health_report | Get Application health
[**GetApplicationCustomer**](ManagementAPI.md#GetApplicationCustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
[**GetApplicationCustomerFriends**](ManagementAPI.md#GetApplicationCustomerFriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
[**GetApplicationCustomers**](ManagementAPI.md#GetApplicationCustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
[**GetApplicationCustomersByAttributes**](ManagementAPI.md#GetApplicationCustomersByAttributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
[**GetApplicationEventTypes**](ManagementAPI.md#GetApplicationEventTypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
[**GetApplicationEventsWithoutTotalCount**](ManagementAPI.md#GetApplicationEventsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
[**GetApplicationSession**](ManagementAPI.md#GetApplicationSession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
[**GetApplicationSessions**](ManagementAPI.md#GetApplicationSessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
[**GetApplications**](ManagementAPI.md#GetApplications) | **Get** /v1/applications | List Applications
[**GetAttribute**](ManagementAPI.md#GetAttribute) | **Get** /v1/attributes/{attributeId} | Get custom attribute
[**GetAttributes**](ManagementAPI.md#GetAttributes) | **Get** /v1/attributes | List custom attributes
[**GetAudienceMemberships**](ManagementAPI.md#GetAudienceMemberships) | **Get** /v1/audiences/{audienceId}/memberships | List audience members
[**GetAudiences**](ManagementAPI.md#GetAudiences) | **Get** /v1/audiences | List audiences
[**GetAudiencesAnalytics**](ManagementAPI.md#GetAudiencesAnalytics) | **Get** /v1/audiences/analytics | List audience analytics
[**GetCampaign**](ManagementAPI.md#GetCampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
[**GetCampaignAnalytics**](ManagementAPI.md#GetCampaignAnalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
[**GetCampaignByAttributes**](ManagementAPI.md#GetCampaignByAttributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
[**GetCampaignGroup**](ManagementAPI.md#GetCampaignGroup) | **Get** /v1/campaign_groups/{campaignGroupId} | Get campaign access group
[**GetCampaignGroups**](ManagementAPI.md#GetCampaignGroups) | **Get** /v1/campaign_groups | List campaign access groups
[**GetCampaignTemplates**](ManagementAPI.md#GetCampaignTemplates) | **Get** /v1/campaign_templates | List campaign templates
[**GetCampaigns**](ManagementAPI.md#GetCampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
[**GetChanges**](ManagementAPI.md#GetChanges) | **Get** /v1/changes | Get audit logs for an account
[**GetCollection**](ManagementAPI.md#GetCollection) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Get campaign-level collection
[**GetCollectionItems**](ManagementAPI.md#GetCollectionItems) | **Get** /v1/collections/{collectionId}/items | Get collection items
[**GetCouponsWithoutTotalCount**](ManagementAPI.md#GetCouponsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
[**GetCustomerActivityReport**](ManagementAPI.md#GetCustomerActivityReport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
[**GetCustomerActivityReportsWithoutTotalCount**](ManagementAPI.md#GetCustomerActivityReportsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
[**GetCustomerAnalytics**](ManagementAPI.md#GetCustomerAnalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
[**GetCustomerProfile**](ManagementAPI.md#GetCustomerProfile) | **Get** /v1/customers/{customerId} | Get customer profile
[**GetCustomerProfileAchievementProgress**](ManagementAPI.md#GetCustomerProfileAchievementProgress) | **Get** /v1/applications/{applicationId}/achievement_progress/{integrationId} | List customer achievements
[**GetCustomerProfiles**](ManagementAPI.md#GetCustomerProfiles) | **Get** /v1/customers/no_total | List customer profiles
[**GetCustomersByAttributes**](ManagementAPI.md#GetCustomersByAttributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
[**GetDashboardStatistics**](ManagementAPI.md#GetDashboardStatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/dashboard | Get statistics for loyalty dashboard
[**GetEventTypes**](ManagementAPI.md#GetEventTypes) | **Get** /v1/event_types | List event types
[**GetExports**](ManagementAPI.md#GetExports) | **Get** /v1/exports | Get exports
[**GetLoyaltyCard**](ManagementAPI.md#GetLoyaltyCard) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Get loyalty card
[**GetLoyaltyCardTransactionLogs**](ManagementAPI.md#GetLoyaltyCardTransactionLogs) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/logs | List card&#39;s transactions
[**GetLoyaltyCards**](ManagementAPI.md#GetLoyaltyCards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards | List loyalty cards
[**GetLoyaltyLedgerBalances**](ManagementAPI.md#GetLoyaltyLedgerBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/ledger_balances | Get customer&#39;s loyalty balances
[**GetLoyaltyPoints**](ManagementAPI.md#GetLoyaltyPoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId} | Get customer&#39;s full loyalty ledger
[**GetLoyaltyProgram**](ManagementAPI.md#GetLoyaltyProgram) | **Get** /v1/loyalty_programs/{loyaltyProgramId} | Get loyalty program
[**GetLoyaltyProgramProfileLedgerTransactions**](ManagementAPI.md#GetLoyaltyProgramProfileLedgerTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/ledger_transactions | List customer&#39;s loyalty transactions
[**GetLoyaltyProgramTransactions**](ManagementAPI.md#GetLoyaltyProgramTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/transactions | List loyalty program transactions
[**GetLoyaltyPrograms**](ManagementAPI.md#GetLoyaltyPrograms) | **Get** /v1/loyalty_programs | List loyalty programs
[**GetLoyaltyStatistics**](ManagementAPI.md#GetLoyaltyStatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/statistics | Get loyalty program statistics
[**GetMessageLogs**](ManagementAPI.md#GetMessageLogs) | **Get** /v1/message_logs | List message log entries
[**GetReferralsWithoutTotalCount**](ManagementAPI.md#GetReferralsWithoutTotalCount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
[**GetRoleV2**](ManagementAPI.md#GetRoleV2) | **Get** /v2/roles/{roleId} | Get role
[**GetRuleset**](ManagementAPI.md#GetRuleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
[**GetRulesets**](ManagementAPI.md#GetRulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
[**GetStore**](ManagementAPI.md#GetStore) | **Get** /v1/applications/{applicationId}/stores/{storeId} | Get store
[**GetUser**](ManagementAPI.md#GetUser) | **Get** /v1/users/{userId} | Get user
[**GetUsers**](ManagementAPI.md#GetUsers) | **Get** /v1/users | List users in account
[**GetWebhook**](ManagementAPI.md#GetWebhook) | **Get** /v1/webhooks/{webhookId} | Get webhook
[**GetWebhooks**](ManagementAPI.md#GetWebhooks) | **Get** /v1/webhooks | List webhooks
[**ImportAccountCollection**](ManagementAPI.md#ImportAccountCollection) | **Post** /v1/collections/{collectionId}/import | Import data into existing account-level collection
[**ImportAllowedList**](ManagementAPI.md#ImportAllowedList) | **Post** /v1/attributes/{attributeId}/allowed_list/import | Import allowed values for attribute
[**ImportAudiencesMemberships**](ManagementAPI.md#ImportAudiencesMemberships) | **Post** /v1/audiences/{audienceId}/memberships/import | Import audience members
[**ImportCampaignStoreBudget**](ManagementAPI.md#ImportCampaignStoreBudget) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/import | Import campaign store budgets
[**ImportCampaignStores**](ManagementAPI.md#ImportCampaignStores) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/import | Import stores
[**ImportCollection**](ManagementAPI.md#ImportCollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import data into existing campaign-level collection
[**ImportCoupons**](ManagementAPI.md#ImportCoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons
[**ImportLoyaltyCards**](ManagementAPI.md#ImportLoyaltyCards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_cards | Import loyalty cards
[**ImportLoyaltyCustomersTiers**](ManagementAPI.md#ImportLoyaltyCustomersTiers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_customers_tiers | Import customers into loyalty tiers
[**ImportLoyaltyPoints**](ManagementAPI.md#ImportLoyaltyPoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_points | Import loyalty points
[**ImportPoolGiveaways**](ManagementAPI.md#ImportPoolGiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
[**ImportReferrals**](ManagementAPI.md#ImportReferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals
[**InviteUserExternal**](ManagementAPI.md#InviteUserExternal) | **Post** /v1/users/invite | Invite user from identity provider
[**ListAccountCollections**](ManagementAPI.md#ListAccountCollections) | **Get** /v1/collections | List collections in account
[**ListAchievements**](ManagementAPI.md#ListAchievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | List achievements
[**ListAllRolesV2**](ManagementAPI.md#ListAllRolesV2) | **Get** /v2/roles | List roles
[**ListCampaignStoreBudgetLimits**](ManagementAPI.md#ListCampaignStoreBudgetLimits) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | List campaign store budget limits
[**ListCatalogItems**](ManagementAPI.md#ListCatalogItems) | **Get** /v1/catalogs/{catalogId}/items | List items in a catalog
[**ListCollections**](ManagementAPI.md#ListCollections) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | List collections in campaign
[**ListCollectionsInApplication**](ManagementAPI.md#ListCollectionsInApplication) | **Get** /v1/applications/{applicationId}/collections | List collections in Application
[**ListStores**](ManagementAPI.md#ListStores) | **Get** /v1/applications/{applicationId}/stores | List stores
[**OktaEventHandlerChallenge**](ManagementAPI.md#OktaEventHandlerChallenge) | **Get** /v1/provisioning/okta | Validate Okta API ownership
[**RemoveLoyaltyPoints**](ManagementAPI.md#RemoveLoyaltyPoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/deduct_points | Deduct points from customer profile
[**ResetPassword**](ManagementAPI.md#ResetPassword) | **Post** /v1/reset_password | Reset password
[**ScimCreateGroup**](ManagementAPI.md#ScimCreateGroup) | **Post** /v1/provisioning/scim/Groups | Create SCIM group
[**ScimCreateUser**](ManagementAPI.md#ScimCreateUser) | **Post** /v1/provisioning/scim/Users | Create SCIM user
[**ScimDeleteGroup**](ManagementAPI.md#ScimDeleteGroup) | **Delete** /v1/provisioning/scim/Groups/{groupId} | Delete SCIM group
[**ScimDeleteUser**](ManagementAPI.md#ScimDeleteUser) | **Delete** /v1/provisioning/scim/Users/{userId} | Delete SCIM user
[**ScimGetGroup**](ManagementAPI.md#ScimGetGroup) | **Get** /v1/provisioning/scim/Groups/{groupId} | Get SCIM group
[**ScimGetGroups**](ManagementAPI.md#ScimGetGroups) | **Get** /v1/provisioning/scim/Groups | List SCIM groups
[**ScimGetResourceTypes**](ManagementAPI.md#ScimGetResourceTypes) | **Get** /v1/provisioning/scim/ResourceTypes | List supported SCIM resource types
[**ScimGetSchemas**](ManagementAPI.md#ScimGetSchemas) | **Get** /v1/provisioning/scim/Schemas | List supported SCIM schemas
[**ScimGetServiceProviderConfig**](ManagementAPI.md#ScimGetServiceProviderConfig) | **Get** /v1/provisioning/scim/ServiceProviderConfig | Get SCIM service provider configuration
[**ScimGetUser**](ManagementAPI.md#ScimGetUser) | **Get** /v1/provisioning/scim/Users/{userId} | Get SCIM user
[**ScimGetUsers**](ManagementAPI.md#ScimGetUsers) | **Get** /v1/provisioning/scim/Users | List SCIM users
[**ScimPatchGroup**](ManagementAPI.md#ScimPatchGroup) | **Patch** /v1/provisioning/scim/Groups/{groupId} | Update SCIM group attributes
[**ScimPatchUser**](ManagementAPI.md#ScimPatchUser) | **Patch** /v1/provisioning/scim/Users/{userId} | Update SCIM user attributes
[**ScimReplaceGroupAttributes**](ManagementAPI.md#ScimReplaceGroupAttributes) | **Put** /v1/provisioning/scim/Groups/{groupId} | Update SCIM group
[**ScimReplaceUserAttributes**](ManagementAPI.md#ScimReplaceUserAttributes) | **Put** /v1/provisioning/scim/Users/{userId} | Update SCIM user
[**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](ManagementAPI.md#SearchCouponsAdvancedApplicationWideWithoutTotalCount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes (without total count)
[**SearchCouponsAdvancedWithoutTotalCount**](ManagementAPI.md#SearchCouponsAdvancedWithoutTotalCount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes in campaign (without total count)
[**SummarizeCampaignStoreBudget**](ManagementAPI.md#SummarizeCampaignStoreBudget) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/summary | Get summary of campaign store budgets
[**TransferLoyaltyCard**](ManagementAPI.md#TransferLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transfer | Transfer card data
[**UpdateAccountCollection**](ManagementAPI.md#UpdateAccountCollection) | **Put** /v1/collections/{collectionId} | Update account-level collection
[**UpdateAchievement**](ManagementAPI.md#UpdateAchievement) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Update achievement
[**UpdateAdditionalCost**](ManagementAPI.md#UpdateAdditionalCost) | **Put** /v1/additional_costs/{additionalCostId} | Update additional cost
[**UpdateAttribute**](ManagementAPI.md#UpdateAttribute) | **Put** /v1/attributes/{attributeId} | Update custom attribute
[**UpdateCampaign**](ManagementAPI.md#UpdateCampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
[**UpdateCollection**](ManagementAPI.md#UpdateCollection) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Update campaign-level collection&#39;s description
[**UpdateCoupon**](ManagementAPI.md#UpdateCoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
[**UpdateCouponBatch**](ManagementAPI.md#UpdateCouponBatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update coupons
[**UpdateLoyaltyCard**](ManagementAPI.md#UpdateLoyaltyCard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Update loyalty card status
[**UpdateReferral**](ManagementAPI.md#UpdateReferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
[**UpdateRoleV2**](ManagementAPI.md#UpdateRoleV2) | **Put** /v2/roles/{roleId} | Update role
[**UpdateStore**](ManagementAPI.md#UpdateStore) | **Put** /v1/applications/{applicationId}/stores/{storeId} | Update store
[**UpdateUser**](ManagementAPI.md#UpdateUser) | **Put** /v1/users/{userId} | Update user



## ActivateUserByEmail

> ActivateUserByEmail(ctx).ActivateUserRequest(activateUserRequest).Execute()

Enable user by email address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	activateUserRequest := *openapiclient.NewActivateUserRequest("john.doe@example.com") // ActivateUserRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.ActivateUserByEmail(context.Background()).ActivateUserRequest(activateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ActivateUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activateUserRequest** | [**ActivateUserRequest**](ActivateUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyCardPoints

> AddLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).AddLoyaltyPoints(addLoyaltyPoints).Execute()

Add points to card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	addLoyaltyPoints := *openapiclient.NewAddLoyaltyPoints(float32(300)) // AddLoyaltyPoints | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.AddLoyaltyCardPoints(context.Background(), loyaltyProgramId, loyaltyCardId).AddLoyaltyPoints(addLoyaltyPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.AddLoyaltyCardPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addLoyaltyPoints** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLoyaltyPoints

> AddLoyaltyPoints(ctx, loyaltyProgramId, integrationId).AddLoyaltyPoints(addLoyaltyPoints).Execute()

Add points to customer profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	addLoyaltyPoints := *openapiclient.NewAddLoyaltyPoints(float32(300)) // AddLoyaltyPoints | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.AddLoyaltyPoints(context.Background(), loyaltyProgramId, integrationId).AddLoyaltyPoints(addLoyaltyPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.AddLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addLoyaltyPoints** | [**AddLoyaltyPoints**](AddLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyCampaignToApplications

> GetCampaigns200Response CopyCampaignToApplications(ctx, applicationId, campaignId).CampaignCopy(campaignCopy).Execute()

Copy the campaign into the specified Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	campaignCopy := *openapiclient.NewCampaignCopy([]int64{int64(123)}) // CampaignCopy | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CopyCampaignToApplications(context.Background(), applicationId, campaignId).CampaignCopy(campaignCopy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CopyCampaignToApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyCampaignToApplications`: GetCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CopyCampaignToApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyCampaignToApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **campaignCopy** | [**CampaignCopy**](CampaignCopy.md) | body | 

### Return type

[**GetCampaigns200Response**](GetCampaigns200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccountCollection

> Collection CreateAccountCollection(ctx).NewCollection(newCollection).Execute()

Create account-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newCollection := *openapiclient.NewNewCollection("My collection") // NewCollection | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateAccountCollection(context.Background()).NewCollection(newCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateAccountCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccountCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateAccountCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newCollection** | [**NewCollection**](NewCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAchievement

> Achievement CreateAchievement(ctx, applicationId, campaignId).CreateAchievement(createAchievement).Execute()

Create achievement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	createAchievement := *openapiclient.NewCreateAchievement("Order50Discount", "50% off on 50th purchase.", "50% off for every 50th purchase in a year.", float32(50)) // CreateAchievement | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateAchievement(context.Background(), applicationId, campaignId).CreateAchievement(createAchievement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateAchievement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAchievement`: Achievement
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateAchievement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createAchievement** | [**CreateAchievement**](CreateAchievement.md) | body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdditionalCost

> AccountAdditionalCost CreateAdditionalCost(ctx).NewAdditionalCost(newAdditionalCost).Execute()

Create additional cost



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newAdditionalCost := *openapiclient.NewNewAdditionalCost("shippingFee", "Shipping fee", "A shipping fee") // NewAdditionalCost | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateAdditionalCost(context.Background()).NewAdditionalCost(newAdditionalCost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateAdditionalCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAdditionalCost`: AccountAdditionalCost
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateAdditionalCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAdditionalCost** | [**NewAdditionalCost**](NewAdditionalCost.md) | body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAttribute

> Attribute CreateAttribute(ctx).NewAttribute(newAttribute).Execute()

Create custom attribute



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newAttribute := *openapiclient.NewNewAttribute("Event", "pageViewed", "Page view event", "string", "Event triggered when a customer displays a page.", []string{"Suggestions_example"}, true) // NewAttribute | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateAttribute(context.Background()).NewAttribute(newAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAttribute`: Attribute
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateAttribute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAttribute** | [**NewAttribute**](NewAttribute.md) | body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBatchLoyaltyCards

> LoyaltyCardBatchResponse CreateBatchLoyaltyCards(ctx, loyaltyProgramId).LoyaltyCardBatch(loyaltyCardBatch).Execute()

Create loyalty cards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardBatch := *openapiclient.NewLoyaltyCardBatch(int64(5000)) // LoyaltyCardBatch | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateBatchLoyaltyCards(context.Background(), loyaltyProgramId).LoyaltyCardBatch(loyaltyCardBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateBatchLoyaltyCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBatchLoyaltyCards`: LoyaltyCardBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateBatchLoyaltyCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBatchLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loyaltyCardBatch** | [**LoyaltyCardBatch**](LoyaltyCardBatch.md) | body | 

### Return type

[**LoyaltyCardBatchResponse**](LoyaltyCardBatchResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignFromTemplate

> CreateTemplateCampaignResponse CreateCampaignFromTemplate(ctx, applicationId).CreateTemplateCampaign(createTemplateCampaign).Execute()

Create campaign from campaign template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	createTemplateCampaign := *openapiclient.NewCreateTemplateCampaign("Discount campaign", int64(4)) // CreateTemplateCampaign | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCampaignFromTemplate(context.Background(), applicationId).CreateTemplateCampaign(createTemplateCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCampaignFromTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaignFromTemplate`: CreateTemplateCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCampaignFromTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignFromTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTemplateCampaign** | [**CreateTemplateCampaign**](CreateTemplateCampaign.md) | body | 

### Return type

[**CreateTemplateCampaignResponse**](CreateTemplateCampaignResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignStoreBudget

> CreateCampaignStoreBudget(ctx, applicationId, campaignId).NewCampaignStoreBudget(newCampaignStoreBudget).Execute()

Create campaign store budget



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCampaignStoreBudget := *openapiclient.NewNewCampaignStoreBudget("Action_example", []openapiclient.NewCampaignStoreBudgetStoreLimit{*openapiclient.NewNewCampaignStoreBudgetStoreLimit(int64(17), float32(1000))}) // NewCampaignStoreBudget | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.CreateCampaignStoreBudget(context.Background(), applicationId, campaignId).NewCampaignStoreBudget(newCampaignStoreBudget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCampaignStoreBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignStoreBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCampaignStoreBudget** | [**NewCampaignStoreBudget**](NewCampaignStoreBudget.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollection

> Collection CreateCollection(ctx, applicationId, campaignId).NewCampaignCollection(newCampaignCollection).Execute()

Create campaign-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCampaignCollection := *openapiclient.NewNewCampaignCollection("My collection") // NewCampaignCollection | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCollection(context.Background(), applicationId, campaignId).NewCampaignCollection(newCampaignCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCampaignCollection** | [**NewCampaignCollection**](NewCampaignCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCoupons

> CreateCoupons200Response CreateCoupons(ctx, applicationId, campaignId).NewCoupons(newCoupons).Silent(silent).Execute()

Create coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCoupons := *openapiclient.NewNewCoupons(int64(1)) // NewCoupons | body
	silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the performance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains the updated customer profiles.  (optional) (default to "yes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCoupons(context.Background(), applicationId, campaignId).NewCoupons(newCoupons).Silent(silent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCoupons`: CreateCoupons200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCoupons** | [**NewCoupons**](NewCoupons.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to &quot;yes&quot;]

### Return type

[**CreateCoupons200Response**](CreateCoupons200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsAsync

> AsyncCouponCreationResponse CreateCouponsAsync(ctx, applicationId, campaignId).NewCouponCreationJob(newCouponCreationJob).Execute()

Create coupons asynchronously



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCouponCreationJob := *openapiclient.NewNewCouponCreationJob(int64(200000), map[string]interface{}(123)) // NewCouponCreationJob | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCouponsAsync(context.Background(), applicationId, campaignId).NewCouponCreationJob(newCouponCreationJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCouponsAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCouponsAsync`: AsyncCouponCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCouponsAsync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCouponCreationJob** | [**NewCouponCreationJob**](NewCouponCreationJob.md) | body | 

### Return type

[**AsyncCouponCreationResponse**](AsyncCouponCreationResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsDeletionJob

> AsyncCouponDeletionJobResponse CreateCouponsDeletionJob(ctx, applicationId, campaignId).NewCouponDeletionJob(newCouponDeletionJob).Execute()

Creates a coupon deletion job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCouponDeletionJob := *openapiclient.NewNewCouponDeletionJob(*openapiclient.NewCouponDeletionFilters()) // NewCouponDeletionJob | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCouponsDeletionJob(context.Background(), applicationId, campaignId).NewCouponDeletionJob(newCouponDeletionJob).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCouponsDeletionJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCouponsDeletionJob`: AsyncCouponDeletionJobResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCouponsDeletionJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsDeletionJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCouponDeletionJob** | [**NewCouponDeletionJob**](NewCouponDeletionJob.md) | body | 

### Return type

[**AsyncCouponDeletionJobResponse**](AsyncCouponDeletionJobResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponsForMultipleRecipients

> CreateCoupons200Response CreateCouponsForMultipleRecipients(ctx, applicationId, campaignId).NewCouponsForMultipleRecipients(newCouponsForMultipleRecipients).Silent(silent).Execute()

Create coupons for multiple recipients



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	newCouponsForMultipleRecipients := *openapiclient.NewNewCouponsForMultipleRecipients([]string{"RecipientsIntegrationIds_example"}) // NewCouponsForMultipleRecipients | body
	silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the performance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains the updated customer profiles.  (optional) (default to "yes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateCouponsForMultipleRecipients(context.Background(), applicationId, campaignId).NewCouponsForMultipleRecipients(newCouponsForMultipleRecipients).Silent(silent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateCouponsForMultipleRecipients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCouponsForMultipleRecipients`: CreateCoupons200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateCouponsForMultipleRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponsForMultipleRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newCouponsForMultipleRecipients** | [**NewCouponsForMultipleRecipients**](NewCouponsForMultipleRecipients.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to &quot;yes&quot;]

### Return type

[**CreateCoupons200Response**](CreateCoupons200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteEmail

> NewInviteEmail CreateInviteEmail(ctx).NewInviteEmail(newInviteEmail).Execute()

Resend invitation email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newInviteEmail := *openapiclient.NewNewInviteEmail("john.doe@example.com", "Gy9b8w1irmQtEPo5RmbMmSPheL5h4") // NewInviteEmail | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateInviteEmail(context.Background()).NewInviteEmail(newInviteEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateInviteEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInviteEmail`: NewInviteEmail
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateInviteEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newInviteEmail** | [**NewInviteEmail**](NewInviteEmail.md) | body | 

### Return type

[**NewInviteEmail**](NewInviteEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInviteV2

> User CreateInviteV2(ctx).NewInvitation(newInvitation).Execute()

Invite user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newInvitation := *openapiclient.NewNewInvitation("john.doe@example.com") // NewInvitation | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateInviteV2(context.Background()).NewInvitation(newInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateInviteV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInviteV2`: User
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateInviteV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInviteV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newInvitation** | [**NewInvitation**](NewInvitation.md) | body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePasswordRecoveryEmail

> NewPasswordEmail CreatePasswordRecoveryEmail(ctx).NewPasswordEmail(newPasswordEmail).Execute()

Request a password reset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newPasswordEmail := *openapiclient.NewNewPasswordEmail("Email_example") // NewPasswordEmail | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreatePasswordRecoveryEmail(context.Background()).NewPasswordEmail(newPasswordEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreatePasswordRecoveryEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePasswordRecoveryEmail`: NewPasswordEmail
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreatePasswordRecoveryEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePasswordRecoveryEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newPasswordEmail** | [**NewPasswordEmail**](NewPasswordEmail.md) | body | 

### Return type

[**NewPasswordEmail**](NewPasswordEmail.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSession

> Session CreateSession(ctx).LoginParams(loginParams).Execute()

Create session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loginParams := *openapiclient.NewLoginParams("john.doe@example.com", "admin123456") // LoginParams | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateSession(context.Background()).LoginParams(loginParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSession`: Session
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateSession`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginParams** | [**LoginParams**](LoginParams.md) | body | 

### Return type

[**Session**](Session.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStore

> Store CreateStore(ctx, applicationId).NewStore(newStore).Execute()

Create store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	newStore := *openapiclient.NewNewStore("South US store", "This is the description of the store in south US.", "STORE-001") // NewStore | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.CreateStore(context.Background(), applicationId).NewStore(newStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CreateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStore`: Store
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.CreateStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newStore** | [**NewStore**](NewStore.md) | body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateUserByEmail

> DeactivateUserByEmail(ctx).DeactivateUserRequest(deactivateUserRequest).Execute()

Disable user by email address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	deactivateUserRequest := *openapiclient.NewDeactivateUserRequest("john.doe@example.com") // DeactivateUserRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeactivateUserByEmail(context.Background()).DeactivateUserRequest(deactivateUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeactivateUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deactivateUserRequest** | [**DeactivateUserRequest**](DeactivateUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeductLoyaltyCardPoints

> DeductLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).DeductLoyaltyPoints(deductLoyaltyPoints).Execute()

Deduct points from card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	deductLoyaltyPoints := *openapiclient.NewDeductLoyaltyPoints(float32(300)) // DeductLoyaltyPoints | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeductLoyaltyCardPoints(context.Background(), loyaltyProgramId, loyaltyCardId).DeductLoyaltyPoints(deductLoyaltyPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeductLoyaltyCardPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeductLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deductLoyaltyPoints** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccountCollection

> DeleteAccountCollection(ctx, collectionId).Execute()

Delete account-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteAccountCollection(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteAccountCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAchievement

> DeleteAchievement(ctx, applicationId, campaignId, achievementId).Execute()

Delete achievement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	achievementId := int64(789) // int64 | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteAchievement(context.Background(), applicationId, campaignId, achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteAchievement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, applicationId, campaignId).Execute()

Delete campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteCampaign(context.Background(), applicationId, campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignStoreBudgets

> DeleteCampaignStoreBudgets(ctx, applicationId, campaignId).Action(action).Period(period).Execute()

Delete campaign store budgets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	action := "action_example" // string | The action that this budget is limiting. (optional)
	period := "period_example" // string | The period to which the limit applies.  **Note**: For budgets with no period, set this to `overall`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteCampaignStoreBudgets(context.Background(), applicationId, campaignId).Action(action).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteCampaignStoreBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignStoreBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **string** | The action that this budget is limiting. | 
 **period** | **string** | The period to which the limit applies.  **Note**: For budgets with no period, set this to &#x60;overall&#x60;.  | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, applicationId, campaignId, collectionId).Execute()

Delete campaign-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteCollection(context.Background(), applicationId, campaignId, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupon

> DeleteCoupon(ctx, applicationId, campaignId, couponId).Execute()

Delete coupon



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	couponId := "couponId_example" // string | The internal ID of the coupon code. You can find this value in the `id` property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteCoupon(context.Background(), applicationId, campaignId, couponId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string** | The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCoupons

> DeleteCoupons(ctx, applicationId, campaignId).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartsAfter(startsAfter).StartsBefore(startsBefore).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).Valid(valid).BatchId(batchId).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).Execute()

Delete coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	value := "value_example" // string | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startsAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startsBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	expiresAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	expiresBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | - `expired`: Matches coupons in which the expiration date is set and in the past. - `validNow`: Matches coupons in which start date is null or in the past and expiration date is null or in the future. - `validFuture`: Matches coupons in which start date is set and in the future.  (optional)
	batchId := "batchId_example" // string | Filter results by batches of coupons (optional)
	usable := "usable_example" // string | - `true`: only coupons where `usageCounter < usageLimit` will be returned. - `false`: only coupons where `usageCounter >= usageLimit` will be returned.  (optional)
	referralId := int64(789) // int64 | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. (optional)
	recipientIntegrationId := "recipientIntegrationId_example" // string | Filter results by match with a profile ID specified in the coupon's `RecipientIntegrationId` field.  (optional)
	exactMatch := true // bool | Filter results to an exact case-insensitive matching against the coupon code (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteCoupons(context.Background(), applicationId, campaignId).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartsAfter(startsAfter).StartsBefore(startsBefore).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).Valid(valid).BatchId(batchId).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | - &#x60;expired&#x60;: Matches coupons in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches coupons in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches coupons in which start date is set and in the future.  | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **usable** | **string** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **referralId** | **int64** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile ID specified in the coupon&#39;s &#x60;RecipientIntegrationId&#x60; field.  | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code | [default to false]

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoyaltyCard

> DeleteLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Execute()

Delete loyalty card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteLoyaltyCard(context.Background(), loyaltyProgramId, loyaltyCardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteLoyaltyCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteReferral

> DeleteReferral(ctx, applicationId, campaignId, referralId).Execute()

Delete referral



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	referralId := "referralId_example" // string | The ID of the referral code.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteReferral(context.Background(), applicationId, campaignId, referralId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string** | The ID of the referral code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStore

> DeleteStore(ctx, applicationId, storeId).Execute()

Delete store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	storeId := "storeId_example" // string | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteStore(context.Background(), applicationId, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, userId).Execute()

Delete user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserByEmail

> DeleteUserByEmail(ctx).DeleteUserRequest(deleteUserRequest).Execute()

Delete user by email address



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	deleteUserRequest := *openapiclient.NewDeleteUserRequest("john.doe@example.com") // DeleteUserRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DeleteUserByEmail(context.Background()).DeleteUserRequest(deleteUserRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DeleteUserByEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserByEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteUserRequest** | [**DeleteUserRequest**](DeleteUserRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySession

> DestroySession(ctx).Execute()

Destroy session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DestroySession(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DestroySession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySessionRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectCampaignStores

> DisconnectCampaignStores(ctx, applicationId, campaignId).Execute()

Disconnect stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.DisconnectCampaignStores(context.Background(), applicationId, campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.DisconnectCampaignStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAccountCollectionItems

> string ExportAccountCollectionItems(ctx, collectionId).Execute()

Export account-level collection's items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportAccountCollectionItems(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportAccountCollectionItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAccountCollectionItems`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportAccountCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAccountCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAchievements

> string ExportAchievements(ctx, applicationId, campaignId, achievementId).Execute()

Export achievement customer data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	achievementId := int64(789) // int64 | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportAchievements(context.Background(), applicationId, campaignId, achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportAchievements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAchievements`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportAchievements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAchievementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportAudiencesMemberships

> string ExportAudiencesMemberships(ctx, audienceId).Execute()

Export audience members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	audienceId := int64(789) // int64 | The ID of the audience.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportAudiencesMemberships(context.Background(), audienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportAudiencesMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportAudiencesMemberships`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportAudiencesMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAudiencesMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCampaignStoreBudgets

> string ExportCampaignStoreBudgets(ctx, applicationId, campaignId).Action(action).Period(period).Execute()

Export campaign store budgets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	action := "action_example" // string | The action that this budget is limiting. (optional)
	period := "period_example" // string | The period to which the limit applies.  **Note**: For budgets with no period, set this to `overall`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCampaignStoreBudgets(context.Background(), applicationId, campaignId).Action(action).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCampaignStoreBudgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCampaignStoreBudgets`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCampaignStoreBudgets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCampaignStoreBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **string** | The action that this budget is limiting. | 
 **period** | **string** | The period to which the limit applies.  **Note**: For budgets with no period, set this to &#x60;overall&#x60;.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCampaignStores

> string ExportCampaignStores(ctx, applicationId, campaignId).Execute()

Export stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCampaignStores(context.Background(), applicationId, campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCampaignStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCampaignStores`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCampaignStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCollectionItems

> string ExportCollectionItems(ctx, applicationId, campaignId, collectionId).Execute()

Export campaign-level collection's items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCollectionItems(context.Background(), applicationId, campaignId, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCollectionItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCollectionItems`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCoupons

> string ExportCoupons(ctx, applicationId).CampaignId(campaignId).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).DateFormat(dateFormat).CampaignState(campaignState).ValuesOnly(valuesOnly).Execute()

Export coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := float32(8.14) // float32 | Filter results by campaign ID. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	value := "value_example" // string | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | Either \"expired\", \"validNow\", or \"validFuture\". The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | Either \"true\" or \"false\". If \"true\", only coupons where `usageCounter < usageLimit` will be returned, \"false\" will return only coupons where `usageCounter >= usageLimit`.  (optional)
	referralId := int64(789) // int64 | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. (optional)
	recipientIntegrationId := "recipientIntegrationId_example" // string | Filter results by match with a profile id specified in the coupon's RecipientIntegrationId field. (optional)
	batchId := "batchId_example" // string | Filter results by batches of coupons (optional)
	exactMatch := true // bool | Filter results to an exact case-insensitive matching against the coupon code. (optional) (default to false)
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)
	campaignState := "campaignState_example" // string | Filter results by the state of the campaign.  - `enabled`: Campaigns that are scheduled, running (activated), or expired. - `running`: Campaigns that are running (activated). - `disabled`: Campaigns that are disabled. - `expired`: Campaigns that are expired. - `archived`: Campaigns that are archived.  (optional)
	valuesOnly := true // bool | Filter results to only return the coupon codes (`value` column) without the associated coupon data. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCoupons(context.Background(), applicationId).CampaignId(campaignId).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).DateFormat(dateFormat).CampaignState(campaignState).ValuesOnly(valuesOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCoupons`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign ID. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int64** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile id specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **dateFormat** | **string** | Determines the format of dates in the export document. | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **valuesOnly** | **bool** | Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomerSessions

> string ExportCustomerSessions(ctx, applicationId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ProfileIntegrationId(profileIntegrationId).DateFormat(dateFormat).CustomerSessionState(customerSessionState).Execute()

Export customer sessions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. (optional)
	profileIntegrationId := "profileIntegrationId_example" // string | Only return sessions for the customer that matches this customer integration ID. (optional)
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)
	customerSessionState := "customerSessionState_example" // string | Filter results by state. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCustomerSessions(context.Background(), applicationId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).ProfileIntegrationId(profileIntegrationId).DateFormat(dateFormat).CustomerSessionState(customerSessionState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCustomerSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCustomerSessions`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCustomerSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCustomerSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. | 
 **profileIntegrationId** | **string** | Only return sessions for the customer that matches this customer integration ID. | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 
 **customerSessionState** | **string** | Filter results by state. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportCustomersTiers

> string ExportCustomersTiers(ctx, loyaltyProgramId).SubledgerIds(subledgerIds).TierNames(tierNames).Execute()

Export customers' tier data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	subledgerIds := []string{"Inner_example"} // []string | An array of subledgers IDs to filter the export by. (optional)
	tierNames := []string{"Inner_example"} // []string | An array of tier names to filter the export by. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportCustomersTiers(context.Background(), loyaltyProgramId).SubledgerIds(subledgerIds).TierNames(tierNames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportCustomersTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportCustomersTiers`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportCustomersTiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportCustomersTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subledgerIds** | **[]string** | An array of subledgers IDs to filter the export by. | 
 **tierNames** | **[]string** | An array of tier names to filter the export by. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportEffects

> string ExportEffects(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()

Export triggered effects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := float32(8.14) // float32 | Filter results by campaign ID. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportEffects(context.Background(), applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportEffects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportEffects`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportEffects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportEffectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign ID. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalance

> string ExportLoyaltyBalance(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export customer loyalty balance to CSV



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	endDate := time.Now() // time.Time | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyBalance(context.Background(), loyaltyProgramId).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyBalance`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyBalances

> string ExportLoyaltyBalances(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export customer loyalty balances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	endDate := time.Now() // time.Time | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. - This parameter does not affect the `currentTier` field in the CSV file,  which shows the customer's tier at the time of export.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyBalances(context.Background(), loyaltyProgramId).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyBalances`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered. - This parameter does not affect the &#x60;currentTier&#x60; field in the CSV file,  which shows the customer&#39;s tier at the time of export.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardBalances

> string ExportLoyaltyCardBalances(ctx, loyaltyProgramId).EndDate(endDate).Execute()

Export all card transaction logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	endDate := time.Now() // time.Time | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyCardBalances(context.Background(), loyaltyProgramId).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyCardBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyCardBalances`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyCardBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCardLedger

> string ExportLoyaltyCardLedger(ctx, loyaltyProgramId, loyaltyCardId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()

Export card's ledger log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyCardLedger(context.Background(), loyaltyProgramId, loyaltyCardId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyCardLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyCardLedger`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyCardLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyCards

> string ExportLoyaltyCards(ctx, loyaltyProgramId).BatchId(batchId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()

Export loyalty cards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	batchId := "batchId_example" // string | Filter results by loyalty card batch ID. (optional)
	createdBefore := time.Now() // time.Time | Only return loyalty cards created before this timestamp.  **Note:** - This must be an RFC3339 timestamp string.  (optional)
	createdAfter := time.Now() // time.Time | Only return loyalty cards created after this timestamp.  **Note:** - This must be an RFC3339 timestamp string.  (optional)
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyCards(context.Background(), loyaltyProgramId).BatchId(batchId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).DateFormat(dateFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyCards`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchId** | **string** | Filter results by loyalty card batch ID. | 
 **createdBefore** | **time.Time** | Only return loyalty cards created before this timestamp.  **Note:** - This must be an RFC3339 timestamp string.  | 
 **createdAfter** | **time.Time** | Only return loyalty cards created after this timestamp.  **Note:** - This must be an RFC3339 timestamp string.  | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLoyaltyLedger

> string ExportLoyaltyLedger(ctx, loyaltyProgramId, integrationId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()

Export customer's transaction logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportLoyaltyLedger(context.Background(), loyaltyProgramId, integrationId).RangeStart(rangeStart).RangeEnd(rangeEnd).DateFormat(dateFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportLoyaltyLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportLoyaltyLedger`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportLoyaltyLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportLoyaltyLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 


 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPoolGiveaways

> string ExportPoolGiveaways(ctx, poolId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()

Export giveaway codes of a giveaway pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	poolId := int64(789) // int64 | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section.
	createdBefore := time.Now() // time.Time | Timestamp that filters the results to only contain giveaways created before this date. Must be an RFC3339 timestamp string. (optional)
	createdAfter := time.Now() // time.Time | Timestamp that filters the results to only contain giveaways created after this date. Must be an RFC3339 timestamp string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportPoolGiveaways(context.Background(), poolId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportPoolGiveaways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportPoolGiveaways`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportPoolGiveaways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPoolGiveawaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createdBefore** | **time.Time** | Timestamp that filters the results to only contain giveaways created before this date. Must be an RFC3339 timestamp string. | 
 **createdAfter** | **time.Time** | Timestamp that filters the results to only contain giveaways created after this date. Must be an RFC3339 timestamp string. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportReferrals

> string ExportReferrals(ctx, applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).BatchId(batchId).DateFormat(dateFormat).Execute()

Export referrals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := float32(8.14) // float32 | Filter results by campaign ID. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | - `expired`: Matches referrals in which the expiration date is set and in the past. - `validNow`: Matches referrals in which start date is null or in the past and expiration date is null or in the future. - `validFuture`: Matches referrals in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | - `true`, only referrals where `usageCounter < usageLimit` will be returned. - `false`, only referrals where `usageCounter >= usageLimit` will be returned.  (optional)
	batchId := "batchId_example" // string | Filter results by batches of referrals (optional)
	dateFormat := "dateFormat_example" // string | Determines the format of dates in the export document. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ExportReferrals(context.Background(), applicationId).CampaignId(campaignId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).BatchId(batchId).DateFormat(dateFormat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ExportReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportReferrals`: string
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ExportReferrals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignId** | **float32** | Filter results by campaign ID. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | - &#x60;expired&#x60;: Matches referrals in which the expiration date is set and in the past. - &#x60;validNow&#x60;: Matches referrals in which start date is null or in the past and expiration date is null or in the future. - &#x60;validFuture&#x60;: Matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | - &#x60;true&#x60;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned. - &#x60;false&#x60;, only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60; will be returned.  | 
 **batchId** | **string** | Filter results by batches of referrals | 
 **dateFormat** | **string** | Determines the format of dates in the export document. | 

### Return type

**string**

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateCouponRejections

> GenerateCouponRejections200Response GenerateCouponRejections(ctx).SessionIntegrationId(sessionIntegrationId).ApplicationId(applicationId).Language(language).CouponCode(couponCode).Execute()

Summarize coupon redemption failures in session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	sessionIntegrationId := "sessionIntegrationId_example" // string | The integration ID of the session to summarize.
	applicationId := float32(8.14) // float32 | Filter results by Application ID. (optional)
	language := "language_example" // string | The [ISO-639](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) code of the language in which the summary will be generated.  (optional)
	couponCode := "couponCode_example" // string | The coupon code for which to get the rejection reason. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GenerateCouponRejections(context.Background()).SessionIntegrationId(sessionIntegrationId).ApplicationId(applicationId).Language(language).CouponCode(couponCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GenerateCouponRejections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateCouponRejections`: GenerateCouponRejections200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GenerateCouponRejections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateCouponRejectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionIntegrationId** | **string** | The integration ID of the session to summarize. | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **language** | **string** | The [ISO-639](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes) code of the language in which the summary will be generated.  | 
 **couponCode** | **string** | The coupon code for which to get the rejection reason. | 

### Return type

[**GenerateCouponRejections200Response**](GenerateCouponRejections200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessLogsWithoutTotalCount

> GetAccessLogsWithoutTotalCount200Response GetAccessLogsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get access logs for Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	path := "path_example" // string | Only return results where the request path matches the given regular expression. (optional)
	method := "method_example" // string | Only return results where the request method matches the given regular expression. (optional)
	status := "status_example" // string | Filter results by HTTP status codes. (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAccessLogsWithoutTotalCount(context.Background(), applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).Path(path).Method(method).Status(status).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAccessLogsWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessLogsWithoutTotalCount`: GetAccessLogsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAccessLogsWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessLogsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **path** | **string** | Only return results where the request path matches the given regular expression. | 
 **method** | **string** | Only return results where the request method matches the given regular expression. | 
 **status** | **string** | Filter results by HTTP status codes. | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetAccessLogsWithoutTotalCount200Response**](GetAccessLogsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Get account details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	accountId := int64(789) // int64 | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the `accountId` property. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountAnalytics

> AccountAnalytics GetAccountAnalytics(ctx, accountId).Execute()

Get account analytics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	accountId := int64(789) // int64 | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the `accountId` property. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAccountAnalytics(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAccountAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountAnalytics`: AccountAnalytics
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAccountAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **int64** | The identifier of the account. Retrieve it via the [List users in account](https://docs.talon.one/management-api#operation/getUsers) endpoint in the &#x60;accountId&#x60; property.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAnalytics**](AccountAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountCollection

> Collection GetAccountCollection(ctx, collectionId).Execute()

Get account-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAccountCollection(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAccountCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAccountCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAchievement

> Achievement GetAchievement(ctx, applicationId, campaignId, achievementId).Execute()

Get achievement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	achievementId := int64(789) // int64 | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAchievement(context.Background(), applicationId, campaignId, achievementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAchievement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAchievement`: Achievement
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAchievement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCost

> AccountAdditionalCost GetAdditionalCost(ctx, additionalCostId).Execute()

Get additional cost



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	additionalCostId := int64(789) // int64 | The ID of the additional cost. You can find the ID the the Campaign Manager's URL when you display the details of the cost in **Account** > **Tools** > **Additional costs**. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAdditionalCost(context.Background(), additionalCostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAdditionalCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdditionalCost`: AccountAdditionalCost
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAdditionalCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int64** | The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdditionalCosts

> GetAdditionalCosts200Response GetAdditionalCosts(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List additional costs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAdditionalCosts(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAdditionalCosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdditionalCosts`: GetAdditionalCosts200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAdditionalCosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdditionalCostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetAdditionalCosts200Response**](GetAdditionalCosts200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplication

> Application GetApplication(ctx, applicationId).Execute()

Get Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: Application
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Application**](Application.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApiHealth

> ApplicationApiHealth GetApplicationApiHealth(ctx, applicationId).Execute()

Get Application health



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationApiHealth(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationApiHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationApiHealth`: ApplicationApiHealth
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationApiHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApiHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationApiHealth**](ApplicationApiHealth.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomer

> ApplicationCustomer GetApplicationCustomer(ctx, applicationId, customerId).Execute()

Get application's customer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	customerId := int64(789) // int64 | The value of the `id` property of a customer profile. Get it with the [List Application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationCustomer(context.Background(), applicationId, customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationCustomer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCustomer`: ApplicationCustomer
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int64** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationCustomer**](ApplicationCustomer.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomerFriends

> GetApplicationCustomerFriends200Response GetApplicationCustomerFriends(ctx, applicationId, integrationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List friends referred by customer profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	integrationId := "integrationId_example" // string | The Integration ID of the Advocate's Profile.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationCustomerFriends(context.Background(), applicationId, integrationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationCustomerFriends``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCustomerFriends`: GetApplicationCustomerFriends200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationCustomerFriends`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string** | The Integration ID of the Advocate&#39;s Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomerFriendsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**GetApplicationCustomerFriends200Response**](GetApplicationCustomerFriends200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomers

> GetApplicationCustomers200Response GetApplicationCustomers(ctx, applicationId).IntegrationId(integrationId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application's customers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	integrationId := "integrationId_example" // string | Filter results performing an exact matching against the profile integration identifier. (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationCustomers(context.Background(), applicationId).IntegrationId(integrationId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCustomers`: GetApplicationCustomers200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationId** | **string** | Filter results performing an exact matching against the profile integration identifier. | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**GetApplicationCustomers200Response**](GetApplicationCustomers200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationCustomersByAttributes

> GetApplicationCustomersByAttributes200Response GetApplicationCustomersByAttributes(ctx, applicationId).CustomerProfileSearchQuery(customerProfileSearchQuery).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()

List application customers matching the given attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	customerProfileSearchQuery := *openapiclient.NewCustomerProfileSearchQuery() // CustomerProfileSearchQuery | body
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationCustomersByAttributes(context.Background(), applicationId).CustomerProfileSearchQuery(customerProfileSearchQuery).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationCustomersByAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationCustomersByAttributes`: GetApplicationCustomersByAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationCustomersByAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerProfileSearchQuery** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) | body | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**GetApplicationCustomersByAttributes200Response**](GetApplicationCustomersByAttributes200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventTypes

> GetApplicationEventTypes200Response GetApplicationEventTypes(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications event types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationEventTypes(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationEventTypes`: GetApplicationEventTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationEventTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetApplicationEventTypes200Response**](GetApplicationEventTypes200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEventsWithoutTotalCount

> GetApplicationEventsWithoutTotalCount200Response GetApplicationEventsWithoutTotalCount(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).EffectType(effectType).Execute()

List Applications events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	type_ := "type__example" // string | Comma-separated list of types by which to filter events. Must be exact match(es). (optional)
	createdBefore := time.Now() // time.Time | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	session := "session_example" // string | Session integration ID filter for events. Must be exact match. (optional)
	profile := "profile_example" // string | Profile integration ID filter for events. Must be exact match. (optional)
	customerName := "customerName_example" // string | Customer name filter for events. Will match substrings case-insensitively. (optional)
	customerEmail := "customerEmail_example" // string | Customer e-mail address filter for events. Will match substrings case-insensitively. (optional)
	couponCode := "couponCode_example" // string | Coupon code (optional)
	referralCode := "referralCode_example" // string | Referral code (optional)
	ruleQuery := "ruleQuery_example" // string | Rule name filter for events (optional)
	campaignQuery := "campaignQuery_example" // string | Campaign name filter for events (optional)
	effectType := "effectType_example" // string | The type of effect that was triggered. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationEventsWithoutTotalCount(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Type_(type_).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Session(session).Profile(profile).CustomerName(customerName).CustomerEmail(customerEmail).CouponCode(couponCode).ReferralCode(referralCode).RuleQuery(ruleQuery).CampaignQuery(campaignQuery).EffectType(effectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationEventsWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationEventsWithoutTotalCount`: GetApplicationEventsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationEventsWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEventsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **type_** | **string** | Comma-separated list of types by which to filter events. Must be exact match(es). | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **session** | **string** | Session integration ID filter for events. Must be exact match. | 
 **profile** | **string** | Profile integration ID filter for events. Must be exact match. | 
 **customerName** | **string** | Customer name filter for events. Will match substrings case-insensitively. | 
 **customerEmail** | **string** | Customer e-mail address filter for events. Will match substrings case-insensitively. | 
 **couponCode** | **string** | Coupon code | 
 **referralCode** | **string** | Referral code | 
 **ruleQuery** | **string** | Rule name filter for events | 
 **campaignQuery** | **string** | Campaign name filter for events | 
 **effectType** | **string** | The type of effect that was triggered. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). | 

### Return type

[**GetApplicationEventsWithoutTotalCount200Response**](GetApplicationEventsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSession

> ApplicationSession GetApplicationSession(ctx, applicationId, sessionId).Execute()

Get Application session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	sessionId := int64(789) // int64 | The **internal** ID of the session. You can get the ID with the [List Application sessions](https://docs.talon.one/management-api#tag/Customer-data/operation/getApplicationSessions) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationSession(context.Background(), applicationId, sessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationSession`: ApplicationSession
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**sessionId** | **int64** | The **internal** ID of the session. You can get the ID with the [List Application sessions](https://docs.talon.one/management-api#tag/Customer-data/operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicationSession**](ApplicationSession.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationSessions

> GetApplicationSessions200Response GetApplicationSessions(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Profile(profile).State(state).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Coupon(coupon).Referral(referral).IntegrationId(integrationId).StoreIntegrationId(storeIntegrationId).Execute()

List Application sessions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	profile := "profile_example" // string | Profile integration ID filter for sessions. Must be exact match. (optional)
	state := "state_example" // string | Filter by sessions with this state. Must be exact match. (optional)
	createdBefore := time.Now() // time.Time | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	coupon := "coupon_example" // string | Filter by sessions with this coupon. Must be exact match. (optional)
	referral := "referral_example" // string | Filter by sessions with this referral. Must be exact match. (optional)
	integrationId := "integrationId_example" // string | Filter by sessions with this integration ID. Must be exact match. (optional)
	storeIntegrationId := "storeIntegrationId_example" // string | The integration ID of the store. You choose this ID when you create a store. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplicationSessions(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).Profile(profile).State(state).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Coupon(coupon).Referral(referral).IntegrationId(integrationId).StoreIntegrationId(storeIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplicationSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationSessions`: GetApplicationSessions200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplicationSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **profile** | **string** | Profile integration ID filter for sessions. Must be exact match. | 
 **state** | **string** | Filter by sessions with this state. Must be exact match. | 
 **createdBefore** | **time.Time** | Only return events created before this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Only return events created after this date. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **coupon** | **string** | Filter by sessions with this coupon. Must be exact match. | 
 **referral** | **string** | Filter by sessions with this referral. Must be exact match. | 
 **integrationId** | **string** | Filter by sessions with this integration ID. Must be exact match. | 
 **storeIntegrationId** | **string** | The integration ID of the store. You choose this ID when you create a store. | 

### Return type

[**GetApplicationSessions200Response**](GetApplicationSessions200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> GetApplications200Response GetApplications(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List Applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetApplications(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: GetApplications200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetApplications200Response**](GetApplications200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttribute

> Attribute GetAttribute(ctx, attributeId).Execute()

Get custom attribute



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	attributeId := int64(789) // int64 | The ID of the attribute. You can find the ID in the Campaign Manager's URL when you display the details of an attribute in **Account** > **Tools** > **Attributes**.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAttribute(context.Background(), attributeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttribute`: Attribute
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int64** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributes

> GetAttributes200Response GetAttributes(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Entity(entity).ApplicationIds(applicationIds).Type_(type_).Kind(kind).Search(search).Execute()

List custom attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	entity := "entity_example" // string | Returned attributes will be filtered by supplied entity. (optional)
	applicationIds := "applicationIds_example" // string | Returned attributes will be filtered by supplied application ids (optional)
	type_ := "type__example" // string | Returned attributes will be filtered by supplied type (optional)
	kind := "kind_example" // string | Returned attributes will be filtered by supplied kind (builtin or custom) (optional)
	search := "search_example" // string | Returned attributes will be filtered by searching case insensitive through Attribute name, description and type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAttributes(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).Entity(entity).ApplicationIds(applicationIds).Type_(type_).Kind(kind).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributes`: GetAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **entity** | **string** | Returned attributes will be filtered by supplied entity. | 
 **applicationIds** | **string** | Returned attributes will be filtered by supplied application ids | 
 **type_** | **string** | Returned attributes will be filtered by supplied type | 
 **kind** | **string** | Returned attributes will be filtered by supplied kind (builtin or custom) | 
 **search** | **string** | Returned attributes will be filtered by searching case insensitive through Attribute name, description and type | 

### Return type

[**GetAttributes200Response**](GetAttributes200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudienceMemberships

> GetAudienceMemberships200Response GetAudienceMemberships(ctx, audienceId).PageSize(pageSize).Skip(skip).Sort(sort).ProfileQuery(profileQuery).Execute()

List audience members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	audienceId := int64(789) // int64 | The ID of the audience.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	profileQuery := "profileQuery_example" // string | The filter to select a profile. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAudienceMemberships(context.Background(), audienceId).PageSize(pageSize).Skip(skip).Sort(sort).ProfileQuery(profileQuery).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAudienceMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudienceMemberships`: GetAudienceMemberships200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAudienceMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAudienceMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **profileQuery** | **string** | The filter to select a profile. | 

### Return type

[**GetAudienceMemberships200Response**](GetAudienceMemberships200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiences

> GetAudiences200Response GetAudiences(ctx).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()

List audiences



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAudiences(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudiences`: GetAudiences200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAudiences`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 

### Return type

[**GetAudiences200Response**](GetAudiences200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudiencesAnalytics

> GetAudiencesAnalytics200Response GetAudiencesAnalytics(ctx).AudienceIds(audienceIds).Sort(sort).Execute()

List audience analytics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	audienceIds := "audienceIds_example" // string | The IDs of one or more audiences, separated by commas, by which to filter results.
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetAudiencesAnalytics(context.Background()).AudienceIds(audienceIds).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetAudiencesAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAudiencesAnalytics`: GetAudiencesAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetAudiencesAnalytics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAudiencesAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **audienceIds** | **string** | The IDs of one or more audiences, separated by commas, by which to filter results. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetAudiencesAnalytics200Response**](GetAudiencesAnalytics200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> Campaign GetCampaign(ctx, applicationId, campaignId).Execute()

Get campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaign(context.Background(), applicationId, campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignAnalytics

> GetCampaignAnalytics200Response GetCampaignAnalytics(ctx, applicationId, campaignId).RangeStart(rangeStart).RangeEnd(rangeEnd).Granularity(granularity).Execute()

Get analytics of campaigns



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	granularity := "granularity_example" // string | The time interval between the results in the returned time-series. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaignAnalytics(context.Background(), applicationId, campaignId).RangeStart(rangeStart).RangeEnd(rangeEnd).Granularity(granularity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaignAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignAnalytics`: GetCampaignAnalytics200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaignAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **granularity** | **string** | The time interval between the results in the returned time-series. | 

### Return type

[**GetCampaignAnalytics200Response**](GetCampaignAnalytics200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignByAttributes

> GetCampaigns200Response GetCampaignByAttributes(ctx, applicationId).CampaignSearch(campaignSearch).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Execute()

List campaigns that match the given attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignSearch := *openapiclient.NewCampaignSearch(map[string]interface{}(123)) // CampaignSearch | body
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	campaignState := "campaignState_example" // string | Filter results by the state of the campaign.  - `enabled`: Campaigns that are scheduled, running (activated), or expired. - `running`: Campaigns that are running (activated). - `disabled`: Campaigns that are disabled. - `expired`: Campaigns that are expired. - `archived`: Campaigns that are archived.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaignByAttributes(context.Background(), applicationId).CampaignSearch(campaignSearch).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaignByAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignByAttributes`: GetCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaignByAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignSearch** | [**CampaignSearch**](CampaignSearch.md) | body | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**GetCampaigns200Response**](GetCampaigns200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroup

> CampaignGroup GetCampaignGroup(ctx, campaignGroupId).Execute()

Get campaign access group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	campaignGroupId := int64(789) // int64 | The ID of the campaign access group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaignGroup(context.Background(), campaignGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaignGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignGroup`: CampaignGroup
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaignGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignGroupId** | **int64** | The ID of the campaign access group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignGroup**](CampaignGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignGroups

> GetCampaignGroups200Response GetCampaignGroups(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List campaign access groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaignGroups(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaignGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignGroups`: GetCampaignGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaignGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetCampaignGroups200Response**](GetCampaignGroups200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignTemplates

> GetCampaignTemplates200Response GetCampaignTemplates(ctx).PageSize(pageSize).Skip(skip).Sort(sort).State(state).Name(name).Tags(tags).UserId(userId).Execute()

List campaign templates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	state := "state_example" // string | Filter results by the state of the campaign template. (optional)
	name := "name_example" // string | Filter results performing case-insensitive matching against the name of the campaign template. (optional)
	tags := "tags_example" // string | Filter results performing case-insensitive matching against the tags of the campaign template. When used in conjunction with the \"name\" query parameter, a logical OR will be performed to search both tags and name for the provided values.  (optional)
	userId := int64(789) // int64 | Filter results by user ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaignTemplates(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).State(state).Name(name).Tags(tags).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaignTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignTemplates`: GetCampaignTemplates200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaignTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **state** | **string** | Filter results by the state of the campaign template. | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign template. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign template. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values.  | 
 **userId** | **int64** | Filter results by user ID. | 

### Return type

[**GetCampaignTemplates200Response**](GetCampaignTemplates200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> GetCampaigns200Response GetCampaigns(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Name(name).Tags(tags).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartBefore(startBefore).StartAfter(startAfter).EndBefore(endBefore).EndAfter(endAfter).CampaignGroupId(campaignGroupId).TemplateId(templateId).StoreId(storeId).Execute()

List campaigns



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	campaignState := "campaignState_example" // string | Filter results by the state of the campaign.  - `enabled`: Campaigns that are scheduled, running (activated), or expired. - `running`: Campaigns that are running (activated). - `disabled`: Campaigns that are disabled. - `expired`: Campaigns that are expired. - `archived`: Campaigns that are archived.  (optional)
	name := "name_example" // string | Filter results performing case-insensitive matching against the name of the campaign. (optional)
	tags := "tags_example" // string | Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \"name\" query parameter, a logical OR will be performed to search both tags and name for the provided values  (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign start time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign start time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	endBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign end time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	endAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign end time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	campaignGroupId := int64(789) // int64 | Filter results to campaigns owned by the specified campaign access group ID. (optional)
	templateId := int64(789) // int64 | The ID of the campaign template this campaign was created from. (optional)
	storeId := int64(789) // int64 | Filter results to campaigns linked to the specified store ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCampaigns(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).CampaignState(campaignState).Name(name).Tags(tags).CreatedBefore(createdBefore).CreatedAfter(createdAfter).StartBefore(startBefore).StartAfter(startAfter).EndBefore(endBefore).EndAfter(endAfter).CampaignGroupId(campaignGroupId).TemplateId(templateId).StoreId(storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: GetCampaigns200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 
 **name** | **string** | Filter results performing case-insensitive matching against the name of the campaign. | 
 **tags** | **string** | Filter results performing case-insensitive matching against the tags of the campaign. When used in conjunction with the \&quot;name\&quot; query parameter, a logical OR will be performed to search both tags and name for the provided values  | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign start time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign start time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **endBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign end time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **endAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the campaign end time timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **campaignGroupId** | **int64** | Filter results to campaigns owned by the specified campaign access group ID. | 
 **templateId** | **int64** | The ID of the campaign template this campaign was created from. | 
 **storeId** | **int64** | Filter results to campaigns linked to the specified store ID. | 

### Return type

[**GetCampaigns200Response**](GetCampaigns200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChanges

> GetChanges200Response GetChanges(ctx).PageSize(pageSize).Skip(skip).Sort(sort).ApplicationId(applicationId).EntityPath(entityPath).UserId(userId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).WithTotalResultSize(withTotalResultSize).ManagementKeyId(managementKeyId).IncludeOld(includeOld).Execute()

Get audit logs for an account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	applicationId := float32(8.14) // float32 | Filter results by Application ID. (optional)
	entityPath := "entityPath_example" // string | Filter results on a case insensitive matching of the url path of the entity (optional)
	userId := int64(789) // int64 | Filter results by user ID. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	managementKeyId := int64(789) // int64 | Filter results that match the given management key ID. (optional)
	includeOld := true // bool | When this flag is set to false, the state without the change will not be returned. The default value is true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetChanges(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).ApplicationId(applicationId).EntityPath(entityPath).UserId(userId).CreatedBefore(createdBefore).CreatedAfter(createdAfter).WithTotalResultSize(withTotalResultSize).ManagementKeyId(managementKeyId).IncludeOld(includeOld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChanges`: GetChanges200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **entityPath** | **string** | Filter results on a case insensitive matching of the url path of the entity | 
 **userId** | **int64** | Filter results by user ID. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the change creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **managementKeyId** | **int64** | Filter results that match the given management key ID. | 
 **includeOld** | **bool** | When this flag is set to false, the state without the change will not be returned. The default value is true. | 

### Return type

[**GetChanges200Response**](GetChanges200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> Collection GetCollection(ctx, applicationId, campaignId, collectionId).Execute()

Get campaign-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCollection(context.Background(), applicationId, campaignId, collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionItems

> GetCollectionItems200Response GetCollectionItems(ctx, collectionId).PageSize(pageSize).Skip(skip).Execute()

Get collection items



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCollectionItems(context.Background(), collectionId).PageSize(pageSize).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCollectionItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionItems`: GetCollectionItems200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCollectionItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 

### Return type

[**GetCollectionItems200Response**](GetCollectionItems200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCouponsWithoutTotalCount

> GetCouponsWithoutTotalCount200Response GetCouponsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Redeemed(redeemed).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).ExpiresBefore(expiresBefore).ExpiresAfter(expiresAfter).StartsBefore(startsBefore).StartsAfter(startsAfter).ValuesOnly(valuesOnly).Execute()

List coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	value := "value_example" // string | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | Either \"expired\", \"validNow\", or \"validFuture\". The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | Either \"true\" or \"false\". If \"true\", only coupons where `usageCounter < usageLimit` will be returned, \"false\" will return only coupons where `usageCounter >= usageLimit`.  (optional)
	redeemed := "redeemed_example" // string | - `true`: only coupons where `usageCounter > 0` will be returned. - `false`: only coupons where `usageCounter = 0` will be returned. - This field cannot be used in conjunction with the `usable` query parameter.  (optional)
	referralId := int64(789) // int64 | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. (optional)
	recipientIntegrationId := "recipientIntegrationId_example" // string | Filter results by match with a profile ID specified in the coupon's RecipientIntegrationId field. (optional)
	batchId := "batchId_example" // string | Filter results by batches of coupons (optional)
	exactMatch := true // bool | Filter results to an exact case-insensitive matching against the coupon code. (optional) (default to false)
	expiresBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	expiresAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startsBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	startsAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valuesOnly := true // bool | Filter results to only return the coupon codes (`value` column) without the associated coupon data. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCouponsWithoutTotalCount(context.Background(), applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Redeemed(redeemed).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).ExpiresBefore(expiresBefore).ExpiresAfter(expiresAfter).StartsBefore(startsBefore).StartsAfter(startsAfter).ValuesOnly(valuesOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCouponsWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCouponsWithoutTotalCount`: GetCouponsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCouponsWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCouponsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **redeemed** | **string** | - &#x60;true&#x60;: only coupons where &#x60;usageCounter &gt; 0&#x60; will be returned. - &#x60;false&#x60;: only coupons where &#x60;usageCounter &#x3D; 0&#x60; will be returned. - This field cannot be used in conjunction with the &#x60;usable&#x60; query parameter.  | 
 **referralId** | **int64** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **expiresBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **expiresAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon expiration date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **startsAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon start date timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valuesOnly** | **bool** | Filter results to only return the coupon codes (&#x60;value&#x60; column) without the associated coupon data. | [default to false]

### Return type

[**GetCouponsWithoutTotalCount200Response**](GetCouponsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReport

> CustomerActivityReport GetCustomerActivityReport(ctx, applicationId, customerId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Execute()

Get customer's activity report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	customerId := int64(789) // int64 | The value of the `id` property of a customer profile. Get it with the [List Application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerActivityReport(context.Background(), applicationId, customerId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerActivityReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerActivityReport`: CustomerActivityReport
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerActivityReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int64** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 

### Return type

[**CustomerActivityReport**](CustomerActivityReport.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerActivityReportsWithoutTotalCount

> GetCustomerActivityReportsWithoutTotalCount200Response GetCustomerActivityReportsWithoutTotalCount(ctx, applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()

Get Activity Reports for Application Customers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	name := "name_example" // string | Only return reports matching the customer name. (optional)
	integrationId := "integrationId_example" // string | Filter results performing an exact matching against the profile integration identifier. (optional)
	campaignName := "campaignName_example" // string | Only return reports matching the campaign name. (optional)
	advocateName := "advocateName_example" // string | Only return reports matching the current customer referrer name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerActivityReportsWithoutTotalCount(context.Background(), applicationId).RangeStart(rangeStart).RangeEnd(rangeEnd).PageSize(pageSize).Skip(skip).Sort(sort).Name(name).IntegrationId(integrationId).CampaignName(campaignName).AdvocateName(advocateName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerActivityReportsWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerActivityReportsWithoutTotalCount`: GetCustomerActivityReportsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerActivityReportsWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerActivityReportsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **name** | **string** | Only return reports matching the customer name. | 
 **integrationId** | **string** | Filter results performing an exact matching against the profile integration identifier. | 
 **campaignName** | **string** | Only return reports matching the campaign name. | 
 **advocateName** | **string** | Only return reports matching the current customer referrer name. | 

### Return type

[**GetCustomerActivityReportsWithoutTotalCount200Response**](GetCustomerActivityReportsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAnalytics

> CustomerAnalytics GetCustomerAnalytics(ctx, applicationId, customerId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

Get customer's analytics report



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	customerId := int64(789) // int64 | The value of the `id` property of a customer profile. Get it with the [List Application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerAnalytics(context.Background(), applicationId, customerId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerAnalytics`: CustomerAnalytics
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**customerId** | **int64** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**CustomerAnalytics**](CustomerAnalytics.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfile

> CustomerProfile GetCustomerProfile(ctx, customerId).Execute()

Get customer profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	customerId := int64(789) // int64 | The value of the `id` property of a customer profile. Get it with the [List Application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerProfile(context.Background(), customerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerProfile`: CustomerProfile
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **int64** | The value of the &#x60;id&#x60; property of a customer profile. Get it with the [List Application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomerProfile**](CustomerProfile.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfileAchievementProgress

> GetCustomerProfileAchievementProgress200Response GetCustomerProfileAchievementProgress(ctx, applicationId, integrationId).PageSize(pageSize).Skip(skip).AchievementId(achievementId).Title(title).Execute()

List customer achievements



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	achievementId := int64(789) // int64 | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. (optional)
	title := "title_example" // string | Filter results by the `title` of an achievement. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerProfileAchievementProgress(context.Background(), applicationId, integrationId).PageSize(pageSize).Skip(skip).AchievementId(achievementId).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerProfileAchievementProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerProfileAchievementProgress`: GetCustomerProfileAchievementProgress200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerProfileAchievementProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfileAchievementProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **achievementId** | **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 
 **title** | **string** | Filter results by the &#x60;title&#x60; of an achievement. | 

### Return type

[**GetCustomerProfileAchievementProgress200Response**](GetCustomerProfileAchievementProgress200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerProfiles

> GetCustomerProfiles200Response GetCustomerProfiles(ctx).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()

List customer profiles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sandbox := true // bool | Indicates whether you are pointing to a sandbox or live customer. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomerProfiles(context.Background()).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomerProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerProfiles`: GetCustomerProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomerProfiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sandbox** | **bool** | Indicates whether you are pointing to a sandbox or live customer. | [default to false]

### Return type

[**GetCustomerProfiles200Response**](GetCustomerProfiles200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomersByAttributes

> GetCustomersByAttributes200Response GetCustomersByAttributes(ctx).CustomerProfileSearchQuery(customerProfileSearchQuery).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()

List customer profiles matching the given attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	customerProfileSearchQuery := *openapiclient.NewCustomerProfileSearchQuery() // CustomerProfileSearchQuery | body
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sandbox := true // bool | Indicates whether you are pointing to a sandbox or live customer. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetCustomersByAttributes(context.Background()).CustomerProfileSearchQuery(customerProfileSearchQuery).PageSize(pageSize).Skip(skip).Sandbox(sandbox).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetCustomersByAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomersByAttributes`: GetCustomersByAttributes200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetCustomersByAttributes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomersByAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerProfileSearchQuery** | [**CustomerProfileSearchQuery**](CustomerProfileSearchQuery.md) | body | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sandbox** | **bool** | Indicates whether you are pointing to a sandbox or live customer. | [default to false]

### Return type

[**GetCustomersByAttributes200Response**](GetCustomersByAttributes200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardStatistics

> GetDashboardStatistics200Response GetDashboardStatistics(ctx, loyaltyProgramId).RangeStart(rangeStart).RangeEnd(rangeEnd).SubledgerId(subledgerId).Execute()

Get statistics for loyalty dashboard



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	rangeStart := time.Now() // time.Time | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	rangeEnd := time.Now() // time.Time | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered. 
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetDashboardStatistics(context.Background(), loyaltyProgramId).RangeStart(rangeStart).RangeEnd(rangeEnd).SubledgerId(subledgerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetDashboardStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardStatistics`: GetDashboardStatistics200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetDashboardStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rangeStart** | **time.Time** | Only return results from after this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **rangeEnd** | **time.Time** | Only return results from before this timestamp.  **Note:** - This must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 

### Return type

[**GetDashboardStatistics200Response**](GetDashboardStatistics200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> GetEventTypes200Response GetEventTypes(ctx).Name(name).IncludeOldVersions(includeOldVersions).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List event types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	name := "name_example" // string | Filter results to event types with the given name. This parameter implies `includeOldVersions`. (optional)
	includeOldVersions := true // bool | Include all versions of every event type. (optional) (default to false)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetEventTypes(context.Background()).Name(name).IncludeOldVersions(includeOldVersions).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetEventTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventTypes`: GetEventTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Filter results to event types with the given name. This parameter implies &#x60;includeOldVersions&#x60;. | 
 **includeOldVersions** | **bool** | Include all versions of every event type. | [default to false]
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetEventTypes200Response**](GetEventTypes200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExports

> GetExports200Response GetExports(ctx).PageSize(pageSize).Skip(skip).ApplicationId(applicationId).CampaignId(campaignId).Entity(entity).Execute()

Get exports



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	applicationId := float32(8.14) // float32 | Filter results by Application ID. (optional)
	campaignId := int64(789) // int64 | Filter by the campaign ID on which the limit counters are used. (optional)
	entity := "entity_example" // string | The name of the entity type that was exported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetExports(context.Background()).PageSize(pageSize).Skip(skip).ApplicationId(applicationId).CampaignId(campaignId).Entity(entity).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetExports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetExports`: GetExports200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetExports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **campaignId** | **int64** | Filter by the campaign ID on which the limit counters are used. | 
 **entity** | **string** | The name of the entity type that was exported. | 

### Return type

[**GetExports200Response**](GetExports200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCard

> LoyaltyCard GetLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).Execute()

Get loyalty card



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyCard(context.Background(), loyaltyProgramId, loyaltyCardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCard`: LoyaltyCard
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactionLogs

> GetLoyaltyCardTransactionLogs200Response GetLoyaltyCardTransactionLogs(ctx, loyaltyProgramId, loyaltyCardId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).Execute()

List card's transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	startDate := time.Now() // time.Time | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	endDate := time.Now() // time.Time | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyCardTransactionLogs(context.Background(), loyaltyProgramId, loyaltyCardId).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyCardTransactionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCardTransactionLogs`: GetLoyaltyCardTransactionLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyCardTransactionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardTransactionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 

### Return type

[**GetLoyaltyCardTransactionLogs200Response**](GetLoyaltyCardTransactionLogs200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCards

> GetLoyaltyCards200Response GetLoyaltyCards(ctx, loyaltyProgramId).PageSize(pageSize).Skip(skip).Sort(sort).Identifier(identifier).ProfileId(profileId).BatchId(batchId).Execute()

List loyalty cards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	identifier := "identifier_example" // string | The card code by which to filter loyalty cards in the response. (optional)
	profileId := int64(789) // int64 | Filter results by customer profile ID. (optional)
	batchId := "batchId_example" // string | Filter results by loyalty card batch ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyCards(context.Background(), loyaltyProgramId).PageSize(pageSize).Skip(skip).Sort(sort).Identifier(identifier).ProfileId(profileId).BatchId(batchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCards`: GetLoyaltyCards200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **identifier** | **string** | The card code by which to filter loyalty cards in the response. | 
 **profileId** | **int64** | Filter results by customer profile ID. | 
 **batchId** | **string** | Filter results by loyalty card batch ID. | 

### Return type

[**GetLoyaltyCards200Response**](GetLoyaltyCards200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyLedgerBalances

> LoyaltyBalancesWithTiers GetLoyaltyLedgerBalances(ctx, loyaltyProgramId, integrationId).EndDate(endDate).SubledgerId(subledgerId).IncludeTiers(includeTiers).IncludeProjectedTier(includeProjectedTier).Execute()

Get customer's loyalty balances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	endDate := time.Now() // time.Time | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	subledgerId := "subledgerId_example" // string | The ID of the subledger used to filter the data. Leave this value empty (\"\") to query the main ledger. (optional)
	includeTiers := true // bool | Indicates whether tier information is included in the response.  When set to `true`, the response includes information about the current tier and the number of points required to move to next tier.  (optional) (default to false)
	includeProjectedTier := true // bool | Indicates whether the customer's projected tier information is included in the response.  When set to `true`, the response includes information about the customer's active points and the name of the projected tier.  **Note** We recommend filtering by `subledgerId` for better performance.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyLedgerBalances(context.Background(), loyaltyProgramId, integrationId).EndDate(endDate).SubledgerId(subledgerId).IncludeTiers(includeTiers).IncludeProjectedTier(includeProjectedTier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyLedgerBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyLedgerBalances`: LoyaltyBalancesWithTiers
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyLedgerBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyLedgerBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **string** | The ID of the subledger used to filter the data. Leave this value empty (\&quot;\&quot;) to query the main ledger. | 
 **includeTiers** | **bool** | Indicates whether tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the current tier and the number of points required to move to next tier.  | [default to false]
 **includeProjectedTier** | **bool** | Indicates whether the customer&#39;s projected tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the customer&#39;s active points and the name of the projected tier.  **Note** We recommend filtering by &#x60;subledgerId&#x60; for better performance.  | [default to false]

### Return type

[**LoyaltyBalancesWithTiers**](LoyaltyBalancesWithTiers.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPoints

> LoyaltyLedger GetLoyaltyPoints(ctx, loyaltyProgramId, integrationId).Execute()

Get customer's full loyalty ledger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyPoints(context.Background(), loyaltyProgramId, integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyPoints`: LoyaltyLedger
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LoyaltyLedger**](LoyaltyLedger.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> LoyaltyProgram GetLoyaltyProgram(ctx, loyaltyProgramId).Execute()

Get loyalty program



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyProgram(context.Background(), loyaltyProgramId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyProgram``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgram`: LoyaltyProgram
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyProgram**](LoyaltyProgram.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfileLedgerTransactions

> GetLoyaltyProgramProfileTransactions200Response GetLoyaltyProgramProfileLedgerTransactions(ctx, loyaltyProgramId, integrationId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()

List customer's loyalty transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)
	subledgerId := "subledgerId_example" // string | The ID of the subledger used to filter the data. Leave this value empty (\"\") to query the main ledger. (optional)
	loyaltyTransactionType := "loyaltyTransactionType_example" // string | Filter results by loyalty transaction type: - `manual`: Loyalty transaction that was done manually. - `session`: Loyalty transaction that resulted from a customer session. - `import`: Loyalty transaction that was imported from a CSV file.  (optional)
	startDate := time.Now() // time.Time | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	endDate := time.Now() // time.Time | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	awaitsActivation := true // bool | If `true`: Filters results to include only point transactions that have action-based activation and have not expired.  If `false`: Returns a `400` response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyProgramProfileLedgerTransactions(context.Background(), loyaltyProgramId, integrationId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyProgramProfileLedgerTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramProfileLedgerTransactions`: GetLoyaltyProgramProfileTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyProgramProfileLedgerTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramProfileLedgerTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **subledgerId** | **string** | The ID of the subledger used to filter the data. Leave this value empty (\&quot;\&quot;) to query the main ledger. | 
 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **awaitsActivation** | **bool** | If &#x60;true&#x60;: Filters results to include only point transactions that have action-based activation and have not expired.  If &#x60;false&#x60;: Returns a &#x60;400&#x60; response.  | 

### Return type

[**GetLoyaltyProgramProfileTransactions200Response**](GetLoyaltyProgramProfileTransactions200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramTransactions

> GetLoyaltyProgramTransactions200Response GetLoyaltyProgramTransactions(ctx, loyaltyProgramId).LoyaltyTransactionType(loyaltyTransactionType).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()

List loyalty program transactions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyTransactionType := "loyaltyTransactionType_example" // string | Filter results by loyalty transaction type: - `manual`: Loyalty transaction that was done manually. - `session`: Loyalty transaction that resulted from a customer session. - `import`: Loyalty transaction that was imported from a CSV file.  (optional)
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)
	startDate := time.Now() // time.Time | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	endDate := time.Now() // time.Time | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	awaitsActivation := true // bool | If `true`: Filters results to include only point transactions that have action-based activation and have not expired.  If `false`: Returns a `400` response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyProgramTransactions(context.Background(), loyaltyProgramId).LoyaltyTransactionType(loyaltyTransactionType).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyProgramTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramTransactions`: GetLoyaltyProgramTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyProgramTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **awaitsActivation** | **bool** | If &#x60;true&#x60;: Filters results to include only point transactions that have action-based activation and have not expired.  If &#x60;false&#x60;: Returns a &#x60;400&#x60; response.  | 

### Return type

[**GetLoyaltyProgramTransactions200Response**](GetLoyaltyProgramTransactions200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyPrograms

> GetLoyaltyPrograms200Response GetLoyaltyPrograms(ctx).Execute()

List loyalty programs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyPrograms(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyPrograms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyPrograms`: GetLoyaltyPrograms200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyPrograms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramsRequest struct via the builder pattern


### Return type

[**GetLoyaltyPrograms200Response**](GetLoyaltyPrograms200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyStatistics

> LoyaltyDashboardData GetLoyaltyStatistics(ctx, loyaltyProgramId).Execute()

Get loyalty program statistics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetLoyaltyStatistics(context.Background(), loyaltyProgramId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetLoyaltyStatistics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyStatistics`: LoyaltyDashboardData
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetLoyaltyStatistics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyStatisticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoyaltyDashboardData**](LoyaltyDashboardData.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageLogs

> MessageLogEntries GetMessageLogs(ctx).EntityType(entityType).MessageID(messageID).ChangeType(changeType).NotificationIDs(notificationIDs).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Cursor(cursor).Period(period).IsSuccessful(isSuccessful).ApplicationId(applicationId).CampaignId(campaignId).LoyaltyProgramId(loyaltyProgramId).ResponseCode(responseCode).WebhookIDs(webhookIDs).Execute()

List message log entries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	entityType := "entityType_example" // string | The entity type the log is related to. 
	messageID := "messageID_example" // string | Filter results by message ID. (optional)
	changeType := "changeType_example" // string | Filter results by change type. (optional)
	notificationIDs := "notificationIDs_example" // string | Filter results by notification ID (include up to 30 values, separated by a comma). (optional)
	createdBefore := time.Now() // time.Time | Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. Use UTC time. (optional)
	createdAfter := time.Now() // time.Time | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. Use UTC time. (optional)
	cursor := string(BYTE_ARRAY_DATA_HERE) // string | A specific unique value in the database. If this value is not given, the server fetches results starting with the first record.  (optional)
	period := "period_example" // string | Filter results by time period. Choose between the available relative time frames.  (optional)
	isSuccessful := true // bool | Indicates whether to return log entries with either successful or unsuccessful HTTP response codes. When set to`true`, only log entries with `2xx` response codes are returned. When set to `false`, only log entries with `4xx` and `5xx` response codes are returned.  (optional)
	applicationId := float32(8.14) // float32 | Filter results by Application ID. (optional)
	campaignId := float32(8.14) // float32 | Filter results by campaign ID. (optional)
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. (optional)
	responseCode := int64(789) // int64 | Filter results by response status code. (optional)
	webhookIDs := "webhookIDs_example" // string | Filter results by webhook ID (include up to 30 values, separated by a comma). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetMessageLogs(context.Background()).EntityType(entityType).MessageID(messageID).ChangeType(changeType).NotificationIDs(notificationIDs).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Cursor(cursor).Period(period).IsSuccessful(isSuccessful).ApplicationId(applicationId).CampaignId(campaignId).LoyaltyProgramId(loyaltyProgramId).ResponseCode(responseCode).WebhookIDs(webhookIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetMessageLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageLogs`: MessageLogEntries
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetMessageLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entityType** | **string** | The entity type the log is related to.  | 
 **messageID** | **string** | Filter results by message ID. | 
 **changeType** | **string** | Filter results by change type. | 
 **notificationIDs** | **string** | Filter results by notification ID (include up to 30 values, separated by a comma). | 
 **createdBefore** | **time.Time** | Filter results where request and response times to return entries before parameter value, expected to be an RFC3339 timestamp string. Use UTC time. | 
 **createdAfter** | **time.Time** | Filter results where request and response times to return entries after parameter value, expected to be an RFC3339 timestamp string. Use UTC time. | 
 **cursor** | **string** | A specific unique value in the database. If this value is not given, the server fetches results starting with the first record.  | 
 **period** | **string** | Filter results by time period. Choose between the available relative time frames.  | 
 **isSuccessful** | **bool** | Indicates whether to return log entries with either successful or unsuccessful HTTP response codes. When set to&#x60;true&#x60;, only log entries with &#x60;2xx&#x60; response codes are returned. When set to &#x60;false&#x60;, only log entries with &#x60;4xx&#x60; and &#x60;5xx&#x60; response codes are returned.  | 
 **applicationId** | **float32** | Filter results by Application ID. | 
 **campaignId** | **float32** | Filter results by campaign ID. | 
 **loyaltyProgramId** | **int64** | Identifier of the loyalty program. | 
 **responseCode** | **int64** | Filter results by response status code. | 
 **webhookIDs** | **string** | Filter results by webhook ID (include up to 30 values, separated by a comma). | 

### Return type

[**MessageLogEntries**](MessageLogEntries.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReferralsWithoutTotalCount

> GetReferralsWithoutTotalCount200Response GetReferralsWithoutTotalCount(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()

List referrals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	code := "code_example" // string | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | Either \"expired\", \"validNow\", or \"validFuture\". The first option matches referrals in which the expiration date is set and in the past. The second matches referrals in which start date is null or in the past and expiration date is null or in the future, the third matches referrals in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | Either \"true\" or \"false\". If \"true\", only referrals where `usageCounter < usageLimit` will be returned, \"false\" will return only referrals where `usageCounter >= usageLimit`.  (optional)
	advocate := "advocate_example" // string | Filter results by match with a profile ID specified in the referral's AdvocateProfileIntegrationId field. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetReferralsWithoutTotalCount(context.Background(), applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Code(code).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).Advocate(advocate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetReferralsWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReferralsWithoutTotalCount`: GetReferralsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetReferralsWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReferralsWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **code** | **string** | Filter results performing case-insensitive matching against the referral code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the referral creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches referrals in which the expiration date is set and in the past. The second matches referrals in which start date is null or in the past and expiration date is null or in the future, the third matches referrals in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only referrals where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only referrals where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **advocate** | **string** | Filter results by match with a profile ID specified in the referral&#39;s AdvocateProfileIntegrationId field. | 

### Return type

[**GetReferralsWithoutTotalCount200Response**](GetReferralsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleV2

> RoleV2 GetRoleV2(ctx, roleId).Execute()

Get role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	roleId := int64(789) // int64 | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetRoleV2(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRoleV2`: RoleV2
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetRoleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64** | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRuleset

> Ruleset GetRuleset(ctx, applicationId, campaignId, rulesetId).Execute()

Get ruleset



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	rulesetId := int64(789) // int64 | The ID of the ruleset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetRuleset(context.Background(), applicationId, campaignId, rulesetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetRuleset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRuleset`: Ruleset
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetRuleset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**rulesetId** | **int64** | The ID of the ruleset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Ruleset**](Ruleset.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRulesets

> GetRulesets200Response GetRulesets(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List campaign rulesets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetRulesets(context.Background(), applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetRulesets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRulesets`: GetRulesets200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetRulesets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRulesetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetRulesets200Response**](GetRulesets200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStore

> Store GetStore(ctx, applicationId, storeId).Execute()

Get store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	storeId := "storeId_example" // string | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetStore(context.Background(), applicationId, storeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStore`: Store
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> User GetUser(ctx, userId).Execute()

Get user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> GetUsers200Response GetUsers(ctx).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List users in account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetUsers(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsers`: GetUsers200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 

### Return type

[**GetUsers200Response**](GetUsers200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhook

> Webhook GetWebhook(ctx, webhookId).Execute()

Get webhook



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	webhookId := int64(789) // int64 | The ID of the webhook. You can find the ID in the Campaign Manager's URL when you display the details of the webhook in **Account** > **Webhooks**. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **int64** | The ID of the webhook. You can find the ID in the Campaign Manager&#39;s URL when you display the details of the webhook in **Account** &gt; **Webhooks**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> GetWebhooks200Response GetWebhooks(ctx).ApplicationIds(applicationIds).Sort(sort).PageSize(pageSize).Skip(skip).CreationType(creationType).Visibility(visibility).OutgoingIntegrationsTypeId(outgoingIntegrationsTypeId).Title(title).Execute()

List webhooks



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationIds := "applicationIds_example" // string | Checks if the given catalog or its attributes are referenced in the specified Application ID.  **Note**: If no Application ID is provided, we check for all connected Applications.  (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	creationType := "creationType_example" // string | Filter results by creation type. (optional)
	visibility := "visibility_example" // string | Filter results by visibility. (optional)
	outgoingIntegrationsTypeId := int64(789) // int64 | Filter results by outgoing integration type ID. (optional)
	title := "title_example" // string | Filter results performing case-insensitive matching against the webhook title. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.GetWebhooks(context.Background()).ApplicationIds(applicationIds).Sort(sort).PageSize(pageSize).Skip(skip).CreationType(creationType).Visibility(visibility).OutgoingIntegrationsTypeId(outgoingIntegrationsTypeId).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.GetWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWebhooks`: GetWebhooks200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationIds** | **string** | Checks if the given catalog or its attributes are referenced in the specified Application ID.  **Note**: If no Application ID is provided, we check for all connected Applications.  | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **creationType** | **string** | Filter results by creation type. | 
 **visibility** | **string** | Filter results by visibility. | 
 **outgoingIntegrationsTypeId** | **int64** | Filter results by outgoing integration type ID. | 
 **title** | **string** | Filter results performing case-insensitive matching against the webhook title. | 

### Return type

[**GetWebhooks200Response**](GetWebhooks200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAccountCollection

> Import ImportAccountCollection(ctx, collectionId).UpFile(upFile).Execute()

Import data into existing account-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportAccountCollection(context.Background(), collectionId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportAccountCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAccountCollection`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportAccountCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAllowedList

> Import ImportAllowedList(ctx, attributeId).UpFile(upFile).Execute()

Import allowed values for attribute



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	attributeId := int64(789) // int64 | The ID of the attribute. You can find the ID in the Campaign Manager's URL when you display the details of an attribute in **Account** > **Tools** > **Attributes**.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportAllowedList(context.Background(), attributeId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportAllowedList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAllowedList`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportAllowedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int64** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAllowedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAudiencesMemberships

> Import ImportAudiencesMemberships(ctx, audienceId).UpFile(upFile).Execute()

Import audience members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	audienceId := int64(789) // int64 | The ID of the audience.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportAudiencesMemberships(context.Background(), audienceId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportAudiencesMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAudiencesMemberships`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportAudiencesMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAudiencesMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCampaignStoreBudget

> Import ImportCampaignStoreBudget(ctx, applicationId, campaignId).Action(action).Period(period).UpFile(upFile).Execute()

Import campaign store budgets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	action := "action_example" // string | The action that this budget is limiting. (optional)
	period := "period_example" // string | The period to which the limit applies.  **Note**: For budgets with no period, set this to `overall`.  (optional)
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportCampaignStoreBudget(context.Background(), applicationId, campaignId).Action(action).Period(period).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportCampaignStoreBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCampaignStoreBudget`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportCampaignStoreBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCampaignStoreBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **string** | The action that this budget is limiting. | 
 **period** | **string** | The period to which the limit applies.  **Note**: For budgets with no period, set this to &#x60;overall&#x60;.  | 
 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCampaignStores

> Import ImportCampaignStores(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportCampaignStores(context.Background(), applicationId, campaignId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportCampaignStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCampaignStores`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportCampaignStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCampaignStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCollection

> Import ImportCollection(ctx, applicationId, campaignId, collectionId).UpFile(upFile).Execute()

Import data into existing campaign-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportCollection(context.Background(), applicationId, campaignId, collectionId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCollection`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportCoupons

> Import ImportCoupons(ctx, applicationId, campaignId).SkipDuplicates(skipDuplicates).UpFile(upFile).Execute()

Import coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	skipDuplicates := true // bool | An indicator of whether to skip duplicate coupon values instead of causing an error. Duplicate values are ignored when `skipDuplicates=true`.  (optional)
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportCoupons(context.Background(), applicationId, campaignId).SkipDuplicates(skipDuplicates).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportCoupons`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipDuplicates** | **bool** | An indicator of whether to skip duplicate coupon values instead of causing an error. Duplicate values are ignored when &#x60;skipDuplicates&#x3D;true&#x60;.  | 
 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCards

> Import ImportLoyaltyCards(ctx, loyaltyProgramId).UpFile(upFile).Execute()

Import loyalty cards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportLoyaltyCards(context.Background(), loyaltyProgramId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportLoyaltyCards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportLoyaltyCards`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportLoyaltyCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyCustomersTiers

> Import ImportLoyaltyCustomersTiers(ctx, loyaltyProgramId).UpFile(upFile).Execute()

Import customers into loyalty tiers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportLoyaltyCustomersTiers(context.Background(), loyaltyProgramId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportLoyaltyCustomersTiers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportLoyaltyCustomersTiers`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportLoyaltyCustomersTiers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyCustomersTiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportLoyaltyPoints

> Import ImportLoyaltyPoints(ctx, loyaltyProgramId).NotificationsEnabled(notificationsEnabled).UpFile(upFile).Execute()

Import loyalty points



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	notificationsEnabled := true // bool | Indicates whether the points import triggers notifications about its effects. For example, a notification is sent if the import upgrades a customer's tier or offsets their negative points balance.  This parameter is optional and defaults to `true`.  (optional)
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportLoyaltyPoints(context.Background(), loyaltyProgramId).NotificationsEnabled(notificationsEnabled).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportLoyaltyPoints`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportLoyaltyPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationsEnabled** | **bool** | Indicates whether the points import triggers notifications about its effects. For example, a notification is sent if the import upgrades a customer&#39;s tier or offsets their negative points balance.  This parameter is optional and defaults to &#x60;true&#x60;.  | 
 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportPoolGiveaways

> Import ImportPoolGiveaways(ctx, poolId).UpFile(upFile).Execute()

Import giveaway codes into a giveaway pool



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	poolId := int64(789) // int64 | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportPoolGiveaways(context.Background(), poolId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportPoolGiveaways``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportPoolGiveaways`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportPoolGiveaways`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poolId** | **int64** | The ID of the pool. You can find it in the Campaign Manager, in the **Giveaways** section. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportPoolGiveawaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportReferrals

> Import ImportReferrals(ctx, applicationId, campaignId).UpFile(upFile).Execute()

Import referrals



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	upFile := "upFile_example" // string | The file containing the data that is being imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ImportReferrals(context.Background(), applicationId, campaignId).UpFile(upFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ImportReferrals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportReferrals`: Import
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ImportReferrals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportReferralsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **upFile** | **string** | The file containing the data that is being imported. | 

### Return type

[**Import**](Import.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InviteUserExternal

> InviteUserExternal(ctx).NewExternalInvitation(newExternalInvitation).Execute()

Invite user from identity provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newExternalInvitation := *openapiclient.NewNewExternalInvitation("john.doe@example.com") // NewExternalInvitation | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.InviteUserExternal(context.Background()).NewExternalInvitation(newExternalInvitation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.InviteUserExternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInviteUserExternalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newExternalInvitation** | [**NewExternalInvitation**](NewExternalInvitation.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountCollections

> ListAccountCollections200Response ListAccountCollections(ctx).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	name := "name_example" // string | Filter by collection name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListAccountCollections(context.Background()).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListAccountCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccountCollections`: ListAccountCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListAccountCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**ListAccountCollections200Response**](ListAccountCollections200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAchievements

> ListAchievements200Response ListAchievements(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Title(title).Execute()

List achievements



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	title := "title_example" // string | Filter by the display name for the achievement in the campaign manager.  **Note**: If no `title` is provided, all the achievements from the campaign are returned.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListAchievements(context.Background(), applicationId, campaignId).PageSize(pageSize).Skip(skip).Title(title).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListAchievements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAchievements`: ListAchievements200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListAchievements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAchievementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **title** | **string** | Filter by the display name for the achievement in the campaign manager.  **Note**: If no &#x60;title&#x60; is provided, all the achievements from the campaign are returned.  | 

### Return type

[**ListAchievements200Response**](ListAchievements200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRolesV2

> ListAllRolesV2200Response ListAllRolesV2(ctx).Execute()

List roles



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListAllRolesV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListAllRolesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRolesV2`: ListAllRolesV2200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListAllRolesV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllRolesV2Request struct via the builder pattern


### Return type

[**ListAllRolesV2200Response**](ListAllRolesV2200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaignStoreBudgetLimits

> ListCampaignStoreBudgetLimits200Response ListCampaignStoreBudgetLimits(ctx, applicationId, campaignId).Action(action).Period(period).Execute()

List campaign store budget limits



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	action := "action_example" // string | The action that this budget is limiting. (optional)
	period := "period_example" // string | The period to which the limit applies.  **Note**: For budgets with no period, set this to `overall`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListCampaignStoreBudgetLimits(context.Background(), applicationId, campaignId).Action(action).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListCampaignStoreBudgetLimits``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCampaignStoreBudgetLimits`: ListCampaignStoreBudgetLimits200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListCampaignStoreBudgetLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignStoreBudgetLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **string** | The action that this budget is limiting. | 
 **period** | **string** | The period to which the limit applies.  **Note**: For budgets with no period, set this to &#x60;overall&#x60;.  | 

### Return type

[**ListCampaignStoreBudgetLimits200Response**](ListCampaignStoreBudgetLimits200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCatalogItems

> ListCatalogItems200Response ListCatalogItems(ctx, catalogId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Sku(sku).ProductNames(productNames).Execute()

List items in a catalog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	catalogId := int64(789) // int64 | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** > **Tools** > **Cart item catalogs**.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	sku := []string{"Inner_example"} // []string | Filter results by one or more SKUs. Must be exact match. (optional)
	productNames := []string{"Inner_example"} // []string | Filter results by one or more product names. Must be exact match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListCatalogItems(context.Background(), catalogId).PageSize(pageSize).Skip(skip).WithTotalResultSize(withTotalResultSize).Sku(sku).ProductNames(productNames).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListCatalogItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCatalogItems`: ListCatalogItems200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListCatalogItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int64** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **sku** | **[]string** | Filter results by one or more SKUs. Must be exact match. | 
 **productNames** | **[]string** | Filter results by one or more product names. Must be exact match. | 

### Return type

[**ListCatalogItems200Response**](ListCatalogItems200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> ListAccountCollections200Response ListCollections(ctx, applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	name := "name_example" // string | Filter by collection name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListCollections(context.Background(), applicationId, campaignId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollections`: ListAccountCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListCollections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**ListAccountCollections200Response**](ListAccountCollections200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsInApplication

> ListAccountCollections200Response ListCollectionsInApplication(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()

List collections in Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	name := "name_example" // string | Filter by collection name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListCollectionsInApplication(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListCollectionsInApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionsInApplication`: ListAccountCollections200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListCollectionsInApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsInApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **name** | **string** | Filter by collection name. | 

### Return type

[**ListAccountCollections200Response**](ListAccountCollections200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStores

> ListStores200Response ListStores(ctx, applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).CampaignId(campaignId).Name(name).IntegrationId(integrationId).Query(query).Execute()

List stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	withTotalResultSize := true // bool | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When `true`: `hasMore` is true when there is a next page. `totalResultSize` is always zero. - When `false`: `hasMore` is always false. `totalResultSize` contains the total number of results for this query.  (optional)
	campaignId := float32(8.14) // float32 | Filter results by campaign ID. (optional)
	name := "name_example" // string | The name of the store. (optional)
	integrationId := "integrationId_example" // string | The integration ID of the store. (optional)
	query := "query_example" // string | Filter results by `name` or `integrationId`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ListStores(context.Background(), applicationId).PageSize(pageSize).Skip(skip).Sort(sort).WithTotalResultSize(withTotalResultSize).CampaignId(campaignId).Name(name).IntegrationId(integrationId).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ListStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStores`: ListStores200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ListStores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **withTotalResultSize** | **bool** | When this flag is set, the result includes the total size of the result, across all pages. This might decrease performance on large data sets.  - When &#x60;true&#x60;: &#x60;hasMore&#x60; is true when there is a next page. &#x60;totalResultSize&#x60; is always zero. - When &#x60;false&#x60;: &#x60;hasMore&#x60; is always false. &#x60;totalResultSize&#x60; contains the total number of results for this query.  | 
 **campaignId** | **float32** | Filter results by campaign ID. | 
 **name** | **string** | The name of the store. | 
 **integrationId** | **string** | The integration ID of the store. | 
 **query** | **string** | Filter results by &#x60;name&#x60; or &#x60;integrationId&#x60;. | 

### Return type

[**ListStores200Response**](ListStores200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OktaEventHandlerChallenge

> OktaEventHandlerChallenge(ctx).Execute()

Validate Okta API ownership



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.OktaEventHandlerChallenge(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.OktaEventHandlerChallenge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOktaEventHandlerChallengeRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLoyaltyPoints

> RemoveLoyaltyPoints(ctx, loyaltyProgramId, integrationId).DeductLoyaltyPoints(deductLoyaltyPoints).Execute()

Deduct points from customer profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := "loyaltyProgramId_example" // string | The identifier for the loyalty program.
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	deductLoyaltyPoints := *openapiclient.NewDeductLoyaltyPoints(float32(300)) // DeductLoyaltyPoints | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.RemoveLoyaltyPoints(context.Background(), loyaltyProgramId, integrationId).DeductLoyaltyPoints(deductLoyaltyPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.RemoveLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **string** | The identifier for the loyalty program. | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deductLoyaltyPoints** | [**DeductLoyaltyPoints**](DeductLoyaltyPoints.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetPassword

> NewPassword ResetPassword(ctx).NewPassword(newPassword).Execute()

Reset password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	newPassword := *openapiclient.NewNewPassword("Admin&12943!7", "Z2VgacVNkexLKBUIzsRDDZSGxnIkC53y") // NewPassword | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ResetPassword(context.Background()).NewPassword(newPassword).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ResetPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetPassword`: NewPassword
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ResetPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newPassword** | [**NewPassword**](NewPassword.md) | body | 

### Return type

[**NewPassword**](NewPassword.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimCreateGroup

> ScimGroup ScimCreateGroup(ctx).ScimBaseGroup(scimBaseGroup).Execute()

Create SCIM group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	scimBaseGroup := *openapiclient.NewScimBaseGroup() // ScimBaseGroup | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimCreateGroup(context.Background()).ScimBaseGroup(scimBaseGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimCreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimCreateGroup`: ScimGroup
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimCreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScimCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimBaseGroup** | [**ScimBaseGroup**](ScimBaseGroup.md) | body | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimCreateUser

> ScimUser ScimCreateUser(ctx).ScimNewUser(scimNewUser).Execute()

Create SCIM user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	scimNewUser := *openapiclient.NewScimNewUser("john.doe@example.com") // ScimNewUser | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimCreateUser(context.Background()).ScimNewUser(scimNewUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimCreateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimCreateUser`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimCreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScimCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scimNewUser** | [**ScimNewUser**](ScimNewUser.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimDeleteGroup

> ScimDeleteGroup(ctx, groupId).Execute()

Delete SCIM group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	groupId := int64(789) // int64 | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.ScimDeleteGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimDeleteGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimDeleteGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimDeleteUser

> ScimDeleteUser(ctx, userId).Execute()

Delete SCIM user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.ScimDeleteUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimDeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetGroup

> ScimGroup ScimGetGroup(ctx, groupId).Execute()

Get SCIM group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	groupId := int64(789) // int64 | The ID of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetGroup(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetGroup`: ScimGroup
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetGroups

> ScimGroupsListResponse ScimGetGroups(ctx).Execute()

List SCIM groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetGroups`: ScimGroupsListResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetGroupsRequest struct via the builder pattern


### Return type

[**ScimGroupsListResponse**](ScimGroupsListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetResourceTypes

> ScimResourceTypesListResponse ScimGetResourceTypes(ctx).Execute()

List supported SCIM resource types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetResourceTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetResourceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetResourceTypes`: ScimResourceTypesListResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetResourceTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetResourceTypesRequest struct via the builder pattern


### Return type

[**ScimResourceTypesListResponse**](ScimResourceTypesListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetSchemas

> ScimSchemasListResponse ScimGetSchemas(ctx).Execute()

List supported SCIM schemas



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetSchemas(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetSchemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetSchemas`: ScimSchemasListResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetSchemasRequest struct via the builder pattern


### Return type

[**ScimSchemasListResponse**](ScimSchemasListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetServiceProviderConfig

> ScimServiceProviderConfigResponse ScimGetServiceProviderConfig(ctx).Execute()

Get SCIM service provider configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetServiceProviderConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetServiceProviderConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetServiceProviderConfig`: ScimServiceProviderConfigResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetServiceProviderConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetServiceProviderConfigRequest struct via the builder pattern


### Return type

[**ScimServiceProviderConfigResponse**](ScimServiceProviderConfigResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUser

> ScimUser ScimGetUser(ctx, userId).Execute()

Get SCIM user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetUser(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetUser`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimGetUsers

> ScimUsersListResponse ScimGetUsers(ctx).Execute()

List SCIM users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimGetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimGetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimGetUsers`: ScimUsersListResponse
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimGetUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiScimGetUsersRequest struct via the builder pattern


### Return type

[**ScimUsersListResponse**](ScimUsersListResponse.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimPatchGroup

> ScimGroup ScimPatchGroup(ctx, groupId).ScimPatchRequest(scimPatchRequest).Execute()

Update SCIM group attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	groupId := int64(789) // int64 | The ID of the group.
	scimPatchRequest := *openapiclient.NewScimPatchRequest([]openapiclient.ScimPatchOperation{*openapiclient.NewScimPatchOperation("Op_example")}) // ScimPatchRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimPatchGroup(context.Background(), groupId).ScimPatchRequest(scimPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimPatchGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimPatchGroup`: ScimGroup
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimPatchGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimPatchGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimPatchRequest** | [**ScimPatchRequest**](ScimPatchRequest.md) | body | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimPatchUser

> ScimUser ScimPatchUser(ctx, userId).ScimPatchRequest(scimPatchRequest).Execute()

Update SCIM user attributes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.
	scimPatchRequest := *openapiclient.NewScimPatchRequest([]openapiclient.ScimPatchOperation{*openapiclient.NewScimPatchOperation("Op_example")}) // ScimPatchRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimPatchUser(context.Background(), userId).ScimPatchRequest(scimPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimPatchUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimPatchUser`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimPatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimPatchUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimPatchRequest** | [**ScimPatchRequest**](ScimPatchRequest.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimReplaceGroupAttributes

> ScimGroup ScimReplaceGroupAttributes(ctx, groupId).ScimBaseGroup(scimBaseGroup).Execute()

Update SCIM group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	groupId := int64(789) // int64 | The ID of the group.
	scimBaseGroup := *openapiclient.NewScimBaseGroup() // ScimBaseGroup | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimReplaceGroupAttributes(context.Background(), groupId).ScimBaseGroup(scimBaseGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimReplaceGroupAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimReplaceGroupAttributes`: ScimGroup
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimReplaceGroupAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int64** | The ID of the group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimReplaceGroupAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimBaseGroup** | [**ScimBaseGroup**](ScimBaseGroup.md) | body | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScimReplaceUserAttributes

> ScimUser ScimReplaceUserAttributes(ctx, userId).ScimNewUser(scimNewUser).Execute()

Update SCIM user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.
	scimNewUser := *openapiclient.NewScimNewUser("john.doe@example.com") // ScimNewUser | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.ScimReplaceUserAttributes(context.Background(), userId).ScimNewUser(scimNewUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.ScimReplaceUserAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScimReplaceUserAttributes`: ScimUser
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.ScimReplaceUserAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiScimReplaceUserAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scimNewUser** | [**ScimNewUser**](ScimNewUser.md) | body | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedApplicationWideWithoutTotalCount

> GetCouponsWithoutTotalCount200Response SearchCouponsAdvancedApplicationWideWithoutTotalCount(ctx, applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()

List coupons that match the given attributes (without total count)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	body := map[string]interface{}{ ... } // map[string]interface{} | body
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	value := "value_example" // string | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | Either \"expired\", \"validNow\", or \"validFuture\". The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | Either \"true\" or \"false\". If \"true\", only coupons where `usageCounter < usageLimit` will be returned, \"false\" will return only coupons where `usageCounter >= usageLimit`.  (optional)
	referralId := int64(789) // int64 | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. (optional)
	recipientIntegrationId := "recipientIntegrationId_example" // string | Filter results by match with a profile ID specified in the coupon's RecipientIntegrationId field. (optional)
	batchId := "batchId_example" // string | Filter results by batches of coupons (optional)
	exactMatch := true // bool | Filter results to an exact case-insensitive matching against the coupon code. (optional) (default to false)
	campaignState := "campaignState_example" // string | Filter results by the state of the campaign.  - `enabled`: Campaigns that are scheduled, running (activated), or expired. - `running`: Campaigns that are running (activated). - `disabled`: Campaigns that are disabled. - `expired`: Campaigns that are expired. - `archived`: Campaigns that are archived.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.SearchCouponsAdvancedApplicationWideWithoutTotalCount(context.Background(), applicationId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).BatchId(batchId).ExactMatch(exactMatch).CampaignState(campaignState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.SearchCouponsAdvancedApplicationWideWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCouponsAdvancedApplicationWideWithoutTotalCount`: GetCouponsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.SearchCouponsAdvancedApplicationWideWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedApplicationWideWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | body | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int64** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **batchId** | **string** | Filter results by batches of coupons | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **campaignState** | **string** | Filter results by the state of the campaign.  - &#x60;enabled&#x60;: Campaigns that are scheduled, running (activated), or expired. - &#x60;running&#x60;: Campaigns that are running (activated). - &#x60;disabled&#x60;: Campaigns that are disabled. - &#x60;expired&#x60;: Campaigns that are expired. - &#x60;archived&#x60;: Campaigns that are archived.  | 

### Return type

[**GetCouponsWithoutTotalCount200Response**](GetCouponsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchCouponsAdvancedWithoutTotalCount

> GetCouponsWithoutTotalCount200Response SearchCouponsAdvancedWithoutTotalCount(ctx, applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()

List coupons that match the given attributes in campaign (without total count)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	body := map[string]interface{}{ ... } // map[string]interface{} | body
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with `-`.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  (optional)
	value := "value_example" // string | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. (optional)
	createdBefore := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	createdAfter := time.Now() // time.Time | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. (optional)
	valid := "valid_example" // string | Either \"expired\", \"validNow\", or \"validFuture\". The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  (optional)
	usable := "usable_example" // string | Either \"true\" or \"false\". If \"true\", only coupons where `usageCounter < usageLimit` will be returned, \"false\" will return only coupons where `usageCounter >= usageLimit`.  (optional)
	referralId := int64(789) // int64 | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. (optional)
	recipientIntegrationId := "recipientIntegrationId_example" // string | Filter results by match with a profile ID specified in the coupon's RecipientIntegrationId field. (optional)
	exactMatch := true // bool | Filter results to an exact case-insensitive matching against the coupon code. (optional) (default to false)
	batchId := "batchId_example" // string | Filter results by batches of coupons (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.SearchCouponsAdvancedWithoutTotalCount(context.Background(), applicationId, campaignId).Body(body).PageSize(pageSize).Skip(skip).Sort(sort).Value(value).CreatedBefore(createdBefore).CreatedAfter(createdAfter).Valid(valid).Usable(usable).ReferralId(referralId).RecipientIntegrationId(recipientIntegrationId).ExactMatch(exactMatch).BatchId(batchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.SearchCouponsAdvancedWithoutTotalCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchCouponsAdvancedWithoutTotalCount`: GetCouponsWithoutTotalCount200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.SearchCouponsAdvancedWithoutTotalCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchCouponsAdvancedWithoutTotalCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** | body | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. By default, results are sorted in ascending order. To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You may not be able to use all fields for sorting. This is due to performance limitations.  | 
 **value** | **string** | Filter results performing case-insensitive matching against the coupon code. Both the code and the query are folded to remove all non-alpha-numeric characters. | 
 **createdBefore** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **createdAfter** | **time.Time** | Filter results comparing the parameter value, expected to be an RFC3339 timestamp string, to the coupon creation timestamp. You can use any time zone setting. Talon.One will convert to UTC internally. | 
 **valid** | **string** | Either \&quot;expired\&quot;, \&quot;validNow\&quot;, or \&quot;validFuture\&quot;. The first option matches coupons in which the expiration date is set and in the past. The second matches coupons in which start date is null or in the past and expiration date is null or in the future, the third matches coupons in which start date is set and in the future.  | 
 **usable** | **string** | Either \&quot;true\&quot; or \&quot;false\&quot;. If \&quot;true\&quot;, only coupons where &#x60;usageCounter &lt; usageLimit&#x60; will be returned, \&quot;false\&quot; will return only coupons where &#x60;usageCounter &gt;&#x3D; usageLimit&#x60;.  | 
 **referralId** | **int64** | Filter the results by matching them with the ID of a referral. This filter shows the coupons created by redeeming a referral code. | 
 **recipientIntegrationId** | **string** | Filter results by match with a profile ID specified in the coupon&#39;s RecipientIntegrationId field. | 
 **exactMatch** | **bool** | Filter results to an exact case-insensitive matching against the coupon code. | [default to false]
 **batchId** | **string** | Filter results by batches of coupons | 

### Return type

[**GetCouponsWithoutTotalCount200Response**](GetCouponsWithoutTotalCount200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SummarizeCampaignStoreBudget

> SummarizeCampaignStoreBudget200Response SummarizeCampaignStoreBudget(ctx, applicationId, campaignId).Execute()

Get summary of campaign store budgets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.SummarizeCampaignStoreBudget(context.Background(), applicationId, campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.SummarizeCampaignStoreBudget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SummarizeCampaignStoreBudget`: SummarizeCampaignStoreBudget200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.SummarizeCampaignStoreBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSummarizeCampaignStoreBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SummarizeCampaignStoreBudget200Response**](SummarizeCampaignStoreBudget200Response.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferLoyaltyCard

> TransferLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).TransferLoyaltyCard(transferLoyaltyCard).Execute()

Transfer card data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	transferLoyaltyCard := *openapiclient.NewTransferLoyaltyCard("summer-loyalty-card-0543") // TransferLoyaltyCard | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.TransferLoyaltyCard(context.Background(), loyaltyProgramId, loyaltyCardId).TransferLoyaltyCard(transferLoyaltyCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.TransferLoyaltyCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transferLoyaltyCard** | [**TransferLoyaltyCard**](TransferLoyaltyCard.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountCollection

> Collection UpdateAccountCollection(ctx, collectionId).UpdateCollection(updateCollection).Execute()

Update account-level collection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint.
	updateCollection := *openapiclient.NewUpdateCollection() // UpdateCollection | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateAccountCollection(context.Background(), collectionId).UpdateCollection(updateCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateAccountCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccountCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateAccountCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in account](#operation/listAccountCollections) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCollection** | [**UpdateCollection**](UpdateCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAchievement

> Achievement UpdateAchievement(ctx, applicationId, campaignId, achievementId).UpdateAchievement(updateAchievement).Execute()

Update achievement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	achievementId := int64(789) // int64 | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint.
	updateAchievement := *openapiclient.NewUpdateAchievement() // UpdateAchievement | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateAchievement(context.Background(), applicationId, campaignId, achievementId).UpdateAchievement(updateAchievement).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateAchievement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAchievement`: Achievement
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateAchievement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**achievementId** | **int64** | The ID of the achievement. You can get this ID with the [List achievement](https://docs.talon.one/management-api#tag/Achievements/operation/listAchievements) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAchievementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateAchievement** | [**UpdateAchievement**](UpdateAchievement.md) | body | 

### Return type

[**Achievement**](Achievement.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdditionalCost

> AccountAdditionalCost UpdateAdditionalCost(ctx, additionalCostId).NewAdditionalCost(newAdditionalCost).Execute()

Update additional cost



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	additionalCostId := int64(789) // int64 | The ID of the additional cost. You can find the ID the the Campaign Manager's URL when you display the details of the cost in **Account** > **Tools** > **Additional costs**. 
	newAdditionalCost := *openapiclient.NewNewAdditionalCost("shippingFee", "Shipping fee", "A shipping fee") // NewAdditionalCost | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateAdditionalCost(context.Background(), additionalCostId).NewAdditionalCost(newAdditionalCost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateAdditionalCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAdditionalCost`: AccountAdditionalCost
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateAdditionalCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**additionalCostId** | **int64** | The ID of the additional cost. You can find the ID the the Campaign Manager&#39;s URL when you display the details of the cost in **Account** &gt; **Tools** &gt; **Additional costs**.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdditionalCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAdditionalCost** | [**NewAdditionalCost**](NewAdditionalCost.md) | body | 

### Return type

[**AccountAdditionalCost**](AccountAdditionalCost.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAttribute

> Attribute UpdateAttribute(ctx, attributeId).NewAttribute(newAttribute).Execute()

Update custom attribute



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	attributeId := int64(789) // int64 | The ID of the attribute. You can find the ID in the Campaign Manager's URL when you display the details of an attribute in **Account** > **Tools** > **Attributes**.
	newAttribute := *openapiclient.NewNewAttribute("Event", "pageViewed", "Page view event", "string", "Event triggered when a customer displays a page.", []string{"Suggestions_example"}, true) // NewAttribute | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateAttribute(context.Background(), attributeId).NewAttribute(newAttribute).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateAttribute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAttribute`: Attribute
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateAttribute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int64** | The ID of the attribute. You can find the ID in the Campaign Manager&#39;s URL when you display the details of an attribute in **Account** &gt; **Tools** &gt; **Attributes**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAttributeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newAttribute** | [**NewAttribute**](NewAttribute.md) | body | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> Campaign UpdateCampaign(ctx, applicationId, campaignId).UpdateCampaign(updateCampaign).Execute()

Update campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	updateCampaign := *openapiclient.NewUpdateCampaign("Summer promotions", []string{"Tags_example"}, []string{"Features_example"}, []openapiclient.LimitConfig{*openapiclient.NewLimitConfig("createCoupon", float32(1000), []string{"Entities_example"})}) // UpdateCampaign | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateCampaign(context.Background(), applicationId, campaignId).UpdateCampaign(updateCampaign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaign`: Campaign
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCampaign** | [**UpdateCampaign**](UpdateCampaign.md) | body | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, applicationId, campaignId, collectionId).UpdateCampaignCollection(updateCampaignCollection).Execute()

Update campaign-level collection's description



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	collectionId := int64(789) // int64 | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint.
	updateCampaignCollection := *openapiclient.NewUpdateCampaignCollection() // UpdateCampaignCollection | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateCollection(context.Background(), applicationId, campaignId, collectionId).UpdateCampaignCollection(updateCampaignCollection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCollection`: Collection
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**collectionId** | **int64** | The ID of the collection. You can get it with the [List collections in Application](#operation/listCollectionsInApplication) endpoint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCampaignCollection** | [**UpdateCampaignCollection**](UpdateCampaignCollection.md) | body | 

### Return type

[**Collection**](Collection.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCoupon

> Coupon UpdateCoupon(ctx, applicationId, campaignId, couponId).UpdateCoupon(updateCoupon).Execute()

Update coupon



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	couponId := "couponId_example" // string | The internal ID of the coupon code. You can find this value in the `id` property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response. 
	updateCoupon := *openapiclient.NewUpdateCoupon() // UpdateCoupon | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateCoupon(context.Background(), applicationId, campaignId, couponId).UpdateCoupon(updateCoupon).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateCoupon``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCoupon`: Coupon
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**couponId** | **string** | The internal ID of the coupon code. You can find this value in the &#x60;id&#x60; property from the [List coupons](https://docs.talon.one/management-api#tag/Coupons/operation/getCouponsWithoutTotalCount) endpoint response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateCoupon** | [**UpdateCoupon**](UpdateCoupon.md) | body | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCouponBatch

> UpdateCouponBatch(ctx, applicationId, campaignId).UpdateCouponBatch(updateCouponBatch).Execute()

Update coupons



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	updateCouponBatch := *openapiclient.NewUpdateCouponBatch() // UpdateCouponBatch | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.UpdateCouponBatch(context.Background(), applicationId, campaignId).UpdateCouponBatch(updateCouponBatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateCouponBatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCouponBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateCouponBatch** | [**UpdateCouponBatch**](UpdateCouponBatch.md) | body | 

### Return type

 (empty response body)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoyaltyCard

> LoyaltyCard UpdateLoyaltyCard(ctx, loyaltyProgramId, loyaltyCardId).UpdateLoyaltyCard(updateLoyaltyCard).Execute()

Update loyalty card status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	loyaltyProgramId := int64(789) // int64 | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	updateLoyaltyCard := *openapiclient.NewUpdateLoyaltyCard("active") // UpdateLoyaltyCard | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateLoyaltyCard(context.Background(), loyaltyProgramId, loyaltyCardId).UpdateLoyaltyCard(updateLoyaltyCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateLoyaltyCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoyaltyCard`: LoyaltyCard
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateLoyaltyCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateLoyaltyCard** | [**UpdateLoyaltyCard**](UpdateLoyaltyCard.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReferral

> Referral UpdateReferral(ctx, applicationId, campaignId, referralId).UpdateReferral(updateReferral).Execute()

Update referral



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	campaignId := int64(789) // int64 | The ID of the campaign. It is displayed in your Talon.One deployment URL.
	referralId := "referralId_example" // string | The ID of the referral code.
	updateReferral := *openapiclient.NewUpdateReferral() // UpdateReferral | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateReferral(context.Background(), applicationId, campaignId, referralId).UpdateReferral(updateReferral).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReferral`: Referral
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateReferral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**campaignId** | **int64** | The ID of the campaign. It is displayed in your Talon.One deployment URL. | 
**referralId** | **string** | The ID of the referral code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateReferral** | [**UpdateReferral**](UpdateReferral.md) | body | 

### Return type

[**Referral**](Referral.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoleV2

> RoleV2 UpdateRoleV2(ctx, roleId).RoleV2Base(roleV2Base).Execute()

Update role



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	roleId := int64(789) // int64 | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint. 
	roleV2Base := *openapiclient.NewRoleV2Base() // RoleV2Base | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateRoleV2(context.Background(), roleId).RoleV2Base(roleV2Base).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateRoleV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoleV2`: RoleV2
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateRoleV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **int64** | The ID of role.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoleV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleV2Base** | [**RoleV2Base**](RoleV2Base.md) | body | 

### Return type

[**RoleV2**](RoleV2.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStore

> Store UpdateStore(ctx, applicationId, storeId).NewStore(newStore).Execute()

Update store



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	applicationId := int64(789) // int64 | The ID of the Application. It is displayed in your Talon.One deployment URL.
	storeId := "storeId_example" // string | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint. 
	newStore := *openapiclient.NewNewStore("South US store", "This is the description of the store in south US.", "STORE-001") // NewStore | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateStore(context.Background(), applicationId, storeId).NewStore(newStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStore`: Store
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int64** | The ID of the Application. It is displayed in your Talon.One deployment URL. | 
**storeId** | **string** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newStore** | [**NewStore**](NewStore.md) | body | 

### Return type

[**Store**](Store.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> User UpdateUser(ctx, userId).UpdateUser(updateUser).Execute()

Update user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/talon-one/talon-one-go-sdk"
)

func main() {
	userId := int64(789) // int64 | The ID of the user.
	updateUser := *openapiclient.NewUpdateUser() // UpdateUser | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagementAPI.UpdateUser(context.Background(), userId).UpdateUser(updateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: User
	fmt.Fprintf(os.Stdout, "Response from `ManagementAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **int64** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUser** | [**UpdateUser**](UpdateUser.md) | body | 

### Return type

[**User**](User.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

