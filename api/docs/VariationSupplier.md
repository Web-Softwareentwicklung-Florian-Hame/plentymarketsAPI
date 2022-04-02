# VariationSupplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 
**SupplierId** | Pointer to **int32** |  | [optional] 
**PurchasePrice** | Pointer to **float32** |  | [optional] 
**MinimumPurchase** | Pointer to **float32** |  | [optional] 
**ItemNumber** | Pointer to **string** |  | [optional] 
**LastPriceQuery** | Pointer to **bool** |  | [optional] 
**DeliveryTimeInDays** | Pointer to **int32** |  | [optional] 
**Discount** | Pointer to **float32** |  | [optional] 
**IsDiscountable** | Pointer to **bool** |  | [optional] 
**PackagingUnit** | Pointer to **float32** |  | [optional] 
**LastUpdateTimestamp** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrencyPurchasePrice** | Pointer to **float32** |  | [optional] 
**ItemDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewVariationSupplier

`func NewVariationSupplier() *VariationSupplier`

NewVariationSupplier instantiates a new VariationSupplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationSupplierWithDefaults

`func NewVariationSupplierWithDefaults() *VariationSupplier`

NewVariationSupplierWithDefaults instantiates a new VariationSupplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationSupplier) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationSupplier) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationSupplier) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationSupplier) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVariationId

`func (o *VariationSupplier) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *VariationSupplier) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *VariationSupplier) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *VariationSupplier) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetSupplierId

`func (o *VariationSupplier) GetSupplierId() int32`

GetSupplierId returns the SupplierId field if non-nil, zero value otherwise.

### GetSupplierIdOk

`func (o *VariationSupplier) GetSupplierIdOk() (*int32, bool)`

GetSupplierIdOk returns a tuple with the SupplierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplierId

`func (o *VariationSupplier) SetSupplierId(v int32)`

SetSupplierId sets SupplierId field to given value.

### HasSupplierId

`func (o *VariationSupplier) HasSupplierId() bool`

HasSupplierId returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *VariationSupplier) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *VariationSupplier) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *VariationSupplier) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *VariationSupplier) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetMinimumPurchase

`func (o *VariationSupplier) GetMinimumPurchase() float32`

GetMinimumPurchase returns the MinimumPurchase field if non-nil, zero value otherwise.

### GetMinimumPurchaseOk

`func (o *VariationSupplier) GetMinimumPurchaseOk() (*float32, bool)`

GetMinimumPurchaseOk returns a tuple with the MinimumPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPurchase

`func (o *VariationSupplier) SetMinimumPurchase(v float32)`

SetMinimumPurchase sets MinimumPurchase field to given value.

### HasMinimumPurchase

`func (o *VariationSupplier) HasMinimumPurchase() bool`

HasMinimumPurchase returns a boolean if a field has been set.

### GetItemNumber

`func (o *VariationSupplier) GetItemNumber() string`

GetItemNumber returns the ItemNumber field if non-nil, zero value otherwise.

### GetItemNumberOk

`func (o *VariationSupplier) GetItemNumberOk() (*string, bool)`

GetItemNumberOk returns a tuple with the ItemNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNumber

`func (o *VariationSupplier) SetItemNumber(v string)`

SetItemNumber sets ItemNumber field to given value.

### HasItemNumber

`func (o *VariationSupplier) HasItemNumber() bool`

HasItemNumber returns a boolean if a field has been set.

### GetLastPriceQuery

`func (o *VariationSupplier) GetLastPriceQuery() bool`

GetLastPriceQuery returns the LastPriceQuery field if non-nil, zero value otherwise.

### GetLastPriceQueryOk

`func (o *VariationSupplier) GetLastPriceQueryOk() (*bool, bool)`

GetLastPriceQueryOk returns a tuple with the LastPriceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPriceQuery

`func (o *VariationSupplier) SetLastPriceQuery(v bool)`

SetLastPriceQuery sets LastPriceQuery field to given value.

### HasLastPriceQuery

`func (o *VariationSupplier) HasLastPriceQuery() bool`

HasLastPriceQuery returns a boolean if a field has been set.

### GetDeliveryTimeInDays

`func (o *VariationSupplier) GetDeliveryTimeInDays() int32`

GetDeliveryTimeInDays returns the DeliveryTimeInDays field if non-nil, zero value otherwise.

### GetDeliveryTimeInDaysOk

`func (o *VariationSupplier) GetDeliveryTimeInDaysOk() (*int32, bool)`

GetDeliveryTimeInDaysOk returns a tuple with the DeliveryTimeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTimeInDays

`func (o *VariationSupplier) SetDeliveryTimeInDays(v int32)`

SetDeliveryTimeInDays sets DeliveryTimeInDays field to given value.

### HasDeliveryTimeInDays

`func (o *VariationSupplier) HasDeliveryTimeInDays() bool`

HasDeliveryTimeInDays returns a boolean if a field has been set.

### GetDiscount

`func (o *VariationSupplier) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *VariationSupplier) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *VariationSupplier) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *VariationSupplier) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetIsDiscountable

