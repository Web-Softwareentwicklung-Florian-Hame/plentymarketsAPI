# PimVariationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMain** | Pointer to **bool** |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Number** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to **int32** |  | [optional] 
**EstimatedAvailableAt** | Pointer to **time.Time** |  | [optional] 
**PurchasePrice** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RelatedUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**PriceCalculationId** | Pointer to **int32** |  | [optional] 
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
**AvailableUntil** | Pointer to **time.Time** |  | [optional] 
**ReleasedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**WeightG** | Pointer to **int32** |  | [optional] 
**WeightNetG** | Pointer to **int32** |  | [optional] 
**WidthMM** | Pointer to **float32** |  | [optional] 
**LengthMM** | Pointer to **float32** |  | [optional] 
**HeightMM** | Pointer to **float32** |  | [optional] 
**ExtraShippingCharge1** | Pointer to **float32** |  | [optional] 
**ExtraShippingCharge2** | Pointer to **float32** |  | [optional] 
**UnitsContained** | Pointer to **int32** |  | [optional] 
**PalletTypeId** | Pointer to **int32** |  | [optional] 
**PackingUnits** | Pointer to **int32** |  | [optional] 
**PackingUnitTypeId** | Pointer to **int32** |  | [optional] 
**TransportationCosts** | Pointer to **float32** |  | [optional] 
**StorageCosts** | Pointer to **float32** |  | [optional] 
**Customs** | Pointer to **float32** |  | [optional] 
**OperatingCosts** | Pointer to **float32** |  | [optional] 
**VatId** | Pointer to **int32** |  | [optional] 
**BundleType** | Pointer to **string** |  | [optional] 
**AutomaticClientVisibility** | Pointer to **int32** |  | [optional] 
**IsHiddenInCategoryList** | Pointer to **bool** |  | [optional] 

## Methods

### NewPimVariationBase

`func NewPimVariationBase() *PimVariationBase`

NewPimVariationBase instantiates a new PimVariationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimVariationBaseWithDefaults

`func NewPimVariationBaseWithDefaults() *PimVariationBase`

NewPimVariationBaseWithDefaults instantiates a new PimVariationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMain

`func (o *PimVariationBase) GetIsMain() bool`

GetIsMain returns the IsMain field if non-nil, zero value otherwise.

### GetIsMainOk

`func (o *PimVariationBase) GetIsMainOk() (*bool, bool)`

GetIsMainOk returns a tuple with the IsMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMain

`func (o *PimVariationBase) SetIsMain(v bool)`

SetIsMain sets IsMain field to given value.

### HasIsMain

`func (o *PimVariationBase) HasIsMain() bool`

HasIsMain returns a boolean if a field has been set.

### GetItemId

`func (o *PimVariationBase) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PimVariationBase) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PimVariationBase) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *PimVariationBase) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetPosition

`func (o *PimVariationBase) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PimVariationBase) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PimVariationBase) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PimVariationBase) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsActive

`func (o *PimVariationBase) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PimVariationBase) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PimVariationBase) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PimVariationBase) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNumber

`func (o *PimVariationBase) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PimVariationBase) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PimVariationBase) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PimVariationBase) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetModel

`func (o *PimVariationBase) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PimVariationBase) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PimVariationBase) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PimVariationBase) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetExternalId

`func (o *PimVariationBase) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PimVariationBase) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PimVariationBase) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PimVariationBase) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetAvailability

`func (o *PimVariationBase) GetAvailability() int32`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PimVariationBase) GetAvailabilityOk() (*int32, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PimVariationBase) SetAvailability(v int32)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PimVariationBase) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetEstimatedAvailableAt

`func (o *PimVariationBase) GetEstimatedAvailableAt() time.Time`

GetEstimatedAvailableAt returns the EstimatedAvailableAt field if non-nil, zero value otherwise.

### GetEstimatedAvailableAtOk

`func (o *PimVariationBase) GetEstimatedAvailableAtOk() (*time.Time, bool)`

GetEstimatedAvailableAtOk returns a tuple with the EstimatedAvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedAvailableAt

`func (o *PimVariationBase) SetEstimatedAvailableAt(v time.Time)`

