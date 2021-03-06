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

// VariationMarket variation market link model
type VariationMarket struct {
	VariationId *int32 `json:"variationId,omitempty"`
	MarketId *float32 `json:"marketId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewVariationMarket instantiates a new VariationMarket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationMarket() *VariationMarket {
	this := VariationMarket{}
	return &this
}

// NewVariationMarketWithDefaults instantiates a new VariationMarket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationMarketWithDefaults() *VariationMarket {
	this := VariationMarket{}
	return &this
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *VariationMarket) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationMarket) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *VariationMarket) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *VariationMarket) SetVariationId(v int32) {
	o.VariationId = &v
}

// GetMarketId returns the MarketId field value if set, zero value otherwise.
func (o *VariationMarket) GetMarketId() float32 {
	if o == nil || o.MarketId == nil {
		var ret float32
		return ret
	}
	return *o.MarketId
}

// GetMarketIdOk returns a tuple with the MarketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationMarket) GetMarketIdOk() (*float32, bool) {
	if o == nil || o.MarketId == nil {
		return nil, false
	}
	return o.MarketId, true
}

// HasMarketId returns a boolean if a field has been set.
func (o *VariationMarket) HasMarketId() bool {
	if o != nil && o.MarketId != nil {
		return true
	}

	return false
}

// SetMarketId gets a reference to the given float32 and assigns it to the MarketId field.
func (o *VariationMarket) SetMarketId(v float32) {
	o.MarketId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VariationMarket) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationMarket) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VariationMarket) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *VariationMarket) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o VariationMarket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	if o.MarketId != nil {
		toSerialize["marketId"] = o.MarketId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableVariationMarket struct {
	value *VariationMarket
	isSet bool
}

func (v NullableVariationMarket) Get() *VariationMarket {
	return v.value
}

func (v *NullableVariationMarket) Set(val *VariationMarket) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationMarket) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationMarket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationMarket(val *VariationMarket) *NullableVariationMarket {
	return &NullableVariationMarket{value: val, isSet: true}
}

func (v NullableVariationMarket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationMarket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


