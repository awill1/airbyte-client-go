# InvalidInputExceptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ExceptionClassName** | Pointer to **string** |  | [optional] 
**ExceptionStack** | Pointer to **[]string** |  | [optional] 
**ValidationErrors** | [**[]InvalidInputProperty**](InvalidInputProperty.md) |  | 

## Methods

### NewInvalidInputExceptionInfo

`func NewInvalidInputExceptionInfo(message string, validationErrors []InvalidInputProperty, ) *InvalidInputExceptionInfo`

NewInvalidInputExceptionInfo instantiates a new InvalidInputExceptionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidInputExceptionInfoWithDefaults

`func NewInvalidInputExceptionInfoWithDefaults() *InvalidInputExceptionInfo`

NewInvalidInputExceptionInfoWithDefaults instantiates a new InvalidInputExceptionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvalidInputExceptionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidInputExceptionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidInputExceptionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExceptionClassName

`func (o *InvalidInputExceptionInfo) GetExceptionClassName() string`

GetExceptionClassName returns the ExceptionClassName field if non-nil, zero value otherwise.

### GetExceptionClassNameOk

`func (o *InvalidInputExceptionInfo) GetExceptionClassNameOk() (*string, bool)`

GetExceptionClassNameOk returns a tuple with the ExceptionClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionClassName

`func (o *InvalidInputExceptionInfo) SetExceptionClassName(v string)`

SetExceptionClassName sets ExceptionClassName field to given value.

### HasExceptionClassName

`func (o *InvalidInputExceptionInfo) HasExceptionClassName() bool`

HasExceptionClassName returns a boolean if a field has been set.

### GetExceptionStack

`func (o *InvalidInputExceptionInfo) GetExceptionStack() []string`

GetExceptionStack returns the ExceptionStack field if non-nil, zero value otherwise.

### GetExceptionStackOk

`func (o *InvalidInputExceptionInfo) GetExceptionStackOk() (*[]string, bool)`

GetExceptionStackOk returns a tuple with the ExceptionStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStack

`func (o *InvalidInputExceptionInfo) SetExceptionStack(v []string)`

SetExceptionStack sets ExceptionStack field to given value.

### HasExceptionStack

`func (o *InvalidInputExceptionInfo) HasExceptionStack() bool`

HasExceptionStack returns a boolean if a field has been set.

### GetValidationErrors

`func (o *InvalidInputExceptionInfo) GetValidationErrors() []InvalidInputProperty`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *InvalidInputExceptionInfo) GetValidationErrorsOk() (*[]InvalidInputProperty, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *InvalidInputExceptionInfo) SetValidationErrors(v []InvalidInputProperty)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


