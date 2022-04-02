# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ParentCategoryId** | Pointer to **int32** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Linklist** | Pointer to **string** |  | [optional] 
**Right** | Pointer to **string** |  | [optional] 
**Sitemap** | Pointer to **string** |  | [optional] 
**HasChildren** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to [**[]CategoryDetails**](CategoryDetails.md) |  | [optional] 

## Methods

### NewCategory

`func NewCategory() *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Category) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentCategoryId

`func (o *Category) GetParentCategoryId() int32`

GetParentCategoryId returns the ParentCategoryId field if non-nil, zero value otherwise.

### GetParentCategoryIdOk

`func (o *Category) GetParentCategoryIdOk() (*int32, bool)`

GetParentCategoryIdOk returns a tuple with the ParentCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCategoryId

`func (o *Category) SetParentCategoryId(v int32)`

SetParentCategoryId sets ParentCategoryId field to given value.

### HasParentCategoryId

`func (o *Category) HasParentCategoryId() bool`

HasParentCategoryId returns a boolean if a field has been set.

### GetLevel

`func (o *Category) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *Category) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *Category) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *Category) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetType

`func (o *Category) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Category) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Category) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Category) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLinklist

`func (o *Category) GetLinklist() string`

GetLinklist returns the Linklist field if non-nil, zero value otherwise.

### GetLinklistOk

`func (o *Category) GetLinklistOk() (*string, bool)`

GetLinklistOk returns a tuple with the Linklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinklist

`func (o *Category) SetLinklist(v string)`

SetLinklist sets Linklist field to given value.

### HasLinklist

`func (o *Category) HasLinklist() bool`

HasLinklist returns a boolean if a field has been set.

### GetRight

`func (o *Category) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *Category) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *Category) SetRight(v string)`

SetRight sets Right field to given value.

### HasRight

`func (o *Category) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetSitemap

`func (o *Category) GetSitemap() string`

GetSitemap returns the Sitemap field if non-nil, zero value otherwise.

### GetSitemapOk

`func (o *Category) GetSitemapOk() (*string, bool)`

GetSitemapOk returns a tuple with the Sitemap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitemap

`func (o *Category) SetSitemap(v string)`

SetSitemap sets Sitemap field to given value.

### HasSitemap

`func (o *Category) HasSitemap() bool`

HasSitemap returns a boolean if a field has been set.

### GetHasChildren

`func (o *Category) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *Category) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *Category) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.

### HasHasChildren

`func (o *Category) HasHasChildren() bool`

HasHasChildren returns a boolean if a field has been set.

### GetDetails

`func (o *Category) GetDetails() []CategoryDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Category) GetDetailsOk() (*[]CategoryDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Category) SetDetails(v []CategoryDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Category) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


