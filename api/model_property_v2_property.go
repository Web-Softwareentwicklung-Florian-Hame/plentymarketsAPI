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

// PropertyV2Property struct for PropertyV2Property
type PropertyV2Property struct {
	Id *int32 `json:"id,omitempty"`
	Cast *string `json:"cast,omitempty"`
	Type *string `json:"type,omitempty"`
	Position *int32 `json:"position,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewPropertyV2Property instantiates a new PropertyV2Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyV2Property() *PropertyV2Property {
	this := PropertyV2Property{}
	return &this
}

// NewPropertyV2PropertyWithDefaults instantiates a new PropertyV2Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyV2PropertyWithDefaults() *PropertyV2Property {
	this := PropertyV2Property{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PropertyV2Property) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PropertyV2Property) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PropertyV2Property) SetId(v int32) {
	o.Id = &v
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *PropertyV2Property) GetCast() string {
	if o == nil || o.Cast == nil {
		var ret string
		return ret
	}
	return *o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetCastOk() (*string, bool) {
	if o == nil || o.Cast == nil {
		return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *PropertyV2Property) HasCast() bool {
	if o != nil && o.Cast != nil {
		return true
	}

	return false
}

// SetCast gets a reference to the given string and assigns it to the Cast field.
func (o *PropertyV2Property) SetCast(v string) {
	o.Cast = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PropertyV2Property) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PropertyV2Property) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PropertyV2Property) SetType(v string) {
	o.Type = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PropertyV2Property) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PropertyV2Property) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *PropertyV2Property) SetPosition(v int32) {
	o.Position = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PropertyV2Property) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PropertyV2Property) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PropertyV2Property) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PropertyV2Property) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyV2Property) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PropertyV2Property) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PropertyV2Property) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PropertyV2Property) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Cast != nil {
		toSerialize["cast"] = o.Cast
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyV2Property struct {
	value *PropertyV2Property
	isSet bool
}

func (v NullablePropertyV2Property) Get() *PropertyV2Property {
	return v.value
}

func (v *NullablePropertyV2Property) Set(val *PropertyV2Property) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyV2Property) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyV2Property) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyV2Property(val *PropertyV2Property) *NullablePropertyV2Property {
	return &NullablePropertyV2Property{value: val, isSet: true}
}

func (v NullablePropertyV2Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyV2Property) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

