# GetCouponsWithoutTotalCount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]Coupon**](Coupon.md) |  | 

## Methods

### NewGetCouponsWithoutTotalCount200Response

`func NewGetCouponsWithoutTotalCount200Response(hasMore bool, data []Coupon, ) *GetCouponsWithoutTotalCount200Response`

NewGetCouponsWithoutTotalCount200Response instantiates a new GetCouponsWithoutTotalCount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCouponsWithoutTotalCount200ResponseWithDefaults

`func NewGetCouponsWithoutTotalCount200ResponseWithDefaults() *GetCouponsWithoutTotalCount200Response`

NewGetCouponsWithoutTotalCount200ResponseWithDefaults instantiates a new GetCouponsWithoutTotalCount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCouponsWithoutTotalCount200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCouponsWithoutTotalCount200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCouponsWithoutTotalCount200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetCouponsWithoutTotalCount200Response) GetData() []Coupon`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCouponsWithoutTotalCount200Response) GetDataOk() (*[]Coupon, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCouponsWithoutTotalCount200Response) SetData(v []Coupon)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


