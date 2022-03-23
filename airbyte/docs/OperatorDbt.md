# OperatorDbt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepoUrl** | **string** |  | 
**GitRepoBranch** | Pointer to **string** |  | [optional] 
**DockerImage** | Pointer to **string** |  | [optional] 
**DbtArguments** | Pointer to **string** |  | [optional] 

## Methods

### NewOperatorDbt

`func NewOperatorDbt(gitRepoUrl string, ) *OperatorDbt`

NewOperatorDbt instantiates a new OperatorDbt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorDbtWithDefaults

`func NewOperatorDbtWithDefaults() *OperatorDbt`

NewOperatorDbtWithDefaults instantiates a new OperatorDbt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepoUrl

`func (o *OperatorDbt) GetGitRepoUrl() string`

GetGitRepoUrl returns the GitRepoUrl field if non-nil, zero value otherwise.

### GetGitRepoUrlOk

`func (o *OperatorDbt) GetGitRepoUrlOk() (*string, bool)`

GetGitRepoUrlOk returns a tuple with the GitRepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoUrl

`func (o *OperatorDbt) SetGitRepoUrl(v string)`

SetGitRepoUrl sets GitRepoUrl field to given value.


### GetGitRepoBranch

`func (o *OperatorDbt) GetGitRepoBranch() string`

GetGitRepoBranch returns the GitRepoBranch field if non-nil, zero value otherwise.

### GetGitRepoBranchOk

`func (o *OperatorDbt) GetGitRepoBranchOk() (*string, bool)`

GetGitRepoBranchOk returns a tuple with the GitRepoBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepoBranch

`func (o *OperatorDbt) SetGitRepoBranch(v string)`

SetGitRepoBranch sets GitRepoBranch field to given value.

### HasGitRepoBranch

`func (o *OperatorDbt) HasGitRepoBranch() bool`

HasGitRepoBranch returns a boolean if a field has been set.

### GetDockerImage

`func (o *OperatorDbt) GetDockerImage() string`

GetDockerImage returns the DockerImage field if non-nil, zero value otherwise.

### GetDockerImageOk

`func (o *OperatorDbt) GetDockerImageOk() (*string, bool)`

GetDockerImageOk returns a tuple with the DockerImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerImage

`func (o *OperatorDbt) SetDockerImage(v string)`

SetDockerImage sets DockerImage field to given value.

### HasDockerImage

`func (o *OperatorDbt) HasDockerImage() bool`

HasDockerImage returns a boolean if a field has been set.

### GetDbtArguments

`func (o *OperatorDbt) GetDbtArguments() string`

GetDbtArguments returns the DbtArguments field if non-nil, zero value otherwise.

### GetDbtArgumentsOk

`func (o *OperatorDbt) GetDbtArgumentsOk() (*string, bool)`

GetDbtArgumentsOk returns a tuple with the DbtArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbtArguments

`func (o *OperatorDbt) SetDbtArguments(v string)`

SetDbtArguments sets DbtArguments field to given value.

### HasDbtArguments

`func (o *OperatorDbt) HasDbtArguments() bool`

HasDbtArguments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


