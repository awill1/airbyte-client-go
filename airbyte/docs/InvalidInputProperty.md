# InvalidInputProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyPath** | **string** |  | 
**InvalidValue** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewInvalidInputProperty

`func NewInvalidInputProperty(propertyPath string, ) *InvalidInputProperty`

NewInvalidInputProperty instantiates a new InvalidInputProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidInputPropertyWithDefaults

`func NewInvalidInputPropertyWithDefaults() *InvalidInputProperty`

NewInvalidInputPropertyWithDefaults instantiates a new InvalidInputProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyPath

`func (o *InvalidInputProperty) GetPropertyPath() string`

GetPropertyPath returns the PropertyPath field if non-nil, zero value otherwise.

### GetPropertyPathOk

`func (o *InvalidInputProperty) GetPropertyPathOk() (*string, bool)`

GetPropertyPathOk returns a tuple with the PropertyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyPath

`func (o *InvalidInputProperty) SetPropertyPath(v string)`

SetPropertyPath sets PropertyPath field to given value.


### GetInvalidValue

`func (o *InvalidInputProperty) GetInvalidValue() string`

GetInvalidValue returns the InvalidValue field if non-nil, zero value otherwise.

### GetInvalidValueOk

`func (o *InvalidInputProperty) GetInvalidValueOk() (*string, bool)`

GetInvalidValueOk returns a tuple with the InvalidValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidValue

`func (o *InvalidInputProperty) SetInvalidValue(v string)`

SetInvalidValue sets InvalidValue field to given value.

### HasInvalidValue

`func (o *InvalidInputProperty) HasInvalidValue() bool`

HasInvalidValue returns a boolean if a field has been set.

### GetMessage

`func (o *InvalidInputProperty) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidInputProperty) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidInputProperty) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvalidInputProperty) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


