# \NotificationsApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TryNotificationConfig**](NotificationsApi.md#TryNotificationConfig) | **Post** /v1/notifications/try | Try sending a notifications



## TryNotificationConfig

> NotificationRead TryNotificationConfig(ctx).Notification(notification).Execute()

Try sending a notifications

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
    notification := *openapiclient.NewNotification(openapiclient.NotificationType("slack"), false, false) // Notification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsApi.TryNotificationConfig(context.Background()).Notification(notification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsApi.TryNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TryNotificationConfig`: NotificationRead
    fmt.Fprintf(os.Stdout, "Response from `NotificationsApi.TryNotificationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTryNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notification** | [**Notification**](Notification.md) |  | 

### Return type

[**NotificationRead**](NotificationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

