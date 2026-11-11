# OutboundLogStreamEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OutboundLogStreamEntryType**](OutboundLogStreamEntryType.md) |  | 
**Data** | [**OutboundLogStreamCursor**](OutboundLogStreamCursor.md) |  | 

## Methods

### NewOutboundLogStreamEntry

`func NewOutboundLogStreamEntry(type_ OutboundLogStreamEntryType, data OutboundLogStreamCursor, ) *OutboundLogStreamEntry`

NewOutboundLogStreamEntry instantiates a new OutboundLogStreamEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogStreamEntryWithDefaults

`func NewOutboundLogStreamEntryWithDefaults() *OutboundLogStreamEntry`

NewOutboundLogStreamEntryWithDefaults instantiates a new OutboundLogStreamEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutboundLogStreamEntry) GetType() OutboundLogStreamEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutboundLogStreamEntry) GetTypeOk() (*OutboundLogStreamEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutboundLogStreamEntry) SetType(v OutboundLogStreamEntryType)`

SetType sets Type field to given value.


### GetData

`func (o *OutboundLogStreamEntry) GetData() OutboundLogStreamCursor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundLogStreamEntry) GetDataOk() (*OutboundLogStreamCursor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundLogStreamEntry) SetData(v OutboundLogStreamCursor)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


