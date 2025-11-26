# GetAdditionalCosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]AccountAdditionalCost**](AccountAdditionalCost.md) |  | 

## Methods

### NewGetAdditionalCosts200Response

`func NewGetAdditionalCosts200Response(totalResultSize int64, data []AccountAdditionalCost, ) *GetAdditionalCosts200Response`

NewGetAdditionalCosts200Response instantiates a new GetAdditionalCosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdditionalCosts200ResponseWithDefaults

`func NewGetAdditionalCosts200ResponseWithDefaults() *GetAdditionalCosts200Response`

NewGetAdditionalCosts200ResponseWithDefaults instantiates a new GetAdditionalCosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetAdditionalCosts200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetAdditionalCosts200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetAdditionalCosts200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetAdditionalCosts200Response) GetData() []AccountAdditionalCost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAdditionalCosts200Response) GetDataOk() (*[]AccountAdditionalCost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAdditionalCosts200Response) SetData(v []AccountAdditionalCost)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


