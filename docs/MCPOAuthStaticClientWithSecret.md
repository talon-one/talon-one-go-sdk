# MCPOAuthStaticClientWithSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Unique identifier for the OAuth2 client. | 
**ClientName** | **string** | Human-readable name for the OAuth2 client. | 
**RedirectUris** | **[]string** | List of allowed redirect URIs for the authorization code flow. | 
**CreatedAt** | **time.Time** | Timestamp of when the client was registered. | 
**Id** | **int64** | Numeric ID of the OAuth client, used to delete it. | 
**CreatedBy** | **int64** | ID of the user who created the client. | 
**ExpiryDate** | Pointer to **time.Time** | The date the client credentials expire. | [optional] 
**ClientSecret** | **string** | The generated client secret. Only returned once on creation. | 

## Methods

### NewMCPOAuthStaticClientWithSecret

`func NewMCPOAuthStaticClientWithSecret(clientId string, clientName string, redirectUris []string, createdAt time.Time, id int64, createdBy int64, clientSecret string, ) *MCPOAuthStaticClientWithSecret`

NewMCPOAuthStaticClientWithSecret instantiates a new MCPOAuthStaticClientWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthStaticClientWithSecretWithDefaults

`func NewMCPOAuthStaticClientWithSecretWithDefaults() *MCPOAuthStaticClientWithSecret`

NewMCPOAuthStaticClientWithSecretWithDefaults instantiates a new MCPOAuthStaticClientWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MCPOAuthStaticClientWithSecret) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MCPOAuthStaticClientWithSecret) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MCPOAuthStaticClientWithSecret) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientName

`func (o *MCPOAuthStaticClientWithSecret) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MCPOAuthStaticClientWithSecret) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MCPOAuthStaticClientWithSecret) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *MCPOAuthStaticClientWithSecret) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *MCPOAuthStaticClientWithSecret) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *MCPOAuthStaticClientWithSecret) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetCreatedAt

`func (o *MCPOAuthStaticClientWithSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MCPOAuthStaticClientWithSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MCPOAuthStaticClientWithSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *MCPOAuthStaticClientWithSecret) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPOAuthStaticClientWithSecret) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPOAuthStaticClientWithSecret) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *MCPOAuthStaticClientWithSecret) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MCPOAuthStaticClientWithSecret) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MCPOAuthStaticClientWithSecret) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetExpiryDate

`func (o *MCPOAuthStaticClientWithSecret) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MCPOAuthStaticClientWithSecret) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MCPOAuthStaticClientWithSecret) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MCPOAuthStaticClientWithSecret) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetClientSecret

`func (o *MCPOAuthStaticClientWithSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MCPOAuthStaticClientWithSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MCPOAuthStaticClientWithSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


