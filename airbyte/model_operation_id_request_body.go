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

// OperationIdRequestBody struct for OperationIdRequestBody
type OperationIdRequestBody struct {
	OperationId string `json:"OperationId"`
}

// NewOperationIdRequestBody instantiates a new OperationIdRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationIdRequestBody(operationId string) *OperationIdRequestBody {
	this := OperationIdRequestBody{}
	this.OperationId = operationId
	return &this
}

// NewOperationIdRequestBodyWithDefaults instantiates a new OperationIdRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationIdRequestBodyWithDefaults() *OperationIdRequestBody {
	this := OperationIdRequestBody{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *OperationIdRequestBody) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *OperationIdRequestBody) GetOperationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *OperationIdRequestBody) SetOperationId(v string) {
	o.OperationId = v
}

func (o OperationIdRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["OperationId"] = o.OperationId
	}
	return json.Marshal(toSerialize)
}

type NullableOperationIdRequestBody struct {
	value *OperationIdRequestBody
	isSet bool
}

func (v NullableOperationIdRequestBody) Get() *OperationIdRequestBody {
	return v.value
}

func (v *NullableOperationIdRequestBody) Set(val *OperationIdRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationIdRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationIdRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationIdRequestBody(val *OperationIdRequestBody) *NullableOperationIdRequestBody {
	return &NullableOperationIdRequestBody{value: val, isSet: true}
}

func (v NullableOperationIdRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationIdRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


