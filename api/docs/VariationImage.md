# VariationImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ItemId** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**FileType** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **int32** |  | [optional] 
**LastUpdate** | Pointer to **string** |  | [optional] 
**Insert** | Pointer to **string** |  | [optional] 
**Md5Checksum** | Pointer to **string** |  | [optional] 
**Width** | Pointer to **int32** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**StorageProviderId** | Pointer to **string** |  | [optional] 
**Md5ChecksumOriginal** | Pointer to **string** |  | [optional] 
**CleanImageName** | Pointer to **string** |  | [optional] 
**UploadUrl** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UrlMiddle** | Pointer to **string** |  | [optional] 
**UrlPreview** | Pointer to **string** |  | [optional] 
**UrlSecondPreview** | Pointer to **string** |  | [optional] 
**DocumentUploadPath** | Pointer to **string** |  | [optional] 
**DocumentUploadPathPreview** | Pointer to **string** |  | [optional] 
**DocumentUploadPreviewWidth** | Pointer to **int32** |  | [optional] 
**DocumentUploadPreviewHeight** | Pointer to **int32** |  | [optional] 
**HasVariationLink** | Pointer to **bool** |  | [optional] 
**Availabilities** | Pointer to [**[]ImageAvailability**](ImageAvailability.md) |  | [optional] 
**Names** | Pointer to [**[]ImageName**](ImageName.md) |  | [optional] 

## Methods

### NewVariationImage

`func NewVariationImage() *VariationImage`

NewVariationImage instantiates a new VariationImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariationImageWithDefaults

`func NewVariationImageWithDefaults() *VariationImage`

NewVariationImageWithDefaults instantiates a new VariationImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VariationImage) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariationImage) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariationImage) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VariationImage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *VariationImage) GetItemId() int32`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VariationImage) GetItemIdOk() (*int32, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VariationImage) SetItemId(v int32)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *VariationImage) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetType

`func (o *VariationImage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariationImage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariationImage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VariationImage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFileType

`func (o *VariationImage) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *VariationImage) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *VariationImage) SetFileType(v string)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *VariationImage) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetPath

`func (o *VariationImage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VariationImage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VariationImage) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VariationImage) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPosition

`func (o *VariationImage) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VariationImage) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VariationImage) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VariationImage) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLastUpdate

`func (o *VariationImage) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *VariationImage) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *VariationImage) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *VariationImage) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.

### GetInsert

`func (o *VariationImage) GetInsert() string`

GetInsert returns the Insert field if non-nil, zero value otherwise.

### GetInsertOk

`func (o *VariationImage) GetInsertOk() (*string, bool)`

GetInsertOk returns a tuple with the Insert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsert

`func (o *VariationImage) SetInsert(v string)`

SetInsert sets Insert field to given value.

### HasInsert

`func (o *VariationImage) HasInsert() bool`

HasInsert returns a boolean if a field has been set.

### GetMd5Checksum

`func (o *VariationImage) GetMd5Checksum() string`

GetMd5Checksum returns the Md5Checksum field if non-nil, zero value otherwise.

### GetMd5ChecksumOk

`func (o *VariationImage) GetMd5ChecksumOk() (*string, bool)`

GetMd5ChecksumOk returns a tuple with the Md5Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Checksum

`func (o *VariationImage) SetMd5Checksum(v string)`

SetMd5Checksum sets Md5Checksum field to given value.

### HasMd5Checksum

`func (o *VariationImage) HasMd5Checksum() bool`

HasMd5Checksum returns a boolean if a field has been set.

### GetWidth

`func (o *VariationImage) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VariationImage) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VariationImage) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VariationImage) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VariationImage) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VariationImage) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VariationImage) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VariationImage) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetSize

`func (o *VariationImage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VariationImage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VariationImage) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VariationImage) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStorageProviderId

`func (o *VariationImage) GetStorageProviderId() string`

GetStorageProviderId returns the StorageProviderId field if non-nil, zero value otherwise.

### GetStorageProviderIdOk

`func (o *VariationImage) GetStorageProviderIdOk() (*string, bool)`

GetStorageProviderIdOk returns a tuple with the StorageProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProviderId

`func (o *VariationImage) SetStorageProviderId(v string)`

SetStorageProviderId sets StorageProviderId field to given value.

### HasStorageProviderId

