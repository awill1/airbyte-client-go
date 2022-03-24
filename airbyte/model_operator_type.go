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
	"fmt"
)

// OperatorType the model 'OperatorType'
type OperatorType string

// List of OperatorType
const (
	OPERATORTYPE_NORMALIZATION OperatorType = "normalization"
	OPERATORTYPE_DBT OperatorType = "dbt"
)

// All allowed values of OperatorType enum
var AllowedOperatorTypeEnumValues = []OperatorType{
	"normalization",
	"dbt",
}

func (v *OperatorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperatorType(value)
	for _, existing := range AllowedOperatorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperatorType", value)
}

// NewOperatorTypeFromValue returns a pointer to a valid OperatorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperatorTypeFromValue(v string) (*OperatorType, error) {
	ev := OperatorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperatorType: valid values are %v", v, AllowedOperatorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperatorType) IsValid() bool {
	for _, existing := range AllowedOperatorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperatorType value
func (v OperatorType) Ptr() *OperatorType {
	return &v
}

type NullableOperatorType struct {
	value *OperatorType
	isSet bool
}

func (v NullableOperatorType) Get() *OperatorType {
	return v.value
}

func (v *NullableOperatorType) Set(val *OperatorType) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorType) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorType(val *OperatorType) *NullableOperatorType {
	return &NullableOperatorType{value: val, isSet: true}
}

func (v NullableOperatorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

