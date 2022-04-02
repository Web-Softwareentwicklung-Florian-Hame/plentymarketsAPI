# VariationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**ManufacturerId** | Pointer to **int32** |  | [optional] 
**StockType** | Pointer to **int32** |  | [optional] 
**AddCmsPage** | Pointer to **string** |  | [optional] 
**StoreSpecial** | Pointer to **int32** |  | [optional] 
**Condition** | Pointer to **int32** |  | [optional] 
**AmazonFedas** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Free1** | Pointer to **string** |  | [optional] 
**Free2** | Pointer to **string** |  | [optional] 
**Free3** | Pointer to **string** |  | [optional] 
**Free4** | Pointer to **string** |  | [optional] 
**Free5** | Pointer to **string** |  | [optional] 
**Free6** | Pointer to **string** |  | [optional] 
**Free7** | Pointer to **string** |  | [optional] 
**Free8** | Pointer to **string** |  | [optional] 
**Free9** | Pointer to **string** |  | [optional] 
**Free10** | Pointer to **string** |  | [optional] 
**Free11** | Pointer to **string** |  | [optional] 
**Free12** | Pointer to **string** |  | [optional] 
**Free13** | Pointer to **string** |  | [optional] 
**Free14** | Pointer to **string** |  | [optional] 
**Free15** | Pointer to **string** |  | [optional] 
**Free16** | Pointer to **string** |  | [optional] 
**Free17** | Pointer to **string** |  | [optional] 
**Free18** | Pointer to **string** |  | [optional] 
**Free19** | Pointer to **string** |  | [optional] 
**Free20** | Pointer to **string** |  | [optional] 
**CustomsTariffNumber** | Pointer to **string** |  | [optional] 
**ProducingCountryId** | Pointer to **int32** |  | [optional] 
**RevenueAccount** | Pointer to **int32** |  | [optional] 
**CouponRestriction** | Pointer to **int32** |  | [optional] 
**FlagOne** | Pointer to **int32** |  | [optional] 
**FlagTwo** | Pointer to **int32** |  | [optional] 
**AgeRestriction** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**AmazonProductType** | Pointer to **int32** |  | [optional] 
**EbayPresetId** | Pointer to **int32** |  | [optional] 
**EbayCategory** | Pointer to **string** |  | [optional] 
**EbayCategory2** | Pointer to **string** |  | [optional] 
**EbayStoreCategory** | Pointer to **string** |  | [optional] 
**EbayStoreCategory2** | Pointer to **string** |  | [optional] 
**AmazonFbaPlatform** | Pointer to **int32** |  | [optional] 
**Feedback** | Pointer to **int32** |  | [optional] 
**Gimahhot** | Pointer to **string** |  | [optional] 
**MaximumOrderQuantity** | Pointer to **float32** |  | [optional] 
**IsSubscribable** | Pointer to **bool** |  | [optional] 
**RakutenCategoryId** | Pointer to **int32** |  | [optional] 
**IsShippingPackage** | Pointer to **bool** |  | [optional] 
**ConditionApi** | Pointer to **float32** |  | [optional] 
**IsSerialNumber** | Pointer to **bool** |  | [optional] 
**IsShippableByAmazon** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **int32** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**MainVariationId** | Pointer to **int32** |  | [optional] 

## Methods

### NewVariationItem

`func NewVariationItem() *VariationItem`

NewVariationItem instantiates a new VariationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationItemWithDefaults

`func NewVariationItemWithDefaults() *VariationItem`

NewVariationItemWithDefaults instantiates a new VariationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *VariationItem) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VariationItem) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VariationItem) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VariationItem) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetManufacturerId

`func (o *VariationItem) GetManufacturerId() int32`

GetManufacturerId returns the ManufacturerId field if non-nil, zero value otherwise.

### GetManufacturerIdOk

`func (o *VariationItem) GetManufacturerIdOk() (*int32, bool)`

GetManufacturerIdOk returns a tuple with the ManufacturerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerId

`func (o *VariationItem) SetManufacturerId(v int32)`

SetManufacturerId sets ManufacturerId field to given value.

### HasManufacturerId

`func (o *VariationItem) HasManufacturerId() bool`

