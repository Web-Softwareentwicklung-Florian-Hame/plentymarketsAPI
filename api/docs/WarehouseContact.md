# WarehouseContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**TypeId** | Pointer to **int32** |  | [optional] 
**OnStockAvailability** | Pointer to **int32** |  | [optional] 
**OutOfStockAvailability** | Pointer to **int32** |  | [optional] 
**SplitByShippingProfile** | Pointer to **bool** |  | [optional] 
**StorageLocationType** | Pointer to **string** |  | [optional] 
**StorageLocationZone** | Pointer to **int32** |  | [optional] 
**RepairWarehouseId** | Pointer to **int32** |  | [optional] 
**IsInventoryModeActive** | Pointer to **bool** |  | [optional] 
**LogisticsType** | Pointer to **string** |  | [optional] 
**AveragePriceSource** | Pointer to **string** |  | [optional] 
**ReorderLevelDynamic** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AllocationReferrerIds** | Pointer to **[]string** |  | [optional] 
**Address** | Pointer to [**WarehouseContactAddress**](WarehouseContactAddress.md) |  | [optional] 

## Methods

### NewWarehouseContact

`func NewWarehouseContact() *WarehouseContact`

NewWarehouseContact instantiates a new WarehouseContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseContactWithDefaults

`func NewWarehouseContactWithDefaults() *WarehouseContact`

NewWarehouseContactWithDefaults instantiates a new WarehouseContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WarehouseContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WarehouseContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WarehouseContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WarehouseContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WarehouseContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WarehouseContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WarehouseContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WarehouseContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *WarehouseContact) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *WarehouseContact) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *WarehouseContact) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *WarehouseContact) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTypeId

