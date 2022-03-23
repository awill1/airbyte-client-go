# AttemptRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**Status** | [**AttemptStatus**](AttemptStatus.md) |  | 
**CreatedAt** | **int64** |  | 
**UpdatedAt** | **int64** |  | 
**EndedAt** | Pointer to **int64** |  | [optional] 
**BytesSynced** | Pointer to **int64** |  | [optional] 
**RecordsSynced** | Pointer to **int64** |  | [optional] 
**TotalStats** | Pointer to [**AttemptStats**](AttemptStats.md) |  | [optional] 
**StreamStats** | Pointer to [**[]AttemptStreamStats**](AttemptStreamStats.md) |  | [optional] 
**FailureSummary** | Pointer to [**AttemptFailureSummary**](AttemptFailureSummary.md) |  | [optional] 

## Methods

### NewAttemptRead

`func NewAttemptRead(id int64, status AttemptStatus, createdAt int64, updatedAt int64, ) *AttemptRead`

NewAttemptRead instantiates a new AttemptRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptReadWithDefaults

`func NewAttemptReadWithDefaults() *AttemptRead`

NewAttemptReadWithDefaults instantiates a new AttemptRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttemptRead) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttemptRead) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttemptRead) SetId(v int64)`

SetId sets Id field to given value.


### GetStatus

`func (o *AttemptRead) GetStatus() AttemptStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttemptRead) GetStatusOk() (*AttemptStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttemptRead) SetStatus(v AttemptStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *AttemptRead) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AttemptRead) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AttemptRead) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AttemptRead) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AttemptRead) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AttemptRead) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEndedAt

`func (o *AttemptRead) GetEndedAt() int64`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *AttemptRead) GetEndedAtOk() (*int64, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *AttemptRead) SetEndedAt(v int64)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *AttemptRead) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetBytesSynced

`func (o *AttemptRead) GetBytesSynced() int64`

GetBytesSynced returns the BytesSynced field if non-nil, zero value otherwise.

### GetBytesSyncedOk

`func (o *AttemptRead) GetBytesSyncedOk() (*int64, bool)`

GetBytesSyncedOk returns a tuple with the BytesSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytesSynced

`func (o *AttemptRead) SetBytesSynced(v int64)`

SetBytesSynced sets BytesSynced field to given value.

### HasBytesSynced

`func (o *AttemptRead) HasBytesSynced() bool`

HasBytesSynced returns a boolean if a field has been set.

### GetRecordsSynced

`func (o *AttemptRead) GetRecordsSynced() int64`

GetRecordsSynced returns the RecordsSynced field if non-nil, zero value otherwise.

### GetRecordsSyncedOk

`func (o *AttemptRead) GetRecordsSyncedOk() (*int64, bool)`

GetRecordsSyncedOk returns a tuple with the RecordsSynced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsSynced

`func (o *AttemptRead) SetRecordsSynced(v int64)`

SetRecordsSynced sets RecordsSynced field to given value.

### HasRecordsSynced

`func (o *AttemptRead) HasRecordsSynced() bool`

HasRecordsSynced returns a boolean if a field has been set.

### GetTotalStats

`func (o *AttemptRead) GetTotalStats() AttemptStats`

GetTotalStats returns the TotalStats field if non-nil, zero value otherwise.

### GetTotalStatsOk

`func (o *AttemptRead) GetTotalStatsOk() (*AttemptStats, bool)`

GetTotalStatsOk returns a tuple with the TotalStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStats

`func (o *AttemptRead) SetTotalStats(v AttemptStats)`

SetTotalStats sets TotalStats field to given value.

### HasTotalStats

`func (o *AttemptRead) HasTotalStats() bool`

HasTotalStats returns a boolean if a field has been set.

### GetStreamStats

`func (o *AttemptRead) GetStreamStats() []AttemptStreamStats`

GetStreamStats returns the StreamStats field if non-nil, zero value otherwise.

### GetStreamStatsOk

`func (o *AttemptRead) GetStreamStatsOk() (*[]AttemptStreamStats, bool)`

GetStreamStatsOk returns a tuple with the StreamStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamStats

`func (o *AttemptRead) SetStreamStats(v []AttemptStreamStats)`

SetStreamStats sets StreamStats field to given value.

### HasStreamStats

`func (o *AttemptRead) HasStreamStats() bool`

HasStreamStats returns a boolean if a field has been set.

### GetFailureSummary

`func (o *AttemptRead) GetFailureSummary() AttemptFailureSummary`

GetFailureSummary returns the FailureSummary field if non-nil, zero value otherwise.

### GetFailureSummaryOk

`func (o *AttemptRead) GetFailureSummaryOk() (*AttemptFailureSummary, bool)`

GetFailureSummaryOk returns a tuple with the FailureSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureSummary

`func (o *AttemptRead) SetFailureSummary(v AttemptFailureSummary)`

SetFailureSummary sets FailureSummary field to given value.

### HasFailureSummary

`func (o *AttemptRead) HasFailureSummary() bool`

HasFailureSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


