# DestinationDefinitionSpecificationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**ConnectionSpecification** | Pointer to **interface{}** | The specification for what values are required to configure the destinationDefinition. | [optional] 
**AuthSpecification** | Pointer to [**AuthSpecification**](AuthSpecification.md) |  | [optional] 
**AdvancedAuth** | Pointer to [**AdvancedAuth**](AdvancedAuth.md) |  | [optional] 
**JobInfo** | [**SynchronousJobRead**](SynchronousJobRead.md) |  | 
**SupportedDestinationSyncModes** | Pointer to [**[]DestinationSyncMode**](DestinationSyncMode.md) |  | [optional] 
**SupportsDbt** | Pointer to **bool** |  | [optional] 
**SupportsNormalization** | Pointer to **bool** |  | [optional] 

## Methods

### NewDestinationDefinitionSpecificationRead

`func NewDestinationDefinitionSpecificationRead(destinationDefinitionId string, jobInfo SynchronousJobRead, ) *DestinationDefinitionSpecificationRead`

NewDestinationDefinitionSpecificationRead instantiates a new DestinationDefinitionSpecificationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDefinitionSpecificationReadWithDefaults

`func NewDestinationDefinitionSpecificationReadWithDefaults() *DestinationDefinitionSpecificationRead`

NewDestinationDefinitionSpecificationReadWithDefaults instantiates a new DestinationDefinitionSpecificationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationDefinitionSpecificationRead) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationDefinitionSpecificationRead) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationDefinitionSpecificationRead) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetDocumentationUrl

`func (o *DestinationDefinitionSpecificationRead) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *DestinationDefinitionSpecificationRead) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *DestinationDefinitionSpecificationRead) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *DestinationDefinitionSpecificationRead) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetConnectionSpecification

`func (o *DestinationDefinitionSpecificationRead) GetConnectionSpecification() interface{}`

GetConnectionSpecification returns the ConnectionSpecification field if non-nil, zero value otherwise.

### GetConnectionSpecificationOk

`func (o *DestinationDefinitionSpecificationRead) GetConnectionSpecificationOk() (*interface{}, bool)`

GetConnectionSpecificationOk returns a tuple with the ConnectionSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSpecification

`func (o *DestinationDefinitionSpecificationRead) SetConnectionSpecification(v interface{})`

SetConnectionSpecification sets ConnectionSpecification field to given value.

### HasConnectionSpecification

`func (o *DestinationDefinitionSpecificationRead) HasConnectionSpecification() bool`

HasConnectionSpecification returns a boolean if a field has been set.

### SetConnectionSpecificationNil

`func (o *DestinationDefinitionSpecificationRead) SetConnectionSpecificationNil(b bool)`

 SetConnectionSpecificationNil sets the value for ConnectionSpecification to be an explicit nil

### UnsetConnectionSpecification
`func (o *DestinationDefinitionSpecificationRead) UnsetConnectionSpecification()`

UnsetConnectionSpecification ensures that no value is present for ConnectionSpecification, not even an explicit nil
### GetAuthSpecification

`func (o *DestinationDefinitionSpecificationRead) GetAuthSpecification() AuthSpecification`

GetAuthSpecification returns the AuthSpecification field if non-nil, zero value otherwise.

### GetAuthSpecificationOk

`func (o *DestinationDefinitionSpecificationRead) GetAuthSpecificationOk() (*AuthSpecification, bool)`

GetAuthSpecificationOk returns a tuple with the AuthSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSpecification

`func (o *DestinationDefinitionSpecificationRead) SetAuthSpecification(v AuthSpecification)`

SetAuthSpecification sets AuthSpecification field to given value.

### HasAuthSpecification

`func (o *DestinationDefinitionSpecificationRead) HasAuthSpecification() bool`

HasAuthSpecification returns a boolean if a field has been set.

### GetAdvancedAuth

`func (o *DestinationDefinitionSpecificationRead) GetAdvancedAuth() AdvancedAuth`

GetAdvancedAuth returns the AdvancedAuth field if non-nil, zero value otherwise.

### GetAdvancedAuthOk

`func (o *DestinationDefinitionSpecificationRead) GetAdvancedAuthOk() (*AdvancedAuth, bool)`

GetAdvancedAuthOk returns a tuple with the AdvancedAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAuth

`func (o *DestinationDefinitionSpecificationRead) SetAdvancedAuth(v AdvancedAuth)`

SetAdvancedAuth sets AdvancedAuth field to given value.

### HasAdvancedAuth

`func (o *DestinationDefinitionSpecificationRead) HasAdvancedAuth() bool`

HasAdvancedAuth returns a boolean if a field has been set.

### GetJobInfo

`func (o *DestinationDefinitionSpecificationRead) GetJobInfo() SynchronousJobRead`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *DestinationDefinitionSpecificationRead) GetJobInfoOk() (*SynchronousJobRead, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *DestinationDefinitionSpecificationRead) SetJobInfo(v SynchronousJobRead)`

SetJobInfo sets JobInfo field to given value.


### GetSupportedDestinationSyncModes

`func (o *DestinationDefinitionSpecificationRead) GetSupportedDestinationSyncModes() []DestinationSyncMode`

GetSupportedDestinationSyncModes returns the SupportedDestinationSyncModes field if non-nil, zero value otherwise.

### GetSupportedDestinationSyncModesOk

`func (o *DestinationDefinitionSpecificationRead) GetSupportedDestinationSyncModesOk() (*[]DestinationSyncMode, bool)`

GetSupportedDestinationSyncModesOk returns a tuple with the SupportedDestinationSyncModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDestinationSyncModes

`func (o *DestinationDefinitionSpecificationRead) SetSupportedDestinationSyncModes(v []DestinationSyncMode)`

SetSupportedDestinationSyncModes sets SupportedDestinationSyncModes field to given value.

### HasSupportedDestinationSyncModes

`func (o *DestinationDefinitionSpecificationRead) HasSupportedDestinationSyncModes() bool`

HasSupportedDestinationSyncModes returns a boolean if a field has been set.

### GetSupportsDbt

`func (o *DestinationDefinitionSpecificationRead) GetSupportsDbt() bool`

GetSupportsDbt returns the SupportsDbt field if non-nil, zero value otherwise.

### GetSupportsDbtOk

`func (o *DestinationDefinitionSpecificationRead) GetSupportsDbtOk() (*bool, bool)`

GetSupportsDbtOk returns a tuple with the SupportsDbt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsDbt

`func (o *DestinationDefinitionSpecificationRead) SetSupportsDbt(v bool)`

SetSupportsDbt sets SupportsDbt field to given value.

### HasSupportsDbt

`func (o *DestinationDefinitionSpecificationRead) HasSupportsDbt() bool`

HasSupportsDbt returns a boolean if a field has been set.

### GetSupportsNormalization

`func (o *DestinationDefinitionSpecificationRead) GetSupportsNormalization() bool`

GetSupportsNormalization returns the SupportsNormalization field if non-nil, zero value otherwise.

### GetSupportsNormalizationOk

`func (o *DestinationDefinitionSpecificationRead) GetSupportsNormalizationOk() (*bool, bool)`

GetSupportsNormalizationOk returns a tuple with the SupportsNormalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsNormalization

`func (o *DestinationDefinitionSpecificationRead) SetSupportsNormalization(v bool)`

SetSupportsNormalization sets SupportsNormalization field to given value.

### HasSupportsNormalization

`func (o *DestinationDefinitionSpecificationRead) HasSupportsNormalization() bool`

HasSupportsNormalization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


