# NewMCPOAuthStaticClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **string** | Human-readable name for the OAuth2 client. | 
**RedirectUris** | **[]string** | List of allowed redirect URIs for the authorization code flow. At least one URI is required. | 
**ExpiryDate** | Pointer to **time.Time** | The date the client credentials expire. | [optional] 

## Methods

### NewNewMCPOAuthStaticClient

`func NewNewMCPOAuthStaticClient(clientName string, redirectUris []string, ) *NewMCPOAuthStaticClient`

NewNewMCPOAuthStaticClient instantiates a new NewMCPOAuthStaticClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMCPOAuthStaticClientWithDefaults

`func NewNewMCPOAuthStaticClientWithDefaults() *NewMCPOAuthStaticClient`

NewNewMCPOAuthStaticClientWithDefaults instantiates a new NewMCPOAuthStaticClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *NewMCPOAuthStaticClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *NewMCPOAuthStaticClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *NewMCPOAuthStaticClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *NewMCPOAuthStaticClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *NewMCPOAuthStaticClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *NewMCPOAuthStaticClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetExpiryDate

`func (o *NewMCPOAuthStaticClient) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewMCPOAuthStaticClient) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewMCPOAuthStaticClient) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewMCPOAuthStaticClient) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


