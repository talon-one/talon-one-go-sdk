# SessionCoupons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionIntegrationId** | **string** | The integration ID of the session in which the coupons were applied. | 
**CouponCode** | Pointer to **string** | The coupon codes for which rejection reason is needed. | [optional] 

## Methods

### NewSessionCoupons

`func NewSessionCoupons(sessionIntegrationId string, ) *SessionCoupons`

NewSessionCoupons instantiates a new SessionCoupons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionCouponsWithDefaults

`func NewSessionCouponsWithDefaults() *SessionCoupons`

NewSessionCouponsWithDefaults instantiates a new SessionCoupons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionIntegrationId

`func (o *SessionCoupons) GetSessionIntegrationId() string`

GetSessionIntegrationId returns the SessionIntegrationId field if non-nil, zero value otherwise.

### GetSessionIntegrationIdOk

`func (o *SessionCoupons) GetSessionIntegrationIdOk() (*string, bool)`

GetSessionIntegrationIdOk returns a tuple with the SessionIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIntegrationId

`func (o *SessionCoupons) SetSessionIntegrationId(v string)`

SetSessionIntegrationId sets SessionIntegrationId field to given value.


### GetCouponCode

`func (o *SessionCoupons) GetCouponCode() string`

GetCouponCode returns the CouponCode field if non-nil, zero value otherwise.

### GetCouponCodeOk

`func (o *SessionCoupons) GetCouponCodeOk() (*string, bool)`

GetCouponCodeOk returns a tuple with the CouponCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponCode

`func (o *SessionCoupons) SetCouponCode(v string)`

SetCouponCode sets CouponCode field to given value.

### HasCouponCode

`func (o *SessionCoupons) HasCouponCode() bool`

HasCouponCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


