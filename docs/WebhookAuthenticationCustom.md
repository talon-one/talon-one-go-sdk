# WebhookAuthenticationCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the webhook authentication. | 
**Type** | **string** | A webhook authentication discriminator of type &#x60;custom&#x60;. | 
**Data** | [**WebhookAuthenticationDataCustom**](WebhookAuthenticationDataCustom.md) | The credentials of the webhook authentication. | 
**CreatedBy** | **string** | The name of the user who created the webhook authentication. | 
**ModifiedBy** | **string** | The name of the user who last modified the webhook authentication. | 
**Webhooks** | [**[]WebhookAuthenticationWebhookRef**](WebhookAuthenticationWebhookRef.md) |  | 
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**Modified** | **time.Time** | The time this entity was last modified. | 

## Methods

### NewWebhookAuthenticationCustom

`func NewWebhookAuthenticationCustom(name string, type_ string, data WebhookAuthenticationDataCustom, createdBy string, modifiedBy string, webhooks []WebhookAuthenticationWebhookRef, id int64, created time.Time, modified time.Time, ) *WebhookAuthenticationCustom`

NewWebhookAuthenticationCustom instantiates a new WebhookAuthenticationCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationCustomWithDefaults

`func NewWebhookAuthenticationCustomWithDefaults() *WebhookAuthenticationCustom`

NewWebhookAuthenticationCustomWithDefaults instantiates a new WebhookAuthenticationCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationCustom) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthenticationCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationCustom) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookAuthenticationCustom) GetData() WebhookAuthenticationDataCustom`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationCustom) GetDataOk() (*WebhookAuthenticationDataCustom, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationCustom) SetData(v WebhookAuthenticationDataCustom)`

SetData sets Data field to given value.


### GetCreatedBy

`func (o *WebhookAuthenticationCustom) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookAuthenticationCustom) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhookAuthenticationCustom) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *WebhookAuthenticationCustom) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhookAuthenticationCustom) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *WebhookAuthenticationCustom) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetWebhooks

`func (o *WebhookAuthenticationCustom) GetWebhooks() []WebhookAuthenticationWebhookRef`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookAuthenticationCustom) GetWebhooksOk() (*[]WebhookAuthenticationWebhookRef, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookAuthenticationCustom) SetWebhooks(v []WebhookAuthenticationWebhookRef)`

SetWebhooks sets Webhooks field to given value.


### GetId

`func (o *WebhookAuthenticationCustom) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthenticationCustom) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAuthenticationCustom) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *WebhookAuthenticationCustom) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookAuthenticationCustom) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookAuthenticationCustom) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *WebhookAuthenticationCustom) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookAuthenticationCustom) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WebhookAuthenticationCustom) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


