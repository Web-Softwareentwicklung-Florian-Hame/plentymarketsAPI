# Variation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IsMain** | Pointer to **bool** |  | [optional] 
**MainVariationId** | Pointer to **int32** |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**CategoryVariationId** | Pointer to **int32** |  | [optional] 
**MarketVariationId** | Pointer to **int32** |  | [optional] 
**ClientVariationId** | Pointer to **int32** |  | [optional] 
**SalesPriceVariationId** | Pointer to **int32** |  | [optional] 
**SupplierVariationId** | Pointer to **int32** |  | [optional] 
**WarehouseVariationId** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**ParentVariationId** | Pointer to **int32** |  | [optional] 
**ParentVariationQuantity** | Pointer to **int32** |  | [optional] 
**Availability** | Pointer to **int32** |  | [optional] 
**EstimatedAvailableAt** | Pointer to **time.Time** |  | [optional] 
**PurchasePrice** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RelatedUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**PriceCalculationId** | Pointer to **int32** |  | [optional] 
**PriceCalculationUUID** | Pointer to **string** |  | [optional] 
**Picking** | Pointer to **string** |  | [optional] 
**StockLimitation** | Pointer to **int32** |  | [optional] 
**IsVisibleIfNetStockIsPositive** | Pointer to **bool** |  | [optional] 
**IsInvisibleIfNetStockIsNotPositive** | Pointer to **bool** |  | [optional] 
**IsAvailableIfNetStockIsPositive** | Pointer to **bool** |  | [optional] 
**IsUnavailableIfNetStockIsNotPositive** | Pointer to **bool** |  | [optional] 
**MainWarehouseId** | Pointer to **int32** |  | [optional] 
**MaximumOrderQuantity** | Pointer to **float32** |  | [optional] 
**MinimumOrderQuantity** | Pointer to **float32** |  | [optional] 
**IntervalOrderQuantity** | Pointer to **float32** |  | [optional] 
**AvailableUntil** | Pointer to **string** |  | [optional] 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**UnitCombinationId** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WeightG** | Pointer to **int32** |  | [optional] 
**WeightNetG** | Pointer to **int32** |  | [optional] 
**WidthMM** | Pointer to **int32** |  | [optional] 
**LengthMM** | Pointer to **int32** |  | [optional] 
**HeightMM** | Pointer to **int32** |  | [optional] 
**ExtraShippingCharge1** | Pointer to **float32** |  | [optional] 
**ExtraShippingCharge2** | Pointer to **float32** |  | [optional] 
**UnitsContained** | Pointer to **float32** |  | [optional] 
**PalletTypeId** | Pointer to **int32** |  | [optional] 
**PackingUnits** | Pointer to **float32** |  | [optional] 
**PackingUnitTypeId** | Pointer to **int32** |  | [optional] 
**TransportationCosts** | Pointer to **float32** |  | [optional] 
**StorageCosts** | Pointer to **float32** |  | [optional] 
**OperatingCosts** | Pointer to **float32** |  | [optional] 
**VatId** | Pointer to **int32** |  | [optional] 
**AutomaticClientVisibility** | Pointer to **int32** |  | [optional] 
**IsHiddenInCategoryList** | Pointer to **bool** |  | [optional] 
**DefaultShippingCosts** | Pointer to **float32** |  | [optional] 
**MayShowUnitPrice** | Pointer to **bool** |  | [optional] 
**MovingAveragePrice** | Pointer to **float32** |  | [optional] 
**PropertyVariationId** | Pointer to **string** |  | [optional] 
**AutomaticListVisibility** | Pointer to **int32** |  | [optional] 
**IsVisibleInListIfNetStockIsPositive** | Pointer to **bool** |  | [optional] 
**IsInvisibleInListIfNetStockIsNotPositive** | Pointer to **bool** |  | [optional] 
**SingleItemCount** | Pointer to **int32** |  | [optional] 
**AvailabilityUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TagVariationId** | Pointer to **string** |  | [optional] 
**HasCalculatedBundleWeight** | Pointer to **bool** |  | [optional] 
**HasCalculatedBundleNetWeight** | Pointer to **bool** |  | [optional] 
**HasCalculatedBundlePurchasePrice** | Pointer to **bool** |  | [optional] 
**HasCalculatedBundleMovingAveragePrice** | Pointer to **bool** |  | [optional] 
**CustomsTariffNumber** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to [**[]Property**](Property.md) |  | [optional] 
**VariationProperties** | Pointer to [**[]VariationProperty**](VariationProperty.md) |  | [optional] 
**VariationBarcodes** | Pointer to [**[]VariationBarcode**](VariationBarcode.md) |  | [optional] 
**VariationSalesPrices** | Pointer to [**[]VariationSalesPrice**](VariationSalesPrice.md) |  | [optional] 
**MarketItemNumbers** | Pointer to [**[]VariationMarketItemNumber**](VariationMarketItemNumber.md) |  | [optional] 
**VariationCategories** | Pointer to [**[]VariationCategory**](VariationCategory.md) |  | [optional] 
**VariationClients** | Pointer to [**[]VariationClient**](VariationClient.md) |  | [optional] 
**VariationMarkets** | Pointer to [**[]VariationMarket**](VariationMarket.md) |  | [optional] 
**VariationDefaultCategory** | Pointer to [**[]VariationDefaultCategory**](VariationDefaultCategory.md) |  | [optional] 
**VariationSuppliers** | Pointer to [**[]VariationSupplier**](VariationSupplier.md) |  | [optional] 
**Images** | Pointer to [**[]VariationImage**](VariationImage.md) |  | [optional] 
**ItemImages** | Pointer to [**[]VariationImage**](VariationImage.md) |  | [optional] 
**VariationAttributeValues** | Pointer to [**[]VariationAttributeValue**](VariationAttributeValue.md) |  | [optional] 
**VariationSkus** | Pointer to [**[]VariationSku**](VariationSku.md) |  | [optional] 
**Unit** | Pointer to [**VariationUnit**](VariationUnit.md) |  | [optional] 
**Parent** | Pointer to [**Variation**](Variation.md) |  | [optional] 
**Item** | Pointer to [**VariationItem**](VariationItem.md) |  | [optional] 
**Stock** | Pointer to [**[]VariationStock**](VariationStock.md) |  | [optional] 

## Methods

### NewVariation

`func NewVariation() *Variation`

NewVariation instantiates a new Variation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationWithDefaults

`func NewVariationWithDefaults() *Variation`

NewVariationWithDefaults instantiates a new Variation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Variation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsMain

`func (o *Variation) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *Variation) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *Variation) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.

### HasIsMain

`func (o *Variation) HasIsMain() bool`

HasIsMain returns a boolean if a field has been set.

### GetMainVariationId

`func (o *Variation) GetMainVariationId() int32`

GetMainVariationId returns the MainVariationId field if non-nil, zero value otherwise.

### GetMainVariationIdOk

`func (o *Variation) GetMainVariationIdOk() (*int32, bool)`

GetMainVariationIdOk returns a tuple with the MainVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainVariationId

`func (o *Variation) SetMainVariationId(v int32)`

SetMainVariationId sets MainVariationId field to given value.

### HasMainVariationId

`func (o *Variation) HasMainVariationId() bool`

HasMainVariationId returns a boolean if a field has been set.

### GetItemId

`func (o *Variation) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Variation) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Variation) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *Variation) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetCategoryVariationId

`func (o *Variation) GetCategoryVariationId() int32`

GetCategoryVariationId returns the CategoryVariationId field if non-nil, zero value otherwise.

### GetCategoryVariationIdOk

`func (o *Variation) GetCategoryVariationIdOk() (*int32, bool)`

GetCategoryVariationIdOk returns a tuple with the CategoryVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryVariationId

`func (o *Variation) SetCategoryVariationId(v int32)`

SetCategoryVariationId sets CategoryVariationId field to given value.

### HasCategoryVariationId

`func (o *Variation) HasCategoryVariationId() bool`

HasCategoryVariationId returns a boolean if a field has been set.

### GetMarketVariationId

`func (o *Variation) GetMarketVariationId() int32`

GetMarketVariationId returns the MarketVariationId field if non-nil, zero value otherwise.

### GetMarketVariationIdOk

`func (o *Variation) GetMarketVariationIdOk() (*int32, bool)`

GetMarketVariationIdOk returns a tuple with the MarketVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketVariationId

`func (o *Variation) SetMarketVariationId(v int32)`

SetMarketVariationId sets MarketVariationId field to given value.

### HasMarketVariationId

`func (o *Variation) HasMarketVariationId() bool`

HasMarketVariationId returns a boolean if a field has been set.

### GetClientVariationId

`func (o *Variation) GetClientVariationId() int32`

GetClientVariationId returns the ClientVariationId field if non-nil, zero value otherwise.

### GetClientVariationIdOk

`func (o *Variation) GetClientVariationIdOk() (*int32, bool)`

GetClientVariationIdOk returns a tuple with the ClientVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVariationId

`func (o *Variation) SetClientVariationId(v int32)`

SetClientVariationId sets ClientVariationId field to given value.

### HasClientVariationId

`func (o *Variation) HasClientVariationId() bool`

HasClientVariationId returns a boolean if a field has been set.

### GetSalesPriceVariationId

`func (o *Variation) GetSalesPriceVariationId() int32`

GetSalesPriceVariationId returns the SalesPriceVariationId field if non-nil, zero value otherwise.

### GetSalesPriceVariationIdOk

`func (o *Variation) GetSalesPriceVariationIdOk() (*int32, bool)`

GetSalesPriceVariationIdOk returns a tuple with the SalesPriceVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPriceVariationId

`func (o *Variation) SetSalesPriceVariationId(v int32)`

SetSalesPriceVariationId sets SalesPriceVariationId field to given value.

### HasSalesPriceVariationId

`func (o *Variation) HasSalesPriceVariationId() bool`

HasSalesPriceVariationId returns a boolean if a field has been set.

### GetSupplierVariationId

`func (o *Variation) GetSupplierVariationId() int32`

GetSupplierVariationId returns the SupplierVariationId field if non-nil, zero value otherwise.

### GetSupplierVariationIdOk

`func (o *Variation) GetSupplierVariationIdOk() (*int32, bool)`

GetSupplierVariationIdOk returns a tuple with the SupplierVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierVariationId

`func (o *Variation) SetSupplierVariationId(v int32)`

SetSupplierVariationId sets SupplierVariationId field to given value.

### HasSupplierVariationId

`func (o *Variation) HasSupplierVariationId() bool`

HasSupplierVariationId returns a boolean if a field has been set.

### GetWarehouseVariationId

`func (o *Variation) GetWarehouseVariationId() int32`

GetWarehouseVariationId returns the WarehouseVariationId field if non-nil, zero value otherwise.

### GetWarehouseVariationIdOk

`func (o *Variation) GetWarehouseVariationIdOk() (*int32, bool)`

GetWarehouseVariationIdOk returns a tuple with the WarehouseVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseVariationId

`func (o *Variation) SetWarehouseVariationId(v int32)`

SetWarehouseVariationId sets WarehouseVariationId field to given value.

### HasWarehouseVariationId

`func (o *Variation) HasWarehouseVariationId() bool`

HasWarehouseVariationId returns a boolean if a field has been set.

### GetPosition

`func (o *Variation) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Variation) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Variation) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Variation) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsActive

`func (o *Variation) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Variation) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Variation) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Variation) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNumber

`func (o *Variation) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Variation) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Variation) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Variation) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetModel

`func (o *Variation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Variation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Variation) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *Variation) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetExternalId

`func (o *Variation) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Variation) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Variation) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Variation) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetParentVariationId

`func (o *Variation) GetParentVariationId() int32`

GetParentVariationId returns the ParentVariationId field if non-nil, zero value otherwise.

### GetParentVariationIdOk

`func (o *Variation) GetParentVariationIdOk() (*int32, bool)`

GetParentVariationIdOk returns a tuple with the ParentVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVariationId

`func (o *Variation) SetParentVariationId(v int32)`

SetParentVariationId sets ParentVariationId field to given value.

### HasParentVariationId

`func (o *Variation) HasParentVariationId() bool`

HasParentVariationId returns a boolean if a field has been set.

### GetParentVariationQuantity

`func (o *Variation) GetParentVariationQuantity() int32`

GetParentVariationQuantity returns the ParentVariationQuantity field if non-nil, zero value otherwise.

### GetParentVariationQuantityOk

`func (o *Variation) GetParentVariationQuantityOk() (*int32, bool)`

GetParentVariationQuantityOk returns a tuple with the ParentVariationQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVariationQuantity

`func (o *Variation) SetParentVariationQuantity(v int32)`

SetParentVariationQuantity sets ParentVariationQuantity field to given value.

### HasParentVariationQuantity

`func (o *Variation) HasParentVariationQuantity() bool`

HasParentVariationQuantity returns a boolean if a field has been set.

### GetAvailability

`func (o *Variation) GetAvailability() int32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Variation) GetAvailabilityOk() (*int32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Variation) SetAvailability(v int32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Variation) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetEstimatedAvailableAt

`func (o *Variation) GetEstimatedAvailableAt() time.Time`

GetEstimatedAvailableAt returns the EstimatedAvailableAt field if non-nil, zero value otherwise.

### GetEstimatedAvailableAtOk

`func (o *Variation) GetEstimatedAvailableAtOk() (*time.Time, bool)`

GetEstimatedAvailableAtOk returns a tuple with the EstimatedAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAvailableAt

`func (o *Variation) SetEstimatedAvailableAt(v time.Time)`

SetEstimatedAvailableAt sets EstimatedAvailableAt field to given value.

### HasEstimatedAvailableAt

`func (o *Variation) HasEstimatedAvailableAt() bool`

HasEstimatedAvailableAt returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *Variation) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *Variation) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *Variation) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *Variation) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Variation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Variation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Variation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Variation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Variation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Variation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Variation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Variation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRelatedUpdatedAt

`func (o *Variation) GetRelatedUpdatedAt() time.Time`

GetRelatedUpdatedAt returns the RelatedUpdatedAt field if non-nil, zero value otherwise.

### GetRelatedUpdatedAtOk

`func (o *Variation) GetRelatedUpdatedAtOk() (*time.Time, bool)`

GetRelatedUpdatedAtOk returns a tuple with the RelatedUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUpdatedAt

`func (o *Variation) SetRelatedUpdatedAt(v time.Time)`

SetRelatedUpdatedAt sets RelatedUpdatedAt field to given value.

### HasRelatedUpdatedAt

`func (o *Variation) HasRelatedUpdatedAt() bool`

HasRelatedUpdatedAt returns a boolean if a field has been set.

### GetPriceCalculationId

`func (o *Variation) GetPriceCalculationId() int32`

GetPriceCalculationId returns the PriceCalculationId field if non-nil, zero value otherwise.

### GetPriceCalculationIdOk

`func (o *Variation) GetPriceCalculationIdOk() (*int32, bool)`

GetPriceCalculationIdOk returns a tuple with the PriceCalculationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCalculationId

`func (o *Variation) SetPriceCalculationId(v int32)`

SetPriceCalculationId sets PriceCalculationId field to given value.

### HasPriceCalculationId

`func (o *Variation) HasPriceCalculationId() bool`

HasPriceCalculationId returns a boolean if a field has been set.

### GetPriceCalculationUUID

`func (o *Variation) GetPriceCalculationUUID() string`

GetPriceCalculationUUID returns the PriceCalculationUUID field if non-nil, zero value otherwise.

### GetPriceCalculationUUIDOk

`func (o *Variation) GetPriceCalculationUUIDOk() (*string, bool)`

GetPriceCalculationUUIDOk returns a tuple with the PriceCalculationUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCalculationUUID

`func (o *Variation) SetPriceCalculationUUID(v string)`

SetPriceCalculationUUID sets PriceCalculationUUID field to given value.

### HasPriceCalculationUUID

`func (o *Variation) HasPriceCalculationUUID() bool`

HasPriceCalculationUUID returns a boolean if a field has been set.

### GetPicking

`func (o *Variation) GetPicking() string`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *Variation) GetPickingOk() (*string, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *Variation) SetPicking(v string)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *Variation) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetStockLimitation

`func (o *Variation) GetStockLimitation() int32`

GetStockLimitation returns the StockLimitation field if non-nil, zero value otherwise.

### GetStockLimitationOk

`func (o *Variation) GetStockLimitationOk() (*int32, bool)`

GetStockLimitationOk returns a tuple with the StockLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLimitation

`func (o *Variation) SetStockLimitation(v int32)`

SetStockLimitation sets StockLimitation field to given value.

### HasStockLimitation

`func (o *Variation) HasStockLimitation() bool`

HasStockLimitation returns a boolean if a field has been set.

### GetIsVisibleIfNetStockIsPositive

`func (o *Variation) GetIsVisibleIfNetStockIsPositive() bool`

GetIsVisibleIfNetStockIsPositive returns the IsVisibleIfNetStockIsPositive field if non-nil, zero value otherwise.

### GetIsVisibleIfNetStockIsPositiveOk

`func (o *Variation) GetIsVisibleIfNetStockIsPositiveOk() (*bool, bool)`

GetIsVisibleIfNetStockIsPositiveOk returns a tuple with the IsVisibleIfNetStockIsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisibleIfNetStockIsPositive

`func (o *Variation) SetIsVisibleIfNetStockIsPositive(v bool)`

SetIsVisibleIfNetStockIsPositive sets IsVisibleIfNetStockIsPositive field to given value.

### HasIsVisibleIfNetStockIsPositive

`func (o *Variation) HasIsVisibleIfNetStockIsPositive() bool`

HasIsVisibleIfNetStockIsPositive returns a boolean if a field has been set.

### GetIsInvisibleIfNetStockIsNotPositive

`func (o *Variation) GetIsInvisibleIfNetStockIsNotPositive() bool`

GetIsInvisibleIfNetStockIsNotPositive returns the IsInvisibleIfNetStockIsNotPositive field if non-nil, zero value otherwise.

### GetIsInvisibleIfNetStockIsNotPositiveOk

`func (o *Variation) GetIsInvisibleIfNetStockIsNotPositiveOk() (*bool, bool)`

GetIsInvisibleIfNetStockIsNotPositiveOk returns a tuple with the IsInvisibleIfNetStockIsNotPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvisibleIfNetStockIsNotPositive

`func (o *Variation) SetIsInvisibleIfNetStockIsNotPositive(v bool)`

SetIsInvisibleIfNetStockIsNotPositive sets IsInvisibleIfNetStockIsNotPositive field to given value.

### HasIsInvisibleIfNetStockIsNotPositive

`func (o *Variation) HasIsInvisibleIfNetStockIsNotPositive() bool`

HasIsInvisibleIfNetStockIsNotPositive returns a boolean if a field has been set.

### GetIsAvailableIfNetStockIsPositive

`func (o *Variation) GetIsAvailableIfNetStockIsPositive() bool`

GetIsAvailableIfNetStockIsPositive returns the IsAvailableIfNetStockIsPositive field if non-nil, zero value otherwise.

### GetIsAvailableIfNetStockIsPositiveOk

`func (o *Variation) GetIsAvailableIfNetStockIsPositiveOk() (*bool, bool)`

GetIsAvailableIfNetStockIsPositiveOk returns a tuple with the IsAvailableIfNetStockIsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableIfNetStockIsPositive

`func (o *Variation) SetIsAvailableIfNetStockIsPositive(v bool)`

SetIsAvailableIfNetStockIsPositive sets IsAvailableIfNetStockIsPositive field to given value.

### HasIsAvailableIfNetStockIsPositive

`func (o *Variation) HasIsAvailableIfNetStockIsPositive() bool`

HasIsAvailableIfNetStockIsPositive returns a boolean if a field has been set.

### GetIsUnavailableIfNetStockIsNotPositive

`func (o *Variation) GetIsUnavailableIfNetStockIsNotPositive() bool`

GetIsUnavailableIfNetStockIsNotPositive returns the IsUnavailableIfNetStockIsNotPositive field if non-nil, zero value otherwise.

### GetIsUnavailableIfNetStockIsNotPositiveOk

`func (o *Variation) GetIsUnavailableIfNetStockIsNotPositiveOk() (*bool, bool)`

GetIsUnavailableIfNetStockIsNotPositiveOk returns a tuple with the IsUnavailableIfNetStockIsNotPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnavailableIfNetStockIsNotPositive

`func (o *Variation) SetIsUnavailableIfNetStockIsNotPositive(v bool)`

SetIsUnavailableIfNetStockIsNotPositive sets IsUnavailableIfNetStockIsNotPositive field to given value.

### HasIsUnavailableIfNetStockIsNotPositive

`func (o *Variation) HasIsUnavailableIfNetStockIsNotPositive() bool`

HasIsUnavailableIfNetStockIsNotPositive returns a boolean if a field has been set.

### GetMainWarehouseId

`func (o *Variation) GetMainWarehouseId() int32`

GetMainWarehouseId returns the MainWarehouseId field if non-nil, zero value otherwise.

### GetMainWarehouseIdOk

`func (o *Variation) GetMainWarehouseIdOk() (*int32, bool)`

GetMainWarehouseIdOk returns a tuple with the MainWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWarehouseId

`func (o *Variation) SetMainWarehouseId(v int32)`

SetMainWarehouseId sets MainWarehouseId field to given value.

