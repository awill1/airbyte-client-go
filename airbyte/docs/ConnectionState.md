# ConnectionState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** |  | 
**State** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConnectionState

`func NewConnectionState(connectionId string, ) *ConnectionState`

NewConnectionState instantiates a new ConnectionState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionStateWithDefaults

`func NewConnectionStateWithDefaults() *ConnectionState`

NewConnectionStateWithDefaults instantiates a new ConnectionState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ConnectionState) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConnectionState) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConnectionState) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetState

`func (o *ConnectionState) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectionState) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectionState) SetState(v map[string]interface{})`

SetState sets State field to given value.

### HasState

`func (o *ConnectionState) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


