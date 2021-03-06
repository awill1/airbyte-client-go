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

// WebBackendConnectionCreate struct for WebBackendConnectionCreate
type WebBackendConnectionCreate struct {
	// Optional name of the connection
	Name *string `json:"name,omitempty"`
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If \"${SOURCE_NAMESPACE}\" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `json:"namespaceFormat,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	SourceId string `json:"sourceId"`
	DestinationId string `json:"destinationId"`
	OperationIds []string `json:"operationIds,omitempty"`
	SyncCatalog *AirbyteCatalog `json:"syncCatalog,omitempty"`
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	Status ConnectionStatus `json:"status"`
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
	Operations []OperationCreate `json:"operations,omitempty"`
}

// NewWebBackendConnectionCreate instantiates a new WebBackendConnectionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebBackendConnectionCreate(sourceId string, destinationId string, status ConnectionStatus) *WebBackendConnectionCreate {
	this := WebBackendConnectionCreate{}
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	this.SourceId = sourceId
	this.DestinationId = destinationId
	this.Status = status
	return &this
}

// NewWebBackendConnectionCreateWithDefaults instantiates a new WebBackendConnectionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebBackendConnectionCreateWithDefaults() *WebBackendConnectionCreate {
	this := WebBackendConnectionCreate{}
	var namespaceDefinition NamespaceDefinitionType = NAMESPACEDEFINITIONTYPE_SOURCE
	this.NamespaceDefinition = &namespaceDefinition
	var namespaceFormat string = "null"
	this.NamespaceFormat = &namespaceFormat
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebBackendConnectionCreate) SetName(v string) {
	o.Name = &v
}

// GetNamespaceDefinition returns the NamespaceDefinition field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetNamespaceDefinition() NamespaceDefinitionType {
	if o == nil || o.NamespaceDefinition == nil {
		var ret NamespaceDefinitionType
		return ret
	}
	return *o.NamespaceDefinition
}

// GetNamespaceDefinitionOk returns a tuple with the NamespaceDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetNamespaceDefinitionOk() (*NamespaceDefinitionType, bool) {
	if o == nil || o.NamespaceDefinition == nil {
		return nil, false
	}
	return o.NamespaceDefinition, true
}

// HasNamespaceDefinition returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasNamespaceDefinition() bool {
	if o != nil && o.NamespaceDefinition != nil {
		return true
	}

	return false
}

// SetNamespaceDefinition gets a reference to the given NamespaceDefinitionType and assigns it to the NamespaceDefinition field.
func (o *WebBackendConnectionCreate) SetNamespaceDefinition(v NamespaceDefinitionType) {
	o.NamespaceDefinition = &v
}

// GetNamespaceFormat returns the NamespaceFormat field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetNamespaceFormat() string {
	if o == nil || o.NamespaceFormat == nil {
		var ret string
		return ret
	}
	return *o.NamespaceFormat
}

// GetNamespaceFormatOk returns a tuple with the NamespaceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetNamespaceFormatOk() (*string, bool) {
	if o == nil || o.NamespaceFormat == nil {
		return nil, false
	}
	return o.NamespaceFormat, true
}

// HasNamespaceFormat returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasNamespaceFormat() bool {
	if o != nil && o.NamespaceFormat != nil {
		return true
	}

	return false
}

// SetNamespaceFormat gets a reference to the given string and assigns it to the NamespaceFormat field.
func (o *WebBackendConnectionCreate) SetNamespaceFormat(v string) {
	o.NamespaceFormat = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasPrefix() bool {
	if o != nil && o.Prefix != nil {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *WebBackendConnectionCreate) SetPrefix(v string) {
	o.Prefix = &v
}

// GetSourceId returns the SourceId field value
func (o *WebBackendConnectionCreate) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetSourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *WebBackendConnectionCreate) SetSourceId(v string) {
	o.SourceId = v
}

// GetDestinationId returns the DestinationId field value
func (o *WebBackendConnectionCreate) GetDestinationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationId
}

// GetDestinationIdOk returns a tuple with the DestinationId field value
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetDestinationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationId, true
}

// SetDestinationId sets field value
func (o *WebBackendConnectionCreate) SetDestinationId(v string) {
	o.DestinationId = v
}

// GetOperationIds returns the OperationIds field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetOperationIds() []string {
	if o == nil || o.OperationIds == nil {
		var ret []string
		return ret
	}
	return o.OperationIds
}

