# AirbyteStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Stream&#39;s name. | 
**JsonSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**SupportedSyncModes** | Pointer to [**[]SyncMode**](SyncMode.md) |  | [optional] 
**SourceDefinedCursor** | Pointer to **bool** | If the source defines the cursor field, then any other cursor field inputs will be ignored. If it does not, either the user_provided one is used, or the default one is used as a backup. | [optional] 
**DefaultCursorField** | Pointer to **[]string** | Path to the field that will be used to determine if a record is new or modified since the last sync. If not provided by the source, the end user will have to specify the comparable themselves. | [optional] 
**SourceDefinedPrimaryKey** | Pointer to **[][]string** | If the source defines the primary key, paths to the fields that will be used as a primary key. If not provided by the source, the end user will have to specify the primary key themselves. | [optional] 
**Namespace** | Pointer to **string** | Optional Source-defined namespace. Airbyte streams from the same sources should have the same namespace. Currently only used by JDBC destinations to determine what schema to write to. | [optional] 

## Methods

### NewAirbyteStream

`func NewAirbyteStream(name string, ) *AirbyteStream`

NewAirbyteStream instantiates a new AirbyteStream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirbyteStreamWithDefaults

`func NewAirbyteStreamWithDefaults() *AirbyteStream`

NewAirbyteStreamWithDefaults instantiates a new AirbyteStream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AirbyteStream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AirbyteStream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AirbyteStream) SetName(v string)`

SetName sets Name field to given value.


### GetJsonSchema

`func (o *AirbyteStream) GetJsonSchema() map[string]interface{}`

GetJsonSchema returns the JsonSchema field if non-nil, zero value otherwise.

### GetJsonSchemaOk

`func (o *AirbyteStream) GetJsonSchemaOk() (*map[string]interface{}, bool)`

GetJsonSchemaOk returns a tuple with the JsonSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonSchema

`func (o *AirbyteStream) SetJsonSchema(v map[string]interface{})`

SetJsonSchema sets JsonSchema field to given value.

### HasJsonSchema

`func (o *AirbyteStream) HasJsonSchema() bool`

HasJsonSchema returns a boolean if a field has been set.

### GetSupportedSyncModes

`func (o *AirbyteStream) GetSupportedSyncModes() []SyncMode`

GetSupportedSyncModes returns the SupportedSyncModes field if non-nil, zero value otherwise.

### GetSupportedSyncModesOk

`func (o *AirbyteStream) GetSupportedSyncModesOk() (*[]SyncMode, bool)`

GetSupportedSyncModesOk returns a tuple with the SupportedSyncModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedSyncModes

`func (o *AirbyteStream) SetSupportedSyncModes(v []SyncMode)`

SetSupportedSyncModes sets SupportedSyncModes field to given value.

### HasSupportedSyncModes

`func (o *AirbyteStream) HasSupportedSyncModes() bool`

HasSupportedSyncModes returns a boolean if a field has been set.

### GetSourceDefinedCursor

`func (o *AirbyteStream) GetSourceDefinedCursor() bool`

GetSourceDefinedCursor returns the SourceDefinedCursor field if non-nil, zero value otherwise.

### GetSourceDefinedCursorOk

`func (o *AirbyteStream) GetSourceDefinedCursorOk() (*bool, bool)`

GetSourceDefinedCursorOk returns a tuple with the SourceDefinedCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinedCursor

`func (o *AirbyteStream) SetSourceDefinedCursor(v bool)`

SetSourceDefinedCursor sets SourceDefinedCursor field to given value.

### HasSourceDefinedCursor

`func (o *AirbyteStream) HasSourceDefinedCursor() bool`

HasSourceDefinedCursor returns a boolean if a field has been set.

### GetDefaultCursorField

`func (o *AirbyteStream) GetDefaultCursorField() []string`

GetDefaultCursorField returns the DefaultCursorField field if non-nil, zero value otherwise.

### GetDefaultCursorFieldOk

`func (o *AirbyteStream) GetDefaultCursorFieldOk() (*[]string, bool)`

GetDefaultCursorFieldOk returns a tuple with the DefaultCursorField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCursorField

`func (o *AirbyteStream) SetDefaultCursorField(v []string)`

SetDefaultCursorField sets DefaultCursorField field to given value.

### HasDefaultCursorField

`func (o *AirbyteStream) HasDefaultCursorField() bool`

HasDefaultCursorField returns a boolean if a field has been set.

### GetSourceDefinedPrimaryKey

`func (o *AirbyteStream) GetSourceDefinedPrimaryKey() [][]string`

GetSourceDefinedPrimaryKey returns the SourceDefinedPrimaryKey field if non-nil, zero value otherwise.

### GetSourceDefinedPrimaryKeyOk

`func (o *AirbyteStream) GetSourceDefinedPrimaryKeyOk() (*[][]string, bool)`

GetSourceDefinedPrimaryKeyOk returns a tuple with the SourceDefinedPrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinedPrimaryKey

`func (o *AirbyteStream) SetSourceDefinedPrimaryKey(v [][]string)`

SetSourceDefinedPrimaryKey sets SourceDefinedPrimaryKey field to given value.

### HasSourceDefinedPrimaryKey

`func (o *AirbyteStream) HasSourceDefinedPrimaryKey() bool`

HasSourceDefinedPrimaryKey returns a boolean if a field has been set.

### GetNamespace

`func (o *AirbyteStream) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *AirbyteStream) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *AirbyteStream) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *AirbyteStream) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


