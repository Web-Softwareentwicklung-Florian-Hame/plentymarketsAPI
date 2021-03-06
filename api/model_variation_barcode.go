/*
Plentymarkets-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// VariationBarcode Variation barcode model
type VariationBarcode struct {
	BarcodeId *int32 `json:"barcodeId,omitempty"`
	VariationId *int32 `json:"variationId,omitempty"`
	Code *string `json:"code,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewVariationBarcode instantiates a new VariationBarcode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationBarcode() *VariationBarcode {
	this := VariationBarcode{}
	return &this
}

// NewVariationBarcodeWithDefaults instantiates a new VariationBarcode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationBarcodeWithDefaults() *VariationBarcode {
	this := VariationBarcode{}
	return &this
}

// GetBarcodeId returns the BarcodeId field value if set, zero value otherwise.
func (o *VariationBarcode) GetBarcodeId() int32 {
	if o == nil || o.BarcodeId == nil {
		var ret int32
		return ret
	}
	return *o.BarcodeId
}

// GetBarcodeIdOk returns a tuple with the BarcodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationBarcode) GetBarcodeIdOk() (*int32, bool) {
	if o == nil || o.BarcodeId == nil {
		return nil, false
	}
	return o.BarcodeId, true
}

// HasBarcodeId returns a boolean if a field has been set.
func (o *VariationBarcode) HasBarcodeId() bool {
	if o != nil && o.BarcodeId != nil {
		return true
	}

	return false
}

// SetBarcodeId gets a reference to the given int32 and assigns it to the BarcodeId field.
func (o *VariationBarcode) SetBarcodeId(v int32) {
	o.BarcodeId = &v
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *VariationBarcode) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationBarcode) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *VariationBarcode) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *VariationBarcode) SetVariationId(v int32) {
	o.VariationId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *VariationBarcode) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationBarcode) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *VariationBarcode) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *VariationBarcode) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VariationBarcode) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationBarcode) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VariationBarcode) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VariationBarcode) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VariationBarcode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BarcodeId != nil {
		toSerialize["barcodeId"] = o.BarcodeId
	}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableVariationBarcode struct {
	value *VariationBarcode
	isSet bool
}

func (v NullableVariationBarcode) Get() *VariationBarcode {
	return v.value
}

func (v *NullableVariationBarcode) Set(val *VariationBarcode) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationBarcode) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationBarcode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationBarcode(val *VariationBarcode) *NullableVariationBarcode {
	return &NullableVariationBarcode{value: val, isSet: true}
}

func (v NullableVariationBarcode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationBarcode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


