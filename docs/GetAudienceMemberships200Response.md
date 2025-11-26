# GetAudienceMemberships200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Data** | [**[]CustomerProfile**](CustomerProfile.md) |  | 

## Methods

### NewGetAudienceMemberships200Response

`func NewGetAudienceMemberships200Response(data []CustomerProfile, ) *GetAudienceMemberships200Response`

NewGetAudienceMemberships200Response instantiates a new GetAudienceMemberships200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAudienceMemberships200ResponseWithDefaults

`func NewGetAudienceMemberships200ResponseWithDefaults() *GetAudienceMemberships200Response`

NewGetAudienceMemberships200ResponseWithDefaults instantiates a new GetAudienceMemberships200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetAudienceMemberships200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetAudienceMemberships200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetAudienceMemberships200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetAudienceMemberships200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetData

`func (o *GetAudienceMemberships200Response) GetData() []CustomerProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAudienceMemberships200Response) GetDataOk() (*[]CustomerProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAudienceMemberships200Response) SetData(v []CustomerProfile)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


