# GetAccessLogsWithoutTotalCount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]AccessLogEntry**](AccessLogEntry.md) |  | 

## Methods

### NewGetAccessLogsWithoutTotalCount200Response

`func NewGetAccessLogsWithoutTotalCount200Response(hasMore bool, data []AccessLogEntry, ) *GetAccessLogsWithoutTotalCount200Response`

NewGetAccessLogsWithoutTotalCount200Response instantiates a new GetAccessLogsWithoutTotalCount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessLogsWithoutTotalCount200ResponseWithDefaults

`func NewGetAccessLogsWithoutTotalCount200ResponseWithDefaults() *GetAccessLogsWithoutTotalCount200Response`

NewGetAccessLogsWithoutTotalCount200ResponseWithDefaults instantiates a new GetAccessLogsWithoutTotalCount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAccessLogsWithoutTotalCount200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAccessLogsWithoutTotalCount200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAccessLogsWithoutTotalCount200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetAccessLogsWithoutTotalCount200Response) GetData() []AccessLogEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAccessLogsWithoutTotalCount200Response) GetDataOk() (*[]AccessLogEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAccessLogsWithoutTotalCount200Response) SetData(v []AccessLogEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


