# VariationSku

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 
**MarketId** | Pointer to **float32** |  | [optional] 
**AccountId** | Pointer to **int32** |  | [optional] 
**InitialSku** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**ParentSku** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ExportedAt** | Pointer to **time.Time** |  | [optional] 
**StockUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AdditionalInformation** | Pointer to **string** |  | [optional] 

## Methods

### NewVariationSku

`func NewVariationSku() *VariationSku`

NewVariationSku instantiates a new VariationSku object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationSkuWithDefaults

`func NewVariationSkuWithDefaults() *VariationSku`

NewVariationSkuWithDefaults instantiates a new VariationSku object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationSku) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationSku) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationSku) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationSku) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVariationId

`func (o *VariationSku) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *VariationSku) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *VariationSku) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *VariationSku) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetMarketId

`func (o *VariationSku) GetMarketId() float32`

GetMarketId returns the MarketId field if non-nil, zero value otherwise.

### GetMarketIdOk

`func (o *VariationSku) GetMarketIdOk() (*float32, bool)`

GetMarketIdOk returns a tuple with the MarketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketId

`func (o *VariationSku) SetMarketId(v float32)`

SetMarketId sets MarketId field to given value.

### HasMarketId

`func (o *VariationSku) HasMarketId() bool`

HasMarketId returns a boolean if a field has been set.

### GetAccountId

`func (o *VariationSku) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VariationSku) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VariationSku) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VariationSku) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetInitialSku

`func (o *VariationSku) GetInitialSku() string`

GetInitialSku returns the InitialSku field if non-nil, zero value otherwise.

### GetInitialSkuOk

`func (o *VariationSku) GetInitialSkuOk() (*string, bool)`

GetInitialSkuOk returns a tuple with the InitialSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSku

`func (o *VariationSku) SetInitialSku(v string)`

SetInitialSku sets InitialSku field to given value.

### HasInitialSku

`func (o *VariationSku) HasInitialSku() bool`

HasInitialSku returns a boolean if a field has been set.

### GetSku

`func (o *VariationSku) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *VariationSku) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *VariationSku) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *VariationSku) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetParentSku

`func (o *VariationSku) GetParentSku() string`

GetParentSku returns the ParentSku field if non-nil, zero value otherwise.

### GetParentSkuOk

`func (o *VariationSku) GetParentSkuOk() (*string, bool)`

GetParentSkuOk returns a tuple with the ParentSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSku

`func (o *VariationSku) SetParentSku(v string)`

SetParentSku sets ParentSku field to given value.

### HasParentSku

`func (o *VariationSku) HasParentSku() bool`

HasParentSku returns a boolean if a field has been set.

### GetIsActive

`func (o *VariationSku) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *VariationSku) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *VariationSku) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *VariationSku) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariationSku) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariationSku) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariationSku) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VariationSku) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VariationSku) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariationSku) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariationSku) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VariationSku) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExportedAt

`func (o *VariationSku) GetExportedAt() time.Time`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *VariationSku) GetExportedAtOk() (*time.Time, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *VariationSku) SetExportedAt(v time.Time)`

SetExportedAt sets ExportedAt field to given value.

### HasExportedAt

`func (o *VariationSku) HasExportedAt() bool`

HasExportedAt returns a boolean if a field has been set.

### GetStockUpdatedAt

`func (o *VariationSku) GetStockUpdatedAt() time.Time`

GetStockUpdatedAt returns the StockUpdatedAt field if non-nil, zero value otherwise.

### GetStockUpdatedAtOk

`func (o *VariationSku) GetStockUpdatedAtOk() (*time.Time, bool)`

GetStockUpdatedAtOk returns a tuple with the StockUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockUpdatedAt

`func (o *VariationSku) SetStockUpdatedAt(v time.Time)`

SetStockUpdatedAt sets StockUpdatedAt field to given value.

### HasStockUpdatedAt

`func (o *VariationSku) HasStockUpdatedAt() bool`

HasStockUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *VariationSku) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VariationSku) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VariationSku) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VariationSku) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetStatus

`func (o *VariationSku) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VariationSku) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VariationSku) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VariationSku) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdditionalInformation

`func (o *VariationSku) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *VariationSku) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *VariationSku) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *VariationSku) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


