# PimVariationProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to **interface{}** | can have any type | [optional] 
**SelectionValues** | Pointer to [**[]PimVariationPropertySelectionValues**](PimVariationPropertySelectionValues.md) |  | [optional] 

## Methods

### NewPimVariationProperty

`func NewPimVariationProperty() *PimVariationProperty`

NewPimVariationProperty instantiates a new PimVariationProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimVariationPropertyWithDefaults

`func NewPimVariationPropertyWithDefaults() *PimVariationProperty`

NewPimVariationPropertyWithDefaults instantiates a new PimVariationProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *PimVariationProperty) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *PimVariationProperty) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *PimVariationProperty) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *PimVariationProperty) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetValue

`func (o *PimVariationProperty) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PimVariationProperty) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PimVariationProperty) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PimVariationProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PimVariationProperty) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PimVariationProperty) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetSelectionValues

`func (o *PimVariationProperty) GetSelectionValues() []PimVariationPropertySelectionValues`

GetSelectionValues returns the SelectionValues field if non-nil, zero value otherwise.

### GetSelectionValuesOk

`func (o *PimVariationProperty) GetSelectionValuesOk() (*[]PimVariationPropertySelectionValues, bool)`

GetSelectionValuesOk returns a tuple with the SelectionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionValues

`func (o *PimVariationProperty) SetSelectionValues(v []PimVariationPropertySelectionValues)`

SetSelectionValues sets SelectionValues field to given value.

### HasSelectionValues

`func (o *PimVariationProperty) HasSelectionValues() bool`

HasSelectionValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