`func (o *VariationSupplier) GetIsDiscountable() bool`

GetIsDiscountable returns the IsDiscountable field if non-nil, zero value otherwise.

### GetIsDiscountableOk

`func (o *VariationSupplier) GetIsDiscountableOk() (*bool, bool)`

GetIsDiscountableOk returns a tuple with the IsDiscountable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDiscountable

`func (o *VariationSupplier) SetIsDiscountable(v bool)`

SetIsDiscountable sets IsDiscountable field to given value.

### HasIsDiscountable

`func (o *VariationSupplier) HasIsDiscountable() bool`

HasIsDiscountable returns a boolean if a field has been set.

### GetPackagingUnit

`func (o *VariationSupplier) GetPackagingUnit() float32`

GetPackagingUnit returns the PackagingUnit field if non-nil, zero value otherwise.

### GetPackagingUnitOk

`func (o *VariationSupplier) GetPackagingUnitOk() (*float32, bool)`

GetPackagingUnitOk returns a tuple with the PackagingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagingUnit

`func (o *VariationSupplier) SetPackagingUnit(v float32)`

SetPackagingUnit sets PackagingUnit field to given value.

### HasPackagingUnit

`func (o *VariationSupplier) HasPackagingUnit() bool`

HasPackagingUnit returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *VariationSupplier) GetLastUpdateTimestamp() string`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *VariationSupplier) GetLastUpdateTimestampOk() (*string, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *VariationSupplier) SetLastUpdateTimestamp(v string)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *VariationSupplier) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariationSupplier) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariationSupplier) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariationSupplier) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VariationSupplier) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrencyPurchasePrice

`func (o *VariationSupplier) GetCurrencyPurchasePrice() float32`

GetCurrencyPurchasePrice returns the CurrencyPurchasePrice field if non-nil, zero value otherwise.

### GetCurrencyPurchasePriceOk

`func (o *VariationSupplier) GetCurrencyPurchasePriceOk() (*float32, bool)`

GetCurrencyPurchasePriceOk returns a tuple with the CurrencyPurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyPurchasePrice

`func (o *VariationSupplier) SetCurrencyPurchasePrice(v float32)`

SetCurrencyPurchasePrice sets CurrencyPurchasePrice field to given value.

### HasCurrencyPurchasePrice

`func (o *VariationSupplier) HasCurrencyPurchasePrice() bool`

HasCurrencyPurchasePrice returns a boolean if a field has been set.

### GetItemDescription

`func (o *VariationSupplier) GetItemDescription() string`

GetItemDescription returns the ItemDescription field if non-nil, zero value otherwise.

### GetItemDescriptionOk

`func (o *VariationSupplier) GetItemDescriptionOk() (*string, bool)`

GetItemDescriptionOk returns a tuple with the ItemDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemDescription

`func (o *VariationSupplier) SetItemDescription(v string)`

SetItemDescription sets ItemDescription field to given value.

### HasItemDescription

`func (o *VariationSupplier) HasItemDescription() bool`

HasItemDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


