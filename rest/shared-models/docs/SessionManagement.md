# SessionManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SliceName** | **string** |  | 
**PlmnId** | [**PlmnId**](PlmnId.md) |  | 
**Snssai** | [**Snssai**](Snssai.md) |  | 
**IpDomain** | Pointer to [**[]IpDomain**](IpDomain.md) |  | [optional] 
**Upf** | Pointer to [**Upf**](Upf.md) |  | [optional] 
**GnbNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSessionManagement

`func NewSessionManagement(sliceName string, plmnId PlmnId, snssai Snssai, ) *SessionManagement`

NewSessionManagement instantiates a new SessionManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionManagementWithDefaults

`func NewSessionManagementWithDefaults() *SessionManagement`

NewSessionManagementWithDefaults instantiates a new SessionManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSliceName

`func (o *SessionManagement) GetSliceName() string`

GetSliceName returns the SliceName field if non-nil, zero value otherwise.

### GetSliceNameOk

`func (o *SessionManagement) GetSliceNameOk() (*string, bool)`

GetSliceNameOk returns a tuple with the SliceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSliceName

`func (o *SessionManagement) SetSliceName(v string)`

SetSliceName sets SliceName field to given value.


### GetPlmnId

`func (o *SessionManagement) GetPlmnId() PlmnId`

GetPlmnId returns the PlmnId field if non-nil, zero value otherwise.

### GetPlmnIdOk

`func (o *SessionManagement) GetPlmnIdOk() (*PlmnId, bool)`

GetPlmnIdOk returns a tuple with the PlmnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlmnId

`func (o *SessionManagement) SetPlmnId(v PlmnId)`

SetPlmnId sets PlmnId field to given value.


### GetSnssai

`func (o *SessionManagement) GetSnssai() Snssai`

GetSnssai returns the Snssai field if non-nil, zero value otherwise.

### GetSnssaiOk

`func (o *SessionManagement) GetSnssaiOk() (*Snssai, bool)`

GetSnssaiOk returns a tuple with the Snssai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnssai

`func (o *SessionManagement) SetSnssai(v Snssai)`

SetSnssai sets Snssai field to given value.


### GetIpDomain

`func (o *SessionManagement) GetIpDomain() []IpDomain`

GetIpDomain returns the IpDomain field if non-nil, zero value otherwise.

### GetIpDomainOk

`func (o *SessionManagement) GetIpDomainOk() (*[]IpDomain, bool)`

GetIpDomainOk returns a tuple with the IpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpDomain

`func (o *SessionManagement) SetIpDomain(v []IpDomain)`

SetIpDomain sets IpDomain field to given value.

### HasIpDomain

`func (o *SessionManagement) HasIpDomain() bool`

HasIpDomain returns a boolean if a field has been set.

### GetUpf

`func (o *SessionManagement) GetUpf() Upf`

GetUpf returns the Upf field if non-nil, zero value otherwise.

### GetUpfOk

`func (o *SessionManagement) GetUpfOk() (*Upf, bool)`

GetUpfOk returns a tuple with the Upf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpf

`func (o *SessionManagement) SetUpf(v Upf)`

SetUpf sets Upf field to given value.

### HasUpf

`func (o *SessionManagement) HasUpf() bool`

HasUpf returns a boolean if a field has been set.

### GetGnbNames

`func (o *SessionManagement) GetGnbNames() []string`

GetGnbNames returns the GnbNames field if non-nil, zero value otherwise.

### GetGnbNamesOk

`func (o *SessionManagement) GetGnbNamesOk() (*[]string, bool)`

GetGnbNamesOk returns a tuple with the GnbNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnbNames

`func (o *SessionManagement) SetGnbNames(v []string)`

SetGnbNames sets GnbNames field to given value.

### HasGnbNames

`func (o *SessionManagement) HasGnbNames() bool`

HasGnbNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


