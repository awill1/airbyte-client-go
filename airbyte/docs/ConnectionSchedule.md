# ConnectionSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Units** | **int64** |  | 
**TimeUnit** | **string** |  | 

## Methods

### NewConnectionSchedule

`func NewConnectionSchedule(units int64, timeUnit string, ) *ConnectionSchedule`

NewConnectionSchedule instantiates a new ConnectionSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionScheduleWithDefaults

`func NewConnectionScheduleWithDefaults() *ConnectionSchedule`

NewConnectionScheduleWithDefaults instantiates a new ConnectionSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnits

`func (o *ConnectionSchedule) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *ConnectionSchedule) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *ConnectionSchedule) SetUnits(v int64)`

SetUnits sets Units field to given value.


### GetTimeUnit

`func (o *ConnectionSchedule) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *ConnectionSchedule) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *ConnectionSchedule) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


