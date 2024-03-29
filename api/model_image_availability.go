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

// ImageAvailability image availability model
type ImageAvailability struct {
	ImageId *int32 `json:"imageId,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewImageAvailability instantiates a new ImageAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageAvailability() *ImageAvailability {
	this := ImageAvailability{}
	return &this
}

// NewImageAvailabilityWithDefaults instantiates a new ImageAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageAvailabilityWithDefaults() *ImageAvailability {
	this := ImageAvailability{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *ImageAvailability) GetImageId() int32 {
	if o == nil || o.ImageId == nil {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAvailability) GetImageIdOk() (*int32, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *ImageAvailability) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *ImageAvailability) SetImageId(v int32) {
	o.ImageId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ImageAvailability) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAvailability) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ImageAvailability) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ImageAvailability) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ImageAvailability) GetValue() float32 {
	if o == nil || o.Value == nil {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageAvailability) GetValueOk() (*float32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ImageAvailability) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *ImageAvailability) SetValue(v float32) {
	o.Value = &v
}

func (o ImageAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageId != nil {
		toSerialize["imageId"] = o.ImageId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableImageAvailability struct {
	value *ImageAvailability
	isSet bool
}

func (v NullableImageAvailability) Get() *ImageAvailability {
	return v.value
}

func (v *NullableImageAvailability) Set(val *ImageAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableImageAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableImageAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageAvailability(val *ImageAvailability) *NullableImageAvailability {
	return &NullableImageAvailability{value: val, isSet: true}
}

func (v NullableImageAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


