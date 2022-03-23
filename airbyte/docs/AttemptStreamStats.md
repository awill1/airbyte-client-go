# AttemptStreamStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamName** | **string** |  | 
**Stats** | [**AttemptStats**](AttemptStats.md) |  | 

## Methods

### NewAttemptStreamStats

`func NewAttemptStreamStats(streamName string, stats AttemptStats, ) *AttemptStreamStats`

NewAttemptStreamStats instantiates a new AttemptStreamStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttemptStreamStatsWithDefaults

`func NewAttemptStreamStatsWithDefaults() *AttemptStreamStats`

NewAttemptStreamStatsWithDefaults instantiates a new AttemptStreamStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamName

`func (o *AttemptStreamStats) GetStreamName() string`

GetStreamName returns the StreamName field if non-nil, zero value otherwise.

### GetStreamNameOk

`func (o *AttemptStreamStats) GetStreamNameOk() (*string, bool)`

GetStreamNameOk returns a tuple with the StreamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamName

`func (o *AttemptStreamStats) SetStreamName(v string)`

SetStreamName sets StreamName field to given value.


### GetStats

`func (o *AttemptStreamStats) GetStats() AttemptStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AttemptStreamStats) GetStatsOk() (*AttemptStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AttemptStreamStats) SetStats(v AttemptStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


