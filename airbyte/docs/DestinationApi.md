# \DestinationApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckConnectionToDestination**](DestinationApi.md#CheckConnectionToDestination) | **Post** /v1/destinations/check_connection | Check connection to the destination
[**CheckConnectionToDestinationForUpdate**](DestinationApi.md#CheckConnectionToDestinationForUpdate) | **Post** /v1/destinations/check_connection_for_update | Check connection for a proposed update to a destination
[**CreateDestination**](DestinationApi.md#CreateDestination) | **Post** /v1/destinations/create | Create a destination
[**DeleteDestination**](DestinationApi.md#DeleteDestination) | **Post** /v1/destinations/delete | Delete the destination
[**GetDestination**](DestinationApi.md#GetDestination) | **Post** /v1/destinations/get | Get configured destination
[**ListDestinationsForWorkspace**](DestinationApi.md#ListDestinationsForWorkspace) | **Post** /v1/destinations/list | List configured destinations for a workspace
[**SearchDestinations**](DestinationApi.md#SearchDestinations) | **Post** /v1/destinations/search | Search destinations
[**UpdateDestination**](DestinationApi.md#UpdateDestination) | **Post** /v1/destinations/update | Update a destination



## CheckConnectionToDestination

> CheckConnectionRead CheckConnectionToDestination(ctx).DestinationIdRequestBody(destinationIdRequestBody).Execute()

Check connection to the destination

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
    destinationIdRequestBody := *openapiclient.NewDestinationIdRequestBody("DestinationId_example") // DestinationIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.CheckConnectionToDestination(context.Background()).DestinationIdRequestBody(destinationIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.CheckConnectionToDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckConnectionToDestination`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.CheckConnectionToDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckConnectionToDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationIdRequestBody** | [**DestinationIdRequestBody**](DestinationIdRequestBody.md) |  | 

### Return type

[**CheckConnectionRead**](CheckConnectionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckConnectionToDestinationForUpdate

> CheckConnectionRead CheckConnectionToDestinationForUpdate(ctx).DestinationUpdate(destinationUpdate).Execute()

Check connection for a proposed update to a destination

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
    destinationUpdate := *openapiclient.NewDestinationUpdate("DestinationId_example", interface{}({"user":"charles"}), "Name_example") // DestinationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.CheckConnectionToDestinationForUpdate(context.Background()).DestinationUpdate(destinationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.CheckConnectionToDestinationForUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckConnectionToDestinationForUpdate`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.CheckConnectionToDestinationForUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckConnectionToDestinationForUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationUpdate** | [**DestinationUpdate**](DestinationUpdate.md) |  | 

### Return type

[**CheckConnectionRead**](CheckConnectionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDestination

> DestinationRead CreateDestination(ctx).DestinationCreate(destinationCreate).Execute()

Create a destination

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
    destinationCreate := *openapiclient.NewDestinationCreate("WorkspaceId_example", "Name_example", "DestinationDefinitionId_example", interface{}({"user":"charles"})) // DestinationCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.CreateDestination(context.Background()).DestinationCreate(destinationCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.CreateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDestination`: DestinationRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.CreateDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationCreate** | [**DestinationCreate**](DestinationCreate.md) |  | 

### Return type

[**DestinationRead**](DestinationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestination

> DeleteDestination(ctx).DestinationIdRequestBody(destinationIdRequestBody).Execute()

Delete the destination

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
    destinationIdRequestBody := *openapiclient.NewDestinationIdRequestBody("DestinationId_example") // DestinationIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.DeleteDestination(context.Background()).DestinationIdRequestBody(destinationIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.DeleteDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationIdRequestBody** | [**DestinationIdRequestBody**](DestinationIdRequestBody.md) |  | 

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


## GetDestination

> DestinationRead GetDestination(ctx).DestinationIdRequestBody(destinationIdRequestBody).Execute()

Get configured destination

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
    destinationIdRequestBody := *openapiclient.NewDestinationIdRequestBody("DestinationId_example") // DestinationIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.GetDestination(context.Background()).DestinationIdRequestBody(destinationIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.GetDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestination`: DestinationRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.GetDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationIdRequestBody** | [**DestinationIdRequestBody**](DestinationIdRequestBody.md) |  | 

### Return type

[**DestinationRead**](DestinationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationsForWorkspace

> DestinationReadList ListDestinationsForWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

List configured destinations for a workspace

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
    resp, r, err := apiClient.DestinationApi.ListDestinationsForWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.ListDestinationsForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDestinationsForWorkspace`: DestinationReadList
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.ListDestinationsForWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationsForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[**DestinationReadList**](DestinationReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchDestinations

> DestinationReadList SearchDestinations(ctx).DestinationSearch(destinationSearch).Execute()

Search destinations

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
    destinationSearch := *openapiclient.NewDestinationSearch() // DestinationSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.SearchDestinations(context.Background()).DestinationSearch(destinationSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.SearchDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDestinations`: DestinationReadList
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.SearchDestinations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationSearch** | [**DestinationSearch**](DestinationSearch.md) |  | 

### Return type

[**DestinationReadList**](DestinationReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestination

> DestinationRead UpdateDestination(ctx).DestinationUpdate(destinationUpdate).Execute()

Update a destination

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
    destinationUpdate := *openapiclient.NewDestinationUpdate("DestinationId_example", interface{}({"user":"charles"}), "Name_example") // DestinationUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationApi.UpdateDestination(context.Background()).DestinationUpdate(destinationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationApi.UpdateDestination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDestination`: DestinationRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationApi.UpdateDestination`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationUpdate** | [**DestinationUpdate**](DestinationUpdate.md) |  | 

### Return type

[**DestinationRead**](DestinationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

