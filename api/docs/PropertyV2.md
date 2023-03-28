# PropertyV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PropertyId** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TargetId** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Markup** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Values** | Pointer to [**[]PropertiesV2Values**](PropertiesV2Values.md) |  | [optional] 
**SelectionValues** | Pointer to [**[]PropertiesV2SelectionValues**](PropertiesV2SelectionValues.md) |  | [optional] 
**Property** | Pointer to [**PropertyV2Property**](PropertyV2Property.md) |  | [optional] 

## Methods

### NewPropertyV2

`func NewPropertyV2() *PropertyV2`

NewPropertyV2 instantiates a new PropertyV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyV2WithDefaults

`func NewPropertyV2WithDefaults() *PropertyV2`

NewPropertyV2WithDefaults instantiates a new PropertyV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PropertyV2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PropertyV2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PropertyV2) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PropertyV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPropertyId

`func (o *PropertyV2) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *PropertyV2) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *PropertyV2) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *PropertyV2) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetGroupId

`func (o *PropertyV2) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *PropertyV2) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *PropertyV2) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *PropertyV2) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetType

`func (o *PropertyV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PropertyV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTargetId

`func (o *PropertyV2) GetTargetId() int32`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PropertyV2) GetTargetIdOk() (*int32, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PropertyV2) SetTargetId(v int32)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *PropertyV2) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### GetValue

`func (o *PropertyV2) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PropertyV2) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PropertyV2) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PropertyV2) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMarkup

`func (o *PropertyV2) GetMarkup() int32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *PropertyV2) GetMarkupOk() (*int32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *PropertyV2) SetMarkup(v int32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *PropertyV2) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PropertyV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PropertyV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PropertyV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PropertyV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PropertyV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PropertyV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PropertyV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PropertyV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValues

`func (o *PropertyV2) GetValues() []PropertiesV2Values`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PropertyV2) GetValuesOk() (*[]PropertiesV2Values, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PropertyV2) SetValues(v []PropertiesV2Values)`

SetValues sets Values field to given value.

### HasValues

`func (o *PropertyV2) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetSelectionValues

`func (o *PropertyV2) GetSelectionValues() []PropertiesV2SelectionValues`

GetSelectionValues returns the SelectionValues field if non-nil, zero value otherwise.

### GetSelectionValuesOk

`func (o *PropertyV2) GetSelectionValuesOk() (*[]PropertiesV2SelectionValues, bool)`

GetSelectionValuesOk returns a tuple with the SelectionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionValues

`func (o *PropertyV2) SetSelectionValues(v []PropertiesV2SelectionValues)`

SetSelectionValues sets SelectionValues field to given value.

### HasSelectionValues

`func (o *PropertyV2) HasSelectionValues() bool`

HasSelectionValues returns a boolean if a field has been set.

### GetProperty

`func (o *PropertyV2) GetProperty() PropertyV2Property`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PropertyV2) GetPropertyOk() (*PropertyV2Property, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PropertyV2) SetProperty(v PropertyV2Property)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *PropertyV2) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


