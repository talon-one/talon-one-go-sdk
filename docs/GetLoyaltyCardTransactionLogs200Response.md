# GetLoyaltyCardTransactionLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** | true means there is more data in the source collection to request.. | 
**Data** | [**[]CardLedgerTransactionLogEntry**](CardLedgerTransactionLogEntry.md) | List of loyalty card transaction logs. | 

## Methods

### NewGetLoyaltyCardTransactionLogs200Response

`func NewGetLoyaltyCardTransactionLogs200Response(hasMore bool, data []CardLedgerTransactionLogEntry, ) *GetLoyaltyCardTransactionLogs200Response`

NewGetLoyaltyCardTransactionLogs200Response instantiates a new GetLoyaltyCardTransactionLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoyaltyCardTransactionLogs200ResponseWithDefaults

`func NewGetLoyaltyCardTransactionLogs200ResponseWithDefaults() *GetLoyaltyCardTransactionLogs200Response`

NewGetLoyaltyCardTransactionLogs200ResponseWithDefaults instantiates a new GetLoyaltyCardTransactionLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLoyaltyCardTransactionLogs200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLoyaltyCardTransactionLogs200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLoyaltyCardTransactionLogs200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetLoyaltyCardTransactionLogs200Response) GetData() []CardLedgerTransactionLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetLoyaltyCardTransactionLogs200Response) GetDataOk() (*[]CardLedgerTransactionLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetLoyaltyCardTransactionLogs200Response) SetData(v []CardLedgerTransactionLogEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