### HasMainWarehouseId

`func (o *Variation) HasMainWarehouseId() bool`

HasMainWarehouseId returns a boolean if a field has been set.

### GetMaximumOrderQuantity

`func (o *Variation) GetMaximumOrderQuantity() float32`

GetMaximumOrderQuantity returns the MaximumOrderQuantity field if non-nil, zero value otherwise.

### GetMaximumOrderQuantityOk

`func (o *Variation) GetMaximumOrderQuantityOk() (*float32, bool)`

GetMaximumOrderQuantityOk returns a tuple with the MaximumOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOrderQuantity

`func (o *Variation) SetMaximumOrderQuantity(v float32)`

SetMaximumOrderQuantity sets MaximumOrderQuantity field to given value.

### HasMaximumOrderQuantity

`func (o *Variation) HasMaximumOrderQuantity() bool`

HasMaximumOrderQuantity returns a boolean if a field has been set.

### GetMinimumOrderQuantity

`func (o *Variation) GetMinimumOrderQuantity() float32`

GetMinimumOrderQuantity returns the MinimumOrderQuantity field if non-nil, zero value otherwise.

### GetMinimumOrderQuantityOk

`func (o *Variation) GetMinimumOrderQuantityOk() (*float32, bool)`

GetMinimumOrderQuantityOk returns a tuple with the MinimumOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOrderQuantity

`func (o *Variation) SetMinimumOrderQuantity(v float32)`

SetMinimumOrderQuantity sets MinimumOrderQuantity field to given value.

### HasMinimumOrderQuantity

`func (o *Variation) HasMinimumOrderQuantity() bool`

HasMinimumOrderQuantity returns a boolean if a field has been set.

### GetIntervalOrderQuantity

`func (o *Variation) GetIntervalOrderQuantity() float32`

GetIntervalOrderQuantity returns the IntervalOrderQuantity field if non-nil, zero value otherwise.

### GetIntervalOrderQuantityOk

`func (o *Variation) GetIntervalOrderQuantityOk() (*float32, bool)`

GetIntervalOrderQuantityOk returns a tuple with the IntervalOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalOrderQuantity

`func (o *Variation) SetIntervalOrderQuantity(v float32)`

SetIntervalOrderQuantity sets IntervalOrderQuantity field to given value.

### HasIntervalOrderQuantity

`func (o *Variation) HasIntervalOrderQuantity() bool`

HasIntervalOrderQuantity returns a boolean if a field has been set.

### GetAvailableUntil

`func (o *Variation) GetAvailableUntil() string`

GetAvailableUntil returns the AvailableUntil field if non-nil, zero value otherwise.

### GetAvailableUntilOk

`func (o *Variation) GetAvailableUntilOk() (*string, bool)`

GetAvailableUntilOk returns a tuple with the AvailableUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUntil

`func (o *Variation) SetAvailableUntil(v string)`

SetAvailableUntil sets AvailableUntil field to given value.

### HasAvailableUntil

`func (o *Variation) HasAvailableUntil() bool`

HasAvailableUntil returns a boolean if a field has been set.

### GetReleasedAt

