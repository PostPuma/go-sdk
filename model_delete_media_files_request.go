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
	"bytes"
	"fmt"
)

// checks if the DeleteMediaFilesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMediaFilesRequest{}

// DeleteMediaFilesRequest struct for DeleteMediaFilesRequest
type DeleteMediaFilesRequest struct {
	// File IDs
	Items []int32 `json:"items"`
}

type _DeleteMediaFilesRequest DeleteMediaFilesRequest

// NewDeleteMediaFilesRequest instantiates a new DeleteMediaFilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMediaFilesRequest(items []int32) *DeleteMediaFilesRequest {
	this := DeleteMediaFilesRequest{}
	this.Items = items
	return &this
}

// NewDeleteMediaFilesRequestWithDefaults instantiates a new DeleteMediaFilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMediaFilesRequestWithDefaults() *DeleteMediaFilesRequest {
	this := DeleteMediaFilesRequest{}
	return &this
}

// GetItems returns the Items field value
func (o *DeleteMediaFilesRequest) GetItems() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DeleteMediaFilesRequest) GetItemsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *DeleteMediaFilesRequest) SetItems(v []int32) {
	o.Items = v
}

func (o DeleteMediaFilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMediaFilesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *DeleteMediaFilesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeleteMediaFilesRequest := _DeleteMediaFilesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteMediaFilesRequest)

	if err != nil {
		return err
	}

	*o = DeleteMediaFilesRequest(varDeleteMediaFilesRequest)

	return err
}

type NullableDeleteMediaFilesRequest struct {
	value *DeleteMediaFilesRequest
	isSet bool
}

func (v NullableDeleteMediaFilesRequest) Get() *DeleteMediaFilesRequest {
	return v.value
}

func (v *NullableDeleteMediaFilesRequest) Set(val *DeleteMediaFilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMediaFilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMediaFilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMediaFilesRequest(val *DeleteMediaFilesRequest) *NullableDeleteMediaFilesRequest {
	return &NullableDeleteMediaFilesRequest{value: val, isSet: true}
}

func (v NullableDeleteMediaFilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMediaFilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


