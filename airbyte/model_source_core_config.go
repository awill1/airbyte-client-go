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

// SourceCoreConfig struct for SourceCoreConfig
type SourceCoreConfig struct {
	SourceDefinitionId string `json:"sourceDefinitionId"`
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
}

// NewSourceCoreConfig instantiates a new SourceCoreConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceCoreConfig(sourceDefinitionId string, connectionConfiguration interface{}) *SourceCoreConfig {
	this := SourceCoreConfig{}
	this.SourceDefinitionId = sourceDefinitionId
	this.ConnectionConfiguration = connectionConfiguration
	return &this
}

// NewSourceCoreConfigWithDefaults instantiates a new SourceCoreConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceCoreConfigWithDefaults() *SourceCoreConfig {
	this := SourceCoreConfig{}
	return &this
}

// GetSourceDefinitionId returns the SourceDefinitionId field value
func (o *SourceCoreConfig) GetSourceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceDefinitionId
}

// GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *SourceCoreConfig) GetSourceDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceDefinitionId, true
}

// SetSourceDefinitionId sets field value
func (o *SourceCoreConfig) SetSourceDefinitionId(v string) {
	o.SourceDefinitionId = v
}

// GetConnectionConfiguration returns the ConnectionConfiguration field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *SourceCoreConfig) GetConnectionConfiguration() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.ConnectionConfiguration
}

// GetConnectionConfigurationOk returns a tuple with the ConnectionConfiguration field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SourceCoreConfig) GetConnectionConfigurationOk() (*interface{}, bool) {
	if o == nil || o.ConnectionConfiguration == nil {
		return nil, false
	}
	return &o.ConnectionConfiguration, true
}

// SetConnectionConfiguration sets field value
func (o *SourceCoreConfig) SetConnectionConfiguration(v interface{}) {
	o.ConnectionConfiguration = v
}

func (o SourceCoreConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceDefinitionId"] = o.SourceDefinitionId
	}
	if o.ConnectionConfiguration != nil {
		toSerialize["connectionConfiguration"] = o.ConnectionConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableSourceCoreConfig struct {
	value *SourceCoreConfig
	isSet bool
}

func (v NullableSourceCoreConfig) Get() *SourceCoreConfig {
	return v.value
}

func (v *NullableSourceCoreConfig) Set(val *SourceCoreConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceCoreConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceCoreConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceCoreConfig(val *SourceCoreConfig) *NullableSourceCoreConfig {
	return &NullableSourceCoreConfig{value: val, isSet: true}
}

func (v NullableSourceCoreConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceCoreConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


