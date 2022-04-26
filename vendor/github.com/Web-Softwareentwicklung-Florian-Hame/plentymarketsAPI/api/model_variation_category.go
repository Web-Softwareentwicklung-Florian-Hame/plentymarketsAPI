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

// VariationCategory variation category model
type VariationCategory struct {
	VariationId *int32 `json:"variationId,omitempty"`
	CategoryId *int32 `json:"categoryId,omitempty"`
	Position *int32 `json:"position,omitempty"`
	IsNeckermannPrimary *bool `json:"isNeckermannPrimary,omitempty"`
}

// NewVariationCategory instantiates a new VariationCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationCategory() *VariationCategory {
	this := VariationCategory{}
	return &this
}

// NewVariationCategoryWithDefaults instantiates a new VariationCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationCategoryWithDefaults() *VariationCategory {
	this := VariationCategory{}
	return &this
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *VariationCategory) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationCategory) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *VariationCategory) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *VariationCategory) SetVariationId(v int32) {
	o.VariationId = &v
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *VariationCategory) GetCategoryId() int32 {
	if o == nil || o.CategoryId == nil {
		var ret int32
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationCategory) GetCategoryIdOk() (*int32, bool) {
	if o == nil || o.CategoryId == nil {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *VariationCategory) HasCategoryId() bool {
	if o != nil && o.CategoryId != nil {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int32 and assigns it to the CategoryId field.
func (o *VariationCategory) SetCategoryId(v int32) {
	o.CategoryId = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *VariationCategory) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationCategory) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *VariationCategory) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *VariationCategory) SetPosition(v int32) {
	o.Position = &v
}

// GetIsNeckermannPrimary returns the IsNeckermannPrimary field value if set, zero value otherwise.
func (o *VariationCategory) GetIsNeckermannPrimary() bool {
	if o == nil || o.IsNeckermannPrimary == nil {
		var ret bool
		return ret
	}
	return *o.IsNeckermannPrimary
}

// GetIsNeckermannPrimaryOk returns a tuple with the IsNeckermannPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationCategory) GetIsNeckermannPrimaryOk() (*bool, bool) {
	if o == nil || o.IsNeckermannPrimary == nil {
		return nil, false
	}
	return o.IsNeckermannPrimary, true
}

// HasIsNeckermannPrimary returns a boolean if a field has been set.
func (o *VariationCategory) HasIsNeckermannPrimary() bool {
	if o != nil && o.IsNeckermannPrimary != nil {
		return true
	}

	return false
}

// SetIsNeckermannPrimary gets a reference to the given bool and assigns it to the IsNeckermannPrimary field.
func (o *VariationCategory) SetIsNeckermannPrimary(v bool) {
	o.IsNeckermannPrimary = &v
}

func (o VariationCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	if o.CategoryId != nil {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.IsNeckermannPrimary != nil {
		toSerialize["isNeckermannPrimary"] = o.IsNeckermannPrimary
	}
	return json.Marshal(toSerialize)
}

type NullableVariationCategory struct {
	value *VariationCategory
	isSet bool
}

func (v NullableVariationCategory) Get() *VariationCategory {
	return v.value
}

func (v *NullableVariationCategory) Set(val *VariationCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationCategory(val *VariationCategory) *NullableVariationCategory {
	return &NullableVariationCategory{value: val, isSet: true}
}

func (v NullableVariationCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

