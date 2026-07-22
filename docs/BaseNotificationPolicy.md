# BaseNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the notification. | 
**Triggers** | [**[]TierWillDowngradeNotificationTrigger**](TierWillDowngradeNotificationTrigger.md) |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int64** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] [default to 1000]
**Scopes** | **[]string** |  | 
**IncludeData** | Pointer to **bool** | Indicates whether to include all generated coupons. If &#x60;false&#x60;, only the &#x60;batchId&#x60; of the generated coupons is included. | [optional] 
**AheadOfDaysTrigger** | Pointer to **int64** | The number of days in advance that strikethrough pricing updates should be sent. | [optional] 

## Methods

### NewBaseNotificationPolicy

`func NewBaseNotificationPolicy(name string, triggers []TierWillDowngradeNotificationTrigger, scopes []string, ) *BaseNotificationPolicy`

NewBaseNotificationPolicy instantiates a new BaseNotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseNotificationPolicyWithDefaults

`func NewBaseNotificationPolicyWithDefaults() *BaseNotificationPolicy`

NewBaseNotificationPolicyWithDefaults instantiates a new BaseNotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseNotificationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseNotificationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTriggers

`func (o *BaseNotificationPolicy) GetTriggers() []TierWillDowngradeNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *BaseNotificationPolicy) GetTriggersOk() (*[]TierWillDowngradeNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *BaseNotificationPolicy) SetTriggers(v []TierWillDowngradeNotificationTrigger)`

SetTriggers sets Triggers field to given value.


### GetBatchingEnabled

`func (o *BaseNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *BaseNotificationPolicy) GetBatchingEnabledOk() (*bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchingEnabled

`func (o *BaseNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled sets BatchingEnabled field to given value.

### HasBatchingEnabled

`func (o *BaseNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### GetBatchSize

`func (o *BaseNotificationPolicy) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *BaseNotificationPolicy) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *BaseNotificationPolicy) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *BaseNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetScopes

`func (o *BaseNotificationPolicy) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *BaseNotificationPolicy) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *BaseNotificationPolicy) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetIncludeData

`func (o *BaseNotificationPolicy) GetIncludeData() bool`

GetIncludeData returns the IncludeData field if non-nil, zero value otherwise.

### GetIncludeDataOk

`func (o *BaseNotificationPolicy) GetIncludeDataOk() (*bool, bool)`

GetIncludeDataOk returns a tuple with the IncludeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeData

`func (o *BaseNotificationPolicy) SetIncludeData(v bool)`

SetIncludeData sets IncludeData field to given value.

### HasIncludeData

`func (o *BaseNotificationPolicy) HasIncludeData() bool`

HasIncludeData returns a boolean if a field has been set.

### GetAheadOfDaysTrigger

`func (o *BaseNotificationPolicy) GetAheadOfDaysTrigger() int64`

GetAheadOfDaysTrigger returns the AheadOfDaysTrigger field if non-nil, zero value otherwise.

### GetAheadOfDaysTriggerOk

`func (o *BaseNotificationPolicy) GetAheadOfDaysTriggerOk() (*int64, bool)`

GetAheadOfDaysTriggerOk returns a tuple with the AheadOfDaysTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAheadOfDaysTrigger

`func (o *BaseNotificationPolicy) SetAheadOfDaysTrigger(v int64)`

SetAheadOfDaysTrigger sets AheadOfDaysTrigger field to given value.

### HasAheadOfDaysTrigger

`func (o *BaseNotificationPolicy) HasAheadOfDaysTrigger() bool`

HasAheadOfDaysTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


