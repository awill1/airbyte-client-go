# JobRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**ConfigType** | [**JobConfigType**](JobConfigType.md) |  | 
**ConfigId** | **string** |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 

## Methods

### NewJobRead

`func NewJobRead(id int64, configType JobConfigType, configId string, createdAt int64, updatedAt int64, status JobStatus, ) *JobRead`

NewJobRead instantiates a new JobRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobReadWithDefaults

`func NewJobReadWithDefaults() *JobRead`

NewJobReadWithDefaults instantiates a new JobRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobRead) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobRead) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobRead) SetId(v int64)`

SetId sets Id field to given value.


### GetConfigType

`func (o *JobRead) GetConfigType() JobConfigType`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *JobRead) GetConfigTypeOk() (*JobConfigType, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *JobRead) SetConfigType(v JobConfigType)`

SetConfigType sets ConfigType field to given value.


### GetConfigId

`func (o *JobRead) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *JobRead) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *JobRead) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetCreatedAt

`func (o *JobRead) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobRead) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobRead) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *JobRead) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *JobRead) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *JobRead) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *JobRead) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobRead) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobRead) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


