# GetCustomerActivityReportsWithoutTotalCount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]CustomerActivityReport**](CustomerActivityReport.md) |  | 

## Methods

### NewGetCustomerActivityReportsWithoutTotalCount200Response

`func NewGetCustomerActivityReportsWithoutTotalCount200Response(hasMore bool, data []CustomerActivityReport, ) *GetCustomerActivityReportsWithoutTotalCount200Response`

NewGetCustomerActivityReportsWithoutTotalCount200Response instantiates a new GetCustomerActivityReportsWithoutTotalCount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerActivityReportsWithoutTotalCount200ResponseWithDefaults

`func NewGetCustomerActivityReportsWithoutTotalCount200ResponseWithDefaults() *GetCustomerActivityReportsWithoutTotalCount200Response`

NewGetCustomerActivityReportsWithoutTotalCount200ResponseWithDefaults instantiates a new GetCustomerActivityReportsWithoutTotalCount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) GetData() []CustomerActivityReport`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) GetDataOk() (*[]CustomerActivityReport, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerActivityReportsWithoutTotalCount200Response) SetData(v []CustomerActivityReport)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


