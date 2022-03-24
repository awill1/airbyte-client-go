# JobDebugRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**ConfigType** | [**JobConfigType**](JobConfigType.md) |  | 
**ConfigId** | **string** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**AirbyteVersion** | **string** |  | 
**SourceDefinition** | [**SourceDefinitionRead**](SourceDefinitionRead.md) |  | 
**DestinationDefinition** | [**DestinationDefinitionRead**](DestinationDefinitionRead.md) |  | 

## Methods

### NewJobDebugRead

`func NewJobDebugRead(id int64, configType JobConfigType, configId string, status JobStatus, airbyteVersion string, sourceDefinition SourceDefinitionRead, destinationDefinition DestinationDefinitionRead, ) *JobDebugRead`

NewJobDebugRead instantiates a new JobDebugRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDebugReadWithDefaults

`func NewJobDebugReadWithDefaults() *JobDebugRead`

NewJobDebugReadWithDefaults instantiates a new JobDebugRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobDebugRead) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobDebugRead) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobDebugRead) SetId(v int64)`

SetId sets Id field to given value.


### GetConfigType

`func (o *JobDebugRead) GetConfigType() JobConfigType`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *JobDebugRead) GetConfigTypeOk() (*JobConfigType, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *JobDebugRead) SetConfigType(v JobConfigType)`

SetConfigType sets ConfigType field to given value.


### GetConfigId

`func (o *JobDebugRead) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *JobDebugRead) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *JobDebugRead) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetStatus

`func (o *JobDebugRead) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobDebugRead) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobDebugRead) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetAirbyteVersion

`func (o *JobDebugRead) GetAirbyteVersion() string`

GetAirbyteVersion returns the AirbyteVersion field if non-nil, zero value otherwise.

### GetAirbyteVersionOk

`func (o *JobDebugRead) GetAirbyteVersionOk() (*string, bool)`

GetAirbyteVersionOk returns a tuple with the AirbyteVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirbyteVersion

`func (o *JobDebugRead) SetAirbyteVersion(v string)`

SetAirbyteVersion sets AirbyteVersion field to given value.


### GetSourceDefinition

`func (o *JobDebugRead) GetSourceDefinition() SourceDefinitionRead`

GetSourceDefinition returns the SourceDefinition field if non-nil, zero value otherwise.

### GetSourceDefinitionOk

`func (o *JobDebugRead) GetSourceDefinitionOk() (*SourceDefinitionRead, bool)`

GetSourceDefinitionOk returns a tuple with the SourceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinition

`func (o *JobDebugRead) SetSourceDefinition(v SourceDefinitionRead)`

SetSourceDefinition sets SourceDefinition field to given value.


### GetDestinationDefinition

`func (o *JobDebugRead) GetDestinationDefinition() DestinationDefinitionRead`

GetDestinationDefinition returns the DestinationDefinition field if non-nil, zero value otherwise.

### GetDestinationDefinitionOk

`func (o *JobDebugRead) GetDestinationDefinitionOk() (*DestinationDefinitionRead, bool)`

GetDestinationDefinitionOk returns a tuple with the DestinationDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinition

`func (o *JobDebugRead) SetDestinationDefinition(v DestinationDefinitionRead)`

SetDestinationDefinition sets DestinationDefinition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


