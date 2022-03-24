# OperationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**OperationId** | **string** |  | 
**Name** | **string** |  | 
**OperatorConfiguration** | [**OperatorConfiguration**](OperatorConfiguration.md) |  | 

## Methods

### NewOperationRead

`func NewOperationRead(workspaceId string, operationId string, name string, operatorConfiguration OperatorConfiguration, ) *OperationRead`

NewOperationRead instantiates a new OperationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationReadWithDefaults

`func NewOperationReadWithDefaults() *OperationRead`

NewOperationReadWithDefaults instantiates a new OperationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *OperationRead) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OperationRead) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OperationRead) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetOperationId

`func (o *OperationRead) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *OperationRead) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *OperationRead) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.


### GetName

`func (o *OperationRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationRead) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorConfiguration

`func (o *OperationRead) GetOperatorConfiguration() OperatorConfiguration`

GetOperatorConfiguration returns the OperatorConfiguration field if non-nil, zero value otherwise.

### GetOperatorConfigurationOk

`func (o *OperationRead) GetOperatorConfigurationOk() (*OperatorConfiguration, bool)`

GetOperatorConfigurationOk returns a tuple with the OperatorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorConfiguration

`func (o *OperationRead) SetOperatorConfiguration(v OperatorConfiguration)`

SetOperatorConfiguration sets OperatorConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


