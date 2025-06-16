# PccQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiveQi** | **int32** |  | 
**MaxBrUl** | **string** |  | 
**MaxBrDl** | **string** |  | 
**Arp** | [**Arp**](Arp.md) |  | 

## Methods

### NewPccQos

`func NewPccQos(fiveQi int32, maxBrUl string, maxBrDl string, arp Arp, ) *PccQos`

NewPccQos instantiates a new PccQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPccQosWithDefaults

`func NewPccQosWithDefaults() *PccQos`

NewPccQosWithDefaults instantiates a new PccQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiveQi

`func (o *PccQos) GetFiveQi() int32`

GetFiveQi returns the FiveQi field if non-nil, zero value otherwise.

### GetFiveQiOk

`func (o *PccQos) GetFiveQiOk() (*int32, bool)`

GetFiveQiOk returns a tuple with the FiveQi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQi

`func (o *PccQos) SetFiveQi(v int32)`

SetFiveQi sets FiveQi field to given value.


### GetMaxBrUl

`func (o *PccQos) GetMaxBrUl() string`

GetMaxBrUl returns the MaxBrUl field if non-nil, zero value otherwise.

### GetMaxBrUlOk

`func (o *PccQos) GetMaxBrUlOk() (*string, bool)`

GetMaxBrUlOk returns a tuple with the MaxBrUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBrUl

`func (o *PccQos) SetMaxBrUl(v string)`

SetMaxBrUl sets MaxBrUl field to given value.


### GetMaxBrDl

`func (o *PccQos) GetMaxBrDl() string`

GetMaxBrDl returns the MaxBrDl field if non-nil, zero value otherwise.

### GetMaxBrDlOk

`func (o *PccQos) GetMaxBrDlOk() (*string, bool)`

GetMaxBrDlOk returns a tuple with the MaxBrDl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBrDl

`func (o *PccQos) SetMaxBrDl(v string)`

SetMaxBrDl sets MaxBrDl field to given value.


### GetArp

`func (o *PccQos) GetArp() Arp`

GetArp returns the Arp field if non-nil, zero value otherwise.

### GetArpOk

`func (o *PccQos) GetArpOk() (*Arp, bool)`

GetArpOk returns a tuple with the Arp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArp

`func (o *PccQos) SetArp(v Arp)`

SetArp sets Arp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


