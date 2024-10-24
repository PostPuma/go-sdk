# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int32** |  | 
**IsOriginal** | **bool** |  | 
**Content** | [**VersionContent**](VersionContent.md) |  | 
**Options** | Pointer to [**VersionOptions**](VersionOptions.md) |  | [optional] 

## Methods

### NewVersion

`func NewVersion(accountId int32, isOriginal bool, content VersionContent, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Version) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Version) GetAccountIdOk() (*int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Version) SetAccountId(v int32)`

SetAccountId sets AccountId field to given value.


### GetIsOriginal

`func (o *Version) GetIsOriginal() bool`

GetIsOriginal returns the IsOriginal field if non-nil, zero value otherwise.

### GetIsOriginalOk

`func (o *Version) GetIsOriginalOk() (*bool, bool)`

GetIsOriginalOk returns a tuple with the IsOriginal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOriginal

`func (o *Version) SetIsOriginal(v bool)`

SetIsOriginal sets IsOriginal field to given value.


### GetContent

`func (o *Version) GetContent() VersionContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Version) GetContentOk() (*VersionContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Version) SetContent(v VersionContent)`

SetContent sets Content field to given value.


### GetOptions

`func (o *Version) GetOptions() VersionOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Version) GetOptionsOk() (*VersionOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Version) SetOptions(v VersionOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Version) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