// GetOperationIdsOk returns a tuple with the OperationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetOperationIdsOk() ([]string, bool) {
	if o == nil || o.OperationIds == nil {
		return nil, false
	}
	return o.OperationIds, true
}

// HasOperationIds returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasOperationIds() bool {
	if o != nil && o.OperationIds != nil {
		return true
	}

	return false
}

// SetOperationIds gets a reference to the given []string and assigns it to the OperationIds field.
func (o *WebBackendConnectionCreate) SetOperationIds(v []string) {
	o.OperationIds = v
}

// GetSyncCatalog returns the SyncCatalog field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetSyncCatalog() AirbyteCatalog {
	if o == nil || o.SyncCatalog == nil {
		var ret AirbyteCatalog
		return ret
	}
	return *o.SyncCatalog
}

// GetSyncCatalogOk returns a tuple with the SyncCatalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetSyncCatalogOk() (*AirbyteCatalog, bool) {
	if o == nil || o.SyncCatalog == nil {
		return nil, false
	}
	return o.SyncCatalog, true
}

// HasSyncCatalog returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasSyncCatalog() bool {
	if o != nil && o.SyncCatalog != nil {
		return true
	}

	return false
}

// SetSyncCatalog gets a reference to the given AirbyteCatalog and assigns it to the SyncCatalog field.
func (o *WebBackendConnectionCreate) SetSyncCatalog(v AirbyteCatalog) {
	o.SyncCatalog = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetSchedule() ConnectionSchedule {
	if o == nil || o.Schedule == nil {
		var ret ConnectionSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetScheduleOk() (*ConnectionSchedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ConnectionSchedule and assigns it to the Schedule field.
func (o *WebBackendConnectionCreate) SetSchedule(v ConnectionSchedule) {
	o.Schedule = &v
}

// GetStatus returns the Status field value
func (o *WebBackendConnectionCreate) GetStatus() ConnectionStatus {
	if o == nil {
		var ret ConnectionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetStatusOk() (*ConnectionStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WebBackendConnectionCreate) SetStatus(v ConnectionStatus) {
	o.Status = v
}

// GetResourceRequirements returns the ResourceRequirements field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetResourceRequirements() ResourceRequirements {
	if o == nil || o.ResourceRequirements == nil {
		var ret ResourceRequirements
		return ret
	}
	return *o.ResourceRequirements
}

// GetResourceRequirementsOk returns a tuple with the ResourceRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetResourceRequirementsOk() (*ResourceRequirements, bool) {
	if o == nil || o.ResourceRequirements == nil {
		return nil, false
	}
	return o.ResourceRequirements, true
}

// HasResourceRequirements returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasResourceRequirements() bool {
	if o != nil && o.ResourceRequirements != nil {
		return true
	}

	return false
}

// SetResourceRequirements gets a reference to the given ResourceRequirements and assigns it to the ResourceRequirements field.
func (o *WebBackendConnectionCreate) SetResourceRequirements(v ResourceRequirements) {
	o.ResourceRequirements = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *WebBackendConnectionCreate) GetOperations() []OperationCreate {
	if o == nil || o.Operations == nil {
		var ret []OperationCreate
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebBackendConnectionCreate) GetOperationsOk() ([]OperationCreate, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *WebBackendConnectionCreate) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []OperationCreate and assigns it to the Operations field.
func (o *WebBackendConnectionCreate) SetOperations(v []OperationCreate) {
	o.Operations = v
}

func (o WebBackendConnectionCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["destinationId"] = o.DestinationId
	}
	if o.OperationIds != nil {
		toSerialize["operationIds"] = o.OperationIds
	}
	if o.SyncCatalog != nil {
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
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	return json.Marshal(toSerialize)
}

type NullableWebBackendConnectionCreate struct {
	value *WebBackendConnectionCreate
	isSet bool
}

func (v NullableWebBackendConnectionCreate) Get() *WebBackendConnectionCreate {
	return v.value
}

func (v *NullableWebBackendConnectionCreate) Set(val *WebBackendConnectionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableWebBackendConnectionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableWebBackendConnectionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebBackendConnectionCreate(val *WebBackendConnectionCreate) *NullableWebBackendConnectionCreate {
	return &NullableWebBackendConnectionCreate{value: val, isSet: true}
}

func (v NullableWebBackendConnectionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebBackendConnectionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


