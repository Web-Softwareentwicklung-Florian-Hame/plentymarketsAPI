# PimVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the variation. optional | [optional] 
**AdditionalSkus** | Pointer to [**[]PimVariationAdditionalSku**](PimVariationAdditionalSku.md) |  | [optional] 
**AttributeValues** | Pointer to [**[]VariationAttributeValue**](VariationAttributeValue.md) |  | [optional] 
**Barcodes** | Pointer to [**[]VariationBarcode**](VariationBarcode.md) |  | [optional] 
**BundleComponents** | Pointer to [**[]PimVariationBundleComponent**](PimVariationBundleComponent.md) |  | [optional] 
**Categories** | Pointer to [**[]VariationCategory**](VariationCategory.md) |  | [optional] 
**Clients** | Pointer to [**[]VariationClient**](VariationClient.md) |  | [optional] 
**DefaultCategories** | Pointer to [**[]VariationDefaultCategory**](VariationDefaultCategory.md) |  | [optional] 
**Images** | Pointer to [**[]VariationImage**](VariationImage.md) |  | [optional] 
**Markets** | Pointer to [**[]VariationMarket**](VariationMarket.md) |  | [optional] 
**MarketIdentNumbers** | Pointer to [**[]VariationMarketItemNumber**](VariationMarketItemNumber.md) |  | [optional] 
**SalesPrices** | Pointer to [**[]VariationSalesPrice**](VariationSalesPrice.md) |  | [optional] 
**Skus** | Pointer to [**[]VariationSku**](VariationSku.md) |  | [optional] 
**Supplier** | Pointer to [**[]VariationSupplier**](VariationSupplier.md) |  | [optional] 
**Warehouses** | Pointer to [**[]PimVariationWarehouse**](PimVariationWarehouse.md) |  | [optional] 
**Properties** | Pointer to [**[]PimVariationProperty**](PimVariationProperty.md) |  | [optional] 
**Tags** | Pointer to [**[]PimVariationTag**](PimVariationTag.md) |  | [optional] 
**Comments** | Pointer to [**[]PimVariationComment**](PimVariationComment.md) |  | [optional] 
**Unit** | Pointer to [**PimVariationUnit**](PimVariationUnit.md) |  | [optional] 
**Base** | Pointer to [**PimVariationBase**](PimVariationBase.md) |  | [optional] 

## Methods

### NewPimVariation

`func NewPimVariation() *PimVariation`

NewPimVariation instantiates a new PimVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimVariationWithDefaults

`func NewPimVariationWithDefaults() *PimVariation`

NewPimVariationWithDefaults instantiates a new PimVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PimVariation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PimVariation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PimVariation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PimVariation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAdditionalSkus

`func (o *PimVariation) GetAdditionalSkus() []PimVariationAdditionalSku`

GetAdditionalSkus returns the AdditionalSkus field if non-nil, zero value otherwise.

### GetAdditionalSkusOk

`func (o *PimVariation) GetAdditionalSkusOk() (*[]PimVariationAdditionalSku, bool)`

GetAdditionalSkusOk returns a tuple with the AdditionalSkus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalSkus

`func (o *PimVariation) SetAdditionalSkus(v []PimVariationAdditionalSku)`

SetAdditionalSkus sets AdditionalSkus field to given value.

### HasAdditionalSkus

`func (o *PimVariation) HasAdditionalSkus() bool`

HasAdditionalSkus returns a boolean if a field has been set.

### GetAttributeValues

`func (o *PimVariation) GetAttributeValues() []VariationAttributeValue`

GetAttributeValues returns the AttributeValues field if non-nil, zero value otherwise.

### GetAttributeValuesOk

`func (o *PimVariation) GetAttributeValuesOk() (*[]VariationAttributeValue, bool)`

GetAttributeValuesOk returns a tuple with the AttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeValues

`func (o *PimVariation) SetAttributeValues(v []VariationAttributeValue)`

SetAttributeValues sets AttributeValues field to given value.

