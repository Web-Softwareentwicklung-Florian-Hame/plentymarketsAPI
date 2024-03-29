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

// Attribute Variation attribute model that is place in VariationAttributeValue model
type Attribute struct {
	Id *int32 `json:"id,omitempty"`
	BackendName *string `json:"backendName,omitempty"`
	Position *int32 `json:"position,omitempty"`
	IsSurchargePercental *bool `json:"isSurchargePercental,omitempty"`
	IsLinkableToImage *bool `json:"isLinkableToImage,omitempty"`
	AmazonAttribute *string `json:"amazonAttribute,omitempty"`
	FruugoAttribute *string `json:"fruugoAttribute,omitempty"`
	PixmaniaAttribute *int32 `json:"pixmaniaAttribute,omitempty"`
	OttoAttribute *string `json:"ottoAttribute,omitempty"`
	GoogleShoppingAttribute *string `json:"googleShoppingAttribute,omitempty"`
	NeckermannAtEpAttribute *int32 `json:"neckermannAtEpAttribute,omitempty"`
	TypeOfSelectionInOnlineStore *string `json:"typeOfSelectionInOnlineStore,omitempty"`
	LaRedouteAttribute *int32 `json:"laRedouteAttribute,omitempty"`
	IsGroupable *bool `json:"isGroupable,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// only accessible for get attributes call when with flag is set
	Values *[]AttributeValue `json:"values,omitempty"`
}

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute() *Attribute {
	this := Attribute{}
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Attribute) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Attribute) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Attribute) SetId(v int32) {
	o.Id = &v
}

// GetBackendName returns the BackendName field value if set, zero value otherwise.
func (o *Attribute) GetBackendName() string {
	if o == nil || o.BackendName == nil {
		var ret string
		return ret
	}
	return *o.BackendName
}

// GetBackendNameOk returns a tuple with the BackendName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetBackendNameOk() (*string, bool) {
	if o == nil || o.BackendName == nil {
		return nil, false
	}
	return o.BackendName, true
}

// HasBackendName returns a boolean if a field has been set.
func (o *Attribute) HasBackendName() bool {
	if o != nil && o.BackendName != nil {
		return true
	}

	return false
}

// SetBackendName gets a reference to the given string and assigns it to the BackendName field.
func (o *Attribute) SetBackendName(v string) {
	o.BackendName = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Attribute) GetPosition() int32 {
	if o == nil || o.Position == nil {
		var ret int32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetPositionOk() (*int32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Attribute) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int32 and assigns it to the Position field.
func (o *Attribute) SetPosition(v int32) {
	o.Position = &v
}

// GetIsSurchargePercental returns the IsSurchargePercental field value if set, zero value otherwise.
func (o *Attribute) GetIsSurchargePercental() bool {
	if o == nil || o.IsSurchargePercental == nil {
		var ret bool
		return ret
	}
	return *o.IsSurchargePercental
}

// GetIsSurchargePercentalOk returns a tuple with the IsSurchargePercental field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetIsSurchargePercentalOk() (*bool, bool) {
	if o == nil || o.IsSurchargePercental == nil {
		return nil, false
	}
	return o.IsSurchargePercental, true
}

// HasIsSurchargePercental returns a boolean if a field has been set.
func (o *Attribute) HasIsSurchargePercental() bool {
	if o != nil && o.IsSurchargePercental != nil {
		return true
	}

	return false
}

// SetIsSurchargePercental gets a reference to the given bool and assigns it to the IsSurchargePercental field.
func (o *Attribute) SetIsSurchargePercental(v bool) {
	o.IsSurchargePercental = &v
}

// GetIsLinkableToImage returns the IsLinkableToImage field value if set, zero value otherwise.
func (o *Attribute) GetIsLinkableToImage() bool {
	if o == nil || o.IsLinkableToImage == nil {
		var ret bool
		return ret
	}
	return *o.IsLinkableToImage
}

// GetIsLinkableToImageOk returns a tuple with the IsLinkableToImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetIsLinkableToImageOk() (*bool, bool) {
	if o == nil || o.IsLinkableToImage == nil {
		return nil, false
	}
	return o.IsLinkableToImage, true
}

// HasIsLinkableToImage returns a boolean if a field has been set.
func (o *Attribute) HasIsLinkableToImage() bool {
	if o != nil && o.IsLinkableToImage != nil {
		return true
	}

	return false
}

// SetIsLinkableToImage gets a reference to the given bool and assigns it to the IsLinkableToImage field.
func (o *Attribute) SetIsLinkableToImage(v bool) {
	o.IsLinkableToImage = &v
}

// GetAmazonAttribute returns the AmazonAttribute field value if set, zero value otherwise.
func (o *Attribute) GetAmazonAttribute() string {
	if o == nil || o.AmazonAttribute == nil {
		var ret string
		return ret
	}
	return *o.AmazonAttribute
}

// GetAmazonAttributeOk returns a tuple with the AmazonAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetAmazonAttributeOk() (*string, bool) {
	if o == nil || o.AmazonAttribute == nil {
		return nil, false
	}
	return o.AmazonAttribute, true
}

// HasAmazonAttribute returns a boolean if a field has been set.
func (o *Attribute) HasAmazonAttribute() bool {
	if o != nil && o.AmazonAttribute != nil {
		return true
	}

	return false
}

// SetAmazonAttribute gets a reference to the given string and assigns it to the AmazonAttribute field.
func (o *Attribute) SetAmazonAttribute(v string) {
	o.AmazonAttribute = &v
}

// GetFruugoAttribute returns the FruugoAttribute field value if set, zero value otherwise.
func (o *Attribute) GetFruugoAttribute() string {
	if o == nil || o.FruugoAttribute == nil {
		var ret string
		return ret
	}
	return *o.FruugoAttribute
}

// GetFruugoAttributeOk returns a tuple with the FruugoAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetFruugoAttributeOk() (*string, bool) {
	if o == nil || o.FruugoAttribute == nil {
		return nil, false
	}
	return o.FruugoAttribute, true
}

// HasFruugoAttribute returns a boolean if a field has been set.
func (o *Attribute) HasFruugoAttribute() bool {
	if o != nil && o.FruugoAttribute != nil {
		return true
	}

	return false
}

// SetFruugoAttribute gets a reference to the given string and assigns it to the FruugoAttribute field.
func (o *Attribute) SetFruugoAttribute(v string) {
	o.FruugoAttribute = &v
}

// GetPixmaniaAttribute returns the PixmaniaAttribute field value if set, zero value otherwise.
func (o *Attribute) GetPixmaniaAttribute() int32 {
	if o == nil || o.PixmaniaAttribute == nil {
		var ret int32
		return ret
	}
	return *o.PixmaniaAttribute
}

// GetPixmaniaAttributeOk returns a tuple with the PixmaniaAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetPixmaniaAttributeOk() (*int32, bool) {
	if o == nil || o.PixmaniaAttribute == nil {
		return nil, false
	}
	return o.PixmaniaAttribute, true
}

// HasPixmaniaAttribute returns a boolean if a field has been set.
func (o *Attribute) HasPixmaniaAttribute() bool {
	if o != nil && o.PixmaniaAttribute != nil {
		return true
	}

	return false
}

// SetPixmaniaAttribute gets a reference to the given int32 and assigns it to the PixmaniaAttribute field.
func (o *Attribute) SetPixmaniaAttribute(v int32) {
	o.PixmaniaAttribute = &v
}

// GetOttoAttribute returns the OttoAttribute field value if set, zero value otherwise.
func (o *Attribute) GetOttoAttribute() string {
	if o == nil || o.OttoAttribute == nil {
		var ret string
		return ret
	}
	return *o.OttoAttribute
}

// GetOttoAttributeOk returns a tuple with the OttoAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetOttoAttributeOk() (*string, bool) {
	if o == nil || o.OttoAttribute == nil {
		return nil, false
	}
	return o.OttoAttribute, true
}

// HasOttoAttribute returns a boolean if a field has been set.
func (o *Attribute) HasOttoAttribute() bool {
	if o != nil && o.OttoAttribute != nil {
		return true
	}

	return false
}

// SetOttoAttribute gets a reference to the given string and assigns it to the OttoAttribute field.
func (o *Attribute) SetOttoAttribute(v string) {
	o.OttoAttribute = &v
}

// GetGoogleShoppingAttribute returns the GoogleShoppingAttribute field value if set, zero value otherwise.
func (o *Attribute) GetGoogleShoppingAttribute() string {
	if o == nil || o.GoogleShoppingAttribute == nil {
		var ret string
		return ret
	}
	return *o.GoogleShoppingAttribute
}

// GetGoogleShoppingAttributeOk returns a tuple with the GoogleShoppingAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetGoogleShoppingAttributeOk() (*string, bool) {
	if o == nil || o.GoogleShoppingAttribute == nil {
		return nil, false
	}
	return o.GoogleShoppingAttribute, true
}

// HasGoogleShoppingAttribute returns a boolean if a field has been set.
func (o *Attribute) HasGoogleShoppingAttribute() bool {
	if o != nil && o.GoogleShoppingAttribute != nil {
		return true
	}

	return false
}

// SetGoogleShoppingAttribute gets a reference to the given string and assigns it to the GoogleShoppingAttribute field.
func (o *Attribute) SetGoogleShoppingAttribute(v string) {
	o.GoogleShoppingAttribute = &v
}

// GetNeckermannAtEpAttribute returns the NeckermannAtEpAttribute field value if set, zero value otherwise.
func (o *Attribute) GetNeckermannAtEpAttribute() int32 {
	if o == nil || o.NeckermannAtEpAttribute == nil {
		var ret int32
		return ret
	}
	return *o.NeckermannAtEpAttribute
}

// GetNeckermannAtEpAttributeOk returns a tuple with the NeckermannAtEpAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetNeckermannAtEpAttributeOk() (*int32, bool) {
	if o == nil || o.NeckermannAtEpAttribute == nil {
		return nil, false
	}
	return o.NeckermannAtEpAttribute, true
}

// HasNeckermannAtEpAttribute returns a boolean if a field has been set.
func (o *Attribute) HasNeckermannAtEpAttribute() bool {
	if o != nil && o.NeckermannAtEpAttribute != nil {
		return true
	}

	return false
}

// SetNeckermannAtEpAttribute gets a reference to the given int32 and assigns it to the NeckermannAtEpAttribute field.
func (o *Attribute) SetNeckermannAtEpAttribute(v int32) {
	o.NeckermannAtEpAttribute = &v
}

// GetTypeOfSelectionInOnlineStore returns the TypeOfSelectionInOnlineStore field value if set, zero value otherwise.
func (o *Attribute) GetTypeOfSelectionInOnlineStore() string {
	if o == nil || o.TypeOfSelectionInOnlineStore == nil {
		var ret string
		return ret
	}
	return *o.TypeOfSelectionInOnlineStore
}

// GetTypeOfSelectionInOnlineStoreOk returns a tuple with the TypeOfSelectionInOnlineStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetTypeOfSelectionInOnlineStoreOk() (*string, bool) {
	if o == nil || o.TypeOfSelectionInOnlineStore == nil {
		return nil, false
	}
	return o.TypeOfSelectionInOnlineStore, true
}

// HasTypeOfSelectionInOnlineStore returns a boolean if a field has been set.
func (o *Attribute) HasTypeOfSelectionInOnlineStore() bool {
	if o != nil && o.TypeOfSelectionInOnlineStore != nil {
		return true
	}

	return false
}

// SetTypeOfSelectionInOnlineStore gets a reference to the given string and assigns it to the TypeOfSelectionInOnlineStore field.
func (o *Attribute) SetTypeOfSelectionInOnlineStore(v string) {
	o.TypeOfSelectionInOnlineStore = &v
}

// GetLaRedouteAttribute returns the LaRedouteAttribute field value if set, zero value otherwise.
func (o *Attribute) GetLaRedouteAttribute() int32 {
	if o == nil || o.LaRedouteAttribute == nil {
		var ret int32
		return ret
	}
	return *o.LaRedouteAttribute
}

// GetLaRedouteAttributeOk returns a tuple with the LaRedouteAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetLaRedouteAttributeOk() (*int32, bool) {
	if o == nil || o.LaRedouteAttribute == nil {
		return nil, false
	}
	return o.LaRedouteAttribute, true
}

// HasLaRedouteAttribute returns a boolean if a field has been set.
func (o *Attribute) HasLaRedouteAttribute() bool {
	if o != nil && o.LaRedouteAttribute != nil {
		return true
	}

	return false
}

// SetLaRedouteAttribute gets a reference to the given int32 and assigns it to the LaRedouteAttribute field.
func (o *Attribute) SetLaRedouteAttribute(v int32) {
	o.LaRedouteAttribute = &v
}

// GetIsGroupable returns the IsGroupable field value if set, zero value otherwise.
func (o *Attribute) GetIsGroupable() bool {
	if o == nil || o.IsGroupable == nil {
		var ret bool
		return ret
	}
	return *o.IsGroupable
}

// GetIsGroupableOk returns a tuple with the IsGroupable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetIsGroupableOk() (*bool, bool) {
	if o == nil || o.IsGroupable == nil {
		return nil, false
	}
	return o.IsGroupable, true
}

// HasIsGroupable returns a boolean if a field has been set.
func (o *Attribute) HasIsGroupable() bool {
	if o != nil && o.IsGroupable != nil {
		return true
	}

	return false
}

// SetIsGroupable gets a reference to the given bool and assigns it to the IsGroupable field.
func (o *Attribute) SetIsGroupable(v bool) {
	o.IsGroupable = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Attribute) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Attribute) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Attribute) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Attribute) GetValues() []AttributeValue {
	if o == nil || o.Values == nil {
		var ret []AttributeValue
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Attribute) GetValuesOk() (*[]AttributeValue, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Attribute) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []AttributeValue and assigns it to the Values field.
func (o *Attribute) SetValues(v []AttributeValue) {
	o.Values = &v
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.BackendName != nil {
		toSerialize["backendName"] = o.BackendName
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if o.IsSurchargePercental != nil {
		toSerialize["isSurchargePercental"] = o.IsSurchargePercental
	}
	if o.IsLinkableToImage != nil {
		toSerialize["isLinkableToImage"] = o.IsLinkableToImage
	}
	if o.AmazonAttribute != nil {
		toSerialize["amazonAttribute"] = o.AmazonAttribute
	}
	if o.FruugoAttribute != nil {
		toSerialize["fruugoAttribute"] = o.FruugoAttribute
	}
	if o.PixmaniaAttribute != nil {
		toSerialize["pixmaniaAttribute"] = o.PixmaniaAttribute
	}
	if o.OttoAttribute != nil {
		toSerialize["ottoAttribute"] = o.OttoAttribute
	}
	if o.GoogleShoppingAttribute != nil {
		toSerialize["googleShoppingAttribute"] = o.GoogleShoppingAttribute
	}
	if o.NeckermannAtEpAttribute != nil {
		toSerialize["neckermannAtEpAttribute"] = o.NeckermannAtEpAttribute
	}
	if o.TypeOfSelectionInOnlineStore != nil {
		toSerialize["typeOfSelectionInOnlineStore"] = o.TypeOfSelectionInOnlineStore
	}
	if o.LaRedouteAttribute != nil {
		toSerialize["laRedouteAttribute"] = o.LaRedouteAttribute
	}
	if o.IsGroupable != nil {
		toSerialize["isGroupable"] = o.IsGroupable
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


