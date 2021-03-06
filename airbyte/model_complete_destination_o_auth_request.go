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

// CompleteDestinationOAuthRequest struct for CompleteDestinationOAuthRequest
type CompleteDestinationOAuthRequest struct {
	DestinationDefinitionId string `json:"destinationDefinitionId"`
	WorkspaceId string `json:"workspaceId"`
	// When completing OAuth flow to gain an access token, some API sometimes requires to verify that the app re-send the redirectUrl that was used when consent was given.
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	// The query parameters present in the redirect URL after a user granted consent e.g auth code
	QueryParams map[string]map[string]interface{} `json:"queryParams,omitempty"`
	// OAuth specific blob.
	OAuthInputConfiguration interface{} `json:"oAuthInputConfiguration,omitempty"`
}

// NewCompleteDestinationOAuthRequest instantiates a new CompleteDestinationOAuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteDestinationOAuthRequest(destinationDefinitionId string, workspaceId string) *CompleteDestinationOAuthRequest {
	this := CompleteDestinationOAuthRequest{}
	this.DestinationDefinitionId = destinationDefinitionId
	this.WorkspaceId = workspaceId
	return &this
}

// NewCompleteDestinationOAuthRequestWithDefaults instantiates a new CompleteDestinationOAuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteDestinationOAuthRequestWithDefaults() *CompleteDestinationOAuthRequest {
	this := CompleteDestinationOAuthRequest{}
	return &this
}

// GetDestinationDefinitionId returns the DestinationDefinitionId field value
func (o *CompleteDestinationOAuthRequest) GetDestinationDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationDefinitionId
}

// GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field value
// and a boolean to check if the value has been set.
func (o *CompleteDestinationOAuthRequest) GetDestinationDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationDefinitionId, true
}

// SetDestinationDefinitionId sets field value
func (o *CompleteDestinationOAuthRequest) SetDestinationDefinitionId(v string) {
	o.DestinationDefinitionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *CompleteDestinationOAuthRequest) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *CompleteDestinationOAuthRequest) GetWorkspaceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *CompleteDestinationOAuthRequest) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *CompleteDestinationOAuthRequest) GetRedirectUrl() string {
	if o == nil || o.RedirectUrl == nil {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteDestinationOAuthRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || o.RedirectUrl == nil {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *CompleteDestinationOAuthRequest) HasRedirectUrl() bool {
	if o != nil && o.RedirectUrl != nil {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *CompleteDestinationOAuthRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetQueryParams returns the QueryParams field value if set, zero value otherwise.
func (o *CompleteDestinationOAuthRequest) GetQueryParams() map[string]map[string]interface{} {
	if o == nil || o.QueryParams == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.QueryParams
}

// GetQueryParamsOk returns a tuple with the QueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteDestinationOAuthRequest) GetQueryParamsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.QueryParams == nil {
		return nil, false
	}
	return o.QueryParams, true
}

// HasQueryParams returns a boolean if a field has been set.
func (o *CompleteDestinationOAuthRequest) HasQueryParams() bool {
	if o != nil && o.QueryParams != nil {
		return true
	}

	return false
}

// SetQueryParams gets a reference to the given map[string]map[string]interface{} and assigns it to the QueryParams field.
func (o *CompleteDestinationOAuthRequest) SetQueryParams(v map[string]map[string]interface{}) {
	o.QueryParams = v
}

// GetOAuthInputConfiguration returns the OAuthInputConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CompleteDestinationOAuthRequest) GetOAuthInputConfiguration() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.OAuthInputConfiguration
}

// GetOAuthInputConfigurationOk returns a tuple with the OAuthInputConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CompleteDestinationOAuthRequest) GetOAuthInputConfigurationOk() (*interface{}, bool) {
	if o == nil || o.OAuthInputConfiguration == nil {
		return nil, false
	}
	return &o.OAuthInputConfiguration, true
}

// HasOAuthInputConfiguration returns a boolean if a field has been set.
func (o *CompleteDestinationOAuthRequest) HasOAuthInputConfiguration() bool {
	if o != nil && o.OAuthInputConfiguration != nil {
		return true
	}

	return false
}

// SetOAuthInputConfiguration gets a reference to the given interface{} and assigns it to the OAuthInputConfiguration field.
func (o *CompleteDestinationOAuthRequest) SetOAuthInputConfiguration(v interface{}) {
	o.OAuthInputConfiguration = v
}

func (o CompleteDestinationOAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destinationDefinitionId"] = o.DestinationDefinitionId
	}
	if true {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if o.RedirectUrl != nil {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if o.QueryParams != nil {
		toSerialize["queryParams"] = o.QueryParams
	}
	if o.OAuthInputConfiguration != nil {
		toSerialize["oAuthInputConfiguration"] = o.OAuthInputConfiguration
	}
	return json.Marshal(toSerialize)
}

type NullableCompleteDestinationOAuthRequest struct {
	value *CompleteDestinationOAuthRequest
	isSet bool
}

func (v NullableCompleteDestinationOAuthRequest) Get() *CompleteDestinationOAuthRequest {
	return v.value
}

func (v *NullableCompleteDestinationOAuthRequest) Set(val *CompleteDestinationOAuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteDestinationOAuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteDestinationOAuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteDestinationOAuthRequest(val *CompleteDestinationOAuthRequest) *NullableCompleteDestinationOAuthRequest {
	return &NullableCompleteDestinationOAuthRequest{value: val, isSet: true}
}

func (v NullableCompleteDestinationOAuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteDestinationOAuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


