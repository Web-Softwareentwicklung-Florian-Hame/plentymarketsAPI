# CreateItemsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProducingCountryId** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**StockType** | Pointer to **int32** | values: 0 &#x3D; Stocked item (default)1 &#x3D; Production item2 &#x3D; Colli 3 &#x3D; Special order item | [optional] 
**StoreSpecial** | Pointer to **int32** | Option to present items more prominently in the online store. 1 &#x3D; Special offer 2 &#x3D; New items 3 &#x3D; Top items | [optional] 
**OwnerId** | Pointer to **int32** | The plentymarkets user that is assigned as owner of this item. | [optional] 
**ManufacturerId** | Pointer to **int32** | The ID of the manufacturer of the item | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CustomsTariffNumber** | Pointer to **string** | Deprecated: The customs tariff number of the item; usually a 11 digit code number based on the Harmonised System | [optional] 
**RevenueAccount** | Pointer to **int32** | The revenue account the item is linked to. An individual revenue account can be saved for each item in plentymarkets. If this is not done, plentymarkets automatically determines a revenue account based on the VAT rate. | [optional] 
**CouponRestriction** | Pointer to **int32** | Indicates if the item can be purchased using a promotional coupon. 0 &#x3D; Permitted 1 &#x3D; Not permitted 2 &#x3D; Purchasable with coupon only | [optional] 
**Condition** | Pointer to **int32** | The condition of the item. Possible values: 0 &#x3D; New 1 &#x3D; Used 2 &#x3D; Boxed as new 3 &#x3D; New with label 4 &#x3D; Factory seconds | [optional] 
**ConditionApi** | Pointer to **int32** | The condition of the item that is transferred to markets via API. 0 &#x3D; New 1 &#x3D; Used but as new 2 &#x3D; Used but very good 3 &#x3D; Used but good 4 &#x3D; Used but acceptable 5 &#x3D; Factory seconds | [optional] 
**IsSubscribable** | Pointer to **bool** | Flag that indicates if the item can be ordered as a subscription item. If yes, the item can be ordered for delivery at regular intervals. | [optional] 
**IsShippingPackage** | Pointer to **bool** | Flag that indicates if a shipping package is to be used for this item. If yes and the variation&#39;s dimensions are entered in the Settings tab of a variation, the correct shipping package is assigned automatically. | [optional] 
**AmazonFbaPlatform** | Pointer to **int32** | Indicates the platform used for Fulfilment by Amazon (FBA). 0 &#x3D; Do not use 1 &#x3D; AMAZON EU (Europe) 2 &#x3D; AMAZON FE (Far East) 3 &#x3D; AMAZON NA (North America) | [optional] 
**IsShippableByAmazon** | Pointer to **bool** |  | [optional] 
**AmazonProductType** | Pointer to **int32** | The Amazon product type of the item. List of IDs: https://www.plentymarkets.co.uk/manual/data-exchange/data-formats/item/ | [optional] 
**AmazonFedas** | Pointer to **string** | The FEDAS product classification key of the item. | [optional] 
**EbayPresetId** | Pointer to **int32** | The eBay preset ID. This plentymarkets ID must be specified to save values for $ebayCategory, $ebayCategory2, $ebayStoreCategory and $ebayStoreCategory2 | [optional] 
**EbayCategory** | Pointer to **int32** | The eBay category 1 of the item. This category is used when a new listing is created. | [optional] 
**EbayCategory2** | Pointer to **int32** | The eBay category 2 of the item. This category is used when a new listing is created. | [optional] 
**RakutenCategoryId** | Pointer to **int32** | The ID of the Rakuten category of this item. | [optional] 
**FlagOne** | Pointer to **int32** | Flag 1 of the item. Flags can be used to organise and filter items. Each item can be assigned up to two flags. 0 &#x3D; no flag | [optional] 
**FlagTwo** | Pointer to **int32** | Flag 2 of the item. Flags can be used to organise and filter items. Each item can be assigned up to two flags. 0 &#x3D; no flag | [optional] 
**AgeRestriction** | **int32** | The age customers must be to purchase the item. Items with an age rating of 18+ must be linked to a shipping profile for which the PostIdent option is activated. 0 &#x3D; None available 3 &#x3D; Released for ages 3 and up 6 &#x3D; Ages 6 and up 9 &#x3D; Ages 9 and up 12 &#x3D; Ages 12 and up 14 &#x3D; Ages 14 and up 16 &#x3D; Ages 16 and up 18 &#x3D; Ages 18 and up 50 &#x3D; Not marked 88 &#x3D; Not 99 &#x3D; Unknown | 
**Feedback** | Pointer to **int32** | The feedback, i.e. rating, that this item received. Possible values are 1 to 5 or 1 to 10 depending on the maximum rating setting. An initial feedback can be saved for items. The saved value will then be displayed as the initial feedback. Every time new feedback is submitted, the average value will be recalculated automatically. | [optional] 
**ItemType** | Pointer to **string** | The type of the item. Because Set items are managed using a separate route, this value is always Default. optional allowed values are default, multiPack, set | [optional] 
**EbayTitles** | Pointer to [**[]ItemEbayTitle**](ItemEbayTitle.md) |  | [optional] 
**ItemShippingProfiles** | Pointer to [**[]ItemShippingProfile**](ItemShippingProfile.md) |  | [optional] 
**ProfileId** | **int32** | The unique ID of the shipping profile | 
**ItemProperties** | Pointer to [**[]ItemProperty**](ItemProperty.md) |  | [optional] 
**PropertyId** | **int32** | The id of the property item | 
**PropertySelectionId** | Pointer to **int32** | The id of the property selection | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Texts** | Pointer to [**[]ItemTexts**](ItemTexts.md) |  | [optional] 
**Variations** | [**[]Variation**](Variation.md) |  | 
**Success** | Pointer to [**map[string]Item**](Item.md) |  | [optional] 
**Failed** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewCreateItemsResponse

