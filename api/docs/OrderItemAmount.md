# OrderItemAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrderItemId** | Pointer to **int32** |  | [optional] 
**IsSystemCurrency** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**ExchangeRate** | Pointer to **float32** |  | [optional] 
**PurchasePrice** | Pointer to **float32** |  | [optional] 
**PriceOriginalGross** | Pointer to **float32** |  | [optional] 
**PriceOriginalNet** | Pointer to **float32** |  | [optional] 
**PriceGross** | Pointer to **float32** |  | [optional] 
**PriceNet** | Pointer to **float32** |  | [optional] 
**Surcharge** | Pointer to **float32** |  | [optional] 
**Discount** | Pointer to **float32** |  | [optional] 
**IsPercentage** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOrderItemAmount

`func NewOrderItemAmount() *OrderItemAmount`

NewOrderItemAmount instantiates a new OrderItemAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemAmountWithDefaults

`func NewOrderItemAmountWithDefaults() *OrderItemAmount`

NewOrderItemAmountWithDefaults instantiates a new OrderItemAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItemAmount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItemAmount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItemAmount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderItemAmount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderItemId

`func (o *OrderItemAmount) GetOrderItemId() int32`

GetOrderItemId returns the OrderItemId field if non-nil, zero value otherwise.

### GetOrderItemIdOk

`func (o *OrderItemAmount) GetOrderItemIdOk() (*int32, bool)`

GetOrderItemIdOk returns a tuple with the OrderItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemId

`func (o *OrderItemAmount) SetOrderItemId(v int32)`

SetOrderItemId sets OrderItemId field to given value.

### HasOrderItemId

`func (o *OrderItemAmount) HasOrderItemId() bool`

HasOrderItemId returns a boolean if a field has been set.

### GetIsSystemCurrency

`func (o *OrderItemAmount) GetIsSystemCurrency() bool`

GetIsSystemCurrency returns the IsSystemCurrency field if non-nil, zero value otherwise.

### GetIsSystemCurrencyOk

`func (o *OrderItemAmount) GetIsSystemCurrencyOk() (*bool, bool)`

GetIsSystemCurrencyOk returns a tuple with the IsSystemCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemCurrency

`func (o *OrderItemAmount) SetIsSystemCurrency(v bool)`

SetIsSystemCurrency sets IsSystemCurrency field to given value.

### HasIsSystemCurrency

`func (o *OrderItemAmount) HasIsSystemCurrency() bool`

HasIsSystemCurrency returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderItemAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderItemAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderItemAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderItemAmount) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchangeRate

`func (o *OrderItemAmount) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *OrderItemAmount) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *OrderItemAmount) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *OrderItemAmount) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetPurchasePrice

`func (o *OrderItemAmount) GetPurchasePrice() float32`

GetPurchasePrice returns the PurchasePrice field if non-nil, zero value otherwise.

### GetPurchasePriceOk

`func (o *OrderItemAmount) GetPurchasePriceOk() (*float32, bool)`

GetPurchasePriceOk returns a tuple with the PurchasePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasePrice

`func (o *OrderItemAmount) SetPurchasePrice(v float32)`

SetPurchasePrice sets PurchasePrice field to given value.

### HasPurchasePrice

`func (o *OrderItemAmount) HasPurchasePrice() bool`

HasPurchasePrice returns a boolean if a field has been set.

### GetPriceOriginalGross

`func (o *OrderItemAmount) GetPriceOriginalGross() float32`

GetPriceOriginalGross returns the PriceOriginalGross field if non-nil, zero value otherwise.

### GetPriceOriginalGrossOk

`func (o *OrderItemAmount) GetPriceOriginalGrossOk() (*float32, bool)`

GetPriceOriginalGrossOk returns a tuple with the PriceOriginalGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOriginalGross

`func (o *OrderItemAmount) SetPriceOriginalGross(v float32)`

SetPriceOriginalGross sets PriceOriginalGross field to given value.

### HasPriceOriginalGross

`func (o *OrderItemAmount) HasPriceOriginalGross() bool`

HasPriceOriginalGross returns a boolean if a field has been set.

### GetPriceOriginalNet

`func (o *OrderItemAmount) GetPriceOriginalNet() float32`

GetPriceOriginalNet returns the PriceOriginalNet field if non-nil, zero value otherwise.

### GetPriceOriginalNetOk

`func (o *OrderItemAmount) GetPriceOriginalNetOk() (*float32, bool)`

GetPriceOriginalNetOk returns a tuple with the PriceOriginalNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceOriginalNet

`func (o *OrderItemAmount) SetPriceOriginalNet(v float32)`

SetPriceOriginalNet sets PriceOriginalNet field to given value.

### HasPriceOriginalNet

`func (o *OrderItemAmount) HasPriceOriginalNet() bool`

HasPriceOriginalNet returns a boolean if a field has been set.

### GetPriceGross

`func (o *OrderItemAmount) GetPriceGross() float32`

GetPriceGross returns the PriceGross field if non-nil, zero value otherwise.

### GetPriceGrossOk

`func (o *OrderItemAmount) GetPriceGrossOk() (*float32, bool)`

GetPriceGrossOk returns a tuple with the PriceGross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceGross

`func (o *OrderItemAmount) SetPriceGross(v float32)`

SetPriceGross sets PriceGross field to given value.

### HasPriceGross

`func (o *OrderItemAmount) HasPriceGross() bool`

HasPriceGross returns a boolean if a field has been set.

### GetPriceNet

`func (o *OrderItemAmount) GetPriceNet() float32`

GetPriceNet returns the PriceNet field if non-nil, zero value otherwise.

### GetPriceNetOk

`func (o *OrderItemAmount) GetPriceNetOk() (*float32, bool)`

GetPriceNetOk returns a tuple with the PriceNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceNet

`func (o *OrderItemAmount) SetPriceNet(v float32)`

SetPriceNet sets PriceNet field to given value.

### HasPriceNet

`func (o *OrderItemAmount) HasPriceNet() bool`

HasPriceNet returns a boolean if a field has been set.

### GetSurcharge

`func (o *OrderItemAmount) GetSurcharge() float32`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *OrderItemAmount) GetSurchargeOk() (*float32, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *OrderItemAmount) SetSurcharge(v float32)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *OrderItemAmount) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetDiscount

`func (o *OrderItemAmount) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *OrderItemAmount) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *OrderItemAmount) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *OrderItemAmount) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetIsPercentage

`func (o *OrderItemAmount) GetIsPercentage() bool`

GetIsPercentage returns the IsPercentage field if non-nil, zero value otherwise.

### GetIsPercentageOk

`func (o *OrderItemAmount) GetIsPercentageOk() (*bool, bool)`

GetIsPercentageOk returns a tuple with the IsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPercentage

`func (o *OrderItemAmount) SetIsPercentage(v bool)`

SetIsPercentage sets IsPercentage field to given value.

### HasIsPercentage

`func (o *OrderItemAmount) HasIsPercentage() bool`

HasIsPercentage returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderItemAmount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderItemAmount) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderItemAmount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderItemAmount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderItemAmount) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderItemAmount) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderItemAmount) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderItemAmount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