`func (o *VariationImage) HasStorageProviderId() bool`

HasStorageProviderId returns a boolean if a field has been set.

### GetMd5ChecksumOriginal

`func (o *VariationImage) GetMd5ChecksumOriginal() string`

GetMd5ChecksumOriginal returns the Md5ChecksumOriginal field if non-nil, zero value otherwise.

### GetMd5ChecksumOriginalOk

`func (o *VariationImage) GetMd5ChecksumOriginalOk() (*string, bool)`

GetMd5ChecksumOriginalOk returns a tuple with the Md5ChecksumOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5ChecksumOriginal

`func (o *VariationImage) SetMd5ChecksumOriginal(v string)`

SetMd5ChecksumOriginal sets Md5ChecksumOriginal field to given value.

### HasMd5ChecksumOriginal

`func (o *VariationImage) HasMd5ChecksumOriginal() bool`

HasMd5ChecksumOriginal returns a boolean if a field has been set.

### GetCleanImageName

`func (o *VariationImage) GetCleanImageName() string`

GetCleanImageName returns the CleanImageName field if non-nil, zero value otherwise.

### GetCleanImageNameOk

`func (o *VariationImage) GetCleanImageNameOk() (*string, bool)`

GetCleanImageNameOk returns a tuple with the CleanImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanImageName

`func (o *VariationImage) SetCleanImageName(v string)`

SetCleanImageName sets CleanImageName field to given value.

### HasCleanImageName

`func (o *VariationImage) HasCleanImageName() bool`

HasCleanImageName returns a boolean if a field has been set.

### GetUploadUrl

