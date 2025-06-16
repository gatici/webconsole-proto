# PccRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** |  | 
**Flows** | [**[]PccFlow**](PccFlow.md) |  | 
**Qos** | [**PccQos**](PccQos.md) |  | 
**Precedence** | **int32** |  | 

## Methods

### NewPccRule

`func NewPccRule(ruleId string, flows []PccFlow, qos PccQos, precedence int32, ) *PccRule`

NewPccRule instantiates a new PccRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccRuleWithDefaults

`func NewPccRuleWithDefaults() *PccRule`

NewPccRuleWithDefaults instantiates a new PccRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *PccRule) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *PccRule) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *PccRule) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetFlows

`func (o *PccRule) GetFlows() []PccFlow`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *PccRule) GetFlowsOk() (*[]PccFlow, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *PccRule) SetFlows(v []PccFlow)`

SetFlows sets Flows field to given value.


### GetQos

`func (o *PccRule) GetQos() PccQos`

GetQos returns the Qos field if non-nil, zero value otherwise.

### GetQosOk

`func (o *PccRule) GetQosOk() (*PccQos, bool)`

GetQosOk returns a tuple with the Qos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQos

`func (o *PccRule) SetQos(v PccQos)`

SetQos sets Qos field to given value.


### GetPrecedence

`func (o *PccRule) GetPrecedence() int32`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *PccRule) GetPrecedenceOk() (*int32, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *PccRule) SetPrecedence(v int32)`

SetPrecedence sets Precedence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


