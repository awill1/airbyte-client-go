# SourceDefinitionSpecificationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | **string** |  | 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**ConnectionSpecification** | Pointer to **map[string]interface{}** | The specification for what values are required to configure the sourceDefinition. | [optional] 
**AuthSpecification** | Pointer to [**AuthSpecification**](AuthSpecification.md) |  | [optional] 
**AdvancedAuth** | Pointer to [**AdvancedAuth**](AdvancedAuth.md) |  | [optional] 
**JobInfo** | [**SynchronousJobRead**](SynchronousJobRead.md) |  | 

## Methods

### NewSourceDefinitionSpecificationRead

`func NewSourceDefinitionSpecificationRead(sourceDefinitionId string, jobInfo SynchronousJobRead, ) *SourceDefinitionSpecificationRead`

NewSourceDefinitionSpecificationRead instantiates a new SourceDefinitionSpecificationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDefinitionSpecificationReadWithDefaults

`func NewSourceDefinitionSpecificationReadWithDefaults() *SourceDefinitionSpecificationRead`

NewSourceDefinitionSpecificationReadWithDefaults instantiates a new SourceDefinitionSpecificationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *SourceDefinitionSpecificationRead) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *SourceDefinitionSpecificationRead) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *SourceDefinitionSpecificationRead) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.


### GetDocumentationUrl

`func (o *SourceDefinitionSpecificationRead) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *SourceDefinitionSpecificationRead) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *SourceDefinitionSpecificationRead) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *SourceDefinitionSpecificationRead) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetConnectionSpecification

`func (o *SourceDefinitionSpecificationRead) GetConnectionSpecification() map[string]interface{}`

GetConnectionSpecification returns the ConnectionSpecification field if non-nil, zero value otherwise.

### GetConnectionSpecificationOk

`func (o *SourceDefinitionSpecificationRead) GetConnectionSpecificationOk() (*map[string]interface{}, bool)`

GetConnectionSpecificationOk returns a tuple with the ConnectionSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSpecification

`func (o *SourceDefinitionSpecificationRead) SetConnectionSpecification(v map[string]interface{})`

SetConnectionSpecification sets ConnectionSpecification field to given value.

### HasConnectionSpecification

`func (o *SourceDefinitionSpecificationRead) HasConnectionSpecification() bool`

HasConnectionSpecification returns a boolean if a field has been set.

### GetAuthSpecification

`func (o *SourceDefinitionSpecificationRead) GetAuthSpecification() AuthSpecification`

GetAuthSpecification returns the AuthSpecification field if non-nil, zero value otherwise.

### GetAuthSpecificationOk

`func (o *SourceDefinitionSpecificationRead) GetAuthSpecificationOk() (*AuthSpecification, bool)`

GetAuthSpecificationOk returns a tuple with the AuthSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSpecification

`func (o *SourceDefinitionSpecificationRead) SetAuthSpecification(v AuthSpecification)`

SetAuthSpecification sets AuthSpecification field to given value.

### HasAuthSpecification

`func (o *SourceDefinitionSpecificationRead) HasAuthSpecification() bool`

HasAuthSpecification returns a boolean if a field has been set.

### GetAdvancedAuth

`func (o *SourceDefinitionSpecificationRead) GetAdvancedAuth() AdvancedAuth`

GetAdvancedAuth returns the AdvancedAuth field if non-nil, zero value otherwise.

### GetAdvancedAuthOk

`func (o *SourceDefinitionSpecificationRead) GetAdvancedAuthOk() (*AdvancedAuth, bool)`

GetAdvancedAuthOk returns a tuple with the AdvancedAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAuth

`func (o *SourceDefinitionSpecificationRead) SetAdvancedAuth(v AdvancedAuth)`

SetAdvancedAuth sets AdvancedAuth field to given value.

### HasAdvancedAuth

`func (o *SourceDefinitionSpecificationRead) HasAdvancedAuth() bool`

HasAdvancedAuth returns a boolean if a field has been set.

### GetJobInfo

`func (o *SourceDefinitionSpecificationRead) GetJobInfo() SynchronousJobRead`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *SourceDefinitionSpecificationRead) GetJobInfoOk() (*SynchronousJobRead, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *SourceDefinitionSpecificationRead) SetJobInfo(v SynchronousJobRead)`

SetJobInfo sets JobInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