HasManufacturerId returns a boolean if a field has been set.

### GetStockType

`func (o *VariationItem) GetStockType() int32`

GetStockType returns the StockType field if non-nil, zero value otherwise.

### GetStockTypeOk

`func (o *VariationItem) GetStockTypeOk() (*int32, bool)`

GetStockTypeOk returns a tuple with the StockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockType

`func (o *VariationItem) SetStockType(v int32)`

SetStockType sets StockType field to given value.

### HasStockType

`func (o *VariationItem) HasStockType() bool`

HasStockType returns a boolean if a field has been set.

### GetAddCmsPage

`func (o *VariationItem) GetAddCmsPage() string`

GetAddCmsPage returns the AddCmsPage field if non-nil, zero value otherwise.

### GetAddCmsPageOk

`func (o *VariationItem) GetAddCmsPageOk() (*string, bool)`

GetAddCmsPageOk returns a tuple with the AddCmsPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddCmsPage

`func (o *VariationItem) SetAddCmsPage(v string)`

SetAddCmsPage sets AddCmsPage field to given value.

### HasAddCmsPage

`func (o *VariationItem) HasAddCmsPage() bool`

HasAddCmsPage returns a boolean if a field has been set.

### GetStoreSpecial

`func (o *VariationItem) GetStoreSpecial() int32`

GetStoreSpecial returns the StoreSpecial field if non-nil, zero value otherwise.

### GetStoreSpecialOk

`func (o *VariationItem) GetStoreSpecialOk() (*int32, bool)`

GetStoreSpecialOk returns a tuple with the StoreSpecial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSpecial

`func (o *VariationItem) SetStoreSpecial(v int32)`

SetStoreSpecial sets StoreSpecial field to given value.

### HasStoreSpecial

`func (o *VariationItem) HasStoreSpecial() bool`

HasStoreSpecial returns a boolean if a field has been set.

### GetCondition

`func (o *VariationItem) GetCondition() int32`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *VariationItem) GetConditionOk() (*int32, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *VariationItem) SetCondition(v int32)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *VariationItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetAmazonFedas

`func (o *VariationItem) GetAmazonFedas() string`

GetAmazonFedas returns the AmazonFedas field if non-nil, zero value otherwise.

### GetAmazonFedasOk

`func (o *VariationItem) GetAmazonFedasOk() (*string, bool)`

GetAmazonFedasOk returns a tuple with the AmazonFedas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFedas

`func (o *VariationItem) SetAmazonFedas(v string)`

SetAmazonFedas sets AmazonFedas field to given value.

### HasAmazonFedas

`func (o *VariationItem) HasAmazonFedas() bool`

HasAmazonFedas returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VariationItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VariationItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VariationItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VariationItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFree1

`func (o *VariationItem) GetFree1() string`

GetFree1 returns the Free1 field if non-nil, zero value otherwise.

### GetFree1Ok

`func (o *VariationItem) GetFree1Ok() (*string, bool)`

GetFree1Ok returns a tuple with the Free1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree1

`func (o *VariationItem) SetFree1(v string)`

SetFree1 sets Free1 field to given value.

### HasFree1

`func (o *VariationItem) HasFree1() bool`

HasFree1 returns a boolean if a field has been set.

### GetFree2

`func (o *VariationItem) GetFree2() string`

GetFree2 returns the Free2 field if non-nil, zero value otherwise.

### GetFree2Ok

`func (o *VariationItem) GetFree2Ok() (*string, bool)`

GetFree2Ok returns a tuple with the Free2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree2

`func (o *VariationItem) SetFree2(v string)`

SetFree2 sets Free2 field to given value.

### HasFree2

`func (o *VariationItem) HasFree2() bool`

HasFree2 returns a boolean if a field has been set.

### GetFree3

`func (o *VariationItem) GetFree3() string`

GetFree3 returns the Free3 field if non-nil, zero value otherwise.

### GetFree3Ok

`func (o *VariationItem) GetFree3Ok() (*string, bool)`

GetFree3Ok returns a tuple with the Free3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree3

`func (o *VariationItem) SetFree3(v string)`

SetFree3 sets Free3 field to given value.

