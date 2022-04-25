# MultipleVariationsUpdatedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to [**map[string]Variation**](Variation.md) |  | [optional] 
**Failed** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewMultipleVariationsUpdatedResponse

`func NewMultipleVariationsUpdatedResponse() *MultipleVariationsUpdatedResponse`

NewMultipleVariationsUpdatedResponse instantiates a new MultipleVariationsUpdatedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipleVariationsUpdatedResponseWithDefaults

`func NewMultipleVariationsUpdatedResponseWithDefaults() *MultipleVariationsUpdatedResponse`

NewMultipleVariationsUpdatedResponseWithDefaults instantiates a new MultipleVariationsUpdatedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MultipleVariationsUpdatedResponse) GetSuccess() map[string]Variation`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MultipleVariationsUpdatedResponse) GetSuccessOk() (*map[string]Variation, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MultipleVariationsUpdatedResponse) SetSuccess(v map[string]Variation)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MultipleVariationsUpdatedResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *MultipleVariationsUpdatedResponse) GetFailed() map[string][]string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MultipleVariationsUpdatedResponse) GetFailedOk() (*map[string][]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MultipleVariationsUpdatedResponse) SetFailed(v map[string][]string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MultipleVariationsUpdatedResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


