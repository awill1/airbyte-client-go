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

// SynchronousJobRead struct for SynchronousJobRead
type SynchronousJobRead struct {
	Id string `json:"id"`
	ConfigType JobConfigType `json:"configType"`
	// only present if a config id was provided.
	ConfigId *string `json:"configId,omitempty"`
	CreatedAt int64 `json:"createdAt"`
	EndedAt int64 `json:"endedAt"`
	Succeeded bool `json:"succeeded"`
	Logs *LogRead `json:"logs,omitempty"`
}

// NewSynchronousJobRead instantiates a new SynchronousJobRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSynchronousJobRead(id string, configType JobConfigType, createdAt int64, endedAt int64, succeeded bool) *SynchronousJobRead {
	this := SynchronousJobRead{}
	this.Id = id
	this.ConfigType = configType
	this.CreatedAt = createdAt
	this.EndedAt = endedAt
	this.Succeeded = succeeded
	return &this
}

// NewSynchronousJobReadWithDefaults instantiates a new SynchronousJobRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSynchronousJobReadWithDefaults() *SynchronousJobRead {
	this := SynchronousJobRead{}
	return &this
}

// GetId returns the Id field value
func (o *SynchronousJobRead) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SynchronousJobRead) SetId(v string) {
	o.Id = v
}

// GetConfigType returns the ConfigType field value
func (o *SynchronousJobRead) GetConfigType() JobConfigType {
	if o == nil {
		var ret JobConfigType
		return ret
	}

	return o.ConfigType
}

// GetConfigTypeOk returns a tuple with the ConfigType field value
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetConfigTypeOk() (*JobConfigType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigType, true
}

// SetConfigType sets field value
func (o *SynchronousJobRead) SetConfigType(v JobConfigType) {
	o.ConfigType = v
}

// GetConfigId returns the ConfigId field value if set, zero value otherwise.
func (o *SynchronousJobRead) GetConfigId() string {
	if o == nil || o.ConfigId == nil {
		var ret string
		return ret
	}
	return *o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetConfigIdOk() (*string, bool) {
	if o == nil || o.ConfigId == nil {
		return nil, false
	}
	return o.ConfigId, true
}

// HasConfigId returns a boolean if a field has been set.
func (o *SynchronousJobRead) HasConfigId() bool {
	if o != nil && o.ConfigId != nil {
		return true
	}

	return false
}

// SetConfigId gets a reference to the given string and assigns it to the ConfigId field.
func (o *SynchronousJobRead) SetConfigId(v string) {
	o.ConfigId = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SynchronousJobRead) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetCreatedAtOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SynchronousJobRead) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetEndedAt returns the EndedAt field value
func (o *SynchronousJobRead) GetEndedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetEndedAtOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EndedAt, true
}

// SetEndedAt sets field value
func (o *SynchronousJobRead) SetEndedAt(v int64) {
	o.EndedAt = v
}

// GetSucceeded returns the Succeeded field value
func (o *SynchronousJobRead) GetSucceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetSucceededOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Succeeded, true
}

// SetSucceeded sets field value
func (o *SynchronousJobRead) SetSucceeded(v bool) {
	o.Succeeded = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *SynchronousJobRead) GetLogs() LogRead {
	if o == nil || o.Logs == nil {
		var ret LogRead
		return ret
	}
	return *o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SynchronousJobRead) GetLogsOk() (*LogRead, bool) {
	if o == nil || o.Logs == nil {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *SynchronousJobRead) HasLogs() bool {
	if o != nil && o.Logs != nil {
		return true
	}

	return false
}

// SetLogs gets a reference to the given LogRead and assigns it to the Logs field.
func (o *SynchronousJobRead) SetLogs(v LogRead) {
	o.Logs = &v
}

func (o SynchronousJobRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["configType"] = o.ConfigType
	}
	if o.ConfigId != nil {
		toSerialize["configId"] = o.ConfigId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["endedAt"] = o.EndedAt
	}
	if true {
		toSerialize["succeeded"] = o.Succeeded
	}
	if o.Logs != nil {
		toSerialize["logs"] = o.Logs
	}
	return json.Marshal(toSerialize)
}

type NullableSynchronousJobRead struct {
	value *SynchronousJobRead
	isSet bool
}

func (v NullableSynchronousJobRead) Get() *SynchronousJobRead {
	return v.value
}

func (v *NullableSynchronousJobRead) Set(val *SynchronousJobRead) {
	v.value = val
	v.isSet = true
}

func (v NullableSynchronousJobRead) IsSet() bool {
	return v.isSet
}

func (v *NullableSynchronousJobRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSynchronousJobRead(val *SynchronousJobRead) *NullableSynchronousJobRead {
	return &NullableSynchronousJobRead{value: val, isSet: true}
}

func (v NullableSynchronousJobRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSynchronousJobRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


