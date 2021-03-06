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

// StockCorrectionEntry struct for StockCorrectionEntry
type StockCorrectionEntry struct {
	Variationid *int32 `json:"variationid,omitempty"`
	ReasonId *int32 `json:"reasonId,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	StorageLocationId *int32 `json:"storageLocationId,omitempty"`
}

// NewStockCorrectionEntry instantiates a new StockCorrectionEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStockCorrectionEntry() *StockCorrectionEntry {
	this := StockCorrectionEntry{}
	return &this
}

// NewStockCorrectionEntryWithDefaults instantiates a new StockCorrectionEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStockCorrectionEntryWithDefaults() *StockCorrectionEntry {
	this := StockCorrectionEntry{}
	return &this
}

// GetVariationid returns the Variationid field value if set, zero value otherwise.
func (o *StockCorrectionEntry) GetVariationid() int32 {
	if o == nil || o.Variationid == nil {
		var ret int32
		return ret
	}
	return *o.Variationid
}

// GetVariationidOk returns a tuple with the Variationid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockCorrectionEntry) GetVariationidOk() (*int32, bool) {
	if o == nil || o.Variationid == nil {
		return nil, false
	}
	return o.Variationid, true
}

// HasVariationid returns a boolean if a field has been set.
func (o *StockCorrectionEntry) HasVariationid() bool {
	if o != nil && o.Variationid != nil {
		return true
	}

	return false
}

// SetVariationid gets a reference to the given int32 and assigns it to the Variationid field.
func (o *StockCorrectionEntry) SetVariationid(v int32) {
	o.Variationid = &v
}

// GetReasonId returns the ReasonId field value if set, zero value otherwise.
func (o *StockCorrectionEntry) GetReasonId() int32 {
	if o == nil || o.ReasonId == nil {
		var ret int32
		return ret
	}
	return *o.ReasonId
}

// GetReasonIdOk returns a tuple with the ReasonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockCorrectionEntry) GetReasonIdOk() (*int32, bool) {
	if o == nil || o.ReasonId == nil {
		return nil, false
	}
	return o.ReasonId, true
}

// HasReasonId returns a boolean if a field has been set.
func (o *StockCorrectionEntry) HasReasonId() bool {
	if o != nil && o.ReasonId != nil {
		return true
	}

	return false
}

// SetReasonId gets a reference to the given int32 and assigns it to the ReasonId field.
func (o *StockCorrectionEntry) SetReasonId(v int32) {
	o.ReasonId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *StockCorrectionEntry) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockCorrectionEntry) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *StockCorrectionEntry) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *StockCorrectionEntry) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetStorageLocationId returns the StorageLocationId field value if set, zero value otherwise.
func (o *StockCorrectionEntry) GetStorageLocationId() int32 {
	if o == nil || o.StorageLocationId == nil {
		var ret int32
		return ret
	}
	return *o.StorageLocationId
}

// GetStorageLocationIdOk returns a tuple with the StorageLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StockCorrectionEntry) GetStorageLocationIdOk() (*int32, bool) {
	if o == nil || o.StorageLocationId == nil {
		return nil, false
	}
	return o.StorageLocationId, true
}

// HasStorageLocationId returns a boolean if a field has been set.
func (o *StockCorrectionEntry) HasStorageLocationId() bool {
	if o != nil && o.StorageLocationId != nil {
		return true
	}

	return false
}

// SetStorageLocationId gets a reference to the given int32 and assigns it to the StorageLocationId field.
func (o *StockCorrectionEntry) SetStorageLocationId(v int32) {
	o.StorageLocationId = &v
}

func (o StockCorrectionEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Variationid != nil {
		toSerialize["variationid"] = o.Variationid
	}
	if o.ReasonId != nil {
		toSerialize["reasonId"] = o.ReasonId
	}
	if o.Quantity != nil {
		toSerialize["quantity"] = o.Quantity
	}
	if o.StorageLocationId != nil {
		toSerialize["storageLocationId"] = o.StorageLocationId
	}
	return json.Marshal(toSerialize)
}

type NullableStockCorrectionEntry struct {
	value *StockCorrectionEntry
	isSet bool
}

func (v NullableStockCorrectionEntry) Get() *StockCorrectionEntry {
	return v.value
}

func (v *NullableStockCorrectionEntry) Set(val *StockCorrectionEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableStockCorrectionEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableStockCorrectionEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStockCorrectionEntry(val *StockCorrectionEntry) *NullableStockCorrectionEntry {
	return &NullableStockCorrectionEntry{value: val, isSet: true}
}

func (v NullableStockCorrectionEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStockCorrectionEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