`func (o *WarehouseContact) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *WarehouseContact) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *WarehouseContact) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *WarehouseContact) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetOnStockAvailability

`func (o *WarehouseContact) GetOnStockAvailability() int32`

GetOnStockAvailability returns the OnStockAvailability field if non-nil, zero value otherwise.

### GetOnStockAvailabilityOk

`func (o *WarehouseContact) GetOnStockAvailabilityOk() (*int32, bool)`

GetOnStockAvailabilityOk returns a tuple with the OnStockAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnStockAvailability

`func (o *WarehouseContact) SetOnStockAvailability(v int32)`

SetOnStockAvailability sets OnStockAvailability field to given value.

### HasOnStockAvailability

`func (o *WarehouseContact) HasOnStockAvailability() bool`

HasOnStockAvailability returns a boolean if a field has been set.

### GetOutOfStockAvailability

`func (o *WarehouseContact) GetOutOfStockAvailability() int32`

GetOutOfStockAvailability returns the OutOfStockAvailability field if non-nil, zero value otherwise.

### GetOutOfStockAvailabilityOk

`func (o *WarehouseContact) GetOutOfStockAvailabilityOk() (*int32, bool)`

GetOutOfStockAvailabilityOk returns a tuple with the OutOfStockAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfStockAvailability

`func (o *WarehouseContact) SetOutOfStockAvailability(v int32)`

SetOutOfStockAvailability sets OutOfStockAvailability field to given value.

### HasOutOfStockAvailability

`func (o *WarehouseContact) HasOutOfStockAvailability() bool`

HasOutOfStockAvailability returns a boolean if a field has been set.

### GetSplitByShippingProfile

`func (o *WarehouseContact) GetSplitByShippingProfile() bool`

GetSplitByShippingProfile returns the SplitByShippingProfile field if non-nil, zero value otherwise.

### GetSplitByShippingProfileOk

`func (o *WarehouseContact) GetSplitByShippingProfileOk() (*bool, bool)`

GetSplitByShippingProfileOk returns a tuple with the SplitByShippingProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitByShippingProfile

`func (o *WarehouseContact) SetSplitByShippingProfile(v bool)`

SetSplitByShippingProfile sets SplitByShippingProfile field to given value.

### HasSplitByShippingProfile

`func (o *WarehouseContact) HasSplitByShippingProfile() bool`

HasSplitByShippingProfile returns a boolean if a field has been set.

### GetStorageLocationType

`func (o *WarehouseContact) GetStorageLocationType() string`

GetStorageLocationType returns the StorageLocationType field if non-nil, zero value otherwise.

### GetStorageLocationTypeOk

`func (o *WarehouseContact) GetStorageLocationTypeOk() (*string, bool)`

GetStorageLocationTypeOk returns a tuple with the StorageLocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocationType

`func (o *WarehouseContact) SetStorageLocationType(v string)`

SetStorageLocationType sets StorageLocationType field to given value.

### HasStorageLocationType

`func (o *WarehouseContact) HasStorageLocationType() bool`

HasStorageLocationType returns a boolean if a field has been set.

### GetStorageLocationZone

`func (o *WarehouseContact) GetStorageLocationZone() int32`

GetStorageLocationZone returns the StorageLocationZone field if non-nil, zero value otherwise.

### GetStorageLocationZoneOk

`func (o *WarehouseContact) GetStorageLocationZoneOk() (*int32, bool)`

GetStorageLocationZoneOk returns a tuple with the StorageLocationZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocationZone

`func (o *WarehouseContact) SetStorageLocationZone(v int32)`

SetStorageLocationZone sets StorageLocationZone field to given value.

### HasStorageLocationZone

`func (o *WarehouseContact) HasStorageLocationZone() bool`

HasStorageLocationZone returns a boolean if a field has been set.

### GetRepairWarehouseId

`func (o *WarehouseContact) GetRepairWarehouseId() int32`

GetRepairWarehouseId returns the RepairWarehouseId field if non-nil, zero value otherwise.

### GetRepairWarehouseIdOk

`func (o *WarehouseContact) GetRepairWarehouseIdOk() (*int32, bool)`

GetRepairWarehouseIdOk returns a tuple with the RepairWarehouseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairWarehouseId

`func (o *WarehouseContact) SetRepairWarehouseId(v int32)`

SetRepairWarehouseId sets RepairWarehouseId field to given value.

### HasRepairWarehouseId

`func (o *WarehouseContact) HasRepairWarehouseId() bool`

HasRepairWarehouseId returns a boolean if a field has been set.

### GetIsInventoryModeActive

`func (o *WarehouseContact) GetIsInventoryModeActive() bool`

GetIsInventoryModeActive returns the IsInventoryModeActive field if non-nil, zero value otherwise.

### GetIsInventoryModeActiveOk

`func (o *WarehouseContact) GetIsInventoryModeActiveOk() (*bool, bool)`

GetIsInventoryModeActiveOk returns a tuple with the IsInventoryModeActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInventoryModeActive

`func (o *WarehouseContact) SetIsInventoryModeActive(v bool)`

SetIsInventoryModeActive sets IsInventoryModeActive field to given value.

### HasIsInventoryModeActive

`func (o *WarehouseContact) HasIsInventoryModeActive() bool`

HasIsInventoryModeActive returns a boolean if a field has been set.

### GetLogisticsType

`func (o *WarehouseContact) GetLogisticsType() string`

GetLogisticsType returns the LogisticsType field if non-nil, zero value otherwise.

### GetLogisticsTypeOk

`func (o *WarehouseContact) GetLogisticsTypeOk() (*string, bool)`

GetLogisticsTypeOk returns a tuple with the LogisticsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogisticsType

`func (o *WarehouseContact) SetLogisticsType(v string)`

SetLogisticsType sets LogisticsType field to given value.

### HasLogisticsType

`func (o *WarehouseContact) HasLogisticsType() bool`

HasLogisticsType returns a boolean if a field has been set.

### GetAveragePriceSource

`func (o *WarehouseContact) GetAveragePriceSource() string`

GetAveragePriceSource returns the AveragePriceSource field if non-nil, zero value otherwise.

### GetAveragePriceSourceOk

`func (o *WarehouseContact) GetAveragePriceSourceOk() (*string, bool)`

GetAveragePriceSourceOk returns a tuple with the AveragePriceSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePriceSource

`func (o *WarehouseContact) SetAveragePriceSource(v string)`

SetAveragePriceSource sets AveragePriceSource field to given value.

### HasAveragePriceSource

`func (o *WarehouseContact) HasAveragePriceSource() bool`

HasAveragePriceSource returns a boolean if a field has been set.

### GetReorderLevelDynamic

`func (o *WarehouseContact) GetReorderLevelDynamic() string`

GetReorderLevelDynamic returns the ReorderLevelDynamic field if non-nil, zero value otherwise.

### GetReorderLevelDynamicOk

`func (o *WarehouseContact) GetReorderLevelDynamicOk() (*string, bool)`

GetReorderLevelDynamicOk returns a tuple with the ReorderLevelDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReorderLevelDynamic

`func (o *WarehouseContact) SetReorderLevelDynamic(v string)`

SetReorderLevelDynamic sets ReorderLevelDynamic field to given value.

### HasReorderLevelDynamic

`func (o *WarehouseContact) HasReorderLevelDynamic() bool`

HasReorderLevelDynamic returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WarehouseContact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WarehouseContact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WarehouseContact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WarehouseContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WarehouseContact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WarehouseContact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WarehouseContact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WarehouseContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAllocationReferrerIds

`func (o *WarehouseContact) GetAllocationReferrerIds() []string`

GetAllocationReferrerIds returns the AllocationReferrerIds field if non-nil, zero value otherwise.

### GetAllocationReferrerIdsOk

`func (o *WarehouseContact) GetAllocationReferrerIdsOk() (*[]string, bool)`

GetAllocationReferrerIdsOk returns a tuple with the AllocationReferrerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationReferrerIds

`func (o *WarehouseContact) SetAllocationReferrerIds(v []string)`

SetAllocationReferrerIds sets AllocationReferrerIds field to given value.

### HasAllocationReferrerIds

`func (o *WarehouseContact) HasAllocationReferrerIds() bool`

HasAllocationReferrerIds returns a boolean if a field has been set.

### GetAddress

`func (o *WarehouseContact) GetAddress() WarehouseContactAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WarehouseContact) GetAddressOk() (*WarehouseContactAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WarehouseContact) SetAddress(v WarehouseContactAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *WarehouseContact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


