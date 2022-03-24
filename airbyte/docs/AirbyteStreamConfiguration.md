# AirbyteStreamConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncMode** | [**SyncMode**](SyncMode.md) |  | 
**CursorField** | Pointer to **[]string** | Path to the field that will be used to determine if a record is new or modified since the last sync. This field is REQUIRED if &#x60;sync_mode&#x60; is &#x60;incremental&#x60;. Otherwise it is ignored. | [optional] 
**DestinationSyncMode** | [**DestinationSyncMode**](DestinationSyncMode.md) |  | 
**PrimaryKey** | Pointer to **[][]string** | Paths to the fields that will be used as primary key. This field is REQUIRED if &#x60;destination_sync_mode&#x60; is &#x60;*_dedup&#x60;. Otherwise it is ignored. | [optional] 
**AliasName** | Pointer to **string** | Alias name to the stream to be used in the destination | [optional] 
**Selected** | Pointer to **bool** |  | [optional] 

## Methods

### NewAirbyteStreamConfiguration

`func NewAirbyteStreamConfiguration(syncMode SyncMode, destinationSyncMode DestinationSyncMode, ) *AirbyteStreamConfiguration`

NewAirbyteStreamConfiguration instantiates a new AirbyteStreamConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirbyteStreamConfigurationWithDefaults

`func NewAirbyteStreamConfigurationWithDefaults() *AirbyteStreamConfiguration`

NewAirbyteStreamConfigurationWithDefaults instantiates a new AirbyteStreamConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncMode

`func (o *AirbyteStreamConfiguration) GetSyncMode() SyncMode`

GetSyncMode returns the SyncMode field if non-nil, zero value otherwise.

### GetSyncModeOk

`func (o *AirbyteStreamConfiguration) GetSyncModeOk() (*SyncMode, bool)`

GetSyncModeOk returns a tuple with the SyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncMode

`func (o *AirbyteStreamConfiguration) SetSyncMode(v SyncMode)`

SetSyncMode sets SyncMode field to given value.


### GetCursorField

`func (o *AirbyteStreamConfiguration) GetCursorField() []string`

GetCursorField returns the CursorField field if non-nil, zero value otherwise.

### GetCursorFieldOk

`func (o *AirbyteStreamConfiguration) GetCursorFieldOk() (*[]string, bool)`

GetCursorFieldOk returns a tuple with the CursorField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursorField

`func (o *AirbyteStreamConfiguration) SetCursorField(v []string)`

SetCursorField sets CursorField field to given value.

### HasCursorField

`func (o *AirbyteStreamConfiguration) HasCursorField() bool`

HasCursorField returns a boolean if a field has been set.

### GetDestinationSyncMode

`func (o *AirbyteStreamConfiguration) GetDestinationSyncMode() DestinationSyncMode`

GetDestinationSyncMode returns the DestinationSyncMode field if non-nil, zero value otherwise.

### GetDestinationSyncModeOk

`func (o *AirbyteStreamConfiguration) GetDestinationSyncModeOk() (*DestinationSyncMode, bool)`

GetDestinationSyncModeOk returns a tuple with the DestinationSyncMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSyncMode

`func (o *AirbyteStreamConfiguration) SetDestinationSyncMode(v DestinationSyncMode)`

SetDestinationSyncMode sets DestinationSyncMode field to given value.


### GetPrimaryKey

`func (o *AirbyteStreamConfiguration) GetPrimaryKey() [][]string`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *AirbyteStreamConfiguration) GetPrimaryKeyOk() (*[][]string, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *AirbyteStreamConfiguration) SetPrimaryKey(v [][]string)`

SetPrimaryKey sets PrimaryKey field to given value.

### HasPrimaryKey

`func (o *AirbyteStreamConfiguration) HasPrimaryKey() bool`

HasPrimaryKey returns a boolean if a field has been set.

### GetAliasName

`func (o *AirbyteStreamConfiguration) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *AirbyteStreamConfiguration) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *AirbyteStreamConfiguration) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *AirbyteStreamConfiguration) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetSelected

`func (o *AirbyteStreamConfiguration) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *AirbyteStreamConfiguration) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *AirbyteStreamConfiguration) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *AirbyteStreamConfiguration) HasSelected() bool`

HasSelected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


