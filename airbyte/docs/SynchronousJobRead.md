# SynchronousJobRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ConfigType** | [**JobConfigType**](JobConfigType.md) |  | 
**ConfigId** | Pointer to **string** | only present if a config id was provided. | [optional] 
**CreatedAt** | **int64** |  | 
**EndedAt** | **int64** |  | 
**Succeeded** | **bool** |  | 
**Logs** | Pointer to [**LogRead**](LogRead.md) |  | [optional] 

## Methods

### NewSynchronousJobRead

`func NewSynchronousJobRead(id string, configType JobConfigType, createdAt int64, endedAt int64, succeeded bool, ) *SynchronousJobRead`

NewSynchronousJobRead instantiates a new SynchronousJobRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSynchronousJobReadWithDefaults

`func NewSynchronousJobReadWithDefaults() *SynchronousJobRead`

NewSynchronousJobReadWithDefaults instantiates a new SynchronousJobRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SynchronousJobRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SynchronousJobRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SynchronousJobRead) SetId(v string)`

SetId sets Id field to given value.


### GetConfigType

`func (o *SynchronousJobRead) GetConfigType() JobConfigType`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SynchronousJobRead) GetConfigTypeOk() (*JobConfigType, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SynchronousJobRead) SetConfigType(v JobConfigType)`

SetConfigType sets ConfigType field to given value.


### GetConfigId

`func (o *SynchronousJobRead) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *SynchronousJobRead) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *SynchronousJobRead) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *SynchronousJobRead) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SynchronousJobRead) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SynchronousJobRead) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SynchronousJobRead) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetEndedAt

`func (o *SynchronousJobRead) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *SynchronousJobRead) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *SynchronousJobRead) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.


### GetSucceeded

`func (o *SynchronousJobRead) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *SynchronousJobRead) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *SynchronousJobRead) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.


### GetLogs

`func (o *SynchronousJobRead) GetLogs() LogRead`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *SynchronousJobRead) GetLogsOk() (*LogRead, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *SynchronousJobRead) SetLogs(v LogRead)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *SynchronousJobRead) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


