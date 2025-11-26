# ListAllRolesV2200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalResultSize** | **int64** |  | 
**Data** | [**[]RoleV2**](RoleV2.md) |  | 

## Methods

### NewListAllRolesV2200Response

`func NewListAllRolesV2200Response(totalResultSize int64, data []RoleV2, ) *ListAllRolesV2200Response`

NewListAllRolesV2200Response instantiates a new ListAllRolesV2200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAllRolesV2200ResponseWithDefaults

`func NewListAllRolesV2200ResponseWithDefaults() *ListAllRolesV2200Response`

NewListAllRolesV2200ResponseWithDefaults instantiates a new ListAllRolesV2200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalResultSize

`func (o *ListAllRolesV2200Response) GetTotalResultSize() int64`

GetTotalResultSize returns the TotalResultSize field if non-nil, zero value otherwise.

### GetTotalResultSizeOk

`func (o *ListAllRolesV2200Response) GetTotalResultSizeOk() (*int64, bool)`

GetTotalResultSizeOk returns a tuple with the TotalResultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultSize

`func (o *ListAllRolesV2200Response) SetTotalResultSize(v int64)`

SetTotalResultSize sets TotalResultSize field to given value.


### GetData

`func (o *ListAllRolesV2200Response) GetData() []RoleV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAllRolesV2200Response) GetDataOk() (*[]RoleV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAllRolesV2200Response) SetData(v []RoleV2)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


