# AttemptInfoRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attempt** | [**AttemptRead**](AttemptRead.md) |  | 
**Logs** | [**LogRead**](LogRead.md) |  | 

## Methods

### NewAttemptInfoRead

`func NewAttemptInfoRead(attempt AttemptRead, logs LogRead, ) *AttemptInfoRead`

NewAttemptInfoRead instantiates a new AttemptInfoRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptInfoReadWithDefaults

`func NewAttemptInfoReadWithDefaults() *AttemptInfoRead`

NewAttemptInfoReadWithDefaults instantiates a new AttemptInfoRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttempt

`func (o *AttemptInfoRead) GetAttempt() AttemptRead`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *AttemptInfoRead) GetAttemptOk() (*AttemptRead, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *AttemptInfoRead) SetAttempt(v AttemptRead)`

SetAttempt sets Attempt field to given value.


### GetLogs

`func (o *AttemptInfoRead) GetLogs() LogRead`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *AttemptInfoRead) GetLogsOk() (*LogRead, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *AttemptInfoRead) SetLogs(v LogRead)`

SetLogs sets Logs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


