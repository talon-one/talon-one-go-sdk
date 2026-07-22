# WebhookAuthenticationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the webhook authentication. | 
**Type** | **interface{}** |  | 
**Data** | [**WebhookAuthenticationDataCustom**](WebhookAuthenticationDataCustom.md) |  | 

## Methods

### NewWebhookAuthenticationBase

`func NewWebhookAuthenticationBase(name string, type_ interface{}, data WebhookAuthenticationDataCustom, ) *WebhookAuthenticationBase`

NewWebhookAuthenticationBase instantiates a new WebhookAuthenticationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBaseWithDefaults

`func NewWebhookAuthenticationBaseWithDefaults() *WebhookAuthenticationBase`

NewWebhookAuthenticationBaseWithDefaults instantiates a new WebhookAuthenticationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBase) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthenticationBase) GetType() interface{}`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBase) GetTypeOk() (*interface{}, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBase) SetType(v interface{})`

SetType sets Type field to given value.


### SetTypeNil

`func (o *WebhookAuthenticationBase) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *WebhookAuthenticationBase) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetData

`func (o *WebhookAuthenticationBase) GetData() WebhookAuthenticationDataCustom`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBase) GetDataOk() (*WebhookAuthenticationDataCustom, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBase) SetData(v WebhookAuthenticationDataCustom)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


