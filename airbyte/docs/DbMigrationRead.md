# DbMigrationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MigrationType** | **string** |  | 
**MigrationVersion** | **string** |  | 
**MigrationDescription** | **string** |  | 
**MigrationState** | Pointer to [**DbMigrationState**](DbMigrationState.md) |  | [optional] 
**MigratedBy** | Pointer to **string** |  | [optional] 
**MigratedAt** | Pointer to **int64** |  | [optional] 
**MigrationScript** | Pointer to **string** |  | [optional] 

## Methods

### NewDbMigrationRead

`func NewDbMigrationRead(migrationType string, migrationVersion string, migrationDescription string, ) *DbMigrationRead`

NewDbMigrationRead instantiates a new DbMigrationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbMigrationReadWithDefaults

`func NewDbMigrationReadWithDefaults() *DbMigrationRead`

NewDbMigrationReadWithDefaults instantiates a new DbMigrationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrationType

`func (o *DbMigrationRead) GetMigrationType() string`

GetMigrationType returns the MigrationType field if non-nil, zero value otherwise.

### GetMigrationTypeOk

`func (o *DbMigrationRead) GetMigrationTypeOk() (*string, bool)`

GetMigrationTypeOk returns a tuple with the MigrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationType

`func (o *DbMigrationRead) SetMigrationType(v string)`

SetMigrationType sets MigrationType field to given value.


### GetMigrationVersion

`func (o *DbMigrationRead) GetMigrationVersion() string`

GetMigrationVersion returns the MigrationVersion field if non-nil, zero value otherwise.

### GetMigrationVersionOk

`func (o *DbMigrationRead) GetMigrationVersionOk() (*string, bool)`

GetMigrationVersionOk returns a tuple with the MigrationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationVersion

`func (o *DbMigrationRead) SetMigrationVersion(v string)`

SetMigrationVersion sets MigrationVersion field to given value.


### GetMigrationDescription

`func (o *DbMigrationRead) GetMigrationDescription() string`

GetMigrationDescription returns the MigrationDescription field if non-nil, zero value otherwise.

### GetMigrationDescriptionOk

`func (o *DbMigrationRead) GetMigrationDescriptionOk() (*string, bool)`

GetMigrationDescriptionOk returns a tuple with the MigrationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationDescription

`func (o *DbMigrationRead) SetMigrationDescription(v string)`

SetMigrationDescription sets MigrationDescription field to given value.


### GetMigrationState

`func (o *DbMigrationRead) GetMigrationState() DbMigrationState`

GetMigrationState returns the MigrationState field if non-nil, zero value otherwise.

### GetMigrationStateOk

`func (o *DbMigrationRead) GetMigrationStateOk() (*DbMigrationState, bool)`

GetMigrationStateOk returns a tuple with the MigrationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationState

`func (o *DbMigrationRead) SetMigrationState(v DbMigrationState)`

SetMigrationState sets MigrationState field to given value.

### HasMigrationState

`func (o *DbMigrationRead) HasMigrationState() bool`

HasMigrationState returns a boolean if a field has been set.

### GetMigratedBy

`func (o *DbMigrationRead) GetMigratedBy() string`

GetMigratedBy returns the MigratedBy field if non-nil, zero value otherwise.

### GetMigratedByOk

`func (o *DbMigrationRead) GetMigratedByOk() (*string, bool)`

GetMigratedByOk returns a tuple with the MigratedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedBy

`func (o *DbMigrationRead) SetMigratedBy(v string)`

SetMigratedBy sets MigratedBy field to given value.

### HasMigratedBy

`func (o *DbMigrationRead) HasMigratedBy() bool`

HasMigratedBy returns a boolean if a field has been set.

### GetMigratedAt

`func (o *DbMigrationRead) GetMigratedAt() int64`

GetMigratedAt returns the MigratedAt field if non-nil, zero value otherwise.

### GetMigratedAtOk

`func (o *DbMigrationRead) GetMigratedAtOk() (*int64, bool)`

GetMigratedAtOk returns a tuple with the MigratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratedAt

`func (o *DbMigrationRead) SetMigratedAt(v int64)`

SetMigratedAt sets MigratedAt field to given value.

### HasMigratedAt

`func (o *DbMigrationRead) HasMigratedAt() bool`

HasMigratedAt returns a boolean if a field has been set.

### GetMigrationScript

`func (o *DbMigrationRead) GetMigrationScript() string`

GetMigrationScript returns the MigrationScript field if non-nil, zero value otherwise.

### GetMigrationScriptOk

`func (o *DbMigrationRead) GetMigrationScriptOk() (*string, bool)`

GetMigrationScriptOk returns a tuple with the MigrationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationScript

`func (o *DbMigrationRead) SetMigrationScript(v string)`

SetMigrationScript sets MigrationScript field to given value.

### HasMigrationScript

`func (o *DbMigrationRead) HasMigrationScript() bool`

HasMigrationScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


