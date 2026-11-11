# MCPOAuthStaticClient

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

## Methods

### NewMCPOAuthStaticClient

`func NewMCPOAuthStaticClient(clientId string, clientName string, redirectUris []string, createdAt time.Time, id int64, createdBy int64, ) *MCPOAuthStaticClient`

NewMCPOAuthStaticClient instantiates a new MCPOAuthStaticClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthStaticClientWithDefaults

`func NewMCPOAuthStaticClientWithDefaults() *MCPOAuthStaticClient`

NewMCPOAuthStaticClientWithDefaults instantiates a new MCPOAuthStaticClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *MCPOAuthStaticClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MCPOAuthStaticClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MCPOAuthStaticClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientName

`func (o *MCPOAuthStaticClient) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MCPOAuthStaticClient) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MCPOAuthStaticClient) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetRedirectUris

`func (o *MCPOAuthStaticClient) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *MCPOAuthStaticClient) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *MCPOAuthStaticClient) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetCreatedAt

`func (o *MCPOAuthStaticClient) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MCPOAuthStaticClient) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MCPOAuthStaticClient) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *MCPOAuthStaticClient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPOAuthStaticClient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPOAuthStaticClient) SetId(v int64)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *MCPOAuthStaticClient) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MCPOAuthStaticClient) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MCPOAuthStaticClient) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetExpiryDate

`func (o *MCPOAuthStaticClient) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MCPOAuthStaticClient) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MCPOAuthStaticClient) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MCPOAuthStaticClient) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