`func NewCreateItemsResponse(ageRestriction int32, profileId int32, propertyId int32, variations []Variation, ) *CreateItemsResponse`

NewCreateItemsResponse instantiates a new CreateItemsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateItemsResponseWithDefaults

`func NewCreateItemsResponseWithDefaults() *CreateItemsResponse`

NewCreateItemsResponseWithDefaults instantiates a new CreateItemsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducingCountryId

`func (o *CreateItemsResponse) GetProducingCountryId() int32`

GetProducingCountryId returns the ProducingCountryId field if non-nil, zero value otherwise.

### GetProducingCountryIdOk

`func (o *CreateItemsResponse) GetProducingCountryIdOk() (*int32, bool)`

GetProducingCountryIdOk returns a tuple with the ProducingCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducingCountryId

`func (o *CreateItemsResponse) SetProducingCountryId(v int32)`

SetProducingCountryId sets ProducingCountryId field to given value.

### HasProducingCountryId

`func (o *CreateItemsResponse) HasProducingCountryId() bool`

HasProducingCountryId returns a boolean if a field has been set.

### GetPosition

`func (o *CreateItemsResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CreateItemsResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CreateItemsResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CreateItemsResponse) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStockType

`func (o *CreateItemsResponse) GetStockType() int32`

GetStockType returns the StockType field if non-nil, zero value otherwise.

### GetStockTypeOk

`func (o *CreateItemsResponse) GetStockTypeOk() (*int32, bool)`

GetStockTypeOk returns a tuple with the StockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockType

`func (o *CreateItemsResponse) SetStockType(v int32)`

SetStockType sets StockType field to given value.

### HasStockType

`func (o *CreateItemsResponse) HasStockType() bool`

HasStockType returns a boolean if a field has been set.

### GetStoreSpecial

`func (o *CreateItemsResponse) GetStoreSpecial() int32`

GetStoreSpecial returns the StoreSpecial field if non-nil, zero value otherwise.

### GetStoreSpecialOk

`func (o *CreateItemsResponse) GetStoreSpecialOk() (*int32, bool)`

GetStoreSpecialOk returns a tuple with the StoreSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSpecial

`func (o *CreateItemsResponse) SetStoreSpecial(v int32)`

SetStoreSpecial sets StoreSpecial field to given value.

### HasStoreSpecial

`func (o *CreateItemsResponse) HasStoreSpecial() bool`

HasStoreSpecial returns a boolean if a field has been set.

### GetOwnerId

`func (o *CreateItemsResponse) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *CreateItemsResponse) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *CreateItemsResponse) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *CreateItemsResponse) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetManufacturerId

