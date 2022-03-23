# \OauthApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteDestinationOAuth**](OauthApi.md#CompleteDestinationOAuth) | **Post** /v1/destination_oauths/complete_oauth | Given a destination def ID generate an access/refresh token etc.
[**CompleteSourceOAuth**](OauthApi.md#CompleteSourceOAuth) | **Post** /v1/source_oauths/complete_oauth | Given a source def ID generate an access/refresh token etc.
[**GetDestinationOAuthConsent**](OauthApi.md#GetDestinationOAuthConsent) | **Post** /v1/destination_oauths/get_consent_url | Given a destination connector definition ID, return the URL to the consent screen where to redirect the user to.
[**GetSourceOAuthConsent**](OauthApi.md#GetSourceOAuthConsent) | **Post** /v1/source_oauths/get_consent_url | Given a source connector definition ID, return the URL to the consent screen where to redirect the user to.
[**SetInstancewideDestinationOauthParams**](OauthApi.md#SetInstancewideDestinationOauthParams) | **Post** /v1/destination_oauths/oauth_params/create | Sets instancewide variables to be used for the oauth flow when creating this destination. When set, these variables will be injected into a connector&#39;s configuration before any interaction with the connector image itself. This enables running oauth flows with consistent variables e.g: the company&#39;s Google Ads developer_token, client_id, and client_secret without the user having to know about these variables. 
[**SetInstancewideSourceOauthParams**](OauthApi.md#SetInstancewideSourceOauthParams) | **Post** /v1/source_oauths/oauth_params/create | Sets instancewide variables to be used for the oauth flow when creating this source. When set, these variables will be injected into a connector&#39;s configuration before any interaction with the connector image itself. This enables running oauth flows with consistent variables e.g: the company&#39;s Google Ads developer_token, client_id, and client_secret without the user having to know about these variables. 



## CompleteDestinationOAuth

> map[string]map[string]interface{} CompleteDestinationOAuth(ctx).CompleteDestinationOAuthRequest(completeDestinationOAuthRequest).Execute()

Given a destination def ID generate an access/refresh token etc.

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
    completeDestinationOAuthRequest := *openapiclient.NewCompleteDestinationOAuthRequest("DestinationDefinitionId_example", "WorkspaceId_example") // CompleteDestinationOAuthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.CompleteDestinationOAuth(context.Background()).CompleteDestinationOAuthRequest(completeDestinationOAuthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.CompleteDestinationOAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteDestinationOAuth`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.CompleteDestinationOAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteDestinationOAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeDestinationOAuthRequest** | [**CompleteDestinationOAuthRequest**](CompleteDestinationOAuthRequest.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteSourceOAuth

> map[string]map[string]interface{} CompleteSourceOAuth(ctx).CompleteSourceOauthRequest(completeSourceOauthRequest).Execute()

Given a source def ID generate an access/refresh token etc.

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
    completeSourceOauthRequest := *openapiclient.NewCompleteSourceOauthRequest("SourceDefinitionId_example", "WorkspaceId_example") // CompleteSourceOauthRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.CompleteSourceOAuth(context.Background()).CompleteSourceOauthRequest(completeSourceOauthRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.CompleteSourceOAuth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteSourceOAuth`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.CompleteSourceOAuth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteSourceOAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **completeSourceOauthRequest** | [**CompleteSourceOauthRequest**](CompleteSourceOauthRequest.md) |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDestinationOAuthConsent

> OAuthConsentRead GetDestinationOAuthConsent(ctx).DestinationOauthConsentRequest(destinationOauthConsentRequest).Execute()

Given a destination connector definition ID, return the URL to the consent screen where to redirect the user to.

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
    destinationOauthConsentRequest := *openapiclient.NewDestinationOauthConsentRequest("DestinationDefinitionId_example", "WorkspaceId_example", "RedirectUrl_example") // DestinationOauthConsentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.GetDestinationOAuthConsent(context.Background()).DestinationOauthConsentRequest(destinationOauthConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.GetDestinationOAuthConsent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDestinationOAuthConsent`: OAuthConsentRead
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.GetDestinationOAuthConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDestinationOAuthConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationOauthConsentRequest** | [**DestinationOauthConsentRequest**](DestinationOauthConsentRequest.md) |  | 

### Return type

[**OAuthConsentRead**](OAuthConsentRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceOAuthConsent

> OAuthConsentRead GetSourceOAuthConsent(ctx).SourceOauthConsentRequest(sourceOauthConsentRequest).Execute()

Given a source connector definition ID, return the URL to the consent screen where to redirect the user to.

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
    sourceOauthConsentRequest := *openapiclient.NewSourceOauthConsentRequest("SourceDefinitionId_example", "WorkspaceId_example", "RedirectUrl_example") // SourceOauthConsentRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.GetSourceOAuthConsent(context.Background()).SourceOauthConsentRequest(sourceOauthConsentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.GetSourceOAuthConsent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceOAuthConsent`: OAuthConsentRead
    fmt.Fprintf(os.Stdout, "Response from `OauthApi.GetSourceOAuthConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceOAuthConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceOauthConsentRequest** | [**SourceOauthConsentRequest**](SourceOauthConsentRequest.md) |  | 

### Return type

[**OAuthConsentRead**](OAuthConsentRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInstancewideDestinationOauthParams

> SetInstancewideDestinationOauthParams(ctx).SetInstancewideDestinationOauthParamsRequestBody(setInstancewideDestinationOauthParamsRequestBody).Execute()

Sets instancewide variables to be used for the oauth flow when creating this destination. When set, these variables will be injected into a connector's configuration before any interaction with the connector image itself. This enables running oauth flows with consistent variables e.g: the company's Google Ads developer_token, client_id, and client_secret without the user having to know about these variables. 

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
    setInstancewideDestinationOauthParamsRequestBody := *openapiclient.NewSetInstancewideDestinationOauthParamsRequestBody("DestinationDefinitionId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // SetInstancewideDestinationOauthParamsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.SetInstancewideDestinationOauthParams(context.Background()).SetInstancewideDestinationOauthParamsRequestBody(setInstancewideDestinationOauthParamsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.SetInstancewideDestinationOauthParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetInstancewideDestinationOauthParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setInstancewideDestinationOauthParamsRequestBody** | [**SetInstancewideDestinationOauthParamsRequestBody**](SetInstancewideDestinationOauthParamsRequestBody.md) |  | 

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


## SetInstancewideSourceOauthParams

> SetInstancewideSourceOauthParams(ctx).SetInstancewideSourceOauthParamsRequestBody(setInstancewideSourceOauthParamsRequestBody).Execute()

Sets instancewide variables to be used for the oauth flow when creating this source. When set, these variables will be injected into a connector's configuration before any interaction with the connector image itself. This enables running oauth flows with consistent variables e.g: the company's Google Ads developer_token, client_id, and client_secret without the user having to know about these variables. 

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
    setInstancewideSourceOauthParamsRequestBody := *openapiclient.NewSetInstancewideSourceOauthParamsRequestBody("SourceDefinitionId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // SetInstancewideSourceOauthParamsRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OauthApi.SetInstancewideSourceOauthParams(context.Background()).SetInstancewideSourceOauthParamsRequestBody(setInstancewideSourceOauthParamsRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OauthApi.SetInstancewideSourceOauthParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetInstancewideSourceOauthParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setInstancewideSourceOauthParamsRequestBody** | [**SetInstancewideSourceOauthParamsRequestBody**](SetInstancewideSourceOauthParamsRequestBody.md) |  | 

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

