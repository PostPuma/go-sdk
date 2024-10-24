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

// checks if the VersionOptionsPinterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptionsPinterest{}

// VersionOptionsPinterest struct for VersionOptionsPinterest
type VersionOptionsPinterest struct {
	Title *string `json:"title,omitempty"`
	Link *string `json:"link,omitempty"`
	Boards *VersionOptionsPinterestBoards `json:"boards,omitempty"`
}

// NewVersionOptionsPinterest instantiates a new VersionOptionsPinterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptionsPinterest() *VersionOptionsPinterest {
	this := VersionOptionsPinterest{}
	return &this
}

// NewVersionOptionsPinterestWithDefaults instantiates a new VersionOptionsPinterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsPinterestWithDefaults() *VersionOptionsPinterest {
	this := VersionOptionsPinterest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *VersionOptionsPinterest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsPinterest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *VersionOptionsPinterest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *VersionOptionsPinterest) SetTitle(v string) {
	o.Title = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *VersionOptionsPinterest) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsPinterest) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *VersionOptionsPinterest) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *VersionOptionsPinterest) SetLink(v string) {
	o.Link = &v
}

// GetBoards returns the Boards field value if set, zero value otherwise.
func (o *VersionOptionsPinterest) GetBoards() VersionOptionsPinterestBoards {
	if o == nil || IsNil(o.Boards) {
		var ret VersionOptionsPinterestBoards
		return ret
	}
	return *o.Boards
}

// GetBoardsOk returns a tuple with the Boards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptionsPinterest) GetBoardsOk() (*VersionOptionsPinterestBoards, bool) {
	if o == nil || IsNil(o.Boards) {
		return nil, false
	}
	return o.Boards, true
}

// HasBoards returns a boolean if a field has been set.
func (o *VersionOptionsPinterest) HasBoards() bool {
	if o != nil && !IsNil(o.Boards) {
		return true
	}

	return false
}

// SetBoards gets a reference to the given VersionOptionsPinterestBoards and assigns it to the Boards field.
func (o *VersionOptionsPinterest) SetBoards(v VersionOptionsPinterestBoards) {
	o.Boards = &v
}

func (o VersionOptionsPinterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptionsPinterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.Boards) {
		toSerialize["boards"] = o.Boards
	}
	return toSerialize, nil
}

type NullableVersionOptionsPinterest struct {
	value *VersionOptionsPinterest
	isSet bool
}

func (v NullableVersionOptionsPinterest) Get() *VersionOptionsPinterest {
	return v.value
}

func (v *NullableVersionOptionsPinterest) Set(val *VersionOptionsPinterest) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptionsPinterest) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptionsPinterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptionsPinterest(val *VersionOptionsPinterest) *NullableVersionOptionsPinterest {
	return &NullableVersionOptionsPinterest{value: val, isSet: true}
}

func (v NullableVersionOptionsPinterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptionsPinterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

