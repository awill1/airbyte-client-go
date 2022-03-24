# AirbyteCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Streams** | [**[]AirbyteStreamAndConfiguration**](AirbyteStreamAndConfiguration.md) |  | 

## Methods

### NewAirbyteCatalog

`func NewAirbyteCatalog(streams []AirbyteStreamAndConfiguration, ) *AirbyteCatalog`

NewAirbyteCatalog instantiates a new AirbyteCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirbyteCatalogWithDefaults

`func NewAirbyteCatalogWithDefaults() *AirbyteCatalog`

NewAirbyteCatalogWithDefaults instantiates a new AirbyteCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreams

`func (o *AirbyteCatalog) GetStreams() []AirbyteStreamAndConfiguration`

GetStreams returns the Streams field if non-nil, zero value otherwise.

### GetStreamsOk

`func (o *AirbyteCatalog) GetStreamsOk() (*[]AirbyteStreamAndConfiguration, bool)`

GetStreamsOk returns a tuple with the Streams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreams

`func (o *AirbyteCatalog) SetStreams(v []AirbyteStreamAndConfiguration)`

SetStreams sets Streams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


