# GetCustomerProfiles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]CustomerProfile**](CustomerProfile.md) |  | 

## Methods

### NewGetCustomerProfiles200Response

`func NewGetCustomerProfiles200Response(hasMore bool, data []CustomerProfile, ) *GetCustomerProfiles200Response`

NewGetCustomerProfiles200Response instantiates a new GetCustomerProfiles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerProfiles200ResponseWithDefaults

`func NewGetCustomerProfiles200ResponseWithDefaults() *GetCustomerProfiles200Response`

NewGetCustomerProfiles200ResponseWithDefaults instantiates a new GetCustomerProfiles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCustomerProfiles200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCustomerProfiles200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCustomerProfiles200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetCustomerProfiles200Response) GetData() []CustomerProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerProfiles200Response) GetDataOk() (*[]CustomerProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerProfiles200Response) SetData(v []CustomerProfile)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


