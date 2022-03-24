# NotificationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewNotificationRead

`func NewNotificationRead(status string, ) *NotificationRead`

NewNotificationRead instantiates a new NotificationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationReadWithDefaults

`func NewNotificationReadWithDefaults() *NotificationRead`

NewNotificationReadWithDefaults instantiates a new NotificationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NotificationRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotificationRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotificationRead) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *NotificationRead) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotificationRead) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotificationRead) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *NotificationRead) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


