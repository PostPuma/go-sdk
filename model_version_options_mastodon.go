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

// checks if the VersionOptionsMastodon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptionsMastodon{}

// VersionOptionsMastodon struct for VersionOptionsMastodon
type VersionOptionsMastodon struct {
	Sensitive *bool `json:"sensitive,omitempty"`
}

// NewVersionOptionsMastodon instantiates a new VersionOptionsMastodon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptionsMastodon() *VersionOptionsMastodon {
	this := VersionOptionsMastodon{}
	return &this
}

// NewVersionOptionsMastodonWithDefaults instantiates a new VersionOptionsMastodon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsMastodonWithDefaults() *VersionOptionsMastodon {
	this := VersionOptionsMastodon{}
	return &this
}

// GetSensitive returns the Sensitive field value if set, zero value otherwise.
func (o *VersionOptionsMastodon) GetSensitive() bool {
	if o == nil || IsNil(o.Sensitive) {
		var ret bool
		return ret
	}
	return *o.Sensitive
}

// GetSensitiveOk returns a tuple with the Sensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsMastodon) GetSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Sensitive) {
		return nil, false
	}
	return o.Sensitive, true
}

// HasSensitive returns a boolean if a field has been set.
func (o *VersionOptionsMastodon) HasSensitive() bool {
	if o != nil && !IsNil(o.Sensitive) {
		return true
	}

	return false
}

// SetSensitive gets a reference to the given bool and assigns it to the Sensitive field.
func (o *VersionOptionsMastodon) SetSensitive(v bool) {
	o.Sensitive = &v
}

func (o VersionOptionsMastodon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptionsMastodon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sensitive) {
		toSerialize["sensitive"] = o.Sensitive
	}
	return toSerialize, nil
}

type NullableVersionOptionsMastodon struct {
	value *VersionOptionsMastodon
	isSet bool
}

func (v NullableVersionOptionsMastodon) Get() *VersionOptionsMastodon {
	return v.value
}

func (v *NullableVersionOptionsMastodon) Set(val *VersionOptionsMastodon) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptionsMastodon) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptionsMastodon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptionsMastodon(val *VersionOptionsMastodon) *NullableVersionOptionsMastodon {
	return &NullableVersionOptionsMastodon{value: val, isSet: true}
}

func (v NullableVersionOptionsMastodon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptionsMastodon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

