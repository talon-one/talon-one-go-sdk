# LabelTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Audience** | [**AudienceReference**](AudienceReference.md) |  | 

## Methods

### NewLabelTarget

`func NewLabelTarget(type_ string, audience AudienceReference, ) *LabelTarget`

NewLabelTarget instantiates a new LabelTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelTargetWithDefaults

`func NewLabelTargetWithDefaults() *LabelTarget`

NewLabelTargetWithDefaults instantiates a new LabelTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LabelTarget) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelTarget) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelTarget) SetType(v string)`

SetType sets Type field to given value.


### GetAudience

`func (o *LabelTarget) GetAudience() AudienceReference`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *LabelTarget) GetAudienceOk() (*AudienceReference, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *LabelTarget) SetAudience(v AudienceReference)`

SetAudience sets Audience field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


