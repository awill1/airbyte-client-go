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

// ConnectionState struct for ConnectionState
type ConnectionState struct {
	ConnectionId string `json:"connectionId"`
	State map[string]interface{} `json:"state,omitempty"`
}

// NewConnectionState instantiates a new ConnectionState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionState(connectionId string) *ConnectionState {
	this := ConnectionState{}
	this.ConnectionId = connectionId
	return &this
}

// NewConnectionStateWithDefaults instantiates a new ConnectionState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionStateWithDefaults() *ConnectionState {
	this := ConnectionState{}
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *ConnectionState) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ConnectionState) GetConnectionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *ConnectionState) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ConnectionState) GetState() map[string]interface{} {
	if o == nil || o.State == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionState) GetStateOk() (map[string]interface{}, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ConnectionState) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given map[string]interface{} and assigns it to the State field.
func (o *ConnectionState) SetState(v map[string]interface{}) {
	o.State = v
}

func (o ConnectionState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionState struct {
	value *ConnectionState
	isSet bool
}

func (v NullableConnectionState) Get() *ConnectionState {
	return v.value
}

func (v *NullableConnectionState) Set(val *ConnectionState) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionState) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionState(val *ConnectionState) *NullableConnectionState {
	return &NullableConnectionState{value: val, isSet: true}
}

func (v NullableConnectionState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

