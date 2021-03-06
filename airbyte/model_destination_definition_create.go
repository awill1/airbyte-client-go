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

// DestinationDefinitionCreate struct for DestinationDefinitionCreate
type DestinationDefinitionCreate struct {
	Name string `json:"name"`
	DockerRepository string `json:"dockerRepository"`
	DockerImageTag string `json:"dockerImageTag"`
	DocumentationUrl string `json:"documentationUrl"`
	Icon *string `json:"icon,omitempty"`
}

// NewDestinationDefinitionCreate instantiates a new DestinationDefinitionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationDefinitionCreate(name string, dockerRepository string, dockerImageTag string, documentationUrl string) *DestinationDefinitionCreate {
	this := DestinationDefinitionCreate{}
	this.Name = name
	this.DockerRepository = dockerRepository
	this.DockerImageTag = dockerImageTag
	this.DocumentationUrl = documentationUrl
	return &this
}

// NewDestinationDefinitionCreateWithDefaults instantiates a new DestinationDefinitionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationDefinitionCreateWithDefaults() *DestinationDefinitionCreate {
	this := DestinationDefinitionCreate{}
	return &this
}

// GetName returns the Name field value
func (o *DestinationDefinitionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationDefinitionCreate) SetName(v string) {
	o.Name = v
}

// GetDockerRepository returns the DockerRepository field value
func (o *DestinationDefinitionCreate) GetDockerRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerRepository
}

// GetDockerRepositoryOk returns a tuple with the DockerRepository field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionCreate) GetDockerRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerRepository, true
}

// SetDockerRepository sets field value
func (o *DestinationDefinitionCreate) SetDockerRepository(v string) {
	o.DockerRepository = v
}

// GetDockerImageTag returns the DockerImageTag field value
func (o *DestinationDefinitionCreate) GetDockerImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionCreate) GetDockerImageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerImageTag, true
}

// SetDockerImageTag sets field value
func (o *DestinationDefinitionCreate) SetDockerImageTag(v string) {
	o.DockerImageTag = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *DestinationDefinitionCreate) GetDocumentationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionCreate) GetDocumentationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *DestinationDefinitionCreate) SetDocumentationUrl(v string) {
	o.DocumentationUrl = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *DestinationDefinitionCreate) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionCreate) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *DestinationDefinitionCreate) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *DestinationDefinitionCreate) SetIcon(v string) {
	o.Icon = &v
}

func (o DestinationDefinitionCreate) MarshalJSON() ([]byte, error) {
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

type NullableDestinationDefinitionCreate struct {
	value *DestinationDefinitionCreate
	isSet bool
}

func (v NullableDestinationDefinitionCreate) Get() *DestinationDefinitionCreate {
	return v.value
}

func (v *NullableDestinationDefinitionCreate) Set(val *DestinationDefinitionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationDefinitionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationDefinitionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationDefinitionCreate(val *DestinationDefinitionCreate) *NullableDestinationDefinitionCreate {
	return &NullableDestinationDefinitionCreate{value: val, isSet: true}
}

func (v NullableDestinationDefinitionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationDefinitionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


