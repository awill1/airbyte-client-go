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

// JobConfigType the model 'JobConfigType'
type JobConfigType string

// List of JobConfigType
const (
	JOBCONFIGTYPE_CHECK_CONNECTION_SOURCE JobConfigType = "check_connection_source"
	JOBCONFIGTYPE_CHECK_CONNECTION_DESTINATION JobConfigType = "check_connection_destination"
	JOBCONFIGTYPE_DISCOVER_SCHEMA JobConfigType = "discover_schema"
	JOBCONFIGTYPE_GET_SPEC JobConfigType = "get_spec"
	JOBCONFIGTYPE_SYNC JobConfigType = "sync"
	JOBCONFIGTYPE_RESET_CONNECTION JobConfigType = "reset_connection"
)

// All allowed values of JobConfigType enum
var AllowedJobConfigTypeEnumValues = []JobConfigType{
	"check_connection_source",
	"check_connection_destination",
	"discover_schema",
	"get_spec",
	"sync",
	"reset_connection",
}

func (v *JobConfigType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobConfigType(value)
	for _, existing := range AllowedJobConfigTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobConfigType", value)
}

// NewJobConfigTypeFromValue returns a pointer to a valid JobConfigType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobConfigTypeFromValue(v string) (*JobConfigType, error) {
	ev := JobConfigType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobConfigType: valid values are %v", v, AllowedJobConfigTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobConfigType) IsValid() bool {
	for _, existing := range AllowedJobConfigTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JobConfigType value
func (v JobConfigType) Ptr() *JobConfigType {
	return &v
}

type NullableJobConfigType struct {
	value *JobConfigType
	isSet bool
}

func (v NullableJobConfigType) Get() *JobConfigType {
	return v.value
}

func (v *NullableJobConfigType) Set(val *JobConfigType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobConfigType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobConfigType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobConfigType(val *JobConfigType) *NullableJobConfigType {
	return &NullableJobConfigType{value: val, isSet: true}
}

func (v NullableJobConfigType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobConfigType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

