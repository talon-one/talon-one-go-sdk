# GetAudiencesAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | [**[]AudienceAnalytics**](AudienceAnalytics.md) |  | 

## Methods

### NewGetAudiencesAnalytics200Response

`func NewGetAudiencesAnalytics200Response(data []AudienceAnalytics, ) *GetAudiencesAnalytics200Response`

NewGetAudiencesAnalytics200Response instantiates a new GetAudiencesAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAudiencesAnalytics200ResponseWithDefaults

`func NewGetAudiencesAnalytics200ResponseWithDefaults() *GetAudiencesAnalytics200Response`

NewGetAudiencesAnalytics200ResponseWithDefaults instantiates a new GetAudiencesAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAudiencesAnalytics200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAudiencesAnalytics200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAudiencesAnalytics200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetAudiencesAnalytics200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *GetAudiencesAnalytics200Response) GetData() []AudienceAnalytics`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAudiencesAnalytics200Response) GetDataOk() (*[]AudienceAnalytics, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAudiencesAnalytics200Response) SetData(v []AudienceAnalytics)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


