# ResourceRequirements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuRequest** | Pointer to **string** |  | [optional] 
**CpuLimit** | Pointer to **string** |  | [optional] 
**MemoryRequest** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceRequirements

`func NewResourceRequirements() *ResourceRequirements`

NewResourceRequirements instantiates a new ResourceRequirements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRequirementsWithDefaults

`func NewResourceRequirementsWithDefaults() *ResourceRequirements`

NewResourceRequirementsWithDefaults instantiates a new ResourceRequirements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuRequest

`func (o *ResourceRequirements) GetCpuRequest() string`

GetCpuRequest returns the CpuRequest field if non-nil, zero value otherwise.

### GetCpuRequestOk

`func (o *ResourceRequirements) GetCpuRequestOk() (*string, bool)`

GetCpuRequestOk returns a tuple with the CpuRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuRequest

`func (o *ResourceRequirements) SetCpuRequest(v string)`

SetCpuRequest sets CpuRequest field to given value.

### HasCpuRequest

`func (o *ResourceRequirements) HasCpuRequest() bool`

HasCpuRequest returns a boolean if a field has been set.

### GetCpuLimit

`func (o *ResourceRequirements) GetCpuLimit() string`

GetCpuLimit returns the CpuLimit field if non-nil, zero value otherwise.

### GetCpuLimitOk

`func (o *ResourceRequirements) GetCpuLimitOk() (*string, bool)`

GetCpuLimitOk returns a tuple with the CpuLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLimit

`func (o *ResourceRequirements) SetCpuLimit(v string)`

SetCpuLimit sets CpuLimit field to given value.

### HasCpuLimit

`func (o *ResourceRequirements) HasCpuLimit() bool`

HasCpuLimit returns a boolean if a field has been set.

### GetMemoryRequest

`func (o *ResourceRequirements) GetMemoryRequest() string`

GetMemoryRequest returns the MemoryRequest field if non-nil, zero value otherwise.

### GetMemoryRequestOk

`func (o *ResourceRequirements) GetMemoryRequestOk() (*string, bool)`

GetMemoryRequestOk returns a tuple with the MemoryRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequest

`func (o *ResourceRequirements) SetMemoryRequest(v string)`

SetMemoryRequest sets MemoryRequest field to given value.

### HasMemoryRequest

`func (o *ResourceRequirements) HasMemoryRequest() bool`

HasMemoryRequest returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *ResourceRequirements) GetMemoryLimit() string`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *ResourceRequirements) GetMemoryLimitOk() (*string, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *ResourceRequirements) SetMemoryLimit(v string)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *ResourceRequirements) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


