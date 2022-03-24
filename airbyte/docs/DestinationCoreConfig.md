# DestinationCoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**ConnectionConfiguration** | **interface{}** | The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition. | 

## Methods

### NewDestinationCoreConfig

`func NewDestinationCoreConfig(destinationDefinitionId string, connectionConfiguration interface{}, ) *DestinationCoreConfig`

NewDestinationCoreConfig instantiates a new DestinationCoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationCoreConfigWithDefaults

`func NewDestinationCoreConfigWithDefaults() *DestinationCoreConfig`

NewDestinationCoreConfigWithDefaults instantiates a new DestinationCoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationCoreConfig) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationCoreConfig) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationCoreConfig) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetConnectionConfiguration

`func (o *DestinationCoreConfig) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *DestinationCoreConfig) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *DestinationCoreConfig) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.


### SetConnectionConfigurationNil

`func (o *DestinationCoreConfig) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *DestinationCoreConfig) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


