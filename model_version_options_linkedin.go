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

// checks if the VersionOptionsLinkedin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptionsLinkedin{}

// VersionOptionsLinkedin struct for VersionOptionsLinkedin
type VersionOptionsLinkedin struct {
	Visibility *string `json:"visibility,omitempty"`
}

// NewVersionOptionsLinkedin instantiates a new VersionOptionsLinkedin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptionsLinkedin() *VersionOptionsLinkedin {
	this := VersionOptionsLinkedin{}
	return &this
}

// NewVersionOptionsLinkedinWithDefaults instantiates a new VersionOptionsLinkedin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsLinkedinWithDefaults() *VersionOptionsLinkedin {
	this := VersionOptionsLinkedin{}
	return &this
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *VersionOptionsLinkedin) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsLinkedin) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *VersionOptionsLinkedin) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *VersionOptionsLinkedin) SetVisibility(v string) {
	o.Visibility = &v
}

func (o VersionOptionsLinkedin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptionsLinkedin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	return toSerialize, nil
}

type NullableVersionOptionsLinkedin struct {
	value *VersionOptionsLinkedin
	isSet bool
}

func (v NullableVersionOptionsLinkedin) Get() *VersionOptionsLinkedin {
	return v.value
}

func (v *NullableVersionOptionsLinkedin) Set(val *VersionOptionsLinkedin) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptionsLinkedin) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptionsLinkedin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptionsLinkedin(val *VersionOptionsLinkedin) *NullableVersionOptionsLinkedin {
	return &NullableVersionOptionsLinkedin{value: val, isSet: true}
}

func (v NullableVersionOptionsLinkedin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptionsLinkedin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


