# OutboundLogsDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | UUID of the outbound message. | 
**NotificationId** | Pointer to **int64** | ID of the notification that triggered the outbound message. | [optional] 
**NotificationName** | Pointer to **string** | Name of the notification that triggered the outbound message. | [optional] 
**WebhookId** | Pointer to **int64** | ID of the webhook that triggered the outbound message. | [optional] 
**WebhookName** | Pointer to **string** | Name of the webhook that triggered the outbound message. | [optional] 
**NotificationType** | **string** | Type of notification that triggered the outbound message. | 
**ApplicationId** | Pointer to **int64** | ID of the Application associated with the outbound message. | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | ID of the loyalty program associated with the outbound message. | [optional] 
**Request** | Pointer to [**OutboundLogRequest**](OutboundLogRequest.md) |  | [optional] 
**CreatedAt** | **time.Time** | Timestamp when the log was created. | 
**ProcessingTimeMs** | **int64** | Time in milliseconds between sending the outbound request and receiving the response. | 
**Response** | Pointer to [**OutboundLogResponse**](OutboundLogResponse.md) |  | [optional] 
**Error** | **string** | Reason the stream ended. | 
**Id** | Pointer to **string** | ID of the error. Include it when reporting the error. | [optional] 

## Methods

### NewOutboundLogsDataInner

`func NewOutboundLogsDataInner(uuid string, notificationType string, createdAt time.Time, processingTimeMs int64, error_ string, ) *OutboundLogsDataInner`

NewOutboundLogsDataInner instantiates a new OutboundLogsDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogsDataInnerWithDefaults

`func NewOutboundLogsDataInnerWithDefaults() *OutboundLogsDataInner`

NewOutboundLogsDataInnerWithDefaults instantiates a new OutboundLogsDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *OutboundLogsDataInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OutboundLogsDataInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OutboundLogsDataInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetNotificationId

`func (o *OutboundLogsDataInner) GetNotificationId() int64`

GetNotificationId returns the NotificationId field if non-nil, zero value otherwise.

### GetNotificationIdOk

`func (o *OutboundLogsDataInner) GetNotificationIdOk() (*int64, bool)`

GetNotificationIdOk returns a tuple with the NotificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationId

`func (o *OutboundLogsDataInner) SetNotificationId(v int64)`

SetNotificationId sets NotificationId field to given value.

### HasNotificationId

`func (o *OutboundLogsDataInner) HasNotificationId() bool`

HasNotificationId returns a boolean if a field has been set.

### GetNotificationName

`func (o *OutboundLogsDataInner) GetNotificationName() string`

GetNotificationName returns the NotificationName field if non-nil, zero value otherwise.

### GetNotificationNameOk

`func (o *OutboundLogsDataInner) GetNotificationNameOk() (*string, bool)`

GetNotificationNameOk returns a tuple with the NotificationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationName

`func (o *OutboundLogsDataInner) SetNotificationName(v string)`

SetNotificationName sets NotificationName field to given value.

### HasNotificationName

`func (o *OutboundLogsDataInner) HasNotificationName() bool`

HasNotificationName returns a boolean if a field has been set.

### GetWebhookId

`func (o *OutboundLogsDataInner) GetWebhookId() int64`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *OutboundLogsDataInner) GetWebhookIdOk() (*int64, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *OutboundLogsDataInner) SetWebhookId(v int64)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *OutboundLogsDataInner) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.

### GetWebhookName

`func (o *OutboundLogsDataInner) GetWebhookName() string`

GetWebhookName returns the WebhookName field if non-nil, zero value otherwise.

### GetWebhookNameOk

`func (o *OutboundLogsDataInner) GetWebhookNameOk() (*string, bool)`

GetWebhookNameOk returns a tuple with the WebhookName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookName

`func (o *OutboundLogsDataInner) SetWebhookName(v string)`

SetWebhookName sets WebhookName field to given value.

### HasWebhookName

`func (o *OutboundLogsDataInner) HasWebhookName() bool`

HasWebhookName returns a boolean if a field has been set.

### GetNotificationType

`func (o *OutboundLogsDataInner) GetNotificationType() string`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *OutboundLogsDataInner) GetNotificationTypeOk() (*string, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *OutboundLogsDataInner) SetNotificationType(v string)`

SetNotificationType sets NotificationType field to given value.


### GetApplicationId

`func (o *OutboundLogsDataInner) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *OutboundLogsDataInner) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *OutboundLogsDataInner) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *OutboundLogsDataInner) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *OutboundLogsDataInner) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *OutboundLogsDataInner) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *OutboundLogsDataInner) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *OutboundLogsDataInner) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetRequest

`func (o *OutboundLogsDataInner) GetRequest() OutboundLogRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OutboundLogsDataInner) GetRequestOk() (*OutboundLogRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OutboundLogsDataInner) SetRequest(v OutboundLogRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OutboundLogsDataInner) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OutboundLogsDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutboundLogsDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutboundLogsDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProcessingTimeMs

`func (o *OutboundLogsDataInner) GetProcessingTimeMs() int64`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *OutboundLogsDataInner) GetProcessingTimeMsOk() (*int64, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *OutboundLogsDataInner) SetProcessingTimeMs(v int64)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.


### GetResponse

`func (o *OutboundLogsDataInner) GetResponse() OutboundLogResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *OutboundLogsDataInner) GetResponseOk() (*OutboundLogResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *OutboundLogsDataInner) SetResponse(v OutboundLogResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *OutboundLogsDataInner) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetError

`func (o *OutboundLogsDataInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OutboundLogsDataInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OutboundLogsDataInner) SetError(v string)`

SetError sets Error field to given value.


### GetId

`func (o *OutboundLogsDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutboundLogsDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutboundLogsDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutboundLogsDataInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


