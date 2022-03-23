# AirbyteStreamAndConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | Pointer to [**AirbyteStream**](AirbyteStream.md) |  | [optional] 
**Config** | Pointer to [**AirbyteStreamConfiguration**](AirbyteStreamConfiguration.md) |  | [optional] 

## Methods

### NewAirbyteStreamAndConfiguration

`func NewAirbyteStreamAndConfiguration() *AirbyteStreamAndConfiguration`

NewAirbyteStreamAndConfiguration instantiates a new AirbyteStreamAndConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAirbyteStreamAndConfigurationWithDefaults

`func NewAirbyteStreamAndConfigurationWithDefaults() *AirbyteStreamAndConfiguration`

NewAirbyteStreamAndConfigurationWithDefaults instantiates a new AirbyteStreamAndConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *AirbyteStreamAndConfiguration) GetStream() AirbyteStream`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AirbyteStreamAndConfiguration) GetStreamOk() (*AirbyteStream, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AirbyteStreamAndConfiguration) SetStream(v AirbyteStream)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AirbyteStreamAndConfiguration) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetConfig

`func (o *AirbyteStreamAndConfiguration) GetConfig() AirbyteStreamConfiguration`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AirbyteStreamAndConfiguration) GetConfigOk() (*AirbyteStreamConfiguration, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AirbyteStreamAndConfiguration) SetConfig(v AirbyteStreamConfiguration)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AirbyteStreamAndConfiguration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


