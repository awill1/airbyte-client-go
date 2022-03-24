# WebBackendConnectionRead

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
**SyncCatalog** | [**AirbyteCatalog**](AirbyteCatalog.md) |  | 
**Schedule** | Pointer to [**ConnectionSchedule**](ConnectionSchedule.md) |  | [optional] 
**Status** | [**ConnectionStatus**](ConnectionStatus.md) |  | 
**OperationIds** | Pointer to **[]string** |  | [optional] 
**Source** | [**SourceRead**](SourceRead.md) |  | 
**Destination** | [**DestinationRead**](DestinationRead.md) |  | 
**Operations** | Pointer to [**[]OperationRead**](OperationRead.md) |  | [optional] 
**LatestSyncJobCreatedAt** | Pointer to **int64** | epoch time of the latest sync job. null if no sync job has taken place. | [optional] 
**LatestSyncJobStatus** | Pointer to [**JobStatus**](JobStatus.md) |  | [optional] 
**IsSyncing** | **bool** |  | 
**ResourceRequirements** | Pointer to [**ResourceRequirements**](ResourceRequirements.md) |  | [optional] 

## Methods

### NewWebBackendConnectionRead

`func NewWebBackendConnectionRead(connectionId string, name string, sourceId string, destinationId string, syncCatalog AirbyteCatalog, status ConnectionStatus, source SourceRead, destination DestinationRead, isSyncing bool, ) *WebBackendConnectionRead`

NewWebBackendConnectionRead instantiates a new WebBackendConnectionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBackendConnectionReadWithDefaults

`func NewWebBackendConnectionReadWithDefaults() *WebBackendConnectionRead`

NewWebBackendConnectionReadWithDefaults instantiates a new WebBackendConnectionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *WebBackendConnectionRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *WebBackendConnectionRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *WebBackendConnectionRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetName

`func (o *WebBackendConnectionRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebBackendConnectionRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebBackendConnectionRead) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceDefinition

`func (o *WebBackendConnectionRead) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *WebBackendConnectionRead) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *WebBackendConnectionRead) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *WebBackendConnectionRead) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *WebBackendConnectionRead) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *WebBackendConnectionRead) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *WebBackendConnectionRead) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *WebBackendConnectionRead) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *WebBackendConnectionRead) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WebBackendConnectionRead) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WebBackendConnectionRead) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WebBackendConnectionRead) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceId

`func (o *WebBackendConnectionRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WebBackendConnectionRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WebBackendConnectionRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDestinationId

`func (o *WebBackendConnectionRead) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *WebBackendConnectionRead) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *WebBackendConnectionRead) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetSyncCatalog

`func (o *WebBackendConnectionRead) GetSyncCatalog() AirbyteCatalog`

GetSyncCatalog returns the SyncCatalog field if non-nil, zero value otherwise.

### GetSyncCatalogOk

`func (o *WebBackendConnectionRead) GetSyncCatalogOk() (*AirbyteCatalog, bool)`

GetSyncCatalogOk returns a tuple with the SyncCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCatalog

`func (o *WebBackendConnectionRead) SetSyncCatalog(v AirbyteCatalog)`

SetSyncCatalog sets SyncCatalog field to given value.


### GetSchedule

`func (o *WebBackendConnectionRead) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WebBackendConnectionRead) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WebBackendConnectionRead) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *WebBackendConnectionRead) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *WebBackendConnectionRead) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebBackendConnectionRead) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebBackendConnectionRead) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.


### GetOperationIds

`func (o *WebBackendConnectionRead) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *WebBackendConnectionRead) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *WebBackendConnectionRead) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.

### HasOperationIds

`func (o *WebBackendConnectionRead) HasOperationIds() bool`

HasOperationIds returns a boolean if a field has been set.

### GetSource

`func (o *WebBackendConnectionRead) GetSource() SourceRead`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebBackendConnectionRead) GetSourceOk() (*SourceRead, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebBackendConnectionRead) SetSource(v SourceRead)`

SetSource sets Source field to given value.


### GetDestination

`func (o *WebBackendConnectionRead) GetDestination() DestinationRead`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *WebBackendConnectionRead) GetDestinationOk() (*DestinationRead, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *WebBackendConnectionRead) SetDestination(v DestinationRead)`

SetDestination sets Destination field to given value.


### GetOperations

`func (o *WebBackendConnectionRead) GetOperations() []OperationRead`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *WebBackendConnectionRead) GetOperationsOk() (*[]OperationRead, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *WebBackendConnectionRead) SetOperations(v []OperationRead)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *WebBackendConnectionRead) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetLatestSyncJobCreatedAt

`func (o *WebBackendConnectionRead) GetLatestSyncJobCreatedAt() int64`

GetLatestSyncJobCreatedAt returns the LatestSyncJobCreatedAt field if non-nil, zero value otherwise.

### GetLatestSyncJobCreatedAtOk

`func (o *WebBackendConnectionRead) GetLatestSyncJobCreatedAtOk() (*int64, bool)`

GetLatestSyncJobCreatedAtOk returns a tuple with the LatestSyncJobCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSyncJobCreatedAt

`func (o *WebBackendConnectionRead) SetLatestSyncJobCreatedAt(v int64)`

SetLatestSyncJobCreatedAt sets LatestSyncJobCreatedAt field to given value.

### HasLatestSyncJobCreatedAt

`func (o *WebBackendConnectionRead) HasLatestSyncJobCreatedAt() bool`

HasLatestSyncJobCreatedAt returns a boolean if a field has been set.

### GetLatestSyncJobStatus

`func (o *WebBackendConnectionRead) GetLatestSyncJobStatus() JobStatus`

GetLatestSyncJobStatus returns the LatestSyncJobStatus field if non-nil, zero value otherwise.

### GetLatestSyncJobStatusOk

`func (o *WebBackendConnectionRead) GetLatestSyncJobStatusOk() (*JobStatus, bool)`

GetLatestSyncJobStatusOk returns a tuple with the LatestSyncJobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSyncJobStatus

`func (o *WebBackendConnectionRead) SetLatestSyncJobStatus(v JobStatus)`

SetLatestSyncJobStatus sets LatestSyncJobStatus field to given value.

### HasLatestSyncJobStatus

`func (o *WebBackendConnectionRead) HasLatestSyncJobStatus() bool`

HasLatestSyncJobStatus returns a boolean if a field has been set.

### GetIsSyncing

`func (o *WebBackendConnectionRead) GetIsSyncing() bool`

GetIsSyncing returns the IsSyncing field if non-nil, zero value otherwise.

### GetIsSyncingOk

`func (o *WebBackendConnectionRead) GetIsSyncingOk() (*bool, bool)`

GetIsSyncingOk returns a tuple with the IsSyncing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSyncing

`func (o *WebBackendConnectionRead) SetIsSyncing(v bool)`

SetIsSyncing sets IsSyncing field to given value.


### GetResourceRequirements

`func (o *WebBackendConnectionRead) GetResourceRequirements() ResourceRequirements`

GetResourceRequirements returns the ResourceRequirements field if non-nil, zero value otherwise.

### GetResourceRequirementsOk

`func (o *WebBackendConnectionRead) GetResourceRequirementsOk() (*ResourceRequirements, bool)`

GetResourceRequirementsOk returns a tuple with the ResourceRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceRequirements

`func (o *WebBackendConnectionRead) SetResourceRequirements(v ResourceRequirements)`

SetResourceRequirements sets ResourceRequirements field to given value.

### HasResourceRequirements

`func (o *WebBackendConnectionRead) HasResourceRequirements() bool`

HasResourceRequirements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


