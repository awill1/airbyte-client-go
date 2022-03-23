# \SchedulerApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteDestinationCheckConnection**](SchedulerApi.md#ExecuteDestinationCheckConnection) | **Post** /v1/scheduler/destinations/check_connection | Run check connection for a given destination configuration
[**ExecuteSourceCheckConnection**](SchedulerApi.md#ExecuteSourceCheckConnection) | **Post** /v1/scheduler/sources/check_connection | Run check connection for a given source configuration
[**ExecuteSourceDiscoverSchema**](SchedulerApi.md#ExecuteSourceDiscoverSchema) | **Post** /v1/scheduler/sources/discover_schema | Run discover schema for a given source a source configuration



## ExecuteDestinationCheckConnection

> CheckConnectionRead ExecuteDestinationCheckConnection(ctx).DestinationCoreConfig(destinationCoreConfig).Execute()

Run check connection for a given destination configuration

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
    destinationCoreConfig := *openapiclient.NewDestinationCoreConfig("DestinationDefinitionId_example", interface{}({"user":"charles"})) // DestinationCoreConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerApi.ExecuteDestinationCheckConnection(context.Background()).DestinationCoreConfig(destinationCoreConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.ExecuteDestinationCheckConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteDestinationCheckConnection`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.ExecuteDestinationCheckConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteDestinationCheckConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationCoreConfig** | [**DestinationCoreConfig**](DestinationCoreConfig.md) |  | 

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


## ExecuteSourceCheckConnection

> CheckConnectionRead ExecuteSourceCheckConnection(ctx).SourceCoreConfig(sourceCoreConfig).Execute()

Run check connection for a given source configuration

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
    sourceCoreConfig := *openapiclient.NewSourceCoreConfig("SourceDefinitionId_example", interface{}({"user":"charles"})) // SourceCoreConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerApi.ExecuteSourceCheckConnection(context.Background()).SourceCoreConfig(sourceCoreConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.ExecuteSourceCheckConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteSourceCheckConnection`: CheckConnectionRead
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.ExecuteSourceCheckConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteSourceCheckConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceCoreConfig** | [**SourceCoreConfig**](SourceCoreConfig.md) |  | 

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


## ExecuteSourceDiscoverSchema

> SourceDiscoverSchemaRead ExecuteSourceDiscoverSchema(ctx).SourceCoreConfig(sourceCoreConfig).Execute()

Run discover schema for a given source a source configuration

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
    sourceCoreConfig := *openapiclient.NewSourceCoreConfig("SourceDefinitionId_example", interface{}({"user":"charles"})) // SourceCoreConfig | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SchedulerApi.ExecuteSourceDiscoverSchema(context.Background()).SourceCoreConfig(sourceCoreConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SchedulerApi.ExecuteSourceDiscoverSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteSourceDiscoverSchema`: SourceDiscoverSchemaRead
    fmt.Fprintf(os.Stdout, "Response from `SchedulerApi.ExecuteSourceDiscoverSchema`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteSourceDiscoverSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceCoreConfig** | [**SourceCoreConfig**](SourceCoreConfig.md) |  | 

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

