# GetCollectionItems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]CollectionItem**](CollectionItem.md) |  | 

## Methods

### NewGetCollectionItems200Response

`func NewGetCollectionItems200Response(hasMore bool, data []CollectionItem, ) *GetCollectionItems200Response`

NewGetCollectionItems200Response instantiates a new GetCollectionItems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectionItems200ResponseWithDefaults

`func NewGetCollectionItems200ResponseWithDefaults() *GetCollectionItems200Response`

NewGetCollectionItems200ResponseWithDefaults instantiates a new GetCollectionItems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCollectionItems200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCollectionItems200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCollectionItems200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetCollectionItems200Response) GetData() []CollectionItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCollectionItems200Response) GetDataOk() (*[]CollectionItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCollectionItems200Response) SetData(v []CollectionItem)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


