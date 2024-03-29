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

// PimVariationPropertyValues struct for PimVariationPropertyValues
type PimVariationPropertyValues struct {
	Lang *string `json:"lang,omitempty"`
	Value *string `json:"value,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewPimVariationPropertyValues instantiates a new PimVariationPropertyValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimVariationPropertyValues() *PimVariationPropertyValues {
	this := PimVariationPropertyValues{}
	return &this
}

// NewPimVariationPropertyValuesWithDefaults instantiates a new PimVariationPropertyValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimVariationPropertyValuesWithDefaults() *PimVariationPropertyValues {
	this := PimVariationPropertyValues{}
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *PimVariationPropertyValues) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationPropertyValues) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *PimVariationPropertyValues) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *PimVariationPropertyValues) SetLang(v string) {
	o.Lang = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PimVariationPropertyValues) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationPropertyValues) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PimVariationPropertyValues) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PimVariationPropertyValues) SetValue(v string) {
	o.Value = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PimVariationPropertyValues) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationPropertyValues) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PimVariationPropertyValues) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PimVariationPropertyValues) SetDescription(v string) {
	o.Description = &v
}

func (o PimVariationPropertyValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePimVariationPropertyValues struct {
	value *PimVariationPropertyValues
	isSet bool
}

func (v NullablePimVariationPropertyValues) Get() *PimVariationPropertyValues {
	return v.value
}

func (v *NullablePimVariationPropertyValues) Set(val *PimVariationPropertyValues) {
	v.value = val
	v.isSet = true
}

func (v NullablePimVariationPropertyValues) IsSet() bool {
	return v.isSet
}

func (v *NullablePimVariationPropertyValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimVariationPropertyValues(val *PimVariationPropertyValues) *NullablePimVariationPropertyValues {
	return &NullablePimVariationPropertyValues{value: val, isSet: true}
}

func (v NullablePimVariationPropertyValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePimVariationPropertyValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


