# \WebBackendApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebBackendCreateConnection**](WebBackendApi.md#WebBackendCreateConnection) | **Post** /v1/web_backend/connections/create | Create a connection
[**WebBackendGetConnection**](WebBackendApi.md#WebBackendGetConnection) | **Post** /v1/web_backend/connections/get | Get a connection
[**WebBackendListAllConnectionsForWorkspace**](WebBackendApi.md#WebBackendListAllConnectionsForWorkspace) | **Post** /v1/web_backend/connections/list_all | Returns all connections for a workspace.
[**WebBackendListConnectionsForWorkspace**](WebBackendApi.md#WebBackendListConnectionsForWorkspace) | **Post** /v1/web_backend/connections/list | Returns all non-deleted connections for a workspace.
[**WebBackendSearchConnections**](WebBackendApi.md#WebBackendSearchConnections) | **Post** /v1/web_backend/connections/search | Search connections
[**WebBackendUpdateConnection**](WebBackendApi.md#WebBackendUpdateConnection) | **Post** /v1/web_backend/connections/update | Update a connection



## WebBackendCreateConnection

> WebBackendConnectionRead WebBackendCreateConnection(ctx).WebBackendConnectionCreate(webBackendConnectionCreate).Execute()

Create a connection

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
    webBackendConnectionCreate := *openapiclient.NewWebBackendConnectionCreate("SourceId_example", "DestinationId_example", openapiclient.ConnectionStatus("active")) // WebBackendConnectionCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendCreateConnection(context.Background()).WebBackendConnectionCreate(webBackendConnectionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendCreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendCreateConnection`: WebBackendConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendCreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webBackendConnectionCreate** | [**WebBackendConnectionCreate**](WebBackendConnectionCreate.md) |  | 

### Return type

[**WebBackendConnectionRead**](WebBackendConnectionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebBackendGetConnection

> WebBackendConnectionRead WebBackendGetConnection(ctx).WebBackendConnectionRequestBody(webBackendConnectionRequestBody).Execute()

Get a connection

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
    webBackendConnectionRequestBody := *openapiclient.NewWebBackendConnectionRequestBody("ConnectionId_example") // WebBackendConnectionRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendGetConnection(context.Background()).WebBackendConnectionRequestBody(webBackendConnectionRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendGetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendGetConnection`: WebBackendConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendGetConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webBackendConnectionRequestBody** | [**WebBackendConnectionRequestBody**](WebBackendConnectionRequestBody.md) |  | 

### Return type

[**WebBackendConnectionRead**](WebBackendConnectionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebBackendListAllConnectionsForWorkspace

> WebBackendConnectionReadList WebBackendListAllConnectionsForWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

Returns all connections for a workspace.

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
    workspaceIdRequestBody := *openapiclient.NewWorkspaceIdRequestBody("WorkspaceId_example") // WorkspaceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendListAllConnectionsForWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendListAllConnectionsForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendListAllConnectionsForWorkspace`: WebBackendConnectionReadList
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendListAllConnectionsForWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendListAllConnectionsForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[**WebBackendConnectionReadList**](WebBackendConnectionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebBackendListConnectionsForWorkspace

> WebBackendConnectionReadList WebBackendListConnectionsForWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

Returns all non-deleted connections for a workspace.

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
    workspaceIdRequestBody := *openapiclient.NewWorkspaceIdRequestBody("WorkspaceId_example") // WorkspaceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendListConnectionsForWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendListConnectionsForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendListConnectionsForWorkspace`: WebBackendConnectionReadList
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendListConnectionsForWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendListConnectionsForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[**WebBackendConnectionReadList**](WebBackendConnectionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebBackendSearchConnections

> WebBackendConnectionReadList WebBackendSearchConnections(ctx).WebBackendConnectionSearch(webBackendConnectionSearch).Execute()

Search connections

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
    webBackendConnectionSearch := *openapiclient.NewWebBackendConnectionSearch() // WebBackendConnectionSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendSearchConnections(context.Background()).WebBackendConnectionSearch(webBackendConnectionSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendSearchConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendSearchConnections`: WebBackendConnectionReadList
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendSearchConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendSearchConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webBackendConnectionSearch** | [**WebBackendConnectionSearch**](WebBackendConnectionSearch.md) |  | 

### Return type

[**WebBackendConnectionReadList**](WebBackendConnectionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebBackendUpdateConnection

> WebBackendConnectionRead WebBackendUpdateConnection(ctx).WebBackendConnectionUpdate(webBackendConnectionUpdate).Execute()

Update a connection

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
    webBackendConnectionUpdate := *openapiclient.NewWebBackendConnectionUpdate("ConnectionId_example", *openapiclient.NewAirbyteCatalog([]openapiclient.AirbyteStreamAndConfiguration{*openapiclient.NewAirbyteStreamAndConfiguration()}), openapiclient.ConnectionStatus("active")) // WebBackendConnectionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebBackendApi.WebBackendUpdateConnection(context.Background()).WebBackendConnectionUpdate(webBackendConnectionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebBackendApi.WebBackendUpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebBackendUpdateConnection`: WebBackendConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `WebBackendApi.WebBackendUpdateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebBackendUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webBackendConnectionUpdate** | [**WebBackendConnectionUpdate**](WebBackendConnectionUpdate.md) |  | 

### Return type

[**WebBackendConnectionRead**](WebBackendConnectionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

