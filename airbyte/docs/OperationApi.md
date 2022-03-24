# \OperationApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOperation**](OperationApi.md#CheckOperation) | **Post** /v1/operations/check | Check if an operation to be created is valid
[**CreateOperation**](OperationApi.md#CreateOperation) | **Post** /v1/operations/create | Create an operation to be applied as part of a connection pipeline
[**DeleteOperation**](OperationApi.md#DeleteOperation) | **Post** /v1/operations/delete | Delete an operation
[**GetOperation**](OperationApi.md#GetOperation) | **Post** /v1/operations/get | Returns an operation
[**ListOperationsForConnection**](OperationApi.md#ListOperationsForConnection) | **Post** /v1/operations/list | Returns all operations for a connection.
[**UpdateOperation**](OperationApi.md#UpdateOperation) | **Post** /v1/operations/update | Update an operation



## CheckOperation

> CheckOperationRead CheckOperation(ctx).OperatorConfiguration(operatorConfiguration).Execute()

Check if an operation to be created is valid

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operatorConfiguration := *openapiclient.NewOperatorConfiguration(openapiclient.OperatorType("normalization")) // OperatorConfiguration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.CheckOperation(context.Background()).OperatorConfiguration(operatorConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.CheckOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOperation`: CheckOperationRead
    fmt.Fprintf(os.Stdout, "Response from `OperationApi.CheckOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operatorConfiguration** | [**OperatorConfiguration**](OperatorConfiguration.md) |  | 

### Return type

[**CheckOperationRead**](CheckOperationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOperation

> OperationRead CreateOperation(ctx).OperationCreate(operationCreate).Execute()

Create an operation to be applied as part of a connection pipeline

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operationCreate := *openapiclient.NewOperationCreate("WorkspaceId_example", "Name_example", *openapiclient.NewOperatorConfiguration(openapiclient.OperatorType("normalization"))) // OperationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.CreateOperation(context.Background()).OperationCreate(operationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.CreateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOperation`: OperationRead
    fmt.Fprintf(os.Stdout, "Response from `OperationApi.CreateOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationCreate** | [**OperationCreate**](OperationCreate.md) |  | 

### Return type

[**OperationRead**](OperationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOperation

> DeleteOperation(ctx).OperationIdRequestBody(operationIdRequestBody).Execute()

Delete an operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operationIdRequestBody := *openapiclient.NewOperationIdRequestBody("OperationId_example") // OperationIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.DeleteOperation(context.Background()).OperationIdRequestBody(operationIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.DeleteOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationIdRequestBody** | [**OperationIdRequestBody**](OperationIdRequestBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperation

> OperationRead GetOperation(ctx).OperationIdRequestBody(operationIdRequestBody).Execute()

Returns an operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operationIdRequestBody := *openapiclient.NewOperationIdRequestBody("OperationId_example") // OperationIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.GetOperation(context.Background()).OperationIdRequestBody(operationIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.GetOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperation`: OperationRead
    fmt.Fprintf(os.Stdout, "Response from `OperationApi.GetOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationIdRequestBody** | [**OperationIdRequestBody**](OperationIdRequestBody.md) |  | 

### Return type

[**OperationRead**](OperationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperationsForConnection

> OperationReadList ListOperationsForConnection(ctx).ConnectionIdRequestBody(connectionIdRequestBody).Execute()

Returns all operations for a connection.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    connectionIdRequestBody := *openapiclient.NewConnectionIdRequestBody("ConnectionId_example") // ConnectionIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.ListOperationsForConnection(context.Background()).ConnectionIdRequestBody(connectionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.ListOperationsForConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperationsForConnection`: OperationReadList
    fmt.Fprintf(os.Stdout, "Response from `OperationApi.ListOperationsForConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOperationsForConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionIdRequestBody** | [**ConnectionIdRequestBody**](ConnectionIdRequestBody.md) |  | 

### Return type

[**OperationReadList**](OperationReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOperation

> OperationRead UpdateOperation(ctx).OperationUpdate(operationUpdate).Execute()

Update an operation

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    operationUpdate := *openapiclient.NewOperationUpdate("OperationId_example", "Name_example", *openapiclient.NewOperatorConfiguration(openapiclient.OperatorType("normalization"))) // OperationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperationApi.UpdateOperation(context.Background()).OperationUpdate(operationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperationApi.UpdateOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOperation`: OperationRead
    fmt.Fprintf(os.Stdout, "Response from `OperationApi.UpdateOperation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **operationUpdate** | [**OperationUpdate**](OperationUpdate.md) |  | 

### Return type

[**OperationRead**](OperationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

