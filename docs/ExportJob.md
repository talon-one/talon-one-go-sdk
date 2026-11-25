# ExportJob

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
**Id** | **int64** | The internal ID of this entity. | 
**AccountId** | **int64** | The ID of the account that owns this entity. | 
**Status** | [**ExportJobStatus**](ExportJobStatus.md) |  | [default to IDLE]
**Created** | **time.Time** | The time this entity was created. | 
**Modified** | Pointer to **time.Time** | The time this entity was modified. | [optional] 

## Methods

### NewExportJob

`func NewExportJob(type_ ExportJobType, args map[string]interface{}, destination []ExportJobDestination, cronSchedule string, id int64, accountId int64, status ExportJobStatus, created time.Time, ) *ExportJob`

NewExportJob instantiates a new ExportJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobWithDefaults

`func NewExportJobWithDefaults() *ExportJob`

NewExportJobWithDefaults instantiates a new ExportJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportJob) GetType() ExportJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportJob) GetTypeOk() (*ExportJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportJob) SetType(v ExportJobType)`

SetType sets Type field to given value.


### GetArgs

`func (o *ExportJob) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *ExportJob) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *ExportJob) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.


### GetDestination

`func (o *ExportJob) GetDestination() []ExportJobDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ExportJob) GetDestinationOk() (*[]ExportJobDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ExportJob) SetDestination(v []ExportJobDestination)`

SetDestination sets Destination field to given value.


### GetCronSchedule

`func (o *ExportJob) GetCronSchedule() string`

GetCronSchedule returns the CronSchedule field if non-nil, zero value otherwise.

### GetCronScheduleOk

`func (o *ExportJob) GetCronScheduleOk() (*string, bool)`

GetCronScheduleOk returns a tuple with the CronSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronSchedule

`func (o *ExportJob) SetCronSchedule(v string)`

SetCronSchedule sets CronSchedule field to given value.


### GetTimezone

`func (o *ExportJob) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ExportJob) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ExportJob) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ExportJob) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetEnabled

`func (o *ExportJob) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExportJob) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExportJob) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExportJob) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAllowConcurrentRuns

`func (o *ExportJob) GetAllowConcurrentRuns() bool`

GetAllowConcurrentRuns returns the AllowConcurrentRuns field if non-nil, zero value otherwise.

### GetAllowConcurrentRunsOk

`func (o *ExportJob) GetAllowConcurrentRunsOk() (*bool, bool)`

GetAllowConcurrentRunsOk returns a tuple with the AllowConcurrentRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConcurrentRuns

`func (o *ExportJob) SetAllowConcurrentRuns(v bool)`

SetAllowConcurrentRuns sets AllowConcurrentRuns field to given value.

### HasAllowConcurrentRuns

`func (o *ExportJob) HasAllowConcurrentRuns() bool`

HasAllowConcurrentRuns returns a boolean if a field has been set.

### GetOverrideFilename

`func (o *ExportJob) GetOverrideFilename() string`

GetOverrideFilename returns the OverrideFilename field if non-nil, zero value otherwise.

### GetOverrideFilenameOk

`func (o *ExportJob) GetOverrideFilenameOk() (*string, bool)`

GetOverrideFilenameOk returns a tuple with the OverrideFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFilename

`func (o *ExportJob) SetOverrideFilename(v string)`

SetOverrideFilename sets OverrideFilename field to given value.

### HasOverrideFilename

`func (o *ExportJob) HasOverrideFilename() bool`

HasOverrideFilename returns a boolean if a field has been set.

### GetId

`func (o *ExportJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExportJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExportJob) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *ExportJob) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExportJob) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExportJob) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetStatus

`func (o *ExportJob) GetStatus() ExportJobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExportJob) GetStatusOk() (*ExportJobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExportJob) SetStatus(v ExportJobStatus)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *ExportJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ExportJob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ExportJob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *ExportJob) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *ExportJob) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *ExportJob) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *ExportJob) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


