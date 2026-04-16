# IntegrationGetAllCampaigns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]IntegrationCampaign**](IntegrationCampaign.md) |  | 

## Methods

### NewIntegrationGetAllCampaigns200Response

`func NewIntegrationGetAllCampaigns200Response(hasMore bool, data []IntegrationCampaign, ) *IntegrationGetAllCampaigns200Response`

NewIntegrationGetAllCampaigns200Response instantiates a new IntegrationGetAllCampaigns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationGetAllCampaigns200ResponseWithDefaults

`func NewIntegrationGetAllCampaigns200ResponseWithDefaults() *IntegrationGetAllCampaigns200Response`

NewIntegrationGetAllCampaigns200ResponseWithDefaults instantiates a new IntegrationGetAllCampaigns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *IntegrationGetAllCampaigns200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *IntegrationGetAllCampaigns200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *IntegrationGetAllCampaigns200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *IntegrationGetAllCampaigns200Response) GetData() []IntegrationCampaign`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IntegrationGetAllCampaigns200Response) GetDataOk() (*[]IntegrationCampaign, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IntegrationGetAllCampaigns200Response) SetData(v []IntegrationCampaign)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


