# JobDebugInfoRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**JobDebugRead**](JobDebugRead.md) |  | 
**Attempts** | [**[]AttemptInfoRead**](AttemptInfoRead.md) |  | 

## Methods

### NewJobDebugInfoRead

`func NewJobDebugInfoRead(job JobDebugRead, attempts []AttemptInfoRead, ) *JobDebugInfoRead`

NewJobDebugInfoRead instantiates a new JobDebugInfoRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDebugInfoReadWithDefaults

`func NewJobDebugInfoReadWithDefaults() *JobDebugInfoRead`

NewJobDebugInfoReadWithDefaults instantiates a new JobDebugInfoRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *JobDebugInfoRead) GetJob() JobDebugRead`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *JobDebugInfoRead) GetJobOk() (*JobDebugRead, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *JobDebugInfoRead) SetJob(v JobDebugRead)`

SetJob sets Job field to given value.


### GetAttempts

`func (o *JobDebugInfoRead) GetAttempts() []AttemptInfoRead`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *JobDebugInfoRead) GetAttemptsOk() (*[]AttemptInfoRead, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *JobDebugInfoRead) SetAttempts(v []AttemptInfoRead)`

SetAttempts sets Attempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


