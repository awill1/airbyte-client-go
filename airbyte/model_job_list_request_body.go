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

// JobListRequestBody struct for JobListRequestBody
type JobListRequestBody struct {
	ConfigTypes []JobConfigType `json:"configTypes"`
	ConfigId string `json:"configId"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// NewJobListRequestBody instantiates a new JobListRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobListRequestBody(configTypes []JobConfigType, configId string) *JobListRequestBody {
	this := JobListRequestBody{}
	this.ConfigTypes = configTypes
	this.ConfigId = configId
	return &this
}

// NewJobListRequestBodyWithDefaults instantiates a new JobListRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobListRequestBodyWithDefaults() *JobListRequestBody {
	this := JobListRequestBody{}
	return &this
}

// GetConfigTypes returns the ConfigTypes field value
func (o *JobListRequestBody) GetConfigTypes() []JobConfigType {
	if o == nil {
		var ret []JobConfigType
		return ret
	}

	return o.ConfigTypes
}

// GetConfigTypesOk returns a tuple with the ConfigTypes field value
// and a boolean to check if the value has been set.
func (o *JobListRequestBody) GetConfigTypesOk() ([]JobConfigType, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConfigTypes, true
}

// SetConfigTypes sets field value
func (o *JobListRequestBody) SetConfigTypes(v []JobConfigType) {
	o.ConfigTypes = v
}

// GetConfigId returns the ConfigId field value
func (o *JobListRequestBody) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *JobListRequestBody) GetConfigIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value
func (o *JobListRequestBody) SetConfigId(v string) {
	o.ConfigId = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *JobListRequestBody) GetPagination() Pagination {
	if o == nil || o.Pagination == (Pagination{}) {
		var ret Pagination
		return ret
	}
	return o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobListRequestBody) GetPaginationOk() (Pagination, bool) {
	if o == nil || o.Pagination == (Pagination{}) {
		return Pagination{}, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *JobListRequestBody) HasPagination() bool {
	if o != nil && o.Pagination != (Pagination{}) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *JobListRequestBody) SetPagination(v Pagination) {
	o.Pagination = v
}

func (o JobListRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["configTypes"] = o.ConfigTypes
	}
	if true {
		toSerialize["configId"] = o.ConfigId
	}
	if o.Pagination != (Pagination{}) {
		toSerialize["pagination"] = o.Pagination
	}
	return json.Marshal(toSerialize)
}

type NullableJobListRequestBody struct {
	value *JobListRequestBody
	isSet bool
}

func (v NullableJobListRequestBody) Get() *JobListRequestBody {
	return v.value
}

func (v *NullableJobListRequestBody) Set(val *JobListRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableJobListRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableJobListRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobListRequestBody(val *JobListRequestBody) *NullableJobListRequestBody {
	return &NullableJobListRequestBody{value: val, isSet: true}
}

func (v NullableJobListRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobListRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


