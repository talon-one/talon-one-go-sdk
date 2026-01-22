# \IntegrationAPI

All URIs are relative to *https://yourbaseurl.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateLoyaltyPoints**](IntegrationAPI.md#ActivateLoyaltyPoints) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/activate_points | Activate loyalty points
[**BestPriorPrice**](IntegrationAPI.md#BestPriorPrice) | **Post** /v1/best_prior_price | Fetch best prior price
[**CreateAudienceV2**](IntegrationAPI.md#CreateAudienceV2) | **Post** /v2/audiences | Create audience
[**CreateCouponReservation**](IntegrationAPI.md#CreateCouponReservation) | **Post** /v1/coupon_reservations/{couponValue} | Create coupon reservation
[**CreateReferral**](IntegrationAPI.md#CreateReferral) | **Post** /v1/referrals | Create referral code for an advocate
[**CreateReferralsForMultipleAdvocates**](IntegrationAPI.md#CreateReferralsForMultipleAdvocates) | **Post** /v1/referrals_for_multiple_advocates | Create referral codes for multiple advocates
[**DeleteAudienceMembershipsV2**](IntegrationAPI.md#DeleteAudienceMembershipsV2) | **Delete** /v2/audiences/{audienceId}/memberships | Delete audience memberships
[**DeleteAudienceV2**](IntegrationAPI.md#DeleteAudienceV2) | **Delete** /v2/audiences/{audienceId} | Delete audience
[**DeleteCouponReservation**](IntegrationAPI.md#DeleteCouponReservation) | **Delete** /v1/coupon_reservations/{couponValue} | Delete coupon reservations
[**DeleteCustomerData**](IntegrationAPI.md#DeleteCustomerData) | **Delete** /v1/customer_data/{integrationId} | Delete customer&#39;s personal data
[**DeleteLoyaltyTransactionsFromLedgers**](IntegrationAPI.md#DeleteLoyaltyTransactionsFromLedgers) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/delete_transactions | Delete customer&#39;s transactions from loyalty ledgers
[**GenerateLoyaltyCard**](IntegrationAPI.md#GenerateLoyaltyCard) | **Post** /v1/loyalty_programs/{loyaltyProgramId}/cards | Generate loyalty card
[**GetCustomerAchievementHistory**](IntegrationAPI.md#GetCustomerAchievementHistory) | **Get** /v1/customer_profiles/{integrationId}/achievements/{achievementId} | List customer&#39;s achievement history
[**GetCustomerAchievements**](IntegrationAPI.md#GetCustomerAchievements) | **Get** /v1/customer_profiles/{integrationId}/achievements | List customer&#39;s available achievements
[**GetCustomerInventory**](IntegrationAPI.md#GetCustomerInventory) | **Get** /v1/customer_profiles/{integrationId}/inventory | List customer data
[**GetCustomerSession**](IntegrationAPI.md#GetCustomerSession) | **Get** /v2/customer_sessions/{customerSessionId} | Get customer session
[**GetLoyaltyBalances**](IntegrationAPI.md#GetLoyaltyBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/balances | Get customer&#39;s loyalty balances
[**GetLoyaltyCardBalances**](IntegrationAPI.md#GetLoyaltyCardBalances) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/balances | Get card&#39;s point balances
[**GetLoyaltyCardPoints**](IntegrationAPI.md#GetLoyaltyCardPoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/points | List card&#39;s unused loyalty points
[**GetLoyaltyCardTransactions**](IntegrationAPI.md#GetLoyaltyCardTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/transactions | List card&#39;s transactions
[**GetLoyaltyProgramProfilePoints**](IntegrationAPI.md#GetLoyaltyProgramProfilePoints) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/points | List customer&#39;s unused loyalty points
[**GetLoyaltyProgramProfileTransactions**](IntegrationAPI.md#GetLoyaltyProgramProfileTransactions) | **Get** /v1/loyalty_programs/{loyaltyProgramId}/profile/{integrationId}/transactions | List customer&#39;s loyalty transactions
[**GetReservedCustomers**](IntegrationAPI.md#GetReservedCustomers) | **Get** /v1/coupon_reservations/customerprofiles/{couponValue} | List customers that have this coupon reserved
[**LinkLoyaltyCardToProfile**](IntegrationAPI.md#LinkLoyaltyCardToProfile) | **Post** /v2/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/link_profile | Link customer profile to card
[**ReopenCustomerSession**](IntegrationAPI.md#ReopenCustomerSession) | **Put** /v2/customer_sessions/{customerSessionId}/reopen | Reopen customer session
[**ReturnCartItems**](IntegrationAPI.md#ReturnCartItems) | **Post** /v2/customer_sessions/{customerSessionId}/returns | Return cart items
[**SyncCatalog**](IntegrationAPI.md#SyncCatalog) | **Put** /v1/catalogs/{catalogId}/sync | Sync cart item catalog
[**TrackEventV2**](IntegrationAPI.md#TrackEventV2) | **Post** /v2/events | Track event
[**UnlinkLoyaltyCardFromProfile**](IntegrationAPI.md#UnlinkLoyaltyCardFromProfile) | **Post** /v2/loyalty_programs/{loyaltyProgramId}/cards/{loyaltyCardId}/unlink_profile | Unlink customer profile from a loyalty card
[**UpdateAudienceCustomersAttributes**](IntegrationAPI.md#UpdateAudienceCustomersAttributes) | **Put** /v2/audience_customers/{audienceId}/attributes | Update profile attributes for all customers in audience
[**UpdateAudienceV2**](IntegrationAPI.md#UpdateAudienceV2) | **Put** /v2/audiences/{audienceId} | Update audience name
[**UpdateCustomerProfileAudiences**](IntegrationAPI.md#UpdateCustomerProfileAudiences) | **Post** /v2/customer_audiences | Update multiple customer profiles&#39; audiences
[**UpdateCustomerProfileV2**](IntegrationAPI.md#UpdateCustomerProfileV2) | **Put** /v2/customer_profiles/{integrationId} | Update customer profile
[**UpdateCustomerProfilesV2**](IntegrationAPI.md#UpdateCustomerProfilesV2) | **Put** /v2/customer_profiles | Update multiple customer profiles
[**UpdateCustomerSessionV2**](IntegrationAPI.md#UpdateCustomerSessionV2) | **Put** /v2/customer_sessions/{customerSessionId} | Update customer session



## ActivateLoyaltyPoints

> ActivateLoyaltyPointsResponse ActivateLoyaltyPoints(ctx, loyaltyProgramId).ActivateLoyaltyPoints(activateLoyaltyPoints).Execute()

Activate loyalty points



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
	loyaltyProgramId := int64(789) // int64 | The identifier for the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	activateLoyaltyPoints := *openapiclient.NewActivateLoyaltyPoints() // ActivateLoyaltyPoints | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.ActivateLoyaltyPoints(context.Background(), loyaltyProgramId).ActivateLoyaltyPoints(activateLoyaltyPoints).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ActivateLoyaltyPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateLoyaltyPoints`: ActivateLoyaltyPointsResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ActivateLoyaltyPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | The identifier for the loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateLoyaltyPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activateLoyaltyPoints** | [**ActivateLoyaltyPoints**](ActivateLoyaltyPoints.md) | body | 

### Return type

[**ActivateLoyaltyPointsResponse**](ActivateLoyaltyPointsResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BestPriorPrice

> []BestPriorPrice BestPriorPrice(ctx).BestPriorPriceRequest(bestPriorPriceRequest).Execute()

Fetch best prior price



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
	bestPriorPriceRequest := *openapiclient.NewBestPriorPriceRequest([]string{"Skus_example"}, time.Now(), "30", true) // BestPriorPriceRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.BestPriorPrice(context.Background()).BestPriorPriceRequest(bestPriorPriceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.BestPriorPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BestPriorPrice`: []BestPriorPrice
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.BestPriorPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBestPriorPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bestPriorPriceRequest** | [**BestPriorPriceRequest**](BestPriorPriceRequest.md) | body | 

### Return type

[**[]BestPriorPrice**](BestPriorPrice.md)

### Authorization

[management_key](../README.md#management_key), [manager_auth](../README.md#manager_auth), [api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAudienceV2

> Audience CreateAudienceV2(ctx).NewAudience(newAudience).Execute()

Create audience



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
	newAudience := *openapiclient.NewNewAudience("Travel audience") // NewAudience | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.CreateAudienceV2(context.Background()).NewAudience(newAudience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CreateAudienceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAudienceV2`: Audience
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CreateAudienceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newAudience** | [**NewAudience**](NewAudience.md) | body | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponReservation

> Coupon CreateCouponReservation(ctx, couponValue).CouponReservations(couponReservations).Execute()

Create coupon reservation



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
	couponValue := "couponValue_example" // string | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode `SUMMER25%OFF` as `SUMMER25%25OFF`. 
	couponReservations := *openapiclient.NewCouponReservations([]string{"IntegrationIDs_example"}) // CouponReservations | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.CreateCouponReservation(context.Background(), couponValue).CouponReservations(couponReservations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CreateCouponReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCouponReservation`: Coupon
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CreateCouponReservation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode &#x60;SUMMER25%OFF&#x60; as &#x60;SUMMER25%25OFF&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **couponReservations** | [**CouponReservations**](CouponReservations.md) | body | 

### Return type

[**Coupon**](Coupon.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferral

> Referral CreateReferral(ctx).NewReferral(newReferral).Execute()

Create referral code for an advocate



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
	newReferral := *openapiclient.NewNewReferral(int64(78), "URNGV8294NV") // NewReferral | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.CreateReferral(context.Background()).NewReferral(newReferral).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CreateReferral``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReferral`: Referral
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CreateReferral`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newReferral** | [**NewReferral**](NewReferral.md) | body | 

### Return type

[**Referral**](Referral.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateReferralsForMultipleAdvocates

> CreateReferralsForMultipleAdvocates201Response CreateReferralsForMultipleAdvocates(ctx).NewReferralsForMultipleAdvocates(newReferralsForMultipleAdvocates).Silent(silent).Execute()

Create referral codes for multiple advocates



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
	newReferralsForMultipleAdvocates := *openapiclient.NewNewReferralsForMultipleAdvocates(int64(45), []string{"AdvocateProfileIntegrationIds_example"}) // NewReferralsForMultipleAdvocates | body
	silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the performance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains the updated customer profiles.  (optional) (default to "yes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.CreateReferralsForMultipleAdvocates(context.Background()).NewReferralsForMultipleAdvocates(newReferralsForMultipleAdvocates).Silent(silent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.CreateReferralsForMultipleAdvocates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateReferralsForMultipleAdvocates`: CreateReferralsForMultipleAdvocates201Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.CreateReferralsForMultipleAdvocates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReferralsForMultipleAdvocatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newReferralsForMultipleAdvocates** | [**NewReferralsForMultipleAdvocates**](NewReferralsForMultipleAdvocates.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to &quot;yes&quot;]

### Return type

[**CreateReferralsForMultipleAdvocates201Response**](CreateReferralsForMultipleAdvocates201Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceMembershipsV2

> DeleteAudienceMembershipsV2(ctx, audienceId).Execute()

Delete audience memberships



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
	r, err := apiClient.IntegrationAPI.DeleteAudienceMembershipsV2(context.Background(), audienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteAudienceMembershipsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceMembershipsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAudienceV2

> DeleteAudienceV2(ctx, audienceId).Execute()

Delete audience



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
	r, err := apiClient.IntegrationAPI.DeleteAudienceV2(context.Background(), audienceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteAudienceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCouponReservation

> DeleteCouponReservation(ctx, couponValue).CouponReservations(couponReservations).Execute()

Delete coupon reservations



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
	couponValue := "couponValue_example" // string | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode `SUMMER25%OFF` as `SUMMER25%25OFF`. 
	couponReservations := *openapiclient.NewCouponReservations([]string{"IntegrationIDs_example"}) // CouponReservations | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.DeleteCouponReservation(context.Background(), couponValue).CouponReservations(couponReservations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteCouponReservation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode &#x60;SUMMER25%OFF&#x60; as &#x60;SUMMER25%25OFF&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCouponReservationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **couponReservations** | [**CouponReservations**](CouponReservations.md) | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomerData

> DeleteCustomerData(ctx, integrationId).Execute()

Delete customer's personal data



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
	integrationId := "integrationId_example" // string | The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.DeleteCustomerData(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteCustomerData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomerDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoyaltyTransactionsFromLedgers

> DeleteLoyaltyTransactionsFromLedgers(ctx, loyaltyProgramId, integrationId).DeleteLoyaltyTransactionsRequest(deleteLoyaltyTransactionsRequest).Execute()

Delete customer's transactions from loyalty ledgers



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
	loyaltyProgramId := int64(789) // int64 | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	integrationId := "integrationId_example" // string | The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 
	deleteLoyaltyTransactionsRequest := *openapiclient.NewDeleteLoyaltyTransactionsRequest("SelectedSubledgers") // DeleteLoyaltyTransactionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.DeleteLoyaltyTransactionsFromLedgers(context.Background(), loyaltyProgramId, integrationId).DeleteLoyaltyTransactionsRequest(deleteLoyaltyTransactionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.DeleteLoyaltyTransactionsFromLedgers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoyaltyTransactionsFromLedgersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **deleteLoyaltyTransactionsRequest** | [**DeleteLoyaltyTransactionsRequest**](DeleteLoyaltyTransactionsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateLoyaltyCard

> LoyaltyCard GenerateLoyaltyCard(ctx, loyaltyProgramId).GenerateLoyaltyCard(generateLoyaltyCard).Execute()

Generate loyalty card



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
	generateLoyaltyCard := *openapiclient.NewGenerateLoyaltyCard() // GenerateLoyaltyCard | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GenerateLoyaltyCard(context.Background(), loyaltyProgramId).GenerateLoyaltyCard(generateLoyaltyCard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GenerateLoyaltyCard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateLoyaltyCard`: LoyaltyCard
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GenerateLoyaltyCard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateLoyaltyCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateLoyaltyCard** | [**GenerateLoyaltyCard**](GenerateLoyaltyCard.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAchievementHistory

> GetCustomerAchievementHistory200Response GetCustomerAchievementHistory(ctx, integrationId, achievementId).ProgressStatus(progressStatus).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).Execute()

List customer's achievement history



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
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	achievementId := int64(789) // int64 | The achievement identifier. 
	progressStatus := []string{"ProgressStatus_example"} // []string | Filter by customer progress status in the achievement.  (optional)
	startDate := time.Now() // time.Time | Timestamp that filters the results to only contain achievements created on or after the start date. (optional)
	endDate := time.Now() // time.Time | Timestamp that filters the results to only contain achievements created before or on the end date. (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetCustomerAchievementHistory(context.Background(), integrationId, achievementId).ProgressStatus(progressStatus).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetCustomerAchievementHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerAchievementHistory`: GetCustomerAchievementHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetCustomerAchievementHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 
**achievementId** | **int64** | The achievement identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAchievementHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **progressStatus** | **[]string** | Filter by customer progress status in the achievement.  | 
 **startDate** | **time.Time** | Timestamp that filters the results to only contain achievements created on or after the start date. | 
 **endDate** | **time.Time** | Timestamp that filters the results to only contain achievements created before or on the end date. | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 

### Return type

[**GetCustomerAchievementHistory200Response**](GetCustomerAchievementHistory200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerAchievements

> GetCustomerAchievements200Response GetCustomerAchievements(ctx, integrationId).CampaignIds(campaignIds).AchievementIds(achievementIds).AchievementStatus(achievementStatus).CurrentProgressStatus(currentProgressStatus).PageSize(pageSize).Skip(skip).Execute()

List customer's available achievements



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
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	campaignIds := []string{"Inner_example"} // []string | Filter by one or more Campaign IDs, separated by a comma.  **Note:** If no campaigns are specified, data for all the campaigns in the Application is returned.  (optional)
	achievementIds := []string{"Inner_example"} // []string | Filter by one or more Achievement IDs, separated by a comma.  **Note:** If no achievements are specified, data for all the achievements in the Application is returned.  (optional)
	achievementStatus := []string{"AchievementStatus_example"} // []string | Filter by status of the achievement.  **Note:** If the achievement status is not specified, only data for all active achievements in the Application is returned.  (optional)
	currentProgressStatus := []string{"CurrentProgressStatus_example"} // []string | Filter by customer progress status in the achievement.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 1000)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetCustomerAchievements(context.Background(), integrationId).CampaignIds(campaignIds).AchievementIds(achievementIds).AchievementStatus(achievementStatus).CurrentProgressStatus(currentProgressStatus).PageSize(pageSize).Skip(skip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetCustomerAchievements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerAchievements`: GetCustomerAchievements200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetCustomerAchievements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerAchievementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignIds** | **[]string** | Filter by one or more Campaign IDs, separated by a comma.  **Note:** If no campaigns are specified, data for all the campaigns in the Application is returned.  | 
 **achievementIds** | **[]string** | Filter by one or more Achievement IDs, separated by a comma.  **Note:** If no achievements are specified, data for all the achievements in the Application is returned.  | 
 **achievementStatus** | **[]string** | Filter by status of the achievement.  **Note:** If the achievement status is not specified, only data for all active achievements in the Application is returned.  | 
 **currentProgressStatus** | **[]string** | Filter by customer progress status in the achievement.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 1000]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 

### Return type

[**GetCustomerAchievements200Response**](GetCustomerAchievements200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerInventory

> CustomerInventory GetCustomerInventory(ctx, integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).Achievements(achievements).Execute()

List customer data



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
	integrationId := "integrationId_example" // string | The integration ID of the customer profile. You can get the `integrationId` of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application's customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint. 
	profile := true // bool | Set to `true` to include customer profile information in the response. (optional)
	referrals := true // bool | Set to `true` to include referral information in the response. (optional)
	coupons := true // bool | Set to `true` to include coupon information in the response. (optional)
	loyalty := true // bool | Set to `true` to include loyalty information in the response. (optional)
	giveaways := true // bool | Set to `true` to include giveaways information in the response. (optional)
	achievements := true // bool | Set to `true` to include achievement information in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetCustomerInventory(context.Background(), integrationId).Profile(profile).Referrals(referrals).Coupons(coupons).Loyalty(loyalty).Giveaways(giveaways).Achievements(achievements).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetCustomerInventory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerInventory`: CustomerInventory
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetCustomerInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration ID of the customer profile. You can get the &#x60;integrationId&#x60; of a profile using: - A customer session integration ID with the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint. - The Management API with the [List application&#39;s customers](https://docs.talon.one/management-api#operation/getApplicationCustomers) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **profile** | **bool** | Set to &#x60;true&#x60; to include customer profile information in the response. | 
 **referrals** | **bool** | Set to &#x60;true&#x60; to include referral information in the response. | 
 **coupons** | **bool** | Set to &#x60;true&#x60; to include coupon information in the response. | 
 **loyalty** | **bool** | Set to &#x60;true&#x60; to include loyalty information in the response. | 
 **giveaways** | **bool** | Set to &#x60;true&#x60; to include giveaways information in the response. | 
 **achievements** | **bool** | Set to &#x60;true&#x60; to include achievement information in the response. | 

### Return type

[**CustomerInventory**](CustomerInventory.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomerSession

> IntegrationCustomerSessionResponse GetCustomerSession(ctx, customerSessionId).Execute()

Get customer session



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
	customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetCustomerSession(context.Background(), customerSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetCustomerSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomerSession`: IntegrationCustomerSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetCustomerSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IntegrationCustomerSessionResponse**](IntegrationCustomerSessionResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyBalances

> LoyaltyBalancesWithTiers GetLoyaltyBalances(ctx, loyaltyProgramId, integrationId).EndDate(endDate).SubledgerId(subledgerId).IncludeTiers(includeTiers).IncludeProjectedTier(includeProjectedTier).Execute()

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
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)
	includeTiers := true // bool | Indicates whether tier information is included in the response.  When set to `true`, the response includes information about the current tier and the number of points required to move to next tier.  (optional) (default to false)
	includeProjectedTier := true // bool | Indicates whether the customer's projected tier information is included in the response.  When set to `true`, the response includes information about the customer's active points and the name of the projected tier.  **Note** We recommend filtering by `subledgerId` for better performance.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyBalances(context.Background(), loyaltyProgramId, integrationId).EndDate(endDate).SubledgerId(subledgerId).IncludeTiers(includeTiers).IncludeProjectedTier(includeProjectedTier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyBalances`: LoyaltyBalancesWithTiers
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **includeTiers** | **bool** | Indicates whether tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the current tier and the number of points required to move to next tier.  | [default to false]
 **includeProjectedTier** | **bool** | Indicates whether the customer&#39;s projected tier information is included in the response.  When set to &#x60;true&#x60;, the response includes information about the customer&#39;s active points and the name of the projected tier.  **Note** We recommend filtering by &#x60;subledgerId&#x60; for better performance.  | [default to false]

### Return type

[**LoyaltyBalancesWithTiers**](LoyaltyBalancesWithTiers.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardBalances

> LoyaltyCardBalances GetLoyaltyCardBalances(ctx, loyaltyProgramId, loyaltyCardId).EndDate(endDate).SubledgerId(subledgerId).Execute()

Get card's point balances



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
	endDate := time.Now() // time.Time | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	subledgerId := []string{"Inner_example"} // []string | Filter results by one or more subledger IDs. Must be exact match. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyCardBalances(context.Background(), loyaltyProgramId, loyaltyCardId).EndDate(endDate).SubledgerId(subledgerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyCardBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCardBalances`: LoyaltyCardBalances
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyCardBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endDate** | **time.Time** | Used to return expired, active, and pending loyalty balances before this timestamp. You can enter any past, present, or future timestamp value.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **subledgerId** | **[]string** | Filter results by one or more subledger IDs. Must be exact match. | 

### Return type

[**LoyaltyCardBalances**](LoyaltyCardBalances.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardPoints

> GetLoyaltyCardPoints200Response GetLoyaltyCardPoints(ctx, loyaltyProgramId, loyaltyCardId).Status(status).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List card's unused loyalty points



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
	status := "status_example" // string | Filter points based on their status. (optional) (default to "active")
	subledgerId := []string{"Inner_example"} // []string | Filter results by one or more subledger IDs. Must be exact match. (optional)
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. You can enter one of the following values:  - `startDate`: Sorts the results by the start date of the points. - `expiryDate`: Sorts the results by the expiry date of the points.  By default, results are sorted in ascending order.  To sort them in descending order, prefix the field name with `-`.  **Note:** You can only sort by one field at a time.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyCardPoints(context.Background(), loyaltyProgramId, loyaltyCardId).Status(status).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyCardPoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCardPoints`: GetLoyaltyCardPoints200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyCardPoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardPointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter points based on their status. | [default to &quot;active&quot;]
 **subledgerId** | **[]string** | Filter results by one or more subledger IDs. Must be exact match. | 
 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. You can enter one of the following values:  - &#x60;startDate&#x60;: Sorts the results by the start date of the points. - &#x60;expiryDate&#x60;: Sorts the results by the expiry date of the points.  By default, results are sorted in ascending order.  To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You can only sort by one field at a time.  | 

### Return type

[**GetLoyaltyCardPoints200Response**](GetLoyaltyCardPoints200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyCardTransactions

> GetLoyaltyCardTransactions200Response GetLoyaltyCardTransactions(ctx, loyaltyProgramId, loyaltyCardId).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()

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
	subledgerId := []string{"Inner_example"} // []string | Filter results by one or more subledger IDs. Must be exact match. (optional)
	loyaltyTransactionType := "loyaltyTransactionType_example" // string | Filter results by loyalty transaction type: - `manual`: Loyalty transaction that was done manually. - `session`: Loyalty transaction that resulted from a customer session. - `import`: Loyalty transaction that was imported from a CSV file.  (optional)
	startDate := time.Now() // time.Time | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	endDate := time.Now() // time.Time | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	awaitsActivation := true // bool | If `true`: Filters results to include only point transactions that have action-based activation and have not expired.  If `false`: Returns a `400` response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyCardTransactions(context.Background(), loyaltyProgramId, loyaltyCardId).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyCardTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyCardTransactions`: GetLoyaltyCardTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyCardTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyCardTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **subledgerId** | **[]string** | Filter results by one or more subledger IDs. Must be exact match. | 
 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **awaitsActivation** | **bool** | If &#x60;true&#x60;: Filters results to include only point transactions that have action-based activation and have not expired.  If &#x60;false&#x60;: Returns a &#x60;400&#x60; response.  | 

### Return type

[**GetLoyaltyCardTransactions200Response**](GetLoyaltyCardTransactions200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfilePoints

> GetLoyaltyProgramProfilePoints200Response GetLoyaltyProgramProfilePoints(ctx, loyaltyProgramId, integrationId).Status(status).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).Sort(sort).Execute()

List customer's unused loyalty points



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
	loyaltyProgramId := int64(789) // int64 | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	status := "status_example" // string | Filter points based on their status. (optional) (default to "active")
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)
	customerSessionIDs := []string{"Inner_example"} // []string | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  `?customerSessionIDs=id1&customerSessionIDs=id2`.  The response contains only data associated with the specified sessions.  (optional)
	transactionUUIDs := []string{"Inner_example"} // []string | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  `?transactionUUIDs=uuid1&transactionUUIDs=uuid2`.  The response contains only data associated with the specified transactions.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	sort := "sort_example" // string | The field by which results should be sorted. You can enter one of the following values:  - `startDate`: Sorts the results by the start date of the points. - `expiryDate`: Sorts the results by the expiry date of the points.  By default, results are sorted in ascending order.  To sort them in descending order, prefix the field name with `-`.  **Note:** You can only sort by one field at a time.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyProgramProfilePoints(context.Background(), loyaltyProgramId, integrationId).Status(status).SubledgerId(subledgerId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).PageSize(pageSize).Skip(skip).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyProgramProfilePoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramProfilePoints`: GetLoyaltyProgramProfilePoints200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyProgramProfilePoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramProfilePointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **string** | Filter points based on their status. | [default to &quot;active&quot;]
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **sort** | **string** | The field by which results should be sorted. You can enter one of the following values:  - &#x60;startDate&#x60;: Sorts the results by the start date of the points. - &#x60;expiryDate&#x60;: Sorts the results by the expiry date of the points.  By default, results are sorted in ascending order.  To sort them in descending order, prefix the field name with &#x60;-&#x60;.  **Note:** You can only sort by one field at a time.  | 

### Return type

[**GetLoyaltyProgramProfilePoints200Response**](GetLoyaltyProgramProfilePoints200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoyaltyProgramProfileTransactions

> GetLoyaltyProgramProfileTransactions200Response GetLoyaltyProgramProfileTransactions(ctx, loyaltyProgramId, integrationId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()

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
	subledgerId := "subledgerId_example" // string | The ID of the subledger by which we filter the data. (optional)
	loyaltyTransactionType := "loyaltyTransactionType_example" // string | Filter results by loyalty transaction type: - `manual`: Loyalty transaction that was done manually. - `session`: Loyalty transaction that resulted from a customer session. - `import`: Loyalty transaction that was imported from a CSV file.  (optional)
	startDate := time.Now() // time.Time | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	endDate := time.Now() // time.Time | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, `T23:59:59` to specify the end of the day. The time zone setting considered is `UTC`. If you do not include a time component, a default time value of `T00:00:00` (midnight) in `UTC` is considered.  (optional)
	pageSize := int64(789) // int64 | The number of items in the response. (optional) (default to 50)
	skip := int64(789) // int64 | The number of items to skip when paging through large result sets. (optional)
	awaitsActivation := true // bool | If `true`: Filters results to include only point transactions that have action-based activation and have not expired.  If `false`: Returns a `400` response.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetLoyaltyProgramProfileTransactions(context.Background(), loyaltyProgramId, integrationId).CustomerSessionIDs(customerSessionIDs).TransactionUUIDs(transactionUUIDs).SubledgerId(subledgerId).LoyaltyTransactionType(loyaltyTransactionType).StartDate(startDate).EndDate(endDate).PageSize(pageSize).Skip(skip).AwaitsActivation(awaitsActivation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetLoyaltyProgramProfileTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoyaltyProgramProfileTransactions`: GetLoyaltyProgramProfileTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetLoyaltyProgramProfileTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the profile-based loyalty program. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramProfileTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customerSessionIDs** | **[]string** | Filter the results by a list of customer session IDs.   To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?customerSessionIDs&#x3D;id1&amp;customerSessionIDs&#x3D;id2&#x60;.  The response contains only data associated with the specified sessions.  | 
 **transactionUUIDs** | **[]string** | Filter the results by a list of transaction UUIDs.  To include multiple IDs, repeat the parameter for each one, for example,  &#x60;?transactionUUIDs&#x3D;uuid1&amp;transactionUUIDs&#x3D;uuid2&#x60;.  The response contains only data associated with the specified transactions.  | 
 **subledgerId** | **string** | The ID of the subledger by which we filter the data. | 
 **loyaltyTransactionType** | **string** | Filter results by loyalty transaction type: - &#x60;manual&#x60;: Loyalty transaction that was done manually. - &#x60;session&#x60;: Loyalty transaction that resulted from a customer session. - &#x60;import&#x60;: Loyalty transaction that was imported from a CSV file.  | 
 **startDate** | **time.Time** | Date and time from which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **endDate** | **time.Time** | Date and time by which results are returned. Results are filtered by transaction creation date.  **Note:**  - It must be an RFC3339 timestamp string. - You can include a time component in your string, for example, &#x60;T23:59:59&#x60; to specify the end of the day. The time zone setting considered is &#x60;UTC&#x60;. If you do not include a time component, a default time value of &#x60;T00:00:00&#x60; (midnight) in &#x60;UTC&#x60; is considered.  | 
 **pageSize** | **int64** | The number of items in the response. | [default to 50]
 **skip** | **int64** | The number of items to skip when paging through large result sets. | 
 **awaitsActivation** | **bool** | If &#x60;true&#x60;: Filters results to include only point transactions that have action-based activation and have not expired.  If &#x60;false&#x60;: Returns a &#x60;400&#x60; response.  | 

### Return type

[**GetLoyaltyProgramProfileTransactions200Response**](GetLoyaltyProgramProfileTransactions200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReservedCustomers

> GetReservedCustomers200Response GetReservedCustomers(ctx, couponValue).Execute()

List customers that have this coupon reserved



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
	couponValue := "couponValue_example" // string | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode `SUMMER25%OFF` as `SUMMER25%25OFF`. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.GetReservedCustomers(context.Background(), couponValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.GetReservedCustomers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReservedCustomers`: GetReservedCustomers200Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.GetReservedCustomers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponValue** | **string** | The code of the coupon.  **Important:** The coupon code requires [URL encoding](https://www.w3schools.com/tags//ref_urlencode.asp)  if it contains special characters. For example, you must encode &#x60;SUMMER25%OFF&#x60; as &#x60;SUMMER25%25OFF&#x60;.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReservedCustomersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReservedCustomers200Response**](GetReservedCustomers200Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkLoyaltyCardToProfile

> LoyaltyCard LinkLoyaltyCardToProfile(ctx, loyaltyProgramId, loyaltyCardId).LoyaltyCardRegistration(loyaltyCardRegistration).Execute()

Link customer profile to card



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
	loyaltyCardRegistration := *openapiclient.NewLoyaltyCardRegistration("R195412") // LoyaltyCardRegistration | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.LinkLoyaltyCardToProfile(context.Background(), loyaltyProgramId, loyaltyCardId).LoyaltyCardRegistration(loyaltyCardRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.LinkLoyaltyCardToProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkLoyaltyCardToProfile`: LoyaltyCard
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.LinkLoyaltyCardToProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | Identifier of the card-based loyalty program containing the loyalty card. You can get the ID with the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | Identifier of the loyalty card. You can get the identifier with the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiLinkLoyaltyCardToProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loyaltyCardRegistration** | [**LoyaltyCardRegistration**](LoyaltyCardRegistration.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReopenCustomerSession

> ReopenSessionResponse ReopenCustomerSession(ctx, customerSessionId).Execute()

Reopen customer session



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
	customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.ReopenCustomerSession(context.Background(), customerSessionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ReopenCustomerSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReopenCustomerSession`: ReopenSessionResponse
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ReopenCustomerSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReopenCustomerSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReopenSessionResponse**](ReopenSessionResponse.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReturnCartItems

> IntegrationStateV2 ReturnCartItems(ctx, customerSessionId).ReturnIntegrationRequest(returnIntegrationRequest).Dry(dry).Execute()

Return cart items



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
	customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint. 
	returnIntegrationRequest := *openapiclient.NewReturnIntegrationRequest(*openapiclient.NewNewReturn([]openapiclient.ReturnedCartItem{*openapiclient.NewReturnedCartItem()})) // ReturnIntegrationRequest | body
	dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.ReturnCartItems(context.Background(), customerSessionId).ReturnIntegrationRequest(returnIntegrationRequest).Dry(dry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.ReturnCartItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReturnCartItems`: IntegrationStateV2
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.ReturnCartItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReturnCartItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnIntegrationRequest** | [**ReturnIntegrationRequest**](ReturnIntegrationRequest.md) | body | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncCatalog

> Catalog SyncCatalog(ctx, catalogId).CatalogSyncRequest(catalogSyncRequest).Execute()

Sync cart item catalog



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
	catalogSyncRequest := *openapiclient.NewCatalogSyncRequest([]map[string]interface{}{map[string]interface{}(123)}) // CatalogSyncRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.SyncCatalog(context.Background(), catalogId).CatalogSyncRequest(catalogSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.SyncCatalog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncCatalog`: Catalog
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.SyncCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**catalogId** | **int64** | The ID of the catalog. You can find the ID in the Campaign Manager in **Account** &gt; **Tools** &gt; **Cart item catalogs**. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catalogSyncRequest** | [**CatalogSyncRequest**](CatalogSyncRequest.md) | body | 

### Return type

[**Catalog**](Catalog.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrackEventV2

> TrackEventV2Response TrackEventV2(ctx).IntegrationEventV2Request(integrationEventV2Request).Silent(silent).Dry(dry).ForceCompleteEvaluation(forceCompleteEvaluation).Execute()

Track event



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
	integrationEventV2Request := *openapiclient.NewIntegrationEventV2Request("pageViewed") // IntegrationEventV2Request | body
	silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the performance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains the updated customer profiles.  (optional) (default to "yes")
	dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`.  (optional)
	forceCompleteEvaluation := true // bool | Forces evaluation for all matching campaigns regardless of the [campaign evaluation mode](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation#setting-campaign-evaluation-mode). Requires `dry=true`.  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.TrackEventV2(context.Background()).IntegrationEventV2Request(integrationEventV2Request).Silent(silent).Dry(dry).ForceCompleteEvaluation(forceCompleteEvaluation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.TrackEventV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrackEventV2`: TrackEventV2Response
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.TrackEventV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrackEventV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationEventV2Request** | [**IntegrationEventV2Request**](IntegrationEventV2Request.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to &quot;yes&quot;]
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  | 
 **forceCompleteEvaluation** | **bool** | Forces evaluation for all matching campaigns regardless of the [campaign evaluation mode](https://docs.talon.one/docs/product/applications/managing-campaign-evaluation#setting-campaign-evaluation-mode). Requires &#x60;dry&#x3D;true&#x60;.  | [default to false]

### Return type

[**TrackEventV2Response**](TrackEventV2Response.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkLoyaltyCardFromProfile

> LoyaltyCard UnlinkLoyaltyCardFromProfile(ctx, loyaltyProgramId, loyaltyCardId).LoyaltyCardRegistration(loyaltyCardRegistration).Execute()

Unlink customer profile from a loyalty card



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
	loyaltyProgramId := int64(789) // int64 | The identifier of the card-based loyalty program containing the loyalty card. You can get this ID using the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint. 
	loyaltyCardId := "loyaltyCardId_example" // string | The identifier of the loyalty card. You can get this ID using the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint. 
	loyaltyCardRegistration := *openapiclient.NewLoyaltyCardRegistration("R195412") // LoyaltyCardRegistration | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UnlinkLoyaltyCardFromProfile(context.Background(), loyaltyProgramId, loyaltyCardId).LoyaltyCardRegistration(loyaltyCardRegistration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UnlinkLoyaltyCardFromProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnlinkLoyaltyCardFromProfile`: LoyaltyCard
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UnlinkLoyaltyCardFromProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loyaltyProgramId** | **int64** | The identifier of the card-based loyalty program containing the loyalty card. You can get this ID using the [List loyalty programs](https://docs.talon.one/management-api#tag/Loyalty/operation/getLoyaltyPrograms) endpoint.  | 
**loyaltyCardId** | **string** | The identifier of the loyalty card. You can get this ID using the [List loyalty cards](https://docs.talon.one/management-api#tag/Loyalty-cards/operation/getLoyaltyCards) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkLoyaltyCardFromProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **loyaltyCardRegistration** | [**LoyaltyCardRegistration**](LoyaltyCardRegistration.md) | body | 

### Return type

[**LoyaltyCard**](LoyaltyCard.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceCustomersAttributes

> UpdateAudienceCustomersAttributes(ctx, audienceId).Body(body).Execute()

Update profile attributes for all customers in audience



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
	body := map[string]interface{}{ ... } // map[string]interface{} | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.UpdateAudienceCustomersAttributes(context.Background(), audienceId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateAudienceCustomersAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceCustomersAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAudienceV2

> Audience UpdateAudienceV2(ctx, audienceId).UpdateAudience(updateAudience).Execute()

Update audience name



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
	updateAudience := *openapiclient.NewUpdateAudience("Travel audience") // UpdateAudience | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UpdateAudienceV2(context.Background(), audienceId).UpdateAudience(updateAudience).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateAudienceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAudienceV2`: Audience
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateAudienceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audienceId** | **int64** | The ID of the audience. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAudienceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAudience** | [**UpdateAudience**](UpdateAudience.md) | body | 

### Return type

[**Audience**](Audience.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileAudiences

> UpdateCustomerProfileAudiences(ctx).CustomerProfileAudienceRequest(customerProfileAudienceRequest).Execute()

Update multiple customer profiles' audiences



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
	customerProfileAudienceRequest := *openapiclient.NewCustomerProfileAudienceRequest() // CustomerProfileAudienceRequest | body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IntegrationAPI.UpdateCustomerProfileAudiences(context.Background()).CustomerProfileAudienceRequest(customerProfileAudienceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateCustomerProfileAudiences``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileAudiencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customerProfileAudienceRequest** | [**CustomerProfileAudienceRequest**](CustomerProfileAudienceRequest.md) | body | 

### Return type

 (empty response body)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfileV2

> CustomerProfileIntegrationResponseV2 UpdateCustomerProfileV2(ctx, integrationId).CustomerProfileIntegrationRequestV2(customerProfileIntegrationRequestV2).RunRuleEngine(runRuleEngine).Dry(dry).Execute()

Update customer profile



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
	integrationId := "integrationId_example" // string | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier. 
	customerProfileIntegrationRequestV2 := *openapiclient.NewCustomerProfileIntegrationRequestV2() // CustomerProfileIntegrationRequestV2 | body
	runRuleEngine := true // bool | Indicates whether to run the Rule Engine.  If `true`, the response includes: - The effects generated by the triggered campaigns are returned in the `effects` property. - The created coupons and referral objects.  If `false`: - The rules are not executed and the `effects` property is always empty. - The response time improves. - You cannot use `responseContent` in the body.  (optional) (default to false)
	dry := true // bool | (Only works when `runRuleEngine=true`) Indicates whether to persist the changes. Changes are ignored when `dry=true`.  When set to `true`, you can use the `evaluableCampaignIds` body property to select specific campaigns to run.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UpdateCustomerProfileV2(context.Background(), integrationId).CustomerProfileIntegrationRequestV2(customerProfileIntegrationRequestV2).RunRuleEngine(runRuleEngine).Dry(dry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateCustomerProfileV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomerProfileV2`: CustomerProfileIntegrationResponseV2
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateCustomerProfileV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | The integration identifier for this customer profile. Must be: - Unique within the deployment. - Stable for the customer. Do not use an ID that the customer can update themselves. For example, you can use a database ID.  Once set, you cannot update this identifier.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfileV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customerProfileIntegrationRequestV2** | [**CustomerProfileIntegrationRequestV2**](CustomerProfileIntegrationRequestV2.md) | body | 
 **runRuleEngine** | **bool** | Indicates whether to run the Rule Engine.  If &#x60;true&#x60;, the response includes: - The effects generated by the triggered campaigns are returned in the &#x60;effects&#x60; property. - The created coupons and referral objects.  If &#x60;false&#x60;: - The rules are not executed and the &#x60;effects&#x60; property is always empty. - The response time improves. - You cannot use &#x60;responseContent&#x60; in the body.  | [default to false]
 **dry** | **bool** | (Only works when &#x60;runRuleEngine&#x3D;true&#x60;) Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;, you can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  | 

### Return type

[**CustomerProfileIntegrationResponseV2**](CustomerProfileIntegrationResponseV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerProfilesV2

> MultipleCustomerProfileIntegrationResponseV2 UpdateCustomerProfilesV2(ctx).MultipleCustomerProfileIntegrationRequest(multipleCustomerProfileIntegrationRequest).Silent(silent).Execute()

Update multiple customer profiles



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
	multipleCustomerProfileIntegrationRequest := *openapiclient.NewMultipleCustomerProfileIntegrationRequest() // MultipleCustomerProfileIntegrationRequest | body
	silent := "silent_example" // string | Possible values: `yes` or `no`. - `yes`: Increases the performance of the API call by returning a 204 response. - `no`: Returns a 200 response that contains the updated customer profiles.  (optional) (default to "yes")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UpdateCustomerProfilesV2(context.Background()).MultipleCustomerProfileIntegrationRequest(multipleCustomerProfileIntegrationRequest).Silent(silent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateCustomerProfilesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomerProfilesV2`: MultipleCustomerProfileIntegrationResponseV2
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateCustomerProfilesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerProfilesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleCustomerProfileIntegrationRequest** | [**MultipleCustomerProfileIntegrationRequest**](MultipleCustomerProfileIntegrationRequest.md) | body | 
 **silent** | **string** | Possible values: &#x60;yes&#x60; or &#x60;no&#x60;. - &#x60;yes&#x60;: Increases the performance of the API call by returning a 204 response. - &#x60;no&#x60;: Returns a 200 response that contains the updated customer profiles.  | [default to &quot;yes&quot;]

### Return type

[**MultipleCustomerProfileIntegrationResponseV2**](MultipleCustomerProfileIntegrationResponseV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomerSessionV2

> IntegrationStateV2 UpdateCustomerSessionV2(ctx, customerSessionId).IntegrationRequest(integrationRequest).Dry(dry).Now(now).Execute()

Update customer session



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
	customerSessionId := "customerSessionId_example" // string | The `integration ID` of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager's **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint. 
	integrationRequest := *openapiclient.NewIntegrationRequest(*openapiclient.NewNewCustomerSessionV2()) // IntegrationRequest | body
	dry := true // bool | Indicates whether to persist the changes. Changes are ignored when `dry=true`.  When set to `true`: - The endpoint considers **only** the payload that you pass when **closing** the session.   When you do not use the `dry` parameter, the endpoint behaves as a typical PUT endpoint. Each update builds upon the previous ones. - You can use the `evaluableCampaignIds` body property to select specific campaigns to run.  [See the docs](https://docs.talon.one/docs/dev/integration-api/dry-requests).  (optional)
	now := time.Now() // time.Time | A timestamp value of a future date that acts as a current date when included in the query.  Use this parameter, for example, to test campaigns that would be evaluated for this customer session in the future (say, [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule)).  **Note:**  - It must be an RFC3339 timestamp string. - It can **only** be a date in the future. - It can **only** be used if the `dry` parameter in the query is set to `true`.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IntegrationAPI.UpdateCustomerSessionV2(context.Background(), customerSessionId).IntegrationRequest(integrationRequest).Dry(dry).Now(now).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IntegrationAPI.UpdateCustomerSessionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomerSessionV2`: IntegrationStateV2
	fmt.Fprintf(os.Stdout, "Response from `IntegrationAPI.UpdateCustomerSessionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerSessionId** | **string** | The &#x60;integration ID&#x60; of the customer session. You set this ID when you create a customer session.  You can see existing customer session integration IDs in the Campaign Manager&#39;s **Sessions** menu, or via the [List Application session](https://docs.talon.one/management-api#operation/getApplicationSessions) endpoint.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerSessionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **integrationRequest** | [**IntegrationRequest**](IntegrationRequest.md) | body | 
 **dry** | **bool** | Indicates whether to persist the changes. Changes are ignored when &#x60;dry&#x3D;true&#x60;.  When set to &#x60;true&#x60;: - The endpoint considers **only** the payload that you pass when **closing** the session.   When you do not use the &#x60;dry&#x60; parameter, the endpoint behaves as a typical PUT endpoint. Each update builds upon the previous ones. - You can use the &#x60;evaluableCampaignIds&#x60; body property to select specific campaigns to run.  [See the docs](https://docs.talon.one/docs/dev/integration-api/dry-requests).  | 
 **now** | **time.Time** | A timestamp value of a future date that acts as a current date when included in the query.  Use this parameter, for example, to test campaigns that would be evaluated for this customer session in the future (say, [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule)).  **Note:**  - It must be an RFC3339 timestamp string. - It can **only** be a date in the future. - It can **only** be used if the &#x60;dry&#x60; parameter in the query is set to &#x60;true&#x60;.  | 

### Return type

[**IntegrationStateV2**](IntegrationStateV2.md)

### Authorization

[api_key_v1](../README.md#api_key_v1)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

