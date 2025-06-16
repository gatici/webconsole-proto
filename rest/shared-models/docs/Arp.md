# Arp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityLevel** | **int32** |  | 
**PreemptCap** | [**PreemptCap**](PreemptCap.md) |  | 
**PreemptVuln** | [**PreemptVuln**](PreemptVuln.md) |  | 

## Methods

### NewArp

`func NewArp(priorityLevel int32, preemptCap PreemptCap, preemptVuln PreemptVuln, ) *Arp`

NewArp instantiates a new Arp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArpWithDefaults

`func NewArpWithDefaults() *Arp`

NewArpWithDefaults instantiates a new Arp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityLevel

`func (o *Arp) GetPriorityLevel() int32`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *Arp) GetPriorityLevelOk() (*int32, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *Arp) SetPriorityLevel(v int32)`

SetPriorityLevel sets PriorityLevel field to given value.


### GetPreemptCap

`func (o *Arp) GetPreemptCap() PreemptCap`

GetPreemptCap returns the PreemptCap field if non-nil, zero value otherwise.

### GetPreemptCapOk

`func (o *Arp) GetPreemptCapOk() (*PreemptCap, bool)`

GetPreemptCapOk returns a tuple with the PreemptCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptCap

`func (o *Arp) SetPreemptCap(v PreemptCap)`

SetPreemptCap sets PreemptCap field to given value.


### GetPreemptVuln

`func (o *Arp) GetPreemptVuln() PreemptVuln`

GetPreemptVuln returns the PreemptVuln field if non-nil, zero value otherwise.

### GetPreemptVulnOk

`func (o *Arp) GetPreemptVulnOk() (*PreemptVuln, bool)`

GetPreemptVulnOk returns a tuple with the PreemptVuln field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptVuln

`func (o *Arp) SetPreemptVuln(v PreemptVuln)`

SetPreemptVuln sets PreemptVuln field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


