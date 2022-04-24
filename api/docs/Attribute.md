# Attribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BackendName** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**IsSurchargePercental** | Pointer to **bool** |  | [optional] 
**IsLinkableToImage** | Pointer to **bool** |  | [optional] 
**AmazonAttribute** | Pointer to **string** |  | [optional] 
**FruugoAttribute** | Pointer to **string** |  | [optional] 
**PixmaniaAttribute** | Pointer to **string** |  | [optional] 
**OttoAttribute** | Pointer to **string** |  | [optional] 
**GoogleShoppingAttribute** | Pointer to **string** |  | [optional] 
**NeckermannAtEpAttribute** | Pointer to **int32** |  | [optional] 
**TypeOfSelectionInOnlineStore** | Pointer to **string** |  | [optional] 
**LaRedouteAttribute** | Pointer to **int32** |  | [optional] 
**IsGroupable** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Values** | Pointer to [**[]AttributeValue**](AttributeValue.md) | only accessible for get attributes call when with flag is set | [optional] 

## Methods

### NewAttribute

`func NewAttribute() *Attribute`

NewAttribute instantiates a new Attribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeWithDefaults

`func NewAttributeWithDefaults() *Attribute`

NewAttributeWithDefaults instantiates a new Attribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Attribute) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Attribute) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Attribute) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Attribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackendName

`func (o *Attribute) GetBackendName() string`

GetBackendName returns the BackendName field if non-nil, zero value otherwise.

### GetBackendNameOk

`func (o *Attribute) GetBackendNameOk() (*string, bool)`

GetBackendNameOk returns a tuple with the BackendName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendName

`func (o *Attribute) SetBackendName(v string)`

SetBackendName sets BackendName field to given value.

### HasBackendName

`func (o *Attribute) HasBackendName() bool`

HasBackendName returns a boolean if a field has been set.

### GetPosition

