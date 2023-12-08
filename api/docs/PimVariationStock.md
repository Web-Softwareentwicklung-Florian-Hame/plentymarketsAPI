# PimVariationStock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | Pointer to **int32** |  | [optional] 
**WarehouseId** | Pointer to **int32** |  | [optional] 
**StockPhysical** | Pointer to **float32** |  | [optional] 
**ReservedStock** | Pointer to **float32** |  | [optional] 
**ReservedEbay** | Pointer to **float32** |  | [optional] 
**ReorderDelta** | Pointer to **float32** |  | [optional] 
**StockNet** | Pointer to **float32** |  | [optional] 
**StorehouseType** | Pointer to **string** |  | [optional] 
**Reordered** | Pointer to **float32** |  | [optional] 
**ReservedBundle** | Pointer to **float32** |  | [optional] 
**AveragePurchasePrice** | Pointer to **float32** |  | [optional] 
**WarehousePriority** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewPimVariationStock

`func NewPimVariationStock() *PimVariationStock`

NewPimVariationStock instantiates a new PimVariationStock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimVariationStockWithDefaults

`func NewPimVariationStockWithDefaults() *PimVariationStock`

NewPimVariationStockWithDefaults instantiates a new PimVariationStock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *PimVariationStock) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PimVariationStock) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PimVariationStock) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PimVariationStock) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetWarehouseId

`func (o *PimVariationStock) GetWarehouseId() int32`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *PimVariationStock) GetWarehouseIdOk() (*int32, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *PimVariationStock) SetWarehouseId(v int32)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *PimVariationStock) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetStockPhysical

`func (o *PimVariationStock) GetStockPhysical() float32`

GetStockPhysical returns the StockPhysical field if non-nil, zero value otherwise.

### GetStockPhysicalOk

`func (o *PimVariationStock) GetStockPhysicalOk() (*float32, bool)`

GetStockPhysicalOk returns a tuple with the StockPhysical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockPhysical

`func (o *PimVariationStock) SetStockPhysical(v float32)`

SetStockPhysical sets StockPhysical field to given value.

### HasStockPhysical

`func (o *PimVariationStock) HasStockPhysical() bool`

HasStockPhysical returns a boolean if a field has been set.

### GetReservedStock

`func (o *PimVariationStock) GetReservedStock() float32`

GetReservedStock returns the ReservedStock field if non-nil, zero value otherwise.

### GetReservedStockOk

`func (o *PimVariationStock) GetReservedStockOk() (*float32, bool)`

GetReservedStockOk returns a tuple with the ReservedStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedStock

`func (o *PimVariationStock) SetReservedStock(v float32)`

SetReservedStock sets ReservedStock field to given value.

### HasReservedStock

`func (o *PimVariationStock) HasReservedStock() bool`

HasReservedStock returns a boolean if a field has been set.

### GetReservedEbay

`func (o *PimVariationStock) GetReservedEbay() float32`

GetReservedEbay returns the ReservedEbay field if non-nil, zero value otherwise.

### GetReservedEbayOk

`func (o *PimVariationStock) GetReservedEbayOk() (*float32, bool)`

GetReservedEbayOk returns a tuple with the ReservedEbay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedEbay

`func (o *PimVariationStock) SetReservedEbay(v float32)`

SetReservedEbay sets ReservedEbay field to given value.

### HasReservedEbay

`func (o *PimVariationStock) HasReservedEbay() bool`

HasReservedEbay returns a boolean if a field has been set.

### GetReorderDelta

`func (o *PimVariationStock) GetReorderDelta() float32`

GetReorderDelta returns the ReorderDelta field if non-nil, zero value otherwise.

### GetReorderDeltaOk

`func (o *PimVariationStock) GetReorderDeltaOk() (*float32, bool)`

GetReorderDeltaOk returns a tuple with the ReorderDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderDelta

`func (o *PimVariationStock) SetReorderDelta(v float32)`

SetReorderDelta sets ReorderDelta field to given value.

### HasReorderDelta

`func (o *PimVariationStock) HasReorderDelta() bool`

