# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**Variations** | [**[]Variation**](Variation.md) |  | 

## Methods

### NewItem

`func NewItem(ageRestriction int32, profileId int32, propertyId int32, variations []Variation, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *Item) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Item) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Item) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Item) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetStockType

`func (o *Item) GetStockType() int32`

GetStockType returns the StockType field if non-nil, zero value otherwise.

### GetStockTypeOk

`func (o *Item) GetStockTypeOk() (*int32, bool)`

GetStockTypeOk returns a tuple with the StockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockType

`func (o *Item) SetStockType(v int32)`

SetStockType sets StockType field to given value.

### HasStockType

`func (o *Item) HasStockType() bool`

HasStockType returns a boolean if a field has been set.

### GetStoreSpecial

`func (o *Item) GetStoreSpecial() int32`

GetStoreSpecial returns the StoreSpecial field if non-nil, zero value otherwise.

### GetStoreSpecialOk

`func (o *Item) GetStoreSpecialOk() (*int32, bool)`

GetStoreSpecialOk returns a tuple with the StoreSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSpecial

`func (o *Item) SetStoreSpecial(v int32)`

SetStoreSpecial sets StoreSpecial field to given value.

### HasStoreSpecial

`func (o *Item) HasStoreSpecial() bool`

HasStoreSpecial returns a boolean if a field has been set.

### GetOwnerId

`func (o *Item) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Item) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Item) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Item) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetManufacturerId

`func (o *Item) GetManufacturerId() int32`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *Item) GetManufacturerIdOk() (*int32, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *Item) SetManufacturerId(v int32)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *Item) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Item) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Item) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Item) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Item) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Item) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Item) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Item) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Item) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCustomsTariffNumber

`func (o *Item) GetCustomsTariffNumber() string`

GetCustomsTariffNumber returns the CustomsTariffNumber field if non-nil, zero value otherwise.

### GetCustomsTariffNumberOk

`func (o *Item) GetCustomsTariffNumberOk() (*string, bool)`

GetCustomsTariffNumberOk returns a tuple with the CustomsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsTariffNumber

`func (o *Item) SetCustomsTariffNumber(v string)`

SetCustomsTariffNumber sets CustomsTariffNumber field to given value.

### HasCustomsTariffNumber

`func (o *Item) HasCustomsTariffNumber() bool`

HasCustomsTariffNumber returns a boolean if a field has been set.

### GetRevenueAccount

`func (o *Item) GetRevenueAccount() int32`

GetRevenueAccount returns the RevenueAccount field if non-nil, zero value otherwise.

### GetRevenueAccountOk

`func (o *Item) GetRevenueAccountOk() (*int32, bool)`

GetRevenueAccountOk returns a tuple with the RevenueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAccount

`func (o *Item) SetRevenueAccount(v int32)`

SetRevenueAccount sets RevenueAccount field to given value.

### HasRevenueAccount

`func (o *Item) HasRevenueAccount() bool`

HasRevenueAccount returns a boolean if a field has been set.

### GetCouponRestriction

`func (o *Item) GetCouponRestriction() int32`

GetCouponRestriction returns the CouponRestriction field if non-nil, zero value otherwise.

### GetCouponRestrictionOk

`func (o *Item) GetCouponRestrictionOk() (*int32, bool)`

GetCouponRestrictionOk returns a tuple with the CouponRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRestriction

`func (o *Item) SetCouponRestriction(v int32)`

SetCouponRestriction sets CouponRestriction field to given value.

### HasCouponRestriction

`func (o *Item) HasCouponRestriction() bool`

HasCouponRestriction returns a boolean if a field has been set.

### GetCondition

