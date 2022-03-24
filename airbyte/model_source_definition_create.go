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

// SourceDefinitionCreate struct for SourceDefinitionCreate
type SourceDefinitionCreate struct {
	Name string `json:"name"`
	DockerRepository string `json:"dockerRepository"`
	DockerImageTag string `json:"dockerImageTag"`
	DocumentationUrl string `json:"documentationUrl"`
	Icon *string `json:"icon,omitempty"`
}

// NewSourceDefinitionCreate instantiates a new SourceDefinitionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceDefinitionCreate(name string, dockerRepository string, dockerImageTag string, documentationUrl string) *SourceDefinitionCreate {
	this := SourceDefinitionCreate{}
	this.Name = name
	this.DockerRepository = dockerRepository
	this.DockerImageTag = dockerImageTag
	this.DocumentationUrl = documentationUrl
	return &this
}

// NewSourceDefinitionCreateWithDefaults instantiates a new SourceDefinitionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceDefinitionCreateWithDefaults() *SourceDefinitionCreate {
	this := SourceDefinitionCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SourceDefinitionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SourceDefinitionCreate) SetName(v string) {
	o.Name = v
}

// GetDockerRepository returns the DockerRepository field value
func (o *SourceDefinitionCreate) GetDockerRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerRepository
}

// GetDockerRepositoryOk returns a tuple with the DockerRepository field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionCreate) GetDockerRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerRepository, true
}

// SetDockerRepository sets field value
func (o *SourceDefinitionCreate) SetDockerRepository(v string) {
	o.DockerRepository = v
}

// GetDockerImageTag returns the DockerImageTag field value
func (o *SourceDefinitionCreate) GetDockerImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionCreate) GetDockerImageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerImageTag, true
}

// SetDockerImageTag sets field value
func (o *SourceDefinitionCreate) SetDockerImageTag(v string) {
	o.DockerImageTag = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *SourceDefinitionCreate) GetDocumentationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *SourceDefinitionCreate) GetDocumentationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *SourceDefinitionCreate) SetDocumentationUrl(v string) {
	o.DocumentationUrl = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *SourceDefinitionCreate) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceDefinitionCreate) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *SourceDefinitionCreate) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *SourceDefinitionCreate) SetIcon(v string) {
	o.Icon = &v
}

func (o SourceDefinitionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["dockerRepository"] = o.DockerRepository
	}
	if true {
		toSerialize["dockerImageTag"] = o.DockerImageTag
	}
	if true {
		toSerialize["documentationUrl"] = o.DocumentationUrl
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	return json.Marshal(toSerialize)
}

type NullableSourceDefinitionCreate struct {
	value *SourceDefinitionCreate
	isSet bool
}

func (v NullableSourceDefinitionCreate) Get() *SourceDefinitionCreate {
	return v.value
}

func (v *NullableSourceDefinitionCreate) Set(val *SourceDefinitionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceDefinitionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceDefinitionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceDefinitionCreate(val *SourceDefinitionCreate) *NullableSourceDefinitionCreate {
	return &NullableSourceDefinitionCreate{value: val, isSet: true}
}

func (v NullableSourceDefinitionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceDefinitionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


