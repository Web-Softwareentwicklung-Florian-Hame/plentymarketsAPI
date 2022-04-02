# VariationPropertyProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Unit** | Pointer to **int32** |  | [optional] 
**PropertyGroupId** | Pointer to **int32** |  | [optional] 
**BackendName** | Pointer to **string** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**IsSearchable** | Pointer to **bool** |  | [optional] 
**IsOderProperty** | Pointer to **bool** |  | [optional] 
**IsShownOnItemPage** | Pointer to **bool** |  | [optional] 
**IsShownOnItemList** | Pointer to **bool** |  | [optional] 
**IsShownAtCheckout** | Pointer to **bool** |  | [optional] 
**IsShownInPdf** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Surcharge** | Pointer to **float32** |  | [optional] 
**IsShownAsAdditionalCosts** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVariationPropertyProperty

`func NewVariationPropertyProperty() *VariationPropertyProperty`

NewVariationPropertyProperty instantiates a new VariationPropertyProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationPropertyPropertyWithDefaults

`func NewVariationPropertyPropertyWithDefaults() *VariationPropertyProperty`

NewVariationPropertyPropertyWithDefaults instantiates a new VariationPropertyProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationPropertyProperty) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationPropertyProperty) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationPropertyProperty) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationPropertyProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *VariationPropertyProperty) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VariationPropertyProperty) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VariationPropertyProperty) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VariationPropertyProperty) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetUnit

`func (o *VariationPropertyProperty) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VariationPropertyProperty) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VariationPropertyProperty) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *VariationPropertyProperty) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetPropertyGroupId

`func (o *VariationPropertyProperty) GetPropertyGroupId() int32`

GetPropertyGroupId returns the PropertyGroupId field if non-nil, zero value otherwise.

### GetPropertyGroupIdOk

`func (o *VariationPropertyProperty) GetPropertyGroupIdOk() (*int32, bool)`

GetPropertyGroupIdOk returns a tuple with the PropertyGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyGroupId

`func (o *VariationPropertyProperty) SetPropertyGroupId(v int32)`

SetPropertyGroupId sets PropertyGroupId field to given value.

### HasPropertyGroupId

`func (o *VariationPropertyProperty) HasPropertyGroupId() bool`

HasPropertyGroupId returns a boolean if a field has been set.

### GetBackendName

`func (o *VariationPropertyProperty) GetBackendName() string`

GetBackendName returns the BackendName field if non-nil, zero value otherwise.

### GetBackendNameOk

`func (o *VariationPropertyProperty) GetBackendNameOk() (*string, bool)`

GetBackendNameOk returns a tuple with the BackendName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendName

`func (o *VariationPropertyProperty) SetBackendName(v string)`

SetBackendName sets BackendName field to given value.

### HasBackendName

`func (o *VariationPropertyProperty) HasBackendName() bool`

HasBackendName returns a boolean if a field has been set.

### GetValueType

`func (o *VariationPropertyProperty) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *VariationPropertyProperty) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *VariationPropertyProperty) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *VariationPropertyProperty) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetIsSearchable

`func (o *VariationPropertyProperty) GetIsSearchable() bool`

GetIsSearchable returns the IsSearchable field if non-nil, zero value otherwise.

### GetIsSearchableOk

`func (o *VariationPropertyProperty) GetIsSearchableOk() (*bool, bool)`

GetIsSearchableOk returns a tuple with the IsSearchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSearchable

`func (o *VariationPropertyProperty) SetIsSearchable(v bool)`

SetIsSearchable sets IsSearchable field to given value.

### HasIsSearchable

`func (o *VariationPropertyProperty) HasIsSearchable() bool`

HasIsSearchable returns a boolean if a field has been set.

### GetIsOderProperty

`func (o *VariationPropertyProperty) GetIsOderProperty() bool`

GetIsOderProperty returns the IsOderProperty field if non-nil, zero value otherwise.

### GetIsOderPropertyOk

`func (o *VariationPropertyProperty) GetIsOderPropertyOk() (*bool, bool)`

GetIsOderPropertyOk returns a tuple with the IsOderProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOderProperty

`func (o *VariationPropertyProperty) SetIsOderProperty(v bool)`

SetIsOderProperty sets IsOderProperty field to given value.

### HasIsOderProperty

`func (o *VariationPropertyProperty) HasIsOderProperty() bool`

HasIsOderProperty returns a boolean if a field has been set.

### GetIsShownOnItemPage

`func (o *VariationPropertyProperty) GetIsShownOnItemPage() bool`

GetIsShownOnItemPage returns the IsShownOnItemPage field if non-nil, zero value otherwise.

### GetIsShownOnItemPageOk

