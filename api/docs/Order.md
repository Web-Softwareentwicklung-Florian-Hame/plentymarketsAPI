# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ReferrerId** | Pointer to **int32** |  | [optional] 
**RoundTotalsOnly** | Pointer to **bool** |  | [optional] 
**NumberOfDecimals** | Pointer to **int32** |  | [optional] 
**ContactReceiver** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactSender** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**WarehouseReceiver** | Pointer to [**WarehouseContact**](WarehouseContact.md) |  | [optional] 
**WarehouseSender** | Pointer to [**WarehouseContact**](WarehouseContact.md) |  | [optional] 
**PlentyId** | Pointer to **int32** |  | [optional] 
**Typeid** | Pointer to **int32** |  | [optional] 
**LockStatus** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**StatusId** | Pointer to **float32** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Relations** | Pointer to [**[]OrderRelation**](OrderRelation.md) |  | [optional] 
**Addresses** | Pointer to [**[]OrderAddress**](OrderAddress.md) |  | [optional] 
**Location** | Pointer to [**OrderLocation**](OrderLocation.md) |  | [optional] 
**Documents** | Pointer to [**[]OrderDocument**](OrderDocument.md) |  | [optional] 
**OrderItems** | Pointer to [**[]OrderItem**](OrderItem.md) |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferrerId

`func (o *Order) GetReferrerId() int32`

GetReferrerId returns the ReferrerId field if non-nil, zero value otherwise.

### GetReferrerIdOk

`func (o *Order) GetReferrerIdOk() (*int32, bool)`

GetReferrerIdOk returns a tuple with the ReferrerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrerId

`func (o *Order) SetReferrerId(v int32)`

SetReferrerId sets ReferrerId field to given value.

### HasReferrerId

`func (o *Order) HasReferrerId() bool`

HasReferrerId returns a boolean if a field has been set.

### GetRoundTotalsOnly

`func (o *Order) GetRoundTotalsOnly() bool`

GetRoundTotalsOnly returns the RoundTotalsOnly field if non-nil, zero value otherwise.

### GetRoundTotalsOnlyOk

`func (o *Order) GetRoundTotalsOnlyOk() (*bool, bool)`

GetRoundTotalsOnlyOk returns a tuple with the RoundTotalsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTotalsOnly

`func (o *Order) SetRoundTotalsOnly(v bool)`

SetRoundTotalsOnly sets RoundTotalsOnly field to given value.

### HasRoundTotalsOnly

`func (o *Order) HasRoundTotalsOnly() bool`

HasRoundTotalsOnly returns a boolean if a field has been set.

### GetNumberOfDecimals

