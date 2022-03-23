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

// InvalidInputProperty struct for InvalidInputProperty
type InvalidInputProperty struct {
	PropertyPath string `json:"propertyPath"`
	InvalidValue *string `json:"invalidValue,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewInvalidInputProperty instantiates a new InvalidInputProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidInputProperty(propertyPath string) *InvalidInputProperty {
	this := InvalidInputProperty{}
	this.PropertyPath = propertyPath
	return &this
}

// NewInvalidInputPropertyWithDefaults instantiates a new InvalidInputProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidInputPropertyWithDefaults() *InvalidInputProperty {
	this := InvalidInputProperty{}
	return &this
}

// GetPropertyPath returns the PropertyPath field value
func (o *InvalidInputProperty) GetPropertyPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyPath
}

// GetPropertyPathOk returns a tuple with the PropertyPath field value
// and a boolean to check if the value has been set.
func (o *InvalidInputProperty) GetPropertyPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyPath, true
}

// SetPropertyPath sets field value
func (o *InvalidInputProperty) SetPropertyPath(v string) {
	o.PropertyPath = v
}

// GetInvalidValue returns the InvalidValue field value if set, zero value otherwise.
func (o *InvalidInputProperty) GetInvalidValue() string {
	if o == nil || o.InvalidValue == nil {
		var ret string
		return ret
	}
	return *o.InvalidValue
}

// GetInvalidValueOk returns a tuple with the InvalidValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidInputProperty) GetInvalidValueOk() (*string, bool) {
	if o == nil || o.InvalidValue == nil {
		return nil, false
	}
	return o.InvalidValue, true
}

// HasInvalidValue returns a boolean if a field has been set.
func (o *InvalidInputProperty) HasInvalidValue() bool {
	if o != nil && o.InvalidValue != nil {
		return true
	}

	return false
}

// SetInvalidValue gets a reference to the given string and assigns it to the InvalidValue field.
func (o *InvalidInputProperty) SetInvalidValue(v string) {
	o.InvalidValue = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InvalidInputProperty) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidInputProperty) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InvalidInputProperty) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InvalidInputProperty) SetMessage(v string) {
	o.Message = &v
}

func (o InvalidInputProperty) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["propertyPath"] = o.PropertyPath
	}
	if o.InvalidValue != nil {
		toSerialize["invalidValue"] = o.InvalidValue
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInvalidInputProperty struct {
	value *InvalidInputProperty
	isSet bool
}

func (v NullableInvalidInputProperty) Get() *InvalidInputProperty {
	return v.value
}

func (v *NullableInvalidInputProperty) Set(val *InvalidInputProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidInputProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidInputProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidInputProperty(val *InvalidInputProperty) *NullableInvalidInputProperty {
	return &NullableInvalidInputProperty{value: val, isSet: true}
}

func (v NullableInvalidInputProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidInputProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


