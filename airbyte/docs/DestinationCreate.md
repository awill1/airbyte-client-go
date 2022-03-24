# DestinationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**DestinationDefinitionId** | **string** |  | 
**ConnectionConfiguration** | **interface{}** | The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition. | 

## Methods

### NewDestinationCreate

`func NewDestinationCreate(workspaceId string, name string, destinationDefinitionId string, connectionConfiguration interface{}, ) *DestinationCreate`

NewDestinationCreate instantiates a new DestinationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationCreateWithDefaults

`func NewDestinationCreateWithDefaults() *DestinationCreate`

NewDestinationCreateWithDefaults instantiates a new DestinationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *DestinationCreate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DestinationCreate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DestinationCreate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *DestinationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationDefinitionId

`func (o *DestinationCreate) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationCreate) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationCreate) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetConnectionConfiguration

`func (o *DestinationCreate) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *DestinationCreate) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *DestinationCreate) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.


### SetConnectionConfigurationNil

`func (o *DestinationCreate) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *DestinationCreate) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


