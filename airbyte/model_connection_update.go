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

// ConnectionUpdate struct for ConnectionUpdate
type ConnectionUpdate struct {
	ConnectionId string `json:"connectionId"`
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If \"${SOURCE_NAMESPACE}\" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `json:"namespaceFormat,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	OperationIds []string `json:"operationIds,omitempty"`
	SyncCatalog AirbyteCatalog `json:"syncCatalog"`
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	Status ConnectionStatus `json:"status"`
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
}

// NewConnectionUpdate instantiates a new ConnectionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUpdate(connectionId string, syncCatalog AirbyteCatalog, status ConnectionStatus) *ConnectionUpdate {
	this := ConnectionUpdate{}
	this.ConnectionId = connectionId
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	this.SyncCatalog = syncCatalog
	this.Status = status
	return &this
}

// NewConnectionUpdateWithDefaults instantiates a new ConnectionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUpdateWithDefaults() *ConnectionUpdate {
	this := ConnectionUpdate{}
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	return &this
}

// GetConnectionId returns the ConnectionId field value
func (o *ConnectionUpdate) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetConnectionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *ConnectionUpdate) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetNamespaceDefinition returns the NamespaceDefinition field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetNamespaceDefinition() NamespaceDefinitionType {
	if o == nil || o.NamespaceDefinition == nil {
		var ret NamespaceDefinitionType
		return ret
	}
	return *o.NamespaceDefinition
}

// GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool) {
	if o == nil || o.NamespaceDefinition == nil {
		return nil, false
	}
	return o.NamespaceDefinition, true
}

// HasNamespaceDefinition returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasNamespaceDefinition() bool {
	if o != nil && o.NamespaceDefinition != nil {
		return true
	}

	return false
}

// SetNamespaceDefinition gets a reference to the given NamespaceDefinitionType and assigns it to the NamespaceDefinition field.
func (o *ConnectionUpdate) SetNamespaceDefinition(v NamespaceDefinitionType) {
	o.NamespaceDefinition = &v
}

// GetNamespaceFormat returns the NamespaceFormat field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetNamespaceFormat() string {
	if o == nil || o.NamespaceFormat == nil {
		var ret string
		return ret
	}
	return *o.NamespaceFormat
}

// GetNamespaceFormatOk returns a tuple with the NamespaceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetNamespaceFormatOk() (*string, bool) {
	if o == nil || o.NamespaceFormat == nil {
		return nil, false
	}
	return o.NamespaceFormat, true
}

// HasNamespaceFormat returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasNamespaceFormat() bool {
	if o != nil && o.NamespaceFormat != nil {
		return true
	}

	return false
}

// SetNamespaceFormat gets a reference to the given string and assigns it to the NamespaceFormat field.
func (o *ConnectionUpdate) SetNamespaceFormat(v string) {
	o.NamespaceFormat = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ConnectionUpdate) SetPrefix(v string) {
	o.Prefix = &v
}

// GetOperationIds returns the OperationIds field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetOperationIds() []string {
	if o == nil || o.OperationIds == nil {
		var ret []string
		return ret
	}
	return o.OperationIds
}

// GetOperationIdsOk returns a tuple with the OperationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetOperationIdsOk() ([]string, bool) {
	if o == nil || o.OperationIds == nil {
		return nil, false
	}
	return o.OperationIds, true
}

// HasOperationIds returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasOperationIds() bool {
	if o != nil && o.OperationIds != nil {
		return true
	}

	return false
}

// SetOperationIds gets a reference to the given []string and assigns it to the OperationIds field.
func (o *ConnectionUpdate) SetOperationIds(v []string) {
	o.OperationIds = v
}

// GetSyncCatalog returns the SyncCatalog field value
func (o *ConnectionUpdate) GetSyncCatalog() AirbyteCatalog {
	if o == nil {
		var ret AirbyteCatalog
		return ret
	}

	return o.SyncCatalog
}

// GetSyncCatalogOk returns a tuple with the SyncCatalog field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetSyncCatalogOk() (*AirbyteCatalog, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SyncCatalog, true
}

// SetSyncCatalog sets field value
func (o *ConnectionUpdate) SetSyncCatalog(v AirbyteCatalog) {
	o.SyncCatalog = v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetSchedule() ConnectionSchedule {
	if o == nil || o.Schedule == nil {
		var ret ConnectionSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetScheduleOk() (*ConnectionSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ConnectionSchedule and assigns it to the Schedule field.
func (o *ConnectionUpdate) SetSchedule(v ConnectionSchedule) {
	o.Schedule = &v
}

// GetStatus returns the Status field value
func (o *ConnectionUpdate) GetStatus() ConnectionStatus {
	if o == nil {
		var ret ConnectionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetStatusOk() (*ConnectionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ConnectionUpdate) SetStatus(v ConnectionStatus) {
	o.Status = v
}

// GetResourceRequirements returns the ResourceRequirements field value if set, zero value otherwise.
func (o *ConnectionUpdate) GetResourceRequirements() ResourceRequirements {
	if o == nil || o.ResourceRequirements == nil {
		var ret ResourceRequirements
		return ret
	}
	return *o.ResourceRequirements
}

// GetResourceRequirementsOk returns a tuple with the ResourceRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionUpdate) GetResourceRequirementsOk() (*ResourceRequirements, bool) {
	if o == nil || o.ResourceRequirements == nil {
		return nil, false
	}
	return o.ResourceRequirements, true
}

// HasResourceRequirements returns a boolean if a field has been set.
func (o *ConnectionUpdate) HasResourceRequirements() bool {
	if o != nil && o.ResourceRequirements != nil {
		return true
	}

	return false
}

// SetResourceRequirements gets a reference to the given ResourceRequirements and assigns it to the ResourceRequirements field.
func (o *ConnectionUpdate) SetResourceRequirements(v ResourceRequirements) {
	o.ResourceRequirements = &v
}

func (o ConnectionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["connectionId"] = o.ConnectionId
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
	if o.OperationIds != nil {
		toSerialize["operationIds"] = o.OperationIds
	}
	if true {
		toSerialize["syncCatalog"] = o.SyncCatalog
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.ResourceRequirements != nil {
		toSerialize["resourceRequirements"] = o.ResourceRequirements
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionUpdate struct {
	value *ConnectionUpdate
	isSet bool
}

func (v NullableConnectionUpdate) Get() *ConnectionUpdate {
	return v.value
}

func (v *NullableConnectionUpdate) Set(val *ConnectionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUpdate(val *ConnectionUpdate) *NullableConnectionUpdate {
	return &NullableConnectionUpdate{value: val, isSet: true}
}

func (v NullableConnectionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


