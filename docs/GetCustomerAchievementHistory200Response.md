# GetCustomerAchievementHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]AchievementProgress**](AchievementProgress.md) |  | 

## Methods

### NewGetCustomerAchievementHistory200Response

`func NewGetCustomerAchievementHistory200Response(totalResultSize int64, data []AchievementProgress, ) *GetCustomerAchievementHistory200Response`

NewGetCustomerAchievementHistory200Response instantiates a new GetCustomerAchievementHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerAchievementHistory200ResponseWithDefaults

`func NewGetCustomerAchievementHistory200ResponseWithDefaults() *GetCustomerAchievementHistory200Response`

NewGetCustomerAchievementHistory200ResponseWithDefaults instantiates a new GetCustomerAchievementHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetCustomerAchievementHistory200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetCustomerAchievementHistory200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetCustomerAchievementHistory200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetCustomerAchievementHistory200Response) GetData() []AchievementProgress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerAchievementHistory200Response) GetDataOk() (*[]AchievementProgress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerAchievementHistory200Response) SetData(v []AchievementProgress)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


