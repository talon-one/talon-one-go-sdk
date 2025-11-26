# GetRulesets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]Ruleset**](Ruleset.md) |  | 

## Methods

### NewGetRulesets200Response

`func NewGetRulesets200Response(totalResultSize int64, data []Ruleset, ) *GetRulesets200Response`

NewGetRulesets200Response instantiates a new GetRulesets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRulesets200ResponseWithDefaults

`func NewGetRulesets200ResponseWithDefaults() *GetRulesets200Response`

NewGetRulesets200ResponseWithDefaults instantiates a new GetRulesets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetRulesets200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetRulesets200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetRulesets200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetRulesets200Response) GetData() []Ruleset`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRulesets200Response) GetDataOk() (*[]Ruleset, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRulesets200Response) SetData(v []Ruleset)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