`func (o *CreateItemsResponse) GetManufacturerId() int32`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *CreateItemsResponse) GetManufacturerIdOk() (*int32, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *CreateItemsResponse) SetManufacturerId(v int32)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *CreateItemsResponse) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateItemsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateItemsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateItemsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateItemsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateItemsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateItemsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateItemsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateItemsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCustomsTariffNumber

`func (o *CreateItemsResponse) GetCustomsTariffNumber() string`

GetCustomsTariffNumber returns the CustomsTariffNumber field if non-nil, zero value otherwise.

### GetCustomsTariffNumberOk

`func (o *CreateItemsResponse) GetCustomsTariffNumberOk() (*string, bool)`

GetCustomsTariffNumberOk returns a tuple with the CustomsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsTariffNumber

`func (o *CreateItemsResponse) SetCustomsTariffNumber(v string)`

SetCustomsTariffNumber sets CustomsTariffNumber field to given value.

### HasCustomsTariffNumber

`func (o *CreateItemsResponse) HasCustomsTariffNumber() bool`

HasCustomsTariffNumber returns a boolean if a field has been set.

### GetRevenueAccount

`func (o *CreateItemsResponse) GetRevenueAccount() int32`

GetRevenueAccount returns the RevenueAccount field if non-nil, zero value otherwise.

### GetRevenueAccountOk

`func (o *CreateItemsResponse) GetRevenueAccountOk() (*int32, bool)`

GetRevenueAccountOk returns a tuple with the RevenueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAccount

`func (o *CreateItemsResponse) SetRevenueAccount(v int32)`

SetRevenueAccount sets RevenueAccount field to given value.

### HasRevenueAccount

`func (o *CreateItemsResponse) HasRevenueAccount() bool`

HasRevenueAccount returns a boolean if a field has been set.

### GetCouponRestriction

`func (o *CreateItemsResponse) GetCouponRestriction() int32`

GetCouponRestriction returns the CouponRestriction field if non-nil, zero value otherwise.

### GetCouponRestrictionOk

`func (o *CreateItemsResponse) GetCouponRestrictionOk() (*int32, bool)`

GetCouponRestrictionOk returns a tuple with the CouponRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRestriction

`func (o *CreateItemsResponse) SetCouponRestriction(v int32)`

SetCouponRestriction sets CouponRestriction field to given value.

### HasCouponRestriction

`func (o *CreateItemsResponse) HasCouponRestriction() bool`

HasCouponRestriction returns a boolean if a field has been set.

### GetCondition

