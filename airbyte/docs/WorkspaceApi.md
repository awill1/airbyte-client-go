# \WorkspaceApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspace**](WorkspaceApi.md#CreateWorkspace) | **Post** /v1/workspaces/create | Creates a workspace
[**DeleteWorkspace**](WorkspaceApi.md#DeleteWorkspace) | **Post** /v1/workspaces/delete | Deletes a workspace
[**GetWorkspace**](WorkspaceApi.md#GetWorkspace) | **Post** /v1/workspaces/get | Find workspace by ID
[**GetWorkspaceBySlug**](WorkspaceApi.md#GetWorkspaceBySlug) | **Post** /v1/workspaces/get_by_slug | Find workspace by slug
[**ListWorkspaces**](WorkspaceApi.md#ListWorkspaces) | **Post** /v1/workspaces/list | List all workspaces registered in the current Airbyte deployment
[**UpdateWorkspace**](WorkspaceApi.md#UpdateWorkspace) | **Post** /v1/workspaces/update | Update workspace state
[**UpdateWorkspaceFeedback**](WorkspaceApi.md#UpdateWorkspaceFeedback) | **Post** /v1/workspaces/tag_feedback_status_as_done | Update workspace feedback state
[**UpdateWorkspaceName**](WorkspaceApi.md#UpdateWorkspaceName) | **Post** /v1/workspaces/update_name | Update workspace name



## CreateWorkspace

> WorkspaceRead CreateWorkspace(ctx).WorkspaceCreate(workspaceCreate).Execute()

Creates a workspace

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
    workspaceCreate := *openapiclient.NewWorkspaceCreate("Name_example") // WorkspaceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.CreateWorkspace(context.Background()).WorkspaceCreate(workspaceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: WorkspaceRead
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceCreate** | [**WorkspaceCreate**](WorkspaceCreate.md) |  | 

### Return type

[**WorkspaceRead**](WorkspaceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

Deletes a workspace

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
    resp, r, err := apiClient.WorkspaceApi.DeleteWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.DeleteWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

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


## GetWorkspace

> WorkspaceRead GetWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

Find workspace by ID

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
    resp, r, err := apiClient.WorkspaceApi.GetWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspace`: WorkspaceRead
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[**WorkspaceRead**](WorkspaceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceBySlug

> WorkspaceRead GetWorkspaceBySlug(ctx).SlugRequestBody(slugRequestBody).Execute()

Find workspace by slug

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
    slugRequestBody := *openapiclient.NewSlugRequestBody("Slug_example") // SlugRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.GetWorkspaceBySlug(context.Background()).SlugRequestBody(slugRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.GetWorkspaceBySlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspaceBySlug`: WorkspaceRead
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.GetWorkspaceBySlug`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceBySlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slugRequestBody** | [**SlugRequestBody**](SlugRequestBody.md) |  | 

### Return type

[**WorkspaceRead**](WorkspaceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> WorkspaceReadList ListWorkspaces(ctx).Execute()

List all workspaces registered in the current Airbyte deployment

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.ListWorkspaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: WorkspaceReadList
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


### Return type

[**WorkspaceReadList**](WorkspaceReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspace

> WorkspaceRead UpdateWorkspace(ctx).WorkspaceUpdate(workspaceUpdate).Execute()

Update workspace state

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
    workspaceUpdate := *openapiclient.NewWorkspaceUpdate("WorkspaceId_example", false, false, false, false) // WorkspaceUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.UpdateWorkspace(context.Background()).WorkspaceUpdate(workspaceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.UpdateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkspace`: WorkspaceRead
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.UpdateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceUpdate** | [**WorkspaceUpdate**](WorkspaceUpdate.md) |  | 

### Return type

[**WorkspaceRead**](WorkspaceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkspaceFeedback

> UpdateWorkspaceFeedback(ctx).WorkspaceGiveFeedback(workspaceGiveFeedback).Execute()

Update workspace feedback state

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
    workspaceGiveFeedback := *openapiclient.NewWorkspaceGiveFeedback("WorkspaceId_example") // WorkspaceGiveFeedback | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.UpdateWorkspaceFeedback(context.Background()).WorkspaceGiveFeedback(workspaceGiveFeedback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.UpdateWorkspaceFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceGiveFeedback** | [**WorkspaceGiveFeedback**](WorkspaceGiveFeedback.md) |  | 

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


## UpdateWorkspaceName

> WorkspaceRead UpdateWorkspaceName(ctx).WorkspaceUpdateName(workspaceUpdateName).Execute()

Update workspace name

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
    workspaceUpdateName := *openapiclient.NewWorkspaceUpdateName("WorkspaceId_example", "Name_example") // WorkspaceUpdateName | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspaceApi.UpdateWorkspaceName(context.Background()).WorkspaceUpdateName(workspaceUpdateName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspaceApi.UpdateWorkspaceName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkspaceName`: WorkspaceRead
    fmt.Fprintf(os.Stdout, "Response from `WorkspaceApi.UpdateWorkspaceName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkspaceNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceUpdateName** | [**WorkspaceUpdateName**](WorkspaceUpdateName.md) |  | 

### Return type

[**WorkspaceRead**](WorkspaceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

