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

// ItemShippingProfileResponse struct for ItemShippingProfileResponse
type ItemShippingProfileResponse struct {
	Id *int32 `json:"id,omitempty"`
	ItemId *int32 `json:"itemId,omitempty"`
	ProfileId *int32 `json:"profileId,omitempty"`
}

// NewItemShippingProfileResponse instantiates a new ItemShippingProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemShippingProfileResponse() *ItemShippingProfileResponse {
	this := ItemShippingProfileResponse{}
	return &this
}

// NewItemShippingProfileResponseWithDefaults instantiates a new ItemShippingProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemShippingProfileResponseWithDefaults() *ItemShippingProfileResponse {
	this := ItemShippingProfileResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ItemShippingProfileResponse) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingProfileResponse) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ItemShippingProfileResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ItemShippingProfileResponse) SetId(v int32) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ItemShippingProfileResponse) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingProfileResponse) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ItemShippingProfileResponse) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *ItemShippingProfileResponse) SetItemId(v int32) {
	o.ItemId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ItemShippingProfileResponse) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemShippingProfileResponse) GetProfileIdOk() (*int32, bool) {
	if o == nil || o.ProfileId == nil {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ItemShippingProfileResponse) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *ItemShippingProfileResponse) SetProfileId(v int32) {
	o.ProfileId = &v
}

func (o ItemShippingProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemId != nil {
		toSerialize["itemId"] = o.ItemId
	}
	if o.ProfileId != nil {
		toSerialize["profileId"] = o.ProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableItemShippingProfileResponse struct {
	value *ItemShippingProfileResponse
	isSet bool
}

func (v NullableItemShippingProfileResponse) Get() *ItemShippingProfileResponse {
	return v.value
}

func (v *NullableItemShippingProfileResponse) Set(val *ItemShippingProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemShippingProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemShippingProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemShippingProfileResponse(val *ItemShippingProfileResponse) *NullableItemShippingProfileResponse {
	return &NullableItemShippingProfileResponse{value: val, isSet: true}
}

func (v NullableItemShippingProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemShippingProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


