# WebhookAuthenticationBaseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the webhook authentication. | [optional] 
**Type** | Pointer to **interface{}** |  | [optional] 
**Data** | Pointer to [**WebhookAuthenticationDataBasic**](WebhookAuthenticationDataBasic.md) |  | [optional] 

## Methods

### NewWebhookAuthenticationBaseOneOf

`func NewWebhookAuthenticationBaseOneOf() *WebhookAuthenticationBaseOneOf`

NewWebhookAuthenticationBaseOneOf instantiates a new WebhookAuthenticationBaseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBaseOneOfWithDefaults

`func NewWebhookAuthenticationBaseOneOfWithDefaults() *WebhookAuthenticationBaseOneOf`

NewWebhookAuthenticationBaseOneOfWithDefaults instantiates a new WebhookAuthenticationBaseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBaseOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBaseOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBaseOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookAuthenticationBaseOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *WebhookAuthenticationBaseOneOf) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBaseOneOf) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBaseOneOf) SetType(v interface{})`

SetType sets Type field to given value.

### HasType

`func (o *WebhookAuthenticationBaseOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *WebhookAuthenticationBaseOneOf) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WebhookAuthenticationBaseOneOf) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *WebhookAuthenticationBaseOneOf) GetData() WebhookAuthenticationDataBasic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBaseOneOf) GetDataOk() (*WebhookAuthenticationDataBasic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBaseOneOf) SetData(v WebhookAuthenticationDataBasic)`

SetData sets Data field to given value.

### HasData

`func (o *WebhookAuthenticationBaseOneOf) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


