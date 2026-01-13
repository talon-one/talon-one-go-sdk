# CouponRejections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionDetails** | [**[]SessionCoupons**](SessionCoupons.md) | Array containing details from session like session id and optional coupon code used in the session. Only the first 15 entries will be processed. | 
**ApplicationId** | **int64** | The application ID for which the coupon was used. | 
**Language** | Pointer to **string** | The language the summary will be generated in. | [optional] 

## Methods

### NewCouponRejections

`func NewCouponRejections(sessionDetails []SessionCoupons, applicationId int64, ) *CouponRejections`

NewCouponRejections instantiates a new CouponRejections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponRejectionsWithDefaults

`func NewCouponRejectionsWithDefaults() *CouponRejections`

NewCouponRejectionsWithDefaults instantiates a new CouponRejections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionDetails

`func (o *CouponRejections) GetSessionDetails() []SessionCoupons`

GetSessionDetails returns the SessionDetails field if non-nil, zero value otherwise.

### GetSessionDetailsOk

`func (o *CouponRejections) GetSessionDetailsOk() (*[]SessionCoupons, bool)`

GetSessionDetailsOk returns a tuple with the SessionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDetails

`func (o *CouponRejections) SetSessionDetails(v []SessionCoupons)`

SetSessionDetails sets SessionDetails field to given value.


### GetApplicationId

`func (o *CouponRejections) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CouponRejections) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CouponRejections) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetLanguage

`func (o *CouponRejections) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CouponRejections) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CouponRejections) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CouponRejections) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