SetEstimatedAvailableAt sets EstimatedAvailableAt field to given value.

### HasEstimatedAvailableAt

`func (o *PimVariationBase) HasEstimatedAvailableAt() bool`

HasEstimatedAvailableAt returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *PimVariationBase) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *PimVariationBase) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *PimVariationBase) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *PimVariationBase) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PimVariationBase) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PimVariationBase) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PimVariationBase) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PimVariationBase) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PimVariationBase) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PimVariationBase) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PimVariationBase) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PimVariationBase) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRelatedUpdatedAt

`func (o *PimVariationBase) GetRelatedUpdatedAt() time.Time`

GetRelatedUpdatedAt returns the RelatedUpdatedAt field if non-nil, zero value otherwise.

### GetRelatedUpdatedAtOk

`func (o *PimVariationBase) GetRelatedUpdatedAtOk() (*time.Time, bool)`

GetRelatedUpdatedAtOk returns a tuple with the RelatedUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedUpdatedAt

`func (o *PimVariationBase) SetRelatedUpdatedAt(v time.Time)`

SetRelatedUpdatedAt sets RelatedUpdatedAt field to given value.

### HasRelatedUpdatedAt

`func (o *PimVariationBase) HasRelatedUpdatedAt() bool`

HasRelatedUpdatedAt returns a boolean if a field has been set.

### GetPriceCalculationId

`func (o *PimVariationBase) GetPriceCalculationId() int32`

GetPriceCalculationId returns the PriceCalculationId field if non-nil, zero value otherwise.

### GetPriceCalculationIdOk

`func (o *PimVariationBase) GetPriceCalculationIdOk() (*int32, bool)`

GetPriceCalculationIdOk returns a tuple with the PriceCalculationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCalculationId

`func (o *PimVariationBase) SetPriceCalculationId(v int32)`

SetPriceCalculationId sets PriceCalculationId field to given value.

### HasPriceCalculationId

`func (o *PimVariationBase) HasPriceCalculationId() bool`

HasPriceCalculationId returns a boolean if a field has been set.

### GetPicking

`func (o *PimVariationBase) GetPicking() string`

GetPicking returns the Picking field if non-nil, zero value otherwise.

### GetPickingOk

`func (o *PimVariationBase) GetPickingOk() (*string, bool)`

GetPickingOk returns a tuple with the Picking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicking

`func (o *PimVariationBase) SetPicking(v string)`

SetPicking sets Picking field to given value.

### HasPicking

`func (o *PimVariationBase) HasPicking() bool`

HasPicking returns a boolean if a field has been set.

### GetStockLimitation

`func (o *PimVariationBase) GetStockLimitation() int32`

GetStockLimitation returns the StockLimitation field if non-nil, zero value otherwise.

### GetStockLimitationOk

`func (o *PimVariationBase) GetStockLimitationOk() (*int32, bool)`

GetStockLimitationOk returns a tuple with the StockLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLimitation

`func (o *PimVariationBase) SetStockLimitation(v int32)`

SetStockLimitation sets StockLimitation field to given value.

### HasStockLimitation

`func (o *PimVariationBase) HasStockLimitation() bool`

HasStockLimitation returns a boolean if a field has been set.

### GetIsVisibleIfNetStockIsPositive

`func (o *PimVariationBase) GetIsVisibleIfNetStockIsPositive() bool`

GetIsVisibleIfNetStockIsPositive returns the IsVisibleIfNetStockIsPositive field if non-nil, zero value otherwise.

### GetIsVisibleIfNetStockIsPositiveOk

`func (o *PimVariationBase) GetIsVisibleIfNetStockIsPositiveOk() (*bool, bool)`

GetIsVisibleIfNetStockIsPositiveOk returns a tuple with the IsVisibleIfNetStockIsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVisibleIfNetStockIsPositive

`func (o *PimVariationBase) SetIsVisibleIfNetStockIsPositive(v bool)`

SetIsVisibleIfNetStockIsPositive sets IsVisibleIfNetStockIsPositive field to given value.

### HasIsVisibleIfNetStockIsPositive

