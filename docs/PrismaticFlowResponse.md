# PrismaticFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | ID of the prismatic flow. | 
**ApplicationID** | Pointer to **int64** | ID of application the flow is registered for. | [optional] 
**EventType** | **string** | The event type we want to register a flow for. | 
**PrismaticFlowUrl** | **string** | The URL of the prismatic flow that we want to trigger for the event. | 
**Config** | [**PrismaticFlowConfigResponse**](PrismaticFlowConfigResponse.md) |  | 

## Methods

### NewPrismaticFlowResponse

`func NewPrismaticFlowResponse(id int64, eventType string, prismaticFlowUrl string, config PrismaticFlowConfigResponse, ) *PrismaticFlowResponse`

NewPrismaticFlowResponse instantiates a new PrismaticFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticFlowResponseWithDefaults

`func NewPrismaticFlowResponseWithDefaults() *PrismaticFlowResponse`

NewPrismaticFlowResponseWithDefaults instantiates a new PrismaticFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrismaticFlowResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrismaticFlowResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrismaticFlowResponse) SetId(v int64)`

SetId sets Id field to given value.


### GetApplicationID

`func (o *PrismaticFlowResponse) GetApplicationID() int64`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *PrismaticFlowResponse) GetApplicationIDOk() (*int64, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *PrismaticFlowResponse) SetApplicationID(v int64)`

SetApplicationID sets ApplicationID field to given value.

### HasApplicationID

`func (o *PrismaticFlowResponse) HasApplicationID() bool`

HasApplicationID returns a boolean if a field has been set.

### GetEventType

`func (o *PrismaticFlowResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PrismaticFlowResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PrismaticFlowResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetPrismaticFlowUrl

`func (o *PrismaticFlowResponse) GetPrismaticFlowUrl() string`

GetPrismaticFlowUrl returns the PrismaticFlowUrl field if non-nil, zero value otherwise.

### GetPrismaticFlowUrlOk

`func (o *PrismaticFlowResponse) GetPrismaticFlowUrlOk() (*string, bool)`

GetPrismaticFlowUrlOk returns a tuple with the PrismaticFlowUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrismaticFlowUrl

`func (o *PrismaticFlowResponse) SetPrismaticFlowUrl(v string)`

SetPrismaticFlowUrl sets PrismaticFlowUrl field to given value.


### GetConfig

`func (o *PrismaticFlowResponse) GetConfig() PrismaticFlowConfigResponse`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PrismaticFlowResponse) GetConfigOk() (*PrismaticFlowConfigResponse, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PrismaticFlowResponse) SetConfig(v PrismaticFlowConfigResponse)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


