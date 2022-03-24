# \SourceDefinitionApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceDefinition**](SourceDefinitionApi.md#CreateSourceDefinition) | **Post** /v1/source_definitions/create | Creates a sourceDefinition
[**DeleteSourceDefinition**](SourceDefinitionApi.md#DeleteSourceDefinition) | **Post** /v1/source_definitions/delete | Delete a source definition
[**GetSourceDefinition**](SourceDefinitionApi.md#GetSourceDefinition) | **Post** /v1/source_definitions/get | Get source
[**ListLatestSourceDefinitions**](SourceDefinitionApi.md#ListLatestSourceDefinitions) | **Post** /v1/source_definitions/list_latest | List the latest sourceDefinitions Airbyte supports
[**ListSourceDefinitions**](SourceDefinitionApi.md#ListSourceDefinitions) | **Post** /v1/source_definitions/list | List all the sourceDefinitions the current Airbyte deployment is configured to use
[**UpdateSourceDefinition**](SourceDefinitionApi.md#UpdateSourceDefinition) | **Post** /v1/source_definitions/update | Update a sourceDefinition



## CreateSourceDefinition

> SourceDefinitionRead CreateSourceDefinition(ctx).SourceDefinitionCreate(sourceDefinitionCreate).Execute()

Creates a sourceDefinition

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
    sourceDefinitionCreate := *openapiclient.NewSourceDefinitionCreate("Name_example", "DockerRepository_example", "DockerImageTag_example", "DocumentationUrl_example") // SourceDefinitionCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDefinitionApi.CreateSourceDefinition(context.Background()).SourceDefinitionCreate(sourceDefinitionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.CreateSourceDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourceDefinition`: SourceDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionApi.CreateSourceDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceDefinitionCreate** | [**SourceDefinitionCreate**](SourceDefinitionCreate.md) |  | 

### Return type

[**SourceDefinitionRead**](SourceDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceDefinition

> DeleteSourceDefinition(ctx).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()

Delete a source definition

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
    sourceDefinitionIdRequestBody := *openapiclient.NewSourceDefinitionIdRequestBody("SourceDefinitionId_example") // SourceDefinitionIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDefinitionApi.DeleteSourceDefinition(context.Background()).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.DeleteSourceDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceDefinitionIdRequestBody** | [**SourceDefinitionIdRequestBody**](SourceDefinitionIdRequestBody.md) |  | 

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


## GetSourceDefinition

> SourceDefinitionRead GetSourceDefinition(ctx).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()

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
    sourceDefinitionIdRequestBody := *openapiclient.NewSourceDefinitionIdRequestBody("SourceDefinitionId_example") // SourceDefinitionIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDefinitionApi.GetSourceDefinition(context.Background()).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.GetSourceDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceDefinition`: SourceDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionApi.GetSourceDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceDefinitionIdRequestBody** | [**SourceDefinitionIdRequestBody**](SourceDefinitionIdRequestBody.md) |  | 

### Return type

[**SourceDefinitionRead**](SourceDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLatestSourceDefinitions

> SourceDefinitionReadList ListLatestSourceDefinitions(ctx).Execute()

List the latest sourceDefinitions Airbyte supports



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
    resp, r, err := apiClient.SourceDefinitionApi.ListLatestSourceDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.ListLatestSourceDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLatestSourceDefinitions`: SourceDefinitionReadList
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionApi.ListLatestSourceDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLatestSourceDefinitionsRequest struct via the builder pattern


### Return type

[**SourceDefinitionReadList**](SourceDefinitionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourceDefinitions

> SourceDefinitionReadList ListSourceDefinitions(ctx).Execute()

List all the sourceDefinitions the current Airbyte deployment is configured to use

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
    resp, r, err := apiClient.SourceDefinitionApi.ListSourceDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.ListSourceDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceDefinitions`: SourceDefinitionReadList
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionApi.ListSourceDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceDefinitionsRequest struct via the builder pattern


### Return type

[**SourceDefinitionReadList**](SourceDefinitionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceDefinition

> SourceDefinitionRead UpdateSourceDefinition(ctx).SourceDefinitionUpdate(sourceDefinitionUpdate).Execute()

Update a sourceDefinition

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
    sourceDefinitionUpdate := *openapiclient.NewSourceDefinitionUpdate("SourceDefinitionId_example", "DockerImageTag_example") // SourceDefinitionUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceDefinitionApi.UpdateSourceDefinition(context.Background()).SourceDefinitionUpdate(sourceDefinitionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionApi.UpdateSourceDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourceDefinition`: SourceDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionApi.UpdateSourceDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceDefinitionUpdate** | [**SourceDefinitionUpdate**](SourceDefinitionUpdate.md) |  | 

### Return type

[**SourceDefinitionRead**](SourceDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