### HasAttributeValues

`func (o *PimVariation) HasAttributeValues() bool`

HasAttributeValues returns a boolean if a field has been set.

### GetBarcodes

`func (o *PimVariation) GetBarcodes() []VariationBarcode`

GetBarcodes returns the Barcodes field if non-nil, zero value otherwise.

### GetBarcodesOk

`func (o *PimVariation) GetBarcodesOk() (*[]VariationBarcode, bool)`

GetBarcodesOk returns a tuple with the Barcodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodes

`func (o *PimVariation) SetBarcodes(v []VariationBarcode)`

SetBarcodes sets Barcodes field to given value.

### HasBarcodes

`func (o *PimVariation) HasBarcodes() bool`

HasBarcodes returns a boolean if a field has been set.

### GetBundleComponents

`func (o *PimVariation) GetBundleComponents() []PimVariationBundleComponent`

GetBundleComponents returns the BundleComponents field if non-nil, zero value otherwise.

### GetBundleComponentsOk

`func (o *PimVariation) GetBundleComponentsOk() (*[]PimVariationBundleComponent, bool)`

GetBundleComponentsOk returns a tuple with the BundleComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleComponents

`func (o *PimVariation) SetBundleComponents(v []PimVariationBundleComponent)`

SetBundleComponents sets BundleComponents field to given value.

### HasBundleComponents

`func (o *PimVariation) HasBundleComponents() bool`

HasBundleComponents returns a boolean if a field has been set.

### GetCategories

`func (o *PimVariation) GetCategories() []VariationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PimVariation) GetCategoriesOk() (*[]VariationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PimVariation) SetCategories(v []VariationCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *PimVariation) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetClients

`func (o *PimVariation) GetClients() []VariationClient`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *PimVariation) GetClientsOk() (*[]VariationClient, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *PimVariation) SetClients(v []VariationClient)`

SetClients sets Clients field to given value.

### HasClients

`func (o *PimVariation) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDefaultCategories

`func (o *PimVariation) GetDefaultCategories() []VariationDefaultCategory`

GetDefaultCategories returns the DefaultCategories field if non-nil, zero value otherwise.

### GetDefaultCategoriesOk

`func (o *PimVariation) GetDefaultCategoriesOk() (*[]VariationDefaultCategory, bool)`

GetDefaultCategoriesOk returns a tuple with the DefaultCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCategories

`func (o *PimVariation) SetDefaultCategories(v []VariationDefaultCategory)`

SetDefaultCategories sets DefaultCategories field to given value.

### HasDefaultCategories

`func (o *PimVariation) HasDefaultCategories() bool`

HasDefaultCategories returns a boolean if a field has been set.

### GetImages

`func (o *PimVariation) GetImages() []VariationImage`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PimVariation) GetImagesOk() (*[]VariationImage, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PimVariation) SetImages(v []VariationImage)`

SetImages sets Images field to given value.

### HasImages

