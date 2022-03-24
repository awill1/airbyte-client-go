# DestinationOauthConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**RedirectUrl** | **string** | The url to redirect to after getting the user consent | 
**OAuthInputConfiguration** | Pointer to **interface{}** | OAuth specific blob. | [optional] 

## Methods

### NewDestinationOauthConsentRequest

`func NewDestinationOauthConsentRequest(destinationDefinitionId string, workspaceId string, redirectUrl string, ) *DestinationOauthConsentRequest`

NewDestinationOauthConsentRequest instantiates a new DestinationOauthConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationOauthConsentRequestWithDefaults

`func NewDestinationOauthConsentRequestWithDefaults() *DestinationOauthConsentRequest`

NewDestinationOauthConsentRequestWithDefaults instantiates a new DestinationOauthConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *DestinationOauthConsentRequest) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *DestinationOauthConsentRequest) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *DestinationOauthConsentRequest) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetWorkspaceId

`func (o *DestinationOauthConsentRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DestinationOauthConsentRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DestinationOauthConsentRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRedirectUrl

`func (o *DestinationOauthConsentRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *DestinationOauthConsentRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *DestinationOauthConsentRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetOAuthInputConfiguration

`func (o *DestinationOauthConsentRequest) GetOAuthInputConfiguration() interface{}`

GetOAuthInputConfiguration returns the OAuthInputConfiguration field if non-nil, zero value otherwise.

### GetOAuthInputConfigurationOk

`func (o *DestinationOauthConsentRequest) GetOAuthInputConfigurationOk() (*interface{}, bool)`

GetOAuthInputConfigurationOk returns a tuple with the OAuthInputConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthInputConfiguration

`func (o *DestinationOauthConsentRequest) SetOAuthInputConfiguration(v interface{})`

SetOAuthInputConfiguration sets OAuthInputConfiguration field to given value.

### HasOAuthInputConfiguration

`func (o *DestinationOauthConsentRequest) HasOAuthInputConfiguration() bool`

HasOAuthInputConfiguration returns a boolean if a field has been set.

### SetOAuthInputConfigurationNil

`func (o *DestinationOauthConsentRequest) SetOAuthInputConfigurationNil(b bool)`

 SetOAuthInputConfigurationNil sets the value for OAuthInputConfiguration to be an explicit nil

### UnsetOAuthInputConfiguration
`func (o *DestinationOauthConsentRequest) UnsetOAuthInputConfiguration()`

UnsetOAuthInputConfiguration ensures that no value is present for OAuthInputConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


