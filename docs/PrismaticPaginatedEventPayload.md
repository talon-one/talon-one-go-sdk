# PrismaticPaginatedEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**EventType** | **string** |  | 
**Data** | **[]interface{}** |  | 

## Methods

### NewPrismaticPaginatedEventPayload

`func NewPrismaticPaginatedEventPayload(totalResultSize int64, eventType string, data []interface{}, ) *PrismaticPaginatedEventPayload`

NewPrismaticPaginatedEventPayload instantiates a new PrismaticPaginatedEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticPaginatedEventPayloadWithDefaults

`func NewPrismaticPaginatedEventPayloadWithDefaults() *PrismaticPaginatedEventPayload`

NewPrismaticPaginatedEventPayloadWithDefaults instantiates a new PrismaticPaginatedEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *PrismaticPaginatedEventPayload) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *PrismaticPaginatedEventPayload) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *PrismaticPaginatedEventPayload) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetEventType

`func (o *PrismaticPaginatedEventPayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PrismaticPaginatedEventPayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PrismaticPaginatedEventPayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetData

`func (o *PrismaticPaginatedEventPayload) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrismaticPaginatedEventPayload) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrismaticPaginatedEventPayload) SetData(v []interface{})`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


