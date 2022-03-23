# DestinationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**DestinationId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**ConnectionConfiguration** | **interface{}** | The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition. | 
**Name** | **string** |  | 
**DestinationName** | **string** |  | 

## Methods

### NewDestinationRead

`func NewDestinationRead(destinationDefinitionId string, destinationId string, workspaceId string, connectionConfiguration interface{}, name string, destinationName string, ) *DestinationRead`

NewDestinationRead instantiates a new DestinationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationReadWithDefaults

`func NewDestinationReadWithDefaults() *DestinationRead`

NewDestinationReadWithDefaults instantiates a new DestinationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationRead) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationRead) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationRead) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetDestinationId

`func (o *DestinationRead) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationRead) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationRead) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetWorkspaceId

`func (o *DestinationRead) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DestinationRead) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DestinationRead) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetConnectionConfiguration

`func (o *DestinationRead) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *DestinationRead) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *DestinationRead) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.


### SetConnectionConfigurationNil

`func (o *DestinationRead) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *DestinationRead) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil
### GetName

`func (o *DestinationRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationRead) SetName(v string)`

SetName sets Name field to given value.


### GetDestinationName

`func (o *DestinationRead) GetDestinationName() string`

GetDestinationName returns the DestinationName field if non-nil, zero value otherwise.

### GetDestinationNameOk

`func (o *DestinationRead) GetDestinationNameOk() (*string, bool)`

GetDestinationNameOk returns a tuple with the DestinationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationName

`func (o *DestinationRead) SetDestinationName(v string)`

SetDestinationName sets DestinationName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