`func (o *Order) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *Order) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *Order) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *Order) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### GetContactReceiver

`func (o *Order) GetContactReceiver() Contact`

GetContactReceiver returns the ContactReceiver field if non-nil, zero value otherwise.

### GetContactReceiverOk

`func (o *Order) GetContactReceiverOk() (*Contact, bool)`

GetContactReceiverOk returns a tuple with the ContactReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactReceiver

`func (o *Order) SetContactReceiver(v Contact)`

SetContactReceiver sets ContactReceiver field to given value.

### HasContactReceiver

`func (o *Order) HasContactReceiver() bool`

HasContactReceiver returns a boolean if a field has been set.

### GetContactSender

`func (o *Order) GetContactSender() Contact`

GetContactSender returns the ContactSender field if non-nil, zero value otherwise.

### GetContactSenderOk

`func (o *Order) GetContactSenderOk() (*Contact, bool)`

GetContactSenderOk returns a tuple with the ContactSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSender

`func (o *Order) SetContactSender(v Contact)`

SetContactSender sets ContactSender field to given value.

### HasContactSender

`func (o *Order) HasContactSender() bool`

HasContactSender returns a boolean if a field has been set.

### GetWarehouseReceiver

`func (o *Order) GetWarehouseReceiver() WarehouseContact`

GetWarehouseReceiver returns the WarehouseReceiver field if non-nil, zero value otherwise.

### GetWarehouseReceiverOk

`func (o *Order) GetWarehouseReceiverOk() (*WarehouseContact, bool)`

GetWarehouseReceiverOk returns a tuple with the WarehouseReceiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseReceiver

`func (o *Order) SetWarehouseReceiver(v WarehouseContact)`

SetWarehouseReceiver sets WarehouseReceiver field to given value.

### HasWarehouseReceiver

`func (o *Order) HasWarehouseReceiver() bool`

HasWarehouseReceiver returns a boolean if a field has been set.

### GetWarehouseSender

`func (o *Order) GetWarehouseSender() WarehouseContact`

GetWarehouseSender returns the WarehouseSender field if non-nil, zero value otherwise.

### GetWarehouseSenderOk

`func (o *Order) GetWarehouseSenderOk() (*WarehouseContact, bool)`

GetWarehouseSenderOk returns a tuple with the WarehouseSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseSender

`func (o *Order) SetWarehouseSender(v WarehouseContact)`

SetWarehouseSender sets WarehouseSender field to given value.

### HasWarehouseSender

`func (o *Order) HasWarehouseSender() bool`

HasWarehouseSender returns a boolean if a field has been set.

### GetPlentyId

`func (o *Order) GetPlentyId() int32`

GetPlentyId returns the PlentyId field if non-nil, zero value otherwise.

### GetPlentyIdOk

`func (o *Order) GetPlentyIdOk() (*int32, bool)`

GetPlentyIdOk returns a tuple with the PlentyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlentyId

`func (o *Order) SetPlentyId(v int32)`

SetPlentyId sets PlentyId field to given value.

### HasPlentyId

`func (o *Order) HasPlentyId() bool`

HasPlentyId returns a boolean if a field has been set.

### GetTypeid

`func (o *Order) GetTypeid() int32`

GetTypeid returns the Typeid field if non-nil, zero value otherwise.

### GetTypeidOk

`func (o *Order) GetTypeidOk() (*int32, bool)`

GetTypeidOk returns a tuple with the Typeid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeid

`func (o *Order) SetTypeid(v int32)`

SetTypeid sets Typeid field to given value.

### HasTypeid

`func (o *Order) HasTypeid() bool`

HasTypeid returns a boolean if a field has been set.

### GetLockStatus

`func (o *Order) GetLockStatus() string`

GetLockStatus returns the LockStatus field if non-nil, zero value otherwise.

### GetLockStatusOk

`func (o *Order) GetLockStatusOk() (*string, bool)`

GetLockStatusOk returns a tuple with the LockStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockStatus

`func (o *Order) SetLockStatus(v string)`

SetLockStatus sets LockStatus field to given value.

### HasLockStatus

`func (o *Order) HasLockStatus() bool`

HasLockStatus returns a boolean if a field has been set.

### GetLocationId

`func (o *Order) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Order) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Order) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Order) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Order) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Order) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatusId

`func (o *Order) GetStatusId() float32`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *Order) GetStatusIdOk() (*float32, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *Order) SetStatusId(v float32)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *Order) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetOwnerId

`func (o *Order) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Order) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Order) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Order) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetRelations

`func (o *Order) GetRelations() []OrderRelation`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *Order) GetRelationsOk() (*[]OrderRelation, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *Order) SetRelations(v []OrderRelation)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *Order) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetAddresses

`func (o *Order) GetAddresses() []OrderAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Order) GetAddressesOk() (*[]OrderAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Order) SetAddresses(v []OrderAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *Order) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetLocation

`func (o *Order) GetLocation() OrderLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Order) GetLocationOk() (*OrderLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Order) SetLocation(v OrderLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Order) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDocuments

`func (o *Order) GetDocuments() []OrderDocument`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *Order) GetDocumentsOk() (*[]OrderDocument, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *Order) SetDocuments(v []OrderDocument)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *Order) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetOrderItems

`func (o *Order) GetOrderItems() []OrderItem`

GetOrderItems returns the OrderItems field if non-nil, zero value otherwise.

### GetOrderItemsOk

`func (o *Order) GetOrderItemsOk() (*[]OrderItem, bool)`

GetOrderItemsOk returns a tuple with the OrderItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderItems

`func (o *Order) SetOrderItems(v []OrderItem)`

SetOrderItems sets OrderItems field to given value.

### HasOrderItems

`func (o *Order) HasOrderItems() bool`

HasOrderItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


