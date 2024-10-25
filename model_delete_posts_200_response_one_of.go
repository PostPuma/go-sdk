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

// checks if the DeletePosts200ResponseOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePosts200ResponseOneOf{}

// DeletePosts200ResponseOneOf struct for DeletePosts200ResponseOneOf
type DeletePosts200ResponseOneOf struct {
	Deleted *bool `json:"deleted,omitempty"`
}

// NewDeletePosts200ResponseOneOf instantiates a new DeletePosts200ResponseOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePosts200ResponseOneOf() *DeletePosts200ResponseOneOf {
	this := DeletePosts200ResponseOneOf{}
	return &this
}

// NewDeletePosts200ResponseOneOfWithDefaults instantiates a new DeletePosts200ResponseOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePosts200ResponseOneOfWithDefaults() *DeletePosts200ResponseOneOf {
	this := DeletePosts200ResponseOneOf{}
	return &this
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *DeletePosts200ResponseOneOf) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePosts200ResponseOneOf) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *DeletePosts200ResponseOneOf) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *DeletePosts200ResponseOneOf) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o DeletePosts200ResponseOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePosts200ResponseOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableDeletePosts200ResponseOneOf struct {
	value *DeletePosts200ResponseOneOf
	isSet bool
}

func (v NullableDeletePosts200ResponseOneOf) Get() *DeletePosts200ResponseOneOf {
	return v.value
}

func (v *NullableDeletePosts200ResponseOneOf) Set(val *DeletePosts200ResponseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePosts200ResponseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePosts200ResponseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePosts200ResponseOneOf(val *DeletePosts200ResponseOneOf) *NullableDeletePosts200ResponseOneOf {
	return &NullableDeletePosts200ResponseOneOf{value: val, isSet: true}
}

func (v NullableDeletePosts200ResponseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePosts200ResponseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


