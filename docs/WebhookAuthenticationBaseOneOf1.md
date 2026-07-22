# WebhookAuthenticationBaseOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the webhook authentication. | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Data** | Pointer to [**WebhookAuthenticationDataCustom**](WebhookAuthenticationDataCustom.md) |  | [optional] 

## Methods

### NewWebhookAuthenticationBaseOneOf1

`func NewWebhookAuthenticationBaseOneOf1() *WebhookAuthenticationBaseOneOf1`

NewWebhookAuthenticationBaseOneOf1 instantiates a new WebhookAuthenticationBaseOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBaseOneOf1WithDefaults

`func NewWebhookAuthenticationBaseOneOf1WithDefaults() *WebhookAuthenticationBaseOneOf1`

NewWebhookAuthenticationBaseOneOf1WithDefaults instantiates a new WebhookAuthenticationBaseOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBaseOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBaseOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBaseOneOf1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookAuthenticationBaseOneOf1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WebhookAuthenticationBaseOneOf1) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBaseOneOf1) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBaseOneOf1) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *WebhookAuthenticationBaseOneOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WebhookAuthenticationBaseOneOf1) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WebhookAuthenticationBaseOneOf1) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *WebhookAuthenticationBaseOneOf1) GetData() WebhookAuthenticationDataCustom`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBaseOneOf1) GetDataOk() (*WebhookAuthenticationDataCustom, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBaseOneOf1) SetData(v WebhookAuthenticationDataCustom)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookAuthenticationBaseOneOf1) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


