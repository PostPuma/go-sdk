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

// checks if the UpdatePostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePostRequest{}

// UpdatePostRequest struct for UpdatePostRequest
type UpdatePostRequest struct {
	Date *string `json:"date,omitempty"`
	Time *string `json:"time,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Accounts []int32 `json:"accounts,omitempty"`
	Tags []int32 `json:"tags,omitempty"`
	Versions []Version `json:"versions,omitempty"`
}

// NewUpdatePostRequest instantiates a new UpdatePostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePostRequest() *UpdatePostRequest {
	this := UpdatePostRequest{}
	return &this
}

// NewUpdatePostRequestWithDefaults instantiates a new UpdatePostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePostRequestWithDefaults() *UpdatePostRequest {
	this := UpdatePostRequest{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *UpdatePostRequest) SetDate(v string) {
	o.Date = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *UpdatePostRequest) SetTime(v string) {
	o.Time = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *UpdatePostRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetAccounts() []int32 {
	if o == nil || IsNil(o.Accounts) {
		var ret []int32
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetAccountsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []int32 and assigns it to the Accounts field.
func (o *UpdatePostRequest) SetAccounts(v []int32) {
	o.Accounts = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetTags() []int32 {
	if o == nil || IsNil(o.Tags) {
		var ret []int32
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetTagsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []int32 and assigns it to the Tags field.
func (o *UpdatePostRequest) SetTags(v []int32) {
	o.Tags = v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *UpdatePostRequest) GetVersions() []Version {
	if o == nil || IsNil(o.Versions) {
		var ret []Version
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePostRequest) GetVersionsOk() ([]Version, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *UpdatePostRequest) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []Version and assigns it to the Versions field.
func (o *UpdatePostRequest) SetVersions(v []Version) {
	o.Versions = v
}

func (o UpdatePostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	return toSerialize, nil
}

type NullableUpdatePostRequest struct {
	value *UpdatePostRequest
	isSet bool
}

func (v NullableUpdatePostRequest) Get() *UpdatePostRequest {
	return v.value
}

func (v *NullableUpdatePostRequest) Set(val *UpdatePostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePostRequest(val *UpdatePostRequest) *NullableUpdatePostRequest {
	return &NullableUpdatePostRequest{value: val, isSet: true}
}

func (v NullableUpdatePostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