`func (o *Attribute) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Attribute) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Attribute) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Attribute) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetIsSurchargePercental

`func (o *Attribute) GetIsSurchargePercental() bool`

GetIsSurchargePercental returns the IsSurchargePercental field if non-nil, zero value otherwise.

### GetIsSurchargePercentalOk

`func (o *Attribute) GetIsSurchargePercentalOk() (*bool, bool)`

GetIsSurchargePercentalOk returns a tuple with the IsSurchargePercental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSurchargePercental

`func (o *Attribute) SetIsSurchargePercental(v bool)`

SetIsSurchargePercental sets IsSurchargePercental field to given value.

### HasIsSurchargePercental

`func (o *Attribute) HasIsSurchargePercental() bool`

HasIsSurchargePercental returns a boolean if a field has been set.

### GetIsLinkableToImage

`func (o *Attribute) GetIsLinkableToImage() bool`

GetIsLinkableToImage returns the IsLinkableToImage field if non-nil, zero value otherwise.

### GetIsLinkableToImageOk

`func (o *Attribute) GetIsLinkableToImageOk() (*bool, bool)`

GetIsLinkableToImageOk returns a tuple with the IsLinkableToImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLinkableToImage

`func (o *Attribute) SetIsLinkableToImage(v bool)`

SetIsLinkableToImage sets IsLinkableToImage field to given value.

### HasIsLinkableToImage

`func (o *Attribute) HasIsLinkableToImage() bool`

HasIsLinkableToImage returns a boolean if a field has been set.

### GetAmazonAttribute

`func (o *Attribute) GetAmazonAttribute() string`

GetAmazonAttribute returns the AmazonAttribute field if non-nil, zero value otherwise.

### GetAmazonAttributeOk

`func (o *Attribute) GetAmazonAttributeOk() (*string, bool)`

GetAmazonAttributeOk returns a tuple with the AmazonAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAttribute

`func (o *Attribute) SetAmazonAttribute(v string)`

SetAmazonAttribute sets AmazonAttribute field to given value.

### HasAmazonAttribute

`func (o *Attribute) HasAmazonAttribute() bool`

HasAmazonAttribute returns a boolean if a field has been set.

### GetFruugoAttribute

`func (o *Attribute) GetFruugoAttribute() string`

GetFruugoAttribute returns the FruugoAttribute field if non-nil, zero value otherwise.

### GetFruugoAttributeOk

`func (o *Attribute) GetFruugoAttributeOk() (*string, bool)`

GetFruugoAttributeOk returns a tuple with the FruugoAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFruugoAttribute

`func (o *Attribute) SetFruugoAttribute(v string)`

SetFruugoAttribute sets FruugoAttribute field to given value.

### HasFruugoAttribute

`func (o *Attribute) HasFruugoAttribute() bool`

HasFruugoAttribute returns a boolean if a field has been set.

### GetPixmaniaAttribute

`func (o *Attribute) GetPixmaniaAttribute() string`

GetPixmaniaAttribute returns the PixmaniaAttribute field if non-nil, zero value otherwise.

### GetPixmaniaAttributeOk

`func (o *Attribute) GetPixmaniaAttributeOk() (*string, bool)`

GetPixmaniaAttributeOk returns a tuple with the PixmaniaAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixmaniaAttribute

`func (o *Attribute) SetPixmaniaAttribute(v string)`

SetPixmaniaAttribute sets PixmaniaAttribute field to given value.

### HasPixmaniaAttribute

`func (o *Attribute) HasPixmaniaAttribute() bool`

HasPixmaniaAttribute returns a boolean if a field has been set.

### GetOttoAttribute

`func (o *Attribute) GetOttoAttribute() string`

GetOttoAttribute returns the OttoAttribute field if non-nil, zero value otherwise.

### GetOttoAttributeOk

`func (o *Attribute) GetOttoAttributeOk() (*string, bool)`

GetOttoAttributeOk returns a tuple with the OttoAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOttoAttribute

`func (o *Attribute) SetOttoAttribute(v string)`

SetOttoAttribute sets OttoAttribute field to given value.

### HasOttoAttribute

`func (o *Attribute) HasOttoAttribute() bool`

HasOttoAttribute returns a boolean if a field has been set.

### GetGoogleShoppingAttribute

`func (o *Attribute) GetGoogleShoppingAttribute() string`

GetGoogleShoppingAttribute returns the GoogleShoppingAttribute field if non-nil, zero value otherwise.

### GetGoogleShoppingAttributeOk

`func (o *Attribute) GetGoogleShoppingAttributeOk() (*string, bool)`

GetGoogleShoppingAttributeOk returns a tuple with the GoogleShoppingAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleShoppingAttribute

`func (o *Attribute) SetGoogleShoppingAttribute(v string)`

SetGoogleShoppingAttribute sets GoogleShoppingAttribute field to given value.

### HasGoogleShoppingAttribute

`func (o *Attribute) HasGoogleShoppingAttribute() bool`

HasGoogleShoppingAttribute returns a boolean if a field has been set.

### GetNeckermannAtEpAttribute

`func (o *Attribute) GetNeckermannAtEpAttribute() int32`

GetNeckermannAtEpAttribute returns the NeckermannAtEpAttribute field if non-nil, zero value otherwise.

### GetNeckermannAtEpAttributeOk

`func (o *Attribute) GetNeckermannAtEpAttributeOk() (*int32, bool)`

GetNeckermannAtEpAttributeOk returns a tuple with the NeckermannAtEpAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeckermannAtEpAttribute

`func (o *Attribute) SetNeckermannAtEpAttribute(v int32)`

SetNeckermannAtEpAttribute sets NeckermannAtEpAttribute field to given value.

### HasNeckermannAtEpAttribute

`func (o *Attribute) HasNeckermannAtEpAttribute() bool`

HasNeckermannAtEpAttribute returns a boolean if a field has been set.

### GetTypeOfSelectionInOnlineStore

`func (o *Attribute) GetTypeOfSelectionInOnlineStore() string`

GetTypeOfSelectionInOnlineStore returns the TypeOfSelectionInOnlineStore field if non-nil, zero value otherwise.

### GetTypeOfSelectionInOnlineStoreOk

`func (o *Attribute) GetTypeOfSelectionInOnlineStoreOk() (*string, bool)`

GetTypeOfSelectionInOnlineStoreOk returns a tuple with the TypeOfSelectionInOnlineStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeOfSelectionInOnlineStore

`func (o *Attribute) SetTypeOfSelectionInOnlineStore(v string)`

SetTypeOfSelectionInOnlineStore sets TypeOfSelectionInOnlineStore field to given value.

### HasTypeOfSelectionInOnlineStore

`func (o *Attribute) HasTypeOfSelectionInOnlineStore() bool`

HasTypeOfSelectionInOnlineStore returns a boolean if a field has been set.

### GetLaRedouteAttribute

`func (o *Attribute) GetLaRedouteAttribute() int32`

GetLaRedouteAttribute returns the LaRedouteAttribute field if non-nil, zero value otherwise.

### GetLaRedouteAttributeOk

`func (o *Attribute) GetLaRedouteAttributeOk() (*int32, bool)`

GetLaRedouteAttributeOk returns a tuple with the LaRedouteAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaRedouteAttribute

`func (o *Attribute) SetLaRedouteAttribute(v int32)`

SetLaRedouteAttribute sets LaRedouteAttribute field to given value.

### HasLaRedouteAttribute

`func (o *Attribute) HasLaRedouteAttribute() bool`

HasLaRedouteAttribute returns a boolean if a field has been set.

### GetIsGroupable

`func (o *Attribute) GetIsGroupable() bool`

GetIsGroupable returns the IsGroupable field if non-nil, zero value otherwise.

### GetIsGroupableOk

`func (o *Attribute) GetIsGroupableOk() (*bool, bool)`

GetIsGroupableOk returns a tuple with the IsGroupable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroupable

`func (o *Attribute) SetIsGroupable(v bool)`

SetIsGroupable sets IsGroupable field to given value.

### HasIsGroupable

`func (o *Attribute) HasIsGroupable() bool`

HasIsGroupable returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Attribute) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Attribute) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Attribute) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Attribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValues

`func (o *Attribute) GetValues() []AttributeValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Attribute) GetValuesOk() (*[]AttributeValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Attribute) SetValues(v []AttributeValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *Attribute) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


