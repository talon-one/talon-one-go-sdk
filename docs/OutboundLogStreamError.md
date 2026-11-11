# OutboundLogStreamError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Reason the stream ended. | 
**Id** | Pointer to **string** | ID of the error. Include it when reporting the error. | [optional] 

## Methods

### NewOutboundLogStreamError

`func NewOutboundLogStreamError(error_ string, ) *OutboundLogStreamError`

NewOutboundLogStreamError instantiates a new OutboundLogStreamError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogStreamErrorWithDefaults

`func NewOutboundLogStreamErrorWithDefaults() *OutboundLogStreamError`

NewOutboundLogStreamErrorWithDefaults instantiates a new OutboundLogStreamError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OutboundLogStreamError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OutboundLogStreamError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OutboundLogStreamError) SetError(v string)`

SetError sets Error field to given value.


### GetId

`func (o *OutboundLogStreamError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboundLogStreamError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboundLogStreamError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutboundLogStreamError) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


