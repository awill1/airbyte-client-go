# \DestinationDefinitionApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDestinationDefinition**](DestinationDefinitionApi.md#CreateDestinationDefinition) | **Post** /v1/destination_definitions/create | Creates a destinationsDefinition
[**DeleteDestinationDefinition**](DestinationDefinitionApi.md#DeleteDestinationDefinition) | **Post** /v1/destination_definitions/delete | Delete a destination definition
[**GetDestinationDefinition**](DestinationDefinitionApi.md#GetDestinationDefinition) | **Post** /v1/destination_definitions/get | Get destinationDefinition
[**ListDestinationDefinitions**](DestinationDefinitionApi.md#ListDestinationDefinitions) | **Post** /v1/destination_definitions/list | List all the destinationDefinitions the current Airbyte deployment is configured to use
[**ListLatestDestinationDefinitions**](DestinationDefinitionApi.md#ListLatestDestinationDefinitions) | **Post** /v1/destination_definitions/list_latest | List the latest destinationDefinitions Airbyte supports
[**UpdateDestinationDefinition**](DestinationDefinitionApi.md#UpdateDestinationDefinition) | **Post** /v1/destination_definitions/update | Update destinationDefinition



## CreateDestinationDefinition

> DestinationDefinitionRead CreateDestinationDefinition(ctx).DestinationDefinitionCreate(destinationDefinitionCreate).Execute()

Creates a destinationsDefinition

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
    destinationDefinitionCreate := *openapiclient.NewDestinationDefinitionCreate("Name_example", "DockerRepository_example", "DockerImageTag_example", "DocumentationUrl_example") // DestinationDefinitionCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationDefinitionApi.CreateDestinationDefinition(context.Background()).DestinationDefinitionCreate(destinationDefinitionCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.CreateDestinationDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDestinationDefinition`: DestinationDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionApi.CreateDestinationDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDestinationDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationDefinitionCreate** | [**DestinationDefinitionCreate**](DestinationDefinitionCreate.md) |  | 

### Return type

[**DestinationDefinitionRead**](DestinationDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDestinationDefinition

> DeleteDestinationDefinition(ctx).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()

Delete a destination definition

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
    destinationDefinitionIdRequestBody := *openapiclient.NewDestinationDefinitionIdRequestBody("DestinationDefinitionId_example") // DestinationDefinitionIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationDefinitionApi.DeleteDestinationDefinition(context.Background()).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.DeleteDestinationDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDestinationDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationDefinitionIdRequestBody** | [**DestinationDefinitionIdRequestBody**](DestinationDefinitionIdRequestBody.md) |  | 

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


## GetDestinationDefinition

> DestinationDefinitionRead GetDestinationDefinition(ctx).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()

Get destinationDefinition

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
    destinationDefinitionIdRequestBody := *openapiclient.NewDestinationDefinitionIdRequestBody("DestinationDefinitionId_example") // DestinationDefinitionIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationDefinitionApi.GetDestinationDefinition(context.Background()).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.GetDestinationDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestinationDefinition`: DestinationDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionApi.GetDestinationDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationDefinitionIdRequestBody** | [**DestinationDefinitionIdRequestBody**](DestinationDefinitionIdRequestBody.md) |  | 

### Return type

[**DestinationDefinitionRead**](DestinationDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDestinationDefinitions

> DestinationDefinitionReadList ListDestinationDefinitions(ctx).Execute()

List all the destinationDefinitions the current Airbyte deployment is configured to use

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
    resp, r, err := apiClient.DestinationDefinitionApi.ListDestinationDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.ListDestinationDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDestinationDefinitions`: DestinationDefinitionReadList
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionApi.ListDestinationDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationDefinitionsRequest struct via the builder pattern


### Return type

[**DestinationDefinitionReadList**](DestinationDefinitionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLatestDestinationDefinitions

> DestinationDefinitionReadList ListLatestDestinationDefinitions(ctx).Execute()

List the latest destinationDefinitions Airbyte supports



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
    resp, r, err := apiClient.DestinationDefinitionApi.ListLatestDestinationDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.ListLatestDestinationDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLatestDestinationDefinitions`: DestinationDefinitionReadList
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionApi.ListLatestDestinationDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLatestDestinationDefinitionsRequest struct via the builder pattern


### Return type

[**DestinationDefinitionReadList**](DestinationDefinitionReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDestinationDefinition

> DestinationDefinitionRead UpdateDestinationDefinition(ctx).DestinationDefinitionUpdate(destinationDefinitionUpdate).Execute()

Update destinationDefinition

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
    destinationDefinitionUpdate := *openapiclient.NewDestinationDefinitionUpdate("DestinationDefinitionId_example") // DestinationDefinitionUpdate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DestinationDefinitionApi.UpdateDestinationDefinition(context.Background()).DestinationDefinitionUpdate(destinationDefinitionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionApi.UpdateDestinationDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDestinationDefinition`: DestinationDefinitionRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionApi.UpdateDestinationDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDestinationDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationDefinitionUpdate** | [**DestinationDefinitionUpdate**](DestinationDefinitionUpdate.md) |  | 

### Return type

[**DestinationDefinitionRead**](DestinationDefinitionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