HasReorderDelta returns a boolean if a field has been set.

### GetStockNet

`func (o *PimVariationStock) GetStockNet() float32`

GetStockNet returns the StockNet field if non-nil, zero value otherwise.

### GetStockNetOk

`func (o *PimVariationStock) GetStockNetOk() (*float32, bool)`

GetStockNetOk returns a tuple with the StockNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNet

`func (o *PimVariationStock) SetStockNet(v float32)`

SetStockNet sets StockNet field to given value.

### HasStockNet

`func (o *PimVariationStock) HasStockNet() bool`

HasStockNet returns a boolean if a field has been set.

### GetStorehouseType

`func (o *PimVariationStock) GetStorehouseType() string`

GetStorehouseType returns the StorehouseType field if non-nil, zero value otherwise.

### GetStorehouseTypeOk

`func (o *PimVariationStock) GetStorehouseTypeOk() (*string, bool)`

GetStorehouseTypeOk returns a tuple with the StorehouseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorehouseType

`func (o *PimVariationStock) SetStorehouseType(v string)`

SetStorehouseType sets StorehouseType field to given value.

### HasStorehouseType

`func (o *PimVariationStock) HasStorehouseType() bool`

HasStorehouseType returns a boolean if a field has been set.

### GetReordered

`func (o *PimVariationStock) GetReordered() float32`

GetReordered returns the Reordered field if non-nil, zero value otherwise.

### GetReorderedOk

`func (o *PimVariationStock) GetReorderedOk() (*float32, bool)`

GetReorderedOk returns a tuple with the Reordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReordered

`func (o *PimVariationStock) SetReordered(v float32)`

SetReordered sets Reordered field to given value.

### HasReordered

`func (o *PimVariationStock) HasReordered() bool`

HasReordered returns a boolean if a field has been set.

### GetReservedBundle

`func (o *PimVariationStock) GetReservedBundle() float32`

GetReservedBundle returns the ReservedBundle field if non-nil, zero value otherwise.

### GetReservedBundleOk

`func (o *PimVariationStock) GetReservedBundleOk() (*float32, bool)`

GetReservedBundleOk returns a tuple with the ReservedBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedBundle

`func (o *PimVariationStock) SetReservedBundle(v float32)`

SetReservedBundle sets ReservedBundle field to given value.

### HasReservedBundle

`func (o *PimVariationStock) HasReservedBundle() bool`

HasReservedBundle returns a boolean if a field has been set.

### GetAveragePurchasePrice

`func (o *PimVariationStock) GetAveragePurchasePrice() float32`

GetAveragePurchasePrice returns the AveragePurchasePrice field if non-nil, zero value otherwise.

### GetAveragePurchasePriceOk

`func (o *PimVariationStock) GetAveragePurchasePriceOk() (*float32, bool)`

GetAveragePurchasePriceOk returns a tuple with the AveragePurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePurchasePrice

`func (o *PimVariationStock) SetAveragePurchasePrice(v float32)`

SetAveragePurchasePrice sets AveragePurchasePrice field to given value.

### HasAveragePurchasePrice

`func (o *PimVariationStock) HasAveragePurchasePrice() bool`

HasAveragePurchasePrice returns a boolean if a field has been set.

### GetWarehousePriority

`func (o *PimVariationStock) GetWarehousePriority() string`

GetWarehousePriority returns the WarehousePriority field if non-nil, zero value otherwise.

### GetWarehousePriorityOk

`func (o *PimVariationStock) GetWarehousePriorityOk() (*string, bool)`

GetWarehousePriorityOk returns a tuple with the WarehousePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehousePriority

`func (o *PimVariationStock) SetWarehousePriority(v string)`

SetWarehousePriority sets WarehousePriority field to given value.

### HasWarehousePriority

`func (o *PimVariationStock) HasWarehousePriority() bool`

HasWarehousePriority returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PimVariationStock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PimVariationStock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PimVariationStock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PimVariationStock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVariationId

`func (o *PimVariationStock) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *PimVariationStock) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *PimVariationStock) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *PimVariationStock) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


