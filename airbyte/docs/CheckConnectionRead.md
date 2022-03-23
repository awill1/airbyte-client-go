# CheckConnectionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 
**JobInfo** | [**SynchronousJobRead**](SynchronousJobRead.md) |  | 

## Methods

### NewCheckConnectionRead

`func NewCheckConnectionRead(status string, jobInfo SynchronousJobRead, ) *CheckConnectionRead`

NewCheckConnectionRead instantiates a new CheckConnectionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckConnectionReadWithDefaults

`func NewCheckConnectionReadWithDefaults() *CheckConnectionRead`

NewCheckConnectionReadWithDefaults instantiates a new CheckConnectionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CheckConnectionRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckConnectionRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckConnectionRead) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *CheckConnectionRead) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CheckConnectionRead) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CheckConnectionRead) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CheckConnectionRead) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetJobInfo

`func (o *CheckConnectionRead) GetJobInfo() SynchronousJobRead`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *CheckConnectionRead) GetJobInfoOk() (*SynchronousJobRead, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *CheckConnectionRead) SetJobInfo(v SynchronousJobRead)`

SetJobInfo sets JobInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


