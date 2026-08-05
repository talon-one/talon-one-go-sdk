# CheckAchievementBlock1Achievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the achievement. | 
**Title** | **string** | The display name for the achievement in the Campaign Manager. | 
**Name** | **string** | The internal name of the achievement used in API requests. | 
**Target** | **float32** | The required number of actions or the transactional milestone to complete the achievement. | 

## Methods

### NewCheckAchievementBlock1Achievement

`func NewCheckAchievementBlock1Achievement(id int64, title string, name string, target float32, ) *CheckAchievementBlock1Achievement`

NewCheckAchievementBlock1Achievement instantiates a new CheckAchievementBlock1Achievement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAchievementBlock1AchievementWithDefaults

`func NewCheckAchievementBlock1AchievementWithDefaults() *CheckAchievementBlock1Achievement`

NewCheckAchievementBlock1AchievementWithDefaults instantiates a new CheckAchievementBlock1Achievement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAchievementBlock1Achievement) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAchievementBlock1Achievement) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAchievementBlock1Achievement) SetId(v int64)`

SetId sets Id field to given value.


### GetTitle

`func (o *CheckAchievementBlock1Achievement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CheckAchievementBlock1Achievement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CheckAchievementBlock1Achievement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetName

`func (o *CheckAchievementBlock1Achievement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckAchievementBlock1Achievement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckAchievementBlock1Achievement) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *CheckAchievementBlock1Achievement) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *CheckAchievementBlock1Achievement) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *CheckAchievementBlock1Achievement) SetTarget(v float32)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


