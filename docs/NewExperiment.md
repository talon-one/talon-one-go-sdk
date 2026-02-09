# NewExperiment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsVariantAssignmentExternal** | **bool** | The source of the assignment. - false - The assignment to the variant is handled internally by the Talon.Oneandled internally by the Talon.One. - true - The assignment to the variant handled externally.  | 
**Activated** | Pointer to **time.Time** | The date and time the experiment was activated.  | [optional] 
**State** | Pointer to **string** | A disabled experiment is not evaluated for rules or coupons.  | [optional] [default to "disabled"]
**Campaign** | [**NewCampaign**](NewCampaign.md) |  | 

## Methods

### NewNewExperiment

`func NewNewExperiment(isVariantAssignmentExternal bool, campaign NewCampaign, ) *NewExperiment`

NewNewExperiment instantiates a new NewExperiment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExperimentWithDefaults

`func NewNewExperimentWithDefaults() *NewExperiment`

NewNewExperimentWithDefaults instantiates a new NewExperiment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsVariantAssignmentExternal

`func (o *NewExperiment) GetIsVariantAssignmentExternal() bool`

GetIsVariantAssignmentExternal returns the IsVariantAssignmentExternal field if non-nil, zero value otherwise.

### GetIsVariantAssignmentExternalOk

`func (o *NewExperiment) GetIsVariantAssignmentExternalOk() (*bool, bool)`

GetIsVariantAssignmentExternalOk returns a tuple with the IsVariantAssignmentExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVariantAssignmentExternal

`func (o *NewExperiment) SetIsVariantAssignmentExternal(v bool)`

SetIsVariantAssignmentExternal sets IsVariantAssignmentExternal field to given value.


### GetActivated

`func (o *NewExperiment) GetActivated() time.Time`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *NewExperiment) GetActivatedOk() (*time.Time, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *NewExperiment) SetActivated(v time.Time)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *NewExperiment) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetState

`func (o *NewExperiment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NewExperiment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NewExperiment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NewExperiment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCampaign

`func (o *NewExperiment) GetCampaign() NewCampaign`

GetCampaign returns the Campaign field if non-nil, zero value otherwise.

### GetCampaignOk

`func (o *NewExperiment) GetCampaignOk() (*NewCampaign, bool)`

GetCampaignOk returns a tuple with the Campaign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaign

`func (o *NewExperiment) SetCampaign(v NewCampaign)`

SetCampaign sets Campaign field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


