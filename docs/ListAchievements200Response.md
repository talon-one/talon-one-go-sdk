# ListAchievements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | [**[]Achievement**](Achievement.md) |  | 

## Methods

### NewListAchievements200Response

`func NewListAchievements200Response(data []Achievement, ) *ListAchievements200Response`

NewListAchievements200Response instantiates a new ListAchievements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAchievements200ResponseWithDefaults

`func NewListAchievements200ResponseWithDefaults() *ListAchievements200Response`

NewListAchievements200ResponseWithDefaults instantiates a new ListAchievements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListAchievements200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListAchievements200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListAchievements200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListAchievements200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *ListAchievements200Response) GetData() []Achievement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAchievements200Response) GetDataOk() (*[]Achievement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAchievements200Response) SetData(v []Achievement)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


