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

// JobWithAttemptsRead struct for JobWithAttemptsRead
type JobWithAttemptsRead struct {
	Job *JobRead `json:"job,omitempty"`
	Attempts []AttemptRead `json:"attempts,omitempty"`
}

// NewJobWithAttemptsRead instantiates a new JobWithAttemptsRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobWithAttemptsRead() *JobWithAttemptsRead {
	this := JobWithAttemptsRead{}
	return &this
}

// NewJobWithAttemptsReadWithDefaults instantiates a new JobWithAttemptsRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithAttemptsReadWithDefaults() *JobWithAttemptsRead {
	this := JobWithAttemptsRead{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *JobWithAttemptsRead) GetJob() JobRead {
	if o == nil || o.Job == nil {
		var ret JobRead
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobWithAttemptsRead) GetJobOk() (*JobRead, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *JobWithAttemptsRead) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given JobRead and assigns it to the Job field.
func (o *JobWithAttemptsRead) SetJob(v JobRead) {
	o.Job = &v
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *JobWithAttemptsRead) GetAttempts() []AttemptRead {
	if o == nil || o.Attempts == nil {
		var ret []AttemptRead
		return ret
	}
	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobWithAttemptsRead) GetAttemptsOk() ([]AttemptRead, bool) {
	if o == nil || o.Attempts == nil {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *JobWithAttemptsRead) HasAttempts() bool {
	if o != nil && o.Attempts != nil {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given []AttemptRead and assigns it to the Attempts field.
func (o *JobWithAttemptsRead) SetAttempts(v []AttemptRead) {
	o.Attempts = v
}

func (o JobWithAttemptsRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	if o.Attempts != nil {
		toSerialize["attempts"] = o.Attempts
	}
	return json.Marshal(toSerialize)
}

type NullableJobWithAttemptsRead struct {
	value *JobWithAttemptsRead
	isSet bool
}

func (v NullableJobWithAttemptsRead) Get() *JobWithAttemptsRead {
	return v.value
}

func (v *NullableJobWithAttemptsRead) Set(val *JobWithAttemptsRead) {
	v.value = val
	v.isSet = true
}

func (v NullableJobWithAttemptsRead) IsSet() bool {
	return v.isSet
}

func (v *NullableJobWithAttemptsRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobWithAttemptsRead(val *JobWithAttemptsRead) *NullableJobWithAttemptsRead {
	return &NullableJobWithAttemptsRead{value: val, isSet: true}
}

func (v NullableJobWithAttemptsRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobWithAttemptsRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


