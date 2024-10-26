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

// checks if the VersionOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionOptions{}

// VersionOptions struct for VersionOptions
type VersionOptions struct {
	Tiktok *VersionOptionsTiktok `json:"tiktok,omitempty"`
	Youtube *VersionOptionsYoutube `json:"youtube,omitempty"`
	Linkedin *VersionOptionsLinkedin `json:"linkedin,omitempty"`
	Mastodon *VersionOptionsMastodon `json:"mastodon,omitempty"`
	Instagram *VersionOptionsInstagram `json:"instagram,omitempty"`
	Pinterest *VersionOptionsPinterest `json:"pinterest,omitempty"`
	FacebookPage *VersionOptionsInstagram `json:"facebook_page,omitempty"`
}

// NewVersionOptions instantiates a new VersionOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionOptions() *VersionOptions {
	this := VersionOptions{}
	return &this
}

// NewVersionOptionsWithDefaults instantiates a new VersionOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionOptionsWithDefaults() *VersionOptions {
	this := VersionOptions{}
	return &this
}

// GetTiktok returns the Tiktok field value if set, zero value otherwise.
func (o *VersionOptions) GetTiktok() VersionOptionsTiktok {
	if o == nil || IsNil(o.Tiktok) {
		var ret VersionOptionsTiktok
		return ret
	}
	return *o.Tiktok
}

// GetTiktokOk returns a tuple with the Tiktok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetTiktokOk() (*VersionOptionsTiktok, bool) {
	if o == nil || IsNil(o.Tiktok) {
		return nil, false
	}
	return o.Tiktok, true
}

// HasTiktok returns a boolean if a field has been set.
func (o *VersionOptions) HasTiktok() bool {
	if o != nil && !IsNil(o.Tiktok) {
		return true
	}

	return false
}

// SetTiktok gets a reference to the given VersionOptionsTiktok and assigns it to the Tiktok field.
func (o *VersionOptions) SetTiktok(v VersionOptionsTiktok) {
	o.Tiktok = &v
}

// GetYoutube returns the Youtube field value if set, zero value otherwise.
func (o *VersionOptions) GetYoutube() VersionOptionsYoutube {
	if o == nil || IsNil(o.Youtube) {
		var ret VersionOptionsYoutube
		return ret
	}
	return *o.Youtube
}

// GetYoutubeOk returns a tuple with the Youtube field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetYoutubeOk() (*VersionOptionsYoutube, bool) {
	if o == nil || IsNil(o.Youtube) {
		return nil, false
	}
	return o.Youtube, true
}

// HasYoutube returns a boolean if a field has been set.
func (o *VersionOptions) HasYoutube() bool {
	if o != nil && !IsNil(o.Youtube) {
		return true
	}

	return false
}

// SetYoutube gets a reference to the given VersionOptionsYoutube and assigns it to the Youtube field.
func (o *VersionOptions) SetYoutube(v VersionOptionsYoutube) {
	o.Youtube = &v
}

// GetLinkedin returns the Linkedin field value if set, zero value otherwise.
func (o *VersionOptions) GetLinkedin() VersionOptionsLinkedin {
	if o == nil || IsNil(o.Linkedin) {
		var ret VersionOptionsLinkedin
		return ret
	}
	return *o.Linkedin
}

// GetLinkedinOk returns a tuple with the Linkedin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetLinkedinOk() (*VersionOptionsLinkedin, bool) {
	if o == nil || IsNil(o.Linkedin) {
		return nil, false
	}
	return o.Linkedin, true
}

// HasLinkedin returns a boolean if a field has been set.
func (o *VersionOptions) HasLinkedin() bool {
	if o != nil && !IsNil(o.Linkedin) {
		return true
	}

	return false
}

// SetLinkedin gets a reference to the given VersionOptionsLinkedin and assigns it to the Linkedin field.
func (o *VersionOptions) SetLinkedin(v VersionOptionsLinkedin) {
	o.Linkedin = &v
}

// GetMastodon returns the Mastodon field value if set, zero value otherwise.
func (o *VersionOptions) GetMastodon() VersionOptionsMastodon {
	if o == nil || IsNil(o.Mastodon) {
		var ret VersionOptionsMastodon
		return ret
	}
	return *o.Mastodon
}

// GetMastodonOk returns a tuple with the Mastodon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetMastodonOk() (*VersionOptionsMastodon, bool) {
	if o == nil || IsNil(o.Mastodon) {
		return nil, false
	}
	return o.Mastodon, true
}

