# AsyncCouponsCreatedNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchID** | **string** | The ID of the batch to which the coupon belongs.  **Note:** The Batch ID is generated when coupons are created.  | 
**TypeOfChange** | **string** |  | 
**Operation** | **string** |  | 
**EmployeeName** | **string** |  | 
**NotificationType** | **string** | The type of the notification | 

## Methods

### NewAsyncCouponsCreatedNotification

`func NewAsyncCouponsCreatedNotification(batchID string, typeOfChange string, operation string, employeeName string, notificationType string, ) *AsyncCouponsCreatedNotification`

NewAsyncCouponsCreatedNotification instantiates a new AsyncCouponsCreatedNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsyncCouponsCreatedNotificationWithDefaults

`func NewAsyncCouponsCreatedNotificationWithDefaults() *AsyncCouponsCreatedNotification`

NewAsyncCouponsCreatedNotificationWithDefaults instantiates a new AsyncCouponsCreatedNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchID

`func (o *AsyncCouponsCreatedNotification) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *AsyncCouponsCreatedNotification) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *AsyncCouponsCreatedNotification) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.


### GetTypeOfChange

`func (o *AsyncCouponsCreatedNotification) GetTypeOfChange() string`

GetTypeOfChange returns the TypeOfChange field if non-nil, zero value otherwise.

### GetTypeOfChangeOk

`func (o *AsyncCouponsCreatedNotification) GetTypeOfChangeOk() (*string, bool)`

GetTypeOfChangeOk returns a tuple with the TypeOfChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfChange

`func (o *AsyncCouponsCreatedNotification) SetTypeOfChange(v string)`

SetTypeOfChange sets TypeOfChange field to given value.


### GetOperation

`func (o *AsyncCouponsCreatedNotification) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *AsyncCouponsCreatedNotification) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *AsyncCouponsCreatedNotification) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetEmployeeName

`func (o *AsyncCouponsCreatedNotification) GetEmployeeName() string`

GetEmployeeName returns the EmployeeName field if non-nil, zero value otherwise.

### GetEmployeeNameOk

`func (o *AsyncCouponsCreatedNotification) GetEmployeeNameOk() (*string, bool)`

GetEmployeeNameOk returns a tuple with the EmployeeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeName

`func (o *AsyncCouponsCreatedNotification) SetEmployeeName(v string)`

SetEmployeeName sets EmployeeName field to given value.


### GetNotificationType

`func (o *AsyncCouponsCreatedNotification) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *AsyncCouponsCreatedNotification) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *AsyncCouponsCreatedNotification) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


