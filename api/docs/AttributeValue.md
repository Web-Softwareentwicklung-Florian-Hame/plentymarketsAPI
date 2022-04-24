# AttributeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AttributeId** | Pointer to **int32** |  | [optional] 
**BackendName** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**AmazonValue** | Pointer to **string** |  | [optional] 
**OttoValue** | Pointer to **string** |  | [optional] 
**NeckermannAtEpValue** | Pointer to **string** |  | [optional] 
**LaRedouteValue** | Pointer to **string** |  | [optional] 
**TracdelightValue** | Pointer to **string** |  | [optional] 
**PercentageDistribution** | Pointer to **int32** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Attribute** | Pointer to [**Attribute**](Attribute.md) |  | [optional] 

## Methods

### NewAttributeValue

`func NewAttributeValue() *AttributeValue`

NewAttributeValue instantiates a new AttributeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeValueWithDefaults

`func NewAttributeValueWithDefaults() *AttributeValue`

NewAttributeValueWithDefaults instantiates a new AttributeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AttributeValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributeValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributeValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AttributeValue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributeId

`func (o *AttributeValue) GetAttributeId() int32`

GetAttributeId returns the AttributeId field if non-nil, zero value otherwise.

### GetAttributeIdOk

`func (o *AttributeValue) GetAttributeIdOk() (*int32, bool)`

GetAttributeIdOk returns a tuple with the AttributeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeId

`func (o *AttributeValue) SetAttributeId(v int32)`

SetAttributeId sets AttributeId field to given value.

### HasAttributeId

`func (o *AttributeValue) HasAttributeId() bool`

HasAttributeId returns a boolean if a field has been set.

### GetBackendName

`func (o *AttributeValue) GetBackendName() string`

GetBackendName returns the BackendName field if non-nil, zero value otherwise.

### GetBackendNameOk

`func (o *AttributeValue) GetBackendNameOk() (*string, bool)`

GetBackendNameOk returns a tuple with the BackendName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendName

`func (o *AttributeValue) SetBackendName(v string)`

SetBackendName sets BackendName field to given value.

### HasBackendName

`func (o *AttributeValue) HasBackendName() bool`

HasBackendName returns a boolean if a field has been set.

### GetPosition

`func (o *AttributeValue) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AttributeValue) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AttributeValue) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AttributeValue) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetImage

`func (o *AttributeValue) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AttributeValue) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AttributeValue) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AttributeValue) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetComment

`func (o *AttributeValue) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AttributeValue) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AttributeValue) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AttributeValue) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetAmazonValue

`func (o *AttributeValue) GetAmazonValue() string`

GetAmazonValue returns the AmazonValue field if non-nil, zero value otherwise.

### GetAmazonValueOk

`func (o *AttributeValue) GetAmazonValueOk() (*string, bool)`

GetAmazonValueOk returns a tuple with the AmazonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonValue

`func (o *AttributeValue) SetAmazonValue(v string)`

SetAmazonValue sets AmazonValue field to given value.

### HasAmazonValue

`func (o *AttributeValue) HasAmazonValue() bool`

HasAmazonValue returns a boolean if a field has been set.

### GetOttoValue

`func (o *AttributeValue) GetOttoValue() string`

GetOttoValue returns the OttoValue field if non-nil, zero value otherwise.

### GetOttoValueOk

`func (o *AttributeValue) GetOttoValueOk() (*string, bool)`

GetOttoValueOk returns a tuple with the OttoValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoValue

`func (o *AttributeValue) SetOttoValue(v string)`

SetOttoValue sets OttoValue field to given value.

### HasOttoValue

`func (o *AttributeValue) HasOttoValue() bool`

HasOttoValue returns a boolean if a field has been set.

### GetNeckermannAtEpValue

`func (o *AttributeValue) GetNeckermannAtEpValue() string`

GetNeckermannAtEpValue returns the NeckermannAtEpValue field if non-nil, zero value otherwise.

### GetNeckermannAtEpValueOk

`func (o *AttributeValue) GetNeckermannAtEpValueOk() (*string, bool)`

GetNeckermannAtEpValueOk returns a tuple with the NeckermannAtEpValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeckermannAtEpValue

`func (o *AttributeValue) SetNeckermannAtEpValue(v string)`

SetNeckermannAtEpValue sets NeckermannAtEpValue field to given value.

### HasNeckermannAtEpValue

`func (o *AttributeValue) HasNeckermannAtEpValue() bool`

HasNeckermannAtEpValue returns a boolean if a field has been set.

### GetLaRedouteValue

`func (o *AttributeValue) GetLaRedouteValue() string`

GetLaRedouteValue returns the LaRedouteValue field if non-nil, zero value otherwise.

### GetLaRedouteValueOk

`func (o *AttributeValue) GetLaRedouteValueOk() (*string, bool)`

GetLaRedouteValueOk returns a tuple with the LaRedouteValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaRedouteValue

`func (o *AttributeValue) SetLaRedouteValue(v string)`

SetLaRedouteValue sets LaRedouteValue field to given value.

### HasLaRedouteValue

`func (o *AttributeValue) HasLaRedouteValue() bool`

HasLaRedouteValue returns a boolean if a field has been set.

### GetTracdelightValue

`func (o *AttributeValue) GetTracdelightValue() string`

GetTracdelightValue returns the TracdelightValue field if non-nil, zero value otherwise.

### GetTracdelightValueOk

`func (o *AttributeValue) GetTracdelightValueOk() (*string, bool)`

GetTracdelightValueOk returns a tuple with the TracdelightValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracdelightValue

`func (o *AttributeValue) SetTracdelightValue(v string)`

SetTracdelightValue sets TracdelightValue field to given value.

### HasTracdelightValue

`func (o *AttributeValue) HasTracdelightValue() bool`

HasTracdelightValue returns a boolean if a field has been set.

### GetPercentageDistribution

`func (o *AttributeValue) GetPercentageDistribution() int32`

GetPercentageDistribution returns the PercentageDistribution field if non-nil, zero value otherwise.

### GetPercentageDistributionOk

`func (o *AttributeValue) GetPercentageDistributionOk() (*int32, bool)`

GetPercentageDistributionOk returns a tuple with the PercentageDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageDistribution

`func (o *AttributeValue) SetPercentageDistribution(v int32)`

SetPercentageDistribution sets PercentageDistribution field to given value.

### HasPercentageDistribution

`func (o *AttributeValue) HasPercentageDistribution() bool`

HasPercentageDistribution returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AttributeValue) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AttributeValue) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AttributeValue) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AttributeValue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAttribute

`func (o *AttributeValue) GetAttribute() Attribute`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *AttributeValue) GetAttributeOk() (*Attribute, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *AttributeValue) SetAttribute(v Attribute)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *AttributeValue) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


