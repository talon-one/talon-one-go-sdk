# OutboundLogStreamCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | **string** | Cursor for the next page of results. Pass it as the &#x60;cursor&#x60; parameter to fetch the next page. | 

## Methods

### NewOutboundLogStreamCursor

`func NewOutboundLogStreamCursor(nextCursor string, ) *OutboundLogStreamCursor`

NewOutboundLogStreamCursor instantiates a new OutboundLogStreamCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogStreamCursorWithDefaults

`func NewOutboundLogStreamCursorWithDefaults() *OutboundLogStreamCursor`

NewOutboundLogStreamCursorWithDefaults instantiates a new OutboundLogStreamCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *OutboundLogStreamCursor) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OutboundLogStreamCursor) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OutboundLogStreamCursor) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