`func (o *PimVariationBase) HasIsVisibleIfNetStockIsPositive() bool`

HasIsVisibleIfNetStockIsPositive returns a boolean if a field has been set.

### GetIsInvisibleIfNetStockIsNotPositive

`func (o *PimVariationBase) GetIsInvisibleIfNetStockIsNotPositive() bool`

GetIsInvisibleIfNetStockIsNotPositive returns the IsInvisibleIfNetStockIsNotPositive field if non-nil, zero value otherwise.

### GetIsInvisibleIfNetStockIsNotPositiveOk

`func (o *PimVariationBase) GetIsInvisibleIfNetStockIsNotPositiveOk() (*bool, bool)`

GetIsInvisibleIfNetStockIsNotPositiveOk returns a tuple with the IsInvisibleIfNetStockIsNotPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvisibleIfNetStockIsNotPositive

`func (o *PimVariationBase) SetIsInvisibleIfNetStockIsNotPositive(v bool)`

SetIsInvisibleIfNetStockIsNotPositive sets IsInvisibleIfNetStockIsNotPositive field to given value.

### HasIsInvisibleIfNetStockIsNotPositive

`func (o *PimVariationBase) HasIsInvisibleIfNetStockIsNotPositive() bool`

HasIsInvisibleIfNetStockIsNotPositive returns a boolean if a field has been set.

### GetIsAvailableIfNetStockIsPositive

`func (o *PimVariationBase) GetIsAvailableIfNetStockIsPositive() bool`

GetIsAvailableIfNetStockIsPositive returns the IsAvailableIfNetStockIsPositive field if non-nil, zero value otherwise.

### GetIsAvailableIfNetStockIsPositiveOk

`func (o *PimVariationBase) GetIsAvailableIfNetStockIsPositiveOk() (*bool, bool)`

GetIsAvailableIfNetStockIsPositiveOk returns a tuple with the IsAvailableIfNetStockIsPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAvailableIfNetStockIsPositive

`func (o *PimVariationBase) SetIsAvailableIfNetStockIsPositive(v bool)`

SetIsAvailableIfNetStockIsPositive sets IsAvailableIfNetStockIsPositive field to given value.

### HasIsAvailableIfNetStockIsPositive

`func (o *PimVariationBase) HasIsAvailableIfNetStockIsPositive() bool`

HasIsAvailableIfNetStockIsPositive returns a boolean if a field has been set.

### GetIsUnavailableIfNetStockIsNotPositive

`func (o *PimVariationBase) GetIsUnavailableIfNetStockIsNotPositive() bool`

GetIsUnavailableIfNetStockIsNotPositive returns the IsUnavailableIfNetStockIsNotPositive field if non-nil, zero value otherwise.

### GetIsUnavailableIfNetStockIsNotPositiveOk

`func (o *PimVariationBase) GetIsUnavailableIfNetStockIsNotPositiveOk() (*bool, bool)`

GetIsUnavailableIfNetStockIsNotPositiveOk returns a tuple with the IsUnavailableIfNetStockIsNotPositive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnavailableIfNetStockIsNotPositive

`func (o *PimVariationBase) SetIsUnavailableIfNetStockIsNotPositive(v bool)`

SetIsUnavailableIfNetStockIsNotPositive sets IsUnavailableIfNetStockIsNotPositive field to given value.

### HasIsUnavailableIfNetStockIsNotPositive

`func (o *PimVariationBase) HasIsUnavailableIfNetStockIsNotPositive() bool`

HasIsUnavailableIfNetStockIsNotPositive returns a boolean if a field has been set.

### GetMainWarehouseId

`func (o *PimVariationBase) GetMainWarehouseId() int32`

GetMainWarehouseId returns the MainWarehouseId field if non-nil, zero value otherwise.

### GetMainWarehouseIdOk

`func (o *PimVariationBase) GetMainWarehouseIdOk() (*int32, bool)`

GetMainWarehouseIdOk returns a tuple with the MainWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainWarehouseId

`func (o *PimVariationBase) SetMainWarehouseId(v int32)`

SetMainWarehouseId sets MainWarehouseId field to given value.

### HasMainWarehouseId

