# AccessAndMobility

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**Tacs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAccessAndMobility

`func NewAccessAndMobility(plmnId PlmnId, snssai Snssai, ) *AccessAndMobility`

NewAccessAndMobility instantiates a new AccessAndMobility object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessAndMobilityWithDefaults

`func NewAccessAndMobilityWithDefaults() *AccessAndMobility`

NewAccessAndMobilityWithDefaults instantiates a new AccessAndMobility object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlmnId

`func (o *AccessAndMobility) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *AccessAndMobility) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *AccessAndMobility) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSnssai

`func (o *AccessAndMobility) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *AccessAndMobility) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *AccessAndMobility) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetTacs

`func (o *AccessAndMobility) GetTacs() []string`

GetTacs returns the Tacs field if non-nil, zero value otherwise.

### GetTacsOk

`func (o *AccessAndMobility) GetTacsOk() (*[]string, bool)`

GetTacsOk returns a tuple with the Tacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTacs

`func (o *AccessAndMobility) SetTacs(v []string)`

SetTacs sets Tacs field to given value.

### HasTacs

`func (o *AccessAndMobility) HasTacs() bool`

HasTacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


