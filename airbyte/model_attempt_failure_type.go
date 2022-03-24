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

// AttemptFailureType Categorizes well known errors into types for programmatic handling. If not set, the type of error is not well known.
type AttemptFailureType string

// List of AttemptFailureType
const (
	ATTEMPTFAILURETYPE_UNKNOWN AttemptFailureType = "unknown"
	ATTEMPTFAILURETYPE_CONFIG_ERROR AttemptFailureType = "config_error"
	ATTEMPTFAILURETYPE_SYSTEM_ERROR AttemptFailureType = "system_error"
	ATTEMPTFAILURETYPE_MANUAL_CANCELLATION AttemptFailureType = "manual_cancellation"
)

// All allowed values of AttemptFailureType enum
var AllowedAttemptFailureTypeEnumValues = []AttemptFailureType{
	"unknown",
	"config_error",
	"system_error",
	"manual_cancellation",
}

func (v *AttemptFailureType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AttemptFailureType(value)
	for _, existing := range AllowedAttemptFailureTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AttemptFailureType", value)
}

// NewAttemptFailureTypeFromValue returns a pointer to a valid AttemptFailureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAttemptFailureTypeFromValue(v string) (*AttemptFailureType, error) {
	ev := AttemptFailureType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AttemptFailureType: valid values are %v", v, AllowedAttemptFailureTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AttemptFailureType) IsValid() bool {
	for _, existing := range AllowedAttemptFailureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttemptFailureType value
func (v AttemptFailureType) Ptr() *AttemptFailureType {
	return &v
}

type NullableAttemptFailureType struct {
	value *AttemptFailureType
	isSet bool
}

func (v NullableAttemptFailureType) Get() *AttemptFailureType {
	return v.value
}

func (v *NullableAttemptFailureType) Set(val *AttemptFailureType) {
	v.value = val
	v.isSet = true
}

func (v NullableAttemptFailureType) IsSet() bool {
	return v.isSet
}

func (v *NullableAttemptFailureType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttemptFailureType(val *AttemptFailureType) *NullableAttemptFailureType {
	return &NullableAttemptFailureType{value: val, isSet: true}
}

func (v NullableAttemptFailureType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttemptFailureType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

