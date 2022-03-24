# WebBackendConnectionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithRefreshedCatalog** | Pointer to **bool** |  | [optional] 
**ConnectionId** | **string** |  | 

## Methods

### NewWebBackendConnectionRequestBody

`func NewWebBackendConnectionRequestBody(connectionId string, ) *WebBackendConnectionRequestBody`

NewWebBackendConnectionRequestBody instantiates a new WebBackendConnectionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebBackendConnectionRequestBodyWithDefaults

`func NewWebBackendConnectionRequestBodyWithDefaults() *WebBackendConnectionRequestBody`

NewWebBackendConnectionRequestBodyWithDefaults instantiates a new WebBackendConnectionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWithRefreshedCatalog

`func (o *WebBackendConnectionRequestBody) GetWithRefreshedCatalog() bool`

GetWithRefreshedCatalog returns the WithRefreshedCatalog field if non-nil, zero value otherwise.

### GetWithRefreshedCatalogOk

`func (o *WebBackendConnectionRequestBody) GetWithRefreshedCatalogOk() (*bool, bool)`

GetWithRefreshedCatalogOk returns a tuple with the WithRefreshedCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithRefreshedCatalog

`func (o *WebBackendConnectionRequestBody) SetWithRefreshedCatalog(v bool)`

SetWithRefreshedCatalog sets WithRefreshedCatalog field to given value.

### HasWithRefreshedCatalog

`func (o *WebBackendConnectionRequestBody) HasWithRefreshedCatalog() bool`

HasWithRefreshedCatalog returns a boolean if a field has been set.

### GetConnectionId

`func (o *WebBackendConnectionRequestBody) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *WebBackendConnectionRequestBody) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *WebBackendConnectionRequestBody) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


