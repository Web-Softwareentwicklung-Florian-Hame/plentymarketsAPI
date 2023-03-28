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

// OrderDocument struct for OrderDocument
type OrderDocument struct {
	Id *int32 `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	Number *string `json:"number,omitempty"`
	NumberWithPrefix *string `json:"numberWithPrefix,omitempty"`
	Path *string `json:"path,omitempty"`
	UserId *string `json:"userId,omitempty"`
	Source *string `json:"source,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	DisplayDate *time.Time `json:"displayDate,omitempty"`
	NewOrderArchitecture *string `json:"newOrderArchitecture,omitempty"`
	Status *string `json:"status,omitempty"`
	Pivot *OrderDocumentPivot `json:"pivot,omitempty"`
}

// NewOrderDocument instantiates a new OrderDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDocument() *OrderDocument {
	this := OrderDocument{}
	return &this
}

// NewOrderDocumentWithDefaults instantiates a new OrderDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDocumentWithDefaults() *OrderDocument {
	this := OrderDocument{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrderDocument) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrderDocument) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrderDocument) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderDocument) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderDocument) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderDocument) SetType(v string) {
	o.Type = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *OrderDocument) GetNumber() string {
	if o == nil || o.Number == nil {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetNumberOk() (*string, bool) {
	if o == nil || o.Number == nil {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *OrderDocument) HasNumber() bool {
	if o != nil && o.Number != nil {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *OrderDocument) SetNumber(v string) {
	o.Number = &v
}

// GetNumberWithPrefix returns the NumberWithPrefix field value if set, zero value otherwise.
func (o *OrderDocument) GetNumberWithPrefix() string {
	if o == nil || o.NumberWithPrefix == nil {
		var ret string
		return ret
	}
	return *o.NumberWithPrefix
}

// GetNumberWithPrefixOk returns a tuple with the NumberWithPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetNumberWithPrefixOk() (*string, bool) {
	if o == nil || o.NumberWithPrefix == nil {
		return nil, false
	}
	return o.NumberWithPrefix, true
}

// HasNumberWithPrefix returns a boolean if a field has been set.
func (o *OrderDocument) HasNumberWithPrefix() bool {
	if o != nil && o.NumberWithPrefix != nil {
		return true
	}

	return false
}

// SetNumberWithPrefix gets a reference to the given string and assigns it to the NumberWithPrefix field.
func (o *OrderDocument) SetNumberWithPrefix(v string) {
	o.NumberWithPrefix = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *OrderDocument) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *OrderDocument) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *OrderDocument) SetPath(v string) {
	o.Path = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OrderDocument) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OrderDocument) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *OrderDocument) SetUserId(v string) {
	o.UserId = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *OrderDocument) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *OrderDocument) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *OrderDocument) SetSource(v string) {
	o.Source = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OrderDocument) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OrderDocument) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OrderDocument) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *OrderDocument) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *OrderDocument) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *OrderDocument) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDisplayDate returns the DisplayDate field value if set, zero value otherwise.
func (o *OrderDocument) GetDisplayDate() time.Time {
	if o == nil || o.DisplayDate == nil {
		var ret time.Time
		return ret
	}
	return *o.DisplayDate
}

// GetDisplayDateOk returns a tuple with the DisplayDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetDisplayDateOk() (*time.Time, bool) {
	if o == nil || o.DisplayDate == nil {
		return nil, false
	}
	return o.DisplayDate, true
}

// HasDisplayDate returns a boolean if a field has been set.
func (o *OrderDocument) HasDisplayDate() bool {
	if o != nil && o.DisplayDate != nil {
		return true
	}

	return false
}

// SetDisplayDate gets a reference to the given time.Time and assigns it to the DisplayDate field.
func (o *OrderDocument) SetDisplayDate(v time.Time) {
	o.DisplayDate = &v
}

// GetNewOrderArchitecture returns the NewOrderArchitecture field value if set, zero value otherwise.
func (o *OrderDocument) GetNewOrderArchitecture() string {
	if o == nil || o.NewOrderArchitecture == nil {
		var ret string
		return ret
	}
	return *o.NewOrderArchitecture
}

// GetNewOrderArchitectureOk returns a tuple with the NewOrderArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetNewOrderArchitectureOk() (*string, bool) {
	if o == nil || o.NewOrderArchitecture == nil {
		return nil, false
	}
	return o.NewOrderArchitecture, true
}

// HasNewOrderArchitecture returns a boolean if a field has been set.
func (o *OrderDocument) HasNewOrderArchitecture() bool {
	if o != nil && o.NewOrderArchitecture != nil {
		return true
	}

	return false
}

// SetNewOrderArchitecture gets a reference to the given string and assigns it to the NewOrderArchitecture field.
func (o *OrderDocument) SetNewOrderArchitecture(v string) {
	o.NewOrderArchitecture = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrderDocument) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrderDocument) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrderDocument) SetStatus(v string) {
	o.Status = &v
}

// GetPivot returns the Pivot field value if set, zero value otherwise.
func (o *OrderDocument) GetPivot() OrderDocumentPivot {
	if o == nil || o.Pivot == nil {
		var ret OrderDocumentPivot
		return ret
	}
	return *o.Pivot
}

// GetPivotOk returns a tuple with the Pivot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDocument) GetPivotOk() (*OrderDocumentPivot, bool) {
	if o == nil || o.Pivot == nil {
		return nil, false
	}
	return o.Pivot, true
}

// HasPivot returns a boolean if a field has been set.
func (o *OrderDocument) HasPivot() bool {
	if o != nil && o.Pivot != nil {
		return true
	}

	return false
}

// SetPivot gets a reference to the given OrderDocumentPivot and assigns it to the Pivot field.
func (o *OrderDocument) SetPivot(v OrderDocumentPivot) {
	o.Pivot = &v
}

func (o OrderDocument) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Number != nil {
		toSerialize["number"] = o.Number
	}
	if o.NumberWithPrefix != nil {
		toSerialize["numberWithPrefix"] = o.NumberWithPrefix
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.DisplayDate != nil {
		toSerialize["displayDate"] = o.DisplayDate
	}
	if o.NewOrderArchitecture != nil {
		toSerialize["newOrderArchitecture"] = o.NewOrderArchitecture
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Pivot != nil {
		toSerialize["pivot"] = o.Pivot
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDocument struct {
	value *OrderDocument
	isSet bool
}

func (v NullableOrderDocument) Get() *OrderDocument {
	return v.value
}

func (v *NullableOrderDocument) Set(val *OrderDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDocument(val *OrderDocument) *NullableOrderDocument {
	return &NullableOrderDocument{value: val, isSet: true}
}

func (v NullableOrderDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


