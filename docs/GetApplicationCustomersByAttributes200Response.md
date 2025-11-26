# GetApplicationCustomersByAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**TotalResultSize** | Pointer to **int64** |  | [optional] 
**Data** | [**[]ApplicationCustomer**](ApplicationCustomer.md) |  | 

## Methods

### NewGetApplicationCustomersByAttributes200Response

`func NewGetApplicationCustomersByAttributes200Response(data []ApplicationCustomer, ) *GetApplicationCustomersByAttributes200Response`

NewGetApplicationCustomersByAttributes200Response instantiates a new GetApplicationCustomersByAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetApplicationCustomersByAttributes200ResponseWithDefaults

`func NewGetApplicationCustomersByAttributes200ResponseWithDefaults() *GetApplicationCustomersByAttributes200Response`

NewGetApplicationCustomersByAttributes200ResponseWithDefaults instantiates a new GetApplicationCustomersByAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetApplicationCustomersByAttributes200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetApplicationCustomersByAttributes200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetApplicationCustomersByAttributes200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetApplicationCustomersByAttributes200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetTotalResultSize

`func (o *GetApplicationCustomersByAttributes200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetApplicationCustomersByAttributes200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetApplicationCustomersByAttributes200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.

### HasTotalResultSize

`func (o *GetApplicationCustomersByAttributes200Response) HasTotalResultSize() bool`

HasTotalResultSize returns a boolean if a field has been set.

### GetData

`func (o *GetApplicationCustomersByAttributes200Response) GetData() []ApplicationCustomer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetApplicationCustomersByAttributes200Response) GetDataOk() (*[]ApplicationCustomer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetApplicationCustomersByAttributes200Response) SetData(v []ApplicationCustomer)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


