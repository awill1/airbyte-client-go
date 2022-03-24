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

// OperatorDbt struct for OperatorDbt
type OperatorDbt struct {
	GitRepoUrl string `json:"gitRepoUrl"`
	GitRepoBranch *string `json:"gitRepoBranch,omitempty"`
	DockerImage *string `json:"dockerImage,omitempty"`
	DbtArguments *string `json:"dbtArguments,omitempty"`
}

// NewOperatorDbt instantiates a new OperatorDbt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorDbt(gitRepoUrl string) *OperatorDbt {
	this := OperatorDbt{}
	this.GitRepoUrl = gitRepoUrl
	return &this
}

// NewOperatorDbtWithDefaults instantiates a new OperatorDbt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorDbtWithDefaults() *OperatorDbt {
	this := OperatorDbt{}
	return &this
}

// GetGitRepoUrl returns the GitRepoUrl field value
func (o *OperatorDbt) GetGitRepoUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitRepoUrl
}

// GetGitRepoUrlOk returns a tuple with the GitRepoUrl field value
// and a boolean to check if the value has been set.
func (o *OperatorDbt) GetGitRepoUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GitRepoUrl, true
}

// SetGitRepoUrl sets field value
func (o *OperatorDbt) SetGitRepoUrl(v string) {
	o.GitRepoUrl = v
}

// GetGitRepoBranch returns the GitRepoBranch field value if set, zero value otherwise.
func (o *OperatorDbt) GetGitRepoBranch() string {
	if o == nil || o.GitRepoBranch == nil {
		var ret string
		return ret
	}
	return *o.GitRepoBranch
}

// GetGitRepoBranchOk returns a tuple with the GitRepoBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDbt) GetGitRepoBranchOk() (*string, bool) {
	if o == nil || o.GitRepoBranch == nil {
		return nil, false
	}
	return o.GitRepoBranch, true
}

// HasGitRepoBranch returns a boolean if a field has been set.
func (o *OperatorDbt) HasGitRepoBranch() bool {
	if o != nil && o.GitRepoBranch != nil {
		return true
	}

	return false
}

// SetGitRepoBranch gets a reference to the given string and assigns it to the GitRepoBranch field.
func (o *OperatorDbt) SetGitRepoBranch(v string) {
	o.GitRepoBranch = &v
}

// GetDockerImage returns the DockerImage field value if set, zero value otherwise.
func (o *OperatorDbt) GetDockerImage() string {
	if o == nil || o.DockerImage == nil {
		var ret string
		return ret
	}
	return *o.DockerImage
}

// GetDockerImageOk returns a tuple with the DockerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDbt) GetDockerImageOk() (*string, bool) {
	if o == nil || o.DockerImage == nil {
		return nil, false
	}
	return o.DockerImage, true
}

// HasDockerImage returns a boolean if a field has been set.
func (o *OperatorDbt) HasDockerImage() bool {
	if o != nil && o.DockerImage != nil {
		return true
	}

	return false
}

// SetDockerImage gets a reference to the given string and assigns it to the DockerImage field.
func (o *OperatorDbt) SetDockerImage(v string) {
	o.DockerImage = &v
}

// GetDbtArguments returns the DbtArguments field value if set, zero value otherwise.
func (o *OperatorDbt) GetDbtArguments() string {
	if o == nil || o.DbtArguments == nil {
		var ret string
		return ret
	}
	return *o.DbtArguments
}

// GetDbtArgumentsOk returns a tuple with the DbtArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorDbt) GetDbtArgumentsOk() (*string, bool) {
	if o == nil || o.DbtArguments == nil {
		return nil, false
	}
	return o.DbtArguments, true
}

// HasDbtArguments returns a boolean if a field has been set.
func (o *OperatorDbt) HasDbtArguments() bool {
	if o != nil && o.DbtArguments != nil {
		return true
	}

	return false
}

// SetDbtArguments gets a reference to the given string and assigns it to the DbtArguments field.
func (o *OperatorDbt) SetDbtArguments(v string) {
	o.DbtArguments = &v
}

func (o OperatorDbt) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gitRepoUrl"] = o.GitRepoUrl
	}
	if o.GitRepoBranch != nil {
		toSerialize["gitRepoBranch"] = o.GitRepoBranch
	}
	if o.DockerImage != nil {
		toSerialize["dockerImage"] = o.DockerImage
	}
	if o.DbtArguments != nil {
		toSerialize["dbtArguments"] = o.DbtArguments
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorDbt struct {
	value *OperatorDbt
	isSet bool
}

func (v NullableOperatorDbt) Get() *OperatorDbt {
	return v.value
}

func (v *NullableOperatorDbt) Set(val *OperatorDbt) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorDbt) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorDbt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorDbt(val *OperatorDbt) *NullableOperatorDbt {
	return &NullableOperatorDbt{value: val, isSet: true}
}

func (v NullableOperatorDbt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorDbt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


