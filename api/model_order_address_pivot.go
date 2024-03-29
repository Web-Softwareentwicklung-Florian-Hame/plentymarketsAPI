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

// OrderAddressPivot struct for OrderAddressPivot
type OrderAddressPivot struct {
	OrderId *string `json:"orderId,omitempty"`
	AddressId *int32 `json:"addressId,omitempty"`
	Typeid *int32 `json:"typeid,omitempty"`
}

// NewOrderAddressPivot instantiates a new OrderAddressPivot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderAddressPivot() *OrderAddressPivot {
	this := OrderAddressPivot{}
	return &this
}

// NewOrderAddressPivotWithDefaults instantiates a new OrderAddressPivot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderAddressPivotWithDefaults() *OrderAddressPivot {
	this := OrderAddressPivot{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *OrderAddressPivot) GetOrderId() string {
	if o == nil || o.OrderId == nil {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddressPivot) GetOrderIdOk() (*string, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *OrderAddressPivot) HasOrderId() bool {
	if o != nil && o.OrderId != nil {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *OrderAddressPivot) SetOrderId(v string) {
	o.OrderId = &v
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *OrderAddressPivot) GetAddressId() int32 {
	if o == nil || o.AddressId == nil {
		var ret int32
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddressPivot) GetAddressIdOk() (*int32, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *OrderAddressPivot) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given int32 and assigns it to the AddressId field.
func (o *OrderAddressPivot) SetAddressId(v int32) {
	o.AddressId = &v
}

// GetTypeid returns the Typeid field value if set, zero value otherwise.
func (o *OrderAddressPivot) GetTypeid() int32 {
	if o == nil || o.Typeid == nil {
		var ret int32
		return ret
	}
	return *o.Typeid
}

// GetTypeidOk returns a tuple with the Typeid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderAddressPivot) GetTypeidOk() (*int32, bool) {
	if o == nil || o.Typeid == nil {
		return nil, false
	}
	return o.Typeid, true
}

// HasTypeid returns a boolean if a field has been set.
func (o *OrderAddressPivot) HasTypeid() bool {
	if o != nil && o.Typeid != nil {
		return true
	}

	return false
}

// SetTypeid gets a reference to the given int32 and assigns it to the Typeid field.
func (o *OrderAddressPivot) SetTypeid(v int32) {
	o.Typeid = &v
}

func (o OrderAddressPivot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.OrderId != nil {
		toSerialize["orderId"] = o.OrderId
	}
	if o.AddressId != nil {
		toSerialize["addressId"] = o.AddressId
	}
	if o.Typeid != nil {
		toSerialize["typeid"] = o.Typeid
	}
	return json.Marshal(toSerialize)
}

type NullableOrderAddressPivot struct {
	value *OrderAddressPivot
	isSet bool
}

func (v NullableOrderAddressPivot) Get() *OrderAddressPivot {
	return v.value
}

func (v *NullableOrderAddressPivot) Set(val *OrderAddressPivot) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderAddressPivot) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderAddressPivot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderAddressPivot(val *OrderAddressPivot) *NullableOrderAddressPivot {
	return &NullableOrderAddressPivot{value: val, isSet: true}
}

func (v NullableOrderAddressPivot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderAddressPivot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


