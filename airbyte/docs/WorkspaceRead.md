# WorkspaceRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** |  | 
**CustomerId** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**InitialSetupComplete** | **bool** |  | 
**DisplaySetupWizard** | Pointer to **bool** |  | [optional] 
**AnonymousDataCollection** | Pointer to **bool** |  | [optional] 
**News** | Pointer to **bool** |  | [optional] 
**SecurityUpdates** | Pointer to **bool** |  | [optional] 
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**FirstCompletedSync** | Pointer to **bool** |  | [optional] 
**FeedbackDone** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkspaceRead

`func NewWorkspaceRead(workspaceId string, customerId string, name string, slug string, initialSetupComplete bool, ) *WorkspaceRead`

NewWorkspaceRead instantiates a new WorkspaceRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceReadWithDefaults

`func NewWorkspaceReadWithDefaults() *WorkspaceRead`

NewWorkspaceReadWithDefaults instantiates a new WorkspaceRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *WorkspaceRead) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceRead) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceRead) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetCustomerId

`func (o *WorkspaceRead) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *WorkspaceRead) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *WorkspaceRead) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetEmail

`func (o *WorkspaceRead) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WorkspaceRead) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WorkspaceRead) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WorkspaceRead) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceRead) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *WorkspaceRead) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *WorkspaceRead) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *WorkspaceRead) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetInitialSetupComplete

`func (o *WorkspaceRead) GetInitialSetupComplete() bool`

GetInitialSetupComplete returns the InitialSetupComplete field if non-nil, zero value otherwise.

### GetInitialSetupCompleteOk

`func (o *WorkspaceRead) GetInitialSetupCompleteOk() (*bool, bool)`

GetInitialSetupCompleteOk returns a tuple with the InitialSetupComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSetupComplete

`func (o *WorkspaceRead) SetInitialSetupComplete(v bool)`

SetInitialSetupComplete sets InitialSetupComplete field to given value.


### GetDisplaySetupWizard

`func (o *WorkspaceRead) GetDisplaySetupWizard() bool`

GetDisplaySetupWizard returns the DisplaySetupWizard field if non-nil, zero value otherwise.

### GetDisplaySetupWizardOk

`func (o *WorkspaceRead) GetDisplaySetupWizardOk() (*bool, bool)`

GetDisplaySetupWizardOk returns a tuple with the DisplaySetupWizard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySetupWizard

`func (o *WorkspaceRead) SetDisplaySetupWizard(v bool)`

SetDisplaySetupWizard sets DisplaySetupWizard field to given value.

### HasDisplaySetupWizard

`func (o *WorkspaceRead) HasDisplaySetupWizard() bool`

HasDisplaySetupWizard returns a boolean if a field has been set.

### GetAnonymousDataCollection

`func (o *WorkspaceRead) GetAnonymousDataCollection() bool`

GetAnonymousDataCollection returns the AnonymousDataCollection field if non-nil, zero value otherwise.

### GetAnonymousDataCollectionOk

`func (o *WorkspaceRead) GetAnonymousDataCollectionOk() (*bool, bool)`

GetAnonymousDataCollectionOk returns a tuple with the AnonymousDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousDataCollection

`func (o *WorkspaceRead) SetAnonymousDataCollection(v bool)`

SetAnonymousDataCollection sets AnonymousDataCollection field to given value.

### HasAnonymousDataCollection

`func (o *WorkspaceRead) HasAnonymousDataCollection() bool`

HasAnonymousDataCollection returns a boolean if a field has been set.

### GetNews

`func (o *WorkspaceRead) GetNews() bool`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *WorkspaceRead) GetNewsOk() (*bool, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *WorkspaceRead) SetNews(v bool)`

SetNews sets News field to given value.

### HasNews

`func (o *WorkspaceRead) HasNews() bool`

HasNews returns a boolean if a field has been set.

### GetSecurityUpdates

`func (o *WorkspaceRead) GetSecurityUpdates() bool`

GetSecurityUpdates returns the SecurityUpdates field if non-nil, zero value otherwise.

### GetSecurityUpdatesOk

`func (o *WorkspaceRead) GetSecurityUpdatesOk() (*bool, bool)`

GetSecurityUpdatesOk returns a tuple with the SecurityUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUpdates

`func (o *WorkspaceRead) SetSecurityUpdates(v bool)`

SetSecurityUpdates sets SecurityUpdates field to given value.

### HasSecurityUpdates

`func (o *WorkspaceRead) HasSecurityUpdates() bool`

HasSecurityUpdates returns a boolean if a field has been set.

### GetNotifications

`func (o *WorkspaceRead) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *WorkspaceRead) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *WorkspaceRead) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *WorkspaceRead) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetFirstCompletedSync

`func (o *WorkspaceRead) GetFirstCompletedSync() bool`

GetFirstCompletedSync returns the FirstCompletedSync field if non-nil, zero value otherwise.

### GetFirstCompletedSyncOk

`func (o *WorkspaceRead) GetFirstCompletedSyncOk() (*bool, bool)`

GetFirstCompletedSyncOk returns a tuple with the FirstCompletedSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstCompletedSync

`func (o *WorkspaceRead) SetFirstCompletedSync(v bool)`

SetFirstCompletedSync sets FirstCompletedSync field to given value.

### HasFirstCompletedSync

`func (o *WorkspaceRead) HasFirstCompletedSync() bool`

HasFirstCompletedSync returns a boolean if a field has been set.

### GetFeedbackDone

`func (o *WorkspaceRead) GetFeedbackDone() bool`

GetFeedbackDone returns the FeedbackDone field if non-nil, zero value otherwise.

### GetFeedbackDoneOk

`func (o *WorkspaceRead) GetFeedbackDoneOk() (*bool, bool)`

GetFeedbackDoneOk returns a tuple with the FeedbackDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackDone

`func (o *WorkspaceRead) SetFeedbackDone(v bool)`

SetFeedbackDone sets FeedbackDone field to given value.

### HasFeedbackDone

`func (o *WorkspaceRead) HasFeedbackDone() bool`

HasFeedbackDone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