### HasFree3

`func (o *VariationItem) HasFree3() bool`

HasFree3 returns a boolean if a field has been set.

### GetFree4

`func (o *VariationItem) GetFree4() string`

GetFree4 returns the Free4 field if non-nil, zero value otherwise.

### GetFree4Ok

`func (o *VariationItem) GetFree4Ok() (*string, bool)`

GetFree4Ok returns a tuple with the Free4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree4

`func (o *VariationItem) SetFree4(v string)`

SetFree4 sets Free4 field to given value.

### HasFree4

`func (o *VariationItem) HasFree4() bool`

HasFree4 returns a boolean if a field has been set.

### GetFree5

`func (o *VariationItem) GetFree5() string`

GetFree5 returns the Free5 field if non-nil, zero value otherwise.

### GetFree5Ok

`func (o *VariationItem) GetFree5Ok() (*string, bool)`

GetFree5Ok returns a tuple with the Free5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree5

`func (o *VariationItem) SetFree5(v string)`

SetFree5 sets Free5 field to given value.

### HasFree5

`func (o *VariationItem) HasFree5() bool`

HasFree5 returns a boolean if a field has been set.

### GetFree6

`func (o *VariationItem) GetFree6() string`

GetFree6 returns the Free6 field if non-nil, zero value otherwise.

### GetFree6Ok

`func (o *VariationItem) GetFree6Ok() (*string, bool)`

GetFree6Ok returns a tuple with the Free6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree6

`func (o *VariationItem) SetFree6(v string)`

SetFree6 sets Free6 field to given value.

### HasFree6

`func (o *VariationItem) HasFree6() bool`

HasFree6 returns a boolean if a field has been set.

### GetFree7

`func (o *VariationItem) GetFree7() string`

GetFree7 returns the Free7 field if non-nil, zero value otherwise.

### GetFree7Ok

`func (o *VariationItem) GetFree7Ok() (*string, bool)`

GetFree7Ok returns a tuple with the Free7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree7

`func (o *VariationItem) SetFree7(v string)`

SetFree7 sets Free7 field to given value.

### HasFree7

`func (o *VariationItem) HasFree7() bool`

HasFree7 returns a boolean if a field has been set.

### GetFree8

`func (o *VariationItem) GetFree8() string`

GetFree8 returns the Free8 field if non-nil, zero value otherwise.

### GetFree8Ok

`func (o *VariationItem) GetFree8Ok() (*string, bool)`

GetFree8Ok returns a tuple with the Free8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree8

`func (o *VariationItem) SetFree8(v string)`

SetFree8 sets Free8 field to given value.

### HasFree8

`func (o *VariationItem) HasFree8() bool`

HasFree8 returns a boolean if a field has been set.

### GetFree9

`func (o *VariationItem) GetFree9() string`

GetFree9 returns the Free9 field if non-nil, zero value otherwise.

### GetFree9Ok

`func (o *VariationItem) GetFree9Ok() (*string, bool)`

GetFree9Ok returns a tuple with the Free9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree9

`func (o *VariationItem) SetFree9(v string)`

SetFree9 sets Free9 field to given value.

### HasFree9

`func (o *VariationItem) HasFree9() bool`

HasFree9 returns a boolean if a field has been set.

### GetFree10

`func (o *VariationItem) GetFree10() string`

GetFree10 returns the Free10 field if non-nil, zero value otherwise.

### GetFree10Ok

`func (o *VariationItem) GetFree10Ok() (*string, bool)`

GetFree10Ok returns a tuple with the Free10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree10

`func (o *VariationItem) SetFree10(v string)`

SetFree10 sets Free10 field to given value.

### HasFree10

`func (o *VariationItem) HasFree10() bool`

HasFree10 returns a boolean if a field has been set.

### GetFree11

`func (o *VariationItem) GetFree11() string`

GetFree11 returns the Free11 field if non-nil, zero value otherwise.

### GetFree11Ok

`func (o *VariationItem) GetFree11Ok() (*string, bool)`

GetFree11Ok returns a tuple with the Free11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree11

`func (o *VariationItem) SetFree11(v string)`

SetFree11 sets Free11 field to given value.

### HasFree11

