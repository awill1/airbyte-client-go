# WorkspaceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**InitialSetupComplete** | **bool** |  | 
**DisplaySetupWizard** | Pointer to **bool** |  | [optional] 
**AnonymousDataCollection** | **bool** |  | 
**News** | **bool** |  | 
**SecurityUpdates** | **bool** |  | 
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 

## Methods

### NewWorkspaceUpdate

`func NewWorkspaceUpdate(workspaceId string, initialSetupComplete bool, anonymousDataCollection bool, news bool, securityUpdates bool, ) *WorkspaceUpdate`

NewWorkspaceUpdate instantiates a new WorkspaceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceUpdateWithDefaults

`func NewWorkspaceUpdateWithDefaults() *WorkspaceUpdate`

NewWorkspaceUpdateWithDefaults instantiates a new WorkspaceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *WorkspaceUpdate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceUpdate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceUpdate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetEmail

`func (o *WorkspaceUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WorkspaceUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WorkspaceUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WorkspaceUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInitialSetupComplete

`func (o *WorkspaceUpdate) GetInitialSetupComplete() bool`

GetInitialSetupComplete returns the InitialSetupComplete field if non-nil, zero value otherwise.

### GetInitialSetupCompleteOk

`func (o *WorkspaceUpdate) GetInitialSetupCompleteOk() (*bool, bool)`

GetInitialSetupCompleteOk returns a tuple with the InitialSetupComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSetupComplete

`func (o *WorkspaceUpdate) SetInitialSetupComplete(v bool)`

SetInitialSetupComplete sets InitialSetupComplete field to given value.


### GetDisplaySetupWizard

`func (o *WorkspaceUpdate) GetDisplaySetupWizard() bool`

GetDisplaySetupWizard returns the DisplaySetupWizard field if non-nil, zero value otherwise.

### GetDisplaySetupWizardOk

`func (o *WorkspaceUpdate) GetDisplaySetupWizardOk() (*bool, bool)`

GetDisplaySetupWizardOk returns a tuple with the DisplaySetupWizard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySetupWizard

`func (o *WorkspaceUpdate) SetDisplaySetupWizard(v bool)`

SetDisplaySetupWizard sets DisplaySetupWizard field to given value.

### HasDisplaySetupWizard

`func (o *WorkspaceUpdate) HasDisplaySetupWizard() bool`

HasDisplaySetupWizard returns a boolean if a field has been set.

### GetAnonymousDataCollection

`func (o *WorkspaceUpdate) GetAnonymousDataCollection() bool`

GetAnonymousDataCollection returns the AnonymousDataCollection field if non-nil, zero value otherwise.

### GetAnonymousDataCollectionOk

`func (o *WorkspaceUpdate) GetAnonymousDataCollectionOk() (*bool, bool)`

GetAnonymousDataCollectionOk returns a tuple with the AnonymousDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousDataCollection

`func (o *WorkspaceUpdate) SetAnonymousDataCollection(v bool)`

SetAnonymousDataCollection sets AnonymousDataCollection field to given value.


### GetNews

`func (o *WorkspaceUpdate) GetNews() bool`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *WorkspaceUpdate) GetNewsOk() (*bool, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *WorkspaceUpdate) SetNews(v bool)`

SetNews sets News field to given value.


### GetSecurityUpdates

`func (o *WorkspaceUpdate) GetSecurityUpdates() bool`

GetSecurityUpdates returns the SecurityUpdates field if non-nil, zero value otherwise.

### GetSecurityUpdatesOk

`func (o *WorkspaceUpdate) GetSecurityUpdatesOk() (*bool, bool)`

GetSecurityUpdatesOk returns a tuple with the SecurityUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUpdates

`func (o *WorkspaceUpdate) SetSecurityUpdates(v bool)`

SetSecurityUpdates sets SecurityUpdates field to given value.


### GetNotifications

`func (o *WorkspaceUpdate) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *WorkspaceUpdate) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *WorkspaceUpdate) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *WorkspaceUpdate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


