# GetCampaignTemplates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]CampaignTemplate**](CampaignTemplate.md) |  | 

## Methods

### NewGetCampaignTemplates200Response

`func NewGetCampaignTemplates200Response(hasMore bool, data []CampaignTemplate, ) *GetCampaignTemplates200Response`

NewGetCampaignTemplates200Response instantiates a new GetCampaignTemplates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCampaignTemplates200ResponseWithDefaults

`func NewGetCampaignTemplates200ResponseWithDefaults() *GetCampaignTemplates200Response`

NewGetCampaignTemplates200ResponseWithDefaults instantiates a new GetCampaignTemplates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetCampaignTemplates200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetCampaignTemplates200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetCampaignTemplates200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetCampaignTemplates200Response) GetData() []CampaignTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCampaignTemplates200Response) GetDataOk() (*[]CampaignTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCampaignTemplates200Response) SetData(v []CampaignTemplate)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


