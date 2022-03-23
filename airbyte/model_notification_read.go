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

// NotificationRead struct for NotificationRead
type NotificationRead struct {
	Status string `json:"status"`
	Message *string `json:"message,omitempty"`
}

// NewNotificationRead instantiates a new NotificationRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationRead(status string) *NotificationRead {
	this := NotificationRead{}
	this.Status = status
	return &this
}

// NewNotificationReadWithDefaults instantiates a new NotificationRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationReadWithDefaults() *NotificationRead {
	this := NotificationRead{}
	return &this
}

// GetStatus returns the Status field value
func (o *NotificationRead) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *NotificationRead) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *NotificationRead) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotificationRead) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationRead) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotificationRead) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotificationRead) SetMessage(v string) {
	o.Message = &v
}

func (o NotificationRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableNotificationRead struct {
	value *NotificationRead
	isSet bool
}

func (v NullableNotificationRead) Get() *NotificationRead {
	return v.value
}

func (v *NullableNotificationRead) Set(val *NotificationRead) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationRead) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationRead(val *NotificationRead) *NullableNotificationRead {
	return &NullableNotificationRead{value: val, isSet: true}
}

func (v NullableNotificationRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

