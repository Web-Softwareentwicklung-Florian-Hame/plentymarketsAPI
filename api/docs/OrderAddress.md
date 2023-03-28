# OrderAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Gender** | Pointer to **string** |  | [optional] 
**Name1** | Pointer to **string** |  | [optional] 
**Name2** | Pointer to **string** |  | [optional] 
**Name3** | Pointer to **string** |  | [optional] 
**Name4** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**Address3** | Pointer to **string** |  | [optional] 
**Address4** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Town** | Pointer to **string** |  | [optional] 
**CountryId** | Pointer to **int32** |  | [optional] 
**StateId** | Pointer to **int32** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**CheckedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ContactPerson** | Pointer to **string** |  | [optional] 
**TaxIdNumber** | Pointer to **string** |  | [optional] 
**Pivot** | Pointer to [**OrderAddressPivot**](OrderAddressPivot.md) |  | [optional] 

## Methods

### NewOrderAddress

`func NewOrderAddress() *OrderAddress`

NewOrderAddress instantiates a new OrderAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderAddressWithDefaults

`func NewOrderAddressWithDefaults() *OrderAddress`

NewOrderAddressWithDefaults instantiates a new OrderAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderAddress) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderAddress) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderAddress) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OrderAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGender

`func (o *OrderAddress) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *OrderAddress) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *OrderAddress) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *OrderAddress) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetName1

`func (o *OrderAddress) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *OrderAddress) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *OrderAddress) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *OrderAddress) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetName2

`func (o *OrderAddress) GetName2() string`

GetName2 returns the Name2 field if non-nil, zero value otherwise.

### GetName2Ok

`func (o *OrderAddress) GetName2Ok() (*string, bool)`

GetName2Ok returns a tuple with the Name2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName2

`func (o *OrderAddress) SetName2(v string)`

SetName2 sets Name2 field to given value.

### HasName2

`func (o *OrderAddress) HasName2() bool`

HasName2 returns a boolean if a field has been set.

### GetName3

`func (o *OrderAddress) GetName3() string`

GetName3 returns the Name3 field if non-nil, zero value otherwise.

### GetName3Ok

`func (o *OrderAddress) GetName3Ok() (*string, bool)`

GetName3Ok returns a tuple with the Name3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName3

`func (o *OrderAddress) SetName3(v string)`

SetName3 sets Name3 field to given value.

### HasName3

`func (o *OrderAddress) HasName3() bool`

HasName3 returns a boolean if a field has been set.

### GetName4

`func (o *OrderAddress) GetName4() string`

GetName4 returns the Name4 field if non-nil, zero value otherwise.

### GetName4Ok

`func (o *OrderAddress) GetName4Ok() (*string, bool)`

GetName4Ok returns a tuple with the Name4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName4

`func (o *OrderAddress) SetName4(v string)`

SetName4 sets Name4 field to given value.

### HasName4

`func (o *OrderAddress) HasName4() bool`

HasName4 returns a boolean if a field has been set.

### GetAddress1

`func (o *OrderAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *OrderAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *OrderAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *OrderAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *OrderAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *OrderAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *OrderAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *OrderAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetAddress3

`func (o *OrderAddress) GetAddress3() string`

GetAddress3 returns the Address3 field if non-nil, zero value otherwise.

### GetAddress3Ok

`func (o *OrderAddress) GetAddress3Ok() (*string, bool)`

GetAddress3Ok returns a tuple with the Address3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress3

`func (o *OrderAddress) SetAddress3(v string)`

SetAddress3 sets Address3 field to given value.

### HasAddress3

`func (o *OrderAddress) HasAddress3() bool`

HasAddress3 returns a boolean if a field has been set.

### GetAddress4

`func (o *OrderAddress) GetAddress4() string`

GetAddress4 returns the Address4 field if non-nil, zero value otherwise.

### GetAddress4Ok

`func (o *OrderAddress) GetAddress4Ok() (*string, bool)`

GetAddress4Ok returns a tuple with the Address4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress4

`func (o *OrderAddress) SetAddress4(v string)`

SetAddress4 sets Address4 field to given value.

### HasAddress4

`func (o *OrderAddress) HasAddress4() bool`

HasAddress4 returns a boolean if a field has been set.

### GetPostalCode

`func (o *OrderAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *OrderAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *OrderAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *OrderAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetTown

`func (o *OrderAddress) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *OrderAddress) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *OrderAddress) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *OrderAddress) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetCountryId

`func (o *OrderAddress) GetCountryId() int32`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *OrderAddress) GetCountryIdOk() (*int32, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *OrderAddress) SetCountryId(v int32)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *OrderAddress) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetStateId

`func (o *OrderAddress) GetStateId() int32`

GetStateId returns the StateId field if non-nil, zero value otherwise.

### GetStateIdOk

`func (o *OrderAddress) GetStateIdOk() (*int32, bool)`

GetStateIdOk returns a tuple with the StateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateId

`func (o *OrderAddress) SetStateId(v int32)`

SetStateId sets StateId field to given value.

### HasStateId

`func (o *OrderAddress) HasStateId() bool`

HasStateId returns a boolean if a field has been set.

### GetReadOnly

`func (o *OrderAddress) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *OrderAddress) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *OrderAddress) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *OrderAddress) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetCheckedAt

`func (o *OrderAddress) GetCheckedAt() time.Time`

GetCheckedAt returns the CheckedAt field if non-nil, zero value otherwise.

### GetCheckedAtOk

`func (o *OrderAddress) GetCheckedAtOk() (*time.Time, bool)`

GetCheckedAtOk returns a tuple with the CheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckedAt

`func (o *OrderAddress) SetCheckedAt(v time.Time)`

SetCheckedAt sets CheckedAt field to given value.

### HasCheckedAt

`func (o *OrderAddress) HasCheckedAt() bool`

HasCheckedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OrderAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrderAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrderAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OrderAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OrderAddress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrderAddress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrderAddress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OrderAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTitle

`func (o *OrderAddress) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrderAddress) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrderAddress) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OrderAddress) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContactPerson

`func (o *OrderAddress) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *OrderAddress) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *OrderAddress) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *OrderAddress) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetTaxIdNumber

`func (o *OrderAddress) GetTaxIdNumber() string`

GetTaxIdNumber returns the TaxIdNumber field if non-nil, zero value otherwise.

### GetTaxIdNumberOk

`func (o *OrderAddress) GetTaxIdNumberOk() (*string, bool)`

GetTaxIdNumberOk returns a tuple with the TaxIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdNumber

`func (o *OrderAddress) SetTaxIdNumber(v string)`

SetTaxIdNumber sets TaxIdNumber field to given value.

### HasTaxIdNumber

`func (o *OrderAddress) HasTaxIdNumber() bool`

HasTaxIdNumber returns a boolean if a field has been set.

### GetPivot

`func (o *OrderAddress) GetPivot() OrderAddressPivot`

GetPivot returns the Pivot field if non-nil, zero value otherwise.

### GetPivotOk

`func (o *OrderAddress) GetPivotOk() (*OrderAddressPivot, bool)`

GetPivotOk returns a tuple with the Pivot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivot

`func (o *OrderAddress) SetPivot(v OrderAddressPivot)`

SetPivot sets Pivot field to given value.

### HasPivot

`func (o *OrderAddress) HasPivot() bool`

HasPivot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


