# JobInfoRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**JobRead**](JobRead.md) |  | 
**Attempts** | [**[]AttemptInfoRead**](AttemptInfoRead.md) |  | 

## Methods

### NewJobInfoRead

`func NewJobInfoRead(job JobRead, attempts []AttemptInfoRead, ) *JobInfoRead`

NewJobInfoRead instantiates a new JobInfoRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobInfoReadWithDefaults

`func NewJobInfoReadWithDefaults() *JobInfoRead`

NewJobInfoReadWithDefaults instantiates a new JobInfoRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *JobInfoRead) GetJob() JobRead`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobInfoRead) GetJobOk() (*JobRead, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobInfoRead) SetJob(v JobRead)`

SetJob sets Job field to given value.


### GetAttempts

`func (o *JobInfoRead) GetAttempts() []AttemptInfoRead`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *JobInfoRead) GetAttemptsOk() (*[]AttemptInfoRead, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *JobInfoRead) SetAttempts(v []AttemptInfoRead)`

SetAttempts sets Attempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


