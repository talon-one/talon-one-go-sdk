# ExportJobLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this entity. | 
**ExportJobId** | **int64** | The internal ID of the connected export job. | 
**Type** | [**ExportJobType**](ExportJobType.md) |  | 
**Result** | [**ExportJobLogResult**](ExportJobLogResult.md) |  | 
**Filename** | **string** | Name of the exported file, without extension. | 
**Log** | **string** | The log text. | 
**StartedAt** | **time.Time** | Time when the export job run started. | 
**FinishedAt** | Pointer to **time.Time** | Time when the export job run finished. | [optional] 

## Methods

### NewExportJobLog

`func NewExportJobLog(id int64, exportJobId int64, type_ ExportJobType, result ExportJobLogResult, filename string, log string, startedAt time.Time, ) *ExportJobLog`

NewExportJobLog instantiates a new ExportJobLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobLogWithDefaults

`func NewExportJobLogWithDefaults() *ExportJobLog`

NewExportJobLogWithDefaults instantiates a new ExportJobLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExportJobLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportJobLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportJobLog) SetId(v int64)`

SetId sets Id field to given value.


### GetExportJobId

`func (o *ExportJobLog) GetExportJobId() int64`

GetExportJobId returns the ExportJobId field if non-nil, zero value otherwise.

### GetExportJobIdOk

`func (o *ExportJobLog) GetExportJobIdOk() (*int64, bool)`

GetExportJobIdOk returns a tuple with the ExportJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportJobId

`func (o *ExportJobLog) SetExportJobId(v int64)`

SetExportJobId sets ExportJobId field to given value.


### GetType

`func (o *ExportJobLog) GetType() ExportJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportJobLog) GetTypeOk() (*ExportJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportJobLog) SetType(v ExportJobType)`

SetType sets Type field to given value.


### GetResult

`func (o *ExportJobLog) GetResult() ExportJobLogResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ExportJobLog) GetResultOk() (*ExportJobLogResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ExportJobLog) SetResult(v ExportJobLogResult)`

SetResult sets Result field to given value.


### GetFilename

`func (o *ExportJobLog) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExportJobLog) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExportJobLog) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetLog

`func (o *ExportJobLog) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *ExportJobLog) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *ExportJobLog) SetLog(v string)`

SetLog sets Log field to given value.


### GetStartedAt

`func (o *ExportJobLog) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ExportJobLog) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ExportJobLog) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetFinishedAt

`func (o *ExportJobLog) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ExportJobLog) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ExportJobLog) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ExportJobLog) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


