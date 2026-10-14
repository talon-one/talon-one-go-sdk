# ListGiveawaysPool200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | [**[]GiveawaysPool**](GiveawaysPool.md) |  | 

## Methods

### NewListGiveawaysPool200Response

`func NewListGiveawaysPool200Response(data []GiveawaysPool, ) *ListGiveawaysPool200Response`

NewListGiveawaysPool200Response instantiates a new ListGiveawaysPool200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGiveawaysPool200ResponseWithDefaults

`func NewListGiveawaysPool200ResponseWithDefaults() *ListGiveawaysPool200Response`

NewListGiveawaysPool200ResponseWithDefaults instantiates a new ListGiveawaysPool200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListGiveawaysPool200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListGiveawaysPool200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListGiveawaysPool200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListGiveawaysPool200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *ListGiveawaysPool200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *ListGiveawaysPool200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *ListGiveawaysPool200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *ListGiveawaysPool200Response) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *ListGiveawaysPool200Response) GetData() []GiveawaysPool`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListGiveawaysPool200Response) GetDataOk() (*[]GiveawaysPool, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListGiveawaysPool200Response) SetData(v []GiveawaysPool)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


