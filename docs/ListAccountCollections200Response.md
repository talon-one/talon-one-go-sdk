# ListAccountCollections200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | [**[]CollectionWithoutPayload**](CollectionWithoutPayload.md) |  | 

## Methods

### NewListAccountCollections200Response

`func NewListAccountCollections200Response(data []CollectionWithoutPayload, ) *ListAccountCollections200Response`

NewListAccountCollections200Response instantiates a new ListAccountCollections200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountCollections200ResponseWithDefaults

`func NewListAccountCollections200ResponseWithDefaults() *ListAccountCollections200Response`

NewListAccountCollections200ResponseWithDefaults instantiates a new ListAccountCollections200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListAccountCollections200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListAccountCollections200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListAccountCollections200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListAccountCollections200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *ListAccountCollections200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *ListAccountCollections200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *ListAccountCollections200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *ListAccountCollections200Response) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *ListAccountCollections200Response) GetData() []CollectionWithoutPayload`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAccountCollections200Response) GetDataOk() (*[]CollectionWithoutPayload, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAccountCollections200Response) SetData(v []CollectionWithoutPayload)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


