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

// DbMigrationExecutionRead struct for DbMigrationExecutionRead
type DbMigrationExecutionRead struct {
	InitialVersion *string `json:"initialVersion,omitempty"`
	TargetVersion *string `json:"targetVersion,omitempty"`
	ExecutedMigrations []DbMigrationRead `json:"executedMigrations,omitempty"`
}

// NewDbMigrationExecutionRead instantiates a new DbMigrationExecutionRead object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbMigrationExecutionRead() *DbMigrationExecutionRead {
	this := DbMigrationExecutionRead{}
	return &this
}

// NewDbMigrationExecutionReadWithDefaults instantiates a new DbMigrationExecutionRead object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbMigrationExecutionReadWithDefaults() *DbMigrationExecutionRead {
	this := DbMigrationExecutionRead{}
	return &this
}

// GetInitialVersion returns the InitialVersion field value if set, zero value otherwise.
func (o *DbMigrationExecutionRead) GetInitialVersion() string {
	if o == nil || o.InitialVersion == nil {
		var ret string
		return ret
	}
	return *o.InitialVersion
}

// GetInitialVersionOk returns a tuple with the InitialVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbMigrationExecutionRead) GetInitialVersionOk() (*string, bool) {
	if o == nil || o.InitialVersion == nil {
		return nil, false
	}
	return o.InitialVersion, true
}

// HasInitialVersion returns a boolean if a field has been set.
func (o *DbMigrationExecutionRead) HasInitialVersion() bool {
	if o != nil && o.InitialVersion != nil {
		return true
	}

	return false
}

// SetInitialVersion gets a reference to the given string and assigns it to the InitialVersion field.
func (o *DbMigrationExecutionRead) SetInitialVersion(v string) {
	o.InitialVersion = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *DbMigrationExecutionRead) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbMigrationExecutionRead) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *DbMigrationExecutionRead) HasTargetVersion() bool {
	if o != nil && o.TargetVersion != nil {
		return true
	}

	return false
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *DbMigrationExecutionRead) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// GetExecutedMigrations returns the ExecutedMigrations field value if set, zero value otherwise.
func (o *DbMigrationExecutionRead) GetExecutedMigrations() []DbMigrationRead {
	if o == nil || o.ExecutedMigrations == nil {
		var ret []DbMigrationRead
		return ret
	}
	return o.ExecutedMigrations
}

// GetExecutedMigrationsOk returns a tuple with the ExecutedMigrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbMigrationExecutionRead) GetExecutedMigrationsOk() ([]DbMigrationRead, bool) {
	if o == nil || o.ExecutedMigrations == nil {
		return nil, false
	}
	return o.ExecutedMigrations, true
}

// HasExecutedMigrations returns a boolean if a field has been set.
func (o *DbMigrationExecutionRead) HasExecutedMigrations() bool {
	if o != nil && o.ExecutedMigrations != nil {
		return true
	}

	return false
}

// SetExecutedMigrations gets a reference to the given []DbMigrationRead and assigns it to the ExecutedMigrations field.
func (o *DbMigrationExecutionRead) SetExecutedMigrations(v []DbMigrationRead) {
	o.ExecutedMigrations = v
}

func (o DbMigrationExecutionRead) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InitialVersion != nil {
		toSerialize["initialVersion"] = o.InitialVersion
	}
	if o.TargetVersion != nil {
		toSerialize["targetVersion"] = o.TargetVersion
	}
	if o.ExecutedMigrations != nil {
		toSerialize["executedMigrations"] = o.ExecutedMigrations
	}
	return json.Marshal(toSerialize)
}

type NullableDbMigrationExecutionRead struct {
	value *DbMigrationExecutionRead
	isSet bool
}

func (v NullableDbMigrationExecutionRead) Get() *DbMigrationExecutionRead {
	return v.value
}

func (v *NullableDbMigrationExecutionRead) Set(val *DbMigrationExecutionRead) {
	v.value = val
	v.isSet = true
}

func (v NullableDbMigrationExecutionRead) IsSet() bool {
	return v.isSet
}

func (v *NullableDbMigrationExecutionRead) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbMigrationExecutionRead(val *DbMigrationExecutionRead) *NullableDbMigrationExecutionRead {
	return &NullableDbMigrationExecutionRead{value: val, isSet: true}
}

func (v NullableDbMigrationExecutionRead) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbMigrationExecutionRead) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


