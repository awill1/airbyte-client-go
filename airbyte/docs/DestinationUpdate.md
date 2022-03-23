# DestinationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationId** | **string** |  | 
**ConnectionConfiguration** | **interface{}** | The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition. | 
**Name** | **string** |  | 

## Methods

### NewDestinationUpdate

`func NewDestinationUpdate(destinationId string, connectionConfiguration interface{}, name string, ) *DestinationUpdate`

NewDestinationUpdate instantiates a new DestinationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationUpdateWithDefaults

`func NewDestinationUpdateWithDefaults() *DestinationUpdate`

NewDestinationUpdateWithDefaults instantiates a new DestinationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationId

`func (o *DestinationUpdate) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *DestinationUpdate) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *DestinationUpdate) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetConnectionConfiguration

`func (o *DestinationUpdate) GetConnectionConfiguration() interface{}`

GetConnectionConfiguration returns the ConnectionConfiguration field if non-nil, zero value otherwise.

### GetConnectionConfigurationOk

`func (o *DestinationUpdate) GetConnectionConfigurationOk() (*interface{}, bool)`

GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfiguration

`func (o *DestinationUpdate) SetConnectionConfiguration(v interface{})`

SetConnectionConfiguration sets ConnectionConfiguration field to given value.


### SetConnectionConfigurationNil

`func (o *DestinationUpdate) SetConnectionConfigurationNil(b bool)`

 SetConnectionConfigurationNil sets the value for ConnectionConfiguration to be an explicit nil

### UnsetConnectionConfiguration
`func (o *DestinationUpdate) UnsetConnectionConfiguration()`

UnsetConnectionConfiguration ensures that no value is present for ConnectionConfiguration, not even an explicit nil
### GetName

`func (o *DestinationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DestinationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DestinationUpdate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