`func (o *VariationImage) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *VariationImage) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *VariationImage) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *VariationImage) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUrl

`func (o *VariationImage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VariationImage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VariationImage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VariationImage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlMiddle

`func (o *VariationImage) GetUrlMiddle() string`

GetUrlMiddle returns the UrlMiddle field if non-nil, zero value otherwise.

### GetUrlMiddleOk

`func (o *VariationImage) GetUrlMiddleOk() (*string, bool)`

GetUrlMiddleOk returns a tuple with the UrlMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlMiddle

`func (o *VariationImage) SetUrlMiddle(v string)`

SetUrlMiddle sets UrlMiddle field to given value.

### HasUrlMiddle

`func (o *VariationImage) HasUrlMiddle() bool`

HasUrlMiddle returns a boolean if a field has been set.

### GetUrlPreview

`func (o *VariationImage) GetUrlPreview() string`

GetUrlPreview returns the UrlPreview field if non-nil, zero value otherwise.

### GetUrlPreviewOk

`func (o *VariationImage) GetUrlPreviewOk() (*string, bool)`

GetUrlPreviewOk returns a tuple with the UrlPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPreview

`func (o *VariationImage) SetUrlPreview(v string)`

SetUrlPreview sets UrlPreview field to given value.

### HasUrlPreview

`func (o *VariationImage) HasUrlPreview() bool`

HasUrlPreview returns a boolean if a field has been set.

### GetUrlSecondPreview

`func (o *VariationImage) GetUrlSecondPreview() string`

GetUrlSecondPreview returns the UrlSecondPreview field if non-nil, zero value otherwise.

### GetUrlSecondPreviewOk

`func (o *VariationImage) GetUrlSecondPreviewOk() (*string, bool)`

GetUrlSecondPreviewOk returns a tuple with the UrlSecondPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlSecondPreview

`func (o *VariationImage) SetUrlSecondPreview(v string)`

SetUrlSecondPreview sets UrlSecondPreview field to given value.

### HasUrlSecondPreview

`func (o *VariationImage) HasUrlSecondPreview() bool`

HasUrlSecondPreview returns a boolean if a field has been set.

### GetDocumentUploadPath

`func (o *VariationImage) GetDocumentUploadPath() string`

GetDocumentUploadPath returns the DocumentUploadPath field if non-nil, zero value otherwise.

### GetDocumentUploadPathOk

`func (o *VariationImage) GetDocumentUploadPathOk() (*string, bool)`

GetDocumentUploadPathOk returns a tuple with the DocumentUploadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUploadPath

`func (o *VariationImage) SetDocumentUploadPath(v string)`

SetDocumentUploadPath sets DocumentUploadPath field to given value.

### HasDocumentUploadPath

`func (o *VariationImage) HasDocumentUploadPath() bool`

HasDocumentUploadPath returns a boolean if a field has been set.

### GetDocumentUploadPathPreview

`func (o *VariationImage) GetDocumentUploadPathPreview() string`

GetDocumentUploadPathPreview returns the DocumentUploadPathPreview field if non-nil, zero value otherwise.

### GetDocumentUploadPathPreviewOk

`func (o *VariationImage) GetDocumentUploadPathPreviewOk() (*string, bool)`

GetDocumentUploadPathPreviewOk returns a tuple with the DocumentUploadPathPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUploadPathPreview

`func (o *VariationImage) SetDocumentUploadPathPreview(v string)`

SetDocumentUploadPathPreview sets DocumentUploadPathPreview field to given value.

### HasDocumentUploadPathPreview

`func (o *VariationImage) HasDocumentUploadPathPreview() bool`

HasDocumentUploadPathPreview returns a boolean if a field has been set.

### GetDocumentUploadPreviewWidth

`func (o *VariationImage) GetDocumentUploadPreviewWidth() int32`

GetDocumentUploadPreviewWidth returns the DocumentUploadPreviewWidth field if non-nil, zero value otherwise.

### GetDocumentUploadPreviewWidthOk

`func (o *VariationImage) GetDocumentUploadPreviewWidthOk() (*int32, bool)`

GetDocumentUploadPreviewWidthOk returns a tuple with the DocumentUploadPreviewWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUploadPreviewWidth

`func (o *VariationImage) SetDocumentUploadPreviewWidth(v int32)`

SetDocumentUploadPreviewWidth sets DocumentUploadPreviewWidth field to given value.

### HasDocumentUploadPreviewWidth

`func (o *VariationImage) HasDocumentUploadPreviewWidth() bool`

HasDocumentUploadPreviewWidth returns a boolean if a field has been set.

### GetDocumentUploadPreviewHeight

`func (o *VariationImage) GetDocumentUploadPreviewHeight() int32`

GetDocumentUploadPreviewHeight returns the DocumentUploadPreviewHeight field if non-nil, zero value otherwise.

### GetDocumentUploadPreviewHeightOk

`func (o *VariationImage) GetDocumentUploadPreviewHeightOk() (*int32, bool)`

GetDocumentUploadPreviewHeightOk returns a tuple with the DocumentUploadPreviewHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUploadPreviewHeight

`func (o *VariationImage) SetDocumentUploadPreviewHeight(v int32)`

SetDocumentUploadPreviewHeight sets DocumentUploadPreviewHeight field to given value.

### HasDocumentUploadPreviewHeight

`func (o *VariationImage) HasDocumentUploadPreviewHeight() bool`

HasDocumentUploadPreviewHeight returns a boolean if a field has been set.

### GetHasVariationLink

`func (o *VariationImage) GetHasVariationLink() bool`

GetHasVariationLink returns the HasVariationLink field if non-nil, zero value otherwise.

### GetHasVariationLinkOk

`func (o *VariationImage) GetHasVariationLinkOk() (*bool, bool)`

GetHasVariationLinkOk returns a tuple with the HasVariationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVariationLink

`func (o *VariationImage) SetHasVariationLink(v bool)`

SetHasVariationLink sets HasVariationLink field to given value.

### HasHasVariationLink

`func (o *VariationImage) HasHasVariationLink() bool`

HasHasVariationLink returns a boolean if a field has been set.

### GetAvailabilities

`func (o *VariationImage) GetAvailabilities() []ImageAvailability`

GetAvailabilities returns the Availabilities field if non-nil, zero value otherwise.

### GetAvailabilitiesOk

`func (o *VariationImage) GetAvailabilitiesOk() (*[]ImageAvailability, bool)`

GetAvailabilitiesOk returns a tuple with the Availabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilities

`func (o *VariationImage) SetAvailabilities(v []ImageAvailability)`

SetAvailabilities sets Availabilities field to given value.

### HasAvailabilities

`func (o *VariationImage) HasAvailabilities() bool`

HasAvailabilities returns a boolean if a field has been set.

### GetNames

`func (o *VariationImage) GetNames() []ImageName`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VariationImage) GetNamesOk() (*[]ImageName, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VariationImage) SetNames(v []ImageName)`

SetNames sets Names field to given value.

### HasNames

`func (o *VariationImage) HasNames() bool`

HasNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


