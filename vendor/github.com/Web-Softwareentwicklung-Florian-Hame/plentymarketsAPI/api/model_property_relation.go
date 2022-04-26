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

// PropertyRelation propertyRelation model
type PropertyRelation struct {
	Cast *string `json:"cast,omitempty"`
	TypeIdentifier *string `json:"typeIdentifier,omitempty"`
	Position *int32 `json:"position,omitempty"`
	Id *int32 `json:"id,omitempty"`
	PropertyId *int32 `json:"propertyId,omitempty"`
	PropertyGroupId *int32 `json:"propertyGroupId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewPropertyRelation instantiates a new PropertyRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyRelation() *PropertyRelation {
	this := PropertyRelation{}
	return &this
}

// NewPropertyRelationWithDefaults instantiates a new PropertyRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyRelationWithDefaults() *PropertyRelation {
	this := PropertyRelation{}
	return &this
}

// GetCast returns the Cast field value if set, zero value otherwise.
func (o *PropertyRelation) GetCast() string {
	if o == nil || o.Cast == nil {
		var ret string
		return ret
	}
	return *o.Cast
}

// GetCastOk returns a tuple with the Cast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetCastOk() (*string, bool) {
	if o == nil || o.Cast == nil {
		return nil, false
	}
	return o.Cast, true
}

// HasCast returns a boolean if a field has been set.
func (o *PropertyRelation) HasCast() bool {
	if o != nil && o.Cast != nil {
		return true
	}

	return false
}

// SetCast gets a reference to the given string and assigns it to the Cast field.
func (o *PropertyRelation) SetCast(v string) {
	o.Cast = &v
}

// GetTypeIdentifier returns the TypeIdentifier field value if set, zero value otherwise.
func (o *PropertyRelation) GetTypeIdentifier() string {
	if o == nil || o.TypeIdentifier == nil {
		var ret string
		return ret
	}
	return *o.TypeIdentifier
}

// GetTypeIdentifierOk returns a tuple with the TypeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetTypeIdentifierOk() (*string, bool) {
	if o == nil || o.TypeIdentifier == nil {
		return nil, false
	}
	return o.TypeIdentifier, true
}

// HasTypeIdentifier returns a boolean if a field has been set.
func (o *PropertyRelation) HasTypeIdentifier() bool {
	if o != nil && o.TypeIdentifier != nil {
		return true
	}

	return false
}

// SetTypeIdentifier gets a reference to the given string and assigns it to the TypeIdentifier field.
func (o *PropertyRelation) SetTypeIdentifier(v string) {
	o.TypeIdentifier = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PropertyRelation) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PropertyRelation) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *PropertyRelation) SetPosition(v int32) {
	o.Position = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PropertyRelation) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PropertyRelation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PropertyRelation) SetId(v int32) {
	o.Id = &v
}

// GetPropertyId returns the PropertyId field value if set, zero value otherwise.
func (o *PropertyRelation) GetPropertyId() int32 {
	if o == nil || o.PropertyId == nil {
		var ret int32
		return ret
	}
	return *o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetPropertyIdOk() (*int32, bool) {
	if o == nil || o.PropertyId == nil {
		return nil, false
	}
	return o.PropertyId, true
}

// HasPropertyId returns a boolean if a field has been set.
func (o *PropertyRelation) HasPropertyId() bool {
	if o != nil && o.PropertyId != nil {
		return true
	}

	return false
}

// SetPropertyId gets a reference to the given int32 and assigns it to the PropertyId field.
func (o *PropertyRelation) SetPropertyId(v int32) {
	o.PropertyId = &v
}

// GetPropertyGroupId returns the PropertyGroupId field value if set, zero value otherwise.
func (o *PropertyRelation) GetPropertyGroupId() int32 {
	if o == nil || o.PropertyGroupId == nil {
		var ret int32
		return ret
	}
	return *o.PropertyGroupId
}

// GetPropertyGroupIdOk returns a tuple with the PropertyGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetPropertyGroupIdOk() (*int32, bool) {
	if o == nil || o.PropertyGroupId == nil {
		return nil, false
	}
	return o.PropertyGroupId, true
}

// HasPropertyGroupId returns a boolean if a field has been set.
func (o *PropertyRelation) HasPropertyGroupId() bool {
	if o != nil && o.PropertyGroupId != nil {
		return true
	}

	return false
}

// SetPropertyGroupId gets a reference to the given int32 and assigns it to the PropertyGroupId field.
func (o *PropertyRelation) SetPropertyGroupId(v int32) {
	o.PropertyGroupId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PropertyRelation) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PropertyRelation) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PropertyRelation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PropertyRelation) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyRelation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PropertyRelation) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PropertyRelation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PropertyRelation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cast != nil {
		toSerialize["cast"] = o.Cast
	}
	if o.TypeIdentifier != nil {
		toSerialize["typeIdentifier"] = o.TypeIdentifier
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PropertyId != nil {
		toSerialize["propertyId"] = o.PropertyId
	}
	if o.PropertyGroupId != nil {
		toSerialize["propertyGroupId"] = o.PropertyGroupId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyRelation struct {
	value *PropertyRelation
	isSet bool
}

func (v NullablePropertyRelation) Get() *PropertyRelation {
	return v.value
}

func (v *NullablePropertyRelation) Set(val *PropertyRelation) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyRelation) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyRelation(val *PropertyRelation) *NullablePropertyRelation {
	return &NullablePropertyRelation{value: val, isSet: true}
}

func (v NullablePropertyRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


