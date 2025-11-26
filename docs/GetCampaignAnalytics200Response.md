# GetCampaignAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]CampaignAnalytics**](CampaignAnalytics.md) |  | 

## Methods

### NewGetCampaignAnalytics200Response

`func NewGetCampaignAnalytics200Response(totalResultSize int64, data []CampaignAnalytics, ) *GetCampaignAnalytics200Response`

NewGetCampaignAnalytics200Response instantiates a new GetCampaignAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignAnalytics200ResponseWithDefaults

`func NewGetCampaignAnalytics200ResponseWithDefaults() *GetCampaignAnalytics200Response`

NewGetCampaignAnalytics200ResponseWithDefaults instantiates a new GetCampaignAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *GetCampaignAnalytics200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *GetCampaignAnalytics200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *GetCampaignAnalytics200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *GetCampaignAnalytics200Response) GetData() []CampaignAnalytics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignAnalytics200Response) GetDataOk() (*[]CampaignAnalytics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignAnalytics200Response) SetData(v []CampaignAnalytics)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


