# LoyaltyTierDowngradeEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**BatchedAt** | Pointer to **time.Time** | Timestamp when the batch was created. | [optional] 
**EventType** | **string** |  | 
**Data** | **[]interface{}** |  | 

## Methods

### NewLoyaltyTierDowngradeEventRequest

`func NewLoyaltyTierDowngradeEventRequest(totalResultSize int64, eventType string, data []interface{}, ) *LoyaltyTierDowngradeEventRequest`

NewLoyaltyTierDowngradeEventRequest instantiates a new LoyaltyTierDowngradeEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyTierDowngradeEventRequestWithDefaults

`func NewLoyaltyTierDowngradeEventRequestWithDefaults() *LoyaltyTierDowngradeEventRequest`

NewLoyaltyTierDowngradeEventRequestWithDefaults instantiates a new LoyaltyTierDowngradeEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *LoyaltyTierDowngradeEventRequest) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *LoyaltyTierDowngradeEventRequest) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *LoyaltyTierDowngradeEventRequest) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetBatchedAt

`func (o *LoyaltyTierDowngradeEventRequest) GetBatchedAt() time.Time`

GetBatchedAt returns the BatchedAt field if non-nil, zero value otherwise.

### GetBatchedAtOk

`func (o *LoyaltyTierDowngradeEventRequest) GetBatchedAtOk() (*time.Time, bool)`

GetBatchedAtOk returns a tuple with the BatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedAt

`func (o *LoyaltyTierDowngradeEventRequest) SetBatchedAt(v time.Time)`

SetBatchedAt sets BatchedAt field to given value.

### HasBatchedAt

`func (o *LoyaltyTierDowngradeEventRequest) HasBatchedAt() bool`

HasBatchedAt returns a boolean if a field has been set.

### GetEventType

`func (o *LoyaltyTierDowngradeEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *LoyaltyTierDowngradeEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *LoyaltyTierDowngradeEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetData

`func (o *LoyaltyTierDowngradeEventRequest) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoyaltyTierDowngradeEventRequest) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoyaltyTierDowngradeEventRequest) SetData(v []interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


