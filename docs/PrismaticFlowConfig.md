# PrismaticFlowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**BatchSize** | Pointer to **int64** | Available for LoyaltyCardPointsExpiring, LoyaltyPointsExpiring, TierWillDowngrade, TierDowngrade, TierUpgrade, CouponCodeExpiring, CampaignNotification | [optional] 
**PeriodsInDays** | Pointer to **[]int64** | Available for LoyaltyCardPointsExpiring, LoyaltyPointsExpiring, TierWillDowngrade | [optional] 

## Methods

### NewPrismaticFlowConfig

`func NewPrismaticFlowConfig(apiKey string, ) *PrismaticFlowConfig`

NewPrismaticFlowConfig instantiates a new PrismaticFlowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticFlowConfigWithDefaults

`func NewPrismaticFlowConfigWithDefaults() *PrismaticFlowConfig`

NewPrismaticFlowConfigWithDefaults instantiates a new PrismaticFlowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *PrismaticFlowConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *PrismaticFlowConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *PrismaticFlowConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetBatchSize

`func (o *PrismaticFlowConfig) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *PrismaticFlowConfig) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *PrismaticFlowConfig) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *PrismaticFlowConfig) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetPeriodsInDays

`func (o *PrismaticFlowConfig) GetPeriodsInDays() []int64`

GetPeriodsInDays returns the PeriodsInDays field if non-nil, zero value otherwise.

### GetPeriodsInDaysOk

`func (o *PrismaticFlowConfig) GetPeriodsInDaysOk() (*[]int64, bool)`

GetPeriodsInDaysOk returns a tuple with the PeriodsInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodsInDays

`func (o *PrismaticFlowConfig) SetPeriodsInDays(v []int64)`

SetPeriodsInDays sets PeriodsInDays field to given value.

### HasPeriodsInDays

`func (o *PrismaticFlowConfig) HasPeriodsInDays() bool`

HasPeriodsInDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


