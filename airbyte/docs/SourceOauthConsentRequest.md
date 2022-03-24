# SourceOauthConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDefinitionId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**RedirectUrl** | **string** | The url to redirect to after getting the user consent | 
**OAuthInputConfiguration** | Pointer to **interface{}** | OAuth specific blob. | [optional] 

## Methods

### NewSourceOauthConsentRequest

`func NewSourceOauthConsentRequest(sourceDefinitionId string, workspaceId string, redirectUrl string, ) *SourceOauthConsentRequest`

NewSourceOauthConsentRequest instantiates a new SourceOauthConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceOauthConsentRequestWithDefaults

`func NewSourceOauthConsentRequestWithDefaults() *SourceOauthConsentRequest`

NewSourceOauthConsentRequestWithDefaults instantiates a new SourceOauthConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDefinitionId

`func (o *SourceOauthConsentRequest) GetSourceDefinitionId() string`

GetSourceDefinitionId returns the SourceDefinitionId field if non-nil, zero value otherwise.

### GetSourceDefinitionIdOk

`func (o *SourceOauthConsentRequest) GetSourceDefinitionIdOk() (*string, bool)`

GetSourceDefinitionIdOk returns a tuple with the SourceDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinitionId

`func (o *SourceOauthConsentRequest) SetSourceDefinitionId(v string)`

SetSourceDefinitionId sets SourceDefinitionId field to given value.


### GetWorkspaceId

`func (o *SourceOauthConsentRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SourceOauthConsentRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SourceOauthConsentRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRedirectUrl

`func (o *SourceOauthConsentRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SourceOauthConsentRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SourceOauthConsentRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetOAuthInputConfiguration

`func (o *SourceOauthConsentRequest) GetOAuthInputConfiguration() interface{}`

GetOAuthInputConfiguration returns the OAuthInputConfiguration field if non-nil, zero value otherwise.

### GetOAuthInputConfigurationOk

`func (o *SourceOauthConsentRequest) GetOAuthInputConfigurationOk() (*interface{}, bool)`

GetOAuthInputConfigurationOk returns a tuple with the OAuthInputConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthInputConfiguration

`func (o *SourceOauthConsentRequest) SetOAuthInputConfiguration(v interface{})`

SetOAuthInputConfiguration sets OAuthInputConfiguration field to given value.

### HasOAuthInputConfiguration

`func (o *SourceOauthConsentRequest) HasOAuthInputConfiguration() bool`

HasOAuthInputConfiguration returns a boolean if a field has been set.

### SetOAuthInputConfigurationNil

`func (o *SourceOauthConsentRequest) SetOAuthInputConfigurationNil(b bool)`

 SetOAuthInputConfigurationNil sets the value for OAuthInputConfiguration to be an explicit nil

### UnsetOAuthInputConfiguration
`func (o *SourceOauthConsentRequest) UnsetOAuthInputConfiguration()`

UnsetOAuthInputConfiguration ensures that no value is present for OAuthInputConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