`func (o *VariationItem) HasFree11() bool`

HasFree11 returns a boolean if a field has been set.

### GetFree12

`func (o *VariationItem) GetFree12() string`

GetFree12 returns the Free12 field if non-nil, zero value otherwise.

### GetFree12Ok

`func (o *VariationItem) GetFree12Ok() (*string, bool)`

GetFree12Ok returns a tuple with the Free12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree12

`func (o *VariationItem) SetFree12(v string)`

SetFree12 sets Free12 field to given value.

### HasFree12

`func (o *VariationItem) HasFree12() bool`

HasFree12 returns a boolean if a field has been set.

### GetFree13

`func (o *VariationItem) GetFree13() string`

GetFree13 returns the Free13 field if non-nil, zero value otherwise.

### GetFree13Ok

`func (o *VariationItem) GetFree13Ok() (*string, bool)`

GetFree13Ok returns a tuple with the Free13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree13

`func (o *VariationItem) SetFree13(v string)`

SetFree13 sets Free13 field to given value.

### HasFree13

`func (o *VariationItem) HasFree13() bool`

HasFree13 returns a boolean if a field has been set.

### GetFree14

`func (o *VariationItem) GetFree14() string`

GetFree14 returns the Free14 field if non-nil, zero value otherwise.

### GetFree14Ok

`func (o *VariationItem) GetFree14Ok() (*string, bool)`

GetFree14Ok returns a tuple with the Free14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree14

`func (o *VariationItem) SetFree14(v string)`

SetFree14 sets Free14 field to given value.

### HasFree14

`func (o *VariationItem) HasFree14() bool`

HasFree14 returns a boolean if a field has been set.

### GetFree15

`func (o *VariationItem) GetFree15() string`

GetFree15 returns the Free15 field if non-nil, zero value otherwise.

### GetFree15Ok

`func (o *VariationItem) GetFree15Ok() (*string, bool)`

GetFree15Ok returns a tuple with the Free15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree15

`func (o *VariationItem) SetFree15(v string)`

SetFree15 sets Free15 field to given value.

### HasFree15

`func (o *VariationItem) HasFree15() bool`

HasFree15 returns a boolean if a field has been set.

### GetFree16

`func (o *VariationItem) GetFree16() string`

GetFree16 returns the Free16 field if non-nil, zero value otherwise.

### GetFree16Ok

`func (o *VariationItem) GetFree16Ok() (*string, bool)`

GetFree16Ok returns a tuple with the Free16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree16

`func (o *VariationItem) SetFree16(v string)`

SetFree16 sets Free16 field to given value.

### HasFree16

`func (o *VariationItem) HasFree16() bool`

HasFree16 returns a boolean if a field has been set.

### GetFree17

`func (o *VariationItem) GetFree17() string`

GetFree17 returns the Free17 field if non-nil, zero value otherwise.

### GetFree17Ok

`func (o *VariationItem) GetFree17Ok() (*string, bool)`

GetFree17Ok returns a tuple with the Free17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree17

`func (o *VariationItem) SetFree17(v string)`

SetFree17 sets Free17 field to given value.

### HasFree17

`func (o *VariationItem) HasFree17() bool`

HasFree17 returns a boolean if a field has been set.

### GetFree18

`func (o *VariationItem) GetFree18() string`

GetFree18 returns the Free18 field if non-nil, zero value otherwise.

### GetFree18Ok

`func (o *VariationItem) GetFree18Ok() (*string, bool)`

GetFree18Ok returns a tuple with the Free18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree18

`func (o *VariationItem) SetFree18(v string)`

SetFree18 sets Free18 field to given value.

### HasFree18

`func (o *VariationItem) HasFree18() bool`

HasFree18 returns a boolean if a field has been set.

### GetFree19

`func (o *VariationItem) GetFree19() string`

GetFree19 returns the Free19 field if non-nil, zero value otherwise.

### GetFree19Ok

`func (o *VariationItem) GetFree19Ok() (*string, bool)`

GetFree19Ok returns a tuple with the Free19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree19

`func (o *VariationItem) SetFree19(v string)`

SetFree19 sets Free19 field to given value.

### HasFree19

