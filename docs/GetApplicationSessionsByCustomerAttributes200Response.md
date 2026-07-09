# GetApplicationSessionsByCustomerAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | [**[]ApplicationSession**](ApplicationSession.md) |  | 

## Methods

### NewGetApplicationSessionsByCustomerAttributes200Response

`func NewGetApplicationSessionsByCustomerAttributes200Response(data []ApplicationSession, ) *GetApplicationSessionsByCustomerAttributes200Response`

NewGetApplicationSessionsByCustomerAttributes200Response instantiates a new GetApplicationSessionsByCustomerAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationSessionsByCustomerAttributes200ResponseWithDefaults

`func NewGetApplicationSessionsByCustomerAttributes200ResponseWithDefaults() *GetApplicationSessionsByCustomerAttributes200Response`

NewGetApplicationSessionsByCustomerAttributes200ResponseWithDefaults instantiates a new GetApplicationSessionsByCustomerAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetApplicationSessionsByCustomerAttributes200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetApplicationSessionsByCustomerAttributes200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetApplicationSessionsByCustomerAttributes200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *GetApplicationSessionsByCustomerAttributes200Response) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetData() []ApplicationSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApplicationSessionsByCustomerAttributes200Response) GetDataOk() (*[]ApplicationSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApplicationSessionsByCustomerAttributes200Response) SetData(v []ApplicationSession)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


