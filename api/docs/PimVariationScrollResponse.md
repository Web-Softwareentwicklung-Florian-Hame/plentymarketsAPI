# PimVariationScrollResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | The cursor for the next page of results. | [optional] 
**Entries** | Pointer to [**[]PimVariation**](PimVariation.md) | List of Variation | [optional] 

## Methods

### NewPimVariationScrollResponse

`func NewPimVariationScrollResponse() *PimVariationScrollResponse`

NewPimVariationScrollResponse instantiates a new PimVariationScrollResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPimVariationScrollResponseWithDefaults

`func NewPimVariationScrollResponseWithDefaults() *PimVariationScrollResponse`

NewPimVariationScrollResponseWithDefaults instantiates a new PimVariationScrollResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *PimVariationScrollResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PimVariationScrollResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PimVariationScrollResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PimVariationScrollResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetEntries

`func (o *PimVariationScrollResponse) GetEntries() []PimVariation`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *PimVariationScrollResponse) GetEntriesOk() (*[]PimVariation, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *PimVariationScrollResponse) SetEntries(v []PimVariation)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *PimVariationScrollResponse) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


