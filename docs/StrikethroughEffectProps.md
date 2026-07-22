# StrikethroughEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The effect name. | 
**Value** | **interface{}** |  | 
**ExcludedFromPriceHistory** | Pointer to **bool** | When set to &#x60;true&#x60;, the applied discount is excluded from the item&#39;s price history. | [optional] 
**EffectId** | **int64** | ID of the effect. | 
**Payload** | **map[string]interface{}** | The JSON payload of the custom effect. | 

## Methods

### NewStrikethroughEffectProps

`func NewStrikethroughEffectProps(name string, value interface{}, effectId int64, payload map[string]interface{}, ) *StrikethroughEffectProps`

NewStrikethroughEffectProps instantiates a new StrikethroughEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughEffectPropsWithDefaults

`func NewStrikethroughEffectPropsWithDefaults() *StrikethroughEffectProps`

NewStrikethroughEffectPropsWithDefaults instantiates a new StrikethroughEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StrikethroughEffectProps) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StrikethroughEffectProps) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StrikethroughEffectProps) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *StrikethroughEffectProps) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StrikethroughEffectProps) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StrikethroughEffectProps) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *StrikethroughEffectProps) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *StrikethroughEffectProps) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetExcludedFromPriceHistory

`func (o *StrikethroughEffectProps) GetExcludedFromPriceHistory() bool`

GetExcludedFromPriceHistory returns the ExcludedFromPriceHistory field if non-nil, zero value otherwise.

### GetExcludedFromPriceHistoryOk

`func (o *StrikethroughEffectProps) GetExcludedFromPriceHistoryOk() (*bool, bool)`

GetExcludedFromPriceHistoryOk returns a tuple with the ExcludedFromPriceHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFromPriceHistory

`func (o *StrikethroughEffectProps) SetExcludedFromPriceHistory(v bool)`

SetExcludedFromPriceHistory sets ExcludedFromPriceHistory field to given value.

### HasExcludedFromPriceHistory

`func (o *StrikethroughEffectProps) HasExcludedFromPriceHistory() bool`

HasExcludedFromPriceHistory returns a boolean if a field has been set.

### GetEffectId

`func (o *StrikethroughEffectProps) GetEffectId() int64`

GetEffectId returns the EffectId field if non-nil, zero value otherwise.

### GetEffectIdOk

`func (o *StrikethroughEffectProps) GetEffectIdOk() (*int64, bool)`

GetEffectIdOk returns a tuple with the EffectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectId

`func (o *StrikethroughEffectProps) SetEffectId(v int64)`

SetEffectId sets EffectId field to given value.


### GetPayload

`func (o *StrikethroughEffectProps) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StrikethroughEffectProps) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *StrikethroughEffectProps) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


