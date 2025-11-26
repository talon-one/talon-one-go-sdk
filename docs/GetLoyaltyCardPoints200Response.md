# GetLoyaltyCardPoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]CardLedgerPointsEntryIntegrationAPI**](CardLedgerPointsEntryIntegrationAPI.md) |  | 

## Methods

### NewGetLoyaltyCardPoints200Response

`func NewGetLoyaltyCardPoints200Response(hasMore bool, data []CardLedgerPointsEntryIntegrationAPI, ) *GetLoyaltyCardPoints200Response`

NewGetLoyaltyCardPoints200Response instantiates a new GetLoyaltyCardPoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoyaltyCardPoints200ResponseWithDefaults

`func NewGetLoyaltyCardPoints200ResponseWithDefaults() *GetLoyaltyCardPoints200Response`

NewGetLoyaltyCardPoints200ResponseWithDefaults instantiates a new GetLoyaltyCardPoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLoyaltyCardPoints200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLoyaltyCardPoints200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLoyaltyCardPoints200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetLoyaltyCardPoints200Response) GetData() []CardLedgerPointsEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLoyaltyCardPoints200Response) GetDataOk() (*[]CardLedgerPointsEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLoyaltyCardPoints200Response) SetData(v []CardLedgerPointsEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


