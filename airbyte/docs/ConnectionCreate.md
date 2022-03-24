# ConnectionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Optional name of the connection | [optional] 
**NamespaceDefinition** | Pointer to [**NamespaceDefinitionType**](NamespaceDefinitionType.md) |  | [optional] [default to NAMESPACEDEFINITIONTYPE_SOURCE]
**NamespaceFormat** | Pointer to **string** | Used when namespaceDefinition is &#39;customformat&#39;. If blank then behaves like namespaceDefinition &#x3D; &#39;destination&#39;. If \&quot;${SOURCE_NAMESPACE}\&quot; then behaves like namespaceDefinition &#x3D; &#39;source&#39;. | [optional] [default to "null"]
**Prefix** | Pointer to **string** | Prefix that will be prepended to the name of each stream when it is written to the destination. | [optional] 
**SourceId** | **string** |  | 
**DestinationId** | **string** |  | 
**OperationIds** | Pointer to **[]string** |  | [optional] 
**SyncCatalog** | Pointer to [**AirbyteCatalog**](AirbyteCatalog.md) |  | [optional] 
**Schedule** | Pointer to [**ConnectionSchedule**](ConnectionSchedule.md) |  | [optional] 
**Status** | [**ConnectionStatus**](ConnectionStatus.md) |  | 
**ResourceRequirements** | Pointer to [**ResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewConnectionCreate

`func NewConnectionCreate(sourceId string, destinationId string, status ConnectionStatus, ) *ConnectionCreate`

NewConnectionCreate instantiates a new ConnectionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCreateWithDefaults

`func NewConnectionCreateWithDefaults() *ConnectionCreate`

NewConnectionCreateWithDefaults instantiates a new ConnectionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceDefinition

`func (o *ConnectionCreate) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *ConnectionCreate) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *ConnectionCreate) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *ConnectionCreate) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *ConnectionCreate) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *ConnectionCreate) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *ConnectionCreate) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *ConnectionCreate) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *ConnectionCreate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConnectionCreate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConnectionCreate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConnectionCreate) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionCreate) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionCreate) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionCreate) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *ConnectionCreate) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionCreate) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionCreate) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetOperationIds

`func (o *ConnectionCreate) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *ConnectionCreate) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *ConnectionCreate) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.

### HasOperationIds

`func (o *ConnectionCreate) HasOperationIds() bool`

HasOperationIds returns a boolean if a field has been set.

### GetSyncCatalog

`func (o *ConnectionCreate) GetSyncCatalog() AirbyteCatalog`

GetSyncCatalog returns the SyncCatalog field if non-nil, zero value otherwise.

### GetSyncCatalogOk

`func (o *ConnectionCreate) GetSyncCatalogOk() (*AirbyteCatalog, bool)`

GetSyncCatalogOk returns a tuple with the SyncCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCatalog

`func (o *ConnectionCreate) SetSyncCatalog(v AirbyteCatalog)`

SetSyncCatalog sets SyncCatalog field to given value.

### HasSyncCatalog

`func (o *ConnectionCreate) HasSyncCatalog() bool`

HasSyncCatalog returns a boolean if a field has been set.

### GetSchedule

`func (o *ConnectionCreate) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConnectionCreate) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConnectionCreate) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ConnectionCreate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionCreate) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionCreate) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionCreate) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.


### GetResourceRequirements

`func (o *ConnectionCreate) GetResourceRequirements() ResourceRequirements`

GetResourceRequirements returns the ResourceRequirements field if non-nil, zero value otherwise.

### GetResourceRequirementsOk

`func (o *ConnectionCreate) GetResourceRequirementsOk() (*ResourceRequirements, bool)`

GetResourceRequirementsOk returns a tuple with the ResourceRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRequirements

`func (o *ConnectionCreate) SetResourceRequirements(v ResourceRequirements)`

SetResourceRequirements sets ResourceRequirements field to given value.

### HasResourceRequirements

`func (o *ConnectionCreate) HasResourceRequirements() bool`

HasResourceRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


