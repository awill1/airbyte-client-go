/*
Airbyte Configuration API

Airbyte Configuration API [https://airbyte.io](https://airbyte.io).  This API is a collection of HTTP RPC-style methods. While it is not a REST API, those familiar with REST should find the conventions of this API recognizable.  Here are some conventions that this API follows: * All endpoints are http POST methods. * All endpoints accept data via `application/json` request bodies. The API does not accept any data via query params. * The naming convention for endpoints is: localhost:8000/{VERSION}/{METHOD_FAMILY}/{METHOD_NAME} e.g. `localhost:8000/v1/connections/create`. * For all `update` methods, the whole object must be passed in, even the fields that did not change.  Change Management: * The major version of the API endpoint can be determined / specified in the URL `localhost:8080/v1/connections/create` * Minor version bumps will be invisible to the end user. The user cannot specify minor versions in requests. * All backwards incompatible changes will happen in major version bumps. We will not make backwards incompatible changes in minor version bumps. Examples of non-breaking changes (includes but not limited to...):   * Adding fields to request or response bodies.   * Adding new HTTP endpoints. * All `web_backend` APIs are not considered public APIs and are not guaranteeing backwards compatibility. 

API version: 1.0.0
Contact: contact@airbyte.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package airbyte

import (
	"encoding/json"
)

// AuthSpecification struct for AuthSpecification
type AuthSpecification struct {
	AuthType *string `json:"auth_type,omitempty"`
	Oauth2Specification *OAuth2Specification `json:"oauth2Specification,omitempty"`
}

// NewAuthSpecification instantiates a new AuthSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthSpecification() *AuthSpecification {
	this := AuthSpecification{}
	return &this
}

// NewAuthSpecificationWithDefaults instantiates a new AuthSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthSpecificationWithDefaults() *AuthSpecification {
	this := AuthSpecification{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *AuthSpecification) GetAuthType() string {
	if o == nil || o.AuthType == nil {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSpecification) GetAuthTypeOk() (*string, bool) {
	if o == nil || o.AuthType == nil {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *AuthSpecification) HasAuthType() bool {
	if o != nil && o.AuthType != nil {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *AuthSpecification) SetAuthType(v string) {
	o.AuthType = &v
}

// GetOauth2Specification returns the Oauth2Specification field value if set, zero value otherwise.
func (o *AuthSpecification) GetOauth2Specification() OAuth2Specification {
	if o == nil || o.Oauth2Specification == nil {
		var ret OAuth2Specification
		return ret
	}
	return *o.Oauth2Specification
}

// GetOauth2SpecificationOk returns a tuple with the Oauth2Specification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthSpecification) GetOauth2SpecificationOk() (*OAuth2Specification, bool) {
	if o == nil || o.Oauth2Specification == nil {
		return nil, false
	}
	return o.Oauth2Specification, true
}

// HasOauth2Specification returns a boolean if a field has been set.
func (o *AuthSpecification) HasOauth2Specification() bool {
	if o != nil && o.Oauth2Specification != nil {
		return true
	}

	return false
}

// SetOauth2Specification gets a reference to the given OAuth2Specification and assigns it to the Oauth2Specification field.
func (o *AuthSpecification) SetOauth2Specification(v OAuth2Specification) {
	o.Oauth2Specification = &v
}

func (o AuthSpecification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthType != nil {
		toSerialize["auth_type"] = o.AuthType
	}
	if o.Oauth2Specification != nil {
		toSerialize["oauth2Specification"] = o.Oauth2Specification
	}
	return json.Marshal(toSerialize)
}

type NullableAuthSpecification struct {
	value *AuthSpecification
	isSet bool
}

func (v NullableAuthSpecification) Get() *AuthSpecification {
	return v.value
}

func (v *NullableAuthSpecification) Set(val *AuthSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthSpecification(val *AuthSpecification) *NullableAuthSpecification {
	return &NullableAuthSpecification{value: val, isSet: true}
}

func (v NullableAuthSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


