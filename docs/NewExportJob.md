# NewExportJob

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

### NewNewExportJob

`func NewNewExportJob(type_ ExportJobType, args map[string]interface{}, destination []ExportJobDestination, cronSchedule string, ) *NewExportJob`

NewNewExportJob instantiates a new NewExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewExportJobWithDefaults

`func NewNewExportJobWithDefaults() *NewExportJob`

NewNewExportJobWithDefaults instantiates a new NewExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NewExportJob) GetType() ExportJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewExportJob) GetTypeOk() (*ExportJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewExportJob) SetType(v ExportJobType)`

SetType sets Type field to given value.


### GetArgs

`func (o *NewExportJob) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *NewExportJob) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *NewExportJob) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.


### GetDestination

`func (o *NewExportJob) GetDestination() []ExportJobDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *NewExportJob) GetDestinationOk() (*[]ExportJobDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *NewExportJob) SetDestination(v []ExportJobDestination)`

SetDestination sets Destination field to given value.


### GetCronSchedule

`func (o *NewExportJob) GetCronSchedule() string`

GetCronSchedule returns the CronSchedule field if non-nil, zero value otherwise.

### GetCronScheduleOk

`func (o *NewExportJob) GetCronScheduleOk() (*string, bool)`

GetCronScheduleOk returns a tuple with the CronSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronSchedule

`func (o *NewExportJob) SetCronSchedule(v string)`

SetCronSchedule sets CronSchedule field to given value.


### GetTimezone

`func (o *NewExportJob) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewExportJob) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NewExportJob) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *NewExportJob) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *NewExportJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewExportJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewExportJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewExportJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowConcurrentRuns

`func (o *NewExportJob) GetAllowConcurrentRuns() bool`

GetAllowConcurrentRuns returns the AllowConcurrentRuns field if non-nil, zero value otherwise.

### GetAllowConcurrentRunsOk

`func (o *NewExportJob) GetAllowConcurrentRunsOk() (*bool, bool)`

GetAllowConcurrentRunsOk returns a tuple with the AllowConcurrentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConcurrentRuns

`func (o *NewExportJob) SetAllowConcurrentRuns(v bool)`

SetAllowConcurrentRuns sets AllowConcurrentRuns field to given value.

### HasAllowConcurrentRuns

`func (o *NewExportJob) HasAllowConcurrentRuns() bool`

HasAllowConcurrentRuns returns a boolean if a field has been set.

### GetOverrideFilename

`func (o *NewExportJob) GetOverrideFilename() string`

GetOverrideFilename returns the OverrideFilename field if non-nil, zero value otherwise.

### GetOverrideFilenameOk

`func (o *NewExportJob) GetOverrideFilenameOk() (*string, bool)`

GetOverrideFilenameOk returns a tuple with the OverrideFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFilename

`func (o *NewExportJob) SetOverrideFilename(v string)`

SetOverrideFilename sets OverrideFilename field to given value.

### HasOverrideFilename

`func (o *NewExportJob) HasOverrideFilename() bool`

HasOverrideFilename returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


