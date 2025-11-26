# Talon.One Go SDK

> [!WARNING]
This project is in beta. Use is "AS IS" and at your own risk.
The stable version of this SDK is available [here](https://github.com/talon-one/talon_go).

This SDK supports all of the operations of Talon.One's Integration API and Management API.

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import "./talon"
```

## Getting started

### Integration API

The following code shows an example of using the Integration API:

```golang
package main

import (
    "context"
    "encoding/json"
    "fmt"

    talon "github.com/talon-one/talon_go/v8"
)

func main() {
    configuration := talon.NewConfiguration()
    // Set API base path
    configuration.Servers = talon.ServerConfigurations{
        {
            // Notice that there is no trailing '/'
            URL:         "https://yourbaseurl.talon.one",
            Description: "Talon.One's API base URL",
        },
    }
    // If you wish to inject a custom implementation of HTTPClient
    // configuration.HTTPClient = &customHTTPClient

    integrationClient := talon.NewAPIClient(configuration)

    // Create integration authentication context using api key
    integrationAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
        "Authorization": {
            Prefix: "ApiKey-v1",
            Key:    "fd1fd219b1e953a6b2700e8034de5bfc877462ae106127311ddd710978654312",
        },
    })

    // Instantiating a NewCustomerSessionV2 struct
    newCustomerSession := talon.NewCustomerSessionV2{
        // You can use either struct literals
        ProfileId:   talon.PtrString("DEADBEEF"),
        CouponCodes: &[]string{"Cool-Stuff!"},
    }

    // Or alternatively, using the relevant setter in a later stage in the code
    newCustomerSession.SetCartItems([]talon.CartItem{
        {
            Name:     talon.PtrString("Pad Thai - Veggie"),
            Sku:      "pad-332",
            Quantity: 1,
            Price:    talon.PtrFloat32(5.5),
            Category: talon.PtrString("Noodles"),
        },
        {
            Name:     talon.PtrString("Chang"),
            Sku:      "chang-br-42",
            Quantity: 1,
            Price:    talon.PtrFloat32(2.3),
            Category: talon.PtrString("Beverages"),
        },
    })

    // Instantiating a new IntegrationRequest
    integrationRequest := talon.IntegrationRequest{
        CustomerSession: newCustomerSession,
    }

    // Optional list of requested information to be present on the response.
    // See docs/IntegrationRequest.md for full list of supported values
    // integrationRequest.SetResponseContent([]string{
    //  "customerSession",
    //  "customerProfile",
    //  "loyalty",
    // })

    // Create/update a customer session using `UpdateCustomerSessionV2` function
    integrationState, _, err := integrationClient.IntegrationApi.
        UpdateCustomerSessionV2(integrationAuthContext, "deetdoot_2").
        Body(integrationRequest).
        Execute()

    if err != nil {
        fmt.Printf("ERROR while calling UpdateCustomerSessionV2: %s\n", err)
        return
    }
    fmt.Printf("%#v\n", integrationState)

    // Parsing the returned effects list, please consult https://developers.talon.one/Integration-API/handling-effects-v2 for the full list of effects and their corresponding properties
    for _, effect := range integrationState.GetEffects() {
        effectType := effect.GetEffectType()
        switch {
        case "setDiscount" == effectType:
            // Initiating right props instance according to the effect type
            effectProps := talon.SetDiscountEffectProps{}
            if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
                fmt.Printf("ERROR while decoding 'setDiscount' props: %s\n", err)
                continue
            }

            // Access the specific effect's properties
            fmt.Printf("Set a discount '%s' of %2.3f\n", effectProps.GetName(), effectProps.GetValue())
        case "acceptCoupon" == effectType:
            // Initiating right props instance according to the effect type
            effectProps := talon.AcceptCouponEffectProps{}
            if err := decodeHelper(effect.GetProps(), &effectProps); err != nil {
                fmt.Printf("ERROR while decoding props: %s\n", err)
                continue
            }

            // Work with AcceptCouponEffectProps' properties
            // ...
        default:
            fmt.Printf("Encounter unknown effect type: %s\n", effectType)
        }
    }
}

// quick decoding of props-map into our library structures using JSON marshaling,
// or alternatively using a library like https://github.com/mitchellh/mapstructure
func decodeHelper(propsMap map[string]interface{}, v interface{}) error {
    propsJSON, err := json.Marshal(propsMap)
    if err != nil {
        return err
    }
    return json.Unmarshal(propsJSON, v)
}
```

### Management API

The following code shows an example of using the Management API:

```golang
package main

import (
    "context"
    "fmt"

    talon "github.com/talon-one/talon_go/v8"
)

