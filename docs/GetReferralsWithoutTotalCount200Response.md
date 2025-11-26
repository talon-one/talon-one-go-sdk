# GetReferralsWithoutTotalCount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Data** | [**[]Referral**](Referral.md) |  | 

## Methods

### NewGetReferralsWithoutTotalCount200Response

`func NewGetReferralsWithoutTotalCount200Response(hasMore bool, data []Referral, ) *GetReferralsWithoutTotalCount200Response`

NewGetReferralsWithoutTotalCount200Response instantiates a new GetReferralsWithoutTotalCount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReferralsWithoutTotalCount200ResponseWithDefaults

`func NewGetReferralsWithoutTotalCount200ResponseWithDefaults() *GetReferralsWithoutTotalCount200Response`

NewGetReferralsWithoutTotalCount200ResponseWithDefaults instantiates a new GetReferralsWithoutTotalCount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetReferralsWithoutTotalCount200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetReferralsWithoutTotalCount200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetReferralsWithoutTotalCount200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetData

`func (o *GetReferralsWithoutTotalCount200Response) GetData() []Referral`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetReferralsWithoutTotalCount200Response) GetDataOk() (*[]Referral, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetReferralsWithoutTotalCount200Response) SetData(v []Referral)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