`func (o *Variation) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *Variation) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *Variation) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *Variation) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetUnitCombinationId

`func (o *Variation) GetUnitCombinationId() float32`

GetUnitCombinationId returns the UnitCombinationId field if non-nil, zero value otherwise.

### GetUnitCombinationIdOk

`func (o *Variation) GetUnitCombinationIdOk() (*float32, bool)`

GetUnitCombinationIdOk returns a tuple with the UnitCombinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCombinationId

`func (o *Variation) SetUnitCombinationId(v float32)`

SetUnitCombinationId sets UnitCombinationId field to given value.

### HasUnitCombinationId

`func (o *Variation) HasUnitCombinationId() bool`

HasUnitCombinationId returns a boolean if a field has been set.

### GetName

`func (o *Variation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Variation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeightG

`func (o *Variation) GetWeightG() int32`

GetWeightG returns the WeightG field if non-nil, zero value otherwise.

### GetWeightGOk

`func (o *Variation) GetWeightGOk() (*int32, bool)`

GetWeightGOk returns a tuple with the WeightG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightG

`func (o *Variation) SetWeightG(v int32)`

SetWeightG sets WeightG field to given value.

### HasWeightG

`func (o *Variation) HasWeightG() bool`

HasWeightG returns a boolean if a field has been set.

### GetWeightNetG

`func (o *Variation) GetWeightNetG() int32`

GetWeightNetG returns the WeightNetG field if non-nil, zero value otherwise.

### GetWeightNetGOk

`func (o *Variation) GetWeightNetGOk() (*int32, bool)`

GetWeightNetGOk returns a tuple with the WeightNetG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightNetG

`func (o *Variation) SetWeightNetG(v int32)`

SetWeightNetG sets WeightNetG field to given value.

### HasWeightNetG

`func (o *Variation) HasWeightNetG() bool`

HasWeightNetG returns a boolean if a field has been set.

### GetWidthMM

`func (o *Variation) GetWidthMM() int32`

GetWidthMM returns the WidthMM field if non-nil, zero value otherwise.

### GetWidthMMOk

`func (o *Variation) GetWidthMMOk() (*int32, bool)`

GetWidthMMOk returns a tuple with the WidthMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMM

`func (o *Variation) SetWidthMM(v int32)`

SetWidthMM sets WidthMM field to given value.

### HasWidthMM

`func (o *Variation) HasWidthMM() bool`

HasWidthMM returns a boolean if a field has been set.

### GetLengthMM

`func (o *Variation) GetLengthMM() int32`

GetLengthMM returns the LengthMM field if non-nil, zero value otherwise.

### GetLengthMMOk

`func (o *Variation) GetLengthMMOk() (*int32, bool)`

GetLengthMMOk returns a tuple with the LengthMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMM

`func (o *Variation) SetLengthMM(v int32)`

SetLengthMM sets LengthMM field to given value.

### HasLengthMM

`func (o *Variation) HasLengthMM() bool`

HasLengthMM returns a boolean if a field has been set.

### GetHeightMM

`func (o *Variation) GetHeightMM() int32`

GetHeightMM returns the HeightMM field if non-nil, zero value otherwise.

### GetHeightMMOk

`func (o *Variation) GetHeightMMOk() (*int32, bool)`

GetHeightMMOk returns a tuple with the HeightMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMM

`func (o *Variation) SetHeightMM(v int32)`

SetHeightMM sets HeightMM field to given value.

### HasHeightMM

`func (o *Variation) HasHeightMM() bool`

HasHeightMM returns a boolean if a field has been set.

### GetExtraShippingCharge1

`func (o *Variation) GetExtraShippingCharge1() float32`

GetExtraShippingCharge1 returns the ExtraShippingCharge1 field if non-nil, zero value otherwise.

### GetExtraShippingCharge1Ok

`func (o *Variation) GetExtraShippingCharge1Ok() (*float32, bool)`

GetExtraShippingCharge1Ok returns a tuple with the ExtraShippingCharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraShippingCharge1

`func (o *Variation) SetExtraShippingCharge1(v float32)`

SetExtraShippingCharge1 sets ExtraShippingCharge1 field to given value.

### HasExtraShippingCharge1

`func (o *Variation) HasExtraShippingCharge1() bool`

HasExtraShippingCharge1 returns a boolean if a field has been set.

### GetExtraShippingCharge2

`func (o *Variation) GetExtraShippingCharge2() float32`

GetExtraShippingCharge2 returns the ExtraShippingCharge2 field if non-nil, zero value otherwise.

### GetExtraShippingCharge2Ok

`func (o *Variation) GetExtraShippingCharge2Ok() (*float32, bool)`

GetExtraShippingCharge2Ok returns a tuple with the ExtraShippingCharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraShippingCharge2

`func (o *Variation) SetExtraShippingCharge2(v float32)`

SetExtraShippingCharge2 sets ExtraShippingCharge2 field to given value.

### HasExtraShippingCharge2

`func (o *Variation) HasExtraShippingCharge2() bool`

HasExtraShippingCharge2 returns a boolean if a field has been set.

### GetUnitsContained

`func (o *Variation) GetUnitsContained() float32`

GetUnitsContained returns the UnitsContained field if non-nil, zero value otherwise.

### GetUnitsContainedOk

`func (o *Variation) GetUnitsContainedOk() (*float32, bool)`

GetUnitsContainedOk returns a tuple with the UnitsContained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsContained

`func (o *Variation) SetUnitsContained(v float32)`

SetUnitsContained sets UnitsContained field to given value.

### HasUnitsContained

`func (o *Variation) HasUnitsContained() bool`

HasUnitsContained returns a boolean if a field has been set.

### GetPalletTypeId

`func (o *Variation) GetPalletTypeId() int32`

GetPalletTypeId returns the PalletTypeId field if non-nil, zero value otherwise.

### GetPalletTypeIdOk

`func (o *Variation) GetPalletTypeIdOk() (*int32, bool)`

GetPalletTypeIdOk returns a tuple with the PalletTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalletTypeId

`func (o *Variation) SetPalletTypeId(v int32)`

SetPalletTypeId sets PalletTypeId field to given value.

### HasPalletTypeId

`func (o *Variation) HasPalletTypeId() bool`

HasPalletTypeId returns a boolean if a field has been set.

### GetPackingUnits

`func (o *Variation) GetPackingUnits() float32`

GetPackingUnits returns the PackingUnits field if non-nil, zero value otherwise.

### GetPackingUnitsOk

`func (o *Variation) GetPackingUnitsOk() (*float32, bool)`

GetPackingUnitsOk returns a tuple with the PackingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingUnits

`func (o *Variation) SetPackingUnits(v float32)`

SetPackingUnits sets PackingUnits field to given value.

### HasPackingUnits

`func (o *Variation) HasPackingUnits() bool`

HasPackingUnits returns a boolean if a field has been set.

### GetPackingUnitTypeId

`func (o *Variation) GetPackingUnitTypeId() int32`

GetPackingUnitTypeId returns the PackingUnitTypeId field if non-nil, zero value otherwise.

### GetPackingUnitTypeIdOk

`func (o *Variation) GetPackingUnitTypeIdOk() (*int32, bool)`

GetPackingUnitTypeIdOk returns a tuple with the PackingUnitTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingUnitTypeId

`func (o *Variation) SetPackingUnitTypeId(v int32)`

SetPackingUnitTypeId sets PackingUnitTypeId field to given value.

### HasPackingUnitTypeId

`func (o *Variation) HasPackingUnitTypeId() bool`

HasPackingUnitTypeId returns a boolean if a field has been set.

### GetTransportationCosts

`func (o *Variation) GetTransportationCosts() float32`

GetTransportationCosts returns the TransportationCosts field if non-nil, zero value otherwise.

### GetTransportationCostsOk

`func (o *Variation) GetTransportationCostsOk() (*float32, bool)`

GetTransportationCostsOk returns a tuple with the TransportationCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportationCosts

`func (o *Variation) SetTransportationCosts(v float32)`

SetTransportationCosts sets TransportationCosts field to given value.

### HasTransportationCosts

`func (o *Variation) HasTransportationCosts() bool`

HasTransportationCosts returns a boolean if a field has been set.

### GetStorageCosts

`func (o *Variation) GetStorageCosts() float32`

GetStorageCosts returns the StorageCosts field if non-nil, zero value otherwise.

### GetStorageCostsOk

`func (o *Variation) GetStorageCostsOk() (*float32, bool)`

GetStorageCostsOk returns a tuple with the StorageCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCosts

`func (o *Variation) SetStorageCosts(v float32)`

SetStorageCosts sets StorageCosts field to given value.

### HasStorageCosts

`func (o *Variation) HasStorageCosts() bool`

HasStorageCosts returns a boolean if a field has been set.

### GetOperatingCosts

`func (o *Variation) GetOperatingCosts() float32`

GetOperatingCosts returns the OperatingCosts field if non-nil, zero value otherwise.

### GetOperatingCostsOk

`func (o *Variation) GetOperatingCostsOk() (*float32, bool)`

GetOperatingCostsOk returns a tuple with the OperatingCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingCosts

`func (o *Variation) SetOperatingCosts(v float32)`

SetOperatingCosts sets OperatingCosts field to given value.

### HasOperatingCosts

`func (o *Variation) HasOperatingCosts() bool`

HasOperatingCosts returns a boolean if a field has been set.

### GetVatId

`func (o *Variation) GetVatId() int32`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *Variation) GetVatIdOk() (*int32, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *Variation) SetVatId(v int32)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *Variation) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetAutomaticClientVisibility

`func (o *Variation) GetAutomaticClientVisibility() int32`

GetAutomaticClientVisibility returns the AutomaticClientVisibility field if non-nil, zero value otherwise.

### GetAutomaticClientVisibilityOk

`func (o *Variation) GetAutomaticClientVisibilityOk() (*int32, bool)`

GetAutomaticClientVisibilityOk returns a tuple with the AutomaticClientVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticClientVisibility

`func (o *Variation) SetAutomaticClientVisibility(v int32)`

SetAutomaticClientVisibility sets AutomaticClientVisibility field to given value.

### HasAutomaticClientVisibility

`func (o *Variation) HasAutomaticClientVisibility() bool`

HasAutomaticClientVisibility returns a boolean if a field has been set.

### GetIsHiddenInCategoryList

`func (o *Variation) GetIsHiddenInCategoryList() bool`

GetIsHiddenInCategoryList returns the IsHiddenInCategoryList field if non-nil, zero value otherwise.

### GetIsHiddenInCategoryListOk

`func (o *Variation) GetIsHiddenInCategoryListOk() (*bool, bool)`

GetIsHiddenInCategoryListOk returns a tuple with the IsHiddenInCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenInCategoryList

`func (o *Variation) SetIsHiddenInCategoryList(v bool)`

SetIsHiddenInCategoryList sets IsHiddenInCategoryList field to given value.

### HasIsHiddenInCategoryList

`func (o *Variation) HasIsHiddenInCategoryList() bool`

HasIsHiddenInCategoryList returns a boolean if a field has been set.

### GetDefaultShippingCosts

`func (o *Variation) GetDefaultShippingCosts() float32`

GetDefaultShippingCosts returns the DefaultShippingCosts field if non-nil, zero value otherwise.

### GetDefaultShippingCostsOk

