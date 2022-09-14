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

// PimVariationWarehouse struct for PimVariationWarehouse
type PimVariationWarehouse struct {
	WarehouseId *int32 `json:"warehouseId,omitempty"`
	ZoneId *int32 `json:"zoneId,omitempty"`
	StorageLocationType *string `json:"storageLocationType,omitempty"`
	ReorderLevel *int32 `json:"reorderLevel,omitempty"`
	MaximumStock *int32 `json:"maximumStock,omitempty"`
	StockTurnoverInDays *int32 `json:"stockTurnoverInDays,omitempty"`
	StorageLocation *int32 `json:"storageLocation,omitempty"`
	StockBuffer *int32 `json:"stockBuffer,omitempty"`
	IsBatch *bool `json:"isBatch,omitempty"`
	IsBestBeforeDate *bool `json:"isBestBeforeDate,omitempty"`
}

// NewPimVariationWarehouse instantiates a new PimVariationWarehouse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPimVariationWarehouse() *PimVariationWarehouse {
	this := PimVariationWarehouse{}
	return &this
}

// NewPimVariationWarehouseWithDefaults instantiates a new PimVariationWarehouse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPimVariationWarehouseWithDefaults() *PimVariationWarehouse {
	this := PimVariationWarehouse{}
	return &this
}

// GetWarehouseId returns the WarehouseId field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetWarehouseId() int32 {
	if o == nil || o.WarehouseId == nil {
		var ret int32
		return ret
	}
	return *o.WarehouseId
}

// GetWarehouseIdOk returns a tuple with the WarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetWarehouseIdOk() (*int32, bool) {
	if o == nil || o.WarehouseId == nil {
		return nil, false
	}
	return o.WarehouseId, true
}

// HasWarehouseId returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasWarehouseId() bool {
	if o != nil && o.WarehouseId != nil {
		return true
	}

	return false
}

// SetWarehouseId gets a reference to the given int32 and assigns it to the WarehouseId field.
func (o *PimVariationWarehouse) SetWarehouseId(v int32) {
	o.WarehouseId = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetZoneId() int32 {
	if o == nil || o.ZoneId == nil {
		var ret int32
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetZoneIdOk() (*int32, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int32 and assigns it to the ZoneId field.
func (o *PimVariationWarehouse) SetZoneId(v int32) {
	o.ZoneId = &v
}

// GetStorageLocationType returns the StorageLocationType field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetStorageLocationType() string {
	if o == nil || o.StorageLocationType == nil {
		var ret string
		return ret
	}
	return *o.StorageLocationType
}

// GetStorageLocationTypeOk returns a tuple with the StorageLocationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetStorageLocationTypeOk() (*string, bool) {
	if o == nil || o.StorageLocationType == nil {
		return nil, false
	}
	return o.StorageLocationType, true
}

// HasStorageLocationType returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasStorageLocationType() bool {
	if o != nil && o.StorageLocationType != nil {
		return true
	}

	return false
}

// SetStorageLocationType gets a reference to the given string and assigns it to the StorageLocationType field.
func (o *PimVariationWarehouse) SetStorageLocationType(v string) {
	o.StorageLocationType = &v
}

// GetReorderLevel returns the ReorderLevel field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetReorderLevel() int32 {
	if o == nil || o.ReorderLevel == nil {
		var ret int32
		return ret
	}
	return *o.ReorderLevel
}

// GetReorderLevelOk returns a tuple with the ReorderLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetReorderLevelOk() (*int32, bool) {
	if o == nil || o.ReorderLevel == nil {
		return nil, false
	}
	return o.ReorderLevel, true
}

// HasReorderLevel returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasReorderLevel() bool {
	if o != nil && o.ReorderLevel != nil {
		return true
	}

	return false
}

// SetReorderLevel gets a reference to the given int32 and assigns it to the ReorderLevel field.
func (o *PimVariationWarehouse) SetReorderLevel(v int32) {
	o.ReorderLevel = &v
}

// GetMaximumStock returns the MaximumStock field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetMaximumStock() int32 {
	if o == nil || o.MaximumStock == nil {
		var ret int32
		return ret
	}
	return *o.MaximumStock
}

// GetMaximumStockOk returns a tuple with the MaximumStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetMaximumStockOk() (*int32, bool) {
	if o == nil || o.MaximumStock == nil {
		return nil, false
	}
	return o.MaximumStock, true
}

// HasMaximumStock returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasMaximumStock() bool {
	if o != nil && o.MaximumStock != nil {
		return true
	}

	return false
}

// SetMaximumStock gets a reference to the given int32 and assigns it to the MaximumStock field.
func (o *PimVariationWarehouse) SetMaximumStock(v int32) {
	o.MaximumStock = &v
}

// GetStockTurnoverInDays returns the StockTurnoverInDays field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetStockTurnoverInDays() int32 {
	if o == nil || o.StockTurnoverInDays == nil {
		var ret int32
		return ret
	}
	return *o.StockTurnoverInDays
}

// GetStockTurnoverInDaysOk returns a tuple with the StockTurnoverInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetStockTurnoverInDaysOk() (*int32, bool) {
	if o == nil || o.StockTurnoverInDays == nil {
		return nil, false
	}
	return o.StockTurnoverInDays, true
}

// HasStockTurnoverInDays returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasStockTurnoverInDays() bool {
	if o != nil && o.StockTurnoverInDays != nil {
		return true
	}

	return false
}

