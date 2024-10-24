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

// checks if the AccountData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountData{}

// AccountData struct for AccountData
type AccountData struct {
	UnionId *string `json:"union_id,omitempty"`
	IsPrivate *bool `json:"is_private,omitempty"`
	DuetDisabled *bool `json:"duet_disabled,omitempty"`
	PrivacyLevels []string `json:"privacy_levels,omitempty"`
	StitchDisabled *bool `json:"stitch_disabled,omitempty"`
	CommentDisabled *bool `json:"comment_disabled,omitempty"`
	MaxVideoPostDurationSec *int32 `json:"max_video_post_duration_sec,omitempty"`
}

// NewAccountData instantiates a new AccountData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountData() *AccountData {
	this := AccountData{}
	return &this
}

// NewAccountDataWithDefaults instantiates a new AccountData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDataWithDefaults() *AccountData {
	this := AccountData{}
	return &this
}

// GetUnionId returns the UnionId field value if set, zero value otherwise.
func (o *AccountData) GetUnionId() string {
	if o == nil || IsNil(o.UnionId) {
		var ret string
		return ret
	}
	return *o.UnionId
}

// GetUnionIdOk returns a tuple with the UnionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetUnionIdOk() (*string, bool) {
	if o == nil || IsNil(o.UnionId) {
		return nil, false
	}
	return o.UnionId, true
}

// HasUnionId returns a boolean if a field has been set.
func (o *AccountData) HasUnionId() bool {
	if o != nil && !IsNil(o.UnionId) {
		return true
	}

	return false
}

// SetUnionId gets a reference to the given string and assigns it to the UnionId field.
func (o *AccountData) SetUnionId(v string) {
	o.UnionId = &v
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise.
func (o *AccountData) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate) {
		var ret bool
		return ret
	}
	return *o.IsPrivate
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetIsPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPrivate) {
		return nil, false
	}
	return o.IsPrivate, true
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *AccountData) HasIsPrivate() bool {
	if o != nil && !IsNil(o.IsPrivate) {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given bool and assigns it to the IsPrivate field.
func (o *AccountData) SetIsPrivate(v bool) {
	o.IsPrivate = &v
}

// GetDuetDisabled returns the DuetDisabled field value if set, zero value otherwise.
func (o *AccountData) GetDuetDisabled() bool {
	if o == nil || IsNil(o.DuetDisabled) {
		var ret bool
		return ret
	}
	return *o.DuetDisabled
}

// GetDuetDisabledOk returns a tuple with the DuetDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetDuetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DuetDisabled) {
		return nil, false
	}
	return o.DuetDisabled, true
}

// HasDuetDisabled returns a boolean if a field has been set.
func (o *AccountData) HasDuetDisabled() bool {
	if o != nil && !IsNil(o.DuetDisabled) {
		return true
	}

	return false
}

// SetDuetDisabled gets a reference to the given bool and assigns it to the DuetDisabled field.
func (o *AccountData) SetDuetDisabled(v bool) {
	o.DuetDisabled = &v
}

// GetPrivacyLevels returns the PrivacyLevels field value if set, zero value otherwise.
func (o *AccountData) GetPrivacyLevels() []string {
	if o == nil || IsNil(o.PrivacyLevels) {
		var ret []string
		return ret
	}
	return o.PrivacyLevels
}

// GetPrivacyLevelsOk returns a tuple with the PrivacyLevels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetPrivacyLevelsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrivacyLevels) {
		return nil, false
	}
	return o.PrivacyLevels, true
}

// HasPrivacyLevels returns a boolean if a field has been set.
func (o *AccountData) HasPrivacyLevels() bool {
	if o != nil && !IsNil(o.PrivacyLevels) {
		return true
	}

	return false
}

// SetPrivacyLevels gets a reference to the given []string and assigns it to the PrivacyLevels field.
func (o *AccountData) SetPrivacyLevels(v []string) {
	o.PrivacyLevels = v
}

