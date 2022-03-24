# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NotificationType** | [**NotificationType**](NotificationType.md) |  | 
**SendOnSuccess** | **bool** |  | [default to false]
**SendOnFailure** | **bool** |  | [default to true]
**SlackConfiguration** | Pointer to [**SlackNotificationConfiguration**](SlackNotificationConfiguration.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification(notificationType NotificationType, sendOnSuccess bool, sendOnFailure bool, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotificationType

`func (o *Notification) GetNotificationType() NotificationType`

GetNotificationType returns the NotificationType field if non-nil, zero value otherwise.

### GetNotificationTypeOk

`func (o *Notification) GetNotificationTypeOk() (*NotificationType, bool)`

GetNotificationTypeOk returns a tuple with the NotificationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationType

`func (o *Notification) SetNotificationType(v NotificationType)`

SetNotificationType sets NotificationType field to given value.


### GetSendOnSuccess

`func (o *Notification) GetSendOnSuccess() bool`

GetSendOnSuccess returns the SendOnSuccess field if non-nil, zero value otherwise.

### GetSendOnSuccessOk

`func (o *Notification) GetSendOnSuccessOk() (*bool, bool)`

GetSendOnSuccessOk returns a tuple with the SendOnSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnSuccess

`func (o *Notification) SetSendOnSuccess(v bool)`

SetSendOnSuccess sets SendOnSuccess field to given value.


### GetSendOnFailure

`func (o *Notification) GetSendOnFailure() bool`

GetSendOnFailure returns the SendOnFailure field if non-nil, zero value otherwise.

### GetSendOnFailureOk

`func (o *Notification) GetSendOnFailureOk() (*bool, bool)`

GetSendOnFailureOk returns a tuple with the SendOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnFailure

`func (o *Notification) SetSendOnFailure(v bool)`

SetSendOnFailure sets SendOnFailure field to given value.


### GetSlackConfiguration

`func (o *Notification) GetSlackConfiguration() SlackNotificationConfiguration`

GetSlackConfiguration returns the SlackConfiguration field if non-nil, zero value otherwise.

### GetSlackConfigurationOk

`func (o *Notification) GetSlackConfigurationOk() (*SlackNotificationConfiguration, bool)`

GetSlackConfigurationOk returns a tuple with the SlackConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackConfiguration

`func (o *Notification) SetSlackConfiguration(v SlackNotificationConfiguration)`

SetSlackConfiguration sets SlackConfiguration field to given value.

### HasSlackConfiguration

`func (o *Notification) HasSlackConfiguration() bool`

HasSlackConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


