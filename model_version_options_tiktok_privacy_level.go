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

// checks if the VersionOptionsTiktokPrivacyLevel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptionsTiktokPrivacyLevel{}

// VersionOptionsTiktokPrivacyLevel struct for VersionOptionsTiktokPrivacyLevel
type VersionOptionsTiktokPrivacyLevel struct {
	Account0 *bool `json:"account-0,omitempty"`
}

// NewVersionOptionsTiktokPrivacyLevel instantiates a new VersionOptionsTiktokPrivacyLevel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptionsTiktokPrivacyLevel() *VersionOptionsTiktokPrivacyLevel {
	this := VersionOptionsTiktokPrivacyLevel{}
	return &this
}

// NewVersionOptionsTiktokPrivacyLevelWithDefaults instantiates a new VersionOptionsTiktokPrivacyLevel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsTiktokPrivacyLevelWithDefaults() *VersionOptionsTiktokPrivacyLevel {
	this := VersionOptionsTiktokPrivacyLevel{}
	return &this
}

// GetAccount0 returns the Account0 field value if set, zero value otherwise.
func (o *VersionOptionsTiktokPrivacyLevel) GetAccount0() bool {
	if o == nil || IsNil(o.Account0) {
		var ret bool
		return ret
	}
	return *o.Account0
}

// GetAccount0Ok returns a tuple with the Account0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsTiktokPrivacyLevel) GetAccount0Ok() (*bool, bool) {
	if o == nil || IsNil(o.Account0) {
		return nil, false
	}
	return o.Account0, true
}

// HasAccount0 returns a boolean if a field has been set.
func (o *VersionOptionsTiktokPrivacyLevel) HasAccount0() bool {
	if o != nil && !IsNil(o.Account0) {
		return true
	}

	return false
}

// SetAccount0 gets a reference to the given bool and assigns it to the Account0 field.
func (o *VersionOptionsTiktokPrivacyLevel) SetAccount0(v bool) {
	o.Account0 = &v
}

func (o VersionOptionsTiktokPrivacyLevel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptionsTiktokPrivacyLevel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account0) {
		toSerialize["account-0"] = o.Account0
	}
	return toSerialize, nil
}

type NullableVersionOptionsTiktokPrivacyLevel struct {
	value *VersionOptionsTiktokPrivacyLevel
	isSet bool
}

func (v NullableVersionOptionsTiktokPrivacyLevel) Get() *VersionOptionsTiktokPrivacyLevel {
	return v.value
}

func (v *NullableVersionOptionsTiktokPrivacyLevel) Set(val *VersionOptionsTiktokPrivacyLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptionsTiktokPrivacyLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptionsTiktokPrivacyLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptionsTiktokPrivacyLevel(val *VersionOptionsTiktokPrivacyLevel) *NullableVersionOptionsTiktokPrivacyLevel {
	return &NullableVersionOptionsTiktokPrivacyLevel{value: val, isSet: true}
}

func (v NullableVersionOptionsTiktokPrivacyLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptionsTiktokPrivacyLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