`func (o *Item) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *Item) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *Item) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *Item) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetConditionApi

`func (o *Item) GetConditionApi() int32`

GetConditionApi returns the ConditionApi field if non-nil, zero value otherwise.

### GetConditionApiOk

`func (o *Item) GetConditionApiOk() (*int32, bool)`

GetConditionApiOk returns a tuple with the ConditionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionApi

`func (o *Item) SetConditionApi(v int32)`

SetConditionApi sets ConditionApi field to given value.

### HasConditionApi

`func (o *Item) HasConditionApi() bool`

HasConditionApi returns a boolean if a field has been set.

### GetIsSubscribable

`func (o *Item) GetIsSubscribable() bool`

GetIsSubscribable returns the IsSubscribable field if non-nil, zero value otherwise.

### GetIsSubscribableOk

`func (o *Item) GetIsSubscribableOk() (*bool, bool)`

GetIsSubscribableOk returns a tuple with the IsSubscribable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribable

`func (o *Item) SetIsSubscribable(v bool)`

SetIsSubscribable sets IsSubscribable field to given value.

### HasIsSubscribable

`func (o *Item) HasIsSubscribable() bool`

HasIsSubscribable returns a boolean if a field has been set.

### GetIsShippingPackage

`func (o *Item) GetIsShippingPackage() bool`

GetIsShippingPackage returns the IsShippingPackage field if non-nil, zero value otherwise.

### GetIsShippingPackageOk

`func (o *Item) GetIsShippingPackageOk() (*bool, bool)`

GetIsShippingPackageOk returns a tuple with the IsShippingPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingPackage

`func (o *Item) SetIsShippingPackage(v bool)`

SetIsShippingPackage sets IsShippingPackage field to given value.

### HasIsShippingPackage

`func (o *Item) HasIsShippingPackage() bool`

HasIsShippingPackage returns a boolean if a field has been set.

### GetAmazonFbaPlatform

`func (o *Item) GetAmazonFbaPlatform() int32`

GetAmazonFbaPlatform returns the AmazonFbaPlatform field if non-nil, zero value otherwise.

### GetAmazonFbaPlatformOk

`func (o *Item) GetAmazonFbaPlatformOk() (*int32, bool)`

GetAmazonFbaPlatformOk returns a tuple with the AmazonFbaPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFbaPlatform

`func (o *Item) SetAmazonFbaPlatform(v int32)`

SetAmazonFbaPlatform sets AmazonFbaPlatform field to given value.

### HasAmazonFbaPlatform

`func (o *Item) HasAmazonFbaPlatform() bool`

HasAmazonFbaPlatform returns a boolean if a field has been set.

### GetIsShippableByAmazon

`func (o *Item) GetIsShippableByAmazon() bool`

GetIsShippableByAmazon returns the IsShippableByAmazon field if non-nil, zero value otherwise.

### GetIsShippableByAmazonOk

`func (o *Item) GetIsShippableByAmazonOk() (*bool, bool)`

GetIsShippableByAmazonOk returns a tuple with the IsShippableByAmazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippableByAmazon

`func (o *Item) SetIsShippableByAmazon(v bool)`

SetIsShippableByAmazon sets IsShippableByAmazon field to given value.

### HasIsShippableByAmazon

`func (o *Item) HasIsShippableByAmazon() bool`

HasIsShippableByAmazon returns a boolean if a field has been set.

### GetAmazonProductType

`func (o *Item) GetAmazonProductType() int32`

GetAmazonProductType returns the AmazonProductType field if non-nil, zero value otherwise.

### GetAmazonProductTypeOk

`func (o *Item) GetAmazonProductTypeOk() (*int32, bool)`

GetAmazonProductTypeOk returns a tuple with the AmazonProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductType

`func (o *Item) SetAmazonProductType(v int32)`

SetAmazonProductType sets AmazonProductType field to given value.

### HasAmazonProductType

`func (o *Item) HasAmazonProductType() bool`

HasAmazonProductType returns a boolean if a field has been set.

### GetAmazonFedas

`func (o *Item) GetAmazonFedas() string`

GetAmazonFedas returns the AmazonFedas field if non-nil, zero value otherwise.

### GetAmazonFedasOk

`func (o *Item) GetAmazonFedasOk() (*string, bool)`

GetAmazonFedasOk returns a tuple with the AmazonFedas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFedas

`func (o *Item) SetAmazonFedas(v string)`

SetAmazonFedas sets AmazonFedas field to given value.

### HasAmazonFedas

`func (o *Item) HasAmazonFedas() bool`

HasAmazonFedas returns a boolean if a field has been set.

### GetEbayPresetId

`func (o *Item) GetEbayPresetId() int32`

GetEbayPresetId returns the EbayPresetId field if non-nil, zero value otherwise.

### GetEbayPresetIdOk

`func (o *Item) GetEbayPresetIdOk() (*int32, bool)`

GetEbayPresetIdOk returns a tuple with the EbayPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayPresetId

`func (o *Item) SetEbayPresetId(v int32)`

SetEbayPresetId sets EbayPresetId field to given value.

### HasEbayPresetId

`func (o *Item) HasEbayPresetId() bool`

HasEbayPresetId returns a boolean if a field has been set.

### GetEbayCategory

`func (o *Item) GetEbayCategory() int32`

GetEbayCategory returns the EbayCategory field if non-nil, zero value otherwise.

### GetEbayCategoryOk

`func (o *Item) GetEbayCategoryOk() (*int32, bool)`

GetEbayCategoryOk returns a tuple with the EbayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory

`func (o *Item) SetEbayCategory(v int32)`

SetEbayCategory sets EbayCategory field to given value.

### HasEbayCategory

`func (o *Item) HasEbayCategory() bool`

HasEbayCategory returns a boolean if a field has been set.

### GetEbayCategory2

`func (o *Item) GetEbayCategory2() int32`

GetEbayCategory2 returns the EbayCategory2 field if non-nil, zero value otherwise.

### GetEbayCategory2Ok

`func (o *Item) GetEbayCategory2Ok() (*int32, bool)`

GetEbayCategory2Ok returns a tuple with the EbayCategory2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory2

`func (o *Item) SetEbayCategory2(v int32)`

SetEbayCategory2 sets EbayCategory2 field to given value.

### HasEbayCategory2

`func (o *Item) HasEbayCategory2() bool`

HasEbayCategory2 returns a boolean if a field has been set.

### GetRakutenCategoryId

`func (o *Item) GetRakutenCategoryId() int32`

GetRakutenCategoryId returns the RakutenCategoryId field if non-nil, zero value otherwise.

### GetRakutenCategoryIdOk

`func (o *Item) GetRakutenCategoryIdOk() (*int32, bool)`

GetRakutenCategoryIdOk returns a tuple with the RakutenCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRakutenCategoryId

`func (o *Item) SetRakutenCategoryId(v int32)`

SetRakutenCategoryId sets RakutenCategoryId field to given value.

### HasRakutenCategoryId

`func (o *Item) HasRakutenCategoryId() bool`

HasRakutenCategoryId returns a boolean if a field has been set.

### GetFlagOne

`func (o *Item) GetFlagOne() int32`

GetFlagOne returns the FlagOne field if non-nil, zero value otherwise.

### GetFlagOneOk

`func (o *Item) GetFlagOneOk() (*int32, bool)`

GetFlagOneOk returns a tuple with the FlagOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagOne

`func (o *Item) SetFlagOne(v int32)`

SetFlagOne sets FlagOne field to given value.

### HasFlagOne

`func (o *Item) HasFlagOne() bool`

HasFlagOne returns a boolean if a field has been set.

### GetFlagTwo

`func (o *Item) GetFlagTwo() int32`

GetFlagTwo returns the FlagTwo field if non-nil, zero value otherwise.

### GetFlagTwoOk

`func (o *Item) GetFlagTwoOk() (*int32, bool)`

GetFlagTwoOk returns a tuple with the FlagTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagTwo

`func (o *Item) SetFlagTwo(v int32)`

SetFlagTwo sets FlagTwo field to given value.

### HasFlagTwo

`func (o *Item) HasFlagTwo() bool`

HasFlagTwo returns a boolean if a field has been set.

### GetAgeRestriction

`func (o *Item) GetAgeRestriction() int32`

GetAgeRestriction returns the AgeRestriction field if non-nil, zero value otherwise.

### GetAgeRestrictionOk

`func (o *Item) GetAgeRestrictionOk() (*int32, bool)`

GetAgeRestrictionOk returns a tuple with the AgeRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRestriction

`func (o *Item) SetAgeRestriction(v int32)`

SetAgeRestriction sets AgeRestriction field to given value.


### GetFeedback

`func (o *Item) GetFeedback() int32`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *Item) GetFeedbackOk() (*int32, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *Item) SetFeedback(v int32)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *Item) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetItemType

`func (o *Item) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *Item) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *Item) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *Item) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetEbayTitles

