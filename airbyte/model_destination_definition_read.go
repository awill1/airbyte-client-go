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

// DestinationDefinitionRead struct for DestinationDefinitionRead
type DestinationDefinitionRead struct {
	DestinationDefinitionId string `json:"destinationDefinitionId"`
	Name string `json:"name"`
	DockerRepository string `json:"dockerRepository"`
	DockerImageTag string `json:"dockerImageTag"`
	DocumentationUrl string `json:"documentationUrl"`
	Icon *string `json:"icon,omitempty"`
	ReleaseStage *ReleaseStage `json:"releaseStage,omitempty"`
	// The date when this connector was first released, in yyyy-mm-dd format.
	ReleaseDate *string `json:"releaseDate,omitempty"`
}

// NewDestinationDefinitionRead instantiates a new DestinationDefinitionRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationDefinitionRead(destinationDefinitionId string, name string, dockerRepository string, dockerImageTag string, documentationUrl string) *DestinationDefinitionRead {
	this := DestinationDefinitionRead{}
	this.DestinationDefinitionId = destinationDefinitionId
	this.Name = name
	this.DockerRepository = dockerRepository
	this.DockerImageTag = dockerImageTag
	this.DocumentationUrl = documentationUrl
	return &this
}

// NewDestinationDefinitionReadWithDefaults instantiates a new DestinationDefinitionRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationDefinitionReadWithDefaults() *DestinationDefinitionRead {
	this := DestinationDefinitionRead{}
	return &this
}

// GetDestinationDefinitionId returns the DestinationDefinitionId field value
func (o *DestinationDefinitionRead) GetDestinationDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationDefinitionId
}

// GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetDestinationDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationDefinitionId, true
}

// SetDestinationDefinitionId sets field value
func (o *DestinationDefinitionRead) SetDestinationDefinitionId(v string) {
	o.DestinationDefinitionId = v
}

// GetName returns the Name field value
func (o *DestinationDefinitionRead) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DestinationDefinitionRead) SetName(v string) {
	o.Name = v
}

// GetDockerRepository returns the DockerRepository field value
func (o *DestinationDefinitionRead) GetDockerRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerRepository
}

// GetDockerRepositoryOk returns a tuple with the DockerRepository field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetDockerRepositoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerRepository, true
}

// SetDockerRepository sets field value
func (o *DestinationDefinitionRead) SetDockerRepository(v string) {
	o.DockerRepository = v
}

// GetDockerImageTag returns the DockerImageTag field value
func (o *DestinationDefinitionRead) GetDockerImageTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageTag
}

// GetDockerImageTagOk returns a tuple with the DockerImageTag field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetDockerImageTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DockerImageTag, true
}

// SetDockerImageTag sets field value
func (o *DestinationDefinitionRead) SetDockerImageTag(v string) {
	o.DockerImageTag = v
}

// GetDocumentationUrl returns the DocumentationUrl field value
func (o *DestinationDefinitionRead) GetDocumentationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetDocumentationUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DocumentationUrl, true
}

// SetDocumentationUrl sets field value
func (o *DestinationDefinitionRead) SetDocumentationUrl(v string) {
	o.DocumentationUrl = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *DestinationDefinitionRead) GetIcon() string {
	if o == nil || o.Icon == nil {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetIconOk() (*string, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *DestinationDefinitionRead) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *DestinationDefinitionRead) SetIcon(v string) {
	o.Icon = &v
}

// GetReleaseStage returns the ReleaseStage field value if set, zero value otherwise.
func (o *DestinationDefinitionRead) GetReleaseStage() ReleaseStage {
	if o == nil || o.ReleaseStage == nil {
		var ret ReleaseStage
		return ret
	}
	return *o.ReleaseStage
}

// GetReleaseStageOk returns a tuple with the ReleaseStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetReleaseStageOk() (*ReleaseStage, bool) {
	if o == nil || o.ReleaseStage == nil {
		return nil, false
	}
	return o.ReleaseStage, true
}

// HasReleaseStage returns a boolean if a field has been set.
func (o *DestinationDefinitionRead) HasReleaseStage() bool {
	if o != nil && o.ReleaseStage != nil {
		return true
	}

	return false
}

// SetReleaseStage gets a reference to the given ReleaseStage and assigns it to the ReleaseStage field.
func (o *DestinationDefinitionRead) SetReleaseStage(v ReleaseStage) {
	o.ReleaseStage = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *DestinationDefinitionRead) GetReleaseDate() string {
	if o == nil || o.ReleaseDate == nil {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationDefinitionRead) GetReleaseDateOk() (*string, bool) {
	if o == nil || o.ReleaseDate == nil {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *DestinationDefinitionRead) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate != nil {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *DestinationDefinitionRead) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

func (o DestinationDefinitionRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinationDefinitionId"] = o.DestinationDefinitionId
	}
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
	if o.ReleaseStage != nil {
		toSerialize["releaseStage"] = o.ReleaseStage
	}
	if o.ReleaseDate != nil {
		toSerialize["releaseDate"] = o.ReleaseDate
	}
	return json.Marshal(toSerialize)
}

type NullableDestinationDefinitionRead struct {
	value *DestinationDefinitionRead
	isSet bool
}

func (v NullableDestinationDefinitionRead) Get() *DestinationDefinitionRead {
	return v.value
}

func (v *NullableDestinationDefinitionRead) Set(val *DestinationDefinitionRead) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationDefinitionRead) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationDefinitionRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationDefinitionRead(val *DestinationDefinitionRead) *NullableDestinationDefinitionRead {
	return &NullableDestinationDefinitionRead{value: val, isSet: true}
}

func (v NullableDestinationDefinitionRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationDefinitionRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


