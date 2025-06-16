# IpDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnnName** | **string** |  | 
**DnsIpv4** | **string** |  | 
**UeSubnet** | **string** |  | 
**Mtu** | **int32** |  | 

## Methods

### NewIpDomain

`func NewIpDomain(dnnName string, dnsIpv4 string, ueSubnet string, mtu int32, ) *IpDomain`

NewIpDomain instantiates a new IpDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpDomainWithDefaults

`func NewIpDomainWithDefaults() *IpDomain`

NewIpDomainWithDefaults instantiates a new IpDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnnName

`func (o *IpDomain) GetDnnName() string`

GetDnnName returns the DnnName field if non-nil, zero value otherwise.

### GetDnnNameOk

`func (o *IpDomain) GetDnnNameOk() (*string, bool)`

GetDnnNameOk returns a tuple with the DnnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnnName

`func (o *IpDomain) SetDnnName(v string)`

SetDnnName sets DnnName field to given value.


### GetDnsIpv4

`func (o *IpDomain) GetDnsIpv4() string`

GetDnsIpv4 returns the DnsIpv4 field if non-nil, zero value otherwise.

### GetDnsIpv4Ok

`func (o *IpDomain) GetDnsIpv4Ok() (*string, bool)`

GetDnsIpv4Ok returns a tuple with the DnsIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIpv4

`func (o *IpDomain) SetDnsIpv4(v string)`

SetDnsIpv4 sets DnsIpv4 field to given value.


### GetUeSubnet

`func (o *IpDomain) GetUeSubnet() string`

GetUeSubnet returns the UeSubnet field if non-nil, zero value otherwise.

### GetUeSubnetOk

`func (o *IpDomain) GetUeSubnetOk() (*string, bool)`

GetUeSubnetOk returns a tuple with the UeSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUeSubnet

`func (o *IpDomain) SetUeSubnet(v string)`

SetUeSubnet sets UeSubnet field to given value.


### GetMtu

`func (o *IpDomain) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *IpDomain) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *IpDomain) SetMtu(v int32)`

SetMtu sets Mtu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


