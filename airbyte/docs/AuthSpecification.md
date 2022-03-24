# AuthSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **string** |  | [optional] 
**Oauth2Specification** | Pointer to [**OAuth2Specification**](OAuth2Specification.md) |  | [optional] 

## Methods

### NewAuthSpecification

`func NewAuthSpecification() *AuthSpecification`

NewAuthSpecification instantiates a new AuthSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSpecificationWithDefaults

`func NewAuthSpecificationWithDefaults() *AuthSpecification`

NewAuthSpecificationWithDefaults instantiates a new AuthSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *AuthSpecification) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthSpecification) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthSpecification) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthSpecification) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetOauth2Specification

`func (o *AuthSpecification) GetOauth2Specification() OAuth2Specification`

GetOauth2Specification returns the Oauth2Specification field if non-nil, zero value otherwise.

### GetOauth2SpecificationOk

`func (o *AuthSpecification) GetOauth2SpecificationOk() (*OAuth2Specification, bool)`

GetOauth2SpecificationOk returns a tuple with the Oauth2Specification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Specification

`func (o *AuthSpecification) SetOauth2Specification(v OAuth2Specification)`

SetOauth2Specification sets Oauth2Specification field to given value.

### HasOauth2Specification

`func (o *AuthSpecification) HasOauth2Specification() bool`

HasOauth2Specification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