`func (o *PimVariationBase) HasMainWarehouseId() bool`

HasMainWarehouseId returns a boolean if a field has been set.

### GetMaximumOrderQuantity

`func (o *PimVariationBase) GetMaximumOrderQuantity() float32`

GetMaximumOrderQuantity returns the MaximumOrderQuantity field if non-nil, zero value otherwise.

### GetMaximumOrderQuantityOk

`func (o *PimVariationBase) GetMaximumOrderQuantityOk() (*float32, bool)`

GetMaximumOrderQuantityOk returns a tuple with the MaximumOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOrderQuantity

`func (o *PimVariationBase) SetMaximumOrderQuantity(v float32)`

SetMaximumOrderQuantity sets MaximumOrderQuantity field to given value.

### HasMaximumOrderQuantity

`func (o *PimVariationBase) HasMaximumOrderQuantity() bool`

HasMaximumOrderQuantity returns a boolean if a field has been set.

### GetMinimumOrderQuantity

`func (o *PimVariationBase) GetMinimumOrderQuantity() float32`

GetMinimumOrderQuantity returns the MinimumOrderQuantity field if non-nil, zero value otherwise.

### GetMinimumOrderQuantityOk

`func (o *PimVariationBase) GetMinimumOrderQuantityOk() (*float32, bool)`

GetMinimumOrderQuantityOk returns a tuple with the MinimumOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumOrderQuantity

`func (o *PimVariationBase) SetMinimumOrderQuantity(v float32)`

SetMinimumOrderQuantity sets MinimumOrderQuantity field to given value.

### HasMinimumOrderQuantity

`func (o *PimVariationBase) HasMinimumOrderQuantity() bool`

HasMinimumOrderQuantity returns a boolean if a field has been set.

### GetIntervalOrderQuantity

`func (o *PimVariationBase) GetIntervalOrderQuantity() float32`

GetIntervalOrderQuantity returns the IntervalOrderQuantity field if non-nil, zero value otherwise.

### GetIntervalOrderQuantityOk

`func (o *PimVariationBase) GetIntervalOrderQuantityOk() (*float32, bool)`

GetIntervalOrderQuantityOk returns a tuple with the IntervalOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalOrderQuantity

`func (o *PimVariationBase) SetIntervalOrderQuantity(v float32)`

SetIntervalOrderQuantity sets IntervalOrderQuantity field to given value.

### HasIntervalOrderQuantity

`func (o *PimVariationBase) HasIntervalOrderQuantity() bool`

HasIntervalOrderQuantity returns a boolean if a field has been set.

### GetAvailableUntil

`func (o *PimVariationBase) GetAvailableUntil() time.Time`

GetAvailableUntil returns the AvailableUntil field if non-nil, zero value otherwise.

### GetAvailableUntilOk

`func (o *PimVariationBase) GetAvailableUntilOk() (*time.Time, bool)`

GetAvailableUntilOk returns a tuple with the AvailableUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableUntil

`func (o *PimVariationBase) SetAvailableUntil(v time.Time)`

SetAvailableUntil sets AvailableUntil field to given value.

### HasAvailableUntil

`func (o *PimVariationBase) HasAvailableUntil() bool`

HasAvailableUntil returns a boolean if a field has been set.

### GetReleasedAt

`func (o *PimVariationBase) GetReleasedAt() time.Time`

GetReleasedAt returns the ReleasedAt field if non-nil, zero value otherwise.

### GetReleasedAtOk

`func (o *PimVariationBase) GetReleasedAtOk() (*time.Time, bool)`

GetReleasedAtOk returns a tuple with the ReleasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedAt

`func (o *PimVariationBase) SetReleasedAt(v time.Time)`

SetReleasedAt sets ReleasedAt field to given value.

### HasReleasedAt

`func (o *PimVariationBase) HasReleasedAt() bool`

HasReleasedAt returns a boolean if a field has been set.

### GetName