`func (o *Variation) GetDefaultShippingCostsOk() (*float32, bool)`

GetDefaultShippingCostsOk returns a tuple with the DefaultShippingCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShippingCosts

`func (o *Variation) SetDefaultShippingCosts(v float32)`

SetDefaultShippingCosts sets DefaultShippingCosts field to given value.

### HasDefaultShippingCosts

`func (o *Variation) HasDefaultShippingCosts() bool`

HasDefaultShippingCosts returns a boolean if a field has been set.

### GetMayShowUnitPrice

`func (o *Variation) GetMayShowUnitPrice() bool`

GetMayShowUnitPrice returns the MayShowUnitPrice field if non-nil, zero value otherwise.

### GetMayShowUnitPriceOk

`func (o *Variation) GetMayShowUnitPriceOk() (*bool, bool)`

GetMayShowUnitPriceOk returns a tuple with the MayShowUnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayShowUnitPrice

`func (o *Variation) SetMayShowUnitPrice(v bool)`

SetMayShowUnitPrice sets MayShowUnitPrice field to given value.

### HasMayShowUnitPrice

`func (o *Variation) HasMayShowUnitPrice() bool`

HasMayShowUnitPrice returns a boolean if a field has been set.

### GetMovingAveragePrice

`func (o *Variation) GetMovingAveragePrice() float32`

GetMovingAveragePrice returns the MovingAveragePrice field if non-nil, zero value otherwise.

### GetMovingAveragePriceOk

`func (o *Variation) GetMovingAveragePriceOk() (*float32, bool)`

GetMovingAveragePriceOk returns a tuple with the MovingAveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovingAveragePrice

`func (o *Variation) SetMovingAveragePrice(v float32)`

SetMovingAveragePrice sets MovingAveragePrice field to given value.

### HasMovingAveragePrice

`func (o *Variation) HasMovingAveragePrice() bool`

HasMovingAveragePrice returns a boolean if a field has been set.

### GetPropertyVariationId

`func (o *Variation) GetPropertyVariationId() string`

GetPropertyVariationId returns the PropertyVariationId field if non-nil, zero value otherwise.

### GetPropertyVariationIdOk

`func (o *Variation) GetPropertyVariationIdOk() (*string, bool)`

GetPropertyVariationIdOk returns a tuple with the PropertyVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyVariationId

`func (o *Variation) SetPropertyVariationId(v string)`

SetPropertyVariationId sets PropertyVariationId field to given value.

### HasPropertyVariationId

`func (o *Variation) HasPropertyVariationId() bool`

HasPropertyVariationId returns a boolean if a field has been set.

### GetAutomaticListVisibility

`func (o *Variation) GetAutomaticListVisibility() int32`

GetAutomaticListVisibility returns the AutomaticListVisibility field if non-nil, zero value otherwise.

### GetAutomaticListVisibilityOk

`func (o *Variation) GetAutomaticListVisibilityOk() (*int32, bool)`

GetAutomaticListVisibilityOk returns a tuple with the AutomaticListVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticListVisibility

`func (o *Variation) SetAutomaticListVisibility(v int32)`

SetAutomaticListVisibility sets AutomaticListVisibility field to given value.

### HasAutomaticListVisibility

`func (o *Variation) HasAutomaticListVisibility() bool`

HasAutomaticListVisibility returns a boolean if a field has been set.

### GetIsVisibleInListIfNetStockIsPositive

`func (o *Variation) GetIsVisibleInListIfNetStockIsPositive() bool`

GetIsVisibleInListIfNetStockIsPositive returns the IsVisibleInListIfNetStockIsPositive field if non-nil, zero value otherwise.

### GetIsVisibleInListIfNetStockIsPositiveOk

`func (o *Variation) GetIsVisibleInListIfNetStockIsPositiveOk() (*bool, bool)`

GetIsVisibleInListIfNetStockIsPositiveOk returns a tuple with the IsVisibleInListIfNetStockIsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisibleInListIfNetStockIsPositive

`func (o *Variation) SetIsVisibleInListIfNetStockIsPositive(v bool)`

SetIsVisibleInListIfNetStockIsPositive sets IsVisibleInListIfNetStockIsPositive field to given value.

### HasIsVisibleInListIfNetStockIsPositive

`func (o *Variation) HasIsVisibleInListIfNetStockIsPositive() bool`

HasIsVisibleInListIfNetStockIsPositive returns a boolean if a field has been set.

### GetIsInvisibleInListIfNetStockIsNotPositive

`func (o *Variation) GetIsInvisibleInListIfNetStockIsNotPositive() bool`

GetIsInvisibleInListIfNetStockIsNotPositive returns the IsInvisibleInListIfNetStockIsNotPositive field if non-nil, zero value otherwise.

### GetIsInvisibleInListIfNetStockIsNotPositiveOk

`func (o *Variation) GetIsInvisibleInListIfNetStockIsNotPositiveOk() (*bool, bool)`

GetIsInvisibleInListIfNetStockIsNotPositiveOk returns a tuple with the IsInvisibleInListIfNetStockIsNotPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvisibleInListIfNetStockIsNotPositive

`func (o *Variation) SetIsInvisibleInListIfNetStockIsNotPositive(v bool)`

SetIsInvisibleInListIfNetStockIsNotPositive sets IsInvisibleInListIfNetStockIsNotPositive field to given value.

### HasIsInvisibleInListIfNetStockIsNotPositive

`func (o *Variation) HasIsInvisibleInListIfNetStockIsNotPositive() bool`

HasIsInvisibleInListIfNetStockIsNotPositive returns a boolean if a field has been set.

### GetSingleItemCount

`func (o *Variation) GetSingleItemCount() int32`

GetSingleItemCount returns the SingleItemCount field if non-nil, zero value otherwise.

### GetSingleItemCountOk

`func (o *Variation) GetSingleItemCountOk() (*int32, bool)`

GetSingleItemCountOk returns a tuple with the SingleItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleItemCount

`func (o *Variation) SetSingleItemCount(v int32)`

SetSingleItemCount sets SingleItemCount field to given value.

### HasSingleItemCount

`func (o *Variation) HasSingleItemCount() bool`

HasSingleItemCount returns a boolean if a field has been set.

### GetAvailabilityUpdatedAt

`func (o *Variation) GetAvailabilityUpdatedAt() time.Time`

GetAvailabilityUpdatedAt returns the AvailabilityUpdatedAt field if non-nil, zero value otherwise.

### GetAvailabilityUpdatedAtOk

`func (o *Variation) GetAvailabilityUpdatedAtOk() (*time.Time, bool)`

GetAvailabilityUpdatedAtOk returns a tuple with the AvailabilityUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityUpdatedAt

`func (o *Variation) SetAvailabilityUpdatedAt(v time.Time)`

SetAvailabilityUpdatedAt sets AvailabilityUpdatedAt field to given value.

### HasAvailabilityUpdatedAt

`func (o *Variation) HasAvailabilityUpdatedAt() bool`

HasAvailabilityUpdatedAt returns a boolean if a field has been set.

### GetTagVariationId

`func (o *Variation) GetTagVariationId() string`

GetTagVariationId returns the TagVariationId field if non-nil, zero value otherwise.

### GetTagVariationIdOk

