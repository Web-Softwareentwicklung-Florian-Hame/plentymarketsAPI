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

// MultipleVariationsUpdatedResponse format of the response when multiple variations gots updated
type MultipleVariationsUpdatedResponse struct {
	Success *map[string]Variation `json:"success,omitempty"`
	Failed *map[string][]string `json:"failed,omitempty"`
}

// NewMultipleVariationsUpdatedResponse instantiates a new MultipleVariationsUpdatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleVariationsUpdatedResponse() *MultipleVariationsUpdatedResponse {
	this := MultipleVariationsUpdatedResponse{}
	return &this
}

// NewMultipleVariationsUpdatedResponseWithDefaults instantiates a new MultipleVariationsUpdatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleVariationsUpdatedResponseWithDefaults() *MultipleVariationsUpdatedResponse {
	this := MultipleVariationsUpdatedResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MultipleVariationsUpdatedResponse) GetSuccess() map[string]Variation {
	if o == nil || o.Success == nil {
		var ret map[string]Variation
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleVariationsUpdatedResponse) GetSuccessOk() (*map[string]Variation, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MultipleVariationsUpdatedResponse) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given map[string]Variation and assigns it to the Success field.
func (o *MultipleVariationsUpdatedResponse) SetSuccess(v map[string]Variation) {
	o.Success = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *MultipleVariationsUpdatedResponse) GetFailed() map[string][]string {
	if o == nil || o.Failed == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleVariationsUpdatedResponse) GetFailedOk() (*map[string][]string, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *MultipleVariationsUpdatedResponse) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given map[string][]string and assigns it to the Failed field.
func (o *MultipleVariationsUpdatedResponse) SetFailed(v map[string][]string) {
	o.Failed = &v
}

func (o MultipleVariationsUpdatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	return json.Marshal(toSerialize)
}

type NullableMultipleVariationsUpdatedResponse struct {
	value *MultipleVariationsUpdatedResponse
	isSet bool
}

func (v NullableMultipleVariationsUpdatedResponse) Get() *MultipleVariationsUpdatedResponse {
	return v.value
}

func (v *NullableMultipleVariationsUpdatedResponse) Set(val *MultipleVariationsUpdatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleVariationsUpdatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleVariationsUpdatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleVariationsUpdatedResponse(val *MultipleVariationsUpdatedResponse) *NullableMultipleVariationsUpdatedResponse {
	return &NullableMultipleVariationsUpdatedResponse{value: val, isSet: true}
}

func (v NullableMultipleVariationsUpdatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleVariationsUpdatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


