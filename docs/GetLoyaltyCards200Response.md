# GetLoyaltyCards200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]LoyaltyCard**](LoyaltyCard.md) |  | 

## Methods

### NewGetLoyaltyCards200Response

`func NewGetLoyaltyCards200Response(hasMore bool, data []LoyaltyCard, ) *GetLoyaltyCards200Response`

NewGetLoyaltyCards200Response instantiates a new GetLoyaltyCards200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoyaltyCards200ResponseWithDefaults

`func NewGetLoyaltyCards200ResponseWithDefaults() *GetLoyaltyCards200Response`

NewGetLoyaltyCards200ResponseWithDefaults instantiates a new GetLoyaltyCards200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLoyaltyCards200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLoyaltyCards200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLoyaltyCards200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetLoyaltyCards200Response) GetData() []LoyaltyCard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLoyaltyCards200Response) GetDataOk() (*[]LoyaltyCard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLoyaltyCards200Response) SetData(v []LoyaltyCard)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


