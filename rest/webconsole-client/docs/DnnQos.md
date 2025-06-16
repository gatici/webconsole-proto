# DnnQos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnnName** | **string** |  | 
**MbrUplink** | **string** |  | 
**MbrDownlink** | **string** |  | 
**FiveQi** | Pointer to **int32** |  | [optional] 
**ArpPriorityLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnnQos

`func NewDnnQos(dnnName string, mbrUplink string, mbrDownlink string, ) *DnnQos`

NewDnnQos instantiates a new DnnQos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnnQosWithDefaults

`func NewDnnQosWithDefaults() *DnnQos`

NewDnnQosWithDefaults instantiates a new DnnQos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnnName

`func (o *DnnQos) GetDnnName() string`

GetDnnName returns the DnnName field if non-nil, zero value otherwise.

### GetDnnNameOk

`func (o *DnnQos) GetDnnNameOk() (*string, bool)`

GetDnnNameOk returns a tuple with the DnnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnName

`func (o *DnnQos) SetDnnName(v string)`

SetDnnName sets DnnName field to given value.


### GetMbrUplink

`func (o *DnnQos) GetMbrUplink() string`

GetMbrUplink returns the MbrUplink field if non-nil, zero value otherwise.

### GetMbrUplinkOk

`func (o *DnnQos) GetMbrUplinkOk() (*string, bool)`

GetMbrUplinkOk returns a tuple with the MbrUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbrUplink

`func (o *DnnQos) SetMbrUplink(v string)`

SetMbrUplink sets MbrUplink field to given value.


### GetMbrDownlink

`func (o *DnnQos) GetMbrDownlink() string`

GetMbrDownlink returns the MbrDownlink field if non-nil, zero value otherwise.

### GetMbrDownlinkOk

`func (o *DnnQos) GetMbrDownlinkOk() (*string, bool)`

GetMbrDownlinkOk returns a tuple with the MbrDownlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMbrDownlink

`func (o *DnnQos) SetMbrDownlink(v string)`

SetMbrDownlink sets MbrDownlink field to given value.


### GetFiveQi

`func (o *DnnQos) GetFiveQi() int32`

GetFiveQi returns the FiveQi field if non-nil, zero value otherwise.

### GetFiveQiOk

`func (o *DnnQos) GetFiveQiOk() (*int32, bool)`

GetFiveQiOk returns a tuple with the FiveQi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiveQi

`func (o *DnnQos) SetFiveQi(v int32)`

SetFiveQi sets FiveQi field to given value.

### HasFiveQi

`func (o *DnnQos) HasFiveQi() bool`

HasFiveQi returns a boolean if a field has been set.

### GetArpPriorityLevel

`func (o *DnnQos) GetArpPriorityLevel() int32`

GetArpPriorityLevel returns the ArpPriorityLevel field if non-nil, zero value otherwise.

### GetArpPriorityLevelOk

`func (o *DnnQos) GetArpPriorityLevelOk() (*int32, bool)`

GetArpPriorityLevelOk returns a tuple with the ArpPriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArpPriorityLevel

`func (o *DnnQos) SetArpPriorityLevel(v int32)`

SetArpPriorityLevel sets ArpPriorityLevel field to given value.

### HasArpPriorityLevel

`func (o *DnnQos) HasArpPriorityLevel() bool`

HasArpPriorityLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


