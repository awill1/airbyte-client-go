# AdvancedAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthFlowType** | Pointer to **string** |  | [optional] 
**PredicateKey** | Pointer to **[]string** | Json Path to a field in the connectorSpecification that should exist for the advanced auth to be applicable. | [optional] 
**PredicateValue** | Pointer to **string** | Value of the predicate_key fields for the advanced auth to be applicable. | [optional] 
**OauthConfigSpecification** | Pointer to [**OAuthConfigSpecification**](OAuthConfigSpecification.md) |  | [optional] 

## Methods

### NewAdvancedAuth

`func NewAdvancedAuth() *AdvancedAuth`

NewAdvancedAuth instantiates a new AdvancedAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdvancedAuthWithDefaults

`func NewAdvancedAuthWithDefaults() *AdvancedAuth`

NewAdvancedAuthWithDefaults instantiates a new AdvancedAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthFlowType

`func (o *AdvancedAuth) GetAuthFlowType() string`

GetAuthFlowType returns the AuthFlowType field if non-nil, zero value otherwise.

### GetAuthFlowTypeOk

`func (o *AdvancedAuth) GetAuthFlowTypeOk() (*string, bool)`

GetAuthFlowTypeOk returns a tuple with the AuthFlowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFlowType

`func (o *AdvancedAuth) SetAuthFlowType(v string)`

SetAuthFlowType sets AuthFlowType field to given value.

### HasAuthFlowType

`func (o *AdvancedAuth) HasAuthFlowType() bool`

HasAuthFlowType returns a boolean if a field has been set.

### GetPredicateKey

`func (o *AdvancedAuth) GetPredicateKey() []string`

GetPredicateKey returns the PredicateKey field if non-nil, zero value otherwise.

### GetPredicateKeyOk

`func (o *AdvancedAuth) GetPredicateKeyOk() (*[]string, bool)`

GetPredicateKeyOk returns a tuple with the PredicateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateKey

`func (o *AdvancedAuth) SetPredicateKey(v []string)`

SetPredicateKey sets PredicateKey field to given value.

### HasPredicateKey

`func (o *AdvancedAuth) HasPredicateKey() bool`

HasPredicateKey returns a boolean if a field has been set.

### GetPredicateValue

`func (o *AdvancedAuth) GetPredicateValue() string`

GetPredicateValue returns the PredicateValue field if non-nil, zero value otherwise.

### GetPredicateValueOk

`func (o *AdvancedAuth) GetPredicateValueOk() (*string, bool)`

GetPredicateValueOk returns a tuple with the PredicateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicateValue

`func (o *AdvancedAuth) SetPredicateValue(v string)`

SetPredicateValue sets PredicateValue field to given value.

### HasPredicateValue

`func (o *AdvancedAuth) HasPredicateValue() bool`

HasPredicateValue returns a boolean if a field has been set.

### GetOauthConfigSpecification

`func (o *AdvancedAuth) GetOauthConfigSpecification() OAuthConfigSpecification`

GetOauthConfigSpecification returns the OauthConfigSpecification field if non-nil, zero value otherwise.

### GetOauthConfigSpecificationOk

`func (o *AdvancedAuth) GetOauthConfigSpecificationOk() (*OAuthConfigSpecification, bool)`

GetOauthConfigSpecificationOk returns a tuple with the OauthConfigSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthConfigSpecification

`func (o *AdvancedAuth) SetOauthConfigSpecification(v OAuthConfigSpecification)`

SetOauthConfigSpecification sets OauthConfigSpecification field to given value.

### HasOauthConfigSpecification

`func (o *AdvancedAuth) HasOauthConfigSpecification() bool`

HasOauthConfigSpecification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


