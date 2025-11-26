# PrismaticFlowWithConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | **string** | The event type we want to register a flow for. | 
**PrismaticFlowUrl** | **string** | The URL of the prismatic flow that we want to trigger for the event. | 
**Config** | [**PrismaticFlowConfig**](PrismaticFlowConfig.md) |  | 

## Methods

### NewPrismaticFlowWithConfig

`func NewPrismaticFlowWithConfig(eventType string, prismaticFlowUrl string, config PrismaticFlowConfig, ) *PrismaticFlowWithConfig`

NewPrismaticFlowWithConfig instantiates a new PrismaticFlowWithConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticFlowWithConfigWithDefaults

`func NewPrismaticFlowWithConfigWithDefaults() *PrismaticFlowWithConfig`

NewPrismaticFlowWithConfigWithDefaults instantiates a new PrismaticFlowWithConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *PrismaticFlowWithConfig) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PrismaticFlowWithConfig) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PrismaticFlowWithConfig) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetPrismaticFlowUrl

`func (o *PrismaticFlowWithConfig) GetPrismaticFlowUrl() string`

GetPrismaticFlowUrl returns the PrismaticFlowUrl field if non-nil, zero value otherwise.

### GetPrismaticFlowUrlOk

`func (o *PrismaticFlowWithConfig) GetPrismaticFlowUrlOk() (*string, bool)`

GetPrismaticFlowUrlOk returns a tuple with the PrismaticFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrismaticFlowUrl

`func (o *PrismaticFlowWithConfig) SetPrismaticFlowUrl(v string)`

SetPrismaticFlowUrl sets PrismaticFlowUrl field to given value.


### GetConfig

`func (o *PrismaticFlowWithConfig) GetConfig() PrismaticFlowConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PrismaticFlowWithConfig) GetConfigOk() (*PrismaticFlowConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PrismaticFlowWithConfig) SetConfig(v PrismaticFlowConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