`func (o *PimVariationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PimVariationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PimVariationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PimVariationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetWeightG

`func (o *PimVariationBase) GetWeightG() int32`

GetWeightG returns the WeightG field if non-nil, zero value otherwise.

### GetWeightGOk

`func (o *PimVariationBase) GetWeightGOk() (*int32, bool)`

GetWeightGOk returns a tuple with the WeightG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightG

`func (o *PimVariationBase) SetWeightG(v int32)`

SetWeightG sets WeightG field to given value.

### HasWeightG

`func (o *PimVariationBase) HasWeightG() bool`

HasWeightG returns a boolean if a field has been set.

### GetWeightNetG

`func (o *PimVariationBase) GetWeightNetG() int32`

GetWeightNetG returns the WeightNetG field if non-nil, zero value otherwise.

### GetWeightNetGOk

`func (o *PimVariationBase) GetWeightNetGOk() (*int32, bool)`

GetWeightNetGOk returns a tuple with the WeightNetG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightNetG

`func (o *PimVariationBase) SetWeightNetG(v int32)`

SetWeightNetG sets WeightNetG field to given value.

### HasWeightNetG

`func (o *PimVariationBase) HasWeightNetG() bool`

HasWeightNetG returns a boolean if a field has been set.

### GetWidthMM

`func (o *PimVariationBase) GetWidthMM() float32`

GetWidthMM returns the WidthMM field if non-nil, zero value otherwise.

### GetWidthMMOk

`func (o *PimVariationBase) GetWidthMMOk() (*float32, bool)`

GetWidthMMOk returns a tuple with the WidthMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthMM

`func (o *PimVariationBase) SetWidthMM(v float32)`

SetWidthMM sets WidthMM field to given value.

### HasWidthMM

`func (o *PimVariationBase) HasWidthMM() bool`

HasWidthMM returns a boolean if a field has been set.

### GetLengthMM

`func (o *PimVariationBase) GetLengthMM() float32`

GetLengthMM returns the LengthMM field if non-nil, zero value otherwise.

### GetLengthMMOk

`func (o *PimVariationBase) GetLengthMMOk() (*float32, bool)`

GetLengthMMOk returns a tuple with the LengthMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLengthMM

`func (o *PimVariationBase) SetLengthMM(v float32)`

SetLengthMM sets LengthMM field to given value.

### HasLengthMM

`func (o *PimVariationBase) HasLengthMM() bool`

HasLengthMM returns a boolean if a field has been set.

### GetHeightMM

`func (o *PimVariationBase) GetHeightMM() float32`

GetHeightMM returns the HeightMM field if non-nil, zero value otherwise.

### GetHeightMMOk

`func (o *PimVariationBase) GetHeightMMOk() (*float32, bool)`

GetHeightMMOk returns a tuple with the HeightMM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeightMM

`func (o *PimVariationBase) SetHeightMM(v float32)`

SetHeightMM sets HeightMM field to given value.

### HasHeightMM

`func (o *PimVariationBase) HasHeightMM() bool`

HasHeightMM returns a boolean if a field has been set.

### GetExtraShippingCharge1

`func (o *PimVariationBase) GetExtraShippingCharge1() float32`

GetExtraShippingCharge1 returns the ExtraShippingCharge1 field if non-nil, zero value otherwise.

### GetExtraShippingCharge1Ok

`func (o *PimVariationBase) GetExtraShippingCharge1Ok() (*float32, bool)`

GetExtraShippingCharge1Ok returns a tuple with the ExtraShippingCharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraShippingCharge1

`func (o *PimVariationBase) SetExtraShippingCharge1(v float32)`

SetExtraShippingCharge1 sets ExtraShippingCharge1 field to given value.

### HasExtraShippingCharge1

`func (o *PimVariationBase) HasExtraShippingCharge1() bool`

HasExtraShippingCharge1 returns a boolean if a field has been set.

### GetExtraShippingCharge2

`func (o *PimVariationBase) GetExtraShippingCharge2() float32`

GetExtraShippingCharge2 returns the ExtraShippingCharge2 field if non-nil, zero value otherwise.

### GetExtraShippingCharge2Ok

`func (o *PimVariationBase) GetExtraShippingCharge2Ok() (*float32, bool)`

GetExtraShippingCharge2Ok returns a tuple with the ExtraShippingCharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraShippingCharge2

`func (o *PimVariationBase) SetExtraShippingCharge2(v float32)`

SetExtraShippingCharge2 sets ExtraShippingCharge2 field to given value.

### HasExtraShippingCharge2

`func (o *PimVariationBase) HasExtraShippingCharge2() bool`

HasExtraShippingCharge2 returns a boolean if a field has been set.

### GetUnitsContained

`func (o *PimVariationBase) GetUnitsContained() int32`

GetUnitsContained returns the UnitsContained field if non-nil, zero value otherwise.

### GetUnitsContainedOk

`func (o *PimVariationBase) GetUnitsContainedOk() (*int32, bool)`

GetUnitsContainedOk returns a tuple with the UnitsContained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsContained

`func (o *PimVariationBase) SetUnitsContained(v int32)`

SetUnitsContained sets UnitsContained field to given value.

### HasUnitsContained

`func (o *PimVariationBase) HasUnitsContained() bool`

HasUnitsContained returns a boolean if a field has been set.

### GetPalletTypeId

`func (o *PimVariationBase) GetPalletTypeId() int32`

GetPalletTypeId returns the PalletTypeId field if non-nil, zero value otherwise.

### GetPalletTypeIdOk

`func (o *PimVariationBase) GetPalletTypeIdOk() (*int32, bool)`

GetPalletTypeIdOk returns a tuple with the PalletTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPalletTypeId

`func (o *PimVariationBase) SetPalletTypeId(v int32)`

SetPalletTypeId sets PalletTypeId field to given value.

### HasPalletTypeId

`func (o *PimVariationBase) HasPalletTypeId() bool`

HasPalletTypeId returns a boolean if a field has been set.

### GetPackingUnits

`func (o *PimVariationBase) GetPackingUnits() int32`

GetPackingUnits returns the PackingUnits field if non-nil, zero value otherwise.

### GetPackingUnitsOk

`func (o *PimVariationBase) GetPackingUnitsOk() (*int32, bool)`

GetPackingUnitsOk returns a tuple with the PackingUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingUnits

`func (o *PimVariationBase) SetPackingUnits(v int32)`

SetPackingUnits sets PackingUnits field to given value.

### HasPackingUnits

`func (o *PimVariationBase) HasPackingUnits() bool`

HasPackingUnits returns a boolean if a field has been set.

### GetPackingUnitTypeId

`func (o *PimVariationBase) GetPackingUnitTypeId() int32`

GetPackingUnitTypeId returns the PackingUnitTypeId field if non-nil, zero value otherwise.

### GetPackingUnitTypeIdOk

`func (o *PimVariationBase) GetPackingUnitTypeIdOk() (*int32, bool)`

GetPackingUnitTypeIdOk returns a tuple with the PackingUnitTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingUnitTypeId

`func (o *PimVariationBase) SetPackingUnitTypeId(v int32)`

SetPackingUnitTypeId sets PackingUnitTypeId field to given value.

### HasPackingUnitTypeId

`func (o *PimVariationBase) HasPackingUnitTypeId() bool`

HasPackingUnitTypeId returns a boolean if a field has been set.

### GetTransportationCosts

`func (o *PimVariationBase) GetTransportationCosts() float32`

GetTransportationCosts returns the TransportationCosts field if non-nil, zero value otherwise.

### GetTransportationCostsOk

`func (o *PimVariationBase) GetTransportationCostsOk() (*float32, bool)`

GetTransportationCostsOk returns a tuple with the TransportationCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportationCosts

`func (o *PimVariationBase) SetTransportationCosts(v float32)`

SetTransportationCosts sets TransportationCosts field to given value.

### HasTransportationCosts

`func (o *PimVariationBase) HasTransportationCosts() bool`

HasTransportationCosts returns a boolean if a field has been set.

### GetStorageCosts

`func (o *PimVariationBase) GetStorageCosts() float32`

GetStorageCosts returns the StorageCosts field if non-nil, zero value otherwise.

### GetStorageCostsOk

`func (o *PimVariationBase) GetStorageCostsOk() (*float32, bool)`

GetStorageCostsOk returns a tuple with the StorageCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCosts

`func (o *PimVariationBase) SetStorageCosts(v float32)`

SetStorageCosts sets StorageCosts field to given value.

### HasStorageCosts

`func (o *PimVariationBase) HasStorageCosts() bool`

HasStorageCosts returns a boolean if a field has been set.

### GetCustoms

`func (o *PimVariationBase) GetCustoms() float32`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *PimVariationBase) GetCustomsOk() (*float32, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *PimVariationBase) SetCustoms(v float32)`

