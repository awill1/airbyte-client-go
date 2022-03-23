# JobListRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigTypes** | [**[]JobConfigType**](JobConfigType.md) |  | 
**ConfigId** | **string** |  | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewJobListRequestBody

`func NewJobListRequestBody(configTypes []JobConfigType, configId string, ) *JobListRequestBody`

NewJobListRequestBody instantiates a new JobListRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobListRequestBodyWithDefaults

`func NewJobListRequestBodyWithDefaults() *JobListRequestBody`

NewJobListRequestBodyWithDefaults instantiates a new JobListRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigTypes

`func (o *JobListRequestBody) GetConfigTypes() []JobConfigType`

GetConfigTypes returns the ConfigTypes field if non-nil, zero value otherwise.

### GetConfigTypesOk

`func (o *JobListRequestBody) GetConfigTypesOk() (*[]JobConfigType, bool)`

GetConfigTypesOk returns a tuple with the ConfigTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigTypes

`func (o *JobListRequestBody) SetConfigTypes(v []JobConfigType)`

SetConfigTypes sets ConfigTypes field to given value.


### GetConfigId

`func (o *JobListRequestBody) GetConfigId() string`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *JobListRequestBody) GetConfigIdOk() (*string, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *JobListRequestBody) SetConfigId(v string)`

SetConfigId sets ConfigId field to given value.


### GetPagination

`func (o *JobListRequestBody) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *JobListRequestBody) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *JobListRequestBody) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *JobListRequestBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


