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

// PimVariationComment struct for PimVariationComment
type PimVariationComment struct {
	Id *int32 `json:"id,omitempty"`
	UserId *int32 `json:"userId,omitempty"`
	Text *string `json:"text,omitempty"`
	IsVisibleForContact *bool `json:"isVisibleForContact,omitempty"`
}

// NewPimVariationComment instantiates a new PimVariationComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimVariationComment() *PimVariationComment {
	this := PimVariationComment{}
	return &this
}

// NewPimVariationCommentWithDefaults instantiates a new PimVariationComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimVariationCommentWithDefaults() *PimVariationComment {
	this := PimVariationComment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PimVariationComment) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationComment) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PimVariationComment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PimVariationComment) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PimVariationComment) GetUserId() int32 {
	if o == nil || o.UserId == nil {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationComment) GetUserIdOk() (*int32, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PimVariationComment) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *PimVariationComment) SetUserId(v int32) {
	o.UserId = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PimVariationComment) GetText() string {
	if o == nil || o.Text == nil {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationComment) GetTextOk() (*string, bool) {
	if o == nil || o.Text == nil {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PimVariationComment) HasText() bool {
	if o != nil && o.Text != nil {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PimVariationComment) SetText(v string) {
	o.Text = &v
}

// GetIsVisibleForContact returns the IsVisibleForContact field value if set, zero value otherwise.
func (o *PimVariationComment) GetIsVisibleForContact() bool {
	if o == nil || o.IsVisibleForContact == nil {
		var ret bool
		return ret
	}
	return *o.IsVisibleForContact
}

// GetIsVisibleForContactOk returns a tuple with the IsVisibleForContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationComment) GetIsVisibleForContactOk() (*bool, bool) {
	if o == nil || o.IsVisibleForContact == nil {
		return nil, false
	}
	return o.IsVisibleForContact, true
}

// HasIsVisibleForContact returns a boolean if a field has been set.
func (o *PimVariationComment) HasIsVisibleForContact() bool {
	if o != nil && o.IsVisibleForContact != nil {
		return true
	}

	return false
}

// SetIsVisibleForContact gets a reference to the given bool and assigns it to the IsVisibleForContact field.
func (o *PimVariationComment) SetIsVisibleForContact(v bool) {
	o.IsVisibleForContact = &v
}

func (o PimVariationComment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Text != nil {
		toSerialize["text"] = o.Text
	}
	if o.IsVisibleForContact != nil {
		toSerialize["isVisibleForContact"] = o.IsVisibleForContact
	}
	return json.Marshal(toSerialize)
}

type NullablePimVariationComment struct {
	value *PimVariationComment
	isSet bool
}

func (v NullablePimVariationComment) Get() *PimVariationComment {
	return v.value
}

func (v *NullablePimVariationComment) Set(val *PimVariationComment) {
	v.value = val
	v.isSet = true
}

func (v NullablePimVariationComment) IsSet() bool {
	return v.isSet
}

func (v *NullablePimVariationComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimVariationComment(val *PimVariationComment) *NullablePimVariationComment {
	return &NullablePimVariationComment{value: val, isSet: true}
}

func (v NullablePimVariationComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePimVariationComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