SetCustoms sets Customs field to given value.

### HasCustoms

`func (o *PimVariationBase) HasCustoms() bool`

HasCustoms returns a boolean if a field has been set.

### GetOperatingCosts

`func (o *PimVariationBase) GetOperatingCosts() float32`

GetOperatingCosts returns the OperatingCosts field if non-nil, zero value otherwise.

### GetOperatingCostsOk

`func (o *PimVariationBase) GetOperatingCostsOk() (*float32, bool)`

GetOperatingCostsOk returns a tuple with the OperatingCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingCosts

`func (o *PimVariationBase) SetOperatingCosts(v float32)`

SetOperatingCosts sets OperatingCosts field to given value.

### HasOperatingCosts

`func (o *PimVariationBase) HasOperatingCosts() bool`

HasOperatingCosts returns a boolean if a field has been set.

### GetVatId

`func (o *PimVariationBase) GetVatId() int32`

GetVatId returns the VatId field if non-nil, zero value otherwise.

### GetVatIdOk

`func (o *PimVariationBase) GetVatIdOk() (*int32, bool)`

GetVatIdOk returns a tuple with the VatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatId

`func (o *PimVariationBase) SetVatId(v int32)`

SetVatId sets VatId field to given value.

### HasVatId

`func (o *PimVariationBase) HasVatId() bool`

