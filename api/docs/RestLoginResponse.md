# RestLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConcurrentSessions** | Pointer to **int32** |  | [optional] 
**MaxCurrentSessions** | Pointer to **int32** |  | [optional] 
**UserId** | Pointer to **int32** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**AccessToken** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 

## Methods

### NewRestLoginResponse

`func NewRestLoginResponse() *RestLoginResponse`

NewRestLoginResponse instantiates a new RestLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestLoginResponseWithDefaults

`func NewRestLoginResponseWithDefaults() *RestLoginResponse`

NewRestLoginResponseWithDefaults instantiates a new RestLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConcurrentSessions

`func (o *RestLoginResponse) GetConcurrentSessions() int32`

GetConcurrentSessions returns the ConcurrentSessions field if non-nil, zero value otherwise.

### GetConcurrentSessionsOk

`func (o *RestLoginResponse) GetConcurrentSessionsOk() (*int32, bool)`

GetConcurrentSessionsOk returns a tuple with the ConcurrentSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrentSessions

`func (o *RestLoginResponse) SetConcurrentSessions(v int32)`

SetConcurrentSessions sets ConcurrentSessions field to given value.

### HasConcurrentSessions

`func (o *RestLoginResponse) HasConcurrentSessions() bool`

HasConcurrentSessions returns a boolean if a field has been set.

### GetMaxCurrentSessions

`func (o *RestLoginResponse) GetMaxCurrentSessions() int32`

GetMaxCurrentSessions returns the MaxCurrentSessions field if non-nil, zero value otherwise.

### GetMaxCurrentSessionsOk

`func (o *RestLoginResponse) GetMaxCurrentSessionsOk() (*int32, bool)`

GetMaxCurrentSessionsOk returns a tuple with the MaxCurrentSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCurrentSessions

`func (o *RestLoginResponse) SetMaxCurrentSessions(v int32)`

SetMaxCurrentSessions sets MaxCurrentSessions field to given value.

### HasMaxCurrentSessions

`func (o *RestLoginResponse) HasMaxCurrentSessions() bool`

HasMaxCurrentSessions returns a boolean if a field has been set.

### GetUserId

`func (o *RestLoginResponse) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RestLoginResponse) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RestLoginResponse) SetUserId(v int32)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RestLoginResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTokenType

`func (o *RestLoginResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *RestLoginResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *RestLoginResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *RestLoginResponse) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *RestLoginResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *RestLoginResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *RestLoginResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *RestLoginResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetAccessToken

`func (o *RestLoginResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *RestLoginResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *RestLoginResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *RestLoginResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RestLoginResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RestLoginResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RestLoginResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *RestLoginResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