`func (o *VariationItem) HasFree19() bool`

HasFree19 returns a boolean if a field has been set.

### GetFree20

`func (o *VariationItem) GetFree20() string`

GetFree20 returns the Free20 field if non-nil, zero value otherwise.

### GetFree20Ok

`func (o *VariationItem) GetFree20Ok() (*string, bool)`

GetFree20Ok returns a tuple with the Free20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree20

`func (o *VariationItem) SetFree20(v string)`

SetFree20 sets Free20 field to given value.

### HasFree20

`func (o *VariationItem) HasFree20() bool`

HasFree20 returns a boolean if a field has been set.

### GetCustomsTariffNumber

`func (o *VariationItem) GetCustomsTariffNumber() string`

GetCustomsTariffNumber returns the CustomsTariffNumber field if non-nil, zero value otherwise.

### GetCustomsTariffNumberOk

`func (o *VariationItem) GetCustomsTariffNumberOk() (*string, bool)`

GetCustomsTariffNumberOk returns a tuple with the CustomsTariffNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomsTariffNumber

`func (o *VariationItem) SetCustomsTariffNumber(v string)`

SetCustomsTariffNumber sets CustomsTariffNumber field to given value.

### HasCustomsTariffNumber

`func (o *VariationItem) HasCustomsTariffNumber() bool`

HasCustomsTariffNumber returns a boolean if a field has been set.

### GetProducingCountryId

`func (o *VariationItem) GetProducingCountryId() int32`

GetProducingCountryId returns the ProducingCountryId field if non-nil, zero value otherwise.

### GetProducingCountryIdOk

`func (o *VariationItem) GetProducingCountryIdOk() (*int32, bool)`

GetProducingCountryIdOk returns a tuple with the ProducingCountryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducingCountryId

`func (o *VariationItem) SetProducingCountryId(v int32)`

SetProducingCountryId sets ProducingCountryId field to given value.

### HasProducingCountryId

`func (o *VariationItem) HasProducingCountryId() bool`

HasProducingCountryId returns a boolean if a field has been set.

### GetRevenueAccount

`func (o *VariationItem) GetRevenueAccount() int32`

GetRevenueAccount returns the RevenueAccount field if non-nil, zero value otherwise.

### GetRevenueAccountOk

`func (o *VariationItem) GetRevenueAccountOk() (*int32, bool)`

GetRevenueAccountOk returns a tuple with the RevenueAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueAccount

`func (o *VariationItem) SetRevenueAccount(v int32)`

SetRevenueAccount sets RevenueAccount field to given value.

### HasRevenueAccount

`func (o *VariationItem) HasRevenueAccount() bool`

HasRevenueAccount returns a boolean if a field has been set.

### GetCouponRestriction

`func (o *VariationItem) GetCouponRestriction() int32`

GetCouponRestriction returns the CouponRestriction field if non-nil, zero value otherwise.

### GetCouponRestrictionOk

`func (o *VariationItem) GetCouponRestrictionOk() (*int32, bool)`

GetCouponRestrictionOk returns a tuple with the CouponRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRestriction

`func (o *VariationItem) SetCouponRestriction(v int32)`

SetCouponRestriction sets CouponRestriction field to given value.

### HasCouponRestriction

`func (o *VariationItem) HasCouponRestriction() bool`

HasCouponRestriction returns a boolean if a field has been set.

### GetFlagOne

`func (o *VariationItem) GetFlagOne() int32`

GetFlagOne returns the FlagOne field if non-nil, zero value otherwise.

### GetFlagOneOk

`func (o *VariationItem) GetFlagOneOk() (*int32, bool)`

GetFlagOneOk returns a tuple with the FlagOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagOne

`func (o *VariationItem) SetFlagOne(v int32)`

SetFlagOne sets FlagOne field to given value.

### HasFlagOne

`func (o *VariationItem) HasFlagOne() bool`

HasFlagOne returns a boolean if a field has been set.

### GetFlagTwo

`func (o *VariationItem) GetFlagTwo() int32`

GetFlagTwo returns the FlagTwo field if non-nil, zero value otherwise.

### GetFlagTwoOk

`func (o *VariationItem) GetFlagTwoOk() (*int32, bool)`