HasVatId returns a boolean if a field has been set.

### GetBundleType

`func (o *PimVariationBase) GetBundleType() string`

GetBundleType returns the BundleType field if non-nil, zero value otherwise.

### GetBundleTypeOk

`func (o *PimVariationBase) GetBundleTypeOk() (*string, bool)`

GetBundleTypeOk returns a tuple with the BundleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleType

`func (o *PimVariationBase) SetBundleType(v string)`

SetBundleType sets BundleType field to given value.

### HasBundleType

`func (o *PimVariationBase) HasBundleType() bool`

HasBundleType returns a boolean if a field has been set.

### GetAutomaticClientVisibility

`func (o *PimVariationBase) GetAutomaticClientVisibility() int32`

GetAutomaticClientVisibility returns the AutomaticClientVisibility field if non-nil, zero value otherwise.

### GetAutomaticClientVisibilityOk

`func (o *PimVariationBase) GetAutomaticClientVisibilityOk() (*int32, bool)`

GetAutomaticClientVisibilityOk returns a tuple with the AutomaticClientVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticClientVisibility

`func (o *PimVariationBase) SetAutomaticClientVisibility(v int32)`

SetAutomaticClientVisibility sets AutomaticClientVisibility field to given value.

### HasAutomaticClientVisibility

`func (o *PimVariationBase) HasAutomaticClientVisibility() bool`

HasAutomaticClientVisibility returns a boolean if a field has been set.

### GetIsHiddenInCategoryList

`func (o *PimVariationBase) GetIsHiddenInCategoryList() bool`

GetIsHiddenInCategoryList returns the IsHiddenInCategoryList field if non-nil, zero value otherwise.

### GetIsHiddenInCategoryListOk

`func (o *PimVariationBase) GetIsHiddenInCategoryListOk() (*bool, bool)`

GetIsHiddenInCategoryListOk returns a tuple with the IsHiddenInCategoryList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHiddenInCategoryList

`func (o *PimVariationBase) SetIsHiddenInCategoryList(v bool)`

SetIsHiddenInCategoryList sets IsHiddenInCategoryList field to given value.

### HasIsHiddenInCategoryList

`func (o *PimVariationBase) HasIsHiddenInCategoryList() bool`

HasIsHiddenInCategoryList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