// SetStockTurnoverInDays gets a reference to the given int32 and assigns it to the StockTurnoverInDays field.
func (o *PimVariationWarehouse) SetStockTurnoverInDays(v int32) {
	o.StockTurnoverInDays = &v
}

// GetStorageLocation returns the StorageLocation field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetStorageLocation() int32 {
	if o == nil || o.StorageLocation == nil {
		var ret int32
		return ret
	}
	return *o.StorageLocation
}

// GetStorageLocationOk returns a tuple with the StorageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetStorageLocationOk() (*int32, bool) {
	if o == nil || o.StorageLocation == nil {
		return nil, false
	}
	return o.StorageLocation, true
}

// HasStorageLocation returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasStorageLocation() bool {
	if o != nil && o.StorageLocation != nil {
		return true
	}

	return false
}

// SetStorageLocation gets a reference to the given int32 and assigns it to the StorageLocation field.
func (o *PimVariationWarehouse) SetStorageLocation(v int32) {
	o.StorageLocation = &v
}

// GetStockBuffer returns the StockBuffer field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetStockBuffer() int32 {
	if o == nil || o.StockBuffer == nil {
		var ret int32
		return ret
	}
	return *o.StockBuffer
}

// GetStockBufferOk returns a tuple with the StockBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetStockBufferOk() (*int32, bool) {
	if o == nil || o.StockBuffer == nil {
		return nil, false
	}
	return o.StockBuffer, true
}

// HasStockBuffer returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasStockBuffer() bool {
	if o != nil && o.StockBuffer != nil {
		return true
	}

	return false
}

// SetStockBuffer gets a reference to the given int32 and assigns it to the StockBuffer field.
func (o *PimVariationWarehouse) SetStockBuffer(v int32) {
	o.StockBuffer = &v
}

// GetIsBatch returns the IsBatch field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetIsBatch() bool {
	if o == nil || o.IsBatch == nil {
		var ret bool
		return ret
	}
	return *o.IsBatch
}

// GetIsBatchOk returns a tuple with the IsBatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetIsBatchOk() (*bool, bool) {
	if o == nil || o.IsBatch == nil {
		return nil, false
	}
	return o.IsBatch, true
}

// HasIsBatch returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasIsBatch() bool {
	if o != nil && o.IsBatch != nil {
		return true
	}

	return false
}

// SetIsBatch gets a reference to the given bool and assigns it to the IsBatch field.
func (o *PimVariationWarehouse) SetIsBatch(v bool) {
	o.IsBatch = &v
}

// GetIsBestBeforeDate returns the IsBestBeforeDate field value if set, zero value otherwise.
func (o *PimVariationWarehouse) GetIsBestBeforeDate() bool {
	if o == nil || o.IsBestBeforeDate == nil {
		var ret bool
		return ret
	}
	return *o.IsBestBeforeDate
}

// GetIsBestBeforeDateOk returns a tuple with the IsBestBeforeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PimVariationWarehouse) GetIsBestBeforeDateOk() (*bool, bool) {
	if o == nil || o.IsBestBeforeDate == nil {
		return nil, false
	}
	return o.IsBestBeforeDate, true
}

// HasIsBestBeforeDate returns a boolean if a field has been set.
func (o *PimVariationWarehouse) HasIsBestBeforeDate() bool {
	if o != nil && o.IsBestBeforeDate != nil {
		return true
	}

	return false
}

// SetIsBestBeforeDate gets a reference to the given bool and assigns it to the IsBestBeforeDate field.
func (o *PimVariationWarehouse) SetIsBestBeforeDate(v bool) {
	o.IsBestBeforeDate = &v
}

func (o PimVariationWarehouse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.WarehouseId != nil {
		toSerialize["warehouseId"] = o.WarehouseId
	}
	if o.ZoneId != nil {
		toSerialize["zoneId"] = o.ZoneId
	}
	if o.StorageLocationType != nil {
		toSerialize["storageLocationType"] = o.StorageLocationType
	}
	if o.ReorderLevel != nil {
		toSerialize["reorderLevel"] = o.ReorderLevel
	}
	if o.MaximumStock != nil {
		toSerialize["maximumStock"] = o.MaximumStock
	}
	if o.StockTurnoverInDays != nil {
		toSerialize["stockTurnoverInDays"] = o.StockTurnoverInDays
	}
	if o.StorageLocation != nil {
		toSerialize["storageLocation"] = o.StorageLocation
	}
	if o.StockBuffer != nil {
		toSerialize["stockBuffer"] = o.StockBuffer
	}
	if o.IsBatch != nil {
		toSerialize["isBatch"] = o.IsBatch
	}
	if o.IsBestBeforeDate != nil {
		toSerialize["isBestBeforeDate"] = o.IsBestBeforeDate
	}
	return json.Marshal(toSerialize)
}

type NullablePimVariationWarehouse struct {
	value *PimVariationWarehouse
	isSet bool
}

func (v NullablePimVariationWarehouse) Get() *PimVariationWarehouse {
	return v.value
}

func (v *NullablePimVariationWarehouse) Set(val *PimVariationWarehouse) {
	v.value = val
	v.isSet = true
}

func (v NullablePimVariationWarehouse) IsSet() bool {
	return v.isSet
}

func (v *NullablePimVariationWarehouse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePimVariationWarehouse(val *PimVariationWarehouse) *NullablePimVariationWarehouse {
	return &NullablePimVariationWarehouse{value: val, isSet: true}
}

func (v NullablePimVariationWarehouse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePimVariationWarehouse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

