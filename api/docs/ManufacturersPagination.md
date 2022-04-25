# ManufacturersPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** |  | [optional] 
**TotalsCount** | Pointer to **int32** |  | [optional] 
**IsLastPage** | Pointer to **bool** |  | [optional] 
**LastPageNumber** | Pointer to **int32** |  | [optional] 
**FirstOnPage** | Pointer to **int32** |  | [optional] 
**LastOnPage** | Pointer to **int32** |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Entries** | Pointer to [**[]Manufacturer**](Manufacturer.md) |  | [optional] 

## Methods

### NewManufacturersPagination

`func NewManufacturersPagination() *ManufacturersPagination`

NewManufacturersPagination instantiates a new ManufacturersPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManufacturersPaginationWithDefaults

`func NewManufacturersPaginationWithDefaults() *ManufacturersPagination`

NewManufacturersPaginationWithDefaults instantiates a new ManufacturersPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ManufacturersPagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ManufacturersPagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ManufacturersPagination) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ManufacturersPagination) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalsCount

`func (o *ManufacturersPagination) GetTotalsCount() int32`

GetTotalsCount returns the TotalsCount field if non-nil, zero value otherwise.

### GetTotalsCountOk

`func (o *ManufacturersPagination) GetTotalsCountOk() (*int32, bool)`

GetTotalsCountOk returns a tuple with the TotalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsCount

`func (o *ManufacturersPagination) SetTotalsCount(v int32)`

SetTotalsCount sets TotalsCount field to given value.

### HasTotalsCount

`func (o *ManufacturersPagination) HasTotalsCount() bool`

HasTotalsCount returns a boolean if a field has been set.

### GetIsLastPage

`func (o *ManufacturersPagination) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *ManufacturersPagination) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *ManufacturersPagination) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *ManufacturersPagination) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetLastPageNumber

`func (o *ManufacturersPagination) GetLastPageNumber() int32`

GetLastPageNumber returns the LastPageNumber field if non-nil, zero value otherwise.

### GetLastPageNumberOk

`func (o *ManufacturersPagination) GetLastPageNumberOk() (*int32, bool)`

GetLastPageNumberOk returns a tuple with the LastPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageNumber

`func (o *ManufacturersPagination) SetLastPageNumber(v int32)`

SetLastPageNumber sets LastPageNumber field to given value.

### HasLastPageNumber

`func (o *ManufacturersPagination) HasLastPageNumber() bool`

HasLastPageNumber returns a boolean if a field has been set.

### GetFirstOnPage

`func (o *ManufacturersPagination) GetFirstOnPage() int32`

GetFirstOnPage returns the FirstOnPage field if non-nil, zero value otherwise.

### GetFirstOnPageOk

`func (o *ManufacturersPagination) GetFirstOnPageOk() (*int32, bool)`

GetFirstOnPageOk returns a tuple with the FirstOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOnPage

`func (o *ManufacturersPagination) SetFirstOnPage(v int32)`

SetFirstOnPage sets FirstOnPage field to given value.

### HasFirstOnPage

`func (o *ManufacturersPagination) HasFirstOnPage() bool`

HasFirstOnPage returns a boolean if a field has been set.

### GetLastOnPage

`func (o *ManufacturersPagination) GetLastOnPage() int32`

GetLastOnPage returns the LastOnPage field if non-nil, zero value otherwise.

### GetLastOnPageOk

`func (o *ManufacturersPagination) GetLastOnPageOk() (*int32, bool)`

GetLastOnPageOk returns a tuple with the LastOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnPage

`func (o *ManufacturersPagination) SetLastOnPage(v int32)`

SetLastOnPage sets LastOnPage field to given value.

### HasLastOnPage

`func (o *ManufacturersPagination) HasLastOnPage() bool`

HasLastOnPage returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *ManufacturersPagination) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ManufacturersPagination) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ManufacturersPagination) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *ManufacturersPagination) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEntries

`func (o *ManufacturersPagination) GetEntries() []Manufacturer`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *ManufacturersPagination) GetEntriesOk() (*[]Manufacturer, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *ManufacturersPagination) SetEntries(v []Manufacturer)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *ManufacturersPagination) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


