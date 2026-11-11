# WebhookAuthenticationBasic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the webhook authentication. | 
**Type** | **string** | A webhook authentication discriminator of type &#x60;basic&#x60;. | 
**Data** | [**WebhookAuthenticationDataBasic**](WebhookAuthenticationDataBasic.md) | The credentials of the webhook authentication. | 
**CreatedBy** | **string** | The name of the user who created the webhook authentication. | 
**ModifiedBy** | **string** | The name of the user who last modified the webhook authentication. | 
**Webhooks** | [**[]WebhookAuthenticationWebhookRef**](WebhookAuthenticationWebhookRef.md) |  | 
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**Modified** | **time.Time** | The time this entity was last modified. | 

## Methods

### NewWebhookAuthenticationBasic

`func NewWebhookAuthenticationBasic(name string, type_ string, data WebhookAuthenticationDataBasic, createdBy string, modifiedBy string, webhooks []WebhookAuthenticationWebhookRef, id int64, created time.Time, modified time.Time, ) *WebhookAuthenticationBasic`

NewWebhookAuthenticationBasic instantiates a new WebhookAuthenticationBasic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookAuthenticationBasicWithDefaults

`func NewWebhookAuthenticationBasicWithDefaults() *WebhookAuthenticationBasic`

NewWebhookAuthenticationBasicWithDefaults instantiates a new WebhookAuthenticationBasic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookAuthenticationBasic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookAuthenticationBasic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookAuthenticationBasic) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WebhookAuthenticationBasic) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookAuthenticationBasic) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookAuthenticationBasic) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *WebhookAuthenticationBasic) GetData() WebhookAuthenticationDataBasic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebhookAuthenticationBasic) GetDataOk() (*WebhookAuthenticationDataBasic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebhookAuthenticationBasic) SetData(v WebhookAuthenticationDataBasic)`

SetData sets Data field to given value.


### GetCreatedBy

`func (o *WebhookAuthenticationBasic) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WebhookAuthenticationBasic) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WebhookAuthenticationBasic) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedBy

`func (o *WebhookAuthenticationBasic) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *WebhookAuthenticationBasic) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *WebhookAuthenticationBasic) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetWebhooks

`func (o *WebhookAuthenticationBasic) GetWebhooks() []WebhookAuthenticationWebhookRef`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *WebhookAuthenticationBasic) GetWebhooksOk() (*[]WebhookAuthenticationWebhookRef, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *WebhookAuthenticationBasic) SetWebhooks(v []WebhookAuthenticationWebhookRef)`

SetWebhooks sets Webhooks field to given value.


### GetId

`func (o *WebhookAuthenticationBasic) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookAuthenticationBasic) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookAuthenticationBasic) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *WebhookAuthenticationBasic) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *WebhookAuthenticationBasic) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *WebhookAuthenticationBasic) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *WebhookAuthenticationBasic) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *WebhookAuthenticationBasic) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *WebhookAuthenticationBasic) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


