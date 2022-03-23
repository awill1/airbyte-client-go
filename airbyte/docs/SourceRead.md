# SourceRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | **string** |  | 
**SourceId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**ConnectionConfiguration** | **interface{}** | The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source. | 
**Name** | **string** |  | 
**SourceName** | **string** |  | 

## Methods

### NewSourceRead

`func NewSourceRead(sourceDefinitionId string, sourceId string, workspaceId string, connectionConfiguration interface{}, name string, sourceName string, ) *SourceRead`

NewSourceRead instantiates a new SourceRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceReadWithDefaults

`func NewSourceReadWithDefaults() *SourceRead`

NewSourceReadWithDefaults instantiates a new SourceRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *SourceRead) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *SourceRead) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *SourceRead) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.


### GetSourceId

`func (o *SourceRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetWorkspaceId

`func (o *SourceRead) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SourceRead) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SourceRead) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetConnectionConfiguration

`func (o *SourceRead) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *SourceRead) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *SourceRead) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.


### SetConnectionConfigurationNil

`func (o *SourceRead) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *SourceRead) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil
### GetName

`func (o *SourceRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceRead) SetName(v string)`

SetName sets Name field to given value.


### GetSourceName

`func (o *SourceRead) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SourceRead) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SourceRead) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


