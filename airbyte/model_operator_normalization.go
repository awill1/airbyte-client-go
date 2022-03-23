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

// OperatorNormalization struct for OperatorNormalization
type OperatorNormalization struct {
	Option *string `json:"option,omitempty"`
}

// NewOperatorNormalization instantiates a new OperatorNormalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorNormalization() *OperatorNormalization {
	this := OperatorNormalization{}
	return &this
}

// NewOperatorNormalizationWithDefaults instantiates a new OperatorNormalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorNormalizationWithDefaults() *OperatorNormalization {
	this := OperatorNormalization{}
	return &this
}

// GetOption returns the Option field value if set, zero value otherwise.
func (o *OperatorNormalization) GetOption() string {
	if o == nil || o.Option == nil {
		var ret string
		return ret
	}
	return *o.Option
}

// GetOptionOk returns a tuple with the Option field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorNormalization) GetOptionOk() (*string, bool) {
	if o == nil || o.Option == nil {
		return nil, false
	}
	return o.Option, true
}

// HasOption returns a boolean if a field has been set.
func (o *OperatorNormalization) HasOption() bool {
	if o != nil && o.Option != nil {
		return true
	}

	return false
}

// SetOption gets a reference to the given string and assigns it to the Option field.
func (o *OperatorNormalization) SetOption(v string) {
	o.Option = &v
}

func (o OperatorNormalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Option != nil {
		toSerialize["option"] = o.Option
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorNormalization struct {
	value *OperatorNormalization
	isSet bool
}

func (v NullableOperatorNormalization) Get() *OperatorNormalization {
	return v.value
}

func (v *NullableOperatorNormalization) Set(val *OperatorNormalization) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorNormalization) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorNormalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorNormalization(val *OperatorNormalization) *NullableOperatorNormalization {
	return &NullableOperatorNormalization{value: val, isSet: true}
}

func (v NullableOperatorNormalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorNormalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