`func (o *Variation) GetTagVariationIdOk() (*string, bool)`

GetTagVariationIdOk returns a tuple with the TagVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagVariationId

`func (o *Variation) SetTagVariationId(v string)`

SetTagVariationId sets TagVariationId field to given value.

### HasTagVariationId

`func (o *Variation) HasTagVariationId() bool`

HasTagVariationId returns a boolean if a field has been set.

### GetHasCalculatedBundleWeight

`func (o *Variation) GetHasCalculatedBundleWeight() bool`

GetHasCalculatedBundleWeight returns the HasCalculatedBundleWeight field if non-nil, zero value otherwise.

### GetHasCalculatedBundleWeightOk

`func (o *Variation) GetHasCalculatedBundleWeightOk() (*bool, bool)`

GetHasCalculatedBundleWeightOk returns a tuple with the HasCalculatedBundleWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCalculatedBundleWeight

`func (o *Variation) SetHasCalculatedBundleWeight(v bool)`

SetHasCalculatedBundleWeight sets HasCalculatedBundleWeight field to given value.

### HasHasCalculatedBundleWeight

`func (o *Variation) HasHasCalculatedBundleWeight() bool`

HasHasCalculatedBundleWeight returns a boolean if a field has been set.

### GetHasCalculatedBundleNetWeight

`func (o *Variation) GetHasCalculatedBundleNetWeight() bool`

GetHasCalculatedBundleNetWeight returns the HasCalculatedBundleNetWeight field if non-nil, zero value otherwise.

### GetHasCalculatedBundleNetWeightOk

`func (o *Variation) GetHasCalculatedBundleNetWeightOk() (*bool, bool)`

GetHasCalculatedBundleNetWeightOk returns a tuple with the HasCalculatedBundleNetWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCalculatedBundleNetWeight

`func (o *Variation) SetHasCalculatedBundleNetWeight(v bool)`

SetHasCalculatedBundleNetWeight sets HasCalculatedBundleNetWeight field to given value.

### HasHasCalculatedBundleNetWeight

`func (o *Variation) HasHasCalculatedBundleNetWeight() bool`

HasHasCalculatedBundleNetWeight returns a boolean if a field has been set.

### GetHasCalculatedBundlePurchasePrice

`func (o *Variation) GetHasCalculatedBundlePurchasePrice() bool`

GetHasCalculatedBundlePurchasePrice returns the HasCalculatedBundlePurchasePrice field if non-nil, zero value otherwise.

### GetHasCalculatedBundlePurchasePriceOk

`func (o *Variation) GetHasCalculatedBundlePurchasePriceOk() (*bool, bool)`

GetHasCalculatedBundlePurchasePriceOk returns a tuple with the HasCalculatedBundlePurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCalculatedBundlePurchasePrice

`func (o *Variation) SetHasCalculatedBundlePurchasePrice(v bool)`

SetHasCalculatedBundlePurchasePrice sets HasCalculatedBundlePurchasePrice field to given value.

### HasHasCalculatedBundlePurchasePrice

`func (o *Variation) HasHasCalculatedBundlePurchasePrice() bool`

HasHasCalculatedBundlePurchasePrice returns a boolean if a field has been set.

### GetHasCalculatedBundleMovingAveragePrice

`func (o *Variation) GetHasCalculatedBundleMovingAveragePrice() bool`

GetHasCalculatedBundleMovingAveragePrice returns the HasCalculatedBundleMovingAveragePrice field if non-nil, zero value otherwise.

### GetHasCalculatedBundleMovingAveragePriceOk

`func (o *Variation) GetHasCalculatedBundleMovingAveragePriceOk() (*bool, bool)`

GetHasCalculatedBundleMovingAveragePriceOk returns a tuple with the HasCalculatedBundleMovingAveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCalculatedBundleMovingAveragePrice

`func (o *Variation) SetHasCalculatedBundleMovingAveragePrice(v bool)`

SetHasCalculatedBundleMovingAveragePrice sets HasCalculatedBundleMovingAveragePrice field to given value.

### HasHasCalculatedBundleMovingAveragePrice

`func (o *Variation) HasHasCalculatedBundleMovingAveragePrice() bool`

HasHasCalculatedBundleMovingAveragePrice returns a boolean if a field has been set.

### GetCustomsTariffNumber

`func (o *Variation) GetCustomsTariffNumber() string`

GetCustomsTariffNumber returns the CustomsTariffNumber field if non-nil, zero value otherwise.

### GetCustomsTariffNumberOk

`func (o *Variation) GetCustomsTariffNumberOk() (*string, bool)`

GetCustomsTariffNumberOk returns a tuple with the CustomsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsTariffNumber

`func (o *Variation) SetCustomsTariffNumber(v string)`

SetCustomsTariffNumber sets CustomsTariffNumber field to given value.

### HasCustomsTariffNumber

`func (o *Variation) HasCustomsTariffNumber() bool`

HasCustomsTariffNumber returns a boolean if a field has been set.

### GetProperties

`func (o *Variation) GetProperties() []Property`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Variation) GetPropertiesOk() (*[]Property, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Variation) SetProperties(v []Property)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Variation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetVariationProperties

`func (o *Variation) GetVariationProperties() []VariationProperty`

GetVariationProperties returns the VariationProperties field if non-nil, zero value otherwise.

### GetVariationPropertiesOk

`func (o *Variation) GetVariationPropertiesOk() (*[]VariationProperty, bool)`

GetVariationPropertiesOk returns a tuple with the VariationProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationProperties

`func (o *Variation) SetVariationProperties(v []VariationProperty)`

SetVariationProperties sets VariationProperties field to given value.

### HasVariationProperties

`func (o *Variation) HasVariationProperties() bool`

HasVariationProperties returns a boolean if a field has been set.

### GetVariationBarcodes

`func (o *Variation) GetVariationBarcodes() []VariationBarcode`

GetVariationBarcodes returns the VariationBarcodes field if non-nil, zero value otherwise.

### GetVariationBarcodesOk

`func (o *Variation) GetVariationBarcodesOk() (*[]VariationBarcode, bool)`

GetVariationBarcodesOk returns a tuple with the VariationBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationBarcodes

`func (o *Variation) SetVariationBarcodes(v []VariationBarcode)`

SetVariationBarcodes sets VariationBarcodes field to given value.

### HasVariationBarcodes

`func (o *Variation) HasVariationBarcodes() bool`

HasVariationBarcodes returns a boolean if a field has been set.

### GetVariationSalesPrices

`func (o *Variation) GetVariationSalesPrices() []VariationSalesPrice`

GetVariationSalesPrices returns the VariationSalesPrices field if non-nil, zero value otherwise.

### GetVariationSalesPricesOk

`func (o *Variation) GetVariationSalesPricesOk() (*[]VariationSalesPrice, bool)`

GetVariationSalesPricesOk returns a tuple with the VariationSalesPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationSalesPrices

`func (o *Variation) SetVariationSalesPrices(v []VariationSalesPrice)`

SetVariationSalesPrices sets VariationSalesPrices field to given value.

### HasVariationSalesPrices

`func (o *Variation) HasVariationSalesPrices() bool`

HasVariationSalesPrices returns a boolean if a field has been set.

### GetMarketItemNumbers

