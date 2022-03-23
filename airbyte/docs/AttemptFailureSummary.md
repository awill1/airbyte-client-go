# AttemptFailureSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failures** | [**[]AttemptFailureReason**](AttemptFailureReason.md) |  | 
**PartialSuccess** | Pointer to **bool** | True if the number of committed records for this attempt was greater than 0. False if 0 records were committed. If not set, the number of committed records is unknown. | [optional] 

## Methods

### NewAttemptFailureSummary

`func NewAttemptFailureSummary(failures []AttemptFailureReason, ) *AttemptFailureSummary`

NewAttemptFailureSummary instantiates a new AttemptFailureSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptFailureSummaryWithDefaults

`func NewAttemptFailureSummaryWithDefaults() *AttemptFailureSummary`

NewAttemptFailureSummaryWithDefaults instantiates a new AttemptFailureSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailures

`func (o *AttemptFailureSummary) GetFailures() []AttemptFailureReason`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *AttemptFailureSummary) GetFailuresOk() (*[]AttemptFailureReason, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *AttemptFailureSummary) SetFailures(v []AttemptFailureReason)`

SetFailures sets Failures field to given value.


### GetPartialSuccess

`func (o *AttemptFailureSummary) GetPartialSuccess() bool`

GetPartialSuccess returns the PartialSuccess field if non-nil, zero value otherwise.

### GetPartialSuccessOk

`func (o *AttemptFailureSummary) GetPartialSuccessOk() (*bool, bool)`

GetPartialSuccessOk returns a tuple with the PartialSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialSuccess

`func (o *AttemptFailureSummary) SetPartialSuccess(v bool)`

SetPartialSuccess sets PartialSuccess field to given value.

### HasPartialSuccess

`func (o *AttemptFailureSummary) HasPartialSuccess() bool`

HasPartialSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


