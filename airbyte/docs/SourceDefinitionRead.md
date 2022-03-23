# SourceDefinitionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | **string** |  | 
**Name** | **string** |  | 
**DockerRepository** | **string** |  | 
**DockerImageTag** | **string** |  | 
**DocumentationUrl** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**ReleaseStage** | Pointer to [**ReleaseStage**](ReleaseStage.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** | The date when this connector was first released, in yyyy-mm-dd format. | [optional] 

## Methods

### NewSourceDefinitionRead

`func NewSourceDefinitionRead(sourceDefinitionId string, name string, dockerRepository string, dockerImageTag string, ) *SourceDefinitionRead`

NewSourceDefinitionRead instantiates a new SourceDefinitionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDefinitionReadWithDefaults

`func NewSourceDefinitionReadWithDefaults() *SourceDefinitionRead`

NewSourceDefinitionReadWithDefaults instantiates a new SourceDefinitionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *SourceDefinitionRead) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *SourceDefinitionRead) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *SourceDefinitionRead) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.


### GetName

`func (o *SourceDefinitionRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceDefinitionRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceDefinitionRead) SetName(v string)`

SetName sets Name field to given value.


### GetDockerRepository

`func (o *SourceDefinitionRead) GetDockerRepository() string`

GetDockerRepository returns the DockerRepository field if non-nil, zero value otherwise.

### GetDockerRepositoryOk

`func (o *SourceDefinitionRead) GetDockerRepositoryOk() (*string, bool)`

GetDockerRepositoryOk returns a tuple with the DockerRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepository

`func (o *SourceDefinitionRead) SetDockerRepository(v string)`

SetDockerRepository sets DockerRepository field to given value.


### GetDockerImageTag

`func (o *SourceDefinitionRead) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *SourceDefinitionRead) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *SourceDefinitionRead) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.


### GetDocumentationUrl

`func (o *SourceDefinitionRead) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *SourceDefinitionRead) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *SourceDefinitionRead) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *SourceDefinitionRead) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### GetIcon

`func (o *SourceDefinitionRead) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *SourceDefinitionRead) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *SourceDefinitionRead) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *SourceDefinitionRead) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetReleaseStage

`func (o *SourceDefinitionRead) GetReleaseStage() ReleaseStage`

GetReleaseStage returns the ReleaseStage field if non-nil, zero value otherwise.

### GetReleaseStageOk

`func (o *SourceDefinitionRead) GetReleaseStageOk() (*ReleaseStage, bool)`

GetReleaseStageOk returns a tuple with the ReleaseStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStage

`func (o *SourceDefinitionRead) SetReleaseStage(v ReleaseStage)`

SetReleaseStage sets ReleaseStage field to given value.

### HasReleaseStage

`func (o *SourceDefinitionRead) HasReleaseStage() bool`

HasReleaseStage returns a boolean if a field has been set.

### GetReleaseDate

`func (o *SourceDefinitionRead) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *SourceDefinitionRead) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *SourceDefinitionRead) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *SourceDefinitionRead) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


