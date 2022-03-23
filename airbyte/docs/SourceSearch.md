# SourceSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | Pointer to **string** |  | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**ConnectionConfiguration** | Pointer to **interface{}** | The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SourceName** | Pointer to **string** |  | [optional] 

## Methods

### NewSourceSearch

`func NewSourceSearch() *SourceSearch`

NewSourceSearch instantiates a new SourceSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSearchWithDefaults

`func NewSourceSearchWithDefaults() *SourceSearch`

NewSourceSearchWithDefaults instantiates a new SourceSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *SourceSearch) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *SourceSearch) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *SourceSearch) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.

### HasSourceDefinitionId

`func (o *SourceSearch) HasSourceDefinitionId() bool`

HasSourceDefinitionId returns a boolean if a field has been set.

### GetSourceId

`func (o *SourceSearch) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SourceSearch) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SourceSearch) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *SourceSearch) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *SourceSearch) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SourceSearch) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SourceSearch) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *SourceSearch) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetConnectionConfiguration

`func (o *SourceSearch) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *SourceSearch) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *SourceSearch) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.

### HasConnectionConfiguration

`func (o *SourceSearch) HasConnectionConfiguration() bool`

HasConnectionConfiguration returns a boolean if a field has been set.

### SetConnectionConfigurationNil

`func (o *SourceSearch) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *SourceSearch) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil
### GetName

`func (o *SourceSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourceSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourceSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourceSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSourceName

`func (o *SourceSearch) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *SourceSearch) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *SourceSearch) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *SourceSearch) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


