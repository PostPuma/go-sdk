# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]Account**](Account.md) |  | [optional] 
**Versions** | Pointer to [**[]Version**](Version.md) |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**User** | Pointer to [**PostUser**](PostUser.md) |  | [optional] 
**ScheduledAt** | Pointer to **string** |  | [optional] 
**PublishedAt** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Trashed** | Pointer to **bool** |  | [optional] 

## Methods

### NewPost

`func NewPost() *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Post) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Post) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Post) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Post) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Post) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Post) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Post) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Post) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *Post) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Post) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Post) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Post) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAccounts

`func (o *Post) GetAccounts() []Account`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Post) GetAccountsOk() (*[]Account, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Post) SetAccounts(v []Account)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Post) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetVersions

`func (o *Post) GetVersions() []Version`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Post) GetVersionsOk() (*[]Version, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Post) SetVersions(v []Version)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Post) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetTags

`func (o *Post) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Post) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Post) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Post) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUser

`func (o *Post) GetUser() PostUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Post) GetUserOk() (*PostUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Post) SetUser(v PostUser)`

SetUser sets User field to given value.

### HasUser

`func (o *Post) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetScheduledAt

`func (o *Post) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Post) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Post) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *Post) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Post) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Post) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Post) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Post) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Post) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Post) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Post) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Post) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetTrashed

`func (o *Post) GetTrashed() bool`

GetTrashed returns the Trashed field if non-nil, zero value otherwise.

### GetTrashedOk

`func (o *Post) GetTrashedOk() (*bool, bool)`

GetTrashedOk returns a tuple with the Trashed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrashed

`func (o *Post) SetTrashed(v bool)`

SetTrashed sets Trashed field to given value.

### HasTrashed

`func (o *Post) HasTrashed() bool`

HasTrashed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


