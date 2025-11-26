# GetLoyaltyProgramProfileTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]LedgerTransactionLogEntryIntegrationAPI**](LedgerTransactionLogEntryIntegrationAPI.md) |  | 

## Methods

### NewGetLoyaltyProgramProfileTransactions200Response

`func NewGetLoyaltyProgramProfileTransactions200Response(hasMore bool, data []LedgerTransactionLogEntryIntegrationAPI, ) *GetLoyaltyProgramProfileTransactions200Response`

NewGetLoyaltyProgramProfileTransactions200Response instantiates a new GetLoyaltyProgramProfileTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoyaltyProgramProfileTransactions200ResponseWithDefaults

`func NewGetLoyaltyProgramProfileTransactions200ResponseWithDefaults() *GetLoyaltyProgramProfileTransactions200Response`

NewGetLoyaltyProgramProfileTransactions200ResponseWithDefaults instantiates a new GetLoyaltyProgramProfileTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLoyaltyProgramProfileTransactions200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLoyaltyProgramProfileTransactions200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLoyaltyProgramProfileTransactions200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetLoyaltyProgramProfileTransactions200Response) GetData() []LedgerTransactionLogEntryIntegrationAPI`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLoyaltyProgramProfileTransactions200Response) GetDataOk() (*[]LedgerTransactionLogEntryIntegrationAPI, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLoyaltyProgramProfileTransactions200Response) SetData(v []LedgerTransactionLogEntryIntegrationAPI)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


