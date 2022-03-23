# ConnectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**NamespaceDefinition** | Pointer to [**NamespaceDefinitionType**](NamespaceDefinitionType.md) |  | [optional] [default to NAMESPACEDEFINITIONTYPE_SOURCE]
**NamespaceFormat** | Pointer to **string** | Used when namespaceDefinition is &#39;customformat&#39;. If blank then behaves like namespaceDefinition &#x3D; &#39;destination&#39;. If \&quot;${SOURCE_NAMESPACE}\&quot; then behaves like namespaceDefinition &#x3D; &#39;source&#39;. | [optional] [default to "null"]
**Prefix** | Pointer to **string** | Prefix that will be prepended to the name of each stream when it is written to the destination. | [optional] 
**OperationIds** | Pointer to **[]string** |  | [optional] 
**SyncCatalog** | [**AirbyteCatalog**](AirbyteCatalog.md) |  | 
**Schedule** | Pointer to [**ConnectionSchedule**](ConnectionSchedule.md) |  | [optional] 
**Status** | [**ConnectionStatus**](ConnectionStatus.md) |  | 
**ResourceRequirements** | Pointer to [**ResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewConnectionUpdate

`func NewConnectionUpdate(connectionId string, syncCatalog AirbyteCatalog, status ConnectionStatus, ) *ConnectionUpdate`

NewConnectionUpdate instantiates a new ConnectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionUpdateWithDefaults

`func NewConnectionUpdateWithDefaults() *ConnectionUpdate`

NewConnectionUpdateWithDefaults instantiates a new ConnectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionUpdate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionUpdate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionUpdate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetNamespaceDefinition

`func (o *ConnectionUpdate) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *ConnectionUpdate) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *ConnectionUpdate) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *ConnectionUpdate) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *ConnectionUpdate) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *ConnectionUpdate) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *ConnectionUpdate) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *ConnectionUpdate) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *ConnectionUpdate) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConnectionUpdate) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConnectionUpdate) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConnectionUpdate) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetOperationIds

`func (o *ConnectionUpdate) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *ConnectionUpdate) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *ConnectionUpdate) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.

### HasOperationIds

`func (o *ConnectionUpdate) HasOperationIds() bool`

HasOperationIds returns a boolean if a field has been set.

### GetSyncCatalog

`func (o *ConnectionUpdate) GetSyncCatalog() AirbyteCatalog`

GetSyncCatalog returns the SyncCatalog field if non-nil, zero value otherwise.

### GetSyncCatalogOk

`func (o *ConnectionUpdate) GetSyncCatalogOk() (*AirbyteCatalog, bool)`

GetSyncCatalogOk returns a tuple with the SyncCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCatalog

`func (o *ConnectionUpdate) SetSyncCatalog(v AirbyteCatalog)`

SetSyncCatalog sets SyncCatalog field to given value.


### GetSchedule

`func (o *ConnectionUpdate) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConnectionUpdate) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConnectionUpdate) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ConnectionUpdate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionUpdate) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionUpdate) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionUpdate) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.


### GetResourceRequirements

`func (o *ConnectionUpdate) GetResourceRequirements() ResourceRequirements`

GetResourceRequirements returns the ResourceRequirements field if non-nil, zero value otherwise.

### GetResourceRequirementsOk

`func (o *ConnectionUpdate) GetResourceRequirementsOk() (*ResourceRequirements, bool)`

GetResourceRequirementsOk returns a tuple with the ResourceRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRequirements

`func (o *ConnectionUpdate) SetResourceRequirements(v ResourceRequirements)`

SetResourceRequirements sets ResourceRequirements field to given value.

### HasResourceRequirements

`func (o *ConnectionUpdate) HasResourceRequirements() bool`

HasResourceRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


