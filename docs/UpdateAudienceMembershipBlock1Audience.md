# UpdateAudienceMembershipBlock1Audience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the audience. | 
**Name** | **string** | The display name of the audience. | 
**Integration** | Pointer to **string** | The Talon.One-supported [3rd-party platform](https://docs.talon.one/docs/dev/technology-partners/overview) that this audience was created in.  For example, &#x60;mParticle&#x60;, &#x60;Segment&#x60;, &#x60;Shopify&#x60;, &#x60;Braze&#x60;, or &#x60;Iterable&#x60;.  **Note:** If you do not integrate with any of these platforms, do not use this property.  | [optional] 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration.  **Note:** To create an audience that doesn&#39;t come from a 3rd party platform, do not use this property.  | [optional] 

## Methods

### NewUpdateAudienceMembershipBlock1Audience

`func NewUpdateAudienceMembershipBlock1Audience(id int64, name string, ) *UpdateAudienceMembershipBlock1Audience`

NewUpdateAudienceMembershipBlock1Audience instantiates a new UpdateAudienceMembershipBlock1Audience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAudienceMembershipBlock1AudienceWithDefaults

`func NewUpdateAudienceMembershipBlock1AudienceWithDefaults() *UpdateAudienceMembershipBlock1Audience`

NewUpdateAudienceMembershipBlock1AudienceWithDefaults instantiates a new UpdateAudienceMembershipBlock1Audience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAudienceMembershipBlock1Audience) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAudienceMembershipBlock1Audience) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAudienceMembershipBlock1Audience) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateAudienceMembershipBlock1Audience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAudienceMembershipBlock1Audience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAudienceMembershipBlock1Audience) SetName(v string)`

SetName sets Name field to given value.


### GetIntegration

`func (o *UpdateAudienceMembershipBlock1Audience) GetIntegration() string`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateAudienceMembershipBlock1Audience) GetIntegrationOk() (*string, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateAudienceMembershipBlock1Audience) SetIntegration(v string)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateAudienceMembershipBlock1Audience) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationId

`func (o *UpdateAudienceMembershipBlock1Audience) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UpdateAudienceMembershipBlock1Audience) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UpdateAudienceMembershipBlock1Audience) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *UpdateAudienceMembershipBlock1Audience) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


