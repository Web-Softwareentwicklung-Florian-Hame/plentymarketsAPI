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

// AttributeName struct for AttributeName
type AttributeName struct {
	Lang *string `json:"lang,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewAttributeName instantiates a new AttributeName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeName() *AttributeName {
	this := AttributeName{}
	return &this
}

// NewAttributeNameWithDefaults instantiates a new AttributeName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeNameWithDefaults() *AttributeName {
	this := AttributeName{}
	return &this
}

// GetLang returns the Lang field value if set, zero value otherwise.
func (o *AttributeName) GetLang() string {
	if o == nil || o.Lang == nil {
		var ret string
		return ret
	}
	return *o.Lang
}

// GetLangOk returns a tuple with the Lang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeName) GetLangOk() (*string, bool) {
	if o == nil || o.Lang == nil {
		return nil, false
	}
	return o.Lang, true
}

// HasLang returns a boolean if a field has been set.
func (o *AttributeName) HasLang() bool {
	if o != nil && o.Lang != nil {
		return true
	}

	return false
}

// SetLang gets a reference to the given string and assigns it to the Lang field.
func (o *AttributeName) SetLang(v string) {
	o.Lang = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeName) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeName) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeName) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeName) SetName(v string) {
	o.Name = &v
}

func (o AttributeName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lang != nil {
		toSerialize["lang"] = o.Lang
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAttributeName struct {
	value *AttributeName
	isSet bool
}

func (v NullableAttributeName) Get() *AttributeName {
	return v.value
}

func (v *NullableAttributeName) Set(val *AttributeName) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeName) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeName(val *AttributeName) *NullableAttributeName {
	return &NullableAttributeName{value: val, isSet: true}
}

func (v NullableAttributeName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


