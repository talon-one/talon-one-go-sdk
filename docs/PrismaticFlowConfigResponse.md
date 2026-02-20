# PrismaticFlowConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkerCount** | Pointer to **int64** | Number of Prismatic workers to run in parallel for this flow (maximum 500). | [optional] 
**MaxEventsPerMessage** | Pointer to **int64** | Maximum number of events to send in a single message to Prismatic. | [optional] 
**MaxRetries** | Pointer to **int64** | Maximum number of retries for a Prismatic event before it is ignored. | [optional] 

## Methods

### NewPrismaticFlowConfigResponse

`func NewPrismaticFlowConfigResponse() *PrismaticFlowConfigResponse`

NewPrismaticFlowConfigResponse instantiates a new PrismaticFlowConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticFlowConfigResponseWithDefaults

`func NewPrismaticFlowConfigResponseWithDefaults() *PrismaticFlowConfigResponse`

NewPrismaticFlowConfigResponseWithDefaults instantiates a new PrismaticFlowConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkerCount

`func (o *PrismaticFlowConfigResponse) GetWorkerCount() int64`

GetWorkerCount returns the WorkerCount field if non-nil, zero value otherwise.

### GetWorkerCountOk

`func (o *PrismaticFlowConfigResponse) GetWorkerCountOk() (*int64, bool)`

GetWorkerCountOk returns a tuple with the WorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerCount

`func (o *PrismaticFlowConfigResponse) SetWorkerCount(v int64)`

SetWorkerCount sets WorkerCount field to given value.

### HasWorkerCount

`func (o *PrismaticFlowConfigResponse) HasWorkerCount() bool`

HasWorkerCount returns a boolean if a field has been set.

### GetMaxEventsPerMessage

`func (o *PrismaticFlowConfigResponse) GetMaxEventsPerMessage() int64`

GetMaxEventsPerMessage returns the MaxEventsPerMessage field if non-nil, zero value otherwise.

### GetMaxEventsPerMessageOk

`func (o *PrismaticFlowConfigResponse) GetMaxEventsPerMessageOk() (*int64, bool)`

GetMaxEventsPerMessageOk returns a tuple with the MaxEventsPerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEventsPerMessage

`func (o *PrismaticFlowConfigResponse) SetMaxEventsPerMessage(v int64)`

SetMaxEventsPerMessage sets MaxEventsPerMessage field to given value.

### HasMaxEventsPerMessage

`func (o *PrismaticFlowConfigResponse) HasMaxEventsPerMessage() bool`

HasMaxEventsPerMessage returns a boolean if a field has been set.

### GetMaxRetries

`func (o *PrismaticFlowConfigResponse) GetMaxRetries() int64`

GetMaxRetries returns the MaxRetries field if non-nil, zero value otherwise.

### GetMaxRetriesOk

`func (o *PrismaticFlowConfigResponse) GetMaxRetriesOk() (*int64, bool)`

GetMaxRetriesOk returns a tuple with the MaxRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetries

`func (o *PrismaticFlowConfigResponse) SetMaxRetries(v int64)`

SetMaxRetries sets MaxRetries field to given value.

### HasMaxRetries

`func (o *PrismaticFlowConfigResponse) HasMaxRetries() bool`

HasMaxRetries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


