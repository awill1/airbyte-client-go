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

// WorkspaceReadList struct for WorkspaceReadList
type WorkspaceReadList struct {
	Workspaces []WorkspaceRead `json:"workspaces"`
}

// NewWorkspaceReadList instantiates a new WorkspaceReadList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceReadList(workspaces []WorkspaceRead) *WorkspaceReadList {
	this := WorkspaceReadList{}
	this.Workspaces = workspaces
	return &this
}

// NewWorkspaceReadListWithDefaults instantiates a new WorkspaceReadList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceReadListWithDefaults() *WorkspaceReadList {
	this := WorkspaceReadList{}
	return &this
}

// GetWorkspaces returns the Workspaces field value
func (o *WorkspaceReadList) GetWorkspaces() []WorkspaceRead {
	if o == nil {
		var ret []WorkspaceRead
		return ret
	}

	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value
// and a boolean to check if the value has been set.
func (o *WorkspaceReadList) GetWorkspacesOk() ([]WorkspaceRead, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Workspaces, true
}

// SetWorkspaces sets field value
func (o *WorkspaceReadList) SetWorkspaces(v []WorkspaceRead) {
	o.Workspaces = v
}

func (o WorkspaceReadList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workspaces"] = o.Workspaces
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceReadList struct {
	value *WorkspaceReadList
	isSet bool
}

func (v NullableWorkspaceReadList) Get() *WorkspaceReadList {
	return v.value
}

func (v *NullableWorkspaceReadList) Set(val *WorkspaceReadList) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceReadList) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceReadList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceReadList(val *WorkspaceReadList) *NullableWorkspaceReadList {
	return &NullableWorkspaceReadList{value: val, isSet: true}
}

func (v NullableWorkspaceReadList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceReadList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


