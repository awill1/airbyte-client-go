# \DestinationDefinitionSpecificationApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDestinationDefinitionSpecification**](DestinationDefinitionSpecificationApi.md#GetDestinationDefinitionSpecification) | **Post** /v1/destination_definition_specifications/get | Get specification for a destinationDefinition



## GetDestinationDefinitionSpecification

> DestinationDefinitionSpecificationRead GetDestinationDefinitionSpecification(ctx).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()

Get specification for a destinationDefinition

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
    resp, r, err := apiClient.DestinationDefinitionSpecificationApi.GetDestinationDefinitionSpecification(context.Background()).DestinationDefinitionIdRequestBody(destinationDefinitionIdRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DestinationDefinitionSpecificationApi.GetDestinationDefinitionSpecification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestinationDefinitionSpecification`: DestinationDefinitionSpecificationRead
    fmt.Fprintf(os.Stdout, "Response from `DestinationDefinitionSpecificationApi.GetDestinationDefinitionSpecification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationDefinitionSpecificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationDefinitionIdRequestBody** | [**DestinationDefinitionIdRequestBody**](DestinationDefinitionIdRequestBody.md) |  | 

### Return type

[**DestinationDefinitionSpecificationRead**](DestinationDefinitionSpecificationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