GetFlagTwoOk returns a tuple with the FlagTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagTwo

`func (o *VariationItem) SetFlagTwo(v int32)`

SetFlagTwo sets FlagTwo field to given value.

### HasFlagTwo

`func (o *VariationItem) HasFlagTwo() bool`

HasFlagTwo returns a boolean if a field has been set.

### GetAgeRestriction

`func (o *VariationItem) GetAgeRestriction() int32`

GetAgeRestriction returns the AgeRestriction field if non-nil, zero value otherwise.

### GetAgeRestrictionOk

`func (o *VariationItem) GetAgeRestrictionOk() (*int32, bool)`

GetAgeRestrictionOk returns a tuple with the AgeRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRestriction

`func (o *VariationItem) SetAgeRestriction(v int32)`

SetAgeRestriction sets AgeRestriction field to given value.

### HasAgeRestriction

`func (o *VariationItem) HasAgeRestriction() bool`

HasAgeRestriction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VariationItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VariationItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VariationItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VariationItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAmazonProductType

`func (o *VariationItem) GetAmazonProductType() int32`

GetAmazonProductType returns the AmazonProductType field if non-nil, zero value otherwise.

### GetAmazonProductTypeOk

`func (o *VariationItem) GetAmazonProductTypeOk() (*int32, bool)`

GetAmazonProductTypeOk returns a tuple with the AmazonProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonProductType

`func (o *VariationItem) SetAmazonProductType(v int32)`

SetAmazonProductType sets AmazonProductType field to given value.

### HasAmazonProductType

`func (o *VariationItem) HasAmazonProductType() bool`

HasAmazonProductType returns a boolean if a field has been set.

### GetEbayPresetId

`func (o *VariationItem) GetEbayPresetId() int32`

GetEbayPresetId returns the EbayPresetId field if non-nil, zero value otherwise.

### GetEbayPresetIdOk

`func (o *VariationItem) GetEbayPresetIdOk() (*int32, bool)`

GetEbayPresetIdOk returns a tuple with the EbayPresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayPresetId

`func (o *VariationItem) SetEbayPresetId(v int32)`

SetEbayPresetId sets EbayPresetId field to given value.

### HasEbayPresetId

`func (o *VariationItem) HasEbayPresetId() bool`

HasEbayPresetId returns a boolean if a field has been set.

### GetEbayCategory

`func (o *VariationItem) GetEbayCategory() string`

GetEbayCategory returns the EbayCategory field if non-nil, zero value otherwise.

### GetEbayCategoryOk

`func (o *VariationItem) GetEbayCategoryOk() (*string, bool)`

GetEbayCategoryOk returns a tuple with the EbayCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory

`func (o *VariationItem) SetEbayCategory(v string)`

SetEbayCategory sets EbayCategory field to given value.

### HasEbayCategory

`func (o *VariationItem) HasEbayCategory() bool`

HasEbayCategory returns a boolean if a field has been set.

### GetEbayCategory2

`func (o *VariationItem) GetEbayCategory2() string`

GetEbayCategory2 returns the EbayCategory2 field if non-nil, zero value otherwise.

### GetEbayCategory2Ok

`func (o *VariationItem) GetEbayCategory2Ok() (*string, bool)`

GetEbayCategory2Ok returns a tuple with the EbayCategory2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayCategory2

`func (o *VariationItem) SetEbayCategory2(v string)`

SetEbayCategory2 sets EbayCategory2 field to given value.

### HasEbayCategory2

`func (o *VariationItem) HasEbayCategory2() bool`

HasEbayCategory2 returns a boolean if a field has been set.

### GetEbayStoreCategory

`func (o *VariationItem) GetEbayStoreCategory() string`

GetEbayStoreCategory returns the EbayStoreCategory field if non-nil, zero value otherwise.

### GetEbayStoreCategoryOk

`func (o *VariationItem) GetEbayStoreCategoryOk() (*string, bool)`

GetEbayStoreCategoryOk returns a tuple with the EbayStoreCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayStoreCategory

`func (o *VariationItem) SetEbayStoreCategory(v string)`

SetEbayStoreCategory sets EbayStoreCategory field to given value.

