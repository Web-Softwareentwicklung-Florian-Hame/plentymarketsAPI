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

// PimVariationStock struct for PimVariationStock
type PimVariationStock struct {
	ItemId *int32 `json:"itemId,omitempty"`
	WarehouseId *int32 `json:"warehouseId,omitempty"`
	StockPhysical *float32 `json:"stockPhysical,omitempty"`
	ReservedStock *float32 `json:"reservedStock,omitempty"`
	ReservedEbay *float32 `json:"reservedEbay,omitempty"`
	ReorderDelta *float32 `json:"reorderDelta,omitempty"`
	StockNet *float32 `json:"stockNet,omitempty"`
	StorehouseType *string `json:"storehouse_type,omitempty"`
	Reordered *float32 `json:"reordered,omitempty"`
	ReservedBundle *float32 `json:"reservedBundle,omitempty"`
	AveragePurchasePrice *float32 `json:"averagePurchasePrice,omitempty"`
	WarehousePriority *string `json:"warehousePriority,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	VariationId *int32 `json:"variationId,omitempty"`
}

// NewPimVariationStock instantiates a new PimVariationStock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimVariationStock() *PimVariationStock {
	this := PimVariationStock{}
	return &this
}

// NewPimVariationStockWithDefaults instantiates a new PimVariationStock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimVariationStockWithDefaults() *PimVariationStock {
	this := PimVariationStock{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PimVariationStock) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PimVariationStock) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *PimVariationStock) SetItemId(v int32) {
	o.ItemId = &v
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *PimVariationStock) GetWarehouseId() int32 {
	if o == nil || o.WarehouseId == nil {
		var ret int32
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetWarehouseIdOk() (*int32, bool) {
	if o == nil || o.WarehouseId == nil {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *PimVariationStock) HasWarehouseId() bool {
	if o != nil && o.WarehouseId != nil {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int32 and assigns it to the WarehouseId field.
func (o *PimVariationStock) SetWarehouseId(v int32) {
	o.WarehouseId = &v
}

// GetStockPhysical returns the StockPhysical field value if set, zero value otherwise.
func (o *PimVariationStock) GetStockPhysical() float32 {
	if o == nil || o.StockPhysical == nil {
		var ret float32
		return ret
	}
	return *o.StockPhysical
}

// GetStockPhysicalOk returns a tuple with the StockPhysical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetStockPhysicalOk() (*float32, bool) {
	if o == nil || o.StockPhysical == nil {
		return nil, false
	}
	return o.StockPhysical, true
}

// HasStockPhysical returns a boolean if a field has been set.
func (o *PimVariationStock) HasStockPhysical() bool {
	if o != nil && o.StockPhysical != nil {
		return true
	}

	return false
}

// SetStockPhysical gets a reference to the given float32 and assigns it to the StockPhysical field.
func (o *PimVariationStock) SetStockPhysical(v float32) {
	o.StockPhysical = &v
}

// GetReservedStock returns the ReservedStock field value if set, zero value otherwise.
func (o *PimVariationStock) GetReservedStock() float32 {
	if o == nil || o.ReservedStock == nil {
		var ret float32
		return ret
	}
	return *o.ReservedStock
}

// GetReservedStockOk returns a tuple with the ReservedStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetReservedStockOk() (*float32, bool) {
	if o == nil || o.ReservedStock == nil {
		return nil, false
	}
	return o.ReservedStock, true
}

// HasReservedStock returns a boolean if a field has been set.
func (o *PimVariationStock) HasReservedStock() bool {
	if o != nil && o.ReservedStock != nil {
		return true
	}

	return false
}

// SetReservedStock gets a reference to the given float32 and assigns it to the ReservedStock field.
func (o *PimVariationStock) SetReservedStock(v float32) {
	o.ReservedStock = &v
}

// GetReservedEbay returns the ReservedEbay field value if set, zero value otherwise.
func (o *PimVariationStock) GetReservedEbay() float32 {
	if o == nil || o.ReservedEbay == nil {
		var ret float32
		return ret
	}
	return *o.ReservedEbay
}

// GetReservedEbayOk returns a tuple with the ReservedEbay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetReservedEbayOk() (*float32, bool) {
	if o == nil || o.ReservedEbay == nil {
		return nil, false
	}
	return o.ReservedEbay, true
}

// HasReservedEbay returns a boolean if a field has been set.
func (o *PimVariationStock) HasReservedEbay() bool {
	if o != nil && o.ReservedEbay != nil {
		return true
	}

	return false
}

// SetReservedEbay gets a reference to the given float32 and assigns it to the ReservedEbay field.
func (o *PimVariationStock) SetReservedEbay(v float32) {
	o.ReservedEbay = &v
}

// GetReorderDelta returns the ReorderDelta field value if set, zero value otherwise.
func (o *PimVariationStock) GetReorderDelta() float32 {
	if o == nil || o.ReorderDelta == nil {
		var ret float32
		return ret
	}
	return *o.ReorderDelta
}

// GetReorderDeltaOk returns a tuple with the ReorderDelta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetReorderDeltaOk() (*float32, bool) {
	if o == nil || o.ReorderDelta == nil {
		return nil, false
	}
	return o.ReorderDelta, true
}

// HasReorderDelta returns a boolean if a field has been set.
func (o *PimVariationStock) HasReorderDelta() bool {
	if o != nil && o.ReorderDelta != nil {
		return true
	}

	return false
}

// SetReorderDelta gets a reference to the given float32 and assigns it to the ReorderDelta field.
func (o *PimVariationStock) SetReorderDelta(v float32) {
	o.ReorderDelta = &v
}

// GetStockNet returns the StockNet field value if set, zero value otherwise.
func (o *PimVariationStock) GetStockNet() float32 {
	if o == nil || o.StockNet == nil {
		var ret float32
		return ret
	}
	return *o.StockNet
}

// GetStockNetOk returns a tuple with the StockNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetStockNetOk() (*float32, bool) {
	if o == nil || o.StockNet == nil {
		return nil, false
	}
	return o.StockNet, true
}

// HasStockNet returns a boolean if a field has been set.
func (o *PimVariationStock) HasStockNet() bool {
	if o != nil && o.StockNet != nil {
		return true
	}

	return false
}

// SetStockNet gets a reference to the given float32 and assigns it to the StockNet field.
func (o *PimVariationStock) SetStockNet(v float32) {
	o.StockNet = &v
}

// GetStorehouseType returns the StorehouseType field value if set, zero value otherwise.
func (o *PimVariationStock) GetStorehouseType() string {
	if o == nil || o.StorehouseType == nil {
		var ret string
		return ret
	}
	return *o.StorehouseType
}

// GetStorehouseTypeOk returns a tuple with the StorehouseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetStorehouseTypeOk() (*string, bool) {
	if o == nil || o.StorehouseType == nil {
		return nil, false
	}
	return o.StorehouseType, true
}

// HasStorehouseType returns a boolean if a field has been set.
func (o *PimVariationStock) HasStorehouseType() bool {
	if o != nil && o.StorehouseType != nil {
		return true
	}

	return false
}

// SetStorehouseType gets a reference to the given string and assigns it to the StorehouseType field.
func (o *PimVariationStock) SetStorehouseType(v string) {
	o.StorehouseType = &v
}

// GetReordered returns the Reordered field value if set, zero value otherwise.
func (o *PimVariationStock) GetReordered() float32 {
	if o == nil || o.Reordered == nil {
		var ret float32
		return ret
	}
	return *o.Reordered
}

// GetReorderedOk returns a tuple with the Reordered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetReorderedOk() (*float32, bool) {
	if o == nil || o.Reordered == nil {
		return nil, false
	}
	return o.Reordered, true
}

// HasReordered returns a boolean if a field has been set.
func (o *PimVariationStock) HasReordered() bool {
	if o != nil && o.Reordered != nil {
		return true
	}

	return false
}

// SetReordered gets a reference to the given float32 and assigns it to the Reordered field.
func (o *PimVariationStock) SetReordered(v float32) {
	o.Reordered = &v
}

// GetReservedBundle returns the ReservedBundle field value if set, zero value otherwise.
func (o *PimVariationStock) GetReservedBundle() float32 {
	if o == nil || o.ReservedBundle == nil {
		var ret float32
		return ret
	}
	return *o.ReservedBundle
}

// GetReservedBundleOk returns a tuple with the ReservedBundle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetReservedBundleOk() (*float32, bool) {
	if o == nil || o.ReservedBundle == nil {
		return nil, false
	}
	return o.ReservedBundle, true
}

// HasReservedBundle returns a boolean if a field has been set.
func (o *PimVariationStock) HasReservedBundle() bool {
	if o != nil && o.ReservedBundle != nil {
		return true
	}

	return false
}

// SetReservedBundle gets a reference to the given float32 and assigns it to the ReservedBundle field.
func (o *PimVariationStock) SetReservedBundle(v float32) {
	o.ReservedBundle = &v
}

// GetAveragePurchasePrice returns the AveragePurchasePrice field value if set, zero value otherwise.
func (o *PimVariationStock) GetAveragePurchasePrice() float32 {
	if o == nil || o.AveragePurchasePrice == nil {
		var ret float32
		return ret
	}
	return *o.AveragePurchasePrice
}

// GetAveragePurchasePriceOk returns a tuple with the AveragePurchasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetAveragePurchasePriceOk() (*float32, bool) {
	if o == nil || o.AveragePurchasePrice == nil {
		return nil, false
	}
	return o.AveragePurchasePrice, true
}

// HasAveragePurchasePrice returns a boolean if a field has been set.
func (o *PimVariationStock) HasAveragePurchasePrice() bool {
	if o != nil && o.AveragePurchasePrice != nil {
		return true
	}

	return false
}

// SetAveragePurchasePrice gets a reference to the given float32 and assigns it to the AveragePurchasePrice field.
func (o *PimVariationStock) SetAveragePurchasePrice(v float32) {
	o.AveragePurchasePrice = &v
}

// GetWarehousePriority returns the WarehousePriority field value if set, zero value otherwise.
func (o *PimVariationStock) GetWarehousePriority() string {
	if o == nil || o.WarehousePriority == nil {
		var ret string
		return ret
	}
	return *o.WarehousePriority
}

// GetWarehousePriorityOk returns a tuple with the WarehousePriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetWarehousePriorityOk() (*string, bool) {
	if o == nil || o.WarehousePriority == nil {
		return nil, false
	}
	return o.WarehousePriority, true
}

// HasWarehousePriority returns a boolean if a field has been set.
func (o *PimVariationStock) HasWarehousePriority() bool {
	if o != nil && o.WarehousePriority != nil {
		return true
	}

	return false
}

// SetWarehousePriority gets a reference to the given string and assigns it to the WarehousePriority field.
func (o *PimVariationStock) SetWarehousePriority(v string) {
	o.WarehousePriority = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PimVariationStock) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PimVariationStock) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PimVariationStock) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVariationId returns the VariationId field value if set, zero value otherwise.
func (o *PimVariationStock) GetVariationId() int32 {
	if o == nil || o.VariationId == nil {
		var ret int32
		return ret
	}
	return *o.VariationId
}

// GetVariationIdOk returns a tuple with the VariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationStock) GetVariationIdOk() (*int32, bool) {
	if o == nil || o.VariationId == nil {
		return nil, false
	}
	return o.VariationId, true
}

// HasVariationId returns a boolean if a field has been set.
func (o *PimVariationStock) HasVariationId() bool {
	if o != nil && o.VariationId != nil {
		return true
	}

	return false
}

// SetVariationId gets a reference to the given int32 and assigns it to the VariationId field.
func (o *PimVariationStock) SetVariationId(v int32) {
	o.VariationId = &v
}

func (o PimVariationStock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ItemId != nil {
		toSerialize["itemId"] = o.ItemId
	}
	if o.WarehouseId != nil {
		toSerialize["warehouseId"] = o.WarehouseId
	}
	if o.StockPhysical != nil {
		toSerialize["stockPhysical"] = o.StockPhysical
	}
	if o.ReservedStock != nil {
		toSerialize["reservedStock"] = o.ReservedStock
	}
	if o.ReservedEbay != nil {
		toSerialize["reservedEbay"] = o.ReservedEbay
	}
	if o.ReorderDelta != nil {
		toSerialize["reorderDelta"] = o.ReorderDelta
	}
	if o.StockNet != nil {
		toSerialize["stockNet"] = o.StockNet
	}
	if o.StorehouseType != nil {
		toSerialize["storehouse_type"] = o.StorehouseType
	}
	if o.Reordered != nil {
		toSerialize["reordered"] = o.Reordered
	}
	if o.ReservedBundle != nil {
		toSerialize["reservedBundle"] = o.ReservedBundle
	}
	if o.AveragePurchasePrice != nil {
		toSerialize["averagePurchasePrice"] = o.AveragePurchasePrice
	}
	if o.WarehousePriority != nil {
		toSerialize["warehousePriority"] = o.WarehousePriority
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.VariationId != nil {
		toSerialize["variationId"] = o.VariationId
	}
	return json.Marshal(toSerialize)
}

type NullablePimVariationStock struct {
	value *PimVariationStock
	isSet bool
}

func (v NullablePimVariationStock) Get() *PimVariationStock {
	return v.value
}

func (v *NullablePimVariationStock) Set(val *PimVariationStock) {
	v.value = val
	v.isSet = true
}

func (v NullablePimVariationStock) IsSet() bool {
	return v.isSet
}

func (v *NullablePimVariationStock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimVariationStock(val *PimVariationStock) *NullablePimVariationStock {
	return &NullablePimVariationStock{value: val, isSet: true}
}

func (v NullablePimVariationStock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePimVariationStock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


