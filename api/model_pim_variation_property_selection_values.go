/*
Plentymarkets-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// PimVariationPropertySelectionValues struct for PimVariationPropertySelectionValues
type PimVariationPropertySelectionValues struct {
	SelectionId *int32 `json:"selectionId,omitempty"`
}

// NewPimVariationPropertySelectionValues instantiates a new PimVariationPropertySelectionValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimVariationPropertySelectionValues() *PimVariationPropertySelectionValues {
	this := PimVariationPropertySelectionValues{}
	return &this
}

// NewPimVariationPropertySelectionValuesWithDefaults instantiates a new PimVariationPropertySelectionValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimVariationPropertySelectionValuesWithDefaults() *PimVariationPropertySelectionValues {
	this := PimVariationPropertySelectionValues{}
	return &this
}

// GetSelectionId returns the SelectionId field value if set, zero value otherwise.
func (o *PimVariationPropertySelectionValues) GetSelectionId() int32 {
	if o == nil || o.SelectionId == nil {
		var ret int32
		return ret
	}
	return *o.SelectionId
}

// GetSelectionIdOk returns a tuple with the SelectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationPropertySelectionValues) GetSelectionIdOk() (*int32, bool) {
	if o == nil || o.SelectionId == nil {
		return nil, false
	}
	return o.SelectionId, true
}

// HasSelectionId returns a boolean if a field has been set.
func (o *PimVariationPropertySelectionValues) HasSelectionId() bool {
	if o != nil && o.SelectionId != nil {
		return true
	}

	return false
}

// SetSelectionId gets a reference to the given int32 and assigns it to the SelectionId field.
func (o *PimVariationPropertySelectionValues) SetSelectionId(v int32) {
	o.SelectionId = &v
}

func (o PimVariationPropertySelectionValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SelectionId != nil {
		toSerialize["selectionId"] = o.SelectionId
	}
	return json.Marshal(toSerialize)
}

type NullablePimVariationPropertySelectionValues struct {
	value *PimVariationPropertySelectionValues
	isSet bool
}

func (v NullablePimVariationPropertySelectionValues) Get() *PimVariationPropertySelectionValues {
	return v.value
}

func (v *NullablePimVariationPropertySelectionValues) Set(val *PimVariationPropertySelectionValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePimVariationPropertySelectionValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePimVariationPropertySelectionValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimVariationPropertySelectionValues(val *PimVariationPropertySelectionValues) *NullablePimVariationPropertySelectionValues {
	return &NullablePimVariationPropertySelectionValues{value: val, isSet: true}
}

func (v NullablePimVariationPropertySelectionValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePimVariationPropertySelectionValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

