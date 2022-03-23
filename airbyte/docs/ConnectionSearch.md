# ConnectionSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NamespaceDefinition** | Pointer to [**NamespaceDefinitionType**](NamespaceDefinitionType.md) |  | [optional] [default to NAMESPACEDEFINITIONTYPE_SOURCE]
**NamespaceFormat** | Pointer to **string** | Used when namespaceDefinition is &#39;customformat&#39;. If blank then behaves like namespaceDefinition &#x3D; &#39;destination&#39;. If \&quot;${SOURCE_NAMESPACE}\&quot; then behaves like namespaceDefinition &#x3D; &#39;source&#39;. | [optional] [default to "null"]
**Prefix** | Pointer to **string** | Prefix that will be prepended to the name of each stream when it is written to the destination. | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**DestinationId** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**ConnectionSchedule**](ConnectionSchedule.md) |  | [optional] 
**Status** | Pointer to [**ConnectionStatus**](ConnectionStatus.md) |  | [optional] 
**Source** | Pointer to [**SourceSearch**](SourceSearch.md) |  | [optional] 
**Destination** | Pointer to [**DestinationSearch**](DestinationSearch.md) |  | [optional] 

## Methods

### NewConnectionSearch

`func NewConnectionSearch() *ConnectionSearch`

NewConnectionSearch instantiates a new ConnectionSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionSearchWithDefaults

`func NewConnectionSearchWithDefaults() *ConnectionSearch`

NewConnectionSearchWithDefaults instantiates a new ConnectionSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionSearch) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionSearch) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionSearch) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConnectionSearch) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetName

`func (o *ConnectionSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceDefinition

`func (o *ConnectionSearch) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *ConnectionSearch) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *ConnectionSearch) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *ConnectionSearch) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *ConnectionSearch) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *ConnectionSearch) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *ConnectionSearch) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *ConnectionSearch) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *ConnectionSearch) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ConnectionSearch) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ConnectionSearch) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ConnectionSearch) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceId

`func (o *ConnectionSearch) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ConnectionSearch) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ConnectionSearch) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ConnectionSearch) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *ConnectionSearch) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *ConnectionSearch) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *ConnectionSearch) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *ConnectionSearch) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetSchedule

`func (o *ConnectionSearch) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ConnectionSearch) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ConnectionSearch) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ConnectionSearch) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionSearch) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionSearch) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionSearch) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionSearch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSource

`func (o *ConnectionSearch) GetSource() SourceSearch`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ConnectionSearch) GetSourceOk() (*SourceSearch, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ConnectionSearch) SetSource(v SourceSearch)`

SetSource sets Source field to given value.

### HasSource

`func (o *ConnectionSearch) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *ConnectionSearch) GetDestination() DestinationSearch`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ConnectionSearch) GetDestinationOk() (*DestinationSearch, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ConnectionSearch) SetDestination(v DestinationSearch)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ConnectionSearch) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