`func (o *CreateItemsResponse) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateItemsResponse) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateItemsResponse) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateItemsResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionApi

`func (o *CreateItemsResponse) GetConditionApi() int32`

GetConditionApi returns the ConditionApi field if non-nil, zero value otherwise.

### GetConditionApiOk

`func (o *CreateItemsResponse) GetConditionApiOk() (*int32, bool)`

GetConditionApiOk returns a tuple with the ConditionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionApi

`func (o *CreateItemsResponse) SetConditionApi(v int32)`

SetConditionApi sets ConditionApi field to given value.

### HasConditionApi

`func (o *CreateItemsResponse) HasConditionApi() bool`

HasConditionApi returns a boolean if a field has been set.

### GetIsSubscribable

`func (o *CreateItemsResponse) GetIsSubscribable() bool`

GetIsSubscribable returns the IsSubscribable field if non-nil, zero value otherwise.

### GetIsSubscribableOk

`func (o *CreateItemsResponse) GetIsSubscribableOk() (*bool, bool)`

GetIsSubscribableOk returns a tuple with the IsSubscribable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribable

`func (o *CreateItemsResponse) SetIsSubscribable(v bool)`

SetIsSubscribable sets IsSubscribable field to given value.

### HasIsSubscribable

`func (o *CreateItemsResponse) HasIsSubscribable() bool`

HasIsSubscribable returns a boolean if a field has been set.

### GetIsShippingPackage

`func (o *CreateItemsResponse) GetIsShippingPackage() bool`

GetIsShippingPackage returns the IsShippingPackage field if non-nil, zero value otherwise.

### GetIsShippingPackageOk

`func (o *CreateItemsResponse) GetIsShippingPackageOk() (*bool, bool)`

GetIsShippingPackageOk returns a tuple with the IsShippingPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingPackage

`func (o *CreateItemsResponse) SetIsShippingPackage(v bool)`

SetIsShippingPackage sets IsShippingPackage field to given value.

### HasIsShippingPackage

`func (o *CreateItemsResponse) HasIsShippingPackage() bool`

HasIsShippingPackage returns a boolean if a field has been set.

### GetAmazonFbaPlatform

`func (o *CreateItemsResponse) GetAmazonFbaPlatform() int32`

GetAmazonFbaPlatform returns the AmazonFbaPlatform field if non-nil, zero value otherwise.

### GetAmazonFbaPlatformOk

`func (o *CreateItemsResponse) GetAmazonFbaPlatformOk() (*int32, bool)`

GetAmazonFbaPlatformOk returns a tuple with the AmazonFbaPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFbaPlatform

`func (o *CreateItemsResponse) SetAmazonFbaPlatform(v int32)`

SetAmazonFbaPlatform sets AmazonFbaPlatform field to given value.

### HasAmazonFbaPlatform

`func (o *CreateItemsResponse) HasAmazonFbaPlatform() bool`

HasAmazonFbaPlatform returns a boolean if a field has been set.

### GetIsShippableByAmazon

`func (o *CreateItemsResponse) GetIsShippableByAmazon() bool`

GetIsShippableByAmazon returns the IsShippableByAmazon field if non-nil, zero value otherwise.

### GetIsShippableByAmazonOk

`func (o *CreateItemsResponse) GetIsShippableByAmazonOk() (*bool, bool)`

GetIsShippableByAmazonOk returns a tuple with the IsShippableByAmazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippableByAmazon

`func (o *CreateItemsResponse) SetIsShippableByAmazon(v bool)`

SetIsShippableByAmazon sets IsShippableByAmazon field to given value.

### HasIsShippableByAmazon

`func (o *CreateItemsResponse) HasIsShippableByAmazon() bool`

HasIsShippableByAmazon returns a boolean if a field has been set.

### GetAmazonProductType

`func (o *CreateItemsResponse) GetAmazonProductType() int32`

GetAmazonProductType returns the AmazonProductType field if non-nil, zero value otherwise.

### GetAmazonProductTypeOk

`func (o *CreateItemsResponse) GetAmazonProductTypeOk() (*int32, bool)`

GetAmazonProductTypeOk returns a tuple with the AmazonProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductType

`func (o *CreateItemsResponse) SetAmazonProductType(v int32)`

SetAmazonProductType sets AmazonProductType field to given value.

### HasAmazonProductType

`func (o *CreateItemsResponse) HasAmazonProductType() bool`

HasAmazonProductType returns a boolean if a field has been set.

### GetAmazonFedas

`func (o *CreateItemsResponse) GetAmazonFedas() string`

GetAmazonFedas returns the AmazonFedas field if non-nil, zero value otherwise.

### GetAmazonFedasOk

`func (o *CreateItemsResponse) GetAmazonFedasOk() (*string, bool)`

GetAmazonFedasOk returns a tuple with the AmazonFedas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFedas

`func (o *CreateItemsResponse) SetAmazonFedas(v string)`

SetAmazonFedas sets AmazonFedas field to given value.

### HasAmazonFedas

`func (o *CreateItemsResponse) HasAmazonFedas() bool`

HasAmazonFedas returns a boolean if a field has been set.

### GetEbayPresetId

`func (o *CreateItemsResponse) GetEbayPresetId() int32`

GetEbayPresetId returns the EbayPresetId field if non-nil, zero value otherwise.

### GetEbayPresetIdOk

`func (o *CreateItemsResponse) GetEbayPresetIdOk() (*int32, bool)`

GetEbayPresetIdOk returns a tuple with the EbayPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayPresetId

`func (o *CreateItemsResponse) SetEbayPresetId(v int32)`

SetEbayPresetId sets EbayPresetId field to given value.

### HasEbayPresetId

`func (o *CreateItemsResponse) HasEbayPresetId() bool`

HasEbayPresetId returns a boolean if a field has been set.

### GetEbayCategory

`func (o *CreateItemsResponse) GetEbayCategory() int32`

GetEbayCategory returns the EbayCategory field if non-nil, zero value otherwise.

### GetEbayCategoryOk

`func (o *CreateItemsResponse) GetEbayCategoryOk() (*int32, bool)`

GetEbayCategoryOk returns a tuple with the EbayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory

`func (o *CreateItemsResponse) SetEbayCategory(v int32)`

SetEbayCategory sets EbayCategory field to given value.

### HasEbayCategory

`func (o *CreateItemsResponse) HasEbayCategory() bool`

HasEbayCategory returns a boolean if a field has been set.

### GetEbayCategory2

`func (o *CreateItemsResponse) GetEbayCategory2() int32`

GetEbayCategory2 returns the EbayCategory2 field if non-nil, zero value otherwise.

### GetEbayCategory2Ok

`func (o *CreateItemsResponse) GetEbayCategory2Ok() (*int32, bool)`

GetEbayCategory2Ok returns a tuple with the EbayCategory2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory2

`func (o *CreateItemsResponse) SetEbayCategory2(v int32)`

SetEbayCategory2 sets EbayCategory2 field to given value.

### HasEbayCategory2

`func (o *CreateItemsResponse) HasEbayCategory2() bool`

HasEbayCategory2 returns a boolean if a field has been set.

### GetRakutenCategoryId

`func (o *CreateItemsResponse) GetRakutenCategoryId() int32`

GetRakutenCategoryId returns the RakutenCategoryId field if non-nil, zero value otherwise.

### GetRakutenCategoryIdOk

`func (o *CreateItemsResponse) GetRakutenCategoryIdOk() (*int32, bool)`

GetRakutenCategoryIdOk returns a tuple with the RakutenCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRakutenCategoryId

`func (o *CreateItemsResponse) SetRakutenCategoryId(v int32)`

SetRakutenCategoryId sets RakutenCategoryId field to given value.

### HasRakutenCategoryId

`func (o *CreateItemsResponse) HasRakutenCategoryId() bool`

HasRakutenCategoryId returns a boolean if a field has been set.

### GetFlagOne

`func (o *CreateItemsResponse) GetFlagOne() int32`

GetFlagOne returns the FlagOne field if non-nil, zero value otherwise.

### GetFlagOneOk

`func (o *CreateItemsResponse) GetFlagOneOk() (*int32, bool)`

GetFlagOneOk returns a tuple with the FlagOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagOne

`func (o *CreateItemsResponse) SetFlagOne(v int32)`

SetFlagOne sets FlagOne field to given value.

### HasFlagOne

`func (o *CreateItemsResponse) HasFlagOne() bool`

HasFlagOne returns a boolean if a field has been set.

### GetFlagTwo

`func (o *CreateItemsResponse) GetFlagTwo() int32`

GetFlagTwo returns the FlagTwo field if non-nil, zero value otherwise.

### GetFlagTwoOk

`func (o *CreateItemsResponse) GetFlagTwoOk() (*int32, bool)`

GetFlagTwoOk returns a tuple with the FlagTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagTwo

`func (o *CreateItemsResponse) SetFlagTwo(v int32)`

SetFlagTwo sets FlagTwo field to given value.

### HasFlagTwo

`func (o *CreateItemsResponse) HasFlagTwo() bool`

HasFlagTwo returns a boolean if a field has been set.

### GetAgeRestriction

`func (o *CreateItemsResponse) GetAgeRestriction() int32`

GetAgeRestriction returns the AgeRestriction field if non-nil, zero value otherwise.

### GetAgeRestrictionOk

`func (o *CreateItemsResponse) GetAgeRestrictionOk() (*int32, bool)`

GetAgeRestrictionOk returns a tuple with the AgeRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRestriction

`func (o *CreateItemsResponse) SetAgeRestriction(v int32)`

SetAgeRestriction sets AgeRestriction field to given value.


### GetFeedback

`func (o *CreateItemsResponse) GetFeedback() int32`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *CreateItemsResponse) GetFeedbackOk() (*int32, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *CreateItemsResponse) SetFeedback(v int32)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *CreateItemsResponse) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetItemType

`func (o *CreateItemsResponse) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateItemsResponse) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateItemsResponse) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CreateItemsResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetEbayTitles

