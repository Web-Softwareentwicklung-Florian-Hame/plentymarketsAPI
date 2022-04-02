# VariationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**PropertyId** | Pointer to **int32** |  | [optional] 
**PropertySelectionId** | Pointer to **int32** |  | [optional] 
**ValueInt** | Pointer to **int32** |  | [optional] 
**ValueFloat** | Pointer to **float32** |  | [optional] 
**Surcharge** | Pointer to **float32** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 
**Names** | Pointer to [**[]VariationPropertyName**](VariationPropertyName.md) |  | [optional] 
**Property** | Pointer to [**VariationPropertyProperty**](VariationPropertyProperty.md) |  | [optional] 

## Methods

### NewVariationProperty

`func NewVariationProperty() *VariationProperty`

NewVariationProperty instantiates a new VariationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationPropertyWithDefaults

`func NewVariationPropertyWithDefaults() *VariationProperty`

NewVariationPropertyWithDefaults instantiates a new VariationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationProperty) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationProperty) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationProperty) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *VariationProperty) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VariationProperty) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VariationProperty) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VariationProperty) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPropertyId

`func (o *VariationProperty) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *VariationProperty) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *VariationProperty) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *VariationProperty) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetPropertySelectionId

`func (o *VariationProperty) GetPropertySelectionId() int32`

GetPropertySelectionId returns the PropertySelectionId field if non-nil, zero value otherwise.

### GetPropertySelectionIdOk

`func (o *VariationProperty) GetPropertySelectionIdOk() (*int32, bool)`

GetPropertySelectionIdOk returns a tuple with the PropertySelectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelectionId

`func (o *VariationProperty) SetPropertySelectionId(v int32)`

SetPropertySelectionId sets PropertySelectionId field to given value.

### HasPropertySelectionId

`func (o *VariationProperty) HasPropertySelectionId() bool`

HasPropertySelectionId returns a boolean if a field has been set.

### GetValueInt

`func (o *VariationProperty) GetValueInt() int32`

GetValueInt returns the ValueInt field if non-nil, zero value otherwise.

### GetValueIntOk

`func (o *VariationProperty) GetValueIntOk() (*int32, bool)`

GetValueIntOk returns a tuple with the ValueInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInt

`func (o *VariationProperty) SetValueInt(v int32)`

SetValueInt sets ValueInt field to given value.

### HasValueInt

`func (o *VariationProperty) HasValueInt() bool`

HasValueInt returns a boolean if a field has been set.

### GetValueFloat

`func (o *VariationProperty) GetValueFloat() float32`

GetValueFloat returns the ValueFloat field if non-nil, zero value otherwise.

### GetValueFloatOk

`func (o *VariationProperty) GetValueFloatOk() (*float32, bool)`

GetValueFloatOk returns a tuple with the ValueFloat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFloat

`func (o *VariationProperty) SetValueFloat(v float32)`

SetValueFloat sets ValueFloat field to given value.

### HasValueFloat

`func (o *VariationProperty) HasValueFloat() bool`

HasValueFloat returns a boolean if a field has been set.

### GetSurcharge

`func (o *VariationProperty) GetSurcharge() float32`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *VariationProperty) GetSurchargeOk() (*float32, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *VariationProperty) SetSurcharge(v float32)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *VariationProperty) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VariationProperty) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariationProperty) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariationProperty) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VariationProperty) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariationProperty) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariationProperty) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariationProperty) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VariationProperty) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetVariationId

`func (o *VariationProperty) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *VariationProperty) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *VariationProperty) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *VariationProperty) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetNames

`func (o *VariationProperty) GetNames() []VariationPropertyName`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VariationProperty) GetNamesOk() (*[]VariationPropertyName, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VariationProperty) SetNames(v []VariationPropertyName)`

SetNames sets Names field to given value.

### HasNames

`func (o *VariationProperty) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetProperty

`func (o *VariationProperty) GetProperty() VariationPropertyProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *VariationProperty) GetPropertyOk() (*VariationPropertyProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *VariationProperty) SetProperty(v VariationPropertyProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *VariationProperty) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


