# OAuth2Specification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootObject** | **[]interface{}** | A list of strings representing a pointer to the root object which contains any oauth parameters in the ConnectorSpecification. Examples: if oauth parameters were contained inside the top level, rootObject&#x3D;[] If they were nested inside another object {&#39;credentials&#39;: {&#39;app_id&#39; etc...}, rootObject&#x3D;[&#39;credentials&#39;] If they were inside a oneOf {&#39;switch&#39;: {oneOf: [{client_id...}, {non_oauth_param]}},  rootObject&#x3D;[&#39;switch&#39;, 0]  | 
**OauthFlowInitParameters** | **[][]string** | Pointers to the fields in the rootObject needed to obtain the initial refresh/access tokens for the OAuth flow. Each inner array represents the path in the rootObject of the referenced field. For example. Assume the rootObject contains params &#39;app_secret&#39;, &#39;app_id&#39; which are needed to get the initial refresh token. If they are not nested in the rootObject, then the array would look like this [[&#39;app_secret&#39;], [&#39;app_id&#39;]] If they are nested inside an object called &#39;auth_params&#39; then this array would be [[&#39;auth_params&#39;, &#39;app_secret&#39;], [&#39;auth_params&#39;, &#39;app_id&#39;]] | 
**OauthFlowOutputParameters** | **[][]string** | Pointers to the fields in the rootObject which can be populated from successfully completing the oauth flow using the init parameters. This is typically a refresh/access token. Each inner array represents the path in the rootObject of the referenced field. | 

## Methods

### NewOAuth2Specification

`func NewOAuth2Specification(rootObject []interface{}, oauthFlowInitParameters [][]string, oauthFlowOutputParameters [][]string, ) *OAuth2Specification`

NewOAuth2Specification instantiates a new OAuth2Specification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2SpecificationWithDefaults

`func NewOAuth2SpecificationWithDefaults() *OAuth2Specification`

NewOAuth2SpecificationWithDefaults instantiates a new OAuth2Specification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootObject

`func (o *OAuth2Specification) GetRootObject() []interface{}`

GetRootObject returns the RootObject field if non-nil, zero value otherwise.

### GetRootObjectOk

`func (o *OAuth2Specification) GetRootObjectOk() (*[]interface{}, bool)`

GetRootObjectOk returns a tuple with the RootObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootObject

`func (o *OAuth2Specification) SetRootObject(v []interface{})`

SetRootObject sets RootObject field to given value.


### GetOauthFlowInitParameters

`func (o *OAuth2Specification) GetOauthFlowInitParameters() [][]string`

GetOauthFlowInitParameters returns the OauthFlowInitParameters field if non-nil, zero value otherwise.

### GetOauthFlowInitParametersOk

`func (o *OAuth2Specification) GetOauthFlowInitParametersOk() (*[][]string, bool)`

GetOauthFlowInitParametersOk returns a tuple with the OauthFlowInitParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthFlowInitParameters

`func (o *OAuth2Specification) SetOauthFlowInitParameters(v [][]string)`

SetOauthFlowInitParameters sets OauthFlowInitParameters field to given value.


### GetOauthFlowOutputParameters

`func (o *OAuth2Specification) GetOauthFlowOutputParameters() [][]string`

GetOauthFlowOutputParameters returns the OauthFlowOutputParameters field if non-nil, zero value otherwise.

### GetOauthFlowOutputParametersOk

`func (o *OAuth2Specification) GetOauthFlowOutputParametersOk() (*[][]string, bool)`

GetOauthFlowOutputParametersOk returns a tuple with the OauthFlowOutputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthFlowOutputParameters

`func (o *OAuth2Specification) SetOauthFlowOutputParameters(v [][]string)`

SetOauthFlowOutputParameters sets OauthFlowOutputParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


