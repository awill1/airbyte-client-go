# ConnectionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**Name** | **string** |  | 
**NamespaceDefinition** | Pointer to [**NamespaceDefinitionType**](NamespaceDefinitionType.md) |  | [optional] [default to NAMESPACEDEFINITIONTYPE_SOURCE]
**NamespaceFormat** | Pointer to **string** | Used when namespaceDefinition is &#39;customformat&#39;. If blank then behaves like namespaceDefinition &#x3D; &#39;destination&#39;. If \&quot;${SOURCE_NAMESPACE}\&quot; then behaves like namespaceDefinition &#x3D; &#39;source&#39;. | [optional] [default to "null"]
**Prefix** | Pointer to **string** | Prefix that will be prepended to the name of each stream when it is written to the destination. | [optional] 
**SourceId** | **string** |  | 
**DestinationId** | **string** |  | 
**OperationIds** | Pointer to **[]string** |  | [optional] 
**SyncCatalog** | [**AirbyteCatalog**](AirbyteCatalog.md) |  | 
**Schedule** | Pointer to [**ConnectionSchedule**](ConnectionSchedule.md) |  | [optional] 
**Status** | [**ConnectionStatus**](ConnectionStatus.md) |  | 
**ResourceRequirements** | Pointer to [**ResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewConnectionRead

`func NewConnectionRead(connectionId string, name string, sourceId string, destinationId string, syncCatalog AirbyteCatalog, status ConnectionStatus, ) *ConnectionRead`

NewConnectionRead instantiates a new ConnectionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionReadWithDefaults

`func NewConnectionReadWithDefaults() *ConnectionRead`

NewConnectionReadWithDefaults instantiates a new ConnectionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetName

`func (o *ConnectionRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionRead) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceDefinition

`func (o *ConnectionRead) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *ConnectionRead) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *ConnectionRead) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *ConnectionRead) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *ConnectionRead) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *ConnectionRead) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *ConnectionRead) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *ConnectionRead) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *ConnectionRead) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConnectionRead) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConnectionRead) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConnectionRead) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *ConnectionRead) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionRead) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionRead) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetOperationIds

`func (o *ConnectionRead) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *ConnectionRead) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *ConnectionRead) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.

### HasOperationIds

`func (o *ConnectionRead) HasOperationIds() bool`

HasOperationIds returns a boolean if a field has been set.

### GetSyncCatalog

`func (o *ConnectionRead) GetSyncCatalog() AirbyteCatalog`

GetSyncCatalog returns the SyncCatalog field if non-nil, zero value otherwise.

### GetSyncCatalogOk

`func (o *ConnectionRead) GetSyncCatalogOk() (*AirbyteCatalog, bool)`

GetSyncCatalogOk returns a tuple with the SyncCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCatalog

`func (o *ConnectionRead) SetSyncCatalog(v AirbyteCatalog)`

SetSyncCatalog sets SyncCatalog field to given value.


### GetSchedule

`func (o *ConnectionRead) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConnectionRead) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConnectionRead) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ConnectionRead) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *ConnectionRead) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ConnectionRead) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetStatus

`func (o *ConnectionRead) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionRead) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionRead) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.


### GetResourceRequirements

`func (o *ConnectionRead) GetResourceRequirements() ResourceRequirements`

GetResourceRequirements returns the ResourceRequirements field if non-nil, zero value otherwise.

### GetResourceRequirementsOk

`func (o *ConnectionRead) GetResourceRequirementsOk() (*ResourceRequirements, bool)`

GetResourceRequirementsOk returns a tuple with the ResourceRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRequirements

`func (o *ConnectionRead) SetResourceRequirements(v ResourceRequirements)`

SetResourceRequirements sets ResourceRequirements field to given value.

### HasResourceRequirements

`func (o *ConnectionRead) HasResourceRequirements() bool`

HasResourceRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


