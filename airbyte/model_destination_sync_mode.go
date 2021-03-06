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

// DestinationSyncMode the model 'DestinationSyncMode'
type DestinationSyncMode string

// List of DestinationSyncMode
const (
	DESTINATIONSYNCMODE_APPEND DestinationSyncMode = "append"
	DESTINATIONSYNCMODE_OVERWRITE DestinationSyncMode = "overwrite"
	DESTINATIONSYNCMODE_APPEND_DEDUP DestinationSyncMode = "append_dedup"
)

// All allowed values of DestinationSyncMode enum
var AllowedDestinationSyncModeEnumValues = []DestinationSyncMode{
	"append",
	"overwrite",
	"append_dedup",
}

func (v *DestinationSyncMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DestinationSyncMode(value)
	for _, existing := range AllowedDestinationSyncModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DestinationSyncMode", value)
}

// NewDestinationSyncModeFromValue returns a pointer to a valid DestinationSyncMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDestinationSyncModeFromValue(v string) (*DestinationSyncMode, error) {
	ev := DestinationSyncMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DestinationSyncMode: valid values are %v", v, AllowedDestinationSyncModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DestinationSyncMode) IsValid() bool {
	for _, existing := range AllowedDestinationSyncModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DestinationSyncMode value
func (v DestinationSyncMode) Ptr() *DestinationSyncMode {
	return &v
}

type NullableDestinationSyncMode struct {
	value *DestinationSyncMode
	isSet bool
}

func (v NullableDestinationSyncMode) Get() *DestinationSyncMode {
	return v.value
}

func (v *NullableDestinationSyncMode) Set(val *DestinationSyncMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationSyncMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationSyncMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationSyncMode(val *DestinationSyncMode) *NullableDestinationSyncMode {
	return &NullableDestinationSyncMode{value: val, isSet: true}
}

func (v NullableDestinationSyncMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationSyncMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

