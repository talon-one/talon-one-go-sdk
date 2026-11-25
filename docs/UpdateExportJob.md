# UpdateExportJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ExportJobType**](ExportJobType.md) |  | 
**Args** | **map[string]interface{}** | Arguments for the exporter. Their structure depends on &#x60;type&#x60; and is validated by the service. | 
**Destination** | [**[]ExportJobDestination**](ExportJobDestination.md) | Destinations for the export job. | 
**CronSchedule** | **string** | Schedule in standard 5-field cron format. | 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | [optional] [default to "UTC"]
**Enabled** | Pointer to **bool** | Whether the export job is enabled. | [optional] [default to false]
**AllowConcurrentRuns** | Pointer to **bool** | Whether to allow concurrent runs of the job. | [optional] [default to false]
**OverrideFilename** | Pointer to **string** | Filename that replaces the generated export filename, without extension. | [optional] 

## Methods

### NewUpdateExportJob

`func NewUpdateExportJob(type_ ExportJobType, args map[string]interface{}, destination []ExportJobDestination, cronSchedule string, ) *UpdateExportJob`

NewUpdateExportJob instantiates a new UpdateExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExportJobWithDefaults

`func NewUpdateExportJobWithDefaults() *UpdateExportJob`

NewUpdateExportJobWithDefaults instantiates a new UpdateExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateExportJob) GetType() ExportJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateExportJob) GetTypeOk() (*ExportJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateExportJob) SetType(v ExportJobType)`

SetType sets Type field to given value.


### GetArgs

`func (o *UpdateExportJob) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdateExportJob) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdateExportJob) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.


### GetDestination

`func (o *UpdateExportJob) GetDestination() []ExportJobDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *UpdateExportJob) GetDestinationOk() (*[]ExportJobDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *UpdateExportJob) SetDestination(v []ExportJobDestination)`

SetDestination sets Destination field to given value.


### GetCronSchedule

`func (o *UpdateExportJob) GetCronSchedule() string`

GetCronSchedule returns the CronSchedule field if non-nil, zero value otherwise.

### GetCronScheduleOk

`func (o *UpdateExportJob) GetCronScheduleOk() (*string, bool)`

GetCronScheduleOk returns a tuple with the CronSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronSchedule

`func (o *UpdateExportJob) SetCronSchedule(v string)`

SetCronSchedule sets CronSchedule field to given value.


### GetTimezone

`func (o *UpdateExportJob) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateExportJob) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateExportJob) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateExportJob) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateExportJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateExportJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateExportJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateExportJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowConcurrentRuns

`func (o *UpdateExportJob) GetAllowConcurrentRuns() bool`

GetAllowConcurrentRuns returns the AllowConcurrentRuns field if non-nil, zero value otherwise.

### GetAllowConcurrentRunsOk

`func (o *UpdateExportJob) GetAllowConcurrentRunsOk() (*bool, bool)`

GetAllowConcurrentRunsOk returns a tuple with the AllowConcurrentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConcurrentRuns

`func (o *UpdateExportJob) SetAllowConcurrentRuns(v bool)`

SetAllowConcurrentRuns sets AllowConcurrentRuns field to given value.

### HasAllowConcurrentRuns

`func (o *UpdateExportJob) HasAllowConcurrentRuns() bool`

HasAllowConcurrentRuns returns a boolean if a field has been set.

### GetOverrideFilename

`func (o *UpdateExportJob) GetOverrideFilename() string`

GetOverrideFilename returns the OverrideFilename field if non-nil, zero value otherwise.

### GetOverrideFilenameOk

`func (o *UpdateExportJob) GetOverrideFilenameOk() (*string, bool)`

GetOverrideFilenameOk returns a tuple with the OverrideFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFilename

`func (o *UpdateExportJob) SetOverrideFilename(v string)`

SetOverrideFilename sets OverrideFilename field to given value.

### HasOverrideFilename

`func (o *UpdateExportJob) HasOverrideFilename() bool`

HasOverrideFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


