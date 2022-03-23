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

// JobReadList struct for JobReadList
type JobReadList struct {
	Jobs []JobWithAttemptsRead `json:"jobs"`
}

// NewJobReadList instantiates a new JobReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobReadList(jobs []JobWithAttemptsRead) *JobReadList {
	this := JobReadList{}
	this.Jobs = jobs
	return &this
}

// NewJobReadListWithDefaults instantiates a new JobReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobReadListWithDefaults() *JobReadList {
	this := JobReadList{}
	return &this
}

// GetJobs returns the Jobs field value
func (o *JobReadList) GetJobs() []JobWithAttemptsRead {
	if o == nil {
		var ret []JobWithAttemptsRead
		return ret
	}

	return o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value
// and a boolean to check if the value has been set.
func (o *JobReadList) GetJobsOk() ([]JobWithAttemptsRead, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Jobs, true
}

// SetJobs sets field value
func (o *JobReadList) SetJobs(v []JobWithAttemptsRead) {
	o.Jobs = v
}

func (o JobReadList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobs"] = o.Jobs
	}
	return json.Marshal(toSerialize)
}

type NullableJobReadList struct {
	value *JobReadList
	isSet bool
}

func (v NullableJobReadList) Get() *JobReadList {
	return v.value
}

func (v *NullableJobReadList) Set(val *JobReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableJobReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableJobReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobReadList(val *JobReadList) *NullableJobReadList {
	return &NullableJobReadList{value: val, isSet: true}
}

func (v NullableJobReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


