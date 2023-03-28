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

// WarehouseContactAddress struct for WarehouseContactAddress
type WarehouseContactAddress struct {
	Name1 *string `json:"name1,omitempty"`
	Address1 *string `json:"address1,omitempty"`
	Address2 *string `json:"address2,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
	Town *string `json:"town,omitempty"`
	CountryId *int32 `json:"countryId,omitempty"`
	Options *[]WarehouseContactAddressOptions `json:"options,omitempty"`
	Title *string `json:"title,omitempty"`
	ContactPerson *string `json:"contactPerson,omitempty"`
	TaxIdNumber *string `json:"taxIdNumber,omitempty"`
}

// NewWarehouseContactAddress instantiates a new WarehouseContactAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWarehouseContactAddress() *WarehouseContactAddress {
	this := WarehouseContactAddress{}
	return &this
}

// NewWarehouseContactAddressWithDefaults instantiates a new WarehouseContactAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWarehouseContactAddressWithDefaults() *WarehouseContactAddress {
	this := WarehouseContactAddress{}
	return &this
}

// GetName1 returns the Name1 field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetName1() string {
	if o == nil || o.Name1 == nil {
		var ret string
		return ret
	}
	return *o.Name1
}

// GetName1Ok returns a tuple with the Name1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetName1Ok() (*string, bool) {
	if o == nil || o.Name1 == nil {
		return nil, false
	}
	return o.Name1, true
}

// HasName1 returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasName1() bool {
	if o != nil && o.Name1 != nil {
		return true
	}

	return false
}

// SetName1 gets a reference to the given string and assigns it to the Name1 field.
func (o *WarehouseContactAddress) SetName1(v string) {
	o.Name1 = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetAddress1() string {
	if o == nil || o.Address1 == nil {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetAddress1Ok() (*string, bool) {
	if o == nil || o.Address1 == nil {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasAddress1() bool {
	if o != nil && o.Address1 != nil {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *WarehouseContactAddress) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetAddress2() string {
	if o == nil || o.Address2 == nil {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetAddress2Ok() (*string, bool) {
	if o == nil || o.Address2 == nil {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasAddress2() bool {
	if o != nil && o.Address2 != nil {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *WarehouseContactAddress) SetAddress2(v string) {
	o.Address2 = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *WarehouseContactAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetTown returns the Town field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetTown() string {
	if o == nil || o.Town == nil {
		var ret string
		return ret
	}
	return *o.Town
}

// GetTownOk returns a tuple with the Town field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetTownOk() (*string, bool) {
	if o == nil || o.Town == nil {
		return nil, false
	}
	return o.Town, true
}

// HasTown returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasTown() bool {
	if o != nil && o.Town != nil {
		return true
	}

	return false
}

// SetTown gets a reference to the given string and assigns it to the Town field.
func (o *WarehouseContactAddress) SetTown(v string) {
	o.Town = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetCountryId() int32 {
	if o == nil || o.CountryId == nil {
		var ret int32
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetCountryIdOk() (*int32, bool) {
	if o == nil || o.CountryId == nil {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasCountryId() bool {
	if o != nil && o.CountryId != nil {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given int32 and assigns it to the CountryId field.
func (o *WarehouseContactAddress) SetCountryId(v int32) {
	o.CountryId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetOptions() []WarehouseContactAddressOptions {
	if o == nil || o.Options == nil {
		var ret []WarehouseContactAddressOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetOptionsOk() (*[]WarehouseContactAddressOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []WarehouseContactAddressOptions and assigns it to the Options field.
func (o *WarehouseContactAddress) SetOptions(v []WarehouseContactAddressOptions) {
	o.Options = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WarehouseContactAddress) SetTitle(v string) {
	o.Title = &v
}

// GetContactPerson returns the ContactPerson field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetContactPerson() string {
	if o == nil || o.ContactPerson == nil {
		var ret string
		return ret
	}
	return *o.ContactPerson
}

// GetContactPersonOk returns a tuple with the ContactPerson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetContactPersonOk() (*string, bool) {
	if o == nil || o.ContactPerson == nil {
		return nil, false
	}
	return o.ContactPerson, true
}

// HasContactPerson returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasContactPerson() bool {
	if o != nil && o.ContactPerson != nil {
		return true
	}

	return false
}

// SetContactPerson gets a reference to the given string and assigns it to the ContactPerson field.
func (o *WarehouseContactAddress) SetContactPerson(v string) {
	o.ContactPerson = &v
}

// GetTaxIdNumber returns the TaxIdNumber field value if set, zero value otherwise.
func (o *WarehouseContactAddress) GetTaxIdNumber() string {
	if o == nil || o.TaxIdNumber == nil {
		var ret string
		return ret
	}
	return *o.TaxIdNumber
}

// GetTaxIdNumberOk returns a tuple with the TaxIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WarehouseContactAddress) GetTaxIdNumberOk() (*string, bool) {
	if o == nil || o.TaxIdNumber == nil {
		return nil, false
	}
	return o.TaxIdNumber, true
}

// HasTaxIdNumber returns a boolean if a field has been set.
func (o *WarehouseContactAddress) HasTaxIdNumber() bool {
	if o != nil && o.TaxIdNumber != nil {
		return true
	}

	return false
}

// SetTaxIdNumber gets a reference to the given string and assigns it to the TaxIdNumber field.
func (o *WarehouseContactAddress) SetTaxIdNumber(v string) {
	o.TaxIdNumber = &v
}

func (o WarehouseContactAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name1 != nil {
		toSerialize["name1"] = o.Name1
	}
	if o.Address1 != nil {
		toSerialize["address1"] = o.Address1
	}
	if o.Address2 != nil {
		toSerialize["address2"] = o.Address2
	}
	if o.PostalCode != nil {
		toSerialize["postalCode"] = o.PostalCode
	}
	if o.Town != nil {
		toSerialize["town"] = o.Town
	}
	if o.CountryId != nil {
		toSerialize["countryId"] = o.CountryId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ContactPerson != nil {
		toSerialize["contactPerson"] = o.ContactPerson
	}
	if o.TaxIdNumber != nil {
		toSerialize["taxIdNumber"] = o.TaxIdNumber
	}
	return json.Marshal(toSerialize)
}

type NullableWarehouseContactAddress struct {
	value *WarehouseContactAddress
	isSet bool
}

func (v NullableWarehouseContactAddress) Get() *WarehouseContactAddress {
	return v.value
}

func (v *NullableWarehouseContactAddress) Set(val *WarehouseContactAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableWarehouseContactAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableWarehouseContactAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWarehouseContactAddress(val *WarehouseContactAddress) *NullableWarehouseContactAddress {
	return &NullableWarehouseContactAddress{value: val, isSet: true}
}

func (v NullableWarehouseContactAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWarehouseContactAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