### HasEbayStoreCategory

`func (o *VariationItem) HasEbayStoreCategory() bool`

HasEbayStoreCategory returns a boolean if a field has been set.

### GetEbayStoreCategory2

`func (o *VariationItem) GetEbayStoreCategory2() string`

GetEbayStoreCategory2 returns the EbayStoreCategory2 field if non-nil, zero value otherwise.

### GetEbayStoreCategory2Ok

`func (o *VariationItem) GetEbayStoreCategory2Ok() (*string, bool)`

GetEbayStoreCategory2Ok returns a tuple with the EbayStoreCategory2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayStoreCategory2

`func (o *VariationItem) SetEbayStoreCategory2(v string)`

SetEbayStoreCategory2 sets EbayStoreCategory2 field to given value.

### HasEbayStoreCategory2

`func (o *VariationItem) HasEbayStoreCategory2() bool`

HasEbayStoreCategory2 returns a boolean if a field has been set.

### GetAmazonFbaPlatform

`func (o *VariationItem) GetAmazonFbaPlatform() int32`

GetAmazonFbaPlatform returns the AmazonFbaPlatform field if non-nil, zero value otherwise.

### GetAmazonFbaPlatformOk

`func (o *VariationItem) GetAmazonFbaPlatformOk() (*int32, bool)`

GetAmazonFbaPlatformOk returns a tuple with the AmazonFbaPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonFbaPlatform

`func (o *VariationItem) SetAmazonFbaPlatform(v int32)`

SetAmazonFbaPlatform sets AmazonFbaPlatform field to given value.

### HasAmazonFbaPlatform

`func (o *VariationItem) HasAmazonFbaPlatform() bool`

HasAmazonFbaPlatform returns a boolean if a field has been set.

### GetFeedback

`func (o *VariationItem) GetFeedback() int32`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *VariationItem) GetFeedbackOk() (*int32, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *VariationItem) SetFeedback(v int32)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *VariationItem) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.

### GetGimahhot

`func (o *VariationItem) GetGimahhot() string`

GetGimahhot returns the Gimahhot field if non-nil, zero value otherwise.

### GetGimahhotOk

`func (o *VariationItem) GetGimahhotOk() (*string, bool)`

GetGimahhotOk returns a tuple with the Gimahhot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGimahhot

`func (o *VariationItem) SetGimahhot(v string)`

SetGimahhot sets Gimahhot field to given value.

### HasGimahhot

`func (o *VariationItem) HasGimahhot() bool`

HasGimahhot returns a boolean if a field has been set.

### GetMaximumOrderQuantity

`func (o *VariationItem) GetMaximumOrderQuantity() float32`

GetMaximumOrderQuantity returns the MaximumOrderQuantity field if non-nil, zero value otherwise.

### GetMaximumOrderQuantityOk

`func (o *VariationItem) GetMaximumOrderQuantityOk() (*float32, bool)`

GetMaximumOrderQuantityOk returns a tuple with the MaximumOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumOrderQuantity

`func (o *VariationItem) SetMaximumOrderQuantity(v float32)`

SetMaximumOrderQuantity sets MaximumOrderQuantity field to given value.

### HasMaximumOrderQuantity

`func (o *VariationItem) HasMaximumOrderQuantity() bool`

HasMaximumOrderQuantity returns a boolean if a field has been set.

### GetIsSubscribable

`func (o *VariationItem) GetIsSubscribable() bool`

GetIsSubscribable returns the IsSubscribable field if non-nil, zero value otherwise.

### GetIsSubscribableOk

`func (o *VariationItem) GetIsSubscribableOk() (*bool, bool)`

GetIsSubscribableOk returns a tuple with the IsSubscribable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribable

`func (o *VariationItem) SetIsSubscribable(v bool)`

SetIsSubscribable sets IsSubscribable field to given value.

### HasIsSubscribable

`func (o *VariationItem) HasIsSubscribable() bool`

HasIsSubscribable returns a boolean if a field has been set.

### GetRakutenCategoryId

`func (o *VariationItem) GetRakutenCategoryId() int32`

GetRakutenCategoryId returns the RakutenCategoryId field if non-nil, zero value otherwise.

