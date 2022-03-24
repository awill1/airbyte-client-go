# WorkspaceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**AnonymousDataCollection** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**News** | Pointer to **bool** |  | [optional] 
**SecurityUpdates** | Pointer to **bool** |  | [optional] 
**Notifications** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 
**DisplaySetupWizard** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkspaceCreate

`func NewWorkspaceCreate(name string, ) *WorkspaceCreate`

NewWorkspaceCreate instantiates a new WorkspaceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCreateWithDefaults

`func NewWorkspaceCreateWithDefaults() *WorkspaceCreate`

NewWorkspaceCreateWithDefaults instantiates a new WorkspaceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *WorkspaceCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WorkspaceCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WorkspaceCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WorkspaceCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAnonymousDataCollection

`func (o *WorkspaceCreate) GetAnonymousDataCollection() bool`

GetAnonymousDataCollection returns the AnonymousDataCollection field if non-nil, zero value otherwise.

### GetAnonymousDataCollectionOk

`func (o *WorkspaceCreate) GetAnonymousDataCollectionOk() (*bool, bool)`

GetAnonymousDataCollectionOk returns a tuple with the AnonymousDataCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousDataCollection

`func (o *WorkspaceCreate) SetAnonymousDataCollection(v bool)`

SetAnonymousDataCollection sets AnonymousDataCollection field to given value.

### HasAnonymousDataCollection

`func (o *WorkspaceCreate) HasAnonymousDataCollection() bool`

HasAnonymousDataCollection returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceCreate) SetName(v string)`

SetName sets Name field to given value.


### GetNews

`func (o *WorkspaceCreate) GetNews() bool`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *WorkspaceCreate) GetNewsOk() (*bool, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *WorkspaceCreate) SetNews(v bool)`

SetNews sets News field to given value.

### HasNews

`func (o *WorkspaceCreate) HasNews() bool`

HasNews returns a boolean if a field has been set.

### GetSecurityUpdates

`func (o *WorkspaceCreate) GetSecurityUpdates() bool`

GetSecurityUpdates returns the SecurityUpdates field if non-nil, zero value otherwise.

### GetSecurityUpdatesOk

`func (o *WorkspaceCreate) GetSecurityUpdatesOk() (*bool, bool)`

GetSecurityUpdatesOk returns a tuple with the SecurityUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityUpdates

`func (o *WorkspaceCreate) SetSecurityUpdates(v bool)`

SetSecurityUpdates sets SecurityUpdates field to given value.

### HasSecurityUpdates

`func (o *WorkspaceCreate) HasSecurityUpdates() bool`

HasSecurityUpdates returns a boolean if a field has been set.

### GetNotifications

`func (o *WorkspaceCreate) GetNotifications() []Notification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *WorkspaceCreate) GetNotificationsOk() (*[]Notification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *WorkspaceCreate) SetNotifications(v []Notification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *WorkspaceCreate) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetDisplaySetupWizard

`func (o *WorkspaceCreate) GetDisplaySetupWizard() bool`

GetDisplaySetupWizard returns the DisplaySetupWizard field if non-nil, zero value otherwise.

### GetDisplaySetupWizardOk

`func (o *WorkspaceCreate) GetDisplaySetupWizardOk() (*bool, bool)`

GetDisplaySetupWizardOk returns a tuple with the DisplaySetupWizard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySetupWizard

`func (o *WorkspaceCreate) SetDisplaySetupWizard(v bool)`

SetDisplaySetupWizard sets DisplaySetupWizard field to given value.

### HasDisplaySetupWizard

`func (o *WorkspaceCreate) HasDisplaySetupWizard() bool`

HasDisplaySetupWizard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