// HasMastodon returns a boolean if a field has been set.
func (o *VersionOptions) HasMastodon() bool {
	if o != nil && !IsNil(o.Mastodon) {
		return true
	}

	return false
}

// SetMastodon gets a reference to the given VersionOptionsMastodon and assigns it to the Mastodon field.
func (o *VersionOptions) SetMastodon(v VersionOptionsMastodon) {
	o.Mastodon = &v
}

// GetInstagram returns the Instagram field value if set, zero value otherwise.
func (o *VersionOptions) GetInstagram() VersionOptionsInstagram {
	if o == nil || IsNil(o.Instagram) {
		var ret VersionOptionsInstagram
		return ret
	}
	return *o.Instagram
}

// GetInstagramOk returns a tuple with the Instagram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetInstagramOk() (*VersionOptionsInstagram, bool) {
	if o == nil || IsNil(o.Instagram) {
		return nil, false
	}
	return o.Instagram, true
}

// HasInstagram returns a boolean if a field has been set.
func (o *VersionOptions) HasInstagram() bool {
	if o != nil && !IsNil(o.Instagram) {
		return true
	}

	return false
}

// SetInstagram gets a reference to the given VersionOptionsInstagram and assigns it to the Instagram field.
func (o *VersionOptions) SetInstagram(v VersionOptionsInstagram) {
	o.Instagram = &v
}

// GetPinterest returns the Pinterest field value if set, zero value otherwise.
func (o *VersionOptions) GetPinterest() VersionOptionsPinterest {
	if o == nil || IsNil(o.Pinterest) {
		var ret VersionOptionsPinterest
		return ret
	}
	return *o.Pinterest
}

// GetPinterestOk returns a tuple with the Pinterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetPinterestOk() (*VersionOptionsPinterest, bool) {
	if o == nil || IsNil(o.Pinterest) {
		return nil, false
	}
	return o.Pinterest, true
}

// HasPinterest returns a boolean if a field has been set.
func (o *VersionOptions) HasPinterest() bool {
	if o != nil && !IsNil(o.Pinterest) {
		return true
	}

	return false
}

// SetPinterest gets a reference to the given VersionOptionsPinterest and assigns it to the Pinterest field.
func (o *VersionOptions) SetPinterest(v VersionOptionsPinterest) {
	o.Pinterest = &v
}

// GetFacebookPage returns the FacebookPage field value if set, zero value otherwise.
func (o *VersionOptions) GetFacebookPage() VersionOptionsInstagram {
	if o == nil || IsNil(o.FacebookPage) {
		var ret VersionOptionsInstagram
		return ret
	}
	return *o.FacebookPage
}

// GetFacebookPageOk returns a tuple with the FacebookPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionOptions) GetFacebookPageOk() (*VersionOptionsInstagram, bool) {
	if o == nil || IsNil(o.FacebookPage) {
		return nil, false
	}
	return o.FacebookPage, true
}

// HasFacebookPage returns a boolean if a field has been set.
func (o *VersionOptions) HasFacebookPage() bool {
	if o != nil && !IsNil(o.FacebookPage) {
		return true
	}

	return false
}

// SetFacebookPage gets a reference to the given VersionOptionsInstagram and assigns it to the FacebookPage field.
func (o *VersionOptions) SetFacebookPage(v VersionOptionsInstagram) {
	o.FacebookPage = &v
}

func (o VersionOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tiktok) {
		toSerialize["tiktok"] = o.Tiktok
	}
	if !IsNil(o.Youtube) {
		toSerialize["youtube"] = o.Youtube
	}
	if !IsNil(o.Linkedin) {
		toSerialize["linkedin"] = o.Linkedin
	}
	if !IsNil(o.Mastodon) {
		toSerialize["mastodon"] = o.Mastodon
	}
	if !IsNil(o.Instagram) {
		toSerialize["instagram"] = o.Instagram
	}
	if !IsNil(o.Pinterest) {
		toSerialize["pinterest"] = o.Pinterest
	}
	if !IsNil(o.FacebookPage) {
		toSerialize["facebook_page"] = o.FacebookPage
	}
	return toSerialize, nil
}

type NullableVersionOptions struct {
	value *VersionOptions
	isSet bool
}

func (v NullableVersionOptions) Get() *VersionOptions {
	return v.value
}

func (v *NullableVersionOptions) Set(val *VersionOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionOptions(val *VersionOptions) *NullableVersionOptions {
	return &NullableVersionOptions{value: val, isSet: true}
}

func (v NullableVersionOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