// GetStitchDisabled returns the StitchDisabled field value if set, zero value otherwise.
func (o *AccountData) GetStitchDisabled() bool {
	if o == nil || IsNil(o.StitchDisabled) {
		var ret bool
		return ret
	}
	return *o.StitchDisabled
}

// GetStitchDisabledOk returns a tuple with the StitchDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetStitchDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.StitchDisabled) {
		return nil, false
	}
	return o.StitchDisabled, true
}

// HasStitchDisabled returns a boolean if a field has been set.
func (o *AccountData) HasStitchDisabled() bool {
	if o != nil && !IsNil(o.StitchDisabled) {
		return true
	}

	return false
}

// SetStitchDisabled gets a reference to the given bool and assigns it to the StitchDisabled field.
func (o *AccountData) SetStitchDisabled(v bool) {
	o.StitchDisabled = &v
}

// GetCommentDisabled returns the CommentDisabled field value if set, zero value otherwise.
func (o *AccountData) GetCommentDisabled() bool {
	if o == nil || IsNil(o.CommentDisabled) {
		var ret bool
		return ret
	}
	return *o.CommentDisabled
}

// GetCommentDisabledOk returns a tuple with the CommentDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetCommentDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CommentDisabled) {
		return nil, false
	}
	return o.CommentDisabled, true
}

// HasCommentDisabled returns a boolean if a field has been set.
func (o *AccountData) HasCommentDisabled() bool {
	if o != nil && !IsNil(o.CommentDisabled) {
		return true
	}

	return false
}

// SetCommentDisabled gets a reference to the given bool and assigns it to the CommentDisabled field.
func (o *AccountData) SetCommentDisabled(v bool) {
	o.CommentDisabled = &v
}

// GetMaxVideoPostDurationSec returns the MaxVideoPostDurationSec field value if set, zero value otherwise.
func (o *AccountData) GetMaxVideoPostDurationSec() int32 {
	if o == nil || IsNil(o.MaxVideoPostDurationSec) {
		var ret int32
		return ret
	}
	return *o.MaxVideoPostDurationSec
}

// GetMaxVideoPostDurationSecOk returns a tuple with the MaxVideoPostDurationSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountData) GetMaxVideoPostDurationSecOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxVideoPostDurationSec) {
		return nil, false
	}
	return o.MaxVideoPostDurationSec, true
}

// HasMaxVideoPostDurationSec returns a boolean if a field has been set.
func (o *AccountData) HasMaxVideoPostDurationSec() bool {
	if o != nil && !IsNil(o.MaxVideoPostDurationSec) {
		return true
	}

	return false
}

// SetMaxVideoPostDurationSec gets a reference to the given int32 and assigns it to the MaxVideoPostDurationSec field.
func (o *AccountData) SetMaxVideoPostDurationSec(v int32) {
	o.MaxVideoPostDurationSec = &v
}

func (o AccountData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnionId) {
		toSerialize["union_id"] = o.UnionId
	}
	if !IsNil(o.IsPrivate) {
		toSerialize["is_private"] = o.IsPrivate
	}
	if !IsNil(o.DuetDisabled) {
		toSerialize["duet_disabled"] = o.DuetDisabled
	}
	if !IsNil(o.PrivacyLevels) {
		toSerialize["privacy_levels"] = o.PrivacyLevels
	}
	if !IsNil(o.StitchDisabled) {
		toSerialize["stitch_disabled"] = o.StitchDisabled
	}
	if !IsNil(o.CommentDisabled) {
		toSerialize["comment_disabled"] = o.CommentDisabled
	}
	if !IsNil(o.MaxVideoPostDurationSec) {
		toSerialize["max_video_post_duration_sec"] = o.MaxVideoPostDurationSec
	}
	return toSerialize, nil
}

type NullableAccountData struct {
	value *AccountData
	isSet bool
}

func (v NullableAccountData) Get() *AccountData {
	return v.value
}

func (v *NullableAccountData) Set(val *AccountData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountData(val *AccountData) *NullableAccountData {
	return &NullableAccountData{value: val, isSet: true}
}

func (v NullableAccountData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