`func (o *Item) GetEbayTitles() []ItemEbayTitle`

GetEbayTitles returns the EbayTitles field if non-nil, zero value otherwise.

### GetEbayTitlesOk

`func (o *Item) GetEbayTitlesOk() (*[]ItemEbayTitle, bool)`

GetEbayTitlesOk returns a tuple with the EbayTitles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayTitles

`func (o *Item) SetEbayTitles(v []ItemEbayTitle)`

SetEbayTitles sets EbayTitles field to given value.

### HasEbayTitles

`func (o *Item) HasEbayTitles() bool`

HasEbayTitles returns a boolean if a field has been set.

### GetItemShippingProfiles

`func (o *Item) GetItemShippingProfiles() []ItemShippingProfile`

GetItemShippingProfiles returns the ItemShippingProfiles field if non-nil, zero value otherwise.

### GetItemShippingProfilesOk

`func (o *Item) GetItemShippingProfilesOk() (*[]ItemShippingProfile, bool)`

GetItemShippingProfilesOk returns a tuple with the ItemShippingProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemShippingProfiles

`func (o *Item) SetItemShippingProfiles(v []ItemShippingProfile)`

SetItemShippingProfiles sets ItemShippingProfiles field to given value.

### HasItemShippingProfiles

`func (o *Item) HasItemShippingProfiles() bool`

