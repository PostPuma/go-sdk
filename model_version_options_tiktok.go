/*
PostPuma - OpenAPI 3.0

PostPuma API specifications

API version: 1.0.0
Contact: support@postpuma.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package PostPuma

import (
	"encoding/json"
)

// checks if the VersionOptionsTiktok type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptionsTiktok{}

// VersionOptionsTiktok struct for VersionOptionsTiktok
type VersionOptionsTiktok struct {
	PrivacyLevel *VersionOptionsTiktokPrivacyLevel `json:"privacy_level,omitempty"`
	AllowComments *VersionOptionsTiktokPrivacyLevel `json:"allow_comments,omitempty"`
	AllowDuet *VersionOptionsTiktokPrivacyLevel `json:"allow_duet,omitempty"`
	AllowStitch *VersionOptionsTiktokPrivacyLevel `json:"allow_stitch,omitempty"`
	ContentDisclosure *VersionOptionsTiktokPrivacyLevel `json:"content_disclosure,omitempty"`
	BrandOrganicToggle *VersionOptionsTiktokPrivacyLevel `json:"brand_organic_toggle,omitempty"`
	BrandContentToggle *VersionOptionsTiktokPrivacyLevel `json:"brand_content_toggle,omitempty"`
}

// NewVersionOptionsTiktok instantiates a new VersionOptionsTiktok object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptionsTiktok() *VersionOptionsTiktok {
	this := VersionOptionsTiktok{}
	return &this
}

// NewVersionOptionsTiktokWithDefaults instantiates a new VersionOptionsTiktok object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsTiktokWithDefaults() *VersionOptionsTiktok {
	this := VersionOptionsTiktok{}
	return &this
}

// GetPrivacyLevel returns the PrivacyLevel field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetPrivacyLevel() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.PrivacyLevel) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.PrivacyLevel
}

// GetPrivacyLevelOk returns a tuple with the PrivacyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetPrivacyLevelOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.PrivacyLevel) {
		return nil, false
	}
	return o.PrivacyLevel, true
}

// HasPrivacyLevel returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasPrivacyLevel() bool {
	if o != nil && !IsNil(o.PrivacyLevel) {
		return true
	}

	return false
}

// SetPrivacyLevel gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the PrivacyLevel field.
func (o *VersionOptionsTiktok) SetPrivacyLevel(v VersionOptionsTiktokPrivacyLevel) {
	o.PrivacyLevel = &v
}

// GetAllowComments returns the AllowComments field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetAllowComments() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.AllowComments) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.AllowComments
}

// GetAllowCommentsOk returns a tuple with the AllowComments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetAllowCommentsOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.AllowComments) {
		return nil, false
	}
	return o.AllowComments, true
}

// HasAllowComments returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasAllowComments() bool {
	if o != nil && !IsNil(o.AllowComments) {
		return true
	}

	return false
}

// SetAllowComments gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the AllowComments field.
func (o *VersionOptionsTiktok) SetAllowComments(v VersionOptionsTiktokPrivacyLevel) {
	o.AllowComments = &v
}

// GetAllowDuet returns the AllowDuet field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetAllowDuet() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.AllowDuet) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.AllowDuet
}

// GetAllowDuetOk returns a tuple with the AllowDuet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetAllowDuetOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.AllowDuet) {
		return nil, false
	}
	return o.AllowDuet, true
}

// HasAllowDuet returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasAllowDuet() bool {
	if o != nil && !IsNil(o.AllowDuet) {
		return true
	}

	return false
}

// SetAllowDuet gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the AllowDuet field.
func (o *VersionOptionsTiktok) SetAllowDuet(v VersionOptionsTiktokPrivacyLevel) {
	o.AllowDuet = &v
}

// GetAllowStitch returns the AllowStitch field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetAllowStitch() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.AllowStitch) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.AllowStitch
}

// GetAllowStitchOk returns a tuple with the AllowStitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetAllowStitchOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.AllowStitch) {
		return nil, false
	}
	return o.AllowStitch, true
}

// HasAllowStitch returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasAllowStitch() bool {
	if o != nil && !IsNil(o.AllowStitch) {
		return true
	}

	return false
}

// SetAllowStitch gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the AllowStitch field.
func (o *VersionOptionsTiktok) SetAllowStitch(v VersionOptionsTiktokPrivacyLevel) {
	o.AllowStitch = &v
}

// GetContentDisclosure returns the ContentDisclosure field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetContentDisclosure() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.ContentDisclosure) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.ContentDisclosure
}

// GetContentDisclosureOk returns a tuple with the ContentDisclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetContentDisclosureOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.ContentDisclosure) {
		return nil, false
	}
	return o.ContentDisclosure, true
}

// HasContentDisclosure returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasContentDisclosure() bool {
	if o != nil && !IsNil(o.ContentDisclosure) {
		return true
	}

	return false
}

// SetContentDisclosure gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the ContentDisclosure field.
func (o *VersionOptionsTiktok) SetContentDisclosure(v VersionOptionsTiktokPrivacyLevel) {
	o.ContentDisclosure = &v
}

// GetBrandOrganicToggle returns the BrandOrganicToggle field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetBrandOrganicToggle() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.BrandOrganicToggle) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.BrandOrganicToggle
}

// GetBrandOrganicToggleOk returns a tuple with the BrandOrganicToggle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetBrandOrganicToggleOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.BrandOrganicToggle) {
		return nil, false
	}
	return o.BrandOrganicToggle, true
}

// HasBrandOrganicToggle returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasBrandOrganicToggle() bool {
	if o != nil && !IsNil(o.BrandOrganicToggle) {
		return true
	}

	return false
}

// SetBrandOrganicToggle gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the BrandOrganicToggle field.
func (o *VersionOptionsTiktok) SetBrandOrganicToggle(v VersionOptionsTiktokPrivacyLevel) {
	o.BrandOrganicToggle = &v
}

// GetBrandContentToggle returns the BrandContentToggle field value if set, zero value otherwise.
func (o *VersionOptionsTiktok) GetBrandContentToggle() VersionOptionsTiktokPrivacyLevel {
	if o == nil || IsNil(o.BrandContentToggle) {
		var ret VersionOptionsTiktokPrivacyLevel
		return ret
	}
	return *o.BrandContentToggle
}

// GetBrandContentToggleOk returns a tuple with the BrandContentToggle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktok) GetBrandContentToggleOk() (*VersionOptionsTiktokPrivacyLevel, bool) {
	if o == nil || IsNil(o.BrandContentToggle) {
		return nil, false
	}
	return o.BrandContentToggle, true
}

// HasBrandContentToggle returns a boolean if a field has been set.
func (o *VersionOptionsTiktok) HasBrandContentToggle() bool {
	if o != nil && !IsNil(o.BrandContentToggle) {
		return true
	}

	return false
}

// SetBrandContentToggle gets a reference to the given VersionOptionsTiktokPrivacyLevel and assigns it to the BrandContentToggle field.
func (o *VersionOptionsTiktok) SetBrandContentToggle(v VersionOptionsTiktokPrivacyLevel) {
	o.BrandContentToggle = &v
}

func (o VersionOptionsTiktok) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptionsTiktok) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrivacyLevel) {
		toSerialize["privacy_level"] = o.PrivacyLevel
	}
	if !IsNil(o.AllowComments) {
		toSerialize["allow_comments"] = o.AllowComments
	}
	if !IsNil(o.AllowDuet) {
		toSerialize["allow_duet"] = o.AllowDuet
	}
	if !IsNil(o.AllowStitch) {
		toSerialize["allow_stitch"] = o.AllowStitch
	}
	if !IsNil(o.ContentDisclosure) {
		toSerialize["content_disclosure"] = o.ContentDisclosure
	}
	if !IsNil(o.BrandOrganicToggle) {
		toSerialize["brand_organic_toggle"] = o.BrandOrganicToggle
	}
	if !IsNil(o.BrandContentToggle) {
		toSerialize["brand_content_toggle"] = o.BrandContentToggle
	}
	return toSerialize, nil
}

type NullableVersionOptionsTiktok struct {
	value *VersionOptionsTiktok
	isSet bool
}

func (v NullableVersionOptionsTiktok) Get() *VersionOptionsTiktok {
	return v.value
}

func (v *NullableVersionOptionsTiktok) Set(val *VersionOptionsTiktok) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptionsTiktok) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptionsTiktok) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptionsTiktok(val *VersionOptionsTiktok) *NullableVersionOptionsTiktok {
	return &NullableVersionOptionsTiktok{value: val, isSet: true}
}

func (v NullableVersionOptionsTiktok) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptionsTiktok) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

