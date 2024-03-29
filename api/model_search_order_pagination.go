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

// SearchOrderPagination struct for SearchOrderPagination
type SearchOrderPagination struct {
	Page *int32 `json:"page,omitempty"`
	TotalsCount *int32 `json:"totalsCount,omitempty"`
	IsLastPage *bool `json:"isLastPage,omitempty"`
	LastPageNumber *int32 `json:"lastPageNumber,omitempty"`
	FirstOnPage *int32 `json:"firstOnPage,omitempty"`
	LastOnPage *int32 `json:"lastOnPage,omitempty"`
	ItemsPerPage *int32 `json:"itemsPerPage,omitempty"`
	Entries *[]Order `json:"entries,omitempty"`
}

// NewSearchOrderPagination instantiates a new SearchOrderPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchOrderPagination() *SearchOrderPagination {
	this := SearchOrderPagination{}
	return &this
}

// NewSearchOrderPaginationWithDefaults instantiates a new SearchOrderPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchOrderPaginationWithDefaults() *SearchOrderPagination {
	this := SearchOrderPagination{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *SearchOrderPagination) SetPage(v int32) {
	o.Page = &v
}

// GetTotalsCount returns the TotalsCount field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetTotalsCount() int32 {
	if o == nil || o.TotalsCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalsCount
}

// GetTotalsCountOk returns a tuple with the TotalsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetTotalsCountOk() (*int32, bool) {
	if o == nil || o.TotalsCount == nil {
		return nil, false
	}
	return o.TotalsCount, true
}

// HasTotalsCount returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasTotalsCount() bool {
	if o != nil && o.TotalsCount != nil {
		return true
	}

	return false
}

// SetTotalsCount gets a reference to the given int32 and assigns it to the TotalsCount field.
func (o *SearchOrderPagination) SetTotalsCount(v int32) {
	o.TotalsCount = &v
}

// GetIsLastPage returns the IsLastPage field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetIsLastPage() bool {
	if o == nil || o.IsLastPage == nil {
		var ret bool
		return ret
	}
	return *o.IsLastPage
}

// GetIsLastPageOk returns a tuple with the IsLastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetIsLastPageOk() (*bool, bool) {
	if o == nil || o.IsLastPage == nil {
		return nil, false
	}
	return o.IsLastPage, true
}

// HasIsLastPage returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasIsLastPage() bool {
	if o != nil && o.IsLastPage != nil {
		return true
	}

	return false
}

// SetIsLastPage gets a reference to the given bool and assigns it to the IsLastPage field.
func (o *SearchOrderPagination) SetIsLastPage(v bool) {
	o.IsLastPage = &v
}

// GetLastPageNumber returns the LastPageNumber field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetLastPageNumber() int32 {
	if o == nil || o.LastPageNumber == nil {
		var ret int32
		return ret
	}
	return *o.LastPageNumber
}

// GetLastPageNumberOk returns a tuple with the LastPageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetLastPageNumberOk() (*int32, bool) {
	if o == nil || o.LastPageNumber == nil {
		return nil, false
	}
	return o.LastPageNumber, true
}

// HasLastPageNumber returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasLastPageNumber() bool {
	if o != nil && o.LastPageNumber != nil {
		return true
	}

	return false
}

// SetLastPageNumber gets a reference to the given int32 and assigns it to the LastPageNumber field.
func (o *SearchOrderPagination) SetLastPageNumber(v int32) {
	o.LastPageNumber = &v
}

// GetFirstOnPage returns the FirstOnPage field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetFirstOnPage() int32 {
	if o == nil || o.FirstOnPage == nil {
		var ret int32
		return ret
	}
	return *o.FirstOnPage
}

// GetFirstOnPageOk returns a tuple with the FirstOnPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetFirstOnPageOk() (*int32, bool) {
	if o == nil || o.FirstOnPage == nil {
		return nil, false
	}
	return o.FirstOnPage, true
}

// HasFirstOnPage returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasFirstOnPage() bool {
	if o != nil && o.FirstOnPage != nil {
		return true
	}

	return false
}

// SetFirstOnPage gets a reference to the given int32 and assigns it to the FirstOnPage field.
func (o *SearchOrderPagination) SetFirstOnPage(v int32) {
	o.FirstOnPage = &v
}

// GetLastOnPage returns the LastOnPage field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetLastOnPage() int32 {
	if o == nil || o.LastOnPage == nil {
		var ret int32
		return ret
	}
	return *o.LastOnPage
}

// GetLastOnPageOk returns a tuple with the LastOnPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetLastOnPageOk() (*int32, bool) {
	if o == nil || o.LastOnPage == nil {
		return nil, false
	}
	return o.LastOnPage, true
}

// HasLastOnPage returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasLastOnPage() bool {
	if o != nil && o.LastOnPage != nil {
		return true
	}

	return false
}

// SetLastOnPage gets a reference to the given int32 and assigns it to the LastOnPage field.
func (o *SearchOrderPagination) SetLastOnPage(v int32) {
	o.LastOnPage = &v
}

// GetItemsPerPage returns the ItemsPerPage field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetItemsPerPage() int32 {
	if o == nil || o.ItemsPerPage == nil {
		var ret int32
		return ret
	}
	return *o.ItemsPerPage
}

// GetItemsPerPageOk returns a tuple with the ItemsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetItemsPerPageOk() (*int32, bool) {
	if o == nil || o.ItemsPerPage == nil {
		return nil, false
	}
	return o.ItemsPerPage, true
}

// HasItemsPerPage returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasItemsPerPage() bool {
	if o != nil && o.ItemsPerPage != nil {
		return true
	}

	return false
}

// SetItemsPerPage gets a reference to the given int32 and assigns it to the ItemsPerPage field.
func (o *SearchOrderPagination) SetItemsPerPage(v int32) {
	o.ItemsPerPage = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *SearchOrderPagination) GetEntries() []Order {
	if o == nil || o.Entries == nil {
		var ret []Order
		return ret
	}
	return *o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchOrderPagination) GetEntriesOk() (*[]Order, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *SearchOrderPagination) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []Order and assigns it to the Entries field.
func (o *SearchOrderPagination) SetEntries(v []Order) {
	o.Entries = &v
}

func (o SearchOrderPagination) MarshalJSON() ([]byte, error) {
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

type NullableSearchOrderPagination struct {
	value *SearchOrderPagination
	isSet bool
}

func (v NullableSearchOrderPagination) Get() *SearchOrderPagination {
	return v.value
}

func (v *NullableSearchOrderPagination) Set(val *SearchOrderPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchOrderPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchOrderPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchOrderPagination(val *SearchOrderPagination) *NullableSearchOrderPagination {
	return &NullableSearchOrderPagination{value: val, isSet: true}
}

func (v NullableSearchOrderPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchOrderPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


