# WarehouseContactAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name1** | Pointer to **string** |  | [optional] 
**Address1** | Pointer to **string** |  | [optional] 
**Address2** | Pointer to **string** |  | [optional] 
**PostalCode** | Pointer to **string** |  | [optional] 
**Town** | Pointer to **string** |  | [optional] 
**CountryId** | Pointer to **int32** |  | [optional] 
**Options** | Pointer to [**[]WarehouseContactAddressOptions**](WarehouseContactAddressOptions.md) |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ContactPerson** | Pointer to **string** |  | [optional] 
**TaxIdNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewWarehouseContactAddress

`func NewWarehouseContactAddress() *WarehouseContactAddress`

NewWarehouseContactAddress instantiates a new WarehouseContactAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarehouseContactAddressWithDefaults

`func NewWarehouseContactAddressWithDefaults() *WarehouseContactAddress`

NewWarehouseContactAddressWithDefaults instantiates a new WarehouseContactAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName1

`func (o *WarehouseContactAddress) GetName1() string`

GetName1 returns the Name1 field if non-nil, zero value otherwise.

### GetName1Ok

`func (o *WarehouseContactAddress) GetName1Ok() (*string, bool)`

GetName1Ok returns a tuple with the Name1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName1

`func (o *WarehouseContactAddress) SetName1(v string)`

SetName1 sets Name1 field to given value.

### HasName1

`func (o *WarehouseContactAddress) HasName1() bool`

HasName1 returns a boolean if a field has been set.

### GetAddress1

`func (o *WarehouseContactAddress) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *WarehouseContactAddress) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *WarehouseContactAddress) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *WarehouseContactAddress) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *WarehouseContactAddress) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *WarehouseContactAddress) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *WarehouseContactAddress) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *WarehouseContactAddress) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetPostalCode

`func (o *WarehouseContactAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *WarehouseContactAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *WarehouseContactAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *WarehouseContactAddress) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetTown

`func (o *WarehouseContactAddress) GetTown() string`

GetTown returns the Town field if non-nil, zero value otherwise.

### GetTownOk

`func (o *WarehouseContactAddress) GetTownOk() (*string, bool)`

GetTownOk returns a tuple with the Town field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTown

`func (o *WarehouseContactAddress) SetTown(v string)`

SetTown sets Town field to given value.

### HasTown

`func (o *WarehouseContactAddress) HasTown() bool`

HasTown returns a boolean if a field has been set.

### GetCountryId

`func (o *WarehouseContactAddress) GetCountryId() int32`

GetCountryId returns the CountryId field if non-nil, zero value otherwise.

### GetCountryIdOk

`func (o *WarehouseContactAddress) GetCountryIdOk() (*int32, bool)`

GetCountryIdOk returns a tuple with the CountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryId

`func (o *WarehouseContactAddress) SetCountryId(v int32)`

SetCountryId sets CountryId field to given value.

### HasCountryId

`func (o *WarehouseContactAddress) HasCountryId() bool`

HasCountryId returns a boolean if a field has been set.

### GetOptions

`func (o *WarehouseContactAddress) GetOptions() []WarehouseContactAddressOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *WarehouseContactAddress) GetOptionsOk() (*[]WarehouseContactAddressOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *WarehouseContactAddress) SetOptions(v []WarehouseContactAddressOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *WarehouseContactAddress) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTitle

`func (o *WarehouseContactAddress) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarehouseContactAddress) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarehouseContactAddress) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WarehouseContactAddress) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContactPerson

`func (o *WarehouseContactAddress) GetContactPerson() string`

GetContactPerson returns the ContactPerson field if non-nil, zero value otherwise.

### GetContactPersonOk

`func (o *WarehouseContactAddress) GetContactPersonOk() (*string, bool)`

GetContactPersonOk returns a tuple with the ContactPerson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPerson

`func (o *WarehouseContactAddress) SetContactPerson(v string)`

SetContactPerson sets ContactPerson field to given value.

### HasContactPerson

`func (o *WarehouseContactAddress) HasContactPerson() bool`

HasContactPerson returns a boolean if a field has been set.

### GetTaxIdNumber

`func (o *WarehouseContactAddress) GetTaxIdNumber() string`

GetTaxIdNumber returns the TaxIdNumber field if non-nil, zero value otherwise.

### GetTaxIdNumberOk

`func (o *WarehouseContactAddress) GetTaxIdNumberOk() (*string, bool)`

GetTaxIdNumberOk returns a tuple with the TaxIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIdNumber

`func (o *WarehouseContactAddress) SetTaxIdNumber(v string)`

SetTaxIdNumber sets TaxIdNumber field to given value.

### HasTaxIdNumber

`func (o *WarehouseContactAddress) HasTaxIdNumber() bool`

HasTaxIdNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


