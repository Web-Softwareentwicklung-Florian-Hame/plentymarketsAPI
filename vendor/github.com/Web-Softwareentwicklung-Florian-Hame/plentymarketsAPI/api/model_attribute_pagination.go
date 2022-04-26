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

// AttributePagination pagination response containing categories
type AttributePagination struct {
	Page *int32 `json:"page,omitempty"`
	TotalsCount *int32 `json:"totalsCount,omitempty"`
	IsLastPage *bool `json:"isLastPage,omitempty"`
	LastPageNumber *int32 `json:"lastPageNumber,omitempty"`
	FirstOnPage *int32 `json:"firstOnPage,omitempty"`
	LastOnPage *int32 `json:"lastOnPage,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Entries *[]Attribute `json:"entries,omitempty"`
}

// NewAttributePagination instantiates a new AttributePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributePagination() *AttributePagination {
	this := AttributePagination{}
	return &this
}

// NewAttributePaginationWithDefaults instantiates a new AttributePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributePaginationWithDefaults() *AttributePagination {
	this := AttributePagination{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AttributePagination) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AttributePagination) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AttributePagination) SetPage(v int32) {
	o.Page = &v
}

// GetTotalsCount returns the TotalsCount field value if set, zero value otherwise.
func (o *AttributePagination) GetTotalsCount() int32 {
	if o == nil || o.TotalsCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalsCount
}

// GetTotalsCountOk returns a tuple with the TotalsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetTotalsCountOk() (*int32, bool) {
	if o == nil || o.TotalsCount == nil {
		return nil, false
	}
	return o.TotalsCount, true
}

// HasTotalsCount returns a boolean if a field has been set.
func (o *AttributePagination) HasTotalsCount() bool {
	if o != nil && o.TotalsCount != nil {
		return true
	}

	return false
}

// SetTotalsCount gets a reference to the given int32 and assigns it to the TotalsCount field.
func (o *AttributePagination) SetTotalsCount(v int32) {
	o.TotalsCount = &v
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *AttributePagination) GetIsLastPage() bool {
	if o == nil || o.IsLastPage == nil {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetIsLastPageOk() (*bool, bool) {
	if o == nil || o.IsLastPage == nil {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *AttributePagination) HasIsLastPage() bool {
	if o != nil && o.IsLastPage != nil {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *AttributePagination) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetLastPageNumber returns the LastPageNumber field value if set, zero value otherwise.
func (o *AttributePagination) GetLastPageNumber() int32 {
	if o == nil || o.LastPageNumber == nil {
		var ret int32
		return ret
	}
	return *o.LastPageNumber
}

// GetLastPageNumberOk returns a tuple with the LastPageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetLastPageNumberOk() (*int32, bool) {
	if o == nil || o.LastPageNumber == nil {
		return nil, false
	}
	return o.LastPageNumber, true
}

// HasLastPageNumber returns a boolean if a field has been set.
func (o *AttributePagination) HasLastPageNumber() bool {
	if o != nil && o.LastPageNumber != nil {
		return true
	}

	return false
}

// SetLastPageNumber gets a reference to the given int32 and assigns it to the LastPageNumber field.
func (o *AttributePagination) SetLastPageNumber(v int32) {
	o.LastPageNumber = &v
}

// GetFirstOnPage returns the FirstOnPage field value if set, zero value otherwise.
func (o *AttributePagination) GetFirstOnPage() int32 {
	if o == nil || o.FirstOnPage == nil {
		var ret int32
		return ret
	}
	return *o.FirstOnPage
}

// GetFirstOnPageOk returns a tuple with the FirstOnPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetFirstOnPageOk() (*int32, bool) {
	if o == nil || o.FirstOnPage == nil {
		return nil, false
	}
	return o.FirstOnPage, true
}

// HasFirstOnPage returns a boolean if a field has been set.
func (o *AttributePagination) HasFirstOnPage() bool {
	if o != nil && o.FirstOnPage != nil {
		return true
	}

	return false
}

// SetFirstOnPage gets a reference to the given int32 and assigns it to the FirstOnPage field.
func (o *AttributePagination) SetFirstOnPage(v int32) {
	o.FirstOnPage = &v
}

// GetLastOnPage returns the LastOnPage field value if set, zero value otherwise.
func (o *AttributePagination) GetLastOnPage() int32 {
	if o == nil || o.LastOnPage == nil {
		var ret int32
		return ret
	}
	return *o.LastOnPage
}

// GetLastOnPageOk returns a tuple with the LastOnPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetLastOnPageOk() (*int32, bool) {
	if o == nil || o.LastOnPage == nil {
		return nil, false
	}
	return o.LastOnPage, true
}

// HasLastOnPage returns a boolean if a field has been set.
func (o *AttributePagination) HasLastOnPage() bool {
	if o != nil && o.LastOnPage != nil {
		return true
	}

	return false
}

// SetLastOnPage gets a reference to the given int32 and assigns it to the LastOnPage field.
func (o *AttributePagination) SetLastOnPage(v int32) {
	o.LastOnPage = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *AttributePagination) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *AttributePagination) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *AttributePagination) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *AttributePagination) GetEntries() []Attribute {
	if o == nil || o.Entries == nil {
		var ret []Attribute
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributePagination) GetEntriesOk() (*[]Attribute, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *AttributePagination) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []Attribute and assigns it to the Entries field.
func (o *AttributePagination) SetEntries(v []Attribute) {
	o.Entries = &v
}

func (o AttributePagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.TotalsCount != nil {
		toSerialize["totalsCount"] = o.TotalsCount
	}
	if o.IsLastPage != nil {
		toSerialize["isLastPage"] = o.IsLastPage
	}
	if o.LastPageNumber != nil {
		toSerialize["lastPageNumber"] = o.LastPageNumber
	}
	if o.FirstOnPage != nil {
		toSerialize["firstOnPage"] = o.FirstOnPage
	}
	if o.LastOnPage != nil {
		toSerialize["lastOnPage"] = o.LastOnPage
	}
	if o.ItemsPerPage != nil {
		toSerialize["itemsPerPage"] = o.ItemsPerPage
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullableAttributePagination struct {
	value *AttributePagination
	isSet bool
}

func (v NullableAttributePagination) Get() *AttributePagination {
	return v.value
}

func (v *NullableAttributePagination) Set(val *AttributePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributePagination(val *AttributePagination) *NullableAttributePagination {
	return &NullableAttributePagination{value: val, isSet: true}
}

func (v NullableAttributePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


