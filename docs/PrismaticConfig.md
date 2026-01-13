# PrismaticConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrismaticUrl** | **string** | The url used to integrate the Prismatic Marketplace. | 
**AccessToken** | **string** | Access token used to authenticate a user in Talon.One. | 

## Methods

### NewPrismaticConfig

`func NewPrismaticConfig(prismaticUrl string, accessToken string, ) *PrismaticConfig`

NewPrismaticConfig instantiates a new PrismaticConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticConfigWithDefaults

`func NewPrismaticConfigWithDefaults() *PrismaticConfig`

NewPrismaticConfigWithDefaults instantiates a new PrismaticConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrismaticUrl

`func (o *PrismaticConfig) GetPrismaticUrl() string`

GetPrismaticUrl returns the PrismaticUrl field if non-nil, zero value otherwise.

### GetPrismaticUrlOk

`func (o *PrismaticConfig) GetPrismaticUrlOk() (*string, bool)`

GetPrismaticUrlOk returns a tuple with the PrismaticUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrismaticUrl

`func (o *PrismaticConfig) SetPrismaticUrl(v string)`

SetPrismaticUrl sets PrismaticUrl field to given value.


### GetAccessToken

`func (o *PrismaticConfig) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *PrismaticConfig) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *PrismaticConfig) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