`func (o *VariationPropertyProperty) GetIsShownOnItemPageOk() (*bool, bool)`

GetIsShownOnItemPageOk returns a tuple with the IsShownOnItemPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownOnItemPage

`func (o *VariationPropertyProperty) SetIsShownOnItemPage(v bool)`

SetIsShownOnItemPage sets IsShownOnItemPage field to given value.

### HasIsShownOnItemPage

`func (o *VariationPropertyProperty) HasIsShownOnItemPage() bool`

HasIsShownOnItemPage returns a boolean if a field has been set.

### GetIsShownOnItemList

`func (o *VariationPropertyProperty) GetIsShownOnItemList() bool`

GetIsShownOnItemList returns the IsShownOnItemList field if non-nil, zero value otherwise.

### GetIsShownOnItemListOk

`func (o *VariationPropertyProperty) GetIsShownOnItemListOk() (*bool, bool)`

GetIsShownOnItemListOk returns a tuple with the IsShownOnItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownOnItemList

`func (o *VariationPropertyProperty) SetIsShownOnItemList(v bool)`

SetIsShownOnItemList sets IsShownOnItemList field to given value.

### HasIsShownOnItemList

`func (o *VariationPropertyProperty) HasIsShownOnItemList() bool`

HasIsShownOnItemList returns a boolean if a field has been set.

### GetIsShownAtCheckout

`func (o *VariationPropertyProperty) GetIsShownAtCheckout() bool`

GetIsShownAtCheckout returns the IsShownAtCheckout field if non-nil, zero value otherwise.

### GetIsShownAtCheckoutOk

`func (o *VariationPropertyProperty) GetIsShownAtCheckoutOk() (*bool, bool)`

GetIsShownAtCheckoutOk returns a tuple with the IsShownAtCheckout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownAtCheckout

`func (o *VariationPropertyProperty) SetIsShownAtCheckout(v bool)`

SetIsShownAtCheckout sets IsShownAtCheckout field to given value.

### HasIsShownAtCheckout

`func (o *VariationPropertyProperty) HasIsShownAtCheckout() bool`

HasIsShownAtCheckout returns a boolean if a field has been set.

### GetIsShownInPdf

`func (o *VariationPropertyProperty) GetIsShownInPdf() bool`

GetIsShownInPdf returns the IsShownInPdf field if non-nil, zero value otherwise.

### GetIsShownInPdfOk

`func (o *VariationPropertyProperty) GetIsShownInPdfOk() (*bool, bool)`

GetIsShownInPdfOk returns a tuple with the IsShownInPdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownInPdf

`func (o *VariationPropertyProperty) SetIsShownInPdf(v bool)`

SetIsShownInPdf sets IsShownInPdf field to given value.

### HasIsShownInPdf

`func (o *VariationPropertyProperty) HasIsShownInPdf() bool`

HasIsShownInPdf returns a boolean if a field has been set.

### GetComment

`func (o *VariationPropertyProperty) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VariationPropertyProperty) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VariationPropertyProperty) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VariationPropertyProperty) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetSurcharge

`func (o *VariationPropertyProperty) GetSurcharge() float32`

GetSurcharge returns the Surcharge field if non-nil, zero value otherwise.

### GetSurchargeOk

`func (o *VariationPropertyProperty) GetSurchargeOk() (*float32, bool)`

GetSurchargeOk returns a tuple with the Surcharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurcharge

`func (o *VariationPropertyProperty) SetSurcharge(v float32)`

SetSurcharge sets Surcharge field to given value.

### HasSurcharge

`func (o *VariationPropertyProperty) HasSurcharge() bool`

HasSurcharge returns a boolean if a field has been set.

### GetIsShownAsAdditionalCosts

`func (o *VariationPropertyProperty) GetIsShownAsAdditionalCosts() bool`

GetIsShownAsAdditionalCosts returns the IsShownAsAdditionalCosts field if non-nil, zero value otherwise.

### GetIsShownAsAdditionalCostsOk

`func (o *VariationPropertyProperty) GetIsShownAsAdditionalCostsOk() (*bool, bool)`

GetIsShownAsAdditionalCostsOk returns a tuple with the IsShownAsAdditionalCosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShownAsAdditionalCosts

`func (o *VariationPropertyProperty) SetIsShownAsAdditionalCosts(v bool)`

SetIsShownAsAdditionalCosts sets IsShownAsAdditionalCosts field to given value.

### HasIsShownAsAdditionalCosts

`func (o *VariationPropertyProperty) HasIsShownAsAdditionalCosts() bool`

HasIsShownAsAdditionalCosts returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VariationPropertyProperty) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariationPropertyProperty) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariationPropertyProperty) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VariationPropertyProperty) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