### GetRakutenCategoryIdOk

`func (o *VariationItem) GetRakutenCategoryIdOk() (*int32, bool)`

GetRakutenCategoryIdOk returns a tuple with the RakutenCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRakutenCategoryId

`func (o *VariationItem) SetRakutenCategoryId(v int32)`

SetRakutenCategoryId sets RakutenCategoryId field to given value.

### HasRakutenCategoryId

`func (o *VariationItem) HasRakutenCategoryId() bool`

HasRakutenCategoryId returns a boolean if a field has been set.

### GetIsShippingPackage

`func (o *VariationItem) GetIsShippingPackage() bool`

GetIsShippingPackage returns the IsShippingPackage field if non-nil, zero value otherwise.

### GetIsShippingPackageOk

`func (o *VariationItem) GetIsShippingPackageOk() (*bool, bool)`

GetIsShippingPackageOk returns a tuple with the IsShippingPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippingPackage

`func (o *VariationItem) SetIsShippingPackage(v bool)`

SetIsShippingPackage sets IsShippingPackage field to given value.

### HasIsShippingPackage

`func (o *VariationItem) HasIsShippingPackage() bool`

HasIsShippingPackage returns a boolean if a field has been set.

### GetConditionApi

`func (o *VariationItem) GetConditionApi() float32`

GetConditionApi returns the ConditionApi field if non-nil, zero value otherwise.

### GetConditionApiOk

`func (o *VariationItem) GetConditionApiOk() (*float32, bool)`

GetConditionApiOk returns a tuple with the ConditionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionApi

`func (o *VariationItem) SetConditionApi(v float32)`

SetConditionApi sets ConditionApi field to given value.

### HasConditionApi

`func (o *VariationItem) HasConditionApi() bool`

HasConditionApi returns a boolean if a field has been set.

### GetIsSerialNumber

`func (o *VariationItem) GetIsSerialNumber() bool`

GetIsSerialNumber returns the IsSerialNumber field if non-nil, zero value otherwise.

### GetIsSerialNumberOk

`func (o *VariationItem) GetIsSerialNumberOk() (*bool, bool)`

GetIsSerialNumberOk returns a tuple with the IsSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSerialNumber

`func (o *VariationItem) SetIsSerialNumber(v bool)`

SetIsSerialNumber sets IsSerialNumber field to given value.

### HasIsSerialNumber

`func (o *VariationItem) HasIsSerialNumber() bool`

HasIsSerialNumber returns a boolean if a field has been set.

### GetIsShippableByAmazon

`func (o *VariationItem) GetIsShippableByAmazon() bool`

GetIsShippableByAmazon returns the IsShippableByAmazon field if non-nil, zero value otherwise.

### GetIsShippableByAmazonOk

`func (o *VariationItem) GetIsShippableByAmazonOk() (*bool, bool)`

GetIsShippableByAmazonOk returns a tuple with the IsShippableByAmazon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShippableByAmazon

`func (o *VariationItem) SetIsShippableByAmazon(v bool)`

SetIsShippableByAmazon sets IsShippableByAmazon field to given value.

### HasIsShippableByAmazon

`func (o *VariationItem) HasIsShippableByAmazon() bool`

HasIsShippableByAmazon returns a boolean if a field has been set.

### GetOwnerId

`func (o *VariationItem) GetOwnerId() int32`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *VariationItem) GetOwnerIdOk() (*int32, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *VariationItem) SetOwnerId(v int32)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *VariationItem) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetItemType

`func (o *VariationItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *VariationItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *VariationItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *VariationItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetMainVariationId

`func (o *VariationItem) GetMainVariationId() int32`

GetMainVariationId returns the MainVariationId field if non-nil, zero value otherwise.

### GetMainVariationIdOk

`func (o *VariationItem) GetMainVariationIdOk() (*int32, bool)`

GetMainVariationIdOk returns a tuple with the MainVariationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainVariationId

`func (o *VariationItem) SetMainVariationId(v int32)`

SetMainVariationId sets MainVariationId field to given value.

### HasMainVariationId

`func (o *VariationItem) HasMainVariationId() bool`

HasMainVariationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


