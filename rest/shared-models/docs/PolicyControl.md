# PolicyControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**DnnQos** | Pointer to [**[]DnnQos**](DnnQos.md) |  | [optional] 
**PccRules** | [**[]PccRule**](PccRule.md) |  | 

## Methods

### NewPolicyControl

`func NewPolicyControl(plmnId PlmnId, snssai Snssai, pccRules []PccRule, ) *PolicyControl`

NewPolicyControl instantiates a new PolicyControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyControlWithDefaults

`func NewPolicyControlWithDefaults() *PolicyControl`

NewPolicyControlWithDefaults instantiates a new PolicyControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *PolicyControl) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *PolicyControl) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *PolicyControl) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSnssai

`func (o *PolicyControl) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *PolicyControl) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *PolicyControl) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetDnnQos

`func (o *PolicyControl) GetDnnQos() []DnnQos`

GetDnnQos returns the DnnQos field if non-nil, zero value otherwise.

### GetDnnQosOk

`func (o *PolicyControl) GetDnnQosOk() (*[]DnnQos, bool)`

GetDnnQosOk returns a tuple with the DnnQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnQos

`func (o *PolicyControl) SetDnnQos(v []DnnQos)`

SetDnnQos sets DnnQos field to given value.

### HasDnnQos

`func (o *PolicyControl) HasDnnQos() bool`

HasDnnQos returns a boolean if a field has been set.

### GetPccRules

`func (o *PolicyControl) GetPccRules() []PccRule`

GetPccRules returns the PccRules field if non-nil, zero value otherwise.

### GetPccRulesOk

`func (o *PolicyControl) GetPccRulesOk() (*[]PccRule, bool)`

GetPccRulesOk returns a tuple with the PccRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPccRules

`func (o *PolicyControl) SetPccRules(v []PccRule)`

SetPccRules sets PccRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


