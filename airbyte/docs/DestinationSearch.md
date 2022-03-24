# DestinationSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**ConnectionConfiguration** | Pointer to **interface{}** | The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DestinationName** | Pointer to **string** |  | [optional] 

## Methods

### NewDestinationSearch

`func NewDestinationSearch() *DestinationSearch`

NewDestinationSearch instantiates a new DestinationSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationSearchWithDefaults

`func NewDestinationSearchWithDefaults() *DestinationSearch`

NewDestinationSearchWithDefaults instantiates a new DestinationSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationSearch) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationSearch) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationSearch) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.

### HasDestinationDefinitionId

`func (o *DestinationSearch) HasDestinationDefinitionId() bool`

HasDestinationDefinitionId returns a boolean if a field has been set.

### GetDestinationId

`func (o *DestinationSearch) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationSearch) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationSearch) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *DestinationSearch) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *DestinationSearch) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DestinationSearch) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DestinationSearch) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *DestinationSearch) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetConnectionConfiguration

`func (o *DestinationSearch) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *DestinationSearch) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *DestinationSearch) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.

### HasConnectionConfiguration

`func (o *DestinationSearch) HasConnectionConfiguration() bool`

HasConnectionConfiguration returns a boolean if a field has been set.

### SetConnectionConfigurationNil

`func (o *DestinationSearch) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *DestinationSearch) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil
### GetName

`func (o *DestinationSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DestinationSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDestinationName

`func (o *DestinationSearch) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *DestinationSearch) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *DestinationSearch) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.

### HasDestinationName

`func (o *DestinationSearch) HasDestinationName() bool`

HasDestinationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


