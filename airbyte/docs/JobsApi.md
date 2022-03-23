# \JobsApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](JobsApi.md#CancelJob) | **Post** /v1/jobs/cancel | Cancels a job
[**GetJobDebugInfo**](JobsApi.md#GetJobDebugInfo) | **Post** /v1/jobs/get_debug_info | Gets all information needed to debug this job
[**GetJobInfo**](JobsApi.md#GetJobInfo) | **Post** /v1/jobs/get | Get information about a job
[**ListJobsFor**](JobsApi.md#ListJobsFor) | **Post** /v1/jobs/list | Returns recent jobs for a connection. Jobs are returned in descending order by createdAt.



## CancelJob

> JobInfoRead CancelJob(ctx).JobIdRequestBody(jobIdRequestBody).Execute()

Cancels a job

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
    jobIdRequestBody := *openapiclient.NewJobIdRequestBody(int64(123)) // JobIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.CancelJob(context.Background()).JobIdRequestBody(jobIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.CancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelJob`: JobInfoRead
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.CancelJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIdRequestBody** | [**JobIdRequestBody**](JobIdRequestBody.md) |  | 

### Return type

[**JobInfoRead**](JobInfoRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDebugInfo

> JobDebugInfoRead GetJobDebugInfo(ctx).JobIdRequestBody(jobIdRequestBody).Execute()

Gets all information needed to debug this job

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
    jobIdRequestBody := *openapiclient.NewJobIdRequestBody(int64(123)) // JobIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJobDebugInfo(context.Background()).JobIdRequestBody(jobIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobDebugInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobDebugInfo`: JobDebugInfoRead
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobDebugInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDebugInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIdRequestBody** | [**JobIdRequestBody**](JobIdRequestBody.md) |  | 

### Return type

[**JobDebugInfoRead**](JobDebugInfoRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobInfo

> JobInfoRead GetJobInfo(ctx).JobIdRequestBody(jobIdRequestBody).Execute()

Get information about a job

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
    jobIdRequestBody := *openapiclient.NewJobIdRequestBody(int64(123)) // JobIdRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJobInfo(context.Background()).JobIdRequestBody(jobIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobInfo`: JobInfoRead
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobIdRequestBody** | [**JobIdRequestBody**](JobIdRequestBody.md) |  | 

### Return type

[**JobInfoRead**](JobInfoRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListJobsFor

> JobReadList ListJobsFor(ctx).JobListRequestBody(jobListRequestBody).Execute()

Returns recent jobs for a connection. Jobs are returned in descending order by createdAt.

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
    jobListRequestBody := *openapiclient.NewJobListRequestBody([]openapiclient.JobConfigType{openapiclient.JobConfigType("check_connection_source")}, "ConfigId_example") // JobListRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.ListJobsFor(context.Background()).JobListRequestBody(jobListRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.ListJobsFor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJobsFor`: JobReadList
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.ListJobsFor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJobsForRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobListRequestBody** | [**JobListRequestBody**](JobListRequestBody.md) |  | 

### Return type

[**JobReadList**](JobReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

