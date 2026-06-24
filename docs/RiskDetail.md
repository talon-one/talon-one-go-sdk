# RiskDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The internal ID of this entity. | 
**Created** | **time.Time** | The time this entity was created. | 
**NotificationId** | **int64** | The ID of the risk notification rule that flagged this risk. | 
**RunDate** | **string** | The date of the ML pipeline run that detected this risk. | 
**GroupKey** | **string** | The Application group this risk was detected in. Contains the Application ID, or &#x60;__GLOBAL__&#x60; for metrics that are not grouped by Application.  | 
**ApplicationId** | Pointer to **int64** | The ID of the Application this risk belongs to. Absent for global metrics. | [optional] 
**Status** | **string** | The triage lifecycle status of this risk. | 
**Criticality** | **string** | The critical classification bucket of this risk. | 
**Entity** | **string** | The entity type the risk was detected in. | 
**Activity** | **string** | The activity metric the risk was detected in. | 
**TimeFrame** | **string** | The rolling time window of the risk evaluation. | 
**ReportedDate** | **time.Time** | The time the ML service reported this risk. | 
**AffectedEntityCount** | **int64** | The total number of entities affected by this risk. | 
**Description** | Pointer to **string** | Human-readable description of the detected anomaly. | [optional] 
**Modified** | **time.Time** | Timestamp of the most recent update. | 
**AffectedEntities** | [**[]RiskAffectedEntityItem**](RiskAffectedEntityItem.md) | The affected entities with the highest severity ratios, in descending order. | 

## Methods

### NewRiskDetail

`func NewRiskDetail(id int64, created time.Time, notificationId int64, runDate string, groupKey string, status string, criticality string, entity string, activity string, timeFrame string, reportedDate time.Time, affectedEntityCount int64, modified time.Time, affectedEntities []RiskAffectedEntityItem, ) *RiskDetail`

NewRiskDetail instantiates a new RiskDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskDetailWithDefaults

`func NewRiskDetailWithDefaults() *RiskDetail`

NewRiskDetailWithDefaults instantiates a new RiskDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskDetail) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *RiskDetail) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskDetail) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskDetail) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetNotificationId

`func (o *RiskDetail) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *RiskDetail) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *RiskDetail) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.


### GetRunDate

`func (o *RiskDetail) GetRunDate() string`

GetRunDate returns the RunDate field if non-nil, zero value otherwise.

### GetRunDateOk

`func (o *RiskDetail) GetRunDateOk() (*string, bool)`

GetRunDateOk returns a tuple with the RunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunDate

`func (o *RiskDetail) SetRunDate(v string)`

SetRunDate sets RunDate field to given value.


### GetGroupKey

`func (o *RiskDetail) GetGroupKey() string`

GetGroupKey returns the GroupKey field if non-nil, zero value otherwise.

### GetGroupKeyOk

`func (o *RiskDetail) GetGroupKeyOk() (*string, bool)`

GetGroupKeyOk returns a tuple with the GroupKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupKey

`func (o *RiskDetail) SetGroupKey(v string)`

SetGroupKey sets GroupKey field to given value.


### GetApplicationId

`func (o *RiskDetail) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *RiskDetail) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *RiskDetail) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *RiskDetail) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetStatus

`func (o *RiskDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCriticality

`func (o *RiskDetail) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *RiskDetail) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *RiskDetail) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.


### GetEntity

`func (o *RiskDetail) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RiskDetail) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RiskDetail) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetActivity

`func (o *RiskDetail) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *RiskDetail) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *RiskDetail) SetActivity(v string)`

SetActivity sets Activity field to given value.


### GetTimeFrame

`func (o *RiskDetail) GetTimeFrame() string`

GetTimeFrame returns the TimeFrame field if non-nil, zero value otherwise.

### GetTimeFrameOk

`func (o *RiskDetail) GetTimeFrameOk() (*string, bool)`

GetTimeFrameOk returns a tuple with the TimeFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrame

`func (o *RiskDetail) SetTimeFrame(v string)`

SetTimeFrame sets TimeFrame field to given value.


### GetReportedDate

`func (o *RiskDetail) GetReportedDate() time.Time`

GetReportedDate returns the ReportedDate field if non-nil, zero value otherwise.

### GetReportedDateOk

`func (o *RiskDetail) GetReportedDateOk() (*time.Time, bool)`

GetReportedDateOk returns a tuple with the ReportedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedDate

`func (o *RiskDetail) SetReportedDate(v time.Time)`

SetReportedDate sets ReportedDate field to given value.


### GetAffectedEntityCount

`func (o *RiskDetail) GetAffectedEntityCount() int64`

GetAffectedEntityCount returns the AffectedEntityCount field if non-nil, zero value otherwise.

### GetAffectedEntityCountOk

`func (o *RiskDetail) GetAffectedEntityCountOk() (*int64, bool)`

GetAffectedEntityCountOk returns a tuple with the AffectedEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntityCount

`func (o *RiskDetail) SetAffectedEntityCount(v int64)`

SetAffectedEntityCount sets AffectedEntityCount field to given value.


### GetDescription

`func (o *RiskDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetModified

`func (o *RiskDetail) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *RiskDetail) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *RiskDetail) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAffectedEntities

`func (o *RiskDetail) GetAffectedEntities() []RiskAffectedEntityItem`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RiskDetail) GetAffectedEntitiesOk() (*[]RiskAffectedEntityItem, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RiskDetail) SetAffectedEntities(v []RiskAffectedEntityItem)`

SetAffectedEntities sets AffectedEntities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


