# \SourceDefinitionSpecificationApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceDefinitionSpecification**](SourceDefinitionSpecificationApi.md#GetSourceDefinitionSpecification) | **Post** /v1/source_definition_specifications/get | Get specification for a SourceDefinition.



## GetSourceDefinitionSpecification

> SourceDefinitionSpecificationRead GetSourceDefinitionSpecification(ctx).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()

Get specification for a SourceDefinition.

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
    resp, r, err := apiClient.SourceDefinitionSpecificationApi.GetSourceDefinitionSpecification(context.Background()).SourceDefinitionIdRequestBody(sourceDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceDefinitionSpecificationApi.GetSourceDefinitionSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceDefinitionSpecification`: SourceDefinitionSpecificationRead
    fmt.Fprintf(os.Stdout, "Response from `SourceDefinitionSpecificationApi.GetSourceDefinitionSpecification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceDefinitionSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceDefinitionIdRequestBody** | [**SourceDefinitionIdRequestBody**](SourceDefinitionIdRequestBody.md) |  | 

### Return type

[**SourceDefinitionSpecificationRead**](SourceDefinitionSpecificationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

