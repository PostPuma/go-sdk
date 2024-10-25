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

// checks if the PostUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostUser{}

// PostUser struct for PostUser
type PostUser struct {
	Name *string `json:"name,omitempty"`
}

// NewPostUser instantiates a new PostUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostUser() *PostUser {
	this := PostUser{}
	return &this
}

// NewPostUserWithDefaults instantiates a new PostUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostUserWithDefaults() *PostUser {
	this := PostUser{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostUser) SetName(v string) {
	o.Name = &v
}

func (o PostUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePostUser struct {
	value *PostUser
	isSet bool
}

func (v NullablePostUser) Get() *PostUser {
	return v.value
}

func (v *NullablePostUser) Set(val *PostUser) {
	v.value = val
	v.isSet = true
}

func (v NullablePostUser) IsSet() bool {
	return v.isSet
}

func (v *NullablePostUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostUser(val *PostUser) *NullablePostUser {
	return &NullablePostUser{value: val, isSet: true}
}

func (v NullablePostUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


