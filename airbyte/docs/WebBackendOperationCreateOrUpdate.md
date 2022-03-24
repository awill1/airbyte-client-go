# WebBackendOperationCreateOrUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationId** | Pointer to **string** |  | [optional] 
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**OperatorConfiguration** | [**OperatorConfiguration**](OperatorConfiguration.md) |  | 

## Methods

### NewWebBackendOperationCreateOrUpdate

`func NewWebBackendOperationCreateOrUpdate(workspaceId string, name string, operatorConfiguration OperatorConfiguration, ) *WebBackendOperationCreateOrUpdate`

NewWebBackendOperationCreateOrUpdate instantiates a new WebBackendOperationCreateOrUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBackendOperationCreateOrUpdateWithDefaults

`func NewWebBackendOperationCreateOrUpdateWithDefaults() *WebBackendOperationCreateOrUpdate`

NewWebBackendOperationCreateOrUpdateWithDefaults instantiates a new WebBackendOperationCreateOrUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationId

`func (o *WebBackendOperationCreateOrUpdate) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *WebBackendOperationCreateOrUpdate) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *WebBackendOperationCreateOrUpdate) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *WebBackendOperationCreateOrUpdate) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WebBackendOperationCreateOrUpdate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WebBackendOperationCreateOrUpdate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WebBackendOperationCreateOrUpdate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *WebBackendOperationCreateOrUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebBackendOperationCreateOrUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebBackendOperationCreateOrUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetOperatorConfiguration

`func (o *WebBackendOperationCreateOrUpdate) GetOperatorConfiguration() OperatorConfiguration`

GetOperatorConfiguration returns the OperatorConfiguration field if non-nil, zero value otherwise.

### GetOperatorConfigurationOk

`func (o *WebBackendOperationCreateOrUpdate) GetOperatorConfigurationOk() (*OperatorConfiguration, bool)`

GetOperatorConfigurationOk returns a tuple with the OperatorConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorConfiguration

`func (o *WebBackendOperationCreateOrUpdate) SetOperatorConfiguration(v OperatorConfiguration)`

SetOperatorConfiguration sets OperatorConfiguration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


