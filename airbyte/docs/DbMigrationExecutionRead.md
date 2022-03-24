# DbMigrationExecutionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialVersion** | Pointer to **string** |  | [optional] 
**TargetVersion** | Pointer to **string** |  | [optional] 
**ExecutedMigrations** | Pointer to [**[]DbMigrationRead**](DbMigrationRead.md) |  | [optional] 

## Methods

### NewDbMigrationExecutionRead

`func NewDbMigrationExecutionRead() *DbMigrationExecutionRead`

NewDbMigrationExecutionRead instantiates a new DbMigrationExecutionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbMigrationExecutionReadWithDefaults

`func NewDbMigrationExecutionReadWithDefaults() *DbMigrationExecutionRead`

NewDbMigrationExecutionReadWithDefaults instantiates a new DbMigrationExecutionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialVersion

`func (o *DbMigrationExecutionRead) GetInitialVersion() string`

GetInitialVersion returns the InitialVersion field if non-nil, zero value otherwise.

### GetInitialVersionOk

`func (o *DbMigrationExecutionRead) GetInitialVersionOk() (*string, bool)`

GetInitialVersionOk returns a tuple with the InitialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialVersion

`func (o *DbMigrationExecutionRead) SetInitialVersion(v string)`

SetInitialVersion sets InitialVersion field to given value.

### HasInitialVersion

`func (o *DbMigrationExecutionRead) HasInitialVersion() bool`

HasInitialVersion returns a boolean if a field has been set.

### GetTargetVersion

`func (o *DbMigrationExecutionRead) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *DbMigrationExecutionRead) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *DbMigrationExecutionRead) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *DbMigrationExecutionRead) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetExecutedMigrations

`func (o *DbMigrationExecutionRead) GetExecutedMigrations() []DbMigrationRead`

GetExecutedMigrations returns the ExecutedMigrations field if non-nil, zero value otherwise.

### GetExecutedMigrationsOk

`func (o *DbMigrationExecutionRead) GetExecutedMigrationsOk() (*[]DbMigrationRead, bool)`

GetExecutedMigrationsOk returns a tuple with the ExecutedMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedMigrations

`func (o *DbMigrationExecutionRead) SetExecutedMigrations(v []DbMigrationRead)`

SetExecutedMigrations sets ExecutedMigrations field to given value.

### HasExecutedMigrations

`func (o *DbMigrationExecutionRead) HasExecutedMigrations() bool`

HasExecutedMigrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


