# VariationStock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PurchasePrice** | Pointer to **float32** |  | [optional] 
**ReservedListing** | Pointer to **float32** |  | [optional] 
**ReservedBundles** | Pointer to **float32** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**WarehouseId** | Pointer to **int32** |  | [optional] 
**PhysicalStock** | Pointer to **float32** |  | [optional] 
**ReservedStock** | Pointer to **float32** |  | [optional] 
**NetStock** | Pointer to **float32** |  | [optional] 
**ReorderLevel** | Pointer to **float32** |  | [optional] 
**DeltaReorderLevel** | Pointer to **float32** |  | [optional] 

## Methods

### NewVariationStock

`func NewVariationStock() *VariationStock`

NewVariationStock instantiates a new VariationStock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationStockWithDefaults

`func NewVariationStockWithDefaults() *VariationStock`

NewVariationStockWithDefaults instantiates a new VariationStock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchasePrice

`func (o *VariationStock) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *VariationStock) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *VariationStock) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *VariationStock) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetReservedListing

`func (o *VariationStock) GetReservedListing() float32`

GetReservedListing returns the ReservedListing field if non-nil, zero value otherwise.

### GetReservedListingOk

`func (o *VariationStock) GetReservedListingOk() (*float32, bool)`

GetReservedListingOk returns a tuple with the ReservedListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedListing

`func (o *VariationStock) SetReservedListing(v float32)`

SetReservedListing sets ReservedListing field to given value.

### HasReservedListing

`func (o *VariationStock) HasReservedListing() bool`

HasReservedListing returns a boolean if a field has been set.

### GetReservedBundles

`func (o *VariationStock) GetReservedBundles() float32`

GetReservedBundles returns the ReservedBundles field if non-nil, zero value otherwise.

### GetReservedBundlesOk

`func (o *VariationStock) GetReservedBundlesOk() (*float32, bool)`

GetReservedBundlesOk returns a tuple with the ReservedBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedBundles

`func (o *VariationStock) SetReservedBundles(v float32)`

SetReservedBundles sets ReservedBundles field to given value.

### HasReservedBundles

`func (o *VariationStock) HasReservedBundles() bool`

HasReservedBundles returns a boolean if a field has been set.

### GetVariationId

`func (o *VariationStock) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *VariationStock) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *VariationStock) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *VariationStock) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetItemId

`func (o *VariationStock) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VariationStock) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VariationStock) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VariationStock) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *VariationStock) GetWarehouseId() int32`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *VariationStock) GetWarehouseIdOk() (*int32, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *VariationStock) SetWarehouseId(v int32)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *VariationStock) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetPhysicalStock

`func (o *VariationStock) GetPhysicalStock() float32`

GetPhysicalStock returns the PhysicalStock field if non-nil, zero value otherwise.

### GetPhysicalStockOk

`func (o *VariationStock) GetPhysicalStockOk() (*float32, bool)`

GetPhysicalStockOk returns a tuple with the PhysicalStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalStock

`func (o *VariationStock) SetPhysicalStock(v float32)`

SetPhysicalStock sets PhysicalStock field to given value.

### HasPhysicalStock

`func (o *VariationStock) HasPhysicalStock() bool`

HasPhysicalStock returns a boolean if a field has been set.

### GetReservedStock

`func (o *VariationStock) GetReservedStock() float32`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *VariationStock) GetReservedStockOk() (*float32, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *VariationStock) SetReservedStock(v float32)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *VariationStock) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetNetStock

`func (o *VariationStock) GetNetStock() float32`

GetNetStock returns the NetStock field if non-nil, zero value otherwise.

### GetNetStockOk

`func (o *VariationStock) GetNetStockOk() (*float32, bool)`

GetNetStockOk returns a tuple with the NetStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetStock

`func (o *VariationStock) SetNetStock(v float32)`

SetNetStock sets NetStock field to given value.

### HasNetStock

`func (o *VariationStock) HasNetStock() bool`

HasNetStock returns a boolean if a field has been set.

### GetReorderLevel

`func (o *VariationStock) GetReorderLevel() float32`

GetReorderLevel returns the ReorderLevel field if non-nil, zero value otherwise.

### GetReorderLevelOk

`func (o *VariationStock) GetReorderLevelOk() (*float32, bool)`

GetReorderLevelOk returns a tuple with the ReorderLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderLevel

`func (o *VariationStock) SetReorderLevel(v float32)`

SetReorderLevel sets ReorderLevel field to given value.

### HasReorderLevel

`func (o *VariationStock) HasReorderLevel() bool`

HasReorderLevel returns a boolean if a field has been set.

### GetDeltaReorderLevel

`func (o *VariationStock) GetDeltaReorderLevel() float32`

GetDeltaReorderLevel returns the DeltaReorderLevel field if non-nil, zero value otherwise.

### GetDeltaReorderLevelOk

`func (o *VariationStock) GetDeltaReorderLevelOk() (*float32, bool)`

GetDeltaReorderLevelOk returns a tuple with the DeltaReorderLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaReorderLevel

`func (o *VariationStock) SetDeltaReorderLevel(v float32)`

SetDeltaReorderLevel sets DeltaReorderLevel field to given value.

### HasDeltaReorderLevel

`func (o *VariationStock) HasDeltaReorderLevel() bool`

HasDeltaReorderLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


