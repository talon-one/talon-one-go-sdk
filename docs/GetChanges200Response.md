# GetChanges200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | [**[]Change**](Change.md) |  | 

## Methods

### NewGetChanges200Response

`func NewGetChanges200Response(data []Change, ) *GetChanges200Response`

NewGetChanges200Response instantiates a new GetChanges200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChanges200ResponseWithDefaults

`func NewGetChanges200ResponseWithDefaults() *GetChanges200Response`

NewGetChanges200ResponseWithDefaults instantiates a new GetChanges200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetChanges200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetChanges200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetChanges200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *GetChanges200Response) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetHasMore

`func (o *GetChanges200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetChanges200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetChanges200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetChanges200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *GetChanges200Response) GetData() []Change`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetChanges200Response) GetDataOk() (*[]Change, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetChanges200Response) SetData(v []Change)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


