# OrderDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**NumberWithPrefix** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayDate** | Pointer to **time.Time** |  | [optional] 
**NewOrderArchitecture** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Pivot** | Pointer to [**OrderDocumentPivot**](OrderDocumentPivot.md) |  | [optional] 

## Methods

### NewOrderDocument

`func NewOrderDocument() *OrderDocument`

NewOrderDocument instantiates a new OrderDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDocumentWithDefaults

`func NewOrderDocumentWithDefaults() *OrderDocument`

NewOrderDocumentWithDefaults instantiates a new OrderDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderDocument) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderDocument) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderDocument) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *OrderDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderDocument) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderDocument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *OrderDocument) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OrderDocument) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OrderDocument) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OrderDocument) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetNumberWithPrefix

`func (o *OrderDocument) GetNumberWithPrefix() string`

GetNumberWithPrefix returns the NumberWithPrefix field if non-nil, zero value otherwise.

### GetNumberWithPrefixOk

`func (o *OrderDocument) GetNumberWithPrefixOk() (*string, bool)`

GetNumberWithPrefixOk returns a tuple with the NumberWithPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberWithPrefix

`func (o *OrderDocument) SetNumberWithPrefix(v string)`

SetNumberWithPrefix sets NumberWithPrefix field to given value.

### HasNumberWithPrefix

`func (o *OrderDocument) HasNumberWithPrefix() bool`

HasNumberWithPrefix returns a boolean if a field has been set.

### GetPath

`func (o *OrderDocument) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *OrderDocument) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *OrderDocument) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *OrderDocument) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetUserId

`func (o *OrderDocument) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrderDocument) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrderDocument) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OrderDocument) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSource

`func (o *OrderDocument) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OrderDocument) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OrderDocument) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *OrderDocument) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderDocument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderDocument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderDocument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderDocument) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderDocument) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderDocument) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDisplayDate

`func (o *OrderDocument) GetDisplayDate() time.Time`

GetDisplayDate returns the DisplayDate field if non-nil, zero value otherwise.

### GetDisplayDateOk

`func (o *OrderDocument) GetDisplayDateOk() (*time.Time, bool)`

GetDisplayDateOk returns a tuple with the DisplayDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDate

`func (o *OrderDocument) SetDisplayDate(v time.Time)`

SetDisplayDate sets DisplayDate field to given value.

### HasDisplayDate

`func (o *OrderDocument) HasDisplayDate() bool`

HasDisplayDate returns a boolean if a field has been set.

### GetNewOrderArchitecture

`func (o *OrderDocument) GetNewOrderArchitecture() string`

GetNewOrderArchitecture returns the NewOrderArchitecture field if non-nil, zero value otherwise.

### GetNewOrderArchitectureOk

`func (o *OrderDocument) GetNewOrderArchitectureOk() (*string, bool)`

GetNewOrderArchitectureOk returns a tuple with the NewOrderArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOrderArchitecture

`func (o *OrderDocument) SetNewOrderArchitecture(v string)`

SetNewOrderArchitecture sets NewOrderArchitecture field to given value.

### HasNewOrderArchitecture

`func (o *OrderDocument) HasNewOrderArchitecture() bool`

HasNewOrderArchitecture returns a boolean if a field has been set.

### GetStatus

`func (o *OrderDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPivot

`func (o *OrderDocument) GetPivot() OrderDocumentPivot`

GetPivot returns the Pivot field if non-nil, zero value otherwise.

### GetPivotOk

`func (o *OrderDocument) GetPivotOk() (*OrderDocumentPivot, bool)`

GetPivotOk returns a tuple with the Pivot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivot

`func (o *OrderDocument) SetPivot(v OrderDocumentPivot)`

SetPivot sets Pivot field to given value.

### HasPivot

`func (o *OrderDocument) HasPivot() bool`

HasPivot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


