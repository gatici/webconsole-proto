# PccFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Direction** | [**Direction**](Direction.md) |  | 
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 

## Methods

### NewPccFlow

`func NewPccFlow(description string, direction Direction, ) *PccFlow`

NewPccFlow instantiates a new PccFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccFlowWithDefaults

`func NewPccFlowWithDefaults() *PccFlow`

NewPccFlowWithDefaults instantiates a new PccFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PccFlow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PccFlow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PccFlow) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDirection

`func (o *PccFlow) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *PccFlow) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *PccFlow) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetStatus

`func (o *PccFlow) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PccFlow) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PccFlow) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PccFlow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


