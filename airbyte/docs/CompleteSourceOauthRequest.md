# CompleteSourceOauthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**RedirectUrl** | Pointer to **string** | When completing OAuth flow to gain an access token, some API sometimes requires to verify that the app re-send the redirectUrl that was used when consent was given. | [optional] 
**QueryParams** | Pointer to **map[string]map[string]interface{}** | The query parameters present in the redirect URL after a user granted consent e.g auth code | [optional] 
**OAuthInputConfiguration** | Pointer to **interface{}** | OAuth specific blob. | [optional] 

## Methods

### NewCompleteSourceOauthRequest

`func NewCompleteSourceOauthRequest(sourceDefinitionId string, workspaceId string, ) *CompleteSourceOauthRequest`

NewCompleteSourceOauthRequest instantiates a new CompleteSourceOauthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteSourceOauthRequestWithDefaults

`func NewCompleteSourceOauthRequestWithDefaults() *CompleteSourceOauthRequest`

NewCompleteSourceOauthRequestWithDefaults instantiates a new CompleteSourceOauthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *CompleteSourceOauthRequest) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *CompleteSourceOauthRequest) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *CompleteSourceOauthRequest) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.


### GetWorkspaceId

`func (o *CompleteSourceOauthRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CompleteSourceOauthRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CompleteSourceOauthRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRedirectUrl

`func (o *CompleteSourceOauthRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CompleteSourceOauthRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CompleteSourceOauthRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CompleteSourceOauthRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetQueryParams

`func (o *CompleteSourceOauthRequest) GetQueryParams() map[string]map[string]interface{}`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *CompleteSourceOauthRequest) GetQueryParamsOk() (*map[string]map[string]interface{}, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *CompleteSourceOauthRequest) SetQueryParams(v map[string]map[string]interface{})`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *CompleteSourceOauthRequest) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetOAuthInputConfiguration

`func (o *CompleteSourceOauthRequest) GetOAuthInputConfiguration() interface{}`

GetOAuthInputConfiguration returns the OAuthInputConfiguration field if non-nil, zero value otherwise.

### GetOAuthInputConfigurationOk

`func (o *CompleteSourceOauthRequest) GetOAuthInputConfigurationOk() (*interface{}, bool)`

GetOAuthInputConfigurationOk returns a tuple with the OAuthInputConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthInputConfiguration

`func (o *CompleteSourceOauthRequest) SetOAuthInputConfiguration(v interface{})`

SetOAuthInputConfiguration sets OAuthInputConfiguration field to given value.

### HasOAuthInputConfiguration

`func (o *CompleteSourceOauthRequest) HasOAuthInputConfiguration() bool`

HasOAuthInputConfiguration returns a boolean if a field has been set.

### SetOAuthInputConfigurationNil

`func (o *CompleteSourceOauthRequest) SetOAuthInputConfigurationNil(b bool)`

 SetOAuthInputConfigurationNil sets the value for OAuthInputConfiguration to be an explicit nil

### UnsetOAuthInputConfiguration
`func (o *CompleteSourceOauthRequest) UnsetOAuthInputConfiguration()`

UnsetOAuthInputConfiguration ensures that no value is present for OAuthInputConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


