# BestPriorPriceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skus** | **[]string** | List of product SKUs to check when determining the best prior price. | 
**TimeframeEndDate** | **time.Time** | The end date and time that defines the latest time for retrieving historical SKU prices. | 
**Timeframe** | **string** | The number of days prior to the timeframeEndDate. Only prices within this look back period are considered for the best prior price evaluation. | 
**StrictEndDate** | **bool** | Indicates whether the timeframe includes the start of the current sale. - When &#x60;false&#x60;, the timeframe includes the start date of the current sale. - When &#x60;true&#x60;, the timeframe striclty uses the number of days specified in &#x60;timeframe&#x60;.  | 
**Target** | Pointer to [**BestPriorTarget**](BestPriorTarget.md) |  | [optional] 

## Methods

### NewBestPriorPriceRequest

`func NewBestPriorPriceRequest(skus []string, timeframeEndDate time.Time, timeframe string, strictEndDate bool, ) *BestPriorPriceRequest`

NewBestPriorPriceRequest instantiates a new BestPriorPriceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBestPriorPriceRequestWithDefaults

`func NewBestPriorPriceRequestWithDefaults() *BestPriorPriceRequest`

NewBestPriorPriceRequestWithDefaults instantiates a new BestPriorPriceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkus

`func (o *BestPriorPriceRequest) GetSkus() []string`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *BestPriorPriceRequest) GetSkusOk() (*[]string, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *BestPriorPriceRequest) SetSkus(v []string)`

SetSkus sets Skus field to given value.


### GetTimeframeEndDate

`func (o *BestPriorPriceRequest) GetTimeframeEndDate() time.Time`

GetTimeframeEndDate returns the TimeframeEndDate field if non-nil, zero value otherwise.

### GetTimeframeEndDateOk

`func (o *BestPriorPriceRequest) GetTimeframeEndDateOk() (*time.Time, bool)`

GetTimeframeEndDateOk returns a tuple with the TimeframeEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframeEndDate

`func (o *BestPriorPriceRequest) SetTimeframeEndDate(v time.Time)`

SetTimeframeEndDate sets TimeframeEndDate field to given value.


### GetTimeframe

`func (o *BestPriorPriceRequest) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *BestPriorPriceRequest) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *BestPriorPriceRequest) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.


### GetStrictEndDate

`func (o *BestPriorPriceRequest) GetStrictEndDate() bool`

GetStrictEndDate returns the StrictEndDate field if non-nil, zero value otherwise.

### GetStrictEndDateOk

`func (o *BestPriorPriceRequest) GetStrictEndDateOk() (*bool, bool)`

GetStrictEndDateOk returns a tuple with the StrictEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrictEndDate

`func (o *BestPriorPriceRequest) SetStrictEndDate(v bool)`

SetStrictEndDate sets StrictEndDate field to given value.


### GetTarget

`func (o *BestPriorPriceRequest) GetTarget() BestPriorTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BestPriorPriceRequest) GetTargetOk() (*BestPriorTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BestPriorPriceRequest) SetTarget(v BestPriorTarget)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BestPriorPriceRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


