# UpdateAttributeValueBlock1Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of the attribute. Reverts to &#x60;0&#x60; when the attribute is deleted or does not exist. | 
**Entity** | **string** | The entity type that owns the attribute. Reverts to an empty string when the attribute is deleted or does not exist. | 
**Name** | **string** | The attribute name as used in API requests. | 
**Title** | **string** | The human-readable name of the attribute. | 
**Type** | **string** | The data type of the attribute. | 

## Methods

### NewUpdateAttributeValueBlock1Attribute

`func NewUpdateAttributeValueBlock1Attribute(id int64, entity string, name string, title string, type_ string, ) *UpdateAttributeValueBlock1Attribute`

NewUpdateAttributeValueBlock1Attribute instantiates a new UpdateAttributeValueBlock1Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAttributeValueBlock1AttributeWithDefaults

`func NewUpdateAttributeValueBlock1AttributeWithDefaults() *UpdateAttributeValueBlock1Attribute`

NewUpdateAttributeValueBlock1AttributeWithDefaults instantiates a new UpdateAttributeValueBlock1Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAttributeValueBlock1Attribute) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAttributeValueBlock1Attribute) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAttributeValueBlock1Attribute) SetId(v int64)`

SetId sets Id field to given value.


### GetEntity

`func (o *UpdateAttributeValueBlock1Attribute) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *UpdateAttributeValueBlock1Attribute) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *UpdateAttributeValueBlock1Attribute) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetName

`func (o *UpdateAttributeValueBlock1Attribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAttributeValueBlock1Attribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAttributeValueBlock1Attribute) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *UpdateAttributeValueBlock1Attribute) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateAttributeValueBlock1Attribute) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateAttributeValueBlock1Attribute) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *UpdateAttributeValueBlock1Attribute) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateAttributeValueBlock1Attribute) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateAttributeValueBlock1Attribute) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


