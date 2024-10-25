# DeletePostsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posts** | **[]string** | Post UUIDs | 
**Trash** | Pointer to **bool** |  | [optional] 

## Methods

### NewDeletePostsRequest

`func NewDeletePostsRequest(posts []string, ) *DeletePostsRequest`

NewDeletePostsRequest instantiates a new DeletePostsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePostsRequestWithDefaults

`func NewDeletePostsRequestWithDefaults() *DeletePostsRequest`

NewDeletePostsRequestWithDefaults instantiates a new DeletePostsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosts

`func (o *DeletePostsRequest) GetPosts() []string`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *DeletePostsRequest) GetPostsOk() (*[]string, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *DeletePostsRequest) SetPosts(v []string)`

SetPosts sets Posts field to given value.


### GetTrash

`func (o *DeletePostsRequest) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *DeletePostsRequest) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *DeletePostsRequest) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *DeletePostsRequest) HasTrash() bool`

HasTrash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


