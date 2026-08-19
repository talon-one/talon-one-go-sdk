# ListAchievementsV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | [**[]AchievementV2**](AchievementV2.md) |  | 

## Methods

### NewListAchievementsV2200Response

`func NewListAchievementsV2200Response(data []AchievementV2, ) *ListAchievementsV2200Response`

NewListAchievementsV2200Response instantiates a new ListAchievementsV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAchievementsV2200ResponseWithDefaults

`func NewListAchievementsV2200ResponseWithDefaults() *ListAchievementsV2200Response`

NewListAchievementsV2200ResponseWithDefaults instantiates a new ListAchievementsV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListAchievementsV2200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListAchievementsV2200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListAchievementsV2200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListAchievementsV2200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *ListAchievementsV2200Response) GetData() []AchievementV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAchievementsV2200Response) GetDataOk() (*[]AchievementV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAchievementsV2200Response) SetData(v []AchievementV2)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


