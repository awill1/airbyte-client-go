# \DbMigrationApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExecuteMigrations**](DbMigrationApi.md#ExecuteMigrations) | **Post** /v1/db_migrations/migrate | Migrate the database to the latest version
[**ListMigrations**](DbMigrationApi.md#ListMigrations) | **Post** /v1/db_migrations/list | List all database migrations



## ExecuteMigrations

> DbMigrationExecutionRead ExecuteMigrations(ctx).DbMigrationRequestBody(dbMigrationRequestBody).Execute()

Migrate the database to the latest version

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
    dbMigrationRequestBody := *openapiclient.NewDbMigrationRequestBody("Database_example") // DbMigrationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DbMigrationApi.ExecuteMigrations(context.Background()).DbMigrationRequestBody(dbMigrationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DbMigrationApi.ExecuteMigrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteMigrations`: DbMigrationExecutionRead
    fmt.Fprintf(os.Stdout, "Response from `DbMigrationApi.ExecuteMigrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteMigrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbMigrationRequestBody** | [**DbMigrationRequestBody**](DbMigrationRequestBody.md) |  | 

### Return type

[**DbMigrationExecutionRead**](DbMigrationExecutionRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMigrations

> DbMigrationReadList ListMigrations(ctx).DbMigrationRequestBody(dbMigrationRequestBody).Execute()

List all database migrations

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
    dbMigrationRequestBody := *openapiclient.NewDbMigrationRequestBody("Database_example") // DbMigrationRequestBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DbMigrationApi.ListMigrations(context.Background()).DbMigrationRequestBody(dbMigrationRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DbMigrationApi.ListMigrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMigrations`: DbMigrationReadList
    fmt.Fprintf(os.Stdout, "Response from `DbMigrationApi.ListMigrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMigrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dbMigrationRequestBody** | [**DbMigrationRequestBody**](DbMigrationRequestBody.md) |  | 

### Return type

[**DbMigrationReadList**](DbMigrationReadList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