func main() {
    configuration := talon.NewConfiguration()
    // Set API base path
    configuration.Servers = talon.ServerConfigurations{
        {
            // Notice that there is no trailing '/'
            URL:         "https://yourbaseurl.talon.one",
            Description: "Talon.One's API base URL",
        },
    }
    // If you wish to inject a custom implementation of HTTPClient
    // configuration.HTTPClient = &customHTTPClient

    managementClient := talon.NewAPIClient(configuration)

    // Create integration authentication context using the logged-in session
    managerAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
        "Authorization": talon.APIKey{
            Prefix: "ManagementKey-v1",
            Key:    "2f0dce055da01ae595005d7d79154bae7448d319d5fc7c5b2951fadd6ba1ea07",
        },
    })

    // Calling `GetApplication` function with the desired id (7)
    application, response, err := managementClient.ManagementApi.
        GetApplication(managerAuthContext, 7).
        Execute()

    if err != nil {
        fmt.Printf("ERROR while calling GetApplication: %s\n", err)
        return
    }

    fmt.Printf("%#v\n\n", application)
    fmt.Printf("%#v\n\n", response)
}
```

## Documentation for API endpoints

All URLs are relative to `https://yourbaseurl.talon.one`.

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IntegrationAPI* | [**BestPriorPrice**](docs/IntegrationAPI.md#bestpriorprice) | **Post** /v1/best_prior_price | Fetch best prior price
*IntegrationAPI* | [**CreateAudienceV2**](docs/IntegrationAPI.md#createaudiencev2) | **Post** /v2/audiences | Create audience
*IntegrationAPI* | [**CreateCouponReservation**](docs/IntegrationAPI.md#createcouponreservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
*IntegrationAPI* | [**CreateReferral**](docs/IntegrationAPI.md#createreferral) | **Post** /v1/referrals | Create referral code for an advocate
*IntegrationAPI* | [**CreateReferralsForMultipleAdvocates**](docs/IntegrationAPI.md#createreferralsformultipleadvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates
*IntegrationAPI* | [**DeleteAudienceMembershipsV2**](docs/IntegrationAPI.md#deleteaudiencemembershipsv2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
*IntegrationAPI* | [**DeleteAudienceV2**](docs/IntegrationAPI.md#deleteaudiencev2) | **Delete** /v2/audiences/{audienceId} | Delete audience
*IntegrationAPI* | [**DeleteCouponReservation**](docs/IntegrationAPI.md#deletecouponreservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
*IntegrationAPI* | [**DeleteCustomerData**](docs/IntegrationAPI.md#deletecustomerdata) | **Delete** /v1/customer_data/{integrationId} | Delete customer&#39;s personal data
*IntegrationAPI* | [**GenerateLoyaltyCard**](docs/IntegrationAPI.md#generateloyaltycard) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards | Generate loyalty card
*IntegrationAPI* | [**GetCustomerAchievementHistory**](docs/IntegrationAPI.md#getcustomerachievementhistory) | **Get** /v1/customer_profiles/{integrationId}/achievements/{achievementId} | List customer&#39;s achievement history
*IntegrationAPI* | [**GetCustomerAchievements**](docs/IntegrationAPI.md#getcustomerachievements) | **Get** /v1/customer_profiles/{integrationId}/achievements | List customer&#39;s available achievements
*IntegrationAPI* | [**GetCustomerInventory**](docs/IntegrationAPI.md#getcustomerinventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
*IntegrationAPI* | [**GetCustomerSession**](docs/IntegrationAPI.md#getcustomersession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
*IntegrationAPI* | [**GetLoyaltyBalances**](docs/IntegrationAPI.md#getloyaltybalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/balances | Get customer&#39;s loyalty balances
*IntegrationAPI* | [**GetLoyaltyCardBalances**](docs/IntegrationAPI.md#getloyaltycardbalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/balances | Get card&#39;s point balances
*IntegrationAPI* | [**GetLoyaltyCardPoints**](docs/IntegrationAPI.md#getloyaltycardpoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/points | List card&#39;s unused loyalty points
*IntegrationAPI* | [**GetLoyaltyCardTransactions**](docs/IntegrationAPI.md#getloyaltycardtransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transactions | List card&#39;s transactions
*IntegrationAPI* | [**GetLoyaltyProgramProfilePoints**](docs/IntegrationAPI.md#getloyaltyprogramprofilepoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/points | List customer&#39;s unused loyalty points
*IntegrationAPI* | [**GetLoyaltyProgramProfileTransactions**](docs/IntegrationAPI.md#getloyaltyprogramprofiletransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/transactions | List customer&#39;s loyalty transactions
*IntegrationAPI* | [**GetReservedCustomers**](docs/IntegrationAPI.md#getreservedcustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List customers that have this coupon reserved
*IntegrationAPI* | [**LinkLoyaltyCardToProfile**](docs/IntegrationAPI.md#linkloyaltycardtoprofile) | **Post** /v2/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/link_profile | Link customer profile to card
*IntegrationAPI* | [**ReopenCustomerSession**](docs/IntegrationAPI.md#reopencustomersession) | **Put** /v2/customer_sessions/{customerSessionId}/reopen | Reopen customer session
*IntegrationAPI* | [**ReturnCartItems**](docs/IntegrationAPI.md#returncartitems) | **Post** /v2/customer_sessions/{customerSessionId}/returns | Return cart items
*IntegrationAPI* | [**SyncCatalog**](docs/IntegrationAPI.md#synccatalog) | **Put** /v1/catalogs/{catalogId}/sync | Sync cart item catalog
*IntegrationAPI* | [**TrackEventV2**](docs/IntegrationAPI.md#trackeventv2) | **Post** /v2/events | Track event
*IntegrationAPI* | [**UpdateAudienceCustomersAttributes**](docs/IntegrationAPI.md#updateaudiencecustomersattributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
*IntegrationAPI* | [**UpdateAudienceV2**](docs/IntegrationAPI.md#updateaudiencev2) | **Put** /v2/audiences/{audienceId} | Update audience name
*IntegrationAPI* | [**UpdateCustomerProfileAudiences**](docs/IntegrationAPI.md#updatecustomerprofileaudiences) | **Post** /v2/customer_audiences | Update multiple customer profiles&#39; audiences
*IntegrationAPI* | [**UpdateCustomerProfileV2**](docs/IntegrationAPI.md#updatecustomerprofilev2) | **Put** /v2/customer_profiles/{integrationId} | Update customer profile
*IntegrationAPI* | [**UpdateCustomerProfilesV2**](docs/IntegrationAPI.md#updatecustomerprofilesv2) | **Put** /v2/customer_profiles | Update multiple customer profiles
*IntegrationAPI* | [**UpdateCustomerSessionV2**](docs/IntegrationAPI.md#updatecustomersessionv2) | **Put** /v2/customer_sessions/{customerSessionId} | Update customer session
*ManagementAPI* | [**ActivateUserByEmail**](docs/ManagementAPI.md#activateuserbyemail) | **Post** /v1/users/activate | Enable user by email address
*ManagementAPI* | [**AddLoyaltyCardPoints**](docs/ManagementAPI.md#addloyaltycardpoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/add_points | Add points to card
*ManagementAPI* | [**AddLoyaltyPoints**](docs/ManagementAPI.md#addloyaltypoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/add_points | Add points to customer profile
*ManagementAPI* | [**CopyCampaignToApplications**](docs/ManagementAPI.md#copycampaigntoapplications) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/copy | Copy the campaign into the specified Application
*ManagementAPI* | [**CreateAccountCollection**](docs/ManagementAPI.md#createaccountcollection) | **Post** /v1/collections | Create account-level collection
*ManagementAPI* | [**CreateAchievement**](docs/ManagementAPI.md#createachievement) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | Create achievement
*ManagementAPI* | [**CreateAdditionalCost**](docs/ManagementAPI.md#createadditionalcost) | **Post** /v1/additional_costs | Create additional cost
*ManagementAPI* | [**CreateAttribute**](docs/ManagementAPI.md#createattribute) | **Post** /v1/attributes | Create custom attribute
*ManagementAPI* | [**CreateBatchLoyaltyCards**](docs/ManagementAPI.md#createbatchloyaltycards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards/batch | Create loyalty cards
*ManagementAPI* | [**CreateCampaignFromTemplate**](docs/ManagementAPI.md#createcampaignfromtemplate) | **Post** /v1/applications/{applicationId}/create_campaign_from_template | Create campaign from campaign template
*ManagementAPI* | [**CreateCampaignStoreBudget**](docs/ManagementAPI.md#createcampaignstorebudget) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | Create campaign store budget
*ManagementAPI* | [**CreateCollection**](docs/ManagementAPI.md#createcollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | Create campaign-level collection
*ManagementAPI* | [**CreateCoupons**](docs/ManagementAPI.md#createcoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Create coupons
*ManagementAPI* | [**CreateCouponsAsync**](docs/ManagementAPI.md#createcouponsasync) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_async | Create coupons asynchronously
*ManagementAPI* | [**CreateCouponsDeletionJob**](docs/ManagementAPI.md#createcouponsdeletionjob) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_deletion_jobs | Creates a coupon deletion job
*ManagementAPI* | [**CreateCouponsForMultipleRecipients**](docs/ManagementAPI.md#createcouponsformultiplerecipients) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_with_recipients | Create coupons for multiple recipients
*ManagementAPI* | [**CreateInviteEmail**](docs/ManagementAPI.md#createinviteemail) | **Post** /v1/invite_emails | Resend invitation email
*ManagementAPI* | [**CreateInviteV2**](docs/ManagementAPI.md#createinvitev2) | **Post** /v2/invites | Invite user
*ManagementAPI* | [**CreatePasswordRecoveryEmail**](docs/ManagementAPI.md#createpasswordrecoveryemail) | **Post** /v1/password_recovery_emails | Request a password reset
*ManagementAPI* | [**CreateSession**](docs/ManagementAPI.md#createsession) | **Post** /v1/sessions | Create session
*ManagementAPI* | [**CreateStore**](docs/ManagementAPI.md#createstore) | **Post** /v1/applications/{applicationId}/stores | Create store
*ManagementAPI* | [**DeactivateUserByEmail**](docs/ManagementAPI.md#deactivateuserbyemail) | **Post** /v1/users/deactivate | Disable user by email address
*ManagementAPI* | [**DeductLoyaltyCardPoints**](docs/ManagementAPI.md#deductloyaltycardpoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/deduct_points | Deduct points from card
*ManagementAPI* | [**DeleteAccountCollection**](docs/ManagementAPI.md#deleteaccountcollection) | **Delete** /v1/collections/{collectionId} | Delete account-level collection
*ManagementAPI* | [**DeleteAchievement**](docs/ManagementAPI.md#deleteachievement) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Delete achievement
*ManagementAPI* | [**DeleteCampaign**](docs/ManagementAPI.md#deletecampaign) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId} | Delete campaign
*ManagementAPI* | [**DeleteCampaignStoreBudgets**](docs/ManagementAPI.md#deletecampaignstorebudgets) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | Delete campaign store budgets
*ManagementAPI* | [**DeleteCollection**](docs/ManagementAPI.md#deletecollection) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Delete campaign-level collection
*ManagementAPI* | [**DeleteCoupon**](docs/ManagementAPI.md#deletecoupon) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Delete coupon
*ManagementAPI* | [**DeleteCoupons**](docs/ManagementAPI.md#deletecoupons) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Delete coupons
*ManagementAPI* | [**DeleteLoyaltyCard**](docs/ManagementAPI.md#deleteloyaltycard) | **Delete** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Delete loyalty card
*ManagementAPI* | [**DeleteReferral**](docs/ManagementAPI.md#deletereferral) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Delete referral
*ManagementAPI* | [**DeleteStore**](docs/ManagementAPI.md#deletestore) | **Delete** /v1/applications/{applicationId}/stores/{storeId} | Delete store
*ManagementAPI* | [**DeleteUser**](docs/ManagementAPI.md#deleteuser) | **Delete** /v1/users/{userId} | Delete user
*ManagementAPI* | [**DeleteUserByEmail**](docs/ManagementAPI.md#deleteuserbyemail) | **Post** /v1/users/delete | Delete user by email address
*ManagementAPI* | [**DestroySession**](docs/ManagementAPI.md#destroysession) | **Delete** /v1/sessions | Destroy session
*ManagementAPI* | [**DisconnectCampaignStores**](docs/ManagementAPI.md#disconnectcampaignstores) | **Delete** /v1/applications/{applicationId}/campaigns/{campaignId}/stores | Disconnect stores
*ManagementAPI* | [**ExportAccountCollectionItems**](docs/ManagementAPI.md#exportaccountcollectionitems) | **Get** /v1/collections/{collectionId}/export | Export account-level collection&#39;s items
*ManagementAPI* | [**ExportAchievements**](docs/ManagementAPI.md#exportachievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId}/export | Export achievement customer data
*ManagementAPI* | [**ExportAudiencesMemberships**](docs/ManagementAPI.md#exportaudiencesmemberships) | **Get** /v1/audiences/{audienceId}/memberships/export | Export audience members
*ManagementAPI* | [**ExportCampaignStoreBudgets**](docs/ManagementAPI.md#exportcampaignstorebudgets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/export | Export campaign store budgets
*ManagementAPI* | [**ExportCampaignStores**](docs/ManagementAPI.md#exportcampaignstores) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/export | Export stores
*ManagementAPI* | [**ExportCollectionItems**](docs/ManagementAPI.md#exportcollectionitems) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/export | Export campaign-level collection&#39;s items
*ManagementAPI* | [**ExportCoupons**](docs/ManagementAPI.md#exportcoupons) | **Get** /v1/applications/{applicationId}/export_coupons | Export coupons
*ManagementAPI* | [**ExportCustomerSessions**](docs/ManagementAPI.md#exportcustomersessions) | **Get** /v1/applications/{applicationId}/export_customer_sessions | Export customer sessions
*ManagementAPI* | [**ExportCustomersTiers**](docs/ManagementAPI.md#exportcustomerstiers) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customers_tiers | Export customers&#39; tier data
*ManagementAPI* | [**ExportEffects**](docs/ManagementAPI.md#exporteffects) | **Get** /v1/applications/{applicationId}/export_effects | Export triggered effects
*ManagementAPI* | [**ExportLoyaltyBalance**](docs/ManagementAPI.md#exportloyaltybalance) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balance | Export customer loyalty balance to CSV
*ManagementAPI* | [**ExportLoyaltyBalances**](docs/ManagementAPI.md#exportloyaltybalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_customer_balances | Export customer loyalty balances
*ManagementAPI* | [**ExportLoyaltyCardBalances**](docs/ManagementAPI.md#exportloyaltycardbalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/export_card_balances | Export all card transaction logs
*ManagementAPI* | [**ExportLoyaltyCardLedger**](docs/ManagementAPI.md#exportloyaltycardledger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/export_log | Export card&#39;s ledger log
*ManagementAPI* | [**ExportLoyaltyCards**](docs/ManagementAPI.md#exportloyaltycards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/export | Export loyalty cards
*ManagementAPI* | [**ExportLoyaltyLedger**](docs/ManagementAPI.md#exportloyaltyledger) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/export_log | Export customer&#39;s transaction logs
*ManagementAPI* | [**ExportPoolGiveaways**](docs/ManagementAPI.md#exportpoolgiveaways) | **Get** /v1/giveaways/pools/{poolId}/export | Export giveaway codes of a giveaway pool
*ManagementAPI* | [**ExportReferrals**](docs/ManagementAPI.md#exportreferrals) | **Get** /v1/applications/{applicationId}/export_referrals | Export referrals
*ManagementAPI* | [**GetAccessLogsWithoutTotalCount**](docs/ManagementAPI.md#getaccesslogswithouttotalcount) | **Get** /v1/applications/{applicationId}/access_logs/no_total | Get access logs for Application
*ManagementAPI* | [**GetAccount**](docs/ManagementAPI.md#getaccount) | **Get** /v1/accounts/{accountId} | Get account details
*ManagementAPI* | [**GetAccountAnalytics**](docs/ManagementAPI.md#getaccountanalytics) | **Get** /v1/accounts/{accountId}/analytics | Get account analytics
*ManagementAPI* | [**GetAccountCollection**](docs/ManagementAPI.md#getaccountcollection) | **Get** /v1/collections/{collectionId} | Get account-level collection
*ManagementAPI* | [**GetAchievement**](docs/ManagementAPI.md#getachievement) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Get achievement
*ManagementAPI* | [**GetAdditionalCost**](docs/ManagementAPI.md#getadditionalcost) | **Get** /v1/additional_costs/{additionalCostId} | Get additional cost
*ManagementAPI* | [**GetAdditionalCosts**](docs/ManagementAPI.md#getadditionalcosts) | **Get** /v1/additional_costs | List additional costs
*ManagementAPI* | [**GetApplication**](docs/ManagementAPI.md#getapplication) | **Get** /v1/applications/{applicationId} | Get Application
*ManagementAPI* | [**GetApplicationApiHealth**](docs/ManagementAPI.md#getapplicationapihealth) | **Get** /v1/applications/{applicationId}/health_report | Get Application health
*ManagementAPI* | [**GetApplicationCustomer**](docs/ManagementAPI.md#getapplicationcustomer) | **Get** /v1/applications/{applicationId}/customers/{customerId} | Get application&#39;s customer
*ManagementAPI* | [**GetApplicationCustomerFriends**](docs/ManagementAPI.md#getapplicationcustomerfriends) | **Get** /v1/applications/{applicationId}/profile/{integrationId}/friends | List friends referred by customer profile
*ManagementAPI* | [**GetApplicationCustomers**](docs/ManagementAPI.md#getapplicationcustomers) | **Get** /v1/applications/{applicationId}/customers | List application&#39;s customers
*ManagementAPI* | [**GetApplicationCustomersByAttributes**](docs/ManagementAPI.md#getapplicationcustomersbyattributes) | **Post** /v1/applications/{applicationId}/customer_search | List application customers matching the given attributes
*ManagementAPI* | [**GetApplicationEventTypes**](docs/ManagementAPI.md#getapplicationeventtypes) | **Get** /v1/applications/{applicationId}/event_types | List Applications event types
*ManagementAPI* | [**GetApplicationEventsWithoutTotalCount**](docs/ManagementAPI.md#getapplicationeventswithouttotalcount) | **Get** /v1/applications/{applicationId}/events/no_total | List Applications events
*ManagementAPI* | [**GetApplicationSession**](docs/ManagementAPI.md#getapplicationsession) | **Get** /v1/applications/{applicationId}/sessions/{sessionId} | Get Application session
*ManagementAPI* | [**GetApplicationSessions**](docs/ManagementAPI.md#getapplicationsessions) | **Get** /v1/applications/{applicationId}/sessions | List Application sessions
*ManagementAPI* | [**GetApplications**](docs/ManagementAPI.md#getapplications) | **Get** /v1/applications | List Applications
*ManagementAPI* | [**GetAttribute**](docs/ManagementAPI.md#getattribute) | **Get** /v1/attributes/{attributeId} | Get custom attribute
*ManagementAPI* | [**GetAttributes**](docs/ManagementAPI.md#getattributes) | **Get** /v1/attributes | List custom attributes
*ManagementAPI* | [**GetAudienceMemberships**](docs/ManagementAPI.md#getaudiencememberships) | **Get** /v1/audiences/{audienceId}/memberships | List audience members
*ManagementAPI* | [**GetAudiences**](docs/ManagementAPI.md#getaudiences) | **Get** /v1/audiences | List audiences
*ManagementAPI* | [**GetAudiencesAnalytics**](docs/ManagementAPI.md#getaudiencesanalytics) | **Get** /v1/audiences/analytics | List audience analytics
*ManagementAPI* | [**GetCampaign**](docs/ManagementAPI.md#getcampaign) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId} | Get campaign
*ManagementAPI* | [**GetCampaignAnalytics**](docs/ManagementAPI.md#getcampaignanalytics) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/analytics | Get analytics of campaigns
*ManagementAPI* | [**GetCampaignByAttributes**](docs/ManagementAPI.md#getcampaignbyattributes) | **Post** /v1/applications/{applicationId}/campaigns_search | List campaigns that match the given attributes
*ManagementAPI* | [**GetCampaignGroup**](docs/ManagementAPI.md#getcampaigngroup) | **Get** /v1/campaign_groups/{campaignGroupId} | Get campaign access group
*ManagementAPI* | [**GetCampaignGroups**](docs/ManagementAPI.md#getcampaigngroups) | **Get** /v1/campaign_groups | List campaign access groups
*ManagementAPI* | [**GetCampaignTemplates**](docs/ManagementAPI.md#getcampaigntemplates) | **Get** /v1/campaign_templates | List campaign templates
*ManagementAPI* | [**GetCampaigns**](docs/ManagementAPI.md#getcampaigns) | **Get** /v1/applications/{applicationId}/campaigns | List campaigns
*ManagementAPI* | [**GetChanges**](docs/ManagementAPI.md#getchanges) | **Get** /v1/changes | Get audit logs for an account
*ManagementAPI* | [**GetCollection**](docs/ManagementAPI.md#getcollection) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Get campaign-level collection
*ManagementAPI* | [**GetCollectionItems**](docs/ManagementAPI.md#getcollectionitems) | **Get** /v1/collections/{collectionId}/items | Get collection items
*ManagementAPI* | [**GetCouponsWithoutTotalCount**](docs/ManagementAPI.md#getcouponswithouttotalcount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/no_total | List coupons
*ManagementAPI* | [**GetCustomerActivityReport**](docs/ManagementAPI.md#getcustomeractivityreport) | **Get** /v1/applications/{applicationId}/customer_activity_reports/{customerId} | Get customer&#39;s activity report
*ManagementAPI* | [**GetCustomerActivityReportsWithoutTotalCount**](docs/ManagementAPI.md#getcustomeractivityreportswithouttotalcount) | **Get** /v1/applications/{applicationId}/customer_activity_reports/no_total | Get Activity Reports for Application Customers
*ManagementAPI* | [**GetCustomerAnalytics**](docs/ManagementAPI.md#getcustomeranalytics) | **Get** /v1/applications/{applicationId}/customers/{customerId}/analytics | Get customer&#39;s analytics report
*ManagementAPI* | [**GetCustomerProfile**](docs/ManagementAPI.md#getcustomerprofile) | **Get** /v1/customers/{customerId} | Get customer profile
*ManagementAPI* | [**GetCustomerProfileAchievementProgress**](docs/ManagementAPI.md#getcustomerprofileachievementprogress) | **Get** /v1/applications/{applicationId}/achievement_progress/{integrationId} | List customer achievements
*ManagementAPI* | [**GetCustomerProfiles**](docs/ManagementAPI.md#getcustomerprofiles) | **Get** /v1/customers/no_total | List customer profiles
*ManagementAPI* | [**GetCustomersByAttributes**](docs/ManagementAPI.md#getcustomersbyattributes) | **Post** /v1/customer_search/no_total | List customer profiles matching the given attributes
*ManagementAPI* | [**GetDashboardStatistics**](docs/ManagementAPI.md#getdashboardstatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/dashboard | Get statistics for loyalty dashboard
*ManagementAPI* | [**GetEventTypes**](docs/ManagementAPI.md#geteventtypes) | **Get** /v1/event_types | List event types
*ManagementAPI* | [**GetExports**](docs/ManagementAPI.md#getexports) | **Get** /v1/exports | Get exports
*ManagementAPI* | [**GetLoyaltyCard**](docs/ManagementAPI.md#getloyaltycard) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Get loyalty card
*ManagementAPI* | [**GetLoyaltyCardTransactionLogs**](docs/ManagementAPI.md#getloyaltycardtransactionlogs) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/logs | List card&#39;s transactions
*ManagementAPI* | [**GetLoyaltyCards**](docs/ManagementAPI.md#getloyaltycards) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards | List loyalty cards
*ManagementAPI* | [**GetLoyaltyPoints**](docs/ManagementAPI.md#getloyaltypoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId} | Get customer&#39;s full loyalty ledger
*ManagementAPI* | [**GetLoyaltyProgram**](docs/ManagementAPI.md#getloyaltyprogram) | **Get** /v1/loyalty_programs/{loyaltyProgramId} | Get loyalty program
*ManagementAPI* | [**GetLoyaltyProgramTransactions**](docs/ManagementAPI.md#getloyaltyprogramtransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/transactions | List loyalty program transactions
*ManagementAPI* | [**GetLoyaltyPrograms**](docs/ManagementAPI.md#getloyaltyprograms) | **Get** /v1/loyalty_programs | List loyalty programs
*ManagementAPI* | [**GetLoyaltyStatistics**](docs/ManagementAPI.md#getloyaltystatistics) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/statistics | Get loyalty program statistics
*ManagementAPI* | [**GetMessageLogs**](docs/ManagementAPI.md#getmessagelogs) | **Get** /v1/message_logs | List message log entries
*ManagementAPI* | [**GetReferralsWithoutTotalCount**](docs/ManagementAPI.md#getreferralswithouttotalcount) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/no_total | List referrals
*ManagementAPI* | [**GetRoleV2**](docs/ManagementAPI.md#getrolev2) | **Get** /v2/roles/{roleId} | Get role
*ManagementAPI* | [**GetRuleset**](docs/ManagementAPI.md#getruleset) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets/{rulesetId} | Get ruleset
*ManagementAPI* | [**GetRulesets**](docs/ManagementAPI.md#getrulesets) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/rulesets | List campaign rulesets
*ManagementAPI* | [**GetStore**](docs/ManagementAPI.md#getstore) | **Get** /v1/applications/{applicationId}/stores/{storeId} | Get store
*ManagementAPI* | [**GetUser**](docs/ManagementAPI.md#getuser) | **Get** /v1/users/{userId} | Get user
*ManagementAPI* | [**GetUsers**](docs/ManagementAPI.md#getusers) | **Get** /v1/users | List users in account
*ManagementAPI* | [**GetWebhook**](docs/ManagementAPI.md#getwebhook) | **Get** /v1/webhooks/{webhookId} | Get webhook
*ManagementAPI* | [**GetWebhooks**](docs/ManagementAPI.md#getwebhooks) | **Get** /v1/webhooks | List webhooks
*ManagementAPI* | [**ImportAccountCollection**](docs/ManagementAPI.md#importaccountcollection) | **Post** /v1/collections/{collectionId}/import | Import data into existing account-level collection
*ManagementAPI* | [**ImportAllowedList**](docs/ManagementAPI.md#importallowedlist) | **Post** /v1/attributes/{attributeId}/allowed_list/import | Import allowed values for attribute
*ManagementAPI* | [**ImportAudiencesMemberships**](docs/ManagementAPI.md#importaudiencesmemberships) | **Post** /v1/audiences/{audienceId}/memberships/import | Import audience members
*ManagementAPI* | [**ImportCampaignStoreBudget**](docs/ManagementAPI.md#importcampaignstorebudget) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/import | Import campaign store budgets
*ManagementAPI* | [**ImportCampaignStores**](docs/ManagementAPI.md#importcampaignstores) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/import | Import stores
*ManagementAPI* | [**ImportCollection**](docs/ManagementAPI.md#importcollection) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId}/import | Import data into existing campaign-level collection
*ManagementAPI* | [**ImportCoupons**](docs/ManagementAPI.md#importcoupons) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_coupons | Import coupons
*ManagementAPI* | [**ImportLoyaltyCards**](docs/ManagementAPI.md#importloyaltycards) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_cards | Import loyalty cards
*ManagementAPI* | [**ImportLoyaltyCustomersTiers**](docs/ManagementAPI.md#importloyaltycustomerstiers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_customers_tiers | Import customers into loyalty tiers
*ManagementAPI* | [**ImportLoyaltyPoints**](docs/ManagementAPI.md#importloyaltypoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/import_points | Import loyalty points
*ManagementAPI* | [**ImportPoolGiveaways**](docs/ManagementAPI.md#importpoolgiveaways) | **Post** /v1/giveaways/pools/{poolId}/import | Import giveaway codes into a giveaway pool
*ManagementAPI* | [**ImportReferrals**](docs/ManagementAPI.md#importreferrals) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/import_referrals | Import referrals
*ManagementAPI* | [**InviteUserExternal**](docs/ManagementAPI.md#inviteuserexternal) | **Post** /v1/users/invite | Invite user from identity provider
*ManagementAPI* | [**ListAccountCollections**](docs/ManagementAPI.md#listaccountcollections) | **Get** /v1/collections | List collections in account
*ManagementAPI* | [**ListAchievements**](docs/ManagementAPI.md#listachievements) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements | List achievements
*ManagementAPI* | [**ListAllRolesV2**](docs/ManagementAPI.md#listallrolesv2) | **Get** /v2/roles | List roles
*ManagementAPI* | [**ListCampaignStoreBudgetLimits**](docs/ManagementAPI.md#listcampaignstorebudgetlimits) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets | List campaign store budget limits
*ManagementAPI* | [**ListCatalogItems**](docs/ManagementAPI.md#listcatalogitems) | **Get** /v1/catalogs/{catalogId}/items | List items in a catalog
*ManagementAPI* | [**ListCollections**](docs/ManagementAPI.md#listcollections) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/collections | List collections in campaign
*ManagementAPI* | [**ListCollectionsInApplication**](docs/ManagementAPI.md#listcollectionsinapplication) | **Get** /v1/applications/{applicationId}/collections | List collections in Application
*ManagementAPI* | [**ListStores**](docs/ManagementAPI.md#liststores) | **Get** /v1/applications/{applicationId}/stores | List stores
*ManagementAPI* | [**OktaEventHandlerChallenge**](docs/ManagementAPI.md#oktaeventhandlerchallenge) | **Get** /v1/provisioning/okta | Validate Okta API ownership
*ManagementAPI* | [**RemoveLoyaltyPoints**](docs/ManagementAPI.md#removeloyaltypoints) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/deduct_points | Deduct points from customer profile
*ManagementAPI* | [**ResetPassword**](docs/ManagementAPI.md#resetpassword) | **Post** /v1/reset_password | Reset password
*ManagementAPI* | [**ScimCreateGroup**](docs/ManagementAPI.md#scimcreategroup) | **Post** /v1/provisioning/scim/Groups | Create SCIM group
*ManagementAPI* | [**ScimCreateUser**](docs/ManagementAPI.md#scimcreateuser) | **Post** /v1/provisioning/scim/Users | Create SCIM user
*ManagementAPI* | [**ScimDeleteGroup**](docs/ManagementAPI.md#scimdeletegroup) | **Delete** /v1/provisioning/scim/Groups/{groupId} | Delete SCIM group
*ManagementAPI* | [**ScimDeleteUser**](docs/ManagementAPI.md#scimdeleteuser) | **Delete** /v1/provisioning/scim/Users/{userId} | Delete SCIM user
*ManagementAPI* | [**ScimGetGroup**](docs/ManagementAPI.md#scimgetgroup) | **Get** /v1/provisioning/scim/Groups/{groupId} | Get SCIM group
*ManagementAPI* | [**ScimGetGroups**](docs/ManagementAPI.md#scimgetgroups) | **Get** /v1/provisioning/scim/Groups | List SCIM groups
*ManagementAPI* | [**ScimGetResourceTypes**](docs/ManagementAPI.md#scimgetresourcetypes) | **Get** /v1/provisioning/scim/ResourceTypes | List supported SCIM resource types
*ManagementAPI* | [**ScimGetSchemas**](docs/ManagementAPI.md#scimgetschemas) | **Get** /v1/provisioning/scim/Schemas | List supported SCIM schemas
*ManagementAPI* | [**ScimGetServiceProviderConfig**](docs/ManagementAPI.md#scimgetserviceproviderconfig) | **Get** /v1/provisioning/scim/ServiceProviderConfig | Get SCIM service provider configuration
*ManagementAPI* | [**ScimGetUser**](docs/ManagementAPI.md#scimgetuser) | **Get** /v1/provisioning/scim/Users/{userId} | Get SCIM user
*ManagementAPI* | [**ScimGetUsers**](docs/ManagementAPI.md#scimgetusers) | **Get** /v1/provisioning/scim/Users | List SCIM users
*ManagementAPI* | [**ScimPatchGroup**](docs/ManagementAPI.md#scimpatchgroup) | **Patch** /v1/provisioning/scim/Groups/{groupId} | Update SCIM group attributes
*ManagementAPI* | [**ScimPatchUser**](docs/ManagementAPI.md#scimpatchuser) | **Patch** /v1/provisioning/scim/Users/{userId} | Update SCIM user attributes
*ManagementAPI* | [**ScimReplaceGroupAttributes**](docs/ManagementAPI.md#scimreplacegroupattributes) | **Put** /v1/provisioning/scim/Groups/{groupId} | Update SCIM group
*ManagementAPI* | [**ScimReplaceUserAttributes**](docs/ManagementAPI.md#scimreplaceuserattributes) | **Put** /v1/provisioning/scim/Users/{userId} | Update SCIM user
*ManagementAPI* | [**SearchCouponsAdvancedApplicationWideWithoutTotalCount**](docs/ManagementAPI.md#searchcouponsadvancedapplicationwidewithouttotalcount) | **Post** /v1/applications/{applicationId}/coupons_search_advanced/no_total | List coupons that match the given attributes (without total count)
*ManagementAPI* | [**SearchCouponsAdvancedWithoutTotalCount**](docs/ManagementAPI.md#searchcouponsadvancedwithouttotalcount) | **Post** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons_search_advanced/no_total | List coupons that match the given attributes in campaign (without total count)
*ManagementAPI* | [**SummarizeCampaignStoreBudget**](docs/ManagementAPI.md#summarizecampaignstorebudget) | **Get** /v1/applications/{applicationId}/campaigns/{campaignId}/stores/budgets/summary | Get summary of campaign store budgets
*ManagementAPI* | [**TransferLoyaltyCard**](docs/ManagementAPI.md#transferloyaltycard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transfer | Transfer card data
*ManagementAPI* | [**UpdateAccountCollection**](docs/ManagementAPI.md#updateaccountcollection) | **Put** /v1/collections/{collectionId} | Update account-level collection
*ManagementAPI* | [**UpdateAchievement**](docs/ManagementAPI.md#updateachievement) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/achievements/{achievementId} | Update achievement
*ManagementAPI* | [**UpdateAdditionalCost**](docs/ManagementAPI.md#updateadditionalcost) | **Put** /v1/additional_costs/{additionalCostId} | Update additional cost
*ManagementAPI* | [**UpdateAttribute**](docs/ManagementAPI.md#updateattribute) | **Put** /v1/attributes/{attributeId} | Update custom attribute
*ManagementAPI* | [**UpdateCampaign**](docs/ManagementAPI.md#updatecampaign) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId} | Update campaign
*ManagementAPI* | [**UpdateCollection**](docs/ManagementAPI.md#updatecollection) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/collections/{collectionId} | Update campaign-level collection&#39;s description
*ManagementAPI* | [**UpdateCoupon**](docs/ManagementAPI.md#updatecoupon) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons/{couponId} | Update coupon
*ManagementAPI* | [**UpdateCouponBatch**](docs/ManagementAPI.md#updatecouponbatch) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/coupons | Update coupons
*ManagementAPI* | [**UpdateLoyaltyCard**](docs/ManagementAPI.md#updateloyaltycard) | **Put** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId} | Update loyalty card status
*ManagementAPI* | [**UpdateReferral**](docs/ManagementAPI.md#updatereferral) | **Put** /v1/applications/{applicationId}/campaigns/{campaignId}/referrals/{referralId} | Update referral
*ManagementAPI* | [**UpdateRoleV2**](docs/ManagementAPI.md#updaterolev2) | **Put** /v2/roles/{roleId} | Update role
*ManagementAPI* | [**UpdateStore**](docs/ManagementAPI.md#updatestore) | **Put** /v1/applications/{applicationId}/stores/{storeId} | Update store
*ManagementAPI* | [**UpdateUser**](docs/ManagementAPI.md#updateuser) | **Put** /v1/users/{userId} | Update user

## Documentation for models

- [APIError](docs/APIError.md)
- [AcceptCouponEffectProps](docs/AcceptCouponEffectProps.md)
- [AcceptReferralEffectProps](docs/AcceptReferralEffectProps.md)
- [AccessLogEntry](docs/AccessLogEntry.md)
- [Account](docs/Account.md)
- [AccountAdditionalCost](docs/AccountAdditionalCost.md)
- [AccountAnalytics](docs/AccountAnalytics.md)
- [AccountDashboardStatistic](docs/AccountDashboardStatistic.md)
- [AccountDashboardStatisticCampaigns](docs/AccountDashboardStatisticCampaigns.md)
- [AccountDashboardStatisticDiscount](docs/AccountDashboardStatisticDiscount.md)
- [AccountDashboardStatisticLoyaltyPoints](docs/AccountDashboardStatisticLoyaltyPoints.md)
- [AccountDashboardStatisticReferrals](docs/AccountDashboardStatisticReferrals.md)
- [AccountDashboardStatisticRevenue](docs/AccountDashboardStatisticRevenue.md)
- [AccountEntity](docs/AccountEntity.md)
- [AccountLimits](docs/AccountLimits.md)
- [Achievement](docs/Achievement.md)
- [AchievementAdditionalProperties](docs/AchievementAdditionalProperties.md)
- [AchievementBase](docs/AchievementBase.md)
- [AchievementProgress](docs/AchievementProgress.md)
- [AchievementProgressWithDefinition](docs/AchievementProgressWithDefinition.md)
- [AchievementStatusEntry](docs/AchievementStatusEntry.md)
- [ActivateLoyaltyPoints](docs/ActivateLoyaltyPoints.md)
- [ActivateLoyaltyPointsResponse](docs/ActivateLoyaltyPointsResponse.md)
- [ActivateUserRequest](docs/ActivateUserRequest.md)
- [AddFreeItemEffectProps](docs/AddFreeItemEffectProps.md)
- [AddItemCatalogAction](docs/AddItemCatalogAction.md)
- [AddLoyaltyPoints](docs/AddLoyaltyPoints.md)
- [AddLoyaltyPointsEffectProps](docs/AddLoyaltyPointsEffectProps.md)
- [AddPriceAdjustmentCatalogAction](docs/AddPriceAdjustmentCatalogAction.md)
- [AddToAudienceEffectProps](docs/AddToAudienceEffectProps.md)
- [AddedDeductedPointsBalancesNotificationPolicy](docs/AddedDeductedPointsBalancesNotificationPolicy.md)
- [AddedDeductedPointsNotification](docs/AddedDeductedPointsNotification.md)
- [AddedDeductedPointsNotificationPolicy](docs/AddedDeductedPointsNotificationPolicy.md)
- [AdditionalCampaignProperties](docs/AdditionalCampaignProperties.md)
- [AdditionalCost](docs/AdditionalCost.md)
- [AnalyticsDataPoint](docs/AnalyticsDataPoint.md)
- [AnalyticsDataPointWithTrend](docs/AnalyticsDataPointWithTrend.md)
- [AnalyticsDataPointWithTrendAndInfluencedRate](docs/AnalyticsDataPointWithTrendAndInfluencedRate.md)
- [AnalyticsDataPointWithTrendAndUplift](docs/AnalyticsDataPointWithTrendAndUplift.md)
- [AnalyticsProduct](docs/AnalyticsProduct.md)
- [AnalyticsSKU](docs/AnalyticsSKU.md)
- [Application](docs/Application.md)
- [ApplicationAPIKey](docs/ApplicationAPIKey.md)
- [ApplicationAnalyticsDataPoint](docs/ApplicationAnalyticsDataPoint.md)
- [ApplicationApiHealth](docs/ApplicationApiHealth.md)
- [ApplicationCIF](docs/ApplicationCIF.md)
- [ApplicationCIFExpression](docs/ApplicationCIFExpression.md)
- [ApplicationCIFReferences](docs/ApplicationCIFReferences.md)
- [ApplicationCampaignAnalytics](docs/ApplicationCampaignAnalytics.md)
- [ApplicationCampaignStats](docs/ApplicationCampaignStats.md)
- [ApplicationCustomer](docs/ApplicationCustomer.md)
- [ApplicationCustomerEntity](docs/ApplicationCustomerEntity.md)
- [ApplicationEntity](docs/ApplicationEntity.md)
- [ApplicationEvent](docs/ApplicationEvent.md)
- [ApplicationNotification](docs/ApplicationNotification.md)
- [ApplicationReferee](docs/ApplicationReferee.md)
- [ApplicationSession](docs/ApplicationSession.md)
- [ApplicationSessionEntity](docs/ApplicationSessionEntity.md)
- [ApplicationStoreEntity](docs/ApplicationStoreEntity.md)
- [AsyncCouponCreationResponse](docs/AsyncCouponCreationResponse.md)
- [AsyncCouponDeletionJobResponse](docs/AsyncCouponDeletionJobResponse.md)
- [AsyncCouponsData](docs/AsyncCouponsData.md)
- [Attribute](docs/Attribute.md)
- [AttributesMandatory](docs/AttributesMandatory.md)
- [AttributesSettings](docs/AttributesSettings.md)
- [Audience](docs/Audience.md)
- [AudienceAnalytics](docs/AudienceAnalytics.md)
- [AudienceCustomer](docs/AudienceCustomer.md)
- [AudienceIntegrationID](docs/AudienceIntegrationID.md)
- [AudienceMembership](docs/AudienceMembership.md)
- [AudienceReference](docs/AudienceReference.md)
- [AwardGiveawayEffectProps](docs/AwardGiveawayEffectProps.md)
- [BaseCampaign](docs/BaseCampaign.md)
- [BaseLoyaltyProgram](docs/BaseLoyaltyProgram.md)
- [BaseNotification](docs/BaseNotification.md)
- [BaseNotificationEntity](docs/BaseNotificationEntity.md)
- [BaseNotificationWebhook](docs/BaseNotificationWebhook.md)
- [BaseNotifications](docs/BaseNotifications.md)
- [BaseSamlConnection](docs/BaseSamlConnection.md)
- [BestPriorPrice](docs/BestPriorPrice.md)
- [BestPriorPriceMetadata](docs/BestPriorPriceMetadata.md)
- [BestPriorPriceRequest](docs/BestPriorPriceRequest.md)
- [BestPriorPriceRequestTarget](docs/BestPriorPriceRequestTarget.md)
- [Binding](docs/Binding.md)
- [BulkApplicationNotification](docs/BulkApplicationNotification.md)
- [BulkOperationOnCampaigns](docs/BulkOperationOnCampaigns.md)
- [Campaign](docs/Campaign.md)
- [CampaignActivationRequest](docs/CampaignActivationRequest.md)
- [CampaignAnalytics](docs/CampaignAnalytics.md)
- [CampaignBudget](docs/CampaignBudget.md)
- [CampaignCollection](docs/CampaignCollection.md)
- [CampaignCollectionEditedNotification](docs/CampaignCollectionEditedNotification.md)
- [CampaignCollectionEditedNotificationItem](docs/CampaignCollectionEditedNotificationItem.md)
- [CampaignCollectionWithoutPayload](docs/CampaignCollectionWithoutPayload.md)
- [CampaignCopy](docs/CampaignCopy.md)
- [CampaignCreatedNotification](docs/CampaignCreatedNotification.md)
- [CampaignCreatedNotificationItem](docs/CampaignCreatedNotificationItem.md)
- [CampaignDeactivationRequest](docs/CampaignDeactivationRequest.md)
- [CampaignDeletedNotification](docs/CampaignDeletedNotification.md)
- [CampaignDeletedNotificationItem](docs/CampaignDeletedNotificationItem.md)
- [CampaignDetail](docs/CampaignDetail.md)
- [CampaignEditedNotification](docs/CampaignEditedNotification.md)
- [CampaignEditedNotificationItem](docs/CampaignEditedNotificationItem.md)
- [CampaignEntity](docs/CampaignEntity.md)
- [CampaignEvaluationGroup](docs/CampaignEvaluationGroup.md)
- [CampaignEvaluationPosition](docs/CampaignEvaluationPosition.md)
- [CampaignEvaluationTreeChangedMessage](docs/CampaignEvaluationTreeChangedMessage.md)
- [CampaignEvaluationTreeChangedNotification](docs/CampaignEvaluationTreeChangedNotification.md)
- [CampaignGroup](docs/CampaignGroup.md)
- [CampaignGroupEntity](docs/CampaignGroupEntity.md)
- [CampaignLogSummary](docs/CampaignLogSummary.md)
- [CampaignNotificationBase](docs/CampaignNotificationBase.md)
- [CampaignNotificationGeneric](docs/CampaignNotificationGeneric.md)
- [CampaignNotificationItemBase](docs/CampaignNotificationItemBase.md)
- [CampaignNotificationPolicy](docs/CampaignNotificationPolicy.md)
- [CampaignRulesetChangedNotification](docs/CampaignRulesetChangedNotification.md)
- [CampaignRulesetChangedNotificationItem](docs/CampaignRulesetChangedNotificationItem.md)
- [CampaignSearch](docs/CampaignSearch.md)
- [CampaignSet](docs/CampaignSet.md)
- [CampaignSetBranchNode](docs/CampaignSetBranchNode.md)
- [CampaignSetLeafNode](docs/CampaignSetLeafNode.md)
- [CampaignSetNode](docs/CampaignSetNode.md)
- [CampaignStateChangedNotification](docs/CampaignStateChangedNotification.md)
- [CampaignStateChangedNotificationItem](docs/CampaignStateChangedNotificationItem.md)
- [CampaignStoreBudget](docs/CampaignStoreBudget.md)
- [CampaignStoreBudgetLimitConfig](docs/CampaignStoreBudgetLimitConfig.md)
- [CampaignTemplate](docs/CampaignTemplate.md)
- [CampaignTemplateCollection](docs/CampaignTemplateCollection.md)
- [CampaignTemplateCouponReservationSettings](docs/CampaignTemplateCouponReservationSettings.md)
- [CampaignTemplateParams](docs/CampaignTemplateParams.md)
- [CampaignVersions](docs/CampaignVersions.md)
- [CardAddedDeductedPointsBalancesNotificationPolicy](docs/CardAddedDeductedPointsBalancesNotificationPolicy.md)
- [CardAddedDeductedPointsNotification](docs/CardAddedDeductedPointsNotification.md)
- [CardAddedDeductedPointsNotificationPolicy](docs/CardAddedDeductedPointsNotificationPolicy.md)
- [CardExpiringPointsNotificationPolicy](docs/CardExpiringPointsNotificationPolicy.md)
- [CardExpiringPointsNotificationTrigger](docs/CardExpiringPointsNotificationTrigger.md)
- [CardLedgerPointsEntryIntegrationAPI](docs/CardLedgerPointsEntryIntegrationAPI.md)
- [CardLedgerTransactionLogEntry](docs/CardLedgerTransactionLogEntry.md)
- [CardLedgerTransactionLogEntryIntegrationAPI](docs/CardLedgerTransactionLogEntryIntegrationAPI.md)
- [CartItem](docs/CartItem.md)
- [Catalog](docs/Catalog.md)
- [CatalogActionFilter](docs/CatalogActionFilter.md)
- [CatalogItem](docs/CatalogItem.md)
- [CatalogSyncRequest](docs/CatalogSyncRequest.md)
- [CatalogsStrikethroughNotificationPolicy](docs/CatalogsStrikethroughNotificationPolicy.md)
- [Change](docs/Change.md)
- [ChangeLoyaltyTierLevelEffectProps](docs/ChangeLoyaltyTierLevelEffectProps.md)
- [ChangeProfilePassword](docs/ChangeProfilePassword.md)
- [CodeGeneratorSettings](docs/CodeGeneratorSettings.md)
- [Collection](docs/Collection.md)
- [CollectionItem](docs/CollectionItem.md)
- [CollectionWithoutPayload](docs/CollectionWithoutPayload.md)
- [Coupon](docs/Coupon.md)
- [CouponConstraints](docs/CouponConstraints.md)
- [CouponCreatedEffectProps](docs/CouponCreatedEffectProps.md)
- [CouponCreationJob](docs/CouponCreationJob.md)
- [CouponDeletionFilters](docs/CouponDeletionFilters.md)
- [CouponDeletionJob](docs/CouponDeletionJob.md)
- [CouponEntity](docs/CouponEntity.md)
- [CouponLimitConfigs](docs/CouponLimitConfigs.md)
- [CouponRejectionReason](docs/CouponRejectionReason.md)
- [CouponReservations](docs/CouponReservations.md)
- [CouponSearch](docs/CouponSearch.md)
- [CouponValue](docs/CouponValue.md)
- [CouponsNotificationData](docs/CouponsNotificationData.md)
- [CouponsNotificationPolicy](docs/CouponsNotificationPolicy.md)
- [CreateAchievement](docs/CreateAchievement.md)
- [CreateApplicationAPIKey](docs/CreateApplicationAPIKey.md)
- [CreateCouponData](docs/CreateCouponData.md)
- [CreateCoupons200Response](docs/CreateCoupons200Response.md)
- [CreateManagementKey](docs/CreateManagementKey.md)
- [CreateReferralsForMultipleAdvocates201Response](docs/CreateReferralsForMultipleAdvocates201Response.md)
- [CreateTemplateCampaign](docs/CreateTemplateCampaign.md)
- [CreateTemplateCampaignResponse](docs/CreateTemplateCampaignResponse.md)
- [CustomEffect](docs/CustomEffect.md)
- [CustomEffectProps](docs/CustomEffectProps.md)
- [CustomerActivityReport](docs/CustomerActivityReport.md)
- [CustomerAnalytics](docs/CustomerAnalytics.md)
- [CustomerInventory](docs/CustomerInventory.md)
- [CustomerProfile](docs/CustomerProfile.md)
- [CustomerProfileAudienceRequest](docs/CustomerProfileAudienceRequest.md)
- [CustomerProfileAudienceRequestItem](docs/CustomerProfileAudienceRequestItem.md)
- [CustomerProfileEntity](docs/CustomerProfileEntity.md)
- [CustomerProfileIntegrationRequestV2](docs/CustomerProfileIntegrationRequestV2.md)
- [CustomerProfileIntegrationResponseV2](docs/CustomerProfileIntegrationResponseV2.md)
- [CustomerProfileSearchQuery](docs/CustomerProfileSearchQuery.md)
- [CustomerProfileUpdateV2Response](docs/CustomerProfileUpdateV2Response.md)
- [CustomerSession](docs/CustomerSession.md)
- [CustomerSessionV2](docs/CustomerSessionV2.md)
- [DeactivateUserRequest](docs/DeactivateUserRequest.md)
- [DeductLoyaltyPoints](docs/DeductLoyaltyPoints.md)
- [DeductLoyaltyPointsEffectProps](docs/DeductLoyaltyPointsEffectProps.md)
- [DeleteCouponsData](docs/DeleteCouponsData.md)
- [DeleteUserRequest](docs/DeleteUserRequest.md)
- [Effect](docs/Effect.md)
- [EffectEntity](docs/EffectEntity.md)
- [EmailEntity](docs/EmailEntity.md)
- [Endpoint](docs/Endpoint.md)
- [Entity](docs/Entity.md)
- [EntityWithTalangVisibleID](docs/EntityWithTalangVisibleID.md)
- [Environment](docs/Environment.md)
- [ErrorEffectProps](docs/ErrorEffectProps.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [ErrorResponseWithStatus](docs/ErrorResponseWithStatus.md)
- [ErrorSource](docs/ErrorSource.md)
- [EvaluableCampaignIds](docs/EvaluableCampaignIds.md)
- [Event](docs/Event.md)
- [EventType](docs/EventType.md)
- [EventV2](docs/EventV2.md)
- [EventV3](docs/EventV3.md)
- [ExpiringCardPointsData](docs/ExpiringCardPointsData.md)
- [ExpiringCardPointsNotification](docs/ExpiringCardPointsNotification.md)
- [ExpiringCouponsData](docs/ExpiringCouponsData.md)
- [ExpiringCouponsNotification](docs/ExpiringCouponsNotification.md)
- [ExpiringCouponsNotificationPolicy](docs/ExpiringCouponsNotificationPolicy.md)
- [ExpiringCouponsNotificationTrigger](docs/ExpiringCouponsNotificationTrigger.md)
- [ExpiringPointsData](docs/ExpiringPointsData.md)
- [ExpiringPointsNotification](docs/ExpiringPointsNotification.md)
- [ExpiringPointsNotificationPolicy](docs/ExpiringPointsNotificationPolicy.md)
- [ExpiringPointsNotificationTrigger](docs/ExpiringPointsNotificationTrigger.md)
- [Export](docs/Export.md)
- [ExtendLoyaltyPointsExpiryDateEffectProps](docs/ExtendLoyaltyPointsExpiryDateEffectProps.md)
- [ExtendedCoupon](docs/ExtendedCoupon.md)
- [FeatureFlag](docs/FeatureFlag.md)
- [FeaturesFeed](docs/FeaturesFeed.md)
- [FuncArgDef](docs/FuncArgDef.md)
- [FunctionDef](docs/FunctionDef.md)
- [GenerateAuditLogSummary](docs/GenerateAuditLogSummary.md)
- [GenerateCampaignDescription](docs/GenerateCampaignDescription.md)
- [GenerateCampaignTags](docs/GenerateCampaignTags.md)
- [GenerateCouponFailureDetailedSummary](docs/GenerateCouponFailureDetailedSummary.md)
- [GenerateCouponFailureSummary](docs/GenerateCouponFailureSummary.md)
- [GenerateItemFilterDescription](docs/GenerateItemFilterDescription.md)
- [GenerateLoyaltyCard](docs/GenerateLoyaltyCard.md)
- [GenerateRuleTitle](docs/GenerateRuleTitle.md)
- [GenerateRuleTitleRule](docs/GenerateRuleTitleRule.md)
- [GenerateUserSessionSummary](docs/GenerateUserSessionSummary.md)
- [GetAccessLogsWithoutTotalCount200Response](docs/GetAccessLogsWithoutTotalCount200Response.md)
- [GetAdditionalCosts200Response](docs/GetAdditionalCosts200Response.md)
- [GetApplicationCustomerFriends200Response](docs/GetApplicationCustomerFriends200Response.md)
- [GetApplicationCustomers200Response](docs/GetApplicationCustomers200Response.md)
- [GetApplicationCustomersByAttributes200Response](docs/GetApplicationCustomersByAttributes200Response.md)
- [GetApplicationEventTypes200Response](docs/GetApplicationEventTypes200Response.md)
- [GetApplicationEventsWithoutTotalCount200Response](docs/GetApplicationEventsWithoutTotalCount200Response.md)
- [GetApplicationSessions200Response](docs/GetApplicationSessions200Response.md)
- [GetApplications200Response](docs/GetApplications200Response.md)
- [GetAttributes200Response](docs/GetAttributes200Response.md)
- [GetAudienceMemberships200Response](docs/GetAudienceMemberships200Response.md)
- [GetAudiences200Response](docs/GetAudiences200Response.md)
- [GetAudiencesAnalytics200Response](docs/GetAudiencesAnalytics200Response.md)
- [GetCampaignAnalytics200Response](docs/GetCampaignAnalytics200Response.md)
- [GetCampaignGroups200Response](docs/GetCampaignGroups200Response.md)
- [GetCampaignTemplates200Response](docs/GetCampaignTemplates200Response.md)
- [GetCampaigns200Response](docs/GetCampaigns200Response.md)
- [GetChanges200Response](docs/GetChanges200Response.md)
- [GetCollectionItems200Response](docs/GetCollectionItems200Response.md)
- [GetCouponsWithoutTotalCount200Response](docs/GetCouponsWithoutTotalCount200Response.md)
- [GetCustomerAchievementHistory200Response](docs/GetCustomerAchievementHistory200Response.md)
- [GetCustomerAchievements200Response](docs/GetCustomerAchievements200Response.md)
- [GetCustomerActivityReportsWithoutTotalCount200Response](docs/GetCustomerActivityReportsWithoutTotalCount200Response.md)
- [GetCustomerProfileAchievementProgress200Response](docs/GetCustomerProfileAchievementProgress200Response.md)
- [GetCustomerProfiles200Response](docs/GetCustomerProfiles200Response.md)
- [GetCustomersByAttributes200Response](docs/GetCustomersByAttributes200Response.md)
- [GetDashboardStatistics200Response](docs/GetDashboardStatistics200Response.md)
- [GetEventTypes200Response](docs/GetEventTypes200Response.md)
- [GetExports200Response](docs/GetExports200Response.md)
- [GetIntegrationCouponRequest](docs/GetIntegrationCouponRequest.md)
- [GetLoyaltyCardPoints200Response](docs/GetLoyaltyCardPoints200Response.md)
- [GetLoyaltyCardTransactionLogs200Response](docs/GetLoyaltyCardTransactionLogs200Response.md)
- [GetLoyaltyCardTransactions200Response](docs/GetLoyaltyCardTransactions200Response.md)
- [GetLoyaltyCards200Response](docs/GetLoyaltyCards200Response.md)
- [GetLoyaltyProgramProfilePoints200Response](docs/GetLoyaltyProgramProfilePoints200Response.md)
- [GetLoyaltyProgramProfileTransactions200Response](docs/GetLoyaltyProgramProfileTransactions200Response.md)
- [GetLoyaltyProgramTransactions200Response](docs/GetLoyaltyProgramTransactions200Response.md)
- [GetLoyaltyPrograms200Response](docs/GetLoyaltyPrograms200Response.md)
- [GetReferralsWithoutTotalCount200Response](docs/GetReferralsWithoutTotalCount200Response.md)
- [GetReservedCustomers200Response](docs/GetReservedCustomers200Response.md)
- [GetRulesets200Response](docs/GetRulesets200Response.md)
- [GetUsers200Response](docs/GetUsers200Response.md)
- [GetWebhooks200Response](docs/GetWebhooks200Response.md)
- [Giveaway](docs/Giveaway.md)
- [GiveawaysPool](docs/GiveawaysPool.md)
- [HiddenConditionsEffects](docs/HiddenConditionsEffects.md)
- [IdentifiableEntity](docs/IdentifiableEntity.md)
- [Import](docs/Import.md)
- [ImportEntity](docs/ImportEntity.md)
- [IncreaseAchievementProgressEffectProps](docs/IncreaseAchievementProgressEffectProps.md)
- [IntegrationCoupon](docs/IntegrationCoupon.md)
- [IntegrationCustomerProfileAudienceRequest](docs/IntegrationCustomerProfileAudienceRequest.md)
- [IntegrationCustomerProfileAudienceRequestItem](docs/IntegrationCustomerProfileAudienceRequestItem.md)
- [IntegrationCustomerSessionResponse](docs/IntegrationCustomerSessionResponse.md)
- [IntegrationEntity](docs/IntegrationEntity.md)
- [IntegrationEvent](docs/IntegrationEvent.md)
- [IntegrationEventV2Request](docs/IntegrationEventV2Request.md)
- [IntegrationEventV3Request](docs/IntegrationEventV3Request.md)
- [IntegrationEventV3Response](docs/IntegrationEventV3Response.md)
- [IntegrationProfileEntity](docs/IntegrationProfileEntity.md)
- [IntegrationProfileEntityV3](docs/IntegrationProfileEntityV3.md)
- [IntegrationRequest](docs/IntegrationRequest.md)
- [IntegrationState](docs/IntegrationState.md)
- [IntegrationStateV2](docs/IntegrationStateV2.md)
- [IntegrationStoreEntity](docs/IntegrationStoreEntity.md)
- [InventoryCoupon](docs/InventoryCoupon.md)
- [InventoryReferral](docs/InventoryReferral.md)
- [ItemAttribute](docs/ItemAttribute.md)
- [JWT](docs/JWT.md)
- [LabelTargetAudience](docs/LabelTargetAudience.md)
- [LabelTargetNone](docs/LabelTargetNone.md)
- [LedgerEntry](docs/LedgerEntry.md)
- [LedgerInfo](docs/LedgerInfo.md)
- [LedgerPointsEntryIntegrationAPI](docs/LedgerPointsEntryIntegrationAPI.md)
- [LedgerTransactionLogEntryIntegrationAPI](docs/LedgerTransactionLogEntryIntegrationAPI.md)
- [LibraryAttribute](docs/LibraryAttribute.md)
- [LimitConfig](docs/LimitConfig.md)
- [LimitCounter](docs/LimitCounter.md)
- [ListAccountCollections200Response](docs/ListAccountCollections200Response.md)
- [ListAchievements200Response](docs/ListAchievements200Response.md)
- [ListAllRolesV2200Response](docs/ListAllRolesV2200Response.md)
- [ListCampaignStoreBudgetLimits200Response](docs/ListCampaignStoreBudgetLimits200Response.md)
- [ListCampaignStoreBudgets](docs/ListCampaignStoreBudgets.md)
- [ListCampaignStoreBudgetsStore](docs/ListCampaignStoreBudgetsStore.md)
- [ListCatalogItems200Response](docs/ListCatalogItems200Response.md)
- [ListStores200Response](docs/ListStores200Response.md)
- [LoginParams](docs/LoginParams.md)
- [Loyalty](docs/Loyalty.md)
- [LoyaltyBalance](docs/LoyaltyBalance.md)
- [LoyaltyBalanceWithTier](docs/LoyaltyBalanceWithTier.md)
- [LoyaltyBalances](docs/LoyaltyBalances.md)
- [LoyaltyBalancesWithTiers](docs/LoyaltyBalancesWithTiers.md)
- [LoyaltyCard](docs/LoyaltyCard.md)
- [LoyaltyCardBalances](docs/LoyaltyCardBalances.md)
- [LoyaltyCardBatch](docs/LoyaltyCardBatch.md)
- [LoyaltyCardBatchResponse](docs/LoyaltyCardBatchResponse.md)
- [LoyaltyCardProfileRegistration](docs/LoyaltyCardProfileRegistration.md)
- [LoyaltyCardRegistration](docs/LoyaltyCardRegistration.md)
- [LoyaltyDashboardData](docs/LoyaltyDashboardData.md)
- [LoyaltyDashboardPointsBreakdown](docs/LoyaltyDashboardPointsBreakdown.md)
- [LoyaltyLedger](docs/LoyaltyLedger.md)
- [LoyaltyLedgerEntry](docs/LoyaltyLedgerEntry.md)
- [LoyaltyLedgerEntryExpiryDateChange](docs/LoyaltyLedgerEntryExpiryDateChange.md)
- [LoyaltyLedgerEntryFlags](docs/LoyaltyLedgerEntryFlags.md)
- [LoyaltyLedgerTransactions](docs/LoyaltyLedgerTransactions.md)
- [LoyaltyMembership](docs/LoyaltyMembership.md)
- [LoyaltyProgram](docs/LoyaltyProgram.md)
- [LoyaltyProgramBalance](docs/LoyaltyProgramBalance.md)
- [LoyaltyProgramEntity](docs/LoyaltyProgramEntity.md)
- [LoyaltyProgramLedgers](docs/LoyaltyProgramLedgers.md)
- [LoyaltyProgramTransaction](docs/LoyaltyProgramTransaction.md)
- [LoyaltySubLedger](docs/LoyaltySubLedger.md)
- [LoyaltyTier](docs/LoyaltyTier.md)
- [ManagementKey](docs/ManagementKey.md)
- [ManagerConfig](docs/ManagerConfig.md)
- [MessageLogEntries](docs/MessageLogEntries.md)
- [MessageLogEntry](docs/MessageLogEntry.md)
- [MessageLogRequest](docs/MessageLogRequest.md)
- [MessageLogResponse](docs/MessageLogResponse.md)
- [MessageTest](docs/MessageTest.md)
- [Meta](docs/Meta.md)
- [MultiApplicationEntity](docs/MultiApplicationEntity.md)
- [MultipleAttribute](docs/MultipleAttribute.md)
- [MultipleAudiences](docs/MultipleAudiences.md)
- [MultipleAudiencesItem](docs/MultipleAudiencesItem.md)
- [MultipleCustomerProfileIntegrationRequest](docs/MultipleCustomerProfileIntegrationRequest.md)
- [MultipleCustomerProfileIntegrationRequestItem](docs/MultipleCustomerProfileIntegrationRequestItem.md)
- [MultipleCustomerProfileIntegrationResponseV2](docs/MultipleCustomerProfileIntegrationResponseV2.md)
- [MultipleNewAttribute](docs/MultipleNewAttribute.md)
- [MultipleNewAudiences](docs/MultipleNewAudiences.md)
- [MutableEntity](docs/MutableEntity.md)
- [NewAccount](docs/NewAccount.md)
- [NewAccountSignUp](docs/NewAccountSignUp.md)
- [NewAdditionalCost](docs/NewAdditionalCost.md)
- [NewAppWideCouponDeletionJob](docs/NewAppWideCouponDeletionJob.md)
- [NewApplication](docs/NewApplication.md)
- [NewApplicationAPIKey](docs/NewApplicationAPIKey.md)
- [NewApplicationCIF](docs/NewApplicationCIF.md)
- [NewApplicationCIFExpression](docs/NewApplicationCIFExpression.md)
- [NewAttribute](docs/NewAttribute.md)
- [NewAudience](docs/NewAudience.md)
- [NewBaseNotification](docs/NewBaseNotification.md)
- [NewCampaign](docs/NewCampaign.md)
- [NewCampaignCollection](docs/NewCampaignCollection.md)
- [NewCampaignEvaluationGroup](docs/NewCampaignEvaluationGroup.md)
- [NewCampaignGroup](docs/NewCampaignGroup.md)
- [NewCampaignSet](docs/NewCampaignSet.md)
- [NewCampaignStoreBudget](docs/NewCampaignStoreBudget.md)
- [NewCampaignStoreBudgetStoreLimit](docs/NewCampaignStoreBudgetStoreLimit.md)
- [NewCampaignTemplate](docs/NewCampaignTemplate.md)
- [NewCatalog](docs/NewCatalog.md)
- [NewCollection](docs/NewCollection.md)
- [NewCouponCreationJob](docs/NewCouponCreationJob.md)
- [NewCouponDeletionJob](docs/NewCouponDeletionJob.md)
- [NewCoupons](docs/NewCoupons.md)
- [NewCouponsForMultipleRecipients](docs/NewCouponsForMultipleRecipients.md)
- [NewCustomEffect](docs/NewCustomEffect.md)
- [NewCustomerProfile](docs/NewCustomerProfile.md)
- [NewCustomerSession](docs/NewCustomerSession.md)
- [NewCustomerSessionV2](docs/NewCustomerSessionV2.md)
- [NewEvent](docs/NewEvent.md)
- [NewEventType](docs/NewEventType.md)
- [NewExternalInvitation](docs/NewExternalInvitation.md)
- [NewGiveawaysPool](docs/NewGiveawaysPool.md)
- [NewInternalAudience](docs/NewInternalAudience.md)
- [NewInvitation](docs/NewInvitation.md)
- [NewInviteEmail](docs/NewInviteEmail.md)
- [NewLoyaltyProgram](docs/NewLoyaltyProgram.md)
- [NewLoyaltyTier](docs/NewLoyaltyTier.md)
- [NewManagementKey](docs/NewManagementKey.md)
- [NewMessageTest](docs/NewMessageTest.md)
- [NewMultipleAudiencesItem](docs/NewMultipleAudiencesItem.md)
- [NewNotificationWebhook](docs/NewNotificationWebhook.md)
- [NewOutgoingIntegrationWebhook](docs/NewOutgoingIntegrationWebhook.md)
- [NewPassword](docs/NewPassword.md)
- [NewPasswordEmail](docs/NewPasswordEmail.md)
- [NewPicklist](docs/NewPicklist.md)
- [NewPriceAdjustment](docs/NewPriceAdjustment.md)
- [NewPriceType](docs/NewPriceType.md)
- [NewReferral](docs/NewReferral.md)
- [NewReferralsForMultipleAdvocates](docs/NewReferralsForMultipleAdvocates.md)
- [NewReturn](docs/NewReturn.md)
- [NewRevisionVersion](docs/NewRevisionVersion.md)
- [NewRole](docs/NewRole.md)
- [NewRoleV2](docs/NewRoleV2.md)
- [NewRuleset](docs/NewRuleset.md)
- [NewSamlConnection](docs/NewSamlConnection.md)
- [NewSecondaryDeployment](docs/NewSecondaryDeployment.md)
- [NewStore](docs/NewStore.md)
- [NewTemplateDef](docs/NewTemplateDef.md)
- [NewUser](docs/NewUser.md)
- [NewWebhook](docs/NewWebhook.md)
- [Notification](docs/Notification.md)
- [NotificationActivation](docs/NotificationActivation.md)
- [NotificationListItem](docs/NotificationListItem.md)
- [OktaEvent](docs/OktaEvent.md)
- [OktaEventPayload](docs/OktaEventPayload.md)
- [OktaEventPayloadData](docs/OktaEventPayloadData.md)
- [OktaEventTarget](docs/OktaEventTarget.md)
- [OneTimeCode](docs/OneTimeCode.md)
- [OutgoingIntegrationBrazePolicy](docs/OutgoingIntegrationBrazePolicy.md)
- [OutgoingIntegrationCleverTapPolicy](docs/OutgoingIntegrationCleverTapPolicy.md)
- [OutgoingIntegrationConfiguration](docs/OutgoingIntegrationConfiguration.md)
- [OutgoingIntegrationIterablePolicy](docs/OutgoingIntegrationIterablePolicy.md)
- [OutgoingIntegrationMoEngagePolicy](docs/OutgoingIntegrationMoEngagePolicy.md)
- [OutgoingIntegrationTemplate](docs/OutgoingIntegrationTemplate.md)
- [OutgoingIntegrationTemplateWithConfigurationDetails](docs/OutgoingIntegrationTemplateWithConfigurationDetails.md)
- [OutgoingIntegrationTemplates](docs/OutgoingIntegrationTemplates.md)
- [OutgoingIntegrationType](docs/OutgoingIntegrationType.md)
- [OutgoingIntegrationTypes](docs/OutgoingIntegrationTypes.md)
- [PatchItemCatalogAction](docs/PatchItemCatalogAction.md)
- [PatchManyItemsCatalogAction](docs/PatchManyItemsCatalogAction.md)
- [PendingActivePointsData](docs/PendingActivePointsData.md)
- [PendingActivePointsNotification](docs/PendingActivePointsNotification.md)
- [PendingPointsNotificationPolicy](docs/PendingPointsNotificationPolicy.md)
- [Picklist](docs/Picklist.md)
- [PriceDetail](docs/PriceDetail.md)
- [PriceType](docs/PriceType.md)
- [PriceTypeReferenceDetail](docs/PriceTypeReferenceDetail.md)
- [PriceTypeReferences](docs/PriceTypeReferences.md)
- [PrismaticFlow](docs/PrismaticFlow.md)
- [PrismaticFlowConfig](docs/PrismaticFlowConfig.md)
- [PrismaticFlowWithConfig](docs/PrismaticFlowWithConfig.md)
- [Product](docs/Product.md)
- [ProductSearchMatch](docs/ProductSearchMatch.md)
- [ProductUnitAnalytics](docs/ProductUnitAnalytics.md)
- [ProductUnitAnalyticsDataPoint](docs/ProductUnitAnalyticsDataPoint.md)
- [ProductUnitAnalyticsTotals](docs/ProductUnitAnalyticsTotals.md)
- [ProfileAudiencesChanges](docs/ProfileAudiencesChanges.md)
- [ProjectedTier](docs/ProjectedTier.md)
- [RedeemReferralEffectProps](docs/RedeemReferralEffectProps.md)
- [Referral](docs/Referral.md)
- [ReferralConstraints](docs/ReferralConstraints.md)
- [ReferralCreatedEffectProps](docs/ReferralCreatedEffectProps.md)
- [ReferralRejectionReason](docs/ReferralRejectionReason.md)
- [RejectCouponEffectProps](docs/RejectCouponEffectProps.md)
- [RejectReferralEffectProps](docs/RejectReferralEffectProps.md)
- [RemoveFromAudienceEffectProps](docs/RemoveFromAudienceEffectProps.md)
- [RemoveItemCatalogAction](docs/RemoveItemCatalogAction.md)
- [RemoveManyItemsCatalogAction](docs/RemoveManyItemsCatalogAction.md)
- [ReopenSessionResponse](docs/ReopenSessionResponse.md)
- [ReserveCouponEffectProps](docs/ReserveCouponEffectProps.md)
- [Return](docs/Return.md)
- [ReturnIntegrationRequest](docs/ReturnIntegrationRequest.md)
- [ReturnedCartItem](docs/ReturnedCartItem.md)
- [Revision](docs/Revision.md)
- [RevisionActivation](docs/RevisionActivation.md)
- [RevisionActivationRequest](docs/RevisionActivationRequest.md)
- [RevisionVersion](docs/RevisionVersion.md)
- [Role](docs/Role.md)
- [RoleAssign](docs/RoleAssign.md)
- [RoleMembership](docs/RoleMembership.md)
- [RoleV2](docs/RoleV2.md)
- [RoleV2ApplicationDetails](docs/RoleV2ApplicationDetails.md)
- [RoleV2Base](docs/RoleV2Base.md)
- [RoleV2PermissionSet](docs/RoleV2PermissionSet.md)
- [RoleV2Permissions](docs/RoleV2Permissions.md)
- [RoleV2RolesGroup](docs/RoleV2RolesGroup.md)
- [RollbackAddedLoyaltyPointsEffectProps](docs/RollbackAddedLoyaltyPointsEffectProps.md)
- [RollbackCouponEffectProps](docs/RollbackCouponEffectProps.md)
- [RollbackDeductedLoyaltyPointsEffectProps](docs/RollbackDeductedLoyaltyPointsEffectProps.md)
- [RollbackDiscountEffectProps](docs/RollbackDiscountEffectProps.md)
- [RollbackIncreasedAchievementProgressEffectProps](docs/RollbackIncreasedAchievementProgressEffectProps.md)
- [RollbackReferralEffectProps](docs/RollbackReferralEffectProps.md)
- [Rule](docs/Rule.md)
- [RuleFailureReason](docs/RuleFailureReason.md)
- [Ruleset](docs/Ruleset.md)
- [SSOConfig](docs/SSOConfig.md)
- [SamlConnection](docs/SamlConnection.md)
- [SamlConnectionInternal](docs/SamlConnectionInternal.md)
- [SamlConnectionMetadata](docs/SamlConnectionMetadata.md)
- [SamlLoginEndpoint](docs/SamlLoginEndpoint.md)
- [ScimBaseGroup](docs/ScimBaseGroup.md)
- [ScimBaseUser](docs/ScimBaseUser.md)
- [ScimBaseUserName](docs/ScimBaseUserName.md)
- [ScimGroup](docs/ScimGroup.md)
- [ScimGroupMember](docs/ScimGroupMember.md)
- [ScimGroupsListResponse](docs/ScimGroupsListResponse.md)
- [ScimNewUser](docs/ScimNewUser.md)
- [ScimPatchOperation](docs/ScimPatchOperation.md)
- [ScimPatchRequest](docs/ScimPatchRequest.md)
- [ScimResource](docs/ScimResource.md)
- [ScimResourceTypesListResponse](docs/ScimResourceTypesListResponse.md)
- [ScimSchemaResource](docs/ScimSchemaResource.md)
- [ScimSchemasListResponse](docs/ScimSchemasListResponse.md)
- [ScimServiceProviderConfigResponse](docs/ScimServiceProviderConfigResponse.md)
- [ScimServiceProviderConfigResponseBulk](docs/ScimServiceProviderConfigResponseBulk.md)
- [ScimServiceProviderConfigResponseChangePassword](docs/ScimServiceProviderConfigResponseChangePassword.md)
- [ScimServiceProviderConfigResponseFilter](docs/ScimServiceProviderConfigResponseFilter.md)
- [ScimServiceProviderConfigResponsePatch](docs/ScimServiceProviderConfigResponsePatch.md)
- [ScimServiceProviderConfigResponseSort](docs/ScimServiceProviderConfigResponseSort.md)
- [ScimUser](docs/ScimUser.md)
- [ScimUsersListResponse](docs/ScimUsersListResponse.md)
- [SecondaryDeployment](docs/SecondaryDeployment.md)
- [Session](docs/Session.md)
- [SetDiscountEffectProps](docs/SetDiscountEffectProps.md)
- [SetDiscountPerAdditionalCostEffectProps](docs/SetDiscountPerAdditionalCostEffectProps.md)
- [SetDiscountPerAdditionalCostPerItemEffectProps](docs/SetDiscountPerAdditionalCostPerItemEffectProps.md)
- [SetDiscountPerItemEffectProps](docs/SetDiscountPerItemEffectProps.md)
- [SetLoyaltyPointsExpiryDateEffectProps](docs/SetLoyaltyPointsExpiryDateEffectProps.md)
- [ShowBundleMetadataEffectProps](docs/ShowBundleMetadataEffectProps.md)
- [ShowNotificationEffectProps](docs/ShowNotificationEffectProps.md)
- [SkuUnitAnalytics](docs/SkuUnitAnalytics.md)
- [SkuUnitAnalyticsDataPoint](docs/SkuUnitAnalyticsDataPoint.md)
- [SlotDef](docs/SlotDef.md)
- [Store](docs/Store.md)
- [StrikethroughChangedItem](docs/StrikethroughChangedItem.md)
- [StrikethroughCustomEffectPerItemProps](docs/StrikethroughCustomEffectPerItemProps.md)
- [StrikethroughDebugResponse](docs/StrikethroughDebugResponse.md)
- [StrikethroughEffect](docs/StrikethroughEffect.md)
- [StrikethroughLabelingNotification](docs/StrikethroughLabelingNotification.md)
- [StrikethroughSetDiscountPerItemEffectProps](docs/StrikethroughSetDiscountPerItemEffectProps.md)
- [StrikethroughTrigger](docs/StrikethroughTrigger.md)
- [SummarizeCampaignStoreBudget200Response](docs/SummarizeCampaignStoreBudget200Response.md)
- [SummaryCampaignStoreBudget](docs/SummaryCampaignStoreBudget.md)
- [TalangAttribute](docs/TalangAttribute.md)
- [TalangAttributeVisibility](docs/TalangAttributeVisibility.md)
- [TemplateArgDef](docs/TemplateArgDef.md)
- [TemplateDef](docs/TemplateDef.md)
- [TemplateLimitConfig](docs/TemplateLimitConfig.md)
- [Tier](docs/Tier.md)
- [TierDowngradeData](docs/TierDowngradeData.md)
- [TierDowngradeNotification](docs/TierDowngradeNotification.md)
- [TierDowngradeNotificationPolicy](docs/TierDowngradeNotificationPolicy.md)
- [TierUpgradeData](docs/TierUpgradeData.md)
- [TierUpgradeNotification](docs/TierUpgradeNotification.md)
- [TierUpgradeNotificationPolicy](docs/TierUpgradeNotificationPolicy.md)
- [TierWillDowngradeData](docs/TierWillDowngradeData.md)
- [TierWillDowngradeNotification](docs/TierWillDowngradeNotification.md)
- [TierWillDowngradeNotificationPolicy](docs/TierWillDowngradeNotificationPolicy.md)
- [TierWillDowngradeNotificationTrigger](docs/TierWillDowngradeNotificationTrigger.md)
- [TimePoint](docs/TimePoint.md)
- [TrackEventV2Response](docs/TrackEventV2Response.md)
- [TransferLoyaltyCard](docs/TransferLoyaltyCard.md)
- [TriggerWebhookEffectProps](docs/TriggerWebhookEffectProps.md)
- [TwoFAConfig](docs/TwoFAConfig.md)
- [UpdateAccount](docs/UpdateAccount.md)
- [UpdateAchievement](docs/UpdateAchievement.md)
- [UpdateApplication](docs/UpdateApplication.md)
- [UpdateApplicationAPIKey](docs/UpdateApplicationAPIKey.md)
- [UpdateApplicationCIF](docs/UpdateApplicationCIF.md)
- [UpdateAttributeEffectProps](docs/UpdateAttributeEffectProps.md)
- [UpdateAudience](docs/UpdateAudience.md)
- [UpdateCampaign](docs/UpdateCampaign.md)
- [UpdateCampaignCollection](docs/UpdateCampaignCollection.md)
- [UpdateCampaignEvaluationGroup](docs/UpdateCampaignEvaluationGroup.md)
- [UpdateCampaignGroup](docs/UpdateCampaignGroup.md)
- [UpdateCampaignTemplate](docs/UpdateCampaignTemplate.md)
- [UpdateCatalog](docs/UpdateCatalog.md)
- [UpdateCollection](docs/UpdateCollection.md)
- [UpdateCoupon](docs/UpdateCoupon.md)
- [UpdateCouponBatch](docs/UpdateCouponBatch.md)
- [UpdateCouponsData](docs/UpdateCouponsData.md)
- [UpdateCustomEffect](docs/UpdateCustomEffect.md)
- [UpdateCustomerProfileV2409Response](docs/UpdateCustomerProfileV2409Response.md)
- [UpdateCustomerSessionV2409Response](docs/UpdateCustomerSessionV2409Response.md)
- [UpdateLoyaltyCard](docs/UpdateLoyaltyCard.md)
- [UpdateLoyaltyProgram](docs/UpdateLoyaltyProgram.md)
- [UpdateLoyaltyProgramTier](docs/UpdateLoyaltyProgramTier.md)
- [UpdatePicklist](docs/UpdatePicklist.md)
- [UpdatePriceType](docs/UpdatePriceType.md)
- [UpdateReferral](docs/UpdateReferral.md)
- [UpdateReferralBatch](docs/UpdateReferralBatch.md)
- [UpdateRole](docs/UpdateRole.md)
- [UpdateStore](docs/UpdateStore.md)
- [UpdateUser](docs/UpdateUser.md)
- [User](docs/User.md)
- [UserEntity](docs/UserEntity.md)
- [ValueMap](docs/ValueMap.md)
- [Webhook](docs/Webhook.md)
- [WebhookAuthentication](docs/WebhookAuthentication.md)
- [WebhookAuthenticationBase](docs/WebhookAuthenticationBase.md)
- [WebhookAuthenticationDataBasic](docs/WebhookAuthenticationDataBasic.md)
- [WebhookAuthenticationDataCustom](docs/WebhookAuthenticationDataCustom.md)
- [WebhookAuthenticationWebhookRef](docs/WebhookAuthenticationWebhookRef.md)
- [WebhookWithOutgoingIntegrationDetails](docs/WebhookWithOutgoingIntegrationDetails.md)
- [WillAwardGiveawayEffectProps](docs/WillAwardGiveawayEffectProps.md)

## Authorization

### api_key_v1

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

### management_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

### manager_auth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

## Utility methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`
