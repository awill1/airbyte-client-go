# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageSize** | Pointer to **int32** |  | [optional] 
**RowOffset** | Pointer to **int32** |  | [optional] 

## Methods

### NewPagination

`func NewPagination() *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageSize

`func (o *Pagination) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Pagination) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Pagination) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Pagination) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRowOffset

`func (o *Pagination) GetRowOffset() int32`

GetRowOffset returns the RowOffset field if non-nil, zero value otherwise.

### GetRowOffsetOk

`func (o *Pagination) GetRowOffsetOk() (*int32, bool)`

GetRowOffsetOk returns a tuple with the RowOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowOffset

`func (o *Pagination) SetRowOffset(v int32)`

SetRowOffset sets RowOffset field to given value.

### HasRowOffset

`func (o *Pagination) HasRowOffset() bool`

HasRowOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


