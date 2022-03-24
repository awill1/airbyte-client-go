# KnownExceptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ExceptionClassName** | Pointer to **string** |  | [optional] 
**ExceptionStack** | Pointer to **[]string** |  | [optional] 
**RootCauseExceptionClassName** | Pointer to **string** |  | [optional] 
**RootCauseExceptionStack** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKnownExceptionInfo

`func NewKnownExceptionInfo(message string, ) *KnownExceptionInfo`

NewKnownExceptionInfo instantiates a new KnownExceptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownExceptionInfoWithDefaults

`func NewKnownExceptionInfoWithDefaults() *KnownExceptionInfo`

NewKnownExceptionInfoWithDefaults instantiates a new KnownExceptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *KnownExceptionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KnownExceptionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KnownExceptionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExceptionClassName

`func (o *KnownExceptionInfo) GetExceptionClassName() string`

GetExceptionClassName returns the ExceptionClassName field if non-nil, zero value otherwise.

### GetExceptionClassNameOk

`func (o *KnownExceptionInfo) GetExceptionClassNameOk() (*string, bool)`

GetExceptionClassNameOk returns a tuple with the ExceptionClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionClassName

`func (o *KnownExceptionInfo) SetExceptionClassName(v string)`

SetExceptionClassName sets ExceptionClassName field to given value.

### HasExceptionClassName

`func (o *KnownExceptionInfo) HasExceptionClassName() bool`

HasExceptionClassName returns a boolean if a field has been set.

### GetExceptionStack

`func (o *KnownExceptionInfo) GetExceptionStack() []string`

GetExceptionStack returns the ExceptionStack field if non-nil, zero value otherwise.

### GetExceptionStackOk

`func (o *KnownExceptionInfo) GetExceptionStackOk() (*[]string, bool)`

GetExceptionStackOk returns a tuple with the ExceptionStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStack

`func (o *KnownExceptionInfo) SetExceptionStack(v []string)`

SetExceptionStack sets ExceptionStack field to given value.

### HasExceptionStack

`func (o *KnownExceptionInfo) HasExceptionStack() bool`

HasExceptionStack returns a boolean if a field has been set.

### GetRootCauseExceptionClassName

`func (o *KnownExceptionInfo) GetRootCauseExceptionClassName() string`

GetRootCauseExceptionClassName returns the RootCauseExceptionClassName field if non-nil, zero value otherwise.

### GetRootCauseExceptionClassNameOk

`func (o *KnownExceptionInfo) GetRootCauseExceptionClassNameOk() (*string, bool)`

GetRootCauseExceptionClassNameOk returns a tuple with the RootCauseExceptionClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseExceptionClassName

`func (o *KnownExceptionInfo) SetRootCauseExceptionClassName(v string)`

SetRootCauseExceptionClassName sets RootCauseExceptionClassName field to given value.

### HasRootCauseExceptionClassName

`func (o *KnownExceptionInfo) HasRootCauseExceptionClassName() bool`

HasRootCauseExceptionClassName returns a boolean if a field has been set.

### GetRootCauseExceptionStack

`func (o *KnownExceptionInfo) GetRootCauseExceptionStack() []string`

GetRootCauseExceptionStack returns the RootCauseExceptionStack field if non-nil, zero value otherwise.

### GetRootCauseExceptionStackOk

`func (o *KnownExceptionInfo) GetRootCauseExceptionStackOk() (*[]string, bool)`

GetRootCauseExceptionStackOk returns a tuple with the RootCauseExceptionStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseExceptionStack

`func (o *KnownExceptionInfo) SetRootCauseExceptionStack(v []string)`

SetRootCauseExceptionStack sets RootCauseExceptionStack field to given value.

### HasRootCauseExceptionStack

`func (o *KnownExceptionInfo) HasRootCauseExceptionStack() bool`

HasRootCauseExceptionStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


