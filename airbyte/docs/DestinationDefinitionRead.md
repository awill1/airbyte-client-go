# DestinationDefinitionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**Name** | **string** |  | 
**DockerRepository** | **string** |  | 
**DockerImageTag** | **string** |  | 
**DocumentationUrl** | **string** |  | 
**Icon** | Pointer to **string** |  | [optional] 
**ReleaseStage** | Pointer to [**ReleaseStage**](ReleaseStage.md) |  | [optional] 
**ReleaseDate** | Pointer to **string** | The date when this connector was first released, in yyyy-mm-dd format. | [optional] 

## Methods

### NewDestinationDefinitionRead

`func NewDestinationDefinitionRead(destinationDefinitionId string, name string, dockerRepository string, dockerImageTag string, documentationUrl string, ) *DestinationDefinitionRead`

NewDestinationDefinitionRead instantiates a new DestinationDefinitionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDefinitionReadWithDefaults

`func NewDestinationDefinitionReadWithDefaults() *DestinationDefinitionRead`

NewDestinationDefinitionReadWithDefaults instantiates a new DestinationDefinitionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationDefinitionRead) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationDefinitionRead) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationDefinitionRead) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetName

`func (o *DestinationDefinitionRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationDefinitionRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationDefinitionRead) SetName(v string)`

SetName sets Name field to given value.


### GetDockerRepository

`func (o *DestinationDefinitionRead) GetDockerRepository() string`

GetDockerRepository returns the DockerRepository field if non-nil, zero value otherwise.

### GetDockerRepositoryOk

`func (o *DestinationDefinitionRead) GetDockerRepositoryOk() (*string, bool)`

GetDockerRepositoryOk returns a tuple with the DockerRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerRepository

`func (o *DestinationDefinitionRead) SetDockerRepository(v string)`

SetDockerRepository sets DockerRepository field to given value.


### GetDockerImageTag

`func (o *DestinationDefinitionRead) GetDockerImageTag() string`

GetDockerImageTag returns the DockerImageTag field if non-nil, zero value otherwise.

### GetDockerImageTagOk

`func (o *DestinationDefinitionRead) GetDockerImageTagOk() (*string, bool)`

GetDockerImageTagOk returns a tuple with the DockerImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImageTag

`func (o *DestinationDefinitionRead) SetDockerImageTag(v string)`

SetDockerImageTag sets DockerImageTag field to given value.


### GetDocumentationUrl

`func (o *DestinationDefinitionRead) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *DestinationDefinitionRead) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *DestinationDefinitionRead) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.


### GetIcon

`func (o *DestinationDefinitionRead) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *DestinationDefinitionRead) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *DestinationDefinitionRead) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *DestinationDefinitionRead) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetReleaseStage

`func (o *DestinationDefinitionRead) GetReleaseStage() ReleaseStage`

GetReleaseStage returns the ReleaseStage field if non-nil, zero value otherwise.

### GetReleaseStageOk

`func (o *DestinationDefinitionRead) GetReleaseStageOk() (*ReleaseStage, bool)`

GetReleaseStageOk returns a tuple with the ReleaseStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStage

`func (o *DestinationDefinitionRead) SetReleaseStage(v ReleaseStage)`

SetReleaseStage sets ReleaseStage field to given value.

### HasReleaseStage

`func (o *DestinationDefinitionRead) HasReleaseStage() bool`

HasReleaseStage returns a boolean if a field has been set.

### GetReleaseDate

`func (o *DestinationDefinitionRead) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *DestinationDefinitionRead) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *DestinationDefinitionRead) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *DestinationDefinitionRead) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


