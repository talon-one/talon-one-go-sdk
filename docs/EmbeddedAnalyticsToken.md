# EmbeddedAnalyticsToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | JWT for embedded analytics. | 
**Hostname** | **string** | Base URL of the embedded analytics instance. | 

## Methods

### NewEmbeddedAnalyticsToken

`func NewEmbeddedAnalyticsToken(token string, hostname string, ) *EmbeddedAnalyticsToken`

NewEmbeddedAnalyticsToken instantiates a new EmbeddedAnalyticsToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedAnalyticsTokenWithDefaults

`func NewEmbeddedAnalyticsTokenWithDefaults() *EmbeddedAnalyticsToken`

NewEmbeddedAnalyticsTokenWithDefaults instantiates a new EmbeddedAnalyticsToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *EmbeddedAnalyticsToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EmbeddedAnalyticsToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EmbeddedAnalyticsToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetHostname

`func (o *EmbeddedAnalyticsToken) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *EmbeddedAnalyticsToken) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *EmbeddedAnalyticsToken) SetHostname(v string)`

SetHostname sets Hostname field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


