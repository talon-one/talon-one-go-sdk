# GetCustomerAchievements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]AchievementStatusEntry**](AchievementStatusEntry.md) |  | 

## Methods

### NewGetCustomerAchievements200Response

`func NewGetCustomerAchievements200Response(totalResultSize int64, data []AchievementStatusEntry, ) *GetCustomerAchievements200Response`

NewGetCustomerAchievements200Response instantiates a new GetCustomerAchievements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerAchievements200ResponseWithDefaults

`func NewGetCustomerAchievements200ResponseWithDefaults() *GetCustomerAchievements200Response`

NewGetCustomerAchievements200ResponseWithDefaults instantiates a new GetCustomerAchievements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetCustomerAchievements200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetCustomerAchievements200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetCustomerAchievements200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetCustomerAchievements200Response) GetData() []AchievementStatusEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCustomerAchievements200Response) GetDataOk() (*[]AchievementStatusEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCustomerAchievements200Response) SetData(v []AchievementStatusEntry)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


