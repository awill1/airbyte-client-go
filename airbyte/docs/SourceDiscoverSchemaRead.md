# SourceDiscoverSchemaRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Catalog** | Pointer to [**AirbyteCatalog**](AirbyteCatalog.md) |  | [optional] 
**JobInfo** | [**SynchronousJobRead**](SynchronousJobRead.md) |  | 

## Methods

### NewSourceDiscoverSchemaRead

`func NewSourceDiscoverSchemaRead(jobInfo SynchronousJobRead, ) *SourceDiscoverSchemaRead`

NewSourceDiscoverSchemaRead instantiates a new SourceDiscoverSchemaRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceDiscoverSchemaReadWithDefaults

`func NewSourceDiscoverSchemaReadWithDefaults() *SourceDiscoverSchemaRead`

NewSourceDiscoverSchemaReadWithDefaults instantiates a new SourceDiscoverSchemaRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalog

`func (o *SourceDiscoverSchemaRead) GetCatalog() AirbyteCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *SourceDiscoverSchemaRead) GetCatalogOk() (*AirbyteCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *SourceDiscoverSchemaRead) SetCatalog(v AirbyteCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *SourceDiscoverSchemaRead) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetJobInfo

`func (o *SourceDiscoverSchemaRead) GetJobInfo() SynchronousJobRead`

GetJobInfo returns the JobInfo field if non-nil, zero value otherwise.

### GetJobInfoOk

`func (o *SourceDiscoverSchemaRead) GetJobInfoOk() (*SynchronousJobRead, bool)`

GetJobInfoOk returns a tuple with the JobInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobInfo

`func (o *SourceDiscoverSchemaRead) SetJobInfo(v SynchronousJobRead)`

SetJobInfo sets JobInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


