# UpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Versions** | Pointer to [**[]Version**](Version.md) |  | [optional] 

## Methods

### NewUpdatePostRequest

`func NewUpdatePostRequest() *UpdatePostRequest`

NewUpdatePostRequest instantiates a new UpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostRequestWithDefaults

`func NewUpdatePostRequestWithDefaults() *UpdatePostRequest`

NewUpdatePostRequestWithDefaults instantiates a new UpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *UpdatePostRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UpdatePostRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UpdatePostRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *UpdatePostRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTime

`func (o *UpdatePostRequest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *UpdatePostRequest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *UpdatePostRequest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *UpdatePostRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdatePostRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdatePostRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdatePostRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdatePostRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdatePostRequest) GetAccounts() []int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdatePostRequest) GetAccountsOk() (*[]int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdatePostRequest) SetAccounts(v []int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdatePostRequest) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePostRequest) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePostRequest) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePostRequest) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *UpdatePostRequest) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *UpdatePostRequest) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *UpdatePostRequest) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *UpdatePostRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


