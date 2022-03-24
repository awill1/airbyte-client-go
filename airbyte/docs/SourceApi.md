# \SourceApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckConnectionToSource**](SourceApi.md#CheckConnectionToSource) | **Post** /v1/sources/check_connection | Check connection to the source
[**CheckConnectionToSourceForUpdate**](SourceApi.md#CheckConnectionToSourceForUpdate) | **Post** /v1/sources/check_connection_for_update | Check connection for a proposed update to a source
[**CreateSource**](SourceApi.md#CreateSource) | **Post** /v1/sources/create | Create a source
[**DeleteSource**](SourceApi.md#DeleteSource) | **Post** /v1/sources/delete | Delete a source
[**DiscoverSchemaForSource**](SourceApi.md#DiscoverSchemaForSource) | **Post** /v1/sources/discover_schema | Discover the schema catalog of the source
[**GetSource**](SourceApi.md#GetSource) | **Post** /v1/sources/get | Get source
[**ListSourcesForWorkspace**](SourceApi.md#ListSourcesForWorkspace) | **Post** /v1/sources/list | List sources for workspace
[**SearchSources**](SourceApi.md#SearchSources) | **Post** /v1/sources/search | Search sources
[**UpdateSource**](SourceApi.md#UpdateSource) | **Post** /v1/sources/update | Update a source



## CheckConnectionToSource

> CheckConnectionRead CheckConnectionToSource(ctx).SourceIdRequestBody(sourceIdRequestBody).Execute()

Check connection to the source

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
    sourceIdRequestBody := *openapiclient.NewSourceIdRequestBody("SourceId_example") // SourceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.CheckConnectionToSource(context.Background()).SourceIdRequestBody(sourceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.CheckConnectionToSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckConnectionToSource`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.CheckConnectionToSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckConnectionToSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIdRequestBody** | [**SourceIdRequestBody**](SourceIdRequestBody.md) |  | 

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


## CheckConnectionToSourceForUpdate

> CheckConnectionRead CheckConnectionToSourceForUpdate(ctx).SourceUpdate(sourceUpdate).Execute()

Check connection for a proposed update to a source

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
    sourceUpdate := *openapiclient.NewSourceUpdate("SourceId_example", interface{}({"user":"charles"}), "Name_example") // SourceUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.CheckConnectionToSourceForUpdate(context.Background()).SourceUpdate(sourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.CheckConnectionToSourceForUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckConnectionToSourceForUpdate`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.CheckConnectionToSourceForUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckConnectionToSourceForUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceUpdate** | [**SourceUpdate**](SourceUpdate.md) |  | 

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


## CreateSource

> SourceRead CreateSource(ctx).SourceCreate(sourceCreate).Execute()

Create a source

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
    sourceCreate := *openapiclient.NewSourceCreate("SourceDefinitionId_example", interface{}({"user":"charles"}), "WorkspaceId_example", "Name_example") // SourceCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.CreateSource(context.Background()).SourceCreate(sourceCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.CreateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSource`: SourceRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.CreateSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceCreate** | [**SourceCreate**](SourceCreate.md) |  | 

### Return type

[**SourceRead**](SourceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSource

> DeleteSource(ctx).SourceIdRequestBody(sourceIdRequestBody).Execute()

Delete a source

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
    sourceIdRequestBody := *openapiclient.NewSourceIdRequestBody("SourceId_example") // SourceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.DeleteSource(context.Background()).SourceIdRequestBody(sourceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.DeleteSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIdRequestBody** | [**SourceIdRequestBody**](SourceIdRequestBody.md) |  | 

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


## DiscoverSchemaForSource

> SourceDiscoverSchemaRead DiscoverSchemaForSource(ctx).SourceIdRequestBody(sourceIdRequestBody).Execute()

Discover the schema catalog of the source

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
    sourceIdRequestBody := *openapiclient.NewSourceIdRequestBody("SourceId_example") // SourceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.DiscoverSchemaForSource(context.Background()).SourceIdRequestBody(sourceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.DiscoverSchemaForSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DiscoverSchemaForSource`: SourceDiscoverSchemaRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.DiscoverSchemaForSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDiscoverSchemaForSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIdRequestBody** | [**SourceIdRequestBody**](SourceIdRequestBody.md) |  | 

### Return type

[**SourceDiscoverSchemaRead**](SourceDiscoverSchemaRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSource

> SourceRead GetSource(ctx).SourceIdRequestBody(sourceIdRequestBody).Execute()

Get source

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
    sourceIdRequestBody := *openapiclient.NewSourceIdRequestBody("SourceId_example") // SourceIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.GetSource(context.Background()).SourceIdRequestBody(sourceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.GetSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSource`: SourceRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.GetSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIdRequestBody** | [**SourceIdRequestBody**](SourceIdRequestBody.md) |  | 

### Return type

[**SourceRead**](SourceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourcesForWorkspace

> SourceReadList ListSourcesForWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

List sources for workspace



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
    resp, r, err := apiClient.SourceApi.ListSourcesForWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.ListSourcesForWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcesForWorkspace`: SourceReadList
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.ListSourcesForWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcesForWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[**SourceReadList**](SourceReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSources

> SourceReadList SearchSources(ctx).SourceSearch(sourceSearch).Execute()

Search sources

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
    sourceSearch := *openapiclient.NewSourceSearch() // SourceSearch | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.SearchSources(context.Background()).SourceSearch(sourceSearch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.SearchSources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSources`: SourceReadList
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.SearchSources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceSearch** | [**SourceSearch**](SourceSearch.md) |  | 

### Return type

[**SourceReadList**](SourceReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSource

> SourceRead UpdateSource(ctx).SourceUpdate(sourceUpdate).Execute()

Update a source

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
    sourceUpdate := *openapiclient.NewSourceUpdate("SourceId_example", interface{}({"user":"charles"}), "Name_example") // SourceUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceApi.UpdateSource(context.Background()).SourceUpdate(sourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceApi.UpdateSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSource`: SourceRead
    fmt.Fprintf(os.Stdout, "Response from `SourceApi.UpdateSource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceUpdate** | [**SourceUpdate**](SourceUpdate.md) |  | 

### Return type

[**SourceRead**](SourceRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

