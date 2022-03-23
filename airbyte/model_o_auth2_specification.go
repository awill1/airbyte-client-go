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

// OAuth2Specification An object containing any metadata needed to describe this connector's Oauth flow
type OAuth2Specification struct {
	// A list of strings representing a pointer to the root object which contains any oauth parameters in the ConnectorSpecification. Examples: if oauth parameters were contained inside the top level, rootObject=[] If they were nested inside another object {'credentials': {'app_id' etc...}, rootObject=['credentials'] If they were inside a oneOf {'switch': {oneOf: [{client_id...}, {non_oauth_param]}},  rootObject=['switch', 0] 
	RootObject []interface{} `json:"rootObject"`
	// Pointers to the fields in the rootObject needed to obtain the initial refresh/access tokens for the OAuth flow. Each inner array represents the path in the rootObject of the referenced field. For example. Assume the rootObject contains params 'app_secret', 'app_id' which are needed to get the initial refresh token. If they are not nested in the rootObject, then the array would look like this [['app_secret'], ['app_id']] If they are nested inside an object called 'auth_params' then this array would be [['auth_params', 'app_secret'], ['auth_params', 'app_id']]
	OauthFlowInitParameters [][]string `json:"oauthFlowInitParameters"`
	// Pointers to the fields in the rootObject which can be populated from successfully completing the oauth flow using the init parameters. This is typically a refresh/access token. Each inner array represents the path in the rootObject of the referenced field.
	OauthFlowOutputParameters [][]string `json:"oauthFlowOutputParameters"`
}

// NewOAuth2Specification instantiates a new OAuth2Specification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Specification(rootObject []interface{}, oauthFlowInitParameters [][]string, oauthFlowOutputParameters [][]string) *OAuth2Specification {
	this := OAuth2Specification{}
	this.RootObject = rootObject
	this.OauthFlowInitParameters = oauthFlowInitParameters
	this.OauthFlowOutputParameters = oauthFlowOutputParameters
	return &this
}

// NewOAuth2SpecificationWithDefaults instantiates a new OAuth2Specification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2SpecificationWithDefaults() *OAuth2Specification {
	this := OAuth2Specification{}
	return &this
}

// GetRootObject returns the RootObject field value
func (o *OAuth2Specification) GetRootObject() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}

	return o.RootObject
}

// GetRootObjectOk returns a tuple with the RootObject field value
// and a boolean to check if the value has been set.
func (o *OAuth2Specification) GetRootObjectOk() ([]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RootObject, true
}

// SetRootObject sets field value
func (o *OAuth2Specification) SetRootObject(v []interface{}) {
	o.RootObject = v
}

// GetOauthFlowInitParameters returns the OauthFlowInitParameters field value
func (o *OAuth2Specification) GetOauthFlowInitParameters() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.OauthFlowInitParameters
}

// GetOauthFlowInitParametersOk returns a tuple with the OauthFlowInitParameters field value
// and a boolean to check if the value has been set.
func (o *OAuth2Specification) GetOauthFlowInitParametersOk() ([][]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OauthFlowInitParameters, true
}

// SetOauthFlowInitParameters sets field value
func (o *OAuth2Specification) SetOauthFlowInitParameters(v [][]string) {
	o.OauthFlowInitParameters = v
}

// GetOauthFlowOutputParameters returns the OauthFlowOutputParameters field value
func (o *OAuth2Specification) GetOauthFlowOutputParameters() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.OauthFlowOutputParameters
}

// GetOauthFlowOutputParametersOk returns a tuple with the OauthFlowOutputParameters field value
// and a boolean to check if the value has been set.
func (o *OAuth2Specification) GetOauthFlowOutputParametersOk() ([][]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OauthFlowOutputParameters, true
}

// SetOauthFlowOutputParameters sets field value
func (o *OAuth2Specification) SetOauthFlowOutputParameters(v [][]string) {
	o.OauthFlowOutputParameters = v
}

func (o OAuth2Specification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rootObject"] = o.RootObject
	}
	if true {
		toSerialize["oauthFlowInitParameters"] = o.OauthFlowInitParameters
	}
	if true {
		toSerialize["oauthFlowOutputParameters"] = o.OauthFlowOutputParameters
	}
	return json.Marshal(toSerialize)
}

type NullableOAuth2Specification struct {
	value *OAuth2Specification
	isSet bool
}

func (v NullableOAuth2Specification) Get() *OAuth2Specification {
	return v.value
}

func (v *NullableOAuth2Specification) Set(val *OAuth2Specification) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Specification) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Specification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Specification(val *OAuth2Specification) *NullableOAuth2Specification {
	return &NullableOAuth2Specification{value: val, isSet: true}
}

func (v NullableOAuth2Specification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Specification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

