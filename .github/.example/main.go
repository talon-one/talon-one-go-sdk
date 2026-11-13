package main

import (
	"bytes"
	"context"
	"fmt"
	"os"

	talon "github.com/talon-one/talon-one-go-sdk/v26"
)

func main() {
	configuration := talon.NewConfiguration()
	// Set API base path
	configuration.Servers = talon.ServerConfigurations{
		{
			// Notice that there is no trailing '/'
			URL:         "http://localhost:9000",
			Description: "Talon.One's API base URL",
		},
	}
	// If you wish to inject a custom implementation of HTTPClient
	// configuration.HTTPClient = &customHTTPClient

	integrationClient := talon.NewAPIClient(configuration)

	// Create integration authentication context using api key
	integrationAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"api_key_v1": {
			Prefix: "ApiKey-v1",
			Key:    os.Getenv("TALON_API_KEY"),
		},
	})

	// Instantiating a NewCustomerSessionV2 struct
	newCustomerSession := talon.NewCustomerSessionV2{
		// You can use either struct literals
		ProfileId:   talon.PtrString("DEADBEEF"),
		CouponCodes: []string{"Cool-Stuff!"},
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
	// 	"customerSession",
	// 	"customerProfile",
	// 	"loyalty",
	// })

	// Create/update a customer session using `UpdateCustomerSessionV2` function
	integrationState, _, err := integrationClient.IntegrationAPI.
		UpdateCustomerSessionV2(integrationAuthContext, "deetdoot_2").
		IntegrationRequest(integrationRequest).
		Execute()
	if err != nil {
		fmt.Printf("ERROR while calling UpdateCustomerSessionV2: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%#v\n", integrationState)

	// Parsing the returned effects list, please consult https://developers.talon.one/Integration-API/handling-effects-v2 for the full list of effects and their corresponding properties
	for _, effect := range integrationState.GetEffects() {
		switch {
		case effect.EffectSetDiscount != nil:
			effectProps := effect.EffectSetDiscount.GetProps()

			// Access the specific effect's properties
			fmt.Printf("Set a discount '%s' of %2.3f\n", effectProps.GetName(), effectProps.GetValue())
		case effect.EffectAcceptCoupon != nil:
			effectProps := effect.EffectAcceptCoupon.GetProps()
			fmt.Printf("Accepted coupon '%s'\n", effectProps.GetValue())
		default:
			fmt.Printf("Unhandled effect: %#v\n", effect.GetActualInstance())
		}
	}

	// Management API: the same client serves both APIs. Management endpoints accept a user
	// token with the "Bearer" prefix, or a Management API key with the "ManagementKey-v1" prefix.
	managementAuthContext := context.WithValue(context.Background(), talon.ContextAPIKeys, map[string]talon.APIKey{
		"api_key_v1": {
			Prefix: "Bearer",
			Key:    os.Getenv("TALON_USER_TOKEN"),
		},
	})

	collection, _, err := integrationClient.ManagementAPI.
		CreateAccountCollection(managementAuthContext).
		NewCollection(talon.NewCollection{Name: "SDK example collection"}).
		Execute()
	if err != nil {
		fmt.Printf("ERROR while calling CreateAccountCollection: %s\n", err)
		os.Exit(1)
	}

	// Upload setters accept any io.Reader, so the CSV can come straight from memory.
	// talon.NamedReader sets the file name; an *os.File can be passed directly instead.
	items := []byte("item\nSKU-1\nSKU-2\n")
	_, _, err = integrationClient.ManagementAPI.
		ImportAccountCollection(managementAuthContext, collection.Id).
		UpFile(talon.NamedReader(bytes.NewReader(items), "items.csv")).
		Execute()
	if err != nil {
		fmt.Printf("ERROR while calling ImportAccountCollection: %s\n", err)
		os.Exit(1)
	}

	collectionItems, _, err := integrationClient.ManagementAPI.
		GetCollectionItems(managementAuthContext, collection.Id).
		Execute()
	if err != nil {
		fmt.Printf("ERROR while calling GetCollectionItems: %s\n", err)
		os.Exit(1)
	}
	if len(collectionItems.Data) != 2 {
		fmt.Printf("ERROR: expected 2 collection items after import, got %d\n", len(collectionItems.Data))
		os.Exit(1)
	}
	fmt.Printf("Imported %d collection items from memory\n", len(collectionItems.Data))
}
