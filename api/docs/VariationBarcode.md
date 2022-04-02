# VariationBarcode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BarcodeId** | Pointer to **int32** |  | [optional] 
**VariationId** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVariationBarcode

`func NewVariationBarcode() *VariationBarcode`

NewVariationBarcode instantiates a new VariationBarcode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationBarcodeWithDefaults

`func NewVariationBarcodeWithDefaults() *VariationBarcode`

NewVariationBarcodeWithDefaults instantiates a new VariationBarcode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBarcodeId

`func (o *VariationBarcode) GetBarcodeId() int32`

GetBarcodeId returns the BarcodeId field if non-nil, zero value otherwise.

### GetBarcodeIdOk

`func (o *VariationBarcode) GetBarcodeIdOk() (*int32, bool)`

GetBarcodeIdOk returns a tuple with the BarcodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcodeId

`func (o *VariationBarcode) SetBarcodeId(v int32)`

SetBarcodeId sets BarcodeId field to given value.

### HasBarcodeId

`func (o *VariationBarcode) HasBarcodeId() bool`

HasBarcodeId returns a boolean if a field has been set.

### GetVariationId

`func (o *VariationBarcode) GetVariationId() int32`

GetVariationId returns the VariationId field if non-nil, zero value otherwise.

### GetVariationIdOk

`func (o *VariationBarcode) GetVariationIdOk() (*int32, bool)`

GetVariationIdOk returns a tuple with the VariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationId

`func (o *VariationBarcode) SetVariationId(v int32)`

SetVariationId sets VariationId field to given value.

### HasVariationId

`func (o *VariationBarcode) HasVariationId() bool`

HasVariationId returns a boolean if a field has been set.

### GetCode

`func (o *VariationBarcode) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VariationBarcode) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VariationBarcode) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VariationBarcode) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariationBarcode) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariationBarcode) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariationBarcode) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VariationBarcode) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


