# CheckTierBlock1Tier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the tier. | 
**Name** | **string** | The display name of the tier. | 
**MinPoints** | **float32** | The minimum amount of points required to enter the tier. | 
**UpperLimit** | Pointer to **float32** |  | [optional] 

## Methods

### NewCheckTierBlock1Tier

`func NewCheckTierBlock1Tier(id int64, name string, minPoints float32, ) *CheckTierBlock1Tier`

NewCheckTierBlock1Tier instantiates a new CheckTierBlock1Tier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckTierBlock1TierWithDefaults

`func NewCheckTierBlock1TierWithDefaults() *CheckTierBlock1Tier`

NewCheckTierBlock1TierWithDefaults instantiates a new CheckTierBlock1Tier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckTierBlock1Tier) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckTierBlock1Tier) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckTierBlock1Tier) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *CheckTierBlock1Tier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckTierBlock1Tier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckTierBlock1Tier) SetName(v string)`

SetName sets Name field to given value.


### GetMinPoints

`func (o *CheckTierBlock1Tier) GetMinPoints() float32`

GetMinPoints returns the MinPoints field if non-nil, zero value otherwise.

### GetMinPointsOk

`func (o *CheckTierBlock1Tier) GetMinPointsOk() (*float32, bool)`

GetMinPointsOk returns a tuple with the MinPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPoints

`func (o *CheckTierBlock1Tier) SetMinPoints(v float32)`

SetMinPoints sets MinPoints field to given value.


### GetUpperLimit

`func (o *CheckTierBlock1Tier) GetUpperLimit() float32`

GetUpperLimit returns the UpperLimit field if non-nil, zero value otherwise.

### GetUpperLimitOk

`func (o *CheckTierBlock1Tier) GetUpperLimitOk() (*float32, bool)`

GetUpperLimitOk returns a tuple with the UpperLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperLimit

`func (o *CheckTierBlock1Tier) SetUpperLimit(v float32)`

SetUpperLimit sets UpperLimit field to given value.

### HasUpperLimit

`func (o *CheckTierBlock1Tier) HasUpperLimit() bool`

HasUpperLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


