# AttributePagination

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
**Entries** | Pointer to [**[]Attribute**](Attribute.md) |  | [optional] 

## Methods

### NewAttributePagination

`func NewAttributePagination() *AttributePagination`

NewAttributePagination instantiates a new AttributePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributePaginationWithDefaults

`func NewAttributePaginationWithDefaults() *AttributePagination`

NewAttributePaginationWithDefaults instantiates a new AttributePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *AttributePagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AttributePagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AttributePagination) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *AttributePagination) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalsCount

`func (o *AttributePagination) GetTotalsCount() int32`

GetTotalsCount returns the TotalsCount field if non-nil, zero value otherwise.

### GetTotalsCountOk

`func (o *AttributePagination) GetTotalsCountOk() (*int32, bool)`

GetTotalsCountOk returns a tuple with the TotalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsCount

`func (o *AttributePagination) SetTotalsCount(v int32)`

SetTotalsCount sets TotalsCount field to given value.

### HasTotalsCount

`func (o *AttributePagination) HasTotalsCount() bool`

HasTotalsCount returns a boolean if a field has been set.

### GetIsLastPage

`func (o *AttributePagination) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *AttributePagination) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *AttributePagination) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *AttributePagination) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetLastPageNumber

`func (o *AttributePagination) GetLastPageNumber() int32`

GetLastPageNumber returns the LastPageNumber field if non-nil, zero value otherwise.

### GetLastPageNumberOk

`func (o *AttributePagination) GetLastPageNumberOk() (*int32, bool)`

GetLastPageNumberOk returns a tuple with the LastPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageNumber

`func (o *AttributePagination) SetLastPageNumber(v int32)`

SetLastPageNumber sets LastPageNumber field to given value.

### HasLastPageNumber

`func (o *AttributePagination) HasLastPageNumber() bool`

HasLastPageNumber returns a boolean if a field has been set.

### GetFirstOnPage

`func (o *AttributePagination) GetFirstOnPage() int32`

GetFirstOnPage returns the FirstOnPage field if non-nil, zero value otherwise.

### GetFirstOnPageOk

`func (o *AttributePagination) GetFirstOnPageOk() (*int32, bool)`

GetFirstOnPageOk returns a tuple with the FirstOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOnPage

`func (o *AttributePagination) SetFirstOnPage(v int32)`

SetFirstOnPage sets FirstOnPage field to given value.

### HasFirstOnPage

`func (o *AttributePagination) HasFirstOnPage() bool`

HasFirstOnPage returns a boolean if a field has been set.

### GetLastOnPage

`func (o *AttributePagination) GetLastOnPage() int32`

GetLastOnPage returns the LastOnPage field if non-nil, zero value otherwise.

### GetLastOnPageOk

`func (o *AttributePagination) GetLastOnPageOk() (*int32, bool)`

GetLastOnPageOk returns a tuple with the LastOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnPage

`func (o *AttributePagination) SetLastOnPage(v int32)`

SetLastOnPage sets LastOnPage field to given value.

### HasLastOnPage

`func (o *AttributePagination) HasLastOnPage() bool`

HasLastOnPage returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *AttributePagination) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *AttributePagination) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *AttributePagination) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *AttributePagination) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEntries

`func (o *AttributePagination) GetEntries() []Attribute`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *AttributePagination) GetEntriesOk() (*[]Attribute, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *AttributePagination) SetEntries(v []Attribute)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *AttributePagination) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


