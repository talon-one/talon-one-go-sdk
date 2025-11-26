# GetLoyaltyProgramProfilePoints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]LedgerPointsEntryIntegrationAPI**](LedgerPointsEntryIntegrationAPI.md) |  | 

## Methods

### NewGetLoyaltyProgramProfilePoints200Response

`func NewGetLoyaltyProgramProfilePoints200Response(hasMore bool, data []LedgerPointsEntryIntegrationAPI, ) *GetLoyaltyProgramProfilePoints200Response`

NewGetLoyaltyProgramProfilePoints200Response instantiates a new GetLoyaltyProgramProfilePoints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoyaltyProgramProfilePoints200ResponseWithDefaults

`func NewGetLoyaltyProgramProfilePoints200ResponseWithDefaults() *GetLoyaltyProgramProfilePoints200Response`

NewGetLoyaltyProgramProfilePoints200ResponseWithDefaults instantiates a new GetLoyaltyProgramProfilePoints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLoyaltyProgramProfilePoints200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLoyaltyProgramProfilePoints200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLoyaltyProgramProfilePoints200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetLoyaltyProgramProfilePoints200Response) GetData() []LedgerPointsEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLoyaltyProgramProfilePoints200Response) GetDataOk() (*[]LedgerPointsEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLoyaltyProgramProfilePoints200Response) SetData(v []LedgerPointsEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


