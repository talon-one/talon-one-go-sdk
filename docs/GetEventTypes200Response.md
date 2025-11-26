# GetEventTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]EventType**](EventType.md) |  | 

## Methods

### NewGetEventTypes200Response

`func NewGetEventTypes200Response(totalResultSize int64, data []EventType, ) *GetEventTypes200Response`

NewGetEventTypes200Response instantiates a new GetEventTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEventTypes200ResponseWithDefaults

`func NewGetEventTypes200ResponseWithDefaults() *GetEventTypes200Response`

NewGetEventTypes200ResponseWithDefaults instantiates a new GetEventTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetEventTypes200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetEventTypes200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetEventTypes200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetEventTypes200Response) GetData() []EventType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetEventTypes200Response) GetDataOk() (*[]EventType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetEventTypes200Response) SetData(v []EventType)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


