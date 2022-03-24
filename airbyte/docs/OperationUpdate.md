# OperationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | **string** |  | 
**Name** | **string** |  | 
**OperatorConfiguration** | [**OperatorConfiguration**](OperatorConfiguration.md) |  | 

## Methods

### NewOperationUpdate

`func NewOperationUpdate(operationId string, name string, operatorConfiguration OperatorConfiguration, ) *OperationUpdate`

NewOperationUpdate instantiates a new OperationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationUpdateWithDefaults

`func NewOperationUpdateWithDefaults() *OperationUpdate`

NewOperationUpdateWithDefaults instantiates a new OperationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *OperationUpdate) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *OperationUpdate) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *OperationUpdate) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetName

`func (o *OperationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorConfiguration

`func (o *OperationUpdate) GetOperatorConfiguration() OperatorConfiguration`

GetOperatorConfiguration returns the OperatorConfiguration field if non-nil, zero value otherwise.

### GetOperatorConfigurationOk

`func (o *OperationUpdate) GetOperatorConfigurationOk() (*OperatorConfiguration, bool)`

GetOperatorConfigurationOk returns a tuple with the OperatorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorConfiguration

`func (o *OperationUpdate) SetOperatorConfiguration(v OperatorConfiguration)`

SetOperatorConfiguration sets OperatorConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