`func (o *Variation) GetMarketItemNumbers() []VariationMarketItemNumber`

GetMarketItemNumbers returns the MarketItemNumbers field if non-nil, zero value otherwise.

### GetMarketItemNumbersOk

`func (o *Variation) GetMarketItemNumbersOk() (*[]VariationMarketItemNumber, bool)`

GetMarketItemNumbersOk returns a tuple with the MarketItemNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketItemNumbers

`func (o *Variation) SetMarketItemNumbers(v []VariationMarketItemNumber)`

SetMarketItemNumbers sets MarketItemNumbers field to given value.

### HasMarketItemNumbers

`func (o *Variation) HasMarketItemNumbers() bool`

HasMarketItemNumbers returns a boolean if a field has been set.

### GetVariationCategories

`func (o *Variation) GetVariationCategories() []VariationCategory`

GetVariationCategories returns the VariationCategories field if non-nil, zero value otherwise.

### GetVariationCategoriesOk

`func (o *Variation) GetVariationCategoriesOk() (*[]VariationCategory, bool)`

GetVariationCategoriesOk returns a tuple with the VariationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationCategories

`func (o *Variation) SetVariationCategories(v []VariationCategory)`

SetVariationCategories sets VariationCategories field to given value.

### HasVariationCategories

`func (o *Variation) HasVariationCategories() bool`

HasVariationCategories returns a boolean if a field has been set.

### GetVariationClients

`func (o *Variation) GetVariationClients() []VariationClient`

GetVariationClients returns the VariationClients field if non-nil, zero value otherwise.

### GetVariationClientsOk

`func (o *Variation) GetVariationClientsOk() (*[]VariationClient, bool)`

GetVariationClientsOk returns a tuple with the VariationClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationClients

`func (o *Variation) SetVariationClients(v []VariationClient)`

SetVariationClients sets VariationClients field to given value.

### HasVariationClients

`func (o *Variation) HasVariationClients() bool`

HasVariationClients returns a boolean if a field has been set.

### GetVariationMarkets

`func (o *Variation) GetVariationMarkets() []VariationMarket`

GetVariationMarkets returns the VariationMarkets field if non-nil, zero value otherwise.

### GetVariationMarketsOk

`func (o *Variation) GetVariationMarketsOk() (*[]VariationMarket, bool)`

GetVariationMarketsOk returns a tuple with the VariationMarkets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationMarkets

`func (o *Variation) SetVariationMarkets(v []VariationMarket)`

SetVariationMarkets sets VariationMarkets field to given value.

### HasVariationMarkets

`func (o *Variation) HasVariationMarkets() bool`

HasVariationMarkets returns a boolean if a field has been set.

### GetVariationDefaultCategory

`func (o *Variation) GetVariationDefaultCategory() []VariationDefaultCategory`

GetVariationDefaultCategory returns the VariationDefaultCategory field if non-nil, zero value otherwise.

### GetVariationDefaultCategoryOk

`func (o *Variation) GetVariationDefaultCategoryOk() (*[]VariationDefaultCategory, bool)`

GetVariationDefaultCategoryOk returns a tuple with the VariationDefaultCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationDefaultCategory

`func (o *Variation) SetVariationDefaultCategory(v []VariationDefaultCategory)`

SetVariationDefaultCategory sets VariationDefaultCategory field to given value.

### HasVariationDefaultCategory

`func (o *Variation) HasVariationDefaultCategory() bool`

HasVariationDefaultCategory returns a boolean if a field has been set.

### GetVariationSuppliers

`func (o *Variation) GetVariationSuppliers() []VariationSupplier`

GetVariationSuppliers returns the VariationSuppliers field if non-nil, zero value otherwise.

### GetVariationSuppliersOk

`func (o *Variation) GetVariationSuppliersOk() (*[]VariationSupplier, bool)`

GetVariationSuppliersOk returns a tuple with the VariationSuppliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationSuppliers

`func (o *Variation) SetVariationSuppliers(v []VariationSupplier)`

SetVariationSuppliers sets VariationSuppliers field to given value.

### HasVariationSuppliers

`func (o *Variation) HasVariationSuppliers() bool`

HasVariationSuppliers returns a boolean if a field has been set.

### GetImages

`func (o *Variation) GetImages() []VariationImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *Variation) GetImagesOk() (*[]VariationImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *Variation) SetImages(v []VariationImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *Variation) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetItemImages

`func (o *Variation) GetItemImages() []VariationImage`

GetItemImages returns the ItemImages field if non-nil, zero value otherwise.

### GetItemImagesOk

`func (o *Variation) GetItemImagesOk() (*[]VariationImage, bool)`

GetItemImagesOk returns a tuple with the ItemImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemImages

`func (o *Variation) SetItemImages(v []VariationImage)`

SetItemImages sets ItemImages field to given value.

### HasItemImages

`func (o *Variation) HasItemImages() bool`

HasItemImages returns a boolean if a field has been set.

### GetVariationAttributeValues

`func (o *Variation) GetVariationAttributeValues() []VariationAttributeValue`

GetVariationAttributeValues returns the VariationAttributeValues field if non-nil, zero value otherwise.

### GetVariationAttributeValuesOk

`func (o *Variation) GetVariationAttributeValuesOk() (*[]VariationAttributeValue, bool)`

GetVariationAttributeValuesOk returns a tuple with the VariationAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationAttributeValues

`func (o *Variation) SetVariationAttributeValues(v []VariationAttributeValue)`

SetVariationAttributeValues sets VariationAttributeValues field to given value.

### HasVariationAttributeValues

`func (o *Variation) HasVariationAttributeValues() bool`

HasVariationAttributeValues returns a boolean if a field has been set.

### GetVariationSkus

`func (o *Variation) GetVariationSkus() []VariationSku`

GetVariationSkus returns the VariationSkus field if non-nil, zero value otherwise.

### GetVariationSkusOk

`func (o *Variation) GetVariationSkusOk() (*[]VariationSku, bool)`

GetVariationSkusOk returns a tuple with the VariationSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationSkus

`func (o *Variation) SetVariationSkus(v []VariationSku)`

SetVariationSkus sets VariationSkus field to given value.

### HasVariationSkus

`func (o *Variation) HasVariationSkus() bool`

HasVariationSkus returns a boolean if a field has been set.

### GetUnit

`func (o *Variation) GetUnit() VariationUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Variation) GetUnitOk() (*VariationUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Variation) SetUnit(v VariationUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Variation) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetParent

`func (o *Variation) GetParent() Variation`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Variation) GetParentOk() (*Variation, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Variation) SetParent(v Variation)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Variation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetItem

`func (o *Variation) GetItem() VariationItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Variation) GetItemOk() (*VariationItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Variation) SetItem(v VariationItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *Variation) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetStock

`func (o *Variation) GetStock() []VariationStock`

GetStock returns the Stock field if non-nil, zero value otherwise.

### GetStockOk

`func (o *Variation) GetStockOk() (*[]VariationStock, bool)`

GetStockOk returns a tuple with the Stock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStock

`func (o *Variation) SetStock(v []VariationStock)`

SetStock sets Stock field to given value.

### HasStock

`func (o *Variation) HasStock() bool`

HasStock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


