# OperationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**OperatorConfiguration** | [**OperatorConfiguration**](OperatorConfiguration.md) |  | 

## Methods

### NewOperationCreate

`func NewOperationCreate(workspaceId string, name string, operatorConfiguration OperatorConfiguration, ) *OperationCreate`

NewOperationCreate instantiates a new OperationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationCreateWithDefaults

`func NewOperationCreateWithDefaults() *OperationCreate`

NewOperationCreateWithDefaults instantiates a new OperationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *OperationCreate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *OperationCreate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *OperationCreate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *OperationCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperationCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperationCreate) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorConfiguration

`func (o *OperationCreate) GetOperatorConfiguration() OperatorConfiguration`

GetOperatorConfiguration returns the OperatorConfiguration field if non-nil, zero value otherwise.

### GetOperatorConfigurationOk

`func (o *OperationCreate) GetOperatorConfigurationOk() (*OperatorConfiguration, bool)`

GetOperatorConfigurationOk returns a tuple with the OperatorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorConfiguration

`func (o *OperationCreate) SetOperatorConfiguration(v OperatorConfiguration)`

SetOperatorConfiguration sets OperatorConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


