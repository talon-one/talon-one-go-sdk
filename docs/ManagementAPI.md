# \ManagementAPI

All URIs are relative to *https://yourbaseurl.talon.one*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CouponCreatedEvent**](ManagementAPI.md#CouponCreatedEvent) | **Post** /CouponCreatedEvent | Coupon Created Event
[**CouponDeletedEvent**](ManagementAPI.md#CouponDeletedEvent) | **Post** /CouponDeletedEvent | Coupon Deleted Event
[**CouponUpdatedEvent**](ManagementAPI.md#CouponUpdatedEvent) | **Post** /CouponUpdatedEvent | Coupon Updated Event
[**LoyaltyPointsChangedEvent**](ManagementAPI.md#LoyaltyPointsChangedEvent) | **Post** /LoyaltyPointsChangedEvent | Changed Loyalty Points Event
[**LoyaltyTierDowngradeEvent**](ManagementAPI.md#LoyaltyTierDowngradeEvent) | **Post** /LoyaltyTierDowngradeEvent | Loyalty Tier Downgrade Event
[**LoyaltyTierUpgradeEvent**](ManagementAPI.md#LoyaltyTierUpgradeEvent) | **Post** /LoyaltyTierUpgradeEvent | Loyalty Tier Upgrade Event



## CouponCreatedEvent

> CouponCreatedEvent(ctx).CouponCreatedEventRequest(couponCreatedEventRequest).Execute()

Coupon Created Event

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
	couponCreatedEventRequest :=  // CouponCreatedEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.CouponCreatedEvent(context.Background()).CouponCreatedEventRequest(couponCreatedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CouponCreatedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponCreatedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponCreatedEventRequest** | [**CouponCreatedEventRequest**](CouponCreatedEventRequest.md) | ... | 

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


## CouponDeletedEvent

> CouponDeletedEvent(ctx).CouponDeletedEventRequest(couponDeletedEventRequest).Execute()

Coupon Deleted Event

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
	couponDeletedEventRequest :=  // CouponDeletedEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.CouponDeletedEvent(context.Background()).CouponDeletedEventRequest(couponDeletedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CouponDeletedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponDeletedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponDeletedEventRequest** | [**CouponDeletedEventRequest**](CouponDeletedEventRequest.md) | ... | 

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


## CouponUpdatedEvent

> CouponUpdatedEvent(ctx).CouponUpdatedEventRequest(couponUpdatedEventRequest).Execute()

Coupon Updated Event

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
	couponUpdatedEventRequest :=  // CouponUpdatedEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.CouponUpdatedEvent(context.Background()).CouponUpdatedEventRequest(couponUpdatedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.CouponUpdatedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCouponUpdatedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **couponUpdatedEventRequest** | [**CouponUpdatedEventRequest**](CouponUpdatedEventRequest.md) | ... | 

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


## LoyaltyPointsChangedEvent

> LoyaltyPointsChangedEvent(ctx).LoyaltyPointsChangedEventRequest(loyaltyPointsChangedEventRequest).Execute()

Changed Loyalty Points Event

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
	loyaltyPointsChangedEventRequest :=  // LoyaltyPointsChangedEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.LoyaltyPointsChangedEvent(context.Background()).LoyaltyPointsChangedEventRequest(loyaltyPointsChangedEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.LoyaltyPointsChangedEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoyaltyPointsChangedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loyaltyPointsChangedEventRequest** | [**LoyaltyPointsChangedEventRequest**](LoyaltyPointsChangedEventRequest.md) | ... | 

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


## LoyaltyTierDowngradeEvent

> LoyaltyTierDowngradeEvent(ctx).LoyaltyTierDowngradeEventRequest(loyaltyTierDowngradeEventRequest).Execute()

Loyalty Tier Downgrade Event

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
	loyaltyTierDowngradeEventRequest :=  // LoyaltyTierDowngradeEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.LoyaltyTierDowngradeEvent(context.Background()).LoyaltyTierDowngradeEventRequest(loyaltyTierDowngradeEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.LoyaltyTierDowngradeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoyaltyTierDowngradeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loyaltyTierDowngradeEventRequest** | [**LoyaltyTierDowngradeEventRequest**](LoyaltyTierDowngradeEventRequest.md) | ... | 

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


## LoyaltyTierUpgradeEvent

> LoyaltyTierUpgradeEvent(ctx).LoyaltyTierUpgradeEventRequest(loyaltyTierUpgradeEventRequest).Execute()

Loyalty Tier Upgrade Event

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
	loyaltyTierUpgradeEventRequest :=  // LoyaltyTierUpgradeEventRequest | ...

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ManagementAPI.LoyaltyTierUpgradeEvent(context.Background()).LoyaltyTierUpgradeEventRequest(loyaltyTierUpgradeEventRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagementAPI.LoyaltyTierUpgradeEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoyaltyTierUpgradeEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loyaltyTierUpgradeEventRequest** | [**LoyaltyTierUpgradeEventRequest**](LoyaltyTierUpgradeEventRequest.md) | ... | 

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

