# CategoryDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **int32** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Description2** | Pointer to **string** |  | [optional] 
**ShortDescription** | Pointer to **string** |  | [optional] 
**MetaKeywords** | Pointer to **string** |  | [optional] 
**MetaDescription** | Pointer to **string** |  | [optional] 
**NameUrl** | Pointer to **string** |  | [optional] 
**MetaTitle** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedBy** | Pointer to **string** |  | [optional] 
**ItemListView** | Pointer to **string** |  | [optional] 
**SingleItemView** | Pointer to **string** |  | [optional] 
**PageView** | Pointer to **string** |  | [optional] 
**Fulltext** | Pointer to **string** |  | [optional] 
**MetaRobots** | Pointer to **string** |  | [optional] 
**CanonicalLink** | Pointer to **string** |  | [optional] 
**PreviewUrl** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**Image2** | Pointer to **string** |  | [optional] 
**Image2Path** | Pointer to **string** |  | [optional] 
**PlentyId** | Pointer to **int32** |  | [optional] 

## Methods

### NewCategoryDetails

`func NewCategoryDetails() *CategoryDetails`

NewCategoryDetails instantiates a new CategoryDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryDetailsWithDefaults

`func NewCategoryDetailsWithDefaults() *CategoryDetails`

NewCategoryDetailsWithDefaults instantiates a new CategoryDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *CategoryDetails) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CategoryDetails) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CategoryDetails) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CategoryDetails) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetLang

`func (o *CategoryDetails) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *CategoryDetails) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *CategoryDetails) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *CategoryDetails) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetName

`func (o *CategoryDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CategoryDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CategoryDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CategoryDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CategoryDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CategoryDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescription2

`func (o *CategoryDetails) GetDescription2() string`

GetDescription2 returns the Description2 field if non-nil, zero value otherwise.

### GetDescription2Ok

`func (o *CategoryDetails) GetDescription2Ok() (*string, bool)`

GetDescription2Ok returns a tuple with the Description2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription2

`func (o *CategoryDetails) SetDescription2(v string)`

SetDescription2 sets Description2 field to given value.

### HasDescription2

`func (o *CategoryDetails) HasDescription2() bool`

HasDescription2 returns a boolean if a field has been set.

### GetShortDescription

`func (o *CategoryDetails) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *CategoryDetails) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *CategoryDetails) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *CategoryDetails) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetMetaKeywords

`func (o *CategoryDetails) GetMetaKeywords() string`

GetMetaKeywords returns the MetaKeywords field if non-nil, zero value otherwise.

### GetMetaKeywordsOk

`func (o *CategoryDetails) GetMetaKeywordsOk() (*string, bool)`

GetMetaKeywordsOk returns a tuple with the MetaKeywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaKeywords

`func (o *CategoryDetails) SetMetaKeywords(v string)`

SetMetaKeywords sets MetaKeywords field to given value.

### HasMetaKeywords

`func (o *CategoryDetails) HasMetaKeywords() bool`

HasMetaKeywords returns a boolean if a field has been set.

### GetMetaDescription

`func (o *CategoryDetails) GetMetaDescription() string`

GetMetaDescription returns the MetaDescription field if non-nil, zero value otherwise.

### GetMetaDescriptionOk

`func (o *CategoryDetails) GetMetaDescriptionOk() (*string, bool)`

GetMetaDescriptionOk returns a tuple with the MetaDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaDescription

`func (o *CategoryDetails) SetMetaDescription(v string)`

SetMetaDescription sets MetaDescription field to given value.

### HasMetaDescription

`func (o *CategoryDetails) HasMetaDescription() bool`

HasMetaDescription returns a boolean if a field has been set.

### GetNameUrl

`func (o *CategoryDetails) GetNameUrl() string`

GetNameUrl returns the NameUrl field if non-nil, zero value otherwise.

### GetNameUrlOk

`func (o *CategoryDetails) GetNameUrlOk() (*string, bool)`

GetNameUrlOk returns a tuple with the NameUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameUrl

`func (o *CategoryDetails) SetNameUrl(v string)`

SetNameUrl sets NameUrl field to given value.

### HasNameUrl

`func (o *CategoryDetails) HasNameUrl() bool`

HasNameUrl returns a boolean if a field has been set.

### GetMetaTitle

`func (o *CategoryDetails) GetMetaTitle() string`

GetMetaTitle returns the MetaTitle field if non-nil, zero value otherwise.

### GetMetaTitleOk

`func (o *CategoryDetails) GetMetaTitleOk() (*string, bool)`

GetMetaTitleOk returns a tuple with the MetaTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaTitle

`func (o *CategoryDetails) SetMetaTitle(v string)`

SetMetaTitle sets MetaTitle field to given value.

### HasMetaTitle

`func (o *CategoryDetails) HasMetaTitle() bool`

HasMetaTitle returns a boolean if a field has been set.

### GetPosition

`func (o *CategoryDetails) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *CategoryDetails) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *CategoryDetails) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *CategoryDetails) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CategoryDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CategoryDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CategoryDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CategoryDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *CategoryDetails) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *CategoryDetails) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *CategoryDetails) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *CategoryDetails) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetItemListView

`func (o *CategoryDetails) GetItemListView() string`

GetItemListView returns the ItemListView field if non-nil, zero value otherwise.

### GetItemListViewOk

`func (o *CategoryDetails) GetItemListViewOk() (*string, bool)`

GetItemListViewOk returns a tuple with the ItemListView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemListView

`func (o *CategoryDetails) SetItemListView(v string)`

SetItemListView sets ItemListView field to given value.

### HasItemListView

`func (o *CategoryDetails) HasItemListView() bool`

HasItemListView returns a boolean if a field has been set.

### GetSingleItemView

`func (o *CategoryDetails) GetSingleItemView() string`

GetSingleItemView returns the SingleItemView field if non-nil, zero value otherwise.

### GetSingleItemViewOk

`func (o *CategoryDetails) GetSingleItemViewOk() (*string, bool)`

GetSingleItemViewOk returns a tuple with the SingleItemView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleItemView

`func (o *CategoryDetails) SetSingleItemView(v string)`

SetSingleItemView sets SingleItemView field to given value.

### HasSingleItemView

`func (o *CategoryDetails) HasSingleItemView() bool`

HasSingleItemView returns a boolean if a field has been set.

### GetPageView

`func (o *CategoryDetails) GetPageView() string`

GetPageView returns the PageView field if non-nil, zero value otherwise.

### GetPageViewOk

`func (o *CategoryDetails) GetPageViewOk() (*string, bool)`

GetPageViewOk returns a tuple with the PageView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageView

`func (o *CategoryDetails) SetPageView(v string)`

SetPageView sets PageView field to given value.

### HasPageView

`func (o *CategoryDetails) HasPageView() bool`

HasPageView returns a boolean if a field has been set.

### GetFulltext

`func (o *CategoryDetails) GetFulltext() string`

GetFulltext returns the Fulltext field if non-nil, zero value otherwise.

### GetFulltextOk

`func (o *CategoryDetails) GetFulltextOk() (*string, bool)`

GetFulltextOk returns a tuple with the Fulltext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulltext

`func (o *CategoryDetails) SetFulltext(v string)`

SetFulltext sets Fulltext field to given value.

### HasFulltext

`func (o *CategoryDetails) HasFulltext() bool`

HasFulltext returns a boolean if a field has been set.

### GetMetaRobots

`func (o *CategoryDetails) GetMetaRobots() string`

GetMetaRobots returns the MetaRobots field if non-nil, zero value otherwise.

### GetMetaRobotsOk

`func (o *CategoryDetails) GetMetaRobotsOk() (*string, bool)`

GetMetaRobotsOk returns a tuple with the MetaRobots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaRobots

`func (o *CategoryDetails) SetMetaRobots(v string)`

SetMetaRobots sets MetaRobots field to given value.

### HasMetaRobots

`func (o *CategoryDetails) HasMetaRobots() bool`

HasMetaRobots returns a boolean if a field has been set.

### GetCanonicalLink

`func (o *CategoryDetails) GetCanonicalLink() string`

GetCanonicalLink returns the CanonicalLink field if non-nil, zero value otherwise.

### GetCanonicalLinkOk

`func (o *CategoryDetails) GetCanonicalLinkOk() (*string, bool)`

GetCanonicalLinkOk returns a tuple with the CanonicalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalLink

`func (o *CategoryDetails) SetCanonicalLink(v string)`

SetCanonicalLink sets CanonicalLink field to given value.

### HasCanonicalLink

`func (o *CategoryDetails) HasCanonicalLink() bool`

HasCanonicalLink returns a boolean if a field has been set.

### GetPreviewUrl

`func (o *CategoryDetails) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *CategoryDetails) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *CategoryDetails) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.

### HasPreviewUrl

`func (o *CategoryDetails) HasPreviewUrl() bool`

HasPreviewUrl returns a boolean if a field has been set.

### GetImage

`func (o *CategoryDetails) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CategoryDetails) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CategoryDetails) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CategoryDetails) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImagePath

`func (o *CategoryDetails) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *CategoryDetails) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *CategoryDetails) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *CategoryDetails) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetImage2

`func (o *CategoryDetails) GetImage2() string`

GetImage2 returns the Image2 field if non-nil, zero value otherwise.

### GetImage2Ok

`func (o *CategoryDetails) GetImage2Ok() (*string, bool)`

GetImage2Ok returns a tuple with the Image2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2

`func (o *CategoryDetails) SetImage2(v string)`

SetImage2 sets Image2 field to given value.

### HasImage2

`func (o *CategoryDetails) HasImage2() bool`

HasImage2 returns a boolean if a field has been set.

### GetImage2Path

`func (o *CategoryDetails) GetImage2Path() string`

GetImage2Path returns the Image2Path field if non-nil, zero value otherwise.

### GetImage2PathOk

`func (o *CategoryDetails) GetImage2PathOk() (*string, bool)`

GetImage2PathOk returns a tuple with the Image2Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage2Path

`func (o *CategoryDetails) SetImage2Path(v string)`

SetImage2Path sets Image2Path field to given value.

### HasImage2Path

`func (o *CategoryDetails) HasImage2Path() bool`

HasImage2Path returns a boolean if a field has been set.

### GetPlentyId

`func (o *CategoryDetails) GetPlentyId() int32`

GetPlentyId returns the PlentyId field if non-nil, zero value otherwise.

### GetPlentyIdOk

`func (o *CategoryDetails) GetPlentyIdOk() (*int32, bool)`

GetPlentyIdOk returns a tuple with the PlentyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlentyId

`func (o *CategoryDetails) SetPlentyId(v int32)`

SetPlentyId sets PlentyId field to given value.

### HasPlentyId

`func (o *CategoryDetails) HasPlentyId() bool`

HasPlentyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