`func (o *CreateItemsResponse) GetEbayTitles() []ItemEbayTitle`

GetEbayTitles returns the EbayTitles field if non-nil, zero value otherwise.

### GetEbayTitlesOk

`func (o *CreateItemsResponse) GetEbayTitlesOk() (*[]ItemEbayTitle, bool)`

GetEbayTitlesOk returns a tuple with the EbayTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayTitles

`func (o *CreateItemsResponse) SetEbayTitles(v []ItemEbayTitle)`

SetEbayTitles sets EbayTitles field to given value.

### HasEbayTitles

`func (o *CreateItemsResponse) HasEbayTitles() bool`

HasEbayTitles returns a boolean if a field has been set.

### GetItemShippingProfiles

`func (o *CreateItemsResponse) GetItemShippingProfiles() []ItemShippingProfile`

GetItemShippingProfiles returns the ItemShippingProfiles field if non-nil, zero value otherwise.

### GetItemShippingProfilesOk

`func (o *CreateItemsResponse) GetItemShippingProfilesOk() (*[]ItemShippingProfile, bool)`

GetItemShippingProfilesOk returns a tuple with the ItemShippingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShippingProfiles

`func (o *CreateItemsResponse) SetItemShippingProfiles(v []ItemShippingProfile)`

SetItemShippingProfiles sets ItemShippingProfiles field to given value.

