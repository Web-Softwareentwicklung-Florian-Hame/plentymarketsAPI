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

// ItemEbayTitle struct for ItemEbayTitle
type ItemEbayTitle struct {
	Title *string `json:"title,omitempty"`
	ItemId *int32 `json:"itemId,omitempty"`
	Id *int32 `json:"id,omitempty"`
}

// NewItemEbayTitle instantiates a new ItemEbayTitle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemEbayTitle() *ItemEbayTitle {
	this := ItemEbayTitle{}
	return &this
}

// NewItemEbayTitleWithDefaults instantiates a new ItemEbayTitle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemEbayTitleWithDefaults() *ItemEbayTitle {
	this := ItemEbayTitle{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ItemEbayTitle) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemEbayTitle) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ItemEbayTitle) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ItemEbayTitle) SetTitle(v string) {
	o.Title = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *ItemEbayTitle) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemEbayTitle) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *ItemEbayTitle) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *ItemEbayTitle) SetItemId(v int32) {
	o.ItemId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ItemEbayTitle) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemEbayTitle) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ItemEbayTitle) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ItemEbayTitle) SetId(v int32) {
	o.Id = &v
}

func (o ItemEbayTitle) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ItemId != nil {
		toSerialize["itemId"] = o.ItemId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableItemEbayTitle struct {
	value *ItemEbayTitle
	isSet bool
}

func (v NullableItemEbayTitle) Get() *ItemEbayTitle {
	return v.value
}

func (v *NullableItemEbayTitle) Set(val *ItemEbayTitle) {
	v.value = val
	v.isSet = true
}

func (v NullableItemEbayTitle) IsSet() bool {
	return v.isSet
}

func (v *NullableItemEbayTitle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemEbayTitle(val *ItemEbayTitle) *NullableItemEbayTitle {
	return &NullableItemEbayTitle{value: val, isSet: true}
}

func (v NullableItemEbayTitle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemEbayTitle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


