# OrderItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**TypeId** | Pointer to **int32** |  | [optional] 
**ReferrerId** | Pointer to **float32** |  | [optional] 
**ItemVariationId** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **float32** |  | [optional] 
**OrderItemName** | Pointer to **string** |  | [optional] 
**AttributeValues** | Pointer to **string** |  | [optional] 
**ShippingProfileId** | Pointer to **int32** |  | [optional] 
**CountryVatId** | Pointer to **int32** |  | [optional] 
**VatField** | Pointer to **int32** |  | [optional] 
**VatRate** | Pointer to **float32** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**WarehouseId** | Pointer to **int32** |  | [optional] 
**Variation** | Pointer to [**Variation**](Variation.md) |  | [optional] 
**VariationBarcodes** | Pointer to [**[]VariationBarcode**](VariationBarcode.md) |  | [optional] 

## Methods

### NewOrderItem

`func NewOrderItem() *OrderItem`

NewOrderItem instantiates a new OrderItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderItemWithDefaults

`func NewOrderItemWithDefaults() *OrderItem`

NewOrderItemWithDefaults instantiates a new OrderItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderItem) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderItem) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderItem) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetTypeId

`func (o *OrderItem) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *OrderItem) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *OrderItem) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *OrderItem) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetReferrerId

`func (o *OrderItem) GetReferrerId() float32`

GetReferrerId returns the ReferrerId field if non-nil, zero value otherwise.

### GetReferrerIdOk

`func (o *OrderItem) GetReferrerIdOk() (*float32, bool)`

GetReferrerIdOk returns a tuple with the ReferrerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerId

`func (o *OrderItem) SetReferrerId(v float32)`

SetReferrerId sets ReferrerId field to given value.

### HasReferrerId

`func (o *OrderItem) HasReferrerId() bool`

HasReferrerId returns a boolean if a field has been set.

### GetItemVariationId

`func (o *OrderItem) GetItemVariationId() int32`

GetItemVariationId returns the ItemVariationId field if non-nil, zero value otherwise.

### GetItemVariationIdOk

`func (o *OrderItem) GetItemVariationIdOk() (*int32, bool)`

GetItemVariationIdOk returns a tuple with the ItemVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemVariationId

`func (o *OrderItem) SetItemVariationId(v int32)`

SetItemVariationId sets ItemVariationId field to given value.

### HasItemVariationId

`func (o *OrderItem) HasItemVariationId() bool`

HasItemVariationId returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOrderItemName

`func (o *OrderItem) GetOrderItemName() string`

GetOrderItemName returns the OrderItemName field if non-nil, zero value otherwise.

### GetOrderItemNameOk

`func (o *OrderItem) GetOrderItemNameOk() (*string, bool)`

GetOrderItemNameOk returns a tuple with the OrderItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItemName

`func (o *OrderItem) SetOrderItemName(v string)`

SetOrderItemName sets OrderItemName field to given value.

### HasOrderItemName

`func (o *OrderItem) HasOrderItemName() bool`

HasOrderItemName returns a boolean if a field has been set.

### GetAttributeValues

`func (o *OrderItem) GetAttributeValues() string`

GetAttributeValues returns the AttributeValues field if non-nil, zero value otherwise.

### GetAttributeValuesOk

`func (o *OrderItem) GetAttributeValuesOk() (*string, bool)`

GetAttributeValuesOk returns a tuple with the AttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValues

`func (o *OrderItem) SetAttributeValues(v string)`

SetAttributeValues sets AttributeValues field to given value.

### HasAttributeValues

`func (o *OrderItem) HasAttributeValues() bool`

HasAttributeValues returns a boolean if a field has been set.

### GetShippingProfileId

`func (o *OrderItem) GetShippingProfileId() int32`

GetShippingProfileId returns the ShippingProfileId field if non-nil, zero value otherwise.

### GetShippingProfileIdOk

`func (o *OrderItem) GetShippingProfileIdOk() (*int32, bool)`

GetShippingProfileIdOk returns a tuple with the ShippingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingProfileId

`func (o *OrderItem) SetShippingProfileId(v int32)`

SetShippingProfileId sets ShippingProfileId field to given value.

### HasShippingProfileId

`func (o *OrderItem) HasShippingProfileId() bool`

HasShippingProfileId returns a boolean if a field has been set.

### GetCountryVatId

`func (o *OrderItem) GetCountryVatId() int32`

GetCountryVatId returns the CountryVatId field if non-nil, zero value otherwise.

### GetCountryVatIdOk

`func (o *OrderItem) GetCountryVatIdOk() (*int32, bool)`

GetCountryVatIdOk returns a tuple with the CountryVatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryVatId

`func (o *OrderItem) SetCountryVatId(v int32)`

SetCountryVatId sets CountryVatId field to given value.

### HasCountryVatId

`func (o *OrderItem) HasCountryVatId() bool`

HasCountryVatId returns a boolean if a field has been set.

### GetVatField

`func (o *OrderItem) GetVatField() int32`

GetVatField returns the VatField field if non-nil, zero value otherwise.

### GetVatFieldOk

`func (o *OrderItem) GetVatFieldOk() (*int32, bool)`

GetVatFieldOk returns a tuple with the VatField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatField

`func (o *OrderItem) SetVatField(v int32)`

SetVatField sets VatField field to given value.

### HasVatField

`func (o *OrderItem) HasVatField() bool`

HasVatField returns a boolean if a field has been set.

### GetVatRate

`func (o *OrderItem) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *OrderItem) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *OrderItem) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *OrderItem) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetPosition

`func (o *OrderItem) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *OrderItem) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *OrderItem) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *OrderItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWarehouseId

`func (o *OrderItem) GetWarehouseId() int32`

GetWarehouseId returns the WarehouseId field if non-nil, zero value otherwise.

### GetWarehouseIdOk

`func (o *OrderItem) GetWarehouseIdOk() (*int32, bool)`

GetWarehouseIdOk returns a tuple with the WarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseId

`func (o *OrderItem) SetWarehouseId(v int32)`

SetWarehouseId sets WarehouseId field to given value.

### HasWarehouseId

`func (o *OrderItem) HasWarehouseId() bool`

HasWarehouseId returns a boolean if a field has been set.

### GetVariation

`func (o *OrderItem) GetVariation() Variation`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *OrderItem) GetVariationOk() (*Variation, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *OrderItem) SetVariation(v Variation)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *OrderItem) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetVariationBarcodes

`func (o *OrderItem) GetVariationBarcodes() []VariationBarcode`

GetVariationBarcodes returns the VariationBarcodes field if non-nil, zero value otherwise.

### GetVariationBarcodesOk

`func (o *OrderItem) GetVariationBarcodesOk() (*[]VariationBarcode, bool)`

GetVariationBarcodesOk returns a tuple with the VariationBarcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationBarcodes

`func (o *OrderItem) SetVariationBarcodes(v []VariationBarcode)`

SetVariationBarcodes sets VariationBarcodes field to given value.

### HasVariationBarcodes

`func (o *OrderItem) HasVariationBarcodes() bool`

HasVariationBarcodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


