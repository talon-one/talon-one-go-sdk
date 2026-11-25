# UpdateLoyaltyProgramMembershipBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | **string** | The action to perform. &#x60;join&#x60; adds the customer to the loyalty program. | 
**Program** | [**UpdateLoyaltyProgramMembershipBlock1Program**](UpdateLoyaltyProgramMembershipBlock1Program.md) |  | 
**JoinDate** | **interface{}** | The date and time the customer joins the program. Cannot be in the future. Accepts one of: - An absolute timestamp in RFC 3339 format (e.g. &#x60;2026-01-15T10:00:00Z&#x60;). - A reference to a value, such as &#x60;{{$Now}}&#x60; or an attribute. - A relative time: an object with a &#x60;baseline&#x60; time, a &#x60;duration&#x60;, and an &#x60;operator&#x60;   that is &#x60;laterBy&#x60; or &#x60;earlierBy&#x60;. - A rounded time: an object with a &#x60;baseline&#x60; time, a &#x60;unit&#x60; (such as &#x60;day&#x60;, &#x60;hour&#x60; or   &#x60;month&#x60;), and a &#x60;direction&#x60; that is &#x60;start&#x60; or &#x60;end&#x60;. The baseline must be a relative   time whose duration uses the same unit.  See the examples for each form. | 

## Methods

### NewUpdateLoyaltyProgramMembershipBlock

`func NewUpdateLoyaltyProgramMembershipBlock(type_ string, operator string, program UpdateLoyaltyProgramMembershipBlock1Program, joinDate interface{}, ) *UpdateLoyaltyProgramMembershipBlock`

NewUpdateLoyaltyProgramMembershipBlock instantiates a new UpdateLoyaltyProgramMembershipBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoyaltyProgramMembershipBlockWithDefaults

`func NewUpdateLoyaltyProgramMembershipBlockWithDefaults() *UpdateLoyaltyProgramMembershipBlock`

NewUpdateLoyaltyProgramMembershipBlockWithDefaults instantiates a new UpdateLoyaltyProgramMembershipBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateLoyaltyProgramMembershipBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoyaltyProgramMembershipBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoyaltyProgramMembershipBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateLoyaltyProgramMembershipBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateLoyaltyProgramMembershipBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *UpdateLoyaltyProgramMembershipBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateLoyaltyProgramMembershipBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateLoyaltyProgramMembershipBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateLoyaltyProgramMembershipBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateLoyaltyProgramMembershipBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetProgram

`func (o *UpdateLoyaltyProgramMembershipBlock) GetProgram() UpdateLoyaltyProgramMembershipBlock1Program`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetProgramOk() (*UpdateLoyaltyProgramMembershipBlock1Program, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *UpdateLoyaltyProgramMembershipBlock) SetProgram(v UpdateLoyaltyProgramMembershipBlock1Program)`

SetProgram sets Program field to given value.


### GetJoinDate

`func (o *UpdateLoyaltyProgramMembershipBlock) GetJoinDate() interface{}`

GetJoinDate returns the JoinDate field if non-nil, zero value otherwise.

### GetJoinDateOk

`func (o *UpdateLoyaltyProgramMembershipBlock) GetJoinDateOk() (*interface{}, bool)`

GetJoinDateOk returns a tuple with the JoinDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinDate

`func (o *UpdateLoyaltyProgramMembershipBlock) SetJoinDate(v interface{})`

SetJoinDate sets JoinDate field to given value.


### SetJoinDateNil

`func (o *UpdateLoyaltyProgramMembershipBlock) SetJoinDateNil(b bool)`

 SetJoinDateNil sets the value for JoinDate to be an explicit nil

### UnsetJoinDate
`func (o *UpdateLoyaltyProgramMembershipBlock) UnsetJoinDate()`

UnsetJoinDate ensures that no value is present for JoinDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


