# WebBackendConnectionSearch

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

### NewWebBackendConnectionSearch

`func NewWebBackendConnectionSearch() *WebBackendConnectionSearch`

NewWebBackendConnectionSearch instantiates a new WebBackendConnectionSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBackendConnectionSearchWithDefaults

`func NewWebBackendConnectionSearchWithDefaults() *WebBackendConnectionSearch`

NewWebBackendConnectionSearchWithDefaults instantiates a new WebBackendConnectionSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *WebBackendConnectionSearch) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *WebBackendConnectionSearch) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *WebBackendConnectionSearch) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *WebBackendConnectionSearch) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetName

`func (o *WebBackendConnectionSearch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebBackendConnectionSearch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebBackendConnectionSearch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebBackendConnectionSearch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespaceDefinition

`func (o *WebBackendConnectionSearch) GetNamespaceDefinition() NamespaceDefinitionType`

GetNamespaceDefinition returns the NamespaceDefinition field if non-nil, zero value otherwise.

### GetNamespaceDefinitionOk

`func (o *WebBackendConnectionSearch) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool)`

GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceDefinition

`func (o *WebBackendConnectionSearch) SetNamespaceDefinition(v NamespaceDefinitionType)`

SetNamespaceDefinition sets NamespaceDefinition field to given value.

### HasNamespaceDefinition

`func (o *WebBackendConnectionSearch) HasNamespaceDefinition() bool`

HasNamespaceDefinition returns a boolean if a field has been set.

### GetNamespaceFormat

`func (o *WebBackendConnectionSearch) GetNamespaceFormat() string`

GetNamespaceFormat returns the NamespaceFormat field if non-nil, zero value otherwise.

### GetNamespaceFormatOk

`func (o *WebBackendConnectionSearch) GetNamespaceFormatOk() (*string, bool)`

GetNamespaceFormatOk returns a tuple with the NamespaceFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceFormat

`func (o *WebBackendConnectionSearch) SetNamespaceFormat(v string)`

SetNamespaceFormat sets NamespaceFormat field to given value.

### HasNamespaceFormat

`func (o *WebBackendConnectionSearch) HasNamespaceFormat() bool`

HasNamespaceFormat returns a boolean if a field has been set.

### GetPrefix

`func (o *WebBackendConnectionSearch) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *WebBackendConnectionSearch) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *WebBackendConnectionSearch) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *WebBackendConnectionSearch) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetSourceId

`func (o *WebBackendConnectionSearch) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WebBackendConnectionSearch) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WebBackendConnectionSearch) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *WebBackendConnectionSearch) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetDestinationId

`func (o *WebBackendConnectionSearch) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *WebBackendConnectionSearch) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *WebBackendConnectionSearch) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *WebBackendConnectionSearch) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### GetSchedule

`func (o *WebBackendConnectionSearch) GetSchedule() ConnectionSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WebBackendConnectionSearch) GetScheduleOk() (*ConnectionSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WebBackendConnectionSearch) SetSchedule(v ConnectionSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *WebBackendConnectionSearch) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *WebBackendConnectionSearch) GetStatus() ConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebBackendConnectionSearch) GetStatusOk() (*ConnectionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebBackendConnectionSearch) SetStatus(v ConnectionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebBackendConnectionSearch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSource

`func (o *WebBackendConnectionSearch) GetSource() SourceSearch`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebBackendConnectionSearch) GetSourceOk() (*SourceSearch, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebBackendConnectionSearch) SetSource(v SourceSearch)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebBackendConnectionSearch) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *WebBackendConnectionSearch) GetDestination() DestinationSearch`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *WebBackendConnectionSearch) GetDestinationOk() (*DestinationSearch, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *WebBackendConnectionSearch) SetDestination(v DestinationSearch)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *WebBackendConnectionSearch) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


