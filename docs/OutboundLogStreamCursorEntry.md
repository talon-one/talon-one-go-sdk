# OutboundLogStreamCursorEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**OutboundLogStreamEntryType**](OutboundLogStreamEntryType.md) |  | 
**Data** | [**OutboundLogStreamCursor**](OutboundLogStreamCursor.md) |  | 

## Methods

### NewOutboundLogStreamCursorEntry

`func NewOutboundLogStreamCursorEntry(type_ OutboundLogStreamEntryType, data OutboundLogStreamCursor, ) *OutboundLogStreamCursorEntry`

NewOutboundLogStreamCursorEntry instantiates a new OutboundLogStreamCursorEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogStreamCursorEntryWithDefaults

`func NewOutboundLogStreamCursorEntryWithDefaults() *OutboundLogStreamCursorEntry`

NewOutboundLogStreamCursorEntryWithDefaults instantiates a new OutboundLogStreamCursorEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OutboundLogStreamCursorEntry) GetType() OutboundLogStreamEntryType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutboundLogStreamCursorEntry) GetTypeOk() (*OutboundLogStreamEntryType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutboundLogStreamCursorEntry) SetType(v OutboundLogStreamEntryType)`

SetType sets Type field to given value.


### GetData

`func (o *OutboundLogStreamCursorEntry) GetData() OutboundLogStreamCursor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboundLogStreamCursorEntry) GetDataOk() (*OutboundLogStreamCursor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboundLogStreamCursorEntry) SetData(v OutboundLogStreamCursor)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


