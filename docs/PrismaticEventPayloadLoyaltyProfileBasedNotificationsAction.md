# PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float32** |  | 
**Reason** | Pointer to **string** |  | [optional] 
**Operation** | **string** |  | 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ExpiryDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsAction

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsAction(amount float32, operation string, ) *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction`

NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsAction instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsActionWithDefaults

`func NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsActionWithDefaults() *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction`

NewPrismaticEventPayloadLoyaltyProfileBasedNotificationsActionWithDefaults instantiates a new PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetReason

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetOperation

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetStartDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *PrismaticEventPayloadLoyaltyProfileBasedNotificationsAction) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


