# CreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **bool** |  | [optional] 
**ScheduleNow** | Pointer to **bool** |  | [optional] 
**Queue** | Pointer to **bool** |  | [optional] 
**Accounts** | Pointer to **[]int32** |  | [optional] 
**Tags** | Pointer to **[]int32** |  | [optional] 
**Versions** | Pointer to [**[]Version**](Version.md) |  | [optional] 

## Methods

### NewCreatePostRequest

`func NewCreatePostRequest() *CreatePostRequest`

NewCreatePostRequest instantiates a new CreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePostRequestWithDefaults

`func NewCreatePostRequestWithDefaults() *CreatePostRequest`

NewCreatePostRequestWithDefaults instantiates a new CreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *CreatePostRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreatePostRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreatePostRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CreatePostRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTime

`func (o *CreatePostRequest) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreatePostRequest) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreatePostRequest) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *CreatePostRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTimezone

`func (o *CreatePostRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreatePostRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreatePostRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreatePostRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSchedule

`func (o *CreatePostRequest) GetSchedule() bool`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreatePostRequest) GetScheduleOk() (*bool, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreatePostRequest) SetSchedule(v bool)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreatePostRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScheduleNow

`func (o *CreatePostRequest) GetScheduleNow() bool`

GetScheduleNow returns the ScheduleNow field if non-nil, zero value otherwise.

### GetScheduleNowOk

`func (o *CreatePostRequest) GetScheduleNowOk() (*bool, bool)`

GetScheduleNowOk returns a tuple with the ScheduleNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleNow

`func (o *CreatePostRequest) SetScheduleNow(v bool)`

SetScheduleNow sets ScheduleNow field to given value.

### HasScheduleNow

`func (o *CreatePostRequest) HasScheduleNow() bool`

HasScheduleNow returns a boolean if a field has been set.

### GetQueue

`func (o *CreatePostRequest) GetQueue() bool`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CreatePostRequest) GetQueueOk() (*bool, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CreatePostRequest) SetQueue(v bool)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CreatePostRequest) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetAccounts

`func (o *CreatePostRequest) GetAccounts() []int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CreatePostRequest) GetAccountsOk() (*[]int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CreatePostRequest) SetAccounts(v []int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CreatePostRequest) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetTags

`func (o *CreatePostRequest) GetTags() []int32`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePostRequest) GetTagsOk() (*[]int32, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePostRequest) SetTags(v []int32)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersions

`func (o *CreatePostRequest) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *CreatePostRequest) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *CreatePostRequest) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *CreatePostRequest) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


