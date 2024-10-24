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

// checks if the CreateTagRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTagRequest{}

// CreateTagRequest struct for CreateTagRequest
type CreateTagRequest struct {
	Name string `json:"name"`
	HexColor string `json:"hex_color"`
}

type _CreateTagRequest CreateTagRequest

// NewCreateTagRequest instantiates a new CreateTagRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTagRequest(name string, hexColor string) *CreateTagRequest {
	this := CreateTagRequest{}
	this.Name = name
	this.HexColor = hexColor
	return &this
}

// NewCreateTagRequestWithDefaults instantiates a new CreateTagRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTagRequestWithDefaults() *CreateTagRequest {
	this := CreateTagRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateTagRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateTagRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateTagRequest) SetName(v string) {
	o.Name = v
}

// GetHexColor returns the HexColor field value
func (o *CreateTagRequest) GetHexColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HexColor
}

// GetHexColorOk returns a tuple with the HexColor field value
// and a boolean to check if the value has been set.
func (o *CreateTagRequest) GetHexColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HexColor, true
}

// SetHexColor sets field value
func (o *CreateTagRequest) SetHexColor(v string) {
	o.HexColor = v
}

func (o CreateTagRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTagRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["hex_color"] = o.HexColor
	return toSerialize, nil
}

func (o *CreateTagRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"hex_color",
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

	varCreateTagRequest := _CreateTagRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateTagRequest)

	if err != nil {
		return err
	}

	*o = CreateTagRequest(varCreateTagRequest)

	return err
}

type NullableCreateTagRequest struct {
	value *CreateTagRequest
	isSet bool
}

func (v NullableCreateTagRequest) Get() *CreateTagRequest {
	return v.value
}

func (v *NullableCreateTagRequest) Set(val *CreateTagRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTagRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTagRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTagRequest(val *CreateTagRequest) *NullableCreateTagRequest {
	return &NullableCreateTagRequest{value: val, isSet: true}
}

func (v NullableCreateTagRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTagRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


