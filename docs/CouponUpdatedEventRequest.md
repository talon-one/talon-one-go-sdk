# CouponUpdatedEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**BatchedAt** | Pointer to **time.Time** | Timestamp when the batch was created. | [optional] 
**EventType** | **string** |  | 
**Data** | **[]interface{}** |  | 

## Methods

### NewCouponUpdatedEventRequest

`func NewCouponUpdatedEventRequest(totalResultSize int64, eventType string, data []interface{}, ) *CouponUpdatedEventRequest`

NewCouponUpdatedEventRequest instantiates a new CouponUpdatedEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponUpdatedEventRequestWithDefaults

`func NewCouponUpdatedEventRequestWithDefaults() *CouponUpdatedEventRequest`

NewCouponUpdatedEventRequestWithDefaults instantiates a new CouponUpdatedEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *CouponUpdatedEventRequest) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *CouponUpdatedEventRequest) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *CouponUpdatedEventRequest) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetBatchedAt

`func (o *CouponUpdatedEventRequest) GetBatchedAt() time.Time`

GetBatchedAt returns the BatchedAt field if non-nil, zero value otherwise.

### GetBatchedAtOk

`func (o *CouponUpdatedEventRequest) GetBatchedAtOk() (*time.Time, bool)`

GetBatchedAtOk returns a tuple with the BatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedAt

`func (o *CouponUpdatedEventRequest) SetBatchedAt(v time.Time)`

SetBatchedAt sets BatchedAt field to given value.

### HasBatchedAt

`func (o *CouponUpdatedEventRequest) HasBatchedAt() bool`

HasBatchedAt returns a boolean if a field has been set.

### GetEventType

`func (o *CouponUpdatedEventRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CouponUpdatedEventRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CouponUpdatedEventRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetData

`func (o *CouponUpdatedEventRequest) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CouponUpdatedEventRequest) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CouponUpdatedEventRequest) SetData(v []interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