### HasItemShippingProfiles

`func (o *CreateItemsResponse) HasItemShippingProfiles() bool`

HasItemShippingProfiles returns a boolean if a field has been set.

### GetProfileId

`func (o *CreateItemsResponse) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateItemsResponse) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateItemsResponse) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.


### GetItemProperties

`func (o *CreateItemsResponse) GetItemProperties() []ItemProperty`

GetItemProperties returns the ItemProperties field if non-nil, zero value otherwise.

### GetItemPropertiesOk

`func (o *CreateItemsResponse) GetItemPropertiesOk() (*[]ItemProperty, bool)`

GetItemPropertiesOk returns a tuple with the ItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemProperties

`func (o *CreateItemsResponse) SetItemProperties(v []ItemProperty)`

SetItemProperties sets ItemProperties field to given value.

### HasItemProperties

`func (o *CreateItemsResponse) HasItemProperties() bool`

HasItemProperties returns a boolean if a field has been set.

### GetPropertyId

`func (o *CreateItemsResponse) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *CreateItemsResponse) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *CreateItemsResponse) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetPropertySelectionId

`func (o *CreateItemsResponse) GetPropertySelectionId() int32`

GetPropertySelectionId returns the PropertySelectionId field if non-nil, zero value otherwise.

### GetPropertySelectionIdOk

`func (o *CreateItemsResponse) GetPropertySelectionIdOk() (*int32, bool)`

GetPropertySelectionIdOk returns a tuple with the PropertySelectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelectionId

`func (o *CreateItemsResponse) SetPropertySelectionId(v int32)`

SetPropertySelectionId sets PropertySelectionId field to given value.

### HasPropertySelectionId

`func (o *CreateItemsResponse) HasPropertySelectionId() bool`

HasPropertySelectionId returns a boolean if a field has been set.

### GetId

`func (o *CreateItemsResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateItemsResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateItemsResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateItemsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTexts

`func (o *CreateItemsResponse) GetTexts() []ItemTexts`

GetTexts returns the Texts field if non-nil, zero value otherwise.

### GetTextsOk

`func (o *CreateItemsResponse) GetTextsOk() (*[]ItemTexts, bool)`

GetTextsOk returns a tuple with the Texts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTexts

`func (o *CreateItemsResponse) SetTexts(v []ItemTexts)`

SetTexts sets Texts field to given value.

### HasTexts

`func (o *CreateItemsResponse) HasTexts() bool`

HasTexts returns a boolean if a field has been set.

### GetVariations

`func (o *CreateItemsResponse) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *CreateItemsResponse) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *CreateItemsResponse) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.


### GetSuccess

`func (o *CreateItemsResponse) GetSuccess() map[string]Item`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateItemsResponse) GetSuccessOk() (*map[string]Item, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateItemsResponse) SetSuccess(v map[string]Item)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateItemsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFailed

`func (o *CreateItemsResponse) GetFailed() map[string][]string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CreateItemsResponse) GetFailedOk() (*map[string][]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CreateItemsResponse) SetFailed(v map[string][]string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CreateItemsResponse) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


