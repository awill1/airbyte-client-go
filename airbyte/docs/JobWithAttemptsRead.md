# JobWithAttemptsRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**JobRead**](JobRead.md) |  | [optional] 
**Attempts** | Pointer to [**[]AttemptRead**](AttemptRead.md) |  | [optional] 

## Methods

### NewJobWithAttemptsRead

`func NewJobWithAttemptsRead() *JobWithAttemptsRead`

NewJobWithAttemptsRead instantiates a new JobWithAttemptsRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithAttemptsReadWithDefaults

`func NewJobWithAttemptsReadWithDefaults() *JobWithAttemptsRead`

NewJobWithAttemptsReadWithDefaults instantiates a new JobWithAttemptsRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *JobWithAttemptsRead) GetJob() JobRead`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobWithAttemptsRead) GetJobOk() (*JobRead, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobWithAttemptsRead) SetJob(v JobRead)`

SetJob sets Job field to given value.

### HasJob

`func (o *JobWithAttemptsRead) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetAttempts

`func (o *JobWithAttemptsRead) GetAttempts() []AttemptRead`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *JobWithAttemptsRead) GetAttemptsOk() (*[]AttemptRead, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *JobWithAttemptsRead) SetAttempts(v []AttemptRead)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *JobWithAttemptsRead) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


