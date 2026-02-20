# PrismaticEventRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**FlowId** | **int64** |  | 
**EventType** | **string** |  | 
**EventData** | **interface{}** |  | 
**PublishedAt** | **time.Time** |  | 
**ProcessedAt** | Pointer to **time.Time** |  | [optional] 
**ProcessAfter** | **time.Time** |  | 
**Retry** | **int64** |  | 

## Methods

### NewPrismaticEventRecord

`func NewPrismaticEventRecord(id int64, flowId int64, eventType string, eventData interface{}, publishedAt time.Time, processAfter time.Time, retry int64, ) *PrismaticEventRecord`

NewPrismaticEventRecord instantiates a new PrismaticEventRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventRecordWithDefaults

`func NewPrismaticEventRecordWithDefaults() *PrismaticEventRecord`

NewPrismaticEventRecordWithDefaults instantiates a new PrismaticEventRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrismaticEventRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrismaticEventRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrismaticEventRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetFlowId

`func (o *PrismaticEventRecord) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *PrismaticEventRecord) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *PrismaticEventRecord) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.


### GetEventType

`func (o *PrismaticEventRecord) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PrismaticEventRecord) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PrismaticEventRecord) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventData

`func (o *PrismaticEventRecord) GetEventData() interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *PrismaticEventRecord) GetEventDataOk() (*interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *PrismaticEventRecord) SetEventData(v interface{})`

SetEventData sets EventData field to given value.


### SetEventDataNil

`func (o *PrismaticEventRecord) SetEventDataNil(b bool)`

 SetEventDataNil sets the value for EventData to be an explicit nil

### UnsetEventData
`func (o *PrismaticEventRecord) UnsetEventData()`

UnsetEventData ensures that no value is present for EventData, not even an explicit nil
### GetPublishedAt

`func (o *PrismaticEventRecord) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PrismaticEventRecord) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PrismaticEventRecord) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetProcessedAt

`func (o *PrismaticEventRecord) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *PrismaticEventRecord) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *PrismaticEventRecord) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *PrismaticEventRecord) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetProcessAfter

`func (o *PrismaticEventRecord) GetProcessAfter() time.Time`

GetProcessAfter returns the ProcessAfter field if non-nil, zero value otherwise.

### GetProcessAfterOk

`func (o *PrismaticEventRecord) GetProcessAfterOk() (*time.Time, bool)`

GetProcessAfterOk returns a tuple with the ProcessAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessAfter

`func (o *PrismaticEventRecord) SetProcessAfter(v time.Time)`

SetProcessAfter sets ProcessAfter field to given value.


### GetRetry

`func (o *PrismaticEventRecord) GetRetry() int64`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *PrismaticEventRecord) GetRetryOk() (*int64, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *PrismaticEventRecord) SetRetry(v int64)`

SetRetry sets Retry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


