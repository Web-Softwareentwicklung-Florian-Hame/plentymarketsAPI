# InlineResponse200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page of the response | [optional] 
**TotalsCount** | Pointer to **int32** | The total number of entries in the response | [optional] 
**IsLastPage** | Pointer to **bool** | Flag that indicates if the page shown is the last page of the response | [optional] 
**LastPageNumber** | Pointer to **int32** | The last page number | [optional] 
**FirstOnPage** | Pointer to **int32** | The index of the first item of the current page result | [optional] 
**LastOnPage** | Pointer to **int32** | The index of the last item of the current page result | [optional] 
**ItemsPerPage** | Pointer to **int32** | The requested amount of items per result page | [optional] 
**Entries** | Pointer to [**[]PimVariation**](PimVariation.md) | List of Variation | [optional] 

## Methods

### NewInlineResponse200

`func NewInlineResponse200() *InlineResponse200`

NewInlineResponse200 instantiates a new InlineResponse200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200WithDefaults

`func NewInlineResponse200WithDefaults() *InlineResponse200`

NewInlineResponse200WithDefaults instantiates a new InlineResponse200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *InlineResponse200) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *InlineResponse200) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *InlineResponse200) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *InlineResponse200) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetTotalsCount

`func (o *InlineResponse200) GetTotalsCount() int32`

GetTotalsCount returns the TotalsCount field if non-nil, zero value otherwise.

### GetTotalsCountOk

`func (o *InlineResponse200) GetTotalsCountOk() (*int32, bool)`

GetTotalsCountOk returns a tuple with the TotalsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalsCount

`func (o *InlineResponse200) SetTotalsCount(v int32)`

SetTotalsCount sets TotalsCount field to given value.

### HasTotalsCount

`func (o *InlineResponse200) HasTotalsCount() bool`

HasTotalsCount returns a boolean if a field has been set.

### GetIsLastPage

`func (o *InlineResponse200) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *InlineResponse200) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *InlineResponse200) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *InlineResponse200) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetLastPageNumber

`func (o *InlineResponse200) GetLastPageNumber() int32`

GetLastPageNumber returns the LastPageNumber field if non-nil, zero value otherwise.

### GetLastPageNumberOk

`func (o *InlineResponse200) GetLastPageNumberOk() (*int32, bool)`

GetLastPageNumberOk returns a tuple with the LastPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPageNumber

`func (o *InlineResponse200) SetLastPageNumber(v int32)`

SetLastPageNumber sets LastPageNumber field to given value.

### HasLastPageNumber

`func (o *InlineResponse200) HasLastPageNumber() bool`

HasLastPageNumber returns a boolean if a field has been set.

### GetFirstOnPage

`func (o *InlineResponse200) GetFirstOnPage() int32`

GetFirstOnPage returns the FirstOnPage field if non-nil, zero value otherwise.

### GetFirstOnPageOk

`func (o *InlineResponse200) GetFirstOnPageOk() (*int32, bool)`

GetFirstOnPageOk returns a tuple with the FirstOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstOnPage

`func (o *InlineResponse200) SetFirstOnPage(v int32)`

SetFirstOnPage sets FirstOnPage field to given value.

### HasFirstOnPage

`func (o *InlineResponse200) HasFirstOnPage() bool`

HasFirstOnPage returns a boolean if a field has been set.

### GetLastOnPage

`func (o *InlineResponse200) GetLastOnPage() int32`

GetLastOnPage returns the LastOnPage field if non-nil, zero value otherwise.

### GetLastOnPageOk

`func (o *InlineResponse200) GetLastOnPageOk() (*int32, bool)`

GetLastOnPageOk returns a tuple with the LastOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOnPage

`func (o *InlineResponse200) SetLastOnPage(v int32)`

SetLastOnPage sets LastOnPage field to given value.

### HasLastOnPage

`func (o *InlineResponse200) HasLastOnPage() bool`

HasLastOnPage returns a boolean if a field has been set.

### GetItemsPerPage

`func (o *InlineResponse200) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *InlineResponse200) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *InlineResponse200) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### HasItemsPerPage

`func (o *InlineResponse200) HasItemsPerPage() bool`

HasItemsPerPage returns a boolean if a field has been set.

### GetEntries

`func (o *InlineResponse200) GetEntries() []PimVariation`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *InlineResponse200) GetEntriesOk() (*[]PimVariation, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *InlineResponse200) SetEntries(v []PimVariation)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *InlineResponse200) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


