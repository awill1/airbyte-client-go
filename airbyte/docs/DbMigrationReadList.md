# DbMigrationReadList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Migrations** | Pointer to [**[]DbMigrationRead**](DbMigrationRead.md) |  | [optional] 

## Methods

### NewDbMigrationReadList

`func NewDbMigrationReadList() *DbMigrationReadList`

NewDbMigrationReadList instantiates a new DbMigrationReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbMigrationReadListWithDefaults

`func NewDbMigrationReadListWithDefaults() *DbMigrationReadList`

NewDbMigrationReadListWithDefaults instantiates a new DbMigrationReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMigrations

`func (o *DbMigrationReadList) GetMigrations() []DbMigrationRead`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *DbMigrationReadList) GetMigrationsOk() (*[]DbMigrationRead, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *DbMigrationReadList) SetMigrations(v []DbMigrationRead)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *DbMigrationReadList) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


