/*
Plentymarkets-API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// VariationImage variation image model
type VariationImage struct {
	ImageId *int32 `json:"imageId,omitempty"`
	Id *int32 `json:"id,omitempty"`
	ItemId *int32 `json:"itemId,omitempty"`
	Type *string `json:"type,omitempty"`
	FileType *string `json:"fileType,omitempty"`
	Path *string `json:"path,omitempty"`
	Position *int32 `json:"position,omitempty"`
	LastUpdate *string `json:"lastUpdate,omitempty"`
	Insert *string `json:"insert,omitempty"`
	Md5Checksum *string `json:"md5Checksum,omitempty"`
	Width *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
	Size *int32 `json:"size,omitempty"`
	StorageProviderId *string `json:"storageProviderId,omitempty"`
	Md5ChecksumOriginal *string `json:"md5ChecksumOriginal,omitempty"`
	CleanImageName *string `json:"cleanImageName,omitempty"`
	UploadUrl *string `json:"uploadUrl,omitempty"`
	Url *string `json:"url,omitempty"`
	UrlMiddle *string `json:"urlMiddle,omitempty"`
	UrlPreview *string `json:"urlPreview,omitempty"`
	UrlSecondPreview *string `json:"urlSecondPreview,omitempty"`
	DocumentUploadPath *string `json:"documentUploadPath,omitempty"`
	DocumentUploadPathPreview *string `json:"documentUploadPathPreview,omitempty"`
	DocumentUploadPreviewWidth *int32 `json:"documentUploadPreviewWidth,omitempty"`
	DocumentUploadPreviewHeight *int32 `json:"documentUploadPreviewHeight,omitempty"`
	HasVariationLink *bool `json:"hasVariationLink,omitempty"`
	Availabilities *[]ImageAvailability `json:"availabilities,omitempty"`
	Names *[]ImageName `json:"names,omitempty"`
	AttributeValueImages *[]ImageMarketAttributeValue `json:"attributeValueImages,omitempty"`
}

// NewVariationImage instantiates a new VariationImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariationImage() *VariationImage {
	this := VariationImage{}
	return &this
}

// NewVariationImageWithDefaults instantiates a new VariationImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariationImageWithDefaults() *VariationImage {
	this := VariationImage{}
	return &this
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *VariationImage) GetImageId() int32 {
	if o == nil || o.ImageId == nil {
		var ret int32
		return ret
	}
	return *o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetImageIdOk() (*int32, bool) {
	if o == nil || o.ImageId == nil {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *VariationImage) HasImageId() bool {
	if o != nil && o.ImageId != nil {
		return true
	}

	return false
}

// SetImageId gets a reference to the given int32 and assigns it to the ImageId field.
func (o *VariationImage) SetImageId(v int32) {
	o.ImageId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariationImage) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariationImage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VariationImage) SetId(v int32) {
	o.Id = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *VariationImage) GetItemId() int32 {
	if o == nil || o.ItemId == nil {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetItemIdOk() (*int32, bool) {
	if o == nil || o.ItemId == nil {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *VariationImage) HasItemId() bool {
	if o != nil && o.ItemId != nil {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *VariationImage) SetItemId(v int32) {
	o.ItemId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VariationImage) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VariationImage) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VariationImage) SetType(v string) {
	o.Type = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *VariationImage) GetFileType() string {
	if o == nil || o.FileType == nil {
		var ret string
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetFileTypeOk() (*string, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *VariationImage) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given string and assigns it to the FileType field.
func (o *VariationImage) SetFileType(v string) {
	o.FileType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VariationImage) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VariationImage) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VariationImage) SetPath(v string) {
	o.Path = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *VariationImage) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *VariationImage) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *VariationImage) SetPosition(v int32) {
	o.Position = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *VariationImage) GetLastUpdate() string {
	if o == nil || o.LastUpdate == nil {
		var ret string
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetLastUpdateOk() (*string, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *VariationImage) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given string and assigns it to the LastUpdate field.
func (o *VariationImage) SetLastUpdate(v string) {
	o.LastUpdate = &v
}

// GetInsert returns the Insert field value if set, zero value otherwise.
func (o *VariationImage) GetInsert() string {
	if o == nil || o.Insert == nil {
		var ret string
		return ret
	}
	return *o.Insert
}

// GetInsertOk returns a tuple with the Insert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetInsertOk() (*string, bool) {
	if o == nil || o.Insert == nil {
		return nil, false
	}
	return o.Insert, true
}

// HasInsert returns a boolean if a field has been set.
func (o *VariationImage) HasInsert() bool {
	if o != nil && o.Insert != nil {
		return true
	}

	return false
}

// SetInsert gets a reference to the given string and assigns it to the Insert field.
func (o *VariationImage) SetInsert(v string) {
	o.Insert = &v
}

// GetMd5Checksum returns the Md5Checksum field value if set, zero value otherwise.
func (o *VariationImage) GetMd5Checksum() string {
	if o == nil || o.Md5Checksum == nil {
		var ret string
		return ret
	}
	return *o.Md5Checksum
}

// GetMd5ChecksumOk returns a tuple with the Md5Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetMd5ChecksumOk() (*string, bool) {
	if o == nil || o.Md5Checksum == nil {
		return nil, false
	}
	return o.Md5Checksum, true
}

// HasMd5Checksum returns a boolean if a field has been set.
func (o *VariationImage) HasMd5Checksum() bool {
	if o != nil && o.Md5Checksum != nil {
		return true
	}

	return false
}

// SetMd5Checksum gets a reference to the given string and assigns it to the Md5Checksum field.
func (o *VariationImage) SetMd5Checksum(v string) {
	o.Md5Checksum = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VariationImage) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VariationImage) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *VariationImage) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VariationImage) GetHeight() int32 {
	if o == nil || o.Height == nil {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetHeightOk() (*int32, bool) {
	if o == nil || o.Height == nil {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VariationImage) HasHeight() bool {
	if o != nil && o.Height != nil {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *VariationImage) SetHeight(v int32) {
	o.Height = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *VariationImage) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *VariationImage) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *VariationImage) SetSize(v int32) {
	o.Size = &v
}

// GetStorageProviderId returns the StorageProviderId field value if set, zero value otherwise.
func (o *VariationImage) GetStorageProviderId() string {
	if o == nil || o.StorageProviderId == nil {
		var ret string
		return ret
	}
	return *o.StorageProviderId
}

// GetStorageProviderIdOk returns a tuple with the StorageProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetStorageProviderIdOk() (*string, bool) {
	if o == nil || o.StorageProviderId == nil {
		return nil, false
	}
	return o.StorageProviderId, true
}

// HasStorageProviderId returns a boolean if a field has been set.
func (o *VariationImage) HasStorageProviderId() bool {
	if o != nil && o.StorageProviderId != nil {
		return true
	}

	return false
}

// SetStorageProviderId gets a reference to the given string and assigns it to the StorageProviderId field.
func (o *VariationImage) SetStorageProviderId(v string) {
	o.StorageProviderId = &v
}

// GetMd5ChecksumOriginal returns the Md5ChecksumOriginal field value if set, zero value otherwise.
func (o *VariationImage) GetMd5ChecksumOriginal() string {
	if o == nil || o.Md5ChecksumOriginal == nil {
		var ret string
		return ret
	}
	return *o.Md5ChecksumOriginal
}

// GetMd5ChecksumOriginalOk returns a tuple with the Md5ChecksumOriginal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetMd5ChecksumOriginalOk() (*string, bool) {
	if o == nil || o.Md5ChecksumOriginal == nil {
		return nil, false
	}
	return o.Md5ChecksumOriginal, true
}

// HasMd5ChecksumOriginal returns a boolean if a field has been set.
func (o *VariationImage) HasMd5ChecksumOriginal() bool {
	if o != nil && o.Md5ChecksumOriginal != nil {
		return true
	}

	return false
}

// SetMd5ChecksumOriginal gets a reference to the given string and assigns it to the Md5ChecksumOriginal field.
func (o *VariationImage) SetMd5ChecksumOriginal(v string) {
	o.Md5ChecksumOriginal = &v
}

// GetCleanImageName returns the CleanImageName field value if set, zero value otherwise.
func (o *VariationImage) GetCleanImageName() string {
	if o == nil || o.CleanImageName == nil {
		var ret string
		return ret
	}
	return *o.CleanImageName
}

// GetCleanImageNameOk returns a tuple with the CleanImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetCleanImageNameOk() (*string, bool) {
	if o == nil || o.CleanImageName == nil {
		return nil, false
	}
	return o.CleanImageName, true
}

// HasCleanImageName returns a boolean if a field has been set.
func (o *VariationImage) HasCleanImageName() bool {
	if o != nil && o.CleanImageName != nil {
		return true
	}

	return false
}

// SetCleanImageName gets a reference to the given string and assigns it to the CleanImageName field.
func (o *VariationImage) SetCleanImageName(v string) {
	o.CleanImageName = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *VariationImage) GetUploadUrl() string {
	if o == nil || o.UploadUrl == nil {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetUploadUrlOk() (*string, bool) {
	if o == nil || o.UploadUrl == nil {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *VariationImage) HasUploadUrl() bool {
	if o != nil && o.UploadUrl != nil {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *VariationImage) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *VariationImage) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *VariationImage) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *VariationImage) SetUrl(v string) {
	o.Url = &v
}

// GetUrlMiddle returns the UrlMiddle field value if set, zero value otherwise.
func (o *VariationImage) GetUrlMiddle() string {
	if o == nil || o.UrlMiddle == nil {
		var ret string
		return ret
	}
	return *o.UrlMiddle
}

// GetUrlMiddleOk returns a tuple with the UrlMiddle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetUrlMiddleOk() (*string, bool) {
	if o == nil || o.UrlMiddle == nil {
		return nil, false
	}
	return o.UrlMiddle, true
}

// HasUrlMiddle returns a boolean if a field has been set.
func (o *VariationImage) HasUrlMiddle() bool {
	if o != nil && o.UrlMiddle != nil {
		return true
	}

	return false
}

// SetUrlMiddle gets a reference to the given string and assigns it to the UrlMiddle field.
func (o *VariationImage) SetUrlMiddle(v string) {
	o.UrlMiddle = &v
}

// GetUrlPreview returns the UrlPreview field value if set, zero value otherwise.
func (o *VariationImage) GetUrlPreview() string {
	if o == nil || o.UrlPreview == nil {
		var ret string
		return ret
	}
	return *o.UrlPreview
}

// GetUrlPreviewOk returns a tuple with the UrlPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetUrlPreviewOk() (*string, bool) {
	if o == nil || o.UrlPreview == nil {
		return nil, false
	}
	return o.UrlPreview, true
}

// HasUrlPreview returns a boolean if a field has been set.
func (o *VariationImage) HasUrlPreview() bool {
	if o != nil && o.UrlPreview != nil {
		return true
	}

	return false
}

// SetUrlPreview gets a reference to the given string and assigns it to the UrlPreview field.
func (o *VariationImage) SetUrlPreview(v string) {
	o.UrlPreview = &v
}

// GetUrlSecondPreview returns the UrlSecondPreview field value if set, zero value otherwise.
func (o *VariationImage) GetUrlSecondPreview() string {
	if o == nil || o.UrlSecondPreview == nil {
		var ret string
		return ret
	}
	return *o.UrlSecondPreview
}

// GetUrlSecondPreviewOk returns a tuple with the UrlSecondPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetUrlSecondPreviewOk() (*string, bool) {
	if o == nil || o.UrlSecondPreview == nil {
		return nil, false
	}
	return o.UrlSecondPreview, true
}

// HasUrlSecondPreview returns a boolean if a field has been set.
func (o *VariationImage) HasUrlSecondPreview() bool {
	if o != nil && o.UrlSecondPreview != nil {
		return true
	}

	return false
}

// SetUrlSecondPreview gets a reference to the given string and assigns it to the UrlSecondPreview field.
func (o *VariationImage) SetUrlSecondPreview(v string) {
	o.UrlSecondPreview = &v
}

// GetDocumentUploadPath returns the DocumentUploadPath field value if set, zero value otherwise.
func (o *VariationImage) GetDocumentUploadPath() string {
	if o == nil || o.DocumentUploadPath == nil {
		var ret string
		return ret
	}
	return *o.DocumentUploadPath
}

// GetDocumentUploadPathOk returns a tuple with the DocumentUploadPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetDocumentUploadPathOk() (*string, bool) {
	if o == nil || o.DocumentUploadPath == nil {
		return nil, false
	}
	return o.DocumentUploadPath, true
}

// HasDocumentUploadPath returns a boolean if a field has been set.
func (o *VariationImage) HasDocumentUploadPath() bool {
	if o != nil && o.DocumentUploadPath != nil {
		return true
	}

	return false
}

// SetDocumentUploadPath gets a reference to the given string and assigns it to the DocumentUploadPath field.
func (o *VariationImage) SetDocumentUploadPath(v string) {
	o.DocumentUploadPath = &v
}

// GetDocumentUploadPathPreview returns the DocumentUploadPathPreview field value if set, zero value otherwise.
func (o *VariationImage) GetDocumentUploadPathPreview() string {
	if o == nil || o.DocumentUploadPathPreview == nil {
		var ret string
		return ret
	}
	return *o.DocumentUploadPathPreview
}

// GetDocumentUploadPathPreviewOk returns a tuple with the DocumentUploadPathPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetDocumentUploadPathPreviewOk() (*string, bool) {
	if o == nil || o.DocumentUploadPathPreview == nil {
		return nil, false
	}
	return o.DocumentUploadPathPreview, true
}

// HasDocumentUploadPathPreview returns a boolean if a field has been set.
func (o *VariationImage) HasDocumentUploadPathPreview() bool {
	if o != nil && o.DocumentUploadPathPreview != nil {
		return true
	}

	return false
}

// SetDocumentUploadPathPreview gets a reference to the given string and assigns it to the DocumentUploadPathPreview field.
func (o *VariationImage) SetDocumentUploadPathPreview(v string) {
	o.DocumentUploadPathPreview = &v
}

// GetDocumentUploadPreviewWidth returns the DocumentUploadPreviewWidth field value if set, zero value otherwise.
func (o *VariationImage) GetDocumentUploadPreviewWidth() int32 {
	if o == nil || o.DocumentUploadPreviewWidth == nil {
		var ret int32
		return ret
	}
	return *o.DocumentUploadPreviewWidth
}

// GetDocumentUploadPreviewWidthOk returns a tuple with the DocumentUploadPreviewWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetDocumentUploadPreviewWidthOk() (*int32, bool) {
	if o == nil || o.DocumentUploadPreviewWidth == nil {
		return nil, false
	}
	return o.DocumentUploadPreviewWidth, true
}

// HasDocumentUploadPreviewWidth returns a boolean if a field has been set.
func (o *VariationImage) HasDocumentUploadPreviewWidth() bool {
	if o != nil && o.DocumentUploadPreviewWidth != nil {
		return true
	}

	return false
}

// SetDocumentUploadPreviewWidth gets a reference to the given int32 and assigns it to the DocumentUploadPreviewWidth field.
func (o *VariationImage) SetDocumentUploadPreviewWidth(v int32) {
	o.DocumentUploadPreviewWidth = &v
}

// GetDocumentUploadPreviewHeight returns the DocumentUploadPreviewHeight field value if set, zero value otherwise.
func (o *VariationImage) GetDocumentUploadPreviewHeight() int32 {
	if o == nil || o.DocumentUploadPreviewHeight == nil {
		var ret int32
		return ret
	}
	return *o.DocumentUploadPreviewHeight
}

// GetDocumentUploadPreviewHeightOk returns a tuple with the DocumentUploadPreviewHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetDocumentUploadPreviewHeightOk() (*int32, bool) {
	if o == nil || o.DocumentUploadPreviewHeight == nil {
		return nil, false
	}
	return o.DocumentUploadPreviewHeight, true
}

// HasDocumentUploadPreviewHeight returns a boolean if a field has been set.
func (o *VariationImage) HasDocumentUploadPreviewHeight() bool {
	if o != nil && o.DocumentUploadPreviewHeight != nil {
		return true
	}

	return false
}

// SetDocumentUploadPreviewHeight gets a reference to the given int32 and assigns it to the DocumentUploadPreviewHeight field.
func (o *VariationImage) SetDocumentUploadPreviewHeight(v int32) {
	o.DocumentUploadPreviewHeight = &v
}

// GetHasVariationLink returns the HasVariationLink field value if set, zero value otherwise.
func (o *VariationImage) GetHasVariationLink() bool {
	if o == nil || o.HasVariationLink == nil {
		var ret bool
		return ret
	}
	return *o.HasVariationLink
}

// GetHasVariationLinkOk returns a tuple with the HasVariationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetHasVariationLinkOk() (*bool, bool) {
	if o == nil || o.HasVariationLink == nil {
		return nil, false
	}
	return o.HasVariationLink, true
}

// HasHasVariationLink returns a boolean if a field has been set.
func (o *VariationImage) HasHasVariationLink() bool {
	if o != nil && o.HasVariationLink != nil {
		return true
	}

	return false
}

// SetHasVariationLink gets a reference to the given bool and assigns it to the HasVariationLink field.
func (o *VariationImage) SetHasVariationLink(v bool) {
	o.HasVariationLink = &v
}

// GetAvailabilities returns the Availabilities field value if set, zero value otherwise.
func (o *VariationImage) GetAvailabilities() []ImageAvailability {
	if o == nil || o.Availabilities == nil {
		var ret []ImageAvailability
		return ret
	}
	return *o.Availabilities
}

// GetAvailabilitiesOk returns a tuple with the Availabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetAvailabilitiesOk() (*[]ImageAvailability, bool) {
	if o == nil || o.Availabilities == nil {
		return nil, false
	}
	return o.Availabilities, true
}

// HasAvailabilities returns a boolean if a field has been set.
func (o *VariationImage) HasAvailabilities() bool {
	if o != nil && o.Availabilities != nil {
		return true
	}

	return false
}

// SetAvailabilities gets a reference to the given []ImageAvailability and assigns it to the Availabilities field.
func (o *VariationImage) SetAvailabilities(v []ImageAvailability) {
	o.Availabilities = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *VariationImage) GetNames() []ImageName {
	if o == nil || o.Names == nil {
		var ret []ImageName
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetNamesOk() (*[]ImageName, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *VariationImage) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []ImageName and assigns it to the Names field.
func (o *VariationImage) SetNames(v []ImageName) {
	o.Names = &v
}

// GetAttributeValueImages returns the AttributeValueImages field value if set, zero value otherwise.
func (o *VariationImage) GetAttributeValueImages() []ImageMarketAttributeValue {
	if o == nil || o.AttributeValueImages == nil {
		var ret []ImageMarketAttributeValue
		return ret
	}
	return *o.AttributeValueImages
}

// GetAttributeValueImagesOk returns a tuple with the AttributeValueImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariationImage) GetAttributeValueImagesOk() (*[]ImageMarketAttributeValue, bool) {
	if o == nil || o.AttributeValueImages == nil {
		return nil, false
	}
	return o.AttributeValueImages, true
}

// HasAttributeValueImages returns a boolean if a field has been set.
func (o *VariationImage) HasAttributeValueImages() bool {
	if o != nil && o.AttributeValueImages != nil {
		return true
	}

	return false
}

// SetAttributeValueImages gets a reference to the given []ImageMarketAttributeValue and assigns it to the AttributeValueImages field.
func (o *VariationImage) SetAttributeValueImages(v []ImageMarketAttributeValue) {
	o.AttributeValueImages = &v
}

func (o VariationImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImageId != nil {
		toSerialize["imageId"] = o.ImageId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ItemId != nil {
		toSerialize["itemId"] = o.ItemId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.FileType != nil {
		toSerialize["fileType"] = o.FileType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.LastUpdate != nil {
		toSerialize["lastUpdate"] = o.LastUpdate
	}
	if o.Insert != nil {
		toSerialize["insert"] = o.Insert
	}
	if o.Md5Checksum != nil {
		toSerialize["md5Checksum"] = o.Md5Checksum
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Height != nil {
		toSerialize["height"] = o.Height
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.StorageProviderId != nil {
		toSerialize["storageProviderId"] = o.StorageProviderId
	}
	if o.Md5ChecksumOriginal != nil {
		toSerialize["md5ChecksumOriginal"] = o.Md5ChecksumOriginal
	}
	if o.CleanImageName != nil {
		toSerialize["cleanImageName"] = o.CleanImageName
	}
	if o.UploadUrl != nil {
		toSerialize["uploadUrl"] = o.UploadUrl
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UrlMiddle != nil {
		toSerialize["urlMiddle"] = o.UrlMiddle
	}
	if o.UrlPreview != nil {
		toSerialize["urlPreview"] = o.UrlPreview
	}
	if o.UrlSecondPreview != nil {
		toSerialize["urlSecondPreview"] = o.UrlSecondPreview
	}
	if o.DocumentUploadPath != nil {
		toSerialize["documentUploadPath"] = o.DocumentUploadPath
	}
	if o.DocumentUploadPathPreview != nil {
		toSerialize["documentUploadPathPreview"] = o.DocumentUploadPathPreview
	}
	if o.DocumentUploadPreviewWidth != nil {
		toSerialize["documentUploadPreviewWidth"] = o.DocumentUploadPreviewWidth
	}
	if o.DocumentUploadPreviewHeight != nil {
		toSerialize["documentUploadPreviewHeight"] = o.DocumentUploadPreviewHeight
	}
	if o.HasVariationLink != nil {
		toSerialize["hasVariationLink"] = o.HasVariationLink
	}
	if o.Availabilities != nil {
		toSerialize["availabilities"] = o.Availabilities
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.AttributeValueImages != nil {
		toSerialize["attributeValueImages"] = o.AttributeValueImages
	}
	return json.Marshal(toSerialize)
}

type NullableVariationImage struct {
	value *VariationImage
	isSet bool
}

func (v NullableVariationImage) Get() *VariationImage {
	return v.value
}

func (v *NullableVariationImage) Set(val *VariationImage) {
	v.value = val
	v.isSet = true
}

func (v NullableVariationImage) IsSet() bool {
	return v.isSet
}

func (v *NullableVariationImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariationImage(val *VariationImage) *NullableVariationImage {
	return &NullableVariationImage{value: val, isSet: true}
}

func (v NullableVariationImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariationImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


