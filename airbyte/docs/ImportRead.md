# ImportRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewImportRead

`func NewImportRead(status string, ) *ImportRead`

NewImportRead instantiates a new ImportRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportReadWithDefaults

`func NewImportReadWithDefaults() *ImportRead`

NewImportReadWithDefaults instantiates a new ImportRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ImportRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImportRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImportRead) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ImportRead) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ImportRead) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ImportRead) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ImportRead) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


