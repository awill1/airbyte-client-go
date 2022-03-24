# \DeploymentApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportArchive**](DeploymentApi.md#ExportArchive) | **Post** /v1/deployment/export | Export Airbyte Configuration and Data Archive
[**ExportWorkspace**](DeploymentApi.md#ExportWorkspace) | **Post** /v1/deployment/export_workspace | Export Airbyte Workspace Configuration
[**ImportArchive**](DeploymentApi.md#ImportArchive) | **Post** /v1/deployment/import | Import Airbyte Configuration and Data Archive
[**ImportIntoWorkspace**](DeploymentApi.md#ImportIntoWorkspace) | **Post** /v1/deployment/import_into_workspace | Import Airbyte Configuration into Workspace (this operation might change ids of imported configurations). Note, in order to use this api endpoint, you might need to upload a temporary archive resource with &#39;deployment/upload_archive_resource&#39; first 
[**UploadArchiveResource**](DeploymentApi.md#UploadArchiveResource) | **Post** /v1/deployment/upload_archive_resource | Upload a GZIP archive tarball and stage it in the server&#39;s cache as a temporary resource



## ExportArchive

> *os.File ExportArchive(ctx).Execute()

Export Airbyte Configuration and Data Archive

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
    resp, r, err := apiClient.DeploymentApi.ExportArchive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.ExportArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportArchive`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.ExportArchive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportArchiveRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/x-gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportWorkspace

> *os.File ExportWorkspace(ctx).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()

Export Airbyte Workspace Configuration

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
    resp, r, err := apiClient.DeploymentApi.ExportWorkspace(context.Background()).WorkspaceIdRequestBody(workspaceIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.ExportWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportWorkspace`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.ExportWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceIdRequestBody** | [**WorkspaceIdRequestBody**](WorkspaceIdRequestBody.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-gzip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportArchive

> ImportRead ImportArchive(ctx).Body(body).Execute()

Import Airbyte Configuration and Data Archive

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
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.ImportArchive(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.ImportArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportArchive`: ImportRead
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.ImportArchive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**ImportRead**](ImportRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-gzip
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportIntoWorkspace

> ImportRead ImportIntoWorkspace(ctx).ImportRequestBody(importRequestBody).Execute()

Import Airbyte Configuration into Workspace (this operation might change ids of imported configurations). Note, in order to use this api endpoint, you might need to upload a temporary archive resource with 'deployment/upload_archive_resource' first 

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
    importRequestBody := *openapiclient.NewImportRequestBody("ResourceId_example", "WorkspaceId_example") // ImportRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.ImportIntoWorkspace(context.Background()).ImportRequestBody(importRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.ImportIntoWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportIntoWorkspace`: ImportRead
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.ImportIntoWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportIntoWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importRequestBody** | [**ImportRequestBody**](ImportRequestBody.md) |  | 

### Return type

[**ImportRead**](ImportRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadArchiveResource

> UploadRead UploadArchiveResource(ctx).Body(body).Execute()

Upload a GZIP archive tarball and stage it in the server's cache as a temporary resource

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
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentApi.UploadArchiveResource(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentApi.UploadArchiveResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadArchiveResource`: UploadRead
    fmt.Fprintf(os.Stdout, "Response from `DeploymentApi.UploadArchiveResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadArchiveResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**UploadRead**](UploadRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-gzip
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

