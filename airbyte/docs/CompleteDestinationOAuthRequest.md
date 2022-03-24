# CompleteDestinationOAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationDefinitionId** | **string** |  | 
**WorkspaceId** | **string** |  | 
**RedirectUrl** | Pointer to **string** | When completing OAuth flow to gain an access token, some API sometimes requires to verify that the app re-send the redirectUrl that was used when consent was given. | [optional] 
**QueryParams** | Pointer to **map[string]map[string]interface{}** | The query parameters present in the redirect URL after a user granted consent e.g auth code | [optional] 
**OAuthInputConfiguration** | Pointer to **interface{}** | OAuth specific blob. | [optional] 

## Methods

### NewCompleteDestinationOAuthRequest

`func NewCompleteDestinationOAuthRequest(destinationDefinitionId string, workspaceId string, ) *CompleteDestinationOAuthRequest`

NewCompleteDestinationOAuthRequest instantiates a new CompleteDestinationOAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteDestinationOAuthRequestWithDefaults

`func NewCompleteDestinationOAuthRequestWithDefaults() *CompleteDestinationOAuthRequest`

NewCompleteDestinationOAuthRequestWithDefaults instantiates a new CompleteDestinationOAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationDefinitionId

`func (o *CompleteDestinationOAuthRequest) GetDestinationDefinitionId() string`

GetDestinationDefinitionId returns the DestinationDefinitionId field if non-nil, zero value otherwise.

### GetDestinationDefinitionIdOk

`func (o *CompleteDestinationOAuthRequest) GetDestinationDefinitionIdOk() (*string, bool)`

GetDestinationDefinitionIdOk returns a tuple with the DestinationDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDefinitionId

`func (o *CompleteDestinationOAuthRequest) SetDestinationDefinitionId(v string)`

SetDestinationDefinitionId sets DestinationDefinitionId field to given value.


### GetWorkspaceId

`func (o *CompleteDestinationOAuthRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *CompleteDestinationOAuthRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *CompleteDestinationOAuthRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetRedirectUrl

`func (o *CompleteDestinationOAuthRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CompleteDestinationOAuthRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CompleteDestinationOAuthRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CompleteDestinationOAuthRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetQueryParams

`func (o *CompleteDestinationOAuthRequest) GetQueryParams() map[string]map[string]interface{}`

GetQueryParams returns the QueryParams field if non-nil, zero value otherwise.

### GetQueryParamsOk

`func (o *CompleteDestinationOAuthRequest) GetQueryParamsOk() (*map[string]map[string]interface{}, bool)`

GetQueryParamsOk returns a tuple with the QueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParams

`func (o *CompleteDestinationOAuthRequest) SetQueryParams(v map[string]map[string]interface{})`

SetQueryParams sets QueryParams field to given value.

### HasQueryParams

`func (o *CompleteDestinationOAuthRequest) HasQueryParams() bool`

HasQueryParams returns a boolean if a field has been set.

### GetOAuthInputConfiguration

`func (o *CompleteDestinationOAuthRequest) GetOAuthInputConfiguration() interface{}`

GetOAuthInputConfiguration returns the OAuthInputConfiguration field if non-nil, zero value otherwise.

### GetOAuthInputConfigurationOk

`func (o *CompleteDestinationOAuthRequest) GetOAuthInputConfigurationOk() (*interface{}, bool)`

GetOAuthInputConfigurationOk returns a tuple with the OAuthInputConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOAuthInputConfiguration

`func (o *CompleteDestinationOAuthRequest) SetOAuthInputConfiguration(v interface{})`

SetOAuthInputConfiguration sets OAuthInputConfiguration field to given value.

### HasOAuthInputConfiguration

`func (o *CompleteDestinationOAuthRequest) HasOAuthInputConfiguration() bool`

HasOAuthInputConfiguration returns a boolean if a field has been set.

### SetOAuthInputConfigurationNil

`func (o *CompleteDestinationOAuthRequest) SetOAuthInputConfigurationNil(b bool)`

 SetOAuthInputConfigurationNil sets the value for OAuthInputConfiguration to be an explicit nil

### UnsetOAuthInputConfiguration
`func (o *CompleteDestinationOAuthRequest) UnsetOAuthInputConfiguration()`

UnsetOAuthInputConfiguration ensures that no value is present for OAuthInputConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