HasItemShippingProfiles returns a boolean if a field has been set.

### GetProfileId

`func (o *Item) GetProfileId() int32`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *Item) GetProfileIdOk() (*int32, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *Item) SetProfileId(v int32)`

SetProfileId sets ProfileId field to given value.


### GetItemProperties

`func (o *Item) GetItemProperties() []ItemProperty`

GetItemProperties returns the ItemProperties field if non-nil, zero value otherwise.

### GetItemPropertiesOk

`func (o *Item) GetItemPropertiesOk() (*[]ItemProperty, bool)`

GetItemPropertiesOk returns a tuple with the ItemProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemProperties

`func (o *Item) SetItemProperties(v []ItemProperty)`

SetItemProperties sets ItemProperties field to given value.

### HasItemProperties

`func (o *Item) HasItemProperties() bool`

HasItemProperties returns a boolean if a field has been set.

### GetPropertyId

`func (o *Item) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *Item) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *Item) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.


### GetPropertySelectionId

`func (o *Item) GetPropertySelectionId() int32`

GetPropertySelectionId returns the PropertySelectionId field if non-nil, zero value otherwise.

### GetPropertySelectionIdOk

`func (o *Item) GetPropertySelectionIdOk() (*int32, bool)`

GetPropertySelectionIdOk returns a tuple with the PropertySelectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertySelectionId

`func (o *Item) SetPropertySelectionId(v int32)`

SetPropertySelectionId sets PropertySelectionId field to given value.

### HasPropertySelectionId

`func (o *Item) HasPropertySelectionId() bool`

HasPropertySelectionId returns a boolean if a field has been set.

### GetId

`func (o *Item) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Item) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Item) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Item) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVariations

`func (o *Item) GetVariations() []Variation`

GetVariations returns the Variations field if non-nil, zero value otherwise.

### GetVariationsOk

`func (o *Item) GetVariationsOk() (*[]Variation, bool)`

GetVariationsOk returns a tuple with the Variations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariations

`func (o *Item) SetVariations(v []Variation)`

SetVariations sets Variations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


