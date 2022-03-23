# OperatorConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorType** | [**OperatorType**](OperatorType.md) |  | 
**Normalization** | Pointer to [**OperatorNormalization**](OperatorNormalization.md) |  | [optional] 
**Dbt** | Pointer to [**OperatorDbt**](OperatorDbt.md) |  | [optional] 

## Methods

### NewOperatorConfiguration

`func NewOperatorConfiguration(operatorType OperatorType, ) *OperatorConfiguration`

NewOperatorConfiguration instantiates a new OperatorConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorConfigurationWithDefaults

`func NewOperatorConfigurationWithDefaults() *OperatorConfiguration`

NewOperatorConfigurationWithDefaults instantiates a new OperatorConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatorType

`func (o *OperatorConfiguration) GetOperatorType() OperatorType`

GetOperatorType returns the OperatorType field if non-nil, zero value otherwise.

### GetOperatorTypeOk

`func (o *OperatorConfiguration) GetOperatorTypeOk() (*OperatorType, bool)`

GetOperatorTypeOk returns a tuple with the OperatorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorType

`func (o *OperatorConfiguration) SetOperatorType(v OperatorType)`

SetOperatorType sets OperatorType field to given value.


### GetNormalization

`func (o *OperatorConfiguration) GetNormalization() OperatorNormalization`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *OperatorConfiguration) GetNormalizationOk() (*OperatorNormalization, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *OperatorConfiguration) SetNormalization(v OperatorNormalization)`

SetNormalization sets Normalization field to given value.

### HasNormalization

`func (o *OperatorConfiguration) HasNormalization() bool`

HasNormalization returns a boolean if a field has been set.

### GetDbt

`func (o *OperatorConfiguration) GetDbt() OperatorDbt`

GetDbt returns the Dbt field if non-nil, zero value otherwise.

### GetDbtOk

`func (o *OperatorConfiguration) GetDbtOk() (*OperatorDbt, bool)`

GetDbtOk returns a tuple with the Dbt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbt

`func (o *OperatorConfiguration) SetDbt(v OperatorDbt)`

SetDbt sets Dbt field to given value.

### HasDbt

`func (o *OperatorConfiguration) HasDbt() bool`

HasDbt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


