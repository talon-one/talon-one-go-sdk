# OutboundLogStreamLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OutboundLogStreamEntryType**](OutboundLogStreamEntryType.md) |  | 
**Data** | [**OutboundLog**](OutboundLog.md) |  | 

## Methods

### NewOutboundLogStreamLogEntry

`func NewOutboundLogStreamLogEntry(type_ OutboundLogStreamEntryType, data OutboundLog, ) *OutboundLogStreamLogEntry`

NewOutboundLogStreamLogEntry instantiates a new OutboundLogStreamLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogStreamLogEntryWithDefaults

`func NewOutboundLogStreamLogEntryWithDefaults() *OutboundLogStreamLogEntry`

NewOutboundLogStreamLogEntryWithDefaults instantiates a new OutboundLogStreamLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutboundLogStreamLogEntry) GetType() OutboundLogStreamEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutboundLogStreamLogEntry) GetTypeOk() (*OutboundLogStreamEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutboundLogStreamLogEntry) SetType(v OutboundLogStreamEntryType)`

SetType sets Type field to given value.


### GetData

`func (o *OutboundLogStreamLogEntry) GetData() OutboundLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundLogStreamLogEntry) GetDataOk() (*OutboundLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundLogStreamLogEntry) SetData(v OutboundLog)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


