# SourceDefinitionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**DockerRepository** | **string** |  | 
**DockerImageTag** | **string** |  | 
**DocumentationUrl** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 

## Methods

### NewSourceDefinitionCreate

`func NewSourceDefinitionCreate(name string, dockerRepository string, dockerImageTag string, documentationUrl string, ) *SourceDefinitionCreate`

NewSourceDefinitionCreate instantiates a new SourceDefinitionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDefinitionCreateWithDefaults

`func NewSourceDefinitionCreateWithDefaults() *SourceDefinitionCreate`

NewSourceDefinitionCreateWithDefaults instantiates a new SourceDefinitionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourceDefinitionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceDefinitionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceDefinitionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDockerRepository

`func (o *SourceDefinitionCreate) GetDockerRepository() string`

GetDockerRepository returns the DockerRepository field if non-nil, zero value otherwise.

### GetDockerRepositoryOk

`func (o *SourceDefinitionCreate) GetDockerRepositoryOk() (*string, bool)`

GetDockerRepositoryOk returns a tuple with the DockerRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepository

`func (o *SourceDefinitionCreate) SetDockerRepository(v string)`

SetDockerRepository sets DockerRepository field to given value.


### GetDockerImageTag

`func (o *SourceDefinitionCreate) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *SourceDefinitionCreate) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *SourceDefinitionCreate) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.


### GetDocumentationUrl

`func (o *SourceDefinitionCreate) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *SourceDefinitionCreate) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *SourceDefinitionCreate) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.


### GetIcon

`func (o *SourceDefinitionCreate) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *SourceDefinitionCreate) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *SourceDefinitionCreate) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *SourceDefinitionCreate) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


