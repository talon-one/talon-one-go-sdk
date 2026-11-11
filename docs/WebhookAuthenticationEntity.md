# WebhookAuthenticationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**Modified** | **time.Time** | The time this entity was last modified. | 
**CreatedBy** | **string** | The name of the user who created the webhook authentication. | 
**ModifiedBy** | **string** | The name of the user who last modified the webhook authentication. | 
**Webhooks** | [**[]WebhookAuthenticationWebhookRef**](WebhookAuthenticationWebhookRef.md) |  | 

## Methods

### NewWebhookAuthenticationEntity

`func NewWebhookAuthenticationEntity(id int64, created time.Time, modified time.Time, createdBy string, modifiedBy string, webhooks []WebhookAuthenticationWebhookRef, ) *WebhookAuthenticationEntity`

NewWebhookAuthenticationEntity instantiates a new WebhookAuthenticationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationEntityWithDefaults

`func NewWebhookAuthenticationEntityWithDefaults() *WebhookAuthenticationEntity`

NewWebhookAuthenticationEntityWithDefaults instantiates a new WebhookAuthenticationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookAuthenticationEntity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthenticationEntity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAuthenticationEntity) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *WebhookAuthenticationEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookAuthenticationEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookAuthenticationEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *WebhookAuthenticationEntity) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookAuthenticationEntity) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WebhookAuthenticationEntity) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetCreatedBy

`func (o *WebhookAuthenticationEntity) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookAuthenticationEntity) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhookAuthenticationEntity) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *WebhookAuthenticationEntity) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhookAuthenticationEntity) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *WebhookAuthenticationEntity) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetWebhooks

`func (o *WebhookAuthenticationEntity) GetWebhooks() []WebhookAuthenticationWebhookRef`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookAuthenticationEntity) GetWebhooksOk() (*[]WebhookAuthenticationWebhookRef, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookAuthenticationEntity) SetWebhooks(v []WebhookAuthenticationWebhookRef)`

SetWebhooks sets Webhooks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