`func (o *PimVariation) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetMarkets

`func (o *PimVariation) GetMarkets() []VariationMarket`

GetMarkets returns the Markets field if non-nil, zero value otherwise.

### GetMarketsOk

`func (o *PimVariation) GetMarketsOk() (*[]VariationMarket, bool)`

GetMarketsOk returns a tuple with the Markets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkets

`func (o *PimVariation) SetMarkets(v []VariationMarket)`

SetMarkets sets Markets field to given value.

### HasMarkets

`func (o *PimVariation) HasMarkets() bool`

HasMarkets returns a boolean if a field has been set.

### GetMarketIdentNumbers

`func (o *PimVariation) GetMarketIdentNumbers() []VariationMarketItemNumber`

GetMarketIdentNumbers returns the MarketIdentNumbers field if non-nil, zero value otherwise.

### GetMarketIdentNumbersOk

`func (o *PimVariation) GetMarketIdentNumbersOk() (*[]VariationMarketItemNumber, bool)`

GetMarketIdentNumbersOk returns a tuple with the MarketIdentNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketIdentNumbers

`func (o *PimVariation) SetMarketIdentNumbers(v []VariationMarketItemNumber)`

SetMarketIdentNumbers sets MarketIdentNumbers field to given value.

### HasMarketIdentNumbers

`func (o *PimVariation) HasMarketIdentNumbers() bool`

HasMarketIdentNumbers returns a boolean if a field has been set.

### GetSalesPrices

`func (o *PimVariation) GetSalesPrices() []VariationSalesPrice`

GetSalesPrices returns the SalesPrices field if non-nil, zero value otherwise.

### GetSalesPricesOk

`func (o *PimVariation) GetSalesPricesOk() (*[]VariationSalesPrice, bool)`

GetSalesPricesOk returns a tuple with the SalesPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesPrices

`func (o *PimVariation) SetSalesPrices(v []VariationSalesPrice)`

SetSalesPrices sets SalesPrices field to given value.

### HasSalesPrices

`func (o *PimVariation) HasSalesPrices() bool`

HasSalesPrices returns a boolean if a field has been set.

### GetSkus

`func (o *PimVariation) GetSkus() []VariationSku`

GetSkus returns the Skus field if non-nil, zero value otherwise.

### GetSkusOk

`func (o *PimVariation) GetSkusOk() (*[]VariationSku, bool)`

GetSkusOk returns a tuple with the Skus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkus

`func (o *PimVariation) SetSkus(v []VariationSku)`

SetSkus sets Skus field to given value.

### HasSkus

`func (o *PimVariation) HasSkus() bool`

HasSkus returns a boolean if a field has been set.

### GetSupplier

`func (o *PimVariation) GetSupplier() []VariationSupplier`

GetSupplier returns the Supplier field if non-nil, zero value otherwise.

### GetSupplierOk

`func (o *PimVariation) GetSupplierOk() (*[]VariationSupplier, bool)`

GetSupplierOk returns a tuple with the Supplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplier

`func (o *PimVariation) SetSupplier(v []VariationSupplier)`

SetSupplier sets Supplier field to given value.

### HasSupplier

`func (o *PimVariation) HasSupplier() bool`

HasSupplier returns a boolean if a field has been set.

### GetWarehouses

`func (o *PimVariation) GetWarehouses() []PimVariationWarehouse`

GetWarehouses returns the Warehouses field if non-nil, zero value otherwise.

### GetWarehousesOk

`func (o *PimVariation) GetWarehousesOk() (*[]PimVariationWarehouse, bool)`

GetWarehousesOk returns a tuple with the Warehouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouses

`func (o *PimVariation) SetWarehouses(v []PimVariationWarehouse)`

SetWarehouses sets Warehouses field to given value.

### HasWarehouses

`func (o *PimVariation) HasWarehouses() bool`

HasWarehouses returns a boolean if a field has been set.

### GetProperties

`func (o *PimVariation) GetProperties() []PimVariationProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PimVariation) GetPropertiesOk() (*[]PimVariationProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PimVariation) SetProperties(v []PimVariationProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PimVariation) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *PimVariation) GetTags() []PimVariationTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PimVariation) GetTagsOk() (*[]PimVariationTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PimVariation) SetTags(v []PimVariationTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PimVariation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetComments

`func (o *PimVariation) GetComments() []PimVariationComment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PimVariation) GetCommentsOk() (*[]PimVariationComment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PimVariation) SetComments(v []PimVariationComment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PimVariation) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetUnit

`func (o *PimVariation) GetUnit() PimVariationUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PimVariation) GetUnitOk() (*PimVariationUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PimVariation) SetUnit(v PimVariationUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PimVariation) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetBase

`func (o *PimVariation) GetBase() PimVariationBase`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PimVariation) GetBaseOk() (*PimVariationBase, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PimVariation) SetBase(v PimVariationBase)`

SetBase sets Base field to given value.

### HasBase

`func (o *PimVariation) HasBase() bool`

HasBase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


