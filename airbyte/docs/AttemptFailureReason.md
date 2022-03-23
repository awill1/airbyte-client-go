# AttemptFailureReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureOrigin** | Pointer to [**AttemptFailureOrigin**](AttemptFailureOrigin.md) |  | [optional] 
**FailureType** | Pointer to [**AttemptFailureType**](AttemptFailureType.md) |  | [optional] 
**ExternalMessage** | Pointer to **string** |  | [optional] 
**Stacktrace** | Pointer to **string** |  | [optional] 
**Retryable** | Pointer to **bool** | True if it is known that retrying may succeed, e.g. for a transient failure. False if it is known that a retry will not succeed, e.g. for a configuration issue. If not set, retryable status is not well known. | [optional] 
**Timestamp** | **int64** |  | 

## Methods

### NewAttemptFailureReason

`func NewAttemptFailureReason(timestamp int64, ) *AttemptFailureReason`

NewAttemptFailureReason instantiates a new AttemptFailureReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptFailureReasonWithDefaults

`func NewAttemptFailureReasonWithDefaults() *AttemptFailureReason`

NewAttemptFailureReasonWithDefaults instantiates a new AttemptFailureReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureOrigin

`func (o *AttemptFailureReason) GetFailureOrigin() AttemptFailureOrigin`

GetFailureOrigin returns the FailureOrigin field if non-nil, zero value otherwise.

### GetFailureOriginOk

`func (o *AttemptFailureReason) GetFailureOriginOk() (*AttemptFailureOrigin, bool)`

GetFailureOriginOk returns a tuple with the FailureOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureOrigin

`func (o *AttemptFailureReason) SetFailureOrigin(v AttemptFailureOrigin)`

SetFailureOrigin sets FailureOrigin field to given value.

### HasFailureOrigin

`func (o *AttemptFailureReason) HasFailureOrigin() bool`

HasFailureOrigin returns a boolean if a field has been set.

### GetFailureType

`func (o *AttemptFailureReason) GetFailureType() AttemptFailureType`

GetFailureType returns the FailureType field if non-nil, zero value otherwise.

### GetFailureTypeOk

`func (o *AttemptFailureReason) GetFailureTypeOk() (*AttemptFailureType, bool)`

GetFailureTypeOk returns a tuple with the FailureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureType

`func (o *AttemptFailureReason) SetFailureType(v AttemptFailureType)`

SetFailureType sets FailureType field to given value.

### HasFailureType

`func (o *AttemptFailureReason) HasFailureType() bool`

HasFailureType returns a boolean if a field has been set.

### GetExternalMessage

`func (o *AttemptFailureReason) GetExternalMessage() string`

GetExternalMessage returns the ExternalMessage field if non-nil, zero value otherwise.

### GetExternalMessageOk

`func (o *AttemptFailureReason) GetExternalMessageOk() (*string, bool)`

GetExternalMessageOk returns a tuple with the ExternalMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalMessage

`func (o *AttemptFailureReason) SetExternalMessage(v string)`

SetExternalMessage sets ExternalMessage field to given value.

### HasExternalMessage

`func (o *AttemptFailureReason) HasExternalMessage() bool`

HasExternalMessage returns a boolean if a field has been set.

### GetStacktrace

`func (o *AttemptFailureReason) GetStacktrace() string`

GetStacktrace returns the Stacktrace field if non-nil, zero value otherwise.

### GetStacktraceOk

`func (o *AttemptFailureReason) GetStacktraceOk() (*string, bool)`

GetStacktraceOk returns a tuple with the Stacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacktrace

`func (o *AttemptFailureReason) SetStacktrace(v string)`

SetStacktrace sets Stacktrace field to given value.

### HasStacktrace

`func (o *AttemptFailureReason) HasStacktrace() bool`

HasStacktrace returns a boolean if a field has been set.

### GetRetryable

`func (o *AttemptFailureReason) GetRetryable() bool`

GetRetryable returns the Retryable field if non-nil, zero value otherwise.

### GetRetryableOk

`func (o *AttemptFailureReason) GetRetryableOk() (*bool, bool)`

GetRetryableOk returns a tuple with the Retryable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryable

`func (o *AttemptFailureReason) SetRetryable(v bool)`

SetRetryable sets Retryable field to given value.

### HasRetryable

`func (o *AttemptFailureReason) HasRetryable() bool`

HasRetryable returns a boolean if a field has been set.

### GetTimestamp

`func (o *AttemptFailureReason) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AttemptFailureReason) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AttemptFailureReason) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


