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

// ConnectionSearch struct for ConnectionSearch
type ConnectionSearch struct {
	ConnectionId *string `json:"connectionId,omitempty"`
	Name *string `json:"name,omitempty"`
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If \"${SOURCE_NAMESPACE}\" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `json:"namespaceFormat,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	SourceId *string `json:"sourceId,omitempty"`
	DestinationId *string `json:"destinationId,omitempty"`
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	Status *ConnectionStatus `json:"status,omitempty"`
	Source *SourceSearch `json:"source,omitempty"`
	Destination *DestinationSearch `json:"destination,omitempty"`
}

// NewConnectionSearch instantiates a new ConnectionSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionSearch() *ConnectionSearch {
	this := ConnectionSearch{}
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	return &this
}

// NewConnectionSearchWithDefaults instantiates a new ConnectionSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionSearchWithDefaults() *ConnectionSearch {
	this := ConnectionSearch{}
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *ConnectionSearch) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *ConnectionSearch) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *ConnectionSearch) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionSearch) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionSearch) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionSearch) SetName(v string) {
	o.Name = &v
}

// GetNamespaceDefinition returns the NamespaceDefinition field value if set, zero value otherwise.
func (o *ConnectionSearch) GetNamespaceDefinition() NamespaceDefinitionType {
	if o == nil || o.NamespaceDefinition == nil {
		var ret NamespaceDefinitionType
		return ret
	}
	return *o.NamespaceDefinition
}

// GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool) {
	if o == nil || o.NamespaceDefinition == nil {
		return nil, false
	}
	return o.NamespaceDefinition, true
}

// HasNamespaceDefinition returns a boolean if a field has been set.
func (o *ConnectionSearch) HasNamespaceDefinition() bool {
	if o != nil && o.NamespaceDefinition != nil {
		return true
	}

	return false
}

// SetNamespaceDefinition gets a reference to the given NamespaceDefinitionType and assigns it to the NamespaceDefinition field.
func (o *ConnectionSearch) SetNamespaceDefinition(v NamespaceDefinitionType) {
	o.NamespaceDefinition = &v
}

// GetNamespaceFormat returns the NamespaceFormat field value if set, zero value otherwise.
func (o *ConnectionSearch) GetNamespaceFormat() string {
	if o == nil || o.NamespaceFormat == nil {
		var ret string
		return ret
	}
	return *o.NamespaceFormat
}

// GetNamespaceFormatOk returns a tuple with the NamespaceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetNamespaceFormatOk() (*string, bool) {
	if o == nil || o.NamespaceFormat == nil {
		return nil, false
	}
	return o.NamespaceFormat, true
}

// HasNamespaceFormat returns a boolean if a field has been set.
func (o *ConnectionSearch) HasNamespaceFormat() bool {
	if o != nil && o.NamespaceFormat != nil {
		return true
	}

	return false
}

// SetNamespaceFormat gets a reference to the given string and assigns it to the NamespaceFormat field.
func (o *ConnectionSearch) SetNamespaceFormat(v string) {
	o.NamespaceFormat = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ConnectionSearch) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ConnectionSearch) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ConnectionSearch) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *ConnectionSearch) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *ConnectionSearch) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *ConnectionSearch) SetSourceId(v string) {
	o.SourceId = &v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise.
func (o *ConnectionSearch) GetDestinationId() string {
	if o == nil || o.DestinationId == nil {
		var ret string
		return ret
	}
	return *o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetDestinationIdOk() (*string, bool) {
	if o == nil || o.DestinationId == nil {
		return nil, false
	}
	return o.DestinationId, true
}

// HasDestinationId returns a boolean if a field has been set.
func (o *ConnectionSearch) HasDestinationId() bool {
	if o != nil && o.DestinationId != nil {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given string and assigns it to the DestinationId field.
func (o *ConnectionSearch) SetDestinationId(v string) {
	o.DestinationId = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ConnectionSearch) GetSchedule() ConnectionSchedule {
	if o == nil || o.Schedule == nil {
		var ret ConnectionSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetScheduleOk() (*ConnectionSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ConnectionSearch) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ConnectionSchedule and assigns it to the Schedule field.
func (o *ConnectionSearch) SetSchedule(v ConnectionSchedule) {
	o.Schedule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectionSearch) GetStatus() ConnectionStatus {
	if o == nil || o.Status == nil {
		var ret ConnectionStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetStatusOk() (*ConnectionStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionSearch) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ConnectionStatus and assigns it to the Status field.
func (o *ConnectionSearch) SetStatus(v ConnectionStatus) {
	o.Status = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ConnectionSearch) GetSource() SourceSearch {
	if o == nil || o.Source == nil {
		var ret SourceSearch
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetSourceOk() (*SourceSearch, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ConnectionSearch) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given SourceSearch and assigns it to the Source field.
func (o *ConnectionSearch) SetSource(v SourceSearch) {
	o.Source = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *ConnectionSearch) GetDestination() DestinationSearch {
	if o == nil || o.Destination == nil {
		var ret DestinationSearch
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionSearch) GetDestinationOk() (*DestinationSearch, bool) {
	if o == nil || o.Destination == nil {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *ConnectionSearch) HasDestination() bool {
	if o != nil && o.Destination != nil {
		return true
	}

	return false
}

// SetDestination gets a reference to the given DestinationSearch and assigns it to the Destination field.
func (o *ConnectionSearch) SetDestination(v DestinationSearch) {
	o.Destination = &v
}

func (o ConnectionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NamespaceDefinition != nil {
		toSerialize["namespaceDefinition"] = o.NamespaceDefinition
	}
	if o.NamespaceFormat != nil {
		toSerialize["namespaceFormat"] = o.NamespaceFormat
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.DestinationId != nil {
		toSerialize["destinationId"] = o.DestinationId
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Destination != nil {
		toSerialize["destination"] = o.Destination
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionSearch struct {
	value *ConnectionSearch
	isSet bool
}

func (v NullableConnectionSearch) Get() *ConnectionSearch {
	return v.value
}

func (v *NullableConnectionSearch) Set(val *ConnectionSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionSearch(val *ConnectionSearch) *NullableConnectionSearch {
	return &NullableConnectionSearch{value: val, isSet: true}
}

func (v NullableConnectionSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


