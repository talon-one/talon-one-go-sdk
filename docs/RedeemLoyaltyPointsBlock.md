# RedeemLoyaltyPointsBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Program** | [**RedeemLoyaltyPointsBlock1Program**](RedeemLoyaltyPointsBlock1Program.md) |  | 
**Subledger** | **string** | The name of the subledger to deduct points from. Can be empty if this block deducts from the loyalty program&#39;s main ledger instead of a subledger. | 
**Value** | [**RedeemLoyaltyPointsBlock1Value**](RedeemLoyaltyPointsBlock1Value.md) |  | 
**Name** | Pointer to **string** | A custom description recorded as the reason for the point deduction. | [optional] 
**OnFailure** | Pointer to [**[]Block**](Block.md) | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewRedeemLoyaltyPointsBlock

`func NewRedeemLoyaltyPointsBlock(type_ string, program RedeemLoyaltyPointsBlock1Program, subledger string, value RedeemLoyaltyPointsBlock1Value, ) *RedeemLoyaltyPointsBlock`

NewRedeemLoyaltyPointsBlock instantiates a new RedeemLoyaltyPointsBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedeemLoyaltyPointsBlockWithDefaults

`func NewRedeemLoyaltyPointsBlockWithDefaults() *RedeemLoyaltyPointsBlock`

NewRedeemLoyaltyPointsBlockWithDefaults instantiates a new RedeemLoyaltyPointsBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedeemLoyaltyPointsBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedeemLoyaltyPointsBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedeemLoyaltyPointsBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RedeemLoyaltyPointsBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RedeemLoyaltyPointsBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RedeemLoyaltyPointsBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RedeemLoyaltyPointsBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *RedeemLoyaltyPointsBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RedeemLoyaltyPointsBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RedeemLoyaltyPointsBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RedeemLoyaltyPointsBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProgram

`func (o *RedeemLoyaltyPointsBlock) GetProgram() RedeemLoyaltyPointsBlock1Program`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *RedeemLoyaltyPointsBlock) GetProgramOk() (*RedeemLoyaltyPointsBlock1Program, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *RedeemLoyaltyPointsBlock) SetProgram(v RedeemLoyaltyPointsBlock1Program)`

SetProgram sets Program field to given value.


### GetSubledger

`func (o *RedeemLoyaltyPointsBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *RedeemLoyaltyPointsBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *RedeemLoyaltyPointsBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetValue

`func (o *RedeemLoyaltyPointsBlock) GetValue() RedeemLoyaltyPointsBlock1Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RedeemLoyaltyPointsBlock) GetValueOk() (*RedeemLoyaltyPointsBlock1Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RedeemLoyaltyPointsBlock) SetValue(v RedeemLoyaltyPointsBlock1Value)`

SetValue sets Value field to given value.


### GetName

`func (o *RedeemLoyaltyPointsBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedeemLoyaltyPointsBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedeemLoyaltyPointsBlock) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RedeemLoyaltyPointsBlock) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnFailure

`func (o *RedeemLoyaltyPointsBlock) GetOnFailure() []Block`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *RedeemLoyaltyPointsBlock) GetOnFailureOk() (*[]Block, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *RedeemLoyaltyPointsBlock) SetOnFailure(v []Block)`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *RedeemLoyaltyPointsBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


