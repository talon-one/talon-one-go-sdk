# IntegrationHubEventRecord

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

### NewIntegrationHubEventRecord

`func NewIntegrationHubEventRecord(id int64, flowId int64, eventType string, eventData interface{}, publishedAt time.Time, processAfter time.Time, retry int64, ) *IntegrationHubEventRecord`

NewIntegrationHubEventRecord instantiates a new IntegrationHubEventRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventRecordWithDefaults

`func NewIntegrationHubEventRecordWithDefaults() *IntegrationHubEventRecord`

NewIntegrationHubEventRecordWithDefaults instantiates a new IntegrationHubEventRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationHubEventRecord) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationHubEventRecord) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationHubEventRecord) SetId(v int64)`

SetId sets Id field to given value.


### GetFlowId

`func (o *IntegrationHubEventRecord) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *IntegrationHubEventRecord) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *IntegrationHubEventRecord) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.


### GetEventType

`func (o *IntegrationHubEventRecord) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IntegrationHubEventRecord) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IntegrationHubEventRecord) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventData

`func (o *IntegrationHubEventRecord) GetEventData() interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *IntegrationHubEventRecord) GetEventDataOk() (*interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *IntegrationHubEventRecord) SetEventData(v interface{})`

SetEventData sets EventData field to given value.


### SetEventDataNil

`func (o *IntegrationHubEventRecord) SetEventDataNil(b bool)`

 SetEventDataNil sets the value for EventData to be an explicit nil

### UnsetEventData
`func (o *IntegrationHubEventRecord) UnsetEventData()`

UnsetEventData ensures that no value is present for EventData, not even an explicit nil
### GetPublishedAt

`func (o *IntegrationHubEventRecord) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *IntegrationHubEventRecord) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *IntegrationHubEventRecord) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetProcessedAt

`func (o *IntegrationHubEventRecord) GetProcessedAt() time.Time`

GetProcessedAt returns the ProcessedAt field if non-nil, zero value otherwise.

### GetProcessedAtOk

`func (o *IntegrationHubEventRecord) GetProcessedAtOk() (*time.Time, bool)`

GetProcessedAtOk returns a tuple with the ProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedAt

`func (o *IntegrationHubEventRecord) SetProcessedAt(v time.Time)`

SetProcessedAt sets ProcessedAt field to given value.

### HasProcessedAt

`func (o *IntegrationHubEventRecord) HasProcessedAt() bool`

HasProcessedAt returns a boolean if a field has been set.

### GetProcessAfter

`func (o *IntegrationHubEventRecord) GetProcessAfter() time.Time`

GetProcessAfter returns the ProcessAfter field if non-nil, zero value otherwise.

### GetProcessAfterOk

`func (o *IntegrationHubEventRecord) GetProcessAfterOk() (*time.Time, bool)`

GetProcessAfterOk returns a tuple with the ProcessAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessAfter

`func (o *IntegrationHubEventRecord) SetProcessAfter(v time.Time)`

SetProcessAfter sets ProcessAfter field to given value.


### GetRetry

`func (o *IntegrationHubEventRecord) GetRetry() int64`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *IntegrationHubEventRecord) GetRetryOk() (*int64, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *IntegrationHubEventRecord) SetRetry(v int64)`

SetRetry sets Retry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


